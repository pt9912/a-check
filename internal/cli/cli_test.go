package cli_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pt9912/a-check/internal/cli"
)

const cfg = `version: 1
languages:
  go: ["**/*.go"]
layers:
  core: ["internal/core/**"]
  adapters: ["internal/adapters/**"]
edges:
  - {from: adapters, to: core}
`

func writeRepo(t *testing.T, files map[string]string) string {
	t.Helper()
	dir := t.TempDir()
	for rel, body := range files {
		p := filepath.Join(dir, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func TestPrintConfig(t *testing.T) { // SPEC-DIST-001: read-only Gerüst
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-config"}, &out, &errb); code != 0 {
		t.Fatalf("exit %d", code)
	}
	if !strings.Contains(out.String(), "version: 1") {
		t.Fatalf("print-config: %q", out.String())
	}
}

func TestPrintMk(t *testing.T) { // SPEC-DIST-001: includebares Fragment
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-mk"}, &out, &errb); code != 0 {
		t.Fatalf("exit %d", code)
	}
	o := out.String()
	if !strings.Contains(o, "A_CHECK_IMAGE") || !strings.Contains(o, "a-check:") {
		t.Fatalf("print-mk: %q", o)
	}
}

func TestUnknownFlag(t *testing.T) { // SPEC-CLI-001: Exit 2
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--bogus"}, &out, &errb); code != 2 {
		t.Fatalf("expected 2, got %d", code)
	}
}

