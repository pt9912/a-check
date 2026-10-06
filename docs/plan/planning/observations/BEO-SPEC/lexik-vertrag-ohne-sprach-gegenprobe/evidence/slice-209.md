**Vorgang:** slice-209

**Fund:** Drei Review-Runden fanden je einen Desync zwischen dem geschriebenen Kotlin-Lexik-Vertrag
([SPEC-EXTRACT-001](../../../../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)) und Kotlin selbst: Zeichenketten mit Dollar-Präfix fehlten (F-3), ein maskiertes
`\$` wurde mitgezählt (N-1), und das Beispiel zur Korrektur zählte die Bytes falsch (K-1). Jeder
Fall erlaubte, eine Abhängigkeit als Kommentar zu verstecken. Drei Funde, **ein** Vorgang.

**Was half:** Zählen Zeichen für Zeichen mit `od -c` statt Lesen; die Folgepflicht „Tests je
Lexik-Form" steht in der DoD des implementierenden Slice.
