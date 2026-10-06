package extract

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// This file is the `kotlin` dialect of the shapes rule (AC-FA-RULE-012,
// ADR-0041): a lexer that knows strings completely — a bracket inside a string
// would otherwise shift the split — plus the normalization and the statement
// split of SPEC-EXTRACT-001. It is deliberately separate from stripComments:
// that one is a comment stripper that errs toward keeping text, this one must
// know where every literal ends, because a literal it ends too late or too early
// can swallow code as comment — the one failure direction that is not
// fail-safe (ADR-0041 §Konsequenzen).

type tokKind int

const (
	tkCode  tokKind = iota // one code rune
	tkLit                  // a whole string / char literal or backtick identifier, verbatim
	tkSpace                // whitespace and comments; nl marks a contained line end
)

type token struct {
	kind tokKind
	text string
	line int
	nl   bool
}

// ktLexer turns a Kotlin source into tokens. It never drops code: everything
// that is not whitespace or a comment ends up in a token.
type ktLexer struct {
	src  string
	i    int
	line int
	toks []token
}

// lexKotlin lexes src. A literal, comment or template still open at EOF is an
// error — the file cannot be split, and it never turns silently green.
func lexKotlin(src string) ([]token, int, error) {
	l := &ktLexer{src: strings.TrimPrefix(src, "\uFEFF"), line: 1}
	if strings.HasPrefix(l.src, "#!") {
		l.skipLineComment()
	}
	for l.i < len(l.src) {
		if err := l.step(true); err != nil {
			return nil, 0, err
		}
	}
	return l.toks, l.line, nil
}

// step consumes one element at l.i. With emit=false it only advances — the mode
// a template expression is scanned in, because its text is already part of the
// enclosing literal.
func (l *ktLexer) step(emit bool) error {
	c := l.src[l.i]
	switch {
	case isSpaceByte(c):
		l.space(emit, l.consumeSpace())
	case strings.HasPrefix(l.src[l.i:], "/*"):
		nl, err := l.skipBlockComment()
		if err != nil {
			return err
		}
		l.space(emit, nl)
	case strings.HasPrefix(l.src[l.i:], "//"):
		l.skipLineComment()
		l.space(emit, false)
	case c == '$' && l.dollarPrefixedString():
		return l.literal(emit, l.lexString)
	case c == '"':
		return l.literal(emit, l.lexString)
	case c == '\'':
		return l.literal(emit, l.lexChar)
	case c == '`':
		return l.literal(emit, l.lexBacktick)
	default:
		r, size := utf8.DecodeRuneInString(l.src[l.i:])
		if emit {
			l.toks = append(l.toks, token{kind: tkCode, text: string(r), line: l.line})
		}
		l.i += size
	}
	return nil
}

// literal runs a literal scanner and records its verbatim text as one token.
func (l *ktLexer) literal(emit bool, scan func() error) error {
	start, line := l.i, l.line
	if err := scan(); err != nil {
		return err
	}
	if emit {
		l.toks = append(l.toks, token{kind: tkLit, text: l.src[start:l.i], line: line})
	}
	return nil
}

// space records whitespace, merging with a preceding space token.
func (l *ktLexer) space(emit, nl bool) {
	if !emit {
		return
	}
	if n := len(l.toks); n > 0 && l.toks[n-1].kind == tkSpace {
		l.toks[n-1].nl = l.toks[n-1].nl || nl
		return
	}
	l.toks = append(l.toks, token{kind: tkSpace, line: l.line, nl: nl})
}

func isSpaceByte(c byte) bool { return c == ' ' || c == '\t' || c == '\f' || c == '\r' || c == '\n' }

// lineEnd consumes one line end at l.i (LF, CRLF or a lone CR) and reports
// whether there was one.
func (l *ktLexer) lineEnd() bool {
	switch {
	case strings.HasPrefix(l.src[l.i:], "\r\n"):
		l.i += 2
	case l.src[l.i] == '\r' || l.src[l.i] == '\n':
		l.i++
	default:
		return false
	}
	l.line++
	return true
}

func (l *ktLexer) consumeSpace() bool {
	nl := false
	for l.i < len(l.src) && isSpaceByte(l.src[l.i]) {
		if l.lineEnd() {
			nl = true
			continue
		}
		l.i++
	}
	return nl
}

func (l *ktLexer) skipLineComment() {
	for l.i < len(l.src) && l.src[l.i] != '\n' && l.src[l.i] != '\r' {
		l.i++
	}
}

// skipBlockComment consumes a NESTED block comment (every /* opens, every */
// closes a level).
func (l *ktLexer) skipBlockComment() (bool, error) {
	startLine, depth, nl := l.line, 0, false
	for l.i < len(l.src) {
		switch {
		case strings.HasPrefix(l.src[l.i:], "/*"):
			depth++
			l.i += 2
		case strings.HasPrefix(l.src[l.i:], "*/"):
			depth--
			l.i += 2
			if depth == 0 {
				return nl, nil
			}
		case l.lineEnd():
			nl = true
		default:
			l.i++
		}
	}
	return false, fmt.Errorf("in Zeile %d: Block-Kommentar nicht geschlossen", startLine)
}

// dollarPrefixedString reports whether the `$` run at l.i directly precedes a
// `"` — then it is a string's dollar prefix (SPEC-EXTRACT-001), not code.
func (l *ktLexer) dollarPrefixedString() bool {
	j := l.i
	for j < len(l.src) && l.src[j] == '$' {
		j++
	}
	return j < len(l.src) && l.src[j] == '"'
}

