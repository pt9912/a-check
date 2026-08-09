# slice-055 — Etappe C (1/2): `MR`-Bestand gegen `v3.5.2` klären

**Status:** *(der Zustand ist das Verzeichnis dieser Datei, nicht dieses Feld — korrigiert in slice-063)* — erster Schnitt der **Etappe C** aus
[slice-046 §6](../done/slice-046-regelwerk-v352-migration-analyse.md) und
[slice-048 §5](../done/slice-048-modul-delta-lesen.md).
**Deckt:** die **Nummern-Kollision** aus slice-046 §4.2, Fund **B-7** (veraltete
ADR-Vorlagen-Version) und Fund **B-19** (Replay/Telemetrie ohne deklarierte Abweichung).
**Nicht hier:** B-2 (Modus pro Sub-Area) und die offenen LOW aus dem Etappe-A-Zweit-Review — 2/2.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Die Nummern-Kollision gibt es nicht

slice-046 §4.2 führte als Migrations-Brocken, dass das `conventions.template.md` von `v3.5.2`
eine **Nummer vorschreibe** — sein abgedruckter Vendoring-Eintrag trägt eine feste Nummer — während
die dortige Nummer in diesem Repo belegt ist und das Vendoring in slice-047 als
[MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
entstand.

Das Template selbst widerlegt die Lesart. Sein Adaptions-Block schreibt als **Disziplin** vor:

> „chronologisch nummeriert, keine nachträglichen inhaltlichen Änderungen an akzeptierten
> Einträgen — nur neue Einträge oder explizite Aufhebungen via neuen MR."

Eine Nummerierung, die *chronologisch pro Repo* zu sein hat, kann nicht zugleich pro Titel
vorgegeben sein. Die im Template abgedruckten Beispiel-Einträge sind **Instanzen**
neben dem generischen `### MR-NNN — <Titel der Adaption>`; nur die **Pflichtfelder** und die
Disziplin sind normativ. a-checks Vendoring-Eintrag ist damit **konform**, nicht abweichend — es ist die
sechste Adaption dieses Repos, und das ist die Regel.

Das Zweit-Review zu Etappe A hatte das bereits angedeutet („Abweichung ist ausschließlich die
Nummer und ist begründet"); hier wird es zur belegten Aussage. **Es bleibt nichts umzunummerieren.**
Der Befund wandert von „zu entscheiden" nach „geprüft, ohne Befund" — festgehalten im
Adaptions-Block, damit die Frage nicht ein drittes Mal gestellt wird.

## 2. Zwei echte Einträge

**B-7** — [MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration)
deklariert `ADR-NNNN` „gemäß Kurs-ADR-Vorlage `v1.3.0`". Das ist eine **aktuell behauptende**
Versionsaussage, und sie ist seit der Migration falsch. Sie wird **nicht am Ort korrigiert**:
die Disziplin verbietet nachträgliche inhaltliche Änderungen an akzeptierten Einträgen — dieselbe
Logik wie bei `Accepted`-ADRs. Stattdessen ein neuer Eintrag, der die Aussage ablöst.

**B-19** — a-check hat **kein** Replay (`modul-12`) und **keine** Agenten-Telemetrie
(`modul-15`). Für ein deterministisches CLI ohne Agenten-Laufzeit ist der volle Apparat
unverhältnismäßig; das ist vertretbar, aber es ist eine **Abweichung** und gehört deklariert.
Eine stillschweigende Auslassung wäre dieselbe Klasse wie ein undeklariertes Gate. Folgewirkung,
die mitdeklariert wird: das Baseline-Closure-Kriterium „Replay-Lauf grün" ist für a-check
**unerfüllbar** und wird durch die Akzeptanztests ersetzt.

## 3. Betroffene Module

- [`harness/conventions.md`](../../../../harness/conventions.md) — Notiz zur Nummern-Disziplin im
  Adaptions-Block, zwei neue `MR`-Einträge.

Eine Schicht.

## 4. Was bewusst nicht getan wird

- **Keine Umnummerierung** (§1) und **keine Korrektur am bestehenden Eintrag** (§2): beides verstieße gegen
  die Nummern- bzw. Änderungs-Disziplin des Adaptions-Blocks.
- **Kein Aufbau von Replay/Telemetrie.** Der neue Eintrag *deklariert* die Abweichung, er hebt
  sie nicht auf. Ob a-check je ein Golden Set braucht, ist eine eigene Entscheidung mit eigenem
  Trigger.

## 5. DoD

- [x] Der Adaptions-Block hält fest, dass `MR`-Nummern **chronologisch pro Repo** vergeben werden
      und die Template-Nummern Beispiele sind — die Kollision aus slice-046 §4.2 ist damit als
      Nicht-Befund belegt.
- [x] Ein neuer `MR` löst die `v1.3.0`-Aussage aus
      [MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration) ab, ohne den Eintrag
      zu verändern (B-7).
- [x] Ein neuer `MR` deklariert das Fehlen von Replay und Telemetrie als bewusste Abweichung samt
      Folgewirkung auf das Closure-Kriterium (B-19); `make gates` und `make verify` grün.

## 6. Closure-Notiz

**Geliefert:** die Nummern-Kollision als Nicht-Befund belegt und im Adaptions-Block festgehalten,
zwei neue Einträge ([MR-007](../../../../harness/conventions.md#mr-007--adr-vorlagen-version-v352-statt-v130) ADR-Vorlagen-Version,
[MR-008](../../../../harness/conventions.md#mr-008--kein-replay-keine-agenten-telemetrie) Replay/Telemetrie-Abweichung), und
eine selbst eingebaute Duplikation zurückgebaut.

**Lerneintrag — Form: geschärfte Regel.**
> **Ein Template, das Beispiele abdruckt, schreibt keine Beispiele vor.** slice-046 las eine
> abgedruckte Nummer als Vorgabe und leitete daraus einen ganzen Migrations-Brocken ab —
> Umnummerierung des Bestands, Bruch der Verweis-Identität, Kollisions-Auflösung. Widerlegt hat
> das nicht eine Diskussion, sondern **derselbe Text zwei Absätze höher**: „chronologisch
> nummeriert". Prüfsatz: *bevor eine Vorgabe aus einem Beispiel abgeleitet wird, die
> Disziplin-Sätze desselben Abschnitts lesen — sie sagen, was am Beispiel normativ ist.*

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. [MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration) ist im Diff **unverändert**; die Korrektur seiner Versionsaussage steht ausschließlich
   im neuen Eintrag — die Änderungs-Disziplin ist also nicht nur beschrieben, sondern eingehalten.

**Zweiter Fund, nicht geplant:** die AC-Drei-Pfad-Regel stand längst in `conventions.md`
§Anforderungs-Anlege-Prozess. Meine Deklaration in [`AGENTS.md`](../../../../AGENTS.md) §5 aus
slice-054 war damit eine **zweite Quelle für dieselbe Wahrheit** — zurückgebaut auf einen Verweis.
Fund B-15 bleibt gültig, aber seine Diagnose war falsch: die Regel war *deklariert und nie
durchgesetzt*, nicht *nicht deklariert*.

**Folge-Slices:** [slice-056](slice-056-sub-area-modus.md) (Etappe C, 2/2 — B-2 und die offenen
LOW).

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
