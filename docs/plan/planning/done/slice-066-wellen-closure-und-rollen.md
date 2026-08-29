# slice-066 — Etappe F (3/3): Wellen-Closure und Rollen-Übergaben

**Status:** der Zustand ist das **Verzeichnis** dieser Datei
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5) — dieses Feld führt ihn bewusst **nicht** doppelt.
**Deckt:** **B-13** (Wellen werden nie auditierbar geschlossen) und **B-9** (Rollen ohne
Übergabe-Artefakte) aus [slice-048 §5](../done/slice-048-modul-delta-lesen.md) — die letzten
beiden offenen Funde der `v3.5.2`-Migration.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**B-13 — zwölf Wellen, null Belege.** `modul-06` gibt eine **Fünf-Schritt-Prozedur** vor, jeder
Schritt mit Beleg: Trigger prüfen · Carveout-Audit · Closure-Notiz `done/welle-NN-results.md`
plus `git mv` der Welle-Plan-Datei · **Wave-Self-Close-Commit** · Roadmap fortschreiben. Gemessen:
`welle-00` … `welle-11` existieren ausschließlich als **Prosa-Überschriften** in der Roadmap —
null Ergebnis-Notizen, null Plan-Dateien. Was gelebt wird, ist Schritt 5 zur Hälfte (die Tabelle
*Abgeschlossene Wellen*).

**B-9 — die Übergaben existieren, ihre Namen nicht.** `modul-08` nennt sechs Rollen und **neun**
Übergabe-Artefakte; ohne Artefakt gibt es „keinen Rollenwechsel, nur einen Kontext-Switch". Die
Messung ist günstiger als der Fund vermuten ließ: **sieben der neun Artefakte existieren real**,
nur unter anderen Namen — ein Slice in `in-progress/` *ist* die Planner→Implementation-Übergabe,
ein Review-Report *ist* Reviewer→Implementation. Es fehlt die **Validator**-Kante (zwei Artefakte),
und niemand hat die Zuordnung je aufgeschrieben.

## 2. Betroffene Module

- `docs/plan/planning/README.md` — neu: die Fünf-Schritt-Prozedur, mit a-checks Ersetzungen
  (`make ci` statt Replay-Lauf) und dem Bootstrap-Schnitt (B-13).
- [`harness/README.md`](../../../../harness/README.md) — neue Sektion: die neun Übergaben gegen
  a-checks reale Artefakte, mit benannten Lücken (B-9).

**Zwei Schichten:** Planungs-Doku und Harness-Doku.

## 3. Auszuführende Gates

`make gates` und `make verify`, Ausgabe je in eine Datei, Exit-Code getrennt geprüft. Kein neuer
Sensor: beide Funde sind Prozess- und Zuordnungs-Arbeit. Die Prozedur wird beim **nächsten**
Wellen-Abschluss zum ersten Mal angewandt — bis dahin ist sie deklariert, nicht belegt, und das
steht so da.

## 4. Was bewusst nicht getan wird

- **Keine retroaktiven Ergebnis-Notizen für `welle-00` … `welle-11`.** Eine nachträglich
  rekonstruierte Wellen-Closure wäre eine Nacherzählung, kein Beleg — dieselbe Entscheidung, die
  [`docs/reviews/README.md`](../../../reviews/README.md) für die Review-Lücke slice-027…041
  getroffen hat. Die Prozedur gilt **ab der nächsten Welle**.
- **Keine Welle-Plan-Dateien für den Bestand.** Schritt 3 verlangt ein `git mv` der Plan-Datei;
  wo es nie eine gab, gibt es nichts zu verschieben. Für künftige Wellen ist die Datei Teil der
  Prozedur.
- **Kein Verschieben des Steering-Loop-Kanals.** slice-057
  hatte ihn als Zwischenschritt angelegt, „bis es Wellen-Closures gibt". Es gibt jetzt die
  *Prozedur*, aber noch keine geschlossene Welle — der Kanal bliebe ohne Ort. Er bleibt der
  laufende Zähl-Ort; die Closure-Notiz **zieht** ihre Einträge daraus (Schritt 3).
