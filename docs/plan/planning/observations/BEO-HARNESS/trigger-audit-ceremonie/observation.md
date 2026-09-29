# Die Trigger-Audit-Zeile läuft Gefahr, zur Ceremonie zu werden

**Sub-Area:** Harness-Einstieg

Der Sensor `verify-trigger-audit` (seit slice-208) prüft nur die **dass**-Hälfte:
dass eine Closure die Zeile „Trigger-Audit der aktiven MR:" trägt. Er kann nicht
prüfen, ob die Sichtung stattfand — eine immer gleiche „0 offen"-Zeile ohne Blick
in die aktiven Einträge wäre eine Formulierung ohne Beobachtung und läfe mechanisch
grün. Die Grenze ist in der Sensor-Datei benannt (Grenze 1); die Gefahr bleibt
trotz bestehen, weil kein Urteil über die Sichtung maschinell möglich ist.
