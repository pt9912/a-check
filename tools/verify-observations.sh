#!/usr/bin/env bash
# verify-observations.sh — Deckung des Beobachtungs-Registers (slice-102,
# auf die Verzeichnisform gehoben mit slice-139).
#
# Verifikations-Schicht, nicht Gate-Schicht: `make gates` beantwortet Code-/
# Architektur-Fragen, `make verify` DoD-/Closure-Fragen (Regelwerk Modul 11).
# Darum haengt dieses Skript an `verify`, nicht an `gates`.
#
# Regelwerk `modul-06-roadmap.md` §Das Beobachtungs-Register teilt die Arbeit
# ausdruecklich: "Mensch urteilt, Maschine prueft Deckung." Das Urteil *ist das
# dieselbe Beobachtung?* faellt beim Schreiben. Maschinell entscheidbar ist nur:
#
#   (1) jeder in done/ zitierte Pfad existiert als Verzeichnis in observations/
#   (2) jedes Verzeichnis traegt ein nicht leeres evidence/
#   (3) Beleg-FORM: Dateiname ist die Kennung eines Vorgangs (slice-NNN,
#       welle-NN, review-YYYY-MM-DD-<kurz>), kein Freitext
#
# Der Zaehler wird nicht mehr gefuehrt (er ist die Zahl der Evidence-Dateien) —
# damit entfaellt die dritte Pruefung der Tabellenform (Zaehler == Beleg-Anzahl,
# slice-102): ein abgeleiteter Zaehler kann seiner Belegliste nicht
# widersprechen (slice-135 §4.1, grundlagen-traceability.md v6.0.0).
#
# ZWEI ZITAT-FORMEN, EIN SENSOR. Die Migration (slice-139) hat die Kennung von
# `BEO-<NNN>` auf den Pfad `BEO-<KUERZEL>/<slug>` gehoben, aber `done/`-Prosa
# wird nicht umgeschrieben (AGENTS.md §3.7 sinngemaess) — Dutzende historische
# Zitate nennen weiterhin die alte Nummer. Jedes migrierte Verzeichnis traegt
# darum in `observation.md` eine `**Ehemals:** \`BEO-<NNN>\`` -Zeile; dieser
# Sensor loest beide Formen gegen denselben Bestand auf.
#
# NICHT geprueft, und beides mit Grund:
#
#   - Die UMKEHRUNG von (1) — "jedes Verzeichnis ist irgendwo zitiert". Die
#     allermeisten stehen unter der Schwelle und sind nirgends zitiert; eine
#     solche Pruefung liefe auf jedem gefuellten Register rot.
#   - Die LAGE eines Belegs (liegt die Slice-Datei in done/). Modul 06 erlaubt
#     sie erst NACH dem `git mv`; auf dem Schreib-Commit laege die Datei noch
#     nicht dort, und der Sensor meldete bei jeder korrekten Closure rot.
#     Dieselbe Reihenfolge-Falle steht als Ehemals-`BEO-006` im Register.
#     Ebenso bleibt die EXISTENZ der Vorgangs-Datei ungeprueft — ein Repo darf
#     Slices fuehren, die es nicht als Plan-Datei ablegt. Ein erfundenes
#     `slice-999` bleibt damit unentdeckt; das ist die Grenze der Deklaration
#     und gehoert benannt, nicht versteckt.
set -euo pipefail
cd "$(dirname "$0")/.."

REGISTER_DIR="docs/plan/planning/observations"
DONE_DIR="docs/plan/planning/done"
VORGANG_RE='^(slice-[0-9]{3}|welle-[0-9]{2}|review-[0-9]{4}-[0-9]{2}-[0-9]{2}-[a-z0-9-]+)\.md$'

# Alle Beobachtungs-Verzeichnisse: BEO-<KUERZEL>/<slug>/ (Tiefe 2 unter
# REGISTER_DIR). README.md liegt flach daneben und ist keins.
entry_dirs() {
  find "$REGISTER_DIR" -mindepth 2 -maxdepth 2 -type d 2>/dev/null | sort || true
}

# Pfad-Kennung eines Verzeichnisses: BEO-<KUERZEL>/<slug>.
entry_id() {
  printf '%s' "${1#"$REGISTER_DIR"/}"
}

# Alte Nummer -> Pfad-Kennung, aus der "**Ehemals:**"-Zeile in observation.md.
legacy_map() {
  local d id ehemals
  while IFS= read -r d; do
    [ -n "$d" ] || continue
    id="$(entry_id "$d")"
    ehemals="$(grep -oE '\*\*Ehemals:\*\* `BEO-[0-9]{3}`' "$d/observation.md" 2>/dev/null \
      | grep -oE 'BEO-[0-9]{3}' || true)"
    [ -n "$ehemals" ] && printf '%s %s\n' "$ehemals" "$id"
  done <<< "$(entry_dirs)"
}

