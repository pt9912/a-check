#!/usr/bin/env bash
# suppression-check.sh — Fitness Function für die Hard Rule AGENTS.md §3.2
# (Suppression-Verbot), slice-049 / Fund B-11 aus slice-048.
#
# Warum ein eigenes Skript und nicht `nolintlint`: der Linter prüft die
# *Wohlgeformtheit* von Direktiven, nicht deren Existenz. Empirisch belegt in
# slice-049: ein wohlgeformtes, wirksames `//nolint:gochecknoglobals // Why: …`
# unterdrückt einen echten Verstoß und lässt `make lint` grün. Die Regel
# "Inline-Suppressions sind verboten; Ausnahmen leben zentral in .golangci.yml"
# (ADR-0005) braucht darum einen Sensor, der jede Direktive ablehnt.
#
# Erfasst beide in Go gebräuchlichen Formen:
#   //nolint            (golangci-lint, mit und ohne :linter und //-Begründung)
#   //lint:ignore       (staticcheck-Altform)
# Bewusste Grenze (ehrliche Heuristik, AC-QA-02): geprüft wird der Quelltext der
# .go-Dateien dieses Repos, nicht generierter oder vendored Fremdcode.
set -euo pipefail
cd "$(dirname "$0")/.."

# Dass dieses Skript und die erklärenden Kommentare in .golangci.yml/AGENTS.md
# nicht selbst getroffen werden, folgt allein daraus, dass nur *.go gescannt wird
# — nicht aus dem Muster (Korrektur aus dem Review 2026-07-26, R-049-F5).
#
# Geprüft wird der Text nach dem ERSTEN `//` einer Zeile. Nur dort kann eine
# wirksame Direktive stehen: ein `//nolint` weiter hinten liegt bereits INNERHALB
# eines Kommentars und ist für den Compiler wie für golangci-lint bloßer Text.
# Ein zeilenweites Muster meldete darum die Zeile
#   // Hinweis: bitte kein //nolint verwenden.
# als Direktive — real aufgetreten und Anlass dieser Korrektur (R-049-F3).
# EHRLICHE GRENZE (AC-QA-02): die Zeichenfolge in einem String-Literal
# (`s := "//nolint"`) wird weiterhin gemeldet. Fail-closed und selten; ein
# Go-Parser wäre die einzige saubere Antwort und steht in keinem Verhältnis.
# FAIL-CLOSED (slice-069, Fund R-068-F3): frueher endete diese Funktion auf
# `2>/dev/null || true`. Damit wurde auch ein Traversierungsfehler verschluckt —
# eine fehlende oder nicht lesbare Wurzel lieferte null Treffer und Exit 0, also
# dasselbe Ergebnis wie ein sauberer Lauf ohne Fund. `set -o pipefail` (Kopf)
# sorgt dafuer, dass der find-Fehler die Pipe rot macht.
# Keine Treffer bleiben Exit 0: `xargs -r` startet awk gar nicht erst.
# Orte ohne eigene Go-Quellen, deklariert statt stillschweigend: `.git` ist
# Metadaten, `.harness/baseline/` ist committet vendorter Fremdtext (MR-006).
# Die Liste ist der Ort fuer eine kuenftige Ausnahme — leer zu lassen waere
# ebenso eine Aussage.
SCAN_EXCLUDES=( -not -path './.git/*' -not -path './.harness/baseline/*' )

scan() {  # $1 = Wurzelverzeichnis; Treffer auf stdout, !=0 bei Traversierungsfehler
  if [ ! -d "$1" ]; then
    echo "suppression-check: Scan-Wurzel '$1' existiert nicht oder ist kein Verzeichnis" >&2
    return 1
  fi
  find "$1" "${SCAN_EXCLUDES[@]}" -name '*.go' -type f -print0 \
    | xargs -0 -r awk -F'//' \
        'NF > 1 && $2 ~ /^[[:space:]]*(nolint|lint:ignore)/ {
           print FILENAME ":" FNR ":" $0
         }'
}

