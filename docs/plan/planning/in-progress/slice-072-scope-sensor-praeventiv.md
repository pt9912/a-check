# slice-072 — der Scope-Sensor greift, bevor der Fehler veröffentlicht ist

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** den Nebenbefund aus [slice-069](../done/slice-069-sensor-fehler-propagierung.md) §4 —
dort ausdrücklich **nicht** aufgenommen, weil er ein anderer Fehlermechanismus ist.
Bezug zu [`SL-003`](../../steering-loop.md) und [ADR-0021](../../adr/0021-commits-modul-trace-check.md).
**Bezug:** realer Vorfall am 2026-08-09 (CI-Run `31301467076`); Roadmap-Zeile *Aktuelle Welle* in
der [Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**Mechanismus: der Sensor misst richtig, aber zu spät.** Er ist nicht falsch-grün wie die Funde
der Gruppe A — er meldet korrekt. Nur meldet er, wenn der Fehler bereits veröffentlicht ist.

Am 2026-08-09 machte der Push `6d8bbe7..2e12c58` die CI rot: drei `docs(planning)`-Commits fassten
zusätzlich `docs/reviews/` an. `commit-scope-check` hat das **korrekt** erkannt — aber erst in
[`ci.yml:70`](../../../../.github/workflows/ci.yml), also nach dem Push. Zu diesem Zeitpunkt ist
die Korrektur ein Rebase auf einem veröffentlichten Branch, und der Maintainer musste zwischen
Historie-Umschreiben und einem dauerhaft roten Lauf wählen.

Vor jedem der drei Commits war `make gates` grün gelaufen. **Der Sensor hängt in keinem
Aggregat**, das lokal vor dem Commit greift:

| Ort | trägt `trace-check` | trägt `commit-scope-check` |
|---|---|---|
| [`.githooks/commit-msg`](../../../../.githooks/commit-msg) | ja | **nein** |
| `make gates` | nein | nein |
| [`ci.yml:70`](../../../../.github/workflows/ci.yml) | ja | ja |

Die Architektur ist bereits richtig gedacht — der Hook-Kommentar nennt sie ausdrücklich: der Hook
ist die **opt-in-Prävention pro Klon**, die CI die **klon-unabhängige Kontrolle**. `trace-check`
hat beide Hälften, `commit-scope-check` nur die zweite. Es fehlt keine Idee, nur eine Verdrahtung.

## 2. Betroffene Module

Zwei Schichten:

1. **[`tools/commit-scope-check.sh`](../../../../tools/commit-scope-check.sh)** — braucht einen
   Modus, der einen **noch nicht existierenden** Commit prüft: Scope aus der Message-Datei, Pfade
   aus dem Index (`git diff --cached --name-only`) statt aus `git show <sha>`.
2. **[`.githooks/commit-msg`](../../../../.githooks/commit-msg)** und das
   [`Makefile`](../../../../Makefile) — Verdrahtung analog zu `trace-check`, das die
   `MSGFILE=`/`RANGE=`-Umschaltung bereits vormacht.

## 3. Auszuführende Gates

`make commit-scope-check` in beiden Modi, `make gates`, und ein echter Commit-Versuch über den
installierten Hook (`make hooks`).

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| Commit-Versuch: Scope `(planning)`, gestagte Datei außerhalb `docs/plan/planning/` | Hook lehnt ab, Commit entsteht **nicht** |
| derselbe Versuch mit passendem Scope | Commit geht durch |
| Range-Modus über eine reale Range | unverändert wie heute — die CI-Hälfte darf sich nicht ändern |

Die dritte Zeile ist Pflicht: dieser Slice fügt eine Hälfte hinzu, er ersetzt die andere nicht.

## 4. Was bewusst nicht getan wird

- **Die Regel selbst ändern.** Welcher Scope welche Pfade fassen darf, bleibt unverändert
  ([`AGENTS.md`](../../../../AGENTS.md) §5). Dieser Slice ändert nur *wann* gemessen wird.
- **Den Hook zur Pflicht machen.** Er bleibt opt-in pro Klon (`make hooks`) — das ist die
  bestehende Architektur, und die CI bleibt die klon-unabhängige Kontrolle. Ein Hook, der sich
  nicht umgehen lässt, existiert in git nicht; das als Sicherheit zu verkaufen wäre eine
  Harness-Lüge.
- **Andere Sensoren in den Hook hängen.** Ob `doc-immutable` oder weitere Range-Gates dort
  hingehören, ist eine eigene Messung. Ein Hook, der alles fährt, wird abgeschaltet statt befolgt.
- **Die rote CI vom 2026-08-09 heilen.** Maintainer-Entscheid: stehen lassen. Ein roter Lauf, der
  einen echten Verstoß anzeigt, ist ein Beleg — kein Schaden.

## 5. DoD

- [x] Ein Commit mit Scope `(planning)` und Pfaden außerhalb `docs/plan/planning/` wird **lokal
      abgelehnt**, bevor er entsteht. Beleg: echter `git commit`-Versuch mit installiertem Hook →
      Exit 1, `FAIL — Scope (planning), aber ausserhalb … gestagt: docs/reviews/README.md`, und
      `git log -1` zeigt unverändert den Vorgänger-Commit. Derselbe Index mit `docs(reviews):`
      ging durch (Exit 0).
- [x] Der Range-Modus verhält sich unverändert. Beleg: `RANGE=2e12c58..HEAD` → Exit 0,
      „19 (planning)-Commit(s) geprueft"; `RANGE=definitely-not-a-revision` → Exit 2 (die
      Fail-closed-Korrektur aus
      [slice-069](../done/slice-069-sensor-fehler-propagierung.md) bleibt wirksam).
