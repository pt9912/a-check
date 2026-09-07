**Vorgang:** slice-175
**Fund:** Die in `evidence/slice-174.md` benannte Form-Lücke ist **beide Male** wieder aufgetreten,
die dieser Slice den Lifecycle durchlief — `open/` → `in-progress/` und `in-progress/` → `done/`.
Zwei Funde, ein Vorgang, ein Beleg (Baseline `modul-06`).

`make slice-mv` zog jeweils die Verweise in den Lifecycle-Verzeichnissen nach und ließ den in der
**flach** liegenden Welle-Datei stehen: `[slice-175](open/…)` bzw. `[slice-175](in-progress/…)` —
Verzeichnis-Präfix ohne `../`, weil `welle-15-…md` eine Ebene über den Lifecycle-Verzeichnissen
liegt. `make doc-check` meldete beide Male `target-missing`, die Korrektur war ein `sed`.

**Die Vorhersage aus `slice-174` ist eingetroffen** — dort stand: *„welle-15 fährt noch vier
Slice-Übergänge und sieht die Form damit wieder."* Nach diesem Slice sind es **drei** Vorkommen in
zwei Vorgängen; bei den verbleibenden Etappen (C, D, E) kommen je zwei dazu, also **sechs
weitere**, solange das Werkzeug die Form nicht lernt.

**Was das über den Ausgang sagt:** Der Eintrag steht auf *verkörpert* — und das bleibt richtig, die
Verkörperung fängt die Formen, für die sie gebaut wurde. Sie ist nur an einem Betriebsmodus
kalibriert, den a-check bis `welle-15` nicht hatte: eine **offene Welle-Datei**, die auf Slices in
Lifecycle-Verzeichnissen zeigt. `welle-14` lag bei den Übergängen ihrer Slices bereits in `done/`,
und wellenlose Slices haben keine Welle-Datei. Dieselbe Klasse wie
[`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../../../BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md):
kalibriert an den Vorkommen, die es beim Bauen gab.
