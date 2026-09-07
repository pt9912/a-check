# Review-Report: slice-175 — 2026-09-07

**Review-Art:** Implementation — geprüft gegen den Slice-Plan, `AGENTS.md` §3/§4/§5/§6,
`harness/conventions.md` §Baseline und §Adaptions-Block sowie gegen die **selbst nachgerechnete**
Herkunft des vendorten Fremdtexts. Der Slice bringt fremden Text ins Repo; Schwerpunkt war deshalb
die Integrität des Vendorings und die Vollständigkeit der Zeiger-Umstellung.

**Gegenstand:** Commit-Range `97472ab^..f49913a` — `97472ab` (Lifecycle-`git mv`
`open/` → `in-progress/`) und `f49913a` (Etappe A: Vendoring, Stand-Deklaration, 13 lebende
Dateien, 4 Symlinks, `versions`-Muster). **HEAD ist während des Reviews weitergewandert**
(`4760348`, Planungs-Arbeit an slice-176 aus einer parallelen Sitzung); die Range oben bleibt der
Gegenstand, siehe I-3.

**Skill:** `.harness/skills/reviewer.md` @ Stand `HEAD` (in dieser Range nur in der
`Bezug:`-Zeile geändert) · <!-- d-check:ignore -->
**Modell:** claude-opus-5[1m] · **Datum:** 2026-09-07

**Review-Unabhängigkeit:** unabhängiger Lauf — eigenes Kontextfenster, die Umsetzung wurde nicht
von dieser Instanz verfasst. Jede Zahl unten ist selbst gemessen: der vendorte Baum gegen den
git-Tag `v6.5.0` in einem frischen Klon, gegen das Release-ZIP und gegen die GitHub-API; keine
Angabe ist dem Slice oder der Commit-Message entnommen.

**Eingangs-Kontext:**

- `docs/plan/planning/in-progress/slice-175-etappe-a-vendoring-v650.md`
- `docs/plan/planning/welle-15-regelwerk-v650-migration.md`
- `AGENTS.md` §3 (Hard Rules), §4 (Quality Gates), §5 (Dokumentations-Regeln), §6 (Workflow)
- `harness/conventions.md` §Baseline, §Adaptions-Block, §Modus-Deklaration
- `harness/conventions/MR-011`, `MR-012`, `MR-014`, `MR-015`, `MR-016`, `MR-019`, `MR-020`
- `.d-check.yml` (Module `versions`, `links`, `ids`, `reviews`), `tools/regelwerk-check.sh`,
  `tools/symlink-check.sh`
- `docs/plan/planning/observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/`,
  `…/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/`,
  `…/BEO-GATE/versions-sensor-trifft-planungs-vorgriff/`,
  `…/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/`
- `docs/plan/planning/done/slice-167-etappe-a-vendoring-v620.md` (Präzedenz derselben Etappe)
- `docs/reviews/2026-09-07-slice-174-regelwerk-v650-delta-analyse.md` (frühere Findings am Bereich)
- adoptierte Baseline: `v6.5.0` · `regelwerk/modul-05-planning-harness.md`,
  `regelwerk/modul-06-roadmap.md`, `regelwerk/modul-08-agentenrollen.md`,
  `regelwerk/modul-10-review-harness.md`
- Kurs-Repo `pt9912/ai-harness-course`, frischer Klon mit Tags; Release-Assets zu `v6.5.0`

**Eigene Läufe (Repo-Zustand am Ende unverändert):** `make regelwerk-check` Exit 0 ·
`make doc-check` Exit 0 · `make symlink-check` Exit 0 · `make gates` Exit 0 · `make verify` Exit 0 ·
`make commit-scope-check RANGE=199a12c..f49913a` Exit 0 · `make doc-immutable
RANGE=199a12c..f49913a` Exit 0 · `make trace-check` Exit 2 (Umgebungsausfall, siehe I-2) ·
`sha256sum -c` gegen beide `SHA256SUMS` · `gh release download v6.5.0` (ZIP + Manifest) ·
`diff -rq` Tag/ZIP gegen den vendorten Baum · **eine Mutations-Sonde** (temporäre Markdown-Datei,
danach entfernt; `git status --porcelain` vor und nach dem Review leer, siehe F-3).

---

## Findings

### F-1 — `state.md` einer Beobachtung behauptet einen Zustand, den dieser Commit aufgehoben hat

