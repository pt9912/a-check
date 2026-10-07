package extract

import (
	"fmt"
	"strings"

	"github.com/pt9912/a-check/internal/hexagon/core"
)

// This file is the `gomod` dialect of the shapes rule (AC-FA-RULE-012,
// ADR-0042). Its lexis follows the Go Modules Reference, section "Lexical
// elements" of go.mod files: only LF is significant, `//` starts a comment
// where a token can start, `/* */` is not allowed, strings are "…" (backslash
// escapes) and `…` (raw). Where the reference is silent or forbids a form, the
// file is unsplittable (exit 2) — never guessed (SPEC-EXTRACT-001).

// normalizeGomod returns the normalized statements of a go.mod source and its
// line count (LF only, at least 1), or an error if it cannot be split.
func normalizeGomod(src string) ([]core.Statement, int, error) {
	src = strings.TrimPrefix(src, "\uFEFF")
	var (
		out   []core.Statement
		block *gomodBlock
	)
	lines := strings.Split(src, "\n")
	for i, ln := range lines {
		toks, err := gomodTokens(ln)
		if err != nil {
			return nil, 0, fmt.Errorf("in Zeile %d: %w", i+1, err)
		}
		if len(toks) == 0 {
			continue
		}
		st, next, err := gomodLine(toks, i+1, block)
		if err != nil {
			return nil, 0, fmt.Errorf("in Zeile %d: %w", i+1, err)
		}
		block = next
		if st != nil {
			out = append(out, *st)
		}
	}
	if block != nil {
		return nil, 0, fmt.Errorf("in Zeile %d: Block nicht geschlossen", block.line)
	}
	return out, gomodLineCount(src), nil
}

// gomodBlock is an open `Kopf (` block with its entries so far.
type gomodBlock struct {
	head    string
	line    int
	entries []string
}

// gomodLine classifies one line of tokens (SPEC-EXTRACT-001, gomod step 3). It
// returns the finished statement (if any) and the open block after the line.
func gomodLine(toks []string, line int, block *gomodBlock) (*core.Statement, *gomodBlock, error) {
	if block != nil {
		return gomodBlockLine(toks, block)
	}
	last := toks[len(toks)-1]
	switch {
	case last == "(":
		head := toks[:len(toks)-1]
		if len(head) == 0 || hasParenToken(head) {
			return nil, nil, fmt.Errorf("ein Block ohne gültigen Kopf")
		}
		return nil, &gomodBlock{head: strings.Join(head, " "), line: line}, nil
	case last == "()" || (len(toks) >= 2 && last == ")" && toks[len(toks)-2] == "("):
		head := toks[:len(toks)-1]
		if last == ")" {
			head = toks[:len(toks)-2]
		}
		if len(head) == 0 || hasParenToken(head) {
			return nil, nil, fmt.Errorf("ein leerer Block ohne gültigen Kopf")
		}
		st := core.Statement{Text: strings.Join(head, " ") + " ()", Line: line}
		return &st, nil, nil
	case hasParenToken(toks):
		return nil, nil, fmt.Errorf("eine Klammer an unzulässiger Stelle")
	}
	st := core.Statement{Text: strings.Join(toks, " "), Line: line}
	return &st, nil, nil
}

// gomodBlockLine handles a line inside an open block: a lone `)` closes it
// into one statement, any other parenthesis is an error (no nesting), every
// other line is an entry.
func gomodBlockLine(toks []string, block *gomodBlock) (*core.Statement, *gomodBlock, error) {
	if len(toks) == 1 && toks[0] == ")" {
		st := core.Statement{Text: block.head + " (" + strings.Join(block.entries, ";") + ")", Line: block.line}
		return &st, nil, nil
	}
	if hasParenToken(toks) {
		return nil, nil, fmt.Errorf("eine Klammer in einem Block-Eintrag (keine Schachtelung, `)` steht allein)")
	}
	block.entries = append(block.entries, strings.Join(toks, " "))
	return nil, block, nil
}

