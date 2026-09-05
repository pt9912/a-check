# slice-138 — Etappe C1: Sub-Area-Kürzel vergeben

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** erster Schritt von Etappe **C** aus [slice-135 §6](../done/slice-135-regelwerk-v600-delta-analyse.md#6-vorschlag-drei-etappen),
korrigiert und präzisiert in [slice-137 §4](../done/slice-137-adaptions-durchgang-v600.md#4-ein-zweiter-fund-die-kürzel-frage-aus-slice-135-löst-sich-nicht-in-dieser-etappe):
Kürzel-Vergabe ist Vorbedingung für die Beobachtungs-Register-Migration
(`BEO-<KUERZEL>/<slug>/`), nicht Teil von Etappe B. [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Konventions-Doku ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

`grundlagen-harness-dateien.md` §Konventionsspeicher (`v6.0.0`): sobald eine Kennungsklasse ein
Bereichssegment führt, trägt die Modus-Deklaration neben dem Namen eine **Kürzel-Spalte** — kurz,
GROSS, ohne Leerzeichen, ab Vergabe unveränderlich. Die neue Beobachtungs-Kennung
`BEO-<KUERZEL>/<slug>` ist diese Kennungsklasse; ihre Migration (Etappe C, Folge-Slice) braucht acht
Kürzel, bevor der erste Pfad geschrieben werden kann.

## 2. Betroffene Module

`harness/conventions.md` §Modus-Deklaration pro Sub-Area — eine Tabelle, eine Spalte.

## 3. Die acht Kürzel

| Sub-Area | Kürzel |
|---|---|
| Spec-Straten | `SPEC` |
| Entscheidungen | `ADR` |
| Kern und Regeln | `KERN` |
| Adapter | `ADAPT` |
| Planungs-Harness | `PLAN` |
| Gate-/Werkzeug-Schicht | `GATE` |
| Review-Harness | `REVIEW` |
| Harness-Einstieg | `HARNESS` |
| Vendored Baseline | *(kein Kürzel — kein Modus, siehe §4)* |

Vergeben aus dem etablierten Kurznamen jeder Zeile, keine Neuerfindung — dieselbe Regel wie beim
Kürzel selbst: „Das Segment wird nachgeschlagen, nicht formuliert."

## 4. Ein Fund während der Vergabe: drei Register-Sub-Areas sind undeklariert

Beim Nachschlagen, welches Kürzel eine bestehende Beobachtung bekommt (Vorbereitung für den
Folge-Slice), fiel auf: Vier der 33 aktiven `docs/plan/planning/observations.md`-Zeilen tragen eine
Sub-Area, die diese Tabelle **nicht führt** — `Implementierung` (1×), `CI-/Build-Schicht` (1×),
`CI-Schicht` (3×), `Durchsetzungsschicht` (2×). `conventions.md` selbst benennt genau diesen Fall:
*„Steht in der Spalte ein Name, den die Modus-Deklaration […] nicht führt, ist entweder die
Zuordnung falsch oder die Deklaration unvollständig."*

Geprüft gegen die vier: keine trägt eigenen Gegenstand genug für eine **neue** Zeile (Achsen-Test,
§Qualifikation) — alle vier Pfade sind in einer bestehenden Zeile bereits benannt:

- `Durchsetzungsschicht` (`.claude/hooks/`, PreToolUse-Command-Guard) — **Gate-/Werkzeug-Schicht**
  führt `.claude/` bereits explizit in ihrer Pfadliste.
- `CI-Schicht` / `CI-/Build-Schicht` (`.github/workflows/`, `tools/ci-*.sh`, Versions-Pins in
  `Makefile`/`Dockerfile`) — dieselbe Achse wie „jedes Target ist in `AGENTS.md` §4 deklariert" der
  **Gate-/Werkzeug-Schicht**; die Pfadliste wird um `.github/workflows/` ergänzt (§5).
- `Implementierung` (die `dirVocab`/`portFor`-Kopplung aus `BEO-024`) — die Regel-Engine
  (`portFor`) ist der Anker der Kopplung, nicht der Adapter; **Kern und Regeln** führt bereits
  „jede Regel hat eine `AC-FA-RULE-*` als Anker".

**Zuordnung, nicht Deklaration war falsch** — die Fehldiagnose aus dem Baseline-Zitat trifft hier
zu. Die vier Zeilen wandern beim Migrations-Folge-Slice unter den korrigierten Sub-Area-Namen; hier
wird nur die Pfadliste einer bestehenden Zeile ergänzt (§5), keine der acht Zeilen neu geschnitten.

## 5. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor.

## 6. Was bewusst nicht getan wird

- **Keine Beobachtungs-Register-Migration.** Die acht Kürzel sind Vorbedingung, nicht Ausführung —
  eigener Folge-Slice.
- **Keine Umbenennung der vier fehlklassifizierten `observations.md`-Zeilen.** Sie wandert mit der
  Migration selbst, in derselben Bewegung wie die Formänderung.
- **Kein Kürzel für „Vendored Baseline".** Sie führt keinen Modus ([`MR-006`](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert));
  eine Beobachtung über a-checks *eigene* Konvergenz kann dort nie entstehen, also braucht die
  Zeile kein Kürzel.

## 7. DoD

- [x] Acht Kürzel vergeben, aus dem etablierten Kurznamen jeder Sub-Area-Zeile abgeleitet (§3);
      Tabelle in `harness/conventions.md` trägt die neue Spalte.
- [x] Der Fund aus §4 benannt und die Pfadliste der `Gate-/Werkzeug-Schicht`-Zeile um
      `.github/workflows/` ergänzt, ohne eine neue Zeile zu schneiden.
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 8. Closure-Notiz

**Geliefert:** acht Sub-Area-Kürzel, direkt aus den etablierten Kurznamen der Modus-Deklaration
abgeleitet — keine Neuerfindung, keine Kollision. Als Nebenfund beim Nachschlagen: vier von 33
aktiven Beobachtungs-Zeilen tragen eine Sub-Area, die `conventions.md` nicht führt; alle vier
lassen sich ohne neue Zeile auf bestehende Sub-Areas zurückführen (§4).

**Lerneintrag — Form: neuer Sensor (benannt, nicht gebaut).** *Eine Freitext-Spalte, die dieselbe
Bedeutung wie eine geschlossene Deklarations-Liste trägt, driftet lautlos — jede neue Zeile ist ein
neuer Freiheitsgrad, den niemand gegen die Liste hält.* Vier von 33 Zeilen (12 %) trugen einen
Sub-Area-Namen ohne Gegenstück, über neun Slices hinweg entstanden (`slice-121` bis `slice-134`),
und keiner der Autoren — alle derselbe Agent — bemerkte es beim Schreiben. Erst die Kürzel-Vergabe,
die jede Zeile gegen die Tabelle nachschlagen musste, deckte es auf. Ein Sensor „jede
`observations.md`-Sub-Area-Zeile hat eine Entsprechung in `conventions.md`" ist mit den bestehenden
`d-check`-Modulen nicht ohne Weiteres baubar (Freitext-Zelle einer Tabelle gegen die Zellen einer
anderen Tabelle) — das ist der Grund, warum dieser Lerneintrag den Sensor **benennt statt baut**;
er gehört als Beobachtung ins Register, sobald die Migration (Folge-Slice) einen Ort dafür hat.

**Zwei beobachtbare Closure-Kriterien:**

1. `harness/conventions.md` §Modus-Deklaration pro Sub-Area trägt die Kürzel-Spalte mit acht
   befüllten Zellen und einer leeren (Vendored Baseline).
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.**

- *Die vier fehlklassifizierten `observations.md`-Zeilen sind benannt, aber nicht korrigiert* —
  Ausgang: **Folge-Slice** (Beobachtungs-Register-Migration), der sie im selben Zug umsortiert.
- *Kein Sensor gegen künftige Sub-Area-Freitext-Drift* — Ausgang: **weiter offen**, siehe
  Lerneintrag — wird bei der Migration ins Beobachtungs-Register selbst eingetragen (die neue
  Verzeichnis-Form erzwingt dort ohnehin einen Kürzel-Nachschlag pro Eintrag, was genau diese
  Klasse von Drift für Beobachtungen strukturell verhindert; für andere Freitext-Sub-Area-Spalten,
  falls es je welche gibt, bliebe sie bestehen).

**Folge-Slices:** keine vergeben. Die Beobachtungs-Register-Migration ist als nächster Schritt
benannt (§6, §8) und braucht eine eigene Kennung bei ihrer Eröffnung.

## 9. Sub-Area-Modus

Berührt wird ausschließlich **Harness-Einstieg** (`harness/conventions.md`) — Greenfield.
