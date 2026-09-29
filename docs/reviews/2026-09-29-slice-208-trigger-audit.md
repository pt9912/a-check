# Review-Report: slice-208 — 2026-09-29

**Review-Art:** Code-Review — der neue Sensor `verify-trigger-audit` samt
Verdrahtung, geprüft gegen Slice-Plan, Sensor-Datei und die Hard Rules
(`AGENTS.md` §3, §4, §5); adversarial mit Mutations-Proben gegen den
done/-Bestand.

**Gegenstand:** slice-208 (Trigger-Audit für MR-Einträge mechanisieren),
HEAD `5efeeaf` — Plan `docs/plan/planning/in-progress/slice-208-trigger-audit-mr-eintraege.md`,
`tools/verify-trigger-audit.sh`, `Makefile`, `.claude/hooks/pretooluse-command-guard.sh`,
`harness/README.md` §Sensors, `harness/sensors/verify-trigger-audit.md`.

**Skill:** `.harness/skills/reviewer.md` @ `5efeeaf`
**Modell:** glm-5.3-flash · **Datum:** 2026-09-29

> **Zitier-Form** *(dieser Block bleibt stehen — er ist Norm, kein
> Ausfüll-Hinweis; die `<Platzhalter>` darin sind Formbeispiele)*. Dieser
> Report friert ein; was er zitiert, bewegt sich
> weiter. Deshalb: **Kennung, nicht Adresse** — `slice-<Kennung>` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link
> (`v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt>). Der vendored Baum trägt
> genau einen Tag; der Sprung löscht den alten, und ein Link darauf färbt beim
> nächsten Bump ein Artefakt rot, das niemand mehr anfassen darf. Ein `pfad`-Feld
> auf den **geprüften Gegenstand** ist davon nicht betroffen — es zitiert den
> Stand des Laufs und darf ihn festhalten (`v<X.Y.Z>` ·
> `regelwerk/grundlagen-harness-dateien.md` §harness/README.md als
> Einstiegspunkt — diese Zeile ist selbst ein Beispiel der Form).

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde — ohne
diese Liste ist der Lauf nicht reproduzierbar):

- Slice-Plan slice-208 (in `in-progress/`, Stand `5efeeaf`)
- `AGENTS.md` §3 (Hard Rules), §4 (Gate-Index-Disziplin), §5 (Traceability, Mess-Regeln)
- `spec/lastenheft.md` — `AC-QA-02` (Hermetik und ehrliche Heuristik-Grenze)
- Beobachtung `BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter` (3×) und `BEO-HARNESS/trigger-audit-ceremonie` (neu, 1×)
- Vorgänger-Muster: `verify-review-haken` (slice-204), `verify-risiko-ausgaenge` (slice-102/129)
- Nachlauf-Commits `43857b2` und `5efeeaf` (Closure-Notiz, Register, §6-Ausgang korrigiert) — geprüft wurde der Stand danach

---

## Findings

Jedes Finding folgt dem **§Output-Schema des Reviewer-Skills** — der
verbindlichen Single Source of Truth. Die Spalten unten sind nur
**gespiegelt** (Bequemlichkeit beim Ausfüllen), nicht neu definiert; bei
Abweichung gilt der Skill bzw. dessen Quelle
`v6.13.0` · `regelwerk/modul-10-review-harness.md` §Ziel-Form: Reviewer-Skill.

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | Der Lauf prüft in der realen Repo-Struktur **null** Closures: der Glob `done/slice-*.md` erfasst nur die 10 flachen Dateien (154–167, alle grandfathered, Phrase-Prüfung läuft nie), während 134 wellenlose Closures in `done/wellenlos/` — dem archivierten Ort auch von slice-208 selbst — und die Slice-Closures unter `done/welle-*/` außerhalb liegen; eine simulierte verletzende Closure ohne Audit-Zeile in `done/wellenlos/` bleibt grün (Exit 0, „0 Closure(s) geprueft"). Skript-Kommentar („wellenlose wie Welle-Slices in done/ selbst") und Sensor-Datei (Grenze 3 nennt nur `done/welle-*/`, verschweigt `done/wellenlos/`) behaupten die Abdeckung; DoD-Punkt 1 des Plans („meldet rot, wenn …") ist im realen Bestand unerfüllbar. | DoD-Punkt 1 slice-208; `AGENTS.md` §4; `v6.13.0` · `regelwerk/modul-13-gates.md` §Vorhanden ≠ behauptet | `tools/verify-trigger-audit.sh:54` (Glob), `:50-51` (Kommentar), `harness/sensors/verify-trigger-audit.md:5` (Vertrag) | ja — `make verify-trigger-audit` meldet „10 Closure(s) geprueft", alle 10 grandfathered; `done/wellenlos/` zählt 134 Dateien; Mutationsprobe: verletzende Closure in `done/wellenlos/` → Exit 0 | Pruefer-ohne-Gegenstand |
| F-2 | MEDIUM | `aktive_mr` zählt Mentions statt Tabelleneinträge: die ok-Meldung weist MR-000 (Prosa-Absatz „Zu MR-000") und MR-014 (Titel des MR-030-Eintrags „löst MR-014 auf") als aktiv aus — MR-014 ist durch MR-030 aufgelöst; die echte Aktiven-Tabelle führt 8 Einträge (MR-016, MR-019, MR-024, MR-025, MR-027, MR-028, MR-029, MR-030), keiner fehlt, zwei kommen hinzu. Die Meldung ist der einzige Ort, an dem der Sensor die zu sichtende Menge ausweist, und behauptet sie falsch. | Mess-Regeln, `AGENTS.md` §5 | `tools/verify-trigger-audit.sh:33-36` | ja — Sensor-Ausgabe („aktive MR: MR-000 MR-014 …", 10 Kennungen) gegen Zählung der Aktiven-Tabelle (8 Zeilen) | Zaehlung-zaehlt-Mentions-statt-Eintraege |
| F-3 | MEDIUM | Der Sensor prüft den Phrase-Präfix **dateiweit** (`grep -qF`), nicht §7-scoped und nicht auf die im Plan-Ziel versprochene „mit Kennung"-Hälfte: die erste reale Closure (slice-208 selbst) besteht die Prüfung, ohne eine Audit-Zeile mit MR-Kennungen zu tragen — die Phrase steht nur in §1 (Zitier-Beispiel) und §7 (Lerneintrag nennt die Sensor-Zeile), keine MR-Kennung erscheint in §7; die bereits berechnete `aktive_mr`-Menge wird nie gegen die Zeile abgeglichen. Die Sensor-Datei deklariert als Grenze 1 nur das Urteil, nicht die unenforcede Kennung-Hälfte des Plan-Ziels und die dateiweite Trefferfläche. | Plan-Ziel §1 slice-208 („mit Kennung"); `BEO-HARNESS/trigger-audit-ceremonie` (1×) | `tools/verify-trigger-audit.sh:43` | ja — `grep -n "Trigger-Audit der aktiven MR"` auf den Plan trifft nur Z. 28 und Z. 86; §7 enthält keine MR-Kennung; Fixture Z. 100 (Zeile ohne Kennung) ist grün | Trigger-Audit-Ceremonie |
| F-4 | LOW | Fünf Fehlerpfade des Selbsttests referenzieren `$PD_SAVE`, das nie gesetzt ist (korrekt wäre `$DONE_SAVE`); die Mutationsprobe (grünes Fixture entzaubert) zeigt: die Diagnose crasht mit „PD_SAVE ist nicht gesetzt" (Exit 1 statt 2), `rm -rf "$tmp"` läuft nicht und lässt das Fixture-Verzeichnis liegen. Fail-closed bleibt erhalten, der Fehlerpfad als Diagnose nicht. | Maintainability | `tools/verify-trigger-audit.sh:113,119,123,135,140` | ja — Mutationsprobe auf einer Kopie: Exit 1, Meldung „Zeile 119: PD_SAVE ist nicht gesetzt" | Fehlerpfad-ungebundene-Variable |
| F-5 | INFO | Die Produktions-ok-Zeile endet auf „ (Selbsttest gefeuert)." — der Produktionslauf ist nicht der Selbsttest; die Meldung suggeriert einen Selbsttest-Durchlauf, wo nur der reguläre Scan lief. | Maintainability | `tools/verify-trigger-audit.sh:78` | ja — `make verify-trigger-audit`-Ausgabe | Produktionsmeldung-mit-Selbsttest-Suffix |
| F-6 | LOW | Die Closure-Notiz verlinkt als Herkunfts-Adresse `../open/slice-208-trigger-audit-mr-eintraege.md` — die Datei liegt in `in-progress/`, das Ziel existiert nicht; `make doc-check` bleibt grün, weil die Lifecycle-Invariante wandernder Slices nach Kennung auflöst. Nach der Zitier-Form (Kennung statt Adresse in einfrierenden Artefakten) friert die Notiz eine falsche Adresse ein, statt die Kennung allein zu tragen. | Zitier-Form, `AGENTS.md` §5 | `docs/plan/planning/in-progress/slice-208-trigger-audit-mr-eintraege.md:102` | ja — `make doc-check` grün bei nicht existierendem Pfadziel (`docs/plan/planning/open/` enthält slice-208 nicht) | Eingefrorene-Adresse-statt-Kennung |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Lense 1 — Sensor-Lauf und Exit-Codes (`bash tools/verify-trigger-audit.sh`, `make verify-trigger-audit`) | geprüft — Befunde F-1 (Menge), F-2 (Extraktion), F-4 (Fehlerpfad), F-5 (Suffix); Exit 0 und Selbsttest in beiden Richtungen selbst ohne Befund |
| Lense 1 — Grandfathering-Grenze (`AUDIT_FROM=208`) gegen den done/-Bestand | geprüft, ohne Befund — 10 flache Closures (154–167), alle grandfathered, null flache ≥ 208; die Grenze selbst arbeitet korrekt (Grandfathering-Fixture rot-grün-richtig) |
| Lense 2 — Ceremonie-Gefahr, erste reale Closure | geprüft — Befund F-3: die erste reale Closure (slice-208) besteht formgerecht ohne die Kennung-Hälfte; der §6-Ausgang trägt das Restrisiko als `BEO-HARNESS/trigger-audit-ceremonie` weiter |
| Lense 3 — Verdrängung (Beleg-Moment, Welle-Verzeichnisse) | geprüft — Befund in F-1: der Beleg entsteht abschlussbereit in `in-progress/` und liegt nach dem `mv` in `done/wellenlos/` (als Stub neben dem Archiv) — der Scan erreicht keinen der Momente; die Ausgrenzung von `done/welle-*/` ist gegenwärtig folgenlos (das Repo läuft wellenlos, keine wellengebundene Closure ≥ 208), ihre Begründung („deren Form ist eine andere") vermischt aber Welle-Results-Notizen mit den Slice-Closures in den welle-Verzeichnissen |
| Lense 4 — Gate-Index-Disziplin (§4): Target ↔ `.PHONY` ↔ GATES-Liste ↔ Index-Zeile | geprüft, ohne Befund — `Makefile:113` Target + `Makefile:53` `.PHONY` + Hook-Zeile 78 (GATES-Liste) + `harness/README.md:97` (Index-Zeile) vorhanden; `make doc-targets` 0 Befunde (654 Dateien), `make doc-mentions` 18/18, `guard-selftest` grün; Sensor-Datei und Skript-Wortlaut decken sich in den deklarierten Grenzen — die Lücke liegt in den nicht deklarierten (F-1) |
| Lense 5 — §3.7 Kommentar-Regeln (Skript, Sensor-Datei) | geprüft, ohne Befund — die Kommentarblöcke tragen Zusage („WAS HIER GILT"), Grenze („NICHT geprueft") und Herkunfts-Anker als auflösbare Kennungen; keine Chronik, kein Konjunktiv über verworfene Alternativen |
| Lense 6 — Form: Slice-Plan (DoD, §6-Ausgang, §8, Welle-Feld) | geprüft, ohne Befund im Stand `5efeeaf` — §6-Ausgang ist ein Mitglied der geschlossenen Menge („weiter offen" → `BEO-HARNESS/trigger-audit-ceremonie`); die Vorform „Ausgang: bei Closure" (Zeitpunkt statt Ausgang) wurde im Nachlauf korrigiert und ist in der Closure-Notiz benannt; §8 mit beiden Vorschalt-Prüfungen, Welle-Feld korrekt „ohne Welle" |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 2 |
| LOW | 2 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Pruefer-ohne-Gegenstand · Zaehlung-zaehlt-Mentions-statt-Eintraege · Trigger-Audit-Ceremonie · Fehlerpfad-ungebundene-Variable · Produktionsmeldung-mit-Selbsttest-Suffix · Eingefrorene-Adresse-statt-Kennung

**Geltungsbereich der Prüfung:** der committete Stand `5efeeaf` — Skript, Verdrahtung (Makefile, Hook, Gate-Index, Sensor-Datei), Plan inkl. gefüllter Closure-Notiz. Die Plan-Kopfzeilen (DoD, §4, §5) wurden gegen Skript-Verhalten und Repo-Bestand geprüft, nicht gegen eine `make gates`-Docker-Kette (doc-targets, doc-mentions, guard-selftest und die Sensor-Läufe sind gelaufen; `lint`/`test`/`coverage` des Hauptmoduls sind von diesem Gegenstand unberührt). Die Mutations-Proben liefen auf Kopien unter `/tmp` — der Review-Gegenstand blieb unverändert.

## Verdikt

**Merge-blockierend:** ja — F-1 ist eine Harness-Lüge-Klasse: Das Gate behauptet
Abdeckung (Skript-Kommentar, Sensor-Vertrag, DoD-Punkt 1), prüft in der realen
Repo-Struktur aber null Closures, und der erste Fall, den es binden soll
(slice-208 selbst), liegt nach dem `mv` außerhalb seines Geltungsbereichs. Die
„dass"-Mechanisierung ist damit unwirksam: eine künftige Closure ohne die
Audit-Zeile bliebe grün. F-2 und F-3 betreffen die Aussagekraft der Meldung und
die Kennung-Hälfte des Plan-Ziels; F-4 bis F-6 sind nachrangig.

**Übergabe:** Findings gehen an den Implementer (Rückkante
Review → Plan bei Plan-Defekt); die **Finding-Klassen** gehen zusätzlich
in die Slice-Closure §7 und von dort in den Zähler. Dieser Report selbst
ist ein **Lauf-Beleg** (Audit: dieser Diff, dieser Skill, dieses Modell,
dieses Verdikt) — er wird über Läufe hinweg nicht wieder gelesen, und
muss es nicht. Der Report ersetzt keine
Verifikation — DoD-/Spec-Konformität prüft der Verifier separat
(Modul 11; anderes Prüf-Artefakt, anderer Eingabe-Kontext).
