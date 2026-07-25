# slice-NNN — <Titel>

> **Vorlagen-Hinweis.** Kopieren nach `docs/plan/planning/open/slice-<NNN>-<kurztitel>.md`,
> Platzhalter ersetzen, diesen Block löschen. Übersetzt aus
> `.harness/baseline/v3.5.2/templates/docs/plan/planning/slice.template.md`
> ([MR-006](../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert));
> Verweise zeigen auf die vendored Baseline und auf a-checks eigene Regeln statt auf Kurs-URLs.
> Angelegt in slice-052.

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../AGENTS.md) §3.3/§5).
**Deckt:** `AC-*`/`ADR-*`/Fund-IDs, die dieser Slice bedient.
**Bezug:** auslösender Slice, Roadmap-Zeile.

---

## 1. Auslöser

Was ist beobachtet worden — mit Messung, nicht mit Vermutung. Ein Slice, dessen Auslöser sich
nicht belegen lässt, ist ein Wunsch.

## 2. Betroffene Module

Datei-/Komponenten-Ebene.

> **Größen-Regel (Fund B-1 aus [slice-048](done/slice-048-modul-delta-lesen.md)).**
> **Höchstens drei DoD-Punkte und höchstens zwei Schichten.** Passt der Slice nicht hinein,
> ist die richtige Antwort **zerlegen**, nicht dehnen. `make verify` prüft die DoD-Zahl ab
> slice-052 maschinell.

## 3. Auszuführende Gates

Welche Sensoren belegen diesen Slice — und, falls ein neuer Sensor entsteht, wie seine
**Negativ-Probe** aussieht. Ein Sensor ohne Probe, die ihn nachweislich rot macht, ist ein toter
Sensor.

## 4. Was bewusst nicht getan wird

Abgrenzung mit Begründung. Fehlt sie, wandert später jede Nachfrage in den Slice zurück.

## 5. DoD

- [ ] <prüfbares Kriterium, mit Beleg-Art>
- [ ] <prüfbares Kriterium, mit Beleg-Art>
- [ ] `make gates` (und bei Abschluss `make verify`) grün — **Ausgabe in eine Datei**, Exit-Code
      getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../AGENTS.md) §5; `make verify` prüft das.)_

**Geliefert:** ein Satz, was wirklich im Repo steht.

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>.**
Die Form wird **benannt**, nicht impliziert. Genau eine der drei:

- **geschärfte Regel** — eine Regel, die künftig anders gelesen wird, als Prüfsatz formuliert
- **neuer Sensor** — eine Beobachtung, die es vorher nicht gab
- **benannte Spec-Lücke** — eine Stelle, an der die Doku schweigt, wo sie sprechen müsste

> Ein Lerneintrag ohne Ursache („war schwierig") ist eine Floskel. Die Form verlangt das *weil*.

**Zwei beobachtbare Closure-Kriterien:** je eines, das ein anderer Mensch ohne Rückfrage prüfen
kann (Gate-Exit, Datei-Zustand, Messung) — kein Datum, keine Selbsteinschätzung.

**Folge-Slices:** IDs, oder ausdrücklich „keine".

## 7. Sub-Area-Modus

Pflicht-Block **nur**, wenn mindestens eine berührte Sub-Area **Brownfield** oder **Hybrid** ist.
Bei reinem Greenfield genügt der Satz: *„alle berührten Sub-Areas GF."*

Bei BF/Hybrid je Sub-Area vier Kriterien — Konventions-Dichte · Phase-Reife ·
Evidenz-/Diskrepanz-Risiko · Reconciliation-Aufwand. Vorgelagert prüfen, ob die genannte
Sub-Area überhaupt eine ist: drei Inklusions-Achsen (eigene `MR-NNN` denkbar · eigene
Inventur-Linie · eigene Pfad-/Datei-Familie), **Schwelle mindestens zwei**. Zu grobe Schnitte
(„das Backend") vorher ausdifferenzieren.

Quelle: `.harness/baseline/v3.5.2/regelwerk/modul-05-planning-harness.md` und
`grundlagen-konventionen.md` §Was ist eine Sub-Area?
