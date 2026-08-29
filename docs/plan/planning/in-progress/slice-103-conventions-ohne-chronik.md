# slice-103 — `harness/conventions.md` trägt den Ist-Zustand, keine Chronik

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers.
**Berührte Spec-Stellen:** — *(keine; Harness-Konventionen ohne Vertragsberührung)*
**Deckt:** keine `AC-*`/`ADR-*`.
**Bezug:** Maintainer-Vorgabe 2026-08-29: *„Keine Chronik oder Forensik. Damit wird nur der
Kontext des Code-Agenten zugemüllt."* [Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

`harness/conventions.md` liest **jeder** Agentenlauf — das ist der Grund, aus dem der
Adaptions-Speicher in slice-096 die Verzeichnis-Form bekam. Dieselbe Rechnung gilt für den Rest
der Datei, und dort steht Chronik.

**Drei Stellen, gemessen gegen die Pflichtgliederung** in `grundlagen-harness-dateien.md`
§Konventionsspeicher und die Vorlage:

| Stelle | Sollform | Ist |
|---|---|---|
| §Baseline | *„welche Konvention adoptiert, mit Stand/Version"* | zusätzlich die Etappen-Chronik zweier Migrationen |
| §Adoptierte Konventions-Quellen | **drei** Bullets (Extern · Vendored · In-Repo), *„KEINE Wiederholung des Inhalts — nur Verweise"* | **fünf** |
| §Adaptions-Block, Präambel | Index-Regeln | zusätzlich ein Block „Zur Nummern-Frage (geprüft in slice-055, ohne Befund)" |

**Der Beweis, dass Chronik nicht bloß Ballast ist, steht in der Datei selbst.** §Baseline behauptet
heute *„Offen ist allein **D**"* — geschrieben in slice-099, überholt seit slice-102. Eine Sektion,
die Fortschritt erzählt, muss bei jedem Schritt nachgezogen werden und ist zwischendurch falsch.
Der Ist-Zustand dagegen ändert sich nur, wenn er sich ändert.

**Die beiden überzähligen Bullets tragen je einen eigenen Defekt:**

- *Problem-Quellen* nennt die vier `arch-check.sh`-Varianten — und sagt im letzten Satz selbst,
  sie definierten *„die **Anforderung** …, **nicht die Harness-Form**"*. Ein Eintrag unter
  *Konventions-Quellen*, der von sich sagt, keine zu sein, benennt seinen eigenen richtigen Ort.
- *Konventions-Vorbild* sagt, die `d-check`-Muster würden übernommen *„sobald die jeweiligen
  Slices sie anlegen"* — die haben sie mit slice-003 bis slice-005 angelegt. Stehengebliebene
  Zukunftsform über einen seit Monaten erreichten Zustand.

## 2. Betroffene Module

`harness/conventions.md`, drei Abschnitte. Eine Schicht.

## 3. Auszuführende Gates

`make gates` — tragend ist `doc-check`: die Kürzungen entfernen Verweise, und `MR-006`
adressiert `§Baseline` per Anker. Zum Abschluss `make verify`.

**Kein neuer Sensor.** Die Probe ist der Bestand: **niemand** verweist auf die drei Passagen —
gemessen, nicht angenommen; der einzige Fremdtreffer ist eine beiläufige Erwähnung des Wortes
„Konventions-Vorbild" in einer Closure-Notiz, kein Link.

## 4. Was bewusst nicht getan wird

- **Nichts wird „gerettet".** Die Problem-Quellen wandern **nicht** ins Lastenheft — Maintainer-
  Vorgabe: keine Chronik, auch nicht anderswo. `spec/lastenheft.md` §Zweck trägt den Zweck ohnehin;
  der gestrichene Bullet verweist selbst dorthin.
- **Begründungen bleiben.** Eine Adaption ohne ihr *Warum* wäre eine stille Setzung. Gestrichen
  wird, was erzählt, **was mit dem Dokument geschah** — nicht, warum eine Regel gilt.
- **Kein anderer Abschnitt.** §Anforderungs-Anlege-Prozess, §Zusatzklassen und
  §Modus-Deklaration tragen Regeln, keine Chronik; sie bleiben unberührt.

## 5. DoD

- [ ] §Baseline nennt Konvention, Stand, vendored Ort und Integrität — ohne Etappen-Erzählung;
      die überholte Aussage „Offen ist allein D" ist damit weg — Beleg: Diff.
- [ ] §Adoptierte Konventions-Quellen trägt die drei Bullets der Sollform; *Problem-Quellen* und
      *Konventions-Vorbild* sind ersatzlos gestrichen — Beleg: Diff.
- [ ] Die Nummern-Frage-Forensik in der Adaptions-Präambel ist weg; `doc-check` bleibt bei 0
      Befunden — Beleg: Target-Ausgabe.

Pflicht, aber **kein** Liefer-Punkt: `make gates` und zum Abschluss `make verify` grün — Ausgabe
in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Berührt wird **Harness-Einstieg** — die Sub-Area, die slice-101 deklariert hat. Greenfield.
