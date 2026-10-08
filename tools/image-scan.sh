#!/usr/bin/env bash
# Trivy gegen das PUBLIZIERTE Container-Image (ADR-0037).
#
# ZUSAGE: meldet bekannte Schwachstellen in dem Bild, das Anwender ziehen —
# nicht im Arbeitsbaum. Zwischen zwei Releases altert das publizierte Image,
# ohne dass sich das Repo aendert; ein commit-getriggertes Gate ist gegen
# diesen Fall prinzipiell blind.
#
# ABGRENZUNG: NICHT in `gates`. Der Scan braucht Netz fuer die Vuln-Datenbank,
# und das ist hier der ZWECK, nicht ein Zugestaendnis — eine gepinnte DB faende
# nur die CVEs von gestern. `gates` bleibt hermetisch (AGENTS.md §3.1,
# `--network none`); ein netzabhaengiges Ziel darin machte jeden lokalen Lauf
# von fremder Verfuegbarkeit abhaengig.
#
# KOPPLUNG: BEIDE Trivy-Laeufe fahren `--exit-code 0`, und das ist der Kern der
# Fehlerbehandlung. Ein nicht existierendes Image quittiert Trivy mit
# `--exit-code 1` ebenfalls mit 1 — Fehler und Befund waeren nicht zu
# unterscheiden, und der Sensor meldete "behebbare CRITICAL/HIGH", wo gar nicht
# geprueft wurde. Mit `--exit-code 0` heisst ein Nicht-Null-Exit von Trivy
# eindeutig "Scan gescheitert"; ueber Befunde entscheidet die AUSWERTUNG.
#
# GRENZE: kein Docker-Socket. Trivy liest das Image aus der Registry; fuer ein
# publiziertes Bild braucht es ihn nicht, und ein gemounteter Socket waere ein
# Host-Root-Pfad fuer ein Werkzeug, das ihn nicht noetig hat. Ein LOKAL
# gebautes Bild ist ueber dieses Skript NICHT scanbar.
#
# GRENZE: das Runtime-Image ist distroless/static plus statisches Go-Binary.
# Der Fund-Raum ist praktisch die Go-Modul-Liste plus eine kleine OS-Flaeche.
# Ein gruener Lauf sagt "nichts Bekanntes in diesem Raum", nicht "das Image ist
# sicher" (AC-QA-02: die Grenze wird ausgewiesen, nicht als Vollstaendigkeit
# ausgegeben).
#
# PLATTFORMEN: jede Referenz wird je Plattform gescannt (IMAGE_SCAN_PLATFORMS,
# ADR-0043). Trivys `--platform` ist dabei KEIN Nachweis: gemessen faellt Trivy
# bei einem Bild ohne die verlangte Plattform STILL auf die vorhandene zurueck
# (Exit 0, gescannt amd64 statt arm64; slice-221). Darum liest der
# Entscheidungslauf die gescannte Architektur aus Trivys JSON und haelt sie gegen
# die verlangte; Abweichung oder fehlende Angabe heisst GESCHEITERT.
#
# GRENZE: der Zaehl-Pfad zaehlt Zeilen mit dem Schluessel "VulnerabilityID" in
# Trivys JSON, und die Architektur steht unter "architecture" — benennt Trivy
# die Schluessel um, zaehlt er nichts bzw. findet keine Architektur. Das zweite
# ist rot (GESCHEITERT), das erste saehe aus wie ein sauberes Bild. Der
# Digest-Pin haelt das still, solange er steht; --selftest deckt die Auswertung,
# nicht die Schluessel.
#
# Exit-Codes: 0 = keine behebbaren CRITICAL/HIGH, 1 = solche gefunden,
#             2 = Scan gescheitert oder Pruefmenge leer.
# ACHTUNG: das sind die Codes des SKRIPTS. `make image-scan` normalisiert jeden
# fehlgeschlagenen Recipe auf make-Exit 2 — ueber `make` sind 1 und 2 NICHT
# unterscheidbar. Wer den Ausgang braucht, liest die AUSGABE oder ruft das
# Skript direkt. Genau das tut .github/workflows/image-scan.yml.
set -uo pipefail

# Digest-Pin: der Tag bleibt lesbar, der @sha256:-Digest ist die Wahrheit. Ein
# Scanner, der sich unter der Hand aendert, macht Befund-Vergleiche ueber die
# Zeit wertlos — dieselbe Begruendung wie bei jedem Basis-Image dieses Repos.
TRIVY_VERSION="${TRIVY_VERSION:-0.74.0}"
TRIVY_DIGEST="${TRIVY_DIGEST:-sha256:62b1e65e8869bc4b4c6aa4fa2b21595256c7c2f6018a9d9ad61caf87187c1969}"

