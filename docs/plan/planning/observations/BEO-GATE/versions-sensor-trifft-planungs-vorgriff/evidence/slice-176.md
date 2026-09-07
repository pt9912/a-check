**Vorgang:** slice-176 (Planung)
**Fund:** Dieselbe Klasse, **andere Blickrichtung**. Der erste Beleg (slice-174) betraf den
*Ziel*-Stand: Welle-Datei und Etappen-DoD nannten den Pfad, der erst entstehen sollte. Hier ist es
der **alte** Stand: Die DoD trug *„`.harness/baseline/v6.2.0/` ist entfernt"* — eine Aussage über
den Pfad, der **verschwinden** soll. `make doc-check` meldete `version-stale`.

Beide Male ist die Aussage sachlich richtig und der Sensor formal im Recht: Ein Planungsdokument
einer Migration muss über Stände sprechen, die nicht der adoptierte sind — den kommenden **und**
den gehenden. Die Beobachtung ist damit breiter als ihr Titel: nicht nur der *Vorgriff*, sondern
jeder **Bezug auf einen anderen Stand aus planerischer Absicht**.

Umgangen wie beim ersten Mal: Verzeichnis-Literal vermieden („der vorige vendorte Stand"). Das
kostet Präzision — die DoD nennt jetzt keinen prüfbaren Pfad mehr, sondern eine Beschreibung.

**Zweiter Vorgang, zweiter Beleg** (slice-174 · slice-176). Bei 3× wäre zu entscheiden, ob
Planungsdokumente eine eigene Schreibweise für Stände bekommen — etwa die Zitier-Form aus
[T-1](../../../../done/welle-15/slice-174-regelwerk-v650-delta-analyse.md), die genau dafür gebaut ist:
Kennung statt Adresse. Dann wäre der Umweg keine Umgehung mehr, sondern die Regel.

**Nachtrag, dritter Fund im selben Vorgang** (zählt nicht erneut): Auch die *Umsetzung* stolperte
darüber. §2.3 des Slice beschreibt den Beweis — „der vorige Stand entfernt, `make gates` danach
Exit 0" — und nannte den Pfad dabei als Literal. Der Sensor meldete ihn.

Das schärft die Klasse ein drittes Mal: Es trifft nicht nur die *Planung*, sondern jeden Text, der
**über** einen Stand spricht statt **auf** ihn zu verweisen. Ein Slice, der das Löschen eines
Standes dokumentiert, muss ihn benennen können — und genau dafür gibt es seit diesem Slice die
Zitier-Form. Sie wäre hier die richtige Antwort: `` `v6.2.0` `` als Kennung statt als Pfad. Dass
der Slice sie einführt und im eigenen Text nicht anwendet, ist der Befund.
