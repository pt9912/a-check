# Review-Report: slice-205 — 2026-09-29

**Review-Art:** Plan — geprüft gegen die Slice-Ziel-Form
(`v6.13.0` · `templates/docs/plan/planning/slice.template.md`), die Hard Rules
(`AGENTS.md` §3) und die Dokumentations-Regeln (§5), gegen die Register-Form
(`v6.13.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register) sowie
adversarial gegen die Generator-Ausgabe `/tmp/aih-v6.13.0` und den
a-check-Bestand (Pfad-Behauptungen, Zahlen, MR-014-Wortlaut).

**Gegenstand:** `slice-205` (uncommittet, `in-progress/`) · Register-Zeile
`BEO-GATE/init-tool-drift` (1×) · die Folge-Slices `slice-206` und
`slice-207` (`open/`, uncommittet).

**Skill:** `.harness/skills/reviewer.md` @ Commit `2224950` (unverändert im
Working Tree) ·
**Modell:** glm-5.3-flash · **Datum:** 2026-09-29

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- `slice-205` (Hauptgegenstand), `slice-206`, `slice-207`
- `AGENTS.md` §3 (Hard Rules) und §5 (Dokumentations-Regeln)
- [`spec/lastenheft.md`](../../spec/lastenheft.md) für `AC-QA-02` (Ehrliche
  Heuristik-Grenze — im Slice als Bezug und im Lerneintrag zitiert)
- `harness/conventions.md` (Modus-Deklaration, Adaptions-Block) und
  `harness/conventions/MR-014-keine-agenten-telemetrie.md` (Wortlaut)
- Generator-Ausgabe `/tmp/aih-v6.13.0` (Kopie `/tmp/aih`) — Struktur und
  Wrapper-Verhalten verifiziert, nicht dem Slice-Zitat allein vertraut
- `v6.13.0` · `regelwerk/modul-15-observability.md` §Kernidee (als Belegquelle
  zu F-1), `v6.13.0` · `templates/docs/plan/planning/slice.template.md`

---

## Findings

### F-1 — „MR-014 richtet sich gegen Telemetrie nach draußen" — der Wortlaut des Eintrags trägt diese Lesart nicht

- `kategorie`: HIGH
- `quelle`: Wortlaut `harness/conventions/MR-014-keine-agenten-telemetrie.md`
  (Adaption-Feld); `v6.13.0` · `regelwerk/modul-15-observability.md`
  §Kernidee; `AGENTS.md` §3
- `pfad`: `docs/plan/planning/in-progress/slice-205-init-tool-vergleich.md`:57–64
  (§3.2), ebenso Zeile 47 (§3.1, Zeile Erfassungsschicht)
- `befund`: Der Plan behauptet, MR-014 „richtet sich gegen Telemetrie **nach
  draußen**" und die beiden widersprächen sich nicht — das Adaption-Feld
  verbietet aber „keine Tool-Call-Spans" **uneingeschränkt**, und genau
  Tool-Call-Spans („je Werkzeug-Aufruf … eine JSON-Zeile") erzeugt die
  Erfassungsschicht; auch die Begründung des Eintrags („Span-Telemetrie über
  einen Prozess ohne Modellaufruf misst nichts") und der Auflösungs-Trigger
  („sobald Agenten-Läufe **im Repo selbst** abrechenbar werden") kennen die
  Lokal/Draußen-Unterscheidung nicht — der Trigger beschreibt im Gegenteil
  exakt die Situation, die die Übernahme herstellt. Die Abgrenzung ist als
  Deutung des Artefakts formuliert, nicht als dessen Wortlaut, und sie ist
  Premisse des Folge-Slices `slice-206`.
- `verifizierbar`: ja — Wortlautvergleich gegen die MR-014-Datei (kein
  Gate-Lauf; `make doc-check` prüft Links, nicht Lesarten).
- `klasse`: Adaption-Wortlaut umgedeutet

### F-2 — Werkzeugpfad `.harness/state/bin/` existiert im Repo nicht; der Risiko-Ausgang „entfallen" trägt darauf

- `kategorie`: HIGH
- `quelle`: nachweislich falsche Tatsachenbehauptung gegen ein Repo-Artefakt
  (`AGENTS.md` §3.7-Umfeld; Skill-Klassifikation HIGH)
- `pfad`: `docs/plan/planning/in-progress/slice-205-init-tool-vergleich.md`:87
  (§5 Rückführung) und :97–100 (§7 Risiko-Ausgang); in `slice-206`
  (Zeilen 27–28, 87) ist der Pfad dagegen korrekt auf die **Generator-Ausgabe**
  bezogen („des Generators")
- `befund`: In a-check existiert `.harness/state/bin/` nicht —
  `.harness/state/` trägt nur `gates-passed.diffsha`; der Träger-Ort liegt so
  ausschließlich in der Generator-Ausgabe (`/tmp/aih-v6.13.0/.harness/state/bin/ai-harness-init`)
  und wäre damit selbst Teil des Wegwerf-Artefakts, dessen Verlust §5
  abhandelt. Das tatsächlich installierte Werkzeug liegt als Host-Binary
  (linuxbrew, Version 0.2.5) außerhalb des Repo. Damit ist auch der
  Risiko-Ausgang „*entfallen*: das Werkzeug liegt im Bestand" unbelegt — der
  Bestand trägt das Werkzeug nicht, und §5 ist dadurch falsch verknüpft.
- `verifizierbar`: ja — `ls .harness/state/` und `find .harness -type d -name bin`
  im Repo-Baum.
- `klasse`: Werkzeugpfad behauptet statt geprüft

### F-3 — Makefile-Zeilenanzahl „630+" falsch gegen den Bestand

- `kategorie`: HIGH
- `quelle`: nachweislich falsche Tatsachenbehauptung gegen ein Repo-Artefakt
- `pfad`: `docs/plan/planning/in-progress/slice-205-init-tool-vergleich.md`:50
  (§3.1, Zeile `harness/mk/`)
- `befund`: Die Ablehnung begründet mit „a-checks monolithisches Makefile ist
  etabliert (630+ Zeilen, gewachsen)" — das Makefile hat 259 Zeilen; selbst die
  Summe aller Build-Fragmente (Makefile 259 + `d-check.mk` 82 + Dockerfile 77 +
  `a-check.mk` 36) kommt auf 454. Die Ablehnung selbst trägt („monolithisch"
  und „gewachsen" stimmen), die Zahl ist gegen kein Artefakt haltbar.
- `verifizierbar`: ja — `wc -l Makefile`.
- `klasse`: Ungeprüfte Größenangabe

### F-4 — Review-DoD-Haken attestiert einen Report, der zum Prüfzeitpunkt nicht existiert

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §3.7 (Zustandsfelder); Register-Klasse
  `BEO-GATE/attestierung-vor-dem-vorgang` (4×, `slice-204` als Träger)
- `pfad`: `docs/plan/planning/in-progress/slice-205-init-tool-vergleich.md`:68–74
  (§4, alle vier DoD-Punkte angehakt)
- `befund`: Der DoD-Punkt „Unabhängiger Review, Report unter
  `docs/reviews/`" ist im `in-progress/`-Stand angehakt, während kein Report
  zu `slice-205` existiert — der Report entsteht erst mit diesem Lauf. Die
  Bestands-Praxis hält die Häkchen bis zur Closure offen (der
  `in-progress`-Stand von `slice-189` führt sämtliche DoD-Punkte ungehakt);
  die Klasse ist im Register bereits 4× belegt und mit `slice-204` als
  Folge-Slice besetzt.
- `verifizierbar`: ja — `ls docs/reviews/` ohne Treffer für 205; `git show`
  des `in-progress`-Commits von `slice-189`.
- `klasse`: Review-Haken ohne Report

### F-5 — „update oder Nachfolge" für MR-014: „update" widerspricht der Immutabilitäts-Disziplin

- `kategorie`: MEDIUM
- `quelle`: `harness/conventions.md` §Adaptions-Block (Disziplin); `AGENTS.md`
  §3.5 sinngemäß (Analogie im Wortlaut der Disziplin selbst)
- `pfad`: `docs/plan/planning/open/slice-206-erfassungsschicht.md`:65 (Plan-Tabelle,
  „update oder Nachfolge"); `slice-205` Zeile 62–63 („Nachfolge-Eintrag oder
  Zusatz")
- `befund`: Ein Accepted-Eintrag wird nie inhaltlich überschrieben —
  „Korrekturen entstehen als neuer `MR` oder als ausdrückliche Aufhebung";
  die Option „update" im Umsetzungs-Plan liegt außerhalb dieser geschlossenen
  Menge, und „Zusatz" (slice-205) ist ebenso kein Pfad der Disziplin.
  `slice-206` benennt andernorts korrekt „Nachfolge-Eintrag" — die Tabelle
  lässt aber den regelwidrigen Weg offen.
- `verifizierbar`: ja — Wortlaut `harness/conventions.md` §Adaptions-Block.
- `klasse`: Immutabilitäts-Pfad offengelassen

### F-6 — Sub-Area-Wahl in slice-206 unvollständig: `harness/`-Berührung ist HARNESS, nicht GATE

- `kategorie`: MEDIUM
- `quelle`: `harness/conventions.md` §Modus-Deklaration pro Sub-Area
  (HARNESS = `AGENTS.md`, `CLAUDE.md`, `harness/`)
- `pfad`: `docs/plan/planning/open/slice-206-erfassungsschicht.md`:94–101 (§8)
- `befund`: `slice-206` ändert neben `.claude/`-Dateien auch
  `harness/erfassung-feldliste.md` und den MR-014-Eintrag unter
  `harness/conventions/` — beide liegen in der Pfad-Familie HARNESS; §8 prüft
  nur GATE, und der Modus-Begründungsblock nennt nur eine Sub-Area.
- `verifizierbar`: ja — Pfad-Familien gegen die Modus-Tabelle.
- `klasse`: Sub-Area-Berührung unvollständig

### F-7 — Closure-Notiz verweist auf „Risiken aus §6" — die Risiken liegen in §7

- `kategorie`: LOW
- `quelle`: `v6.13.0` · `templates/docs/plan/planning/slice.template.md`
  (dort ist Risiken §6); Planning-README §Beim Kopieren, Punkt 4 („Nicht die
  Nummer kopieren")
- `pfad`: `docs/plan/planning/in-progress/slice-205-init-tool-vergleich.md`:136 (§8)
- `befund`: Der Plan schiebt Analyse-Abschnitte ein (repo-üblich) und verschiebt
  die Risiken nach §7 — den Form-Verweis „Risiken aus §6 … siehe §6" hat er
  trotzdem aus der Vorlage übernommen; `slice-189` hatte dieselbe Stelle
  korrekt auf seine Risiken-Sektion gezogen.
- `verifizierbar`: ja — Textvergleich Plan gegen Vorlage.
- `klasse`: Abschnitts-Nummer aus der Vorlage kopiert

### F-8 — Platzhalter-Syntax „{@placeholder}" falsch zitiert

- `kategorie`: LOW
- `quelle`: unbelegte Detail-Behauptung (Gegenstand /tmp-Ausgabe)
- `pfad`: `docs/plan/planning/in-progress/slice-205-init-tool-vergleich.md`:52 (§3.1,
  Zeile Generiertes Spec-Stratum)
- `befund`: Die Generator-Straten tragen `<Projektname>`-/`YYYY-MM-DD`-Platzhalter;
  die Syntax `{@…}` kommt in `/tmp/aih-v6.13.0/spec/` in keinem der drei Straten
  vor. Die Entscheidung (Ablehnung, Re-Bootstrap wäre ein Reset) bleibt von der
  Korrektur unberührt.
- `verifizierbar`: ja — `grep -rn "{@" /tmp/aih-v6.13.0/spec/` (leer).
- `klasse`: Zitier-Detail falsch zitiert

### F-9 — d-check-Version „v0.78.0" zitiert, das Repo pinnt v0.79.0

- `kategorie`: LOW
- `quelle`: `d-check.mk` (DCHECK_IMAGE `d-check:v0.79.0`)
- `pfad`: `docs/plan/planning/in-progress/slice-205-init-tool-vergleich.md`:53 (§3.1,
  Zeile file-Modul)
- `befund`: Der Plan zitiert das `file`-Modul am Stand `v0.78.0`, während der
  Bestand `v0.79.0` pinnt; die verifizierbare Hälfte der Zeile ist richtig
  (das Modul existiert — sichtbar an den `--disable file`-Flags in
  `d-check.mk` — und steht nicht in `modules:`), nur die Versions-Attribution
  geht am Pin vorbei.
- `verifizierbar`: ja — `d-check.mk`, `modules:`-Zeile in `.d-check.yml`.
- `klasse`: Version-Zitat am Pin vorbei

### F-10 — Tippfehler im Beobachtungs-Text friert ein

- `kategorie`: LOW
- `quelle`: Register-Form (`v6.13.0` · `regelwerk/modul-06-roadmap.md`
  §Das Beobachtungs-Register — `observation.md` ist „unveränderlich ab Anlage")
- `pfad`: `docs/plan/planning/observations/BEO-GATE/init-tool-drift/observation.md`:9
- `befund`: „ein VergleichSlice vergleicht" — fehlendes Leerzeichen/Trennzeichen;
  da `observation.md` ab Anlage unveränderlich ist, würde der Fehler mit der
  Zeile eingefroren.
- `verifizierbar`: ja — Text.
- `klasse`: Tippfehler im einfrierenden Artefakt

## Negativbefunde

| Lense / Bereich | Ergebnis |
|---|---|
| **Harness-Lügen** — Struktur-Behauptungen gegen `/tmp/aih-v6.13.0` und den Bestand: `.claude/hooks/span-emit.sh`, `.claude/agents/` (sechs Rollen), `.claude/commands/{plan-welle,implement-slice,close-welle}`, `harness/mk/` (exakt 10 Module), `tools/harness/`, `harness/erfassung-feldliste.md`; `make gates`/`make verify` existieren als Targets; IDs (`AC-QA-02`, `MR-014`, `BEO-GATE/init-tool-drift`, `slice-206/207`) lösen auf | Befunde F-2, F-3, F-8, F-9; alle übrigen Behauptungen geprüft, ohne Befund |
| **Kommentar-/Zustandsfeld-Regeln** (`AGENTS.md` §3.7) — `state.md` „Stand: offen (1×)" entspricht der Bestands-Praxis; Closure-Notiz §8 trägt Zustand statt Chronik | Befunde F-4 (vorgezogenes DoD-Häkchen), F-10; übriger Bestand geprüft, ohne Befund |
| **Form** — Kopf-Felder, §1 Out-of-Scope mit Klassen-Begründung (beide Punkte: „Es wäre ein anderer Vorgang" · „Schicht-Abgrenzung"), Risiko-Ausgänge aus §7 in der geschlossenen Dreier-Menge (entfallen mit Begründung · weiter offen → Register), Delta-Zählung 7 = 3 Übernahmen + 4 Ablehnungen konsistent, §9 zwei vorgelagerte Prüfungen + GF-Hinweis | Befund F-7; übrige Form geprüft, ohne Befund |
| **Referenzen/Links** — relative Links aus `in-progress/`/`open/`, Anker `#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze` und `#mr-014`, Zitier-Form Baseline als Tag + Pfad in Inline-Code (`v6.13.0`) im Plan | geprüft, ohne Befund |
| **MR-014-Abgrenzung (§3.2)** — Wortlaut, Begründung und Auflösungs-Trigger des Eintrags gegen die Plan-Lesart; modul-15 §Kernidee als Baseline-Beleg | Befunde F-1, F-5 |
| **Rollen-Trennung / Sub-Area §9** — GATE-Achsen 1,2,3 und Modus GF gegen `harness/conventions.md`; Berührungspfad-Familien der drei Pläne gegen die Modus-Tabelle | Befund F-6 (slice-206); slice-205/slice-207 geprüft, ohne Befund |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 4 |
| MEDIUM | 2 |
| LOW | 4 |
| INFO | 0 |

**Finding-Klassen dieses Laufs:** Adaption-Wortlaut umgedeutet · Werkzeugpfad
behauptet statt geprüft · Ungeprüfte Größenangabe · Review-Haken ohne Report ·
Immutabilitäts-Pfad offengelassen · Sub-Area-Berührung unvollständig ·
Abschnitts-Nummer aus der Vorlage kopiert · Zitier-Detail falsch zitiert ·
Version-Zitat am Pin vorbei · Tippfehler im einfrierenden Artefakt

**Geltungsbereich des Laufs:** geprüft sind der Slice-Plan `slice-205`, die
neue Register-Zeile `BEO-GATE/init-tool-drift` und die Folge-Slice-Pläne
`slice-206`/`slice-207` im **uncommitteten** Working Tree — gegen die
Generator-Ausgabe unter `/tmp/aih-v6.13.0` (Struktur, Wrapper-Quelltext,
settings.json) und die Repo-Artefakte. Nicht geprüft: die übrigen,
nicht referenzierten Register-Einträge; das Verhalten des Traegers selbst über
einen instrumentierten Lauf hinaus (nur Quelltext-Lesart); ob `make gates` auf
dem uncommitteten Stand grün ist (Verifikations-Frage, Modul 11).

## Verdikt

**Merge-blockierend:** ja — die vier HIGH-Findings: die MR-014-Lesart (F-1)
ist am Wortlaut des Eintrags nicht haltbar und ist Premisse von `slice-206`;
der Risiko-Ausgang §7 (F-2) und die Zeilenzahl (F-3) sind gegen Repo-Artefakte
falsch; der Review-DoD-Haken (F-4) attestiert einen Report, der erst durch
diesen Lauf entsteht, und wiederholt damit eine im Register 4× belegte Klasse.
Kein Gate fängt diese Befunde (`make doc-check` prüft Links und Anker, nicht
Zahlen und Lesarten) — die Korrektur ist vor der Closure zu leisten: §3.2 und
§7 auf den MR-014-Wortlaut stellen, §5/§7 auf den realen Werkzeug-Stand
(host-brew, Version) korrigieren, die DoD-Häkchen auf den Ist-Stand zurücksetzen,
„update" aus der Umsetzungs-Tabelle von `slice-206` nehmen.

**Übergabe:** Findings gehen an den Implementer (Rückkante Review → Plan bei
Plan-Defekt); die **Finding-Klassen** gehen zusätzlich in die Slice-Closure §7
und von dort in den Zähler — `BEO-GATE/attestierung-vor-dem-vorgang` trägt
F-4 als fünftes Auftreten der bestehenden Klasse. Dieser Report selbst ist ein
**Lauf-Beleg** (Audit: dieser Working-Tree-Stand, dieser Skill, dieses Modell,
dieses Verdikt) — er wird über Läufe hinweg nicht wieder gelesen, und muss es
nicht. Der Report ersetzt keine Verifikation — DoD-/Spec-Konformität prüft der
Verifier separat (Modul 11; anderes Prüf-Artefakt, anderer Eingabe-Kontext).
