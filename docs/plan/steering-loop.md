# Steering-Loop — beobachtete Fehlermuster und ihre Antworten

**Zweck.** Der Harness ist nicht statisch. Wiederholt sich ein Versagen, ist die richtige Antwort
nicht „besser aufpassen", sondern ein **Guide oder ein Sensor**. Dieses Register ist der Ort, an
dem Wiederholungen gezählt werden — ohne ihn ist jeder Vorfall ein Einzelfall, und die Schwelle
wird nie erreicht.

**Schwelle (Baseline `grundlagen-klassifikation` §Steering Loop):**

| Vorfälle | Einordnung | Aktion |
|---|---|---|
| 1× | Vorfall | notieren, weiter |
| 2× | Symptom | kategorisieren, beobachten |
| **3×** | **Lücke im Harness** | **Guide oder Sensor nachziehen** |

**Warum hier und nicht in einer Wellen-Closure-Notiz.** Die Baseline verortet Steering-Loop-Einträge
in `done/welle-NN-results.md`. Als dieses Register entstand, hatte a-check **keine** Welle
auditierbar geschlossen (Fund B-13 aus
[slice-048](planning/done/slice-048-modul-delta-lesen.md)); ein Kanal, der auf ein nicht
existierendes Artefakt wartet, sammelt nichts.

