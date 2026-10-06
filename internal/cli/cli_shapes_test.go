package cli_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pt9912/a-check/internal/cli"
)

// shapesCfg is the CR's lead case: the build file of the domain module may hold
// only the listed statements (AC-FA-RULE-012). The file lies in no layer and no
// language — the rule is independent of both.
const shapesCfg = cfg + `shapes:
  - files: ["domain/build.gradle.kts"]
    dialect: kotlin
    mode: allow-statements
    allow:
      - 'plugins { kotlin("jvm") }'
      - 'kotlin { jvmToolchain(25) }'
      - |
        dependencies {
          testImplementation(kotlin("test"))
        }
      - {pattern: 'tasks\.test\{useJUnitPlatform\(\)\}', match: regex}
      - 'val note = "implementation(\"com.evil:lib:1.0\") steht nur im String"'
`

const goodGradle = `// Fachkern: keine Fremdabhängigkeit — implementation("com.evil:lib:1.0") nur im Kommentar
plugins
{
    kotlin("jvm")   // nur das
}

kotlin { jvmToolchain(25) }

dependencies {
    /* nur Tests */ testImplementation(kotlin("test"))
}
tasks.test { useJUnitPlatform() }
val note = "implementation(\"com.evil:lib:1.0\") steht nur im String"
`

func runShapes(t *testing.T, gradle string, extra map[string]string) (int, string, string) {
	t.Helper()
	files := map[string]string{
		".a-check.yml":            shapesCfg,
		"domain/build.gradle.kts": gradle,
		"internal/core/c.go":      "package core\n",
	}
	for k, v := range extra {
		files[k] = v
	}
	dir := writeRepo(t, files)
	var out, errb bytes.Buffer
	code := cli.Run([]string{dir}, &out, &errb)
	return code, out.String(), errb.String()
}

// AC-FA-RULE-012 happy: die Datei entspricht der Liste — Leerraum, Kommentare,
// Zeilenumbrüche egal; das verbotene Muster steht nur im Kommentar (Zeile 1) und
// in der Zeichenkette einer ERLAUBTEN Anweisung (val note) und meldet nicht.
func TestShapesGreen(t *testing.T) {
	if code, out, errs := runShapes(t, goodGradle, nil); code != 0 || out != "" {
		t.Fatalf("code=%d out=%q err=%q", code, out, errs)
	}
	// Gegenprobe: dieselbe Zeichenkette in einer NICHT gelisteten Anweisung meldet
	src := strings.Replace(goodGradle, "val note =", "val other =", 1)
	if code, out, _ := runShapes(t, src, nil); code != 1 || !strings.Contains(out, "domain/build.gradle.kts:13: shape-unlisted: val other=") {
		t.Fatalf("Gegenprobe: code=%d out=%q", code, out)
	}
}

// AC-FA-RULE-012 boundary: die Gegenprobe-Fälle des CR — je Form ein Befund.
func TestShapesRedCases(t *testing.T) {
	base := strings.Replace(goodGradle, `val note = "implementation(\"com.evil:lib:1.0\") steht nur im String"`, "", 1)
	cases := map[string]struct{ src, want string }{
		"zusaetzliche Anweisung": {base + "apply(plugin = \"evil\")\n", `apply(plugin="evil")`},
		"Anweisung im Block":     {strings.Replace(base, `testImplementation(kotlin("test"))`, "testImplementation(kotlin(\"test\"))\n    implementation(\"com.evil:lib:1.0\")", 1), `dependencies{testImplementation(kotlin("test"));implementation("com.evil:lib:1.0")}`},
		"plugins Umbruch + Plugin": {strings.Replace(base, `kotlin("jvm")   // nur das`, "kotlin(\"jvm\")\n    id(\"evil\")", 1), `plugins{kotlin("jvm");id("evil")}`},
		"compileClasspath":        {base + "compileClasspath += files(\"evil.jar\")\n", `compileClasspath+=files("evil.jar")`},
		"add mit run":             {base + "dependencies.add(\"x\", run { \"com.evil:lib:1.0\" })\n", `dependencies.add("x",run{"com.evil:lib:1.0"})`},
		"im Kommentar begonnen":   {base + "/* dependencies { */ implementation(\"evil\")\n", `implementation("evil")`},
	}
	for name, c := range cases {
		code, out, errs := runShapes(t, c.src, nil)
		if code != 1 || !strings.Contains(out, "shape-unlisted: "+c.want+"\n") {
			t.Errorf("%s: code=%d out=%q err=%q", name, code, out, errs)
		}
	}
}

