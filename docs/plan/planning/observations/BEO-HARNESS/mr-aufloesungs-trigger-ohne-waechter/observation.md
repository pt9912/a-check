# Der Auflösungs-Trigger eines `MR`-Eintrags hat keinen Wächter

**Sub-Area:** Harness-Einstieg

Jeder Adaptions-Eintrag trägt ein Pflichtfeld `Auflösungs-Trigger`. Prüft ihn niemand, ist er eine
Absichtserklärung ohne Verfallsdatum: der Eintrag bleibt aktiv, obwohl seine Bedingung eingetreten
ist, und wird nur gefunden, wenn jemand die Einträge von Hand durchgeht.

Der **Trigger-Audit** der Wellen-Closure
([`modul-06`](../../../../../../.harness/baseline/v6.2.0/regelwerk/modul-06-roadmap.md#wellen-closure-prozedur-modul-6),
Schritt 2) prüft drei Artefaktklassen — Carveout, bootstrap-aware Gate, ADR. `MR`-Einträge stehen
nicht darin, und die Slice-Closure kennt den Schritt in dieser Form ebenfalls nicht.

**Abgrenzung zu [`rueckbau-kandidat-ueberlebt-baseline-migration`](../rueckbau-kandidat-ueberlebt-baseline-migration/observation.md):**
dort wird ein Trigger *gewählt* — der billigere Ersatz statt der sauberen Auflösung; hier wird
keiner gewählt, weil niemand hinsieht. Andere Ursache, anderes Gegenmittel: ein generischer
Verweis ([`MR-020`](../../../../../../harness/conventions/MR-020-adr-vorlage-generisch.md)) hilft
gegen diese Ausprägung nicht.
