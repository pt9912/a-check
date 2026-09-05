# slice-139 — Etappe C2: Beobachtungs-Register auf Verzeichnisform migriert

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** zweiter Schritt von Etappe **C** aus [slice-135 §6](../done/slice-135-regelwerk-v600-delta-analyse.md#6-vorschlag-drei-etappen),
Vorgänger [slice-138](../done/slice-138-sub-area-kuerzel.md) (Kürzel-Vergabe). [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Harness-Mechanik ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

`docs/plan/planning/observations.md` (Tabellenform, 33 aktive + 1 gestrichene Beobachtung) wird zu
`docs/plan/planning/observations/BEO-<KUERZEL>/<slug>/{observation.md,state.md,evidence/}` —
gemäß [`observation.template.md`](../../../../.harness/baseline/v6.0.0/templates/docs/plan/planning/observation.template.md).
Der Zähler ist danach abgeleitet (Zahl der `evidence/`-Dateien), nicht mehr geführt.

## 2. Betroffene Module

- `docs/plan/planning/observations/` — neue Ablage (119 Dateien: 34×2 + 51 Evidence).
- `tools/verify-observations.sh` — auf Verzeichnisform umgeschrieben.
- Live-Referenzen repo-weit (nicht `done/`-Prosa): `AGENTS.md`, `harness/README.md`,
  `harness/conventions.md`, `docs/plan/planning/README.md`, `.claude/commands/slice.md`,
  `.harness/skills/cr-text-reviewer.md`, `README.md`.
- Fünf `MR-<NNN>`-Dateien und `conventions.md`s Index: `Ersetzt-Baseline-Regel`-Pins auf `v6.0.0`
  gehoben (Etappe B, [slice-137](../done/slice-137-adaptions-durchgang-v600.md), hatte die
  Inhalte bereits geprüft — hier nur der Pin-Nachzug).
- `AGENTS.md` §5 „Slice-Form" + vier weitere Template-Provenienz-Zeiger
  (`docs/plan/carveouts/carveout.template.md`, `docs/reviews/README.md`,
  `.harness/skills/reviewer.md`, `harness/README.md` Sensors-Kommentar) — bewusst auf `v5.12.0`
  zurückgehalten seit [slice-136](../done/slice-136-baseline-v600-vendoring.md); ihre Ziele sind
  zwischen den Ständen `diff`-geprüft identisch (`docs/plan/adr/adr.template.md`,
  `docs/plan/carveouts/carveout.template.md`, `docs/reviews/review-report.template.md`,
  `.harness/skills/reviewer.template.md`, `harness/README.template.md` — keiner davon im
  26-Dateien-Delta aus slice-135 §2), jetzt auf `v6.0.0` nachgezogen.
- Drei Schichten (Planungs-Harness, Harness-Einstieg, Gate-/Werkzeug-Schicht) — größer als die
  übliche Zwei-Schichten-Grenze, aber dieselbe Ausnahme wie bei slice-095/[slice-101](wellenlos/slice-101-beobachtungs-register.md):
  eine Formmigration des Registers selbst berührt zwangsläufig jede Stelle, die darauf zeigt.

## 3. Migrationsverfahren — mechanisch, kein Freihand-Edit

Ein Python-Skript (`generate.py`, Sitzungs-Scratch, nicht committet) parst die alte Tabelle
zellenweise (escaped `\|` korrekt behandelt) und schreibt die drei Dateien je Verzeichnis wortgleich
aus den Originalzellen — **kein** Text neu formuliert, nur die Form gewechselt. Ein zweites Skript
(`rewrite_links.py`) ersetzt repo-weit jedes Linkziel auf `observations.md`: ein Link mit Linktext
`BEO-<NNN>` bekommt das neue Verzeichnis als Ziel, jeder andere Link zeigt auf
`observations/README.md`. **Der sichtbare Linktext bleibt unverändert** — auch in `done/`-Prosa,
die nicht umgeschrieben wird (dieselbe Disziplin wie beim `git-mv`-Nachziehen von Slice-Verweisen,
Ehemals-`BEO-008`). 83 spezifische + 18 generische Links in 27 Dateien umgeschrieben,
`make doc-check` bestätigt: **0** verbleibende defekte Verweise außer dem auf diesen Slice selbst
(der beim `git mv` nach `done/` mitzieht).

## 4. Zwei Funde bei der Migration

### 4.1 Ein doppelt gezählter Vorgang (Ehemals-`BEO-022`)

Die Tabellenform zählte `slice-116` zweimal (`slice-080, slice-116, slice-116`, Zähler 3×) — CR
4/1 und CR 4/2 trafen im selben Vorgang auf. Die neue Regel „ein Vorgang zählt einmal, das erzwingt
das Dateisystem" (kein zweites `slice-116.md` möglich) deckt genau diesen Fall: Der abgeleitete
Zähler ist **2×**, nicht 3×. Das ändert nichts an der Verkörperung selbst — der Skill
`.harness/skills/cr-text-reviewer.md` besteht und wirkt —, korrigiert aber die Zählweise; in
[`BEO-GATE/cr-text-behauptet-statt-gemessen/observation.md`](../observations/BEO-GATE/cr-text-behauptet-statt-gemessen/observation.md)
dokumentiert. `cr-text-reviewer.md` selbst nannte „3×, Schwelle überschritten" — korrigiert auf
„2× bei Verkörperung".

### 4.2 Ein Sub-Area-Nebenfund war schon vor diesem Slice benannt

[slice-138 §4](../done/slice-138-sub-area-kuerzel.md#4-ein-fund-während-der-vergabe-drei-register-sub-areas-sind-undeklariert)
hatte vier fehlklassifizierte Zeilen bereits gefunden und die Zuordnung entschieden; diese Migration
setzt sie um (`Implementierung` → Kern und Regeln, `CI-/Build-Schicht`/`CI-Schicht`/
`Durchsetzungsschicht` → Gate-/Werkzeug-Schicht). Kein neuer Fund, nur die Ausführung.

## 5. Was bewusst nicht getan wird

- **`done/`-Prosa wird nicht umgeschrieben.** Nur Linkziele wandern (§3); der sichtbare Text
  „`BEO-022`" bleibt in historischen Closure-Notizen stehen — er war zum Zeitpunkt des Schreibens
  korrekt. `tools/verify-observations.sh` löst beide Zitat-Formen auf.
- **Keine Zeitdokumente-Archivierung.** Optional, ohne Trigger (keine offene Welle) — siehe
  [slice-135 §4.2](../done/slice-135-regelwerk-v600-delta-analyse.md#4-die-echten-brocken).
  `docs/plan/planning/README.md`s eigene Fünf-Schritte-Prozedur benennt die Auslassung jetzt
  ausdrücklich, statt sie stillschweigend hinter sich zu lassen.
- **`.harness/baseline/v5.12.0/` bleibt vorerst vendored.** Der Form-Review dieser Etappe ist mit
  diesem Slice abgeschlossen; der Baum-Fall ist ein eigener, kleiner Folge-Slice (Etappe A
  hinterließ ihn ausdrücklich für „nach dem Form-Review", `modul-02` §Freshness-Audit).
- **Kein neuer [`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md)-Nachfolger.** [slice-137](../done/slice-137-adaptions-durchgang-v600.md)
  benannte ihn als eigenen Folge-Slice — reine Versions-Referenz-Korrektur, kein
  Beobachtungs-Register-Gegenstand.

## 6. Auszuführende Gates

`make gates` und `make verify`. Kein neuer Gate-Sensor; `tools/verify-observations.sh` ist
umgeschrieben, nicht neu — Selbsttest läuft bei jedem Aufruf (vier Fixture-Klassen: gutes
Verzeichnis, fehlendes `evidence/`, leeres `evidence/`, Beleg ohne Vorgangs-Form; plus zwei
Zitat-Fixtures, neue und alte Form).

## 7. DoD

- [x] `docs/plan/planning/observations/` trägt 34 Verzeichnisse (33 aktiv + 1 gestrichen), jedes mit
      `observation.md` + `state.md`, aktive mit nicht leerem `evidence/`; Inhalt wortgleich aus der
      alten Tabelle migriert, kein Text neu formuliert (§3).
- [x] `tools/verify-observations.sh` prüft die neue Form (Verzeichnis-Existenz, nicht-leeres
      `evidence/`, Beleg-Form) und löst beide Zitat-Formen (`BEO-<KUERZEL>/<slug>` und
      `Ehemals: BEO-<NNN>`) auf; Selbsttest deckt beide.
- [x] Alle Live-Referenzen repo-weit auf die neue Form gehoben (Links **und** Prosa, wo sie den
      Mechanismus beschreibt — nicht nur zitiert); `done/`-Zitate bleiben wortgleich, ihre
      Linkziele lösen dennoch auf. `make gates` und `make verify` grün — Ausgabe in eine Datei,
      Exit-Code getrennt geprüft, nie in eine Pipe.

## 8. Closure-Notiz

**Geliefert:** die Beobachtungs-Register-Migration von der Tabellenform (34 Kennungen seit
slice-101) auf die Verzeichnisform der `v6.0.0`-Baseline — mechanisch erzeugt aus den
Original-Zellen (kein Text neu formuliert), 119 neue Dateien, 27 Dateien mit 101 umgeschriebenen
Linkzielen bei unverändertem sichtbaren Zitat-Text. Ein Zähl-Fehler der alten Form aufgedeckt und
korrigiert (§4.1). Fünf `MR`-Pins und vier weitere Template-Provenienz-Zeiger auf `v6.0.0`
nachgezogen — der Formvergleich, auf den Etappe A wartete, ist mit diesem Slice abgeschlossen.
Ein dritter Fund unterwegs: `tools/verify-observations.sh`s erster Entwurf verkettete die bekannten
Pfade beim Zitat-Abgleich ohne Trennzeichen (`printf '%s'` statt `'%s\n'`) — der Selbsttest deckte
es nicht auf, weil seine Fixture nur ein Verzeichnis führte und Verkettung ohne Trennzeichen bei
einem einzelnen Eintrag ununterscheidbar ist. Sichtbar wurde es erst, als dieser Slice selbst nach
`done/` wanderte und sein eigenes `BEO-GATE/…`-Zitat gegen ein zweites reales Verzeichnis prüfte.
Behoben, und der Selbsttest auf zwei Verzeichnisse erweitert (gegen den Fix regressionsgeprüft).

**Lerneintrag — Form: geschärfte Regel.** *Eine Formmigration, die nur Zellinhalte in neue Dateien
kopiert, bleibt nicht neutral — der Wechsel von „ein gepflegter Zähler" auf „eine Zahl abgeleiteter
Dateien" deckt Fehler auf, die die alte Form maskierte, weil sie Widerspruchsfreiheit nie erzwang.*
`BEO-022`s Tabellenzeile trug `3×` und drei Belege — intern konsistent, weil beide Zahlen von
Hand gepflegt wurden und zueinander passten. Erst die neue Regel „ein Vorgang zählt einmal, das
Dateisystem erzwingt es" (kein zweites `slice-116.md` möglich) zeigte: die **3** war nie richtig,
nur widerspruchsfrei zu sich selbst. *Weil* eine Form, die zwei unabhängig gepflegte Werte
(Zähler, Belegliste) nebeneinander erlaubt, Übereinstimmung mit Wahrheit verwechselt — genau das
Muster, das dieselbe Migration beim Beobachtungs-Register selbst als Grund für den Formwechsel
nennt (`grundlagen-traceability.md` `v6.0.0`: „ein gespeicherter Zähler neben einer Belegliste
sind zwei Quellen für denselben Zustand, und Kopien driften"), hier einmal am eigenen Bestand
bestätigt.

**Zwei beobachtbare Closure-Kriterien:**

1. `bash tools/verify-observations.sh` meldet **34** Beobachtungen, Selbsttest gefeuert, Exit 0.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.**

- *`.harness/baseline/v5.12.0/` ist noch vendored* — Ausgang: **Folge-Slice** — Etappe A
  kündigte den Fall des alten Baums für „nach dem Form-Review" an; der ist mit diesem Slice
  abgeschlossen, der Fall selbst ist mechanisch und klein genug für einen eigenen Slice.
- *[`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md) braucht seinen Nachfolge-Eintrag* — Ausgang: **Folge-Slice**, bereits in
  [slice-137](../done/slice-137-adaptions-durchgang-v600.md) benannt, hier nur bestätigt.
- *`evidence/`-Fund-Texte sind aus der alten Form rekonstruiert, nicht per Vorgang unterschieden*
  — Ausgang: **gestrichen mit Begründung**: kein Risiko im Sinn der Dreier-Menge — die alte
  Tabellenform führte nie eine Fund-Granularität pro Beleg, es gibt nichts zu verlieren, das
  vorher da war. Wo der Bestand echte Pro-Vorgang-Details trug (§4.1, `BEO-026`s Drei-Slice-Detail
  in seiner `state.md`), sind sie erhalten.

**Folge-Slices:** keine vergeben. Zwei sind benannt (`.harness/baseline/v5.12.0/`-Fall,
[`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md)-Nachfolger) und brauchen eigene Kennungen bei ihrer eigenen Eröffnung.

## 9. Sub-Area-Modus

Berührt: **Planungs-Harness** (`docs/plan/planning/`), **Harness-Einstieg** (`AGENTS.md`,
`harness/`), **Gate-/Werkzeug-Schicht** (`tools/verify-observations.sh`) — alle drei Greenfield.
