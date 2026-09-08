# Reviewer-Skill — a-check

- **Status:** Accepted
- **Gilt für:** Plan-/Design-/Code-Review der Doku- und (ab slice-003)
  Code-Artefakte dieses Repos.
- **Bezug:** [`AGENTS.md`](../../AGENTS.md) §3 (Hard Rules) + §5 (Traceability);
  Regelwerk **v6.6.0** Modul 10 (vendored: [`.harness/baseline/v6.6.0/regelwerk/modul-10-review-harness.md`](../baseline/v6.6.0/regelwerk/modul-10-review-harness.md)). Baseline: [`harness/conventions.md`](../../harness/conventions.md) §Baseline.

Repo-spezifisches „worauf achtest du", damit ein Reviewer-Agent zwischen
Sessions nicht driftet (Regelwerk Modul 10). Diese Datei wird versioniert,
nicht überschrieben (ADR-Hard-Rule, Modul 4).

## Kontext-Eingang (Pflicht)

Bevor der Reviewer den Gegenstand liest:

- der Review-Gegenstand (Diff, ADR-Entwurf, Slice-Plan)
- [`spec/lastenheft.md`](../../spec/lastenheft.md) für referenzierte `AC-*`-IDs
- ADRs aus [`docs/plan/adr/`](../../docs/plan/adr/), deren ID im Gegenstand
  vorkommt — nur aktive (`Proposed`/`Accepted`), nie `Superseded`
- [`AGENTS.md`](../../AGENTS.md) §3 Hard Rules
- frühere Findings am selben Bereich ([`docs/reviews/`](../../docs/reviews/))

## Klassifikation (für dieses Repo)

**HIGH** — blockiert:
- Verstoß gegen eine Hard Rule ([`AGENTS.md`](../../AGENTS.md) §3.1–§3.6)
- Harness-Lüge: behauptetes Gate ohne Make-Target, erfundene ID, stille Setzung
- Spec-Stratum referenziert abwärts (ADR/Slice) — Referenz-Richtung verletzt
- nachweislich falsche Tatsachenbehauptung (gegen ein Repo-Artefakt verifiziert)
- **Kommentar trägt keine der Kommentar-Klassen** — er beschreibt die verworfene
  Alternative („ohne X wäre …"), einen abwesenden Text („früher stand hier …")
  oder bricht mitten im Satz ab, weil eine Teilersetzung den Rest stehen ließ.
  Gilt für Code, Konfiguration, Skripte
  ([`AGENTS.md`](../../AGENTS.md) §3.7). **Kein Gate fängt das** — die Regel ist
  inferentiell und hängt am Review.
  *Im Bestand belegt:* [`BEO-HARNESS/hard-rule-37-ohne-sensor`](../../docs/plan/planning/observations/BEO-HARNESS/hard-rule-37-ohne-sensor/observation.md)
  (slice-108, slice-169).
