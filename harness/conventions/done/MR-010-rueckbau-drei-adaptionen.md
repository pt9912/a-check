# MR-010 — Rückbau: drei Adaptionen, die die Baseline übernommen hat

- **Status:** Accepted
- **Datum:** 2026-08-29
- **Geltungsbereich:** der Adaptions-Block selbst; keine Repo-Datei ändert ihr Verhalten
- **Ersetzt-Baseline-Regel:** — *(keine; dieser Eintrag **löst** Adaptionen auf, er setzt keine.
  Damit ist er kein Fork im Sinne der Baseline, sondern deren Gegenteil: eine Rückkehr zum Default)*
- **Adaption:** **keine.** Das Repo folgt in allen drei Punkten dem Baseline-Default.
- **Begründung:** je Eintrag der Beleg aus dem Durchgang der Etappe B
  ([slice-095 §3](../../../docs/plan/planning/done/wellenlos/slice-095-adaptions-durchgang-v5120.md)):
  - [`MR-001`](../../conventions.md#mr-001) — `grundlagen-source-precedence.md` führt `spec/spezifikation.md` **selbst**
    als Rang 2 von neun. Die Abweichung „drei statt zwei Spec-Ränge" existiert nicht mehr.
  - [`MR-002`](../../conventions.md#mr-002) — §ID-Schema als Klammer: *„`LH-FA-03` und `LH-FA-IDX-003` sind **beide**
    wohlgeformt."* Bereichskürzel sind Default-konform, das Vertrags-Präfix ist frei wählbar.
  - [`MR-006`](../../conventions.md#mr-006) — `modul-02` Schritt 2 macht das committete Vendoring zum
    Bootstrap-**Default**. Der Eintrag sagte selbst, er weiche nicht ab, sondern *stelle den
    Default her*; nach dem Fork-Test der neuen Baseline war er nie eine Adaption.
- **Auflösungs-Trigger:** entfällt — der Eintrag ist mit seiner Entstehung erledigt und liegt
  darum von Beginn an in `done/`. Eine Nicht-Abweichung gehört nicht in die Tabelle, die jeder
  Agentenlauf liest.
- **Löst auf:** [`MR-001`](../../conventions.md#mr-001), [`MR-002`](../../conventions.md#mr-002), [`MR-006`](../../conventions.md#mr-006)
- **Ausgelöst durch Baseline-Stand:** `v5.12.0`
