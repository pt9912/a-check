# slice-209 — Spec-first: Sollform je Datei (`shapes:`)

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-16 — [Welle-Plan](../welle-16-shapes-sollform.md).

**Bezug:** Change Request „Positivliste von Anweisungen je Datei (Sollform)"
(Maintainer, 2026-10-06). Abgrenzung zu
[AC-FA-RULE-011](../../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak);
Schema in
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml);
Grenze nach
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze);
Determinismus nach
[AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus).

**Berührte Spec-Stellen:** `spec/lastenheft.md` §3 (neue Regel-Anforderung),
§5 (Globale Out-of-Scope-Punkte), Schema in [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) ·
`spezifikation.md` §[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema), §[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung), §[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion) ·
`architecture.md` [ARC-003](../../../../spec/architecture.md#2-komponenten) (Extraktions-Adapter liefert die Anweisungen).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Abnahme der zehn Entscheide durch den Maintainer, 2026-10-06).

**Autor:** Claude. **Datum:** 2026-10-06.

**Lerneintrag — Form:** benannte Spec-Lücke.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die Regelart `shapes:` steht als abnahmefähiger Vertrag im
Lastenheft, ist in einer ADR begründet und in der Spezifikation so präzise
gefasst, dass slice-210 und slice-211 ohne eigene Vertragsentscheidung
implementieren können.

**Abnahme-Entscheide, die dieser Slice festschreibt** (aus der Analyse des CR,
vom Maintainer mit der Welle-Anweisung übernommen):

1. **Kein Warn-Level.** Der CR-Hinweis `shape-unused` wird ein
   Opt-in-**Befund** mit Exit 1 (`unused: fail` je Eintrag-Block), sonst gar
   nichts — die Historie 0.25.0 des Lastenhefts hält fest, dass a-check keinen
   Warn-Level kennt.
2. **`ignore: [version-literals]` entfällt.** Der Ausschluss ist nicht
   umrissen; den Zweck deckt ein Regex-Eintrag.
3. **Regex-Einträge sind voll verankert** (implizit `^…$` über die ganze
   normalisierte Anweisung) — sonst erlaubt `dependencies\{.*` beliebigen
   Blockinhalt.
4. **Leerraum-Faltung:** Leerraum außerhalb von Zeichenketten entfällt, außer
   zwischen zwei Bezeichner-Zeichen — dort wird er zu genau einem Leerzeichen
   (`val a` bleibt von `vala` unterscheidbar).
5. **Anweisungsgrenze `kotlin`:** Ein Zeilenende trennt nur auf oberster
   Ebene (keine offene `(`/`[`/`{`) und nur, wenn die nächste
   Nicht-Leer-Zeile nicht mit `{`, `.` oder `?.` beginnt; `;` trennt immer auf
   oberster Ebene. Zu feines Trennen ist fail-safe (strenger), zu grobes auch
   (der Block wird als Ganzes verglichen).
6. **Zeilen-Mapping:** Jeder Befund nennt die Originalzeile, an der die
   Anweisung beginnt (nach Kommentar-Entfernung: das erste Nicht-Kommentar-
   Zeichen).
7. **`mode: exact`** nimmt `expect: <pfad>`; fehlt die Sollform-Datei, ist das
   Exit 2.
8. **`files` ∩ `exclude`:** Eine Datei, die `shapes` nennt und `exclude`
   ausnimmt, ist ein Konfigurations-Widerspruch — Exit 2. Ein Glob, der
   **keine** Datei trifft, ist Exit 2 („jede Datei muss existieren").
9. **Mengen-Semantik in `allow-statements`:** Reihenfolge und Mehrfach-
   Vorkommen einer erlaubten Anweisung sind egal.
10. **Schlüsselname** `shapes` bleibt, Befundklassen `shape-unlisted`,
    `shape-differs`, `shape-unused` wie im CR.

**Schärfungen beim Schreiben — mit der ADR zur Abnahme vorgelegt, nicht abgenommen** (Plan-Änderung
2026-10-06, nach dem unabhängigen Review; die zehn Entscheide oben bleiben im Wortlaut der
Abnahme stehen):

- zu 4: Leerraum bleibt auch zwischen **zwei Operatorzeichen** (`a - -b` ≠ `a--b`).
- zu 5: Die Grenze gilt auch **innerhalb** von `{…}` und wird dort zu `;`; ein Zeilenende nach
  Operator, `,` oder `.` setzt fort.
- zu 3: Ein Regex muss **für sich** kompilieren, bevor er umhüllt wird; die Grenze „`.*` überspannt
  auch Code" ist benannt, die sichere Form (`"[^"$\\]*"`) dokumentiert, nicht erzwungen.
- neu: `match: literal` statt `substring`; Dollar-Präfix-Zeichenketten (`$$"…"`) in der Lexik;
  `files`/`expect` dürfen die Scan-Wurzel nicht verlassen (Exit 2); `literal`-Einträge
  normalisiert die Extraktion im validierenden Einstieg (auch `--print-graph` Exit 2);
  byte-gleiche Befundzeilen einmal; Befund-Form `pfad:zeile: <klasse>: <meldung>`.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Produkt-Code.** *Schicht-Abgrenzung*: der Slice ist Spec-first; Code
  folgt in slice-210 und slice-211, erst nach Abnahme.
- **Der generische Dialekt.** *Ein Folge-Vorgang übernimmt es*: Roadmap,
  *Nächste Wellen*, mit dem Trigger „zweiter Konsument mit Nicht-Kotlin-
  Manifest". Die Anforderung reserviert `dialect` als offene Menge, beschreibt
  aber nur `kotlin`.
- **Die fail-closed Import-Allowlist je Schicht.** *Bestand bleibt bewusst
  stehen*: sie bleibt gated im Out-of-Scope von [AC-FA-RULE-011](../../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak); die ADR
  grenzt sie nur ab.
- **Benutzerhandbuch.** *Ein Folge-Slice übernimmt es*: slice-210 schreibt den
  Handbuch-Abschnitt mit der ersten lauffähigen Fassung — vorher gäbe es nichts
  zu bedienen.

## 2. Definition of Done

- [x] Lastenheft: neue Regel-Anforderung (nächste freie `AC-FA-RULE`-Kennung)
      mit **explizitem Anker**, Beschreibung, Happy/Boundary/Negative aus den
      Gegenprobe-Fällen des CR, Out-of-Scope; §5 präzisiert (Text-Prüfung
      benannter Dateien neben der Import-Ebene); [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) um den Block
      und dessen Exit-2-Fälle; Versions-Bump, Historie-Zeile, CHANGELOG
      `[Unreleased]`.
- [x] ADR (nächste freie Nummer, Index nachgezogen): Regelart neben
      `constructs`, Normalisierungsvertrag (Entscheide 3–6), Abgrenzung zur
      Import-Allowlist, Platz der Normalisierung im Hexagon — `Accepted` erst
      nach Maintainer-Abnahme.
- [x] Spezifikation: Schema-Abschnitt in §[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) (Schlüssel, Werte,
      strikte Dekodierung, Fehlerfälle), Normalisierung und Zerlegung in
      §[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion), Regel-Semantik in §[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
      (Vergleich, Befund-Format
      `pfad:zeile: <klasse>: <meldung>`, Sortierung).
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `spec/lastenheft.md` | update | neue Regel-Anforderung, §5, [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml), Kopf-Version, Historie |
| `docs/plan/adr/<NNNN>-shapes-sollform-je-datei.md` + ADR-Index | neu / update | Entscheidung und verglichene Alternativen (Verbotsliste in `constructs`, Import-Allowlist, `exact` allein) |
| `spec/spezifikation.md` | update | §[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) und §[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung), §[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion) |
| `spec/architecture.md` | update | [ARC-003](../../../../spec/architecture.md#2-komponenten): die Extraktion liefert zusätzlich die normalisierten Anweisungen der `shapes`-Dateien — **Plan-Änderung vor dem Schreiben** (2026-10-06): die Zerlegung ist Extraktion, nicht Regel-Auswertung |
| `CHANGELOG.md` | update | `[Unreleased]`, Vertrag berührt (`AGENTS.md` §6 Schritt 7) |

## 4. Trigger

**Start** (`next` → `in-progress`): WIP-Limit frei, Welle welle-16 eröffnet.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): wenn die Spezifikation des
  Kotlin-Normalisierers (Roh-Strings, verschachtelte `${…}`, verschachtelte
  Block-Kommentare) einen eigenen Abschnitt verlangt, der in einer
  Review-Sitzung nicht prüfbar ist — dann Normalisierungsvertrag als eigener
  Slice.
- `in-progress` → `open` (blockiert): der Maintainer nimmt einen der zehn
  Abnahme-Entscheide in §1 nicht ab und will erst Adopter-Rückmeldung.

## 5. Closure-Trigger

DoD vollständig; ADR `Accepted`; `make gates` und `make verify` Exit 0;
Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Statement-Grenze in Kotlin ist heuristisch.** Eine Gradle-Schreibweise,
  die Entscheid 5 falsch zerlegt, wird rot statt grün (fail-safe) — aber eine
  Regel, die bei legitimen Dateien rot ist, wird abgeschaltet. — **Ausgang:**
  *entfallen* — gestrichen mit Begründung: als Risiko des **Vertrags** ist es mit der Fassung
  entfallen, die beide Fehlrichtungen der Zerlegung als fail-safe festschreibt
  ([SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
  Schritt 3); die Wache über echte Fehlalarme trägt der Re-Evaluierungs-Trigger von
  [ADR-0041](../../adr/0041-shapes-sollform-je-datei.md), den
  Testfall (Postfix `!!`/`++`/`--`, `>`) §6 von slice-210.
- **Nur ein Adopter belegt.** Der Vertrag könnte an dessen Datei überangepasst
  sein. — **Ausgang:** *weiter offen* → Beobachtungs-Register:
  [`BEO-SPEC/vertrag-auf-einen-konsumenten-belegt`](../observations/BEO-SPEC/vertrag-auf-einen-konsumenten-belegt/observation.md)
  (neu, 1×).
- **§5-Erweiterung verschiebt den Produkt-Umfang.** „Heuristik auf
  Import-Ebene" ist seit `constructs` schon nicht mehr wörtlich wahr; die
  Präzisierung muss das benennen, nicht kaschieren. — **Ausgang:** *entfallen* —
  gestrichen mit Begründung: §5 des Lastenhefts nennt jetzt ausdrücklich drei Prüf-Ebenen
  (Importe, Roh-Text, normalisierte Anweisungsfolge) und dass keine Semantik auswertet; die
  Erweiterung ist benannt, nicht kaschiert.

## 7. Closure-Notiz

**Lerneintrag — Form: benannte Spec-Lücke.** Ein Regex-Eintrag sieht die normalisierte Anweisung
als Text, nicht als Token-Folge: ein `.*` für Zeichenketten-Inhalt überspannt auch Code. Die Lücke
ist benannt in [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema),
im Out-of-Scope von [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012) und als
Konsequenz in [ADR-0041](../../adr/0041-shapes-sollform-je-datei.md); dokumentiert ist die sichere
Klasse `"[^"$\\]*"`, erzwungen wird sie nicht. Ein Platzhalter für „genau eine Zeichenkette" würde
sie schließen und ist als eigener Umfang benannt, nicht geplant.

**Geliefert:** neue Anforderung [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012)
(Lastenheft 0.28.0, mit §5 und Glossar), Schema in
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
Spezifikation 0.33.0, Architektur 0.5.0, [ADR-0041](../../adr/0041-shapes-sollform-je-datei.md)
`Accepted` nach Maintainer-Abnahme einschließlich der vier Schärfungen aus §1.

**Was hat funktioniert:** Spec-first mit unabhängigem Review vor jeder Zeile Code. Vier Runden
fanden zusammen 17 Befunde, darunter drei, mit denen sich eine Abhängigkeit als Kommentar hätte
verstecken lassen — im Vertrag gefunden kosten sie einen Satz, im Lexer einen Rückbau. Die
Schärfungen gegenüber den abgenommenen Entscheiden standen dabei offen als Plan-Änderung in §1,
statt still in der ADR.

**Was ging anders als geplant:** Der Lexik-Vertrag brauchte drei Korrektur-Runden, die letzte
wegen eines Beispiels, das ich aus dem Kopf statt Byte für Byte gezählt hatte. Und
`make doc-structure` lehnte die erste neue Anforderung seit Einführung der AC-Form-Regel ab, weil
die Regel `Boundary:` wörtlich verlangt, der grandfathered Bestand aber `Boundary (…):` schreibt.
Die Plan-Änderung (§[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
und [ARC-003](../../../../spec/architecture.md#2-komponenten) zusätzlich) lag als eigener Commit vor dem Schreiben.

**Steering-Loop-Eintrag:** benannte Spec-Lücke, siehe Lerneintrag — gezählt, nicht verkörpert
(kein Register-Eintrag hat mit diesem Slice 3× erreicht).

**Beobachtungs-Register (`../observations/`):** drei neu angelegt, je 1× —
[`BEO-GATE/marke-enger-als-bestandsform`](../observations/BEO-GATE/marke-enger-als-bestandsform/observation.md),
[`BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe`](../observations/BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe/observation.md)
(drei Funde, ein Vorgang),
[`BEO-SPEC/vertrag-auf-einen-konsumenten-belegt`](../observations/BEO-SPEC/vertrag-auf-einen-konsumenten-belegt/observation.md)
(Risiko 2).

**Folge-Slices:** keine neuen; slice-210 und slice-211 führen den Vertrag aus und tragen die
Folgepflichten aus dem Review (Tests je Lexik-Form, sichere Regex-Klasse in Handbuch und Gerüst,
CHANGELOG umschreiben).

**Risiken aus §6:** alle drei mit Ausgang — zwei *entfallen* mit Begründung, eines *weiter offen*
im Register.

**Drei Paarungen:** getragen von der Closure von welle-16 (offene Welle, `docs/plan/planning/README.md`
§Beim Kopieren, Punkt 3).

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-024`](../../../../harness/conventions.md#mr-024) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-10-06, Konventions-Stand dieses Laufs; kein Release seit `v0.20.0`, kein Golden Set).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `SPEC` (Achsen 1, 2, 3 ✓) ·
`ADR` (1, 2, 3 ✓) — beide in der Modus-Deklaration von
[`harness/conventions.md`](../../../../harness/conventions.md) geführt.

**Vorgelagert — offene Beobachtungen sichten** (2026-10-06, gemergter Stand):

- `SPEC`:
  [`BEO-SPEC/kennung-ohne-expliziten-anker`](../observations/BEO-SPEC/kennung-ohne-expliziten-anker/observation.md)
  ist **verkörpert** — daher der explizite Anker als DoD-Bestandteil.
- `ADR`: keine Treffer.
- Nachbarschaft, nicht berührt, aber inhaltlich verwandt:
  [`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md)
  (2×) beschreibt für die eigenen Sensoren genau das, was der CR für
  `constructs` beim Adopter beobachtet: eine Verbotsliste trifft nur die
  häufige Schreibweise. Kein Zähler-Zuwachs (andere Sub-Area, kein eigener
  Vorgang), aber ein Argument für die ADR.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
