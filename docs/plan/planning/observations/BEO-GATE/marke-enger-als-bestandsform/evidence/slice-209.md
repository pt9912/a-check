**Vorgang:** slice-209

**Fund:** Die neue Anforderung [AC-FA-RULE-012](../../../../../../../spec/lastenheft.md#ac-fa-rule-012) trug ihre
Boundary-Fälle in der Bestandsform `- **Boundary (versteckte Anweisung):**`. `make verify`
(Modul `structure`, Regel (5) in `.d-check.yml`) meldete `section-marker-missing`, weil
`require-all` die Zeichenfolge `Boundary:` verlangt. Behoben, indem die erste Boundary-Zeile die
nackte Marke trägt und die Qualifikation als Kursiv-Zusatz dahinter.

**Wie es auffiel:** am ersten Lauf nach dem Schreiben — die 19 grandfathered Anforderungen hatten
die Regel seit ihrer Einführung nie gegen die qualifizierte Form laufen lassen.
