# slice-167 — Etappe A: Baseline `v6.2.0` committet vendoren

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-14](welle-14-regelwerk-v610-migration.md).

**Bezug:** [slice-161](../done/slice-161-regelwerk-v610-delta-analyse.md)
§6 (Etappe-A-Vorschlag, Ziel-Version durch Retarget auf `v6.2.0` gehoben),
Maintainer-Wort 2026-09-06 ("Etappe A").

**Berührte Spec-Stellen:** — *(keine)* — Harness-/Konventions-Änderung
ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim
Maintainer.

**Autor:** Claude (Sonnet 5). **Datum:** 2026-09-06.

---

## 1. Ziel

`.harness/baseline/v6.2.0/` committet vendoren, die Stand-Deklaration an
den strukturellen Pointern auf `v6.2.0` heben und den in `slice-161` §4.1
benannten Nachzug (Reviewer-Skill-Zeile in `harness/README.md` §Guides)
umsetzen — der letzte Schritt vor der Closure von `welle-14`.

## 2. Definition of Done

- [x] `.harness/baseline/v6.2.0/{regelwerk,templates}/` (53 Dateien) neben
      `v6.0.0` vendored, eigenes `SHA256SUMS` generiert, Provenienz
      unabhängig belegt (ZIP-Digest lokal **und** via GitHub-API
      gegengerechnet, Datei-Baum-Identität zu `v6.0.0` bis auf die neun
      bereits vermessenen Zeilen bestätigt). `make regelwerk-check` grün.
- [x] Stand-Deklaration an den **strukturellen** Pointern (nicht an den
      inhaltlich unveränderlichen `Ersetzt-Baseline-Regel`-Ankern
      akzeptierter `MR`-Dateien, s. §3) auf `v6.2.0` gehoben — inklusive
      der vier `.claude/rules/`-Symlinks (zunächst übersehen, vom
      Maintainer bemerkt, s. §3).
- [x] `harness/README.md` §Guides: Zeile für `.harness/skills/reviewer.md`
      ergänzt (slice-161 §4.1, bisher unterlassen).
- [x] Unabhängiger Review über getrennten Kontext durchgeführt (Report
      unter `docs/reviews/`).
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

## 3. Umsetzung

**Vendoring (Liefer-Punkt 1).** `gh release download v6.2.0 --repo
pt9912/ai-harness-course` lädt `lab-regelwerk.zip` + `SHA256SUMS`;
`sha256sum -c` bestätigt lokal, `gh api
repos/pt9912/ai-harness-course/releases/tags/v6.2.0 --jq
'.assets[].digest'` bestätigt denselben Digest **unabhängig** über die
GitHub-API (zweiter Kanal, nicht dieselbe Quelle wie der Download).
Entpackt nach `.harness/baseline/v6.2.0/`, eigenes `SHA256SUMS` per
`sha256sum` über alle 53 Dateien generiert (`regelwerk-check.sh` prüft
dagegen). Datei-**Pfad**-Baum ist mit `v6.0.0` identisch (`diff <(find
v6.0.0 -type f | sort) <(find v6.2.0 -type f | sort)` leer). **Präzisiert
durch Review:** ein naiver `diff -rq` gegen `v6.0.0` zeigt **31**
inhaltlich verschiedene Dateien, nicht neun — jede `regelwerk/*.md`-Datei
trägt eine `<!-- Quelle: …/blob/<tag>/… -->`-Kommentarzeile, die bei
jedem Versionssprung zwangsläufig mitläuft, unabhängig von echten
Regeländerungen. Nach Herausfiltern dieses mechanischen Tag-Stempels
bleiben exakt die neun bereits in `slice-161`/`slice-164` vermessenen
Dateien mit substanziellen Änderungen übrig — die Zahl bezog sich von
Anfang an auf den *upstream*-`git diff --stat` (slice-164 §2), nicht auf
einen naiven lokalen Verzeichnis-Diff.

**Stand-Deklaration (Liefer-Punkt 2) — was bumpt, was nicht.** Zwei
Klassen von `.harness/baseline/v6.0.0/…`-Referenzen im Repo:

