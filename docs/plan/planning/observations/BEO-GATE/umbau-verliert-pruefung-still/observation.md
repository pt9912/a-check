# Ein Umbau verliert eine bestehende Prüfung, ohne dass es jemand sagt

**Sub-Area:** Gate-/Werkzeug-Schicht

Ein Prüfer oder eine Pipeline wird umgebaut — neue Struktur, neuer Ablauf —, und eine Prüfung,
die die alte Fassung trug, fehlt danach. Kein Plan, keine ADR, keine Commit-Message nennt den
Wegfall; die Tests bleiben grün, weil sie die weggefallene Eigenschaft nie einzeln gebrochen
haben. Gefunden wird es nur, wer alte und neue Fassung Eigenschaft für Eigenschaft nebeneinander
legt.
