#!/usr/bin/env bash
# verify-observations.sh — Deckung des Beobachtungs-Registers (slice-102).
#
# Verifikations-Schicht, nicht Gate-Schicht: `make gates` beantwortet Code-/
# Architektur-Fragen, `make verify` DoD-/Closure-Fragen (Regelwerk Modul 11).
# Darum haengt dieses Skript an `verify`, nicht an `gates`.
#
# Regelwerk `modul-06-roadmap.md` §Das Beobachtungs-Register teilt die Arbeit
# ausdruecklich: "Mensch urteilt, Maschine prueft Deckung." Das Urteil *ist das
# dieselbe Beobachtung?* faellt beim Schreiben. Maschinell entscheidbar ist nur:
#
#   (1) jede in done/ zitierte BEO-<NNN> hat eine Registerzeile
#   (2) jede Registerzeile traegt mindestens einen Beleg
#   (3) Beleg-FORM: Slice-Kennung slice-<NNN>, kein Freitext
#   (4) Beleg-ANZAHL: so viele, wie der Zaehler nennt
#
# NICHT geprueft, und beides mit Grund:
#
#   - Die UMKEHRUNG von (1) — "jede Zeile ist irgendwo zitiert". Die
#     allermeisten Zeilen stehen unter der Schwelle und sind nirgends zitiert;
#     eine solche Pruefung liefe auf jedem gefuellten Register rot.
#   - Die LAGE des Belegs (liegt die Slice-Datei in done/). Modul 06 erlaubt sie
#     erst NACH dem `git mv`; auf dem Schreib-Commit laege die Datei noch nicht
#     dort, und der Sensor meldete bei jeder korrekten Closure rot. Dieselbe
#     Reihenfolge-Falle steht als BEO-006 im Register. Ebenso bleibt die
#     EXISTENZ der Datei ungeprueft — ein Repo darf Slices fuehren, die es nicht
#     als Plan-Datei ablegt. Ein erfundenes `slice-999` bleibt damit unentdeckt;
#     das ist die Grenze der Deklaration und gehoert benannt, nicht versteckt.
set -euo pipefail
cd "$(dirname "$0")/.."

REGISTER="docs/plan/planning/observations.md"
DONE_DIR="docs/plan/planning/done"

# Registerzeilen der AKTIVEN Tabelle: `| BEO-NNN | Beobachtung | Sub-Area |
# N× | Belege | Stand |`. Die Sektion *Gestrichene Eintraege* hat eine ANDERE
# Spaltenzahl (Kennung | Beobachtung | Gestrichen am | Warum) — wer beide
# zusammen liest, sucht dort eine Beleg-Spalte, die es nicht gibt, und meldet
# jede gestrichene Zeile als beleglos. Real aufgetreten beim ersten Gebrauch
# der Tabelle (slice-110).
register_rows() {
  awk '/^## Gestrichene/{exit} /^\| BEO-[0-9]{3} \|/{print}' "$REGISTER" 2>/dev/null || true
}

# Kennungen aus BEIDEN Tabellen — fuer die Zitat-Deckung. Eine gestrichene
# Beobachtung bleibt zitierbar; ihre Zeile wechselt die Tabelle, nicht die
# Existenz.
alle_kennungen() {
  grep -E '^\| BEO-[0-9]{3} \|' "$REGISTER" 2>/dev/null | awk -F'|' '{print $2}' | tr -d ' ' | sort -u
}

check_register() {  # Befunde auf stdout, 1 bei Befund
  local fail=0 line kennung zaehler belege n_belege
  while IFS= read -r line; do
    [ -n "$line" ] || continue
    kennung="$(printf '%s' "$line" | awk -F'|' '{print $2}' | tr -d ' ')"
    zaehler="$(printf '%s' "$line" | awk -F'|' '{print $5}' | grep -oE '[0-9]+' | head -1)"
    belege="$(printf '%s' "$line" | awk -F'|' '{print $6}')"
    # (3) Form: nur slice-<NNN>, kein Freitext.
    if printf '%s' "$belege" | grep -qvE '^[[:space:]]*(slice-[0-9]{3})([[:space:]]*,[[:space:]]*slice-[0-9]{3})*[[:space:]]*$'; then
      echo "$REGISTER: $kennung — Belege sind kein reines 'slice-<NNN>' (Form): '$(printf '%s' "$belege" | sed 's/^ *//;s/ *$//')'"
      fail=1; continue
    fi
    n_belege="$(printf '%s' "$belege" | grep -oE 'slice-[0-9]{3}' | wc -l)"
    # (2) mindestens einer.
    if [ "$n_belege" -eq 0 ]; then
      echo "$REGISTER: $kennung — kein Beleg; eine Zeile ohne Beleg ist eine Behauptung"
      fail=1; continue
    fi
    # (4) so viele wie der Zaehler.
    if [ -n "$zaehler" ] && [ "$n_belege" -ne "$zaehler" ]; then
      echo "$REGISTER: $kennung — $n_belege Beleg(e) bei Zaehler $zaehler; beides muss uebereinstimmen"
      fail=1
    fi
  done <<< "$(register_rows)"
  return "$fail"
}

