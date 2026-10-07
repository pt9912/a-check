**Stand:** verkörpert — als **benannte Spec-Lücke** in der Spezifikation 0.36.0, §[SPEC-RULE-001](../../../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
(Ausgabe-Regel: die Meldung jedes `shape-*`-Befunds ist einzeilig und umkehrbar; `shape-unused`
trägt die Art als Präfix). Eine Spec-Stelle trägt keinen Herkunfts-Anker; ihr Gegenstück ist die
Versions-Zeile 0.36.0 der Spezifikation.

Drei Belege: slice-210 (mehrzeiliger Roh-String in `shape-unlisted`, gelassen), slice-211
(`shape-unused` mit rohem CR, nur dort festgelegt), slice-212 (nicht umkehrbares `\n`, Zusatz
` (regex)` mehrdeutig) — der dritte hat die Regel für alle Befunde ausgelöst. Die Implementierung
liefert slice-213.
