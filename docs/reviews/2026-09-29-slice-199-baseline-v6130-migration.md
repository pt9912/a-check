# Review-Report: slice-199 — 2026-09-29

**Review-Art:** Unabhängiger Lauf — geprüft gegen `slice-199`, die
Baseline-Regeln (`harness/conventions.md` §Baseline; `v6.13.0` ·
`regelwerk/modul-02-harness-bootstrap.md` §Freshness-Audit der vendored
Baseline) und die vier Commits selbst. Kern-Behauptungen sind
Bundle-Reproduzierbarkeit, Zeiger-Messung der MR-Durchgangs, Pin-Abdeckung
und Abschnitts-Gleichheit.

**Gegenstand:** Commits `2ce5da8..4a9ad1f` — `a6d19b6` (build(harness):
Baseline auf v6.13.0 gehoben), `296cf4e` (MR-012/015/022 nach
conventions/done), `18b925f` (Verweise nachziehen), `4a9ad1f`
(Ruhe-Marker).

**Skill:** `.harness/skills/reviewer.md` @ Stand `a6d19b6` (im geprüften
Range um die `v6.13.0`-Referenzen gebumpt) · <!-- d-check:ignore -->
**Modell:** glm-5.3-flash · **Datum:** 2026-09-29

**Eingangs-Kontext:** `.harness/skills/reviewer.md`, `slice-199` (§1–§9),
`harness/conventions.md` §Baseline, `v6.13.0` ·
`regelwerk/modul-02-harness-bootstrap.md` §Freshness-Audit, beide
Committs je Datei (`git show`), eigene Bundle-Bauten beider Tags aus dem
Kurs-Checkout (`git archive` + `tools/build-bundle.sh`), eigene
Abschnitts-Messungen.

---

## Findings

### M1 · Messung — MEDIUM

