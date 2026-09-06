# welle-14-regelwerk-v610-migration — Ergebnis-Notiz

**Abschluss:** 2026-09-06. Retargeted während der Laufzeit (`v6.1.0` → `v6.2.0`, s.
[welle-14-regelwerk-v610-migration.md](welle-14-regelwerk-v610-migration.md) Retarget-Hinweis) —
Dateiname trägt weiterhin die ursprüngliche `v610`-Kennung (stabile ID, kein Nachzug, dieselbe
Praxis wie bei Slice-Dateinamen).

---

## Geliefert

**a-check auf den Kurs-Stand `v6.2.0` gehoben — sieben Slices, alle in `done/`.**

| Slice | Gegenstand | Ergebnis |
|---|---|---|
| [slice-161](slice-161-regelwerk-v610-delta-analyse.md) | Delta-Analyse `v6.0.0`→`v6.1.0` | 6 Dateien, `+38/−1`; ein echter Nachzug identifiziert |
| [slice-162](slice-162-review-pflicht-absatz-v610-wortlaut.md) | Review-Pflicht-Absatz auf `v6.1.0`-Wortlaut zurückgeschnitten | [MR-018](../../../../harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md) |
| [slice-163](slice-163-adaptions-durchgang-v610.md) | Etappe B: alle 18 `MR`-Dateien gegen `v6.1.0` geprüft | [MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md)s Auflösungs-Trigger erkannt, Folge-Slice vorgeschlagen |
| [slice-164](slice-164-regelwerk-v620-delta-analyse.md) | Delta-Analyse `v6.1.0`→`v6.2.0` (Increment) | 4 Dateien, `+16/−5`; Review-Report-Konvergenz gefunden |
| [slice-165](slice-165-review-dod-punkt-opt-in-beibehalten.md) | Review-Checkbox-Punkt: Opt-in beibehalten | [MR-019](../../../../harness/conventions/MR-019-review-dod-opt-in.md); empirisch: eigene DoD-Phrase löste `make doc-reviews` nie aus |
| [slice-166](slice-166-mr020-adr-vorlage-generisch.md) | [MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md) durch generischen Verweis abgelöst | [MR-020](../../../../harness/conventions/MR-020-adr-vorlage-generisch.md); dritter Versions-Bump-Durchlauf vermieden |
| [slice-167](slice-167-etappe-a-vendoring-v620.md) | Etappe A: `.harness/baseline/v6.2.0/` vendored | Stand-Deklaration an allen strukturellen Pointern gehoben |

**Das Welle-Ziel war:** a-check auf den aktuellen Kurs-Stand heben, ohne einen blinden
Pointer-Bump zu fahren, der einen Agenten einen Pfad kopieren lässt, den es hier noch nicht gibt
(dieselbe Sorge wie bei `slice-136`, dem `v6.0.0`-Vorbild). Eingelöst: der einzige echte
Feature-Sprung im gesamten `v6.0.0`→`v6.2.0`-Bereich (Review-Report als konstanter/verpflichtender
DoD-Punkt) wurde **vor** dem Pointer-Bump entschieden ([MR-019](../../../../harness/conventions/MR-019-review-dod-opt-in.md)), nicht danach vertagt.

## Was funktionierte