# Geprueft wird, was Anwender ziehen -- und seit dem Spiegel (AC-FA-DIST-002)
# sind das ZWEI Registries. Der Hub-Ref stand hier zunaechst NICHT, weil er
# kein Bild trug; ein Ref ohne Bild haette den Nachtlauf ab dem ersten Tag rot
# gemacht. Seit 2026-08-30 traegt er eines (slice-127).
#
# GEMESSEN vor der Aufnahme, statt angenommen: beide Refs tragen denselben
# Inhalt -- der CONFIG-Digest ist identisch (sha256:e4f357f0...), der
# MANIFEST-Digest nicht (356aeaea vs. 5bdd40ca), weil er registry-lokal ist.
# Damit ist eine zweite, vom Scan unabhaengige Bestaetigung der
# Inhalts-Gleichheit da.
IMAGE_SCAN_REFS="${IMAGE_SCAN_REFS:-ghcr.io/pt9912/a-check:latest pt9912/a-check:latest}"

# Die Plattformen des Image-Index (ADR-0043). Jede Referenz wird je Plattform
# gescannt; die gescannte Architektur wird nachgeprueft (siehe Kopf).
IMAGE_SCAN_PLATFORMS="${IMAGE_SCAN_PLATFORMS:-linux/amd64 linux/arm64}"

# Cache ausserhalb des Repos: der Arbeitsbaum bleibt sauber, und `git status`
# meldet keine Werkzeug-Artefakte. XDG_CACHE_HOME wird geehrt.
CACHE="${TRIVY_CACHE:-${XDG_CACHE_HOME:-$HOME/.cache}/a-check/trivy}"

# Der Entscheidungslauf liefert Trivys JSON (eingerueckt, ein Schluessel je
# Zeile). Ausgewertet wird es mit grep/sed statt eines JSON-Parsers — ein
# Fremd-Interpreter waere eine vierte Toolchain (AGENTS.md §3.1).

# Die Auswertung als eigene Funktionen, damit sie NETZLOS pruefbar ist. Ohne
# diesen Einstieg waere die Semantik nur mit Netz zu pruefen und damit
# praktisch gar nicht (ADR-0037 §Fitness Function).
zaehle() {
  # Ein Befund = eine Zeile, die mit dem Schluessel "VulnerabilityID" BEGINNT
  # (nach Einrueckung). `grep -c` liefert 1, wenn nichts passt — deshalb
  # `|| true`, sonst risse der Zaehl-Pfad den Lauf ab und ein SAUBERES Image
  # saehe aus wie ein Fehler.
  printf '%s' "$1" | grep -cE '^[[:space:]]*"VulnerabilityID":' || true
}

architektur() {
  # Erste "architecture"-Angabe (Metadata.ImageConfig) oder leer — nie ein
  # Abbruch: leer erreicht die Pruefung, die es als GESCHEITERT meldet.
  printf '%s' "$1" | grep -oE '"architecture":[[:space:]]*"[^"]*"' | sed -n 1p \
    | sed 's/.*"\([^"]*\)"$/\1/' || true
}

befunde() {
  # Lesbare Zeile je Befund: ID und Paket, aus den Schluesseln des JSON.
  printf '%s' "$1" | grep -E '^[[:space:]]*"(VulnerabilityID|PkgName)":' \
    | sed 's/^[[:space:]]*"[A-Za-z]*":[[:space:]]*"\([^"]*\)".*/\1/' | paste -d' ' - - || true
}

if [ "${1:-}" = "--selftest" ]; then
  echo "image-scan: Selbsttest der Auswertung (netzlos)"
  fails=0
  probe() {
    local name="$1" eingabe="$2" erwartet="$3" got
    got="$(zaehle "$eingabe")"
    if [ "$got" = "$erwartet" ]; then
      printf '  ok   %-34s %s\n' "$name" "$erwartet"
    else
      printf '  FAIL %-34s erwartet %s, war: %s\n' "$name" "$erwartet" "$got"
      fails=$((fails + 1))
    fi
  }
  probe "leere Ausgabe"              ''                                                '0'
  probe "nur Leerzeile"              '
'                                                                                      '0'
  probe "ein Befund"                 '      "VulnerabilityID": "CVE-1",'              '1'
  probe "zwei Befunde"               '      "VulnerabilityID": "CVE-1",
      "VulnerabilityID": "CVE-2",'                                                     '2'
  # Die eigentliche Probe des Ankers: der Schluessel MITTEN in einem Text (etwa
  # einer Beschreibung) ist kein Befund. Ohne sie waere der Anker `^` nicht von
  # einem `grep VulnerabilityID` zu unterscheiden.
  probe "Schluessel nur im Text"     '      "Description": "see \"VulnerabilityID\": x",' '0'
  probe "Warnzeile dazwischen"       'WARN irgendwas
      "VulnerabilityID": "CVE-1",'                                                     '1'
  # Ein leerer Wert zaehlt MIT: die Zeile existiert, also gab es einen Befund.
  # Wer hier 0 erwartete, verwechselte fehlende Metadaten mit fehlendem Fund.
  probe "ID leer"                    '      "VulnerabilityID": "",'                   '1'
  arch() {
    local name="$1" eingabe="$2" erwartet="$3" got
    got="$(architektur "$eingabe")"
    if [ "$got" = "$erwartet" ]; then
      printf '  ok   %-34s %s\n' "$name" "${erwartet:-<leer>}"
    else
      printf '  FAIL %-34s erwartet %s, war: %s\n' "$name" "${erwartet:-<leer>}" "${got:-<leer>}"
      fails=$((fails + 1))
    fi
  }
  arch "Architektur arm64"           '      "architecture": "arm64",'                  'arm64'
  arch "Architektur amd64"           '"architecture":"amd64"'                           'amd64'
  # Fehlt die Angabe, ist das Ergebnis LEER — und leer ist im Lauf GESCHEITERT.
  arch "Architektur fehlt"           '      "os": "linux",'                            ''
  arch "erste Angabe gilt"           '      "architecture": "arm64",
      "architecture": "amd64",'                                                         'arm64'
  echo
  echo "== Fehlschlaege: $fails"
  [ "$fails" -eq 0 ]
  exit $?
