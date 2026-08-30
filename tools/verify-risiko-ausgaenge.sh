#!/usr/bin/env bash
# verify-risiko-ausgaenge.sh — der Teil der Closure-Pflicht, der lokal bleibt
# (AGENTS.md §5), hervorgegangen aus verify-closure-notes.sh (slice-050) bei
# dessen Ablösung durch das Modul `structure` (slice-080).
#
# WAS HIER BLEIBT UND WARUM. Jedes in §6 notierte Risiko traegt genau einen
# Ausgang aus der geschlossenen Dreier-Menge (slice-102). Diese Pruefung ist
# ABSCHNITTS-UEBERGREIFEND: sie liest den Risiko-Block und misst ihn an der
# Closure-Notiz. `structure` prueft abschnitts-LOKAL — eine Bedingung gilt
# innerhalb eines Abschnitts, nicht zwischen zweien. Darum ist dies kein Rest,
# den man noch wegkonfigurieren koennte, sondern eine andere Frage.
#
# WAS INS MODUL GEWANDERT IST (.d-check.yml, Regel (2), `make doc-structure`):
# genau ein Closure-Abschnitt, nicht leer, kein Platzhalter, keine Floskel,
# mindestens zwei Saetze. Paritaet ist je Befundklasse gegen dieses Skript
# gemessen worden, bevor sie entfiel (slice-080).
#
# Verifikations-Schicht, nicht Gate-Schicht: `make gates` beantwortet
# Code-/Architektur-Fragen, `make verify` DoD-/Closure-Fragen (Regelwerk
# Modul 11).
#
# NICHT geprueft (ehrliche Grenze, AC-QA-02): die EXISTENZ eines Risiko-Blocks.
# modul-05 bindet die Pflicht an *notierte* Risiken — wer den Block weglaesst,
# wird nicht erwischt. Ebenso wenig, ob der eingetragene Ausgang TRAEGT; das ist
# Urteil (modul-06: "Mensch urteilt, Maschine prueft Deckung").
set -euo pipefail
cd "$(dirname "$0")/.."

DONE_DIR="docs/plan/planning/done"
# Stichtag (slice-102, Baseline v5.12.0 modul-05 §Offene Risiken werden bei
# Closure aufgeloest). Aeltere Notizen entstanden, BEVOR die geschlossene
# Dreier-Menge galt — slice-092 fuehrt etwa "Maintainer-Entscheidung" als
# Ausgang, damals korrekt. Rueckwirkend umzuschreiben waere Geschichts-Politur.
RISK_FROM=102
# Die drei Ausgaenge als FORM, nicht als Inhalt: eingetreten -> Carveout oder
# Folge-Slice mit ID · entfallen -> gestrichen mit Begruendung · weiter offen ->
# Beobachtungs-Register.
AUSGAENGE='(Carveout|Folge-Slice|gestrichen mit Begründung|Beobachtungs-Register|BEO-[0-9]{3})'

slice_num() {  # $1 = Pfad -> Nummer ohne fuehrende Nullen, leer wenn keine
  basename "$1" | sed -nE 's/^slice-0*([0-9]+)-.*/\1/p'
}

# Risiko-Block als EIGENE Sektion (Ziel-Form v5.12.0: "## 6. Risiken und
# offene Punkte"). Die alte a-check-Form fuehrt ihn stattdessen als Block IN
# der Closure-Notiz; beide werden gelesen, damit der Gliederungs-Wechsel
# (slice-107) den Check nicht still abschaltet.
risiko_sektion() {  # $1 = Datei
  awk '
    /^## .*[Rr]isiken/ { inblock=1; next }
    inblock && /^## / { inblock=0 }
    inblock { print }
  ' "$1"
}

closure_body() {  # $1 = Datei
  awk '
    /^## .*[Cc]losure/ && !/[Cc]losure-(Trigger|Kriterien)/ { inblock=1; next }
    inblock && /^## / { inblock=0 }
    inblock { print }
  ' "$1"
}

# Risiko-Block INNERHALB der Closure-Notiz (a-checks Form bis slice-106).
risiko_aus_closure() {  # Closure-Rumpf auf stdin
  awk '
    /^\*\*Offene Risiken/ { inblock=1; next }
    inblock && /^\*\*/ && !/^\*\*Offene Risiken/ { inblock=0 }
    inblock { print }
  '
}

# Aufzaehlungspunkte zusammenkleben: ein Ausgang steht oft erst in der zweiten
# Zeile. Fortsetzungszeilen NORMALISIERT ankleben (Einrueckung raus), sonst
# zerreisst ein Umbruch die Wendung — real aufgetreten an slice-102.
punkte_kleben() {  # Zeilen auf stdin
  awk '
    {
      line = $0
      sub(/^[[:space:]]+/, "", line)
      if ($0 ~ /^- /) { if (buf != "") print buf; buf = line }
      else if (buf != "") { buf = buf " " line }
    }
    END { if (buf != "") print buf }
  '
}

check_file() {  # $1 = Datei; gibt Befunde auf stdout, Rueckgabe 1 bei Befund
  local f="$1" fail=0 num punkt body
  num="$(slice_num "$f")"
  [ -n "$num" ] || return 0
  [ "$num" -ge "$RISK_FROM" ] || return 0
  body="$(closure_body "$f")"
  while IFS= read -r punkt; do
    [ -n "$punkt" ] || continue
    if ! printf '%s' "$punkt" | grep -qE "Ausgang:.*$AUSGAENGE"; then
      echo "$f: Risiko ohne Ausgang aus der geschlossenen Dreier-Menge: ${punkt:0:70}…"
      fail=1
    fi
  done <<< "$( { printf '%s\n' "$body" | risiko_aus_closure; risiko_sektion "$f"; } | punkte_kleben )"
  return "$fail"
}

