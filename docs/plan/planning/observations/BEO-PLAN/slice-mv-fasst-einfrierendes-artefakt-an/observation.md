# `slice-mv` zieht Verweise auch dort nach, wo sie nicht nachgezogen werden dürfen

**Sub-Area:** Planungs-Harness

`make slice-mv` ersetzt den alten Lifecycle-Pfad eines Slice in **jedem** Dokument, das ihn
nennt. Die Menge ist zu weit: Ein **Review-Report** ist ein Lauf-Beleg und friert ein — seine
`pfad`-Felder halten den Stand fest, gegen den geprüft wurde, und der war `in-progress/`. Nach
dem Nachzug behauptet der Report, gegen einen `done/`-Stand gelaufen zu sein, den es beim Lauf
nicht gab.

Das ist die **Gegenrichtung** zu
[`verweis-auf-wandernden-slice`](../verweis-auf-wandernden-slice/observation.md): Dort fehlte
Deckung, hier gibt es zu viel davon. Beide Male ist die Ursache dieselbe — das Werkzeug kennt
die Verweis-*Form*, aber nicht die *Klasse* des Dokuments, in dem sie steht.

Der Reviewer-Skill nennt die Ausnahme ausdrücklich: Die Zitier-Form eines Reports verlangt
Kennung statt Adresse, *„das `pfad`-Feld auf den geprüften Gegenstand ist davon nicht betroffen —
es hält den Stand des Laufs fest und darf das."* Genau dieses erlaubte Feld trifft `slice-mv`.
