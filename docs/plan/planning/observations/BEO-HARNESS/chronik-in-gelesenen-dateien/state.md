**Stand:** geplant → [slice-191](../../../open/slice-191-chronik-phrasen-sensor.md) (3×)

**Bei 3× ist „besser aufpassen" keine Antwort** ([`AGENTS.md`](../../../../../../AGENTS.md) §5).
Der Bestand ist behoben — `conventions.md` (slice-103), `harness/README.md` §Sensors und §Rollen
(slice-182), `AGENTS.md` §3.3 und §5 (slice-187); `grep -n "Bis slice-" AGENTS.md` liefert null
Treffer. Was fehlt, ist der Wächter.

**Die volle Klasse hat keinen** und wird keinen bekommen: *„Ist dieser Satz Chronik?"* ist ein
Urteil, kein Match — [`AGENTS.md`](../../../../../../AGENTS.md) §3.7 sagt das selbst. **Aber die
häufige Schreibweise ist greppbar**, und dreimal dieselbe: *„Bis slice-NNN stand hier …"*, *„bis
slice-NNN tat das …"*, *„Vorher stand hier …"*. Ein `forbid-pattern` auf diese Wendungen in den
**lebenden** Dateien fängt nicht die Klasse, sondern ihre häufigste Form — und genau diese Form
ist dreimal aufgetreten.

**Die Grenze gehört an den Sensor, nicht in die Fußnote:** Er prüft eine **Phrase**, nicht eine
Klasse; eine Umformulierung entkommt ihm. Wer ihn für den Wächter der Regel hält, hat einen halben
Wächter für einen ganzen genommen — dieselbe Falle wie bei
[`muster-trifft-nur-die-haeufige-schreibweise`](../../BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md).
Geplant als [slice-191](../../../open/slice-191-chronik-phrasen-sensor.md).

**Der Moment, an dem es passiert, ist bekannt:** derselbe Commit, der eine Regel ändert. Das ist
die Stelle, an der der Sensor feuern muss — und der Grund, warum ein Review-Guide allein dreimal
nicht gereicht hat.
