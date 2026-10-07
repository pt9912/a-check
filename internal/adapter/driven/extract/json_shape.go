package extract

import (
	"fmt"
	"strings"

	"github.com/pt9912/a-check/internal/hexagon/core"
)

// This file is the `json` dialect of the shapes rule (AC-FA-RULE-012,
// ADR-0042), derived from RFC 8259: no comments, whitespace outside strings is
// dropped, a statement is one member of the root object. Where two sources
// would fall on one normal form, or the input is no JSON at all, the file is
// unsplittable (exit 2) — SPEC-EXTRACT-001, "Dialekt json".

type jsonTokKind int

const (
	jtStruct jsonTokKind = iota // one of { } [ ] : ,
	jtString                    // "…" verbatim
	jtAtom                      // a run of number/literal characters
)

type jsonTok struct {
	kind jsonTokKind
	text string
	line int
}

// jsonAtomChars are the characters of numbers and of true/false/null — the only
// characters allowed outside strings besides the structural ones and whitespace.
const jsonAtomChars = "0123456789+-.eEtrufalsn"

// normalizeJSON returns the members of the root object as statements and the
// line count (LF, CRLF or CR), or an error if the file cannot be split.
func normalizeJSON(src string) ([]core.Statement, int, error) {
	toks, err := jsonTokens(strings.TrimPrefix(src, "\uFEFF"))
	if err != nil {
		return nil, 0, err
	}
	stmts, err := jsonMembers(toks)
	if err != nil {
		return nil, 0, err
	}
	return stmts, countLines(src), nil
}

// jsonTokens lexes the source. Whitespace is dropped, but never between two
// tokens of which neither is structural — that would merge `nu ll` into `null`.
func jsonTokens(src string) ([]jsonTok, error) {
	l := &ktLexer{src: src, line: 1}
	var (
		toks       []jsonTok
		spaceSince bool
	)
	for l.i < len(src) {
		if isJSONSpace(src[l.i]) {
			if !l.lineEnd() {
				l.i++
			}
			spaceSince = true
			continue
		}
		tok, err := jsonNext(l)
		if err != nil {
			return nil, err
		}
		if spaceSince && len(toks) > 0 && toks[len(toks)-1].kind != jtStruct && tok.kind != jtStruct {
			return nil, fmt.Errorf("in Zeile %d: Leerraum zwischen zwei Werten ohne Trennzeichen", tok.line)
		}
		toks = append(toks, tok)
		spaceSince = false
	}
	return toks, nil
}

// isJSONSpace is RFC 8259 §2 whitespace: space, tab, LF, CR — a form feed is
// not, it is an illegal character outside a string.
func isJSONSpace(c byte) bool { return c == ' ' || c == '\t' || c == '\n' || c == '\r' }

// jsonNext reads one token at l.i: a structural character, a string or an atom.
func jsonNext(l *ktLexer) (jsonTok, error) {
	c, line := l.src[l.i], l.line
	switch {
	case strings.IndexByte("{}[]:,", c) >= 0:
		l.i++
		return jsonTok{kind: jtStruct, text: string(c), line: line}, nil
	case c == '"':
		start := l.i
		if err := jsonString(l); err != nil {
			return jsonTok{}, err
		}
		return jsonTok{kind: jtString, text: l.src[start:l.i], line: line}, nil
	case strings.IndexByte(jsonAtomChars, c) >= 0:
		start := l.i
		for l.i < len(l.src) && strings.IndexByte(jsonAtomChars, l.src[l.i]) >= 0 {
			l.i++
		}
		return jsonTok{kind: jtAtom, text: l.src[start:l.i], line: line}, nil
	default:
		return jsonTok{}, fmt.Errorf("in Zeile %d: Zeichen %q ist außerhalb einer Zeichenkette nicht zulässig (JSON kennt keine Kommentare)", line, c)
	}
}

