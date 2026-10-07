# slice-212 — Spec-first: generischer `shapes`-Dialekt

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-17 — [Welle-Plan](../welle-17-shapes-generischer-dialekt.md).

**Bezug:** [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012) (Out-of-Scope:
„ein generischer Dialekt"), [ADR-0041](../../adr/0041-shapes-sollform-je-datei.md) Punkt 10 und
Re-Evaluierungs-Trigger, [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** `spec/lastenheft.md` [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012)
(Out-of-Scope und Dialekt-Menge) · `spezifikation.md`
§[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema),
§[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion),
§[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „machen wir weiter", 2026-10-07).

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der Vertrag für die Prüfung von `go.mod` und `package.json` per `shapes` steht
abnahmefähig — gemessen an realen Manifesten der a-check-Konsumenten, aus der formalen Grammatik
der Formate abgeleitet, als Lastenheft-Änderung, Folge-ADR zu
[ADR-0041](../../adr/0041-shapes-sollform-je-datei.md) Punkt 10 und Spezifikation.

**Fragen, die dieser Slice zur Abnahme vorlegt** (Antwort mit Messung, nicht vorweg):

1. **Konfigurierbar oder benannt?** Ein „generischer" Dialekt mit frei konfigurierbaren
   Kommentar-, Zeichenketten- und Trennzeichen — oder **benannte** Dialekte (`gomod`, `json`) mit
   fest im Werkzeug verankerter Lexik? Der Kotlin-Fall zeigte, dass eine aus dem Gedächtnis
   aufgezählte Lexik Code verschlucken kann (zwei Belege im Register); eine vom Konsumenten
   konfigurierte Lexik verlagert dieses Risiko zu ihm.
2. **Zerlegungstiefe bei JSON.** Ein `package.json` ist **eine** Anweisung auf oberster Ebene;
   verglichen als Ganzes wäre `allow-statements` gleich `exact`. Wo liegt die Einheit — Mitglieder
   des Wurzel-Objekts, Einträge in `dependencies`, ein Pfad-Ausdruck?
3. **Einzeiligkeit der Befundzeile.** Mehrzeilige Anweisungen sind bei JSON der Normalfall. Die
   Frage ist im Register zweimal offen geblieben; dieser Slice **entscheidet** sie für alle
   `shape-*`-Befunde.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Produkt-Code.** *Schicht-Abgrenzung*: Spec-first; der Code folgt in slice-213 nach Abnahme.
- **XML und TOML.** *Bestand bleibt bewusst stehen*: kein belegter Bedarf (Welle-Plan §6).
- **Der Kotlin-Dialekt.** *Bestand bleibt bewusst stehen*: geliefert und reviewt; ändert sich
  hier nur, wenn die Einzeiligkeits-Entscheidung (Frage 3) ihn trifft — dann als benannte
  Plan-Änderung.

## 2. Definition of Done

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md`
§Ziel-Form: Slice — ≤ 3 Liefer-Punkte.

- [ ] Messung abgelegt: die `go.mod`- und `package.json`-Dateien der a-check-Konsumenten
      (Bestand unter `/Development`, mindestens `d-check` und `m-trace`) — welche Konstrukte
      vorkommen, gegen die formale Grammatik (Go-Modulreferenz, RFC 8259) abgeglichen; das
      Instrument liegt mit dem Ergebnis im Slice.
- [ ] Lastenheft (Versions-Bump, Historie, CHANGELOG `[Unreleased]`) und Folge-ADR mit den
      Antworten auf die drei Fragen aus §1 — `Accepted` erst nach Maintainer-Abnahme.
- [ ] Spezifikation: Lexik und Zerlegung je neuem Format, Einzeiligkeit der Befundzeile.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

Regeln dieser Sektion: Baseline-Regelwerk `grundlagen-bootstrap.md` §Was ist eine Sub-Area?

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| Messung (Abschnitt in diesem Slice + Instrument) | neu | Grundlage für die drei Fragen |
| `spec/lastenheft.md` | update | [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012): Dialekt-Menge, Out-of-Scope, Akzeptanzkriterien je Format |
| `docs/plan/adr/<NNNN>-…md` + Index | neu | Folge-ADR zu [ADR-0041](../../adr/0041-shapes-sollform-je-datei.md) Punkt 10 |
| `spec/spezifikation.md` | update | Schema, Lexik/Zerlegung, Befundform |
| `CHANGELOG.md` | update | Vertrag berührt |

## 4. Trigger

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md`
§Trigger je Lifecycle-Übergang und WIP-Limit.

**Start** (`next` → `in-progress`): WIP-Limit frei, welle-17 eröffnet.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): wenn die Messung zeigt, dass `go.mod` und `package.json`
  keine gemeinsame Zerlegung tragen — dann je Format ein eigener Spec-Slice.
- `in-progress` → `open` (blockiert): der Maintainer will vor der Abnahme einen realen
  Konsumenten-Bedarf sehen.

## 5. Closure-Trigger

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md`
§Closure- und Lerneintrag-Regeln.

DoD vollständig; Folge-ADR `Accepted`; `make gates` und `make verify` Exit 0; Closure-Notiz mit
Lerneintrag.

## 6. Risiken und offene Punkte

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md`
§Offene Risiken werden bei Closure aufgelöst.

- **Kein zweiter Konsument belegt den Bedarf.** Die Welle startet auf Maintainer-Wort; der Vertrag
  kann an Formaten gebaut werden, die niemand prüfen will. Verwandt mit
  [`BEO-SPEC/vertrag-auf-einen-konsumenten-belegt`](../observations/BEO-SPEC/vertrag-auf-einen-konsumenten-belegt/observation.md)
  (1×). — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*
- **Dritter Lexik-Fall.** [`BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe`](../observations/BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe/observation.md)
  steht bei 2×; ein dritter Fund in dieser Welle macht ihn zur Lücke. Gegenmittel im Plan: Lexik
  aus der formalen Grammatik ableiten, nicht aufzählen. — **Ausgang:** *(bei Closure zuzuweisen)*
- **Einzeiligkeit trifft den Kotlin-Dialekt.** Entscheidet Frage 3 für eine Escape-Form, ändert
  sich die Ausgabe bestehender `shape-unlisted`-Befunde mit Roh-Strings (Breaking für
  Konsumenten, die Ausgaben vergleichen). — **Ausgang:** *(bei Closure zuzuweisen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `SPEC` (Achsen 1, 2, 3 ✓) · `ADR` (1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07):
`BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe` (2×) und
`BEO-SPEC/vertrag-auf-einen-konsumenten-belegt` (1×) — beide als Risiko in §6;
`BEO-SPEC/ausgabe-einzeilig-nicht-zugesichert` (2×) — als Frage 3 in §1 aufgenommen, damit sie
diesmal entschieden wird. `ADR`: keine Treffer.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
