# slice-121 — Port-Richtung heißt `inbound`/`outbound`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Vorgabe 2026-08-30: **Ports haben `inbound`/`outbound`, Adapter haben
`driving`/`driven`.** Korrigiert
[AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch).

**Berührte Spec-Stellen:**
[AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch) ·
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

- [ ] Die Spec-Kette trägt die Regel: [Lastenheft](../../../../spec/lastenheft.md)
      ([`AC-FA-RULE-008`](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch) neu gefasst, [`AC-FA-CONF-001`](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) im Schema), eine neue ADR mit
      `Supersedes` [`ADR-0012`](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md) samt Index-Eintrag, und die
      [Spezifikation](../../../../spec/spezifikation.md).
- [ ] Die Engine setzt es durch: `direction` wird **rollen-abhängig** validiert (Exit 2 mit
      nennender Meldung bei der falschen Vokabel), und `port-direction-mismatch` prüft eine
      **Paarung** (`driving`↔`inbound`, `driven`↔`outbound`) statt einer String-Gleichheit.
- [ ] Tests und Benutzerhandbuch sind nachgezogen; die Verweise auf die abgelöste ADR sind
      aufgelöst (§3).

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

**Spec-first** ([`AGENTS.md`](../../../../AGENTS.md) §5): Lastenheft → ADR → Spezifikation → Code
→ Tests. Die ADR steht **zwischen** Lastenheft und Spezifikation, weil sie die Spezifikation
schärft und das Lastenheft nur zitiert.

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| [`spec/lastenheft.md`](../../../../spec/lastenheft.md) | update | [`AC-FA-RULE-008`](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch) + [`AC-FA-CONF-001`](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml), Versions-Bump, Historie |
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

- *Bestehende Konsumenten-Configs mit `direction: driving` an Port-Schichten brechen* — das ist
  der **beabsichtigte** Effekt des Entscheids, und die Meldung nennt Schicht, Wert und die
  gültige Menge. **Ausgang:** <bei Closure>
- *Die Supersede-Verweise könnten mehr Arbeit sein als die Sache selbst* — **Ausgang:** <bei Closure>
- *Die Paarung ist eine zweite Stelle, an der die Zuordnung steht (neben der Validierung)* —
  zwei Orte, die auseinanderlaufen können. **Ausgang:** <bei Closure>

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5.)_

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>**

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden die **Spec-Straten** (Lastenheft, ADR,
Spezifikation) und die **Implementierung** (`internal/adapter/driven/config/`,
`internal/hexagon/core/`).

**Vorgelagert — offene Beobachtungen sichten:** keine Treffer in den berührten Sub-Areas;
`BEO-023` (Prüfer mit leerer Prüfmenge) liegt in der Gate-Schicht.

Alle berührten Sub-Areas GF.
