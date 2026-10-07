**Vorgang:** slice-217

**Fund (Review F-1, F-2; Delta-Review D-1):** Der Gate-Index sagte, die Release-Pipeline teste den
gebauten Digest über das neue Target — die Pipeline rief es nicht auf. Der Kopfkommentar von
`tools/multiarch-check.sh` versprach „genau zwei Einträge, keine weitere Plattform": ein
Array-Feld in einem Eintrag ließ eine dritte Plattform durch (F-2), und nach dessen Fix prüfte das
Skript die Plattform-Beschriftung des Index nicht mehr (D-1). Drei Funde, **ein** Vorgang; der
Prüfer lief, war grün und prüfte weniger, als sein Kommentar sagte. Behoben, je mit Gegenprobe.
