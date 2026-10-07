# slice-212 — Spec-first: `shapes`-Dialekte `gomod` und `json`

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

**Lerneintrag — Form:** benannte Spec-Lücke.

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

## 1b. Messung (2026-10-07)

**Gegenstand:** die Manifeste aller lokalen Repos, die eine `.a-check.yml` tragen (also
nachweislich a-check-Konsumenten sind) — gefunden mit
`find /Development -maxdepth 3 -name .a-check.yml`, Manifeste darunter mit `find … -name go.mod
-o -name package.json -o -name build.gradle.kts`, ohne `node_modules`, `vendor`, `.git`,
`testdata`, `build`. Gezählt mit `grep` (go.mod) bzw. einer JSON-Ladung (package.json); die
Originale wurden nur gelesen.

**`go.mod` — 8 Dateien** (a-check ×2, d-check ×2, m-trace/apps/api, pg-change-feed,
pgwire-recorder, claude-ai-harness-init) — a-check selbst eingeschlossen, es trägt eine `.a-check.yml` (Korrektur nach Review F-1).

| Merkmal | Befund |
|---|---|
| Direktiven auf oberster Ebene | nur `module`, `go`, `require` |
| `require` als Block `require ( … )` | 4 Dateien mit je 2 Blöcken, **eine Abhängigkeit je Zeile** |
| `require` einzeilig | 1 Datei (a-check: `require gopkg.in/yaml.v3 v3.0.1`) |
| ohne `require` | 3 Dateien (die beiden `tools/archive-wave` und claude-ai-harness-init) |
| Kommentare | 85 Zeilen-Kommentare, **alle** `// indirect` |
| `replace`, `exclude`, `retract`, `tool`, `toolchain`, `godebug` | 0 |
| Backquote-Zeichenketten | 0 |

**`package.json` — 5 Dateien** (m-trace): Wurzel ist **ein** Objekt (Tiefe 2–3);
Abhängigkeiten stehen in den Wurzel-Mitgliedern `dependencies`, `devDependencies`,
`peerDependencies`; daneben `scripts` (beliebige Befehle) und Metadaten.

**Nebenmessung Kotlin:** derselbe Bestand führt 82 `build.gradle.kts` (d-migrate, belief-agent,
pg-change-feed). Mit dem in `v0.21.0` ausgelieferten Image, je Datei in eine Scratch-Kopie mit
einem `shapes`-Eintrag (`allow: ["nie()"]`) gelegt: **82 zerlegt, 0× Exit 2, 292 Anweisungen**.

**Was die Messung entscheidet:**

1. **Benannte Dialekte, keine konfigurierbare Lexik.** Beide Formate haben eine kleine **formale
   Grammatik** (Go-Modulreferenz §go.mod files; RFC 8259) — die Lexik lässt sich daraus
   **ableiten** statt aufzählen, und genau das fehlte beim Kotlin-Dialekt (zwei Register-Belege).
   Eine vom Konsumenten konfigurierte Lexik verschöbe dieses Risiko zu ihm; kein gemessener Fall
   braucht sie. Vorschlag: `dialect: gomod` und `dialect: json`.
2. **Zerlegungstiefe:** `go.mod` — Anweisung = Direktive auf oberster Ebene; ein Block `verb ( … )`
   ist eine Anweisung, seine Zeilen sind darin durch `;` getrennt (dieselbe Form wie ein
   Kotlin-Block). `json` — Anweisung = **Mitglied des Wurzel-Objekts** (Schlüssel und ganzer
   Wert); eine zusätzliche Abhängigkeit ändert das Mitglied `dependencies` und ist rot. Eine
   Wurzel, die kein Objekt ist, ist Exit 2.
3. **Einzeiligkeit:** JSON-Zeichenketten können kein rohes Zeilenende enthalten (RFC 8259),
   `go.mod`-Zeichenketten in der gemessenen Form auch nicht. Offen bleiben Kotlin-Roh-Strings und
   die nicht gemessenen `go.mod`-Backquotes. Vorschlag: die **Meldung** jedes `shape-*`-Befunds
   schreibt ein Zeilenende als `\n` (wie `shape-unused` seit Spezifikation 0.35.0) — der
   Vergleich bleibt byte-genau, nur die Ausgabe wird einzeilig.