**Stand 2026-08-09: zwei Wellen sind geschlossen** ([`welle-12`](planning/done/welle-12-results.md),
[`welle-13`](planning/done/welle-13-results.md)) — und das Register bleibt trotzdem der Ort. Die
Closure-Prozedur **zieht** die Einträge von hier, sie verschiebt sie nicht dorthin
([`planning/README.md` §Wellen-Closure-Prozedur](planning/README.md#wellen-closure-prozedur),
slice-066). Der Grund ist derselbe wie bei der Entstehung, nur andersherum gelesen: ein Zähler, der
erst bei der nächsten Closure entstünde, zählt zwischen zwei Wellen nichts. Der frühere Satz „es
wandert in die Wellen-Closure, sobald es sie gibt" ist damit überholt.

**Angelegt:** slice-057, Fund **B-21**. Die Praxis existierte in slice-001…008 und schlief danach
**lückenlos** ein (40 Slices ohne einen einzigen Eintrag) — der erste Beleg dafür, dass ein Kanal
ohne Ort nicht überlebt.

---

## SL-001 — Gate-Lauf in einer Pipe verschluckt

- **Beobachtung:** `make <gate> | tail` liefert den Exit-Code des letzten Pipe-Glieds, nicht den
  von `make`. Ein rotes Gate verschwindet spurlos; in einem Fall wurde ein Commit an einen
  ungeprüften Lauf gekettet und ging mit rotem `doc-check` heraus.
- **Vorfälle:** **fünf** am 2026-07-25 — vier beim Doku-Schreiben, einer beim Roadmap-Nachzug
  (behoben in `51d5999`).
- **Klasse:** Schwelle überschritten (3×) ⇒ Harness-Lücke.
- **Bisherige Antwort (unzureichend):** eine Agenten-Memory außerhalb des Repos und, seit
  [slice-051](planning/done/slice-051-workflow-und-freigabe.md), Schritt 6 des
  Workflow-Skeletts — beides *inferential feedforward*. Der fünfte Vorfall passierte **nach**
  dem Guide.
- **Antwort (slice-057):** *computational feedforward* — der PreToolUse-Guard lehnt
  `make <gate> | …` und `make <gate> && git commit` fail-closed ab
  (`.claude/hooks/pretooluse-command-guard.sh`, Regel 2). Richtig ist
  `make <gate> > datei 2>&1` mit getrennt geprüftem Exit-Code.
- **Sechster Vorfall (2026-07-26) — und eine Lücke in der Antwort selbst.** `1a9f270` ging mit
  **rotem** `make gates` heraus (Exit 2). Der Lauf war vorschriftsmäßig in eine Datei umgeleitet
  und der Exit-Code ausgegeben — gelesen wurde er nicht, und der Commit folgte im selben
  Tool-Call nach einem `;`. Die Guard-Regel prüfte den Rest auf `git commit` **nur** nach `&&`;
  bei `;` oder Zeilenumbruch fiel sie durch. Die Lücke traf damit genau die Schreibweise, die der
  Guide selbst nahelegt: wer die Umleitung befolgt, verkettet danach mit `;`.
  **Geschlossen in slice-064** — die Prüfung gilt jetzt für jede Verkettung, belegt durch
  Selbsttest-Fixture *und* Live-Probe im echten Tool-Call.
- **Grenze, ehrlich:** die Regel ist quote-bewusst und greift darum nicht, wenn dasselbe Muster
  in einem Sub-Shell-String steht. Der Guard ist ein *Stolperdraht, keine Sandbox*; er fängt die
  versehentliche Drift, nicht die umgeleitete. **Ebenfalls bewusst abgelehnt** wird seit
  slice-064 die geprüfte Verzweigung (`if make gates; then git commit; fi`): von außen ist
  „Exit-Code gelesen" nicht von „ausgegeben und ignoriert" zu unterscheiden, und ein Guard, der
  das zu erraten versucht, wird unzuverlässig. Der Commit gehört in einen eigenen Aufruf.
- **Beleg:** 13 Diskriminierungs-Proben (fünf müssen greifen, sechs dürfen nicht, zwei für die
  unveränderte Regel 1), plus sieben Fälle im `--selftest` an `make gates`.
- **Nebenbeobachtung:** die Regel feuerte **während ihrer eigenen Entwicklung** zweimal auf den
  Autor — einmal berechtigt (Prüfkommando mit echter Pipe), einmal als Fehlalarm auf das Muster
  in einem Argument-String. Der Fehlalarm ist die Ursache der Quote-Behandlung; ohne ihn wäre die
  Regel mit einer Rausch-Quelle in Betrieb gegangen, und ein Sensor, der rauscht, wird
  abgeschaltet statt repariert.

## SL-002 — Relative Verweise brechen beim `git mv` in den nächsten Zustand

- **Beobachtung:** Slice-Dokumente verweisen relativ (`roadmap.md`, `../in-progress/slice-NNN.md`).
  Wandert die Datei per `git mv` von `in-progress/` nach `done/`, zeigen diese Verweise ins Leere —
  jedes Mal gefangen von `doc-check`, jedes Mal einzeln nachgezogen.
- **Vorfälle:** **sieben** am 2026-07-25, über sechs Slices (`f57289d`, `6c69c04`, `71b8844`,
  `d436da9`, `50ddbc0`, `0495a5f`, `51d5999`).
- **Klasse:** Schwelle überschritten (3×) ⇒ Harness-Lücke.
- **Bisherige Antwort:** keine. Der Befund wird zuverlässig gefangen (`doc-check` ist grün-scharf),
  aber **nach** dem Commit — der Zyklus kostet je einen Nachzieh-Commit.
- **Antwort:** **halb gebaut** (Stand slice-059). Zwei Kandidaten:
  1. *Guide:* **geliefert** — Schritt 9 des Workflow-Skeletts
     ([`.claude/commands/slice.md`](../../.claude/commands/slice.md)) verlangt die Verweis-Prüfung
     **vor** dem `git mv` und nennt die zustandsunabhängige Form (`../in-progress/roadmap.md`),
     die aus beiden Verzeichnissen auflöst. Inferential feedforward — es ordnet an, es erzwingt
     nicht.
  2. *Sensor:* **geliefert** (slice-060) — `make verify-slice-links`, eingehängt in `make verify`.
     Er sagt nicht voraus, wohin verschoben wird, sondern prüft eine **Invariante**: ein relativer
     Verweis muss aus **jedem** Lifecycle-Verzeichnis auflösen. Damit fällt auch das *Anlegen* aus
     einer Vorlage anderer Verzeichnistiefe darunter — der Fall, der den achten und neunten Vorfall
     verursachte und den ein reiner `git mv`-Wächter nie gesehen hätte. Belegt durch beide
     Richtungen: ein präfixloser `roadmap.md`-Verweis macht ihn rot, die Form
     `../in-progress/roadmap.md` grün.
- **Damit ist SL-002 in beiden Quadranten durchgesetzt** (`modul-09`): Guide in Schritt 9 des
  Workflow-Skeletts, Sensor in `make verify`. Der Eintrag bleibt stehen — gelöscht wird nichts,
  ein leerer Steering-Loop bedeutet „nie beobachtet", nicht „nichts passiert".
- **Zwei Vorfälle nach dem Guide** (2026-07-26, in slice-058): das Anlegen aus einer Vorlage
  anderer Verzeichnistiefe erzeugte acht gebrochene Verweise, das Schreiben der Closure-Notiz
  über genau diesen Befund zwei weitere — beide Male von `doc-check` **vor** dem Commit gefangen.
  Damit steht der Vorfallszähler bei **neun**, und die Beobachtung ist allgemeiner als der
  Eintragstitel: relative Verweise brechen nicht nur beim `git mv`, sondern bei jedem Wechsel der
  Verzeichnistiefe. Das wiederholt die Lehre aus [SL-001](#sl-001--gate-lauf-in-einer-pipe-verschluckt)
  — ein Guide wirkt nicht gegen Routine.
- **Vorgabe für den Sensor, aus slice-058 gelernt:** der kritische Fall ist die **präfixlose**
  Nachbardatei (`roadmap.md`), nicht der Pfad mit `../`. Der erste Anlauf der manuellen Prüfung
  sah nur `../`-Verweise an und verfehlte damit ausgerechnet den einzigen brechenden.
- **Warum hier stehen bleiben:** die Antwort auf ein Muster gehört nicht in denselben Slice wie
  seine Erfassung, wenn sie ein eigenes Werkzeug braucht. Der Eintrag hält die Zählung fest,
  damit die Lücke nicht ein achtes Mal als Einzelfall durchgeht.

## SL-003 — Commit-Betreff bezeichnet nicht die enthaltene Arbeit

- **Beobachtung:** Ein Commit trägt substanzielle Arbeit unter einem Betreff, der sie nicht nennt —
  typischerweise wandert die Substanz eines Folge-Slice in einen `docs(planning)`- oder
  `fix(planning)`-Commit des Vorgängers. Spiegelbildlich nennen `feat`-Commits Substanz, die sie
  nicht enthalten. `make trace-check` ist dabei **grün**: eine ID ist genannt, sie bezeichnet nur
  nicht die geleistete Arbeit. Damit wird `git log -S` und die Rückverfolgung „welcher Slice hat
  das gebracht" unzuverlässig — dieselbe Begründung, die hinter Hard Rule
  [`AGENTS.md`](../../AGENTS.md) §3.3 steht.
- **Vorfälle:** **drei** am 2026-07-25, alle in der `v3.5.2`-Migrations-Kette:
  - `615e37f` „docs(planning): slice-049 Closure" — **408** Zeilen, liefert die Substanz von
    slice-050 (`tools/verify-closure-notes.sh`, `make verify`, `closure-note-reviewer.md`, das
    slice-050-Plandokument).
  - `f57289d` „fix(planning): slice-050 Roadmap-Link nach dem Verschieben" — **200** Zeilen,
    liefert die Substanz von slice-051 (`.claude/commands/slice.md`, Freigabe-Checkliste in
    `releasing.md`, das slice-051-Plandokument).
  - `d436da9` „docs(planning): slice-052 Folge-Slice-Link auf den Ziel-Pfad" — **115** Zeilen,
    liefert die Substanz von slice-053 (Lifecycle und WIP-Limit in `AGENTS.md`,
    `next/README.md`).

  Dazu zwei spiegelbildliche Fälle: `0f868d7` („feat(harness): Verifikations-Schicht `make verify`
  + Closure-Note-Reviewer") ändert **2** Zeilen und ist ein Link-Fix; `4f9fa5c` („feat(harness):
  Lifecycle vollstaendig, next/ wiederhergestellt, Drift-Log") liefert von drei genannten Dingen
  nur das Drift-Log. Belegt durch einen Abgleich **aller 36** Commits der Kette; die übrigen 32
  sind konsistent.
- **Klasse:** Schwelle überschritten (3×) ⇒ Harness-Lücke.
- **Antwort (Guide, slice-061):** Schritt 10 des Workflow-Skeletts — der Betreff nennt, was im
  Diff steht; wandert Substanz eines anderen Slice mit, ist es ein eigener Commit mit dessen ID.
- **Fünfter Vorfall — der Eintrag selbst (2026-07-26):** `f0e7805` („docs(planning): SL-003 und
  SL-004") ändert `.claude/commands/slice.md` und `docs/plan/steering-loop.md`. Der Commit, der
  diesen Eintrag anlegt, begeht ihn. Dazu `7faa708` („docs(planning): slice-043 in-progress →
  done") mit einem Review-Report im selben Commit — Zähler damit bei **fünf**. Der Guide aus
  slice-061 hat seinen ersten Vorfall nicht verhindert; dieselbe Lage wie bei `SL-001` vor
  slice-057 und `SL-002` vor slice-060, zum dritten Mal belegt.
- **Antwort (Sensor): gebaut (slice-062)** — `make commit-scope-check`, in der CI über die
  Commit-Range neben `trace-check`. Die Regel steht in [`AGENTS.md`](../../AGENTS.md) §5 und gilt
  **nur für den Scope `(planning)`**: dort ist sie über die ganze Historie rauschfrei (fünf
  Treffer bei 74, alle echt), während sie für `docs(...)` allgemein 31 Treffer bei 193 erzeugte —
  `docs(spec)` ändert legitim `spec/`, `docs(adr)` legitim ADRs. Jeder Commit wird an der Fassung
  gemessen, die zu **seinem** Zeitpunkt galt; ältere sind damit grandfathered, ohne dass ein
  Stichtags-Hash nötig wäre.
- **Frühere Einschätzung, korrigiert:** slice-061 hielt den Sensor für „nachweislich rauschfrei",
  geprüft an acht Commits. Über die volle Historie trug das nur scope-spezifisch. Die Stichprobe
  war zu klein für die Behauptung — festgehalten, weil dieselbe Verkürzung den nächsten
  Sensorentwurf treffen kann.
- **Antwort (Sensor), Entwurfsstand vor slice-062:** Die Hypothese *„ein `docs(planning)`/
  `fix(planning)`-Commit ändert nur `docs/plan/planning/`"* wurde gegen die Historie geprüft und
  **diskriminiert sauber**: sie fängt alle drei Vorfälle und lässt fünf geprüfte legitime
  `docs(planning)`-Commits durch, ohne Rauschen. Sie wird **nicht** in slice-061 gebaut, weil die
  zugrunde liegende Regel — welcher Commit-Typ welche Pfade berühren darf — im Repo noch nicht
  deklariert ist; ein Sensor würde sie erfinden statt durchsetzen. Erst die Konvention, dann ihr
  Sensor (dieselbe Reihenfolge wie bei der Closure-Pflicht vor slice-050).
- **Warum das zählt, obwohl nichts kaputt ist:** die Arbeit war jedes Mal vollständig und die
  Gates grün. Beschädigt ist ausschließlich die **Auffindbarkeit** — und die fällt erst auf, wenn
  jemand Monate später fragt, woher eine Regel kommt.

## SL-004 — Ein neuer Doku-Sensor meldet im ersten Lauf sein eigenes Umfeld

- **Beobachtung:** Ein frisch gebauter Sensor über Markdown beanstandet Text, der *über* seinen
  Prüfgegenstand spricht, statt ihn zu sein — zitierte Muster in Inline-Code, Code-Blöcken oder
  Argument-Strings. Die Korrektur ist jedes Mal dieselbe: Zitat-Kontexte vor der Auswertung
  ausblenden.
- **Vorfälle:** **fünf**. Drei beim ersten scharfen Lauf eines *neuen* Sensors — und zwei am
  2026-08-29 an einem **alten**, der den Vorfilter nie bekommen hat:
  - [slice-050](planning/done/slice-050-verify-schicht.md) — `verify-closure-notes`: drei
    Fehlalarm-Wellen, darunter ein Kursiv-Regex, der substanziellen Text als Platzhalter las.
  - [slice-057](planning/done/slice-057-steering-loop.md) — Guard-Regel 2: Fehlalarm auf das
    Pipe-Muster in einem Argument-String; Ursache der Quote-Behandlung.
  - [slice-060](planning/done/slice-060-slice-link-invariante.md) — `verify-slice-links`:
    Fehlalarm auf einen Verweis, den das eigene Slice-Dokument in Backticks **zitiert**.
  - [slice-099](planning/done/slice-099-form-rest-und-fall-des-alten-baums.md) — zweimal
    `verify-closure-notes`: einmal an einer Closure-Notiz, die eine Platzhalter-Wendung
    **unzitiert** in einem Risiko-Satz führte, und unmittelbar darauf an der Notiz, die **diesen
    Vorfall beschrieb** und die Wendung dabei in Backticks zitierte.

**Was die zwei neuen Vorfälle hinzufügen — und warum sie nicht bloß Wiederholung sind.** Die drei
alten trafen je einen *frisch gebauten* Sensor; der Guide aus slice-061 richtet sich an genau
diesen Moment. `verify-closure-notes` ist aus slice-050 und damit **älter als der Guide** — an ihm
zeigt sich, dass ein Guide für Neubauten den Bestand nicht erreicht. Ausgelöst hat es die
Baseline `v5.12.0`: sie verlangt in **jeder** Closure-Notiz einen Abschnitt über offene Risiken,
und damit spricht der Bestand seit dem 2026-08-29 regelmäßig über genau das Vokabular, aus dem die
Platzhalter-Liste besteht. **Eine Konventions-Änderung kann einen jahrelang stillen Sensor scharf
machen.**

**Zweimal falsch behandelt, bevor es richtig behandelt wurde.** In beiden Fällen wurde zuerst der
*Text* angepasst — genau das, wovor die Nebenbeobachtung unten warnt. Erst der dritte Anlauf hat
den *Sensor* angefasst (slice-100):
Placeholder- und Floskel-Prüfung laufen jetzt über denselben Zitat-Vorfilter wie die Satzzählung,
mit Fixtures in **beide** Richtungen. Die verbogene Formulierung wurde danach auf ihre zitierte
Fassung zurückgesetzt — sonst wäre die Beobachtung trotz Sensor-Fix verloren gewesen.

**Offen, ausdrücklich nicht mitgefixt:** die Platzhalter-Liste enthält eine Wendung, die auch
**unzitiert** legitim in einem Risiko-Block steht (der erste der beiden Vorfälle). Sie zu streichen
wäre eine Schwellen-Senkung und braucht nach [`AGENTS.md`](../../AGENTS.md) §3.6 eine ADR. Die
Frage liegt im Beobachtungs-Register; entschieden wird sie am nächsten unzitierten Vorfall.
- **Klasse:** Schwelle überschritten (3×) ⇒ Harness-Lücke.
- **Antwort (Guide, slice-061):** Schritt 5 des Workflow-Skeletts — wer einen Sensor über Markdown
  baut, blendet Zitat-Kontexte **von Anfang an** aus und nimmt eine Fixture mit zitiertem Muster
  in den Selbsttest auf.
- **Kein Sensor, ausdrücklich:** das Muster betrifft **Bauwissen** über Sensoren, nicht einen
  wiederkehrenden Laufzeit-Fehler. Es gibt keinen Lauf, in dem es sich zeigen könnte — der Guide
  ist hier die vollständige Antwort, nicht die halbe. Das unterscheidet den Eintrag von `SL-001`
  und `SL-002`, wo der Guide nachweislich zu schwach war.
- **Nebenbeobachtung:** in allen drei Fällen war der Fehlalarm **nützlich** — er entstand, weil der
  Sensor scharf genug war, und führte je zu einer präziseren Regel. Gefährlich ist nicht der
  Fehlalarm, sondern seine Behandlung: wird der *Text* angepasst statt des *Sensors*, ist die
  Beobachtung verloren.

## SL-005 — Eine neue Datei wird nicht in ihren handgepflegten Index eingetragen

- **Beobachtung:** Eine ADR entsteht, wird verlinkt, referenziert und geprüft — nur die Zeile im
  Index [`docs/plan/adr/README.md`](adr/README.md) fehlt. `make gates` bleibt **grün**, weil
  `doc-check` prüft, dass jeder Link **auflöst**, nicht dass jede Datei **verlinkt ist**. Was nicht
  verlinkt ist, kann auch nicht ins Leere zeigen; der Index sieht darum immer vollständig aus.
- **Vorfälle:** **zwei**, beide am 2026-08-09:
  - [ADR-0030](adr/0030-kein-digest-im-generierten-fragment.md), angelegt in
    [slice-083](planning/done/slice-083-print-mk-digest-selbstbezug.md) — bemerkt in
    [slice-081](planning/done/slice-081-heuristik-diagnose.md), beim Nachtragen der nächsten ADR.
  - [ADR-0031](adr/0031-heuristik-grenzen-diagnose.md), angelegt in
    [slice-081](planning/done/slice-081-heuristik-diagnose.md) — bemerkt in
    [slice-085](planning/done/slice-085-schicht-ohne-aufloesung.md), wieder nur beim Nachtragen der
    nächsten.
- **Klasse:** Symptom (2×). **Der zweite Vorfall wiegt schwerer als der erste**: Er passierte,
  **nachdem** der Fehler diagnostiziert, benannt und als Folge-Slice festgehalten war — beim
  allernächsten ADR. Dieselbe Lehre wie bei [SL-001](#sl-001--gate-lauf-in-einer-pipe-verschluckt)
  und [SL-002](#sl-002--relative-verweise-brechen-beim-git-mv-in-den-nächsten-zustand), diesmal
  ohne dass überhaupt ein Guide dazwischenstand: **Wissen allein verhindert den Fehler nicht.**
- **Dieselbe Asymmetrie an anderer Stelle:** [slice-071](planning/done/slice-071-sensor-scope-vollstaendig.md)
  fand sie bei `regelwerk-check` — geprüft wurde Manifest → Baum, nicht Baum → Manifest; die Lösung
  war ein `comm -13` über beide Mengen. Zwei Vorkommen in **verschiedenen** Sensoren machen die
  fehlende Gegenrichtung zur Klasse, nicht zum Einzelfall.
- **Antwort:** **geschnitten, nicht gebaut** —
  [slice-087](planning/done/slice-087-index-vollstaendigkeit.md), Trigger „sofort". Der Slice trägt
  die Vorarbeit: Bestandsmessung (genau **ein** handgepflegter Datei-Index im Repo) und die
  d-check-Abdeckung (kein Modul deckt die Richtung Ziel → Verweis; `--trace --require-complete`
  findet **Anforderungs**-Waisen, nicht Datei-Waisen — die RTM listete
  [ADR-0031](adr/0031-heuristik-grenzen-diagnose.md) korrekt, **während** sie im Index fehlte).
  Offen bleibt dort der Entscheid lokaler Sensor gegen d-check-CR.
- **Warum das zählt, obwohl nichts kaputt ist:** wie bei
  [SL-003](#sl-003--commit-betreff-bezeichnet-nicht-die-enthaltene-arbeit) ist der Schaden reine
  **Auffindbarkeit** — und ausgerechnet im Register der immutablen Entscheidungen dieses Repos.

## SL-006 — Dateiname oder Anker aus dem Gedächtnis statt aus der Quelle

- **Beobachtung:** Ein Verweis wird geschrieben, wie er heißen *müsste*, statt wie er heißt —
  Slice-Dateinamen aus der Slice-Nummer erraten, Heading-Anker aus der Überschrift konstruiert,
  in einem Fall ein Platzhalter-Anker (`#ac-fa-fa-cli-001-platzhalter`), der beim Schreiben
  entstand und stehen blieb. `doc-check` fängt jeden dieser Fälle zuverlässig; die Kosten sind
  eine Korrekturrunde je Vorfall.
- **Vorfälle:** **mindestens sieben**, über zwei Wellen:
  - `welle-12`: zwei geratene Dateinamen, ein falscher Anker
    ([`welle-12-results.md`](planning/done/welle-12-results.md) §Was funktionierte).
  - `welle-13`: `slice-071-gegenrichtung-und-scope.md` statt
    `slice-071-sensor-scope-vollstaendig.md`; `#mr-001--adrs-schärfen-die-spezifikation-nie-das-lastenheft`
    statt `#mr-001--source-precedence-mit-eigener-spezifikations-schicht`; der Platzhalter-Anker in
    [ADR-0033](adr/0033-forbidden-constructs-fail-closed.md); eine unverlinkte Spezifikations-Kennung
    in [slice-081](planning/done/slice-081-heuristik-diagnose.md).
- **Klasse:** Schwelle überschritten (3×) ⇒ Harness-Lücke.
- **Antwort: der Sensor existiert und wirkt** — `doc-check` ist grün-scharf und hat **jeden** Fall
  gefangen, keiner ging heraus. Was fehlt, ist die Vermeidung: der Verweis wird geschrieben, bevor
  die Quelle angesehen wurde. Ein `ls` über das Zielverzeichnis bzw. ein `grep` nach der Überschrift
  in der Zieldatei kostet einen Aufruf und ersetzt die Korrekturrunde.
- **Dieser Eintrag hat sich beim Schreiben selbst ausgelöst** — zweimal `id-unlinked`, weil er
  Kennungen **zitiert**, um über sie zu sprechen (eine im Fließtext, eine in einem
  Beispiel-Kommando). Das ist die [SL-004](#sl-004--ein-neuer-doku-sensor-meldet-im-ersten-lauf-sein-eigenes-umfeld)-Klasse
  aus der Gegenrichtung: dort meldet ein neuer Sensor sein Umfeld, hier meldet ein bestehender
  Sensor einen Text, der über seinen Prüfgegenstand spricht. Behoben durch Umformulieren — die
  Kennungen sind jetzt benannt statt zitiert. **Der Sensor blieb unangetastet**; das ist die Regel
  aus `SL-004`: wird der Text angepasst statt des Sensors, ist die Beobachtung gerettet, wird der
  Sensor entschärft, ist sie verloren.
- **Kein zweiter Sensor, ausdrücklich:** ein Sensor, der dasselbe wie `doc-check` prüft, nur früher,
  wäre Doppelung. Der Unterschied zu [SL-002](#sl-002--relative-verweise-brechen-beim-git-mv-in-den-nächsten-zustand)
  ist wesentlich: dort brach der Verweis **nach** dem Commit bei einem Verzeichniswechsel, hier ist
  er von Anfang an falsch und wird vor dem Commit gefangen. Der Eintrag zählt, was die Zyklen
  kostet — er verlangt kein Werkzeug.
- **Nebenbeobachtung:** die Fehlerrate ist bei **Ankern** höher als bei Dateinamen, und dort am
  höchsten, wo der Anker aus einer Überschrift mit Bindestrich-Kette entsteht. Die
  Konstruktionsregel ist ableitbar — genau deshalb wird sie geraten statt nachgesehen.

---

## Pflege

- Ein Eintrag entsteht ab dem **zweiten** gleichartigen Vorfall (Symptom), nicht erst ab dem
  dritten — sonst fehlt beim Erreichen der Schwelle die Zählung.
- Ein Eintrag ohne benannte Antwort ist zulässig (siehe SL-002), ein Eintrag ohne **Vorfallszahl**
  nicht: die Zahl ist das Einzige, was die Schwelle prüfbar macht.
- Wird eine Antwort gebaut, bleibt der Eintrag stehen und bekommt den Beleg. Gelöscht wird nichts —
  ein leerer Steering-Loop bedeutet „nie beobachtet", nicht „nichts passiert".