// jsonString scans "…" from l.i: a backslash masks the next character, an
// unmasked control character (U+0000–U+001F, a line end too) is an error.
func jsonString(l *ktLexer) error {
	line := l.line
	for l.i++; l.i < len(l.src); l.i++ {
		switch c := l.src[l.i]; {
		case c == '\\':
			l.i++
		case c == '"':
			l.i++
			return nil
		case c < 0x20:
			return fmt.Errorf("in Zeile %d: unmaskiertes Steuerzeichen in einer Zeichenkette", line)
		}
	}
	return fmt.Errorf("in Zeile %d: Zeichenkette nicht geschlossen", line)
}

// jsonMembers splits the token stream into the members of the root object.
func jsonMembers(toks []jsonTok) ([]core.Statement, error) {
	if len(toks) == 0 || toks[0].text != "{" || toks[0].kind != jtStruct {
		return nil, fmt.Errorf("die Wurzel ist kein Objekt")
	}
	sp := &jsonSplit{stack: []string{"{"}}
	for _, t := range toks[1:] {
		if err := sp.take(t); err != nil {
			return nil, err
		}
	}
	if len(sp.stack) != 0 {
		return nil, fmt.Errorf("das Wurzel-Objekt ist nicht geschlossen")
	}
	return sp.out, nil
}

// jsonSplit collects the members of the root object while tracking brackets.
type jsonSplit struct {
	out   []core.Statement
	cur   []jsonTok
	stack []string
}

// take processes one token after the opening brace of the root: a `,` on depth
// 1 or the closing brace of the root ends a member; `{}` has none.
func (sp *jsonSplit) take(t jsonTok) error {
	if len(sp.stack) == 0 {
		return fmt.Errorf("in Zeile %d: Inhalt nach dem Ende des Wurzel-Objekts", t.line)
	}
	closed, err := jsonStructure(t, &sp.stack)
	if err != nil {
		return err
	}
	rootEnd := closed && len(sp.stack) == 0
	sep := t.kind == jtStruct && t.text == "," && len(sp.stack) == 1
	if !sep && !rootEnd {
		sp.cur = append(sp.cur, t)
		return nil
	}
	if rootEnd && len(sp.out) == 0 && len(sp.cur) == 0 {
		return nil // {} has no members
	}
	st, err := jsonMember(sp.cur, t.line)
	if err != nil {
		return err
	}
	sp.out = append(sp.out, st)
	sp.cur = nil
	return nil
}

// jsonStructure updates the bracket stack for a structural token and reports
// whether it closed a bracket; a mismatched or unmatched bracket is an error.
func jsonStructure(t jsonTok, stack *[]string) (bool, error) {
	if t.kind != jtStruct {
		return false, nil
	}
	switch t.text {
	case "{", "[":
		*stack = append(*stack, t.text)
	case "}", "]":
		want := map[string]string{"}": "{", "]": "["}[t.text]
		n := len(*stack)
		if n == 0 || (*stack)[n-1] != want {
			return false, fmt.Errorf("in Zeile %d: Klammer %q ohne passendes Gegenstück", t.line, t.text)
		}
		*stack = (*stack)[:n-1]
		return true, nil
	}
	return false, nil
}

// jsonMember checks one member — non-empty, starting with a string, with a `:`
// on its own level — and renders it without whitespace.
func jsonMember(m []jsonTok, line int) (core.Statement, error) {
	if len(m) == 0 {
		return core.Statement{}, fmt.Errorf("in Zeile %d: leeres Mitglied (doppeltes oder abschließendes Komma)", line)
	}
	if m[0].kind != jtString {
		return core.Statement{}, fmt.Errorf("in Zeile %d: ein Mitglied beginnt nicht mit einer Zeichenkette", m[0].line)
	}
	depth, colon := 0, false
	var b strings.Builder
	for _, t := range m {
		if t.kind == jtStruct {
			switch t.text {
			case "{", "[":
				depth++
			case "}", "]":
				depth--
			case ":":
				colon = colon || depth == 0
			}
		}
		b.WriteString(t.text)
	}
	if !colon {
		return core.Statement{}, fmt.Errorf("in Zeile %d: ein Mitglied trägt kein `:`", m[0].line)
	}
	return core.Statement{Text: b.String(), Line: m[0].line}, nil
}
