**Stand:** offen

der Workflow nennt den CHANGELOG in Schritt 7 („falls ein öffentlicher Vertrag berührt"), aber **kein Gate prüft es**: `gate-consistency` vergleicht nur Versions-Nummern (`version.md#aktuell` == aktuellstes CHANGELOG-**Release**), nicht ob eine Änderung einen Eintrag hat. Aufgefallen erst bei der Release-Prep, also am spätestmöglichen Punkt. **Zweiter Vorfall** (slice-133): 21 Commits seit `v0.18.0` ohne Eintrag, erneut erst bei der Release-Prep gefunden. Bei 3× ist es eine Lücke und verlangt Guide oder Sensor
