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
in `done/welle-NN-results.md`. a-check hat bis heute **keine** Welle auditierbar geschlossen
(Fund B-13 aus [slice-048](planning/done/slice-048-modul-delta-lesen.md), offen); ein Kanal, der
auf ein nicht existierendes Artefakt wartet, sammelt nichts. Dieses Register ist darum der
Zwischenschritt — es wandert in die Wellen-Closure, sobald es sie gibt.

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
- **Grenze, ehrlich:** die Regel ist quote-bewusst und greift darum nicht, wenn dasselbe Muster
  in einem Sub-Shell-String steht. Der Guard ist ein *Stolperdraht, keine Sandbox*; er fängt die
  versehentliche Drift, nicht die umgeleitete.
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
- **Antwort (Sensor): offen, aber vorbereitet.** Die Hypothese *„ein `docs(planning)`/
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
- **Vorfälle:** **drei**, je beim ersten scharfen Lauf:
  - [slice-050](planning/done/slice-050-verify-schicht.md) — `verify-closure-notes`: drei
    Fehlalarm-Wellen, darunter ein Kursiv-Regex, der substanziellen Text als Platzhalter las.
  - [slice-057](planning/done/slice-057-steering-loop.md) — Guard-Regel 2: Fehlalarm auf das
    Pipe-Muster in einem Argument-String; Ursache der Quote-Behandlung.
  - [slice-060](planning/done/slice-060-slice-link-invariante.md) — `verify-slice-links`:
    Fehlalarm auf einen Verweis, den das eigene Slice-Dokument in Backticks **zitiert**.
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

---

## Pflege

- Ein Eintrag entsteht ab dem **zweiten** gleichartigen Vorfall (Symptom), nicht erst ab dem
  dritten — sonst fehlt beim Erreichen der Schwelle die Zählung.
- Ein Eintrag ohne benannte Antwort ist zulässig (siehe SL-002), ein Eintrag ohne **Vorfallszahl**
  nicht: die Zahl ist das Einzige, was die Schwelle prüfbar macht.
- Wird eine Antwort gebaut, bleibt der Eintrag stehen und bekommt den Beleg. Gelöscht wird nichts —
  ein leerer Steering-Loop bedeutet „nie beobachtet", nicht „nichts passiert".
