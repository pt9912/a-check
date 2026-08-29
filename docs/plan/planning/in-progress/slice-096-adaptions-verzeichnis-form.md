# slice-096 — Etappe C1: Adaptions-Speicher in die Verzeichnis-Form

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** keine `AC-*`/`ADR-*` — Formwechsel im Konventionsspeicher.
**Bezug:** Etappe **C** aus [slice-092 §6](../done/slice-092-regelwerk-v5120-delta-analyse.md),
erste Hälfte. Vorgänger [slice-095](../done/slice-095-adaptions-durchgang-v5120.md) (Etappe B).
Form-Entscheid des Maintainers 2026-08-29: **Verzeichnis-Form**.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

`grundlagen-harness-dateien.md` §Konventionsspeicher macht die Verzeichnis-Form zum Default:
ein Eintrag je Datei, `conventions.md` trägt nur noch den Index. Der Grund ist Kontext, nicht
Ästhetik — *„`conventions.md` liest **jeder** Agentenlauf"*, und aufgelöste Einträge lesen sich
wie geltende. Die Form ist laut Baseline **Wahl**; der Maintainer hat den Default gewählt.

Dieser Slice wechselt **nur die Form**. Die Urteile aus Etappe B werden hier **nicht** ausgeführt —
das ist C2, und die Trennung hat einen Grund: der Formwechsel ist anker-kritisch und muss für sich
prüfbar sein. Wandert er zusammen mit acht neuen Einträgen, sagt kein Gate mehr, woran ein
gebrochener Verweis lag.

## 2. Betroffene Module

- `harness/conventions.md` §Adaptions-Block — wird Index (zwei Tabellen).
- `harness/conventions/` — neun Eintrags-Dateien; `conventions/done/` für die eine bereits
  aufgelöste.

Eine Schicht, ein Konventionsspeicher.

## 3. Die anker-kritische Hälfte

Jeder bestehende Verweis im Repo adressiert eine Adaption über den **Überschriften-Slug** der
Inline-Form (`#mr-006--baseline-committet-vendored-statt-per-url-referenziert`). Verschwinden die
`### MR-NNN — Titel`-Überschriften, verschwinden ihre Auto-Anker mit ihnen.

Die Baseline schreibt die Gegenmaßnahme wörtlich vor: *„Kam dieses Repo von der Inline-Form, trägt
die Zeile den alten Überschriften-Slug als **zweiten** Anker daneben, sonst rotten die bereits
veröffentlichten Verweise."* Jede Index-Zeile bekommt also **zwei** Anker — die stabile Kennung
`mr-NNN` für neue Verweise, den alten Slug für die bestehenden.

Belegen kann das nur `doc-check`: es prüft jeden Link **und** jeden Anker, und der Bestand hat
über 30 solcher Verweise.

**`MR-000` bleibt inline.** Die Vorlage führt ihn ausdrücklich weiter im Index — er ist die
Adoptions-Erklärung, keine Adaption, und gilt für jeden Lauf. Sein Anker bleibt damit unberührt.

## 4. Auszuführende Gates

`make gates` — tragend ist `doc-check` (Links, Anker, Kennungs-Linkpflicht) und
`gate-consistency`. Zum Abschluss `make verify`.

**Kein neuer Sensor**, also keine Negativ-Probe. Die Probe ist der Bestand selbst: über 30
vorhandene Verweise auf die alten Slugs sind die schärfste Prüfung, die dieser Formwechsel
bekommen kann — sie waren vorher grün und müssen es danach sein.

## 5. Was bewusst nicht getan wird

- **Die B-Urteile werden nicht ausgeführt.** Kein Rückbau, keine Nachfolge-Einträge, keine
  Bewegung nach `done/` außer der einen, die schon vor diesem Slice aufgelöst war. Das ist C2.
- **Das Pflichtfeld `Ersetzt-Baseline-Regel` wird nicht in die Bestandseinträge geschrieben** —
  „Einträge werden nie überschrieben"
  ([slice-095 §4](../done/slice-095-adaptions-durchgang-v5120.md)). Der Index trägt die Spalte;
  wo der Bestandseintrag sie nicht hat, steht dort ausdrücklich `—` mit Verweis auf C2, statt eine
  Regel zu erfinden.
- **`.d-check.yml` bleibt unangetastet.** Das `ids`-Muster zeigt auf `harness/conventions.md`, und
  genau dort bleibt der Index — die Linkpflicht trifft weiterhin die richtige Datei.

## 6. DoD

- [ ] Neun Eintrags-Dateien unter `harness/conventions/` (die aufgelöste unter `done/`),
      Inhalt unverändert übernommen; `conventions.md` trägt Präambel, `MR-000` inline und **zwei**
      Index-Tabellen — Beleg: Diff und Verzeichnis-Zustand.
- [ ] Jede Index-Zeile trägt **beide** Anker (Kennung + alter Überschriften-Slug); `doc-check`
      meldet **0** Befunde, obwohl über 30 Verweise auf die alten Slugs im Bestand stehen —
      Beleg: Target-Ausgabe mit Exit-Code.
- [ ] `make gates` (und bei Abschluss `make verify`) grün — **Ausgabe in eine Datei**, Exit-Code
      getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 8. Sub-Area-Modus

Berührt wird der Konventionsspeicher — weiterhin ohne eigene Zeile in der Modus-Deklaration pro
Sub-Area ([slice-091 §7](../done/slice-091-claude-md-auf-verweis-reduzieren.md)). Alle berührten
Sub-Areas mit Modus sind GF.
