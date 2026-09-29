#!/usr/bin/env bash
# verify-trigger-audit.sh — die „dass"-Haelfte des Trigger-Audits fuer MR-Einträge
# (AGENTS.md §5; hervorgegangen aus BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter,
# 3x: slice-170, slice-197, slice-206).
#
# WAS HIER GILT. Jede abschlussbereite Closure-Notiz traegt eine Zeile
# "Trigger-Audit der aktiven MR:" mit der Kennung JEDES aktiven MR-Eintrags —
# den Beleg, dass die Closure die aktiven Eintraege aus
# harness/conventions.md angesehen hat. Das Feld "Aufloesungs-Trigger" selbst
# ist Prosa und maschinell nicht auswertbar; der Sensor prueft die SICHTUNG
# (Form: Kennung je aktivem MR), nicht das Urteil.
#
# GELTUNGSBEREICH (slice-208, Review F-1: Pruefer ohne Gegenstand). Gescannt
# wird dort, wo der Beleg-Moment liegt:
#   - in-progress/: die abschlussbereite Closure (vor dem `git mv`) — die
#     Notiz mit Platzhalter wird uebersprungen und gemeldet;
#   - done/: die flachen Closures zwischen `git mv` und Archivierung.
# Archivierte Stubs unter done/wellenlos/ und done/welle-*/ tragen per
# Ziel-Form keine Closure-Abschnitte mehr — ihr Beleg-Moment lag in done/
# bzw. in-progress/ und ist durchlaufen; die Stub-Form kann ihn nicht mehr
# tragen (Sensor-Datei, Grenze 3).
#
# NICHT geprueft (ehrliche Grenze, AC-QA-02):
# - Das URTEIL. Ob "0 offen" stimmt oder die genannten MR-IDs wirklich
#   abgelaufen sind, ist Prosa-Auswertung; der Sensor prueft die Form.
# - Slices VOR slice-208 sind grandfathered — ihre Closures entstanden, bevor
#   die Zusage stand (dieselbe Grenze wie verify-risiko-ausgaenge RISK_FROM).
# - Die KENNUNG ist Form: die Zeile nennt jeden aktiven MR; ob die Sichtung
#   stattfand, bleibt Urteil (Grenze 1 der Sensor-Datei; BEO-HARNESS/
#   trigger-audit-ceremonie, 1x).
set -euo pipefail
cd "$(dirname "$0")/.."

PROGRESS_DIR="docs/plan/planning/in-progress"
DONE_DIR="docs/plan/planning/done"
CONVENTIONS="harness/conventions.md"
AUDIT_FROM=208
PHRASE="Trigger-Audit der aktiven MR:"

aktive_mr() {  # aktive MR-Nummern aus den ZEILEN der Aktiven-Tabelle
  sed -n '/^### Aktive Adaptionen/,/^### Aufgelöste Adaptionen/p' "$CONVENTIONS" \
    | grep -E '^\| \[MR-' | grep -oE '^\| \[MR-[0-9]{3}' | grep -oE 'MR-[0-9]{3}' \
    | sort -u || true
}

notiz_offen() {  # $1 = Datei; 0 = Platzhalter (in Arbeit), 1 = ausgefuellt
  grep -q 'wird vor dem `git mv` nach `done/` gefüllt' "$1"
}

check_file() {  # $1 = Datei; Befunde auf stdout, Rueckgabe 1 bei Befund
  local f="$1" num line id fehlt=""
  num="$(basename "$f" | sed -nE 's/^slice-0*([0-9]+)-.*/\1/p')"
  [ -n "$num" ] || return 0
  [ "$num" -ge "$AUDIT_FROM" ] || return 0
  line="$(grep -F "$PHRASE" "$f" | head -1 || true)"
  if [ -z "$line" ]; then
    echo "$f: Closure-Notiz ohne $PHRASE-Zeile — die Sichtung der aktiven MR-Eintraege ist nicht belegt"
    return 1
  fi
  while IFS= read -r id; do
    [ -n "$id" ] || continue
    case "$line" in
      *"$id"*) : ;;
      *) fehlt="$fehlt $id" ;;
    esac
  done < <(aktive_mr)
  if [ -n "$fehlt" ]; then
    echo "$f: Audit-Zeile nennt die aktiven MR-Kennungen nicht:$fehlt"
    return 1
  fi
  return 0
}

# Der Lauf erreicht beide Beleg-Momente: die abschlussbereite Closure in
# in-progress/ (Platzhalter uebersprungen und gemeldet) und die flachen
# Closures in done/. Archivierte Stubs tragen keine Closure-Abschnitte.
run() {
  local fail=0 count=0 offen=0 f
  for f in "$PROGRESS_DIR"/slice-*.md; do
    [ -e "$f" ] || continue
    if notiz_offen "$f"; then
      offen=$((offen + 1))
      continue
    fi
    count=$((count + 1))
    if ! out="$(check_file "$f")"; then
      printf '%s\n' "$out" >&2
      fail=1
    elif [ -n "$out" ]; then
      printf '%s\n' "$out" >&2
      fail=1
    fi
  done
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
    echo "verify-trigger-audit: FAIL — Closure ohne gueltige Trigger-Audit-Zeile (AGENTS.md §5; BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter)." >&2
    return 1
  fi
  local aktive
  aktive="$(aktive_mr | tr '\n' ' ')"
  # Die Leere Aktiven-Menge ist eine MELDUNG, kein stiller Durchlauf: sie
  # sagt, dass kein MR-Eintrag mehr auf eine Sichtung wartet.
  if [ -z "$(aktive_mr)" ]; then
    printf "verify-trigger-audit ok: %s Closure(s) geprueft — aktive MR-Menge LEER, nichts wartet auf eine Sichtung" "$count"
  else
    printf "verify-trigger-audit ok: %s Closure(s) geprueft — aktive MR: %s" "$count" "$aktive"
  fi
  if [ "$offen" -gt 0 ]; then
    printf ", %s in Arbeit uebersprungen" "$offen"
  fi
  echo " (Selbsttest gefeuert)."
  return 0
}

