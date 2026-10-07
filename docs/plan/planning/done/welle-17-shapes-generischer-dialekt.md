# Welle welle-17: Benannte `shapes`-Dialekte `gomod` und `json`

**Lifecycle:** Diese Datei entsteht bei der **Eröffnung** der Welle und liegt
flach unter `docs/plan/planning/`; bei Closure wandert sie per `git mv` nach
`done/` (neben ihre `welle-17-results.md`). Der Zustand ist die
Verzeichnis-Position — kein Status-Feld.

**Zielmeilenstein:** kein Meilenstein-Bezug.

**Verantwortlich:** Claude (Planner); Abnahme beim Maintainer.
**Datum:** 2026-10-07.

---

## 1. Welle-Ziel

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wann Arbeit eine Welle braucht.

Die Sollform je Datei ([AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012)) prüft
außer Kotlin auch die Build-Manifeste der übrigen Konsumenten-Sprachen — Leitfälle `go.mod` und
`package.json` aus realen a-check-Konsumenten. Gespiegelt an den Gegenprobe-Fällen der Anforderung, übertragen auf die neuen Formate:
eine zusätzliche Abhängigkeit ist rot, gleich in welcher Schreibweise; Kommentar und Formatierung
ändern das Urteil nicht; eine nicht zerlegbare Datei ist Exit 2.

## 2. Trigger (Welle startet)

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Roadmap-Regeln — beobachtbar, und kein Ergebnis dieser Welle.

- welle-16-shapes-sollform liegt in `done/` (erste Hälfte des Roadmap-Triggers).
- **Maintainer-Anweisung** (dieses Gespräch, 2026-10-07) **statt** der zweiten Hälfte „ein zweiter
  Konsument mit Nicht-Kotlin-Manifest". Die Ersetzung ist eine Umplanung und steht im Drift-Log der
  [Roadmap](../in-progress/roadmap.md); der fehlende Konsumenten-Beleg ist als Risiko in slice-212
  geführt, nicht übergangen.

## 3. Closure-Trigger (Welle schließt)

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wann Arbeit eine Welle braucht — das *Mehr* ist der Release-Beleg und die Gegenprobe an realen
Manifesten gegen das veröffentlichte Image.

- Alle Slices aus §4 liegen in `done/`.
- `make ci` und `make verify` Exit 0 auf dem finalen Stand — Ausgabe in eine Datei, Exit-Code
  getrennt geprüft (Replay-Ersatz laut [MR-028](../../../../harness/conventions.md#mr-028)).
- Ein Release-Tag mit dem neuen Dialekt ist veröffentlicht und nach
  [`docs/user/releasing.md`](../../../user/releasing.md) re-gepinnt.
- **Gegenprobe an realen Manifesten:** je eine `shapes`-Konfiguration für das `go.mod` und das
  `package.json` eines realen a-check-Konsumenten (Kopie, nicht das Original) lädt gegen das
  veröffentlichte Image, ist grün, und wird rot, sobald eine Fremdabhängigkeit eingefügt wird — der
  **Befund** wird gesehen, nicht nur der Exit-Code.
- Ergebnis-Notiz `done/welle-17-results.md` geschrieben.

## 4. Slices in dieser Welle

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md`
§Lifecycle als State Machine — der Zustand eines Slice ist sein Verzeichnis.

| Slice | Titel | Bezug |
|---|---|---|
| [slice-212](slice-212-generischer-dialekt-spec-first.md) | Spec-first: Messung an realen Manifesten, Lastenheft-CR, Folge-ADR, Spezifikation — zur Abnahme | [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012), [ADR-0041](../../adr/0041-shapes-sollform-je-datei.md) Punkt 10 |
| [slice-213](slice-213-generischer-dialekt-implementierung.md) | Dialekt `gomod` und einzeilige, umkehrbare Meldung | [ADR-0042](../../adr/0042-shapes-benannte-dialekte-gomod-json.md) |
| [slice-214](slice-214-shapes-dialekt-json.md) | Dialekt `json` | [ADR-0042](../../adr/0042-shapes-benannte-dialekte-gomod-json.md) |

**Reihenfolge ist Abhängigkeit:** slice-213 startet erst, wenn slice-212 in `done/` liegt und
der Maintainer den Vertrag abgenommen hat. Ob slice-213 nach der Abnahme in zwei Slices zerfällt
(je Format), entscheidet die Größenregel am abgenommenen Vertrag — nicht dieser Plan.

## 5. Abhängigkeiten

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Roadmap-Struktur: fünf Abschnitte.

- Blockiert: keine.
- Wird blockiert von: keiner Welle (welle-16 ist geschlossen).

## 6. Out-of-Scope für diese Welle

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Eröffnung Schritt 1.

- **XML (`pom.xml`) und TOML (`Cargo.toml`, `pyproject.toml`).** Kein a-check-Konsument hat dafür
  einen belegten Bedarf; XML braucht eine andere Zerlegung (Elemente, nicht Anweisungen). Kommen
  sie, ist das eine eigene Welle.
- **Semantik der Manifeste.** Wie bei Kotlin: geprüft wird Text nach Normalisierung, keine
  Versions-Auflösung, kein Lockfile, kein Paketmanager
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
- **Änderungen an der Lexik und Zerlegung des Kotlin-Dialekts.** Er ist geliefert und reviewt; die
  neuen Dialekte stehen neben ihm. Die **Ausgabeform** der Befunde (einzeilige Meldung) betrifft
  ihn dagegen mit — sie gilt für alle `shape-*`-Befunde.
- **Ein Platzhalter für „genau eine Zeichenkette"** in Literal-Einträgen — in welle-16 als
  möglicher eigener Umfang benannt, nicht beauftragt.

**Plan-Änderung 2026-10-07 (Review slice-212, F-7):** Ziel und Abgrenzung folgen der Messung in
slice-212 — benannte Dialekte mit je eigener, aus der Grammatik abgeleiteter Lexik; die einzeilige,
umkehrbare Meldung gilt für alle `shape-*`-Befunde.

## 7. Closure-Notiz

Ergebnis: `welle-17-results.md` — Geschwister im Ruheort `done/`
Zähler: `../observations/` — das Beobachtungs-Register, eine Ebene über dem Ruheort
