# slice-105 — Der Form-Review, richtig gemacht

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers.
**Berührte Spec-Stellen:** — *(keine; reine Ist-Messung)*
**Deckt:** keine `AC-*`/`ADR-*` — Analyse zur Abnahme, **keine** Artefakt-Änderung.
**Bezug:** Nacharbeit zu [slice-099](../done/slice-099-form-rest-und-fall-des-alten-baums.md),
das Etappe C als abgeschlossen erklärt hat. [Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

Der Form-Review in slice-099 war **zweifach zu eng**, und beide Verengungen standen dort sogar als
Tabelle — was sie plausibel aussehen ließ:

1. Er verglich **Überschriften**, nicht Inhalt. Die Roadmap-Vorlage hat sich um **+65/−11** Zeilen
   geändert; sichtbar war davon eine umbenannte Sektion.
2. Er prüfte nur Singletons **mit Pflichtgliederung**. `modul-02` bindet die Nacharbeit aber an
   *umbenannte Sektionen und neue Pflicht-Felder* — für **jeden** Singleton.

Gefunden hat beides der Maintainer, nicht der Review. Damit war „Etappe C abgeschlossen" eine
Behauptung, keine Messung. Dieser Slice holt die Messung nach — und ändert **nichts**.

## 2. Betroffene Module

Nur dieses Dokument. Die Ausführung folgt in eigenen Slices, je Artefakt geschnitten.

## 3. Was gemessen wurde — und was das Verfahren nicht kann

Verglichen wurden **17** a-check-Artefakte gegen ihre vendored Vorlage, in **zwei** Läufen:
exakter Titel-Vergleich und ein Vergleich über Kernbegriffe.

**Beide irren, in entgegengesetzte Richtungen.** Der exakte meldet Namensvarianten als fehlend
([`MR-000`](../../../../harness/conventions.md#mr-000) trägt im Repo den Zusatz *(inkl. ID-Schema-Deklaration)*). Der unscharfe vergibt zu viel:
er hält `## Aktuelle Welle` und `## Offene Wellen` für dasselbe, weil beide „Welle" enthalten —
und genau diese Umbenennung ist der Befund, der diesen Slice ausgelöst hat. **Kein Skript
entscheidet das; jede Zeile braucht ein Urteil.** Die Tabelle unten trägt darum eine Spalte
*Urteil*, und wo es fehlt, steht das ausdrücklich.

## 4. Befunde

**A — Regelwerk-Zeiger im Rumpf.** `§Template-Schichtung` verlangt *„genau **ein**
Regelwerk-Zeiger pro Pflicht-Sektion"*, und der Rumpf überlebt das Adoptieren. Ist/Soll:

| Artefakt | Zeiger |
|---|---|
| `docs/plan/planning/in-progress/roadmap.md` | 0 / 6 |
| `docs/plan/planning/README.md` | 0 / 4 |
| [`AGENTS.md`](../../../../AGENTS.md) | 0 / 3 |
| `spec/spezifikation.md` | 0 / 6 |
| `spec/architecture.md` | 0 / 3 |
| `spec/lastenheft.md` | 0 / 2 |
| `harness/conventions.md`, `harness/README.md`, `README.md`, `carveout.template.md` | 0 / 1–2 |
| `docs/plan/planning/observations.md` | **2 / 2 ✓** |

Konform ist genau das Artefakt, das **aus** der `v5.12.0`-Vorlage entstand (slice-101). Alles
Ältere trägt die Zeiger nicht — es sind rund **25** Stellen.

**B — Sektionen, je mit Urteil.**

| Artefakt | Befund | Urteil |
|---|---|---|
| `roadmap.md` | `## Aktuelle Welle` → `## Offene Wellen` | **echt** — Modellwechsel, nicht Rename (§5) |
| `slice.template.md` | 8 Abschnitte statt 7, andere Namen und Reihenfolge | **echt** — eigener Slice nötig (§5) |
| `docs/plan/planning/README.md` | vier Sektionen fehlen (*Lifecycle-Bedeutungen · Slices vs. Wellen · Aktueller Stand · Roadmap*); a-check führt stattdessen nur *Wellen-Closure-Prozedur* | **echt** — die Datei trägt heute eine Prozedur statt eines Einstiegs |
| `.harness/skills/closure-note-reviewer.md` | *Prüf-Auftrag (wörtlich aus Modul 11 §Schritt 5)* fehlt | **echt**, klein |
| `docs/plan/adr/README.md`, `docs/plan/carveouts/README.md` | *Konventionen* fehlt | **echt**, klein |
| `carveout.template.md` | *Geschichte* fehlt | **echt**, klein |
| `harness/conventions.md` | *Glossar (optional)* fehlt | **kein Befund** — die Vorlage markiert es selbst als optional |
| `AGENTS.md` | *3.7 Ein Kommentar beschreibt, was da ist* fehlt | **ungeprüft** — neue Hard Rule oder Namensvariante? |
| `README.md` | drei Sektionen fehlen | **ungeprüft** — Projekt-README ist bewusst repo-eigen |
| `spec/lastenheft.md`, `spec/architecture.md`, `spec/spezifikation.md` | 2 · 2 · 6 Sektionen | **ungeprüft** — die Spec-Straten sind auf 145–619 Zeilen gewachsen; Namensvarianten wahrscheinlich, Umfang ist ohnehin kein Kriterium |
| `harness/README.md`, `.harness/skills/reviewer.md` | — | **deckungsgleich** |

**C — Drei wiederkehrende Vorlagen fehlen ganz.** Die Baseline führt fünf (`adr/NNNN-titel`,
`slice`, `welle`, `carveout`, `review-report`); a-check hat **zwei**. Für ADR, Welle und
Review-Report gibt es keine lokale Vorlage — Instanzen entstehen also ohne Ziel-Form.

## 5. Zwei Befunde, die keine Formalie sind

**Die Roadmap-Umbenennung ist ein Modellwechsel.** `## Offene Wellen` ist *derivativ*: ein Zeiger
je flacher Welle-Datei, **Bijektion in beide Richtungen**. Dazu der Ruhe-Marker *Nichts in
Arbeit*, der **zusätzlich** zur Liste steht, wenn `in-progress/` keinen Slice trägt. Der Plural
ist Absicht — `modul-06` warnt, ein Wächter gegen *genau eine* Datei melde legitime Zustände als
Drift. Und der Marker-Wortlaut steht **absichtlich nicht** in der Vorlage, damit ein Doku-Sensor
ihn nicht als Substring der Vorlage matcht.

**Der Slice-Formwechsel schaltet zwei Sensoren still ab.** Gemessen: `verify-slice-form` sucht
`DoD` in der Überschrift und zählt in der neuen Gliederung (`## 2. Definition of Done`) **0** statt
5 Punkte — und läuft grün. Der Risiko-Ausgangs-Check sucht seinen Block *innerhalb* der Closure;
in der neuen Form ist *Risiken und offene Punkte* eine eigene Sektion davor. Ein reiner
Formwechsel wäre damit ein doppeltes halluziniertes Gate.

## 6. Was bewusst nicht getan wird

- **Nichts wird umgesetzt.** Das ist der Fehler von slice-099 in umgekehrter Richtung: dort wurde
  ausgeführt, was der Review sah, und der Review sah zu wenig.
- **Die „ungeprüft"-Zeilen werden nicht geraten.** Sie sind benannt, damit der nächste Schritt sie
  liest, statt sie für erledigt zu halten.
- **Kein Etappen-Name.** Diese Arbeit ist Nacharbeit zu C, keine neue Etappe; sie so zu nennen
  würde die Migrations-Erzählung fortschreiben, die slice-103 gerade aus `conventions.md` entfernt
  hat.

## 7. DoD

- [x] Alle 17 Artefakte mit Vorlage sind verglichen, in zwei Verfahren, und die Grenze beider
      Verfahren ist benannt statt kaschiert (§3) — Beleg: §4.
- [x] Jede Sektions-Zeile trägt ein Urteil aus **echt · kein Befund · ungeprüft · deckungsgleich**;
      keine Zeile bleibt ohne — Beleg: Tabelle §4 B.
- [x] Die zwei Befunde mit Mechanik-Folgen sind belegt statt behauptet: der Modellwechsel der
      Roadmap und die zwei Sensoren, die ein Formwechsel still abschaltet (§5).

Pflicht, aber **kein** Liefer-Punkt: `make gates` und zum Abschluss `make verify` grün — Ausgabe
in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 8. Closure-Notiz

**Geliefert:** der Form-Review, den slice-099 hätte leisten sollen — 17 Artefakte, jede Zeile mit
Urteil, die ungeprüften als ungeprüft benannt. Umgesetzt wurde nichts.

**Lerneintrag — Form: geschärfte Regel.** *Ein Review, der seinen eigenen Geltungsbereich
begründet, hat ihn damit noch nicht geprüft.* slice-099 §3 trug eine saubere Tabelle, die den
Umfang des Reviews aus `modul-02` herleitete — und die Herleitung war an zwei Stellen falsch: sie
las „Singletons **mit Pflichtgliederung**" statt „Singletons", und sie verglich Überschriften statt
Inhalt. *Weil* eine begründete Verengung genau so aussieht wie eine richtige Grenze, und der
Unterschied nur auffällt, wenn jemand **gegen** die Herleitung misst statt mit ihr. Beide Lücken
fand der Maintainer. **Der Prüfsatz:** die Grenze eines Reviews wird an einem Fall geprüft, der
außerhalb liegen soll — findet man dort etwas, war es keine Grenze.

**Zwei beobachtbare Closure-Kriterien:**

1. Die Tabelle in §4 B trägt für **jede** Zeile eines von vier Urteilen; „ungeprüft" steht
   viermal und ist damit sichtbar, statt als stilles Nichts zu erscheinen.
2. Beide Mechanik-Befunde sind gemessen, nicht behauptet: `verify-slice-form` zählt in der neuen
   Gliederung **0** statt 5 DoD-Punkte (Probe gefahren), und die Roadmap-Vorlage hat **+65/−11**
   Zeilen, von denen der Überschriften-Vergleich genau eine sah.

**Offene Risiken und ihr Ausgang:**

- *Vier Zeilen sind ungeprüft* (`AGENTS.md` §3.7, `README.md`, die drei Spec-Straten) — Ausgang:
  **weiter offen**, im Beobachtungs-Register als `BEO-011`; sie brauchen Lesearbeit, keine
  Messung.
- *Der Review selbst könnte wieder zu eng sein* — Ausgang: **weiter offen**, als `BEO-012` im
  Register. Zwei Verfahren mit gegenläufigem Irrtum sind besser als eines, aber kein Beweis; die
  nächste Lücke fände wieder jemand anders.
- *Drei wiederkehrende Vorlagen fehlen ganz* — Ausgang: **Folge-Slice**, eigener Schnitt nach der
  Ausführung der Form-Angleichung.

**Folge-Slices:** je Artefakt geschnitten, nach Abnahme dieser Tabelle. Die Reihenfolge schlägt
§4 vor: Roadmap und Slice-Vorlage zuerst, weil beide Mechanik berühren.

## 9. Sub-Area-Modus

Berührt wird **Planungs-Harness** — Greenfield.
