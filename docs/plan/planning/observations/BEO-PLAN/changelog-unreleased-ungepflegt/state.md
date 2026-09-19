**Stand:** verkörpert in [`AGENTS.md`](../../../../../../AGENTS.md) §6 Schritt 7
(*„Die CHANGELOG-Zeile gehört hierher — nicht in die Release-Vorbereitung"*) `seit slice-196`.

Drei Ausprägungen, alle drei bei der **Release-Vorbereitung** gefunden: slice-127 (über zwölf Slices
ein einziger Eintrag, darunter eine Breaking Change) · slice-133 (21 Commits seit `v0.18.0` ohne
Eintrag) · slice-196 (58 Slices und 265 Commits seit `v0.19.0`, davon **eine** konsumenten-sichtbare
Änderung).

**Kein Sensor**, und das ist benannt statt verschwiegen: Eine Prüfung auf „hat einen Eintrag"
bräuchte als Eingabe das Urteil, ob eine Änderung *eintragspflichtig* ist — `spec/` und die Regeln
ja, ein Sensor-Umbau nein. Sie wäre auf jedem Harness-Slice falsch-positiv oder abgeschaltet. Der
Fehler ist kein Wissens-, sondern ein **Zeitpunkt**-Fehler; die geschärfte Regel bindet den Eintrag
an den Slice, in dem der Vertrag sich ändert.
