package core

import (
	"strings"
	"testing"
)

func mustShape(t *testing.T, allow ...ShapeAllowSpec) Shape {
	t.Helper()
	sh, err := NewShape([]string{"mod/build.gradle.kts"}, "kotlin", "allow-statements", allow, "", "")
	if err != nil {
		t.Fatal(err)
	}
	return sh
}

// AC-FA-RULE-012 negative: die fail-closed-Fälle der Konfiguration (Exit 2).
func TestNewShapeFailClosed(t *testing.T) {
	ok := []ShapeAllowSpec{{Pattern: "a()"}}
	cases := map[string]func() error{
		"leeres files":        func() error { _, e := NewShape(nil, "kotlin", "allow-statements", ok, "", ""); return e },
		"leerer glob":         func() error { _, e := NewShape([]string{""}, "kotlin", "allow-statements", ok, "", ""); return e },
		"absoluter glob":      func() error { _, e := NewShape([]string{"/etc/x"}, "kotlin", "allow-statements", ok, "", ""); return e },
		"glob aus der wurzel": func() error { _, e := NewShape([]string{"a/../../x"}, "kotlin", "allow-statements", ok, "", ""); return e },
		"dialect fehlt":       func() error { _, e := NewShape([]string{"x"}, "", "allow-statements", ok, "", ""); return e },
		"mode unbekannt":      func() error { _, e := NewShape([]string{"x"}, "kotlin", "deny", ok, "", ""); return e },
		"allow leer":          func() error { _, e := NewShape([]string{"x"}, "kotlin", "allow-statements", nil, "", ""); return e },
		"expect bei allow":    func() error { _, e := NewShape([]string{"x"}, "kotlin", "allow-statements", ok, "", "s.kts"); return e },
		"leeres pattern":      func() error { _, e := NewShape([]string{"x"}, "kotlin", "allow-statements", []ShapeAllowSpec{{}}, "", ""); return e },
		"match unbekannt": func() error {
			_, e := NewShape([]string{"x"}, "kotlin", "allow-statements", []ShapeAllowSpec{{Pattern: "a", Match: "substring"}}, "", "")
			return e
		},
		"regex kaputt": func() error {
			_, e := NewShape([]string{"x"}, "kotlin", "allow-statements", []ShapeAllowSpec{{Pattern: "a(", Match: "regex"}}, "", "")
			return e
		},
		// kompiliert nur umhüllt — umhüllt träfe es jede Anweisung (Review F-6)
		"regex nur umhuellt": func() error {
			_, e := NewShape([]string{"x"}, "kotlin", "allow-statements", []ShapeAllowSpec{{Pattern: "a)|(.*", Match: "regex"}}, "", "")
			return e
		},
		// kompiliert nur für sich — umhüllt nicht (Review N-2)
		"regex nur fuer sich": func() error {
			_, e := NewShape([]string{"x"}, "kotlin", "allow-statements", []ShapeAllowSpec{{Pattern: `\Qfoo`, Match: "regex"}}, "", "")
			return e
		},
	}
	for name, f := range cases {
		if f() == nil {
			t.Errorf("%s: Fehler erwartet", name)
		}
	}
}

// Mengen-Semantik, voll verankerter Regex, literal-Gleichheit.
func TestEvaluateShapesUnlisted(t *testing.T) {
	m := Model{Shapes: []Shape{mustShape(t,
		ShapeAllowSpec{Pattern: "plugins { kotlin(\"jvm\") }"},
		ShapeAllowSpec{Pattern: `kotlin\{jvmToolchain\(\d+\)\}`, Match: "regex"},
	)}}
	scan := ShapeScan{
		Literals: [][]string{{`plugins{kotlin("jvm")}`, ""}},
		Files: []ShapeFile{{Entry: 0, Path: "mod/build.gradle.kts", Lines: 9, Statements: []Statement{
			{Text: `kotlin{jvmToolchain(25)}`, Line: 1},
			{Text: `plugins{kotlin("jvm")}`, Line: 3},
			{Text: `plugins{kotlin("jvm")}`, Line: 4}, // Wiederholung erlaubt
			{Text: `kotlin{jvmToolchain(25)};x()`, Line: 5},
			{Text: `compileClasspath+=files("x")`, Line: 7},
		}}},
	}
	fs := EvaluateShapes(m, scan, ".a-check.yml")
	if len(fs) != 2 || fs[0].Line != 5 || fs[1].Line != 7 || fs[0].Rule != "shape-unlisted" || fs[1].Msg != `compileClasspath+=files("x")` {
		t.Fatalf("findings: %+v", fs)
	}
}

