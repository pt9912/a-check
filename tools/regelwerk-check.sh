#!/usr/bin/env bash
# regelwerk-check.sh — Wartungs-Target für die vendored Baseline (MR-006),
# slice-049 / Fund B-8 aus slice-048.
#
# KEIN GATE. Regelwerk Modul 13 nennt genau dieses Target als Beispiel für
# "vorhanden, aber nicht als Gate behauptet"; es haengt darum nicht im
# `gates`-Aggregat und blockiert keinen Handoff.
#
# Zwei Haelften, streng getrennt (Etappe-A-Fund F-6 aus slice-047; Modul 02
# erklaert die Freshness-Pruefung ausdruecklich zur Netz-Operation ausserhalb
# der Gates):
#
#   (1) INTEGRITAET  — stimmen die vendored Dateien noch mit SHA256SUMS?
#                      hermetisch, deterministisch, fail-closed. Wird geprueft.
#   (2) FRESHNESS    — gibt es stromaufwaerts ein neueres Release als der
#                      adoptierte Stand? Netz-Operation. Wird NICHT geprueft,
#                      sondern als offene Handlung ausgegeben.
#
# Die zweite Haelfte wird bewusst nicht behauptet: ein Target, das mehr
# abzudecken vorgibt als es tut, ist nach dem Regelwerk-Abschnitt
# Durchsetzungsschicht selbst eine Harness-Luege.
set -euo pipefail
cd "$(dirname "$0")/.."

BASE_DIR=".harness/baseline"
RELEASES_URL="https://github.com/pt9912/ai-harness-course/releases"

tag="$(find "$BASE_DIR" -mindepth 1 -maxdepth 1 -type d -printf '%f\n' 2>/dev/null | sort | tail -1 || true)"
if [ -z "$tag" ]; then
  echo "regelwerk-check: FAIL — keine vendored Baseline unter $BASE_DIR/ gefunden (MR-006)." >&2
  exit 1
fi

sums="$BASE_DIR/$tag/SHA256SUMS"
if [ ! -f "$sums" ]; then
  echo "regelwerk-check: FAIL — $sums fehlt; Integritaet der Baseline $tag nicht belegbar." >&2
  exit 1
fi

# (1) Integritaet — sha256sum -c gegen die committete Liste.
if ! (cd "$BASE_DIR/$tag" && sha256sum -c --quiet SHA256SUMS); then
  echo "regelwerk-check: FAIL — vendored Baseline $tag weicht von SHA256SUMS ab." >&2
  echo "Entweder wurde Fremdtext veraendert (unzulaessig, MR-006) oder die Liste ist veraltet." >&2
  exit 1
fi

n="$(grep -c . "$sums")"
echo "regelwerk-check: Integritaet ok — $n Datei(en) der Baseline $tag stimmen mit SHA256SUMS."

# (2) Freshness — ausdruecklich NICHT geprueft.
cat <<EOF
regelwerk-check: Freshness NICHT geprueft (Netz-Operation, kein Gate).
  adoptierter Stand : $tag
  Release-Liste     : $RELEASES_URL
  offene Handlung   : Release-Liste ansehen; liegt dort ein neuerer Tag, ist eine
                      Migrations-Analyse faellig (Vorbild: slice-046), nicht ein
                      stiller Austausch der vendored Dateien.
EOF
