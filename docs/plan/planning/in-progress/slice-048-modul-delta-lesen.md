# slice-048 — Etappe B: Modul-Delta gegen `v3.5.2` lesen

**Status:** in-progress — Etappe **B** aus
[slice-046 §6](../open/slice-046-regelwerk-v352-migration-analyse.md), am 2026-07-25 per
Maintainer-Wort gezogen. Ergebnis §2–§4.
**Auslöser:** Etappe A ([slice-047](../done/slice-047-baseline-vendoring.md)) hat die Baseline
vendored; jetzt wird sie **gelesen** — netzlos aus `.harness/baseline/v3.5.2/regelwerk/`.
**Bezug:** liefert die Findings, aus denen Etappe C (`MR-*`-Bereinigung) und D
(Template-Konformität) geschnitten werden. Kein Vertrag der Produkt-Achse berührt.
[Roadmap](../in-progress/roadmap.md).

> **Eigen-Probe zu Fund B-1:** dieser Slice hält bewusst **≤ 3 DoD-Punkte** (§5) — die Regel, deren
> Verletzung er meldet.

---

## 1. Gelesen — und was nicht

**Gelesen (6 Module, 831 Zeilen):** `modul-02` (Bootstrap/Modus), `modul-04` (ADRs), `modul-05`
(Planning/Lifecycle), `modul-08` (Rollen), `modul-11` (Verifikation), `modul-13` (Gates).

