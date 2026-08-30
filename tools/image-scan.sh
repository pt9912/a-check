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
# GRENZE: der Zaehl-Pfad prueft, ob Trivys Template ZEILEN liefert — nicht, ob
# die Feldnamen (.Vulnerabilities, .Severity, .FixedVersion) noch heissen wie
# hier notiert. Benennt Trivy sie um, rendert das Template nichts, und das
# saehe aus wie ein sauberes Bild. Der Digest-Pin haelt das still, solange er
# steht; --selftest deckt die Auswertung, nicht die Feldnamen.
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

# Geprueft wird, was Anwender ziehen. Heute ist das EIN Bezugsweg; kommt ein
# Docker-Hub-Spiegel dazu, gehoert sein Ref hier daneben — und zwar erst, wenn
# er ein Bild traegt: ein Ref ohne Bild machte den Nachtlauf ab dem ersten Tag
# rot.
IMAGE_SCAN_REFS="${IMAGE_SCAN_REFS:-ghcr.io/pt9912/a-check:latest}"

# Cache ausserhalb des Repos: der Arbeitsbaum bleibt sauber, und `git status`
# meldet keine Werkzeug-Artefakte. XDG_CACHE_HOME wird geehrt.
CACHE="${TRIVY_CACHE:-${XDG_CACHE_HOME:-$HOME/.cache}/a-check/trivy}"

# Eine Zeile je Befund plus ein zaehlbarer Marker — Trivys eigenes
# Template-Format statt eines JSON-Parsers. Ein Fremd-Interpreter waere eine
# vierte Toolchain (AGENTS.md §3.1).
TPL='{{ range . }}{{ range .Vulnerabilities }}FINDING {{ .Severity }} {{ .PkgName }} {{ .InstalledVersion }} -> fix {{ .FixedVersion }} {{ .VulnerabilityID }}
{{ end }}{{ end }}'

# Die Auswertung als eigene Funktion, damit sie NETZLOS pruefbar ist. Ohne
# diesen Einstieg waere die Semantik nur mit Netz zu pruefen und damit
# praktisch gar nicht (ADR-0037 §Fitness Function).
zaehle() {
  # `grep -c` liefert 1, wenn nichts passt — deshalb `|| true`, sonst risse der
  # Zaehl-Pfad den Lauf ab und ein SAUBERES Image saehe aus wie ein Fehler.
  printf '%s' "$1" | grep -c '^FINDING ' || true
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
  probe "ein Befund"                 'FINDING HIGH x/net v0.1 -> fix 0.2 CVE-1'         '1'
  probe "zwei Befunde"               'FINDING HIGH a v1 -> fix 2 CVE-1
FINDING CRITICAL b v3 -> fix 4 CVE-2'                                                  '2'
  # Die eigentliche Probe des Ankers: ein Marker MITTEN in der Zeile ist kein
  # Befund. Ohne sie waere `grep FINDING` von `grep '^FINDING '` nicht zu
  # unterscheiden.
  probe "Marker nur am Zeilenende"   'xx FINDING HIGH a v1 -> fix 2 CVE-1'             '0'
  probe "Trivy-Warnzeile dazwischen" 'WARN irgendwas
FINDING HIGH a v1 -> fix 2 CVE-1'                                                      '1'
  # Leer gerenderte Felder zaehlen MIT: die Zeile existiert, also gab es einen
  # Befund. Wer hier 0 erwartete, verwechselte fehlende Metadaten mit fehlendem
  # Fund.
  probe "Feldnamen leer gerendert"   'FINDING    ->  fix  '                            '1'
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

findings=0
errored=0

for ref in ${IMAGE_SCAN_REFS}; do
  echo "=============================================================="
  echo "== Vollbericht (alle Schweregrade): ${ref}"
  echo "=============================================================="
  # Faellt nie an Befunden — beantwortet "was steckt gerade drin", auch wenn
  # nichts davon behebbar ist.
  if ! trivy --severity CRITICAL,HIGH,MEDIUM,LOW,UNKNOWN --format table "${ref}"; then
    echo "image-scan: Scan von ${ref} ist GESCHEITERT (nicht: Befunde gefunden)."
    errored=1
    continue
  fi

  echo
  echo "--------------------------------------------------------------"
  echo "-- Handlungspflichtig (CRITICAL/HIGH mit verfuegbarem Fix): ${ref}"
  echo "--------------------------------------------------------------"
  # Nur DIESER Lauf entscheidet ueber rot (ADR-0037 Punkt 3).
  if ! out="$(trivy --severity CRITICAL,HIGH --ignore-unfixed \
               --format template --template "${TPL}" "${ref}")"; then
    echo "image-scan: Entscheidungslauf fuer ${ref} ist GESCHEITERT."
    errored=1
    continue
  fi

  count="$(zaehle "${out}")"
  if [ "${count}" = "0" ]; then
    echo "OK — keine behebbaren CRITICAL/HIGH in ${ref}."
  else
    printf '%s\n' "${out}" | grep '^FINDING ' | sed 's/^FINDING /  /'
    echo "image-scan: ${ref}: ${count} behebbare CRITICAL/HIGH-Befunde."
    findings=1
  fi
  echo
done

if [ "${errored}" = "1" ]; then
  echo "image-scan: mindestens ein Scan ist GESCHEITERT — der Befundstand ist UNBEKANNT, nicht gruen."
  exit 2
fi
if [ "${findings}" = "1" ]; then
  exit 1
fi
echo "image-scan: keine behebbaren CRITICAL/HIGH in: ${IMAGE_SCAN_REFS}"
exit 0
