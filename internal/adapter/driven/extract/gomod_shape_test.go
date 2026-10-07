package extract

import (
	"strings"
	"testing"
)

func gomodTexts(t *testing.T, src string) []string {
	t.Helper()
	st, _, err := normalizeGomod(src)
	if err != nil {
		t.Fatalf("normalizeGomod(%q): %v", src, err)
	}
	return texts(st)
}

func eqGomod(t *testing.T, src string, want ...string) {
	t.Helper()
	if got := gomodTexts(t, src); strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Fatalf("normalizeGomod(%q)\n got: %q\nwant: %q", src, got, want)
	}
}

// Ein go.mod in realer Form (Block-Form, `// indirect`): Kommentare entfallen,
// Tokens durch genau ein Leerzeichen, ein Block ist eine Anweisung mit
// `;`-getrennten Einträgen (SPEC-EXTRACT-001, gomod Schritt 2/3).
func TestGomodRealForm(t *testing.T) {
	src := "module github.com/x/y\n\ngo 1.27.0\n\nrequire (\n\tgithub.com/a/b v1.2.3\n\tgolang.org/x/c v0.1.0 // indirect\n)\n"
	want := []string{"module github.com/x/y", "go 1.27.0", "require (github.com/a/b v1.2.3;golang.org/x/c v0.1.0)"}
	eqGomod(t, src, want...)
	// Einrückung, Leerraum innerhalb der Zeilen, CRLF und Kommentare ändern nichts
	loose := "// Kopf\r\nmodule   github.com/x/y \r\n\r\ngo\t1.27.0\r\nrequire (   // Block\r\n    github.com/a/b    v1.2.3\r\n\r\n  golang.org/x/c v0.1.0\r\n)\r\n"
	eqGomod(t, loose, want...)
}

// Einzeilige Direktive, `=>` als eigenes Token, Zeichenketten byte-genau.
func TestGomodDirectivesAndStrings(t *testing.T) {
	eqGomod(t, "require gopkg.in/yaml.v3 v3.0.1\n", "require gopkg.in/yaml.v3 v3.0.1")
	eqGomod(t, "replace a/b => ../b\n", "replace a/b => ../b")
	eqGomod(t, "module \"x  y\"\n", `module "x  y"`)
	eqGomod(t, "module `raw\\n`\n", "module `raw\\n`")
	eqGomod(t, `module "a\"b"`+"\n", `module "a\"b"`)
}

// Das Zeilenende ist Grammatik: eine auf zwei Zeilen verteilte Direktive ist eine
// andere Anweisungsfolge (ADR-0042 Punkt 3).
func TestGomodLineEndIsGrammar(t *testing.T) {
	eqGomod(t, "require a/b\n  v1.0.0\n", "require a/b", "v1.0.0")
}

// `/* */` ist in go.mod kein Kommentar, sondern verboten (Referenz: „not allowed").
func TestGomodNoBlockComment(t *testing.T) {
	for _, src := range []string{"/* x */ module a\n", "module a /* x */\n", "module a/*b\n"} {
		if _, _, err := normalizeGomod(src); err == nil {
			t.Errorf("%q: Fehler erwartet", src)
		}
	}
}

// Leere Blöcke: `require ()` und `require ( )` auf oberster Ebene; `()` an
// anderer Stelle ist ein Fehler.
func TestGomodEmptyBlock(t *testing.T) {
	eqGomod(t, "require ()\n", "require ()")
	eqGomod(t, "require ( )\n", "require ()")
	eqGomod(t, "require (\n)\n", "require ()")
}

// Alle fail-closed-Fälle der Spezifikation: Exit 2, nie still grün.
func TestGomodUnsplittable(t *testing.T) {
	for _, src := range []string{
		"module \"offen\n",             // Zeilenende in Zeichenkette
		"module `offen\nx`\n",          // Zeilenende in roher Zeichenkette
		"module \"a\\\n\"\n",           // Zeilenende direkt nach Backslash
		"module a;b\n",                 // `;` außerhalb einer Zeichenkette
		"module a\"b\"\n",              // Zeichenkette mitten im Token
		"module a`b`\n",                // rohe Zeichenkette mitten im Token
		"module ( x\n",                 // `(` als Token mitten in der Zeile
		"()\n",                         // leerer Block ohne Kopf
		"module \"a\"\"b\"\n",          // zweite Zeichenkette direkt hinter der ersten
		"module a//b\n",                // `//` direkt hinter einem Zeichen
		"module a(b\n",                 // Klammer im Token
		"replace a=>b\n",               // verklebtes =>
		"replace a ==> b\n",            // `==>`
		"replace \"a\"=> b\n",          // => an Zeichenkette geklebt
		"require (\n a v1 (\n)\n",      // Schachtelung
		"require (\n a v1\n) x\n",      // `)` mit weiteren Tokens
		")\n",                          // `)` außerhalb eines Blocks
		"require (\n a v1\n",           // offener Block am Dateiende
		"(\n)\n",                       // Block ohne Kopf
		"x () y\n",                     // `()` mitten in der Zeile
		"require (\n a () \n)\n",       // `()` in einem Eintrag
	} {
		if _, _, err := normalizeGomod(src); err == nil {
			t.Errorf("%q: Fehler erwartet", src)
		}
	}
}

// Zeilennummern und Zeilenzahl zählen LF; CR ist Leerraum und zählt nicht.
func TestGomodLines(t *testing.T) {
	st, lines, err := normalizeGomod("\n// c\nmodule a\r\nrequire (\n b v1\n)\n")
	if err != nil || len(st) != 2 || st[0].Line != 3 || st[1].Line != 4 || lines != 6 {
		t.Fatalf("st=%+v lines=%d err=%v", st, lines, err)
	}
	if _, lines, _ := normalizeGomod("module a\rgo 1.2"); lines != 1 {
		t.Fatalf("lone CR: lines=%d, want 1", lines)
	}
	if _, lines, _ := normalizeGomod(""); lines != 1 {
		t.Fatalf("leer: lines=%d, want 1", lines)
	}
}