func TestScanClean(t *testing.T) { // SPEC-CLI-001 happy: Exit 0
	dir := writeRepo(t, map[string]string{
		".a-check.yml":           cfg,
		"internal/core/svc.go":   "package core\nimport \"fmt\"\n\nvar _ = fmt.Sprint\n",
		"internal/adapters/a.go": "package adapters\nimport \"x/internal/core\"\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 0 {
		t.Fatalf("expected 0, got %d (out=%q)", code, out.String())
	}
}

func TestScanViolation(t *testing.T) { // SPEC-CLI-001 negative: Exit 1
	dir := writeRepo(t, map[string]string{
		".a-check.yml":          cfg,
		"internal/core/bad.go":  "package core\nimport \"x/internal/adapters/http\"\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("expected 1, got %d", code)
	}
	if !strings.Contains(out.String(), "core-impurity") {
		t.Fatalf("expected core-impurity: %q", out.String())
	}
}

func TestMissingConfig(t *testing.T) { // SPEC-CLI-001: Exit 2 bei fehlender Config
	var out, errb bytes.Buffer
	if code := cli.Run([]string{t.TempDir()}, &out, &errb); code != 2 {
		t.Fatalf("expected 2, got %d", code)
	}
}

// cfgRegexTech: C++-Repo mit Qt-Muster als RE2-Regex auf dem ui-Adapter (ADR-0015).
const cfgRegexTech = `version: 1
languages:
  cpp: ["**/*.cpp", "**/*.h"]
layers:
  ui:  {globs: ["adapters/ui/**"], role: adapter}
  geo: {globs: ["adapters/geometry/**"], role: adapter}
edges:
  - {from: geo, to: ui}
tech:
  - {pattern: "Q[A-Za-z]", adapter: "adapters/ui", match: regex}
`

func TestTechRegexLeakExit1(t *testing.T) { // AC-FA-RULE-003 / ADR-0015: regex-tech-leak erreicht Exit-Code 1
	dir := writeRepo(t, map[string]string{
		".a-check.yml":            cfgRegexTech,
		"adapters/geometry/g.cpp": "#include <QWidget>\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("expected 1, got %d (out=%q)", code, out.String())
	}
	if !strings.Contains(out.String(), "tech-leak") {
		t.Fatalf("expected tech-leak in output: %q", out.String())
	}
}

func TestTechRegexInvalidMatchExit2(t *testing.T) { // AC-FA-CONF-001 / ADR-0015: ungültiges match erreicht Exit-Code 2
	badCfg := strings.Replace(cfgRegexTech, "match: regex", "match: glob", 1)
	dir := writeRepo(t, map[string]string{
		".a-check.yml":      badCfg,
		"adapters/ui/w.cpp": "#include <QWidget>\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 2 {
		t.Fatalf("expected 2, got %d", code)
	}
}

func TestTechRegexIgnoreSymbols(t *testing.T) { // AC-QA-02 / ADR-0015: markers.ignore_symbols unterdrückt den Q[A-Za-z]/Queue.h-False-Positive
	// Q[A-Za-z] trifft "Queue.h" (der dokumentierte FP); der Marker unterdrückt ihn.
	withMarker := map[string]string{
		".a-check.yml":            cfgRegexTech + "markers:\n  ignore_symbols: [\"Queue.h\"]\n",
		"adapters/geometry/g.cpp": "#include \"Queue.h\"\n",
	}
	var out, errb bytes.Buffer
	if code := cli.Run([]string{writeRepo(t, withMarker)}, &out, &errb); code != 0 {
		t.Fatalf("ignore_symbols muss den Queue.h-FP unterdrücken (Exit 0), got %d (out=%q)", code, out.String())
	}
	// Ohne den Marker schlägt derselbe FP als tech-leak an (Exit 1) — die Heuristik-Grenze ist ausgewiesen, nicht verschwiegen.
	withoutMarker := map[string]string{
		".a-check.yml":            cfgRegexTech,
		"adapters/geometry/g.cpp": "#include \"Queue.h\"\n",
	}
	out.Reset()
	errb.Reset()
	if code := cli.Run([]string{writeRepo(t, withoutMarker)}, &out, &errb); code != 1 {
		t.Fatalf("ohne ignore_symbols erwarte tech-leak (Exit 1), got %d (out=%q)", code, out.String())
	}
}

func TestUnknownLanguageExit2(t *testing.T) { // AC-FA-CONF-001 / slice-017: unbekannte Sprache -> Exit 2 statt still falsch-grün
	cfg := "version: 1\nlanguages:\n  ruby: [\"**/*.rb\"]\nlayers:\n  core: [\"core/**\"]\nedges:\n  - {from: core, to: core}\n"
	dir := writeRepo(t, map[string]string{".a-check.yml": cfg, "core/x.rb": "require 'json'\n"})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 2 {
		t.Fatalf("unbekannte Sprache muss Exit 2 liefern, got %d", code)
	}
	if !strings.Contains(errb.String(), "unbekannte Sprache") || !strings.Contains(errb.String(), "ruby") {
		t.Fatalf("Meldung soll die Sprache nennen: %q", errb.String())
	}
	if out.Len() != 0 {
		t.Fatalf("Config-Fehler gehört auf stderr, stdout muss leer sein: %q", out.String())
	}
}

func TestPythonFixedRootResolution(t *testing.T) { // AC-FA-EXTRACT-001 (Python) + AC-FA-CONF-001 Happy-Auflösung: Backend + fixed-root-Rezept (slice-020 §3.3) greifen zusammen
	pyCfg := `version: 1
languages:
  python: ["**/*.py"]
layers:
  core:     ["src/myapp/domain/**"]
  adapters: ["src/myapp/adapters/**"]
edges:
  - {from: adapters, to: core}
resolution:
  python: {mode: fixed-root, roots: ["src/myapp"], package_base: "myapp"}
`
	// Zeile 1: /*-Bytefolge im String (Review-R1: darf Folge-Imports nicht fressen);
	// Zeile 3: Mehrsegment-Rest `adapters.db` pinnt die .->/-Konvertierung (Review-R1).
	dir := writeRepo(t, map[string]string{
		".a-check.yml":              pyCfg,
		"src/myapp/domain/model.py": "GLOB = \"**/*.py\"\nfrom myapp.adapters import db\nimport myapp.adapters.db\n",
		"src/myapp/adapters/db.py":  "import json\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("Python-Domäne importiert Adapter-Modul: erwarte Exit 1, got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
	if got := strings.Count(out.String(), "core-impurity"); got != 2 {
		t.Fatalf("erwarte 2 core-impurity-Befunde (from-Form + Mehrsegment-dotted), got %d: %q", got, out.String())
	}
	if !strings.Contains(out.String(), "src/myapp/domain/model.py:2") || !strings.Contains(out.String(), "src/myapp/domain/model.py:3") {
		t.Fatalf("erwarte Befunde auf Zeile 2 und 3 der Domänen-Datei: %q", out.String())
	}
}

func TestCsharpFixedRootResolution(t *testing.T) { // AC-FA-EXTRACT-001 (C#) + AC-FA-CONF-001 Happy-Auflösung: using-Backend + fixed-root-Rezept (slice-021 §3.3)
	csCfg := `version: 1
languages:
  csharp: ["**/*.cs"]
layers:
  core:     ["src/MyApp/Domain/**"]
  adapters: ["src/MyApp/Adapters/**"]
edges:
  - {from: adapters, to: core}
resolution:
  csharp: {mode: fixed-root, roots: ["src/MyApp"], package_base: "MyApp"}
`
	// Zeile 1: Direktive mit Mehrsegment-Rest `Adapters.Db` (pinnt .->/-Konvertierung);
	// Zeile 2: Alias-Form end-to-end; Zeile 3: using-DECLARATION mit qualifiziertem
	// Adapter-Typ — ohne das Pflicht-; ergäbe sie einen 3. Befund (Review-R1: pinnt
	// den ;-Kern-Mutanten auch end-to-end, ein `var`-Fehlsymbol löste auf keine Schicht).
	dir := writeRepo(t, map[string]string{
		".a-check.yml":                 csCfg,
		"src/MyApp/Domain/Model.cs":    "using MyApp.Adapters.Db;\nusing Db2 = MyApp.Adapters.Db2;\nusing MyApp.Adapters.Db t = Get();\n",
		"src/MyApp/Adapters/Db/Db.cs":  "using System.Text;\n",
		"src/MyApp/Adapters/Db2/D2.cs": "using System;\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("C#-Domäne nutzt Adapter-Namespaces: erwarte Exit 1, got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
	if got := strings.Count(out.String(), "core-impurity"); got != 2 {
		t.Fatalf("erwarte genau 2 core-impurity-Befunde (Direktive + Alias-Ziel, using-Statement keiner), got %d: %q", got, out.String())
	}
	if !strings.Contains(out.String(), "src/MyApp/Domain/Model.cs:1") || !strings.Contains(out.String(), "src/MyApp/Domain/Model.cs:2") {
		t.Fatalf("erwarte Befunde auf Zeile 1 und 2 der Domänen-Datei: %q", out.String())
	}
}

func TestMonoRepoMixedUnsupportedExit2(t *testing.T) { // slice-017: Mono-Repo go+ruby(unsupported) -> Exit 2, go rettet nicht
	// (typescript ist seit slice-022 ein Backend und taugt nicht mehr als Fixture.)
	cfg := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\n  ruby: [\"**/*.rb\"]\nlayers:\n  core: [\"core/**\"]\nedges:\n  - {from: core, to: core}\n"
	dir := writeRepo(t, map[string]string{".a-check.yml": cfg, "core/x.go": "package core\n"})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 2 {
		t.Fatalf("gemischte Sprachen mit unsupported -> Exit 2, got %d", code)
	}
	if !strings.Contains(errb.String(), "ruby") {
		t.Fatalf("Meldung soll die unsupported Sprache nennen: %q", errb.String())
	}
}

func TestTypescriptRelativeResolution(t *testing.T) { // AC-FA-EXTRACT-001 (TS) + AC-FA-CONF-001 Happy (relative): slice-022 §4.1/§4.2
	tsCfg := `version: 1
languages:
  typescript: ["**/*.ts"]
layers:
  core:     ["core/**"]
  ports:    ["ports/**"]
  adapters: ["adapters/**"]
edges:
  - {from: adapters, to: ports}
  - {from: ports,    to: core}
resolution:
  typescript: {mode: relative}
`
	// Zeile 1: relativer Import -> adapters -> core-impurity. Zeile 2: Bare-Import
	// '@actions/adapters' — bei Roh-Durchreichung segment-matchte 'adapters' (Geister-
	// Befund); muss leer bleiben (ADR-0017). Zeile 3: Ausdrucks-Zeile MIT nacktem
	// `from '…'` — ihr Fehlsymbol '../adapters/db2' löste auf; nur die Mittelteil-
	// Klasse (kein '=') blockiert sie, das pinnt die Beschränkung end-to-end
	// (Review-R1 T-1; ein knex.from(...) bewachte nur den from(-Anker). Zeilen 4-6:
	// Prettier-umbrochener Import — die Schlusszeile } from '../adapters/db3' muss
	// den Befund liefern (Entscheid G).
	dir := writeRepo(t, map[string]string{
		".a-check.yml": tsCfg,
		"core/service.ts": "import { Db } from '../adapters/db';\n" +
			"import * as a from '@actions/adapters';\n" +
			"export const q = from '../adapters/db2';\n" +
			"import {\n  A,\n} from \"../adapters/db3\"\n",
		"adapters/db.ts": "export const db = 1;\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("TS-Domäne importiert Adapter relativ: erwarte Exit 1, got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
	if got := strings.Count(out.String(), "core-impurity"); got != 2 {
		t.Fatalf("erwarte genau 2 core-impurity (relativ + Fortsetzungszeile; Bare-Import und Ausdrucks-Zeile keinen), got %d: %q", got, out.String())
	}
	if !strings.Contains(out.String(), "core/service.ts:1") || !strings.Contains(out.String(), "core/service.ts:6") {
		t.Fatalf("erwarte Befunde auf Zeile 1 und 6 der Domänen-Datei: %q", out.String())
	}
}

func TestMonoRepoGoTypescriptRelative(t *testing.T) { // ADR-0017: Mono-Repo — Go path-Default + TS relative, je eigener Modus (CLI end-to-end)
	cfg := `version: 1
languages:
  go:         ["**/*.go"]
  typescript: ["**/*.ts"]
layers:
  core:     ["core/**"]
  adapters: ["adapters/**"]
edges:
  - {from: adapters, to: core}
resolution:
  typescript: {mode: relative}
`
	dir := writeRepo(t, map[string]string{
		".a-check.yml": cfg,
		"core/x.go":    "package core\nimport \"myrepo/adapters/db\"\n",
		"core/y.ts":    "import { D } from '../adapters/db';\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("Mono-Repo Go+TS: erwarte Exit 1, got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
	if got := strings.Count(out.String(), "core-impurity"); got != 2 {
		t.Fatalf("beide Sprachen müssen je über IHREN Modus auflösen -> 2× core-impurity, got %d: %q", got, out.String())
	}
}

func TestFlatAdapterRootSubunit(t *testing.T) { // AC-FA-RULE-002 0.15.0 / ADR-0019: flacher Adapter (b-cad-Form) — x.cpp->x.h kein lateral, Cross-Layer bleibt
	cfg := `version: 1
languages:
  cpp: ["**/*.cpp", "**/*.h"]
layers:
  io: {globs: ["adapters/io/**"], role: adapter}
  ui: {globs: ["adapters/ui/**"], role: adapter}
edges:
  - {from: io, to: io}
`
	dir := writeRepo(t, map[string]string{
		".a-check.yml": cfg,
		// Zeile 1: eigener Header (Root<->Root) -> KEIN Befund (ADR-0019);
		// Zeile 2: fremder Adapter -> lateral-adapter (kategorisch, unveraendert).
		"adapters/io/dxf_reader.cpp": "#include \"adapters/io/dxf_reader.h\"\n#include \"adapters/ui/window.h\"\n",
		"adapters/io/dxf_reader.h":   "#pragma once\n",
		"adapters/ui/window.h":       "#pragma once\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("erwarte Exit 1 (genau der Cross-Layer-Befund), got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
	if got := strings.Count(out.String(), "lateral-adapter"); got != 1 {
		t.Fatalf("erwarte genau 1 lateral-adapter (Cross-Layer; Root<->Root keiner), got %d: %q", got, out.String())
	}
	if !strings.Contains(out.String(), "adapters/ui/window.h") || strings.Contains(out.String(), "dxf_reader.h") {
		t.Fatalf("nur der ui-Import darf melden, nie der eigene Header: %q", out.String())
	}
}

func TestGoPackageRootSubunitE2E(t *testing.T) { // ADR-0019 Entscheid D (Review-R1 T-1): Go-Paket-Blatt end-to-end — Eigen-Paket ok, Fremd-Paket lateral
	cfg := `version: 1
languages:
  go: ["**/*.go"]
layers:
  adapters: {globs: ["internal/adapter/**"], role: adapter}
edges:
  - {from: adapters, to: adapters}
`
	dir := writeRepo(t, map[string]string{
		".a-check.yml": cfg,
		// Externes Testpaket importiert sein EIGENES Paket-Verzeichnis (der Dogfooding-Fund).
		"internal/adapter/report/report.go":      "package report\n",
		"internal/adapter/report/report_test.go": "package report_test\nimport \"myrepo/internal/adapter/report\"\n",
		// Fremdes Paket-Verzeichnis derselben Schicht -> lateral (Blatt-Klassifikation blendet nicht).
		"internal/adapter/config/config.go": "package config\nimport \"myrepo/internal/adapter/report\"\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("erwarte Exit 1 (nur der Fremd-Paket-Import), got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
	if got := strings.Count(out.String(), "lateral-adapter"); got != 1 {
		t.Fatalf("erwarte genau 1 lateral-adapter (config->report; report_test->report keiner), got %d: %q", got, out.String())
	}
	if !strings.Contains(out.String(), "config/config.go") || strings.Contains(out.String(), "report_test.go") {
		t.Fatalf("nur der Fremd-Paket-Import darf melden, nie der Eigen-Paket-Import: %q", out.String())
	}
}

func TestDcheckPilotDeltas(t *testing.T) { // slice-023 end-to-end (AC-FA-RULE-003/AC-FA-CONF-001 0.14.0): adapter-Liste + composition_root: forbid + exclude
	cfg := `version: 1
languages:
  go: ["**/*.go"]
layers:
  core:    {globs: ["core/**"], role: domain}
  config:  {globs: ["adapters/config/**"], role: adapter}
  report:  {globs: ["adapters/report/**"], role: adapter}
  httpad:  {globs: ["adapters/http/**"], role: adapter}
edges:
  - {from: core, to: core}
composition_root: ["cmd/**"]
tech:
  - {pattern: yaml, adapter: [adapters/config, adapters/report]}
  - {pattern: "net/http", adapter: adapters/http, composition_root: forbid}
exclude:
  - "**/*_test.go"
`
	dir := writeRepo(t, map[string]string{
		".a-check.yml": cfg,
		// yaml in beiden gelisteten Adaptern erlaubt:
		"adapters/config/c.go": "package config\nimport \"gopkg.in/yaml.v3\"\n",
		"adapters/report/r.go": "package report\nimport \"gopkg.in/yaml.v3\"\n",
		// yaml außerhalb ALLER gelisteten -> tech-leak (Meldung nennt beide):
		"adapters/http/h.go": "package http\nimport \"gopkg.in/yaml.v3\"\n",
		// composition_root: net/http (forbid) -> tech-leak trotz Verdrahtungspunkt;
		// yaml (allow, Default) und der Schicht-Querimport bleiben ausgenommen:
		"cmd/main.go": "package main\nimport \"net/http\"\nimport \"gopkg.in/yaml.v3\"\nimport \"myrepo/core/svc\"\n",
		// exclude: die Test-Datei enthielte einen weiteren yaml-Leak — darf nie melden:
		"adapters/http/h_test.go": "package http\nimport \"gopkg.in/yaml.v3\"\n",
		"core/svc.go":             "package core\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("erwarte Exit 1 (zwei tech-leaks), got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
	if got := strings.Count(out.String(), "tech-leak"); got != 2 {
		t.Fatalf("erwarte genau 2 tech-leak (h.go + cmd/main.go), got %d: %q", got, out.String())
	}
	if !strings.Contains(out.String(), "adapters/config|adapters/report") {
		t.Fatalf("Mehr-Adapter-Meldung muss beide gelisteten Adapter nennen: %q", out.String())
	}
	if !strings.Contains(out.String(), "(composition_root: forbid)") {
		t.Fatalf("Composition-Root-Verbots-Meldung muss den forbid-Grund nennen: %q", out.String())
	}
	if strings.Contains(out.String(), "h_test.go") {
		t.Fatalf("exclude-Datei darf nie in Befunden auftauchen: %q", out.String())
	}
}

func TestMonoRepoMultiSupportedRuns(t *testing.T) { // slice-017: Mono-Repo mit nur unterstützten Sprachen (go+cpp) läuft
	cfg := "version: 1\nlanguages:\n  go: [\"**/*.go\"]\n  cpp: [\"**/*.h\", \"**/*.cpp\"]\nlayers:\n  core: [\"core/**\"]\nedges:\n  - {from: core, to: core}\n"
	dir := writeRepo(t, map[string]string{".a-check.yml": cfg, "core/x.go": "package core\n"})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 0 {
		t.Fatalf("go+cpp (beide unterstützt) muss laufen (Exit 0), got %d (err=%q)", code, errb.String())
	}
}

// kmpCfg is a disjoint KMP/Gradle multi-module config: two modules share
// package_base com.ex with disjoint sub-namespaces (domain vs application),
// resolved file-set-aware over two roots (ADR-0022, slice-027).
const kmpCfg = `version: 1
languages:
  kotlin: ["**/*.kt"]
layers:
  domain:      {globs: ["mod-a/**"], role: domain}
  application: {globs: ["mod-b/**"], role: app}
edges:
  - {from: application, to: domain}
resolution:
  kotlin: {mode: fixed-root, package_base: com.ex, roots: ["mod-a/src/commonMain/kotlin/com/ex", "mod-b/src/commonMain/kotlin/com/ex"]}
`

func TestKmpDisjointCleanExit0(t *testing.T) { // AC-FA-CONF-001 / ADR-0022 AC1+AC4: disjunkte Multi-Modul-Config lädt UND löst sauber -> Exit 0
	dir := writeRepo(t, map[string]string{
		".a-check.yml": kmpCfg,
		"mod-a/src/commonMain/kotlin/com/ex/domain/A.kt":      "package com.ex.domain\n",
		"mod-b/src/commonMain/kotlin/com/ex/application/B.kt": "package com.ex.application\nimport com.ex.domain.A\n", // application -> domain erlaubt
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 0 {
		t.Fatalf("sauberes disjunktes Multi-Modul: erwarte Exit 0, got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
}

func TestKmpDisjointDomainToApplicationExit1(t *testing.T) { // AC-FA-CONF-001 / ADR-0022 AC2: verbotene domain->application-Kante gemeldet (vor slice-027: 0 Befunde)
	dir := writeRepo(t, map[string]string{
		".a-check.yml": kmpCfg,
		"mod-a/src/commonMain/kotlin/com/ex/domain/A.kt":      "package com.ex.domain\nimport com.ex.application.B\n", // verboten
		"mod-b/src/commonMain/kotlin/com/ex/application/B.kt": "package com.ex.application\n",
	})
	var out, errb bytes.Buffer
	code := cli.Run([]string{dir}, &out, &errb)
	if code != 1 {
		t.Fatalf("domain importiert application: erwarte Exit 1, got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
	if strings.Count(out.String(), "impurity") != 1 || !strings.Contains(out.String(), "com.ex.application.B") {
		t.Fatalf("erwarte genau 1 impurity-Befund über com.ex.application.B, got %q", out.String())
	}
}

func TestKmpSameFqnTwoRootsExit2(t *testing.T) { // AC-FA-CONF-001 / ADR-0022 AC5: gleicher FQN real in 2 roots mit VERSCHIEDENEN Schichten -> Exit 2 (fail-closed)
	dir := writeRepo(t, map[string]string{
		".a-check.yml": kmpCfg,
		"mod-a/src/commonMain/kotlin/com/ex/util/X.kt":   "package com.ex.util\n", // domain-Schicht
		"mod-b/src/commonMain/kotlin/com/ex/util/X.kt":   "package com.ex.util\n", // application-Schicht — gleicher FQN, andere Schicht
		"mod-a/src/commonMain/kotlin/com/ex/domain/A.kt": "package com.ex.domain\nimport com.ex.util.X\n",
	})
	var out, errb bytes.Buffer
	code := cli.Run([]string{dir}, &out, &errb)
	if code != 2 {
		t.Fatalf("gleicher FQN in 2 roots/2 Schichten: erwarte Exit 2, got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
	if !strings.Contains(errb.String(), "mehrdeutige Mehr-Wurzel-Auflösung") {
		t.Fatalf("stderr muss die Mehrdeutigkeit nennen, got %q", errb.String())
	}
	if out.String() != "" {
		t.Fatalf("bei fail-closed keine Findings auf stdout, got %q", out.String())
	}
}

func TestPrintConfigDocumentsResolution(t *testing.T) { // AC-FA-CONF-001 / ADR-0022 AC6: --print-config dokumentiert die Multi-Modul-Resolution
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-config"}, &out, &errb); code != 0 {
		t.Fatalf("exit %d", code)
	}
	o := out.String()
	for _, want := range []string{"resolution", "fixed-root", "package_base"} {
		if !strings.Contains(o, want) {
			t.Fatalf("print-config muss die Multi-Modul-Resolution dokumentieren (%q fehlt): %q", want, o)
		}
	}
}
