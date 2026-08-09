#!/usr/bin/env bash
# gate-consistency.sh — Meta-Gate gegen Harness-Lügen (Regelwerk Modul 13 /
# §Durchsetzungsschicht; Stack-Vorbild d-check tools/gate-consistency.sh):
#
#   (1) Jedes in AGENTS.md §4 bzw. harness/README.md §Sensors als
#       Tabellenzeile dokumentierte `make`-Target existiert real (Makefile
#       oder das includebare d-check.mk) — kein halluziniertes Gate.
#   (2) Jedes reale Gate-Target (Makefile + d-check.mk, ohne die Utility-
#       Targets help/build) ist in AGENTS.md §4 gelistet — AGENTS' eigene
#       Zusage „Nur hier gelistete Targets existieren im Makefile".
#   (3) Die modules-Liste der .a-check-Doku-Konfig (.d-check.yml) trägt die
#       aktiven Module (links/anchors/ids/matrix) und NICHT external — sonst
#       verliert der netzlose doc-check still seine Beweis-Aussage (AC-QA-02).
#   (6) Jede ADR-Datei ist im ADR-Index verlinkt (slice-087, SL-005).
#
# Vor der echten Prüfung läuft ein Selbsttest: ein Phantom-Target muss das
# Gate nachweislich feuern lassen.
set -euo pipefail
cd "$(dirname "$0")/.."

# Utility-Targets: keine Gates, müssen nicht in AGENTS §4 stehen.
# arch-graph druckt nur den Architektur-Graphen (kein Pass/Fail), daher Utility.
UTILITY_TARGETS='help build compile hooks arch-graph'

# Bewusst NICHT `.PHONY`: Targets, die tatsächlich eine gleichnamige Datei
# erzeugen. Heute keines — jedes Target dieses Makefiles ist ein Kommando.
# Die Liste ist der deklarierte Ort für eine künftige Ausnahme; sie leer zu
# lassen ist eine Aussage, keine Auslassung.
NON_PHONY_TARGETS=''

# Dokumentierte Targets: alle `make <name>`-Tokens in Tabellenzeilen.
doc_targets() {
  grep -E '^\|' "$1" | grep -oE '`make [a-z][a-z0-9_-]*`' \
    | sed -E 's/`make ([a-z0-9_-]+)`/\1/' | sort -u
}

# Reale Targets aus allen Makefile-Fragmenten (Makefile + includebare *.mk):
# Regelzeilen am Zeilenanfang, auch Mehrfach-Targets (`a b: dep`).
# Zuweisungen (`X := y`, `X ?= y`) und `.PHONY`/`.DEFAULT_GOAL` sind
# ausgeschlossen (führendes `.` bzw. `=` nach dem Doppelpunkt).
makefile_targets() {
  cat "$@" | grep -oE '^[a-zA-Z][a-zA-Z0-9 _-]*:([^=]|$)' \
    | sed 's/:.*//' | tr ' ' '\n' | sed '/^$/d' | sort -u
}

# Die `.PHONY`-Menge eines Makefiles: alle `.PHONY:`-Zeilen vereinigt,
# Zeilenfortsetzungen (`\` am Ende) eingeschlossen.
phony_targets() {  # $1 = Makefile
  sed -E ':a; /\\$/ { N; s/\\\n//; ba }' "$1" \
    | grep -E '^\.PHONY:' | sed -E 's/^\.PHONY:[[:space:]]*//' \
    | tr ' \t' '\n\n' | sed '/^$/d' | sort -u
}