- [x] `make gates` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** [`commit-scope-check.sh`](../../../../tools/commit-scope-check.sh) hat einen
`MSGFILE`-Modus, der den **Index** gegen die Pending-Message prüft;
[`Makefile`](../../../../Makefile) macht die `MSGFILE=`/`RANGE=`-Umschaltung wie `trace-check`;
[`.githooks/commit-msg`](../../../../.githooks/commit-msg) ruft beide Pending-Prüfungen —
**beide immer**, auch wenn die erste rot ist.

**Lerneintrag — Form: geschärfte Regel.** Als Prüfsatz: *Ein Sensor ist nicht dadurch wirksam,
dass er richtig misst, sondern dadurch, dass er misst, solange die Korrektur noch billig ist.*

**Die Ursache** war keine fehlende Fähigkeit, sondern eine fehlende Verdrahtung — und die fiel
niemandem auf, weil der Sensor **korrekt** meldete. `F-1`, `F-5` und `F-12` waren falsch-grün und
darum durch Fixtures fangbar; hier war jeder Lauf richtig, nur zu spät. Am 2026-08-09 kostete das
einen dauerhaft roten CI-Lauf: vor jedem der drei Commits war `make gates` grün, weil
`commit-scope-check` in keinem lokalen Aggregat hing. **Kein Gate-Lauf hätte das melden können** —
die Frage „läuft dieser Sensor früh genug?" stellt kein Sensor.

**Bewusst nicht behauptet:** Der Hook bleibt opt-in pro Klon (`make hooks`). Ein Hook, der sich
nicht umgehen ließe, existiert in git nicht; ihn als Sicherheit auszugeben wäre die Harness-Lüge,
gegen die dieser Slice antritt. Die klon-unabhängige Kontrolle bleibt die CI — der Hook macht den
Fehler nur billig.

**Zwei beobachtbare Closure-Kriterien:**

1. `git commit -m "docs(planning): …"` mit einer gestagten Datei außerhalb `docs/plan/planning/`
   endet mit Exit 1 und **ohne neuen Commit**; derselbe Index mit `docs(reviews):` geht durch.
2. Beide Modi koexistieren: `MSGFILE=` prüft den Index, `RANGE=` unverändert die Commit-Range —
   inklusive der Fail-closed-Korrektur für unauflösbare Ranges.

**Folge-Slices:** keine. Offen bleibt die in
[slice-073](../done/slice-073-dcheck-statt-eigenbau.md) gemessene Ablösung von
`gate-consistency` (1)+(2), noch ungeschnitten.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
