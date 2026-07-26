---
description: Slice-Workflow für a-check — der 8-Schritt-Pfad aus AGENTS.md §6, plus Abschluss
---

# Slice-Workflow

Workflow-Skelett der Durchsetzungsschicht (slice-051). Es ist der **schwächste** der drei
Bindepunkte: es *ordnet* den Ablauf, es erzwingt ihn nicht — erzwungen wird über den
PreToolUse-Guard (Tool-Call-Gate) und den Stop-Hook (Handoff-Gate). Dieses Dokument behauptet
darum nichts, was kein Sensor deckt.

Argument (optional): der Slice, um den es geht — `/slice slice-052-…` oder eine kurze Beschreibung.

## Vor der ersten Änderung

1. `harness/README.md` lesen.
2. Die **relevante kanonische Quelle** lesen — Source Precedence beachten, nicht „alles im Repo".
3. Betroffene IDs benennen: Slice-ID, `AC-*`, `ADR-*`, betroffene Module, auszuführende Gates.
   IDs werden **referenziert, nicht erfunden**; neue entstehen nur beim Spec-/ADR-Schreiben nach
   dem Schema in `harness/conventions.md`.
4. Kleinste sinnvolle Änderung planen — **Plan vor Code**.

Spec-first, falls ein Vertrag berührt ist: Lastenheft-CR → ADR → Spezifikation → Code → Tests.
Eine ADR schärft die Spezifikation, **nie** das Lastenheft.

## Während der Arbeit

5. Engsten nützlichen Sensor laufen lassen (eine Testdatei, ein Target) — nicht gleich `gates`.
6. Repo-weiter Gate-Lauf vor dem Handoff:

   ```sh
   make gates > /tmp/gates.log 2>&1; echo "EXIT=$?"
   ```

   **Ausgabe in eine Datei umleiten, Exit-Code getrennt prüfen.** `make gates | tail` liefert den
   Exit-Code von `tail` — ein rotes Gate verschwindet dabei spurlos. Aus demselben Grund den
   Gate-Lauf **nie** mit `&&` an einen Commit ketten.

   Rot? Zurück zu **Schritt 4** (Plan verfeinern), nicht zu Schritt 1. Ein roter `arch-check` ist
   ein Plan-Defekt, kein Kontext-Defekt.

## Zum Abschluss

7. Doku/Indizes nachziehen, falls ein öffentlicher Vertrag berührt ist (Lastenheft, Spezifikation,
   Benutzerhandbuch, ADR-Index, CHANGELOG).
8. Closure-Notiz schreiben — genau eine, ausgefüllt, mit mindestens einem von drei Inhalten
   (Lernsignal *mit Ursache* · konkretes Folge-Slice · beobachtbare Architektur-Aussage), dann:

   ```sh
   make verify > /tmp/verify.log 2>&1; echo "EXIT=$?"
   ```

   Reihenfolge beachten: erst die DoD-Haken und die Notiz schreiben, **dann** `gates`/`verify`
   laufen lassen — sonst gilt der aufgezeichnete Nachweis für einen älteren Inhaltsstand und der
   Stop-Hook blockiert zu Recht.

9. Lifecycle: `git mv` in den nächsten Zustand als **eigener** Commit, ohne Inhaltsänderung.

   **Vorher die relativen Verweise prüfen.** Der Verschiebe-Commit liegt *hinter* dem Gate-Lauf
   aus Schritt 6 und ist damit ungeprüft — genau dort entstanden die roten Slice-Endstände, die
   [`SL-002`](../../docs/plan/steering-loop.md) zählt. Ein Verweis, der aus `in-progress/`
   auflöst, tut es aus `done/` nicht mehr:

   ```sh
   # Welche Verweise bräche der Wechsel nach done/?
   grep -oE '\]\([^)]*\)' <datei> | sed 's/^](//;s/)$//;s/#.*//' \
     | grep -vE '^https?:|^$' | while read -r l; do
         [ -e "docs/plan/planning/done/$l" ] || echo "BRICHT: $l"
       done
   ```

   Der kritische Fall ist die **präfixlose** Nachbardatei (`roadmap.md`), nicht der Pfad mit
   `../` — eine Prüfung, die nur `../`-Verweise ansieht, übersieht ihn. Richtig ist die
   **zustandsunabhängige** Form: `../in-progress/roadmap.md` löst aus beiden Verzeichnissen auf.
   Korrektur als eigener Commit **vor** dem `git mv` (Hard Rule §3.3).

10. Berichten: welche Sensoren liefen, mit echter Ausgabe, und welche Risiken offen bleiben.
    **Keine Erfolgsmeldung ohne Gate-Ausgabe.**

## Hard Rules, die nie gebrochen werden

- Kein Host-Toolchain-Aufruf (`go`, `pip`, `npm`, `cargo`, `apt`, …) — alles über `make`/Docker.
- Keine Inline-Suppression (`//nolint`); Ausnahmen zentral in `.golangci.yml` mit `Why:`.
- Eine `Accepted`-ADR wird nie überschrieben — Korrektur ist eine neue ADR mit `Supersedes`.
- Kein Gate wird ohne ADR gelockert.
- Merge und Push erst auf ausdrückliches Wort des Maintainers.