# Jedes eigene Rezept-Target ist `.PHONY`. Fehlt die Deklaration, überspringt
# make das Rezept, sobald eine gleichnamige Datei existiert — und meldet
# Exit 0, ohne es ausgeführt zu haben (slice-068, Fund F-1).
# Geprüft wird die **Differenzmenge**, nicht ein Beispiel: eine Ein-Target-Probe
# würde einen Teil-Fix als vollständig ausweisen (Plan-Review R-068-F2).
# Grenze, ausdrücklich: nur die Targets, die dieses Repo selbst definiert.
# `d-check.mk` ist Fremdlieferung — ein Befund dort gehört gemeldet, nicht
# lokal gepatcht.
check_phony_complete() {  # $1 = Makefile
  local fail=0 t phony
  phony="$(phony_targets "$1")"
  while IFS= read -r t; do
    [ -z "$t" ] && continue
    grep -qw "$t" <<<"$NON_PHONY_TARGETS" && continue
    if ! grep -qx "$t" <<<"$phony"; then
      echo "gate-consistency: FAIL — Target '$t' fehlt in .PHONY ($1); eine gleichnamige Datei laesst make das Rezept ueberspringen und Exit 0 melden" >&2
      fail=1
    fi
  done <<<"$(makefile_targets "$1")"
  return "$fail"
}

# nutzt globales MK_TARGETS
check_documented_exist() {
  local fail=0 doc t
  for doc in "$@"; do
    while IFS= read -r t; do
      [ -z "$t" ] && continue
      if ! grep -qx "$t" <<<"$MK_TARGETS"; then
        echo "gate-consistency: FAIL — $doc dokumentiert 'make $t', das aber kein reales Target ist" >&2
        fail=1
      fi
    done <<<"$(doc_targets "$doc")"
  done
  return "$fail"
}

self_test() {
  local tmp
  tmp="$(mktemp -d)"
  printf '| `make phantom-target` | x |\n' > "$tmp/doc.md"
  printf 'echtes-target zweites-target: dep\n\ttrue\nVAR := x\n' > "$tmp/Makefile"
  MK_TARGETS="$(makefile_targets "$tmp/Makefile")"
  if check_documented_exist "$tmp/doc.md" 2>/dev/null; then
    echo "gate-consistency: Selbsttest FEHLGESCHLAGEN — Phantom-Target nicht erkannt" >&2
    rm -rf "$tmp"
    exit 2
  fi
  if [ "$(makefile_targets "$tmp/Makefile" | wc -l)" -ne 2 ]; then
    echo "gate-consistency: Selbsttest FEHLGESCHLAGEN — Makefile-Parser (Mehrfach-Targets/Zuweisungen)" >&2
    rm -rf "$tmp"
    exit 2
  fi
  rm -rf "$tmp"
}

# Fitness-Function für den .PHONY-Check: er muss bei einer Lücke feuern **und**
# bei Vollständigkeit schweigen. Ohne die zweite Hälfte wäre ein Sensor, der
# immer rot meldet, von einem korrekten nicht zu unterscheiden.
phony_self_test() {
  local tmp
  tmp="$(mktemp -d)"
  # (a) Lücke: zwei Targets, nur eines deklariert — muss feuern.
  printf '.PHONY: erstes\nerstes:\n\ttrue\nzweites:\n\ttrue\n' > "$tmp/Makefile"
  if check_phony_complete "$tmp/Makefile" 2>/dev/null; then
    echo "gate-consistency: Selbsttest FEHLGESCHLAGEN — fehlendes .PHONY nicht erkannt" >&2
    rm -rf "$tmp"
    exit 2
  fi
  # (b) Vollständig, über eine Zeilenfortsetzung — muss schweigen.
  printf '.PHONY: erstes \\\n        zweites\nerstes:\n\ttrue\nzweites:\n\ttrue\n' > "$tmp/Makefile"
  if ! check_phony_complete "$tmp/Makefile" 2>/dev/null; then
    echo "gate-consistency: Selbsttest FEHLGESCHLAGEN — vollstaendiges .PHONY faelschlich als Luecke gemeldet" >&2
    rm -rf "$tmp"
    exit 2
  fi
  rm -rf "$tmp"
}

