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
scan() {  # $1 = Wurzelverzeichnis
  find "$1" -name '*.go' -type f -print0 \
    | xargs -0 -r awk -F'//' \
        'NF > 1 && $2 ~ /^[[:space:]]*(nolint|lint:ignore)/ {
           print FILENAME ":" FNR ":" $0
         }' 2>/dev/null || true
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
  rm -rf "$tmp"
}

self_test

hits="$(scan ./internal; scan ./cmd 2>/dev/null || true)"
if [ -n "$hits" ]; then
  echo "suppression-check: FAIL — Inline-Suppression verboten (AGENTS.md §3.2, ADR-0005)." >&2
  echo "Ausnahmen gehoeren zentral in .golangci.yml unter 'exclusions' mit Why:-Kommentar." >&2
  printf '%s\n' "$hits" >&2
  exit 1
fi

echo "suppression-check ok: keine Inline-Suppression in den Go-Quellen (Selbsttest gefeuert)."
