# `make version-coherence` — Kohärenz doppelt deklarierter Versions-Angaben

## Vertrag

Zwei Regeln, beide hermetisch: Derselbe `uses:`-SHA trägt unter
`.github/workflows/` überall denselben Tag-Kommentar. Und eine
Versions-Variable, die `Makefile` **und** `Dockerfile` führen, hat an beiden
Orten denselben Wert.

## Grenze — was das Grün nicht abdeckt

1. **Die Wahrheit einer Angabe** — geprüft wird **Divergenz, nicht Unwahrheit**.
   Zwei übereinstimmend *falsche* Angaben bleiben grün. Permanent: Die Registry
   zu fragen wäre Netz, und `gates` ist hermetisch.
2. **Welche Seite recht hat** — der Sensor erklärt **keine** zur führenden. Er
   meldet, dass zwei Orte auseinanderlaufen, und überlässt das Urteil dem
   Leser. Permanent, und Absicht.
3. **Einfach deklarierte Angaben** — eine Version, die nur an einem Ort steht,
   hat keinen Partner und fällt aus der Prüfmenge. Für Baseline-Pins schließt
   `versions` in `doc-check` diese Lücke, für andere niemand.

## Bindung

slice-131 · Antwort auf
[`BEO-GATE/versionsangabe-neben-digest-ungeprueft`](../../docs/plan/planning/observations/BEO-GATE/versionsangabe-neben-digest-ungeprueft/observation.md)
bei 3× · im `gates`-Aggregat.
