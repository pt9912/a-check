# Review-Report: `welle-12` (Regelwerk-Migration, gesamt) — 2026-08-09

**Review-Art:** unabhängiger Code-/Prozess-Review der vollständigen Migration.

**Unabhängigkeit — ausdrücklich:** **unabhängiger Lauf außerhalb der Claude-Modellfamilie.** Er
erfüllt den Roadmap-Trigger „Maintainer startet einen Review-Lauf außerhalb dieser Modell-Familie"
im Wortlaut. Alle sieben vorherigen Reviews dieser Migration sind Selbst-Reviews derselben
Familie; dieser Lauf ist der erste, der das nicht ist.

**Gegenstand:** `182b9ff..6d8bbe7` — 68 Commits, 110 Dateien, +10.943/−195. Die vendored Baseline
unter `.harness/baseline/v3.5.2/` (43 Dateien, +4.822) war per Auftrag ausgenommen.

**Skill:** `.harness/skills/reviewer.md` @ `20ee992` · <!-- d-check:ignore (Adopter-spezifischer Skill-Pfad, existiert im Ziel-Repo ggf. nicht) -->
**Modell:** OpenAI Codex, GPT-5 · **Datum:** 2026-08-09 · **Arbeitsbaum:** unverändert

**Eingangs-Kontext:** [`AGENTS.md`](../../AGENTS.md),
[`harness/README.md`](../../harness/README.md),
[`harness/conventions.md`](../../harness/conventions.md), relevante `AC-*`,
[ADR-0004](../plan/adr/0004-distribution-image-mk.md),
[ADR-0005](../plan/adr/0005-lint-profil.md),
[ADR-0007](../plan/adr/0007-latest-tag-politik.md),
[ADR-0021](../plan/adr/0021-commits-modul-trace-check.md), Regelwerk-Module 02/05/08/10/13/16,
alle bisherigen Migrations-Reviews.