## 2. Definition of Done

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md`
§Ziel-Form: Slice — ≤ 3 Liefer-Punkte.

- [x] Messung abgelegt: die `go.mod`- und `package.json`-Dateien der a-check-Konsumenten
      (Bestand unter `/Development`, mindestens `d-check` und `m-trace`) — welche Konstrukte
      vorkommen, gegen die formale Grammatik (Go-Modulreferenz, RFC 8259) abgeglichen; das
      Instrument liegt mit dem Ergebnis im Slice.
- [x] Lastenheft (Versions-Bump, Historie, CHANGELOG `[Unreleased]`) und Folge-ADR mit den
      Antworten auf die drei Fragen aus §1 — `Accepted` erst nach Maintainer-Abnahme.
- [x] Spezifikation: Lexik und Zerlegung je neuem Format, Einzeiligkeit der Befundzeile.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

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
  (1×). — **Ausgang:** *weiter offen* → Beobachtungs-Register:
  [`BEO-SPEC/vertrag-auf-einen-konsumenten-belegt`](../observations/BEO-SPEC/vertrag-auf-einen-konsumenten-belegt/observation.md)
  (jetzt 2×). Der Vertrag trägt reale Dateien; dass ihn jemand braucht, ist nicht belegt.
- **Dritter Lexik-Fall.** [`BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe`](../observations/BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe/observation.md)
  steht bei 2×; ein dritter Fund in dieser Welle macht ihn zur Lücke. Gegenmittel im Plan: Lexik
  aus der formalen Grammatik ableiten, nicht aufzählen. — **Ausgang:** *entfallen* — gestrichen
  mit Begründung: die Lexik ist diesmal aus Go-Modulreferenz und RFC 8259 abgeleitet und gegen sie
  sowie 3000+ reale Dateien geprüft; kein Review-Lauf fand einen Fall, in dem Code verschluckt
  würde. Die Funde F-8/F-9 waren Lücken der fail-closed-Abgrenzung (Kollisionen, offene Fälle),
  keine aus dem Gedächtnis aufgezählte Lexik — kein dritter Beleg. Das Risiko im **Code** führt
  slice-213 §6 weiter.
- **Einzeiligkeit trifft den Kotlin-Dialekt.** Entscheidet Frage 3 für eine Escape-Form, ändert
  sich die Ausgabe bestehender `shape-unlisted`-Befunde mit Roh-Strings (Breaking für
  Konsumenten, die Ausgaben vergleichen). — **Ausgang:** *eingetreten* → Folge-Slice slice-213
  (implementiert die umkehrbare Meldung; im CHANGELOG unter „Changed“ benannt).

## 7. Closure-Notiz

**Lerneintrag — Form: benannte Spec-Lücke.** Die Einzeiligkeit der Befundzeile war dreimal offen
geblieben; [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) legt sie
in 0.36.0 für alle `shape-*`-Befunde fest — einzeilig **und umkehrbar**, `shape-unused` mit
Art-Präfix. Der erste Entwurf war einzeilig, aber nicht umkehrbar; erst der Review zeigte, dass
dann zwei verschiedene Anweisungen eine Meldung teilen.

**Geliefert:** Vertrag für `gomod` und `json` —
[AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012) (Lastenheft 0.29.0),
[ADR-0042](../../adr/0042-shapes-benannte-dialekte-gomod-json.md) `Accepted`, Spezifikation 0.36.0;
Messung an 8 `go.mod`, 5 `package.json` und 82 `build.gradle.kts` realer Konsumenten.

**Was hat funktioniert:** Die Lexik aus der Primärquelle ableiten statt aus dem Gedächtnis — die
Go-Modulreferenz widerlegte zwei Kotlin-Annahmen (`/* */` kein Kommentar, Zeilenende signifikant),
bevor sie in den Vertrag kamen. Und wo die Quelle schweigt, fail-closed statt geraten: Der Reviewer
fand über 153 `go.mod` und 3005 `package.json` keinen einzigen Fehlalarm der neuen Exit-2-Fälle.

**Was ging anders als geplant:** Die eigene Messung schloss a-check selbst aus (F-1, HIGH) — ein
`-not -path` für das eigene Repo, obwohl es Konsument ist. Und drei Review-Läufe statt einem; die
Funde trafen vor allem **Kollisionen** (zwei Quellen, eine Normalform) — in Lexik und in der
Ausgabe gleichermaßen.

**Steering-Loop-Eintrag:** benannte Spec-Lücke —
[`BEO-SPEC/ausgabe-einzeilig-nicht-zugesichert`](../observations/BEO-SPEC/ausgabe-einzeilig-nicht-zugesichert/observation.md)
erreichte mit diesem Slice 3× und ist **verkörpert** in der Spezifikation 0.36.0, §[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
(Ausgabe-Regel). Eine Spec-Stelle trägt keinen Herkunfts-Anker; ihr Gegenstück ist die
Versions-Zeile.

**Beobachtungs-Register (`../observations/`):** `ausgabe-einzeilig-nicht-zugesichert` → 3×,
*verkörpert*; [`BEO-SPEC/vertrag-auf-einen-konsumenten-belegt`](../observations/BEO-SPEC/vertrag-auf-einen-konsumenten-belegt/observation.md)
→ 2×.

**Folge-Slices:** slice-213 (Implementierung, liegt in `open/`).

**Risiken aus §6:** alle drei mit Ausgang — *weiter offen* (Register), *entfallen* (Begründung),
*eingetreten* (slice-213).

**Drei Paarungen:** getragen von der Closure von welle-17.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-07; kein Release seit `v0.21.0`, kein Golden Set, keine Validator-Kante).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `SPEC` (Achsen 1, 2, 3 ✓) · `ADR` (1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07):
`BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe` (2×) und
`BEO-SPEC/vertrag-auf-einen-konsumenten-belegt` (1×) — beide als Risiko in §6;
`BEO-SPEC/ausgabe-einzeilig-nicht-zugesichert` (2×) — als Frage 3 in §1 aufgenommen, damit sie
diesmal entschieden wird. `ADR`: keine Treffer.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
