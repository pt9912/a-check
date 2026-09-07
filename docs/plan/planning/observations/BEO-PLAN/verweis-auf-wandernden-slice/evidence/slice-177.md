**Vorgang:** slice-177
**Fund:** Vierter Vorgang in Folge. `make slice-mv` ließ den Verweis in der flach liegenden
Welle-Datei wieder stehen (`[slice-177](open/…)`), `make doc-check` meldete ihn, Korrektur per
`sed`.

**Neu an diesem Auftreten — und es ist der schwerere Teil:** Der Lifecycle-`git mv` landete
zusammen mit dem Inhalts-Umbau in **einem** Commit. Git erkannte nur `R040` (40 % Ähnlichkeit,
unter der 50 %-Schwelle), und `git log --follow` auf den neuen Pfad lieferte **einen** statt
**sechs** Commits — fünf Commits Vorgeschichte waren über den Lifecycle-Pfad verloren. Der Schaden,
den `AGENTS.md` §3.3 wörtlich voraussagt, war damit nicht mehr nur möglich, sondern eingetreten;
gefunden hat ihn der unabhängige Review, nicht der Lauf. Nachträglich aufgesplittet, Rename wieder
bei 100 %.

**Dritter Fall derselben Art in dieser Sitzung** (slice-172, slice-176, hier). Bei slice-176 fiel
er mir selbst auf, hier nicht. Kein Gate sieht ihn: `doc-check` prüft Links, nicht Rename-Raten.

**Zum Zähler:** Die `slice-mv`-Form-Lücke steht bei sechs Vorkommen in vier Vorgängen; die
Etappen B und E bringen je zwei weitere. Die §3.3-Verletzung ist eine **andere** Klasse und gehört
nicht in diesen Eintrag — sie hat noch keinen und wäre bei drei Vorfällen in einer Sitzung einen
wert.
