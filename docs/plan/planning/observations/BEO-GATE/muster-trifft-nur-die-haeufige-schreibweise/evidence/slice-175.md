**Vorgang:** slice-175
**Fund:** Das `versions`-Muster für Baseline-Pins (`seit slice-173`) lautete
`\.harness/baseline/(v\d+\.\d+\.\d+)/` und traf damit **nicht** die relative Form
`../baseline/<tag>/`, die eine Datei *innerhalb* von `.harness/` benutzt.

Genau ein Vorkommen im Repo: `.harness/skills/reviewer.md` Zeile 7 verweist als
``v6.2.0` · `regelwerk/modul-10-review-harness.md``. Beim Vendoring von `v6.5.0` blieb
die Zeile stehen — **zweifach unbemerkt**: das `sed`, mit dem der Bestand umgestellt wurde, suchte
dieselbe zu enge Zeichenfolge, und der Sensor, der den vergessenen Nachzug melden soll, sah sie
ebenso wenig. Dieselbe Verengung an beiden Stellen, weil dieselbe Vorstellung dahinterstand.

Aufgefallen ist es an einer **Inkonsistenz im selben Satz**: Nach dem `sed` stand dort
*„Regelwerk **v6.2.0** Modul 10 (vendored: `…/v6.5.0/…`)"* — ein Teil gehoben, der andere nicht.

**Behoben** in diesem Slice: Muster auf `(?:^|[^-a-zA-Z0-9_.])baseline/(v\d+\.\d+\.\d+)/`
erweitert; Mutations-Probe belegt beide Richtungen (`reviewer.md:7 v6.4.0 version-stale` bei
gesetztem Fehler, 0 Befunde nach Rücknahme).

**Was das über den ersten Migrations-Lauf sagt:** Der Sensor wurde in
[slice-173](../../../../done/wellenlos/slice-173-versions-sensor-baseline-pins.md) mit einer
Prüfmenge von 14 Dateien belegt und in beide Richtungen mutations-kalibriert — trotzdem hatte er
eine Form-Lücke. Kalibrierung an vorhandenen Vorkommen prüft, ob der Sensor *diese* findet, nicht
ob er *alle* Schreibweisen kennt.
