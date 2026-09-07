# Der Auflösungs-Trigger eines `MR`-Eintrags hat keinen Wächter

**Sub-Area:** Harness-Einstieg

Jeder Adaptions-Eintrag trägt ein Pflichtfeld `Auflösungs-Trigger`. Prüft ihn niemand, ist er eine
Absichtserklärung ohne Verfallsdatum: der Eintrag bleibt aktiv, obwohl seine Bedingung eingetreten
ist, und wird nur gefunden, wenn jemand die Einträge von Hand durchgeht.

Zwei Lücken übereinander, und die zweite ist die nähere:

1. Der **Trigger-Audit** der Baseline
   (`modul-06` (`v6.2.0` · `regelwerk/modul-06-roadmap.md` §Wellen-Closure-Prozedur),
   Closure-Schritt 2, im wellenlosen Betrieb getragen von der **Slice**-Closure) zählt drei
   Artefaktklassen auf — Carveout, bootstrap-aware Gate, ADR. Die vierte Klasse, die im selben
   Regelwerk ein Pflichtfeld `Auflösungs-Trigger` führt, steht nicht darin: der `MR`-Eintrag.
2. a-check hat den Schritt **für keine** der drei genannten Klassen verkörpert: `Trigger-Audit`
   kommt in [`AGENTS.md`](../../../../../../AGENTS.md),
   [`harness/README.md`](../../../../../../harness/README.md) und
   [`docs/plan/planning/README.md`](../../../README.md) null Mal vor.

Die erste Lücke erklärt, warum niemand an `MR`-Einträge denkt; die zweite, warum auch die drei
genannten Klassen ungeprüft blieben.

**Abgrenzung zu [`rueckbau-kandidat-ueberlebt-baseline-migration`](../rueckbau-kandidat-ueberlebt-baseline-migration/observation.md):**
dort wird ein Trigger *gewählt* — der billigere Ersatz statt der sauberen Auflösung; hier wird
keiner gewählt, weil niemand hinsieht. Andere Ursache, anderes Gegenmittel: ein generischer
Verweis ([`MR-020`](../../../../../../harness/conventions.md#mr-020)) hilft
gegen diese Ausprägung nicht.
