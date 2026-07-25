# slice-048 — Etappe B: Modul-Delta gegen `v3.5.2` lesen

**Status:** in-progress — Etappe **B** aus
[slice-046 §6](../open/slice-046-regelwerk-v352-migration-analyse.md), am 2026-07-25 per
Maintainer-Wort gezogen. Ergebnis §2–§5.
**Auslöser:** Etappe A ([slice-047](../done/slice-047-baseline-vendoring.md)) hat die Baseline
vendored; jetzt wird sie **gelesen** — netzlos aus `.harness/baseline/v3.5.2/regelwerk/`.
**Bezug:** liefert die Findings, aus denen die restlichen Migrations-Etappen geschnitten werden.
Kein Vertrag der Produkt-Achse berührt. [Roadmap](../in-progress/roadmap.md).

> **Lese-Umfang korrigiert (2026-07-25):** die erste Fassung las **sechs** Module und wies den Rest
> als Lücke aus. Auf Maintainer-Hinweis — *„wir wollen nach Regelwerk v3.5.2 migrieren"* — ist der
> Zuschnitt hinfällig: eine vollständige Migration verträgt keine ungelesenen Module. Jetzt sind
> **alle 17 Module und alle drei Grundlagen-Abschnitte gelesen** (2867 Zeilen). Die Erweiterung
> brachte **elf** zusätzliche Funde (B-11 … B-21), darunter drei, die ohne sie unsichtbar
> geblieben wären.

> **Eigen-Probe zu Fund B-1:** dieser Slice hält bewusst **≤ 3 DoD-Punkte** (§6) — die Regel, deren
> Verletzung er meldet.

---

## 1. Gelesen

**Vollständig (20 Dateien, 2867 Zeilen):** `modul-00` … `modul-16` sowie die drei
Grundlagen-Abschnitte (Konventionen, Klassifikation, Durchsetzungsschicht).

Keine ausgewiesene Lese-Lücke mehr. Wo ein Modul **nicht** zu einem Fund führt, steht das als
Negativbefund in §4 — Schweigen wird nicht als Konformität gewertet.

## 2. Der Befund hinter den Befunden

Die auffälligste Regelmäßigkeit ist keine einzelne Regel, sondern ein **Muster im Zeitverlauf**:
die Harness-Praktiken der Bootstrap-Phase sind nach der Fundament-Welle **verfallen**.

| Praktik | vorhanden in | fehlt ab |
|---|---|---|
| **§8 Sub-Area-Modus-Begründung** | slice-001…008, **kurz wiederbelebt** 032…035 | 009…031, 036…047 |
| **Steering-Loop-Eintrag** | slice-001…008 | **009…048 — lückenlos, nie wieder** |

Zwölf von 48 Slices tragen die Modus-Begründung, acht von 48 einen Steering-Loop-Bezug. Das ist
genau der Verfall, den die Baseline unter *Entropy Management* beschreibt — und er trifft
ausgerechnet die Mechanik, die ihn hätte melden sollen. Die Einzelfunde unten sind zu großen
Teilen Symptome davon.