# --- (6) ADR-Index-Vollständigkeit (slice-087, SL-005) ---------------------
# Die GEGENRICHTUNG zu doc-check: das prüft, dass jeder Link AUFLÖST, nicht dass
# jede Datei VERLINKT IST. Was nicht verlinkt ist, kann nicht ins Leere zeigen —
# ein handgepflegter Index sieht darum immer vollständig aus, egal wie viele
# Einträge fehlen. Genau so fehlten ADR-0030 und ADR-0031 im Index, beide bei
# grünem `make gates`, beide nur zufällig bemerkt.
#
# Dieselbe Asymmetrie hatte slice-071 bei regelwerk-check gefunden (Manifest →
# Baum, nicht Baum → Manifest); die Antwort ist dieselbe: `comm -13` über beide
# Mengen. Zwei Vorkommen in verschiedenen Sensoren machen sie zur Klasse.
#
# Geprüft wird NUR die fehlende Richtung. Ein Index-Eintrag ohne Datei ist ein
# toter Link und gehört doc-check — hier ihn mitzuprüfen wäre die Doppelung, die
# slice-079 an anderer Stelle gerade abträgt.
adr_files() {   # $1 = ADR-Verzeichnis -> Dateinamen, sortiert
  find "$1" -maxdepth 1 -name '[0-9]*.md' -printf '%f\n' | sort
}

# Verlinkte Ziele des Index: die Dateinamen aus Markdown-Links. Ein Anker-Suffix
# (`0031-x.md#abschnitt`) zählt als Verlinkung derselben Datei.
adr_linked() { # $1 = Index-Datei -> verlinkte Dateinamen, sortiert
  grep -oE '\]\([0-9][0-9a-z.-]*\.md(#[^)]*)?\)' "$1" \
    | sed -E 's/^\]\(//; s/#.*$//; s/\)$//' | sort -u
}

check_adr_index() { # $1 = ADR-Verzeichnis, $2 = Index-Datei
  local missing fail=0
  missing="$(comm -23 <(adr_files "$1") <(adr_linked "$2"))"
  if [ -n "$missing" ]; then
    while read -r f; do
      [ -n "$f" ] || continue
      echo "gate-consistency: FAIL — ADR '$f' fehlt im Index $2 (SL-005)" >&2
      fail=1
    done <<<"$missing"
  fi
  return "$fail"
}

# Selbsttest beider Richtungen: die Lücke muss feuern, der vollständige Index
# muss schweigen. Ohne die zweite Probe wäre ein immer-rotes Gate unentdeckt.
adr_index_self_test() {
  local tmp
  tmp="$(mktemp -d)"
  mkdir -p "$tmp/adr"
  : > "$tmp/adr/0001-erste.md"
  : > "$tmp/adr/0002-zweite.md"
  # (a) Lücke: nur eine von zwei ADRs verlinkt — muss feuern.
  printf '| [ADR-0001](0001-erste.md) | Titel |\n' > "$tmp/index.md"
  if check_adr_index "$tmp/adr" "$tmp/index.md" 2>/dev/null; then
    echo "gate-consistency: Selbsttest FEHLGESCHLAGEN — fehlender ADR-Index-Eintrag nicht erkannt" >&2
    rm -rf "$tmp"
    exit 2
  fi
  # (b) Vollständig, die zweite Zeile mit Anker — muss schweigen.
  printf '| [ADR-0001](0001-erste.md) |\n| [ADR-0002](0002-zweite.md#geschichte) |\n' > "$tmp/index.md"
  if ! check_adr_index "$tmp/adr" "$tmp/index.md" 2>/dev/null; then
    echo "gate-consistency: Selbsttest FEHLGESCHLAGEN — vollstaendiger ADR-Index faelschlich als Luecke gemeldet" >&2
    rm -rf "$tmp"
    exit 2
  fi
  rm -rf "$tmp"
}