// Dieselbe Datei in zwei Einträgen: byte-gleiche Befundzeilen einmal.
func TestEvaluateShapesDedupe(t *testing.T) {
	sh := mustShape(t, ShapeAllowSpec{Pattern: "a()"})
	m := Model{Shapes: []Shape{sh, sh}}
	f := func(e int) ShapeFile {
		return ShapeFile{Entry: e, Path: "x.kts", Lines: 1, Statements: []Statement{{Text: "b()", Line: 1}}}
	}
	fs := EvaluateShapes(m, ShapeScan{Literals: [][]string{{"a()"}, {"a()"}}, Files: []ShapeFile{f(1), f(0)}}, ".a-check.yml")
	if len(fs) != 1 || !strings.Contains(fs[0].Msg, "b()") {
		t.Fatalf("dedupe: %+v", fs)
	}
}

// SortFindings fasst nichts zusammen — zwei gleiche constructs-Befunde bleiben zwei.
func TestSortFindingsKeepsDuplicates(t *testing.T) {
	fs := []Finding{{"b", 1, "r", "m"}, {"a", 1, "r", "m"}, {"a", 1, "r", "m"}}
	SortFindings(fs)
	if len(fs) != 3 || fs[0].Path != "a" || fs[2].Path != "b" {
		t.Fatalf("%+v", fs)
	}
}

func TestInsideRoot(t *testing.T) {
	for _, p := range []string{"a/b", "**/x.kts", "a/./b"} {
		if err := InsideRoot("files", p); err != nil {
			t.Errorf("%q: %v", p, err)
		}
	}
	for _, p := range []string{"", "/a", "..", "../a", "a/../../b"} {
		if InsideRoot("files", p) == nil {
			t.Errorf("%q: Fehler erwartet", p)
		}
	}
}

// Die dokumentierte sichere Klasse für Zeichenketten-Inhalt hält, ein `.*`
// überspannt Code (benannte Grenze, ADR-0041 Punkt 5; Handbuch §4).
func TestShapeRegexSafeClass(t *testing.T) {
	smuggled := `version="1"+run{dependencies.add("x","evil")}+""`
	safe := mustShape(t, ShapeAllowSpec{Pattern: `version="[^"$\\]*"`, Match: "regex"})
	if !safe.Allow[0].match(`version="1.2.3"`) || safe.Allow[0].match(smuggled) || safe.Allow[0].match(`version="${evil()}"`) {
		t.Fatal("sichere Klasse: erwartet trifft 1.2.3, nicht Code, nicht Vorlage")
	}
	loose := mustShape(t, ShapeAllowSpec{Pattern: `version=".*"`, Match: "regex"})
	if !loose.Allow[0].match(smuggled) {
		t.Fatal("benannte Grenze: .* überspannt Code — trifft sie nicht mehr, ist die Doku zu korrigieren")
	}
}

// Review slice-210 F-7: ein führendes ./ wird entfernt, sonst träfe der Glob nie.
func TestNewShapeTrimsDotSlash(t *testing.T) {
	sh, err := NewShape([]string{"./mod/b.kts", "././c.kts"}, "kotlin", "allow-statements", []ShapeAllowSpec{{Pattern: "a()"}}, "", "")
	if err != nil || sh.Files[0] != "mod/b.kts" || sh.Files[1] != "c.kts" {
		t.Fatalf("files=%v err=%v", sh.Files, err)
	}
}

// Review slice-210 D-1: zwei gleiche unerlaubte Anweisungen auf einer Zeile
// ergeben eine Befundzeile (SPEC-CONF-001, byte-gleich einmal).
func TestEvaluateShapesSameLineOnce(t *testing.T) {
	m := Model{Shapes: []Shape{mustShape(t, ShapeAllowSpec{Pattern: "a()"})}}
	f := ShapeFile{Entry: 0, Path: "x.kts", Lines: 1, Statements: []Statement{{Text: "b()", Line: 1}, {Text: "b()", Line: 1}, {Text: "c()", Line: 1}}}
	fs := EvaluateShapes(m, ShapeScan{Literals: [][]string{{"a()"}}, Files: []ShapeFile{f}}, ".a-check.yml")
	if len(fs) != 2 || fs[0].Msg != "b()" || fs[1].Msg != "c()" {
		t.Fatalf("%+v", fs)
	}
}

