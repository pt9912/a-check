# slice-121 — Port-Richtung heißt `inbound`/`outbound`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Vorgabe 2026-08-30: **Ports haben `inbound`/`outbound`, Adapter haben
`driving`/`driven`.** Korrigiert
[AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch).

**Berührte Spec-Stellen:**
[AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch) ·
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) ·
[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) ·
[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Die Richtungs-Dimension trägt je Rolle das Vokabular ihrer Rolle: an `role: port` gilt
`inbound`/`outbound`, an `role: adapter` `driving`/`driven`.

## 2. Definition of Done

- [x] Die Spec-Kette trägt die Regel: [Lastenheft](../../../../spec/lastenheft.md)
      ([`AC-FA-RULE-008`](../../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch) neu gefasst, [`AC-FA-CONF-001`](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) im Schema), eine neue ADR mit
      `Supersedes` [`ADR-0012`](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) samt Index-Eintrag, und die
      [Spezifikation](../../../../spec/spezifikation.md).
- [x] Die Engine setzt es durch: `direction` wird **rollen-abhängig** validiert (Exit 2 mit
      nennender Meldung bei der falschen Vokabel), und `port-direction-mismatch` prüft eine
      **Paarung** (`driving`↔`inbound`, `driven`↔`outbound`) statt einer String-Gleichheit.
- [x] Tests und Benutzerhandbuch sind nachgezogen; die Verweise auf die abgelöste ADR sind
      aufgelöst (§3).

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

**Spec-first** ([`AGENTS.md`](../../../../AGENTS.md) §5): Lastenheft → ADR → Spezifikation → Code
→ Tests. Die ADR steht **zwischen** Lastenheft und Spezifikation, weil sie die Spezifikation
schärft und das Lastenheft nur zitiert.

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| [`spec/lastenheft.md`](../../../../spec/lastenheft.md) | update | [`AC-FA-RULE-008`](../../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch) + [`AC-FA-CONF-001`](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml), Versions-Bump, Historie |
| `docs/plan/adr/00NN-*.md` | neu | löst [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) ab; die ist `Accepted` und damit immutabel (§3.5) |
| [`docs/plan/adr/README.md`](../../adr/README.md) | update | Index-Pflicht (§5) |
| [`spec/spezifikation.md`](../../../../spec/spezifikation.md) | update | Schema und Regel-Auswertung |
| `internal/adapter/driven/config/config.go` | update | `validDirection` rollen-abhängig |
| `internal/hexagon/core/rules.go` | update | Paarung statt Gleichheit |
| `internal/hexagon/core/model.go` | update | Doc-String der `Direction`-Zusage |
| Tests, [`docs/user/benutzerhandbuch.md`](../../../../docs/user/benutzerhandbuch.md) | update | Belege und Benutzer-Doku |

**Auszuführende Gates:** `make gates` (tragend `test`, `coverage-gate`, `doc-check`,
`arch-check`), `make image-test`, zum Abschluss `make verify`.

### Der Nebenaufwand, der leicht übersehen wird

`matrix` verbietet Verweise auf ein Dokument mit Status `superseded`
([`.d-check.yml`](../../../../.d-check.yml), `status.forbidden`). **Sechs** Dateien nennen die
abzulösende ADR — der Index, die neue ADR selbst, eine weitere ADR, zwei Planning-Dateien und
das CHANGELOG. Sobald der Status wechselt, sind das Befunde, bis sie aufgelöst sind. Das Werkzeug
kennt dafür `matrix.allow-supersede-lineage` (opt-in: die ablösende Datei darf auf ihr abgelöstes
Ziel zeigen); ob das reicht oder ob einzelne Verweise umzubiegen sind, wird **gemessen**, nicht
angenommen — der Lauf sagt es.

### Was die Regel für die Engine bedeutet

Heute ist `port-direction-mismatch` eine **String-Gleichheit**: Adapter-Richtung ≠ Port-Richtung
⇒ Befund. Das funktioniert nur, solange beide Seiten dasselbe Vokabular führen. Mit zwei
Vokabularen wird daraus eine **Paarung** — `driving` gehört zu `inbound`, `driven` zu `outbound`.
Die Regel-Aussage bleibt dieselbe: ein Treiber-Adapter spricht nur eingehende Ports.

## 4. Trigger

**Start:** eingetreten — Maintainer-Vorgabe, keine Vorbedingung.

**Rückführungen:**

- `in-progress` → `next`: falls die Auflösung der Supersede-Verweise (§3) für sich einen Slice
  füllt. Dann sind es zwei.

## 5. Closure-Trigger

Spec-Kette steht, Engine setzt durch, Tests und Handbuch nachgezogen, Gates grün.

**Was bewusst nicht getan wird:** die alte Schreibweise am Port **still weiter akzeptieren**.
Maintainer-Entscheid: `direction: driving` an `role: port` wird **Exit 2** mit nennender Meldung.
Ein Alias wäre nicht durchgesetzt, sondern nur dokumentiert — a-check kennt kein Warn-Level, die
alte Schreibweise bliebe unbemerkt. Ebenso wenig wird
[slice-013](../open/slice-013-driving-driven-vertiefung.md) mitbearbeitet; dass dessen
Vertagungs-Argument sich ändert, ist ein Folge-Slice.

## 6. Risiken und offene Punkte

- *Bestehende Konsumenten-Configs mit `direction: driving` an Port-Schichten brechen* —
  **Ausgang:** **gestrichen mit Begründung**: das war nie ein Risiko, sondern die **Absicht**
  des Maintainer-Entscheids. Die Meldung nennt Schicht, Wert **und** die für diese Rolle gültige
  Menge; die Umstellung ist eine Zeile je Schicht. Belegt an den eigenen Tests: drei brachen beim
  ersten Lauf, genau die mit der alten Vokabel.
- *Die Supersede-Verweise könnten mehr Arbeit sein als die Sache selbst* — **Ausgang:**
  gestrichen mit Begründung: **null** Befunde, gemessen statt angenommen. Die sechs Verweiser
  liegen entweder in keiner `matrix`-Klasse (`CHANGELOG.md`, der ADR-Index) oder in einer ohne
  Regel auf `adr` (die Slice-Klasse hat keine), und der eine, der beides hätte
  ([`ADR-0019`](../../adr/0019-adapterseg-root-subeinheit.md)), ist von `matrix.exempt-paths` grandfathered. **Nebenbefund:** `status.forbidden`
  ist damit im heutigen Bestand für diesen Fall inert.
- *Die Paarung ist eine zweite Stelle, an der die Zuordnung steht* — **Ausgang:** weiter offen
  im **Beobachtungs-Register** als benannte Grenze: `dirVocab` (Config-Seite) und `portFor`
  (Regel-Seite) sind zwei Funktionen in zwei Paketen, die dieselbe Zuordnung tragen. Kein Test
  hält sie gegeneinander — er könnte es nur über die Paket-Grenze hinweg, und die trennt die
  Hexagon-Schichten absichtlich.

## 7. Closure-Notiz

**Geliefert:** Die Richtungs-Dimension trägt je Rolle ihr Vokabular: `inbound`/`outbound` an
`role: port`, `driving`/`driven` an `role: adapter`. Die Spec-Kette steht (Lastenheft 0.25.0,
[ADR-0036](../../adr/0036-port-richtung-inbound-outbound.md) mit `Supersedes` [`ADR-0012`](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md),
Spezifikation), die Engine setzt es durch — rollen-abhängige Validierung mit nennender Meldung
und `port-direction-mismatch` als **Paarung** statt String-Gleichheit —, Tests und Handbuch sind
nachgezogen.

**Lerneintrag — Form: benannte Spec-Lücke.** *Eine Entscheidung, die zwei Alternativen abwägt,
belegt damit nicht, dass die Frage vollständig gestellt war.*
[ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) hat sauber
gearbeitet: orthogonale Dimension gegen Subtyp-Rollen, mit Begründung und verworfener
Alternative. **Beide Alternativen trugen dasselbe Vokabular** — die Wahl lag zwischen zwei
Bauformen, nie zwischen zwei Benennungen. Der Text nannte die richtige Vokabel sogar
(*„`driving` = primär/inbound"*) und ließ trotzdem nur die eine Hälfte als Wert zu. *Weil* eine
Abwägung nach Vollständigkeit aussieht, ist die ungestellte Frage schwerer zu sehen als die
falsch beantwortete: es steht ja etwas da, mit Gründen. Was hier fehlte, war kein Argument,
sondern eine **dritte Option auf der Liste**.

**Drei beobachtbare Closure-Kriterien:**

1. Die Zuordnung liegt je Seite an **einer** Stelle: `dirVocab(role)` speist Validierung und
   Fehlermeldung, `portFor(adapterDir)` die Regel. Beide sind Funktionen, keine Paket-Variablen
   — das Lint-Profil verbietet globale Zustandsträger, und eine Suppression wäre nach
   [`AGENTS.md`](../../../../AGENTS.md) §3.2 keine Option gewesen.
2. Zwei Tabellen-Tests halten **beide** Richtungen der Zusage: sechs Kombinationen für die
   Validierung (drei grün, drei rot, jede rote prüft die genannte Menge) und sechs für die
   Paarung (zwei Paare, zwei Nicht-Paare, zwei mit fehlender Richtung). Ohne die grünen Fälle
   wäre eine Regel, die alles meldet, von einer korrekten nicht zu unterscheiden.
3. Die Validierung greift gegen die **effektive** Rolle, also auch bei Namens-Inferenz. Belegt
   dadurch, dass `core.InferRole` dafür exportiert wurde statt die Abbildung im Config-Adapter
   zu kopieren — zwei Kopien einer Namenskonvention driften, und die Drift wäre still.

**Was über den Auftrag hinausging, und warum es benannt ist:** eine Richtung an einer Schicht
**ohne** Port-/Adapter-Rolle lud bis 0.24.0 stillschweigend („inert"). Mit rollen-abhängigem
Vokabular ist unbeantwortbar, welches dort gälte — sie ist jetzt ein Fehler. Das steht in der
Anforderung und in der ADR, nicht nur im Code; ein Test, der vorher „lädt inert" hieß, heißt
jetzt „wird abgewiesen".

**Offene Risiken und ihr Ausgang:** der erste gestrichen mit Begründung (die Wirkung war die Absicht), der zweite
gestrichen mit Begründung, der dritte weiter offen im Register.

**Beobachtungs-Register:** `BEO-024` neu angelegt (Implementierung, 1×, Beleg slice-121): dieselbe
Zuordnung liegt in zwei Paketen, und die Hexagon-Schichtung verhindert den Test, der sie
gegeneinander hielte.

**Folge-Slices:** [slice-013](../open/slice-013-driving-driven-vertiefung.md) — sein
Vertagungs-Argument für Teil A stützt sich auf eine Messung, die mit dem alten Vokabular erhoben
wurde; ob die Auto-Inferenz unter der neuen Terminologie mehr trifft, ist eine eigene Messung.
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden die **Spec-Straten** (Lastenheft, ADR,
Spezifikation) und die **Implementierung** (`internal/adapter/driven/config/`,
`internal/hexagon/core/`).

**Vorgelagert — offene Beobachtungen sichten:** keine Treffer in den berührten Sub-Areas;
`BEO-023` (Prüfer mit leerer Prüfmenge) liegt in der Gate-Schicht.

Alle berührten Sub-Areas GF.