// AC-FA-RULE-012 negative: Exit 2, nie still grün.
func TestShapesExit2(t *testing.T) {
	missing := strings.Replace(shapesCfg, "domain/build.gradle.kts", "domain/fehlt.gradle.kts", 1)
	excluded := shapesCfg + "exclude: [\"domain/**\"]\n"
	unknownDialect := strings.Replace(shapesCfg, "dialect: kotlin", "dialect: groovy", 1)
	cases := map[string]struct{ cfg, gradle string }{
		"fehlende Datei":    {missing, goodGradle},
		"exclude-Konflikt":  {excluded, goodGradle},
		"unbekannter dialect": {unknownDialect, goodGradle},
		"offener Block":     {shapesCfg, "plugins {\n kotlin(\"jvm\")\n"},
		"offene Zeichenkette": {shapesCfg, "plugins { kotlin(\"jvm) }\n"},
	}
	for name, c := range cases {
		code, out, errs := runShapes(t, c.gradle, map[string]string{".a-check.yml": c.cfg})
		if code != 2 || out != "" || !strings.Contains(errs, "shapes") {
			t.Errorf("%s: code=%d out=%q err=%q", name, code, out, errs)
		}
	}
}

// Der no-scan-Pfad prüft Dialekt und literal-Einträge wie ein Scan (Review F-5).
func TestShapesPrintGraphValidates(t *testing.T) {
	bad := strings.Replace(shapesCfg, `- 'kotlin { jvmToolchain(25) }'`, `- 'kotlin { jvmToolchain(25) }; extra()'`, 1)
	dir := writeRepo(t, map[string]string{".a-check.yml": bad})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-graph", dir}, &out, &errb); code != 2 || !strings.Contains(errb.String(), "genau eine") {
		t.Fatalf("code=%d err=%q", code, errb.String())
	}
}

// AC-QA-01: zwei Läufe, byte-identische Ausgabe.
func TestShapesDeterministic(t *testing.T) {
	src := goodGradle + "b()\na()\n"
	_, out1, _ := runShapes(t, src, nil)
	_, out2, _ := runShapes(t, src, nil)
	if out1 != out2 || strings.Count(out1, "\n") != 2 {
		t.Fatalf("out1=%q out2=%q", out1, out2)
	}
}

// Review slice-210 F-13: das --print-config-Gerüst zeigt den shapes-Block, und
// für Zeichenketten-Inhalt nur die sichere Klasse, nie `.*` (ADR-0041 Punkt 5).
func TestPrintConfigShowsShapes(t *testing.T) {
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-config"}, &out, &errb); code != 0 {
		t.Fatalf("code=%d", code)
	}
	s := out.String()
	if !strings.Contains(s, "# shapes:") || !strings.Contains(s, "mode: allow-statements") || !strings.Contains(s, `"[^"$\\]*"`) || strings.Contains(s, `".*"`) {
		t.Fatalf("Gerüst ohne shapes-Block oder mit .*: %q", s)
	}
}

// AC-FA-RULE-012 boundary (exact) end-to-end: gleich bis auf Leerraum und
// Kommentare ist grün; eine zusätzliche Anweisung ist genau ein shape-differs.
func TestShapesExactEndToEnd(t *testing.T) {
	exactCfg := cfg + `shapes:
  - files: ["domain/build.gradle.kts"]
    dialect: kotlin
    mode: exact
    expect: sollform/build.gradle.kts
`
	sollform := "plugins { kotlin(\"jvm\") }\nkotlin { jvmToolchain(25) }\n"
	run := func(gradle string, withSollform bool) (int, string, string) {
		files := map[string]string{".a-check.yml": exactCfg, "domain/build.gradle.kts": gradle, "internal/core/c.go": "package core\n"}
		if withSollform {
			files["sollform/build.gradle.kts"] = sollform
		}
		var out, errb bytes.Buffer
		code := cli.Run([]string{writeRepo(t, files)}, &out, &errb)
		return code, out.String(), errb.String()
	}
	if code, out, errs := run("// Kopf\nplugins\n{\n  kotlin(\"jvm\")\n}\n\nkotlin {\n  jvmToolchain(25)\n}\n", true); code != 0 || out != "" {
		t.Fatalf("gleich: code=%d out=%q err=%q", code, out, errs)
	}
	code, out, _ := run(sollform+"dependencies { implementation(\"evil\") }\n", true)
	if code != 1 || out != "domain/build.gradle.kts:3: shape-differs: dependencies{implementation(\"evil\")} (nicht in der Sollform)\n" {
		t.Fatalf("zusätzlich: code=%d out=%q", code, out)
	}
	if code, out, errs := run(sollform, false); code != 2 || out != "" || !strings.Contains(errs, "expect-Datei") {
		t.Fatalf("fehlende Sollform: code=%d out=%q err=%q", code, out, errs)
	}
}

// AC-FA-RULE-012 boundary (unused) end-to-end: der Befund zeigt auf die Zeile
// des Eintrags in der .a-check.yml.
func TestShapesUnusedEndToEnd(t *testing.T) {
	c := strings.Replace(shapesCfg, "    mode: allow-statements\n", "    mode: allow-statements\n    unused: fail\n", 1) +
		"      - 'apply(plugin = \"nie-benutzt\")'\n"
	code, out, errs := runShapes(t, goodGradle, map[string]string{".a-check.yml": c})
	want := strings.Count(c[:strings.Index(c, "nie-benutzt")], "\n") + 1 // Zeile des Eintrags
	if code != 1 || !strings.HasPrefix(out, fmt.Sprintf(".a-check.yml:%d: ", want)) || !strings.Contains(out, `shape-unused: apply(plugin = "nie-benutzt")`) || strings.Count(out, "\n") != 1 {
		t.Fatalf("code=%d out=%q err=%q", code, out, errs)
	}
}

