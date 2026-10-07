**Vorgang:** slice-214

**Fund (Delta-Review D-4):** `jsonValidate` bricht bei einer gültigen Schachtelungstiefe von etwa
3·10^6 (rund 6 MB) mit `fatal error: stack overflow` ab; Tiefe 10^6 läuft durch. RFC 8259 §9
erlaubt eine Grenze, Spezifikation und Handbuch nennen keine. **Nicht behoben:** eine
Tiefengrenze ist eine Vertragsänderung (Frage an den Architect); reale Manifeste liegen um
Größenordnungen darunter (233 `package.json` gemessen, keine nahe der Grenze).