check_register() {  # Befunde auf stdout, 1 bei Befund
  local fail=0 d id n bad
  while IFS= read -r d; do
    [ -n "$d" ] || continue
    id="$(entry_id "$d")"
    if [ ! -f "$d/observation.md" ] || [ ! -f "$d/state.md" ]; then
      echo "$d: observation.md oder state.md fehlt"
      fail=1; continue
    fi
    if [ ! -d "$d/evidence" ]; then
      echo "$d: kein evidence/ — eine Beobachtung ohne Beleg ist eine Behauptung"
      fail=1; continue
    fi
    n="$(find "$d/evidence" -maxdepth 1 -type f -name '*.md' | wc -l)"
    if [ "$n" -eq 0 ]; then
      echo "$d: evidence/ ist leer — eine Beobachtung ohne Beleg ist eine Behauptung"
      fail=1; continue
    fi
    bad="$(find "$d/evidence" -maxdepth 1 -type f -name '*.md' -printf '%f\n' \
      | grep -vE "$VORGANG_RE" || true)"
    if [ -n "$bad" ]; then
      echo "$d: Beleg(e) ohne Vorgangs-Form (slice-NNN/welle-NN/review-YYYY-MM-DD-*): $bad"
      fail=1
    fi
  done <<< "$(entry_dirs)"
  return "$fail"
}

check_zitate() {  # (1) zitierte Kennung ohne Verzeichnis, beide Zitat-Formen
  local fail=0 bekannt_pfade k legmap neu sichtbarer_text
  bekannt_pfade="$(entry_dirs | while IFS= read -r d; do [ -n "$d" ] && printf '%s\n' "$(entry_id "$d")"; done)"
  legmap="$(legacy_map)"

  # Nur der SICHTBARE Linktext zaehlt als Zitat — der neue Pfad steckt nach
  # der Migration im Linkziel jeder umgeschriebenen Alt-Form-Referenz
  # (](...BEO-KUERZEL/slug/observation.md)); ohne den Ausschluss meldete jede
  # migrierte Alt-Form-Zeile faelschlich ein zweites, neues Zitat.
  sichtbarer_text="$(find "$DONE_DIR" -type f -name '*.md' -exec cat {} + 2>/dev/null \
    | sed -E 's/\]\([^)]*\)//g' || true)"

  # Neue Form: BEO-<KUERZEL>/<slug> als sichtbarer Text zitiert.
  for k in $(printf '%s\n' "$sichtbarer_text" | grep -oE 'BEO-[A-Z]+/[a-z0-9-]+' | sort -u); do
    if ! printf '%s\n' "$bekannt_pfade" | grep -qxF "$k"; then
      echo "$DONE_DIR: $k zitiert, aber kein solches Verzeichnis in $REGISTER_DIR"
      fail=1
    fi
  done

  # Alte Form: BEO-<NNN> — muss ueber die Ehemals-Zeile auflösen.
  for k in $(printf '%s\n' "$sichtbarer_text" | grep -oE 'BEO-[0-9]{3}' | sort -u); do
    neu="$(printf '%s\n' "$legmap" | awk -v n="$k" '$1==n{print $2; exit}')"
    if [ -z "$neu" ]; then
      echo "$DONE_DIR: $k (alte Form) zitiert, aber kein Verzeichnis mit dieser Ehemals-Kennung"
      fail=1
    fi
  done
  return "$fail"
}

