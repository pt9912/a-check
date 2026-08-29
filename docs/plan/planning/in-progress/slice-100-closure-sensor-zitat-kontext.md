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

- [x] Placeholder- und Floskel-Prüfung laufen über denselben Zitat-Vorfilter wie die Satzzählung
      (Code-Blöcke und Inline-Code entfernt) — Beleg: Diff.
- [x] Der Selbsttest trägt eine Fixture mit zitiertem Muster, die schweigt, **und** behält die mit
      unzitiertem, die feuert — Beleg: Target-Ausgabe.
- [x] [`SL-004`](../../steering-loop.md#sl-004--ein-neuer-doku-sensor-meldet-im-ersten-lauf-sein-eigenes-umfeld)
      nennt die zwei neuen Vorfälle mit Zahl und die Antwort — Beleg: Diff.

Pflicht, aber **kein** Liefer-Punkt: `make gates` und `make verify` grün — Ausgabe in eine Datei,
Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** Placeholder- und Floskel-Prüfung laufen über denselben Zitat-Vorfilter wie die
Satzzählung, mit Fixtures in beide Richtungen — und die in slice-099 verbogene Formulierung ist
auf ihre zitierte Fassung zurückgesetzt.

**Lerneintrag — Form: geschärfte Regel.** *Ein Guide für Neubauten erreicht den Bestand nicht — und
eine Konventions-Änderung kann einen jahrelang stillen Sensor scharf machen.* Der `SL-004`-Guide
steht seit slice-061 in Schritt 5 des Workflow-Skeletts und richtet sich an den Moment, in dem ein
Sensor **gebaut** wird. `verify-closure-notes` ist aus slice-050 und war damit nie sein Adressat;
der Vorfilter fehlte ihm vier Monate lang folgenlos. Folgenlos blieb er, *weil* Closure-Notizen
bis dahin nicht über Offenes sprachen — bis die Baseline `v5.12.0` in **jeder** Notiz einen
Abschnitt über offene Risiken verlangte. Erst diese Konventions-Änderung hat den Sensor in
Vokabular geführt, das er nicht von seinem Prüfgegenstand unterscheiden konnte. **Der Prüfsatz für
künftige Migrationen:** ein Baseline-Sprung ändert nicht nur, was dokumentiert wird, sondern
**worüber** die Dokumente sprechen — und damit, worauf bestehende Textsensoren treffen.

**Zwei beobachtbare Closure-Kriterien:**

1. Der Selbsttest fährt beide Richtungen: dieselbe Wendung **zitiert** muss schweigen,
   **unzitiert** muss feuern. Ohne die zweite Fixture wäre ein Vorfilter, der alles verschluckt,
   von einem korrekten nicht zu unterscheiden.
2. Der eigentliche Beweis ist der Bestand, nicht die Fixture: die Stelle in
   [slice-099](../done/slice-099-form-rest-und-fall-des-alten-baums.md), an der der Sensor
   gefeuert hatte, trägt wieder ihre zitierte Fassung — und `make verify-closure-notes` läuft mit
   Exit 0 über 95 Notizen.

**Offene Risiken und ihr Ausgang:**

- *Die Platzhalter-Liste enthält eine Wendung, die auch unzitiert legitim vorkommt* — Ausgang:
  **weiter offen**, fürs Beobachtungs-Register (Etappe D). Sie zu streichen wäre eine
  Schwellen-Senkung und braucht nach [`AGENTS.md`](../../../../AGENTS.md) §3.6 eine ADR;
  entschieden wird am nächsten unzitierten Vorfall.
- *`SL-004` steht jetzt bei fünf Vorfällen und hat weiterhin keinen Sensor* — Ausgang:
  **gestrichen mit Begründung**. Der Eintrag begründet das selbst: das Muster ist **Bauwissen**
  über Sensoren, kein wiederkehrender Laufzeit-Fehler; es gibt keinen Lauf, in dem es sich zeigen
  könnte. Was die zwei neuen Vorfälle ändern, ist nicht die Antwort, sondern ihre Reichweite — der
  Guide gilt ab jetzt auch für den **Bestand**, nicht nur für Neubauten.

**Folge-Slices:** Etappe D — Risiko-Ausgänge als Sensor und das Beobachtungs-Register.

## 7. Sub-Area-Modus

Berührt wird die **Gate-/Werkzeug-Schicht** (`tools/`) — in der Modus-Deklaration pro Sub-Area als
Greenfield geführt.
