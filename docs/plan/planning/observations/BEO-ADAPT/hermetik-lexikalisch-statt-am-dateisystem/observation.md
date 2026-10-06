# Eine Hermetik-Zusage wird am Pfad-Text geprüft statt am Dateisystem

**Sub-Area:** Adapter

„Zeigt nicht aus der Scan-Wurzel hinaus" wird lexikalisch geprüft — kein `/` am Anfang, kein `..`
nach `path.Clean` —, während der Lesezugriff dem Dateisystem folgt. Ein Symlink irgendwo im Pfad
führt dann trotz bestandener Prüfung hinaus, und der fremde Inhalt erscheint in der Ausgabe. Die
Zusage gilt dem Text, die Lücke liegt im Dateisystem.
