**Vorgang:** slice-214

**Fund (Review F-1, F-3, F-5, F-8, F-9):** Der Dialekt `json` nahm `tru`, `01`, `"\x"`, ein
fehlendes Komma, ein abschließendes Komma in einem verschachtelten Objekt und ungültiges UTF-8 an —
die Spezifikation zählte nur einzelne Fehlerfälle auf. CHANGELOG und Handbuch sagten mehr zu.
Fünf Funde, **ein** Vorgang. **Behoben** als Plan-Änderung vor dem Code: die Datei muss gültiges
JSON nach RFC 8259 sein (Spezifikation 0.38.0), ein Grammatik-Prüfer über die Tokens, Tests je
Form mit Mutations-Gegenprobe.