# Selbsttest (Fitness Function der Fitness Function): eine Fixture mit Direktive
# MUSS gefunden werden, eine ohne MUSS still bleiben. Ohne diesen Beweis wäre ein
# stillschweigend kaputtes Muster ein False-Green — genau die Klasse, die in
# slice-035 real auftrat.
self_test() {
  local tmp
  tmp="$(mktemp -d)"
  mkdir -p "$tmp/pos" "$tmp/neg"
  printf 'package p\n\nvar x = 1 //nolint:gochecknoglobals // Why: Fixture\n' > "$tmp/pos/a.go"
  printf 'package p\n\n//lint:ignore SA1000 Fixture\nvar y = 2\n' > "$tmp/pos/b.go"
  # Negativ-Fixturen. Die zweite ist die eigentliche Probe: sie enthält die
  # Zeichenfolge `//nolint` wörtlich, aber innerhalb eines Kommentars — genau der
  # Fall, den das alte Muster fälschlich meldete. Die erste konnte das Muster nie
  # treffen (kein `//`-Präfix) und belegte darum nichts (R-049-F3).
  printf 'package p\n\n// ein gewoehnlicher Kommentar ueber nolint-Regeln\nvar z = 3\n' > "$tmp/neg/c.go"
  printf 'package p\n\n// Hinweis: bitte kein //nolint verwenden.\nvar w = 4\n' > "$tmp/neg/d.go"

  if [ "$(scan "$tmp/pos" | wc -l)" -ne 2 ]; then
    echo "suppression-check: Selbsttest FEHLGESCHLAGEN — Direktiven nicht erkannt (Muster tot)" >&2
    rm -rf "$tmp"; exit 2
  fi
  if [ -n "$(scan "$tmp/neg")" ]; then
    echo "suppression-check: Selbsttest FEHLGESCHLAGEN — Fehlalarm auf gewoehnlichem Kommentar" >&2
    rm -rf "$tmp"; exit 2
  fi
  # Traversierungsfehler, beide Richtungen (slice-069): eine fehlende Wurzel muss
  # scheitern, eine vorhandene ohne Treffer muss still Exit 0 liefern. Ohne die
  # zweite Haelfte waere ein scan, das immer scheitert, ununterscheidbar.
  if scan "$tmp/gibt-es-nicht" >/dev/null 2>&1; then
    echo "suppression-check: Selbsttest FEHLGESCHLAGEN — fehlende Scan-Wurzel nicht erkannt;" >&2
    echo "  der Sensor wuerde null Treffer und Exit 0 melden (R-068-F3)." >&2
    rm -rf "$tmp"; exit 2
  fi
  if ! scan "$tmp/neg" >/dev/null 2>&1; then
    echo "suppression-check: Selbsttest FEHLGESCHLAGEN — vorhandene Wurzel ohne Treffer faelschlich rot" >&2
    rm -rf "$tmp"; exit 2
  fi
  rm -rf "$tmp"
}

self_test

# Jede Wurzel einzeln pruefen. Frueher stand hier
# `scan ./internal; scan ./cmd 2>/dev/null || true` — eine Substitution, deren
# Exit-Code nur vom LETZTEN Befehl stammt: ein Fehler beim Scan von `internal`
# war strukturell unsichtbar (slice-069).
# Wurzelmenge ABGELEITET statt aufgezaehlt (slice-071, Fund F-2; verschaerft
# durch R-068-F3). Frueher stand hier `./internal ./cmd` — eine fest verdrahtete
# Liste, waehrend Target und Hard Rule AGENTS.md §3.2 von den "Go-Quellen des
# Repos" sprechen. Jede `.go`-Datei ausserhalb der beiden Baeume lief durch, und
# eine zusaetzliche Wurzel einzutragen haette die Luecke nur verschoben.
hits=""
for root in .; do
  if ! out="$(scan "$root")"; then
    echo "suppression-check: FAIL — Scan der Wurzel '$root' fehlgeschlagen; das Ergebnis waere unvollstaendig." >&2
    exit 1
  fi
  hits="${hits}${out}"
done
if [ -n "$hits" ]; then
  echo "suppression-check: FAIL — Inline-Suppression verboten (AGENTS.md §3.2, ADR-0005)." >&2
  echo "Ausnahmen gehoeren zentral in .golangci.yml unter 'exclusions' mit Why:-Kommentar." >&2
  printf '%s\n' "$hits" >&2
  exit 1
fi

echo "suppression-check ok: keine Inline-Suppression in den Go-Quellen (Selbsttest gefeuert)."
