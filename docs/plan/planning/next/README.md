# `next/` — priorisiert für die nächste Welle

Der **Zustand** eines Slice ist das Verzeichnis, in dem seine Datei liegt; er wechselt nur
durch `git mv` als eigenen Commit ([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

`next/` ist die Stufe zwischen „geplant" und „in Arbeit": der Slice ist **geschnitten und
priorisiert**, aber noch nicht begonnen. Er wartet auf einen beobachtbaren Trigger — nicht auf
ein Datum.

## Ein- und Ausgänge

| von → nach | Bedingung |
|---|---|
| `open/` → `next/` | für die nächste Welle priorisiert |
| `next/` → `in-progress/` | Trigger eingetreten, WIP-Limit frei (**genau ein** aktiver Slice) |
| `in-progress/` → `next/` | **Rückführung:** der Slice ist zu groß — zurück zur Zerlegung, nicht dehnen (Größen-Regel: höchstens drei DoD-Punkte) |

Die zweite Rückführung `in-progress/` → `open/` (blockiert) führt **nicht** hierher: ein
blockierter Slice ist nicht priorisierbar, solange der Blocker steht.

## Warum dieses Verzeichnis wieder existiert

Es wurde in 53 Slices genau einmal benutzt (slice-009) und danach mitsamt Ort aufgelöst, während
[`AGENTS.md`](../../../../AGENTS.md) §5 den Zustand weiter nannte — ein deklarierter Zustand ohne
Ort ist eine stille Setzung (Fund **B-18** aus
[slice-048](../done/welle-12/slice-048-modul-delta-lesen.md), behoben in slice-053).

Dass ein Slice `next/` durchläuft, ist **nicht** Pflicht: der direkte Weg `open/ → in-progress/`
bleibt zulässig. Pflicht ist, dass der Zustand einen Ort hat, wenn die Regel ihn nennt.
