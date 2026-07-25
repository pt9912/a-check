# slice-047 — Etappe A: Baseline committet vendoren (`v1.3.0` → `v3.5.2`)

**Status:** in-progress — Etappe **A** aus
[slice-046 §6](../open/slice-046-regelwerk-v352-migration-analyse.md), am 2026-07-25 per
Maintainer-Wort gezogen („A zuerst"). Ergebnis §3.
**Auslöser:** Maintainer-Vorgabe: **vollständige Migration nach `v3.5.2`**, auch gegen bestehende
`MR-*`; Etappen-Schnitt aus der Delta-Analyse.
**Bezug:** Baseline-Deklaration
[`harness/conventions.md` §Baseline](../../../../harness/conventions.md#baseline),
Briefing [`AGENTS.md`](../../../../AGENTS.md) §1, Harness-Einstieg
[`harness/README.md`](../../../../harness/README.md), Doku-Gate
[`.d-check.yml`](../../../../.d-check.yml). [Roadmap](roadmap.md).

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

- `make gates` **grün**: `doc-check` 115 Dateien / **0 Befunde** — dieselbe Dateizahl wie vor dem
  Vendoring, der `scan.ignore` greift also (sonst wären es 158). `gate-consistency`, `arch-check`
  (0 Befunde), `lint`, `coverage-gate` (96,20 %), `guard-selftest` unverändert grün.
- `sha256sum -c SHA256SUMS` **grün** — Integritätsnachweis reproduzierbar.
- Die neuen Verweise lösen **netzlos** auf: `AGENTS.md` → vendored Index → `../templates/…`.

**Was bewusst *nicht* gegatet ist:** die **Aktualität** des vendored Baums gegen den deklarierten
`<tag>` (also „liegt hier wirklich `v3.5.2`?"). Ein solcher Check ist ein Kandidat für
`gate-consistency` — analog zur bestehenden Pin-Konsistenz —, braucht aber eine Quelle für den
Soll-Zustand und gehört damit in Etappe D, nicht hier. Bis dahin trägt `SHA256SUMS` nur die
**innere** Konsistenz, nicht die Herkunft.

## 4. DoD

- [x] `.harness/baseline/v3.5.2/{regelwerk,templates}/` + `SHA256SUMS` materialisiert, Selbsttest grün.
- [x] Drei Pin-Stellen umgestellt (`conventions.md` §Baseline + §Quellen, `AGENTS.md` §1, `harness/README.md` §Guides).
- [x] Adaption als [MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert) deklariert, mit Nummern-Hinweis auf Etappe C.
- [x] Vendored Baum aus dem Doku-Gate genommen, Dateizahl als Beleg.
- [x] `make gates` grün mit echter Ausgabe.
- [ ] **Review-Synthese** unter [`docs/reviews/`](../../../reviews/) (Regelwerk Modul 10).

## 5. Closure-Notiz

_(beim Abschluss.)_
