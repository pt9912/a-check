**Vorgang:** slice-197

**Fund:** Der Haken `- [x] Unabhängiger Review durchgeführt (Report unter docs/reviews/)` wurde mit
dem Arbeits-Commit `010960e` gesetzt, **bevor** der Review gestartet war. `git show 010960e`
zeigt die Zeile als `- [ ]` → `- [x]`; der Report entstand erst danach.

**Zweites Auftreten derselben Klasse.** Das erste ist `slice-169`.

**Wie es auffiel:** im unabhängigen Review des Slice selbst — der Reviewer hat den Commit gegen den
Zeitpunkt des Reports gehalten. Kein Gate deckt das, solange der Slice in `in-progress/` liegt:
`make doc-reviews` greift erst für einen Slice in `done/`, und sein Exit 0 ist für einen laufenden
Slice **kein** Beleg.

**Was daraus folgt:** Der Haken ist am Ende wahr — der Report existiert bei der Closure. Die
Aussage war zum Zeitpunkt ihrer Setzung falsch, und genau das ist die Klasse. Die Closure-Notiz
nennt es.
