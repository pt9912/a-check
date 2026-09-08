**Stand:** offen (2×)

**Die offene Ursachen-Frage ist beantwortet — gemessen, nicht geschlossen.** Der Verdacht der
`observation.md` trifft zu: Der Git-Zugriff von `d-check` findet nur **kanonisch benannte** Packs
(`pack-*.idx`). Die Gegenprobe lief in beide Richtungen an einem Wegwerf-Repo mit zwei Commits:
Pack als `loose-<sha>` benannt ⇒ `Range-Basis "HEAD~1" nicht auflösbar`, dieselbe Datei nach `mv`
auf `pack-<sha>` ⇒ Exit 0, gleiche Version, gleicher Aufruf. Im eigenen Klon lagen **682** Commits
in einem `loose-<sha>.pack` und waren damit für den Prüfer unsichtbar; das Umbenennen der drei
Dateien (`.idx`/`.pack`/`.rev`) macht `trace-check`, `commit-scope-check` und `doc-immutable` über
jeden Range wieder grün, `git fsck --connectivity-only` bleibt sauber.

**Nicht die Version:** `v0.74.1` scheitert mit derselben Meldung wie `v0.75.0` (gegen eine Konfig
ohne den `mentions`-Block gemessen). Die CI war nie betroffen — ein frischer `git clone` benennt
Packs kanonisch.

**Was offen bleibt, ist der Auslöser.** Wer den nicht-kanonischen Namen erzeugt hat, ist unbelegt;
`maintenance.*` und `gc.auto` sind in diesem Klon nicht gesetzt, ein Timer läuft nicht. Damit kann
der Zustand wiederkehren, und der Ausgang steht noch aus: entweder ein CR an `d-check` (kanonische
Pack-Erkennung ist eine Bibliotheks-Eigenschaft, keine Repo-Eigenschaft) oder eine dokumentierte
Handreichung im eigenen Klon. **Der Eintrag bleibt darum offen, obwohl der konkrete Fall behoben
ist.**

**Drittes Vorkommen 2026-09-08, benannt statt gezählt:** Es fiel beim Abschluss von slice-186
erneut an und kostete den Umweg über die Versions-Hypothese. Es trägt **keinen** Beleg, weil kein
eigener Vorgang dahintersteht — die Messung lief innerhalb einer fremden Closure
(`modul-06` §Das Beobachtungs-Register: *„Ein Vorkommen ohne abgeschlossenen Vorgang bekommt
keinen Beleg und bewegt den Zähler nicht"*). Der Zähler steht deshalb weiter auf 2×; die zwei
Belege sind slice-170 und slice-172.

**Korrektur am Stand selbst:** Hier stand „offen (1×)" neben **zwei** Beleg-Dateien. Der Zähler
ist abgeleitet, nicht geführt — eine Zahl im Text, die ihrer Belegliste widerspricht, ist genau
die zweite Quelle, die die Register-Form vermeiden soll.
