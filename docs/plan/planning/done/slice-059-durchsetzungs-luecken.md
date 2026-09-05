# slice-059 — Durchsetzungs-Lücken: Gate-Liste, Befund-Maskierung, Schritt 9

**Status:** der Zustand ist das **Verzeichnis** dieser Datei
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5) — dieses Feld führt ihn bewusst **nicht** doppelt.
**Welle:** welle-12-regelwerk-migration.
**Deckt:** **F-1** aus
[`2026-07-26-etappe-f-slice-057.md`](../../../reviews/2026-07-26-etappe-f-slice-057.md), **F-4**
aus [`2026-07-26-etappe-d-slice-052-053-054.md`](../../../reviews/2026-07-26-etappe-d-slice-052-053-054.md)
und **F-2** aus [`2026-07-26-etappe-e-slice-050-051.md`](../../../reviews/2026-07-26-etappe-e-slice-050-051.md).
**Bezug:** zweiter von drei Fix-Schnitten der Review-Serie; Folge von
[slice-058](../done/slice-058-sensor-praezision.md).
[Roadmap](../in-progress/roadmap.md) — Verweis zustandsunabhängig (Lehre aus slice-058).

---

## 1. Auslöser

Drei Befunde an der Durchsetzungs- und Aggregat-Ebene, jeder mit einem Lauf belegt:

| Fund | Beobachtung (gemessen) |
|---|---|
| **R-057-F1** | Die `GATES`-Liste des PreToolUse-Guard ist hartcodiert und kennt `doc-immutable` nicht — in [`AGENTS.md`](../../../../AGENTS.md) §4 ausdrücklich als **CI-durchgesetzt** geführt. Live belegt: `make doc-immutable \| tail -1` lief **ungehindert** durch; `make` brach mit Fehler 2 ab, der Exit-Code der Pipeline war der von `tail`. Genau der Vorgang, gegen den SL-001 antritt. Nichts gleicht die Liste gegen die Target-Tabelle ab. |
| **R-052-F4** | `verify` hängt seine drei Teil-Sensoren als sequenzielle Prerequisites; `make` bricht beim ersten roten ab. Belegt: mit einem Slice-Form- **und** einem AC-Form-Verstoß zugleich meldete `make verify` nur die Slice-Form-Befunde. Kein False-Green, aber unvollständige Diagnose in der Schicht, die vor der „fertig"-Meldung Auskunft geben soll. |
| **R-050-F2** | Das Workflow-Skelett stellt den Gate-Lauf als **Schritt 6** vor den Lifecycle-`git mv` als **Schritt 9**; der Verschiebe-Commit ist damit per Ablauf ungeprüft. Genau dort entstanden die beiden `doc-check`-roten Slice-Endstände der Kette. Ein Schritt „Verweise prüfen" fehlt, obwohl das Skelett die andere bekannte Fehlerklasse (Pipe) ausdrücklich adressiert. |

[slice-058](../done/slice-058-sensor-praezision.md) hat den dritten Punkt unfreiwillig bestätigt:
derselbe Fehler trat dort **zweimal** auf, einmal davon beim Schreiben der Notiz über ihn.

## 2. Betroffene Module

- [`.claude/hooks/pretooluse-command-guard.sh`](../../../../.claude/hooks/pretooluse-command-guard.sh)
  — Gate-Liste plus Drift-Prüfung im `--selftest` (R-057-F1).
- [`Makefile`](../../../../Makefile) — `verify` sammelt Befunde, statt beim ersten abzubrechen
  (R-052-F4).
- [`.claude/commands/slice.md`](../../../../.claude/commands/slice.md) — Schritt 9 (R-050-F2).
- `docs/plan/steering-loop.md` — `SL-002` bekommt den Beleg für seinen
  Guide-Kandidaten 1; der Eintrag bleibt stehen, wie die Pflege-Regel es verlangt.

**Zwei Schichten:** Gate-/Werkzeug-Schicht (`Makefile`, `.claude/`) und Planungs-Doku
(`steering-loop.md`).

## 3. Auszuführende Gates

`make gates` (enthält `guard-selftest`) und `make verify`, Ausgabe je in eine Datei, Exit-Code
getrennt geprüft.

Negativ-Proben, ohne die keiner der drei Fixes belegt wäre:

1. **Gate-Liste** — `make doc-immutable | tail -1` muss ab jetzt **abgelehnt** werden; ein
   Nicht-Prüf-Target in einer Pipe (`make help | head`) muss weiter durchgehen. Zusätzlich muss
   die Drift-Prüfung im `--selftest` rot werden, wenn ein deklariertes Prüf-Target in der Liste
   fehlt.
2. **Befund-Maskierung** — bei zwei gleichzeitigen Verstößen in verschiedenen Teil-Sensoren muss
   `make verify` **beide** melden und trotzdem mit ≠ 0 enden.
