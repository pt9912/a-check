#!/usr/bin/env bash
# symlink-check.sh — jeder getrackte Symlink loest auf, und ein Symlink in die
# vendored Baseline zeigt auf den ADOPTIERTEN Stand.
#
# Antwort auf BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft bei 3x
# (slice-142 · slice-167 · slice-173). `.claude/rules/` traegt vier Symlinks auf
# einzelne Regelwerk-Module; d-check sieht sie nicht, weil es bei einem Symlink
# den ZIELINHALT liest und dort kein Linkpfad steht.
#
# ZWEI Pruefungen, weil die drei gezaehlten Instanzen zwei verschiedene Formen
# haben — das war der Fund des Reviews zu slice-173 (F-2):
#   (1) Ziel existiert nicht          — der Stand wurde entfernt (slice-173 gemessen)
#   (2) Ziel zeigt auf einen anderen  — der alte Stand liegt noch daneben, der
#       als den adoptierten             Symlink loest auf und ist trotzdem falsch
#                                       (slice-167: Symlinks auf v6.0.0, das
#                                       waehrend der Migration vendored blieb)
# Ein Sensor mit nur (1) haette slice-167 gruen gemeldet und die Beobachtung
# faelschlich als verkoerpert ausgewiesen.
#
# GELTUNGSBEREICH: alle von git getrackten Symlinks. Pruefung (1) gilt fuer
# jeden, (2) nur fuer Ziele unter `.harness/baseline/<tag>/`.
#
# NICHT geprueft (ehrliche Grenze, AC-QA-02): ob ein Symlink AUSSERHALB der
# Baseline auf das inhaltlich richtige Ziel zeigt. Das waere ein Urteil ueber
# Absicht. Und: der adoptierte Stand wird aus einer Prosa-Zeile gelesen (siehe
# unten) — dieselbe Kopplung, die BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf
# fuer `versions.current-from` registriert; sie bricht laut, nicht still.
set -euo pipefail
cd "$(dirname "$0")/.."

KONVENTIONEN="${KONVENTIONEN:-harness/conventions.md}"

# Der adoptierte Stand steht EINMAL im Adaptions-Block — dieselbe Quelle, die
# `versions.current-from` in .d-check.yml liest. Fail-closed: ohne lesbaren
# Stand wird nicht geraten, sondern abgebrochen.
adoptierter_stand() {
  sed -n '/^## Baseline/,/^## Adoptierte/p' "$KONVENTIONEN" \
    | sed -n 's/^- \*\*Stand:\*\* \[`\(v[0-9][^`]*\)`\].*/\1/p' | head -1
}

# Die eine Pruef-Funktion: liest NUL-getrennte Pfade von stdin, gibt je Befund
# eine Zeile "pfad<TAB>grund" aus. BEIDE Aufrufwege (Gate-Lauf und Selbsttest)
# gehen hier durch — ein zweiter Codepfad fuer den Test hiesse, dass der Test
# nicht das prueft, was im Gate laeuft (BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf).
#
# NUL-Trennung statt Feld-Zerlegung: `git ls-files -s | awk '{$1="";…}'` kollabierte
# mehrfache Leerzeichen im Dateinamen und stolperte ueber C-quotierte Nicht-ASCII-
# Pfade (Review slice-173, F-3). `-z` liefert rohe Pfade ohne Quoting.
kaputte() {
  local p ziel tag
  while IFS= read -r -d '' p; do
    [ -L "$p" ] || continue
    if [ ! -e "$p" ]; then
      printf '%s\tZiel existiert nicht\n' "$p"
      continue
    fi
    ziel="$(readlink "$p")"
    case "$ziel" in
      *.harness/baseline/*)
        tag="${ziel#*.harness/baseline/}"
        tag="${tag%%/*}"
        if [ "$tag" != "$STAND" ]; then
          printf '%s\tzeigt auf Baseline %s, adoptiert ist %s\n' "$p" "$tag" "$STAND"
        fi
        ;;
    esac
  done
}

# Selbsttest: vier Kontrollen, zwei je Pruefung. Ohne ihn waere eine leere
# Eingabe von einem sauberen Bestand nicht zu unterscheiden.
selftest() {
  local tmp; tmp="$(mktemp -d)"
  trap 'rm -rf "$tmp"' RETURN
  mkdir -p "$tmp/.harness/baseline/$STAND/regelwerk" "$tmp/.harness/baseline/v0.0.1/regelwerk"
  : > "$tmp/.harness/baseline/$STAND/regelwerk/m.md"
  : > "$tmp/.harness/baseline/v0.0.1/regelwerk/m.md"
  : > "$tmp/ziel"
  ln -s "ziel"                                              "$tmp/heil"
  ln -s "gibt-es-nicht"                                     "$tmp/kaputt"
  ln -s ".harness/baseline/$STAND/regelwerk/m.md"           "$tmp/aktuell"
  ln -s ".harness/baseline/v0.0.1/regelwerk/m.md"           "$tmp/veraltet"
  ln -s "ziel"                                              "$tmp/mit  zwei leerzeichen"

  local got erwartet
  got="$( cd "$tmp" && printf '%s\0' heil kaputt aktuell veraltet "mit  zwei leerzeichen" \
          | { STAND="$STAND"; kaputte; } | sort )"
  erwartet="$(printf '%s\n' \
      "kaputt	Ziel existiert nicht" \
      "veraltet	zeigt auf Baseline v0.0.1, adoptiert ist $STAND" | sort)"
  if [ "$got" != "$erwartet" ]; then
    echo "symlink-check: Selbsttest FEHLGESCHLAGEN" >&2
    echo "  erwartet:" >&2; printf '%s\n' "$erwartet" | sed 's/^/    /' >&2
    echo "  bekommen:" >&2; printf '%s\n' "${got:-<leer>}" | sed 's/^/    /' >&2
    return 1
  fi
}

STAND="$(adoptierter_stand)"
if [ -z "$STAND" ]; then
  echo "symlink-check FAIL — adoptierter Baseline-Stand nicht lesbar aus $KONVENTIONEN §Baseline." >&2
  echo "  Erwartete Form: '- **Stand:** [\`vX.Y.Z\`](…)'. Ohne ihn wird nicht geraten." >&2
  exit 1
fi

export -f kaputte 2>/dev/null || true
selftest

anzahl="$(git ls-files -z | { c=0; while IFS= read -r -d '' p; do [ -L "$p" ] && c=$((c+1)); done; echo "$c"; })"
befunde="$(git ls-files -z | kaputte || true)"

if [ -n "$befunde" ]; then
  echo "symlink-check FAIL — Symlink-Ziel stimmt nicht:" >&2
  printf '%s\n' "$befunde" | sed 's/^/  /' >&2
  echo "  Ursache meist: ein Baseline-Bump zog die Markdown-Verweise nach, den Symlink nicht." >&2
  echo "  doc-check sieht ihn nicht — es liest bei einem Symlink den Zielinhalt." >&2
  exit 1
fi

echo "symlink-check ok: $anzahl getrackte Symlink(s); Ziele existieren, Baseline-Ziele auf $STAND (Selbsttest gefeuert)."
echo "  NICHT geprueft: ob ein Symlink ausserhalb der Baseline inhaltlich das richtige Ziel hat."
