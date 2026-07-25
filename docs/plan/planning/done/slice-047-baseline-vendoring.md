# slice-047 — Etappe A: Baseline committet vendoren (`v1.3.0` → `v3.5.2`)

**Status:** in-progress — Etappe **A** aus
[slice-046 §6](../open/slice-046-regelwerk-v352-migration-analyse.md), am 2026-07-25 per
Maintainer-Wort gezogen („A zuerst"). Ergebnis §3, [Erst-Report](../../../reviews/2026-07-25-slice-047-baseline-vendoring.md) + [Zweit-Review](../../../reviews/2026-07-25-slice-047-baseline-vendoring-zweitreview.md), Closure §5.
**Auslöser:** Maintainer-Vorgabe: **vollständige Migration nach `v3.5.2`**, auch gegen bestehende
`MR-*`; Etappen-Schnitt aus der Delta-Analyse.
**Bezug:** Baseline-Deklaration
[`harness/conventions.md` §Baseline](../../../../harness/conventions.md#baseline),
Briefing [`AGENTS.md`](../../../../AGENTS.md) §1, Harness-Einstieg
[`harness/README.md`](../../../../harness/README.md), Doku-Gate
[`.d-check.yml`](../../../../.d-check.yml). [Roadmap](../in-progress/roadmap.md).

Kein Vertrag der Produkt-Achse berührt: **kein** Lastenheft-/Spezifikations-Bump, **keine** neue
`AC-*`-ID, **keine** ADR — dies ist eine Harness-/Konventions-Änderung. Die Adaption ist als
[MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
deklariert (eine undeklarierte Struktur-Änderung wäre die „stille Setzung", gegen die dieses Repo
sonst fail-closed vorgeht).

---

## 1. Was Etappe A tut

Die Baseline lag bisher **nur als URL** in `AGENTS.md` §1 — das Modell von `v1.3.0`. `v3.5.2`
erwartet sie **committet vendored**, damit die `../templates/…`-Verweise des Regelwerks („so sieht
das Artefakt aus") **netzlos lokal** auflösen und der Nachschlag über den `<tag>` reproduzierbar
ist. Genau das stellt diese Etappe her — **nichts weiter**: kein Modul-Delta, keine
`MR-*`-Bereinigung, keine Template-Konformität (Etappen B–D).

## 2. Umsetzung

| Schritt | Ergebnis |
|---|---|
| Bundle materialisiert | `.harness/baseline/v3.5.2/{regelwerk,templates}/` aus dem `lab-regelwerk.zip` des `v3.5.2`-Releases — **43 Dateien** (21 Regelwerk-Abschnitte, 21 Templates, 1 Manifest) |
| Integrität | `SHA256SUMS` über alle 42 Inhaltsdateien, `sha256sum -c` grün |
| Doku-Gate | [`.d-check.yml`](../../../../.d-check.yml) `scan.ignore: [".harness/baseline/**"]` — der vendored Baum ist **externer, unveränderter Fremdtext** mit eigenen Platzhaltern; ohne den Ausschluss prüfte `doc-check` fremde Template-Kennungen gegen a-checks Linkregeln |
| Deklaration | §Baseline auf `v3.5.2` + vendored Pfad; §Adoptierte Konventions-Quellen auf die vendored Lese-Form umgestellt; neue Adaption [MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert) |
| Briefing | `AGENTS.md` §1: Lese-Form ist jetzt der **vendored Abschnitt**, nicht die ZIP-URL — mit dem Hinweis, den *zur Aufgabe gehörenden* Abschnitt zu laden, nie das ganze Bundle |
| Einstieg | `harness/README.md` §Guides: Zeile auf den vendored Index umgestellt |

**Nummern-Entscheid dieser Etappe:** das Baseline-Template führt die Vendoring-Adaption als
[MR-003](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang); diese Nummer ist hier belegt
([`MR-003`](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang),
aufgelöst 2026-06-21). Etappe A nimmt die **nächste freie** ([MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)) — dieselbe Praxis wie
`ai-harness-init` (dort `MR-007`) <!-- d-check:ignore (fremde MR-Kennung des Repos ai-harness-init, nicht a-checks) --> — und verweist für die Frage der Nummern-Identität ausdrücklich
auf Etappe C. Damit bleibt A frei von Umnummerierungs-Folgen (die [MR-003](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang)-Verweise in
`harness/README.md` und die `d-check`-`ids`-Linkpflicht bleiben intakt).

## 3. Ergebnis und Verifikation

- `make gates` **grün**: `doc-check` **0 Befunde**. Die geprüfte Dateizahl wächst dabei **nur um die
  in diesem Branch neu entstandenen Eigen-Dokumente** (116 nach dem Vendoring-Commit, 117 mit dem
  Review-Report) — **nicht um die 43 vendored Fremddateien**; ohne den `scan.ignore` wären es rund
  158. Das ist der Beleg, dass der Ausschluss greift. `gate-consistency`, `arch-check` (0 Befunde),
  `lint`, `coverage-gate` (96,20 %), `guard-selftest` unverändert grün.
- `sha256sum -c SHA256SUMS` **grün** — Integritätsnachweis reproduzierbar.
- Die neuen Verweise lösen **netzlos** auf: `AGENTS.md` → vendored Index → `../templates/…`.

**Was bewusst *nicht* geprüft wird — und wohin es gehört.** `SHA256SUMS` belegt nur die **innere**
Konsistenz der Arbeitskopie, nicht ihre Herkunft. Die Baseline selbst trennt hier zwei Dinge
(`.harness/baseline/v3.5.2/regelwerk/modul-02-harness-bootstrap.md` §Freshness-Audit), und diese
Trennung ist zu übernehmen:

- **Freshness** („ist `v3.5.2` noch das aktuelle Release?") ist ausdrücklich eine
  **Netz-Operation außerhalb der Gates** — Wartung, kein Feedback-Gate — und prüft die
  **Release-Liste**, nicht das Asset (ein Hash auf den gepinnten Tag meldet „kein Drift", während
  upstream ein neuer Tag steht). Sie gehört **nicht** in `gate-consistency`, das netzlos läuft.
  Offen bleibt der **Auslöser** (ereignis-gebunden statt Kalenderpflicht) — zu deklarieren in
  Etappe C/D.
- **Asset-Integrität/Provenienz** (ist dieser Baum das Release?) ist gate-fähig — Modul 02 nennt
  dafür d-checks `sources`-Modul (`source-pin` auf den `sha256` des Assets); es ist in
  [`.d-check.yml`](../../../../.d-check.yml) heute **nicht** aktiviert. Kandidat für Etappe D.

**Korrektur gegenüber der ersten Fassung dieses Abschnitts:** dort war die Aktualitätsprüfung
pauschal als „Kandidat für `gate-consistency`" an Etappe D verwiesen — das widerspricht der eben
vendored Baseline. Fund des unabhängigen Zweit-Reviews (F-6).

## 4. DoD

- [x] `.harness/baseline/v3.5.2/{regelwerk,templates}/` + `SHA256SUMS` materialisiert, Selbsttest grün.
- [x] Drei Pin-Stellen umgestellt (`conventions.md` §Baseline + §Quellen, `AGENTS.md` §1, `harness/README.md` §Guides).
- [x] Adaption als [MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert) deklariert, mit Nummern-Hinweis auf Etappe C.
- [x] Vendored Baum aus dem Doku-Gate genommen, Dateizahl als Beleg.
- [x] `make gates` grün mit echter Ausgabe.
- [x] **Review-Report** unter [`docs/reviews/`](../../../reviews/2026-07-25-slice-047-baseline-vendoring.md) — erster Report in der **vendored Vorlagen-Form** (Modul 10): 0 HIGH, 2 MEDIUM (F-1 Provenienz nicht gegatet → Etappe D, F-2 Prozess), 2 LOW, 2 INFO; nicht merge-blockierend mit begründeter Abweichung.
- [x] **Unabhängiges Zweit-Review** ([Report](../../../reviews/2026-07-25-slice-047-baseline-vendoring-zweitreview.md), separate Instanz — Rollentrennung nach Modul 8/10 damit erfüllt): 0 HIGH, **6 MEDIUM**, 7 LOW, Verdikt **merge-blockierend**. Alle sechs MEDIUM vor dem Merge abgearbeitet (Auflösungs-Tabelle dort); die sieben LOW sind Etappe C/D zugeordnet.

## 5. Closure-Notiz

**Abgeschlossen und gemergt (2026-07-25).** Etappe A steht: die Baseline ist vendored, deklariert und
netzlos nachschlagbar. Der Weg dorthin ist der eigentliche Lerneintrag.

### Lerneinträge

1. **Ein Selbst-Review findet die eigene Abdeckungslücke nicht.** Das unabhängige Zweit-Review fand
   **sechs** MEDIUM, wo der Erst-Report zwei sah — und die Ursache war nicht Nachlässigkeit im
   Einzelnen, sondern die **Auswahl der betrachteten Bereiche**: der Negativbefund-Block hatte keine
   Zeile für `.harness/skills/**`, `roadmap.md`, `docs/reviews/README.md` und den `.d-check.yml`-Kopf.
   Genau dort lagen die Befunde. Die Negativbefund-Zeile macht die Lücke sichtbar, sie schließt sie
   nicht — das leistet nur eine zweite Instanz.
2. **Eine Versions-Hebung ist ein Sweep, kein Feld.** Zwei Dateien behaupteten nach dem Commit
   weiterhin `v1.3.0` (`reviewer.md`, `.d-check.yml`-Kopf) — beide außerhalb der drei „Pin-Stellen",
   die der Plan benannte. Wer einen Stand hebt, muss `grep` über den ganzen Baum laufen lassen und
   *historische* von *aktuell-behauptenden* Erwähnungen trennen.
3. **Eine Zahl als Beleg muss überall dieselbe sein.** „115" im Plan, „116" in Commit und Report —
   die Zahl war die einzige Evidenz dafür, dass der `scan.ignore` greift, und widersprach sich über
   drei Artefakte. Ersetzt durch eine Aussage, die nicht mitwächst: die Dateizahl steigt nur um die
   neuen **Eigen**-Dokumente, nicht um die 43 vendored.
4. **Die vendored Baseline widerlegt sofort eine eigene Planannahme.** Der Plan verwies die
   Aktualitätsprüfung an `gate-consistency`; Modul 02 — im selben Commit vendored — führt sie
   ausdrücklich als **Netz-Operation außerhalb der Gates**. Wer eine Baseline hereinholt, sollte die
   Aussagen, die sie zum eigenen Vorhaben trifft, **vor** dem Plan lesen, nicht danach.