# --- Pin-Konsistenz (slice-018, Opt 3) -------------------------------------
# Fängt zwei Drift-Klassen, die d-checks tag-basierte Module `versions`/`pins`
# NICHT fangen (a-check pinnt per Digest, nicht per Tag):
#   (A) Versions-Nummer: version.md#aktuell == aktuellstes CHANGELOG-Release.
#   (B) Digest: der volle @sha256-Digest ist in a-check.mk,
#       den `docker run`-Beispielen BEIDER READMEs (README.md + README.de.md) UND
#       version.md#aktuell identisch (die zwei harten, nicht verlinkbaren Pins + die
#       zwei Kommando-Beispiele + die eine Prosa-Wahrheit).
# Jede harte Pin-Datei muss GENAU EINEN a-check-Digest tragen — ein zweiter,
# abweichender (Decoy/Alt-Kommentar) macht den effektiven Pin mehrdeutig und
# könnte echte Drift maskieren ⇒ fail-closed.
# Zusätzlich (Maintainer-Entscheid slice-018): d-check.mk trägt einen wohlgeformten
# Tag (aus der `DCHECK_IMAGE`-Zeile) und einen wohlgeformten DCHECK_DIGEST.
# GRENZE (ehrliche Heuristik, AC-QA-02): dass der Tag :vX.Y.Z tatsächlich auf
# @sha256:… auflöst — und dass Prosa-/Kommentar-Versionsnummern stimmen —, ist
# eine Registry-/Netz-Eigenschaft bzw. Kommentar-Semantik; netzlos nicht prüfbar,
# beim Re-Pin (online) verifiziert. Das Gate prüft die offline-belegbare
# Digest-Deklarations-Konsistenz.

# Eindeutige a-check-Digests einer Datei (je Zeile einer). `|| true`: kein Treffer
# ⇒ Exit 0 (robust unter `set -e`, auch bei bare-Aufruf ohne ||-Kontext).
a_check_digests() {  # $1 = Datei
  grep -oE 'a-check@sha256:[0-9a-f]{64}' "$1" 2>/dev/null | sort -u || true
}
# Zeilenzahl eines Strings (0 für leer).
_nlines() { printf '%s' "$1" | grep -c . || true; }

pin_consistency() {
  local root="$1" fail=0
  local mk="$root/a-check.mk"
  local readme="$root/README.md" dereadme="$root/README.de.md" ver="$root/version.md"
  local chg="$root/CHANGELOG.md" dmk="$root/d-check.mk"

  # (B) Digest-Gleichheit. version.md#aktuell ist die Anker-Wahrheit — auf den
  # `## Aktuell`-Abschnitt verankert, damit `## Verlauf` unbeachtet bleibt.
  local ver_block d_ver
  ver_block="$(sed -n '/^## Aktuell/,/^## /p' "$ver" 2>/dev/null || true)"
  d_ver="$(printf '%s\n' "$ver_block" | grep -oE 'a-check@sha256:[0-9a-f]{64}' | sort -u || true)"
  if [ "$(_nlines "$d_ver")" != "1" ]; then
    echo "gate-consistency: FAIL — version.md#aktuell ohne genau einen a-check@sha256-Release-Digest" >&2
    fail=1; d_ver=""
  fi

  local name path digs n
  for pair in "a-check.mk:$mk" "README.md:$readme" "README.de.md:$dereadme"; do
    name="${pair%%:*}"; path="${pair#*:}"
    digs="$(a_check_digests "$path")"
    n="$(_nlines "$digs")"
    if [ "$n" != "1" ]; then
      echo "gate-consistency: FAIL — $name trägt $n a-check@sha256-Pins (genau 1 erwartet; ein zweiter/abweichender maskiert Drift)" >&2
      fail=1; continue
    fi
    if [ -n "$d_ver" ] && [ "$digs" != "$d_ver" ]; then
      echo "gate-consistency: FAIL — Digest-Drift: $name ($digs) ≠ version.md#aktuell ($d_ver)" >&2
      fail=1
    fi
  done

  # (A) Versions-Nummer: version.md#aktuell == aktuellstes CHANGELOG-Release
  # (Keep-a-Changelog-Konvention: neuestes Release zuoberst).
  local v_ver v_chg
  v_ver="$(printf '%s\n' "$ver_block" | grep -E 'Aktuelle Version:' | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+' | head -1 || true)"
  v_chg="v$(grep -oE '^## \[[0-9]+\.[0-9]+\.[0-9]+\]' "$chg" 2>/dev/null | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1 || true)"
  if [ -z "$v_ver" ]; then
    echo "gate-consistency: FAIL — version.md#aktuell ohne erkennbare 'Aktuelle Version: vX.Y.Z'" >&2
    fail=1
  elif [ "$v_ver" != "$v_chg" ]; then
    echo "gate-consistency: FAIL — Versions-Drift: version.md#aktuell ($v_ver) ≠ CHANGELOG-Release ($v_chg)" >&2
    fail=1
  fi

  # (C) d-check.mk: wohlgeformter Tag aus der DCHECK_IMAGE-Zeile + DCHECK_DIGEST
  # (nur die tragenden Zeilen, nicht die ganze Datei — legitime Kommentar-
  # Versionsnummern lösen so keinen Fehlalarm aus).
  local d_tag
  d_tag="$(grep -E '^DCHECK_IMAGE[[:space:]]*\?=' "$dmk" 2>/dev/null | grep -oE 'd-check:v[0-9]+\.[0-9]+\.[0-9]+' | head -1 || true)"
  if [ -z "$d_tag" ]; then
    echo "gate-consistency: FAIL — d-check.mk: DCHECK_IMAGE ohne wohlgeformten :vX.Y.Z-Tag" >&2
    fail=1
  fi
  if ! grep -qE '^DCHECK_DIGEST[[:space:]]*\?=[[:space:]]*sha256:[0-9a-f]{64}' "$dmk" 2>/dev/null; then
    echo "gate-consistency: FAIL — d-check.mk ohne wohlgeformten DCHECK_DIGEST (sha256:<64hex>)" >&2
    fail=1
  fi

  return "$fail"
}

