# Review-DoD-Haken gesetzt, bevor der Review lief

**Sub-Area:** Gate-/Werkzeug-Schicht

Ein Slice trägt den Haken `- [x] Unabhängiger Review durchgeführt`, während in `docs/reviews/`
kein Report gleicher Kennung liegt. Der Haken ist **selbst-attestiert**: `make doc-reviews`
prüft die Deckung erst für Slices in `done/` — solange der Slice in `in-progress/` liegt, ist
ein falscher Haken von keinem Lauf gedeckt.

Die Lücke ist nicht die Zusage, sondern ihr **Zeitpunkt**: Der Sensor greift am Lifecycle-Ende,
die Attestierung entsteht in der Mitte. Wer den Haken beim Schreiben der Implementierung setzt
und den Review danach vergisst, bemerkt es erst beim `git mv` — oder gar nicht, wenn der Haken
den Blick auf die offene Aufgabe verstellt.