**Empirisches Testen gegen den gepinnten Sensor-Digest statt Dokumentation zu glauben.** `slice-165`
baute ein isoliertes Scratch-Repo und testete fünf DoD-Formulierungen gegen den echten
`d-check`-Digest — und fand dabei, dass a-checks eigene, in vier Slices tatsächlich verwendete
Formulierung („Unabhängiges Plan-Review …") die Trigger-Phrase von `make doc-reviews` **nie**
ausgelöst hatte. Vier Slices hatten reale, aber mechanisch ungeprüfte Reviews. Reine
Dokumentenlektüre hätte das nicht gefunden — `AGENTS.md`s eigene Beschreibung der Phrase war
korrekt, nur die Praxis wich ab.

**Die Zwei-Klassen-Unterscheidung beim Pointer-Bump** (`slice-167` §3: „strukturell" vs.
„`Ersetzt-Baseline-Regel`-Anker akzeptierter `MR`-Dateien") vermied die Sorge der letzten Migration
(sechs bewusst zurückgehaltene Zeiger, `slice-136`) — diesmal ohne einen einzigen Zeiger vertagen zu
müssen, weil die eine echte Lücke bereits vorab entschieden war.

**Retargeting statt Doppelarbeit:** als `v6.2.0` zwei Stunden nach `v6.1.0` erschien, noch bevor
Etappe A vendorte, wurde die laufende Welle umgehängt statt `v6.1.0` zu vendoren und Stunden später
erneut migrieren zu müssen (`slice-164` §7, geschärfte Regel).

**Jeder Slice bekam ein echtes, getrenntes Review** — anders als `welle-12`/`welle-13`, die beide
ausdrücklich nur Selbst-Reviews trugen. Sechs unabhängige Reviews liefen in dieser Welle, jedes über
einen frischen Subagenten-Kontext, jedes mit mindestens einem echten Finding.

## Was anders lief

**Reviews fanden in fünf von sechs Läufen mindestens einen echten Fehler** — kein Lauf war reine
Formsache:

- `slice-163`: Zählfehler im Beobachtungs-Register (7 statt 8 offene Einträge) — Ursache eine
  vorbestehende Sub-Area-Feld-Inkonsistenz, bereits in einem früheren Review benannt, aber nicht
  in die eigene Zählung übernommen.
- `slice-165`: ein `AGENTS.md`-§6-Zitat aus dem Gesprächskontext, das zum Zeitpunkt des Zitats
  bereits durch `slice-162` (selber Tag) überholt war — dieselbe Fehlerklasse wie
  [`BEO-PLAN/slice-provenienz-aus-gedaechtnis-statt-git-log`](../observations/BEO-PLAN/slice-provenienz-aus-gedaechtnis-statt-git-log/observation.md),
  jetzt als eigener, verwandter Eintrag registriert
  ([`dateiinhalt-aus-gedaechtnis-zitiert`](../observations/BEO-PLAN/dateiinhalt-aus-gedaechtnis-zitiert/observation.md)).
- `slice-166`: eine repo-eigene Zählung („zwölf Dateien") war falsch — tatsächlich zehn.
- `slice-167`: ein struktureller Pointer (`harness/conventions.md` §Adoptierte Konventions-Quellen)
  wurde beim Bump übersehen — ein Selbstwiderspruch im selben Dokument, den kein Gate fängt, weil
  der alte Pfad technisch weiter auflöste (`v6.0.0` blieb ja vendored).
- Zusätzlich, vom Maintainer bemerkt statt vom Agenten: die vier `.claude/rules/`-Symlinks blieben
  beim Vendoring auf `v6.0.0` zeigen — zweite Evidenz für ein bereits registriertes Muster
  ([`BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft`](../observations/BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft/observation.md),
  zuvor 1× bei `slice-142`, jetzt 2×).

**Zwei Review-Reports fehlten zunächst ganz.** `slice-163` und `slice-164` bekamen reale,
unabhängige Reviews über getrennte Subagenten, aber die Ergebnisse wurden nicht als Datei unter
`docs/reviews/` persistiert — entdeckt erst durch `slice-165`s eigene Analyse (die genau diese
Lücke als generelles Muster diagnostizierte). Beide Reports nachträglich aus den bereits
durchgeführten Reviews rekonstruiert.

## Steering-Loop-Einträge

- **`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`** erreichte **3×** (`slice-120`, `slice-123`,
  `slice-165`) — Ausgang: **geplant** in
  [`slice-168`](wellenlos/slice-168-pruefer-kalibrierungs-selbsttest.md) (Kalibrierungs-Selbsttest für
  phrasen-/muster-basierte Prüfer; noch keine automatisierte Lösung, die Analyse-Fragen sind offen).
- **`BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration`** — **verkörpert** in
  [`MR-020`](../../../../harness/conventions/MR-020-adr-vorlage-generisch.md) `seit slice-166`: die
  saubere Auflösung (generischer Baseline-Verweis statt einer dritten Versions-`MR`) wurde gefahren,
  bevor das Muster ein drittes Mal auftrat.
- **`BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft`** — weiterhin **offen**, jetzt 2× (nicht
  3×): `slice-142` und `slice-167`.
- **`BEO-PLAN/dateiinhalt-aus-gedaechtnis-zitiert`** — neu angelegt (`slice-165`), 1×, weiterhin
  offen.

## Folge-Slices

- [slice-168](wellenlos/slice-168-pruefer-kalibrierungs-selbsttest.md) — Kalibrierungs-Selbsttest für
  phrasen-/muster-basierte Prüfer (`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`, 3×). Trigger:
  Priorisierung durch den Maintainer, kein akuter Zwang.

**Benannt, aber ausdrücklich nicht geschnitten:**

- **[MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md)/[MR-018](../../../../harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md)/[MR-019](../../../../harness/conventions/MR-019-review-dod-opt-in.md)/[MR-020](../../../../harness/conventions/MR-020-adr-vorlage-generisch.md)-Klasse
  „Ersetzt-Baseline-Regel: keine"** (Repo-Aussage-Korrektur statt
  Baseline-Regel-Ersatz) bleibt ein Rückbau-Kandidat nach dem Fork-Test —
  [MR-020](../../../../harness/conventions/MR-020-adr-vorlage-generisch.md) ist die einzige
  davon mit **permanentem** Trigger, die anderen drei lösen sich erst mit der jeweils nächsten
  Baseline-Migration ([MR-018](../../../../harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md)) oder gar nicht bis zur ID-Schema-Überarbeitung. Kein eigener
  Slice, solange keine dieser Adaptionen selbst wieder fällig wird.
- **`.harness/baseline/v6.0.0/` löschen**, sobald keine `Ersetzt-Baseline-Regel`-Anker mehr darauf
  zeigen (alle sechs betroffenen `MR`-Dateien müssten dafür selbst abgelöst werden — kein
  absehbarer Trigger, da sie inhaltlich unveränderlich sind, solange sie `Accepted` bleiben).

## Verifikation

| Prüfung | Ergebnis |
|---|---|
| Alle sieben Slices der Welle in `done/` | ✅ `slice-161` … `slice-167` |
| Stand-Deklaration an den drei kanonischen Stellen | ✅ `harness/conventions.md` §Baseline, `AGENTS.md` §1, `harness/README.md` §Guides — alle `v6.2.0` |
| `make gates` auf dem finalen Stand | ✅ Exit 0 |
| `make verify` auf dem finalen Stand | ✅ Exit 0 |
| `make ci` (Replay-Ersatz, [`MR-015`](../../../../harness/conventions/MR-015-welle-closure-ohne-replay.md)) | ✅ Exit 0 |
| Carveout-Audit | ✅ Bestand **null** — [`carveouts/README.md`](../../carveouts/README.md) |
| Bootstrap-aware-Gate-Audit | ✅ keine Berührung dieser Welle |
| ADR-Re-Evaluierungs-Audit | ✅ keine ADR berührt |
| Unabhängiges Review je Slice | ✅ sechs von sieben (`slice-161` hatte bereits eines aus dem Vorlauf; `slice-162` ebenso) — anders als `welle-12`/`welle-13`, die beide nur Selbst-Reviews trugen |

**Was diese Verifikation nicht belegt:** ob die Entscheidung, `v6.0.0` dauerhaft neben `v6.2.0`
vendored zu lassen, bei einer dritten Baseline-Migration noch trägt — das ist als Risiko in
`slice-167` §6 registriert, nicht hier entschieden.
