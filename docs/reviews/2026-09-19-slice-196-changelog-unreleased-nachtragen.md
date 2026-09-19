# Review-Report: slice-196 — 2026-09-19

**Review-Art:** Code (Doku-/Vertragsänderung) — geprüft gegen den Slice-Plan, die
Konventionen und die Repo-Artefakte, die der Diff behauptet (Modul 10 §Drei
Review-Arten). **Nicht** gegen die DoD — das ist die Verifikation, eine andere
Rolle.

**Gegenstand:** slice-196, Commits `cc5e10f` … `417f8a0` (Kern: `8cae28a` der
CHANGELOG-Block, `836323f` die `AGENTS.md`-Zeile), Stand `417f8a0`.

**Skill:** `.harness/skills/reviewer.md` @ `60e7b66` (seit slice-193 unverändert) ·
**Modell:** unbekannt (Subagent) · **Datum:** 2026-09-19

> **Zitier-Form** (Norm, `v6.6.0` · `templates/docs/reviews/review-report.template.md`):
> Dieser Report friert ein; was er zitiert, bewegt sich weiter. Deshalb **Kennung,
> nicht Adresse** — `slice-NNN` statt seines Lifecycle-Pfads, `make <target>` statt
> eines Links auf die Sensor-Datei, eine Baseline-Stelle als **Tag + Pfad in
> Inline-Code** statt als Link (`` `v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt> ``).
> Das `pfad`-Feld auf den **geprüften Gegenstand** ist davon nicht betroffen — es
> hält den Stand des Laufs fest und darf das.

**Eingangs-Kontext:**

- Slice-Plan slice-196 (`in-progress/`, Stand `417f8a0`) — Plan-Review gegen Spec/ADR-Bezug
- `AGENTS.md` §3.7 (Kommentar-Regeln), §4, §5 (Doku-Regeln, drei Mess-Regeln), §6 Schritt 7
- `harness/conventions.md` §Baseline, §Modus-Deklaration, §Anforderungs-Anlege-Prozess
- `docs/user/releasing.md` §Versionsquelle, §Freigabe-Checkliste
- `.d-check.yml` (`ids`/`matrix`/`mentions`/`versions`), `Makefile`, `harness/README.md` §Sensors
- Beobachtungs-Register `BEO-PLAN/changelog-unreleased-ungepflegt` (`observation.md`, `state.md`, `evidence/`)
- Folge-Slice slice-197 (`open/`), `git log`/`git show` gegen `v0.19.0`
- `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice ·
  `modul-06-roadmap.md` §Das Beobachtungs-Register und §Wann Arbeit eine Welle braucht ·
  `modul-10-review-harness.md`

---

## Findings

### F-1 — §1 behauptet, das Lastenheft habe sich nicht geändert

- `kategorie`: HIGH
- `quelle`: Reviewer-Skill §Klassifikation („nachweislich falsche Tatsachenbehauptung,
  gegen ein Repo-Artefakt verifiziert") · `AGENTS.md` §5
- `pfad`: `docs/plan/planning/in-progress/slice-196-changelog-unreleased-nachtragen.md:29`
- `befund`: „`spec/` änderte sich in slice-194 und slice-154 (Form-Nachzug), das Lastenheft
  gar nicht" — das Lastenheft steht bei `v0.19.0` auf **0.26.0** und heute auf **0.27.0**;
  geändert hat es slice-154 (`12d90ed`, +36 Zeilen, neue §5/§6 **und** die Historie-Zeile
  0.27.0). Der Satz trägt die zentrale Aussage des Slice („die konsumenten-sichtbare
  Änderung ist genau eine") und widerspricht demselben Halbsatz, der `spec/` als geändert
  ausweist.
- `verifizierbar`: ja — `git log v0.19.0..HEAD -- spec/lastenheft.md` und der
  `**Version:**`-Kopf bei `v0.19.0` gegen HEAD.
- `klasse`: Tatsachenbehauptung über den Änderungsumfang nicht gegengeprüft

### F-2 — Der Register-Beleg zählt 55 statt 56 Harness-Slices

- `kategorie`: MEDIUM
- `quelle`: `v6.6.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
  (Beleg-Form) · Reviewer-Skill §Mess-Regeln („Wer eine Menge zählt, zählt sie zweimal
  verschieden")
- `pfad`: `docs/plan/planning/observations/BEO-PLAN/changelog-unreleased-ungepflegt/evidence/slice-196.md:6`
- `befund`: „die übrigen 55 sind Harness-Arbeit" — von den 58 Slices führt der
  `[Unreleased]`-Abschnitt genau **zwei** einzeln im sichtbaren Teil (194 und 195), also
  bleiben **56**. Der Beleg zählt 194 doppelt (je einmal über den `-- internal/`- und den
  `-- spec/`-Aufruf) und subtrahiert 3 statt 2. Plan §1 („die übrigen 56") und der
  Harness-Block, der slice-154 ausdrücklich als Harness-Arbeit führt, sagen 56 — der
  Block trägt 56 Kennungen.
- `verifizierbar`: ja — die beiden `git log`-Mengen schneiden sich in 194; der Block
  listet 56 Kennungen abzüglich der beiden oberen.
- `klasse`: Summe zweier Schnittmengen statt Vereinigung

### F-3 — Der Guide-Trigger nennt drei Dokumente; die Änderung ohne Doku-Berührung erreicht er nicht

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §3.7 (Grenze der Zusage) · Reviewer-Skill §Klassifikation
  („unbelegte Tatsachenbehauptung")
- `pfad`: `AGENTS.md:368`
- `befund`: „Ein Slice, der Lastenheft, Spezifikation oder Benutzerhandbuch ändert, trägt
  seinen Eintrag in `[Unreleased]` **in sich**" — im Bestand tragen mindestens zwei
  konsumenten-sichtbare Einträge Slices, die **keines** der drei berührt haben: slice-038
  (`internal/adapter/driven/config/`, Tie-Break der Schicht-Zuordnung, unter `[0.15.0]`
  geführt) und slice-041 (`internal/adapter/driven/graph/`, Legenden-Layout, unter
  `[0.16.0]` geführt) — beide ohne `spec/`- und ohne `docs/user/`-Änderung. Der eigene
  Beleg nennt den weiteren Trigger („`spec/` **und die Regeln** ja"); der Träger führt ihn
  nicht.
- `verifizierbar`: ja — `git show --stat` der beiden Slice-Commits gegen die drei im Satz
  genannten Dokumente.
- `klasse`: Zusage enger formuliert als ihr Gegenstand

### F-4 — §1 nennt vier Einträge, §2 und der Vorzustand drei

- `kategorie`: LOW
- `quelle`: Reviewer-Skill §Mess-Regeln
- `pfad`: `docs/plan/planning/in-progress/slice-196-changelog-unreleased-nachtragen.md:26`
- `befund`: „die **vier** Einträge der Slices 194/195" — der Abschnitt trug vor dem Slice
  **drei** Bullet-Einträge (`### Changed — BREAKING` und `### Added` aus 194, `### Fixed`
  aus 195). §2 desselben Plans und `evidence/slice-196.md` nennen drei.
- `verifizierbar`: ja — `git show 8cae28a^:CHANGELOG.md`, Abschnitt zwischen
  `## [Unreleased]` und `## [0.19.0]`.
- `klasse`: Zählung im Dokument widerspricht sich selbst

### F-5 — „Alle stehen im Gate-Index" trifft die Targets, nicht die Konfigurationen

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §4 (der Gate-Index steht in `harness/README.md` §Sensors) ·
  `v6.6.0` · `regelwerk/modul-13-*.md` §Vorhanden ≠ behauptet
- `pfad`: `CHANGELOG.md:71`
- `befund`: Die vier Targets des Bullets haben je eine Zeile im Gate-Index; die im selben
  Bullet genannten „geschärften Konfigurationen" nicht — das `versions`-Muster, die
  Aktivierungs-Regel der `.d-check.yml` und der Gate-Index „einmal" stehen in
  `.d-check.yml` bzw. `AGENTS.md` §4, und nur „Zellengrenzen" ist in der
  `doc-structure`-Zeile erwähnt. Der Abschnitt verspricht, was sonst am Harness steht, „als
  Zustand, den man im Repo nachsehen kann"; „Alle stehen im Gate-Index" führt für die
  Konfigurations-Hälfte an eine Stelle, die sie nicht trägt.
- `verifizierbar`: ja — Suche über `harness/README.md` §Sensors.
- `klasse`: Verweis-Ziel weiter als der belegte Umfang

### F-6 — „den alten ließ slice-172 fallen" zeigt für den genannten Sprung auf den Nachbarvorgang

- `kategorie`: LOW
- `quelle`: Reviewer-Skill §Klassifikation (belegbare Aussage gegen ein Repo-Artefakt)
- `pfad`: `CHANGELOG.md:52`
- `befund`: Im selben Bullet ist der Sprung als `v5.12.0` → `v6.6.0` benannt; „der alte"
  liest sich damit als `v5.12.0` — fallen gelassen hat den aber slice-140 (`ab2ada8`).
  slice-172 ließ **`v6.0.0`** fallen (`18e6696`). Beide Kennungen stehen in derselben
  Aufzählung; wer den Satz nachprüft, landet beim falschen Vorgang.
- `verifizierbar`: ja — `git log --diff-filter=D -- .harness/baseline/<tag>`.
- `klasse`: Herkunfts-Zeiger auf den Nachbarvorgang

### F-7 — §3 nennt `AGENTS.md` §5 als Ablage, gebaut ist §6 Schritt 7

- `kategorie`: LOW
- `quelle`: `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice
  (§3 Umsetzung)
- `pfad`: `docs/plan/planning/in-progress/slice-196-changelog-unreleased-nachtragen.md:61`
- `befund`: Die Umsetzungstabelle führt „die Modus-Deklaration oder `AGENTS.md` §5";
  DoD (Zeile 77), §8 (Zeile 130) und die Umsetzung selbst nennen §6 Schritt 7. §5
  desselben Dokuments trägt die Doku-Regeln und die drei Mess-Regeln — der Ausgang liegt
  dort nicht.
- `verifizierbar`: ja — Sichtprüfung beider Abschnitte.
- `klasse`: Plan-Stelle und Lieferort auseinander

### F-8 — Messungen gegen `v0.19.0..HEAD` tragen ihren Endpunkt nicht

- `kategorie`: LOW
- `quelle`: Reviewer-Skill §Mess-Regeln (Geltungsbereich einer Messung; `seit slice-179`)
- `pfad`: `docs/plan/planning/in-progress/slice-196-changelog-unreleased-nachtragen.md:25`
  und `:128`
- `befund`: „`v0.19.0..HEAD` trägt **265** Commits" (§1, Indikativ) und „265 Commits" in der
  Closure-Notiz §8 — der Anker wandert mit jedem Commit des Slice; zur Closure sind es
  **270** (fünf eigene Commits), ohne dass ein Schnittpunkt genannt ist. Die 58 Slices
  tragen, weil sie an `done/` hängen; die Commitzahl tut es nicht.
- `verifizierbar`: ja — `git log --oneline v0.19.0..HEAD | wc -l` gegen `417f8a0`.
- `klasse`: Messung gegen einen wandernden Anker

## Negativbefunde

- geprüft, ohne Befund: **Vollständigkeit des `[Unreleased]`-Abschnitts** — zwei verschieden
  gebaute Zähler (Kennungs-Präfixe gegen alle `1[0-9]{2}`-Zahlen des Abschnitts) liefern
  beide genau die **58** Nummern 135–188 und 192–195, keine mehr und keine weniger; ein
  dritter Zähler über die Slice-Kennungen der Commit-Messages des Fensters liefert dieselben
  58 plus die eigene (196).
- geprüft, ohne Befund: **„die konsumenten-sichtbare Änderung ist genau eine"** —
  `git log v0.19.0..HEAD -- internal/` nennt genau zwei Commits, beide slice-194.
- geprüft, ohne Befund: **Gate-Index ↔ Makefile** — alle vier genannten Targets existieren
  als Makefile-Regel **und** haben eine Zeile in `harness/README.md` §Sensors; `make
  doc-targets` Exit 0. Ein zweiter, anders gebauter Zähler (Target-Diff `v0.19.0` → HEAD)
  findet genau `archive-wave`, `archive-wave-test`, `dcheck-phrase-selftest`,
  `doc-mentions`, `doc-reviews`, `symlink-check`; die ersten beiden sind im Block genannt.
- geprüft, ohne Befund: **vendored Baseline** — `.harness/baseline/` trägt genau einen Stand
  (`v6.6.0` mit `regelwerk/`, `templates/`, `SHA256SUMS`); `make regelwerk-check` ist als
  „Integrität der vendored Baseline gegen SHA256SUMS" deklariert; `make symlink-check`
  Exit 0, 7 Symlinks, Baseline-Ziele auf `v6.6.0`. Die Herkunft `tools/archive-wave/` „aus
  dem Schwester-Repo" ist im Volltext von slice-144 belegt („aus `d-check` übernommen"),
  der Pin `d-check v0.75.0` steht in `d-check.mk`.
- geprüft, ohne Befund: **Register-Ausgang und Zählung** — `evidence/` trägt genau drei
  Dateien (127, 133, 196) ⇒ 3×; `state.md` trägt `verkörpert` mit Zielort (`AGENTS.md` §6
  Schritt 7) und Herkunfts-Anker `seit slice-196`, also Zielort **und** Anker wie die
  Ziel-Form verlangt.
- geprüft, ohne Befund: **§3.7 auf den neuen Texten** — die `AGENTS.md`-Zeile nennt Regel
  und Grenze im Indikativ, ohne die verworfene Alternative zu beschreiben; `state.md` trägt
  im `Stand:`-Feld Zustand und auflösbaren Anker, und die Prosa darunter folgt der im
  Register eingeführten Form (vgl. `BEO-GATE/probe-liefert-den-gegenstand-mit`: dieselbe
  Wendung „Drei Ausprägungen, alle drei …", dieselbe „Kein Sensor"-Begründung). Kein Satz
  beschreibt einen abwesenden früheren Zustand.
- geprüft, ohne Befund: **Referenz-Richtung und Linkpflicht** — `matrix` führt nur
  `spec/**`-Klassen, der CHANGELOG ist kein Spec-Stratum; er ist für alle `ids`-Muster
  deklariert ausgenommen (`.d-check.yml` §ids), die Ausnahme ist also erklärt und nicht
  übersehen. `make doc-check` über 586 Dateien, 0 Befunde; die relativen Pfade des Plans
  (`../../../../CHANGELOG.md`, `../../adr/0017-…`, `../open/slice-197-…`) lösen aus jedem
  Lifecycle-Verzeichnis auf.
- geprüft, ohne Befund: **Folge-Slice** — slice-197 existiert in `open/`, nennt slice-196
  als Übernahme-Quelle und führt die drei ADR-Kern-Befunde, die §7 ausschließt.
- geprüft, ohne Befund: **die Grenze „Kein Gate deckt das"** — `tools/gate-consistency.sh`
  prüft (A) `version.md#aktuell` gegen das jüngste CHANGELOG-Release; kein Gate verbindet
  eine Änderung mit einem Eintrag. Die Aussage ist gedeckt, nicht behauptet.
- geprüft, ohne Befund: **Roadmap-Ruhe-Marker** — mit dem Slice in `in-progress/` ist der
  Marker entfernt; `make doc-planning` Exit 0.

## Gates

| Lauf | Exit | Geltungsbereich |
|---|---|---|
| `make gates` | 0 | alle inneren Gates (`doc-check` 586 Dateien/0 Befunde, `doc-targets`, `gate-consistency`, `version-coherence`, `symlink-check` 7 Symlinks, `dcheck-phrase-selftest`, `guard-selftest`, `ci-range-selftest`, `suppression-check`, `record-gates`) — repo-weit über den Arbeitsbaum `417f8a0` |
| `make verify` | 0 | `verify-risiko-ausgaenge`, `verify-observations`, `doc-structure`, `doc-complete` (21 Anforderungen, 0 Waisen) |
| `make trace-check` | 0 | Message-Traceability über `HEAD~1..HEAD` |
| `make doc-check` (in `gates`) | 0 | Markdown-Links, Anker, Kennungs-Linkpflicht, Referenzmatrix, Spans, Hostpfade, Versions-Kohärenz, Review-Deckung — **nicht**: ob Prosa wahr ist |

**Nicht gefahren:** `make ci`/`make image-test` (Image-Bau; für einen Diff ohne
`internal/`-Änderung nicht gegenstandsnah), `make doc-immutable` (der Slice schließt die
drei Befunde ausdrücklich aus — Gegenstand von slice-197), `make image-scan` (Netz).
**Was kein Lauf prüft — Geltungsbereich der Belege:** ob die Prosa des Harness-Blocks
*zutrifft*. Genau diese Grenze benennt die `AGENTS.md`-Zeile selbst („Kein Gate deckt das");
sie ist damit für diesen Report der Grund, die Aussagen einzeln gegen `git` und den Baum zu
halten — nicht der Grund, sie für geprüft zu halten.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 2 |
| LOW | 5 |
| INFO | 0 |

**Finding-Klassen dieses Laufs:** Tatsachenbehauptung über den Änderungsumfang nicht
gegengeprüft · Summe zweier Schnittmengen statt Vereinigung · Zusage enger formuliert als
ihr Gegenstand · Zählung im Dokument widerspricht sich selbst · Verweis-Ziel weiter als der
belegte Umfang · Herkunfts-Zeiger auf den Nachbarvorgang · Plan-Stelle und Lieferort
auseinander · Messung gegen einen wandernden Anker

## Verdikt

**Merge-blockierend:** ja — F-1 ist eine gegen `git` verifizierte falsche Aussage im
Plan-Kopf, die die zentrale Einordnung des Slice trägt („genau eine konsumenten-sichtbare
Änderung"); F-3 betrifft die Substanz des gelieferten Guides, nicht seine Form. F-2 ist der
Beleg eines Register-Eintrags bei 3× und damit Zähler-relevant. Die fünf LOW-Befunde
blockieren nicht.

**Übergabe:** Findings gehen an den Implementer. Die **Finding-Klassen** gehen zusätzlich in
die Slice-Closure §7 und von dort in den Zähler. Dieser Report ist ein **Lauf-Beleg** (dieser
Diff, dieser Skill, dieses Modell, dieses Verdikt) und ersetzt keine Verifikation — DoD-/
Spec-Konformität prüft der Verifier separat (Modul 11; anderes Prüf-Artefakt, anderer
Eingabe-Kontext).
