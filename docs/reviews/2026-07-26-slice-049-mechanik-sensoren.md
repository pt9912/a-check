# Review-Report: slice-049 (Etappe E 1/3) — 2026-07-26

**Review-Art:** Code-Review — geprüft gegen den Slice-Plan, `AGENTS.md` §3 Hard Rules / §4 Gates,
[ADR-0005](../plan/adr/0005-lint-profil.md) und
[`MR-006`](../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert).

**Unabhängigkeit — ausdrücklich:** **Selbst-Review**, kein unabhängiger Lauf (neues
Kontextfenster, dieselbe Modell-Familie wie die Autoren-Instanz).

**Gegenstand:** `540f599..894bcf8` (3 Commits: `b319085`, `615e37f`, `894bcf8`) — 673 eingefügte
Zeilen, davon drei neue Skripte unter `tools/`.

**Skill:** `.harness/skills/reviewer.md` @ `20ee992` (2026-07-25) · <!-- d-check:ignore (Adopter-spezifischer Skill-Pfad, existiert im Ziel-Repo ggf. nicht) -->
**Modell:** `claude-opus-5[1m]` · **Datum:** 2026-07-26

**Eingangs-Kontext:**

- [slice-049](../plan/planning/done/slice-049-mechanik-sensoren.md) (Gegenstand),
  [slice-048 §5](../plan/planning/done/slice-048-modul-delta-lesen.md) (Funde B-8, B-11)