func exactShape(t *testing.T) Shape {
	t.Helper()
	sh, err := NewShape([]string{"x.kts"}, "kotlin", "exact", nil, "", "./sollform/x.kts")
	if err != nil || sh.Expect != "sollform/x.kts" {
		t.Fatalf("sh=%+v err=%v", sh, err)
	}
	return sh
}

// AC-FA-RULE-012 boundary (exact): gleich ist grün; anders, zusätzlich, fehlend
// ist je GENAU ein Befund shape-differs mit der ersten Abweichung.
func TestEvaluateShapesExact(t *testing.T) {
	m := Model{Shapes: []Shape{exactShape(t)}}
	want := []Statement{{Text: "a()", Line: 1}, {Text: "b()", Line: 2}}
	run := func(got ...Statement) []Finding {
		f := ShapeFile{Entry: 0, Path: "x.kts", Lines: 7, Statements: got}
		return EvaluateShapes(m, ShapeScan{Literals: [][]string{nil}, Expected: [][]Statement{want}, Files: []ShapeFile{f}}, ".a-check.yml")
	}
	if fs := run(Statement{"a()", 3}, Statement{"b()", 5}); len(fs) != 0 {
		t.Fatalf("gleich: %+v", fs)
	}
	cases := []struct {
		got  []Statement
		line int
		msg  string
	}{
		{[]Statement{{"a()", 1}, {"c()", 4}, {"d()", 6}}, 4, "c() (erwartet: b())"},
		{[]Statement{{"a()", 1}, {"b()", 2}, {"evil()", 3}}, 3, "evil() (nicht in der Sollform)"},
		{[]Statement{{"a()", 1}}, 7, "fehlt: b()"},
	}
	for _, c := range cases {
		fs := run(c.got...)
		if len(fs) != 1 || fs[0].Rule != "shape-differs" || fs[0].Line != c.line || fs[0].Msg != c.msg {
			t.Errorf("got %+v, want line %d msg %q", fs, c.line, c.msg)
		}
	}
}

// AC-FA-RULE-012 boundary (unused): mit unused: fail meldet ein trefferloser
// Eintrag an seiner Config-Zeile; ohne nicht. Ein Eintrag, der nur als zweiter
// passt, gilt als getroffen.
func TestEvaluateShapesUnused(t *testing.T) {
	allow := []ShapeAllowSpec{{Pattern: "a()", Line: 10}, {Pattern: `a\(\)`, Match: "regex", Line: 11}, {Pattern: "x {\n  y()\n}\n", Line: 12}}
	on, err := NewShape([]string{"x.kts"}, "kotlin", "allow-statements", allow, "fail", "")
	if err != nil {
		t.Fatal(err)
	}
	off, _ := NewShape([]string{"x.kts"}, "kotlin", "allow-statements", allow, "", "")
	f := ShapeFile{Entry: 0, Path: "x.kts", Lines: 1, Statements: []Statement{{Text: "a()", Line: 1}}}
	scan := ShapeScan{Literals: [][]string{{"a()", "", "x{y()}"}}, Files: []ShapeFile{f}}
	fs := EvaluateShapes(Model{Shapes: []Shape{on}}, scan, ".a-check.yml")
	if len(fs) != 1 || fs[0].Path != ".a-check.yml" || fs[0].Line != 12 || fs[0].Rule != "shape-unused" || fs[0].Msg != `literal: x {\n  y()\n}` {
		t.Fatalf("unused: %+v", fs)
	}
	if fs := EvaluateShapes(Model{Shapes: []Shape{off}}, scan, ".a-check.yml"); len(fs) != 0 {
		t.Fatalf("ohne unused: %+v", fs)
	}
	if _, err := NewShape([]string{"x.kts"}, "kotlin", "allow-statements", allow, "warn", ""); err == nil {
		t.Fatal("unused: warn — Fehler erwartet (kein Warn-Level)")
	}
}

// exact-Kombinationen fail-closed.
func TestNewShapeExactFailClosed(t *testing.T) {
	ok := []ShapeAllowSpec{{Pattern: "a()"}}
	for name, f := range map[string]func() error{
		"expect fehlt":   func() error { _, e := NewShape([]string{"x"}, "kotlin", "exact", nil, "", ""); return e },
		"allow bei exact": func() error { _, e := NewShape([]string{"x"}, "kotlin", "exact", ok, "", "s"); return e },
		"unused bei exact": func() error { _, e := NewShape([]string{"x"}, "kotlin", "exact", nil, "fail", "s"); return e },
		"expect hinaus":  func() error { _, e := NewShape([]string{"x"}, "kotlin", "exact", nil, "", "../s"); return e },
	} {
		if f() == nil {
			t.Errorf("%s: Fehler erwartet", name)
		}
	}
}

