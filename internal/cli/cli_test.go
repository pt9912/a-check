package cli_test

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
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
	if !strings.Contains(o, "a-check-graph:") || !strings.Contains(o, "--print-graph /src") {
		t.Fatalf("print-mk: a-check-graph-Target fehlt (AC-FA-DIST-001 0.20.0): %q", o)
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
	if !strings.Contains(errb.String(), "existiert real unter mehreren Wurzeln") { // Scan-Zeit-Phrase (der alte Ladezeit-Guard formulierte anders)
		t.Fatalf("stderr muss die scan-zeitliche Mehrdeutigkeit nennen, got %q", errb.String())
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

func TestKmpSharedRootClassNoSpuriousAmbiguity(t *testing.T) { // ADR-0022 (Review A1): Klasse direkt unter package_base, real nur in EINEM Modul -> kein spurioses Exit 2
	dir := writeRepo(t, map[string]string{
		".a-check.yml": kmpCfg,
		"mod-a/src/commonMain/kotlin/com/ex/Shared.kt":        "package com.ex\n",                                     // direkt unter package_base, nur in mod-a (domain)
		"mod-b/src/commonMain/kotlin/com/ex/application/B.kt": "package com.ex.application\nimport com.ex.Shared\n",   // application -> domain (erlaubt); mod-b hat KEIN com.ex.Shared
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 0 {
		t.Fatalf("Klasse real nur unter mod-a (Elternpaket in mod-b existiert nur zufällig) darf keine Mehrdeutigkeit erzeugen: erwarte Exit 0, got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
}

func TestKmpSameFqnSameLayerExit0(t *testing.T) { // ADR-0022 §5 (Review C4): gleicher FQN real in 2 roots DERSELBEN Schicht (expect/actual) -> kein Exit 2
	cfg := `version: 1
languages:
  kotlin: ["**/*.kt"]
layers:
  domain: {globs: ["mod-a/**", "mod-b/**"], role: domain}
edges:
  - {from: domain, to: domain}
resolution:
  kotlin: {mode: fixed-root, package_base: com.ex, roots: ["mod-a/src/commonMain/kotlin/com/ex", "mod-b/src/commonMain/kotlin/com/ex"]}
`
	dir := writeRepo(t, map[string]string{
		".a-check.yml": cfg,
		"mod-a/src/commonMain/kotlin/com/ex/util/X.kt": "package com.ex.util\n",
		"mod-b/src/commonMain/kotlin/com/ex/util/X.kt": "package com.ex.util\n", // gleicher FQN com.ex.util.X, beide in Schicht domain
		"mod-a/src/commonMain/kotlin/com/ex/A.kt":      "package com.ex\nimport com.ex.util.X\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 0 {
		t.Fatalf("gleicher FQN in 2 roots DERSELBEN Schicht (expect/actual) muss sauber auflösen: erwarte Exit 0, got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
}

func TestKmpWildcardImportExit1(t *testing.T) { // ADR-0022 (Review C5): Wildcard-Import a.b.* end-to-end (Symbol mit Trailing-Dot -> Paket-Verzeichnis)
	dir := writeRepo(t, map[string]string{
		".a-check.yml": kmpCfg,
		"mod-a/src/commonMain/kotlin/com/ex/domain/A.kt":      "package com.ex.domain\nimport com.ex.application.*\n", // Wildcard, verboten
		"mod-b/src/commonMain/kotlin/com/ex/application/B.kt": "package com.ex.application\n",
	})
	var out, errb bytes.Buffer
	code := cli.Run([]string{dir}, &out, &errb)
	if code != 1 {
		t.Fatalf("Wildcard domain -> application.*: erwarte Exit 1, got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
	if strings.Count(out.String(), "impurity") != 1 {
		t.Fatalf("erwarte genau 1 impurity-Befund (Wildcard), got %q", out.String())
	}
}

// splitCfg is a JVM-Kotlin split-package geometry (slice-031): the package
// com.ex.conn lives under a ports module (mod-p) AND an adapters module (mod-a).
const splitCfg = `version: 1
languages:
  kotlin: ["**/*.kt"]
layers:
  ports:    {globs: ["mod-p/**"], role: port}
  adapters: {globs: ["mod-a/**"], role: adapter}
edges:
  - {from: adapters, to: ports}
resolution:
  kotlin: {mode: fixed-root, package_base: com.ex, roots: ["mod-p/src/main/kotlin/com/ex", "mod-a/src/main/kotlin/com/ex"]}
`

func TestSplitPackageClassResolvesExit0(t *testing.T) { // slice-031 AC1/AC2: Split-Package-Symbol (Datei≠Name) löst über die Deklaration -> kein spurioses Exit 2
	dir := writeRepo(t, map[string]string{
		".a-check.yml":                              splitCfg,
		"mod-p/src/main/kotlin/com/ex/conn/Api.kt":  "package com.ex.conn\ninterface Pool\n",       // Pool in Api.kt (Datei≠Name), ports
		"mod-a/src/main/kotlin/com/ex/conn/Impl.kt": "package com.ex.conn\nfun Pool.asJdbc() {}\n", // asJdbc in Impl.kt, adapters
		"mod-a/src/main/kotlin/com/ex/svc/Reader.kt": "package com.ex.svc\nimport com.ex.conn.Pool\n", // adapter -> port-Typ Pool (erlaubt)
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 0 {
		t.Fatalf("'Pool' (in Api.kt deklariert) muss auf ports auflösen, adapters->ports erlaubt -> Exit 0 (vor slice-031: Exit 2); got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
}

func TestSplitPackageExtFunForbiddenEdgeExit1(t *testing.T) { // slice-031 AC4: ext-fun asJdbc (Adapter) aus einer ports-Datei importiert -> löst auf adapters -> port-impurity
	dir := writeRepo(t, map[string]string{
		".a-check.yml":                              splitCfg,
		"mod-p/src/main/kotlin/com/ex/conn/Api.kt":  "package com.ex.conn\ninterface Pool\n",
		"mod-a/src/main/kotlin/com/ex/conn/Impl.kt": "package com.ex.conn\nfun Pool.asJdbc() {}\n",
		"mod-p/src/main/kotlin/com/ex/use/Uses.kt":  "package com.ex.use\nimport com.ex.conn.asJdbc\n", // ports -> Adapter-Symbol
	})
	var out, errb bytes.Buffer
	code := cli.Run([]string{dir}, &out, &errb)
	if code != 1 || !strings.Contains(out.String(), "port-impurity") {
		t.Fatalf("asJdbc muss auf adapters auflösen und die ports->adapters-Kante als port-impurity melden (Exit 1); got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
}

func TestSplitPackageUndeclaredResidualExtern(t *testing.T) { // slice-031 AC6: Symbol ohne Deklaration, Paketverzeichnis in 2 versch. Schichten -> extern (fail-open, kein Exit 2)
	dir := writeRepo(t, map[string]string{
		".a-check.yml":                              splitCfg,
		"mod-p/src/main/kotlin/com/ex/conn/Api.kt":  "package com.ex.conn\ninterface Pool\n",
		"mod-a/src/main/kotlin/com/ex/conn/Impl.kt": "package com.ex.conn\nclass HikariPool\n",
		"mod-a/src/main/kotlin/com/ex/svc/R.kt":     "package com.ex.svc\nimport com.ex.conn.Ghost\n", // Ghost nirgends deklariert
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 0 {
		t.Fatalf("'Ghost' (nirgends deklariert, conn-Verzeichnis in ports UND adapters) muss extern bleiben (fail-open), kein Exit 2; got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
}

func TestSplitPackageUndeclaredUniqueResolves(t *testing.T) { // slice-031 AC7: Symbol ohne Deklaration, Paketverzeichnis in GENAU EINEM Root -> löst (Rückwärtskompatibilität)
	dir := writeRepo(t, map[string]string{
		".a-check.yml":                             splitCfg,
		"mod-a/src/main/kotlin/com/ex/only/Box.kt": "package com.ex.only\nclass Box\n",                  // Paket com.ex.only NUR in mod-a (adapters)
		"mod-p/src/main/kotlin/com/ex/use/U.kt":    "package com.ex.use\nimport com.ex.only.Missing\n", // Missing nicht deklariert, only-dir nur in mod-a
	})
	var out, errb bytes.Buffer
	code := cli.Run([]string{dir}, &out, &errb)
	if code != 1 || !strings.Contains(out.String(), "port-impurity") {
		t.Fatalf("'Missing' (nicht deklariert, only-Verzeichnis nur in mod-a=adapters) muss via Paketverzeichnis eindeutig auf adapters auflösen -> ports->adapters port-impurity (Exit 1); got %d (out=%q err=%q)", code, out.String(), errb.String())
	}
}

// --- AC-FA-CLI-002 / SPEC-CLI-002: --print-graph (no-scan-Modus) -------------

const cfgGraph = `version: 1
languages:
  go: ["**/*.go"]
layers:
  core: ["internal/core/**"]
  ports: ["internal/ports/**"]
  adapters: ["internal/adapters/**"]
edges:
  - {from: adapters, to: ports}
  - {from: ports, to: core}
allow:
  - {from: ports, to: ports, reason: "Re-Export"}
composition_root: ["cmd/**"]
`

func countFiles(t *testing.T, dir string) int {
	t.Helper()
	n := 0
	_ = filepath.WalkDir(dir, func(_ string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			n++
		}
		return nil
	})
	return n
}

func TestPrintGraphHappy(t *testing.T) { // AC-FA-CLI-002 Happy
	dir := writeRepo(t, map[string]string{".a-check.yml": cfgGraph})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-graph", dir}, &out, &errb); code != 0 {
		t.Fatalf("exit %d (err=%q)", code, errb.String())
	}
	o := out.String()
	for _, want := range []string{"flowchart TB", ":::adapter", ":::domain", ":::port", "-.->|allow|", "classDef domain"} {
		if !strings.Contains(o, want) {
			t.Fatalf("print-graph fehlt %q:\n%s", want, o)
		}
	}
}

func TestPrintGraphReadOnly(t *testing.T) { // AC-FA-CLI-002 Happy: schreibt nichts
	dir := writeRepo(t, map[string]string{".a-check.yml": cfgGraph})
	before := countFiles(t, dir)
	var out, errb bytes.Buffer
	cli.Run([]string{"--print-graph", dir}, &out, &errb)
	if after := countFiles(t, dir); after != before {
		t.Fatalf("print-graph hat Dateien verändert: %d -> %d", before, after)
	}
}

func TestPrintGraphBoundaryMinimal(t *testing.T) { // AC-FA-CLI-002 Boundary: minimale Config
	minimal := `version: 1
languages:
  go: ["**/*.go"]
layers:
  core: ["core/**"]
  adapters: ["adapters/**"]
edges:
  - {from: adapters, to: core}
`
	dir := writeRepo(t, map[string]string{".a-check.yml": minimal})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-graph", dir}, &out, &errb); code != 0 {
		t.Fatalf("exit %d (err=%q)", code, errb.String())
	}
	o := out.String()
	for _, unwanted := range []string{"C0[", "S0[", ":::dangling", "subgraph"} {
		if strings.Contains(o, unwanted) {
			t.Fatalf("minimale Config: unerwartetes %q:\n%s", unwanted, o)
		}
	}
}

func TestPrintGraphMermaidUnsafeNames(t *testing.T) { // AC-FA-CLI-002 Boundary [Medium-Fix]
	cfgUnsafe := `version: 1
languages:
  go: ["**/*.go"]
layers:
  "a.b/c": ["a/**"]
  "driven-adapters": ["b/**"]
  "q]x|y": ["c/**"]
edges:
  - {from: "driven-adapters", to: "a.b/c"}
`
	dir := writeRepo(t, map[string]string{".a-check.yml": cfgUnsafe})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-graph", dir}, &out, &errb); code != 0 {
		t.Fatalf("gültige Config mit unsafe Namen: exit %d (err=%q)", code, errb.String())
	}
	o := out.String()
	if !strings.Contains(o, "&#93;") || !strings.Contains(o, "&#124;") {
		t.Fatalf("Mermaid-unsafe Namen nicht escaped:\n%s", o)
	}
	if strings.Contains(o, "q]x|y") {
		t.Fatalf("roher unsafe Name im Output:\n%s", o)
	}
}

func TestPrintGraphDangling(t *testing.T) { // AC-FA-CLI-002 Boundary [Medium-Dangling-Fix]
	cfgDangling := `version: 1
languages:
  go: ["**/*.go"]
layers:
  core: ["core/**"]
edges:
  - {from: ghost, to: core}
`
	dir := writeRepo(t, map[string]string{".a-check.yml": cfgDangling})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-graph", dir}, &out, &errb); code != 0 {
		t.Fatalf("Dangling darf kein Exit 2 sein, got %d (err=%q)", code, errb.String())
	}
	if !strings.Contains(out.String(), ":::dangling") {
		t.Fatalf("Dangling-Knoten fehlt:\n%s", out.String())
	}
}

func TestPrintGraphUnknownLanguageExit2(t *testing.T) { // AC-FA-CLI-002 Negative [High-Fix]
	cfgRuby := `version: 1
languages:
  ruby: ["**/*.rb"]
layers:
  core: ["core/**"]
  adapters: ["adapters/**"]
edges:
  - {from: adapters, to: core}
`
	dir := writeRepo(t, map[string]string{".a-check.yml": cfgRuby})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-graph", dir}, &out, &errb); code != 2 {
		t.Fatalf("unbekannte Sprache: expected 2, got %d", code)
	}
	if out.Len() != 0 {
		t.Fatalf("kein halbes Diagramm erwartet, stdout=%q", out.String())
	}
}

func TestPrintGraphInvalidConfigExit2(t *testing.T) { // AC-FA-CLI-002 Negative
	dir := writeRepo(t, map[string]string{".a-check.yml": "version: 1\nbogus_key: 1\n"})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-graph", dir}, &out, &errb); code != 2 {
		t.Fatalf("unbekannter Schlüssel: expected 2, got %d", code)
	}
}

func TestPrintGraphMissingConfig(t *testing.T) { // AC-FA-CLI-002 Negative
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-graph", t.TempDir()}, &out, &errb); code != 2 {
		t.Fatalf("fehlende Config: expected 2, got %d", code)
	}
}

func TestPrintGraphRestArgExit2(t *testing.T) { // AC-FA-CLI-002 Negative: Restargument nach dem Pfad
	dir := writeRepo(t, map[string]string{".a-check.yml": cfgGraph})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-graph", dir, "--bogus"}, &out, &errb); code != 2 {
		t.Fatalf("Restargument nach dem Pfad: expected 2, got %d", code)
	}
	if out.Len() != 0 {
		t.Fatalf("kein halbes Diagramm bei Restargument, stdout=%q", out.String())
	}
}

func TestPrintGraphUnknownFlagExit2(t *testing.T) { // AC-FA-CLI-002 Negative: unbekanntes Flag vor dem Pfad
	var out, errb bytes.Buffer
	if code := cli.Run([]string{"--print-graph", "--bogus"}, &out, &errb); code != 2 {
		t.Fatalf("unbekanntes Flag: expected 2, got %d", code)
	}
}

func TestPrintGraphDeterminism(t *testing.T) { // AC-QA-01: zwei Läufe byte-identisch
	dir := writeRepo(t, map[string]string{".a-check.yml": cfgGraph})
	var o1, o2, e bytes.Buffer
	cli.Run([]string{"--print-graph", dir}, &o1, &e)
	cli.Run([]string{"--print-graph", dir}, &o2, &e)
	if o1.String() != o2.String() {
		t.Fatalf("nicht deterministisch:\n%s\n---\n%s", o1.String(), o2.String())
	}
}

// --- AC-FA-RULE-011: construct-leak End-to-End (ADR-0027) -------------------

// parityCfg bildet die b-cad-Struktur nach und drückt deren grep-Regel P1
// („dlopen/dlsym/dlclose nur in src/adapters/plugin") als constructs-Eintrag aus.
const parityCfg = `version: 1
languages:
  cpp: ["**/*.cpp", "**/*.h"]
layers:
  model:      ["src/model/**"]
  plugin_api: ["src/plugin_api/**"]
  adapters:   ["src/adapters/**"]
  plugins:    ["plugins/**"]
edges:
  - {from: plugins,  to: plugin_api}
  - {from: plugins,  to: model}
  - {from: adapters, to: model}
composition_root: ["src/main.cpp"]
constructs:
  - {pattern: '\b(dlopen|dlsym|dlclose)\s*\(', match: regex, adapter: "src/adapters/plugin", composition_root: forbid}
`

// parityTree: eine Datei je Fall — in der Zone, außerhalb, in einer Datei ohne
// Schicht, in der Composition Root, nur im Kommentar, ohne Vorkommen.
func parityTree() map[string]string {
	return map[string]string{
		".a-check.yml":                   parityCfg,
		"src/adapters/plugin/loader.cpp": "#include \"p.h\"\nvoid load() {\n  h = dlopen(path);\n  s = dlsym(h, n);\n}\n",
		"src/adapters/io/reader.cpp":     "#include \"r.h\"\nvoid read() {\n  h = dlopen(path);\n}\n",
		"plugins/example/p.cpp":          "#include \"a.h\"\nvoid init() { dlsym(h, n); }\n",
		"src/main.cpp":                   "int main() {\n  dlclose(h);\n  return 0;\n}\n",
		"src/model/geometry.cpp":         "// frueher wurde hier dlopen(path) benutzt\nvoid area() {}\n",
		"src/model/clean.cpp":            "void volume() {}\n",
	}
}

// grepReference re-implementiert die bash-grep-Referenz des Konsumenten:
// zeilenweise Regex über den ROHEN Text (grep kennt keine Kommentare), alles
// außerhalb der erlaubten Zone ist ein Treffer.
func grepReference(t *testing.T, dir, zone string) map[string]bool {
	t.Helper()
	re := regexp.MustCompile(`\b(dlopen|dlsym|dlclose)\s*\(`)
	hits := map[string]bool{}
	err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil || d.IsDir() || !strings.HasSuffix(p, ".cpp") {
			return walkErr
		}
		rel, relErr := filepath.Rel(dir, p)
		if relErr != nil {
			return relErr
		}
		rel = filepath.ToSlash(rel)
		if strings.Contains(rel, zone) {
			return nil
		}
		body, readErr := os.ReadFile(p)
		if readErr != nil {
			return readErr
		}
		for i, ln := range strings.Split(string(body), "\n") {
			if re.MatchString(ln) {
				hits[fmt.Sprintf("%s:%d", rel, i+1)] = true
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return hits
}

// constructLeaks parst die Befund-Zeilen (`pfad:zeile: regel: meldung`) und
// liefert die construct-leak-Positionen als "pfad:zeile".
func constructLeaks(t *testing.T, stdout string) map[string]bool {
	t.Helper()
	out := map[string]bool{}
	for _, ln := range strings.Split(strings.TrimSpace(stdout), "\n") {
		if ln == "" {
			continue
		}
		parts := strings.SplitN(ln, ": ", 3)
		if len(parts) < 3 {
			t.Fatalf("unerwartetes Befund-Format: %q", ln)
		}
		if parts[1] == "construct-leak" {
			out[parts[0]] = true
		}
	}
	return out
}

func TestConstructLeakParity(t *testing.T) { // AC-FA-RULE-011: Paritätsprobe gegen die grep-Referenz
	dir := writeRepo(t, parityTree())
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 1 {
		t.Fatalf("erwarte Exit 1, got %d (out=%q)", code, out.String())
	}
	got := constructLeaks(t, out.String())
	want := grepReference(t, dir, "src/adapters/plugin")

	// DEKLARIERTE DIVERGENZ (ADR-0027): grep sieht den Kommentar, a-check nicht.
	const commentOnly = "src/model/geometry.cpp:1"
	if !want[commentOnly] {
		t.Fatalf("die grep-Referenz muss den Kommentar-Treffer sehen — sonst prüft die Probe nichts")
	}
	if got[commentOnly] {
		t.Fatalf("Treffer nur im Kommentar darf a-check NICHT melden (ausgewiesene Divergenz)")
	}
	delete(want, commentOnly)

	for pos := range want {
		if !got[pos] {
			t.Errorf("grep meldet %s, a-check nicht (Paritätslücke)", pos)
		}
	}
	for pos := range got {
		if !want[pos] {
			t.Errorf("a-check meldet %s, grep nicht (Über-Meldung)", pos)
		}
	}
	// Die drei erwarteten Treffer: außerhalb der Zone, ohne Schicht bzw. in der
	// Composition Root (composition_root: forbid).
	for _, pos := range []string{"src/adapters/io/reader.cpp:3", "plugins/example/p.cpp:2", "src/main.cpp:2"} {
		if !got[pos] {
			t.Errorf("erwarteter construct-leak fehlt: %s", pos)
		}
	}
	if len(got) != 3 {
		t.Errorf("erwarte genau 3 construct-leak-Befunde, got %v", got)
	}
}

func TestConstructLeakFitness(t *testing.T) { // AC-FA-RULE-011: Fitness-Probe (injizierter Aufruf)
	clean := map[string]string{
		".a-check.yml":                   parityCfg,
		"src/adapters/plugin/loader.cpp": "void load() { h = dlopen(path); }\n",
		"src/model/geometry.cpp":         "void area() {}\n",
	}
	var out, errb bytes.Buffer
	if code := cli.Run([]string{writeRepo(t, clean)}, &out, &errb); code != 0 {
		t.Fatalf("Gegenprobe: Aufruf INNERHALB der Zone muss grün sein, got %d (out=%q)", code, out.String())
	}

	injected := map[string]string{}
	for k, v := range clean {
		injected[k] = v
	}
	injected["src/model/geometry.cpp"] = "void area() {}\nvoid sneaky() { h = dlopen(path); }\n"
	out.Reset()
	errb.Reset()
	if code := cli.Run([]string{writeRepo(t, injected)}, &out, &errb); code != 1 {
		t.Fatalf("injizierter Aufruf außerhalb der Zone muss Exit 1 geben, got %d (out=%q)", code, out.String())
	}
	if !strings.Contains(out.String(), "src/model/geometry.cpp:2: construct-leak:") {
		t.Fatalf("Befund muss Datei, Zeile und Regel nennen: %q", out.String())
	}
	if !strings.Contains(out.String(), "src/adapters/plugin") {
		t.Fatalf("Meldung muss die erlaubte Zone nennen: %q", out.String())
	}
}

func TestConstructLeakDeterministic(t *testing.T) { // AC-QA-01: zwei Läufe byte-identisch
	dir := writeRepo(t, parityTree())
	var a, b, errb bytes.Buffer
	cli.Run([]string{dir}, &a, &errb)
	errb.Reset()
	cli.Run([]string{dir}, &b, &errb)
	if a.String() != b.String() {
		t.Fatalf("Ausgabe nicht byte-identisch:\n%q\n%q", a.String(), b.String())
	}
}

// --- ADR-0029: Abdeckungs-Diagnose (End-to-End) -----------------------------

const coverageCfg = `version: 1
languages:
  go: ["**/*.go"]
layers:
  core: ["internal/core/**"]
edges:
  - {from: core, to: core}
composition_root: ["cmd/**"]
`

func TestCoverageNotice(t *testing.T) { // ADR-0029: Diagnose ohne Exit-Code-Wechsel
	dir := writeRepo(t, map[string]string{
		".a-check.yml":         coverageCfg,
		"internal/core/svc.go": "package core\n",
		"tools/gen.go":         "package tools\n",
		"cmd/main.go":          "package main\n", // Composition Root: zaehlt nicht
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 0 {
		t.Fatalf("die Diagnose darf den Exit-Code nicht aendern, got %d", code)
	}
	e := errb.String()
	if !strings.Contains(e, "1 gescannte Datei(en) liegen in keiner Schicht") {
		t.Fatalf("Diagnose fehlt: %q", e)
	}
	if !strings.Contains(e, "  tools/gen.go\n") {
		t.Fatalf("Pfad fehlt in der Diagnose: %q", e)
	}
	if strings.Contains(e, "cmd/main.go") {
		t.Fatalf("composition_root darf nicht als ungedeckt gelten: %q", e)
	}
	if out.Len() != 0 {
		t.Fatalf("die Diagnose gehoert auf stderr, stdout war %q", out.String())
	}
}

func TestCoverageNoticeAbsentWhenCovered(t *testing.T) { // ADR-0029: kein Rauschen
	dir := writeRepo(t, map[string]string{
		".a-check.yml":         coverageCfg,
		"internal/core/svc.go": "package core\n",
	})
	var out, errb bytes.Buffer
	cli.Run([]string{dir}, &out, &errb)
	if strings.Contains(errb.String(), "keiner Schicht") {
		t.Fatalf("vollstaendig gedeckter Baum darf keine Diagnose erzeugen: %q", errb.String())
	}
}

func TestCoverageNoticeCapNamesRemainder(t *testing.T) { // ADR-0029: Kuerzung nennt ihre Groesse
	files := map[string]string{".a-check.yml": coverageCfg, "internal/core/svc.go": "package core\n"}
	for i := 0; i < 13; i++ {
		files[fmt.Sprintf("tools/g%02d.go", i)] = "package tools\n"
	}
	var out, errb bytes.Buffer
	cli.Run([]string{writeRepo(t, files)}, &out, &errb)
	e := errb.String()
	if !strings.Contains(e, "13 gescannte Datei(en)") || !strings.Contains(e, "… und 3 weitere") {
		t.Fatalf("Kuerzung muss die Restzahl nennen: %q", e)
	}
}

func TestCoverageNoticeDeterministic(t *testing.T) { // AC-QA-01
	dir := writeRepo(t, map[string]string{
		".a-check.yml":         coverageCfg,
		"internal/core/svc.go": "package core\n",
		"tools/b.go":           "package tools\n",
		"tools/a.go":           "package tools\n",
	})
	var o1, e1, o2, e2 bytes.Buffer
	cli.Run([]string{dir}, &o1, &e1)
	cli.Run([]string{dir}, &o2, &e2)
	if e1.String() != e2.String() {
		t.Fatalf("stderr nicht byte-identisch:\n%q\n%q", e1.String(), e2.String())
	}
	if !strings.Contains(e1.String(), "  tools/a.go\n  tools/b.go\n") {
		t.Fatalf("Pfade muessen stabil sortiert sein: %q", e1.String())
	}
}

// --- ADR-0031: Grenz-Diagnose (End-to-End) ----------------------------------

const pyCfg = `version: 1
languages:
  python: ["**/*.py"]
layers:
  core: ["app/core/**"]
  ui: ["app/ui/**"]
edges:
  - {from: ui, to: core}
`

// cppCfg laesst die Aufloesung beim Default `path`: ein `../`-Include kann so
// kein Ziel-Glob treffen — genau die zweite Grenz-Klasse (ADR-0031).
const cppCfg = `version: 1
languages:
  cpp: ["**/*.h", "**/*.cpp"]
layers:
  core: ["src/core/**"]
  ui: ["src/ui/**"]
edges:
  - {from: ui, to: core}
`

func TestLimitNoticePythonRelativeImport(t *testing.T) { // ADR-0031 Klasse 1
	dir := writeRepo(t, map[string]string{
		".a-check.yml":      pyCfg,
		"app/core/svc.py":   "from ..ui import widget\n",
		"app/ui/widget.py":  "x = 1\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 0 {
		t.Fatalf("die Diagnose darf den Exit-Code nicht aendern, got %d", code)
	}
	e := errb.String()
	if !strings.Contains(e, "app/core/svc.py:1: relativer Import") {
		t.Fatalf("relativer Python-Import nicht gemeldet: %q", e)
	}
	if out.Len() != 0 {
		t.Fatalf("die Diagnose gehoert auf stderr, stdout war %q", out.String())
	}
}

func TestLimitNoticePythonSecondDirective(t *testing.T) { // ADR-0031 Klasse 1
	dir := writeRepo(t, map[string]string{
		".a-check.yml":     pyCfg,
		"app/core/svc.py":  "import os, sys\n",
	})
	var out, errb bytes.Buffer
	cli.Run([]string{dir}, &out, &errb)
	if !strings.Contains(errb.String(), "app/core/svc.py:1: zweite Direktive") {
		t.Fatalf("Mehrfach-Direktive nicht gemeldet: %q", errb.String())
	}
}

// GEGENPROBE zu ADR-0035: hier sind die Glob-Praefixe ZWEIsegmentig
// ("src/core", "src/ui"), und der Kandidat "../ui/widget.h" traegt kein
// "src"-Segment — kein Praefix trifft, die Zeile bleibt also zu Recht eine
// Grenze. Ohne diese Richtung waere eine Klasse, die nie feuert, von einer
// korrekten nicht zu unterscheiden.
func TestLimitNoticeCppParentRelativeInclude(t *testing.T) { // ADR-0035 Klasse 2, Gegenprobe
	dir := writeRepo(t, map[string]string{
		".a-check.yml":    cppCfg,
		"src/core/svc.h":  "#include \"../ui/widget.h\"\n",
		"src/ui/widget.h": "\n",
	})
	var out, errb bytes.Buffer
	cli.Run([]string{dir}, &out, &errb)
	e := errb.String()
	if !strings.Contains(e, "src/core/svc.h:1: relativer Pfad") || !strings.Contains(e, `"path"`) {
		t.Fatalf("elternrelativer Include nicht mit Modus-Nennung gemeldet: %q", e)
	}
}

// ADR-0035: ein "../"-Symbol, dessen Ziel-Segment ein Glob-Praefix TRIFFT, loest
// unter mode "path" auf — der Kandidat behaelt seine Punkte, und die
// Schicht-Zuordnung sucht das Praefix segmentweise an beliebiger Stelle. Vor
// ADR-0035 erschien dieselbe Zeile gleichzeitig als Befund UND als unbeurteilt.
const cppCfgFlach = `version: 1
languages:
  cpp: ["**/*.h", "**/*.cpp"]
layers:
  core: ["core/**"]
  ui: ["ui/**"]
edges:
  - {from: ui, to: core}
`

func TestLimitNoticeNotReportedWhenGlobPrefixHits(t *testing.T) { // ADR-0035
	dir := writeRepo(t, map[string]string{
		".a-check.yml": cppCfgFlach,
		"core/svc.h":   "#include \"../ui/widget.h\"\n",
		"ui/widget.h":  "\n",
	})
	var out, errb bytes.Buffer
	cli.Run([]string{dir}, &out, &errb)
	e := errb.String()
	if strings.Contains(e, "core/svc.h:1: relativer Pfad") {
		t.Fatalf("aufloesende Zeile faelschlich als Grenze gemeldet (ADR-0035): %q", e)
	}
	if !strings.Contains(out.String(), "core/svc.h:1") {
		t.Fatalf("die Kante muss beurteilt werden — erwartet ein Befund: %q", out.String())
	}
}

// Unter `relative` loest dieselbe Zeile auf — dort ist sie keine Grenze, sondern
// der regulaere Fall. Waere die Diagnose modus-blind, spraeche sie hier falsch.
func TestLimitNoticeSilentUnderRelativeMode(t *testing.T) { // ADR-0031 Entscheidung 5
	dir := writeRepo(t, map[string]string{
		".a-check.yml":    cppCfg + "resolution:\n  cpp:\n    mode: relative\n",
		"src/core/svc.h":  "#include \"../ui/widget.h\"\n",
		"src/ui/widget.h": "\n",
	})
	var out, errb bytes.Buffer
	cli.Run([]string{dir}, &out, &errb)
	if strings.Contains(errb.String(), "Heuristik-Grenze") {
		t.Fatalf("unter mode relative darf der Include keine Grenze sein: %q", errb.String())
	}
}

func TestLimitNoticeAbsentWhenClean(t *testing.T) { // ADR-0031: kein Rauschen
	dir := writeRepo(t, map[string]string{
		".a-check.yml":     pyCfg,
		"app/core/svc.py":  "import os\nfrom collections import OrderedDict\n",
	})
	var out, errb bytes.Buffer
	cli.Run([]string{dir}, &out, &errb)
	if strings.Contains(errb.String(), "Heuristik-Grenze") {
		t.Fatalf("Baum ohne Grenz-Formen darf keine Diagnose erzeugen: %q", errb.String())
	}
}

// Zitierte Muster sind keine Muster: ein auskommentierter Import spricht UEBER
// die Form, er ist sie nicht (SL-004 — dreimal hat ein neuer Sensor im ersten
// Lauf sein eigenes Umfeld gemeldet).
func TestLimitNoticeIgnoresComments(t *testing.T) {
	dir := writeRepo(t, map[string]string{
		".a-check.yml":    pyCfg,
		"app/core/svc.py": "# from . import x\n# import a, b\nimport os\n",
		"src/core/x.h":    "",
	})
	var out, errb bytes.Buffer
	cli.Run([]string{dir}, &out, &errb)
	if strings.Contains(errb.String(), "Heuristik-Grenze") {
		t.Fatalf("auskommentierte Formen duerfen nicht melden: %q", errb.String())
	}
}

// AK 2 aus CR-1: gerade der befundfreie Lauf braucht den Hinweis — dort ist er
// das Einzige, was „sauber" von „nicht angesehen" trennt.
func TestLimitNoticeAtZeroFindings(t *testing.T) {
	dir := writeRepo(t, map[string]string{
		".a-check.yml":    pyCfg,
		"app/ui/w.py":     "from ..core import svc\n",
		"app/core/svc.py": "x = 1\n",
	})
	var out, errb bytes.Buffer
	code := cli.Run([]string{dir}, &out, &errb)
	if code != 0 {
		t.Fatalf("erwartet Exit 0 (kein Befund), got %d — stdout %q", code, out.String())
	}
	if !strings.Contains(errb.String(), "Heuristik-Grenze") {
		t.Fatalf("Hinweis fehlt gerade beim befundfreien Lauf: %q", errb.String())
	}
}

// AK 1 aus CR-1: die Diagnose fasst die Regel-Auswertung nicht an. Derselbe Baum
// einmal MIT und einmal OHNE Grenz-Form muss dieselbe Befundmenge liefern.
func TestLimitNoticeLeavesFindingsUntouched(t *testing.T) {
	// fixed-root, damit der gepunktete Modulpfad ueberhaupt auf ein Schicht-Glob
	// auflöst — sonst gaebe es keinen Befund, gegen den die Diagnose sich messen
	// liesse.
	base := map[string]string{
		".a-check.yml":     pyCfg + "resolution:\n  python:\n    mode: fixed-root\n    package_base: app\n    roots: [app]\n",
		"app/core/svc.py":  "import app.ui.widget\n", // echter Verstoss: core -> ui
		"app/ui/widget.py": "x = 1\n",
	}
	var o1, e1 bytes.Buffer
	c1 := cli.Run([]string{writeRepo(t, base)}, &o1, &e1)

	withLimit := map[string]string{}
	for k, v := range base {
		withLimit[k] = v
	}
	withLimit["app/ui/widget.py"] = "from . import other\nx = 1\n"
	var o2, e2 bytes.Buffer
	c2 := cli.Run([]string{writeRepo(t, withLimit)}, &o2, &e2)

	if c1 != c2 || c1 != 1 {
		t.Fatalf("Exit-Code veraendert: %d vs. %d (want 1)", c1, c2)
	}
	if o1.String() != o2.String() {
		t.Fatalf("Befundmenge veraendert:\n%q\n%q", o1.String(), o2.String())
	}
	if !strings.Contains(e2.String(), "Heuristik-Grenze") {
		t.Fatalf("Diagnose fehlt im zweiten Lauf: %q", e2.String())
	}
}

func TestLimitNoticeCapNamesRemainder(t *testing.T) { // ADR-0031: keine stille Kappung
	files := map[string]string{".a-check.yml": pyCfg}
	for i := 0; i < 13; i++ {
		files[fmt.Sprintf("app/core/m%02d.py", i)] = "from . import x\n"
	}
	var out, errb bytes.Buffer
	cli.Run([]string{writeRepo(t, files)}, &out, &errb)
	e := errb.String()
	if !strings.Contains(e, "13 Import-Zeile(n)") || !strings.Contains(e, "… und 3 weitere") {
		t.Fatalf("Kuerzung muss die Restzahl nennen: %q", e)
	}
}

func TestLimitNoticeDeterministic(t *testing.T) { // AC-QA-01
	dir := writeRepo(t, map[string]string{
		".a-check.yml":    pyCfg,
		"app/core/b.py":   "from . import x\n",
		"app/core/a.py":   "import os, sys\n",
	})
	var o1, e1, o2, e2 bytes.Buffer
	cli.Run([]string{dir}, &o1, &e1)
	cli.Run([]string{dir}, &o2, &e2)
	if e1.String() != e2.String() {
		t.Fatalf("stderr nicht byte-identisch:\n%q\n%q", e1.String(), e2.String())
	}
	if strings.Index(e1.String(), "app/core/a.py") > strings.Index(e1.String(), "app/core/b.py") {
		t.Fatalf("Zeilen muessen stabil nach Pfad sortiert sein: %q", e1.String())
	}
}

// --- ADR-0032: Aufloesungs-Diagnose (End-to-End) ----------------------------

// resCfg bildet das Konsumenten-Muster nach: sprach-praefixierte Schicht-Globs
// (`go/internal/…`), waehrend Go-Importe den Modulpfad tragen. Das
// Glob-Literalpraefix kommt darin nicht als Segment-Run vor — jedes Ziel gilt
// als repo-extern, und kein Verstoss wird je beurteilt.
const resCfg = `version: 1
languages:
  go: ["**/*.go"]
layers:
  service: ["go/internal/service/**"]
  ui: ["go/internal/ui/**"]
edges:
  - {from: ui, to: service}
`

// Dieselben Schichten OHNE Sprach-Praefix: hier loest der Modulpfad auf.
const resCfgOK = `version: 1
languages:
  go: ["**/*.go"]
layers:
  service: ["internal/service/**"]
  ui: ["internal/ui/**"]
edges:
  - {from: ui, to: service}
`

func TestResolutionNoticeWhenNothingResolves(t *testing.T) { // ADR-0032
	dir := writeRepo(t, map[string]string{
		".a-check.yml":              resCfg,
		"go/internal/service/s.go":  "package service\n\nimport \"example.com/d/internal/ui\"\n",
		"go/internal/ui/w.go":       "package ui\n\nimport \"fmt\"\n",
	})
	var out, errb bytes.Buffer
	if code := cli.Run([]string{dir}, &out, &errb); code != 0 {
		t.Fatalf("die Diagnose darf den Exit-Code nicht aendern, got %d", code)
	}
	e := errb.String()
	if !strings.Contains(e, "Schicht service: 1 Datei(en), 0 von 1 Import-Symbolen") {
		t.Fatalf("service nicht gemeldet: %q", e)
	}
	if !strings.Contains(e, "Schicht ui: 1 Datei(en), 0 von 1 Import-Symbolen") {
		t.Fatalf("ui nicht gemeldet: %q", e)
	}
	if out.Len() != 0 {
		t.Fatalf("die Diagnose gehoert auf stderr, stdout war %q", out.String())
	}
}

// Die Kern-Entscheidung von ADR-0032: EINE aufloesende Schicht genuegt, um zu
// schweigen. Die per-Schicht-Regel des CR haette hier `service` gemeldet —
// obwohl ein abhaengigkeitsfreier Kern genau so aussieht.
func TestResolutionNoticeSilentWhenOneLayerResolves(t *testing.T) {
	dir := writeRepo(t, map[string]string{
		".a-check.yml":           resCfgOK,
		"internal/service/s.go":  "package service\n\nimport \"fmt\"\n",
		"internal/ui/w.go":       "package ui\n\nimport \"example.com/d/internal/service\"\n",
	})
	var out, errb bytes.Buffer
	cli.Run([]string{dir}, &out, &errb)
	if strings.Contains(errb.String(), "lösen auf eine Schicht auf") {
		t.Fatalf("eine aufloesende Schicht muss die Diagnose verstummen lassen: %q", errb.String())
	}
}

func TestResolutionNoticeSilentWithoutSymbols(t *testing.T) { // ADR-0032 Entscheidung 3
	dir := writeRepo(t, map[string]string{
		".a-check.yml":             resCfg,
		"go/internal/service/s.go": "package service\n\ntype T struct{}\n",
		"go/internal/ui/w.go":      "package ui\n\ntype U struct{}\n",
	})
	var out, errb bytes.Buffer
	cli.Run([]string{dir}, &out, &errb)
	if strings.Contains(errb.String(), "lösen auf eine Schicht auf") {
		t.Fatalf("Schichten ohne Importe sind normal, kein Befund: %q", errb.String())
	}
}

// Eine Schicht ohne Symbole wird nicht MITGELISTET, auch wenn die Diagnose
// wegen der anderen Schicht feuert — sonst laese sich Stille als Auffaelligkeit.
func TestResolutionNoticeOmitsSymbolLessLayer(t *testing.T) {
	dir := writeRepo(t, map[string]string{
		".a-check.yml":             resCfg,
		"go/internal/service/s.go": "package service\n\nimport \"example.com/d/internal/ui\"\n",
		"go/internal/ui/w.go":      "package ui\n\ntype U struct{}\n",
	})
	var out, errb bytes.Buffer
	cli.Run([]string{dir}, &out, &errb)
	e := errb.String()
	if !strings.Contains(e, "Schicht service:") {
		t.Fatalf("service muss gemeldet werden: %q", e)
	}
	if strings.Contains(e, "Schicht ui:") {
		t.Fatalf("Schicht ohne Symbole darf nicht mitgelistet werden: %q", e)
	}
}

func TestResolutionNoticeDeterministic(t *testing.T) { // AC-QA-01
	dir := writeRepo(t, map[string]string{
		".a-check.yml":             resCfg,
		"go/internal/service/s.go": "package service\n\nimport \"example.com/d/internal/ui\"\n",
		"go/internal/ui/w.go":      "package ui\n\nimport \"example.com/d/internal/service\"\n",
	})
	var o1, e1, o2, e2 bytes.Buffer
	cli.Run([]string{dir}, &o1, &e1)
	cli.Run([]string{dir}, &o2, &e2)
	if e1.String() != e2.String() {
		t.Fatalf("stderr nicht byte-identisch:\n%q\n%q", e1.String(), e2.String())
	}
	if strings.Index(e1.String(), "Schicht service") > strings.Index(e1.String(), "Schicht ui") {
		t.Fatalf("Schichten muessen stabil nach Namen sortiert sein: %q", e1.String())
	}
}
