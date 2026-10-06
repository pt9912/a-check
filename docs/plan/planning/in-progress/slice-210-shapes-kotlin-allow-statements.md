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

**Lerneintrag — Form:** wird bei Closure benannt.

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

- [ ] Normalisierer und Anweisungs-Zerleger `kotlin` (Kommentare inkl.
      verschachtelter Block-Kommentare, Zeichenketten inkl. Roh-Strings,
      Dollar-Präfix und `${…}`-Vorlagen samt maskiertem `\$`, Leerraum-Faltung, Anweisungsgrenze, Zeilen-Mapping),
      Exit 2 bei offener Klammer/Zeichenkette/Kommentar — mit Tests je Fall.
- [ ] Regel `shape-unlisted` und Config-Dekodierung (`files`, `dialect`,
      `mode`, `allow` literal/regex voll verankert; strikt, Exit-2-Fälle aus
      der Spezifikation), alle Gegenprobe-Fälle des CR als Tests, deterministische
      Ausgabe.
- [ ] `--print-config` zeigt den Block im Gerüst; Benutzerhandbuch-Abschnitt
      und Regel-Tabelle — beide zeigen für Zeichenketten-Inhalt nur die sichere Klasse
      `"[^"$\\]*"`, nie `.*`; den CHANGELOG-Eintrag aus slice-209 („nicht implementiert") in
      `[Unreleased]` umschreiben.
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

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
  *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*
- **Fehlalarm bei legitimen Dateien** durch die Anweisungsgrenze. Bekannte Quelle aus dem Review von slice-209: ein Zeilenende nach Postfix `!!`/`++`/`--` oder nach `>` verbindet zwei Anweisungen — fail-safe, aber rot; als Testfall aufnehmen. —
  **Ausgang:** *(bei Closure zuzuweisen)*
- **Handbuch führt neue Vokabeln** (`shape-*`) neben `construct-leak`;
  verwandt mit
  [`BEO-USER/handbuch-vokabel-der-adapter-rolle`](../observations/BEO-USER/handbuch-vokabel-der-adapter-rolle/observation.md)
  (1×). — **Ausgang:** *(bei Closure zuzuweisen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

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