**Selbstanwendung, unangenehm aber messbar:** die Baseline setzt die Schwelle bei **3×**
(„1× Vorfall · 2× Symptom · 3× Lücke im Harness"). Der Fehler *„`make …` in eine Pipe gehängt und
den roten Gate-Lauf dadurch verschluckt"* trat allein am 2026-07-25 **viermal** auf. Nach
Baseline-Regel ist das seit dem dritten Mal keine Unachtsamkeit mehr, sondern eine
**Harness-Lücke ohne Sensor** — dokumentiert ist sie bis heute nur in einer Agenten-Memory
außerhalb des Repos, also genau dort, wo der nächste Lauf sie nicht zwingend sieht (B-21).

## 3. Findings

### 3.1 Planung und Lifecycle

| # | Kat. | Quelle | Befund (gemessen) |
|---|---|---|---|
| **B-1** | **HIGH** | `modul-05` §Ziel-Form Slice | **Slice-Größe:** Baseline setzt **≤ 3 DoD-Punkte** und höchstens zwei Schichten als *harte* Schnitt-Regel („zu groß ⇒ zurück zum Schneiden"). Gemessen: slice-047 **7**, slice-046 **6**, slice-044 **4**. Die Regel ist im Repo nirgends abgebildet — weder `AGENTS.md` noch `conventions.md`. |
| **B-2** | **HIGH** | `modul-05` §Sub-Area-Modus, `modul-02` §Kernidee, `grundlagen-konventionen` §Sub-Area | **Modus pro Sub-Area:** Baseline verlangt je berührter Sub-Area einen Begründungsblock (Konventionen-Dichte · Phase-Reife · Evidenz-Risiko · Reconciliation-Aufwand) als §8 des Slice-Plans, und Sub-Areas qualifizieren über **drei Inklusions-Achsen, mind. zwei erfüllt**. `conventions.md` deklariert stattdessen **einen pauschalen Repo-Modus** (`*` → Greenfield) — das ausdrücklich benannte Anti-Pattern. Praxis-Verfall siehe §2. |
| **B-6** | **MEDIUM** | `modul-05` §Lifecycle, `AGENTS.md` §5 | **Lifecycle unvollständig:** Baseline führt **fünf** Übergänge inkl. der Rückführungen `in-progress→next` (zu groß) und `in-progress→open` (Blocker). `AGENTS.md` §5 nennt nur die Vorwärtskette — für die Rückführungen gibt es weder Regel noch Präzedenz. **WIP-Limit = 1** („harte Größe, kein Vorschlag") ist nirgends deklariert (faktisch eingehalten). |
| **B-18** | **LOW** | `modul-05`, `grundlagen-konventionen` §Verzeichniskonvention | **`docs/plan/planning/next/` existiert gar nicht** — obwohl `AGENTS.md` §5 den Zustand nennt. In 48 Slices wurde er **genau einmal** benutzt (slice-009, `9cb5ffa`) und danach mitsamt Verzeichnis aufgelöst. Ein deklarierter Zustand ohne Ort ist eine stille Setzung. |
| **B-5** | **MEDIUM** | `modul-05` §Closure | **Lerneintrag-Form nicht benannt:** Baseline verlangt eine von **drei** Formen (geschärfte Regel · neuer Sensor · benannte Spec-Lücke) plus zwei beobachtbare Closure-Kriterien; a-check schreibt freie Prosa ohne Form-Zuordnung. |
| **B-12** | **MEDIUM** | `modul-06` §Fünf Abschnitte | **Roadmap: vierter Pflicht-Abschnitt fehlt.** Vorhanden sind *Aktuelle Welle · Nächste Wellen · Meilensteine · Abgeschlossene Wellen* (plus ein Extra *Abhängigkeitsgraph*); es fehlt **Historische Trigger-Verschiebungen** — das Drift-Log. Ohne es ist jede Umplanung still, und genau das nennt die Baseline die Hälfte der Auditierbarkeit. |
| **B-13** | **MEDIUM** | `modul-06` §Wellen-Closure | **Wellen werden nie auditierbar geschlossen.** Die Baseline gibt eine 5-Schritt-Prozedur mit Beleg je Schritt und `done/welle-NN-results.md` vor. Gemessen: **null** Welle-Ergebnisnotizen, **null** Welle-Plandateien — Wellen existieren nur als Prosa-Überschrift in der Roadmap (`welle-10a` o. ä.). Damit fehlen Closure-Log-Zeiger, Wave-Self-Close-Commit und Carveout-Audit-Punkt. |

### 3.2 Carveouts und Diskrepanz-Werkzeug

| # | Kat. | Quelle | Befund (gemessen) |
|---|---|---|---|
| **B-14** | **MEDIUM** | `modul-07`, `grundlagen-konventionen` §Verzeichniskonvention | **`docs/plan/carveouts/` existiert nicht.** `CO-NNN` ist seit [MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration) als ID-Reihe deklariert, aber es gibt weder Verzeichnis, noch Vorlage, noch Audit-Slice (`SL-CO-AUDIT-<welle>`), noch die Freshness-Regel (>90 Tage ungeprüft ⇒ HIGH). Ein deklariertes Werkzeug ohne Ort — dieselbe Klasse wie B-18. |
| **B-10** | **LOW** | `modul-07` §Werkzeug-Wahl | **Der Diskrepanz-Trichter ist ungenutzt.** Die Baseline trennt Carveout · BF-Sub-Area-Markierung · permanente ADR über zwei sequenzielle Fragen (Granularität **vor** Temporalität). a-check kennt nur ADRs — jede Diskrepanz landet zwangsläufig im schärfsten Werkzeug, was das Fehlen von B-2/B-14 mitverursacht. *(hochgestuft aus INFO, nachdem `modul-07` gelesen war — die erste Fassung stufte ohne Lektüre.)* |

### 3.3 Verifikation, Durchsetzung, Rollen

| # | Kat. | Quelle | Befund (gemessen) |
|---|---|---|---|
| **B-3** | **MEDIUM** | `modul-11`, `grundlagen-konventionen` §Referenz-Richtung | **Kein `verify`-Target.** Baseline trennt `make gates` (Code-Fragen) von `verify:` (DoD-/Closure-Fragen) und hängt dort u. a. `check-references` fail-closed ein. a-check hat **kein** `verify`; die DoD-Prüfung ist Prosa im Slice — Behauptung ohne Sensor. |
| **B-11** | **MEDIUM** | `modul-09` §Hard Rules („zwei Quadranten") | **Suppression-Verbot ist halb durchgesetzt.** `AGENTS.md` §3.2 verbietet `//nolint`; `.golangci.yml` erklärt das in **Zeile 3 als Kommentar** — aber `nolintlint` ist **nicht** unter den aktivierten Lintern. Die Hard Rule hat den inferential-feedforward-Quadranten und **keinen** computational-feedback-Quadranten; ein `//nolint` im Code liefe heute durch `make lint`. Billigster Fix im ganzen Bericht: einen Linter aktivieren. |
| **B-21** | **MEDIUM** | `grundlagen-klassifikation` §Steering Loop, `modul-06` §Closure | **Kein Steering-Loop-Kanal.** Die 3×-Regel braucht einen Ort, an dem wiederkehrende Fehlermuster gezählt werden — Baseline: Steering-Loop-Einträge in der Wellen-Closure-Notiz. a-check hat den Bezug in slice-001…008 und **seither in keinem** der 40 folgenden Slices (§2). Konkreter Beleg für die Kosten: der Pipe-Fehler mit 4 Vorfällen an einem Tag hat bis heute keinen Sensor. |
| **B-4** | **MEDIUM** | `modul-11`, `modul-05` §Closure | **`closure-note-reviewer` fehlt.** Die Vorlage liegt seit Etappe A vendored (`templates/.harness/skills/closure-note-reviewer.template.md`, `SHA256SUMS` Zeile 35), und `reviewer.md` verweist selbst auf den „Schwester-Skill … (Modul 11)". `.harness/skills/` trägt nur `reviewer.md`. Der fehlende Skill trägt die semantische Prüfung „Lerneintrag vs. Floskel". |
| **B-16** | **LOW** | `grundlagen-durchsetzungsschicht` §Drei Bindepunkte | **Dritter Bindepunkt fehlt.** Tool-Call-Gate (`pretooluse-command-guard.sh`) und Handoff-Gate (`stop-require-gates.sh`) existieren und sind in `.claude/settings.json` verdrahtet; das **Workflow-Skelett** als Slash-Command (`.claude/commands/`) fehlt — das Verzeichnis existiert nicht. Es ist ausdrücklich der *schwächste* der drei (inferential), aber es ist der, der den 8-Schritt-Pfad im Lauf hält. |
| **B-9** | **LOW** | `modul-08` §Neun Übergaben | **Rollen ohne Übergabe-Artefakte.** Sechs Rollen, neun Übergabe-Artefakte; ohne Artefakt „kein Rollenwechsel, nur ein Kontext-Switch". a-check fährt faktisch eine Instanz; am 2026-07-25 wurde **erstmals** eine Rolle real getrennt (unabhängiger Reviewer, slice-047). Verifier und Validator existieren nicht. |

### 3.4 Spec-Form, Betrieb, Altlasten

| # | Kat. | Quelle | Befund (gemessen) |
|---|---|---|---|
| **B-15** | **MEDIUM** | `modul-03` §Ziel-Form Akzeptanzkriterium | **AC-Form weicht ab.** Baseline: drei Pfade im Given/When/Then-Stil — **Happy · Boundary · Negative** — plus Out-of-Scope. Gemessen über 19 `AC-*`: **keines** trägt die Drei-Pfad-Gliederung; die Form ist „Beschreibung"-Prosa mit eingebetteten Grenz- und Negativsätzen. Immerhin **16 von 19** tragen einen Out-of-Scope-Block. Der Negativpfad ist der, den die Baseline als teuerste Auslassung benennt. |
| **B-7** | **MEDIUM** | `modul-04` §ID-Schema-Bezug | **[MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration) verweist auf die „Kurs-ADR-Vorlage `v1.3.0`".** Aktuell-behauptende Versionsaussage, kein historischer Hinweis — das Zweit-Review von Etappe A hatte sie als „historisch legitim" durchgehen lassen, was bei genauem Lesen nicht trägt. |
| **B-17** | **LOW** | `grundlagen-konventionen` §Source Precedence | **[MR-003](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang) ist gegenstandslos geworden.** Die Adaption hieß „Source Precedence **ohne** `docs/user`-Rang". In `v3.5.2` führt die Default-Reihenfolge `docs/user/*.md` selbst auf **Rang 5** — und `AGENTS.md` §2 listet es real auf Rang 6. Die Abweichung beschreibt heute den Normalfall; sie gehört gestrichen, nicht gepflegt. |
| **B-8** | **LOW** | `modul-13` §Vorhanden ≠ behauptet | **Maintenance-Target fehlt.** Modul 13 nennt `regelwerk-check` als Beispiel für „vorhanden, aber nicht als Gate behauptet". Genau dorthin gehört der **Freshness-Audit** aus [slice-047](../done/slice-047-baseline-vendoring.md) (Etappe-A-Fund F-6), der als Netz-Operation nicht in `gates` darf. |
| **B-20** | **LOW** | `modul-16` §Produktionsfreigabe | **`releasing.md` ist Prozedur, keine Freigabe-Checkliste.** Fünf Abschnitte (Stand · Versionsquelle · Auslösen · Konsum · Aufruf-Referenz), **null** Checklisten-Items mit Beleg-Slot, keine Anti-Item-Liste, keine Incident-Klausel. Für ein Repo, das ein digest-gepinntes Image in fremde CI-Läufe ausliefert, ist die Beleg-Pflicht die eigentlich interessante Hälfte. |
| **B-19** | **INFO** | `modul-12`, `modul-15` | **Replay und Telemetrie fehlen vollständig** — kein `evals/golden/`, kein Replay-Manifest, keine Span-/Token-Metrik. Folgewirkung: das Baseline-Closure-Kriterium „Replay-Lauf grün" (`modul-06`) ist für a-check **unerfüllbar**. Für ein deterministisches CLI-Tool ohne Agenten-Laufzeit ist der volle Apparat vermutlich unverhältnismäßig — dann gehört er als **bewusste Abweichung mit `MR-*`** deklariert, nicht stillschweigend ausgelassen. |

## 4. Negativbefunde (geprüft, ohne Befund)

- geprüft, ohne Befund: **Docker-Harness** (`modul-14`) — vorbildlich. Alle drei Fremd-Images per Digest gepinnt (`golang@sha256:…`, `golangci-lint@sha256:…`, `distroless/static-debian12:nonroot@sha256:…`), Stage-Schnitt `deps → compile/test/coverage/build → runtime`, Runtime distroless+nonroot. Erfüllt die Mindestkombination Lock-File (`go.sum`) + Image-Hash.
- geprüft, ohne Befund: **Minimal Agent Workflow** (`modul-09`) — `AGENTS.md` §6 ist Schritt für Schritt deckungsgleich mit den acht Baseline-Schritten, inklusive Schritt 8 („keine Erfolgsmeldung ohne Gate-Ausführung").
- geprüft, ohne Befund: **Source-Precedence-Form** (`modul-01`) — neun Ränge, und „neun Ränge sind ein Maximum"; Konfliktauflösungs-Klausel steht spiegelbildlich in `AGENTS.md` §1 und `harness/README.md`.
- geprüft, ohne Befund: **Doku-Konsistenz-Regel** (`modul-15`) — die eine Pflicht-Regel („keine Befehle behaupten, die es nicht gibt") ist als `make gate-consistency` real durchgesetzt, nicht nur beschrieben.
- geprüft, ohne Befund: **Review-Form** (`modul-10`) — Kategorien, Output-Schema und Negativbefund-Pflicht unverändert; die vier Reports vom 2026-07-25 sind konform. Die neue Auflage „HIGH-Liste mit ≥ 2 repo-spezifischen Regeln" erfüllt [`.harness/skills/reviewer.md`](../../../../.harness/skills/reviewer.md).
- geprüft, ohne Befund: **ADR-Hard-Rule** (`modul-04`) — „`Accepted` nie überschreiben, Korrektur = Folge-ADR mit `Supersedes`" ist in `AGENTS.md` §3.5 verkörpert **und** maschinell durchgesetzt (`doc-immutable`, Modul `vcs`, CI über die Commit-Range). Zwei reale Supersede-Ketten.
- geprüft, ohne Befund: **ADR-Form** (`modul-04`) — ≥ 3 verglichene Alternativen mit Trade-off plus Fitness Function: [ADR-0027](../../adr/0027-constructs-roh-text-monopol.md) (5), [ADR-0028](../../adr/0028-ziel-glob-schattenwurf.md) (3), [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) (4).
- geprüft, ohne Befund: **Gate-ID-Bindung** (`modul-13`) — `lint` → [ADR-0005](../../adr/0005-lint-profil.md), `coverage-gate` → [ADR-0006](../../adr/0006-coverage-gate.md), `arch-check` → [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), `image-test` → [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk), `trace-check` → [ADR-0021](../../adr/0021-commits-modul-trace-check.md).
- geprüft, ohne Befund: **Sensors-Tabelle ohne Lauf-Status** (`modul-13`, `grundlagen-konventionen`) — `harness/README.md` trägt keine Status-Spalte; Rückbau lief bereits am 2026-07-25.
- geprüft, ohne Befund: **Gate-Landschaft** — 13 Targets über sechs Sensor-Klassen; die Modul-13-Warnung „nur die generischen sechs" trifft nicht zu.
- geprüft, ohne Befund: **Traceability-Constraint** (`grundlagen-konventionen`) — Commit-Hook plus CI-Range über `make trace-check`, ID-Pflicht in `AGENTS.md` §5.
- geprüft, ohne Befund: **Spec-Stratifizierung** (`modul-03`, `grundlagen-konventionen`) — drei Straten mit `LH-`/`SPEC-`/`ARC-`-Präfixen vorhanden; „ADR schärft Spezifikation, nie das Lastenheft" ist als `AGENTS.md` §3.4/§5 verkörpert. *(Die maschinelle Hälfte — `check-references` — fehlt, siehe B-3.)*
- geprüft, ohne Befund: **Fehlannahme-Regeln** `modul-00` — keine der vier trifft auf a-check zu; das Repo führt Spec, ADRs und Sensoren statt Prompt-Wissen.

## 5. Etappen-Zuordnung

- **Etappe C** (`MR-*`-Bereinigung): **B-7** (`v1.3.0`-Aussage), **B-17** ([MR-003](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang)
  gegenstandslos),
  **B-2**-Teil (pauschaler Repo-Modus ist gegen die Baseline falsch geschnitten), **B-19**
  (bewusste Abweichung deklarieren statt auslassen) — dazu die Nummern-Kollision aus slice-046 und
  die sieben LOW aus dem Etappe-A-Zweit-Review.
- **Etappe D** (Template-/Form-Konformität): **B-1**, **B-5**, **B-6**, **B-12**, **B-15**,
  **B-18** — Slice-Vorlage mit §8 und Größen-Regel, Lerneintrag-Formen, Lifecycle inkl.
  Rückführungen, Roadmap-Drift-Log, AC-Drei-Pfad-Form, `next/` wiederherstellen oder streichen.
- **Etappe E — Mechanik (neu, aus der vollständigen Lektüre):** **B-3** (`verify`-Target inkl.
  `check-references`), **B-4** (`closure-note-reviewer`), **B-8** (`regelwerk-check` als
  Maintenance-Target, nimmt Etappe-A-Fund F-6 auf), **B-11** (`nolintlint` aktivieren),
  **B-16** (Workflow-Skelett), **B-20** (Freigabe-Checkliste mit Beleg-Slots). Das sind die Funde
  mit echtem Sensor-Gewinn — **B-11 ist der billigste** und schließt eine Hard-Rule-Lücke.
- **Etappe F — Betriebsmodell (neu):** **B-13** (Wellen-Closure-Prozedur), **B-14** (Carveout-Ort
  und -Audit), **B-21** (Steering-Loop-Kanal), **B-9** (Rollen-Übergaben), **B-10**
  (Diskrepanz-Trichter). Diese fünf hängen zusammen: sie sind die in §2 beschriebene verfallene
  Praxis, und sie lassen sich nur gemeinsam sinnvoll wiederherstellen.

Reihenfolge-Empfehlung: **E vor D**. Die Mechanik-Funde sind klein, sofort prüfbar und schaffen
die Sensoren, an denen die Form-Funde aus D danach hängen können — Form ohne Sensor ist genau die
Praxis, die laut §2 zweimal eingeschlafen ist.

## 6. DoD

- [x] Alle 17 Module und drei Grundlagen-Abschnitte gelesen; Funde mit Kategorie, Quelle und
      **gemessener** Repo-Seite belegt (§1, §3).
- [x] Negativbefunde ausgewiesen, Schweigen nicht als Konformität gewertet (§4).
- [x] Zuordnung aller Funde zu C/D/E/F mit begründeter Reihenfolge (§5).

## 7. Closure-Notiz

**Geliefert:** 21 Funde gegen die vollständig gelesene Baseline `v3.5.2` (2867 Zeilen), zwölf
Negativbefunde, vier geschnittene Folge-Etappen. Maintainer-Abnahme der Reihenfolge **E vor D**
am 2026-07-25.

**Lerneintrag — Form: geschärfte Regel.**
> **Ein geerbter Arbeits-Zuschnitt verliert seine Geltung, sobald das Ziel darüber neu gesetzt
> wird.** Der Sechs-Modul-Zuschnitt stammte aus [slice-046 §5](../open/slice-046-regelwerk-v352-migration-analyse.md)
> und war dort richtig — er entstand *vor* der Maintainer-Vorgabe „komplett nach v3.5.2
> migrieren". Ich habe ihn danach weitergetragen, statt ihn gegen das neue Ziel zu prüfen, und
> zusätzlich zwei Funde (B-9/B-10) als „gehört in keine Etappe" zur Disposition gestellt — eine
> Kategorie, die ein vollständiges Migrationsziel gar nicht kennt. Prüfregel für den nächsten
> Etappen-Start: *steht der übernommene Zuschnitt älter als die jüngste Ziel-Aussage? Dann zuerst
> den Zuschnitt neu schneiden, nicht die Arbeit beginnen.*

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` grün auf dem Stand des Slice (Exit 0) — belegt.
2. Jeder der 21 Funde trägt eine Repo-Seite mit Messung *und* eine Etappen-Zuweisung (§3, §5);
   kein Fund bleibt ohne Adressat.

**Was anders lief:** die erste Fassung wies zehn Module als ungelesene Lücke aus und stufte B-10
allein deshalb als INFO. Nach der vollständigen Lektüre wurde daraus LOW — eine Kategorie, die auf
Nichtwissen beruhte, nicht auf Beobachtung. Das ist der zweite Beleg für dieselbe Ursache wie im
Lerneintrag.

**Folge-Slices:** Etappe E (Mechanik) zuerst, darin gemäß Fund **B-1** in Slices von **≤ 3
DoD-Punkten** geschnitten statt als ein Sammel-Slice; danach D, C, F.
