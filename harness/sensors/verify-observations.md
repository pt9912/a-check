# `make verify-observations` — Deckung des Beobachtungs-Registers

## Vertrag

Jeder in `done/` zitierte Beobachtungs-Pfad hat ein Verzeichnis unter
`docs/plan/planning/observations/`, und jedes Verzeichnis trägt ein **nicht
leeres** `evidence/`. Beide Kennungs-Formen werden erkannt: die neue
`BEO-<KUERZEL>/<slug>` und die alte `BEO-NNN` über die `Ehemals:`-Zeile.

Der Zähler wird **nicht geführt, sondern abgeleitet** — er ist die Zahl der
Evidence-Dateien. Die frühere Zähler-Prüfung ist damit strukturell entfallen:
Ein abgeleiteter Zähler kann seiner Belegliste nicht widersprechen.

## Grenze — was das Grün nicht abdeckt

1. **Lage und Existenz der Beleg-Datei** — ob der im Beleg genannte Vorgang
   wirklich existiert und dort liegt, wo seine Klasse abschließt, prüft der Lauf
   nicht (Baseline `modul-06`). Ein erfundenes `slice-999` bliebe unentdeckt.
   Permanent — ein Repo darf Vorgänge führen, die es nicht als Datei ablegt.
2. **Die Umkehrung** — „jedes Verzeichnis ist irgendwo zitiert" wird **nicht**
   geprüft, und das ist Absicht: Die meisten Einträge stehen unter der Schwelle
   und sind nirgends zitiert. Permanent.

## Bindung

Harness-Prozess (Regelwerk `modul-06` §Das Beobachtungs-Register) · slice-102,
Verzeichnisform slice-139 · im `verify`-Aggregat.