// Review slice-211 F-3/F-4/F-9: Sollform-Datei nicht zerlegbar, als Symlink oder
// zugleich von files getroffen — jeweils Exit 2, nie still grün.
func TestShapesExactExit2(t *testing.T) {
	base := cfg + "shapes:\n  - files: [\"domain/build.gradle.kts\"]\n    dialect: kotlin\n    mode: exact\n    expect: sollform/b.kts\n"
	cases := map[string]map[string]string{
		"unzerlegbar":  {"sollform/b.kts": "plugins {\n"},
		"selbst geprueft": {".a-check.yml": strings.Replace(base, "sollform/b.kts", "domain/build.gradle.kts", 1)},
	}
	for name, extra := range cases {
		files := map[string]string{".a-check.yml": base, "domain/build.gradle.kts": "a()\n", "sollform/b.kts": "a()\n", "internal/core/c.go": "package core\n"}
		for k, v := range extra {
			files[k] = v
		}
		var out, errb bytes.Buffer
		if code := cli.Run([]string{writeRepo(t, files)}, &out, &errb); code != 2 || out.String() != "" || !strings.Contains(errb.String(), "expect-Datei") {
			t.Errorf("%s: code=%d out=%q err=%q", name, code, out.String(), errb.String())
		}
	}
	// Symlink auf eine Datei AUSSERHALB der Scan-Wurzel wird nicht verfolgt
	outside := filepath.Join(t.TempDir(), "fremd.kts")
	if err := os.WriteFile(outside, []byte("a()\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	dir := writeRepo(t, map[string]string{".a-check.yml": base, "domain/build.gradle.kts": "a()\n", "internal/core/c.go": "package core\n"})
	if err := os.MkdirAll(filepath.Join(dir, "sollform"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(dir, "sollform", "b.kts")); err != nil {
		t.Fatal(err)
	}
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 2 || !strings.Contains(errb.String(), "Symlink") {
		t.Fatalf("Symlink: code=%d err=%q", code, errb.String())
	}
}

// Review slice-211 F-4: das Gerüst zeigt auch exact und unused.
func TestPrintConfigShowsExactAndUnused(t *testing.T) {
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-config"}, &out, &errb); code != 0 {
		t.Fatalf("code=%d", code)
	}
	if s := out.String(); !strings.Contains(s, "mode: exact") || !strings.Contains(s, "expect:") || !strings.Contains(s, "unused: fail") {
		t.Fatalf("Gerüst: %q", s)
	}
}

// Review slice-211 N-1/N-2/N-5: ein Symlink-VERZEICHNIS im expect-Pfad, andere
// Schreibweisen desselben Pfads und ein Hardlink der geprüften Datei — jeweils
// Exit 2, nie fremder Inhalt auf stdout, nie still grün.
func TestShapesExactPathForms(t *testing.T) {
	mk := func(expect string) (string, string) {
		c := cfg + "shapes:\n  - files: [\"d/b.kts\"]\n    dialect: kotlin\n    mode: exact\n    expect: " + expect + "\n"
		dir := writeRepo(t, map[string]string{".a-check.yml": c, "d/b.kts": "evil()\n", "internal/core/c.go": "package core\n"})
		return dir, c
	}
	for _, e := range []string{"d//b.kts", "x/../d/b.kts", "d/./b.kts"} {
		dir, _ := mk(e)
		var out, errb bytes.Buffer
		if code := cli.Run([]string{dir}, &out, &errb); code != 2 || out.String() != "" {
			t.Errorf("expect %q: code=%d out=%q err=%q", e, code, out.String(), errb.String())
		}
	}
	// Hardlink der geprüften Datei als Sollform
	dir, _ := mk("s/b.kts")
	if err := os.MkdirAll(filepath.Join(dir, "s"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Link(filepath.Join(dir, "d", "b.kts"), filepath.Join(dir, "s", "b.kts")); err != nil {
		t.Fatal(err)
	}
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 2 || !strings.Contains(errb.String(), "zugleich die geprüfte Datei") {
		t.Fatalf("Hardlink: code=%d err=%q", code, errb.String())
	}
	// Symlink-Verzeichnis im Pfad, das aus der Wurzel hinauszeigt
	outside := t.TempDir()
	if err := os.WriteFile(filepath.Join(outside, "b.kts"), []byte("a()\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	dir2, _ := mk("link/b.kts")
	if err := os.Symlink(outside, filepath.Join(dir2, "link")); err != nil {
		t.Fatal(err)
	}
	var out2, errb2 bytes.Buffer
	if code := cli.Run([]string{dir2}, &out2, &errb2); code != 2 || out2.String() != "" || !strings.Contains(errb2.String(), "Symlink (link)") {
		t.Fatalf("Symlink-Verzeichnis: code=%d out=%q err=%q", code, out2.String(), errb2.String())
	}
}
