#!/usr/bin/env bash
# verify-slice-links.sh — die computational Haelfte von SL-002 (slice-060).
#
# PROBLEM. Eine Slice-Datei wandert per `git mv` durch die Lifecycle-Verzeichnisse
# (AGENTS.md §5). Relative Verweise, die aus ihrem heutigen Verzeichnis aufloesen,
# tun das aus dem naechsten nicht mehr — neun Vorfaelle in docs/plan/steering-loop.md
# (SL-002), zwei davon NACH dem Guide in Schritt 9 des Workflow-Skeletts.
#
# ANSATZ. Nicht vorhersagen, wohin verschoben wird, sondern eine INVARIANTE pruefen.
# Alle vier Lifecycle-Verzeichnisse liegen auf derselben Ebene, also gilt:
#
#     Ein relativer Verweis muss aus JEDEM Lifecycle-Verzeichnis aufloesen.
#
# Das deckt mehr ab als der `git mv`: auch das ANLEGEN aus einer Vorlage anderer
# Verzeichnistiefe faellt darunter — der Fall, der slice-058 zuerst traf.
#
# GEGENSTAND: open/, next/, in-progress/. `done/` ist AUSGENOMMEN, und zwar
# sachlich: es ist Endzustand (AGENTS.md §5 kennt fuenf Uebergaenge, keiner fuehrt
# hinaus), ein kuenftiger `git mv` kann dortige Verweise nicht mehr brechen. Ob sie
# HEUTE aufloesen, ist eine andere Frage und wird von `doc-check` beantwortet.
#
# NICHT GEPRUEFT (ehrliche Grenze, AC-QA-02): ob ein Verweisziel inhaltlich das
# richtige ist. Der Sensor prueft Aufloesbarkeit ueber Verzeichniswechsel, nicht
# Bedeutung — das bleibt Review-Sache.
set -euo pipefail
cd "$(dirname "$0")/.."

PLANNING="docs/plan/planning"
LIFECYCLE="open next in-progress done"   # done als ZIEL pruefen, nicht als Quelle

# Relative Verweise einer Markdown-Datei: ohne URLs, ohne reine Anker, ohne
# absolute Pfade. Anker werden abgeschnitten — geprueft wird die Datei-Existenz;
# Anker deckt `doc-check` ab.
# Code-Bloecke und Inline-Code werden VORHER entfernt: ein Link-Muster darin ist
# Text ueber einen Verweis, kein Verweis. Ohne das meldete der Sensor beim ersten
# Lauf sein eigenes Slice-Dokument, das die falsche Form `[Roadmap](roadmap.md)`
# in Backticks ZITIERT — derselbe Fehlalarm-Typ, den slice-050 (Kursiv-Regex) und
# slice-057 (Muster im Argument-String) bei ihren Sensoren gefunden haben. Ein
# Sensor, der rauscht, wird abgeschaltet statt repariert.
links_of() {  # $1 = Datei
  sed -e '/^[[:space:]]*```/,/^[[:space:]]*```/d' "$1" 2>/dev/null \
    | sed 's/`[^`]*`//g' \
    | grep -oE '\]\([^)]*\)' \
    | sed 's/^](//; s/)$//; s/#.*//' \
    | grep -vE '^https?:|^mailto:|^/|^$' \
    | sort -u || true
}

# Befunde einer Datei auf stdout; Rueckgabe 1 bei Befund.
check_file() {  # $1 = Datei
  local f="$1" l d fail=0
  while IFS= read -r l; do
    [ -n "$l" ] || continue
    for d in $LIFECYCLE; do
      if [ ! -e "$PLANNING/$d/$l" ]; then
        echo "$f: '$l' loest aus $PLANNING/$d/ nicht auf — der Verweis ueberlebt den Lifecycle-Wechsel nicht"
        fail=1
        break        # ein Zielverzeichnis genuegt als Beleg
      fi
    done
  done <<<"$(links_of "$f")"
  return "$fail"
}

# Selbsttest: beide Richtungen, dauerhaft. Ohne ihn waere ein totes Muster ein
# False-Green — und ein Sensor, dessen Negativ-Fixture den Pruefgegenstand gar
# nicht treffen kann, belegt nichts (Lehre aus slice-058).
self_test() {
  local tmp real_planning=$PLANNING
  tmp="$(mktemp -d)"
  mkdir -p "$tmp/open" "$tmp/next" "$tmp/in-progress" "$tmp/done"
  : > "$tmp/in-progress/roadmap.md"
  : > "$tmp/done/slice-001-alt.md"

  PLANNING="$tmp"
  # (a) praefixloser Verweis auf die Nachbardatei — der reale slice-058-Fall.
  printf '# s\n\n[Roadmap](roadmap.md)\n' > "$tmp/in-progress/slice-900-schlecht.md"
  # (b) zustandsunabhaengige Form — loest aus allen vier Verzeichnissen auf.
  printf '# s\n\n[Roadmap](../in-progress/roadmap.md)\n' > "$tmp/in-progress/slice-901-gut.md"
  # (c) Verweis auf eine done/-Nachbardatei, korrekt praefixiert.
  printf '# s\n\n[alt](../done/slice-001-alt.md)\n' > "$tmp/in-progress/slice-902-gut.md"
  # (d) die FALSCHE Form, aber nur ZITIERT — in Inline-Code und im Code-Block.
  #     Muss schweigen: Text ueber einen Verweis ist kein Verweis. Diese Fixture
  #     trifft das Muster beinahe, kann es also wirklich pruefen (Lehre slice-058).
  { printf '# s\n\nFalsch ist `[Roadmap](roadmap.md)`, richtig:\n\n'
    printf '```md\n[Roadmap](roadmap.md)\n```\n'
    printf '\n[Roadmap](../in-progress/roadmap.md)\n'; } > "$tmp/in-progress/slice-903-gut.md"

  if check_file "$tmp/in-progress/slice-900-schlecht.md" >/dev/null; then
    echo "verify-slice-links: Selbsttest FEHLGESCHLAGEN — zustandsabhaengiger Verweis nicht erkannt (Pruefung tot)" >&2
    PLANNING=$real_planning; rm -rf "$tmp"; exit 2
  fi
  local good
  for good in slice-901-gut slice-902-gut slice-903-gut; do
    if ! check_file "$tmp/in-progress/$good.md" >/dev/null; then
      echo "verify-slice-links: Selbsttest FEHLGESCHLAGEN — konformer Verweis in '$good' beanstandet" >&2
      PLANNING=$real_planning; rm -rf "$tmp"; exit 2
    fi
  done
  PLANNING=$real_planning
  rm -rf "$tmp"
}

self_test

fail=0
count=0
for d in open next in-progress; do
  for f in "$PLANNING/$d"/slice-*.md; do
    [ -e "$f" ] || continue
    count=$((count + 1))
    if ! out="$(check_file "$f")"; then
      printf '%s\n' "$out" >&2
      fail=1
    fi
  done
done

if [ "$fail" -ne 0 ]; then
  echo "verify-slice-links: FAIL — Verweise ueberleben den Lifecycle-Wechsel nicht (SL-002)." >&2
  echo "Richtig ist die zustandsunabhaengige Form, z. B. '../in-progress/roadmap.md' statt 'roadmap.md'." >&2
  exit 1
fi
echo "verify-slice-links ok: $count wandernde(r) Slice(s) mit lifecycle-festen Verweisen (Selbsttest gefeuert; done/ ist Endzustand und ausgenommen)."