# Selbsttest: je eine Fixture pro Richtung muss feuern bzw. schweigen. Ohne ihn
# waere ein totes awk-Muster ein False-Green.
self_test() {
  local tmp good bad
  tmp="$(mktemp -d)"
  # DREI Richtungen (slice-102). Ohne die dritte Fixture waere ein Muster, das
  # alles durchlaesst, von einem korrekten nicht zu unterscheiden — und ohne die
  # zweite eines, das nur "Ausgang:" sucht.
  { echo '## 6. Closure-Notiz'; echo 'Text hier. Zweiter Satz.'; echo '';
    echo '**Offene Risiken und ihr Ausgang:**'; echo '';
    echo '- *A* — Ausgang: **Folge-Slice**, slice-900.';
    echo '- *B* — Ausgang: **gestrichen mit Begründung**, weil weg.';
    echo '- *C* — Ausgang: **weiter offen**, fürs Beobachtungs-Register.'; } > "$tmp/slice-102-gut.md"
  # Ausgang ueber einen Zeilenumbruch hinweg — die Probe fuer das Ankleben.
  { echo '## 6. Closure-Notiz'; echo 'Text hier. Zweiter Satz.'; echo '';
    echo '**Offene Risiken und ihr Ausgang:**'; echo '';
    echo '- *A* — Ausgang: **gestrichen mit'; echo '  Begründung**, weil weg.'; } > "$tmp/slice-102-umbruch.md"
  # NEUE GLIEDERUNG (slice-107): Risiken als eigene Sektion, nicht in der Closure.
  { echo '## 6. Risiken und offene Punkte'; echo '';
    echo '- *A* — **Ausgang:** weiter offen, `BEO-001` im Register.'; echo '';
    echo '## 7. Closure-Notiz'; echo 'Text hier. Zweiter Satz.'; } > "$tmp/slice-102-sektion-gut.md"
  for good in slice-102-gut slice-102-umbruch slice-102-sektion-gut; do
    if ! check_file "$tmp/$good.md" >/dev/null; then
      echo "verify-risiko-ausgaenge: Selbsttest FEHLGESCHLAGEN — '$good' faelschlich beanstandet" >&2
      rm -rf "$tmp"; exit 2
    fi
  done
  { echo '## 6. Risiken und offene Punkte'; echo '';
    echo '- *A* — bleibt halt so.'; echo '';
    echo '## 7. Closure-Notiz'; echo 'Text hier. Zweiter Satz.'; } > "$tmp/slice-102-sektion-ohne.md"
  { echo '## 6. Closure-Notiz'; echo 'Text hier. Zweiter Satz.'; echo '';
    echo '**Offene Risiken und ihr Ausgang:**'; echo '';
    echo '- *A* — bleibt halt so.'; } > "$tmp/slice-102-ohne.md"
  { echo '## 6. Closure-Notiz'; echo 'Text hier. Zweiter Satz.'; echo '';
    echo '**Offene Risiken und ihr Ausgang:**'; echo '';
    echo '- *A* — Ausgang: **Maintainer-Entscheidung**.'; } > "$tmp/slice-102-fremd.md"
  for bad in slice-102-sektion-ohne slice-102-ohne slice-102-fremd; do
    if check_file "$tmp/$bad.md" >/dev/null; then
      echo "verify-risiko-ausgaenge: Selbsttest FEHLGESCHLAGEN — '$bad' nicht erkannt (Pruefung tot)" >&2
      rm -rf "$tmp"; exit 2
    fi
  done
  # Grandfathering: derselbe Block unter dem Stichtag muss schweigen.
  cp "$tmp/slice-102-ohne.md" "$tmp/slice-091-alt.md"
  if ! check_file "$tmp/slice-091-alt.md" >/dev/null; then
    echo "verify-risiko-ausgaenge: Selbsttest FEHLGESCHLAGEN — Regel greift trotz Grandfathering" >&2
    rm -rf "$tmp"; exit 2
  fi
  rm -rf "$tmp"
}

self_test

fail=0
count=0
for f in "$DONE_DIR"/slice-*.md; do
  [ -e "$f" ] || continue
  count=$((count + 1))
  if ! out="$(check_file "$f")"; then
    printf '%s\n' "$out" >&2
    fail=1
  elif [ -n "$out" ]; then
    printf '%s\n' "$out" >&2
    fail=1
  fi
done

if [ "$fail" -ne 0 ]; then
  echo "verify-risiko-ausgaenge: FAIL — Risiko ohne Ausgang (AGENTS.md §5)." >&2
  exit 1
fi

# ERWARTETE GRUNDGESAMTHEIT (slice-070, Fund F-12): > 0.
# `done/` ist der Endzustand jedes je abgeschlossenen Slice — dort null Dateien
# zu finden heisst nicht "nichts zu pruefen", sondern "der Pruefbestand ist
# verschwunden". Ohne diese Grenze meldete der Sensor "ok: 0 Slice(s)" mit
# Exit 0 und war damit von einem bestandenen Lauf nicht zu unterscheiden.
if [ "$count" -eq 0 ]; then
  echo "verify-risiko-ausgaenge: FAIL — keine Slice-Datei in $DONE_DIR gefunden." >&2
  echo "  Erwartet wird ein nicht leerer done/-Bestand; null Dateien sind Bestandsverlust," >&2
  echo "  nicht 'nichts zu pruefen'." >&2
  exit 1
fi
echo "verify-risiko-ausgaenge ok: $count Slice(s) in done/, jedes notierte Risiko mit Ausgang (Selbsttest gefeuert)."
