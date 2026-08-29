# welle-12-regelwerk-migration — Ergebnis-Notiz

**Abschluss:** 2026-08-09. **Erster realer Durchlauf der Fünf-Schritt-Prozedur**
([`planning/README.md`](../README.md)) — sie war seit slice-066 deklariert, aber nie belegt.

**Keine Welle-Plan-Datei zu verschieben.** Schritt 3 sieht ein `git mv` der flachen Welle-Datei
nach `done/` vor; `welle-12` existiert wie die zwölf Wellen vor ihr nur als Prosa-Überschrift in
der Roadmap. Diese Notiz ist damit das **einzige** Welle-Artefakt — ausgewiesen, nicht übergangen.

---

## Geliefert

**Regelwerk-Migration `v1.3.0` → `v3.5.2`** in sechs Etappen, plus der Nachlauf des ersten
unabhängigen Reviews.

| Etappe | Slices |
|---|---|
| A — Vendoring | [slice-047](slice-047-baseline-vendoring.md) |
| B — Modul-Delta lesen | [slice-048](slice-048-modul-delta-lesen.md) (21 Funde) |
| E — Mechanik/Sensoren | [slice-049](slice-049-mechanik-sensoren.md) … [slice-051](slice-051-workflow-und-freigabe.md) |
| D — Form/Templates | [slice-052](slice-052-slice-form.md) … [slice-054](slice-054-ac-form.md) |
| C — `MR-*`-Bereinigung | [slice-055](slice-055-mr-bestand.md), [slice-056](slice-056-sub-area-modus.md) |
| F — Steering/Closure | slice-057, [slice-065](slice-065-carveout-ort-und-trichter.md), [slice-066](slice-066-wellen-closure-und-rollen.md) |
| Fix-Schnitte | [slice-058](slice-058-sensor-praezision.md) … [slice-064](slice-064-guard-verkettung.md), [slice-067](slice-067-roadmap-form.md) |
| Auslösende Analyse | [slice-046](slice-046-regelwerk-v352-migration-analyse.md) |
| **Review-Nachlauf** | [slice-068](slice-068-phony-vollstaendig.md) … [slice-078](slice-078-rollen-uebergaben.md) |

**Der Review-Nachlauf gehört zur Welle**, nicht daneben: ohne ihn wäre die Migration *behauptet*,
nicht belegt. [slice-079](../done/slice-079-gate-consistency-abloesen.md) und
[slice-080](../open/slice-080-verify-abloesung-dcheck.md) gehören **nicht** dazu — sie sind aus der
Arbeit entstanden, haben aber einen anderen Gegenstand (d-check-Ablösung) und eigene Trigger.

## Was funktionierte

**Die Sensoren haben gehalten, wo sie zuständig waren.** Während des Nachlaufs meldeten sie
wiederholt echte Fehler, bevor Schaden entstand:

- `verify-slice-links` fing **dreimal** präfixlose Geschwister-Verweise — jedes Mal *vor* dem
  `git mv`, an dem sie gebrochen wären.
- `doc-check` fing zwei geratene Dateinamen, einen falschen Anker und drei unverlinkte Kennungen.
- Der PreToolUse-Guard blockierte einen Gate-Lauf, der mit einem Commit im selben Aufruf verkettet
  war (`SL-001`).
- Der Stop-Hook forderte zweimal einen `make gates`-Lauf nach, weil Commits den Inhalts-Hash
  verändert hatten.

**Die Zerlegungs-Regel trug.** Der erste Sammel-Entwurf über fünf Findings wurde vom Plan-Review
blockiert (`R-068-F5`); die Zerlegung nach **Fehlermechanismus** statt nach Fundnummer ergab
Slices, die jeweils eine Ursache trafen — und brachte zwei zusätzliche Funde ans Licht, die in der
Sammelform untergegangen wären.

## Was anders lief

**Der Etappen-Schnitt wuchs von vier auf sechs**, Reihenfolge **E vor D** — die vollständige
Baseline-Lektüre in Etappe B brachte elf zusätzliche Funde, und die Mechanik-Funde schufen erst die
Sensoren, an denen die Form-Funde hingen (Drift-Log der [Roadmap](../in-progress/roadmap.md)).