- **Kein Validator-Apparat.** Die fehlende Kante wird benannt, nicht erfunden: a-check hat keinen
  Abnehmer außerhalb des Repos außer den Konsumenten seiner Image-Releases, und deren Rückmeldung
  läuft heute über Issues, nicht über ein Artefakt.

## 5. DoD

- [x] `docs/plan/planning/README.md` führt die Fünf-Schritt-Prozedur mit Beleg je Schritt,
      benennt die repo-spezifische Ersetzung des Replay-Kriteriums
      ([MR-008](../../../../harness/conventions.md#mr-008--kein-replay-keine-agenten-telemetrie))
      und den Bootstrap-Schnitt „gilt ab der nächsten Welle" (B-13).
- [x] [`harness/README.md`](../../../../harness/README.md) ordnet **alle neun** Übergaben einem
      realen a-check-Artefakt zu oder weist sie ausdrücklich als fehlend aus (B-9).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** die Fünf-Schritt-Wellen-Closure in `docs/plan/planning/README.md` — je Schritt ein
benannter Beleg, mit a-checks Ersetzung des Replay-Kriteriums durch „`make ci` grün"
([MR-008](../../../../harness/conventions.md#mr-008--kein-replay-keine-agenten-telemetrie)) und dem
Bootstrap-Schnitt „gilt ab der nächsten Welle". Dazu in
[`harness/README.md`](../../../../harness/README.md) die Zuordnung **aller neun** Übergaben aus
`modul-08` zu realen a-check-Artefakten, mit der Validator-Kante als ausdrücklicher Lücke.
**Damit sind alle 21 Funde aus slice-048 abgearbeitet oder begründet geschlossen — die
`v3.5.2`-Migration ist vollständig.**

**Lerneintrag — Form: benannte Spec-Lücke.**
> **Die Validator-Kante fehlt, und sie lässt sich nicht durch ein Artefakt herstellen, das man
> selbst schreibt.** Sieben der neun Übergaben existierten längst — sie mussten nur ihren Namen
> bekommen. Die achte und neunte nicht: Validation fragt „bauen wir das Richtige?", und diese
> Frage kann *weil sie gegen den realen Bedarf geht* nicht innerhalb des Repos beantwortet werden.
> Ein selbst verfasster „Validierungsbeleg" wäre Verifikation mit anderem Etikett — genau der
> Fall, den `modul-08` als „Mehrfachzuweisung ohne anderen Eingabe-Kontext" verbietet. Prüfsatz:
> *bevor ein fehlendes Rollen-Artefakt angelegt wird, den Absender benennen — gibt es ihn nicht
> außerhalb des eigenen Kontexts, ist die Lücke zu dokumentieren, nicht zu füllen.*

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. `docs/plan/planning/README.md` nennt für **jeden** der fünf Schritte einen prüfbaren Beleg
   (Exit-Code, `Letzte Prüfung:`-Datum, Notiz-Existenz, Commit-Hash, Roadmap-Diff);
   `harness/README.md` führt **neun** Zeilen, davon sieben mit realem Artefakt und zwei als
   fehlend ausgewiesen.

**Was hier bewusst nicht behauptet wird:** die Prozedur ist **deklariert, nicht belegt**. Ihr
erster Durchlauf ist ihre Probe, und der kommt mit dem nächsten Wellen-Abschluss. Ein Slice, der
eine Prozedur anlegt und sie zugleich für erprobt erklärt, wäre dieselbe Klasse Harness-Lüge wie
ein behauptetes Gate ohne Target.

**Zum Steering-Loop-Kanal:** slice-057 hatte ihn als
Zwischenschritt angelegt — „bis es Wellen-Closures gibt". Diese Bedingung ist **nicht** erfüllt:
es gibt jetzt die Prozedur, aber keine geschlossene Welle. Der Kanal bleibt darum der laufende
Zähl-Ort, und die Closure-Notiz *zieht* ihre Einträge daraus. Ein Register, das erst bei der
nächsten Closure entstünde, würde zwischen zwei Wellen nichts zählen — genau der Fehler, den
slice-057 vermied.

**Folge-Slices:** keine aus der Migration. Offen bleibt, was die Prozedur beim ersten realen
Wellen-Abschluss über sich selbst lernt.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
