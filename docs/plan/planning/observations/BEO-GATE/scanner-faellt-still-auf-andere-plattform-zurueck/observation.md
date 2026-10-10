# Der CVE-Scanner fällt bei Multi-Plattform-Eingabe still auf eine andere Plattform zurück

**Sub-Area:** Gate-/Werkzeug-Schicht

Trivy bekommt ein Multi-Plattform-Bild und eine gewünschte Plattform — und scannt ohne Meldung eine
andere. Der Bericht sieht aus wie ein Bericht über die gewünschte Plattform; ob er es ist, sagt nur
die Architektur in seinen Metadaten. Ein Scan ohne diesen Nachweis kann ein Bild für sauber
erklären, das er nie gesehen hat.
