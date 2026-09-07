# Ein Prüf-Muster trifft nur die häufige Schreibweise seines Gegenstands

**Sub-Area:** Gate-/Werkzeug-Schicht

Ein Sensor, der seinen Gegenstand über ein Textmuster findet, wird an den Vorkommen kalibriert, die
beim Bauen sichtbar sind — und das sind die **häufigen**. Eine seltenere, gleichwertige
Schreibweise desselben Gegenstands fällt dabei nicht auf: Sie erzeugt keinen Fehlalarm, sondern
**Schweigen**, und Schweigen ist von Korrektheit nicht zu unterscheiden.

Das trennt diese Klasse von
[`pruefer-ohne-gegenstand-oder-aufruf`](../pruefer-ohne-gegenstand-oder-aufruf/observation.md):
Dort ist die Prüfmenge **leer** oder der Prüfer läuft nicht; hier läuft er, hat Gegenstand, und
meldet für die Mehrheit der Fälle korrekt. Die Lücke ist ein Teil-Ausfall, und genau deshalb
schwerer zu sehen — der grüne Lauf ist ja größtenteils verdient.

**Woran sie auffällt:** nicht am Sensor, sondern an einer Handarbeit daneben, die dieselbe zu enge
Annahme macht. Wer das Muster schreibt und wer den Bestand von Hand umstellt, ist dieselbe Person
mit derselben Vorstellung davon, wie der Gegenstand aussieht.