3. **Schritt 9** — keine maschinelle Probe; der Guide ist *inferential feedforward* und behauptet
   nichts, was ein Sensor decken müsste. Der Beleg ist die Textstelle selbst.

## 4. Was bewusst nicht getan wird

- **Kein Sensor für SL-002.** Der Guide (Schritt 9) ist die billige Hälfte; die *computational*
  Hälfte — eine Prüfung, die vor dem `git mv` meldet, welche Verweise brechen — bleibt offen und
  gehört zu `make verify`. Sie braucht eine eigene Entscheidung über den Prüfumfang (slice-058 hat
  gezeigt, dass die **präfixlose** Nachbardatei der kritische Fall ist, nicht der Randfall) und
  würde diesen Slice über zwei Schichten hinaus dehnen.
- **Keine dynamische Ableitung der Gate-Liste aus dem `Makefile`.** Der Guard läuft vor jedem
  Tool-Call; ein Dateizugriff pro Aufruf wäre teuer und fehleranfällig. Stattdessen bleibt die
  Liste explizit und bekommt einen Drift-Wächter, der im ohnehin laufenden `--selftest` anschlägt.
- **Keine Änderung an `regelwerk-check`s Gate-Status** und **nichts an der Planungs-Doku außer
  `SL-002`** — Letzteres ist slice-060.

## 5. DoD

- [x] Die Gate-Liste des Guard deckt alle deklarierten Prüf-Targets ab, und eine Drift-Prüfung im
      `--selftest` schlägt an, wenn ein neues hinzukommt; belegt durch je einen Lauf für Ablehnung,
      Durchlass und Drift (R-057-F1).
- [x] `make verify` führt alle drei Teil-Sensoren aus und meldet **alle** Befunde; belegt durch
      einen Lauf mit zwei gleichzeitigen Verstößen in verschiedenen Sensoren (R-052-F4).
- [x] Das Workflow-Skelett nennt die Verweis-Prüfung vor dem Lifecycle-`git mv`, und `SL-002`
      trägt den Beleg für den gelieferten Guide-Kandidaten (R-050-F2). `make gates` und
      `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** die `GATES`-Liste des Guard deckt jetzt alle deklarierten Prüf-Targets ab und hat
einen **Drift-Wächter** im `--selftest`, der jedes neue Target einfordert; `make verify` führt
seine drei Teil-Sensoren als Sequenz im selben Rezept aus und meldet alle Befunde statt nur den
ersten; Schritt 9 des Workflow-Skeletts verlangt die Verweis-Prüfung vor dem `git mv` samt
Kommando und benennt die zustandsunabhängige Form. `SL-002` trägt den
Beleg für den gelieferten Guide-Kandidaten und steht auf **halb gebaut**.

**Lerneintrag — Form: neuer Sensor.**
> **Eine hartcodierte Liste, die einen beweglichen Bestand abbildet, ist eine Momentaufnahme —
> sie braucht einen Wächter, sonst driftet sie lautlos.** Die `GATES`-Liste des Guard war beim
> Anlegen vollständig und wurde es durch spätere Targets nicht mehr; `doc-immutable` fiel heraus,
> *weil* nichts sie gegen die Target-Tabelle hielt. Das Repo hatte den Mechanismus dafür längst —
> `gate-consistency` gleicht Doku gegen `Makefile` ab —, nur eben nicht für diese Liste. Prüfsatz:
> *wenn eine Aufzählung im Code einen Bestand spiegelt, der anderswo wächst, gehört die Prüfung
> „spiegelt sie ihn noch?" in denselben Commit wie die Aufzählung.* Der Wächter kostete zwölf
> Zeilen und hätte den Befund verhindert.

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. Jeder der drei Fixes ändert ein messbares Ergebnis gegenüber dem Vorzustand, je mit Lauf
   belegt: `make doc-immutable | tail -1` wird **abgelehnt** (im Review lief derselbe Aufruf
   ungehindert durch); ein temporäres Prüf-Target `probe059-check` im `Makefile` macht
   `guard-selftest` **rot** (EXIT=2) und nach Rücknahme wieder grün; `make verify` meldet bei
   gleichzeitigem Slice-Form- **und** AC-Form-Verstoß **beide** Bereiche (2 + 4 Befunde) statt
   nur des ersten.

**Was der Guide nicht leistet, ausdrücklich:** Schritt 9 ist *inferential feedforward*. Nach
Modul 09 bleibt SL-002 damit **halb durchgesetzt** — dieselbe Klasse wie B-11 vor slice-049 und
wie SL-001 vor slice-057. Beide Male hat erst die computational Hälfte den Rückfall beendet, und
slice-058 hat den Rückfall hier bereits zweimal vorgeführt. Der Sensor bleibt darum als benannter
offener Punkt stehen, nicht als erledigt.

**Folge-Slices:** slice-060 (Belegtreue der Planungs-Doku: `SL-003` für das Etiketten-Muster,
slice-048-Korrekturen, Status-Felder). Offen aus diesem Slice: der SL-002-Sensor.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
