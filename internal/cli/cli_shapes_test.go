package cli_test

import (
	"bytes"
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
`

const goodGradle = `// Fachkern: keine Fremdabhängigkeit
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
// Zeilenumbrüche egal; ein verbotenes Muster nur im Kommentar oder in der
// Zeichenkette einer Anweisung ist kein eigener Befund. Die String-Zeile selbst
// ist nicht gelistet und muss darum gemeldet werden — als ganze Anweisung.
func TestShapesGreenExceptListedNote(t *testing.T) {
	code, out, errs := runShapes(t, goodGradle, nil)
	if code != 1 || strings.Count(out, "shape-unlisted") != 1 || !strings.Contains(out, "domain/build.gradle.kts:13: shape-unlisted: val note=") {
		t.Fatalf("code=%d out=%q err=%q", code, out, errs)
	}
	clean := strings.Replace(goodGradle, `val note = "implementation(\"com.evil:lib:1.0\") steht nur im String"`, "", 1)
	if code, out, errs := runShapes(t, clean, nil); code != 0 || out != "" {
		t.Fatalf("clean: code=%d out=%q err=%q", code, out, errs)
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
	if out1 != out2 || strings.Count(out1, "\n") != 3 {
		t.Fatalf("out1=%q out2=%q", out1, out2)
	}
}
