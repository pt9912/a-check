# slice-NNN — <Titel>

> **Vorlagen-Hinweis.** Kopieren nach `docs/plan/planning/open/slice-<NNN>-<kurztitel>.md`,
> Platzhalter ersetzen, diesen Block löschen. Übersetzt aus
> `.harness/baseline/v5.12.0/templates/docs/plan/planning/slice.template.md`
> ([MR-006](../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert));
> Verweise zeigen auf die vendored Baseline und auf a-checks eigene Regeln statt auf Kurs-URLs.
> Gliederung auf den Stand `v5.12.0` gebracht in slice-107.
>
> **Vier Felder der Ziel-Form führt a-check nicht:** `Welle:`, das Reconciliation-Register, die
> *drei Paarungen* und der *Herkunfts-Anker*. Sie brauchen je eine eigene Entscheidung und stehen
> als Beobachtung im Register — sie fehlen hier nicht versehentlich.

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** <wer die Arbeit hält — Rolle, nicht Person; `—` bis zur Priorisierung>
**Autor:** <wer diesen Plan geschrieben hat>. **Datum:** YYYY-MM-DD.
**Berührte Spec-Stellen:** <`SPEC-*` / `ARC-*`, sonst der §-Anker; `—` wenn keine>. Der Verweis
zeigt **aufwärts** — die Spec nennt diesen Slice nie.
**Deckt:** `AC-*`/`ADR-*`/Fund-IDs, die dieser Slice bedient.
**Bezug:** auslösender Slice, Roadmap-Zeile.

---

## 1. Ziel

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md` §Ziel-Form: Slice —
Schnitt nach **Lieferwert**, nicht nach Schichten; jeder Slice ist einzeln lieferbar.

Ein Satz. Was ist beobachtet worden — mit Messung, nicht mit Vermutung? Ein Slice, dessen Auslöser
sich nicht belegen lässt, ist ein Wunsch.

## 2. Definition of Done

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md` §Ziel-Form: Slice —
**höchstens drei Liefer-Punkte** und höchstens zwei Schichten; passt der Slice nicht hinein, wird
er **zerlegt, nicht gedehnt**. Gezählt wird nur, was mit dem Umfang wächst; die Gate-Läufe und die
Closure-Pflichten darunter zählen **nicht** mit.

- [ ] <prüfbares Kriterium, mit Beleg-Art>
- [ ] <prüfbares Kriterium, mit Beleg-Art>
- [ ] <prüfbares Kriterium, mit Beleg-Art>
- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register [`observations.md`](observations.md) fortgeschrieben — neue
      `BEO-NNN` oder Zähler +1 mit Beleg; *keine Beobachtung angefallen* ist ebenfalls eine
      Antwort und wird in §7 notiert.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

Regeln dieser Sektion: Baseline-Regelwerk `grundlagen-bootstrap.md` §Was ist eine Sub-Area? —
diese Liste liefert die **Pfad-Kandidaten** für §8, nicht die Antwort: Pfad-Berührung ist nicht
hinreichend, und eine Aussagen-Berührung steht hier gar nicht.

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| <…> | neu / update / refactor | <…> |

**Auszuführende Gates:** <welche Sensoren belegen diesen Slice — und, falls ein neuer entsteht,
wie seine **Negativ-Probe** aussieht. Ein Sensor ohne Probe, die ihn nachweislich rot macht, ist
ein toter Sensor.>

## 4. Trigger

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md` §Trigger je
Lifecycle-Übergang und WIP-Limit.

**Start** (`open`/`next` → `in-progress`): <beobachtbare Bedingung>

**Rückführungen — vorab benennen, nicht erst im Nachhinein begründen:**

- `in-progress` → `next` (zu groß, zurück zur Zerlegung): <Bedingung>
- `in-progress` → `open` (blockiert): <Bedingung>

## 5. Closure-Trigger

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md` §Closure- und
Lerneintrag-Regeln — zwei beobachtbare Kriterien **und** ein Lerneintrag; ohne ihn ist der Slice
nur abgelegt.

<…>

**Was bewusst nicht getan wird:** <Abgrenzung mit Begründung. Fehlt sie, wandert später jede
Nachfrage in den Slice zurück.>

## 6. Risiken und offene Punkte

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md` §Offene Risiken werden
bei Closure aufgelöst — **jedes** Risiko bekommt genau **einen** Ausgang aus der geschlossenen
Dreier-Menge, und kein Slice geht nach `done/`, während eines ohne Ausgang dasteht.

- <Risiko> — **Ausgang:** <eingetreten: Carveout oder Folge-Slice | entfallen: gestrichen mit
  Begründung | weiter offen: `BEO-NNN` im Beobachtungs-Register>

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../AGENTS.md) §5; `make verify` prüft das.)_

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md` §Das Beobachtungs-Register —
vorhandene `BEO-NNN` **zitieren** statt neu formulieren, sonst zählt das Register zwei Namen
getrennt.

**Geliefert:** ein Satz, was wirklich im Repo steht.

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>.** Die Form wird
**benannt**, nicht impliziert; genau eine der drei. Ein Lerneintrag ohne Ursache („war schwierig")
ist eine Floskel — die Form verlangt das *weil*.

**Zwei beobachtbare Closure-Kriterien:** je eines, das ein anderer Mensch ohne Rückfrage prüfen
kann (Gate-Exit, Datei-Zustand, Messung) — kein Datum, keine Selbsteinschätzung.

**Offene Risiken und ihr Ausgang:** <jedes aus §6, mit seinem Ausgang>

**Beobachtungs-Register:** <neue `BEO-NNN` (Sub-Area, 1×, Beleg) | `BEO-NNN` auf N× erhöht |
keine Beobachtung angefallen>

**Folge-Slices:** IDs, oder ausdrücklich „keine".

## 8. Sub-Area-Modus-Begründung

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md` §Ziel-Form:
Sub-Area-Modus-Begründung — die **zwei vorgelagerten Schritte** stehen in jedem Slice-Plan,
unabhängig von Modus und Slice-Typ; die **vier** Pflichtkriterien (Konventionen-Dichte ·
Phase-Reife · Evidenz-/Diskrepanz-Risiko · Reconciliation-Aufwand) nur bei BF oder Hybrid.

**Vorgelagert — Sub-Area-Wahl prüfen:** <je berührter Sub-Area: erfüllt sie die Schwelle von
mindestens zwei der drei Inklusions-Achsen? Zu grobe Schnitte vorher ausdifferenzieren.>

**Vorgelagert — offene Beobachtungen sichten:** <Register durchgegangen; je berührter Sub-Area der
Treffer mit Zähler-Stand — oder „keine Treffer".>

Bei reinem Greenfield genügt danach: *„alle berührten Sub-Areas GF."* Sonst je Sub-Area ein Block:

### Sub-Area: <Name>

- **Modus:** GF | BF | Hybrid
- **Konventionen-Dichte:** <Beleg aus `harness/conventions.md`>
- **Phase-Reife:** Phase 0–5 <Begründung gegen die Phase × Modus-Matrix>
- **Evidenz-/Diskrepanz-Risiko:** <bei BF/Hybrid das Hauptrisiko; bei GF meist niedrig>
- **Reconciliation-Aufwand:** <Slice-Schätzung; Graduation-/Folge-Slice-Trigger>
