# MR-019 — Review-DoD-Punkt bleibt Opt-in statt verpflichtend

- **Status:** Accepted
- **Datum:** 2026-09-06
- **Geltungsbereich:** [`AGENTS.md`](../../AGENTS.md) §5 (Slice-Form, DoD-Posten-Liste),
  [`.d-check.yml`](../../.d-check.yml) (`structure`, `tasks-ignore-pattern`)
- **Ersetzt-Baseline-Regel:** — *(keine. Der Treiber ist `lab/templates/docs/plan/planning/slice.template.md`
  in seiner `v6.2.0`-Fassung — ein **Template**, keine Regelwerk-Regel, und dazu eines, das a-check
  noch nicht vendored hat; kein netzlos auflösbarer Anker möglich. Derselbe Fall wie
  [`MR-017`](../conventions.md#mr-017)/[`MR-018`](../conventions.md#mr-018): Rückbau-Kandidat, sobald
  `v6.2.0` vendored ist und auf das dann vendorte Template gezeigt werden kann)*.
- **Adaption:** `v6.2.0` macht den Review-Report zu einem unbedingten DoD-Checkbox-Punkt jedes
  Slice-Plans (*„Review durchgeführt, Report … liegt vor"*) und führt ihn als vierten pro Slice
  konstanten, nicht gezählten Posten in `modul-05-planning-harness.md`. a-check übernimmt die
  zweite Hälfte (Review-Report zählt nicht zu den drei Liefer-Punkten, `AGENTS.md` §5), aber **nicht**
  die erste: der Review-Report bleibt **Opt-in** über `make doc-reviews`/`DC-FA-RVW-001`, ausgelöst
  durch die exakte Phrase „unabhängiger Review" in einer DoD-Zeile.
- **Begründung:** `AGENTS.md` §6 verlangt Review bereits präziser als die Baseline-Vorlage — *„bei
  jedem Slice mit Code- oder Vertragsänderung"*, nicht bei jedem Slice. Ein unbedingter Checkbox-Punkt
  triggerte entweder auch reine Analyse-Slices unnötig, oder er müsste selbst wieder bedingt formuliert
  werden — dann ist er kein Baseline-Import mehr, sondern a-checks bestehender Opt-in-Mechanismus.
  **Zusätzlicher, empirisch gemessener Befund** ([slice-165](../../docs/plan/planning/in-progress/slice-165-review-dod-punkt-opt-in-beibehalten.md)
  §3/§4): weder der `v6.2.0`-Baseline-Wortlaut noch a-checks eigene, tatsächlich in `slice-161`–`164`
  verwendete Formulierung („Unabhängiges Plan-Review …") lösen die Trigger-Phrase aus — nur die exakte
  Form „unabhängig**er** Review" (isoliert gegen den gepinnten `d-check`-Digest getestet, nicht nur
  dokumentiert). Diese Adaption hält fest, welche Formulierung tatsächlich funktioniert, damit künftige
  Slices nicht denselben Fehlschlag wiederholen.
- **Auflösungs-Trigger:** die Überarbeitung des Review-Mechanismus selbst (Umstieg auf einen
  unbedingten Checkbox-Punkt) — oder die nächste Baseline-Migration, wenn sie das Template erneut
  ändert und diese Adaption dadurch gegenstandslos wird.
- **Ausgelöst durch Baseline-Stand:** `v6.2.0` (Kurs-Welle 119, 2026-09-05); vendored bleibt zum
  Zeitpunkt dieses Eintrags `v6.0.0`.
