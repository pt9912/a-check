# Beobachtungs-Register

Regeln dieser Ablage: Baseline-Regelwerk `modul-06-roadmap.md`
§Das Beobachtungs-Register — wer schreibt, wer liest, wann gestrichen wird, welche Form ein Beleg
hat und welchen der drei Ausgänge ein Eintrag ab 3× trägt.

**Form — drei Dateien, drei Lebensdauern**, je Beobachtung ein Verzeichnis
`BEO-<KUERZEL>/<slug>/`: `observation.md` (unveränderlich ab Anlage: Bezeichnung und Sub-Area) ·
`state.md` (veränderlich: `offen` oder einer der drei Ausgänge) ·
`evidence/<vorgangs-id>.md` (unveränderlich ab Merge, eine je Auftreten). Der Zähler wird
**abgeleitet** — er ist die Zahl der Evidence-Dateien; es gibt kein Feld, in das man ihn schreibt.

**Die Kennung ist der Pfad.** Beide Segmente werden nachgeschlagen, nicht erfunden: `<KUERZEL>` ist
das Sub-Area-Kürzel aus [`harness/conventions.md` §Modus-Deklaration pro Sub-Area](../../../../harness/conventions.md#modus-deklaration-pro-sub-area),
`<slug>` lowercase Kebab-Case.

**Wer schreibt:** die **Slice-Closure** — neues Verzeichnis mit `observation.md`, oder eine weitere
Datei in ein vorhandenes `evidence/`. Erhöht wird nie etwas; es wird nur angelegt, der Zähler folgt.
**Wer liest:** die Welle-Closure, was **3×** erreicht hat — und die Slice-Planung (§8 des Plans),
was darunter steht.

**Migriert mit [slice-139](../done/wellenlos/slice-139-beobachtungsregister-migration.md)** von der
Tabellenform (`observations.md`, angelegt mit slice-101) auf diese Verzeichnisform — Auslöser:
[slice-135](../done/wellenlos/slice-135-regelwerk-v600-delta-analyse.md) (Baseline-Sprung `v5.12.0` →
`v6.0.0`). Jedes migrierte Verzeichnis trägt in `observation.md` eine `**Ehemals:**`-Zeile mit
seiner alten `BEO-<NNN>`-Kennung — historische Zitate in `done/` verwenden weiterhin die alte
Nummer (Prosa wird nicht umgeschrieben, [`AGENTS.md`](../../../../AGENTS.md) §3.7 sinngemäß);
`make verify-observations` löst beide Formen auf.

**Ist nichts offen**, steht hier nur diese Datei — ein leeres Verzeichnis führt `git` nicht. Aktuell
33 aktive Beobachtungen unter [`BEO-GATE/`](BEO-GATE), [`BEO-HARNESS/`](BEO-HARNESS),
[`BEO-PLAN/`](BEO-PLAN), [`BEO-SPEC/`](BEO-SPEC) und [`BEO-KERN/`](BEO-KERN); eine gestrichene
unter [`BEO-PLAN/vier-form-vergleiche-ungeprueft/`](BEO-PLAN/vier-form-vergleiche-ungeprueft)
(`state.md` trägt `gestrichen`, das Verzeichnis bleibt liegen).