# Selbsttest: je eine Fixture pro Befundklasse muss feuern, die guten schweigen.
self_test() {
  local tmp REG_SAVE DONE_SAVE
  tmp="$(mktemp -d)"
  REG_SAVE="$REGISTER_DIR"
  DONE_SAVE="$DONE_DIR"

  mkdir -p "$tmp/reg/BEO-X/gut/evidence" "$tmp/reg/BEO-X/kein-evidence" \
    "$tmp/reg/BEO-X/leeres-evidence/evidence" "$tmp/reg/BEO-X/schlechte-form/evidence"
  printf '# t\n**Sub-Area:** t\n**Ehemals:** `BEO-001`\n\nx\n' > "$tmp/reg/BEO-X/gut/observation.md"
  printf '**Stand:** offen\n' > "$tmp/reg/BEO-X/gut/state.md"
  printf '**Vorgang:** slice-007\n**Fund:** x\n' > "$tmp/reg/BEO-X/gut/evidence/slice-007.md"

  printf '# t\n' > "$tmp/reg/BEO-X/kein-evidence/observation.md"
  printf '**Stand:** offen\n' > "$tmp/reg/BEO-X/kein-evidence/state.md"
  rm -rf "$tmp/reg/BEO-X/kein-evidence/evidence"

  printf '# t\n' > "$tmp/reg/BEO-X/leeres-evidence/observation.md"
  printf '**Stand:** offen\n' > "$tmp/reg/BEO-X/leeres-evidence/state.md"

  printf '# t\n' > "$tmp/reg/BEO-X/schlechte-form/observation.md"
  printf '**Stand:** offen\n' > "$tmp/reg/BEO-X/schlechte-form/state.md"
  printf 'x\n' > "$tmp/reg/BEO-X/schlechte-form/evidence/beleg-1.md"

  REGISTER_DIR="$tmp/reg"
  if check_register >/dev/null; then
    echo "verify-observations: Selbsttest FEHLGESCHLAGEN — defekte Verzeichnisse nicht erkannt" >&2
    REGISTER_DIR="$REG_SAVE"; rm -rf "$tmp"; exit 2
  fi
  REGISTER_DIR="$tmp/reg2"
  mkdir -p "$REGISTER_DIR/BEO-X/gut/evidence" "$REGISTER_DIR/BEO-X/zweites/evidence"
  printf '# t\n**Sub-Area:** t\n**Ehemals:** `BEO-001`\n\nx\n' > "$REGISTER_DIR/BEO-X/gut/observation.md"
  printf '**Stand:** offen\n' > "$REGISTER_DIR/BEO-X/gut/state.md"
  printf '**Vorgang:** slice-007\n**Fund:** x\n' > "$REGISTER_DIR/BEO-X/gut/evidence/slice-007.md"
  printf '# t\n**Sub-Area:** t\n\nx\n' > "$REGISTER_DIR/BEO-X/zweites/observation.md"
  printf '**Stand:** offen\n' > "$REGISTER_DIR/BEO-X/zweites/state.md"
  printf '**Vorgang:** slice-008\n**Fund:** x\n' > "$REGISTER_DIR/BEO-X/zweites/evidence/slice-008.md"
  if ! check_register >/dev/null; then
    echo "verify-observations: Selbsttest FEHLGESCHLAGEN — gutes Verzeichnis als defekt gemeldet" >&2
    REGISTER_DIR="$REG_SAVE"; rm -rf "$tmp"; exit 2
  fi

  mkdir -p "$tmp/done"
  printf 'Siehe BEO-X/nirgends im Register.\n' > "$tmp/done/slice-900-a.md"
  printf 'Siehe BEO-777 (alte Form) im Register.\n' > "$tmp/done/slice-900-b.md"
  DONE_DIR="$tmp/done"
  if check_zitate >/dev/null; then
    echo "verify-observations: Selbsttest FEHLGESCHLAGEN — unaufloesbares Zitat nicht erkannt" >&2
    REGISTER_DIR="$REG_SAVE"; DONE_DIR="$DONE_SAVE"; rm -rf "$tmp"; exit 2
  fi
  # ZWEI Verzeichnisse im Register, ZWEI Zitate — deckt einen fehlenden Separator
  # beim Sammeln der bekannten Pfade auf (Verkettung ohne Newline liesse "gut" und
  # "zweites" zu einem Fantasie-Pfad verschmelzen, den kein Zitat trifft).
  printf 'Siehe BEO-X/gut im Register.\n' > "$tmp/done/slice-901-a.md"
  printf 'Siehe BEO-X/zweites im Register.\n' > "$tmp/done/slice-901-c.md"
  printf 'Siehe BEO-001 (alte Form) im Register.\n' > "$tmp/done/slice-901-b.md"
  rm -f "$tmp/done/slice-900-a.md" "$tmp/done/slice-900-b.md"
  if ! check_zitate >/dev/null; then
    echo "verify-observations: Selbsttest FEHLGESCHLAGEN — aufloesbares Zitat als Befund gemeldet" >&2
    REGISTER_DIR="$REG_SAVE"; DONE_DIR="$DONE_SAVE"; rm -rf "$tmp"; exit 2
  fi

  REGISTER_DIR="$REG_SAVE"
  DONE_DIR="$DONE_SAVE"
  rm -rf "$tmp"
}

self_test

if [ ! -d "$REGISTER_DIR" ]; then
  echo "verify-observations: FAIL — $REGISTER_DIR fehlt (Regelwerk modul-06 §Das Beobachtungs-Register)." >&2
  exit 1
fi

fail=0
check_register || fail=1
check_zitate   || fail=1
if [ "$fail" -ne 0 ]; then
  echo "verify-observations: FAIL — Register-Deckung verletzt." >&2
  exit 1
fi

n="$(entry_dirs | grep -c . || true)"
echo "verify-observations ok: $n Beobachtung(en) mit nicht leerem evidence/, Zitate (alt und neu) gedeckt (Selbsttest gefeuert)."
echo "  NICHT geprueft: Lage und Existenz der Beleg-Datei (modul-06; Reihenfolge-Falle Ehemals-BEO-006)."
