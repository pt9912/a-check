**Vorgang:** slice-187

**Fund:** [`harness/conventions.md`](../../../../../../../harness/conventions.md) §Adaptions-Block
sagt, an einem akzeptierten `MR`-Eintrag werde nichts nachträglich geändert — *„analog zur
ADR-Immutabilität ([`AGENTS.md`](../../../../../../../AGENTS.md) §3.5)"*. Der Verweis liest sich als
Verweis auf **dieselbe Durchsetzung**; `make doc-immutable` prüft aber ausschließlich
`docs/plan/adr/[0-9]*.md`. Die `paths`-Liste des Moduls `vcs` nennt `harness/conventions.md`
nicht, und kein Kommentar sagt das.

**Wie es auffiel:** beim Abgleich selbst. Zwei Ziel-Form-Punkte ließen sich nur durch eine
Änderung an [`MR-000`](../../../../../../../harness/conventions.md#mr-000) übernehmen; die Frage *„was passiert, wenn ich das trotzdem tue?"* wurde
gemessen statt angenommen — `make doc-immutable` über die Commit-Range blieb **grün**. Die
Änderung ist trotzdem unterblieben und steht als Kommentar **über** dem Eintrag; die Regel gilt
ohne Sensor.

**Zweites Mal dieselbe Klasse.** Behoben ist der konkrete Fall: §Disziplin trägt die Grenze am
Satz — *analog* meint die Regel, nicht ihre Durchsetzung — samt dem Hinweis, was ein künftiger
Sensor dort tragen müsste (ein `MR`-Eintrag hat kein `Status:`-Feld, an dem `immutable-when`
greifen könnte).
