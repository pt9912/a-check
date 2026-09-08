**Vorgang:** slice-182

**Fund:** `harness/README.md` schrieb Baseline-Normtext an zwei Stellen nach, statt zu verweisen.
`modul-13` §Vorhanden ≠ behauptet stand **zweimal in derselben Datei**, 22 Zeilen auseinander —
einmal als Einleitung der Nicht-Gates-Tabelle (500 Zeichen), einmal im Absatz darunter (553). Der
Abschnitt §Rollen und ihre Übergabe-Artefakte gab auf 1542 Zeichen Prosa `modul-08` §Die neun
Übergaben und §Rollen-Regeln wieder; der Satz *„Verifikation grün, Validation rot"* stand dreifach
im Repo (`modul-08`:142, [`MR-016`](../../../../../../../harness/conventions/MR-016-validator-unbesetzt.md):19, `harness/README.md`:189).

**Warum es hier zählt:** `harness/README.md` ist Schritt 1 des Minimal Agent Workflow — jeder Lauf
liest sie ganz. Beide betroffenen Dateien verbieten das selbst: `AGENTS.md` §1 („sie dupliziert
deren Inhalt nicht; sonst entsteht Drift"), `harness/README.md` §Purpose („Diese Datei dupliziert
sie nicht").

**Wie es auffiel:** Der Maintainer las den Absatz und fragte „was ist damit? Steht das nicht schon
im Regelwerk?" — zweimal hintereinander, zu zwei verschiedenen Stellen. Nicht durch einen Lauf;
die Regel ist inferentiell (`AGENTS.md` §3.7).

**Behandlung:** beide Stellen auf Verweise gekürzt (1930 → 1278 und 1542 → 584 Zeichen sichtbare
Prosa). Der Eintrag bleibt offen: Ein Sensor dafür wäre ein Urteil über die Herkunft eines Satzes.
