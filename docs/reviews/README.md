# Review-Synthesen

Ein Review-Dokument je Slice mit Code- oder Vertragsänderung. Form nach Regelwerk Modul 10
(Stand siehe [`harness/conventions.md` §Baseline](../../harness/conventions.md#baseline)):

- **Kategorien** `HIGH` / `MEDIUM` / `LOW` / `INFO`.
- **Output-Schema je Finding:** Quelle · Pfad · Befund (beobachtbar, ohne Lösungsvorschlag) ·
  `verifizierbar` (gibt es einen Gate-Lauf, der es bestätigen würde?).
- **Negativbefund-Zeilen** je betrachtetem Bereich („geprüft, ohne Befund") — ohne sie sehen
  „keine Findings in X" und „X nicht angesehen" identisch aus.
- **Abgrenzung zur Verifikation** (Modul 11): eine DoD-Verletzung ist **kein** Review-Finding.
  Der Reviewer prüft gegen Plan/ADR/Konventionen, der Verifier gegen DoD/Spec. Die
  DoD-Verifikation steht darum im Slice-Dokument, nicht hier.

## Zwei bewusste Abweichungen

**1. Kategorien-Drift der Altdokumente.** Die Synthesen bis `2026-07-04-slice-026-phantom-guard.md`
nutzen `BLOCKER` / `MINOR` / `NIT`. Ab `2026-07-25` gelten die Regelwerks-Kategorien. Die
Altdokumente werden **nicht** umgeschrieben — sie sind Belege ihres Zeitpunkts. Als
`MR-*`-Adaption ist die alte Kategorien-Menge nie deklariert worden; wer sie wiederbeleben will,
deklariert sie in [`harness/conventions.md`](../../harness/conventions.md).

**2. Lücke slice-027 … slice-041 — bewusst nicht nachgezogen.** Für diese Slices existiert kein
Review-Dokument; die Reviews fanden statt (die Befunde stehen in den jeweiligen Slice-Notizen und
in den Commit-Messages), wurden aber nicht als Synthese abgelegt. Entscheidung des Maintainers am
**2026-07-25**: **nicht rückwirkend nachtragen** — ein nachträglich rekonstruiertes Review wäre
kein Beleg, sondern eine Nacherzählung, und der Wert eines Review-Dokuments liegt in der
Gleichzeitigkeit mit dem Diff. Nachgetragen wurden ausnahmsweise
[slice-042](2026-07-25-slice-042-constructs-monopol.md) und
[slice-044](2026-07-25-slice-044-ziel-glob-schattenwurf.md), weil ihre Reviews **am selben Tag**
liefen und die Befunde unmittelbar vorlagen.

Ab [slice-043](2026-07-25-slice-043-abdeckungs-diagnose.md) entsteht die Synthese wieder **vor dem
Merge**, und der Slice-DoD führt sie als eigenen Punkt.
