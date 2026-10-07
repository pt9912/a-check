# Ein Shell-Prüfer bricht unter `pipefail` ohne Meldung ab

**Sub-Area:** Gate-/Werkzeug-Schicht

Ein Bash-Prüfer mit `set -euo pipefail` extrahiert Werte per Pipe (`grep … | head -1`,
`tar … | od -N2`, `x=$(grep …)`). Findet `grep` nichts, oder schließt das letzte Glied die Pipe
früh (SIGPIPE), endet die Zuweisung mit einem Fehler — und `set -e` beendet das Skript, bevor die
Prüfung mit ihrer Meldung erreicht ist. Das Ergebnis ist rot, aber stumm, oder ein Exit-Code
(141), der wie ein Werkzeugfehler aussieht statt wie ein Befund.