**Formhinweis zur Ablage:** der Wortlaut des Reviewers ist unverändert übernommen. Ergänzt wurden
ausschließlich (a) Links auf die genannten Kennungen, wie `doc-check` sie verlangt, und (b) der
klar abgegrenzte [Anhang](#anhang--gegenprobe-des-maintainer-agenten) mit der Gegenprobe. Der
Anhang ist **nicht** Teil des unabhängigen Laufs.

---

## Findings

### F-1 — Neue Make-Targets sind nicht `.PHONY`

- `kategorie`: **HIGH**
- `quelle`: [`AGENTS.md`](../../AGENTS.md) §4; Harness-Lügen-Schutz
- `pfad`: [`Makefile:35`](../../Makefile), [`Makefile:73`](../../Makefile)
- `befund`: Keines der acht neu eingeführten Targets ist in `.PHONY` aufgeführt. Eine gleichnamige
  Datei lässt unter anderem `make suppression-check` oder `make verify` mit Exit 0 aussteigen,
  ohne das jeweilige Rezept auszuführen.
- `verifizierbar`: ja — isolierte Fixture mit Dateien `suppression-check` und `verify`; beide
  Aufrufe meldeten lediglich `is up to date`, Exit 0.

### F-2 — `suppression-check` prüft nicht alle Go-Quellen

- `kategorie`: **HIGH**
- `quelle`: Hard Rule [`AGENTS.md`](../../AGENTS.md) §3.2;
  [ADR-0005](../plan/adr/0005-lint-profil.md)
- `pfad`: [`tools/suppression-check.sh:33`](../../tools/suppression-check.sh),
  [`tools/suppression-check.sh:71`](../../tools/suppression-check.sh)
- `befund`: Der Sensor scannt ausschließlich `internal/` und `cmd/`, während Target und
  Dokumentation „Go-Quellen des Repos" behaupten. Eine `.go`-Datei im Repo-Root mit wirksamer
  `//nolint`-Direktive wurde mit Exit 0 akzeptiert.
- `verifizierbar`: ja — isolierte Root-Datei `outside.go`; `bash tools/suppression-check.sh` →
  Exit 0.

### F-3 — AC-Sensor prüft Etiketten, nicht die behauptete Given/When/Then-Form

- `kategorie`: **HIGH**
- `quelle`: [`AGENTS.md`](../../AGENTS.md) §5;
  [`harness/conventions.md`](../../harness/conventions.md) §Anforderungs-Anlege-Prozess
- `pfad`: [`AGENTS.md:198`](../../AGENTS.md),
  [`harness/conventions.md:311`](../../harness/conventions.md),
  [`tools/verify-ac-form.sh:47`](../../tools/verify-ac-form.sh)
- `befund`: Der Vertrag verlangt Happy/Boundary/Negative im Given/When/Then-Stil und behauptet
  deren Durchsetzung. Der Sensor prüft nur vier fettgedruckte Bezeichner; ein neues AC mit viermal
  „beliebiger Text" wurde als vollständig akzeptiert.
- `verifizierbar`: ja — Fixture `AC-FA-PROBE-001`; `verify-ac-form` meldete
  `1 neue AC(s) geprueft`, Exit 0.

### F-4 — Lifecycle-Link-Sensor ignoriert Referenz-Links

- `kategorie`: **HIGH**
- `quelle`: [`SL-002`](../plan/steering-loop.md); [`AGENTS.md`](../../AGENTS.md) §4
- `pfad`: [`tools/verify-slice-links.sh:40`](../../tools/verify-slice-links.sh)
- `befund`: Extrahiert werden nur Inline-Links der Form `](...)`. Ein relativer Referenz-Link
  (`[Roadmap][roadmap]` mit `[roadmap]: roadmap.md`), der nach einem Lifecycle-Wechsel bricht,
  wird nicht gesehen.
- `verifizierbar`: ja — isolierter Slice mit Referenz-Link; `verify-slice-links` → Exit 0.

### F-5 — Ungültiger Commit-Range wird falsch-grün

- `kategorie`: **HIGH**
- `quelle`: Commit-Scope-Regel [`AGENTS.md`](../../AGENTS.md) §5; CI-Durchsetzung
- `pfad`: [`tools/commit-scope-check.sh:77`](../../tools/commit-scope-check.sh),
  [`.github/workflows/ci.yml:70`](../../.github/workflows/ci.yml)
- `befund`: Fehler von `git rev-list` werden verworfen. Für einen nicht auflösbaren Range meldet
  der Sensor „0 Commit(s) geprüft" und Exit 0 statt eines Range-Fehlers.
- `verifizierbar`: ja — `make commit-scope-check RANGE=definitely-not-a-revision` → Exit 0.

### F-6 — Zwei-Schichten-Grenze wird als maschinell geprüft ausgegeben, aber nie geprüft

- `kategorie`: **HIGH**
- `quelle`: Fund `B-1` aus [slice-048](../plan/planning/done/slice-048-modul-delta-lesen.md);
  Planning-Harness
- `pfad`: [`docs/plan/planning/README.md:8`](../plan/planning/README.md),
  [`tools/verify-slice-form.sh:22`](../../tools/verify-slice-form.sh)
- `befund`: Die Planungs-Doku ordnet „höchstens zwei Schichten" ausdrücklich `make verify` zu. Das
  Skript erklärt dieselbe Regel dagegen als ungeprüfte Review-Sache; ein Fixture-Slice mit drei
  benannten Schichten lief grün.
- `verifizierbar`: ja — `verify-slice-form` akzeptierte die Drei-Schichten-Fixture mit Exit 0.

### F-7 — `welle-12` überspringt den ersten nach eigener Regel fälligen Closure-Lauf

- `kategorie`: **HIGH**
- `quelle`: Fund `B-13` aus [slice-048](../plan/planning/done/slice-048-modul-delta-lesen.md);
  Wellen-Closure-Prozedur
- `pfad`: [`docs/plan/planning/README.md:57`](../plan/planning/README.md),
  [`docs/plan/planning/in-progress/roadmap.md:75`](../plan/planning/in-progress/roadmap.md)
- `befund`: Die Grandfather-Klausel nennt ausschließlich `welle-00` bis `welle-11`; damit ist
  `welle-12` die nächste Welle und ihr Abschluss die angekündigte erste Probe. Die Roadmap
  schließt sie dennoch ohne Ergebnisnotiz und verschiebt die Prozedur erneut auf „die nächste
  Welle".
- `verifizierbar`: ja — Lesevergleich plus Commitfolge `6283059..f853652`.

### F-8 — Die Migrationsanalyse bleibt `open`, während „alle Slices in `done`" behauptet wird

- `kategorie`: **HIGH**
- `quelle`: Status-/Currency-Disziplin; Roadmap-Closure
- `pfad`: [`slice-046:3`](../plan/planning/open/slice-046-regelwerk-v352-migration-analyse.md),
  [`slice-046:134`](../plan/planning/open/slice-046-regelwerk-v352-migration-analyse.md),
  [`docs/plan/planning/in-progress/roadmap.md:75`](../plan/planning/in-progress/roadmap.md)
- `befund`: Der auslösende Slice der vollständigen Migration liegt weiterhin in `open/`, nennt sich
  „Analyse zur Abnahme" und trägt eine offene Abnahme sowie Closure-Platzhalter. Das Closure-Log
  behauptet gleichzeitig, alle Slices der Migration lägen in `done/`.
- `verifizierbar`: ja — Pfad- und Inhaltsvergleich.

### F-9 — Mehrere Freigabe-Belege belegen ihr Item nicht

- `kategorie`: **HIGH**
- `quelle`: Fund `B-20` aus [slice-048](../plan/planning/done/slice-048-modul-delta-lesen.md);
  Regelwerk Modul 16 „Kein Häkchen ohne Beleg"
- `pfad`: [`docs/user/releasing.md:87`](../user/releasing.md),
  [`.d-check.yml:79`](../../.d-check.yml)
- `befund`: Item 2 verlangt ausschließlich referenzierte `Accepted`/`Superseded`-ADRs, sein Beleg
  `doc-immutable` prüft aber nur Änderungen an bereits `Accepted`-ADRs. In einer isolierten
  Fixture mit referenzierter `Proposed`-ADR lief `make doc-immutable` mit Exit 0; außerdem enthält
  der Beleg-Slot von Item 5 keinen für die behauptete Image-Identität erforderlichen Image-Hash,
  und `make verify` aus Item 3 kennt die Menge „gelieferte Slices" nicht.
- `verifizierbar`: ja — `Proposed`-ADR-Fixture: `make doc-immutable RANGE=HEAD~1..HEAD` → Exit 0.

Der im Auftrag vorgegebene Defekt von Item 6 wurde nicht nochmals als eigenes Finding gezählt.

### F-10 — „Genau ein" WIP-Slice widerspricht dem zulässigen Leerlauf

- `kategorie`: **HIGH**
- `quelle`: Fund `B-6` aus [slice-048](../plan/planning/done/slice-048-modul-delta-lesen.md);
  Regelwerk Modul 5
- `pfad`: [`AGENTS.md:194`](../../AGENTS.md),
  [`docs/plan/planning/in-progress/roadmap.md:13`](../plan/planning/in-progress/roadmap.md)
- `befund`: Die Baseline setzt ein Maximum von einem Slice pro Implementer; das Repo übersetzt dies
  in „genau ein". Der Endstand erklärt zugleich, dass keine Welle läuft, und enthält null Slices
  in `in-progress/`.
- `verifizierbar`: ja — Dateizählung ergibt 0; `make doc-planning` läuft dennoch mit Exit 0.

### F-11 — Zwei Pflichtübergaben fehlen, die Migration heißt trotzdem vollständig

- `kategorie`: **HIGH**
- `quelle`: Fund `B-9` aus [slice-048](../plan/planning/done/slice-048-modul-delta-lesen.md);
  Regelwerk Modul 8
- `pfad`: [`harness/README.md:141`](../../harness/README.md),
  [`slice-066:80`](../plan/planning/done/slice-066-wellen-closure-und-rollen.md)
- `befund`: Verifier → Validator und Validator → Planner werden ausdrücklich als fehlend geführt;
  Modul 8 verlangt jedes der neun Artefakte. Unmittelbar danach erklärt
  [slice-066](../plan/planning/done/slice-066-wellen-closure-und-rollen.md) dennoch alle 21 Funde
  für geschlossen und die Migration für vollständig, ohne eine entsprechende Adaption.
- `verifizierbar`: ja — Lesevergleich mit Modul 8, Zeilen 40–58.

### F-12 — Vier Verify-Sensoren melden eine leere Grundgesamtheit als „ok"

- `kategorie`: **MEDIUM**
- `quelle`: Sensor-Präzision; Leitfrage A2 des Auftrags
- `pfad`: [`tools/verify-closure-notes.sh:116`](../../tools/verify-closure-notes.sh),
  [`tools/verify-slice-form.sh:114`](../../tools/verify-slice-form.sh),
  [`tools/verify-slice-links.sh:106`](../../tools/verify-slice-links.sh),
  [`tools/verify-ac-form.sh:103`](../../tools/verify-ac-form.sh)
- `befund`: In einer leeren Planning-/Lastenheft-Fixture meldeten alle vier Sensoren Exit 0 und
  ausdrücklich „0 … ok". Ein verschwundener Prüfbestand wird damit als erfüllte Verifikation
  behandelt.
- `verifizierbar`: ja — vier isolierte Läufe, jeweils Exit 0.

### F-13 — Closure-Sensor zählt Satzzeichen statt Sätze

- `kategorie`: **MEDIUM**
- `quelle`: Closure-Pflicht; Leitfrage A1 des Auftrags
- `pfad`: [`tools/verify-closure-notes.sh:80`](../../tools/verify-closure-notes.sh)
- `befund`: Die behauptete Mindestzahl von zwei Sätzen wird über die Anzahl von `.`, `!` und `?`
  ermittelt. Die einzelne Zeile „Geprüft via `foo.go`." wurde wegen des Punkts im Dateinamen als
  zwei Sätze und damit als ausgefüllte Closure akzeptiert.
- `verifizierbar`: ja — isolierte `done`-Fixture; Exit 0.

### F-14 — Integritätsprüfung erkennt nicht manifestierte Baseline-Dateien nicht

- `kategorie`: **MEDIUM**
- `quelle`: [`MR-006`](../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert);
  Leitfragen A1/A2 des Auftrags
- `pfad`: [`tools/regelwerk-check.sh:56`](../../tools/regelwerk-check.sh)
- `befund`: `sha256sum -c` bestätigt ausschließlich gelistete Einträge. Eine zusätzliche Datei im
  vendored Baseline-Baum, die nicht in `SHA256SUMS` stand, blieb ungemessen; der Lauf meldete
  weiterhin „Integritaet ok — 42 Dateien", Exit 0.
- `verifizierbar`: ja — isolierte Datei `.harness/baseline/v3.5.2/UNLISTED.md`.

### F-15 — Makefile spricht weiter von drei Verify-Sensoren

- `kategorie`: **LOW**
- `quelle`: Status-/Currency-Genauigkeit
- `pfad`: [`Makefile:94`](../../Makefile), [`Makefile:102`](../../Makefile)
- `befund`: Der Kommentar beschreibt „drei Teil-Sensoren", das Rezept führt vier aus. Die vierte
  Prüfung kam mit `verify-slice-links` hinzu, ohne die Mengenangabe nachzuziehen.
- `verifizierbar`: ja — statischer Zeilenvergleich.

---

## Deckung der 21 Delta-Funde

Bezug: [slice-048](../plan/planning/done/slice-048-modul-delta-lesen.md).

**Ohne Befund geschlossen:** `B-2`, `B-4`, `B-5`, `B-7`, `B-10`, `B-12`, `B-14`, `B-16`, `B-17`,
`B-18`, `B-19`, `B-21`.

**Nicht durch die behauptete Umsetzung gedeckt oder durch obige Findings entwertet:** `B-1`,
`B-3`, `B-6`, `B-8`, `B-9`, `B-11`, `B-13`, `B-15`, `B-20`.

## Negativbefunde

- geprüft, ohne Befund: deklarierte Grandfathering-Grenzen von `verify-slice-form` und
  `verify-ac-form`; beide Schwellen sind sichtbar und wachsen nicht automatisch mit.
- geprüft, ohne Befund: direkte Fehleraggregation im `verify`-Rezept; ohne den `.PHONY`-Bypass
  laufen alle vier Sensoren und ein roter Teillauf setzt den Gesamtexit auf 1.
- geprüft, ohne Befund: `commit-scope-check` erkennt bei einem gültigen Range tatsächlich fremde
  Pfade; der reale Migrationsrange lief mit 12 geprüften und 35 grandfatherten
  `(planning)`-Commits.
- geprüft, ohne Befund: CI-Verdrahtung von `commit-scope-check` sowie Einhängung von
  `suppression-check` in `gates`.
- geprüft, ohne Befund: Freshness von `regelwerk-check` wird ehrlich als ungeprüfte Netzoperation
  ausgegeben; der Status „kein Gate" ist begründet und nicht still.
- geprüft, ohne Befund: Selbsttests der sieben Shell-Sensoren laufen und die Skripte bestehen
  `bash -n`.
- geprüft, ohne Befund: Command-Guard inklusive Drift-Wächter; `make guard-selftest` → Exit 0.
- geprüft, ohne Befund: Modus-Deklarationen, Carveout-Ort und -Trichter, Drift-Log, `next/`,
  Steering-Loop-Kanal und Workflow-Skelett.
- geprüft, ohne Befund: bisherige Review-Reports und deren bereits behandelte Findings; die
  Findings dieses Laufs sind keine Wiederholung der dort dokumentierten 22 Befunde.
- nicht geprüft: Inhalt der vendored Baseline selbst, Produktcode und DoD-Erfüllung — gemäß
  Auftrag außerhalb des Review-Scopes.

## Ausgeführte Läufe

`make verify`, `make gate-consistency`, `make suppression-check`, `make regelwerk-check`,
`make commit-scope-check RANGE=182b9ff..6d8bbe7`, `make guard-selftest`, `make doc-planning` sowie
die beschriebenen isolierten Negativ-Fixtures.

`make gates`/`make ci` wurden **nicht** ausgeführt: Der Review war read-only, und `make gates`
schreibt über `record-gates` Gate-State. Der Arbeitsbaum blieb sauber.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 11 |
| MEDIUM | 3 |
| LOW | 1 |
| INFO | 0 |

## Verdikt

**Merge-/Acceptance-blockierend: ja.**

Die Behauptung „Migration vollständig und belegt" wird vom Repo nicht getragen. Besonders belastend
sind die False-Greens der Sensoren, die übersprungene erste Wellen-Closure und die
Freigabe-Belege, deren jeweiliger Lauf nicht die Aussage des Checklisten-Items misst.

---

## Anhang — Gegenprobe des Maintainer-Agenten

**Nicht Teil des unabhängigen Laufs.** Am 2026-08-09 wurden sieben Findings vor der Annahme
gegengeprüft, um den Report nicht ungeprüft zu übernehmen. **Sieben von sieben bestätigt:**

| Finding | Gegenprobe |
|---|---|
| `F-1` | bestätigt; unabhängig vor Eingang des Reports gefunden. Neun Targets fehlen in `.PHONY`, acht davon Migrations-Neuzugänge |
| `F-2` | Mechanismus bestätigt (`scan ./internal; scan ./cmd`). **Einschränkung:** heute existiert keine `.go`-Datei außerhalb dieser beiden Bäume — latente Lücke, kein aktueller Blindfleck |
| `F-5` | reproduziert: `make commit-scope-check RANGE=definitely-not-a-revision` → Exit 0, Ausgabe „ok: 0 (planning)-Commit(s) … geprueft" |
| `F-8` | bestätigt; Statuszeile im Dokument lautet „open — Analyse zur Abnahme" |
| `F-10` | bestätigt, schärfer als berichtet: [`AGENTS.md:194`](../../AGENTS.md) sagt „Es liegt **genau ein** Slice in `in-progress/`", die Baseline (`modul-05:44`) nur „WIP-Limit pro Implementer = 1". Die Verschärfung von *Maximum* auf *exakt* ist hausgemacht |
| `F-13` | bestätigt; `grep -oE '[.!?]' \| wc -l` liefert für „Geprueft via foo.go." den Wert 2 |
| `F-15` | bestätigt; der Kommentarblock sagt dreimal „drei", das Rezept führt vier Sensoren |

Die verbleibenden acht Findings wurden nicht einzeln nachgeprüft: bei sieben von sieben Treffern
und je einer mitgelieferten Fixture trägt der Report auch ohne vollständige Nachprüfung.

**Folge:** Die `welle-12`-Closure wurde am 2026-08-09 zurückgezogen (Modul-6-Ausgang **(b)**), die
Welle steht wieder unter *Aktuelle Welle* in der
[Roadmap](../plan/planning/in-progress/roadmap.md), und das Drift-Log trägt den Eintrag.
