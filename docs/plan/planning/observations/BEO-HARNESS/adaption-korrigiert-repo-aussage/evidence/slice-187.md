**Vorgang:** slice-187

**Fund:** Der Ziel-Form-Abgleich fand, dass die ID-Schema-Deklaration in
[`MR-000`](../../../../../../../harness/conventions.md#mr-000) die **Beobachtungs-Kennung** nicht
führt — die Vorlage führt sie. Behoben mit
[`MR-023`](../../../../../../../harness/conventions.md#mr-023), und dessen Feld
`Ersetzt-Baseline-Regel` steht auf `—`: Er korrigiert eine **Repo**-Aussage, keine Baseline-Regel.

**Drittes Mal dasselbe Muster**, und zum zweiten Mal an derselben Zeile-Familie: [`MR-020`](../../../../../../../harness/conventions.md#mr-020) hat die
ID-Schema-Deklaration schon einmal so repariert (veraltete `ADR-NNNN`-Zeile), [`MR-023`](../../../../../../../harness/conventions.md#mr-023) tut es
erneut. Unter dem Fork-Test sind beide Rückbau-Kandidaten — wer a-check forkt, übernimmt
Adaptionen, die nichts adaptieren.

**Der Trigger dieses Eintrags ist damit eingetreten:** `state.md` nannte als Auflösung *„die
Überarbeitung der ID-Schema-Deklaration"*. Genau die ist fällig — und sie ist zu groß für eine
Closure, weil sie einen Eintrag ablöst, dessen Kennungen repo-weit in Commits stehen.

**Wie es auffiel:** durch den unabhängigen Review (Report zu slice-187, F-4). slice-187 hatte
zunächst ein drittes, nirgends vorgesehenes Werkzeug gewählt — einen freistehenden Absatz oberhalb
des Eintrags — und mit *„ließe sich **nur** durch eine inhaltliche Änderung beheben"* begründet.
§Disziplin nennt zwei Instrumente, und [`MR-020`](../../../../../../../harness/conventions.md#mr-020) war der Präzedenzfall.
