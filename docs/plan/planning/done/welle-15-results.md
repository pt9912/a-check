# welle-15 — Ergebnis: Regelwerk-Migration `v6.2.0` → `v6.5.0`

**Geschlossen:** 2026-09-07. **Plan:** [welle-15](welle-15/welle-15-regelwerk-v650-migration.md).

---

## 1. Geliefert

| Slice | Etappe | Ergebnis |
|---|---|---|
| [slice-174](welle-15/slice-174-regelwerk-v650-delta-analyse.md) | — | Delta-Analyse: 30 Dateien roh, **16 mit Inhalt**, 14 nur Formatierung |
| [slice-175](welle-15/slice-175-etappe-a-vendoring-v650.md) | **A** | `v6.5.0` vendored, über vier Kanäle verifiziert; Stand an drei Stellen, vier Symlinks, 13 lebende Zeiger |
| [slice-179](welle-15/slice-179-etappe-b-adaptions-durchgang-v650.md) | **B** | alle sieben aktiven Adaptionen bewertet — keine löst auf; `exempt-paths` eingeordnet |
| [slice-176](welle-15/slice-176-zitier-form-einfrierende-artefakte.md) | **C** | Zitier-Form übernommen, 16 Links umgestellt, voriger Stand entfernt |
| [slice-177](welle-15/slice-177-sensors-struktur-zwei-tabellen.md) | **D** | 14 Sensor-Dateien, zwei Tabellen; Zellen über 250 Zeichen **16 → 4** |
| [slice-178](welle-15/slice-178-slice-form-ziel-und-abgrenzung.md) | **E** | Slice-Form §1/§9 und Workflow-Schritt 4; Kopieranleitung gegen den Bestand geprüft |

## 2. Was funktioniert hat

**Messen vor Planen.** Der Roh-Diff nannte 30 Dateien und `+572`; nach Abzug
zweier Formatierungs-Klassen blieben 16 Dateien und `+452`, und **14** trugen
gar keinen Inhalt. Ohne diesen Schnitt wäre der Etappen-Plan an der falschen
Stelle groß geworden.

**Der Beweis in Etappe C.** Vorigen Stand entfernt, `make gates` unmittelbar
danach Exit 0 — **null** Nachzüge in eingefrorenen Artefakten. slice-172 hatte
für denselben Vorgang **22** gebraucht. Die Zitier-Form ist damit nicht
behauptet, sondern belegt.

**Zwei Sensoren im ersten echten Einsatz.** Die in slice-173 gebauten
`versions`-Baseline-Pins und `make symlink-check` fuhren hier ihren ersten
Migrations-Lauf — und fanden beide eine Fehlkalibrierung an sich selbst
(übersehene relative Form, Über-Treffen auf fremde URLs). Beide behoben,
mutations-belegt.

## 3. Was anders lief

**Jeder der sechs Reviews war merge-blockierend.** Zusammen 11 HIGH und 30
MEDIUM — und die Befunde trafen dreimal nicht eine Zahl, sondern die **Methode**:

- slice-174 nannte das richtige Bereinigungs-Kriterium und wandte es halb an.
- slice-177 nannte das richtige Auswahl-Kriterium und begründete damit, bei vier
  von vierzehn Dateien aufzuhören.
- slice-178 wollte eine Anleitung reparieren, die vom Bestand abwich — und wich
  an drei neuen Stellen selbst ab.

**Dreimal `git mv` mit Inhaltsänderung im selben Commit** (slice-172, 176, 177).
Beim dritten Mal war der Schaden eingetreten statt möglich: `R040`,
`git log --follow` verlor fünf Commits. Alle drei nachträglich aufgesplittet —
möglich nur, weil sie unveröffentlicht waren.

**Ein History-Rewrite auf bereits gepushten Commits.** Bei slice-172 wurde
„nichts gepusht" behauptet, ohne `origin/main` zu befragen; der Maintainer hatte
gepusht. Zurückgebaut, die Regel steht seither im Gedächtnis: Der Remote-Stand
ist eine **Messung**, keine Erinnerung.

## 4. Steering-Loop-Einträge

**Zwei Einträge erreichten die Schwelle und haben ihren Ausgang** (Lese-Schritt):

- [`BEO-PLAN/review-geltungsbereich-zu-eng`](../observations/BEO-PLAN/review-geltungsbereich-zu-eng/observation.md)
  bei **3×** → **verkörpert** als Regel *Geltungsbereich einer Messung* —
  liegt in `AGENTS.md §5`, Anker `seit slice-179`. Kein Sensor: ob ein
  Geltungsbereich weit genug ist, ist ein Urteil.
- [`BEO-PLAN/verweis-auf-wandernden-slice`](../observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md)
  bei **7×** → **verkörpert mit benannter Lücke**; den Ausgang trägt
  [slice-180](../open/slice-180-slice-mv-dritte-verweis-form.md). `make slice-mv`
  kennt zwei der drei Verweis-Formen; die dritte trat in dieser Welle
  **zehnmal** auf.

**Vier Einträge bekamen Belege, ohne die Schwelle zu erreichen:**
`muster-trifft-nur-die-haeufige-schreibweise` (neu, 1×),
`versions-sensor-trifft-planungs-vorgriff` (neu, 2×),
`baseline-regel-nie-erwogen-weil-bestand-sie-verletzt` (2×),
`zwei-baseline-staende-nach-migrationsende` (**verkörpert**, aufgelöst durch C).

## 5. Folge-Slices

- [slice-180](../open/slice-180-slice-mv-dritte-verweis-form.md) — `slice-mv` lernt
  die dritte Verweis-Form.
- Zwei benannt, ohne Datei: die verbliebenen zwei Historie-Zellen in `AGENTS.md`
  §4 (slice-177 §1) und die Zitier-Form im `**Welle:**`-Feld der Archiv-Stubs
  (slice-176 §3). Sie entstehen bei Priorisierung — kein Slice erfindet eine
  Adresse, die niemand annimmt.

## 6. Verifikation

| Prüfung | Ergebnis |
|---|---|
| Alle sechs Slices in `done/` | ✅ |
| Stand-Deklaration an drei Stellen auf `v6.5.0` | ✅ `conventions.md` §Baseline · `AGENTS.md` §1 · `harness/README.md` §Guides; **kein** `v6.2.0` mehr in lebenden Dateien |
| `make doc-check` ohne `version-stale` | ✅ — erster echter Migrations-Lauf des Sensors |
| `make symlink-check` | ✅ sieben Symlinks, Baseline-Ziele auf `v6.5.0` |
| `make regelwerk-check` | ✅ **ein** vendorter Stand, 54 Dateien, kein „ungeprüft" |
| `make ci` (Replay-Ersatz, [`MR-015`](../../../../harness/conventions.md#mr-015)) | ✅ Exit 0 |
| Carveout-Audit | ✅ Bestand **null** |
| Bootstrap-aware-Gate-Audit | ✅ keine gestufte Schwelle im Repo |
| ADR-Re-Evaluierungs-Audit | ✅ keine der 39 hängt an der Baseline-Migration |
| Unabhängiger Review je Slice | ✅ sechs von sechs, jeder merge-blockierend, jeder eingearbeitet |

**Was diese Verifikation nicht belegt:** ob die vier neuen Sensor-Konfigurationen
(`versions`-Baseline-Muster, `symlink-check`, die zwei erweiterten `exempt-paths`)
beim **nächsten** Sprung tragen. Sie sind an `v6.2.0`→`v6.5.0` kalibriert, und
genau das ist die Klasse, die diese Welle zweimal gefunden hat.