**Nicht gelesen** — und damit ausdrücklich **nicht** bewertet: `modul-00/01/03/06/07/09/12/14/15/16`
sowie die drei Grundlagen-Abschnitte (Konventionen, Klassifikation, Durchsetzungsschicht). Etappe B
war auf die fünf in [slice-046 §5](../open/slice-046-regelwerk-v352-migration-analyse.md) benannten
plus `modul-08` zugeschnitten. **Besonders lückenhaft bleibt `modul-07` (Carveouts)** — a-check hat
bis heute **null** Carveouts (`CO-NNN` „bisher ungenutzt"), und mehrere gelesene Module verweisen
für die Werkzeug-Wahl bei Diskrepanz dorthin.

## 2. Findings — abweichende Praxis

| # | Kategorie | Quelle | Befund (gemessen) |
|---|---|---|---|
| **B-1** | **HIGH** | `modul-05` §Ziel-Form Slice | **Slice-Größe:** die Baseline setzt **≤ 3 DoD-Punkte** und **höchstens zwei Schichten** als *harte* Schnitt-Regel („zu groß ⇒ zurück zum Schneiden"). Gemessen: slice-047 **7**, slice-046 **6**, slice-044 **4** DoD-Punkte. Nur slice-040/041 liegen mit 2 darunter. Die Regel ist im Repo nirgends abgebildet — weder in `AGENTS.md` noch in `conventions.md`. |
| **B-2** | **HIGH** | `modul-05` §Ziel-Form Sub-Area-Modus-Begründung, `modul-02` §Kernidee | **Modus pro Sub-Area:** die Baseline verlangt **je berührter Sub-Area** einen Begründungsblock (Modus · Konventionen-Dichte · Phase-Reife · Evidenz-/Diskrepanz-Risiko · Reconciliation-Aufwand) als **§8 des Slice-Plans**. `conventions.md` deklariert stattdessen **einen pauschalen Repo-Modus** (`*` → Greenfield) — genau das Anti-Pattern „ein Repo/Slice hat einen Bootstrap-Modus". Die §8-Praxis existierte und ist **zweimal stillschweigend eingeschlafen**: vorhanden in slice-001…008 und 032…035, **fehlt in 009…031 und 036…047** (12 von 47). |
| **B-3** | **MEDIUM** | `modul-11` §Fitness Function ohne Standard-Tool | **Kein `verify`-Target.** Die Baseline trennt `make gates` (Code-/Architektur-Fragen) von `verify:` (DoD-/Closure-Fragen) und verlangt, dass der Implementer `make verify-*` **selbst vor der „fertig"-Meldung** läuft. a-check hat **kein** `verify`-Target; die DoD-Prüfung läuft als Prosa im Slice-Dokument — also als Behauptung ohne Sensor, genau die „häufigste Verifier-Lücke". |
| **B-4** | **MEDIUM** | `modul-11`, `modul-05` §Closure | **`closure-note-reviewer` fehlt.** Die Baseline führt eine Vorlage dafür (`templates/.harness/skills/closure-note-reviewer.template.md`, seit Etappe A vendored, in `SHA256SUMS` Zeile 35) und der vorhandene `reviewer.md` verweist selbst auf den „Schwester-Skill … (Modul 11)". `.harness/skills/` trägt aber nur `reviewer.md`. Der fehlende Skill trägt die *semantische* Prüfung „Lerneintrag vs. Floskel" — deterministisch prüfbar ist nur die Struktur. |
| **B-5** | **MEDIUM** | `modul-05` §Closure | **Lerneintrag-Form nicht benannt.** Die Baseline verlangt den Lerneintrag in **einer von drei Formen** (geschärfte Regel · neuer Sensor · benannte Spec-Lücke) plus **zwei beobachtbare** Closure-Kriterien. a-checks Closure-Notizen schreiben freie Prosa und ordnen sich keiner Form zu — prüfbar wäre das erst mit Form-Angabe. |
| **B-6** | **MEDIUM** | `modul-05` §Lifecycle, `AGENTS.md` §5 | **Lifecycle unvollständig abgebildet.** Die Baseline führt **fünf** Übergänge, darunter die zwei **Rückführungen** `in-progress→next` (zu groß) und `in-progress→open` (Blocker). `AGENTS.md` §5 nennt nur die Vorwärts-Kette `open → next → in-progress → done` — die Rückführungen fehlen, es gibt für sie also weder Regel noch Präzedenz. Der Zustand `next` wurde in 47 Slices **genau einmal** benutzt (slice-009, `9cb5ffa`); der Normalweg ist `open → in-progress` direkt (slice-047 wurde sogar direkt in `in-progress` angelegt). Das **WIP-Limit = 1** („harte Größe, kein Vorschlag") ist nirgends deklariert. |
| **B-7** | **MEDIUM** | `modul-04` §ID-Schema-Bezug | **[MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration) verweist auf die „Kurs-ADR-Vorlage `v1.3.0`".** Das ist eine **aktuell-behauptende** Versionsaussage, kein historischer Hinweis — das Zweit-Review von Etappe A hatte diese Zeile als „historisch legitim" eingeordnet, was bei genauem Lesen nicht trägt. |
| **B-8** | **LOW** | `modul-13` §Vorhanden ≠ behauptet | **Maintenance-Target fehlt.** Modul 13 nennt ausdrücklich ein `regelwerk-check`-Target als Beispiel für „vorhanden, aber **nicht** als Gate behauptet". Genau diese Klasse fehlt a-check — und genau dorthin gehört der **Freshness-Audit** aus [slice-047 §3](../done/slice-047-baseline-vendoring.md) (Etappe-A-Fund F-6), der als Netz-Operation nicht in `gates` darf. |
| **B-9** | **LOW** | `modul-08` §Neun Übergaben | **Rollen ohne Übergabe-Artefakte.** Die Baseline führt sechs Rollen und **neun** Übergabe-Artefakte; ohne Artefakt gibt es „keinen Rollenwechsel, nur einen Kontext-Switch". a-check fährt faktisch eine Instanz; heute wurde **erstmals** eine Rolle real getrennt (unabhängiger Reviewer, slice-047). Verifier und Validator existieren als Rolle nicht. |
| **B-10** | **INFO** | `modul-07` (nicht gelesen) | **Carveout-Werkzeug ungenutzt.** `CO-NNN` ist seit [MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration) deklariert und **nie** verwendet. Mehrere gelesene Module verweisen für die Werkzeug-Wahl bei Diskrepanz auf Modul 07; ob a-check je einen Carveout gebraucht hätte, ist ohne dessen Lektüre nicht beurteilbar. |

## 3. Negativbefunde (geprüft, ohne Befund)

- geprüft, ohne Befund: **Review-Form** (`modul-10`) — Kategorien, Output-Schema und Negativbefund-Pflicht sind unverändert; die vier Reports vom 2026-07-25 sind konform. Neu ist nur die Auflage „HIGH-Liste mit **≥ 2 repo-spezifischen** Regeln" im Reviewer-Skill — [`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md) erfüllt sie (ADR-Verstoß, Suppression ohne ADR, Docker/make-only).
- geprüft, ohne Befund: **ADR-Hard-Rule** (`modul-04`) — „`Accepted` wird nie überschrieben, Korrektur = Folge-ADR mit `Supersedes`" ist in `AGENTS.md` §3.5 verkörpert **und** maschinell durchgesetzt (`doc-immutable`, `vcs`-Modul über die Commit-Range). Zwei reale Supersede-Ketten im Bestand.
- geprüft, ohne Befund: **ADR-Form** (`modul-04`) — „mindestens drei verglichene Alternativen, jede mit Trade-off" und „jede Entscheidung mit Architektur-Wirkung bekommt eine Fitness Function": [ADR-0027](../../adr/0027-constructs-roh-text-monopol.md) (5 Alternativen), [ADR-0028](../../adr/0028-ziel-glob-schattenwurf.md) (3), [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) (4), alle mit `## Fitness Function`.
- geprüft, ohne Befund: **Gate-ID-Bindung** (`modul-13` §Fitness Function aus einem ADR-Satz) — a-checks Gate-Targets tragen die ID im Help-Text (`lint` → [ADR-0005](../../adr/0005-lint-profil.md), `coverage-gate` → [ADR-0006](../../adr/0006-coverage-gate.md), `arch-check` → [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), `image-test` → [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk), `trace-check` → [ADR-0021](../../adr/0021-commits-modul-trace-check.md)).
- geprüft, ohne Befund: **Sensors-Tabelle ohne Lauf-Status** (`modul-13` §Hard Rule) — `harness/README.md` trägt drei Spalten, keine Status-Spalte; der Rückbau lief bereits am 2026-07-25 ([MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration)-Commit).
- geprüft, ohne Befund: **Gate-Landschaft** — 13 Targets über sechs Sensor-Klassen (Linter, Test, Coverage, Architektur, Meta, Distribution); die Modul-13-Warnung „nur die generischen sechs" trifft nicht zu.
- geprüft, ohne Befund: **WIP faktisch** — `in-progress/` trägt neben der Roadmap genau **einen** Slice; das Limit wird eingehalten, nur nicht deklariert (Teil von B-6).
- **nicht geprüft** (kein Negativbefund, sondern Lücke): `modul-00/01/03/06/07/09/12/14/15/16` + drei Grundlagen-Abschnitte (§1).

## 4. Was daraus für C und D folgt

- **Etappe C** (`MR-*`-Bereinigung) bekommt aus B zwei Punkte über die Nummern-Kollision hinaus:
  **B-2** (der pauschale Repo-Modus in `conventions.md` §Modus-Deklaration ist gegen die Baseline
  falsch geschnitten) und **B-7** ([MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration)-Versionsaussage). Dazu die sieben LOW aus dem
  Etappe-A-Zweit-Review.
- **Etappe D** (Template-Konformität) bekommt **B-1** (Slice-Vorlage mit §8 und Größen-Regel),
  **B-3** (`verify`-Target), **B-4** (`closure-note-reviewer`), **B-5** (Lerneintrag-Form),
  **B-6** (Lifecycle-Rückführungen + WIP in `AGENTS.md`), **B-8** (`regelwerk-check` als
  Maintenance-Target, nimmt den Etappe-A-Fund F-6 auf).
- **Neu aufgeworfen:** **B-9/B-10** sind keine Template-Frage, sondern Betriebsmodell (Rollen,
  Carveouts). Sie gehören **nicht** in C/D, sondern brauchen eine eigene Entscheidung —
  Vorschlag: als **Etappe E** aufsetzen oder ausdrücklich verwerfen.

## 5. DoD

- [ ] Sechs Module gegengelesen, Findings mit Kategorie, Quelle und **gemessener** Repo-Seite (§2).
- [ ] Negativbefunde und **nicht gelesene** Bereiche ausgewiesen (§3, §1).
- [ ] Zuordnung der Findings zu C/D/E (§4) — die Grundlage für den nächsten Etappen-Schnitt.

## 6. Closure-Notiz

_(beim Abschluss.)_