check_zitate() {  # (1) zitierte Kennung ohne Zeile
  local fail=0 kennungen bekannt k
  bekannt="$(alle_kennungen)"
  kennungen="$(grep -rhoE 'BEO-[0-9]{3}' "$DONE_DIR" 2>/dev/null | sort -u || true)"
  for k in $kennungen; do
    if ! printf '%s\n' "$bekannt" | grep -qx "$k"; then
      echo "$DONE_DIR: $k zitiert, aber ohne Zeile im Register — die Kette ist nicht auffindbar"
      fail=1
    fi
  done
  return "$fail"
}

# Selbsttest: je eine Fixture pro Befundklasse muss feuern, die gute schweigen.
self_test() {
  local tmp REG_SAVE
  tmp="$(mktemp -d)"
  REG_SAVE="$REGISTER"

  probe() {  # $1 = Zeile, $2 = erwartet "gruen"|"rot"
    printf '| Kennung | B | S | Zähler | Belege | Stand |\n|---|---|---|---|---|---|\n%s\n' "$1" > "$tmp/reg.md"
    REGISTER="$tmp/reg.md"
    if check_register >/dev/null; then [ "$2" = "gruen" ] && return 0; else [ "$2" = "rot" ] && return 0; fi
    echo "verify-observations: Selbsttest FEHLGESCHLAGEN — '$1' erwartet $2" >&2
    REGISTER="$REG_SAVE"; rm -rf "$tmp"; exit 2
  }
  probe '| BEO-001 | x | Planungs-Harness | 2× | slice-007, slice-012 | offen |' gruen
  probe '| BEO-002 | x | Planungs-Harness | 1× |  | offen |'                     rot
  probe '| BEO-003 | x | Planungs-Harness | 2× | slice-007 | offen |'            rot
  probe '| BEO-004 | x | Planungs-Harness | 1× | irgendwas | offen |'            rot
  # Gestrichene Tabelle (slice-110): ihre Zeilen haben keine Beleg-Spalte und
  # duerfen nicht als aktive gelesen werden.
  printf '| Kennung | B | S | Zähler | Belege | Stand |\n|---|---|---|---|---|---|\n| BEO-001 | x | S | 1× | slice-007 | offen |\n\n## Gestrichene Einträge\n\n| Kennung | B | Gestrichen am | Warum |\n|---|---|---|---|\n| BEO-009 | y | 2026-01-01 | Ursache weg |\n' > "$tmp/reg2.md"
  REGISTER="$tmp/reg2.md"
  if ! check_register >/dev/null; then
    echo "verify-observations: Selbsttest FEHLGESCHLAGEN — gestrichene Zeile als aktive gelesen" >&2
    REGISTER="$REG_SAVE"; rm -rf "$tmp"; exit 2
  fi
  REGISTER="$REG_SAVE"

  # Zitat ohne Registerzeile — die Gegenrichtung.
  local DONE_SAVE="$DONE_DIR"
  mkdir -p "$tmp/done"; printf 'Siehe BEO-777 im Register.\n' > "$tmp/done/slice-900-x.md"
  DONE_DIR="$tmp/done"
  if check_zitate >/dev/null; then
    echo "verify-observations: Selbsttest FEHLGESCHLAGEN — zitierte BEO ohne Zeile nicht erkannt" >&2
    DONE_DIR="$DONE_SAVE"; rm -rf "$tmp"; exit 2
  fi
  DONE_DIR="$DONE_SAVE"
  rm -rf "$tmp"
}

self_test

if [ ! -f "$REGISTER" ]; then
  echo "verify-observations: FAIL — $REGISTER fehlt (Regelwerk modul-06 §Das Beobachtungs-Register)." >&2
  exit 1
fi

fail=0
check_register || fail=1
check_zitate   || fail=1
if [ "$fail" -ne 0 ]; then
  echo "verify-observations: FAIL — Register-Deckung verletzt." >&2
  exit 1
fi

n="$(register_rows | grep -c . || true)"
echo "verify-observations ok: $n Beobachtung(en) mit formgebundenen Belegen, Zitate gedeckt (Selbsttest gefeuert)."
echo "  NICHT geprueft: Lage und Existenz der Beleg-Datei (modul-06; Reihenfolge-Falle BEO-006)."
