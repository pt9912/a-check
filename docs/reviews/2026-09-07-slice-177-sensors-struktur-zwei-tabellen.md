# Review — slice-177: Sensors-Struktur, zwei Tabellen und `harness/sensors/`

**Review-Art:** unabhängiger Lauf (eigenes Kontextfenster, kein `fork`; die Umsetzung
stammt aus einem anderen Kontext — `AGENTS.md` §6, `v6.5.0` ·
`regelwerk/modul-08-agentenrollen.md` §Rollen-Regeln)
**Skill:** `.harness/skills/reviewer.md` @ `3fae6d3` · **Modell:** `claude-opus-5[1m]`
**Datum:** 2026-09-07
**Gegenstand:** Commit-Range `HEAD~1..HEAD` (`9e42685`) — ein Commit, der den
Lifecycle-Wechsel **und** die Umsetzung trägt (siehe F-1)

> **Zitier-Form** *(Norm, bleibt stehen).* Dieser Report friert ein; was er zitiert,
> bewegt sich weiter. Deshalb **Kennung, nicht Adresse** — `slice-NNN` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als `v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt>. Das `pfad`-Feld
> auf den **geprüften Gegenstand** ist davon nicht betroffen — es hält den Stand des
> Laufs fest und darf das.

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- Slice-Plan slice-177 (Stand `9e42685`), Diff der Range `HEAD~1..HEAD`
- `AGENTS.md` §3 (Hard Rules, insbesondere §3.3 und §3.7), §4 (Gate-Tabelle),
  §5 (Dokumentations-Regeln), §6 (Workflow)
- `harness/README.md` §Sensors (beide Tabellen), `harness/conventions.md` §Baseline
- `v6.5.0` · `templates/harness/sensors/gate.template.md` (Ziel-Form je Sensor-Datei)
- `v6.5.0` · `templates/harness/README.template.md` §Sensors,
  `v6.5.0` · `templates/AGENTS.template.md` §4
- `v6.5.0` · `regelwerk/modul-13-quality-gates.md` §Vorhanden ≠ behauptet
- `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Lifecycle als State Machine,
  §Offene Risiken werden bei Closure aufgelöst, §Zwei Schritte vor der Modus-Begründung
