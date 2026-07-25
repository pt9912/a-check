#!/usr/bin/env bash
# verify-slice-form.sh — Form-Regeln für Slice-Pläne (slice-052),
# Funde B-1 (Größen-Regel) und B-5 (Lerneintrag-Form) aus slice-048.
#
# BOOTSTRAP-AWARE GATE (Regelwerk Modul 13). Die Regeln gelten erst ab
# SLICE_FORM_FROM; die 51 älteren Slices sind grandfathered — sie entstanden
# vor der Vorlage, und ein rückwirkendes Umschreiben wäre Geschichts-Politur
# ohne Erkenntnisgewinn. Dieselbe Grandfathering-Mechanik benutzt die Baseline
# fuer die Referenz-Richtungs-Regel ("Der Gate prüft nur ab Einführung neu").
#
# Reifestufe und Hochschalt-Trigger — ohne beides waere die Stufung ein
# "Bootstrap-Schlupfloch":
#   Stufe heute : gilt ab slice-052 (Einfuehrung der Vorlage)
#   Trigger     : die Schwelle wandert NICHT weiter. Sie ist keine Kalibrierung,
#                 sondern ein Stichtag; sie faellt weg, wenn alle Slices
#                 unterhalb ebenfalls konform sind (dann ist die Variable inert).
#
# Geprueft wird STRUKTUR:
#   (1) hoechstens 3 DoD-Punkte (B-1)
#   (2) in done/: die Closure-Notiz benennt eine der drei Lerneintrag-Formen (B-5)
#
# NICHT geprueft: "hoechstens zwei Schichten" aus B-1 — was eine Schicht ist,
# ist eine Ermessensfrage ueber Modul-Grenzen; ein Zaehler darueber waere
# Schein-Genauigkeit. Bleibt Sache des Reviews.
set -euo pipefail
cd "$(dirname "$0")/.."

SLICE_FORM_FROM=52          # Stichtag: erster Slice unter der Vorlage
MAX_DOD=3
FORMEN='geschärfte Regel|neuer Sensor|benannte Spec-Lücke'

slice_num() {  # $1 = Pfad -> Nummer ohne fuehrende Nullen, leer wenn keine
  basename "$1" | sed -nE 's/^slice-0*([0-9]+)-.*/\1/p'
}

# Gilt die Form-Regel fuer diese Datei? Eigene Funktion, damit der Selbsttest
# die Grandfathering-Entscheidung wirklich pruefen kann statt nur die Nummer.
applies() {  # $1 = Pfad -> 0 = geprueft, 1 = grandfathered
  local num; num="$(slice_num "$1")"
  [ -n "$num" ] && [ "$num" -ge "$SLICE_FORM_FROM" ]
}

check_file() {  # $1 = Datei, $2 = "done" oder "offen"; Befunde auf stdout
  local f="$1" phase="$2" n fail=0 body
  n="$(grep -cE '^- \[[ x]\] ' "$f" || true)"
  if [ "$n" -gt "$MAX_DOD" ]; then
    echo "$f: $n DoD-Punkte — hoechstens $MAX_DOD erlaubt (Groessen-Regel B-1); zerlegen statt dehnen"
    fail=1
  fi
  if [ "$phase" = "done" ]; then
    body="$(awk '/^## .*[Cc]losure/ && !/[Cc]losure-(Trigger|Kriterien)/ {i=1;next} i&&/^## /{i=0} i' "$f")"
    if ! printf '%s\n' "$body" | grep -qE "Form: *($FORMEN)"; then
      echo "$f: Closure-Notiz benennt keine der drei Lerneintrag-Formen (B-5); erwartet 'Form: <geschärfte Regel|neuer Sensor|benannte Spec-Lücke>'"
      fail=1
    fi
  fi
  return "$fail"
}

# Selbsttest: je eine Fixture pro Befundklasse muss feuern, die gute schweigen.
self_test() {
  local tmp
  tmp="$(mktemp -d)"
  { echo '## 5. DoD'; echo '- [x] a'; echo '- [x] b'; echo '## 6. Closure-Notiz';
    echo '**Lerneintrag — Form: geschärfte Regel.** X, weil Y.'; } > "$tmp/gut.md"
  { echo '## 5. DoD'; for i in 1 2 3 4; do echo "- [x] p$i"; done;
    echo '## 6. Closure-Notiz'; echo '**Form: neuer Sensor**'; } > "$tmp/zu-gross.md"
  { echo '## 5. DoD'; echo '- [x] a'; echo '## 6. Closure-Notiz'; echo 'Lief gut.'; } > "$tmp/ohne-form.md"

  if ! check_file "$tmp/gut.md" done >/dev/null; then
    echo "verify-slice-form: Selbsttest FEHLGESCHLAGEN — konforme Fixture beanstandet" >&2
    rm -rf "$tmp"; exit 2
  fi
  local bad
  for bad in zu-gross ohne-form; do
    if check_file "$tmp/$bad.md" done >/dev/null; then
      echo "verify-slice-form: Selbsttest FEHLGESCHLAGEN — '$bad' nicht erkannt (Pruefung tot)" >&2
      rm -rf "$tmp"; exit 2
    fi
  done
  # Grandfathering muss in BEIDE Richtungen greifen: unterhalb des Stichtags
  # nicht anwenden, ab dem Stichtag anwenden. Sonst waere die Bootstrap-Stufe
  # entweder wirkungslos oder ein stiller Freibrief.
  if applies "$tmp/slice-003-alt.md"; then
    echo "verify-slice-form: Selbsttest FEHLGESCHLAGEN — alter Slice nicht grandfathered" >&2
    rm -rf "$tmp"; exit 2
  fi
  if ! applies "$tmp/slice-052-neu.md"; then
    echo "verify-slice-form: Selbsttest FEHLGESCHLAGEN — Slice ab Stichtag wird uebersprungen" >&2
    rm -rf "$tmp"; exit 2
  fi
  if applies "$tmp/roadmap.md"; then
    echo "verify-slice-form: Selbsttest FEHLGESCHLAGEN — Nicht-Slice-Datei wuerde geprueft" >&2
    rm -rf "$tmp"; exit 2
  fi
  rm -rf "$tmp"
}

self_test

fail=0; geprueft=0; uebersprungen=0
for dir in open next in-progress done; do
  [ -d "docs/plan/planning/$dir" ] || continue
  phase="offen"; [ "$dir" = "done" ] && phase="done"
  for f in docs/plan/planning/$dir/slice-*.md; do
    [ -e "$f" ] || continue
    if ! applies "$f"; then
      uebersprungen=$((uebersprungen + 1)); continue
    fi
    geprueft=$((geprueft + 1))
    if ! out="$(check_file "$f" "$phase")"; then
      printf '%s\n' "$out" >&2; fail=1
    fi
  done
done

if [ "$fail" -ne 0 ]; then
  echo "verify-slice-form: FAIL — Slice-Form verletzt (docs/plan/planning/slice.template.md)." >&2
  exit 1
fi
echo "verify-slice-form ok: $geprueft Slice(s) ab slice-$SLICE_FORM_FROM geprueft, $uebersprungen aelter (grandfathered)."