**Die Welle wurde einmal fälschlich geschlossen und wieder geöffnet.** Am 2026-07-26 stand sie im
Closure-Log mit „Closure-Kriterium erfüllt: alle Slices in `done/`" — während
[slice-046](slice-046-regelwerk-v352-migration-analyse.md) in `open/` lag. Der unabhängige Review
fand das als `F-8`; die Closure wurde am 2026-08-09 zurückgezogen (Modul-6-Ausgang **(b)**,
Drift-Log-Eintrag). **Das ist der wichtigste Einzelbefund dieser Welle**: die Prozedur, die hier
zum ersten Mal läuft, wurde beim ersten Anlauf übersprungen — und niemandem fiel es auf, weil das
Überspringen selbst als Entscheidung formuliert war („gilt ab der nächsten Welle").

**Sieben Selbst-Reviews fanden 22 Findings, keines blockierend. Ein unabhängiger Lauf fand 15,
davon 11 HIGH, Verdikt blockierend.** Die Differenz ist das Ergebnis dieser Welle, nicht ihr
Betriebsunfall: Der Roadmap-Trigger verlangte ausdrücklich einen Lauf *„außerhalb dieser
Modell-Familie"*, und genau diese Formulierung hat sich als tragend erwiesen.

## Steering-Loop-Einträge

Gezogen aus `docs/plan/steering-loop.md` — das Register bleibt der
laufende Zähl-Ort; hier stehen nur die in dieser Welle real aufgetretenen Vorfälle.

| Eintrag | in dieser Welle |
|---|---|
| `SL-001` — Gate-Lauf in einer Pipe verschluckt | **einmal**: der Command-Guard blockierte `make gates` + `git commit` im selben Aufruf |
| `SL-002` — relative Verweise brechen beim `git mv` | **dreimal** durch `verify-slice-links` gefangen, dazu zweimal durch `doc-check` nach einem Lifecycle-Wechsel |
| `SL-003` — Commit-Betreff bezeichnet nicht die Arbeit | **zweimal**: ein `git mv` rutschte in einen Report-Commit, ein `git add -A` zog eine fremde Datei mit. Beide aufgeteilt. Dazu der reale CI-Rot-Lauf, aus dem [slice-072](slice-072-scope-sensor-praeventiv.md) entstand |
| `SL-004` — neuer Doku-Sensor meldet sein eigenes Umfeld | **nicht aufgetreten** |

`SL-003` ist der einzige, der in dieser Welle eine **Mechanik** erzeugt hat: der Scope-Sensor
greift seit [slice-072](slice-072-scope-sensor-praeventiv.md) im `commit-msg`-Hook statt erst in
der CI.

## Folge-Slices

- [slice-079](../done/slice-079-gate-consistency-abloesen.md) — Ablösung von `gate-consistency`
  (1)+(2) durch d-checks `targets`-Modul. Trigger: sofort, wartet auf nichts.
- [slice-080](../open/slice-080-verify-abloesung-dcheck.md) — Ablösung der vier `verify-*`.
  Trigger: ein d-check-Release trägt `structure` + `links.resolve-from` **und** der Pin ist
  gehoben. Vorbedingung: CR 1/CR 2 aus
  [slice-073 §8](slice-073-dcheck-statt-eigenbau.md) eingereicht — **noch nicht erfolgt**.

**Offen und ausgewiesen:** `F-9` (Freigabe-Belege) sowie die zwei `--print-mk`-Defekte —
eingebackener Vorgänger-Digest und wörtliches `docker` statt `$(DOCKER)`. Sie sind **nicht**
geschnitten; sie gehören zur Produkt-Achse, nicht zur Regelwerk-Migration.

## Verifikation (Schritt 1 der Prozedur)

| Prüfung | Ergebnis |
|---|---|
| Alle 22 Migrations-Slices in `done/` | ✅ `slice-046` … `slice-067` |
| Alle 11 Nachlauf-Slices in `done/` | ✅ `slice-068` … `slice-078` |
| Jedes belegte False-Green behoben, **je mit Gegenprobe die vorher rot war** | ✅ `F-1`, `F-2`, `F-5`, `F-12`, `F-14`, `R-068-F3`, `R-068-F4` |
| `make ci` (Baseline-Ersetzung für den Replay-Lauf, [`MR-008`](../../../../harness/conventions.md#mr-008--kein-replay-keine-agenten-telemetrie)) | ✅ **Exit 0** — „[ci] gates + image-test grün" |
| Carveout-Audit (Schritt 2) | ✅ Bestand **null** — [`carveouts/README.md`](../../carveouts/README.md); das ist eine Aussage, keine Auslassung |

**Was diese Verifikation nicht belegt:** dass die Migration *inhaltlich* vollständig ist. Sie
belegt, dass jeder Fund des unabhängigen Reviews behoben und jede Behebung durch eine Gegenprobe
gedeckt ist. Ein zweiter unabhängiger Lauf gegen den heutigen Stand ist **nicht** erfolgt und wäre
der nächste ehrliche Schritt, wenn dieselbe Sicherheit für den Nachlauf gelten soll wie für die
Migration.
