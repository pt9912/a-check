# Review-Report: slice-190 — ID-Schema-Deklaration als Ganzes — 2026-09-29

**Review-Art:** Code — gegen Plan + Konventionen (unabhängiger Lauf, adversarisches
Review; der Lauf ist nicht der Autor des Gegenstands).

**Gegenstand:** `HEAD~5..HEAD` = 4ea201b · 0dc262c · 25a14d8 · baa393b · 2f8b68a
(Commit-Grenze per `git log --oneline -7` verifiziert: `cd3906b`/`2dc79ff` gehören
zu slice-191 und liegen außerhalb).

**Skill:** `.harness/skills/reviewer.md` @ 2f8b68a · **Modell:** GLM (glm-5.3-flash,
Claude-Agent-SDK) · **Datum:** 2026-09-29

**Eingangs-Kontext:**

- Slice-Plan `docs/plan/planning/in-progress/slice-190-id-schema-deklaration-ueberarbeiten.md`
  (§1 Out-of-Scope, §2 Ausgangsmessung, §3, §4 DoD, §7, §8, §9)
- `harness/conventions/MR-029-id-schema-deklaration-gesamt.md` (neu);
  `harness/conventions/done/MR-020-adr-vorlage-generisch.md`,
  `harness/conventions/done/MR-023-id-schema-beobachtungs-kennung.md` (abgelöst)
- `harness/conventions.md` (Tabellen, MR-000-Kommentar-Absatz),
  `spec/lastenheft.md` §3, `.d-check.yml` (`trace-check`-Formen)
- `AGENTS.md` §3/§5; `v6.13.0` · `regelwerk/modul-05-planning-harness.md`
  §Ziel-Form: Slice und §Zwei Schritte vor der Modus-Begründung;
  `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register;
  `v6.13.0` · `regelwerk/grundlagen-source-precedence.md` §Slice-Kennungen
- Beobachtung `BEO-HARNESS/adaption-korrigiert-repo-aussage` (`observation.md`,
  `state.md`, `evidence/` — 3 Belege: slice-097, slice-162, slice-187)

**Eigene Ausführungen:** `make doc-check` (Exit 0 — 627 Dateien, 0 Befunde) und
`make verify` (Exit 0 — 21 Anforderungen, 0 Waisen) wurden für diesen Lauf selbst
ausgeführt, nicht zitiert. `make gates` (Docker-Build-Kette) wurde **nicht**
wiederholt; die Grünf-Meldung im Commit 2f8b68a bleibt Implementer-/Verifier-Beleg.

---

## Findings

### F-1 — Closure-Notiz behauptet Register-Ausgang *verkörpert*, das Register trägt *geplant*

- `kategorie`: HIGH
- `quelle`: `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register
  (Ausgangszuweisung durch den Lese-Schritt; Herkunfts-Anker `seit slice-<NNN>` am
  Zielort); `AGENTS.md` §5 (Risiko-Ausgang)
- `pfad`: `docs/plan/planning/in-progress/slice-190-id-schema-deklaration-ueberarbeiten.md`:152–154, 163 — Gegenprobe: `docs/plan/planning/observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/state.md`:1
- `befund`: §8 schreibt, der Register-Eintrag sei *verkörpert* „(Anker am
  state.md), die 3×-Kette ist geschlossen" — der `state.md` trägt aber weiterhin
  `**Stand:** geplant → slice-190 (3×)`, also einen anderen der drei Ausgänge und
  keinen `seit slice-190`-Anker. Die Closure-Notiz ist als Bedingung des Übergangs
  nach `done/` formuliert und wäre in dieser Fassung eine gegen das Register-Artefakt
  nachweislich falsche Behauptung; der DoD-Haken „trägt seinen Ausgang" (Zeile 81)
  stützt sich auf denselben Widerspruch.
- `verifizierbar`: nein — kein Gate prüft Aussagen gegen den Registerstand; Lesen
  genügt (`state.md`:1 gegen §8).
- `klasse`: Attestierung vor dem Vorgang

