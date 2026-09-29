# `make verify-trigger-audit` — Sichtung der aktiven MR-Einträge belegt

## Vertrag

Jede Closure-Notiz eines Slices ab `slice-208` in `docs/plan/planning/done/`
trägt eine Zeile **„Trigger-Audit der aktiven MR:"** — den Beleg, dass die
Closure die aktiven `MR`-Einträge aus
[`harness/conventions.md`](../../harness/conventions.md) §Aktive Adaptionen
angesehen hat. Die leere Aktiven-Menge wird **gemeldet**, nicht still
übersprungen.

## Grenze — was das Grün nicht abdeckt

1. **Das Urteil** — ob „0 offen" stimmt oder die genannten `MR`-IDs wirklich
   abgelaufen sind, ist Prosa-Auswertung ([`AGENTS.md`](../../AGENTS.md) §3.7);
   der Sensor prüft die **Form** der Sichtung, nicht ihre Wahrheit. Permanent.
2. **Grandfathering** — Closures vor slice-208 entstanden, bevor die Zusage
   stand; sie werden nicht geadelt, sie sind ausgenommen. Der Grenzstand
   wandert mit (`AUDIT_FROM`).
3. **Welle-Closures in `done/welle-*/`** werden nicht gescannt — deren Form
   ist eine andere (Baseline `modul-06` Closure-Schritt 2); dieses Repo läuft
   wellenlos. Permanent.
4. **Der Scan-Ersatz ist die Klasse** — `structure` hat keine Datei-Existenz-
   Bedingung, das `reviews`-Modul scannt genau ein `done-dir` (d-check
   v0.79.0); darum ein eigenes `verify`-Skript statt eines CR an das
   Fremdwerkzeug (slice-204-Muster). Permanent.

## Bindung

Harness-Prozess ([`AGENTS.md`](../../AGENTS.md) §5) · Antwort auf
[`BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter`](../../docs/plan/planning/observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md)
bei 3× · slice-208 · im `verify`-Aggregat.