// hasParenToken reports whether a token is one of the punctuation forms `(`,
// `)` or `()` — tokens that only the block grammar may place.
func hasParenToken(toks []string) bool {
	for _, t := range toks {
		if t == "(" || t == ")" || t == "()" {
			return true
		}
	}
	return false
}

// gomodTokens splits one line (without its LF) into tokens: comments dropped,
// strings kept byte-exact inside their token (SPEC-EXTRACT-001, gomod step 1).
func gomodTokens(ln string) ([]string, error) {
	var toks []string
	for i := 0; i < len(ln); {
		if isGomodSpace(ln[i]) {
			i++
			continue
		}
		if strings.HasPrefix(ln[i:], "//") {
			break // comment where a token can start
		}
		end, err := gomodTokenEnd(ln, i)
		if err != nil {
			return nil, err
		}
		tok := ln[i:end]
		if err := checkGomodToken(tok); err != nil {
			return nil, err
		}
		toks = append(toks, tok)
		i = end
	}
	return toks, nil
}

// gomodTokenEnd returns the end of the token starting at i. A string may only
// open a token; a quote, a `//` or a `/*` inside a token, and a `;` anywhere
// outside a string, are errors.
func gomodTokenEnd(ln string, i int) (int, error) {
	j := i
	if ln[j] == '"' || ln[j] == '`' {
		end, err := gomodString(ln, j)
		if err != nil {
			return 0, err
		}
		j = end
	}
	for j < len(ln) && !isGomodSpace(ln[j]) {
		switch {
		case ln[j] == '"' || ln[j] == '`':
			return 0, fmt.Errorf("eine Zeichenkette beginnt mitten in einem Token")
		case strings.HasPrefix(ln[j:], "//"):
			return 0, fmt.Errorf("`//` direkt hinter einem Zeichen")
		case strings.HasPrefix(ln[j:], "/*"):
			return 0, fmt.Errorf("`/*` ist in go.mod nicht erlaubt")
		case ln[j] == ';':
			return 0, fmt.Errorf("`;` außerhalb einer Zeichenkette")
		}
		j++
	}
	return j, nil
}

// gomodString scans a string starting at i and returns the index after it. The
// line has no LF, so reaching its end means a line end inside the string — an
// error in both forms, also right after a backslash.
func gomodString(ln string, i int) (int, error) {
	q := ln[i]
	for j := i + 1; j < len(ln); j++ {
		switch {
		case q == '"' && ln[j] == '\\':
			j++
		case ln[j] == q:
			return j + 1, nil
		}
	}
	return 0, fmt.Errorf("eine Zeichenkette ist nicht geschlossen (Zeilenende in einer Zeichenkette)")
}

// checkGomodToken rejects the token forms that would make two sources fall on
// one normal form: a parenthesis inside a longer token (except `()`), and `=>`
// glued to other characters.
func checkGomodToken(tok string) error {
	rest := tok
	if tok[0] == '"' || tok[0] == '`' {
		end, err := gomodString(tok, 0)
		if err != nil {
			return err
		}
		rest = tok[end:] // the characters glued to the string, checked like any token
		if rest == "" {
			return nil
		}
	}
	isPunct := tok == "(" || tok == ")" || tok == "()"
	if !isPunct && strings.ContainsAny(rest, "()") {
		return fmt.Errorf("eine Klammer steht innerhalb eines Tokens (%s)", tok)
	}
	if strings.Contains(rest, "=>") && tok != "=>" {
		return fmt.Errorf("`=>` mit anderen Zeichen verklebt (%s)", tok)
	}
	return nil
}

func isGomodSpace(c byte) bool { return c == ' ' || c == '\t' || c == '\r' }

// gomodLineCount counts lines by LF only; a trailing line without LF counts,
// at least 1 (SPEC-EXTRACT-001, gomod step 3).
func gomodLineCount(src string) int {
	n := strings.Count(src, "\n")
	if src != "" && !strings.HasSuffix(src, "\n") {
		n++
	}
	if n < 1 {
		return 1
	}
	return n
}
