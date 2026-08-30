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
Entsteht ein **CR an ein fremdes Werkzeug**, läuft er vor dem Hinausgehen durch den Skill
[`.harness/skills/cr-text-reviewer.md`](../../.harness/skills/cr-text-reviewer.md) — jeder Satz,
der eine Tatsache behauptet, braucht seinen Handgriff (`BEO-022`, 3×).
Eine ADR schärft die Spezifikation, **nie** das Lastenheft.

## Während der Arbeit

5. Engsten nützlichen Sensor laufen lassen (eine Testdatei, ein Target) — nicht gleich `gates`.

   **Baust du selbst einen Sensor über Markdown?** Dann blende Zitat-Kontexte von Anfang an aus —
   Inline-Code, Code-Blöcke, Argument-Strings. Text, der *über* ein Muster spricht, ist nicht das
   Muster. Dreimal in Folge hat ein neuer Sensor sonst im ersten Lauf sein eigenes Umfeld gemeldet
   ([`SL-004`](../../docs/plan/planning/observations.md)). Eine Fixture mit **zitiertem** Muster gehört in
   den Selbsttest: sie trifft das Muster beinahe und prüft es damit wirklich.
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
8. Closure-Notiz schreiben — genau eine, ausgefüllt, mit mindestens einem von drei Inhalten.
   **Hat der Slice eine `AC-*`-Anforderung angelegt, steht ihre Kennung hier** (im Plan konnte
   sie nicht stehen — sie existierte noch nicht; `make doc-complete` prüft es).
   **Die Risiko-Ausgänge werden ab hier geprüft, nicht erst nach dem `git mv`:**
   `verify-risiko-ausgaenge` sieht `in-progress/` mit, sobald die Closure-Notiz
   ausgefüllt ist (slice-129). Vorher fiel ein Formfehler erst beim **nächsten**
   Slice auf — und zwang zu einer Inhaltsänderung nach dem `mv`, also zu einem
   Verstoß gegen Hard Rule §3.3.
   (Lernsignal *mit Ursache* · konkretes Folge-Slice · beobachtbare Architektur-Aussage), dann:

   ```sh
   make verify > /tmp/verify.log 2>&1; echo "EXIT=$?"
   ```

   Reihenfolge beachten: erst die DoD-Haken und die Notiz schreiben, **dann** `gates`/`verify`
   laufen lassen — sonst gilt der aufgezeichnete Nachweis für einen älteren Inhaltsstand und der
   Stop-Hook blockiert zu Recht.

9. Lifecycle: **`make slice-mv SLICE=<slice-NNN> TO=<zustand>`**. Das Target macht den `git mv`
   **und** zieht die Verweise **auf** die Datei nach — repo-weit, in beiden vorkommenden Formen.
   Beides gehört in **einen** Commit: der Rename bleibt bei 100 %, die Verweise liegen in
   *anderen* Dateien (Hard Rule §3.3).

   **Von Hand ist es dreimal schiefgegangen** ([`BEO-008`](../../docs/plan/planning/observations.md)):
   `git mv` bewegt die Datei, die Verweise auf sie bleiben stehen. Gefangen hat es jedes Mal
   `doc-check` — **nach** dem Wechsel.

   **Die Gegenrichtung läuft schon in Schritt 6 mit:** Verweise *in* wandernden Dateien prüft seit
   slice-080 das Modul `links` (`links.resolve-from`), also `make doc-check` und damit
   `make gates`. Invariante: *„jeder relative Verweis löst aus **jedem** Lifecycle-Verzeichnis
   auf"*. Ist sie rot, ist die Korrektur ein **eigener** Commit **vor** dem `git mv`.
   Der kritische Fall ist die **präfixlose** Nachbardatei (`roadmap.md`), nicht der Pfad mit
   `../`; richtig ist die zustandsunabhängige Form `../in-progress/roadmap.md`. Diese Prüfung von
   Hand nachzubauen lohnt nicht — ein naives `grep` über die Links meldet auch **zitierte**
   Verweise in Backticks, was die Regel ausdrücklich ausnimmt
   ([`SL-002`](../../docs/plan/planning/observations.md)).

10. Berichten: welche Sensoren liefen, mit echter Ausgabe, und welche Risiken offen bleiben.
    **Keine Erfolgsmeldung ohne Gate-Ausgabe.**

    **Der Commit-Betreff nennt, was im Diff steht.** Wandert Substanz eines anderen Slice mit, ist
    das ein eigener Commit mit dessen ID — nicht ein Anhängsel an den gerade offenen. `trace-check`
    fängt das nicht: es prüft, *ob* eine ID genannt ist, nicht ob sie die Arbeit bezeichnet.
    Dreimal ist so die Substanz eines Folge-Slice in einem `docs(planning)`-Commit des Vorgängers
    gelandet ([`SL-003`](../../docs/plan/planning/observations.md)); die Arbeit war jedes Mal in Ordnung,
    verloren ging die Auffindbarkeit.

## Hard Rules, die nie gebrochen werden

- Kein Host-Toolchain-Aufruf (`go`, `pip`, `npm`, `cargo`, `apt`, …) — alles über `make`/Docker.
- Keine Inline-Suppression (`//nolint`); Ausnahmen zentral in `.golangci.yml` mit `Why:`.
- Eine `Accepted`-ADR wird nie überschrieben — Korrektur ist eine neue ADR mit `Supersedes`.
- Kein Gate wird ohne ADR gelockert.
- Merge und Push erst auf ausdrückliches Wort des Maintainers.