| Klasse | Beispiele | Verhalten |
|---|---|---|
| **Strukturell** — verweist auf „den aktuell vendorten Stand", nicht auf einen historischen Zeitpunkt | `AGENTS.md` §1 (Index/Templates/`SHA256SUMS`/Kurs-Tag-Link), `AGENTS.md` §5 (`slice.template.md`-Pointer), `harness/conventions.md` §Baseline „Stand:", `harness/README.md` (Sensors-Kommentar), `docs/reviews/README.md`, `.harness/skills/reviewer.md`, `docs/plan/carveouts/README.md` | **gehoben auf `v6.2.0`** |
| **`Ersetzt-Baseline-Regel`-Anker akzeptierter `MR`-Dateien** — dokumentiert, welcher Baseline-Abschnitt zum Zeitpunkt **dieser** Adaption galt | [MR-011](../../../../harness/conventions.md#mr-011)/[MR-012](../../../../harness/conventions.md#mr-012)/[MR-014](../../../../harness/conventions.md#mr-014)/[MR-015](../../../../harness/conventions.md#mr-015)/[MR-016](../../../../harness/conventions.md#mr-016)/[MR-018](../../../../harness/conventions.md#mr-018) (aktiv), [MR-017](../../../../harness/conventions.md#mr-017) (aufgelöst, `conventions/done/`) | **unverändert** — Adaptions-Block-Disziplin: „an einem akzeptierten Eintrag wird nichts nachträglich inhaltlich geändert" |

**Dritte Klasse, zunächst übersehen:** die vier `.claude/rules/`-Symlinks
auf einzelne Regelwerk-Module (`modul-01/05/06/08`) sind ebenfalls
strukturell, standen aber nicht in der ursprünglichen Liefer-Punkt-2-Liste
— derselbe Fehler, den
[`BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft`](../observations/BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft/observation.md)
bereits bei `slice-142` registriert hatte (dort: „kein Schritt der
Migration prüfte sie, und kein Gate hätte den Bruch gefangen"). Vom
Maintainer bemerkt, nicht vom Agenten — zweite Evidenz für denselben
Eintrag (jetzt 2×, noch nicht 3×). Nachträglich auf `v6.2.0` korrigiert.

Die zweite Klasse hält `.harness/baseline/v6.0.0/` weiterhin für
netzlose Auflösung nötig — deshalb bleibt `v6.0.0` **vendored liegen**,
statt gelöscht zu werden. Das ist keine Ausnahme, sondern von
`harness/conventions.md` §Baseline und `regelwerk-check.sh` selbst
vorgesehen: mehrere vendorte Stände sind zulässig, das Werkzeug meldet
das als Hinweis, nicht als Fehler, und prüft den höchsten. Der
Unterschied zur vorigen Migration (`slice-136`, sechs bewusst
zurückgehaltene Zeiger wegen **ungelöster** Feature-Lücken): hier ist die
einzige inhaltliche Lücke (Review-DoD-Zeile, `slice.template.md`) bereits
durch [MR-019](../../../../harness/conventions.md#mr-019)/`slice-165`
**entschieden und verankert** (fünfter
Kopier-Hinweis in `AGENTS.md` §5) — die Nachbildung ist damit *aktiv*,
nicht *vertagt*, und der Pointer kann trotzdem auf `v6.2.0` zeigen.

**Reviewer-Skill-Zeile (Liefer-Punkt 3).** `harness/README.md` §Guides
fehlte bisher eine Zeile für
[`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md),
obwohl `AGENTS.md` §6 mehrfach darauf verweist (slice-161 §4.1 Fund,
seither nicht umgesetzt). Ergänzt nach dem Wortlaut aus
`lab/templates/harness/README.template.md`s `v6.1.0`-Zuwachs (gemessen in
slice-161 §4.1, seither unverändert bis `v6.2.0`).

## 4. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe (dieses Gespräch,
2026-09-06), WIP-Limit frei (`in-progress/` leer nach `slice-166`-Closure),
keine offenen Folge-Slices mehr (`welle-14` §„Stand").

**Rückführungen:**

- `in-progress` → `next`: entfällt — Umfang ist bereits auf drei
  Liefer-Punkte geschnitten und passt.
- `in-progress` → `open`: falls die ZIP-Provenienz nicht unabhängig
  bestätigt werden kann (Netz-Ausfall, Digest-Abweichung).

## 5. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz
geschrieben. **Nicht** Closure-Trigger für `welle-14` selbst — die Welle
braucht zusätzlich `done/welle-14-results.md` (Modul 8 §Rollen-Sequenz für
eine Welle, Träger Planner).

## 6. Risiken und offene Punkte

- *Zwei vendorte Stände (`v6.0.0` + `v6.2.0`) bleiben dauerhaft liegen,
  weil kein Folge-Slice die `MR`-Anker jemals auf `v6.2.0` migriert (das
  dürfte er wegen der Immutabilität ohnehin nicht) — `v6.0.0` wird zu
  totem Gewicht* — **Ausgang:** weiter offen → Beobachtungs-Register (kein
  bestehender Eintrag trifft genau diesen Fall — neu wäre er erst bei
  einer dritten Baseline-Migration mit demselben Muster anzulegen, nicht
  vorsorglich beim ersten Auftreten).
- *Die Klassifikation „strukturell vs. `Ersetzt-Baseline-Regel`-Anker" in
  §3 ist eine neue Unterscheidung, die kein bestehendes Werkzeug prüft —
  ein künftiger Bump könnte versehentlich einen `MR`-Anker mitziehen* —
  **Ausgang:** gestrichen mit Begründung: `doc-immutable` deckt
  `docs/plan/adr/` ab, nicht `harness/conventions/MR-*.md` — die
  Immutabilitäts-Prüfung für `MR`-Dateien ist rein redaktionell (Review),
  dieselbe Grenze wie bei ADRs vor `doc-immutable`s Einführung; kein neues
  Risiko, sondern eine bestehende, benannte Grenze (`AGENTS.md` §3.7
  Durchsetzung: keine).

## 7. Closure-Notiz

- **Was hat funktioniert:** die Zwei-Klassen-Unterscheidung (strukturell
  vs. `Ersetzt-Baseline-Regel`-Anker, §3) hat die vorige Migrations-Sorge
  („blinder Bump lässt Agenten einen nicht-existenten Pfad kopieren",
  `slice-136`) diesmal vermieden, ohne sechs Zeiger vertagen zu müssen —
  weil die einzige echte Lücke (Review-DoD-Zeile) bereits vorab entschieden
  war ([MR-019](../../../../harness/conventions.md#mr-019)).
- **Was ging anders als geplant:** `slice-161` §6 hatte die Etappe für
  `v6.1.0` vorgeschlagen; durch den Retarget (`welle-14`) lief sie
  stattdessen gegen `v6.2.0` — inhaltlich identisches Verfahren, nur das
  Ziel verschoben. Zwei Korrekturen aus dem unabhängigen Review vor
  Abschluss eingearbeitet: `harness/conventions.md` §Adoptierte
  Konventions-Quellen zeigte trotz der eigenen §Baseline-Angabe direkt
  darüber weiterhin auf `v6.0.0` (Selbstwiderspruch im selben Dokument,
  von keinem Gate gefangen, da der `v6.0.0`-Pfad technisch weiter
  auflöst) — auf `v6.2.0` korrigiert; und die „neun Dateien"-Zahl in §3
  bezog sich unausgesprochen auf den *upstream*-Diff, während ein naiver
  lokaler `diff -rq` wegen der Tag-Stempel-Kommentarzeile in jeder
  Regelwerk-Datei 31 Treffer zeigt — Formulierung präzisiert. Zusätzlich,
  vom Maintainer bemerkt: die vier `.claude/rules/`-Symlinks (s. §3,
  bereits vor diesem Review korrigiert).
- **Lerneintrag — Form: geschärfte Regel.** *Beim Baseline-Vendoring wird
  vor dem Pointer-Bump geprüft, ob jede vom Sprung betroffene
  Template-Zeile bereits eine a-check-eigene Entscheidung (MR oder
  äquivalent) trägt — nur dann darf der strukturelle Pointer sofort auf
  den neuen Stand zeigen. Fehlt die Entscheidung, bleibt der Pointer wie
  bei `slice-136` zurück, bis sie getroffen ist. Die Unterscheidung
  „strukturell" vs. „`Ersetzt-Baseline-Regel`-Anker" (§3) ist dabei
  orthogonal: Anker akzeptierter `MR`-Dateien bumpen nie, unabhängig vom
  Entscheidungsstand.*
- **Beobachtungs-Register (`../observations/`):** kein neuer Eintrag —
  das einzige Risiko mit Register-Ausgang (§6) trifft keinen bestehenden
  Fall und wird erst bei tatsächlichem Wiederauftreten angelegt.
- **Folge-Slices:** keine.
- **Risiken aus §6:** beide mit Ausgang — siehe §6.
- **Drei Paarungen:** verschoben auf die Closure von `welle-14` (dieser
  Slice trägt ein `**Welle:**`-Feld, [Modul 8](../../../../.harness/baseline/v6.0.0/regelwerk/modul-08-agentenrollen.md#rollen-sequenz-für-eine-welle)).

## 8. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** zwei Sub-Areas berührt —
**Vendored Baseline** (`.harness/baseline/`, kein Modus, externer
Fremdtext, [`MR-006`](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert))
und **Harness-Einstieg** (`AGENTS.md`, `harness/README.md`,
`docs/reviews/README.md`, `.harness/skills/reviewer.md`,
`docs/plan/carveouts/README.md` — strukturelle Pointer-Änderungen),
Greenfield, Schwelle ≥ 2/3 erfüllt.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-HARNESS/` über die
Verzeichnisliste geprüft — 8 `offen` vor diesem Slice (unverändert seit
`slice-166`s Auflösung von `rueckbau-kandidat-ueberlebt-baseline-migration`),
keiner erreicht 3×.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig (Vendored
Baseline führt ohnehin keinen Modus).