// Review slice-211 F-1/F-2/F-4: shape-unused-Meldung einzeilig (auch bei CR),
// Regex-Zusatz, und Zählung über mehrere Dateien eines Eintrags.
func TestEvaluateShapesUnusedFormAndFiles(t *testing.T) {
	allow := []ShapeAllowSpec{{Pattern: "q()\r", Line: 3}, {Pattern: `r\(\)`, Match: "regex", Line: 4}, {Pattern: "b()", Line: 5}}
	sh, err := NewShape([]string{"*.kts"}, "kotlin", "allow-statements", allow, "fail", "")
	if err != nil {
		t.Fatal(err)
	}
	f1 := ShapeFile{Entry: 0, Path: "a.kts", Lines: 1, Statements: []Statement{{Text: "x()", Line: 1}}}
	f2 := ShapeFile{Entry: 0, Path: "b.kts", Lines: 1, Statements: []Statement{{Text: "b()", Line: 1}}}
	fs := EvaluateShapes(Model{Shapes: []Shape{sh}}, ShapeScan{Literals: [][]string{{"q()", "", "b()"}}, Files: []ShapeFile{f1, f2}}, ".a-check.yml")
	var unused []string
	for _, f := range fs {
		if f.Rule == "shape-unused" {
			unused = append(unused, f.Msg)
		}
	}
	// b() trifft nur in der ZWEITEN Datei und gilt als getroffen; q() ohne abschließendes CR,
	// die Art als Präfix (ADR-0042 Punkt 5)
	if strings.Join(unused, "|") != `literal: q()|regex: r\\(\\)` {
		t.Fatalf("unused=%q all=%+v", unused, fs)
	}
}

// SPEC-RULE-001 Ausgabe-Regel (ADR-0042 Punkt 5): einzeilig und umkehrbar — eine
// rohes Zeilenende und ein wörtliches `\n` ergeben verschiedene Meldungen.
func TestOneLineInjective(t *testing.T) {
	cases := map[string]string{
		"a\nb":   `a\nb`,
		`a\nb`:   `a\\nb`,
		"a\r\nb": `a\r\nb`,
		`a\b`:    `a\\b`,
		"plain": "plain",
	}
	seen := map[string]string{}
	for in, want := range cases {
		got := oneLine(in)
		if got != want || strings.ContainsAny(got, "\r\n") {
			t.Errorf("oneLine(%q) = %q, want %q", in, got, want)
		}
		if prev, dup := seen[got]; dup {
			t.Errorf("%q und %q ergeben dieselbe Meldung %q", prev, in, got)
		}
		seen[got] = in
	}
}

// shape-unlisted und shape-differs schreiben ihre Anweisungen einzeilig; zwei
// Anweisungen, die sich nur in rohem Zeilenende vs. wörtlichem `\n` unterscheiden,
// bleiben zwei Befunde (die Zusammenfassung byte-gleicher Zeilen greift nicht).
func TestEvaluateShapesMessagesOneLine(t *testing.T) {
	m := Model{Shapes: []Shape{mustShape(t, ShapeAllowSpec{Pattern: "a()"})}}
	f := ShapeFile{Entry: 0, Path: "x.kts", Lines: 3, Statements: []Statement{
		{Text: "s(\"\"\"x\ny\"\"\")", Line: 1}, {Text: `s("""x\ny""")`, Line: 1},
	}}
	fs := EvaluateShapes(m, ShapeScan{Literals: [][]string{{"a()"}}, Files: []ShapeFile{f}}, ".a-check.yml")
	if len(fs) != 2 || fs[0].Msg == fs[1].Msg {
		t.Fatalf("zwei verschiedene Anweisungen, zwei Befunde erwartet: %+v", fs)
	}
	for _, x := range fs {
		if strings.ContainsAny(x.Msg, "\r\n") {
			t.Fatalf("Meldung nicht einzeilig: %q", x.Msg)
		}
	}
	d, ok := firstDifference(ShapeFile{Path: "x", Lines: 1, Statements: []Statement{{Text: "a\nb", Line: 1}}}, []Statement{{Text: `a\nb`, Line: 1}})
	if !ok || d.Msg != `a\nb (erwartet: a\\nb)` {
		t.Fatalf("shape-differs: %+v", d)
	}
}