### F-2 — MR-029 leitet das AC-FA-Bereichskürzel aus der Modus-Deklaration ab — das widerspricht MR-002 und der gelebten Praxis

- `kategorie`: HIGH
- `quelle`: `harness/conventions/done/MR-002-id-schema-bereichskuerzel.md` (Begründung:
  Bereiche `RULE`/`EXTRACT`/`CLI`/`CONF`/`DIST` aus dem Lastenheft);
  `harness/conventions.md` §Modus-Deklaration pro Sub-Area (Kürzel `SPEC`, `ADR`,
  `KERN`, `ADAPT`, `PLAN`, `GATE`, `REVIEW`, `HARNESS`, `USER`);
  `harness/conventions.md` §Anforderungs-Anlege-Prozess (Deklarationsort: Lastenheft §3)
- `pfad`: `harness/conventions/MR-029-id-schema-deklaration-gesamt.md`:8–9
- `befund`: Der AC-FA-Baustein der neuen „vollständigen, aktuellen Deklaration"
  behauptet, „das Bereichskürzel ist aus der Modus-Deklaration abgeleitet (MR-002)"
  — MR-002 leitet die Kürzel aus den Lastenheft-Funktionsbereichen ab, die
  Modus-Deklaration führt keine der AC-FA-Bereiche, und der Deklarationsort ist
  Lastenheft §3. Die Zeile hat das Muster von MR-023 (dessen BEO-Zeile die gleiche
  Herkunft korrekt trägt) auf eine Familie kopiert, für die sie falsch ist; sie
  irrt jeden künftigen Anforderungs-Autor über die Vergabestelle.
- `verifizierbar`: ja — Lesen gegen MR-002, §Modus-Deklaration und Lastenheft §3;
  kein Gate.
- `klasse`: Falsche Herkunfts-Angabe in Deklaration

