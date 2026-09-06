# Review-Report: slice-167 — 2026-09-06

**Review-Art:** Plan — geprüft gegen den Slice-Plan und die Vendoring-Tat selbst
(Modul 10 §Drei Review-Arten): Kern-Behauptungen sind ZIP-Provenienz,
Datei-Integrität und die Vollständigkeit des Pointer-Bumps.

**Gegenstand:** Commits `92e1f64` ("feat(harness): Baseline-Pin auf
v6.2.0 -- Etappe A") und `5649d84` (".claude/rules-Symlinks auf v6.2.0
korrigiert")

**Skill:** `.harness/skills/reviewer.md` @ Stand `5649d84` (unverändert
seit Anlage) · <!-- d-check:ignore -->
**Modell:** claude-sonnet-5 · **Datum:** 2026-09-06

**Eingangs-Kontext:**

- `docs/plan/planning/done/slice-167-etappe-a-vendoring-v620.md`
- `docs/plan/planning/done/slice-161-regelwerk-v610-delta-analyse.md`
- `docs/plan/planning/done/slice-164-regelwerk-v620-delta-analyse.md`
- `harness/conventions.md`, `AGENTS.md` §1/§5
- `.harness/baseline/v6.0.0/`, `.harness/baseline/v6.2.0/`
- Kurs-Repo `pt9912/ai-harness-course`, Tag `v6.2.0` (frischer Download,
  nicht dem Slice-Zitat vertraut)

---

## Findings

### F-1 — `harness/conventions.md` §Adoptierte Konventions-Quellen nicht auf `v6.2.0` gehoben

- `kategorie`: HIGH
- `quelle`: eigener Vergleich gegen §Baseline im selben Dokument
- `pfad`: `harness/conventions.md` (vor Fix) Zeilen 45–48
- `befund`: Die Zeilen zeigten weiterhin auf `.harness/baseline/v6.0.0/regelwerk/README.md`
  bzw. `.../templates/README.md`, während direkt darüber (§Baseline) korrekt `v6.2.0` steht —
  ein Selbstwiderspruch im selben Dokument. Per Slice-eigener Definition („verweist auf den
  aktuell vendorten Stand") ist das eindeutig Klasse „strukturell", nicht der
  `Ersetzt-Baseline-Regel`-Anker-Klasse zuzuordnen. Kein Gate fängt das, weil der `v6.0.0`-Pfad
  technisch weiter auflöst (`v6.0.0` bleibt ja vendored).
- `verifizierbar`: ja — Sichtprüfung beider Abschnitte im selben Dokument.
- `klasse`: Struktureller Pointer beim Bump übersehen (Selbstwiderspruch im Dokument)

### F-2 — „neun Dateien"-Behauptung unpräzise formuliert

- `kategorie`: MEDIUM
- `quelle`: `diff -rq .harness/baseline/v6.0.0 .harness/baseline/v6.2.0`
- `pfad`: `docs/plan/planning/done/slice-167-etappe-a-vendoring-v620.md` (vor Fix) §3
- `befund`: Ein naiver lokaler Verzeichnis-Diff zeigt 31 inhaltlich verschiedene Dateien, nicht
  neun — jede `regelwerk/*.md`-Datei trägt eine `<!-- Quelle: …/blob/<tag>/… -->`-Kommentarzeile,
  die bei jedem Versionssprung zwangsläufig mitläuft, unabhängig von echten Regeländerungen. Die
  Zahl „neun" war in der Sache korrekt (deckungsgleich mit der Vereinigung aus `slice-161`/
  `slice-164`), bezog sich aber unausgesprochen auf den *upstream*-`git diff --stat`, nicht auf
  einen lokalen Verzeichnis-Diff, wie die Formulierung nahelegte.
- `verifizierbar`: ja — `diff -rq` gegen die committeten Bäume, exakt reproduzierbar.
- `klasse`: Präzisions-Lücke zwischen Upstream-Diff und lokalem Datei-Diff

## Negativbefunde

- geprüft, ohne Befund: ZIP-Provenienz — `gh api …/releases/tags/v6.2.0` liefert denselben
  Digest wie ein frischer, unabhängiger Download und `sha256sum`; drei unabhängige Kanäle
  (API, eigener Download, Release-eigene `SHA256SUMS`) stimmen überein.
- geprüft, ohne Befund: Integrität — `sha256sum -c` gegen `.harness/baseline/v6.2.0/SHA256SUMS`:
  alle 53 Einträge OK, keine unmanifestierte/fehlende Datei.
- geprüft, ohne Befund: Datei-**Pfad**-Baum-Identität zu `v6.0.0` — exakt gleich, wie behauptet.
- geprüft, ohne Befund: Zwei-Klassen-Vollständigkeit im Rest des Repos (nach F-1-Fix) — alle
  verbleibenden `v6.0.0`-Vorkommen sind entweder historische Snapshots (Changelogs,
  Beobachtungs-Register, abgeschlossene Slices) oder `Ersetzt-Baseline-Regel`-Anker akzeptierter
  `MR`-Dateien, korrekt unverändert.
- geprüft, ohne Befund: `harness/README.md` §Guides — neue Reviewer-Skill-Zeile zeichengleich mit
  dem in `slice-161` §4.1 vorgeschlagenen Wortlaut.
- geprüft, ohne Befund: keine `harness/conventions/MR-*.md`-Datei wurde in `92e1f64` inhaltlich
  verändert.
- geprüft, ohne Befund: `make regelwerk-check` — grün, meldet zwei vendorte Stände korrekt,
  prüft `v6.2.0`, weist `v6.0.0` als ungeprüft (nicht als Fehler) aus.
- geprüft, ohne Befund: `make gates`/`make verify` — beide Exit 0.
- geprüft, ohne Befund: Kopf-Metadaten, DoD-Größe, Closure-Notiz-Form, Risiko-Ausgänge
  strukturell konsistent zu `slice-165`/`slice-166`.
- geprüft, ohne Befund: `.claude/rules/`-Symlink-Korrektur (Commit `5649d84`) — alle vier
  Symlinks lösen jetzt auf `v6.2.0` auf.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 1 |
| LOW | 0 |
| INFO | 0 |

**Finding-Klassen dieses Laufs:** Struktureller Pointer beim Bump
übersehen (Selbstwiderspruch im Dokument) · Präzisions-Lücke zwischen
Upstream-Diff und lokalem Datei-Diff

## Verdikt

**Merge-blockierend:** nein mehr — beide Findings vor Abschluss von
`slice-167` korrigiert. Die Provenienz- und Integritätsprüfung der
eigentlichen Vendoring-Tat (der riskanteste Teil dieses Slice) ist über
drei unabhängige Kanäle bestätigt und tadellos; die beiden Findings
betreffen Doku-Vollständigkeit bzw. Formulierungspräzision, nicht die
Substanz des Vendorings.

**Übergabe:** Findings sind bereits in `slice-167`s Closure eingearbeitet.
Dieser Report ist Lauf-Beleg, keine Verifikation — DoD-/Spec-Konformität
prüft der Verifier separat (Modul 11).
