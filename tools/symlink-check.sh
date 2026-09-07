#!/usr/bin/env bash
# symlink-check.sh — jeder getrackte Symlink löst auf.
#
# Antwort auf BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft bei 3×
# (slice-142 · slice-167 · slice-173). Das Muster ist jedes Mal dasselbe: ein
# Baseline-Bump zieht die Markdown-Verweise nach, weil `doc-check` sie meldet —
# die vier Symlinks unter `.claude/rules/` auf einzelne Regelwerk-Module sieht
# kein Gate, weil d-check den ZIELINHALT liest, nicht den Linkpfad. Bei
# slice-167 fiel es dem Maintainer auf, nicht dem Lauf; bei slice-173 wurde es
# gemessen: ein auf `v6.1.0` umgebogener Symlink (Ziel existiert nicht mehr)
# ließ `make doc-check` bei 0 Befunden.
#
# GELTUNGSBEREICH: alle von git getrackten Symlinks, nicht nur `.claude/rules/`
# — die Fehler-Familie ist "Symlink zeigt ins Leere", nicht "Baseline-Symlink
# zeigt ins Leere". Untrackte bleiben außen vor: sie sind lokaler Kram und in
# keinem frischen Klon vorhanden.
#
# NICHT geprüft (ehrliche Grenze, AC-QA-02): ob das Ziel das RICHTIGE ist. Ein
# Symlink auf ein existierendes, aber veraltetes Modul bleibt grün — das wäre
# ein Urteil über Absicht. Erkannt wird nur, was nachweisbar ins Leere zeigt.
set -euo pipefail
cd "$(dirname "$0")/.."

# Die eine Prüf-Funktion: liest Pfade von stdin, gibt die kaputten aus.
# BEIDE Aufrufwege (Repo-Lauf und Selbsttest) gehen hier durch — ein zweiter
# Codepfad für den Test hieße, dass der Test nicht das prüft, was im Gate läuft
# (BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf).
# `-e` folgt dem Symlink; fehlt das Ziel, ist die Bedingung falsch.
kaputte() {
  local l
  while IFS= read -r l; do
    [ -e "$l" ] || printf '%s\n' "$l"
  done
}

# Selbsttest: je eine Fixture pro Richtung muss feuern bzw. schweigen. Ohne ihn
# wäre eine leere Eingabe von einem sauberen Bestand nicht zu unterscheiden.
selftest() {
  local tmp; tmp="$(mktemp -d)"
  trap 'rm -rf "$tmp"' RETURN
  : > "$tmp/ziel"
  ln -s "ziel"      "$tmp/heil"
  ln -s "gibt-es-nicht" "$tmp/kaputt"

  local got
  got="$(printf '%s\n' "$tmp/heil" "$tmp/kaputt" | kaputte || true)"
  if [ "$got" != "$tmp/kaputt" ]; then
    echo "symlink-check: Selbsttest FEHLGESCHLAGEN — erwartet genau '$tmp/kaputt', bekam: ${got:-<leer>}" >&2
    return 1
  fi
}

selftest

symlinks="$(git ls-files -s | awk '$1=="120000"{ $1=""; $2=""; $3=""; sub(/^ +/,""); print }')"
anzahl="$(printf '%s' "$symlinks" | grep -c . || true)"

if [ "$anzahl" -eq 0 ]; then
  echo "symlink-check ok: kein getrackter Symlink im Repo (Selbsttest gefeuert)."
  exit 0
fi

kaputt="$(printf '%s\n' "$symlinks" | kaputte || true)"
if [ -n "$kaputt" ]; then
  echo "symlink-check FAIL — Symlink zeigt ins Leere:" >&2
  printf '%s\n' "$kaputt" | sed 's/^/  /' >&2
  echo "  Ursache meist: ein Pfad-Bestandteil wurde umbenannt oder entfernt (z. B. ein" >&2
  echo "  Baseline-Stand), und der Symlink reiste nicht mit. doc-check sieht ihn nicht." >&2
  exit 1
fi

echo "symlink-check ok: $anzahl getrackte Symlink(s) loesen auf (Selbsttest gefeuert)."
echo "  NICHT geprueft: ob das Ziel das richtige ist — nur, dass es existiert."
