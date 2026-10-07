# Ein rekursiver Prüfer über Fremdeingaben hat keine Tiefengrenze

**Sub-Area:** Adapter

Ein Prüfer, der verschachtelte Fremdeingaben rekursiv absteigt, bricht bei großer, aber gültiger
Schachtelungstiefe mit einem Laufzeit-Abbruch (Stack-Überlauf) ab statt mit einer Meldung — und
der Vertrag nennt keine Grenze. Die Eingabe ist gültig; die Antwort ist kein dokumentierter
Ausgang.
