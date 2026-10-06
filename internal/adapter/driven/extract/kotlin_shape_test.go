package extract

import (
	"strings"
	"testing"

	"github.com/pt9912/a-check/internal/hexagon/core"
)

func texts(st []core.Statement) []string {
	out := make([]string, 0, len(st))
	for _, s := range st {
		out = append(out, s.Text)
	}
	return out
}

func mustNorm(t *testing.T, src string) []core.Statement {
	t.Helper()
	st, _, err := normalizeKotlin(src)
	if err != nil {
		t.Fatalf("normalizeKotlin(%q): %v", src, err)
	}
	return st
}

func eqTexts(t *testing.T, src string, want ...string) {
	t.Helper()
	got := texts(mustNorm(t, src))
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Fatalf("normalizeKotlin(%q)\n got: %q\nwant: %q", src, got, want)
	}
}

// AC-FA-RULE-012 happy: Leerraum, Kommentare und Zeilenumbrüche an beliebiger
// Stelle ändern die normalisierte Form nicht.
func TestKotlinNormalizeFormatIrrelevant(t *testing.T) {
	compact := "plugins{kotlin(\"jvm\")}\nkotlin{jvmToolchain(25)}\ndependencies{testImplementation(kotlin(\"test\"))}\n"
	loose := `// Kopf
plugins   /* inline */
{
    kotlin( "jvm" )   // Plugin
}

kotlin {
  jvmToolchain(
    25
  )
}
/* Block
   über Zeilen */
dependencies {
    testImplementation(kotlin("test"))
}
`
	want := []string{`plugins{kotlin("jvm")}`, `kotlin{jvmToolchain(25)}`, `dependencies{testImplementation(kotlin("test"))}`}
	eqTexts(t, compact, want...)
	eqTexts(t, loose, want...)
}

// Ein Block ist EINE Anweisung; Grenzen darin werden zu `;` (SPEC-EXTRACT-001
// Schritt 3) — sonst fielen zwei Aufrufe zu `a()b()` zusammen.
func TestKotlinBlockStatementsSeparated(t *testing.T) {
	eqTexts(t, "dependencies {\n  a()\n\n  b()\n}\n", "dependencies{a();b()}")
	eqTexts(t, "dependencies { a(); b(); }", "dependencies{a();b()}")
	eqTexts(t, "x {\n}\n", "x{}")
}

// Gegenprobe-Fall des CR: plugins mit Zeilenumbruch vor `{` und einem weiteren
// Plugin ist eine Anweisung, die nicht der erlaubten gleicht.
func TestKotlinNewlineBeforeBrace(t *testing.T) {
	eqTexts(t, "plugins\n{\n  kotlin(\"jvm\")\n  id(\"evil\")\n}\n", `plugins{kotlin("jvm");id("evil")}`)
}

// Fortsetzung: `.`/`?.` am Zeilenanfang, Operator/`,`/`.` am Zeilenende, und
// innerhalb von ( [ ist ein Zeilenende Leerraum.
func TestKotlinContinuation(t *testing.T) {
	eqTexts(t, "tasks\n  .test\n  ?.run()\n", "tasks.test?.run()")
	eqTexts(t, "compileClasspath +=\n  files(\"x\")\n", `compileClasspath+=files("x")`)
	eqTexts(t, "add(\"a\",\n \"b\")\n", `add("a","b")`)
	// Postfix !! am Zeilenende verbindet zwei Anweisungen: fail-safe, als
	// Fehlalarm-Quelle bekannt (Review slice-209 F-13).
	eqTexts(t, "x!!\ny()\n", "x!!y()")
}

// Leerraum bleibt genau dann (als ein Leerzeichen), wenn er zwei Wort- oder zwei
// Operatorzeichen trennt (ADR-0041 Punkt 3).
func TestKotlinWhitespaceFolding(t *testing.T) {
	eqTexts(t, "val   a   =   1", "val a=1")
	eqTexts(t, "x = a - -b", "x=a- -b")
	eqTexts(t, "x = a--b", "x=a--b")
	eqTexts(t, `id "x"`, `id"x"`)
}

// Zeichenketten bleiben byte-genau; `//` und `/*` darin öffnen keinen Kommentar,
// Klammern darin zählen nicht.
func TestKotlinStringsVerbatim(t *testing.T) {
	eqTexts(t, `a("// kein  Kommentar { ")`, `a("// kein  Kommentar { ")`)
	eqTexts(t, `a('{')`+"\n"+`b('\'')`, `a('{')`, `b('\'')`)
	eqTexts(t, "a(`x y`)", "a(`x y`)")
	eqTexts(t, "a(\"\"\"\n  } { \" \"\" \n\"\"\")", "a(\"\"\"\n  } { \" \"\" \n\"\"\")")
	// überzählige Quotes vor den letzten drei gehören zum Inhalt
	eqTexts(t, `a("""x"""")`, `a("""x"""")`)
	eqTexts(t, `a("${ "}" + f({ 1 }) }")`, `a("${ "}" + f({ 1 }) }")`)
}

// Verschachtelte Block-Kommentare, Shebang, BOM, CRLF.
func TestKotlinCommentsAndLineEnds(t *testing.T) {
	eqTexts(t, "/* a /* b */ still comment */ x()", "x()")
	eqTexts(t, "#!/usr/bin/env kotlin\nx()", "x()")
	eqTexts(t, "\uFEFFx()", "x()")
	st := mustNorm(t, "a()\r\nb()\rc()\n")
	if got := texts(st); strings.Join(got, ",") != "a(),b(),c()" || st[2].Line != 3 {
		t.Fatalf("CRLF/CR: %+v", st)
	}
}

