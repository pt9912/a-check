**Stand:** verkörpert in [`make slice-mv`](../../../../../../Makefile)
(`tools/slice-mv.sh`) `seit slice-118`, **Lücke geschlossen** `seit slice-180`.

Die Verkörperung deckte bis slice-180 **zwei** der drei im Bestand vorkommenden Verweis-Formen.
Die dritte —
`<lifecycle-verzeichnis>/slice-NNN-….md` **ohne** `../`-Präfix, wie sie eine flach unter
`docs/plan/planning/` liegende Welle-Datei schreibt — kannte sie nicht; seit slice-180 erkennt das Werkzeug sie an ihren zwei Ankern (Link-Ziel,
Inline-Code), und Auswahl wie Ersetzung teilen dafür **eine** Muster-Quelle.

**Warum das erst mit `welle-15` auffiel:** `welle-14` lag bei den Übergängen ihrer Slices bereits
in `done/`, ihre Verweise trugen `../done/…`; wellenlose Slices haben gar keine Welle-Datei, die
auf sie zeigt. Erst eine **offene** Welle-Datei erzeugt die Form. Dieselbe Klasse wie
[`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../../BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md):
kalibriert an den Vorkommen, die es beim Bauen gab.

**Zähler-Stand 7×** — vier davon aus `welle-15` (slice-174 · slice-175 · slice-176 · slice-177),
in der die Lücke bei **jedem** der zehn Lifecycle-Wechsel auftrat. `make doc-check` meldete sie
jedes Mal; kein stiller Ausfall, aber zehnmal Handarbeit.
