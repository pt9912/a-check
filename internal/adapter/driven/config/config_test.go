package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const valid = `version: 1
languages:
  go: ["**/*.go"]
layers:
  core: ["core/**"]
  adapters: ["adapters/**"]
edges:
  - {from: adapters, to: core}
adapter_sink: driver-common
tech:
  - {pattern: "net/http", adapter: "adapters/http"}
markers:
  ignore_symbols: ["Queue.h"]
`

func write(t *testing.T, body string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), ".a-check.yml")
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestValidLoad(t *testing.T) { // AC-FA-CONF-001 happy
	m, err := New().Load(write(t, valid))
	if err != nil {
		t.Fatalf("valid config failed: %v", err)
	}
	if len(m.Layers) != 2 || len(m.Edges) != 1 || m.AdapterSink != "driver-common" {
		t.Fatalf("model not decoded as expected: %+v", m)
	}
	if len(m.IgnoreSymbols) != 1 || m.IgnoreSymbols[0] != "Queue.h" {
		t.Fatalf("markers not decoded: %+v", m.IgnoreSymbols)
	}
}

func TestUnknownKeyFailsClosed(t *testing.T) { // AC-FA-CONF-001 negative (strict decode)
	if _, err := New().Load(write(t, valid+"bogus: 1\n")); err == nil {
		t.Fatal("expected error on unknown key (fail-closed)")
	}
}

func TestMissingRequiredBlock(t *testing.T) { // AC-FA-CONF-001 boundary
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: [\"core/**\"]\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error: 'edges' missing")
	}
}

func TestMissingLanguagesBlock(t *testing.T) { // AC-FA-CONF-001 boundary: fehlender languages-Pflichtblock -> Fehler (Exit 2)
	body := "version: 1\nlayers:\n  core: [\"core/**\"]\nedges:\n  - {from: core, to: core}\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error: 'languages' missing (Pflichtblock)")
	}
}

func TestWrongVersion(t *testing.T) {
	if _, err := New().Load(write(t, "version: 2\nlanguages:\n  go: [\"a\"]\nlayers:\n  c: [\"c\"]\nedges:\n  - {from: c, to: c}\n")); err == nil {
		t.Fatal("expected error on unsupported version")
	}
}

func TestLayerObjectFormWithRole(t *testing.T) { // AC-FA-RULE-006 / AC-FA-CONF-001: Objektform mit role
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  geometry: {globs: [\"geometry/**\"], role: adapter}\n  core: [\"core/**\"]\nedges:\n  - {from: geometry, to: core}\n"
	m, err := New().Load(write(t, body))
	if err != nil {
		t.Fatalf("object-form layer failed: %v", err)
	}
	got := ""
	for _, l := range m.Layers {
		if l.Name == "geometry" {
			got = l.Role
		}
	}
	if got != "adapter" {
		t.Fatalf("expected geometry role 'adapter', got %q (%+v)", got, m.Layers)
	}
}

func TestLayerObjectUnknownKeyFailsClosed(t *testing.T) { // SPEC-CONF-001: strict auch im Objekt (yaml.Node-Gotcha)
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: {globs: [\"core/**\"], bogus: 1}\nedges:\n  - {from: core, to: core}\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error on unknown key inside a layer object (fail-closed)")
	}
}

func TestLayerInvalidRoleFailsClosed(t *testing.T) { // AC-FA-RULE-006: role nur domain|port|adapter
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: {globs: [\"core/**\"], role: domainx}\nedges:\n  - {from: core, to: core}\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error on invalid role")
	}
}

func TestLayerScalarFailsClosed(t *testing.T) { // SPEC-CONF-001: Scalar-Layer ist weder Glob-Liste noch Objekt
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: \"core/**\"\nedges:\n  - {from: core, to: core}\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error for scalar layer value")
	}
}

func TestLayerSeqBadElementFailsClosed(t *testing.T) { // Glob-Liste mit Nicht-String-Element
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: [{x: 1}]\nedges:\n  - {from: core, to: core}\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error for non-string glob element")
	}
}

func TestLayerObjectBadGlobsTypeFailsClosed(t *testing.T) { // globs im Objekt muss eine Liste sein
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: {globs: \"core/**\", role: domain}\nedges:\n  - {from: core, to: core}\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error for non-list globs in object form")
	}
}