- `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
- `.harness/skills/reviewer.md` (Klassifikation, Output-Schema, Zitier-Form)
- Quellen der geprüften Tatsachenbehauptungen: `tools/image-scan.sh`,
  `tools/symlink-check.sh`, `tools/harness/record-gates.sh`, `tools/archive-wave/`,
  `.d-check.yml`, `d-check.mk`, `Makefile`, `.github/workflows/image-scan.yml`
- frühere Findings am selben Bereich: Reports zu slice-174, slice-175, slice-176
- `AC-*`: keine berührt (der Slice-Kopf führt `— (keine)`; geprüft und zutreffend —
  die Änderung fasst kein Spec-Stratum an)

---

## HIGH

### F-1 — Lifecycle-`git mv` und Inhalts-Umschreibung liegen in **einem** Commit; die Rename-Erkennung ist dabei genau so gefallen, wie §3.3 es voraussagt

- **kategorie:** HIGH
- **quelle:** `AGENTS.md` §3.3 (*„git mv + Inhaltsänderung = zwei Commits … Sonst fällt
  die Rename-Detection unter die 50 %-Schwelle und `git log --follow` wird
  unzuverlässig"*) · `AGENTS.md` §5 (Slice-Lifecycle ist reine Datei-Bewegung) ·
  `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Lifecycle als State Machine
  (*„Ein Übergang ist deshalb ein **reiner `git mv`**"*)
- **pfad:** Commit `9e42685` (`docs/plan/planning/open/slice-177-…md` →
  `docs/plan/planning/in-progress/slice-177-…md`, dazu 234 neue und 121 entfernte Zeilen)
- **befund:** Der Commit enthält gleichzeitig den Lifecycle-Wechsel `open/` →
  `in-progress/`, die von `make slice-mv` erzeugten Verweis-Umschreibungen in fünf
  Fremddateien, den Ruhe-Marker-Wechsel der Roadmap und die vollständige Neufassung des
  Slice-Plans. Der Slice-Kopf derselben Datei sagt in seiner Lifecycle-Zeile, der
  Zustand *„wechselt nur durch `make slice-mv`"*; die Beauftragung dieses Reviews nannte
  den `git mv` als „davor liegend" — beides trifft auf den Range nicht zu.
- **Gegen-Messung (adversarisch, HIGH-Pflicht):** `git show 9e42685 --name-status`
  meldet `D` + `A`, **keinen** Rename — die Standard-Schwelle von 50 % ist nicht
  erreicht. Erst `--find-renames=10%` zeigt `R040`, also 40 % Ähnlichkeit.
  Folge, gemessen: `git log --follow` auf den neuen Pfad liefert **genau einen**
  Commit; `git log` auf den alten Pfad liefert **sechs** (`dc0e58b`, `179c44b`,
  `fe5a0df`, `199a12c`, `f49913a`, `9e42685`). Fünf Commits Vorgeschichte — darunter die
  Ausgangsmessung aus §2 (`fe5a0df`) und die Verankerung des Beobachtungs-Belegs
  (`199a12c`) — sind über den Lifecycle-Pfad nicht mehr auffindbar. Der von §3.3
  benannte Schaden ist nicht potenziell, sondern eingetreten.
- **verifizierbar:** ja — `git show <commit> --name-status` und
  `git log --follow <neuer Pfad>`; kein Gate misst es (weder `doc-check` noch
  `commit-scope-check` prüfen Rename-Ähnlichkeit).

### F-2 — Der Schnitt „vier bekommen eine Datei, zwölf nicht" hält seiner eigenen Begründung nicht stand: mindestens acht der zwölf tragen eine ausdrücklich benannte Deckungsgrenze

- **kategorie:** HIGH
- **quelle:** `v6.5.0` · `templates/harness/sensors/gate.template.md` §Template-Hinweis
  (*„entsteht **nur, wenn ein Target mehr braucht als einen Satz** — Deckungsgrenze,
  Ausgabe-Bedeutung, Exit-Codes, Abbruch-Bedingungen"*; **Deckungsgrenze zuerst
  genannt**) · slice-177 §3.1 und §1 · Reviewer-Skill §Klassifikation, *nachweislich
  falsche Tatsachenbehauptung*
- **pfad:** `docs/plan/planning/done/slice-177-sensors-struktur-zwei-tabellen.md`:94–101
  (Tabelle §3.1) und :29–33 (§1, erster „Nicht in diesem Slice"-Punkt)
- **befund:** §3.1 teilt die 16 überlangen Zellen in zwei Gruppen und behauptet für die
  zweite: *„**Vertrag plus Historie** … *woher* die Regel kam, *seit wann* sie im
  Aggregat hängt, *welcher Vorfall* sie auslöste"*, Antwort deshalb *„**kürzen**, nicht
  auslagern"*. Die Commit-Message wiederholt es als Ergebnis einer Sortierung
  (*„Danach sortiert zerfallen die 16 in zwei Gruppen"*). Der Bestand widerspricht: Die
  zwölf tragen ganz überwiegend **genau das erste Kriterium der Ziel-Form** — eine
  benannte Deckungsgrenze, also eine Aussage darüber, was das Grün *nicht* abdeckt.
  Damit ist nicht gesagt, dass zwölf Dateien hätten entstehen müssen; gesagt ist, dass
  die **Begründung** für ihr Ausbleiben gegen den eigenen Bestand nicht trägt. Die
  Auswahl begründet sich selbst, statt gemessen zu sein — und der Slice führt sie in §8
  zugleich als Lerneintrag („geschärfte Regel") und in §7 als Beleg dafür, dass das
  Risiko „nur zum Teil" eingetreten sei.
- **Gegen-Messung (adversarisch, HIGH-Pflicht):** Zellenlängen selbst nachgemessen
  (Feld 3 je Tabellenzeile in `AGENTS.md` §4): **16** Zellen über 250 Zeichen —
  die Zahl aus §2 stimmt. Von den zwölf nicht behandelten enthalten **acht** eine
  wörtlich ausgewiesene Grenze:
  `doc-workflows` (*„Geprüft wird die **Form**, nicht die **Gültigkeit** … der
  Widerspruch aus BEO-026 bleibt damit ungedeckt"*) · `verify-observations` (*„**Nicht**
  geprüft: Lage und Existenz der Beleg-Datei … und die Umkehrung"*) ·
  `dcheck-phrase-selftest` (*„**Nicht** geprüft: ob eine künftig geänderte Formulierung
  ebenfalls greifen würde"*) · `version-coherence` (*„Geprüft wird **Divergenz, nicht
  Unwahrheit** — zwei übereinstimmend falsche Angaben bleiben grün"*) ·
  `ci-range-selftest` (*„**Nicht** geprüft: welchen Wert GitHub in das Feld schreibt"*) ·
  `verify-risiko-ausgaenge` (*„**Nicht** geprüft: die Existenz eines Risiko-Blocks"*) ·
  `regelwerk-check` (*„die Freshness-Hälfte bleibt … ausdrücklich ungeprüft"*) ·
  `doc-reviews` (*„nicht rekursiv"*, *„Opt-in pro Slice über die DoD-Phrase selbst —
  ohne sie ist die Kandidatenmenge für diesen Slice leer"*). Zwei weitere tragen eine
  Vorbedingung bzw. eine bewusst nicht konfigurierte Fähigkeit: `doc-structure`
  (*„Braucht den Pin `v0.69.0`"*) und `doc-planning` (*„Die zweite und dritte
  Modul-Fähigkeit … bleiben bewusst unkonfiguriert"*). Für die Beschreibung
  „Vertrag plus Historie" bleiben nach dieser Messung **zwei** übrig: `doc-targets` und
  `gate-consistency`.
- **verifizierbar:** teilweise — die Zellenlängen sind reproduzierbar auszählbar, die
  Zuordnung „ist das eine Deckungsgrenze?" ist Urteil (`AGENTS.md` §3.7: kein Match).

---

## MEDIUM

### F-3 — In `AGENTS.md` §4 sind 260 Zeichen der alten `archive-wave`-Zelle stehengeblieben; die gemessene Wirkung „728 → 256" beschreibt die Zeile, nicht die Datei

- **kategorie:** MEDIUM
- **quelle:** slice-177 §3.2 (Tabelle *Gemessene Wirkung*) · `AGENTS.md` §3.7
  (der Kommentar nennt den Zustand, nicht die Chronik)
- **pfad:** `AGENTS.md`:197–200
- **befund:** Die alte `archive-wave`-Zelle war im Quelltext auf vier physische Zeilen
  umgebrochen. Ersetzt wurde nur die erste; die drei Fortsetzungszeilen stehen
  unverändert unter der neuen, jetzt vollständigen Tabellenzeile und sagen ein zweites
  Mal, was die neue Zelle schon sagt (*„aber seit slice-157 **Pflichtschritt beim
  Abschluss eines wellenlosen Slice** (§6)"* neben *„Seit slice-157 Pflichtschritt beim
  Abschluss eines wellenlosen Slice (§6)"*). Sie enden auf ein alleinstehendes `|` und
  gehören zu keiner Zeile mehr. §3.2 des Slice-Plans und die Commit-Message führen für
  diese Zelle „728 → 256"; im Dateizustand stehen 256 Zeichen Zeile **plus** 260 Zeichen
  Rest.
- **verifizierbar:** ja — `sed -n '197,200p' AGENTS.md`; kein Gate misst es
  (`doc-targets` prüft Target-Existenz, nicht Tabellen-Wohlgeformtheit).

### F-4 — `harness/sensors/image-scan.md` nennt eine von zwei tatsächlich gescannten Referenzen

- **kategorie:** MEDIUM
- **quelle:** `v6.5.0` · `templates/harness/sensors/gate.template.md` §Vertrag/§Grenze
  (*„Der Prüfbereich, und wo er enger ist als der Bereich, über den das Grün gelesen
  wird"*) · slice-177 §3.2 (*„Der Vertrag steht **einmal** in der Sensor-Datei"*)
- **pfad:** `harness/sensors/image-scan.md`:5–9 gegen `tools/image-scan.sh`:63
- **befund:** Die Datei beschreibt den Lauf als *„gegen `ghcr.io/pt9912/a-check:latest`"*.
  Das Skript setzt `IMAGE_SCAN_REFS="${IMAGE_SCAN_REFS:-ghcr.io/pt9912/a-check:latest
  pt9912/a-check:latest}"` und durchläuft **beide** Referenzen; die zweite ist der
  Docker-Hub-Spiegel. Der Prüfbereich ist damit größer als der genannte — und da die
  Datei ausdrücklich als die **eine** Stelle eingeführt wird, an der der Vertrag steht,
  ist die Untertreibung jetzt die einzige Aussage. Die Ziel-Form verlangt an dieser
  Stelle zusätzlich das Kommando statt einer eingefrorenen Angabe; die Datei nennt
  weder Zahl noch Kommando.
- **verifizierbar:** ja — `grep IMAGE_SCAN_REFS tools/image-scan.sh`; kein Gate.

### F-5 — Der Abschnitt `Sperren` in `image-scan.md` führt zwei Lauf-Ausgänge und lässt die einzige echte Sperre aus

- **kategorie:** MEDIUM
- **quelle:** `v6.5.0` · `templates/harness/sensors/gate.template.md` §Sperren
  (*„Woran der Lauf abbricht, **bevor er etwas tut**. Je Sperre ihr Name, wie ihn die
  Abbruch-Meldung nennt"*)
- **pfad:** `harness/sensors/image-scan.md`:39–42 gegen `tools/image-scan.sh`:123–126
- **befund:** Gelistet sind *„kein Netz → Exit 2"* und *„Image nicht publiziert →
  Exit 2"*. Beides sind Ergebnisse eines bereits gelaufenen Scans (`errored=1` nach dem
  Trivy-Aufruf, ausgewertet erst am Schleifenende) und stehen inhaltlich schon in der
  Exit-Tabelle darüber. Die einzige Bedingung, die das Skript **vor** jeder Arbeit
  abbricht, trägt einen Namen in der Abbruch-Meldung — *„image-scan:
  `IMAGE_SCAN_REFS` ist leer — nichts zu pruefen ist KEIN gruener Befundstand"*, Exit 2
  — und fehlt in der Datei. Damit ist der Fail-closed-Riegel gegen die leere Prüfmenge,
  den das Skript ausdrücklich als Norm-Anleihe bei `verify-risiko-ausgaenge` begründet,
  im Vertrag nicht vertreten.
- **verifizierbar:** ja — `sed -n '118,130p' tools/image-scan.sh`; kein Gate.

### F-6 — `harness/sensors/doc-check.md` führt in der Bindung die Kennung `SL-002`, die es im Beobachtungs-Register seit dessen Migration nicht mehr gibt — unverlinkt und nirgends auflösbar

- **kategorie:** MEDIUM
- **quelle:** `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
  (*„Die Kennung ist der Pfad `BEO-<KUERZEL>/<slug>` … **Eine fortlaufende Nummer gibt
  es nicht mehr**"*) · `AGENTS.md` §5 (Agenten referenzieren IDs, sie erfinden keine)
- **pfad:** `harness/sensors/doc-check.md`:49
- **befund:** Die Bindungs-Zeile lautet *„`SL-002` seit slice-080"*. `SL-002` ist eine
  Kennung der mit slice-139 abgelösten Tabellenform des Registers; die heutige
  `observations/README.md` führt sie nicht, und im Register existiert kein Eintrag
  dieses Namens. In `AGENTS.md` und `harness/README.md` steht dieselbe Kennung wenigstens
  als Link auf die Register-README (die sie ebenfalls nicht definiert); in der neuen
  Datei ist sie nur noch Inline-Code. Eine frisch angelegte Datei führt damit ein
  zurückgezogenes Kennungsschema fort.
- **verifizierbar:** nein — die `ids`-Linkpflicht in `.d-check.yml` deckt
  `ADR-`/`MR-`/`AC-`/`SPEC-`/`ARC-`, nicht `SL-`; `make doc-check` ist grün.

### F-7 — Beide vorgelagerten Schritte in §9 stehen weiter auf „entsteht mit dem Übergang nach `in-progress/`", obwohl der Übergang in diesem Commit stattgefunden hat

- **kategorie:** MEDIUM
- **quelle:** `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Zwei Schritte vor der
  Modus-Begründung (*„Sie hängen weder am Modus noch am Slice-Typ und stehen deshalb in
  **jedem** Slice-Plan — auch bei reinem Refactor, auch wenn am Ende ‚alles GF‘
  dasteht"*)
- **pfad:** `docs/plan/planning/done/slice-177-sensors-struktur-zwei-tabellen.md`:225–234
- **befund:** §9 trägt für *Sub-Area-Wahl prüfen* und für *offene Beobachtungen sichten*
  je einen Vorwärtsverweis auf einen Zeitpunkt, der mit demselben Commit eingetreten
  ist; ausgefüllt ist keiner der beiden. Der Plan ist im selben Commit vollständig
  neugeschrieben und bis zur ausgefüllten Closure-Notiz geführt — die Vertagung
  beschreibt also keinen offenen Rest, sondern einen übersprungenen Schritt. Die
  Sichtung hat inhaltlich in §2 stattgefunden (Zitat des `BEO-HARNESS`-Eintrags), steht
  aber nicht dort, wo der Lese-Schritt sie sucht; die Sub-Area-Wahl ist nirgends
  geprüft, die Schluss-Zeile *„Alle berührten Sub-Areas GF"* nennt sie nicht einmal.
- **verifizierbar:** nein — `make verify` ist grün; `doc-structure` prüft die Existenz
  der Abschnitte, nicht ihre Ausfüllung.

### F-8 — Der Risiko-Ausgang „eingetreten" trägt weder Carveout noch Folge-Slice-ID

- **kategorie:** MEDIUM
- **quelle:** `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Offene Risiken werden
  bei Closure aufgelöst (*„*eingetreten* → Carveout … **oder Folge-Slice mit ID**"*;
  *„**Urteil** bleibt, ob der eingetragene Ausgang *trägt*"*)
- **pfad:** `docs/plan/planning/done/slice-177-sensors-struktur-zwei-tabellen.md`:164–171
  und :214–217
- **befund:** Risiko 1 (*„eine Regel steht da, die der eigene Bestand bricht"*) trägt den
  Ausgang *„eingetreten und **nur zum Teil aufgelöst**, Folge-Slice steht aus"*. Ein
  Carveout unter `docs/plan/carveouts/` gibt es nicht, eine Folge-Slice-ID auch nicht;
  §8 bestätigt das ausdrücklich (*„zwei benannt, beide noch ohne Datei … Sie entstehen,
  wenn der Maintainer sie priorisiert"*). Der Bestand, der die Ziel-Form bricht, ist
  damit nach der Closure ohne Adressaten — genau die Lage, gegen die die geschlossene
  Dreier-Menge steht. §1 sagt *„Bestand bleibt **bis dahin** bewusst stehen"*, ohne dass
  ein „dahin" existiert.
- **verifizierbar:** nein — `make verify-risiko-ausgaenge` ist grün: es prüft, *dass*
  einer der drei Ausgänge dasteht, nicht ob er trägt (so in `AGENTS.md` §4 auch
  ausgewiesen).

### F-9 — Die neue Nicht-Gate-Tabelle lässt `make doc-tracked` aus, das nach ihrem eigenen Kriterium hineingehört

- **kategorie:** MEDIUM
- **quelle:** slice-177 §3.3 (Definition der Tabelle: Targets, die *bewegen · messen ·
  sagen*) · `v6.5.0` · `regelwerk/modul-13-quality-gates.md` §Vorhanden ≠ behauptet
- **pfad:** `harness/README.md`:98–113 (neue Tabelle) gegen `AGENTS.md` §4
  (`make doc-tracked`, `make doc-commits`) und `d-check.mk`
- **befund:** Aufgenommen sind `doc-repair`, `doc-trace`, `doc-doctor`, `doc-usage`,
  `doc-help`. `make doc-tracked` (*„Getrackt-Status auflösbarer Referenz-Ziele"*, in
  keinem Aggregat, advisory) erfüllt dasselbe Kriterium und steht in `AGENTS.md` §4,
  aber in keiner der beiden `harness/README.md`-Tabellen; ebenso `make doc-commits`.
  Die Tabelle, die den Nicht-Gate-Bestand vollständig aufnehmen soll, ist damit an der
  Stelle unvollständig, an der sie neu ist.
- **verifizierbar:** teilweise — `make doc-targets` prüft nur die Richtung
  „dokumentiertes Target existiert real" und „reales **Gate**-Target ist in `AGENTS.md`
  §4 gelistet"; die hier fehlende Richtung ist ungewächtert.

### F-10 — Die Closure-Notiz sagt „keine Beobachtung angefallen", während im Register ein Beleg auf den Namen dieses Slice liegt und §2 desselben Dokuments mit ihm rechnet

- **kategorie:** MEDIUM
- **quelle:** `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
  (*„Der Zähler wird abgeleitet … es gibt kein Feld, in das man ihn schreibt"*) ·
  Reviewer-Skill §Klassifikation
- **pfad:** `docs/plan/planning/done/slice-177-sensors-struktur-zwei-tabellen.md`:207–212
  gegen :57–58 und
  `docs/plan/planning/observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/`
- **befund:** §8 führt *„**Beobachtungs-Register:** keine Beobachtung angefallen"*. §2
  desselben Dokuments zitiert dagegen `evidence/slice-177.md` in eben diesem Eintrag
  und rechnet *„damit 2×"*. Die Datei existiert, ist mit `199a12c` committet, trägt
  `**Vorgang:** slice-177` und benennt sich selbst als *„Zweite Instanz"*. §7 formuliert
  präziser (*„kein **neuer** Beleg"*) — §8 lässt das Wort weg und wird damit unwahr.
  Hinzu kommt: `state.md` des Eintrags steht auf *„offen (1×)"*, während zwei
  Evidence-Dateien darin liegen. Ein gespeicherter Zähler existiert im Repo nur bei
  sechs von 51 Einträgen, und zwei davon sind veraltet — der Slice liest genau einen
  dieser beiden und lässt ihn stehen.
- **verifizierbar:** teilweise — die Zahl der Evidence-Dateien ist auszählbar;
  `make verify-observations` prüft nur, dass zitierte Pfade existieren und `evidence/`
  nicht leer ist, nicht den Zählerstand.

---

## LOW

### F-11 — Drei der vier Dateien tragen die Klasse, die die Ziel-Form ausdrücklich ausschließt: womit das Werkzeug selbst gedeckt ist

- **kategorie:** LOW
- **quelle:** `v6.5.0` · `templates/harness/sensors/gate.template.md` §Regeln dieser
  Datei (*„**Was hier NICHT steht:** womit das Werkzeug selbst gedeckt ist — welcher
  Test welche Hälfte trägt … Diese Datei sagt, wie ein Lauf zu **lesen** ist"*)
- **pfad:** `harness/sensors/symlink-check.md`:12–14 · `harness/sensors/image-scan.md`:37 ·
  `harness/sensors/archive-wave.md`:40–41
- **befund:** `symlink-check.md` erklärt, dass beide Prüfungen durch dieselbe Funktion
  laufen, *„die auch der Selbsttest aufruft"*; `image-scan.md` nennt
  *„Selbsttest der Auswertung: `bash tools/image-scan.sh --selftest`"*;
  `archive-wave.md` führt in der Bindung *„Testsuite über `make archive-wave-test`"*.
  Alle drei beantworten „ist das Werkzeug richtig?", nicht „wie ist ein Lauf zu lesen?".
  Positiv anzumerken ist die Gegenrichtung: die Angabe *„mutations-kalibriert"* aus der
  alten `harness/README.md`-Zelle wurde beim Umbau korrekt **nicht** übernommen.
- **verifizierbar:** nein — Urteil über die Zugehörigkeit einer Aussage.

### F-12 — `doc-check.md` friert eine Zahl ein, die die Ziel-Form an derselben Stelle ausdrücklich ins Kommando verweist

- **kategorie:** LOW
- **quelle:** `v6.5.0` · `templates/harness/sensors/gate.template.md` §Grenze
  (*„Eine eingefrorene Zahl stünde hier falsch, sobald jemand committet"*)
- **pfad:** `harness/sensors/doc-check.md`:23
- **befund:** Loch 2 spricht von *„Ihre 10 Baseline-Pins"* in `harness/conventions.md`.
  Nachgemessen trägt die Datei heute **13** Vorkommen des Pin-Musters auf **10** Zeilen;
  welche der beiden Lesarten die Zahl meint, sagt der Satz nicht. Die Angabe ist aus der
  alten `AGENTS.md`-Zelle übernommen, wo sie zum Messzeitpunkt von slice-173 stand — sie
  ist damit derselbe Fall, gegen den die Ziel-Form drei Zeilen weiter argumentiert.
- **verifizierbar:** ja — `grep -c` auf das Pin-Muster; kein Gate.

### F-13 — `doc-check.md` §Sperren beginnt mit einer Sperre, die für dieses Target nicht gilt

- **kategorie:** LOW
- **quelle:** `v6.5.0` · `templates/harness/sensors/gate.template.md` §Sperren
- **pfad:** `harness/sensors/doc-check.md`:41–42
- **befund:** Der erste Punkt lautet *„`Range-Basis nicht auflösbar` — nur bei den
  Range-gebundenen Modulen, **nicht hier**; `doc-check` läuft ohne Commit-Range."* Eine
  Sperre, die nicht greift, ist in einer Liste von Abbruch-Bedingungen keine Antwort auf
  die Frage des Abschnitts; sie beschreibt ein anderes Target.
- **verifizierbar:** nein.

### F-14 — Drei Aussagen sind beim Kürzen weder in der Zelle noch in der Sensor-Datei angekommen

- **kategorie:** LOW
- **quelle:** slice-177 §3.2 (*„Der Vertrag steht **einmal** in der Sensor-Datei"*)
- **pfad:** `harness/sensors/doc-check.md`:26–30 · `harness/sensors/archive-wave.md`:15 ·
  `harness/sensors/symlink-check.md`:16–21
- **befund:** Verglichen wurden die alten Zellen (`git show HEAD~1:AGENTS.md`,
  `git show HEAD~1:harness/README.md`) gegen neue Zelle **plus** Datei. Nicht
  übernommen: (a) *„`state.md` ist veränderlich und bleibt geprüft"* — die
  Einschränkung, die den Ausnahme-Glob des Beobachtungs-Registers auf zwei der drei
  Datei-Rollen begrenzt; ohne sie liest sich Loch 3 weiter als es ist. Die Aussage steht
  nur noch als Kommentar in `.d-check.yml`. (b) Der Zeiger auf
  `tools/archive-wave/README.md`. (c) Die Messung aus slice-173, dass ein umgebogener
  Symlink `doc-check` bei **0** Befunden ließ — der empirische Grund für die Existenz
  des Sensors; sie steht nur noch im Plan von slice-173 in `done/`.
- **verifizierbar:** ja — Diff der beiden Stände; kein Gate.

### F-15 — `Verantwortlich:` steht auf „noch nicht priorisiert", während die Datei in `in-progress/` liegt

- **kategorie:** LOW
- **quelle:** `v6.5.0` · `regelwerk/modul-05-planning-harness.md` §Lifecycle als State
  Machine (*„`open → next` setzt den Verantwortlichen"*)
- **pfad:** `docs/plan/planning/done/slice-177-sensors-struktur-zwei-tabellen.md`:16
- **befund:** Der Kopf führt `**Verantwortlich:** — *(noch nicht priorisiert)*`. Der
  Slice ist priorisiert, in Arbeit und bis zur Closure-Notiz geschrieben; das Feld sagt
  das Gegenteil. Die Baseline weist es ausdrücklich als Deklaration ohne Sensor aus.
- **verifizierbar:** nein — `doc-structure` prüft die Anwesenheit des Feldes, nicht
  seinen Wert; `make verify` ist grün.

### F-16 — `make record-gates` bleibt ohne ein Wort in der Gate-Tabelle, obwohl es nach dem Kriterium des Slice nicht urteilt

- **kategorie:** LOW
- **quelle:** slice-177 §3.3 (Nicht-Gate = *„urteilt nicht über den Zustand des Repos"*)
- **pfad:** `harness/README.md`:82 · `tools/harness/record-gates.sh`
- **befund:** Das Target schreibt einen Working-Tree-Hash nach
  `.harness/state/gates-passed.diffsha` und meldet immer Erfolg — es *bewegt*, es
  urteilt nicht. Es bleibt in der Gate-Tabelle, ohne dass der Slice die Abgrenzung
  anspricht. Dass es als letzter `gates`-Prerequisite Teil des Handoff-Gates ist, ist
  ein tragfähiges Gegenargument — nur steht es nirgends.
- **verifizierbar:** nein — Einordnungsfrage.

---

## INFO

### F-17 — Fünf reale Gate-Targets fehlen weiterhin ganz in `harness/README.md` §Sensors

- **kategorie:** INFO
- **pfad:** `harness/README.md` §Sensors gegen `Makefile`
- **befund:** `make doc-reviews`, `make suppression-check`,
  `make verify-risiko-ausgaenge`, `make commit-scope-check` und `make archive-wave-test`
  fehlen in beiden Tabellen, obwohl die ersten beiden im `gates`-Aggregat hängen und in dessen Zeile
  namentlich aufgezählt werden. Bestand von vor diesem Slice; genannt, weil er die
  Vollständigkeits-Annahme der beiden Tabellen betrifft.
- **verifizierbar:** teilweise — `make doc-targets` prüft diese Richtung nicht.

### F-18 — Das DoD-Häkchen „Unabhängiger Review durchgeführt" war zum Commit-Zeitpunkt gesetzt, ohne dass ein Report existierte

- **kategorie:** INFO
- **pfad:** `docs/plan/planning/done/slice-177-sensors-struktur-zwei-tabellen.md`:143
- **befund:** Zum Stand `9e42685` trägt `docs/reviews/` Reports zu slice-174, slice-175
  und slice-176, keinen zu slice-177 — dieser hier ist der erste. Gehört zur
  Verifikation (`docs/reviews/README.md` §Abgrenzung zur Verifikation), nicht zum
  Review; hier nur als Hinweis an den Verifier. `make doc-reviews` greift erst in
  `done/`.
- **verifizierbar:** ja — durch `make doc-reviews` nach dem Lifecycle-Wechsel.

### F-19 — „steht seit `v5.12.0` unverändert" ist in diesem Repo nicht nachprüfbar

- **kategorie:** INFO
- **pfad:** `docs/plan/planning/done/slice-177-sensors-struktur-zwei-tabellen.md`:49–56
- **befund:** Der Satz stützt die Kernaussage von §2. Nachprüfbar ist im Repo nur der
  vendorte Stand `v6.5.0` — dort steht der Absatz wortgleich, geprüft. Die Aussage über
  `v5.12.0`, `v6.0.0` und `v6.2.0` beruht auf einer im Evidence-Beleg dokumentierten
  Messung im Kurs-Klon (`git show <tag>:lab/templates/AGENTS.template.md`) und ist hier
  weder bestätigt noch widerlegt — netzlos und ohne zweiten vendorten Stand nicht
  entscheidbar. Herkunft ist benannt, nicht behauptet; deshalb kein MEDIUM.
- **verifizierbar:** nein (ohne Kurs-Klon).

---

## Negativbefunde (geprüft, ohne Befund)

- **Exit-Codes in `image-scan.md`.** `0` / `1` / `2` gegen `tools/image-scan.sh`
  gehalten: `exit 0` nach *„keine behebbaren CRITICAL/HIGH"*, `exit 1` bei
  `findings=1`, `exit 2` bei `errored=1` — die Tabelle stimmt. Auch *„Der Vollbericht
  fällt nie"* trägt: der Berichtslauf läuft mit `--exit-code 0`, und sein Scheitern
  setzt `errored`, nicht `findings`. Die Aussage *„über `make` sind 1 und 2 nicht
  unterscheidbar"* ist die korrekte Lesart von `make`s Normalisierung.
- **Zwei-Prüfungen-Beschreibung in `symlink-check.md`.** Gegen `tools/symlink-check.sh`
  vollständig bestätigt: beide Prüfungen in **einer** Funktion (`kaputte()`), die der
  Selbsttest aufruft; Geltungsbereich = von git getrackte Symlinks; fail-closed über
  `adoptierter_stand()`; der Sperren-Name *„adoptierter Baseline-Stand nicht lesbar"*
  ist der Wortlaut der Abbruch-Meldung; die Schluss-Zeile nennt tatsächlich die Zahl der
  getrackten Symlinks. Auch die Zuordnung der drei Vorfälle (slice-173 = entferntes
  Ziel, slice-167 = noch vorhandener alter Stand) deckt sich mit dem Skriptkopf.
- **`WELLE=`-Sperre in `archive-wave.md`.** Der Wortlaut deckt sich mit der Messung aus
  slice-144, wie sie in der abgelösten `AGENTS.md`-Zelle stand; `RewriteFieldForMove`
  existiert in `tools/archive-wave/archive.go` und trägt die in Loch 2 beschriebene
  Feld-Umschreibung. Der sichere Default ohne `APPLY=1` ist korrekt wiedergegeben.
- **Drei der vier Löcher in `doc-check.md`.** Symlink-Blindheit (Loch 1) deckt sich mit
  der Begründung, aus der `symlink-check` entstand; die Ausnahmeliste (Loch 3) hat in
  `.d-check.yml` tatsächlich **fünf** Einträge; der Digest-Riegel (Loch 4) ist mit
  `DCHECK_DIGEST` in `d-check.mk` konsistent. Die `versions`-Tabelle bildet
  slice-133/slice-173 richtig ab.
- **Die Einordnung von `image-scan` als Gate außerhalb des Aggregats.** `v6.5.0` ·
  `regelwerk/modul-13-quality-gates.md` §Vorhanden ≠ behauptet sagt wörtlich, was §3.3
  des Plans daraus zitiert — und nennt sogar `regelwerk-check` und die advisory
  `doc-*`-Targets als Beispiele. Das Target fällt exit-codiert ein Urteil (1 bei
  behebbaren Funden); es gehört nicht in die Nicht-Gate-Tabelle. Die Entscheidung trägt.
- **Die zitierte Ziel-Form-Regel selbst.** `v6.5.0` ·
  `templates/AGENTS.template.md` §4 trägt den Absatz *„Diese Tabelle **listet auf**;
  definiert wird hier nichts. Die *Bindung* eines Targets … steht in
  `harness/README.md` §Sensors"* wortgleich zum Zitat in §2.
- **Verbote der Ziel-Form.** Keine der vier Dateien führt ein Status- oder Datumsfeld;
  es gibt kein `harness/sensors/done/`; der Template-Hinweis-Block und der Abschnitt
  *Regeln dieser Datei* sind in allen vier gelöscht; alle vier tragen Vertrag, Grenze
  und Bindung, `image-scan.md` zusätzlich die Exit-Tabelle. Die Index-Zeile ist überall
  erhalten und verlinkt die Datei — in **beiden** Tabellen.
- **Größen-Regel.** Drei Liefer-Punkte (Sensor-Dateien · zweite Tabelle · gekürzte
  Zellen), zwei berührte Schichten (`harness/`, `AGENTS.md`) — innerhalb der Schwelle
  aus `AGENTS.md` §5. Der Gate-Lauf steht als feste Zeile unter dem DoD.
- **Kennungs-Linkpflicht in den neuen Dateien.** `ADR-0037` (`image-scan.md`) und
  `MR-018` (`doc-check.md`) sind verlinkt; `MR-006` in der neuen Nicht-Gate-Zeile
  ebenfalls. `make doc-check` läuft grün über die vier neuen Dateien.
- **Commit-Disziplin außer §3.3.** Die Message nennt `slice-177` und `welle-15`
  (`make trace-check`-Muster erfüllt); der Scope `feat(harness)` ist nicht
  `(planning)`, die Scope-Regel greift also nicht.
- **Roadmap-Kopplung.** Der Ruhe-Marker *Nichts in Arbeit* ist durch den Zeiger auf
  slice-177 ersetzt worden; `make doc-planning` bestätigt die Kopplung.
- **Gate-Läufe.** `make gates` Exit **0** · `make verify` Exit **0** · `make doc-check`
  Exit **0** (483 Dateien, 0 Befunde — mit diesem Report im Baum). Alle unpiped in
  Dateien umgeleitet, Exit-Code separat gelesen. Die Behauptung der Commit-Message
  („`make gates` und `make verify` je Exit 0") ist damit unabhängig bestätigt.
- **Arbeitsbaum.** Es wurden **keine** Mess-Eingriffe am Gegenstand vorgenommen — alle
  Messungen sind lesend (`git`, `grep`, `sed`, `awk`, `make`). `git status --short` ist
  am Ende leer; die einzige Schreibwirkung ist `.harness/state/gates-passed.diffsha` aus
  `record-gates`, und dieser Pfad ist über `.gitignore:3` ausgenommen (`git check-ignore`
  bestätigt).

---

## Kategorie-Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 8 |
| LOW | 6 |
| INFO | 3 |

## Verdikt

**Nicht abschlussreif ohne Klärung von F-1 und F-2.**

Die handwerkliche Hälfte des Slice hält: Die vier Sensor-Dateien folgen der Ziel-Form,
ihre Verbote sind eingehalten, und die inhaltlich prüfbaren Aussagen stimmen mit dem
Repo überein — die Exit-Codes gegen `tools/image-scan.sh`, die Zwei-Prüfungen-Struktur
gegen `tools/symlink-check.sh`, die `WELLE=`-Sperre gegen `tools/archive-wave/`, drei
der vier Löcher in `doc-check.md`. Auch die Entscheidung, `image-scan` **nicht** in die
Nicht-Gate-Tabelle zu stellen, trägt und ist mit der richtigen Baseline-Stelle belegt.
Die Ausgangsmessung („16 von 40") ist auf die Zahl reproduzierbar.

Was nicht trägt, sind die beiden Aussagen, an denen der Slice seine Reichweite
festmacht. F-2: Die Zweiteilung „vier mit Ausgängen, zwölf mit Historie" ist als
Sortier-Ergebnis formuliert, hält der Nachmessung aber nicht stand — mindestens acht der
zwölf tragen eine ausdrücklich benannte Deckungsgrenze, also genau das erste Kriterium
der Ziel-Form; für „Vertrag plus Historie" bleiben zwei übrig. Das ist derselbe
Fehlertyp, den der Slice in §8 als Lerneintrag reklamiert, nur eine Ebene höher: nicht
die Länge wurde zum falschen Kriterium, sondern die Gruppierung zur unbelegten
Begründung. F-1 ist unabhängig davon und mechanisch: Der Lifecycle-`git mv` ist mit dem
Inhalts-Umbau in einen Commit gefallen, die Ähnlichkeit auf 40 % gedrückt und
`git log --follow` auf einen einzigen Commit reduziert — die Hard Rule §3.3 beschreibt
diesen Schaden wörtlich, und er ist eingetreten, nicht bloß möglich.

Die acht MEDIUM-Befunde haben zwei Ursachen. Vier betreffen die neuen Dateien selbst
(F-4, F-5, F-6 und, über die Vollständigkeit der neuen Tabelle, F-9): dort, wo der
Vertrag jetzt **einmal** stehen soll, ist er an drei Stellen enger, unvollständig oder
mit einer zurückgezogenen Kennung geschrieben. Vier betreffen den Plan als Dokument
(F-3, F-7, F-8, F-10): ein stehengebliebener Rest, der die eigene Wirkungsmessung
relativiert, zwei unausgefüllte Pflichtschritte und ein Risiko-Ausgang ohne Adressaten.
Auffällig ist, dass **kein** Gate einen dieser Befunde sieht — `make gates` und
`make verify` sind unabhängig bestätigt grün.

Die Entscheidung über Übernahme, Rückstellung oder Folge-Slice liegt beim Implementer;
dieser Report kategorisiert und schlägt nichts vor.