# Selbsttest der Pin-Konsistenz (Fitness-Function): konsistente Fixtures ⇒ grün;
# je eine gedriftete Dimension (B Digest / A Version / C d-check) UND ein
# Decoy-Zweitdigest ⇒ rot.
pin_self_test() {
  local tmp A B F
  tmp="$(mktemp -d)"
  A="sha256:$(printf '0%.0s' $(seq 1 64))"   # konsistenter Digest
  B="sha256:$(printf '1%.0s' $(seq 1 64))"   # abweichender Digest
  F="sha256:$(printf 'f%.0s' $(seq 1 64))"   # d-check DCHECK_DIGEST
  _pin_fixture() {  # $1 README-Digest · $2 version.md-Version · $3 d-check-Digest
    printf 'A_CHECK_IMAGE ?= ghcr.io/pt9912/a-check@%s\n' "$A" > "$tmp/a-check.mk"
    printf 'run ghcr.io/pt9912/a-check@%s /src\n' "$1" > "$tmp/README.md"
    printf 'run ghcr.io/pt9912/a-check@%s /src\n' "$A" > "$tmp/README.de.md"
    printf '## Aktuell\nAktuelle Version: %s\nDigest ghcr.io/pt9912/a-check@%s\n## Verlauf\n' "$2" "$A" > "$tmp/version.md"
    printf '## [0.10.0] - 2026-07-04\n' > "$tmp/CHANGELOG.md"
    printf 'DCHECK_IMAGE ?= ghcr.io/pt9912/d-check:v0.37.1\n# historisch v0.35.0 → v0.37.1\nDCHECK_DIGEST ?= %s\n' "$3" > "$tmp/d-check.mk"
  }
  _assert_red() {  # $1 = Beschreibung der gedrifteten Dimension
    if pin_consistency "$tmp" 2>/dev/null; then
      echo "gate-consistency: Selbsttest FEHLGESCHLAGEN — $1 nicht erkannt (Fitness-Function tot)" >&2
      rm -rf "$tmp"; exit 2
    fi
  }
  # konsistent ⇒ muss grün sein (inkl. legitimer d-check-Kommentar-Version)
  _pin_fixture "$A" "v0.10.0" "$F"
  if ! pin_consistency "$tmp" 2>/dev/null; then
    echo "gate-consistency: Selbsttest FEHLGESCHLAGEN — konsistente Pins fälschlich als Drift gemeldet" >&2
    rm -rf "$tmp"; exit 2
  fi
  _pin_fixture "$B" "v0.10.0" "$F"; _assert_red "(B) Digest-Drift (README ≠ version.md)"
  _pin_fixture "$A" "v0.9.0"  "$F"; _assert_red "(A) Versions-Drift (version.md ≠ CHANGELOG)"
  _pin_fixture "$A" "v0.10.0" "sha256:short"; _assert_red "(C) d-check.mk kaputter DCHECK_DIGEST"
  # A1-Regression: zweiter, abweichender Digest (Decoy) in a-check.mk ⇒ rot.
  _pin_fixture "$A" "v0.10.0" "$F"
  printf '# alt: ghcr.io/pt9912/a-check@%s\nA_CHECK_IMAGE ?= ghcr.io/pt9912/a-check@%s\n' "$B" "$A" > "$tmp/a-check.mk"
  _assert_red "(A1) Decoy-Zweitdigest in a-check.mk"
  unset -f _pin_fixture _assert_red
  rm -rf "$tmp"
}