### F-3 — Slice-Form-Begründung pinnt „Welle 130/131" — ein nicht vendorter, nicht auflösbarer Baseline-Stand

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5 (Zitier-Form: Kennung statt Adresse; für lebende Einträge
  Tag + Pfad in Inline-Code); `v6.13.0` · `regelwerk/grundlagen-source-precedence.md`
  §Slice-Kennungen („Welche Form gilt, deklariert das Repo")
- `pfad`: `harness/conventions/MR-029-id-schema-deklaration-gesamt.md`:16–18
- `befund`: Die Begründung zur numerischen Slice-Form zitiert „die Baseline
  (Welle 130/131, „Kennungen sind Namen")" — vendort ist aber nur `v6.13.0`
  (Kurs-Welle 153), und der Sachverhalt (Repo-Deklaration ausdrücklich erlaubt)
  trägt sich auch dort. Der Eintrag, dessen Kern-Verkaufspunkt die
  Migrationsfestigkeit ist (§7-Ausgang: „versions- und schwellenfrei"; Lerneintrag:
  „generische Formulierung … statt einer Versionsnummer"), trägt damit selbst einen
  veraltenden Stand-Bezug — die Wiederkunft genau des Musters, dessen Entfall §7
  als Begründung für „entfallen" nennt.
- `verifizierbar`: ja — `grep` über `.harness/baseline/v6.13.0/` (nur dort steht
  „Kennungen sind Namen"); „Welle 130/131" ist im Repo nicht auflösbar.
- `klasse`: Versions-Pin in generischer Deklaration

### F-4 — §9 bleibt Platzhalter: beide vorgelagerten Prüfungen sind unbelegt

- `kategorie`: MEDIUM
- `quelle`: `v6.13.0` · `templates/docs/plan/planning/slice.template.md` §8
  („Der Abschnitt selbst entfällt nie"); `v6.13.0` ·
  `regelwerk/modul-05-planning-harness.md` §Zwei Schritte vor der Modus-Begründung
- `pfad`: `docs/plan/planning/in-progress/slice-190-id-schema-deklaration-ueberarbeiten.md`:168–173
- `befund`: „Sub-Area-Wahl prüfen" und „offene Beobachtungen sichten" stehen
  unverändert als Platzhalter „(beim Übergang nach `in-progress/` auszufüllen …)"
  — der Übergang liegt vor dem Review-Range, die Closure-Schritte (baa393b,
  2f8b68a) im Range, und keiner der Commits füllt die Sichtungsergebnisse ein.
  Der Lese-Schritt der Slice-Planung (Register nach `BEO-HARNESS` gelesen,
  Zähler-Stand notiert, Vorgänger-Report geprüft) ist damit im Bestand unbelegt.
- `verifizierbar`: nein — Lesen; `doc-structure` prüft die Kopffelder, nicht die
  Ausfüllung von §9.
- `klasse`: Pflicht-Prüfung unbelegt

### F-5 — Closure-Notiz widerspricht sich zum Steering-Loop-Eintrag

- `kategorie`: MEDIUM
- `quelle`: `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Wellen-Closure-Prozedur
  Schritt 3 (Anker-Paarung gilt dem „liegt in <Zielort>"-Feld)
- `pfad`: `docs/plan/planning/in-progress/slice-190-id-schema-deklaration-ueberarbeiten.md`:147–149 gegen 160–161
- `befund`: Das Feld „Steering-Loop-Eintrag: — (nichts verkörpert …)" und der
  Paarungs-Satz „Anker — der Steering-Loop-Eintrag liegt in `MR-029`" im selben
  §8 behaupten Gegenteile: entweder ist die geschärfte Regel in MR-029 verkörpert
  (dann ist „—" falsch und der Anker zu führen), oder nichts ist verkörpert (dann
  hat die Anker-Paarung kein Objekt, und der Lerneintrag „Form: geschärfte Regel"
  trägt keinen Zielort).
- `verifizierbar`: ja — Lesen der beiden Stellen im selben Dokument.
- `klasse`: Widersprüchliche Closure-Felder

### F-6 — Eingefrorener MR-023-Eintrag zitiert slice-190 über die Lifecycle-Adresse

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §5 (Zitier-Form in einfrierenden Artefakten — seit slice-176,
  gemessen „null Nachzüge"); `harness/conventions.md` §Baseline (done-Einträge sind
  eingefroren, „der Sensor nimmt sie aus")
- `pfad`: `harness/conventions/done/MR-023-id-schema-beobachtungs-kennung.md`:38
- `befund`: Der beim Auflösen eingefrorene Eintrag verlinkt den abzulösenden Slice
  als `…/planning/in-progress/slice-190-id-schema-deklaration-ueberarbeiten.md`
  statt als Kennung — jeder weitere Zug der Slice-Datei (nächster Halt: `done/`)
  erzwingt einen Nachzug in eine eingefrorene Datei oder färbt `doc-check` rot;
  genau die 22-Nachzüge-Klasse, die slice-176 beendet hat. MR-020 zeigt die
  korrekte Form (nur `done/`-Pfade und Kennungen).
- `verifizierbar`: ja — der nächste `git mv` der Slice-Datei (doc-check rot bzw.
  Edit an der done-Datei).
- `klasse`: Adress-Zitat im eingefrorenen Eintrag

### F-7 — „Risiken aus §6" — die Risiken stehen in §7

- `kategorie`: LOW
- `quelle`: Slice-Plan eigene Gliederung (§6 = Closure-Trigger, §7 = Risiken und
  offene Punkte; der DoD-Punkt selbst nennt §7 korrekt)
- `pfad`: `docs/plan/planning/in-progress/slice-190-id-schema-deklaration-ueberarbeiten.md`:158 (ebenso Commit-Message baa393b)
- `befund`: Die Closure-Notiz verweist für die Risiko-Ausgänge zweimal auf „§6";
  die Ausgänge stehen in §7. Die Substanz (jedes Risiko mit genau einem Ausgang)
  ist vollständig — nur der Verweiser ist falsch.
- `verifizierbar`: ja — Lesen.
- `klasse`: Falscher §-Verweiser

### F-8 — Commit-Message 2f8b68a benennt GRENZE-Arbeit, die der Diff nicht trägt

- `kategorie`: LOW
- `quelle`: Provenance — `git show 2f8b68a` gegen die Message; die GRENZE-Arbeit
  stammt aus slice-202 (Commits 2de4aeb „GRENZE des Attestierungs-Sensors
  qualifiziert", e89638d), die vor dem Range liegen
- `pfad`: Commit 2f8b68a (Message) — Diff: nur ID-Links im Plan plus §2-Klammer
- `befund`: Die Message nennt „die GRENZE des neuen Sensors (erste Inhaltszeile)
  benannt, open/-Fall in der GRENZE ergänzt" als Bestandteil des Nachzugs — beide
  Inhalte fehlen im Diff, und „GRENZE" steht im Plan nirgends. Der Audit-Satz
  „dieser Commit hat das getan" trifft für zwei von drei Message-Postulaten nicht.
- `verifizierbar`: ja — `git show 2f8b68a` (Message-Diff-Schere exakt sichtbar).
- `klasse`: Message-Diff-Schere

### F-9 — DoD-Absolut-Aussage „keinen Verweis auf eine Korrektur-Kette"

- `kategorie`: LOW
- `quelle`: Slice-Plan DoD (erster Punkt) gegen MR-029-Titel und -Begründung
- `pfad`: `docs/plan/planning/in-progress/slice-190-id-schema-deklaration-ueberarbeiten.md`:76–77
- `befund`: Der abgehakte DoD-Punkt verspricht, wer die Deklaration nachschläge,
  „findet keinen Verweis auf eine Korrektur-Kette" — MR-029s Titel („löst MR-020
  und MR-023 ab") und Begründung erzählen die Kette ausdrücklich. Der Sinn des
  DoD (eine aktuelle Fassung statt MR-000 plus Korrekturliste) ist erreicht; die
  Absolut-Formulierung im Wortlaut nicht.
- `verifizierbar`: nein — Urteil über die Formulierung; Lesen genügt.
- `klasse`: Absolut-Formulierung im DoD

### F-10 — DC-FA-* ist in Nutzung, aber in keiner ID-Schema-Deklaration

- `kategorie`: INFO
- `quelle`: `harness/conventions.md` §Adaptions-Block (ID-Schema-Deklaration als
  Aufgabe); kein Regress gegen MR-000 — die Form war dort ebenfalls undeklariert
- `pfad`: `.d-check.yml`:36, 128, 178, 191, 228 (`DC-FA-RVW-001`, `DC-FA-VCS-001`, `DC-FA-COMMITS-001`, `DC-FA-TGT-001`, `DC-FA-STRUCT-001`, …)
- `befund`: Das Repo nutzt ein `DC-FA-<BEREICH>-<NNN>`-Schema mit neun Bereichen
  in `.d-check.yml`, das weder MR-000 noch der als Ganzes neue Eintrag MR-029
  führt. Der Vollständigkeits-Claim von MR-029 bezieht sich auf den historischen
  MR-000-Umfang — die Lücke ist vorbestehend und als künftiger Ergänzungs-MR
  benennbar, kein Befund gegen diesen Range.
- `verifizierbar`: nein — `grep` belegt die Nutzung; die Bewertung als Lücke ist
  Urteil.
- `klasse`: Deklarations-Lücke (bestehend)

## Negativbefunde

- Out-of-Scope-Disziplin §1: **geprüft, ohne Befund** — der MR-000-Wortlaut ist im
  Range unverändert geändert (nur der Kommentar-Absatz „über" dem Eintrag und die
  Tabellen in `harness/conventions.md`); kein anderer aktiver MR-Eintrag wurde
  inhaltlich angefasst; kein Sensor auf „korrigiert eine Repo-Aussage" wurde angelegt.
- §2-Ausgangsmessung: **geprüft, ohne Befund** — „vier zeigerfreie aktive Einträge,
  zwei an der ID-Schema-Familie" gegen die aktive Tabelle vor dem Range
  (MR-019/020/023/024) nachgezählt; der zweite Zähler (Anker-Suche in
  `conventions.md`) ist benannt und kommt zum selben Ergebnis.
- Tabellen- und Zähl-Konsistenz: **geprüft, ohne Befund** — aktive Tabelle 8 Zeilen;
  „fünf der acht tragen einen Zeiger, drei ein `—`" zweimal verschieden nachgezählt
  (Zeilen-Scan und Zell-Klassifikation: MR-014/016/025/027/028 gegen MR-019/024/029);
  MR-020/MR-023 stehen in der aufgelösten Tabelle mit Ankern und „aufgelöst durch
  MR-029"; der MR-017-Zeiger folgt auf die done-Lage.
- Link-Layer: **geprüft, ohne Befund** — `make doc-check` eigener Lauf (627 Dateien,
  0 Befunde, Exit 0); keine Referenz auf die alten MR-020/023-Pfade außerhalb
  `conventions/done/`.
- Verifikations-Schicht: **geprüft, ohne Befund** — `make verify` eigener Lauf
  (Exit 0, 21 Anforderungen, 0 Waisen).
- Register-Belege: **geprüft, ohne Befund** — `evidence/` trägt genau die drei
  Vorgänge der 3×-Kette (slice-097, slice-162, slice-187).
- Vollständigkeit MR-029 gegen MR-000: **geprüft, ohne Befund** — alle sechs
  historischen Formen (`AC-FA`, `AC-QA`, `ADR-NNNN`, `MR-NNN`, `CO-NNN`, `slice-NNN`)
  plus die Beobachtungs-Kennung sind abgedeckt; die numerische Slice-Form ist
  ausdrücklich als Repo-Deklaration gegen die Baseline-Namens-Empfehlung gehalten.
- Carveout-Angabe: **geprüft, ohne Befund** — „`CO-<NNN>` (bisher ungenutzt)":
  `docs/plan/carveouts/` trägt nur die README.
- Traceability-Form: **geprüft, ohne Befund** — alle fünf Commits nennen `slice-190`
  (zudem `AC-QA-02`/`AC-QA-03`); die Scope-Regel (nur `(planning)` geregelt) ist
  nicht berührt.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 4 |
| LOW | 3 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Attestierung vor dem Vorgang · Falsche
Herkunfts-Angabe in Deklaration · Versions-Pin in generischer Deklaration ·
Pflicht-Prüfung unbelegt · Widersprüchliche Closure-Felder · Adress-Zitat im
eingefrorenen Eintrag · Falscher §-Verweiser · Message-Diff-Schere ·
Absolut-Formulierung im DoD · Deklarations-Lücke (bestehend)

## Verdikt

**Merge-blockierend:** ja — die beiden HIGH-Findings liegen an der
Closure-Substanz: F-1 (Register-Ausgang im `state.md` noch nicht auf *verkörpert*
mit Anker gezogen, die Notiz behauptet es) muss vor dem Übergang nach `done/`
vollzogen oder die Notiz zurückgenommen werden; F-2 (falsche Herkunfts-Angabe des
Bereichskürzels) steht in dem Eintrag, der künftig **die** Deklaration ist, und
muss korrigiert werden, bevor weitere Autoren ihr folgen. Die Ablösungs-Substanz
selbst — MR-029 angelegt, MR-020/MR-023 sauber aufgelöst, Tabellen und Zählungen
konsistent, Link-Layer und Verifikation in eigenen Läufen grün — ist unstrittig.
Die MEDIUM-Findings F-3/F-6 betreffen dieselbe Generisch-/Zitier-Discipline und
sollten im selben Nachzug mitgelöst werden.

**Übergabe:** Findings gehen an den Implementer (Rückkante Review → Plan bei
Plan-Defekt); die **Finding-Klassen** gehen zusätzlich in die Slice-Closure §7 und
von dort in den Zähler. Dieser Report selbst ist ein **Lauf-Beleg** — er wird
über Läufe hinweg nicht wieder gelesen und muss es nicht. Er ersetzt keine
Verifikation — DoD-/Spec-Konformität prüft der Verifier separat (Modul 11).
