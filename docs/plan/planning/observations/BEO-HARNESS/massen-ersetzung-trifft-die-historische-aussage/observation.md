# Eine Massen-Ersetzung trifft auch die Aussage, die stehenbleiben muss

**Sub-Area:** Harness-Einstieg

Ein Versions- oder Kennungs-Wechsel wird über eine Dateiliste ausgerollt — `sed` über alles, was
den alten Wert trägt. In **lebenden** Dokumenten ist das richtig: Der Wert ist ein Zeiger, und
Zeiger wandern mit.

Dieselben Dokumente tragen aber auch Sätze, in denen der alte Wert **eine historische Tatsache**
ist: *„Migration `A`→`B`"*, *„der alte Stand wird entfernt"*, *„gemessen gegen `A`"*. Die Ersetzung
macht daraus eine falsche Aussage — und zwar eine, die plausibel aussieht, weil sie den aktuellen
Wert nennt.

**Warum kein Lauf das fängt:** Ein Versions-Sensor bindet an die **Form** eines Pins (ein Pfad, ein
Digest, eine `uses:`-Zeile). Eine nackte Kennung im Fließtext ist keine Form, sondern Prosa — und
ob der Satz sie als Zeiger oder als Tatsache führt, ist ein Urteil über seine Bedeutung.

**Die gefährlichste Stelle ist das Dokument, das den Wechsel beschreibt.** Der Migrationsplan nennt
beide Stände häufiger als jedes andere Dokument, und fast jede seiner Nennungen des alten ist
historisch. Er steht trotzdem in der Liste, weil er auch Zeiger trägt.
