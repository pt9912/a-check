**Vorgang:** slice-211

**Fund:** Drei Review-Runden fanden denselben Bruch an drei Stellen der `shapes`-Dateisuche:
die `expect`-Datei als Symlink (F-3), ein Symlink-Verzeichnis im `expect`-Pfad (N-1) und eines im
literalen Präfix eines `files`-Globs (G-1, Code aus slice-210). Jedes Mal las a-check eine Datei
außerhalb der Scan-Wurzel und zeigte ihren Inhalt als Befund. Drei Funde, **ein** Vorgang.

**Behoben** durch `Lstat` je Pfadbestandteil, mit Mutations-Gegenprobe;
[SPEC-CONF-001](../../../../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) nennt die Regel seit
0.35.0. Meine eigene Vorhersage zum `files`-Fall („WalkDir folgt keinem Symlink, also sicher") war
falsch — sie übersah das `os.Stat` auf den Präfix vor dem Walk.