- **Zustandsfeld trägt Chronik** — eine `Stand`-/`Status`-Zelle (Roadmap,
  Beobachtungs-Register, Meilenstein) erzählt, *wie* der Zustand entstand,
  statt Zustand und Beleg als auflösbaren Anker zu nennen; oder ein Drift-Log
  protokolliert Schließungen und erreichte Meilensteine. **Kein Gate fängt
  das.** Zwei Ausprägungen, beide teuer geworden: ein `state.md`, das nach der
  Behebung den behobenen Zustand weiter behauptet, und eines, dessen
  `Stand:`-Zeile bei 3× keinen der drei Ausgänge trägt.
  *Im Bestand belegt:* [`BEO-HARNESS/chronik-in-gelesenen-dateien`](../../docs/plan/planning/observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
  (slice-103, slice-182).
- **Norm nur im Template-Kommentar** — eine Regel steht im `<!-- -->`-Block
  einer vendorten Ziel-Form und nirgends sonst. Sie ist weg, sobald jemand die
  Vorlage kopiert und die Kommentare löscht — und **a-check kopiert sie**
  ([`docs/plan/planning/README.md`](../../docs/plan/planning/README.md) §Beim
  Kopieren der Slice-Ziel-Form; eigene Vorlagen-Dateien führt das Repo nicht).
  *Besetzter Fall:* `v6.6.0` · `templates/docs/reviews/review-report.template.md`
  trägt im vierten Kommentarblock die Norm *„die Klassen-Bezeichnung muss über
  Läufe hinweg stabil sein"* — sie steuert den Steering-Loop-Zähler und steht
  **nur dort**.
  **Zweite Ausprägung, hiesig:** [`.d-check.yml`](../../.d-check.yml) trägt die
  Begründung jeder Regel ausschließlich in Kommentaren. Das ist zulässig — aber
  eine *Zusage* darf dort nicht allein stehen; sie gehört in
  [`AGENTS.md`](../../AGENTS.md) §4 oder eine Sensor-Datei, die der Lauf liest.
  **Kein Gate fängt beides.**

**MEDIUM** — vor Merge/Acceptance klären:
- unbelegte Tatsachenbehauptung (nicht gegen ein Repo-Artefakt belegbar)
- `Bezug:`/`Schärft:` unvollständig oder unpassend zur Entscheidung
- fehlende wesentliche Konsequenz/Risiko in einer ADR; fehlender oder vager
  Fitness-Function-Anker

**LOW** — nice-to-fix: Wording, Bezug-Feinheiten, fehlender Querverweis,
Provenance-Platzierung außerhalb der Historie-Zone.

**INFO** — Hinweis ohne erwartete Aktion (Verweis auf zuständige Rolle oder
Folge-Slice).

## Was dieser Skill NICHT macht

- Keine Lösungsvorschläge — Reviewer kategorisiert, Implementer entscheidet.
- Keine Verifikation gegen DoD (Verifier, Modul 11), keine Validation
  (Validator).
- Kein Schreibzugriff auf den Review-Gegenstand.

## Output-Schema

Pro Finding: `kategorie` · `quelle` (AC-/ADR-ID, Hard-Rule, Konvention) ·
`pfad` (Datei:Zeile) · `befund` (1–2 Sätze, beobachtbar, **ohne
Lösungsvorschlag**) · `verifizierbar` (ja/nein — gäbe es einen Gate-/Tool-Lauf,
der es bestätigt?) · `klasse` (die Finding-Klasse in einem Halbsatz).

**`klasse` speist den Steering-Loop-Zähler**, und darum gilt für sie die Norm
aus der Ziel-Form des Reports: **die Bezeichnung muss über Läufe hinweg stabil
sein.** Leiten zwei Läufe dieselbe Klasse unterschiedlich ab, zählt das
Register sie getrennt, und keine erreicht je 3×. Alle bestehenden Reports
führen das Feld; im Skill fehlte es bis slice-185.

**Zitier-Form** (Norm, `v6.6.0` · `templates/docs/reviews/review-report.template.md`):
Der Report friert ein; was er zitiert, bewegt sich weiter. Deshalb **Kennung,
nicht Adresse** — `slice-NNN` statt seines Lifecycle-Pfads, `make <target>`
statt eines Links auf die Sensor-Datei, eine Baseline-Stelle als **Tag + Pfad
in Inline-Code** statt als Link: `` `v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt> ``.
Der vendored Baum trägt genau einen Tag; der nächste Sprung löscht den alten,
und ein Link darauf färbt einen Report rot, den niemand mehr anfassen darf.
Das `pfad`-Feld auf den **geprüften Gegenstand** ist davon nicht betroffen — es
hält den Stand des Laufs fest und darf das.

Zusätzlich: pro geprüftem Bereich eine **Negativbefund-Zeile** („geprüft, ohne
Befund"), eine **Kategorie-Summary** und ein **Verdikt**. **HIGH-Findings
werden vor Übernahme adversarisch gegen das Repo-Artefakt verifiziert**
(Modul 11). Kontext-Trennung: wer ein Artefakt verfasst hat, reviewt es nicht
im selben Kontextfenster (Modul 8). Report-Ablage: ein Report pro Lauf unter
[`docs/reviews/`](../../docs/reviews/), Folgeläufe als neue Datei.

## Mess-Regeln (umgezogen aus `AGENTS.md` §5, slice-186)

Zwei Urteilsregeln über **Belege**: hier steht ihre **Herleitung** und der
gemessene Bestand, an dem sie hängen — die **Zusage** selbst steht in
[`AGENTS.md`](../../AGENTS.md) §5, weil sie den *Schreibenden* bindet und der
diese Datei nicht liest ([`harness/README.md`](../../harness/README.md) §Guides:
*„nicht Teil der Implementer-Eingabe"*). Zwei Adressaten, zwei Orte: Wer eine
Messung schreibt, braucht den Satz; wer sie prüft, die Fälle darunter. Beide
sagen selbst „kein Sensor" — sie hängen am Review, und `modul-08` §Welche Rolle
braucht welche Artefaktklasse weist die **Urteilsgrundlage** der Skill-Datei zu.

- **Geltungsbereich einer Messung** (`seit slice-179`, Lese-Schritt der
  welle-15-Closure): Wer eine Messung als **Beleg** schreibt — in einem
  Slice-Plan, einer Closure-Notiz, einem Review-Report —, nennt ihren
  **Geltungsbereich** und sagt, ob er den Gegenstand deckt. Nicht *„22 Befunde,
  keine weitere Klasse"*, sondern *„22 Befunde über Markdown-Links; Prosa sieht
  das Instrument nicht"*.
  **Anlass:** [`BEO-PLAN/review-geltungsbereich-zu-eng`](../../docs/plan/planning/observations/BEO-PLAN/review-geltungsbereich-zu-eng/observation.md)
  bei 3× — dreimal war die Begrenzung begründet und trotzdem zu eng, und
  dreimal fand es jemand anderes als der Messende. **Kein Sensor:** ob ein
  Geltungsbereich weit genug ist, ist ein Urteil über eine Absicht
  ([`AGENTS.md`](../../AGENTS.md) §3.7); ein
  zweites Muster, das nach übersehenen Klassen sucht, kann dieselbe Verengung
  haben wie das erste. Was greift, ist die Frage beim **Schreiben** — sie kostet
  einen Halbsatz und hätte alle drei Fälle gefangen.
  **Zweite Hälfte** (`seit slice-182`): Ein Größen-*Vergleich* mit einer Ziel-Form
  ist noch kein Befund. Eine Vorlage ist kürzer als jedes ausgefüllte Dokument,
  und ihre Platzhalter und Bedienhinweise — die beim Kopieren verschwinden —
  zählen in ihr mit. Befund ist nachgeschriebener **Baseline-Normtext**; das
  entscheidet die Herkunft eines Satzes, nicht seine Länge.
  **Und die eigene Spec ist nicht die Baseline:** Ein Einstiegspunkt, der eine
  `AC-*`-Zusage zusammenfasst und auf sie verlinkt, zeigt nach **oben** in der
  Source Precedence — das ist seine Aufgabe, nicht sein Fehler. Dieselbe Zusage
  steht dann mehrfach im Repo, und das ist richtig so.
  Gemessen an [`harness/README.md`](../../harness/README.md): Der Abschnitt mit dem
  größten Faktor (**22 ×** gegen die Ziel-Form, wo zwei Platzhalter-Punkte
  stehen) blieb unverändert — er ist die ausgefüllte Ziel-Form und fasst zwei
  `AC-QA-*` zusammen; gekürzt wurden zwei Abschnitte mit kleinerem Faktor, in
  denen das **Regelwerk** nachgeschrieben war — einer davon zweimal in derselben
  Datei.
  **Dritte Hälfte** (`seit slice-183`, an [`AGENTS.md`](../../AGENTS.md) §5
  gemessen): *„Steht die
  Regel im Regelwerk?"* ist ebenfalls die falsche Frage. In einem Repo, das eine
  Baseline **adoptiert** hat, lautet die Antwort fast immer ja — sechs von sechs
  Substanz-Stichproben fanden eine Fundstelle. Trägt nur: *„schreibt dieser
  Absatz ihre **Begründung** nach?"* Eine Regel zu **nennen** und die repo-eigene
  Ausprägung danebenzustellen ist die Aufgabe eines Briefings; erst die
  nachgeschriebene Herleitung ist der Befund. Gemessen traf das auf **drei von
  18** Blöcken in [`AGENTS.md`](../../AGENTS.md) §5 zu — bei einem
  Größen-Faktor von **23,2 ×** gegen die
  Ziel-Form.
- **Wer eine Menge zählt, zählt sie zweimal verschieden** (`seit slice-193`,
  Register-Eintrag bei 3×): Eine Zählung, die einen **Befund** oder einen
  **Umfang** trägt, wird mit einem zweiten, **anders gebauten** Zähler
  wiederholt. Weichen beide ab, ist der Unterschied der Befund.
  **Drei Ausprägungen, alle belegt:** Die Sammelaussage deckt ihre Menge nicht
  (slice-185: *„alle sechs sind Bedienhinweise"* — einer trug eine Norm) · die
  Zahl stammt aus der falschen Quelle (slice-186: *„alle fünf Abschnitte"* — die
  Ziel-Form führt sechs) · **der Zähler liest je Zeile nur den ersten Treffer**
  (slice-193: eine Tabellenzelle mit vier Targets zählte als eines; zwei Targets
  wären ohne Index-Zeile geblieben und hätten den Lauf rot gemacht).
  **Die Prüf-Frage ist nicht „hast du gezählt?", sondern „was zählt dein Zähler
  als eins?"** — die dritte Ausprägung ist die teuerste, weil ihre Zahl
  plausibel aussieht.
  **Kein Sensor:** Ob eine Zählung ihre Menge trifft, ist ein Urteil über ihre
  Konstruktion ([`AGENTS.md`](../../AGENTS.md) §3.7). Was greift, ist der zweite
  Zähler — er kostet einen Aufruf und hätte alle drei Fälle gefangen.
  Auslöser: [`BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat`](../../docs/plan/planning/observations/BEO-PLAN/kandidaten-klassifikation-groeber-als-der-kandidat/observation.md)
  (slice-185, slice-186, slice-193 — 3×).
- **Eine Mutations-Probe belegt erst, wenn sie rot war** (`seit slice-181`,
  Register-Eintrag bei 3×): Wer einen Prüfer mit einer Probe belegt, zeigt
  **beide** Richtungen — und die **rote** ist die, die zählt. Grün beweist
  nichts: Ein Prüfer, der seinen Gegenstand gar nicht erreicht, ist grün, und
  eine Probe, die ihn verfehlt, ebenso.
  **Zwei Ausprägungen, beide belegt:** Die Probe **liefert den Gegenstand mit**
  — sie baut ihren Fall so, dass eine bereits gedeckte Eigenschaft ihn in die
  Prüfmenge bringt, und die neue wird nie erreicht (slice-180: eine Datei mit
  *allen drei* Verweis-Formen). Oder sie **trifft daneben** — die Mutation
  landet außerhalb des Gegenstands (slice-181: Fülltext hinter dem schließenden
  `|` einer Tabellenzelle; slice-169: das Muster mutiert statt der
  Kandidatenmenge).
  **Die Prüf-Frage ist nicht „hast du eine Probe?", sondern „war sie rot, und
  woran?"** — die Meldung nennen, nicht nur den Exit-Code.
  **Kein Sensor:** Ob eine Probe ihren Gegenstand trifft, ist ein Urteil über
  ihren Aufbau ([`AGENTS.md`](../../AGENTS.md) §3.7). Was greift, ist die Frage
  beim Schreiben — sie kostet
  einen Handgriff und hätte alle drei Fälle gefangen.
  Auslöser: [`BEO-GATE/probe-liefert-den-gegenstand-mit`](../../docs/plan/planning/observations/BEO-GATE/probe-liefert-den-gegenstand-mit/observation.md)
  (slice-169, slice-180, slice-181 — 3×).

## Pflege (Steering-Loop)

Bei dreimaligem gleichem Finding: Klassifikation schärfen → Folge-ADR oder
[`AGENTS.md`](../../AGENTS.md)-Eintrag → Fitness Function (Modul 13).
