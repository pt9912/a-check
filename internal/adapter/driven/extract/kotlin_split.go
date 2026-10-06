package extract

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/pt9912/a-check/internal/hexagon/core"
)

// splitter normalizes the token stream and splits it into top-level statements
// (SPEC-EXTRACT-001 steps 2 and 3). The split is a CONTINUATION RULE, not a
// grammar: splitting too finely makes every piece need its own allow entry,
// splitting too coarsely makes the whole need one — both fail safe (ADR-0041
// point 4). Inside a block every boundary is written as one `;`, so `a()` and
// `b()` on two lines never fold into `a()b()`.
type splitter struct {
	toks     []token
	stmts    []core.Statement
	cur      strings.Builder
	curLine  int
	stack    []token // open brackets, with their line for the error message
	lastCls  charClass
	lastRune rune
	pending  bool // whitespace seen since the last emitted token
	pendNL   bool // ... and it contained a line end
}

// normalizeKotlin returns the normalized top-level statements of a Kotlin
// source and its line count (at least 1), or an error if it cannot be split.
func normalizeKotlin(src string) ([]core.Statement, int, error) {
	toks, _, err := lexKotlin(src)
	if err != nil {
		return nil, 0, err
	}
	s := &splitter{toks: toks}
	for k := range toks {
		if err := s.take(k); err != nil {
			return nil, 0, err
		}
	}
	if n := len(s.stack); n > 0 {
		open := s.stack[n-1]
		return nil, 0, fmt.Errorf("in Zeile %d: Klammer %q nicht geschlossen", open.line, open.text)
	}
	s.finish()
	return s.stmts, countLines(src), nil
}

// take processes token k.
func (s *splitter) take(k int) error {
	t := s.toks[k]
	if t.kind == tkSpace {
		s.pending = true
		s.pendNL = s.pendNL || t.nl
		return nil
	}
	if t.kind == tkCode && t.text == ";" && s.atBlockLevel() {
		s.boundary()
		s.pending, s.pendNL = false, false
		return nil
	}
	switch {
	case s.pendNL && s.lineEndSplits(k):
		s.boundary()
	case s.pending && s.needSpace(t):
		s.cur.WriteByte(' ')
	}
	s.pending, s.pendNL = false, false
	return s.emit(t)
}

// atBlockLevel reports whether a boundary may fall here: at top level or
// directly inside `{…}` — inside `(` and `[` a line end is whitespace.
func (s *splitter) atBlockLevel() bool {
	n := len(s.stack)
	return n == 0 || s.stack[n-1].text == "{"
}

// lineEndSplits applies the continuation rule to the line end before token k:
// no split inside ( or [, before a line starting with `{`, `.` or `?.`, or after
// a line ending on an operator, `,` or `.`.
func (s *splitter) lineEndSplits(k int) bool {
	if !s.atBlockLevel() || s.cur.Len() == 0 {
		return false
	}
	if s.lastCls == clsOp || s.lastRune == ',' || s.lastRune == '.' {
		return false
	}
	t := s.toks[k]
	if t.kind != tkCode {
		return true
	}
	switch t.text {
	case "{", ".":
		return false
	case "?":
		return !s.nextIsDot(k + 1)
	}
	return true
}

// nextIsDot reports whether the next non-space token from k is a code `.`.
func (s *splitter) nextIsDot(k int) bool {
	for ; k < len(s.toks); k++ {
		if s.toks[k].kind != tkSpace {
			return s.toks[k].kind == tkCode && s.toks[k].text == "."
		}
	}
	return false
}

// needSpace keeps one space between two word or two operator characters.
func (s *splitter) needSpace(t token) bool {
	if s.cur.Len() == 0 || t.kind != tkCode {
		return false
	}
	r, _ := utf8.DecodeRuneInString(t.text)
	c := classOf(r)
	return c != clsNone && c == s.lastCls
}

// boundary ends a statement at top level, or writes one `;` inside a block —
// never directly after `{` or a `;`.
func (s *splitter) boundary() {
	if len(s.stack) == 0 {
		s.finish()
		return
	}
	if s.lastRune != '{' && s.lastRune != ';' {
		s.cur.WriteByte(';')
		s.lastCls, s.lastRune = clsNone, ';'
	}
}

// emit appends a token, tracking the bracket stack.
func (s *splitter) emit(t token) error {
	if s.cur.Len() == 0 {
		s.curLine = t.line
	}
	if t.kind == tkLit {
		s.cur.WriteString(t.text)
		s.lastCls, s.lastRune = clsNone, '"'
		return nil
	}
	switch t.text {
	case "(", "[", "{":
		s.stack = append(s.stack, t)
	case ")", "]", "}":
		if err := s.close(t); err != nil {
			return err
		}
	}
	s.cur.WriteString(t.text)
	r, _ := utf8.DecodeRuneInString(t.text)
	s.lastCls, s.lastRune = classOf(r), r
	return nil
}

// close pops the matching opener; a `;` written right before `}` is dropped.
func (s *splitter) close(t token) error {
	n := len(s.stack)
	want := map[string]string{")": "(", "]": "[", "}": "{"}[t.text]
	if n == 0 || s.stack[n-1].text != want {
		return fmt.Errorf("in Zeile %d: Klammer %q ohne passendes Gegenstück", t.line, t.text)
	}
	s.stack = s.stack[:n-1]
	if t.text == "}" && strings.HasSuffix(s.cur.String(), ";") {
		str := s.cur.String()
		s.cur.Reset()
		s.cur.WriteString(str[:len(str)-1])
	}
	return nil
}

// finish closes the current top-level statement, if any.
func (s *splitter) finish() {
	if s.cur.Len() > 0 {
		s.stmts = append(s.stmts, core.Statement{Text: s.cur.String(), Line: s.curLine})
	}
	s.cur.Reset()
	s.lastCls, s.lastRune = clsNone, 0
}

// countLines is the file's line count for messages without a statement of their
// own: the number of lines, at least 1 (SPEC-EXTRACT-001).
func countLines(src string) int {
	src = strings.ReplaceAll(strings.ReplaceAll(src, "\r\n", "\n"), "\r", "\n")
	n := strings.Count(src, "\n")
	if src != "" && !strings.HasSuffix(src, "\n") {
		n++
	}
	if n < 1 {
		return 1
	}
	return n
}
