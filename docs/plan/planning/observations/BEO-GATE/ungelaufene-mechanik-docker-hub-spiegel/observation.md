# Ungelaufene Mechanik: Docker-Hub-Spiegel

**Sub-Area:** Gate-/Werkzeug-Schicht
**Ehemals:** `BEO-031` (Tabellenform, bis slice-139)

Eine geschriebene, aber **ungelaufene** Mechanik ist eine Zusage ohne Beleg — beim Docker-Hub-Spiegel wurde der Fehler dadurch erst beim Handlauf sichtbar: die Config-Digest-Extraktion lieferte **leer**, weil `docker manifest inspect` mehrzeiliges JSON gibt und das übernommene `sed`-Muster einzeilig matcht