- **Quelle:** `AGENTS.md` §5 Mess-Regeln (zweiter Zähler) · `slice-199` §3 Schritt 3
- **Pfad:** `docs/plan/planning/in-progress/slice-199-baseline-v6130-migration.md:109` (Beleg-Zelle) · `a6d19b6` (Commit-Message, Absatz „Pins gemessen")
- **Befund:** Die Commit-Message behauptet „Pins gemessen (17 Dateien/42 Nennungen)"; die im Plan Schritt 3 versprochene „Zähltabelle Dateien/Nennungen, gemessen gegen den Stand vor dem Swap" ist in `slice-199` und im übrigen Repo nicht abgelegt. Die Zahl 42 ist mit keinem von drei anders gebauten Zählern reproduzierbar: 47 Zeilen-Nennungen, 63 Vorkommen, 36 Pfad-Muster-Treffer (jeweils `v6.6.0` in lebenden Dateien bei `2ce5da8`, ohne `.harness/baseline`, `done/`, `docs/reviews`, `conventions/done`); nur die Datei-Anzahl 17 trifft exakt.
- **Verifizierbar:** ja — `git grep -c v6.6.0 2ce5da8` mit den drei Zählern; Abweichung 42 ↔ {47, 63, 36} sichtbar.
- **Klasse:** Zählung ohne zweiten Zähler / Zähltabelle nicht abgelegt.

### M2 · Beleg — MEDIUM

- **Quelle:** `slice-199` §3 Schritte 5–7 (Beleg-Spalten) · `AGENTS.md` §6 Schritt 8
- **Pfad:** `docs/plan/planning/in-progress/slice-199-baseline-v6130-migration.md:111-113`
- **Befund:** Der Plan verspricht als Belege „Abschnitts-Messung je Eintrag im Slice dokumentiert" (Schritt 5), „Befundliste mit Abgleich-Ergebnis" (Schritt 6, Voll-Abgleich) und die Stichprobe gegen den Bestand (Schritt 7); keine der drei Befundlisten ist im Slice oder sonst im Repo abgelegt. Für die vier Einträge ohne `Ersetzt-Baseline-Regel`-Zeiger (`MR-019`, `MR-020`, `MR-023`, `MR-024`) ist der Ausgang aus den fünf Klassen (`v6.13.0` · `regelwerk/modul-02-harness-bootstrap.md` §Freshness-Audit) nirgends benannt — ihr Durchgang ist behauptet, nicht belegt. Die drei Nachfolge-Einträge tragen ihre Abschnitts-Messungen vollständig; die Lücke betrifft den Rest des Durchgangs.
- **Verifizierbar:** ja — `grep -rl "Voll-Abgleich"` über den Bestand; die Treffer beschränken sich auf die Regel-Zitate in `slice-199` und `harness/conventions.md`.
- **Klasse:** Durchgangs-Beleg versprochen, nicht abgelegt.

### M3 · Messungs-Vermerk — MEDIUM

- **Quelle:** `AGENTS.md` §5 Mess-Regeln (Beleg korrekt)
- **Pfad:** `harness/conventions/MR-026-welle-closure-ohne-replay-v6130.md:17-20`
- **Befund:** „Die Schritte 3 und 4 tragen den Planungs-Bestand-Lese-Schritt (Welle 150) und die Prosa-Erschöpfungs-Regel (Welle 152)" — beide neuen Blöcke stehen nachweislich in Closure-Schritt 3 (`v6.13.0` · `regelwerk/modul-06-roadmap.md`, Zeilen 220 ff.); Schritt 4 (Zeitdokumente archivieren, Zeile 296) ist zwischen `v6.6.0` und `v6.13.0` unverändert. Der Eintrag ist `Accepted` und damit nach der Disziplin des Adaptions-Blocks nur über einen Nachfolge-Eintrag korrigierbar. Die tragende Aussage (Replay-Zeile Schritt 1 wortgleich, Ausgang *bleibt gültig*) ist von dem Fehler nicht berührt.
- **Verifizierbar:** ja — `diff` beider `modul-06-roadmap.md`-Fassungen; 11 Hunks, keiner berührt Schritt 4.
- **Klasse:** Schritt-Zuordnung im Messungs-Vermerk falsch.

### M4 · Eigen-Form — MEDIUM

- **Quelle:** `AGENTS.md` §5 (Risiko-Ausgang, DoD-Größen-Regel) · `make verify-risiko-ausgaenge`, `make doc-structure`
- **Pfad:** `docs/plan/planning/in-progress/slice-199-baseline-v6130-migration.md:117` (DoD) · `:163-181` (§7)
- **Befund:** `make verify` ist an `4a9ad1f` rot (Exit 2): `verify-risiko-ausgaenge` meldet drei §7-Risiken ohne Ausgang aus der geschlossenen Dreier-Menge — zwei stehen auf „Ausgang: bei Closure", einer kürzt auf „→ Register", wo der Sensor `Beobachtungs-Register` oder `BEO-<NNN>` verlangt; `doc-structure` meldet 4 Task-Items im DoD (3 ignoriert), erlaubt sind 3. Ursprung ist der Plan-Commit `2ce5da8` (unterhalb des Ranges), aber der Zustand ist an HEAD lebendig: der Closure-Trigger (§6: `make verify` grün) ist solange unerreichbar, und das Risiko 2 („MR-Durchgang … dritter Fund = Lücke") trägt seinen Ausgang nicht.
- **Verifizierbar:** ja — `make verify` (Exit 2, Meldungen im Lauf-Log).
- **Klasse:** Ausgang-Kurzform und DoD-Überbau am eigenen Slice-Artefakt.

### L1 · Form — LOW

- **Quelle:** `v6.13.0` · `templates/harness/conventions/MR-NNN-titel.template.md` (Pflichtfeld „Löst auf")
- **Pfad:** `harness/conventions/MR-025-referenzmatrix-grandfathering-v6130.md:29` · `MR-026-welle-closure-ohne-replay-v6130.md:24` · `MR-027-verfeinerungs-form-v6130.md:27`
- **Befund:** Die drei Nachfolge-Einträge nennen das Ablöse-Feld „**Löst ab:**"; die Ziel-Form nennt es „**Löst auf:**". Der Bestand spaltet sich dadurch weiter (8 Einträge „Löst auf" — `MR-010` bis `MR-017`, `MR-020` —, 5 „Löst ab" — `MR-021`, `MR-022`, die drei neuen); der Link geht jeweils auf die Index-Zeile, wie die Ziel-Form es verlangt.
- **Verifizierbar:** ja — `grep "Löst a" harness/conventions/`.
- **Klasse:** Feldname gegen die Ziel-Form.

### L2 · Zählung — LOW

- **Quelle:** `AGENTS.md` §5 Mess-Regeln (zweiter Zähler) · `slice-199` §2
- **Pfad:** `docs/plan/planning/in-progress/slice-199-baseline-v6130-migration.md:55-56`
- **Befund:** „sämtliche 23 Regelwerk-Dateien und 19 der 20 Vorlagen" — geändert sind 23 von **26** Regelwerk-Dateien (unverändert: `grundlagen-bootstrap.md`, `modul-14-docker-harness.md`, `modul-15-observability.md`), und die Basis „20 Vorlagen" ist gegen den Kurs-Bestand (26 `.md` unter `lab/templates`) nicht reproduzierbar; 19 geänderte Template-Pfade stimmen. Die Summe 42 Dateien, +691/−325, ist exakt.
- **Verifizierbar:** ja — `git diff --stat v6.6.0 v6.13.0 -- lab/regelwerk lab/templates` im Kurs-Checkout; `comm` der Dateilisten.
- **Klasse:** Delta-Zählung ohne bezifferte Basis.

### L3 · Lesart — LOW

- **Quelle:** `AGENTS.md` §5 Zitier-Form · `harness/conventions.md` §Baseline (Zitier-Form in einfrierenden Artefakten)
- **Pfad:** `AGENTS.md:283` · `harness/sensors/doc-mentions.md:10` · `.d-check.yml:10,182,188`
- **Befund:** Drei lebende Dateien zitieren `v6.6.0` ohne Pfad-Link, und die Lesart ist pro Datei undeutlich: `.d-check.yml` deklariert ihre drei Kennungen im Kopf ausdrücklich als Herkunft („das ist Herkunft, keine Angabe über den aktuellen Stand"); `doc-mentions.md` zitiert `v6.6.0` · `templates/harness/README.template.md` ohne solche Deklaration, obwohl die Norm wortidentisch in `v6.13.0` steht (`regelwerk`-Vorlage, Zeile 119); `AGENTS.md` nennt `v6.6.0` neben der Zitier-Form-Regel, während `reviewer.md` dieselbe Norm-Familie nach dem Bump auf `v6.13.0` zitiert — eine der beiden Stellen ist je nach Lesart (Herkunft vs. Stand) veraltet.
- **Verifizierbar:** ja — Textvergleich der zitierten Normen zwischen beiden Template-Fassungen.
- **Klasse:** Herkunft-Lesart undeklariert.

### I1 · INFO

- **Quelle:** —
- **Pfad:** `docs/plan/planning/open/slice-200-adoption-v6130-agents-und-matrix.md`
- **Befund:** Während des Reviews wanderte HEAD von `4a9ad1f` auf `ca27b26` („docs(planning): slice-200 in open — Adoption der v6.13.0-Aspekte"). Damit tragen die in `slice-199` §2 benannten Adoption-Aspekte (Welle 148 WIP-Limit, Welle 149 §5-Form, Welle 138 Matrix-Klassen) jetzt eine Kennung; die Lastenheft-RB-Aspekte laufen über `slice-189`. Der Review-Gegenstand dieses Laufs bleibt `2ce5da8..4a9ad1f`; `ca27b26` ist nicht geprüft.
- **Verifizierbar:** —
- **Klasse:** Lauf-Horizont.

---

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Bundle-Reproduzierbarkeit (DoD-Punkt 1) | geprüft, ohne Befund — `git archive v6.13.0` + `tools/build-bundle.sh` reproduziert `.harness/baseline/v6.13.0` byte-gleich (`diff -rq` leer); dieselbe Gegenprobe auf `v6.6.0` reproduziert den aus `2ce5da8` extrahierten Altbestand byte-gleich; `SHA256SUMS` 54/54 OK |
| Genau ein Stand + Symlinks | geprüft, ohne Befund — keine getrackte Datei unter `v6.6.0` verbleibt, vier `.claude/rules`-Symlinks auf `v6.13.0`; `make regelwerk-check`, `make symlink-check` grün (eigene Ausführung) |
| Zeigerwanderung MR-014/016 (Commit-Message: „Zielabschnitte wortgleich") | geprüft, ohne Befund — `modul-15-observability.md` ändert nur die Quelle-Zeile; `modul-08-agentenrollen.md` §Die neun Übergaben und ihre Artefakte ist wortgleich (`difflib`-Vergleich: identisch) |
| Messzahlen MR-025/026/027 | geprüft, ohne Befund — alle sechs Werte exakt reproduzierbar (23 242→23 083, 10 552→13 522, 10 942→6 574; Bytes inkl. Überschrift-Zeile, ohne Quelle-Zeile); Wortgleichheit der Zeile 45 (`modul-06`, Replay) und Zeile 17 (`grundlagen-source-precedence`, Suffix) bestätigt; „4 Hunks" bei `-U3` bestätigt; kein Hunk berührt die Grandfathering-Aussage |
| Freeze-Disziplin der abgelösten Einträge | geprüft, ohne Befund — `MR-012`/`MR-015`/`MR-022` wurden **nicht** auf `v6.13.0` mitgehoben: der Zeiger steht als Kennung (`v6.6.0` · `regelwerk/…` §Abschnitt) eingefroren in `done/`; der Move-Commit `296cf4e` ist ein reiner `git mv` (100 % Similarity), Inhalt im Vorgänger-Commit — Reihenfolge der Regel §3.3 eingehalten |
| Historische Aussagen (Massen-Ersetzung, Risiko 1) | geprüft, ohne Befund — CHANGELOG-Historie, `docs/plan/planning/done/`, `docs/reviews/`, `conventions/done/`, Beobachtungs-evidence sind im Range unangetastet; keine historische Nennung wurde mitgehoben |
| reviewer.md „Besetzter Fall" nach dem Bump | geprüft, ohne Befund — die Norm „die Klassen-Bezeichnung muss über Läufe hinweg stabil sein" steht in `v6.13.0` · `templates/docs/reviews/review-report.template.md` (Zeile 92) im vierten Kommentarblock, wie in `v6.6.0` — der Bump der Zitier-Form ist belegt, nicht still mitgeschoben |
| §1-Abgrenzung / Scope | geprüft, ohne Befund — kein Produkt-Code, kein Spec-Stratum, kein ADR berührt; alle Nicht-Baseline-Edits sind Pin-Nachzüge, MR-Durchgang, CHANGELOG (Pflicht aus `AGENTS.md` §6 Schritt 7) und Ruhe-Marker |
| Commit-Disziplin | geprüft, ohne Befund — alle vier Messages tragen `slice-199` (drei zusätzlich `AC-QA-02`/`AC-QA-03`); der einzige `(planning)`-Commit des Ranges (`4a9ad1f`) berührt ausschließlich `docs/plan/planning/` |
| Gate-Lauf (eigene Ausführung) | geprüft, ohne Befund — `make gates` Exit 0 (inkl. `doc-check`: 601 Dateien, 0 Befunde; `doc-targets`, `gate-consistency`, `version-coherence` grün) |

---

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 4 |
| LOW | 3 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Zählung ohne zweiten Zähler / Zähltabelle nicht abgelegt · Durchgangs-Beleg versprochen, nicht abgelegt · Schritt-Zuordnung im Messungs-Vermerk falsch · Ausgang-Kurzform und DoD-Überbau am eigenen Slice-Artefakt · Feldname gegen die Ziel-Form · Delta-Zählung ohne bezifferte Basis · Herkunft-Lesart undeklariert · Lauf-Horizont.

## Verdikt

**Merge-blockierend:** nein — die Commits sind getreten, und keine der Stufen ist gegen sie gerichtet; die vier MEDIUM blockieren **die Closure** von `slice-199` und sind dort heilbar: M1/M2 durch Ablage der versprochenen Zähl- und Befundlisten (§3 Schritt 3/5/6/7) vor dem `git mv`, M4 ohnehin durch den roten `make verify`, M3 über einen Nachfolge-`MR` oder eine ausdrückliche Korrektur vor der Immutabilitäts-Wirkung. Der Substanz-Kern der Migration — byte-gleiches Vendoring beider Richtungen, exakt messbare Zeiger-Wanderung, eingefrorene statt mitgehobener Altnennungen, vollständiger Scope — trägt in jeder nachgebauten Gegenprobe.