- [ADR-0005](../plan/adr/0005-lint-profil.md) `Accepted` — Lint-Profil ohne `//nolint`
- [`AGENTS.md`](../../AGENTS.md) §3.2 (Suppression-Verbot), §3.6 (Gate-Lockerung), §4, §5
- [`harness/conventions.md`](../../harness/conventions.md) `MR-006`
- vendored Baseline: `modul-09` (zwei Quadranten), `modul-13` („vorhanden ≠ behauptet")
- frühere Findings am selben Bereich: [Etappe B](2026-07-26-etappe-b-slice-048.md) (F-1, gleiche
  Klasse wie F-1 hier)

---

## Findings

### F-1 — Wegmarken-Endstand ist gate-rot, Closure behauptet grün

- `kategorie`: **HIGH**
- `quelle`: Klasse „Harness-Lüge" ([Reviewer-Skill](../../.harness/skills/reviewer.md) §HIGH);
  [`AGENTS.md`](../../AGENTS.md) §4 (`doc-check` im `gates`-Aggregat), §6 Schritt 8
- `pfad`: [`docs/plan/planning/done/slice-049-mechanik-sensoren.md:141`](../plan/planning/done/slice-049-mechanik-sensoren.md)
  (Closure-Kriterium 1); Ursache in derselben Datei, Zeilen 11 und 145 @ `894bcf8`
- `befund`: Closure-Kriterium 1 lautet „`make gates` grün mit `suppression-check` im Aggregat
  (Exit 0) — belegt". Auf dem Stand `894bcf8` exitet `make doc-check` mit **2**. Die nach `done/`
  verschobene Slice-Datei verlinkt zwei Ziele verzeichnisrelativ, als läge sie noch in
  `in-progress/`: `roadmap.md` (Zeile 11) und `slice-050-verify-schicht.md` (Zeile 145).
  Spiegelbild zu F-1 der [Etappe B](2026-07-26-etappe-b-slice-048.md) — dort zeigte die Roadmap auf
  die verschobene Datei, hier die verschobene Datei auf ihre Nachbarn.
- `verifizierbar`: **ja — belegt.** `make doc-check` auf Worktree bei `894bcf8`:
  `d-check: 122 Datei(en) geprüft, 2 Befund(e)` · beide `target-missing` · `EXIT=2`.
- `gegenprobe`: Die Wegmarken `71b8844` (E), `51d5999` (D), `a4632e7` (C) und die Spitze `cc7ee9b`
  wurden gleichartig geprüft und sind `doc-check`-**grün** (EXIT=0). Der Defekt ist also kein
  Kettenmuster, sondern trifft genau die beiden ersten Wegmarken.

### F-2 — Der Abschluss-Commit liefert die ausdrücklich ausgegrenzte Substanz des Folge-Slice

- `kategorie`: **MEDIUM**
- `quelle`: Slice-Plan §Nicht in diesem Slice (Selbstabgrenzung); `modul-05` §Ziel-Form Slice;
  [`AGENTS.md`](../../AGENTS.md) §5 (Traceability-ID bezeichnet die geleistete Arbeit)
- `pfad`: Commit `615e37f`; Abgrenzung in
  [`docs/plan/planning/done/slice-049-mechanik-sensoren.md:8-10`](../plan/planning/done/slice-049-mechanik-sensoren.md)
- `befund`: Der Plan grenzt aus: „**Nicht in diesem Slice:** B-3 (`verify` + `check-references`),
  B-4 (`closure-note-reviewer`) … sie folgen als 2/3 und 3/3, damit jeder Schnitt die Größen-Regel
  aus **B-1** einhält". Der Commit `615e37f` — Betreff „docs(planning): slice-049 Closure", einzige
  Traceability-ID `slice-049` — liefert genau diese: `tools/verify-closure-notes.sh` (134 Zeilen),
  die Targets `verify` und `verify-closure-notes` im `Makefile`,
  `.harness/skills/closure-note-reviewer.md` (97 Zeilen), den Closure-Pflicht-Block in
  `AGENTS.md` §5 (dort selbst mit „(slice-050)" markiert) sowie das Plandokument
  `slice-050-verify-schicht.md`. Nach diesem Commit liegen **zwei** Slices in `in-progress/`.
  Damit trifft auch die Plan-Aussage §4 „Zwei Schichten … innerhalb der B-1-Grenze" auf den
  ausgelieferten Umfang nicht mehr zu.
- `verifizierbar`: ja — `git log --diff-filter=A 540f599..894bcf8 -- <pfad>` und
  `git ls-tree 615e37f docs/plan/planning/in-progress/`.
- `einordnung` (Modul 10 verlangt Begründung statt stiller Milde): **nicht** HIGH, weil keine der
  vier HIGH-Klassen des Skills getroffen ist — keine Hard Rule §3.1–§3.6, kein behauptetes Gate
  ohne Target, kein Abwärtsverweis, keine falsche Sachaussage über den Code. Verletzt ist die
  **Zuordnung** von Arbeit zu Slice und Commit, nicht die Sache selbst; die Substanz ist in
  slice-050 vollständig dokumentiert. Das WIP-Limit = 1 war zu diesem Zeitpunkt in `AGENTS.md`
  noch **nicht** deklariert (es kommt erst in Etappe D, `d436da9`) — slice-048 hatte die Lücke
  als B-6 gerade selbst gemeldet. Gegen die seit Etappe A adoptierte Baseline `modul-05` gilt sie
  trotzdem.

### F-3 — Der Selbsttest belegt die Fehlalarm-Freiheit nicht, die er behauptet

- `kategorie`: **MEDIUM**
- `quelle`: `modul-13` (Sensor-Beweispflicht); Slice-Plan §5 („ohne Beweis, dass der Sensor bei
  echtem Verstoß rot wird, ist er ein toter Sensor")
- `pfad`: [`tools/suppression-check.sh:36-41`](../../tools/suppression-check.sh) (Negativ-Fixture
  `neg/c.go`)
- `befund`: Der Selbsttest fordert „eine Fixture mit Direktive MUSS gefunden werden, eine ohne MUSS
  still bleiben". Die Negativ-Fixture lautet `// ein gewoehnlicher Kommentar ueber nolint-Regeln`
  — sie enthält `nolint` **ohne** vorangehendes `//` und kann das Muster
  `//[[:space:]]*(nolint|lint:ignore)` per Konstruktion nicht treffen. Die Fehlalarm-Freiheit ist
  damit gegen einen Fall geprüft, der trivial nicht matcht. Real reproduziert: eine gewöhnliche
  Kommentarzeile, die die Zeichenfolge `//nolint` im Fließtext **erwähnt**, wird als Direktive
  gemeldet und macht `make gates` rot.
- `verifizierbar`: **ja — belegt.** `make suppression-check` gegen eine Probe-Datei meldete beide
  Zeilen: `…:4:// Probe D: genutzte globale Variable OHNE //nolint — gochecknoglobals muss feuern.`
  neben der echten Direktive in Zeile 5.

### F-4 — Baseline-Tag wird lexikografisch statt nach Version gewählt

- `kategorie`: **MEDIUM**
- `quelle`: [`MR-006`](../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
  (Integrität der vendored Baseline); `modul-13` §Vorhanden ≠ behauptet
- `pfad`: [`tools/regelwerk-check.sh:31`](../../tools/regelwerk-check.sh)
- `befund`: Der zu prüfende Stand wird als
  `find … -type d -printf '%f\n' | sort | tail -1` bestimmt — lexikografisch. Liegt neben `v3.5.2`
  ein Verzeichnis `v3.10.0`, wählt die Zeile `v3.5.2`; das Target meldet dann „Integrität ok" für
  den **alten** Stand, ohne dass der Lauf das sichtbar macht. Der Fall tritt genau während einer
  Migration ein, also dann, wenn das Target gebraucht wird.
- `verifizierbar`: **ja — belegt.** `printf 'v3.10.0\nv3.5.2\n' | sort | tail -1` → `v3.5.2`.
  Ein zweites Fixture-Verzeichnis unter `.harness/baseline/` würde es im Lauf zeigen.

### F-5 — Kommentar begründet den Selbstschutz mit dem falschen Mechanismus

- `kategorie`: **LOW**
- `quelle`: Maintainability
- `pfad`: [`tools/suppression-check.sh:22-24`](../../tools/suppression-check.sh)
- `befund`: Der Kommentar lautet „Muster als Zeichenklasse, damit dieses Skript nicht sich selbst
  und nicht die erklärenden Kommentare in `.golangci.yml`/`AGENTS.md` trifft". Die Zeichenklasse
  `[[:space:]]` leistet das nicht; dass die drei genannten Dateien nicht getroffen werden, folgt
  allein aus der Beschränkung des Scans auf `*.go` — die im selben Satz als Nebensatz steht.
- `verifizierbar`: nein — Kommentar-Aussage, kein Gate-Lauf.

## Negativbefunde

- geprüft, ohne Befund: **die zentrale Plan-Abweichung (§1.1, `nolintlint` verworfen)** —
  **unabhängig reproduziert**, nicht geglaubt. Probe D (genutzte globale Variable, keine
  Direktive): `make lint` → **EXIT=2**, `gochecknoglobals` feuert. Probe C (dieselbe Stelle mit
  wohlgeformtem `//nolint:gochecknoglobals // Why: …`): `make lint` → **EXIT=0**. Der erste
  C-Lauf war durch `revive: package-comments` kontaminiert — exakt die Kontaminationsklasse, die
  der Slice bei seinen Proben A/B selbst beschreibt; nach Bereinigung grün. Die Begründung für
  die Abweichung vom Plan trägt.
- geprüft, ohne Befund: **Wirksamkeit des Ersatz-Sensors** — `make suppression-check` fängt genau
  die Direktive, die `make lint` in Probe C durchgelassen hat (EXIT ≠ 0 mit Fundstelle). Die
  Hard-Rule-Lücke aus B-11 ist damit real geschlossen, nicht nur behauptet.
- geprüft, ohne Befund: **„42 Dateien" in `SHA256SUMS`** (§3 des Slice) — exakt 42 Zeilen.
- geprüft, ohne Befund: **Gate-Deklaration** ([`AGENTS.md`](../../AGENTS.md) §4) — `suppression-check`
  und `regelwerk-check` stehen beide in der Target-Tabelle; `suppression-check` hängt im
  `gates`-Aggregat, `regelwerk-check` ausdrücklich **nicht** und ist als „kein Gate" ausgewiesen.
  Die Modul-13-Trennung „vorhanden ≠ behauptet" ist eingehalten.
- geprüft, ohne Befund: **Hard Rule §3.6** — keine Schwelle gesenkt; beide Änderungen fügen
  Beobachtung hinzu. Die Begründung, warum `suppression-check` keine neue ADR braucht (ADR-0005
  trägt die Entscheidung bereits, `Accepted` und immutabel nach §3.5), trägt.
- geprüft, ohne Befund: **Sensoren auf der Merge-Spitze** — `make suppression-check` EXIT=0,
  `make regelwerk-check` EXIT=0 auf `cc7ee9b`.
- geprüft, ohne Befund: **Freshness-Ehrlichkeit** — `regelwerk-check` druckt den adoptierten
  Stand, die Release-URL und den Satz, dass Freshness in diesem Lauf *nicht* geprüft wurde; es
  behauptet die Hälfte nicht, die es nicht leistet.
- geprüft, ohne Befund: **Closure-Notiz-Form** ([`AGENTS.md`](../../AGENTS.md) §5) — genau ein
  Abschnitt, Lerneintrag in einer der drei benannten Formen („geschärfte Regel") mit Ursache,
  zwei beobachtbare Kriterien, konkretes Folge-Slice.
- geprüft, ohne Befund: **Traceability** — alle drei Commits nennen mindestens eine ID; `b319085`
  nennt zusätzlich `ADR-0005` und `MR-006` passend zur Substanz.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 3 |
| LOW | 1 |
| INFO | 0 |

## Verdikt

**Merge-blockierend: nein — mit zwei Auflagen.**

F-1 blockiert den Merge dieser Kette aus demselben Grund nicht wie in Etappe B: gemergt wird die
Spitze, und die ist `doc-check`-grün (belegt). Was bleibt, ist die unhaltbare **Behauptung** im
Closure-Kriterium und der fehlende Sensor dahinter — offener Punkt **SL-002** im
Steering-Loop. Mit diesem Report liegen nun **zwei** rote Gate-Läufe
als Beleg vor statt Erinnerungswerte.

F-3 ist die inhaltlich wichtigste Auflage: der Sensor ist zwar fail-closed (ein Fehlalarm
blockiert, er verliert nichts), aber sein Selbsttest belegt eine Eigenschaft, die er nicht prüft —
dieselbe Klasse, gegen die der Slice mit seiner eigenen Negativ-Probe antritt. Der Lerneintrag des
Slice („erst die Negativ-Probe, dann die Erfolgsmeldung — und die Probe muss den Fall treffen, der
wirklich durchrutschen könnte, nicht den bequemsten") beschreibt seinen eigenen Selbsttest
zutreffend.

F-2 und F-4 sind vor dem Merge zu klären, blockieren ihn aber nicht: F-2 ist Zuordnungs-, keine
Sachfrage; F-4 wird erst bei der nächsten Baseline-Migration wirksam — also planbar vor dem
nächsten `regelwerk-check`-Einsatz.

**Übergabe:** Findings gehen an die Implementation. Der Report ersetzt keine Verifikation —
DoD-/Spec-Konformität prüft `make verify` separat (Modul 11).
