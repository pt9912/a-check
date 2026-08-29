# slice-100 — `verify-closure-notes` blendet Zitat-Kontexte aus

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers.
**Berührte Spec-Stellen:** — *(keine; Sensor-Präzision ohne Vertragsberührung)*
**Deckt:** [`SL-004`](../../steering-loop.md#sl-004--ein-neuer-doku-sensor-meldet-im-ersten-lauf-sein-eigenes-umfeld),
zwei neue Vorfälle.
**Bezug:** Schuld aus [slice-099](../done/slice-099-form-rest-und-fall-des-alten-baums.md).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

Beim Abschluss von slice-099 hat `verify-closure-notes` **zweimal** fehlalarmiert. Der zweite Fall
ist der aufschlussreiche: die Notiz beschrieb den ersten Fehlalarm und **zitierte** dabei die
auslösende Wendung in Backticks — der Sensor meldete Text, der *über* sein Muster spricht.

Das ist wörtlich die Klasse
[`SL-004`](../../steering-loop.md#sl-004--ein-neuer-doku-sensor-meldet-im-ersten-lauf-sein-eigenes-umfeld),
und deren Antwort steht seit slice-061 als Guide in Schritt 5 des Workflow-Skeletts: *wer einen
Sensor über Markdown baut, blendet Zitat-Kontexte von Anfang an aus und nimmt eine Fixture mit
zitiertem Muster in den Selbsttest auf.* `verify-closure-notes` stammt aus slice-050 und ist
**älter als der Guide** — der Vorfilter fehlt ihm.

Und er ist im selben Skript schon vorhanden: die Satzzählung blendet Code-Blöcke und Inline-Code
seit slice-075 aus, mit dem Kommentar *„derselbe Vorfilter wie in `verify-slice-links`"*. Nur die
Placeholder- und Floskel-Prüfung sieht ihn nicht.

**Warum das jetzt auffällt:** die Baseline `v5.12.0` verlangt in **jeder** Closure-Notiz einen
Abschnitt über offene Risiken. Notizen sprechen seitdem regelmäßig über Offenes — und damit über
das Vokabular, aus dem die Placeholder-Liste besteht.

## 2. Betroffene Module

`tools/verify-closure-notes.sh` und
[`docs/plan/steering-loop.md`](../../steering-loop.md) (`SL-004`, Vorfallszahl).
Eine Schicht: Gate-/Werkzeug-Schicht.

## 3. Was ausdrücklich **nicht** getan wird — und warum das eine ADR bräuchte

Die Placeholder-Liste enthält eine Wendung, die in einem Risiko-Block **unzitiert** und völlig
legitim vorkommt; genau daran ist der erste Fehlalarm entstanden. Sie zu streichen wäre die
naheliegende zweite Hälfte des Fixes — und sie unterbleibt hier.

**Grund:** eine Prüfregel zu verkleinern ist eine Lockerung, und
[`AGENTS.md`](../../../../AGENTS.md) §3.6 verlangt dafür eine ADR, keinen Sensor-Commit. Der
Vorfilter ist etwas anderes: er *präzisiert*, indem er den Sensor an eine im Repo bereits
etablierte Regel angleicht (`SL-004`-Guide, Satzzählung, `verify-slice-links`). Er senkt keine
Schwelle, er behebt eine Auslassung.

Die Vokabular-Frage wandert damit ins Beobachtungs-Register der Etappe D. Ob sie eine ADR wert
ist, entscheidet der nächste unzitierte Vorfall.

## 4. Auszuführende Gates

`make verify` (der Sensor hängt an der Verifikations-Schicht, nicht an `gates`), dann `make gates`.

**Negativ-Probe, beide Richtungen — sonst wäre der Vorfilter ein Freibrief:** eine Fixture mit
**zitiertem** Muster muss schweigen, die bestehende mit **unzitiertem** Platzhalter muss weiter
feuern. Ohne die zweite ließe sich ein Vorfilter, der alles verschluckt, nicht von einem korrekten
unterscheiden.

## 5. DoD

- [ ] Placeholder- und Floskel-Prüfung laufen über denselben Zitat-Vorfilter wie die Satzzählung
      (Code-Blöcke und Inline-Code entfernt) — Beleg: Diff.
- [ ] Der Selbsttest trägt eine Fixture mit zitiertem Muster, die schweigt, **und** behält die mit
      unzitiertem, die feuert — Beleg: Target-Ausgabe.
- [ ] [`SL-004`](../../steering-loop.md#sl-004--ein-neuer-doku-sensor-meldet-im-ersten-lauf-sein-eigenes-umfeld)
      nennt die zwei neuen Vorfälle mit Zahl und die Antwort — Beleg: Diff.

Pflicht, aber **kein** Liefer-Punkt: `make gates` und `make verify` grün — Ausgabe in eine Datei,
Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Berührt wird die **Gate-/Werkzeug-Schicht** (`tools/`) — in der Modus-Deklaration pro Sub-Area als
Greenfield geführt.
