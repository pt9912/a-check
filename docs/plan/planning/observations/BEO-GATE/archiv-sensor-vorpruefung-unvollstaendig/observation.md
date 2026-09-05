# Sensor-Vorprüfung vor Archivierung deckt nicht jede Modul-Klasse ab

**Sub-Area:** Gate-/Werkzeug-Schicht

Die von der Baseline verlangte Vorprüfung „Geltungsbereich der vorhandenen Sensoren gegen die
Stub-Form" wurde einmal geführt (sieben Fundstellen, zwei `structure`-Regeln angepasst) — und war
beim ersten echten `APPLY=1`-Lauf trotzdem unvollständig: das `ids`-Modul (Linkpflicht für bare
`AC-*`/`ADR-*`-Kennungen im `Hervorgegangen:`-Feld eines Stubs) und `ignore-refs` (ein Review-Report
zitiert einen jetzt archivierten anderen Review-Report) standen nicht auf der geprüften Liste. Das
Vorbild `d-check` hatte für genau diese beiden Fälle bereits eine Lösung (`exempt-paths` je
ID-Muster, `ignore-refs`-Einträge) — aber die eigene Vorprüfung hatte nicht bei `d-check`s
`.d-check.yml` selbst nachgesehen, welche Module dort für den Stub-Pfad bereits eine Ausnahme
tragen. Zusätzlich ein reiner Code-Fund: `ReadWelleField` (übernommen aus `d-check`) bricht die
Feld-Erfassung nur an einer Leerzeile ab — a-checks Kopf-Blöcke reihen `**Welle:**` und
Folgefelder (`**Deckt:**`/`**Bezug:**`) ohne Leerzeile aneinander, was beim ersten Anwendungslauf
Folgefeld-Text mit in den Stub zog.