fi

# Fail-closed bei leerer Pruefmenge — dieselbe Norm wie beim
# Grundgesamtheits-Riegel in verify-risiko-ausgaenge (slice-070, Fund F-12).
# Ohne diese Pruefung liefe die Schleife nullmal und der Schluss-echo
# behauptete Sauberkeit ueber eine nie besuchte Menge.
if [ -z "$(printf '%s' "${IMAGE_SCAN_REFS}" | tr -d '[:space:]')" ]; then
  echo "image-scan: IMAGE_SCAN_REFS ist leer — nichts zu pruefen ist KEIN gruener Befundstand." >&2
  exit 2
fi

mkdir -p "$CACHE"

trivy() {
  docker run --rm \
    -v "${CACHE}:/root/.cache/trivy" \
    "aquasec/trivy:${TRIVY_VERSION}@${TRIVY_DIGEST}" \
    image --no-progress --scanners vuln --exit-code 0 "$@"
}

if [ -z "$(printf '%s' "${IMAGE_SCAN_PLATFORMS}" | tr -d '[:space:]')" ]; then
  echo "image-scan: IMAGE_SCAN_PLATFORMS ist leer — nichts zu pruefen ist KEIN gruener Befundstand." >&2
  exit 2
fi

findings=0
errored=0

for ref in ${IMAGE_SCAN_REFS}; do
 for plat in ${IMAGE_SCAN_PLATFORMS}; do
  want="${plat#*/}"
  echo "=============================================================="
  echo "== Vollbericht (alle Schweregrade): ${ref} (${plat})"
  echo "=============================================================="
  # Faellt nie an Befunden — beantwortet "was steckt gerade drin", auch wenn
  # nichts davon behebbar ist. Welche Architektur er zeigt, weist der
  # Entscheidungslauf darunter nach.
  if ! trivy --platform "${plat}" --severity CRITICAL,HIGH,MEDIUM,LOW,UNKNOWN --format table "${ref}"; then
    echo "image-scan: Scan von ${ref} (${plat}) ist GESCHEITERT (nicht: Befunde gefunden)."
    errored=1
    continue
  fi

  echo
  echo "--------------------------------------------------------------"
  echo "-- Handlungspflichtig (CRITICAL/HIGH mit verfuegbarem Fix): ${ref} (${plat})"
  echo "--------------------------------------------------------------"
  # Nur DIESER Lauf entscheidet ueber rot (ADR-0037 Punkt 3).
  if ! out="$(trivy --platform "${plat}" --severity CRITICAL,HIGH --ignore-unfixed \
               --format json "${ref}")"; then
    echo "image-scan: Entscheidungslauf fuer ${ref} (${plat}) ist GESCHEITERT."
    errored=1
    continue
  fi

  got="$(architektur "${out}")"
  if [ "${got}" != "${want}" ]; then
    echo "image-scan: ${ref} (${plat}): Trivy hat die Architektur '${got:-<keine Angabe>}' gescannt, verlangt war '${want}' — GESCHEITERT (kein Nachweis fuer ${plat})."
    errored=1
    continue
  fi

  count="$(zaehle "${out}")"
  if [ "${count}" = "0" ]; then
    echo "OK — keine behebbaren CRITICAL/HIGH in ${ref} (${plat}, gescannt: ${got})."
  else
    befunde "${out}" | sed 's/^/  /'
    echo "image-scan: ${ref} (${plat}): ${count} behebbare CRITICAL/HIGH-Befunde."
    findings=1
  fi
  echo
 done
done

if [ "${errored}" = "1" ]; then
  echo "image-scan: mindestens ein Scan ist GESCHEITERT — der Befundstand ist UNBEKANNT, nicht gruen."
  exit 2
fi
if [ "${findings}" = "1" ]; then
  exit 1
fi
echo "image-scan: keine behebbaren CRITICAL/HIGH in: ${IMAGE_SCAN_REFS} — je Plattform: ${IMAGE_SCAN_PLATFORMS}"
exit 0
