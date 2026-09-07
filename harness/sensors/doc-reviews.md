# `make doc-reviews` — Review-Report-Deckung

## Vertrag

Eine `done/`-Slice-DoD-Zeile mit der Phrase „unabhängiger Review" braucht
mindestens einen Report unter `docs/reviews/` mit derselben `slice-<NNN>`-Kennung
im Dateinamen. 1:N zulässig, Substring-Match, nicht rekursiv.

## Grenze — was das Grün nicht abdeckt

1. **Slices ohne die Trigger-Phrase** — der Sensor ist **Opt-in pro Slice über
   die DoD-Phrase selbst**. Ohne sie ist die Kandidatenmenge für diesen Slice
   leer, und das Grün sagt nichts über ihn. Kein Fehlalarm, aber auch keine
   Deckung; heilbar nur dadurch, dass jeder Slice die Phrase trägt.
2. **Archivierte Stubs** — sie tragen keine DoD mehr und fallen aus der
   Kandidatenmenge. Beabsichtigt, permanent.
3. **Ob der Report etwas taugt** — geprüft ist die Existenz einer Datei mit
   passender Kennung, nicht ihr Inhalt. Permanent; das Urteil trägt der
   Reviewer-Skill.

**Die Phrase ist der Angelpunkt**, und dass sie noch greift, prüft ein eigener
Sensor: `make dcheck-phrase-selftest`.

## Sperren

- eigenes `doc-*`-Target aus demselben Grund wie `doc-workflows`.

## Bindung

`DC-FA-RVW-001` · konfiguriert in `.d-check.yml`, im `gates`-Aggregat seit
slice-160.
