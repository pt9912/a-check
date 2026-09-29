# MR-030 — Keine Agenten-Telemetrie nach draußen (löst [`MR-014`](../conventions.md#mr-014) auf)

- **Status:** Accepted
- **Datum:** 2026-09-29
- **Geltungsbereich:** gesamtes Repo; betrifft das Baseline-Modul `modul-15`
- **Ersetzt-Baseline-Regel:** [`modul-15-observability.md` §Kernidee](../../.harness/baseline/v6.13.0/regelwerk/modul-15-observability.md#kernidee-modul-15)
- **Adaption:** a-check schreibt **keine** Agenten-Telemetrie **nach draußen** —
  keine Netz-Sendung, kein Abfluss des Bestands. **Lokale Lauf-Erfassung ist
  zulässig:** die Erfassungsschicht
  ([`.claude/hooks/span-emit.sh`](../../.claude/hooks/span-emit.sh) + Traeger +
  [`harness/erfassung-feldliste.md`](../../harness/erfassung-feldliste.md),
  slice-206) schreibt in den gitignorierten Zustands-Bereich unterhalb von
  `.harness/`. Ihr Schema ist geschlossen
  ([`harness/erfassung-feldliste.md`](../../harness/erfassung-feldliste.md)); der
  Bestand ist **nicht verschlüsselt** und **nicht zugriffsbeschränkt**.
- **Begründung:** [`MR-014`](../conventions.md#mr-014) verbot Tool-Call-Spans
  uneingeschränkt — gerechtfertigt aus dem Observability-Modell des Baseline-Moduls,
  nicht aus der Lage dieses Repos. MR-014s eigener Auflösungs-Trigger („sobald
  Agenten-Läufe **im Repo selbst** abrechenbar werden") ist mit der Token-Erfassung
  dieser Schicht eingetreten; die Maintainer-Freigabe (2026-09-29) stellt die Zusage
  auf die Lokal/Draußen-Grenze um: beobachten ja, abfließen lassen nein.
- **Löst auf:** [`MR-014`](../conventions.md#mr-014)
- **Ausgelöst durch Baseline-Stand:** `v6.13.0`
- **Auflösungs-Trigger:** permanent.
