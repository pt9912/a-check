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
	fs := EvaluateShapes(m, scan)
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
	fs := EvaluateShapes(m, ShapeScan{Literals: [][]string{{"a()"}, {"a()"}}, Files: []ShapeFile{f(1), f(0)}})
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
	fs := EvaluateShapes(m, ShapeScan{Literals: [][]string{{"a()"}}, Files: []ShapeFile{f}})
	if len(fs) != 2 || fs[0].Msg != "b()" || fs[1].Msg != "c()" {
		t.Fatalf("%+v", fs)
	}
}
