# slice-173 — Baseline-Pins vom `versions`-Modul wächtern lassen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (der Closure-Trigger wäre die eigene DoD — kein
repo-weites Mehr).

**Bezug:** Fund des unabhängigen Reviews zu
[slice-172](../done/slice-172-baseline-v600-entfernen.md) (F-3,
MEDIUM). Dort §7, Risiko 3: der Wächter für Baseline-Pins ist nicht zu
erfinden, sondern **unkonfiguriert**.

**Berührte Spec-Stellen:** — *(keine)* — Gate-Konfiguration ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

---

## 1. Ziel

Ein nicht nachgezogener `.harness/baseline/<tag>/`-Pfad wird ein Befund,
statt still auf einen alten Stand zu zeigen, solange dieser noch existiert.

## 2. Analyse (vor der Umsetzung)

Die vendored Ziel-Form des Adaptions-Eintrags
([`MR-NNN-titel.template.md`](../../../../.harness/baseline/v6.2.0/templates/harness/conventions/MR-NNN-titel.template.md))
benennt den Wächter ausdrücklich: *„jeder Baseline-Bump entwertet `<tag>` —
der adoptierte Stand steht einmal im Adaptions-Block, ein Versions-Sensor
prüft jeden Pin dagegen; Muster in `.d-check.yml`"*. Die vendored
[`templates/.d-check.yml`](../../../../.harness/baseline/v6.2.0/templates/.d-check.yml)
liefert es fertig aus:

```yaml
# versions:
#   pin-pattern: '\.harness/baseline/(v\d+\.\d+\.\d+)/'
#   current-from: harness/conventions.md#baseline
#   exempt-paths: ["harness/conventions/done/**"]   # aufgeloeste Eintraege sind eingefroren
```

**Ist-Stand gemessen (2026-09-06):** `versions` ist in
[`.d-check.yml`](../../../../.d-check.yml) konfiguriert, trägt aber **ein**
Muster — die Lastenheft-Version. Baseline-Pins deckt es nicht, obwohl **35**
Dateien welche tragen. Ein vorhandener Prüfer ohne Gegenstand, dieselbe
Klasse wie
[`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md).

**Zu klären ist eine Frage, die der Bestand aufwirft, nicht die Ziel-Form:**
`exempt-paths` nimmt `harness/conventions/done/**` aus, weil aufgelöste
Einträge eingefroren sind.
[slice-172](../done/slice-172-baseline-v600-entfernen.md) hat den Pin
in [`MR-018`](../../../../harness/conventions.md#mr-018) dort trotzdem
gebumpt — er musste, sonst wäre der Link ins Leere gezeigt. Die Ausnahme
setzt also voraus, dass alte Stände liegenbleiben; wer löscht, bricht sie.
Der Slice entscheidet, welche der beiden Regeln in a-check gilt, statt die
Vorlage ungeprüft zu übernehmen.

## 3. Umsetzung

*(offen — entsteht mit der Umsetzung)*

## 4. Definition of Done

- [ ] [`.d-check.yml`](../../../../.d-check.yml) trägt ein zweites
      `versions`-Muster für `.harness/baseline/<tag>/`; ein künstlich
      veralteter Pin erzeugt nachweislich einen `version-stale`-Befund, ein
      korrekter nicht.
- [ ] Die `exempt-paths`-Frage aus §2 ist entschieden und die Entscheidung
      steht dort, wo sie beim nächsten Baseline-Sprung gelesen wird.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe und WIP-Limit frei.

**Rückführungen:** stellt sich heraus, dass das Muster den Bestand
massenhaft bricht statt einzelne Nachzüge zu melden — zurück nach `next/`;
eine Regel, die den Bestand flächig rot färbt, wird abgeschaltet statt
befolgt ([`AGENTS.md`](../../../../AGENTS.md) §5, Begründung zum
Commit-Scope).

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz
geschrieben. Danach Archivierung als wellenloser Slice
([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Das Muster trifft auch Prosa-Erwähnungen eines Pfades, die bewusst einen
  historischen Stand nennen — dann meldet der Sensor einen Nachzug, der
  falsch wäre* — Ausgang bei Closure.
- *`current-from` liest den Stand aus einer Prosa-Zeile; ändert deren
  Wortlaut, bricht der Sensor an einer Stelle, die niemand mit ihm in
  Verbindung bringt* — Ausgang bei Closure.

## 8. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice;
Lerneintrag — Form: wird dort benannt.)_

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Gate-/Werkzeug-Schicht** (`.d-check.yml`, `Makefile`), Achsen 1,2,3,
deklariert in
[`conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area).

**Vorgelagert — offene Beobachtungen sichten:** entsteht mit dem Übergang
nach `in-progress/`. Einschlägig ist absehbar
[`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
— dieser Slice ist eine weitere Instanz derselben Klasse, und sie hat die
Schwelle bereits erreicht.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
