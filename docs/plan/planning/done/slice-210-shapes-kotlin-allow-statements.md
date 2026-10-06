# slice-210 — `shapes:` mit Dialekt `kotlin` und `mode: allow-statements`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-16 — [Welle-Plan](../welle-16-shapes-sollform.md).

**Bezug:** [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012),
[ADR-0041](../../adr/0041-shapes-sollform-je-datei.md) (beide aus slice-209);
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(`--print-config`-Gerüst),
[AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus).

**Berührte Spec-Stellen:** `spezifikation.md` §[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema), §[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
(in slice-209 geschrieben, hier umgesetzt — keine Spec-Änderung).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „mach weiter bis zum Release", 2026-10-06).

**Autor:** Claude. **Datum:** 2026-10-06.

**Lerneintrag — Form:** benannte Spec-Lücke.

---

## 1. Ziel und Abgrenzung

**Ziel:** Ein Adopter kann `build.gradle.kts` eines Moduls per `shapes:` mit
`dialect: kotlin` und `mode: allow-statements` auf eine Positivliste festlegen;
jede nicht gelistete Anweisung meldet `shape-unlisted` mit `datei:zeile`,
Exit 1 — alle Gegenprobe-Fälle des CR sind Tests.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **`mode: exact` und `shape-unused`.** *Ein Folge-Slice übernimmt es*:
  slice-211. Beide bauen auf Normalisierer und Zerleger dieses Slice auf, sind
  aber einzeln lieferbar; zusammen wären es mehr als drei Liefer-Punkte.
- **Weitere Dialekte.** *Ein anderer Vorgang*: generischer Dialekt, Roadmap
  *Nächste Wellen*, gated auf einen zweiten Konsumenten.
- **Änderung des Vertrags.** *Schicht-Abgrenzung*: Lastenheft, ADR und
  Spezifikation stehen nach slice-209; wer hier eine Vertragslücke findet,
  hält an und geht über eine Plan-Änderung, nicht über einen still
  erweiterten Code-Pfad.

  **Plan-Änderung 2026-10-06 (nach dem Review, vor dem Code):** Der Review
  fand eine solche Lücke (F-2): In Kotlin ist `$` vor einem
  Backtick-Bezeichner auch **innerhalb** einer Zeichenkette eine Vorlage;
  [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion) Schritt 1 kannte nur `${`. Ein `"` im Bezeichner beendete
  die Zeichenkette zu früh, ein späteres `/*` schluckte Code. Dieser Slice
  nimmt die Spec-Präzisierung mit (Spezifikation 0.34.0, Schritt 1 um die
  Backtick-Vorlage ergänzt) — sie schärft die Lexik, ändert keine Zusage des
  Lastenhefts und berührt [ADR-0041](../../adr/0041-shapes-sollform-je-datei.md) nicht (die nennt „vollständige Lexik"
  als Gegenmittel). Ebenso mit: die Präzisierung, dass byte-gleiche
  Befundzeilen generell einmal ausgegeben werden (F-4).
- **Graph-Ausgabe.** *Bestand bleibt bewusst stehen*: `--print-graph` zeigt
  Schichten und Kanten; eine Datei-Sollform ist keine Kante (dieselbe
  Begründung wie bei `constructs`).

## 2. Definition of Done

- [x] Normalisierer und Anweisungs-Zerleger `kotlin` (Kommentare inkl.
      verschachtelter Block-Kommentare, Zeichenketten inkl. Roh-Strings,
      Dollar-Präfix und `${…}`-Vorlagen samt maskiertem `\$`, Leerraum-Faltung, Anweisungsgrenze, Zeilen-Mapping),
      Exit 2 bei offener Klammer/Zeichenkette/Kommentar — mit Tests je Fall.
- [x] Regel `shape-unlisted` und Config-Dekodierung (`files`, `dialect`,
      `mode`, `allow` literal/regex voll verankert; strikt, Exit-2-Fälle aus
      der Spezifikation), alle Gegenprobe-Fälle des CR als Tests, deterministische
      Ausgabe.
- [x] `--print-config` zeigt den Block im Gerüst; Benutzerhandbuch-Abschnitt
      und Regel-Tabelle — beide zeigen für Zeichenketten-Inhalt nur die sichere Klasse
      `"[^"$\\]*"`, nie `.*`; den CHANGELOG-Eintrag aus slice-209 („nicht implementiert") in
      `[Unreleased]` umschreiben.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` grün, inklusive `make coverage-gate` (Schwelle unverändert).

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| Normalisierer/Zerleger (Ort nach ADR aus slice-209) | neu | reine Textfunktion, keine I/O |
| Regel-Auswertung `shapes` in `internal/hexagon/` | neu | Befund `shape-unlisted` |
| Config-Adapter (`.a-check.yml`-Dekodierung) | update | Block, strikte Dekodierung, Exit-2-Fälle |
| `--print-config`-Gerüst | update | Regelart sichtbar |
| Tests (Normalisierer, Regel, Config, End-to-End-Fixture `build.gradle.kts`) | neu | Happy/Boundary/Negative der neuen Regel-Anforderung; Gegenprobe-Fälle des CR je eigener Testfall |
| `docs/user/benutzerhandbuch.md`, `CHANGELOG.md` | update | öffentlicher Vertrag berührt |

## 4. Trigger

**Start** (`next` → `in-progress`): slice-209 in `done/`, ADR `Accepted`,
WIP-Limit frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): wenn der Normalisierer allein nicht in
  einer Review-Sitzung prüfbar ist — dann Normalisierer/Zerleger als eigener
  Slice vor der Regel.
- `in-progress` → `open` (blockiert): eine Gegenprobe-Anforderung des CR ist
  mit dem Vertrag aus slice-209 nicht erfüllbar (Vertragslücke) — zurück an
  den Planner, Folge-Änderung am Lastenheft.

## 5. Closure-Trigger

DoD vollständig; jeder Gegenprobe-Fall ist ein Test, der ohne die Regel aus
dem richtigen Grund rot liefe (bewusstes Brechen gezeigt); `make gates` Exit 0;
Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Kotlin-Lexik ist reicher als die Heuristik** (`"""…"""` mit `$`-Vorlagen,
  verschachtelte `${ "}" }`, `/* /* */ */`). Ein Lexer-Fehler, der eine
  Zeichenkette zu früh schließt, kann Code als String behandeln. — **Ausgang:**
  *weiter offen* → Beobachtungs-Register:
  [`BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe`](../observations/BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe/observation.md)
  (jetzt 2×). Der Review fand mit der Backtick-Vorlage einen weiteren Fall (behoben); dass die
  Lexik nun vollständig ist, ist nicht belegt.
- **Fehlalarm bei legitimen Dateien** durch die Anweisungsgrenze. Bekannte Quelle aus dem Review von slice-209: ein Zeilenende nach Postfix `!!`/`++`/`--` oder nach `>` verbindet zwei Anweisungen — fail-safe, aber rot; als Testfall aufnehmen. —
  **Ausgang:** *entfallen* — gestrichen mit Begründung: als benannte Grenze ins Benutzerhandbuch
  §4 und als Tests (`!!`, `++`, `--`, `>`) überführt; die Wache über echte Fehlalarme trägt der
  Re-Evaluierungs-Trigger von [ADR-0041](../../adr/0041-shapes-sollform-je-datei.md).
- **Handbuch führt neue Vokabeln** (`shape-*`) neben `construct-leak`;
  verwandt mit
  [`BEO-USER/handbuch-vokabel-der-adapter-rolle`](../observations/BEO-USER/handbuch-vokabel-der-adapter-rolle/observation.md)
  (1×). — **Ausgang:** *entfallen* — gestrichen mit Begründung: es entstand kein zweites
  Vokabular; `shape-unlisted` steht in Regel-Tabelle, Abschnitt und Glossar gleich, und der
  Review prüfte das Handbuch gegen den Code (ohne Vokabel-Befund).

## 7. Closure-Notiz

**Lerneintrag — Form: benannte Spec-Lücke.** [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
kannte die Vorlage mit Backtick-Bezeichner in Zeichenketten nicht (`$` vor `` ` `` ist auch dort eine
Vorlage); ein `"` im Bezeichner schloss die Zeichenkette zu früh, und ein späteres `/*` verschluckte
Code. Benannt und geschlossen in Spezifikation 0.34.0 — als Plan-Änderung vor dem Code, wie §1 es
für eine Vertragslücke verlangt.

**Geliefert:** `shapes` mit Dialekt `kotlin` und `mode: allow-statements`
([AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012),
[ADR-0041](../../adr/0041-shapes-sollform-je-datei.md)): Kern-Modell und Auswertung, Kotlin-Lexer und
Zerleger, Dateisuche mit den Exit-2-Fällen, strikte Dekodierung, Prüfung im validierenden Einstieg
(auch `--print-graph`), `--print-config`-Gerüst, Handbuch 1.42, CHANGELOG. Alle Gegenprobe-Fälle des
CR sind Tests; Coverage 96,4 %.

**Was hat funktioniert:** Die Regressionsfälle aus dem Review von slice-209 standen vor dem Code
als Tests da, und jede Fix-Behauptung bekam eine Mutations-Gegenprobe (Dollar-Präfix, Regex-Verankerung,
Backtick-Vorlage) — rot aus dem richtigen Grund, nicht nur grün.

**Was ging anders als geplant:** Der Review fand eine Vertragslücke (F-2) und damit den vierten
Lexik-Fall in zwei Slices; der Slice nahm die Spec-Präzisierung per Plan-Änderung mit. Die
Glossar-Zahl „zehn Regeln" war ein viertes Auftreten einer schon verkörperten Klasse; behoben durch
einen Zeiger statt einer neuen Zahl. Und eine versehentlich geschriebene Inline-Suppression fiel
beim Durchsehen vor dem ersten Lauf auf, nicht erst im Gate.

**Steering-Loop-Eintrag:** benannte Spec-Lücke, siehe Lerneintrag. Viertes Auftreten von
[`BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher`](../observations/BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher/observation.md)
nach der Verkörperung: **kein Sensor**, begründet im `state.md` des Eintrags — ob eine Zahl in Prosa
eine abschließende Aufzählung behauptet, ist ein Urteil; ein Muster auf „die N …" kennte die
Bezugsmenge nicht.

**Beobachtungs-Register (`../observations/`):**
[`BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe`](../observations/BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe/observation.md)
→ 2× (`evidence/slice-210.md`);
[`BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher`](../observations/BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher/observation.md)
→ 4×, bleibt *verkörpert* mit Begründung ohne Sensor.

**Folge-Slices:** keine neuen; slice-211 liefert `exact` und `unused`.

**Risiken aus §6:** alle drei mit Ausgang — eines *weiter offen* im Register, zwei *entfallen* mit
Begründung.

**Drei Paarungen:** getragen von der Closure von welle-16.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-024`](../../../../harness/conventions.md#mr-024) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-06; kein Release seit `v0.20.0`; der Auflösungs-Trigger des dritten Eintrags ist das nächste Release, das noch aussteht).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `KERN` (Achsen 1, 2, 3 ✓) ·
`ADAPT` (2, 3 ✓) · `USER` (2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-06, gemergter Stand):

- `KERN`:
  [`BEO-KERN/dirvocab-portfor-auseinander`](../observations/BEO-KERN/dirvocab-portfor-auseinander/observation.md)
  (1×) — Vokabular in Config-Adapter und Regel-Engine getrennt geführt; die
  neuen Werte (`dialect`, `mode`) liegen nur im Adapter, die Regel bekommt
  dekodierte Typen. Kein Zuwachs erwartet, beim Review prüfen.
- `ADAPT`: keine Treffer.
- `USER`: siehe §6, dritter Punkt.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
