# MR-003 — Source Precedence ohne `docs/user`-Rang

- **Datum:** 2026-06-20
- **Geltungsbereich:** [`harness/README.md` §Source precedence](../../README.md#source-precedence),
  [`AGENTS.md` §2](../../../AGENTS.md#2-kanonische-quellen-source-precedence)
- **Adaption:** Der Template-Default führt neun Ränge inkl. eines Rangs
  `docs/user/*` (Operations, Quality, Releasing); a-check führt acht
  Ränge ohne `docs/user`, weil noch kein Operations-Doku-Stratum
  existiert (CLI-Tool vor dem ersten Release).
- **Begründung:** Ein Rang für nicht existierende Dateien wäre ein
  halluzinierter Eintrag (gleiche Klasse wie ein behauptetes Gate); die
  Rangordnung ist laut Baseline projektspezifische Wahl, die hier
  deklariert wird.
- **Auflösungs-Trigger:** mit der Release-Pipeline entsteht Betriebs-/
  Releasing-Doku; der `docs/user`-Rang wird dann eingefügt und dieser
  Eintrag als aufgelöst markiert.
- **Aufgelöst (2026-06-21):** Mit dem
  [Benutzerhandbuch](../../../docs/user/benutzerhandbuch.md) existiert Nutzer-/
  Betriebs-Doku; der `docs/user`-Rang ist als Rang 6 (vor `README.md`)
  eingefügt — die Source Precedence führt jetzt neun Ränge.
