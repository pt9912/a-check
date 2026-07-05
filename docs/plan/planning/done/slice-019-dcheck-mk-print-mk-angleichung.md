# slice-019 — `d-check.mk` an `--print-mk` angleichen (Verbatim + Pin)

**Status:** done (2026-07-04 — Weg 3 umgesetzt gegen **v0.37.1**; [Closure](#5-closure-notiz)).
**Bezug:** Reproduzierbarkeit **sinngemäß** [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)
(`d-check.mk` ist eine **konsumierte Dev-Abhängigkeit**, nicht a-checks eigenes Distributions-Artefakt —
daher per Analogie, **nicht direkt** gebunden); `AGENTS.md` §4 ↔ `tools/gate-consistency.sh`
(Target-Invariante); Koordination mit [slice-018](../done/slice-018-versions-register-pin-gate.md)
(Pin-Drift-Gate). [Roadmap](../in-progress/roadmap.md). **Evidenz:** `d-check.mk` weicht vom
v0.35.0-`--print-mk` ab.

## 1. Auslöser & Trade-off

Das committete `d-check.mk` ist von einem älteren `--print-mk` abgeleitet und **bewusst getrimmt** —
`d-check.mk:8-10` sagt es ausdrücklich: die 5 Targets sind weggelassen, „weil jedes reale
d-check.mk-Target sonst in AGENTS §4 stehen müsste (gate-consistency)". Der Status quo ist also
**keine Nachlässigkeit, sondern eine dokumentierte Entscheidung** (schlanke, gate-zentrische
AGENTS-§4-Fläche). Dieser Slice **kehrt sie um** — die eigentliche Review-Frage ist der Trade:

- **Pro Weg 3 (verbatim):** „Erzeugt aus `--print-mk`" wird wieder **wörtlich** wahr; ein Bump ist ohne
  Sonderwissen regenerierbar; die vollen d-check-Fähigkeiten (`doc-doctor`/`-repair`/`-immutable`/
  `-commits`) stehen offen.
- **Kosten:** 5 advisory, **gate-lose** Tools in AGENTS §4 verwässern das „Nur hier gelistete Targets
  existieren"-Signal (die **Stand: advisory**-Spalte mildert, hebt nicht auf); dazu die Restkost aus
  §3 (kein Regressions-Gate).
- **Entscheidung (Maintainer):** **Verbatim-Treue vor schlanker Fläche** — die dokumentierte
  Regenerierbarkeit ist den Briefing-Ballast wert; die `advisory`-Spalte + diese Begründung wahren die
  AGENTS-§4-Ehrlichkeit.

## 2. Geplanter Umfang (Weg 3 — verbatim + lokaler Pin)

1. **`d-check.mk` = die v0.35.0-`--print-mk`-Ausgabe**, mit **einer** lokalen Anpassung: der
   `DCHECK_DIGEST`-Hook wird auf `sha256:9d7b23ac…` **gesetzt** (Digest sticht Tag → Reproduzierbarkeit
   **sinngemäß** [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)). Die Ausgabestruktur (`DCHECK_IMAGE ?= …:v0.35.0` + `DCHECK_DIGEST`/
   `DCHECK_REF`-Logik + `## `-Help-Annotationen) ist **aus einem v0.35.0-`--print-mk`-Lauf beobachtet**
   (in dieser Session); **beim Umsetzungs-Lauf final gegen das dann gepinnte `--print-mk` verifizieren**
   — weicht es ab (z. B. keine `## `-Annotationen), ist das eine **zweite** Anpassung und offen
   auszuweisen (sonst ist „verbatim" selbst nur „lose wahr").
2. **Alle 8 Targets** übernehmen (`doc-check`/`-trace`/`-complete`/`-doctor`/`-repair`/`-immutable`/
   `-commits`/`-help`).
3. **`AGENTS.md` §4** um die 5 neuen Targets erweitern — **als Pipe-Tabellenzeilen**
   (`| `​`make doc-doctor`​` | … | advisory |`), **nicht** als Fließtext: `gate-consistency`s
   `doc_targets()` grept `^\| … make X` (`tools/gate-consistency.sh`) und sähe Fließtext nicht.
   `doc-help` ist **kein** Utility (`help/build/compile/hooks`) → ebenfalls gelistet.
4. **`harness/README.md §Sensors` unverändert** — die 5 sind advisory, keine Feedback-Gates;
   Präzedenz `doc-trace`/`-complete` (schon nur in AGENTS §4). Trennung im Commit **begründen**.
5. **`DCHECK_DIGEST`-Setzung:** verbatim `?=` (override-bar, konsistent mit `DCHECK_IMAGE ?=`) —
   **oder** hart `:=` (strikter Pin, aber eine **zweite** Verbatim-Abweichung). Default: `?=` mit
   Notiz „override-bar" (§3).
6. **Smoke-Check** der 5 (kein Gate deckt sie ab): `doc-doctor`/`-repair`/`-help` ohne Argument;
   `doc-immutable`/`-commits` mit `RANGE=HEAD~1..HEAD` → Recipe/Flags fehlerfrei. DoD-Kriterium.
7. `gates`-Target **unverändert**; Kommentar entschärfen (verbatim + `DCHECK_DIGEST`-Pin);
   `make gates` grün.

## 3. Vor der Umsetzung zu klären

- **Koordination mit [slice-018](../done/slice-018-versions-register-pin-gate.md) (Pin-Drift-Gate):** Weg 3
  **spaltet** den heute driftfreien d-check-Pin (eine Digest-Koordinate) in **Tag + Digest**, die
  gegeneinander driften können (Tag sagt `v0.35.0`, Digest evtl. von woanders) — **genau** die Klasse,
  für die slice-018 `versions`/`pins` einführt. slice-018 listet `d-check.mk` **noch nicht** als
  Pin-Quelle. **Zu entscheiden:** slice-018s Gate deckt nach slice-019 auch `d-check.mk`s
  Tag↔Digest-Konsistenz ab **oder** exemptiert es explizit — sonst schafft slice-019 eine Drift-Quelle,
  die das dafür gebaute Gate nicht sieht. (slice-018 §1 um diesen Punkt ergänzt.)
- **`?=` vs `:=`** für `DCHECK_DIGEST` (§2.5): Verbatim/override-bar vs. harter Reproduzierbarkeits-Pin.
- **Restkost (akzeptiert):** nach dem einmaligen Smoke schützt **kein Gate** die advisory Targets vor
  stillem Verrotten bei künftigen Image-Bumps (geänderte Flags) — dieselbe Klasse, die sie ursprünglich
  draußen hielt. Bewusst getragene Restkost von Weg 3 (§1-Trade).

## 5. Closure-Notiz

Umgesetzt gegen **v0.37.1** (nicht v0.35.0 wie im Plan skizziert — der Pin-Stand ist
inzwischen v0.37.1):

- **`d-check.mk`** = `d-check:v0.37.1 --print-mk` **verbatim**, mit der einen Anpassung
  `DCHECK_DIGEST ?= sha256:3bbdb19b…` (Digest sticht Tag → Reproduzierbarkeit sinngemäß
  [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)); `?=` override-bar
  (Plan §2.5-Default). Struktur (`DCHECK_IMAGE`-Tag + `DCHECK_DIGEST`/`DCHECK_REF` + `## `-Help)
  gegen den gepinnten `--print-mk`-Lauf verifiziert — **keine** zweite Abweichung.
- **Delta zum Plan (§2.2):** v0.37.1 liefert **10** `doc-*`-Targets, nicht 8 — seit v0.35.0
  neu: **`doc-planning`** (DC-FA-PLAN-001) und **`doc-tracked`** (DC-FA-TRK-001). Alle 10 in
  [`AGENTS.md` §4](../../../../AGENTS.md#4-quality-gates) als **advisory** gelistet
  (Pipe-Zeilen — `gate-consistency` grün).
- **Smoke-Check (§2.6):** `doc-help`/`-doctor`/`-repair`/`-immutable`/`-commits`/`-planning`/
  `-tracked` — Recipes/Flags fehlerfrei (rc 0/1, kein Usage-Fehler).
- **`harness/README.md §Sensors` unverändert** (advisory ≠ Feedback-Gate, Plan §2.4).
- **`make gates` grün**; CHANGELOG-`[Unreleased]`-Notiz (Dev-Tooling, kein Image-Release nötig).
- **Offen (slice-018-Koordination, §3):** der d-check-Pin ist jetzt **Tag + Digest** (Drift-Klasse);
  [slice-018](../done/slice-018-versions-register-pin-gate.md)s `versions`/`pins`-Gate muss `d-check.mk`s
  Tag↔Digest-Konsistenz abdecken **oder** explizit exemptieren — sonst ist hier eine Drift-Quelle,
  die das dafür gebaute Gate nicht sieht.
