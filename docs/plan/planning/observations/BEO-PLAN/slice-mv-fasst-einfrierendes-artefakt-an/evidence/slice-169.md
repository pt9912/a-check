**Vorgang:** slice-169

**Fund:** `make slice-mv SLICE=slice-169 TO=done` meldete „4 Datei(en) mit Verweisen
nachgezogen". Zwei davon waren lebende Dokumente (Roadmap, ein offener Slice-Plan), eine ein
Evidence-Beleg — und eine der **Review-Report dieses Slice**, in dem sieben `pfad`-Felder von
`docs/plan/planning/in-progress/…` auf `…/done/…` umgeschrieben wurden. Betroffen war auch der
Kopf-Satz *„plus der uncommittete Arbeitsstand von …"*, der damit einen Zustand behauptet, den
es beim Lauf nicht gab.

**Wie es auffiel:** beim Lesen des `git status` nach dem `slice-mv` — die Datei unter
`docs/reviews/` stach heraus. Nicht durch einen Lauf: `make doc-check` blieb grün, weil der
Report seine Pfade in Inline-Code nennt und nicht als Markdown-Link.

**Behandlung in diesem Vorgang:** Die Änderung am Report ist mit `git checkout` zurückgenommen;
`doc-check` bleibt grün. Der Nachzug in `evidence/slice-173.md` **bleibt** — dort steht ein
echter Markdown-Link, der sonst bräche, und eine Adress-Korrektur ändert die Aussage des Belegs
nicht.
