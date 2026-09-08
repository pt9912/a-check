# Ein Erstdurchgang wird als einmaliger Slice geschnitten, obwohl er eine Bestandsgröße abarbeitet

**Sub-Area:** Planungs-Harness

Ein Vorgang, der einen **Bestand** durchgeht — alle Ziel-Form-Paare, alle ADRs, alle Slices —,
wird als *ein* Slice geplant, weil die **Regel** dahinter eine ist. Der Umfang ergibt sich aber
nicht aus der Regel, sondern aus dem Bestand, und der ist beim Schneiden meist ungemessen.

Die Folge ist eine **Kette von Folge-Slices**, in der jeder das Risiko *„zu groß"* als
*eingetreten* schließt und an den nächsten übergibt. Jeder einzelne ist regelkonform — die
Rückführung ist vorab benannt, der Ausgang trägt eine Kennung —, und trotzdem war die Planung
falsch: Nicht die Slices sind zu groß, sondern der Gegenstand ist keine Slice-Größe.

**Das Unterscheidungsmerkmal:** Eine *Regel* verankern und einen *Bestand* durchgehen sind zwei
Vorgänge. Der erste ist ein Slice; der zweite ist eine Kampagne, deren Fortschritt gezählt wird.