// lexString scans a (raw) string with optional dollar prefix of length n ≥ 1.
func (l *ktLexer) lexString() error {
	n := 0
	for l.src[l.i] == '$' {
		n++
		l.i++
	}
	if n == 0 {
		n = 1
	}
	startLine := l.line
	if strings.HasPrefix(l.src[l.i:], `"""`) {
		l.i += 3
		return l.stringBody(n, true, startLine)
	}
	l.i++
	return l.stringBody(n, false, startLine)
}

// stringBody scans up to the closing delimiter. A template starts only where a
// run of at least n unescaped `$` stands before `{`; the run is counted FORWARD
// while lexing, so an escaped `\$` is content and ends a run (SPEC-EXTRACT-001).
func (l *ktLexer) stringBody(n int, raw bool, startLine int) error {
	unclosed := fmt.Errorf("in Zeile %d: Zeichenkette nicht geschlossen", startLine)
	for l.i < len(l.src) {
		done, ok, err := l.stringStep(n, raw)
		switch {
		case err != nil:
			return err
		case !ok:
			return unclosed
		case done:
			return nil
		}
	}
	return unclosed
}

// stringStep consumes one element of a string body. done reports the closing
// delimiter; ok=false reports a line end or a trailing escape in a one-line
// string — the string is unclosed.
func (l *ktLexer) stringStep(n int, raw bool) (done, ok bool, err error) {
	c := l.src[l.i]
	switch {
	case c == '"' && raw:
		return l.rawClose(), true, nil
	case c == '"':
		l.i++
		return true, true, nil
	case c == '\\' && !raw:
		if l.i+1 >= len(l.src) || l.src[l.i+1] == '\n' || l.src[l.i+1] == '\r' {
			return false, false, nil
		}
		l.i += 2
	case c == '$':
		return false, true, l.dollarRun(n)
	case l.lineEnd():
		return false, raw, nil
	default:
		l.i++
	}
	return false, true, nil
}

// rawClose consumes the quote run at l.i. A run of at least three closes the raw
// string — surplus quotes before the last three belong to the content
// (SPEC-EXTRACT-001); a shorter run is content and the string goes on.
func (l *ktLexer) rawClose() bool {
	k := 0
	for l.i < len(l.src) && l.src[l.i] == '"' {
		k++
		l.i++
	}
	return k >= 3
}

// dollarRun consumes a run of `$`; if it is at least n long and followed by `{`,
// the last n open a template whose expression is scanned to its closing brace.
func (l *ktLexer) dollarRun(n int) error {
	run := 0
	for l.i < len(l.src) && l.src[l.i] == '$' {
		run++
		l.i++
	}
	if run >= n && l.i < len(l.src) && l.src[l.i] == '`' {
		// $`name` is a template too: a `"` inside the backtick identifier does
		// not end the string (Review slice-210 F-2, SPEC-EXTRACT-001)
		return l.lexBacktick()
	}
	if run < n || l.i >= len(l.src) || l.src[l.i] != '{' {
		return nil
	}
	return l.templateExpr()
}

// templateExpr scans a ${…} expression from its opening brace to the matching
// closing one; the expression may hold strings, comments and braces of its own.
func (l *ktLexer) templateExpr() error {
	startLine := l.line
	l.i++
	for depth := 1; ; {
		if l.i >= len(l.src) {
			return fmt.Errorf("in Zeile %d: Vorlage ${…} nicht geschlossen", startLine)
		}
		switch l.src[l.i] {
		case '{':
			depth++
			l.i++
		case '}':
			depth--
			l.i++
			if depth == 0 {
				return nil
			}
		default:
			if err := l.step(false); err != nil {
				return err
			}
		}
	}
}

// lexChar scans a char literal; a backslash escapes the next byte, so '{' and
// '\'' are literals, not code.
func (l *ktLexer) lexChar() error {
	startLine := l.line
	for l.i++; l.i < len(l.src); l.i++ {
		switch l.src[l.i] {
		case '\\':
			if l.i+1 >= len(l.src) || l.src[l.i+1] == '\n' || l.src[l.i+1] == '\r' {
				return fmt.Errorf("in Zeile %d: Zeichen-Literal nicht geschlossen", startLine)
			}
			l.i++
		case '\'':
			l.i++
			return nil
		case '\n', '\r':
			return fmt.Errorf("in Zeile %d: Zeichen-Literal nicht geschlossen", startLine)
		}
	}
	return fmt.Errorf("in Zeile %d: Zeichen-Literal nicht geschlossen", startLine)
}

// lexBacktick scans a backtick identifier as one token.
func (l *ktLexer) lexBacktick() error {
	startLine := l.line
	for l.i++; l.i < len(l.src); l.i++ {
		switch l.src[l.i] {
		case '`':
			l.i++
			return nil
		case '\n', '\r':
			return fmt.Errorf("in Zeile %d: Backtick-Bezeichner nicht geschlossen", startLine)
		}
	}
	return fmt.Errorf("in Zeile %d: Backtick-Bezeichner nicht geschlossen", startLine)
}

// charClass is the whitespace-folding class of SPEC-EXTRACT-001 step 2: word
// characters and operator characters keep ONE space between two of their own
// class; everything else (brackets, `, ; .`, literals) drops it.
type charClass int

const (
	clsNone charClass = iota
	clsWord
	clsOp
)

func classOf(r rune) charClass {
	switch {
	case r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r):
		return clsWord
	case strings.ContainsRune("()[]{},;.\"'`", r):
		return clsNone
	default:
		return clsOp
	}
}