self_test
phony_self_test
pin_self_test
adr_index_self_test
fail=0
MK_TARGETS="$(makefile_targets Makefile d-check.mk)"

# (1) Doku → real
check_documented_exist AGENTS.md harness/README.md || fail=1

# (2) real → AGENTS §4 (ohne Utility-Targets)
agents_targets="$(doc_targets AGENTS.md)"
while IFS= read -r t; do
  [ -z "$t" ] && continue
  if grep -qw "$t" <<<"$UTILITY_TARGETS"; then
    continue
  fi
  if ! grep -qx "$t" <<<"$agents_targets"; then
    echo "gate-consistency: FAIL — reales Target '$t' fehlt in AGENTS.md §4" >&2
    fail=1
  fi
done <<<"$MK_TARGETS"

# (3) .d-check.yml-Modulliste des netzlosen doc-check (AC-QA-02)
modules_line="$(grep -E '^modules:' .d-check.yml || true)"
for m in links anchors ids matrix; do
  if [[ "$modules_line" != *"$m"* ]]; then
    echo "gate-consistency: FAIL — .d-check.yml modules ohne '$m'; der netzlose doc-check beweist AC-QA-02 nur mit den aktiven Modulen" >&2
    fail=1
  fi
done
if [[ "$modules_line" == *external* ]]; then
  echo "gate-consistency: FAIL — .d-check.yml aktiviert external; das doc-check-Gate muss netzlos bleiben (AC-QA-02)" >&2
  fail=1
fi

# (4) Pin-Konsistenz (slice-018): Digest-Gleichheit der harten Pins ==
#     version.md#aktuell, Versions-Nummer == CHANGELOG, d-check.mk-Deklaration.
pin_consistency . || fail=1

# (5) .PHONY-Vollständigkeit (slice-068, Fund F-1): jedes eigene Rezept-Target
#     ist deklariert, sonst meldet make Exit 0 ohne das Rezept auszuführen.
check_phony_complete Makefile || fail=1

# (6) ADR-Index-Vollstaendigkeit (slice-087): jede ADR-Datei ist im Index
#     verlinkt — die Gegenrichtung, die doc-check per Konstruktion nicht sieht.
check_adr_index docs/plan/adr docs/plan/adr/README.md || fail=1

if [ "$fail" -ne 0 ]; then
  exit 1
fi
echo "gate-consistency ok: Doku ↔ Makefile konsistent, .d-check.yml-Module intakt, Pins konsistent, .PHONY vollstaendig, ADR-Index vollstaendig (Selbsttests gefeuert)."
