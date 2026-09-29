# Review-Report: slice-204 — 2026-09-29

**Review-Art:** Plan + Code — geprüft gegen Plan + Konventionen (Modul 10 §Drei Review-Arten); DoD-/Spec-Konformität bleibt Sache des Verifiers.

**Gegenstand:** slice-204 (committed auf `main`, `d216b88`): Plan `docs/plan/planning/in-progress/slice-204-review-haken-an-report-binden.md`, `tools/verify-review-haken.sh`, `Makefile` (Target `verify-review-haken` + `verify`-Anschluss), `harness/README.md` §Sensors (Index-Zeile), `harness/sensors/verify-review-haken.md`.

**Skill:** `.harness/skills/reviewer.md` @ `d216b88` ·
**Modell:** glm-5.3-flash · **Datum:** 2026-09-29

> **Zitier-Form:** Kennung, nicht Adresse — `slice-204` statt Lifecycle-Pfad, `make doc-targets` statt Link auf die Sensor-Datei; Baseline-Stellen als Tag + Pfad in Inline-Code (`v6.13.0` · `regelwerk/<datei>.md` §<Abschnitt>). Ein `pfad`-Feld auf den geprüften Gegenstand hält den Stand des Laufs fest und darf das.

**Eingangs-Kontext:**

- slice-204 (Plan, in-progress) — Review-Gegenstand
- `.d-check.yml` (Module `structure` (1)/(7), `reviews`) — für die Deckungs-Aussagen
- `harness/conventions/MR-019-review-dod-opt-in.md` (Wortlaut der Trigger-Phrase)
- `spec/lastenheft.md` §AC-QA-02 (ehrliche Heuristik-Grenze)
- `AGENTS.md` §3 (Hard Rules) + §5; `harness/README.md` §Sensors (Gate-Index)
- Register: `BEO-GATE/attestierung-vor-dem-vorgang` (observation, state, 5 Belege)

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | Sensor-Datei, Skript-Kommentar und Plan §1 behaupten, die structure-Bedingung 7 decke die Hook-Fälle in `open/`/`next/`; Bedingung 7 bindet per Glob nur `open/**` und deklariert `next/` ausdrücklich als nicht gedeckt (fail-closed leer, gemessen slice-202). Ein Haken in einem `next/`-Slice wäre damit von keinem der beiden Wächter erfasst, während die Grenze-Deklaration des neuen Sensors das Gegenteil zusagt. | Reviewer-Skill HIGH — nachweislich falsche Tatsachenbehauptung; `AGENTS.md` §3.7 | `harness/sensors/verify-review-haken.md:22` · `tools/verify-review-haken.sh:24,59` · `docs/plan/planning/in-progress/slice-204-review-haken-an-report-binden.md:28` | ja — `.d-check.yml:430` (Glob `open/**`) gegen `.d-check.yml:422` (Grenze-Kommentar `next/ ist NICHT gedeckt`); keine weitere Bedingung prüft `planning/next` | Deckungsaussage über den eigenen Geltungsbereich hinaus |
| F-2 | MEDIUM | Die Phrase-Prüfung ist eine Pipeline mit `grep -q` unter `pipefail`: endet der zweite grep nach dem ersten Treffer früh, stirbt der vordere grep an SIGPIPE, die Pipeline kehrt 141 zurück, und der `if`-Zweig behandelt einen vorhandenen Haken als nicht vorhanden — falsches Grün statt Befund. Reproduktion: 53-KB-Fixture mit Haken in Zeile 1 meldet KEIN MATCH (Status 141). | Maintainability (Verifikations-Werkzeug mit nachweislichem False-Green-Pfad) | `tools/verify-review-haken.sh:49` | ja — Reproduktion oben; bei realer Slice-Größe (KB-Bereich) heute nicht erreicht | Prüfpipeline kehrt Treffer in Nicht-Treffer (SIGPIPE unter pipefail) |
| F-3 | MEDIUM | Der Selbsttest-Kommentar sagt, die done/-Fixture präfiere die Abgrenzung; geprüft wird aber nur ein leerer `PROGRESS_DIR` — keine Verzweigung liest die Fixture, und die Abgrenzungs-Richtung von DoD-Punkt 2 (`doc-reviews`-Bedeutung bleibt unberührt) ist damit behauptet, aber nicht pro genehm. | `AGENTS.md` §3.7 (Kommentar-Zusage ohne Prüfwirkung); Mess-Regel „Eine Mutations-Probe belegt erst, wenn sie rot war" | `tools/verify-review-haken.sh:86-87,135-137` | ja — Lesen des Selbsttests gegen die Fixtures; die done/-Fixture taucht in keinem Prüflauf auf | Probe trifft den Prüf-Gegenstand nicht |
| F-4 | LOW | §1 nennt als Bindungsort noch das `structure`-Modul („die Bindung hier läuft über das `structure`-Modul") und §4 rechnet mit „zwei Konfigurations-Blöcken" und einem Modul-Blocker, während §3 die gebaute Abweichung auf die eigene Schalung dokumentiert — die Teil-Ersetzung hat §1/§4 nicht erreicht. | `AGENTS.md` §3.7 | `docs/plan/planning/in-progress/slice-204-review-haken-an-report-binden.md:35,70-73` | ja — Plan §1/§4 gegen §3 und die gebauten Dateien | Plan-Abweichung einseitig nachgezogen |
| F-5 | LOW | Plan-Bezug und §8 zählen die Beobachtung mit 4× (Vorfälle bis slice-191); Skript-Kopf und Sensor-Datei zählen 5× — slice-205 trat während der Laufzeit bei — und das Register (state.md) trägt 5×. Der Zähler-Stand im Plan liegt hinter dem Register zurück. | `AGENTS.md` §3.7 (Zustandsfeld nennt Zustand + Beleg) | `docs/plan/planning/in-progress/slice-204-review-haken-an-report-binden.md:8,101` · `tools/verify-review-haken.sh:4-5` · `harness/sensors/verify-review-haken.md:32` | ja — state.md des Registers gegen Plan und Sensor-Datei | Zähler-Stand hinter Register zurück |
| F-6 | LOW | Report-Namen in Bereichs-Form bedienen nur ihre Endpunkte: `2026-09-05-slice135-157-multi-linsen-review.md` matched slice-135, nicht slice-157 — ein gedeckter Slice ohne eigene Kennung im Dateinamen würde rot gemeldet; diese Namensform des Bestands ist in der Grenze-Liste der Sensor-Datei nicht benannt. | Maintainability | `tools/verify-review-haken.sh:42` | ja — `ls docs/reviews` gegen die Match-Regex (slice-135: Treffer, slice-157: kein Treffer) | Kennungs-Form im Report-Dateinamen unvollständig |
| F-7 | INFO | `MR-019` zitiert die Trigger-Phrase kleingeschrieben („unabhängiger Review"), das Skript prüft case-sensitiv großgeschrieben („Unabhängiger Review") — die im Bestand und im structure-Muster deployte Form; eine kleingeschriebene DoD-Zeile entginge dem Sensor. | `harness/conventions/MR-019-review-dod-opt-in.md` | `tools/verify-review-haken.sh:34` | ja — `grep -qF` ist case-sensitiv; Bestand: 2 abgehakte Zeilen mit Großform, 0 mit Kleinform | Case-Abhängigkeit der Trigger-Phrase |
| F-8 | INFO | Im Selbsttest wird `PROGRESS_DIR="$tmp/leer"` in zwei aufeinanderfolgenden Zeilen doppelt gesetzt; wirkungslos. | Maintainability | `tools/verify-review-haken.sh:136-137` | ja — Lesen | Redundante Zuweisung in Test-Fixture |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Sensor-Korrektheit (Regex-/Phrasen-Formen, Lense 1) | geprüft — Befunde F-1, F-2, F-6, F-7; im Übrigen ohne Befund: die Formen ohne Trenner und mit führenden Nullen (`slice135`) matchen, die Kennungs-Grenze slice-205 gegen slice-2050 ist im Selbsttest getrennt |
| Verdrängung/Fehlalarm (Lense 2) | geprüft, ohne Befund — Lauf auf dem aktuellen Stand: Exit 0, 1 in-progress-Slice geprüft, kein Fehlalarm; Wirksamkeits-Probe: die Vorfälle 4 und 5 (slice-191, slice-205 — beide mit exakter Phrasen-Form im `in-progress`-Stand, beide ohne Report) hätte der Sensor als rot gemeldet; die Vorfälle 1–3 (slice-169, 197, 200) liegen außerhalb seines Scans und sind dort Sache der structure-Bedingung (7) bzw. ihrer Nachfolge |
| Gate-Index-Disziplin (Lense 3) | geprüft, ohne Befund — Target existiert als Makefile-Regel und hängt im `verify`-Aggregat; `make doc-targets` grün (645 Dateien, 0 Befunde, beide Richtungen); Index-Zeile trägt die Bindung; Zellenlängen 73/164/212 ≤ 250 (`make doc-structure` grün) |
| §3.7 Kommentar-Regeln (Lense 4) | geprüft — Befunde F-3, F-4; die übrigen Skript-Kommentare tragen Zusage, Kopplung und Grenze, keine Chronik, keine ANPASSEN-Reste |
| Hard Rules §3.6 (Lense 5) | geprüft, ohne Befund — neues Gate, keine gesenkte Schwelle; die Bindung ist deklariert (Index-Zeile + Sensor-Datei), ein ADR ist dafür nicht erforderlich |
| Plan-Form (Lense 6) | geprüft — Befunde F-4, F-5; im Übrigen ohne Befund: §3 dokumentiert die Plan-Abweichung (structure → eigene Schalung), der §6-Ausgang stammt aus der geschlossenen Dreier-Menge mit Begründung, und die DoD ist durchgehend unabgehakt — keine Selbst-Attestierung im Slice, dessen Gegenstand genau das fängt |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 2 |
| LOW | 3 |
| INFO | 2 |

**Finding-Klassen dieses Laufs:** Deckungsaussage über den eigenen Geltungsbereich hinaus · Prüfpipeline kehrt Treffer in Nicht-Treffer (SIGPIPE unter pipefail) · Probe trifft den Prüf-Gegenstand nicht · Plan-Abweichung einseitig nachgezogen · Zähler-Stand hinter Register zurück · Kennungs-Form im Report-Dateinamen unvollständig · Case-Abhängigkeit der Trigger-Phrase · Redundante Zuweisung in Test-Fixture

## Verdikt

**Merge-blockierend:** ja — F-1 (die Grenze-Deklaration des Sensors behauptet Coverage in `next/`, die das geprüfte Geschwister-Gate nachweislich nicht liefert) blockiert die Übernahme; F-2 und F-3 sind vor der Closure zu klären (False-Green-Pfad im Prüfweg bzw. behauptete Probe ohne Prüfwirkung — für ein Gate der Verifikations-Schicht ist ein belegbarer False-Green-Pfad kein Zustand zum Schließen). F-4 bis F-8 sind vor dem `git mv` nach `done/` günstig mitgebracht.

**Übergabe:** Findings gehen an den Implementer (Rückkante Review → Plan bei Plan-Defekt, F-4/F-5); die **Finding-Klassen** gehen zusätzlich in die Slice-Closure §7 und von dort in den Zähler — F-3 trifft dieselbe Klasse, die als `BEO-GATE/probe-liefert-den-gegenstand-mit` bereits 3× gezählt hat (4. Auflage, im Lese-Schritt der Closure zu prüfen). Dieser Report selbst ist ein **Lauf-Beleg** (Audit: dieser Stand, dieser Skill, dieses Modell, dieses Verdikt) — er wird über Läufe hinweg nicht wieder gelesen und ersetzt keine Verifikation (DoD-/Spec-Konformität prüft der Verifier separat, Modul 11).