- `kategorie`: **HIGH**
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
  („`state.md` — der veränderliche Stand … keine Chronik"); `AGENTS.md` §3.7
  (*Dieselbe Regel für Zustandsfelder*)
- `pfad`: `docs/plan/planning/observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/state.md`:6-7
- `befund`: Die Datei sagt im Präsens *„Der beobachtete **Zustand** ist beseitigt: genau ein Stand
  liegt vendored, `make regelwerk-check` meldet keinen ungeprüften mehr."* Nach `f49913a` liegen
  **zwei** Stände unter `.harness/baseline/` (`v6.2.0`, `v6.5.0`), und `make regelwerk-check` gibt
  aus: *„HINWEIS — 2 vendored Staende gefunden; geprueft wird der hoechste (v6.5.0). ungeprueft
  bleiben: v6.2.0"*. Vor `f49913a` lag genau ein Stand (`git ls-tree -d f49913a^
  .harness/baseline/`), der Satz war also bis zu diesem Commit wahr; der Commit hebt ihn auf und
  fasst die Datei nicht an, obwohl er im selben Zug ein anderes Verzeichnis desselben Registers
  anlegt. Ein Chronik-Lesart („galt bei slice-172") scheidet aus: `state.md` ist ausdrücklich der
  veränderliche Stand — `.d-check.yml` nimmt `observation.md` und `evidence/**` vom
  `versions`-Muster aus und lässt `state.md` bewusst geprüft *(„state.md bleibt geprueft, es ist
  der veraenderliche Stand")*.
- `verifizierbar`: **nein** — `make verify-observations` prüft nur Deckung (Verzeichnis existiert,
  `evidence/` nicht leer), nicht den Wahrheitsgehalt von `state.md`; `versions` greift nicht, weil
  der Satz Prosa und kein Pfad-Literal ist.
- *Adversarische Gegenprobe:* Der **Zähler** bleibt zu Recht bei 1× — die Beobachtung lautet „…
  obwohl die Migration **geschlossen** ist", und `welle-15` ist offen. Das Finding betrifft nicht
  den Zähler, sondern die beiden Tatsachensätze daneben. Dieselbe Defektklasse steht ein zweites
  Mal im Register: `…/BEO-GATE/trace-check-lokal-nicht-befragbar/state.md` sagt „offen (1×)", das
  `evidence/`-Verzeichnis trägt zwei Dateien (bereits als F-16 im Vorgänger-Report notiert).

### F-2 — Fünf `Ersetzt-Baseline-Regel`-Zeiger sind ohne die vorgeschriebene Messung mitgewandert; einer ist nicht wortgleich

- `kategorie`: **HIGH**
- `quelle`: `harness/conventions.md` §Baseline — *„Bedingung ist, dass der referenzierte Abschnitt
  im neuen Stand wortgleich ist, und das ist zu **messen**, nicht anzunehmen. Ist er es nicht,
  trägt die Stelle die Abweichung sichtbar, statt still umzuziehen."*
- `pfad`: `harness/conventions/MR-015-welle-closure-ohne-replay.md`:4 (sowie `MR-011`:3,
  `MR-012`:4, `MR-014`:3, `MR-016`:3 und die Spiegel-Spalte `harness/conventions.md`:120-124)
- `befund`: Alle fünf Zeiger sind in `f49913a` von `v6.2.0` auf `v6.5.0` umgestellt worden. Die
  Messung ist selbst nachgeholt: vier der fünf Abschnitte sind wortgleich (nur die
  Markdown-Trennzeile einer Tabelle wechselt von `|---|` zu `| --- |`), **einer ist es nicht** —
  in `regelwerk/modul-06-roadmap.md` §Wellen-Closure-Prozedur (Ziel von `MR-015`) ersetzt `v6.5.0`
  im Eröffnungs-Absatz *„im Slice-Plan (Modul 9)"* durch *„im Slice-Plan (§1 Ziel und Abgrenzung,
  `modul-05-planning-harness.md` §Ziel-Form: Slice; in der Plan-Ausgabe des Laufs
  `modul-09-implementierung.md`)"* — eine Zeile weniger, drei mehr. Weder Commit-Message noch
  Slice-Plan noch ein `MR`-Eintrag hält eine Messung oder die Abweichung fest; der Zeiger ist
  still umgezogen. Die vorige Etappe hat das anders gehalten: `slice-167`:109 vermerkt
  ausdrücklich *„weil jeder referenzierte Abschnitt dort wortgleich steht (gemessen,
  `slice-172` §2.2)"*.
- `verifizierbar`: **nein** — das `versions`-Muster erzwingt nur das *Mitwandern* des Zeigers, nicht
  die Wortgleichheit des Ziels; kein Target misst sie.
- *Adversarische Gegenprobe:* Die geänderten Sätze berühren die Replay-Zusage nicht, die `MR-015`
  ersetzt (Closure-Schritt 1 ist unverändert) — die *Substanz* der Adaption trägt also weiter. Die
  Konvention stellt die Bedingung jedoch auf den **Abschnitt**, nicht auf den Teilsatz, und
  verlangt bei Abweichung eine sichtbare Notiz; genau die fehlt. Ohne Messung ist zudem nicht
  entscheidbar, ob die Substanz trägt — die Aussage entsteht erst durch das Nachrechnen.

### F-3 — Das erweiterte `versions`-Muster meldet fremde Pfade als `version-stale`

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §4 (`make doc-check`, `versions`); Register-Klasse
  `BEO-GATE/versions-sensor-trifft-planungs-vorgriff`
- `pfad`: `.d-check.yml`:118
- `befund`: Das Muster lautet nach `f49913a` `(?:^|[^-a-zA-Z0-9_.])baseline/(v\d+\.\d+\.\d+)/` und
  verlangt kein `.harness/`-Präfix mehr. Eine Mutations-Sonde (temporäre Markdown-Datei im
  Repo-Wurzelverzeichnis, `make doc-check`, danach entfernt) zeigt: die Fremd-URL
  `https://example.org/baseline/v1.0.0/spec` wird als *„Versions-Pin trägt v1.0.0, erwartet
  v6.5.0"* beanstandet — ein Treffer, den das alte Muster **nicht** hatte. Ebenso trifft das
  Muster das Fixture-Literal `.harness/baseline/v0.0.1/…`, das `tools/symlink-check.sh`:75-82
  führt; heute folgenlos, weil `d-check` nur Markdown liest. Der in
  `…/muster-trifft-nur-die-haeufige-schreibweise/evidence/slice-175.md` festgehaltene Beleg
  („Mutations-Probe belegt beide Richtungen") misst Fehler-setzen und Fehler-zurücknehmen an der
  beabsichtigten Form; die Über-Treffer-Richtung ist nicht gemessen. Negativ mitgemessen: die
  Wort-Abgrenzung wirkt — `my-baseline/v1.2.3/`, `xbaseline/v9.9.9/` und `foo.baseline/v8.8.8/`
  lösen korrekt nicht aus.
- `verifizierbar`: **ja** — `make doc-check` gegen eine Fixture-Datei mit den vier Fällen; heute
  liefe sie rot, ohne dass ein a-check-Pin veraltet wäre.

### F-4 — Der Sensor-Eingriff ist ein Liefer-Punkt, den weder §1 noch die DoD des Slice führt

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5 (Slice-Form, Liefer-Punkte); `v6.5.0` ·
  `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice
- `pfad`: `docs/plan/planning/in-progress/slice-175-etappe-a-vendoring-v650.md` §1 und §4;
  `.d-check.yml`:115-118
- `befund`: §1 sagt zu: *„Der Stand `v6.5.0` liegt vendored, die Stand-Deklaration nennt ihn an
  ihren drei Stellen, und die vier Baseline-Symlinks … zeigen darauf."* Die DoD nennt dieselben
  drei Dinge. Geliefert ist zusätzlich eine Änderung am `versions`-Muster in `.d-check.yml` plus
  ein neues Register-Verzeichnis — ein dritter Liefer-Punkt in einer zweiten Sub-Area
  (`GATE` neben `HARNESS`), den der Plan weder unter „Nicht in diesem Slice" ausschließt noch
  aufnimmt. Die **Größen-Regel selbst hält** (3 Liefer-Punkte, 2 Schichten), aber der Plan ist
  nicht nachgezogen; nach `v6.5.0` · `regelwerk/modul-05-planning-harness.md` ist die Abgrenzung
  „die Grenze, an der ein wachsender Slice sich messen lässt".
- `verifizierbar`: **nein** — `make doc-structure` zählt DoD-Punkte, nicht das, was daneben
  geliefert wurde.

### F-5 — §2 *Analyse* und §3 *Umsetzung* des Slice sind zum Review-Handoff noch `(offen)`

- `kategorie`: LOW
- `quelle`: `v6.5.0` · `regelwerk/modul-08-agentenrollen.md` §Die neun Übergaben
  (Implementer→Reviewer: „PR mit Diff + Plan-Verweis")
- `pfad`: `docs/plan/planning/in-progress/slice-175-etappe-a-vendoring-v650.md` §2, §3
- `befund`: Beide Abschnitte tragen unverändert den Vorlagen-Platzhalter. Sämtliche Messungen des
  Slice (Herkunft des Vendorings, 13 lebende Dateien, 16 Zeitdokumente, Mutations-Probe) stehen
  nur in der Commit-Message und in der Evidence-Datei. Die Vergleichs-Slices derselben Bauart
  führen sie im Plan: `slice-167` §3 *Umsetzung* trägt rund 70 Zeilen, `slice-174` §2/§3
  ebenfalls. Für den Reviewer heißt das, dass es gegen den Plan nichts zu prüfen gibt und jede
  Zahl von außen nachgerechnet werden muss.
- `verifizierbar`: **nein** — `doc-structure` prüft Kopffelder, Größen-Regel und Closure-Struktur,
  nicht den Füllstand von §2/§3.

### F-6 — Eine lebende Planungs-Datei spricht weiter von „fünf `v6.0.0`-Ankern"

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §3.7 (Zustandsaussagen im Indikativ über den geltenden Zustand)
- `pfad`: `docs/plan/planning/open/slice-171-mr011-begruendung-und-rollen-tabelle.md`:159
- `befund`: Das Risiko in §7 lautet *„Die Ablösung von `MR-011` entfernt einen der fünf
  `v6.0.0`-Anker im Feld `Ersetzt-Baseline-Regel`"*. Die fünf Anker standen zuletzt auf `v6.2.0`
  und seit `f49913a` auf `v6.5.0`. Die Angabe war vor diesem Commit bereits veraltet; der Commit
  hat genau diese fünf Zeilen angefasst und die Aussage nicht mitgezogen. `open/` ist keine
  eingefrorene Klasse — die `versions`-`exempt-paths` nehmen `done/`, `docs/reviews/` und
  `conventions/done/` aus, `open/` nicht.
- `verifizierbar`: **nein** — die Zahl steht als Prosa („fünf … `v6.0.0`-Anker"), nicht als
  Pfad-Literal; kein Modul liest sie.

### F-7 — Die Klassen-Grenze der neuen Beobachtung ist eine Menge-Frage, kein Art-Unterschied

- `kategorie`: LOW
- `quelle`: `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register („leiten zwei
  Schreiber denselben Bereich unterschiedlich ab, entstehen zwei Pfade und die Beobachtung teilt
  sich still")
- `pfad`: `docs/plan/planning/observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md`
- `befund`: Die Abgrenzung gegen `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf` lautet: *„Dort ist
  die Prüfmenge **leer** … hier läuft er, hat Gegenstand, und meldet für die Mehrheit der Fälle
  korrekt."* Das Unterscheidungsmerkmal ist damit die **Trefferzahl** (0 vs. > 0), nicht die
  Ursache — und die Ursache ist in beiden Fällen dieselbe: ein Muster, das eine Schreibweise des
  Gegenstands verfehlt. Der erste Beleg der älteren Klasse (`verify-ac-form` suchte
  `^**Happy Path:**`, `evidence/slice-120.md`) ist genau dieser Fall mit Trefferzahl null. Läge
  der nächste Vorfall bei null Treffern, wäre nicht entscheidbar, in welches Verzeichnis er
  gehört; der Zähler teilte sich still — die Lage, gegen die die Pfad-Kennung erfunden wurde.
  Zugunsten der Anlage: die Abgrenzung ist ausdrücklich formuliert und nicht stillschweigend
  vollzogen.
- `verifizierbar`: **nein** — „ist das dieselbe Beobachtung?" ist laut Regelwerk ausdrücklich das
  menschliche Urteil; maschinell prüfbar ist nur die Deckung.

### I-1 — Vendoring-Integrität über vier unabhängige Kanäle bestätigt

- `kategorie`: INFO
- `quelle`: `harness/conventions.md` §Baseline / `MR-006`
- `pfad`: `.harness/baseline/v6.5.0/`
- `befund`: (1) `sha256sum -c SHA256SUMS` im vendorten Baum: 54/54 OK, Exit 0; die Liste deckt den
  Baum exakt (`diff` der Pfadmengen leer) und listet sich selbst nicht — dieselbe Konvention wie
  `v6.2.0` (53 Zeilen, relative Pfade). (2) `gh api …/releases/tags/v6.5.0` nennt für
  `lab-regelwerk.zip` den Digest `80684c17…d18865`. (3) Das mitgelieferte Release-Manifest
  `SHA256SUMS` (84 Byte) trägt denselben Wert. (4) Das heruntergeladene ZIP hashed selbst auf
  denselben Wert, und `diff -rq` seines Inhalts gegen den vendorten Baum ist leer. Zusätzlich
  gegen den git-Tag: Pfadbaum deckungsgleich (54/54); 28 Dateien weichen inhaltlich ab, in
  **allen** 29 geänderten Zeilenpaaren ausschließlich mechanisch — relative Repo-Pfade werden zu
  `https://github.com/pt9912/ai-harness-course/blob/v6.5.0/…`, und zweimal wird
  `releases/latest/download` zu `releases/download/v6.5.0`. **Keine** Zeile mit inhaltlicher
  Abweichung. Neu gegenüber `v6.2.0` ist genau eine Datei
  (`templates/harness/sensors/gate.template.md`), wie behauptet.
- `verifizierbar`: **ja** — `make regelwerk-check` für die Integritäts-Hälfte; die Herkunfts-Hälfte
  ist Netz und bleibt bewusst außerhalb der Gates.

### I-2 — `make trace-check` ist in diesem Arbeitsbaum weiterhin nicht lauffähig

- `kategorie`: INFO
- `quelle`: `AGENTS.md` §5; `BEO-GATE/trace-check-lokal-nicht-befragbar`
- `pfad`: `Makefile` `trace-check`
- `befund`: Jede geprüfte Range bricht mit *„d-check: error: Range-Basis-Vorfahren nicht lesbar:
  object not found"* ab — die Default-Range, `199a12c..f49913a` (kurz und als voller SHA),
  `HEAD~3..HEAD`, `HEAD~20..HEAD~15` und `HEAD~50..HEAD~48`. Ein Verdacht ist ausgeschlossen: das
  Verschieben der inkrementellen Commit-Graph-Kette (`.git/objects/info/commit-graphs`, danach
  zurückgestellt) ändert nichts. Der Ausfall ist damit weder Range- noch Commit-spezifisch und
  nicht von dieser Range verursacht; er ist als F-16 des Vorgänger-Reports bereits notiert. Die
  Traceability wurde per Sichtprüfung bestätigt: `97472ab` nennt `slice-175`, `f49913a` nennt
  `slice-175, welle-15`.
- `verifizierbar`: **ja** (in einem funktionierenden Klon) — der Ausfall selbst ist reproduzierbar.

### I-3 — HEAD ist während des Reviews weitergewandert

- `kategorie`: INFO
- `quelle`: —
- `pfad`: `4760348` *(docs(planning): slice-176 traegt den Loesch-Schritt)*
- `befund`: Beim Start des Reviews war `f49913a` HEAD und der Arbeitsbaum sauber. Während der
  Prüfung entstanden aus einer parallelen Sitzung Änderungen an `open/slice-176-…`, an
  `…/BEO-GATE/versions-sensor-trifft-planungs-vorgriff/state.md` und ein neues
  `evidence/slice-176.md`; sie sind inzwischen als `4760348` committet. `make gates` und
  `make verify` liefen deshalb gegen `4760348`, nicht gegen `f49913a`. Beide Exit 0. Nicht Teil
  des Gegenstands, aber für die Reproduzierbarkeit der Lauf-Belege festgehalten: der neue
  slice-176-Absatz verweist auf *„slice-175 §1 (Maintainer-Entscheidung 2026-09-07)"*, und §1 von
  slice-175 trägt in `f49913a` keinen solchen Punkt.
- `verifizierbar`: **ja** — `git log`, `git status`.

### I-4 — „16 Zeitdokumente" ist die Zahl **vor** der eigenen Register-Datei

- `kategorie`: INFO
- `quelle`: Commit-Message `f49913a`
- `pfad`: —
- `befund`: Bei `f49913a` tragen **17** Markdown-Dateien außerhalb der vendorten Bäume einen Pfad
  `baseline/v6.2.0/` (`git grep -l`). 16 davon bestanden vorher; die siebzehnte ist
  `…/muster-trifft-nur-die-haeufige-schreibweise/evidence/slice-175.md`, die dieser Commit selbst
  anlegt und die als Evidence-Datei ab Merge unveränderlich ist. Die Aussage ist damit zum
  Schreibzeitpunkt richtig und zum Commit-Zeitpunkt um eins zu niedrig — relevant nur, weil
  slice-176 gegen genau diese Zahl arbeiten wird.
- `verifizierbar`: **ja** — `git grep -l 'baseline/v6\.2\.0/' <rev> -- '*.md'`.

### I-5 — Zwei vendored Stände sind gedeckt, nicht geduldet

- `kategorie`: INFO
- `quelle`: `harness/conventions.md` §Baseline
- `pfad`: `.harness/baseline/`, `tools/regelwerk-check.sh`:35-45
- `befund`: Die Zusage lautet *„Genau **ein** Stand liegt vendored; mehrere sind nur während einer
  Migration zulässig, und das Target weist den ungeprüften dann namentlich aus."* Beide Hälften
  sind erfüllt: `welle-15` ist eröffnet und offen (Datei liegt flach, Roadmap führt sie unter
  *Offene Wellen*), und `make regelwerk-check` nennt `v6.2.0` namentlich als ungeprüft. Die
  Zusage ist damit **nicht** unwahr geworden. Sie ist allerdings an einen Ausgang gebunden, den
  erst slice-176 setzt; bis dahin trägt sie die Beobachtung aus F-1.
- `verifizierbar`: **ja** — `make regelwerk-check` (Exit 0, Hinweis-Zeile im Klartext).

---

## Negativbefunde (geprüft, ohne Befund)

- **`SHA256SUMS`-Vollständigkeit und -Konvention.** 54 Einträge, 54 Dateien im Baum, Pfadmengen
  identisch; keine unmanifestierte Datei, kein Eintrag ohne Datei. Relative Pfade, zwei Leerzeichen
  als Trenner, Manifest listet sich selbst nicht — gleich `v6.2.0`. `make regelwerk-check` prüft
  beide Richtungen und ist Exit 0.
- **Inhaltliche Identität mit dem Upstream.** Volldiff Tag ↔ vendored: 29 geänderte Zeilenpaare,
  alle mechanische Link-Absolutierung; kein Satz, kein Tabellen-Feld, keine Regel weicht ab. Diff
  ZIP ↔ vendored: leer.
- **Kurs-Wellen-Stempel.** `regelwerk/README.md` des vendorten Stands sagt *„Kurs-Welle 128 ·
  2026-09-06"*; `harness/conventions.md`:24 zitiert genau das. Der von slice-174 §3.3 benannte
  Nachzug 119 → 128 ist vollzogen.
- **Symlinks.** Sieben getrackte Symlinks, alle lösen auf; die vier Baseline-Ziele stehen auf
  `v6.5.0`. `make symlink-check` Exit 0 mit gefeuertem Selbsttest.
- **Nicht-Markdown-Bestand.** `Makefile`, `*.mk`, `tools/*.sh`, `.d-check.yml`,
  `.github/workflows/*.yml`, `.github/dependabot.yml`, `Dockerfile` — kein hartkodierter
  Baseline-Tag. Die drei Skripte, die den Pfad kennen (`regelwerk-check`, `symlink-check`,
  `suppression-check`), lesen den Stand aus `harness/conventions.md` bzw. arbeiten
  versions-agnostisch. Die CI berührt die Baseline gar nicht.
- **Randformen der Umstellung.** Gesucht wurde zusätzlich nach `baseline/<tag>` **ohne**
  Folge-Slash und nach `v6.2.0` in jeder Schreibweise über alle getrackten Dateien. In lebenden
  Dateien bleiben nur Provenienz-Aussagen („seit Kurs-Welle 119 (`v6.2.0`)", „Regelwerk-Migration
  `v6.2.0` → `v6.5.0`", die `Ausgelöst durch Baseline-Stand`-Zeilen der akzeptierten `MR-019`/
  `MR-020`) — sämtlich Aussagen über die Vergangenheit, die korrekt stehen bleiben. Der vom Slice
  selbst gefundene Sonderfall `../baseline/<tag>/` in `.harness/skills/reviewer.md`:7 ist behoben.
- **Vollzähligkeit der 13+4 Umstellungen.** Die in `f49913a` geänderten Zeiger sind genau
  17: vier Symlinks plus 13 Dateien (`AGENTS.md`, `harness/README.md`, `harness/conventions.md`,
  `MR-011`/`MR-012`/`MR-014`/`MR-015`/`MR-016`, `.harness/skills/reviewer.md`,
  `docs/reviews/README.md`, `docs/plan/planning/README.md`, `docs/plan/carveouts/README.md`,
  `open/slice-177-…`). Es bleibt keine lebende Datei mit einem `v6.2.0`-Zeiger übrig.
- **Kopier-Anweisungen in `AGENTS.md` §5 gegen das neu verlinkte Template.** Alle fünf Punkte
  gelten weiter gegen `v6.5.0` · `templates/docs/plan/planning/slice.template.md`: `**Welle:**`
  (Z. 14), Reconciliation-Register (Z. 88), drei Paarungen (Z. 91), Herkunfts-Anker (Z. 157) und
  die Review-DoD-Zeile (Z. 83-85) existieren dort unverändert als Felder. Die Umbenennungen §1
  *Ziel und Abgrenzung* / §8 *Sub-Area-Prüfungen und Modus-Begründung* sind in
  `welle-15` §1 ausdrücklich als Etappe E (slice-178) ausgenommen.
- **`docs/reviews/README.md`** beschreibt Review-Art, Skill-Version und Modell-ID; diese drei
  Kopffelder sind im `v6.5.0`-Template unverändert. Neu ist dort der *Zitier-Form*-Block, den
  Etappe C (slice-176) aufnimmt — die Beschreibung ist damit unvollständig, nicht falsch.
- **Größen-Regel.** Drei Liefer-Punkte, zwei Schichten (`HARNESS`, `GATE`) — die Obergrenze aus
  `AGENTS.md` §5 ist erreicht, nicht überschritten. `make doc-structure` (in `verify`) Exit 0.
- **Lifecycle-Commit `97472ab`.** Reiner Rename: `git diff --name-status -M` meldet `R100` für die
  bewegte Datei; die vier Mitänderungen sind Verweis-Nachzüge und liegen sämtlich unter
  `docs/plan/planning/`. Scope `docs(planning)` korrekt; `make commit-scope-check` Exit 0.
- **ADR-Immutabilität über die Range.** `make doc-immutable RANGE=199a12c..f49913a` Exit 0 — keine
  `Accepted`-ADR berührt.
- **Aggregate.** `make gates` Exit 0 (inkl. `gate-consistency`, `doc-targets`, `doc-planning`,
  `doc-workflows`, `version-coherence`, `dcheck-phrase-selftest`, `symlink-check`), `make verify`
  Exit 0 (21 Anforderungen, 0 Waisen). Nach dem Review ist `git status --porcelain` leer; die
  Mutations-Sonde aus F-3 wurde entfernt und `make doc-check` danach erneut Exit 0.

---

## Kategorie-Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 2 |
| LOW | 3 |
| INFO | 5 |

## Verdikt

**Änderungen nötig vor Closure.** Der eigentliche Gegenstand des Slice — das Vendoring — ist
sauber und über vier unabhängige Kanäle belegbar; daran ist nichts zu beanstanden, und die
Zeiger-Umstellung ist vollständig, auch in den Formen, die ein naives Muster verfehlt.

Beide HIGH-Findings liegen daneben, in derselben Ecke: eine Aussage über den Repo-Zustand, die
dieser Commit aufgehoben hat (F-1), und fünf Zeiger, die ohne die von der eigenen Konvention
verlangte Messung umgezogen sind — einer davon auf einen Abschnitt, der nicht wortgleich ist
(F-2). Beides ist Register- und Konventions-Pflege, kein Fremdtext-Problem, und beides fällt in
den Zuständigkeitsbereich der Closure dieses Slice.

F-4 und F-5 betreffen den Plan, nicht den Diff: der Slice hat mehr geliefert, als er zugesagt hat,
und hält seine Messungen außerhalb des Plans. Für den Verifier bleibt die DoD damit prüfbar, für
den nächsten Leser die Herleitung nicht.
