# Ein DoD-Punkt attestiert einen Vorgang, der noch nicht stattgefunden hat

**Sub-Area:** Gate-/Werkzeug-Schicht

Ein Slice-Plan trägt eine Aussage in der Vergangenheitsform über einen Schritt, der noch
aussteht — der Haken `- [x] Unabhängiger Review durchgeführt` ohne Report, die Zeile
„nach dem `git mv` geprüft" vor dem `git mv`. Die Aussage ist **selbst-attestiert**: Sie
beschreibt nicht den Zustand, sondern die Absicht, und niemand unterscheidet die beiden beim
Lesen.

Die Lücke ist nicht die Zusage, sondern ihr **Zeitpunkt**. Die Sensoren, die solche Zusagen
decken, greifen am Lifecycle-**Ende**: `make doc-reviews` prüft die Report-Deckung erst für
Slices in `done/`, die drei Paarungen prüfen nach dem `git mv`. Die Attestierung entsteht davor.
Wer den Haken beim Schreiben setzt, hat den ausstehenden Schritt danach nicht mehr im Blick —
der Haken verstellt ihn.