// Gegenprobe-Fall des CR: eine Anweisung, die im Kommentar „beginnt" und im Code
// fortgesetzt wird, ist eine eigene, unerlaubte Anweisung.
func TestKotlinStatementHiddenInComment(t *testing.T) {
	st := mustNorm(t, "plugins { kotlin(\"jvm\") }\n/* dependencies { */ implementation(\"evil\")\n")
	if got := texts(st); len(got) != 2 || got[1] != `implementation("evil")` || st[1].Line != 2 {
		t.Fatalf("hidden statement: %+v", st)
	}
}

// Review slice-209 F-3/N-1/K-1: Dollar-Präfix und maskiertes `\$` dürfen den
// Lexer nicht dazu bringen, eine Zeichenkette zu früh oder zu spät zu schließen
// — sonst verschwindet Code als Kommentar.
func TestKotlinDollarStringsDoNotSwallowCode(t *testing.T) {
	cases := map[string]string{
		"F-3 dollar prefix":    `version = "x" + $$"${" + '"' + '}' + "// "; dependencies.add("implementation", "com.evil:lib:1.0")`,
		"N-1 escaped dollar":   `version = "x" + "\${" + '"' + '}' + "// "; dependencies.add("implementation", "com.evil:lib:1.0")`,
		"K-1 escaped in $$":    `version = "x" + $$"\$${" + '"' + '}' + "// "; dependencies.add("implementation", "com.evil:lib:1.0")`,
		"raw $$ no template":   `version = "x" + $$"""${""" + '"' + '}' + "// "; dependencies.add("implementation", "com.evil:lib:1.0")`,
		"single-dollar string": `version = "x" + "$" + "// "; dependencies.add("implementation", "com.evil:lib:1.0")`,
	}
	for name, src := range cases {
		st := mustNorm(t, src)
		if len(st) != 2 || !strings.HasPrefix(st[1].Text, "dependencies.add(") {
			t.Errorf("%s: add-Aufruf verschluckt: %q", name, texts(st))
		}
	}
	// ein echtes Template mit n=2: `$$"\$$${x}"` — maskiertes $, dann Folge 2
	eqTexts(t, `a($$"\$$${ "}" }")`, `a($$"\$$${ "}" }")`)
}

// Nicht zerlegbare Dateien sind ein Fehler (Exit 2), nie still grün.
func TestKotlinUnsplittable(t *testing.T) {
	for _, src := range []string{
		"plugins {\n  kotlin(\"jvm\")\n",
		"a(\"offen\n)",
		"/* offen",
		"a(]",
		"}",
		`a("${ x ")`,
		"a('x",
		"a(`x",
		`a("""offen`,
	} {
		if _, _, err := normalizeKotlin(src); err == nil {
			t.Errorf("normalizeKotlin(%q): Fehler erwartet", src)
		}
	}
}

func TestKotlinLinesAndStatementLines(t *testing.T) {
	st, lines, err := normalizeKotlin("\n\n  // c\n  a(\n1)\nb()\n")
	if err != nil || lines != 6 || st[0].Line != 4 || st[1].Line != 6 {
		t.Fatalf("lines=%d st=%+v err=%v", lines, st, err)
	}
	if _, lines, _ := normalizeKotlin(""); lines != 1 {
		t.Fatalf("leere Datei: lines=%d, want 1", lines)
	}
	if st, _, _ := normalizeKotlin("// nur Kommentar\n"); len(st) != 0 {
		t.Fatalf("Datei ohne Code: %+v", st)
	}
}

// Review slice-210 F-2: $`name` ist auch in einer Zeichenkette eine Vorlage — ein
// `"` im Backtick-Bezeichner beendet die Zeichenkette nicht, sonst schluckt ein
// späteres /* echten Code.
func TestKotlinBacktickTemplateInString(t *testing.T) {
	// ohne Backtick-Vorlage schlösse der erste String nach `$`, der Backtick
	// liefe bis in den zweiten String, und `/*` schluckte evil() bis zum */
	src := "val s = \"$`\"`\" + \"`/*\"\nevil()\nx() // */\n"
	st := mustNorm(t, src)
	if got := texts(st); len(got) != 3 || got[1] != "evil()" {
		t.Fatalf("Code verschluckt: %q", got)
	}
	eqTexts(t, "a($$\"$`x`\")", "a($$\"$`x`\")") // n=2: ein $ ist Inhalt, kein Backtick-Template
	if _, _, err := normalizeKotlin("a(\"$`offen\n`\")"); err == nil {
		t.Fatal("Zeilenende im Backtick-Bezeichner: Fehler erwartet")
	}
}

// Review slice-210 F-9: Backslash vor Zeilenende im Zeichen-Literal ist ein Fehler.
func TestKotlinCharEscapeBeforeLineEnd(t *testing.T) {
	if _, _, err := normalizeKotlin("a('\\\n')"); err == nil {
		t.Fatal("Fehler erwartet")
	}
}

// Review slice-210 F-5: Zeilenende nach Postfix/Operator verbindet zwei
// Anweisungen — fail-safe (rot), als bekannte Fehlalarm-Quelle belegt.
func TestKotlinOperatorLineEndJoins(t *testing.T) {
	eqTexts(t, "i++\nj()\n", "i++j()")
	eqTexts(t, "i--\nj()\n", "i--j()")
	eqTexts(t, "listOf<A>\nj()\n", "listOf<A>j()")
}
