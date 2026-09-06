# Zwei Baseline-Stände bleiben vendored, obwohl die Migration geschlossen ist

**Sub-Area:** Harness-Einstieg

[`conventions.md`](../../../../../../harness/conventions.md#baseline) §Baseline sagt: „Genau **ein**
Stand liegt vendored; mehrere sind nur während einer Migration zulässig." Nach dem Abschluss von
`welle-14` liegen `v6.0.0` **und** `v6.2.0` unter `.harness/baseline/`.

Der Grund ist benannt, nicht versehentlich: die `Ersetzt-Baseline-Regel`-Felder mehrerer
akzeptierter `MR`-Einträge zeigen inhaltlich unveränderlich auf `v6.0.0`-Pfade, und ein
akzeptierter Eintrag wird nicht nachträglich geändert. Damit hält die Adaptions-Disziplin den alten
Stand fest, und die Zusage „genau ein Stand" gilt auf unbestimmte Zeit nicht.

Beobachtet wird nicht der Zustand — der ist erklärbar —, sondern seine **Dauer ohne
Entscheidung**: es gibt keinen Mechanismus, der die Frage stellt, und keine Frist, nach der sie
gestellt werden müsste.
