**Vorgang:** slice-192

**Fund:** Der Baseline-Sprung wurde mit `sed` über alle Dateien ausgerollt, die den alten Stand
trugen. Darunter waren zwei, in denen er **historisch** stand:

- der **Migrationsplan selbst** — sieben Stellen. Danach behauptete §1 *„`v6.6.0` ist entfernt"*,
  §2 hieß *„Delta-Analyse `v6.6.0` → `v6.6.0`"*, und der dort als Instrument genannte Aufruf
  `git diff v6.6.0 v6.6.0` liefert leere Ausgabe — während *„10 Dateien, +87/−32"* als sein
  Ergebnis danebenstand.
- das **Closure-Log der Roadmap** — die Zeile zu `welle-15` behauptete eine Migration
  `v6.2.0`→`v6.6.0` und widersprach damit der Ergebnisnotiz, auf die sie zeigt, und der
  Welle-Kennung in ihrer eigenen Nachbarzelle.

**Kein Lauf fängt es.** Das Modul `versions` bindet an das Pfad-Muster
`<baseline-verzeichnis>/<tag>/`; eine nackte Kennung im Fließtext trifft es nicht — zu Recht, denn
dort ist sie meistens eine Tatsache. `make doc-check` sieht nur Links.

**Gefunden hat es der unabhängige Review** (Report zu slice-192, F-1 und F-2), nicht der
Schreibende und kein Gate. Behoben sind alle sieben Stellen plus die Roadmap-Zeile.

**Die billige Vorbeugung stand die ganze Zeit da:** Ersetzt man mit dem **Sensor-Muster** (nur
Pfad-Pins) statt mit der nackten Kennung, ist keine der sieben Stellen betroffen.
