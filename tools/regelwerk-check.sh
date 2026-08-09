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

# Stand-Auswahl nach VERSIONS-, nicht Zeichenordnung (`sort -V`, GNU coreutils).
# Mit reinem `sort` gewann `v3.5.2` gegen `v3.10.0`, und das Target haette den
# alten Stand geprueft und "Integritaet ok" gemeldet — belegt im Review
# 2026-07-26 (R-049-F4). Der Fall tritt genau waehrend einer Migration ein, also
# dann, wenn dieses Target gebraucht wird.
tags="$(find "$BASE_DIR" -mindepth 1 -maxdepth 1 -type d -printf '%f\n' 2>/dev/null | sort -V || true)"
tag="$(printf '%s\n' "$tags" | tail -1)"

# Mehrere vendored Staende sind zulaessig (Migration), aber die Auswahl darf
# nicht stillschweigend passieren: sonst bleibt der nicht gewaehlte Stand
# ungeprueft, ohne dass es jemand sieht.
n_tags="$(printf '%s\n' "$tags" | grep -c . || true)"
if [ "$n_tags" -gt 1 ]; then
  echo "regelwerk-check: HINWEIS — $n_tags vendored Staende gefunden; geprueft wird der hoechste ($tag)."
  echo "  ungeprueft bleiben: $(printf '%s\n' "$tags" | grep -v "^$tag$" | tr '\n' ' ')"
fi

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

# (1b) Vollstaendigkeit — die GEGENRICHTUNG Baum -> Manifest (slice-071, Fund F-14).
# `sha256sum -c` bestaetigt ausschliesslich GELISTETE Eintraege. Eine zusaetzliche
# Datei im vendored Baum blieb dadurch ungemessen, und der Lauf meldete weiter
# "Integritaet ok" — die Pruefrichtung war einseitig.
gelistet="$(sed -E 's/^[0-9a-f]+[[:space:]]+//' "$sums" | sort)"
vorhanden="$(cd "$BASE_DIR/$tag" && find . -type f ! -name SHA256SUMS | sed 's|^\./||' | sort)"
unlisted="$(comm -13 <(printf '%s\n' "$gelistet") <(printf '%s\n' "$vorhanden"))"
if [ -n "$unlisted" ]; then
  echo "regelwerk-check: FAIL — Datei(en) im Baseline-Baum ohne Eintrag in SHA256SUMS:" >&2
  printf '%s\n' "$unlisted" | sed 's|^|    |' >&2
  echo "  Fremdtext ist unveraendert zu vendoren (MR-006); eine nicht manifestierte Datei" >&2
  echo "  ist entweder lokal hinzugefuegt oder die Liste ist unvollstaendig." >&2
  exit 1
fi

n="$(grep -c . "$sums")"
echo "regelwerk-check: Integritaet ok — $n Datei(en) der Baseline $tag stimmen mit SHA256SUMS, keine unmanifestierte Datei im Baum."

# (2) Freshness — ausdruecklich NICHT geprueft.
cat <<EOF
regelwerk-check: Freshness NICHT geprueft (Netz-Operation, kein Gate).
  adoptierter Stand : $tag
  Release-Liste     : $RELEASES_URL
  offene Handlung   : Release-Liste ansehen; liegt dort ein neuerer Tag, ist eine
                      Migrations-Analyse faellig (Vorbild: slice-046), nicht ein
                      stiller Austausch der vendored Dateien.
EOF
