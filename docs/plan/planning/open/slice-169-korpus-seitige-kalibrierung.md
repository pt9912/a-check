# slice-169 — Korpus-seitige Kalibrierung phrasen-basierter Prüfer

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (Trigger: Review-Findings F-1/F-2 aus
[`slice-168`](../done/slice-168-pruefer-kalibrierungs-selbsttest.md) —
kein Mehr über die eigene DoD hinaus).

**Bezug:** [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
— die Korpus-Hälfte, die `slice-168` nicht gedeckt hat.

**Berührte Spec-Stellen:** — *(keine)* — Harness-/Werkzeug-Änderung ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

---

## 1. Ziel

`slice-168` kalibriert die **Werkzeug-Seite** phrasen-basierter Prüfer:
reagiert `d-check` noch auf die gewählte Phrase? Ausgefallen ist zweimal
die **Korpus-Seite**: a-checks eigene Dokumente trafen das Muster nicht
mehr (slice-120, slice-165). Dieser Slice deckt sie nach.

## 2. Analyse (vor der Umsetzung)

Zwei Befunde des unabhängigen Reviews zu `slice-168`
([Report](../../../reviews/2026-09-06-slice-168-pruefer-kalibrierungs-selbsttest.md)):

- **F-1 (HIGH):** Der Selbsttest hardcodet die richtige Fixture-Zeile und
  liest den echten Korpus nie an. Gemessen: die Kandidatenmenge des
  `reviews`-Moduls ist heute **nicht leer** (vier Slices), aber drei
  geschlossene Slices mit vorhandenem Review-Report tragen eine
  nicht-auslösende Wortform — die Menge kann jederzeit wieder leer werden,
  ohne dass ein Sensor es sagt.
- **F-2 (MEDIUM):** `TASKS_IGNORE_PATTERN` steht im Selbsttest als Kopie
  neben dem Original in `.d-check.yml`. Gemessen: nach einem Bruch der
  **echten** Konfiguration blieb der Selbsttest bei Exit 0.

**Zu entscheiden vor der Umsetzung:**

1. **Erwartungszahl oder Nichtleerheit?** `.d-check.yml` kennt mit
   `exempt-expect-count` bereits eine Erwartungszahl-Kontrolle. Eine feste
   Zahl bricht bei jedem neuen Slice; eine Nichtleerheits-Schwelle bricht
   nie, sagt aber weniger. Zu klären, welche Form die Ausfallart trifft.
2. **Woher kommt das Muster zur Laufzeit?** Ein `yq`-Zugriff auf
   `.d-check.yml` bräuchte ein weiteres Werkzeug im hermetischen Lauf; ein
   `grep`-Auszug ist fragil gegenüber YAML-Formatierung. Zu klären, welcher
   Weg ohne neue Abhängigkeit trägt.
3. **Reicht der Geltungsbereich auf die anderen phrasen-basierten
   Konfigurationen** (`versions.pin-pattern`, `commits.exempt-pattern`,
   `vcs.immutable-when`, `matrix.exclude-sections`), oder wächst er damit
   über die Größen-Regel? Vermutung: eigener Slice — hier zu belegen, nicht
   zu behaupten.

## 3. Umsetzung

*(offen — entsteht mit der Umsetzung)*

## 4. Definition of Done

- [ ] Analyse-Fragen aus §2 beantwortet und belegt.
- [ ] Korpus-seitige Kontrolle der Kandidatenmenge umgesetzt und gegen eine
      absichtlich leere Menge verifiziert.
- [ ] Muster-Kopie durch eine Kopplung an `.d-check.yml` ersetzt, gegen
      einen Bruch der echten Konfiguration verifiziert.
- [ ] Unabhängiger Review durchgeführt (Report unter `docs/reviews/`).
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe und WIP-Limit frei.

**Rückführungen:** wächst der Geltungsbereich über die beiden Befunde
hinaus (Analyse-Frage 3), zurück nach `next/` zur Zerlegung.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz
geschrieben.

## 7. Risiken und offene Punkte

- *Eine Erwartungszahl über den eigenen Korpus altert mit jedem neuen
  Slice und wird dann routinemäßig hochgezählt statt geprüft* — Ausgang bei
  Closure.
- *Die Kopplung an `.d-check.yml` braucht möglicherweise ein YAML-Werkzeug
  im hermetischen Lauf, das heute nicht da ist* — Ausgang bei Closure.

## 8. Closure-Notiz

*(offen — wird beim Übergang nach `done/` geschrieben; Lerneintrag — Form: wird dort
benannt.)*

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Gate-/Werkzeug-Schicht** (`tools/`, `Makefile`, `.d-check.yml`),
Greenfield, Schwelle ≥ 2/3 erfüllt.

**Vorgelagert — offene Beobachtungen sichten:** entsteht mit dem Übergang
nach `in-progress/` (der Register-Stand ist beim Anlegen ein anderer als
beim Beginn der Arbeit).

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