func TestLayerRoleAppAccepted(t *testing.T) { // AC-FA-RULE-007: role: app wird akzeptiert (Positiv-Test)
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  application: {globs: [\"application/**\"], role: app}\n  core: [\"core/**\"]\nedges:\n  - {from: application, to: core}\n"
	m, err := New().Load(write(t, body))
	if err != nil {
		t.Fatalf("role: app sollte akzeptiert werden, got %v", err)
	}
	got := ""
	for _, l := range m.Layers {
		if l.Name == "application" {
			got = l.Role
		}
	}
	if got != "app" {
		t.Fatalf("expected application role 'app', got %q (%+v)", got, m.Layers)
	}
}

func TestLayerDirectionAccepted(t *testing.T) { // AC-FA-RULE-008: {role, direction} akzeptiert + dekodiert
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  cli: {globs: [\"cli/**\"], role: adapter, direction: driving}\n  api: {globs: [\"api/**\"], role: port, direction: driven}\nedges:\n  - {from: cli, to: api}\n"
	m, err := New().Load(write(t, body))
	if err != nil {
		t.Fatalf("{role, direction} sollte akzeptiert werden, got %v", err)
	}
	got := map[string]string{}
	for _, l := range m.Layers {
		got[l.Name] = l.Direction
	}
	if got["cli"] != "driving" || got["api"] != "driven" {
		t.Fatalf("direction nicht dekodiert: %+v", m.Layers)
	}
}

func TestLayerInvalidDirectionFailsClosed(t *testing.T) { // AC-FA-RULE-008: direction nur driving|driven
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  cli: {globs: [\"cli/**\"], role: adapter, direction: sideways}\nedges:\n  - {from: cli, to: cli}\n"
	if _, err := New().Load(write(t, body)); err == nil {
		t.Fatal("expected error on invalid direction (driving|driven)")
	}
}

func TestLayerDirectionWithoutRoleAccepted(t *testing.T) { // AC-FA-RULE-008: direction ohne role lädt (inert — die Regel braucht role adapter/port)
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  x: {globs: [\"x/**\"], direction: driving}\n  core: [\"core/**\"]\nedges:\n  - {from: x, to: core}\n"
	m, err := New().Load(write(t, body))
	if err != nil {
		t.Fatalf("direction ohne role sollte laden (inert), got %v", err)
	}
	for _, l := range m.Layers {
		if l.Name == "x" && (l.Direction != "driving" || l.Role != "") {
			t.Fatalf("x: erwarte direction=driving, role=\"\" (inert), got %+v", l)
		}
	}
}

// techBody baut eine minimale gültige Config mit einem tech-Eintrag.
func techBody(entry string) string {
	return "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: [\"core/**\"]\n  adapters: [\"adapters/**\"]\nedges:\n  - {from: adapters, to: core}\ntech:\n  - " + entry + "\n"
}

func TestTechMatchRegexValid(t *testing.T) { // AC-FA-CONF-001 / ADR-0015: match: regex mit gültiger Regex lädt
	if _, err := New().Load(write(t, techBody(`{pattern: "Q[A-Za-z]", adapter: "adapters/ui", match: regex}`))); err != nil {
		t.Fatalf("match: regex mit gültiger Regex muss laden, got %v", err)
	}
}

func TestTechMatchSubstringExplicitValid(t *testing.T) { // AC-FA-CONF-001: explizites match: substring lädt
	if _, err := New().Load(write(t, techBody(`{pattern: "net/http", adapter: "adapters/http", match: substring}`))); err != nil {
		t.Fatalf("match: substring muss laden, got %v", err)
	}
}

func TestTechMatchUnknownFailsClosed(t *testing.T) { // AC-FA-CONF-001 negative: unbekannter match-Wert -> Exit 2
	if _, err := New().Load(write(t, techBody(`{pattern: "x", adapter: "adapters/ui", match: glob}`))); err == nil {
		t.Fatal("expected error on unknown match value (fail-closed)")
	}
}

func TestTechMatchInvalidRegexFailsClosed(t *testing.T) { // AC-FA-CONF-001 negative: nicht kompilierbare Regex -> Exit 2
	if _, err := New().Load(write(t, techBody(`{pattern: "Q[A-Za-z", adapter: "adapters/ui", match: regex}`))); err == nil {
		t.Fatal("expected error on uncompilable regex (fail-closed)")
	}
}

func TestTechMatchEmptyRegexFailsClosed(t *testing.T) { // AC-FA-CONF-001 negative: leeres regex-Pattern -> Exit 2 (würde jeden Import treffen)
	if _, err := New().Load(write(t, techBody(`{pattern: "", adapter: "adapters/ui", match: regex}`))); err == nil {
		t.Fatal("expected error on empty regex pattern (fail-closed)")
	}
}

func TestTechUnknownKeyFailsClosed(t *testing.T) { // AC-FA-CONF-001 negative: unbekannter Schlüssel im tech-Eintrag -> Exit 2
	if _, err := New().Load(write(t, techBody(`{pattern: "x", adapter: "adapters/ui", bogus: 1}`))); err == nil {
		t.Fatal("expected error on unknown key in tech entry (strict decode)")
	}
}

// --- slice-023 (AC-FA-RULE-003/AC-FA-CONF-001 0.14.0): adapter-Liste, composition_root, exclude ---

func TestTechAdapterListValid(t *testing.T) { // adapter als Pfad-LISTE dekodiert (Reihenfolge erhalten)
	m, err := New().Load(write(t, techBody(`{pattern: yaml, adapter: [adapters/config, adapters/report]}`)))
	if err != nil {
		t.Fatalf("adapter-Liste muss laden, got %v", err)
	}
	a := m.Techs[0].Adapters
	if len(a) != 2 || a[0] != "adapters/config" || a[1] != "adapters/report" {
		t.Fatalf("adapter-Liste nicht korrekt dekodiert: %v", a)
	}
}

func TestTechAdapterScalarBackCompat(t *testing.T) { // Skalar == einelementige Liste (rückwärtskompatibel)
	m, err := New().Load(write(t, techBody(`{pattern: "net/http", adapter: "adapters/http"}`)))
	if err != nil || len(m.Techs[0].Adapters) != 1 || m.Techs[0].Adapters[0] != "adapters/http" {
		t.Fatalf("Skalar-adapter muss einelementige Liste ergeben, got %v / %v", err, m.Techs)
	}
}

func TestTechAdapterEmptyListFailsClosed(t *testing.T) { // leere Liste -> Exit 2
	if _, err := New().Load(write(t, techBody(`{pattern: yaml, adapter: []}`))); err == nil || !strings.Contains(err.Error(), "leere adapter-Liste") {
		t.Fatalf("leere adapter-Liste muss brechen, got %v", err)
	}
}

func TestTechAdapterEmptyListEntryFailsClosed(t *testing.T) { // leerer Listen-Eintrag -> Exit 2 (strikte neue Form)
	if _, err := New().Load(write(t, techBody(`{pattern: yaml, adapter: [adapters/config, ""]}`))); err == nil || !strings.Contains(err.Error(), "leerer adapter-Listen-Eintrag") {
		t.Fatalf("leerer Listen-Eintrag muss brechen, got %v", err)
	}
}

func TestTechAdapterAbsentFailsClosed(t *testing.T) { // Review-R1 B1: fehlender adapter -> Exit 2 (Alt-Zustand war stiller Never-Leak)
	if _, err := New().Load(write(t, techBody(`{pattern: yaml}`))); err == nil || !strings.Contains(err.Error(), "adapter fehlt") {
		t.Fatalf("fehlender adapter muss brechen (kein stiller Never-Leak-Eintrag), got %v", err)
	}
}

func TestTechAdapterEmptyScalarFailsClosed(t *testing.T) { // Review-R1 B1/M2: leerer Skalar und null -> Exit 2
	if _, err := New().Load(write(t, techBody(`{pattern: yaml, adapter: ""}`))); err == nil || !strings.Contains(err.Error(), "leerer adapter") {
		t.Fatalf("leerer adapter-Skalar muss brechen, got %v", err)
	}
	if _, err := New().Load(write(t, techBody(`{pattern: yaml, adapter: null}`))); err == nil || !strings.Contains(err.Error(), "leerer adapter") {
		t.Fatalf("adapter: null muss brechen, got %v", err)
	}
}

func TestTechAdapterAsMappingFailsClosed(t *testing.T) { // Review-R1: Mapping ist weder Pfad noch Pfad-Liste -> Exit 2
	if _, err := New().Load(write(t, techBody(`{pattern: yaml, adapter: {x: y}}`))); err == nil || !strings.Contains(err.Error(), "Pfad oder Pfad-Liste") {
		t.Fatalf("adapter als Mapping muss brechen, got %v", err)
	}
}

func TestTechAdapterAliasResolved(t *testing.T) { // Review-R1 M1: YAML-Alias auf adapter wird dereferenziert, nicht abgelehnt
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: [\"core/**\"]\n  adapters: [\"adapters/**\"]\nedges:\n  - {from: adapters, to: core}\nmarkers:\n  ignore_symbols: &adp [\"adapters/config\"]\ntech:\n  - {pattern: yaml, adapter: *adp}\n"
	m, err := New().Load(write(t, body))
	if err != nil {
		t.Fatalf("Alias-adapter muss laden (dereferenziert), got %v", err)
	}
	if a := m.Techs[0].Adapters; len(a) != 1 || a[0] != "adapters/config" {
		t.Fatalf("Alias muss zur Liste dereferenzieren, got %v", a)
	}
}

func TestTechCompositionRootDecode(t *testing.T) { // forbid -> Flag gesetzt; allow/weggelassen -> Default false
	m, err := New().Load(write(t, techBody(`{pattern: "net/http", adapter: "adapters/http", composition_root: forbid}`)))
	if err != nil || !m.Techs[0].ForbidCompositionRoot {
		t.Fatalf("composition_root: forbid muss das Flag setzen, got %v / %+v", err, m.Techs)
	}
	m2, err := New().Load(write(t, techBody(`{pattern: "net/http", adapter: "adapters/http", composition_root: allow}`)))
	if err != nil || m2.Techs[0].ForbidCompositionRoot {
		t.Fatalf("composition_root: allow ist der Default (Flag false), got %v / %+v", err, m2.Techs)
	}
	m3, err := New().Load(write(t, techBody(`{pattern: "net/http", adapter: "adapters/http"}`)))
	if err != nil || m3.Techs[0].ForbidCompositionRoot {
		t.Fatalf("ohne composition_root-Feld gilt allow (Flag false), got %v / %+v", err, m3.Techs)
	}
}

func TestTechCompositionRootInvalidFailsClosed(t *testing.T) { // Wert außerhalb {allow, forbid} -> Exit 2
	if _, err := New().Load(write(t, techBody(`{pattern: "net/http", adapter: "adapters/http", composition_root: maybe}`))); err == nil || !strings.Contains(err.Error(), "(allow|forbid)") {
		t.Fatalf("ungültiger composition_root-Wert muss brechen (Enum in der Meldung), got %v", err)
	}
}

func TestExcludeDecode(t *testing.T) { // AC-FA-CONF-001 / ADR-0018: exclude-Globs werden dekodiert
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: [\"core/**\"]\nedges:\n  - {from: core, to: core}\nexclude:\n  - \"**/*_test.go\"\n  - \"**/node_modules/**\"\n"
	m, err := New().Load(write(t, body))
	if err != nil || len(m.Exclude) != 2 || m.Exclude[0] != "**/*_test.go" {
		t.Fatalf("exclude muss dekodiert werden, got %v / %v", err, m.Exclude)
	}
}

func TestExcludeEmptyGlobFailsClosed(t *testing.T) { // ADR-0018: leerer Glob -> Exit 2 (stiller Match-Nichts-Eintrag)
	body := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  core: [\"core/**\"]\nedges:\n  - {from: core, to: core}\nexclude:\n  - \"\"\n"
	if _, err := New().Load(write(t, body)); err == nil || !strings.Contains(err.Error(), "leerer Glob") {
		t.Fatalf("leerer exclude-Glob muss brechen, got %v", err)
	}
}

// resBody baut eine minimale gültige Config mit einem resolution-Eintrag. Es
// deklariert bewusst mehrere languages (config.Load prüft die Backend-Menge
// nicht — das tut extract), damit der resolution-Key-gegen-languages-Check
// (ADR-0016) nur bei absichtlich nicht-deklarierten Keys greift.
func resBody(entry string) string {
	return "version: 1\nlanguages:\n  go: [\"**/*.go\"]\n  cpp: [\"**/*.h\"]\n  kotlin: [\"**/*.kt\"]\n  typescript: [\"**/*.ts\"]\n  csharp: [\"**/*.cs\"]\nlayers:\n  core: [\"core/**\"]\nedges:\n  - {from: core, to: core}\nresolution:\n  " + entry + "\n"
}

func TestResolutionFixedRootValid(t *testing.T) { // AC-FA-CONF-001 / ADR-0016: fixed-root lädt + wird dekodiert
	m, err := New().Load(write(t, resBody(`kotlin: {mode: fixed-root, roots: ["src"], package_base: "com.x"}`)))
	if err != nil {
		t.Fatalf("fixed-root muss laden, got %v", err)
	}
	r := m.Resolution["kotlin"]
	if r.Mode != "fixed-root" || r.PackageBase != "com.x" || len(r.Roots) != 1 || r.Roots[0] != "src" {
		t.Fatalf("resolution nicht dekodiert: %+v", r)
	}
}

func TestResolutionReservedModeFailsClosed(t *testing.T) { // ADR-0016: reservierter mode (nur noch namespace, slice-022) -> Exit 2, Meldung nennt "reserviert"
	if _, err := New().Load(write(t, resBody(`csharp: {mode: namespace}`))); err == nil || !strings.Contains(err.Error(), "reserviert") {
		t.Fatalf("mode 'namespace' muss als reserviert brechen, got %v", err)
	}
}

func TestResolutionRelativeValid(t *testing.T) { // AC-FA-CONF-001 / ADR-0017: relative lädt + wird dekodiert
	m, err := New().Load(write(t, resBody(`typescript: {mode: relative}`)))
	if err != nil {
		t.Fatalf("mode relative muss laden (seit ADR-0017 gültig), got %v", err)
	}
	if r := m.Resolution["typescript"]; r.Mode != "relative" || len(r.Roots) != 0 || r.PackageBase != "" {
		t.Fatalf("relative nicht korrekt dekodiert: %+v", r)
	}
}

func TestResolutionRelativeRejectsRootsAndPackageBase(t *testing.T) { // ADR-0017: relative duldet kein roots/package_base (fail-closed)
	if _, err := New().Load(write(t, resBody(`typescript: {mode: relative, roots: ["src"]}`))); err == nil {
		t.Fatal("relative mit roots muss brechen (roots würden still ignoriert)")
	}
	if _, err := New().Load(write(t, resBody(`typescript: {mode: relative, package_base: "x"}`))); err == nil {
		t.Fatal("relative mit package_base muss brechen (würde still ignoriert)")
	}
}

func TestResolutionUndeclaredLanguageFailsClosed(t *testing.T) { // ADR-0016: resolution-Key ohne languages-Deklaration -> Exit 2 (kein stiller No-Op)
	if _, err := New().Load(write(t, resBody(`rust: {mode: path}`))); err == nil {
		t.Fatal("resolution für nicht deklarierte Sprache (rust) muss brechen — sonst Tippfehler = false-green")
	}
}

func TestResolutionPathWithRootsFailsClosed(t *testing.T) { // ADR-0016 (LOW): mode path duldet kein roots/package_base
	if _, err := New().Load(write(t, resBody(`go: {mode: path, roots: ["src"]}`))); err == nil {
		t.Fatal("mode: path mit roots muss brechen (roots würden still ignoriert)")
	}
}

func TestResolutionDegenerateFixedRootFailsClosed(t *testing.T) { // ADR-0016 (LOW): fixed-root ohne roots UND ohne package_base = No-Op -> Exit 2
	if _, err := New().Load(write(t, resBody(`go: {mode: fixed-root}`))); err == nil {
		t.Fatal("fixed-root ohne roots/package_base ist ein stiller No-Op und muss brechen")
	}
}

func TestResolutionEmptyRootFailsClosed(t *testing.T) { // ADR-0016 (LOW): leerer root -> Exit 2
	if _, err := New().Load(write(t, resBody(`cpp: {mode: fixed-root, roots: [""]}`))); err == nil {
		t.Fatal("leerer root muss brechen")
	}
}

func TestResolutionUnknownKeyFailsClosed(t *testing.T) { // ADR-0016 (F2): unbekannter Schlüssel im resolution-Eintrag -> Exit 2 (strict-decode)
	if _, err := New().Load(write(t, resBody(`go: {mode: path, bogus: 1}`))); err == nil {
		t.Fatal("unbekannter Schlüssel im resolution-Eintrag muss brechen (KnownFields)")
	}
}

func TestResolutionTrailingSlashRootNormalized(t *testing.T) { // ADR-0016 (LOW): "src/" wird zu "src" normalisiert
	m, err := New().Load(write(t, resBody(`cpp: {mode: fixed-root, roots: ["src/"]}`)))
	if err != nil || len(m.Resolution["cpp"].Roots) != 1 || m.Resolution["cpp"].Roots[0] != "src" {
		t.Fatalf("Trailing-Slash-Root muss zu 'src' normalisiert werden, got %v / %+v", err, m.Resolution)
	}
}

func TestResolutionUnknownModeFailsClosed(t *testing.T) { // ADR-0016: unbekannter mode -> Exit 2; Enum-Meldung nennt alle gültigen Modi (slice-022)
	_, err := New().Load(write(t, resBody(`go: {mode: bogus}`)))
	if err == nil {
		t.Fatal("unbekannter mode muss brechen (fail-closed)")
	}
	if !strings.Contains(err.Error(), "(path|fixed-root|relative)") {
		t.Fatalf("Enum-Meldung muss (path|fixed-root|relative) nennen, got %v", err)
	}
}

func TestResolutionPathEqualsOmitted(t *testing.T) { // ADR-0016: mode: path == weggelassen (beide -> mode path)
	m, err := New().Load(write(t, resBody(`go: {mode: path}`)))
	if err != nil || m.Resolution["go"].Mode != "path" {
		t.Fatalf("mode: path muss laden (mode=path), got %v / %+v", err, m.Resolution)
	}
	m2, err := New().Load(write(t, resBody(`go: {}`)))
	if err != nil || m2.Resolution["go"].Mode != "path" {
		t.Fatalf("resolution go: {} muss als mode=path laden, got %v / %+v", err, m2.Resolution)
	}
}

// kmpFlat is Config X from the belief-agent report (ADR-0020): two roots sharing
// package_base with flat source-set globs — phantom-capable, must fail closed.
const kmpFlat = `version: 1
languages:
  kotlin: ["**/*.kt"]
layers:
  core: ["src/commonMain/**"]
  adapters: ["src/jvmMain/**"]
edges:
  - {from: adapters, to: core}
resolution:
  kotlin: {mode: fixed-root, roots: ["src/commonMain/kotlin/myapp", "src/jvmMain/kotlin/myapp"], package_base: "myapp"}
`

func TestResolutionMultiRootPhantomFailsClosed(t *testing.T) { // AC-FA-CONF-001 / ADR-0020: 2 roots je in 2 Schichten -> Exit 2 (sonst stilles Falsch-Negativ)
	_, err := New().Load(write(t, kmpFlat))
	if err == nil {
		t.Fatal("Mehr-Wurzel-Phantom (2 roots in 2 Schichten) muss fail-closed brechen — sonst still falsch-grün")
	}
	if !strings.Contains(err.Error(), "Mehr-Wurzel") || !strings.Contains(err.Error(), "commonMain") || !strings.Contains(err.Error(), "jvmMain") {
		t.Fatalf("Meldung muss die beiden roots nennen, got %v", err)
	}
}

// kmpDeep is Config Y: same roots, but package-specific globs deeper than the
// roots — no root is contained in a layer, so it loads (the documented recipe).
const kmpDeep = `version: 1
languages:
  kotlin: ["**/*.kt"]
layers:
  core: ["src/commonMain/kotlin/myapp/core/**"]
  adapters: ["src/jvmMain/kotlin/myapp/adapters/**"]
edges:
  - {from: adapters, to: core}
resolution:
  kotlin: {mode: fixed-root, roots: ["src/commonMain/kotlin/myapp", "src/jvmMain/kotlin/myapp"], package_base: "myapp"}
`

func TestResolutionMultiRootDeepGlobsValid(t *testing.T) { // ADR-0020 Boundary: paket-tiefe Globs (Rezept) -> kein Root in einer Schicht -> lädt
	if _, err := New().Load(write(t, kmpDeep)); err != nil {
		t.Fatalf("paket-tiefe Globs (das dokumentierte Rezept) müssen laden, got %v", err)
	}
}

// twoRootsSameLayer: two roots, but both fall inside the SAME layer — no
// cross-layer phantom, so it loads (the guard requires DIFFERENT layers).
const twoRootsSameLayer = `version: 1
languages:
  go: ["**/*.go"]
layers:
  core: ["src/**"]
edges:
  - {from: core, to: core}
resolution:
  go: {mode: fixed-root, roots: ["src/a", "src/b"]}
`

func TestResolutionMultiRootSameLayerValid(t *testing.T) { // ADR-0020: 2 roots in DERSELBEN Schicht -> kein Konflikt -> lädt
	if _, err := New().Load(write(t, twoRootsSameLayer)); err != nil {
		t.Fatalf("zwei roots in derselben Schicht sind kein Phantom-Konflikt und müssen laden, got %v", err)
	}
}
