**Vorgang:** slice-187

**Fund:** Der Slice hat eine Hard Rule korrigiert und die Korrektur **im Regeltext erklärt**:
[`AGENTS.md`](../../../../../../../AGENTS.md) §3.3 trug nach dem Commit drei Sätze über den Zustand
davor — *„Bis slice-187 stand hier nur der Regelfall … und das Briefing sagte das Gegenteil."* Das
ist genau die Form, die §3.7 verbietet (*„Falsch: abwesenden Text beschreiben … die vorige Fassung
hält `git`"*), geschrieben in die Datei, die die Regel trägt.

**Der Anlass ist die Wiederholung selbst:** Es passiert im **selben Commit**, der die Regel
ändert. Wer eine Regel korrigiert, hat die alte Fassung im Kopf und will erklären, warum die neue
anders ist — der Ort dafür ist der Slice, nicht das Briefing.

**Drittes Vorkommen.** Gefunden hat es der unabhängige Review (Report zu slice-187, F-3), nicht der
Schreibende und kein Lauf. Behoben sind mit diesem Slice **beide** Stellen in `AGENTS.md`: die neu
entstandene in §3.3 und die im `state.md` seit slice-182 als Reststelle geführte in §5 (WIP-Limit,
*„Bis slice-077 stand hier ‚genau ein'"*). `grep -n "Bis slice-" AGENTS.md` liefert danach null
Treffer.
