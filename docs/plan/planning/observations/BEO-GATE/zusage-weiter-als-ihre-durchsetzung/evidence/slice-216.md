**Vorgang:** slice-216

**Fund (Review F-1, F-3):** Der erste Vertragsentwurf für das Multi-Arch-Image sagte zwei Dinge zu,
die kein geplanter Prüfer hielt: „derselbe Commit ergibt denselben Index-Digest" (die
Gegenmessung im Review widerlegte es) und „byte-identische Ausgabe auf beiden Plattformen"
(weder ADR noch Folge-Slices verglichen die Plattformen je Release). Zwei Funde, **ein** Vorgang.
Behoben vor der Abnahme: die Reproduzierbarkeit ist keine Zusage mehr, der Plattform-Vergleich hat
einen Träger in der Release-Pipeline.