# Selbsttest: je eine Fixture pro Richtung muss feuern, die guten schweigen
# (Modul 11: die rote Haelfte zuerst). Die Kennungs-Probe praefiert die
# "mit Kennung"-Haelfte: eine Zeile, die MR-030 nicht nennt, ist rot, auch
# wenn sie die Phrase traegt.
self_test() {
  local tmp PD_SAVE DD_SAVE CONV_SAVE
  tmp="$(mktemp -d)"
  PD_SAVE="$PROGRESS_DIR"; DD_SAVE="$DONE_DIR"; CONV_SAVE="$CONVENTIONS"

  mkdir -p "$tmp/repo/$PROGRESS_DIR" "$tmp/repo/$DONE_DIR" "$tmp/repo/harness"
  cat > "$tmp/repo/$CONVENTIONS" << 'EOF'
### Aktive Adaptionen
| [MR-030](conventions/MR-030-erfassung-lokal-kein-abfluss.md) <a id="mr-030"></a> | Beispiel | x | y |
| [MR-029](conventions/MR-029-id-schema-deklaration-gesamt.md) <a id="mr-029"></a> | Beispiel | x | y |
### Aufgelöste Adaptionen
| [MR-014](conventions/done/MR-014-keine-agenten-telemetrie.md) <a id="mr-014"></a> | aufgeloest | x | y |
EOF

  cat > "$tmp/repo/$PROGRESS_DIR/slice-990-rot.md" << 'EOF'
# slice-990
## 7. Closure-Notiz
Trigger-Audit der aktiven MR: MR-029 0 offen.
EOF
  cat > "$tmp/repo/$PROGRESS_DIR/slice-991-gruen.md" << 'EOF'
# slice-991
## 7. Closure-Notiz
Trigger-Audit der aktiven MR: MR-029 MR-030 0 offen.
EOF
  cat > "$tmp/repo/$PROGRESS_DIR/slice-992-offen.md" << 'EOF'
# slice-992
## 7. Closure-Notiz
*(wird vor dem `git mv` nach `done/` gefüllt)*
EOF

  PROGRESS_DIR="$tmp/repo/$PROGRESS_DIR"
  DONE_DIR="$tmp/repo/$DONE_DIR"
  CONVENTIONS="$tmp/repo/$CONVENTIONS"

  # ROT: Audit-Zeile ohne die Kennung MR-030.
  if check_file "$PROGRESS_DIR/slice-990-rot.md" >/dev/null 2>&1; then
    echo "verify-trigger-audit: Selbsttest FEHLGESCHLAGEN — Audit-Zeile ohne Kennung blieb gruen" >&2
    PROGRESS_DIR="$PD_SAVE"; DONE_DIR="$DD_SAVE"; CONVENTIONS="$CONV_SAVE"; rm -rf "$tmp"; exit 2
  fi
  echo "selftest: Audit-Zeile ohne aktive Kennung -> rot (gefiret)"

  # GRUEN: Zeile traegt beide Kennungen.
  if ! check_file "$PROGRESS_DIR/slice-991-gruen.md" >/dev/null 2>&1; then
    echo "verify-trigger-audit: Selbsttest FEHLGESCHLAGEN — vollstaendige Audit-Zeile meldete rot" >&2
    PROGRESS_DIR="$PD_SAVE"; DONE_DIR="$DD_SAVE"; CONVENTIONS="$CONV_SAVE"; rm -rf "$tmp"; exit 2
  fi

  # GRUEN: Platzhalter-Closure ist in Arbeit (gemeldet, nicht geprueft).
  # Der rot-Fixture-Case ist zuvor zu entfernen — run() prueft alles
  # verbleibende in in-progress/.
  rm "$PROGRESS_DIR/slice-990-rot.md"
  if ! run >/dev/null 2>&1; then
    echo "verify-trigger-audit: Selbsttest FEHLGESCHLAGEN — Platzhalter-Closure meldete rot" >&2
    PROGRESS_DIR="$PD_SAVE"; DONE_DIR="$DD_SAVE"; CONVENTIONS="$CONV_SAVE"; rm -rf "$tmp"; exit 2
  fi
  echo "selftest: Platzhalter-Closure -> uebersprungen und gemeldet"

  # GRUEN: leere Aktiven-Menge — gemeldet, nicht still.
  : > "$CONVENTIONS"
  if ! run >/dev/null 2>&1; then
    echo "verify-trigger-audit: Selbsttest FEHLGESCHLAGEN — leere Aktiven-Menge meldete rot" >&2
    PROGRESS_DIR="$PD_SAVE"; DONE_DIR="$DD_SAVE"; CONVENTIONS="$CONV_SAVE"; rm -rf "$tmp"; exit 2
  fi

  PROGRESS_DIR="$PD_SAVE"; DONE_DIR="$DD_SAVE"; CONVENTIONS="$CONV_SAVE"; rm -rf "$tmp"
}

self_test

run
