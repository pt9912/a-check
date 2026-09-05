# Review-Report: slice-135…157 (v6.0.0-Migration + Nachfolgearbeiten) — 2026-09-05

**Review-Art:** Code + Design — geprüft gegen Plan/ADR/Spec-Konsistenz **und** Konventionen
(Modul 10 §Drei Review-Arten, hier kombiniert wegen vier paralleler Linsen).

**Gegenstand:** `slice-135`…`slice-157` (Commits `e589554^`…`629bdba`) — Baseline-Migration
v5.12.0→v6.0.0, `archive-wave`-Werkzeugfixes, Spec-Nachträge (`spec/lastenheft.md`,
`spec/architecture.md`), Selbst-Archivierungsroutine (`AGENTS.md` §6).

**Skill:** `.harness/skills/reviewer.md` @ Stand 2026-09-05, angewandt über vier unabhängige,
**kontext-getrennte** Subagenten (Modul 8 §Kontext-Trennung: keiner der vier hat den geprüften
Code/Text selbst geschrieben) statt eines einzelnen Laufs.

**Modell:** claude-sonnet-5 (Implementer-Kontext) + vier unabhängige Subagenten-Läufe (drei
`general-purpose`, einer `code-documentation:code-reviewer`) · **Datum:** 2026-09-05.

**Auslöser:** Maintainer-Frage „Warum ist der reviews-Ordner leer? Machen wir keine Reviews
mehr?" — real: seit `slice-127` (2026-08-15) fand kein unabhängiges Review mehr statt, weil die
Reviewer-Rolle in dieser Session nie über einen getrennten Kontext ausgelöst wurde (siehe
Finding F-6 unten — das Muster ist selbst ein Fund dieses Reports).

**Eingangs-Kontext:**

- `AGENTS.md` (Hard Rules, §6 Minimal Agent Workflow)
- `docs/reviews/README.md` (Review-Pflicht „ab slice-043 … eigener DoD-Punkt")
- `.harness/skills/reviewer.md`
- `harness/conventions.md` (MR-Bestand)
- geprüfte Slices selbst (aus den Archiv-`.zip`-Dateien extrahiert, wo bereits archiviert)

**Vier Linsen:** (1) Code-Korrektheit `tools/archive-wave/` — **noch laufend**, Folge-Report
folgt als eigene Datei (Modul-10-Konvention: ein Report pro Lauf). (2) Vertrag/Spec-Konsistenz
`spec/lastenheft.md`+`spec/architecture.md`. (3) Regelwerk/Konvention — Stichprobe über 17 der 23
Slice-Closures. (4) Test-Abdeckung der fünf neuen `archive-wave`-Regressionstests.

---

## Findings

### F-1 — Zwei neue technische Behauptungen in `spec/architecture.md` trotz „kein neuer Fakt"

- `kategorie`: HIGH
- `quelle`: Closure-Anspruch `slice-154` §3/§4 („kein neuer Fakt behauptet")
- `pfad`: `spec/architecture.md:112` (ARC-008 „austauschbar … Pin-Mechanik bindet an den Digest,
  nicht an die Registry-Identität"); `spec/architecture.md:168-171` („kein Teilergebnis, kein
  Retry … terminal für den Lauf")
- `befund`: Beide Aussagen stehen an keiner anderen Stelle in `lastenheft.md`/`architecture.md`
  (repo-weiter `grep` nach „Retry"/„Teilergebnis"/„terminal"/„austauschbar" liefert nur diese
  Fundstellen) — echte neue architektonische Festlegungen, keine Konsolidierung bestehender
  Aussagen.
- `verifizierbar`: ja — `grep -rn "Retry\|Teilergebnis\|terminal" spec/` bzw. `grep -rn
  "austauschbar" spec/` gegen den Rest des Dokuments.
- `klasse`: „kein-neuer-Fakt"-Closure-Anspruch nicht gegen den eigenen Text geprüft

### F-2 — Glossar-Eintrag „Composition Root" führt unbelegte Kürzel-Ergänzung

- `kategorie`: MEDIUM (herabgestuft von der Subagent-Einstufung HIGH — Gegenprobe unten)
- `quelle`: `spec/lastenheft.md` §6 Glossar
- `pfad`: `spec/lastenheft.md:715`
- `befund`: Der Eintrag definiert Composition Root als „Ort, an dem Verdrahtung (DI) stattfinden
  darf". Das Konzept selbst ist über `spec/architecture.md`s wiederholtes „verdrahtet Adapter an
  den Kern"/„verdrahtet konkrete Adapter an die Ports" (ARC-006, §3) gedeckt — die
  Kürzel-Ergänzung „(DI)" (Dependency Injection) steht dagegen an keiner Stelle im Repo sonst.
  **Gegenprobe:** der ursprüngliche Subagent-Fund (HIGH, „komplett unbelegt") prüfte nur auf das
  Substantiv „Verdrahtung", nicht auf die Verbform „verdrahtet" — dadurch ein Teil-Fehlalarm;
  der Kern-Begriff ist belegt, nur „(DI)" nicht.
- `verifizierbar`: ja — `grep -n "verdraht" spec/architecture.md` zeigt drei Treffer.
- `klasse`: Glossar-Definition ergänzt Fachbegriff-Kürzel ohne Repo-Beleg

### F-3 — Faktenfehler „16 aktive MR-Einträge" in zwei `done/`-Slices

- `kategorie`: HIGH
- `quelle`: `slice-135` §4.1/§6/§8, `slice-136` §4/§8
- `pfad`: `harness/conventions.md` §Aktive Adaptionen (aktuell: sechs — MR-011/012/014/015/016/017)
- `befund`: Beide Slices behaupten „16 aktive MR-Einträge" als gemessenen Fakt und leiten daraus
  den Etappe-B-Zuschnitt sowie einen Risiko-Ausgang ab. 16 ist die höchste je vergebene
  MR-Kennung, nicht die Zahl aktiver Einträge — zehn liegen bereits in `conventions/done/`.
- `verifizierbar`: ja — `grep -c` gegen `harness/conventions.md` §Aktive Adaptionen.
- `klasse`: höchste-vergebene-Kennung mit aktiver-Bestand-Zahl verwechselt

### F-4 — Dasselbe Muster dreifach: als vollständig behauptete Prüfung war unvollständig, extern gefangen

- `kategorie`: HIGH
- `quelle`: Steering-Loop-Regel (`AGENTS.md`: 2. Vorfall ⇒ Registereintrag, 3. ⇒ Harness-Lücke)
- `pfad`: `slice-142` (vier `.claude/rules/`-Symlinks von `grep`-Bereinigung übersehen);
  `slice-135`/`139`→`143` (Zeitdokumente-Archivierung fälschlich „optional" — nur die
  Wellen-Hälfte der Baseline-Regel gelesen, `slice-143` §1 räumt das wörtlich ein);
  `slice-150`→`151` („Offene Slices ohne Welle"-Tabelle nach eigener Prüfung behalten, vom
  Maintainer korrigiert)
- `befund`: In allen drei Fällen erklärte dieselbe Session eine Prüfung für vollständig, die es
  nicht war — und in allen drei Fällen war es der Maintainer, der die Lücke fing, nicht ein
  interner Blick. Die Session benennt die Parallele sogar selbst (`slice-143` zieht sie explizit
  zu `slice-142`), aber **kein** Beobachtungs-Register-Eintrag trägt dieses Muster, obwohl die
  eigene Steering-Loop-Regel das ab dem zweiten Vorfall verlangt.
- `verifizierbar`: ja — Zitate in den drei genannten Dokumenten, Register-Verzeichnis
  `docs/plan/planning/observations/` zeigt keinen passenden Eintrag.
- `klasse`: als-vollständig-behauptete-Prüfung war unvollständig, extern statt intern gefangen

### F-5 — Risiko ohne der drei zulässigen Ausgänge

- `kategorie`: MEDIUM
- `quelle`: Modul 5 §Offene Risiken werden bei Closure aufgelöst
- `pfad`: `slice-152`, Abschnitt „Offene Risiken und ihr Ausgang"
- `befund`: Der Text benennt einen zweiten Fund explizit als offen („trägt keinen eigenen
  Ausgang — er ist durch den belassenen Puffer umgangen, nicht behoben") und lässt ihn ohne einen
  der drei laut Modul 5 verlangten Ausgänge (Carveout/Folge-Slice · gestrichen mit Begründung ·
  Beobachtungs-Register) stehen. Offen benannt, aber formal ein Verstoß.
- `verifizierbar`: ja — Volltext in `slice-152` (archiviert, per `unzip -p
  docs/plan/planning/done/wellenlos/slice-152-archiv.zip … ` lesbar).
- `klasse`: Risiko-Ausgang außerhalb der geschlossenen Dreier-Menge

### F-6 — Reviewer-Rolle nie über getrennten Kontext ausgelöst (Ursprungs-Fund dieses Reports)

- `kategorie`: HIGH
- `quelle`: `docs/reviews/README.md` („ab slice-043 … Slice-DoD führt sie als eigenen Punkt"),
  `.harness/skills/reviewer.md` §Kontext-Trennung, Modul 8
- `pfad`: alle 23 Slices `slice-135`…`slice-157`
- `befund`: Keiner dieser Slices hat einen Review-DoD-Punkt oder einen Review-Report. Ursache:
  Planung, Implementierung und Verifikation liefen die ganze Session über im selben
  Kontextfenster; die Reviewer-Rolle wurde nie über einen echten frischen Subagenten ausgelöst.
  Dieser Report selbst ist die Korrektur — die vier Linsen laufen zum ersten Mal wieder
  kontext-getrennt.
- `verifizierbar`: ja — kein Review-DoD-Punkt in einem der 23 Slice-Dokumente, keine
  `docs/reviews/*.md`-Datei zwischen 2026-08-15 und diesem Report.
- `klasse`: Reviewer-Rolle nicht kontext-getrennt ausgelöst

## Negativbefunde

- geprüft, ohne Befund: Verifikations-Behauptungen in 15 von 17 gestichprobten Slice-Closures
  (konkreter Befehl/Befund im selben Dokument, stichprobenartig gegen den heutigen Repo-Stand
  nachvollzogen — passt).
- geprüft, ohne Befund: „kein neuer Fakt"-Anspruch in `slice-136`/`139`/`146`/`150` (im
  Unterschied zu F-1, das nur `slice-154` betrifft).
- geprüft, ohne Befund: Lerneintrag-Form in allen 17 gestichprobten Slices — durchweg eine der
  drei zulässigen Formen, inhaltlich substantiell.
- geprüft, ohne Befund: Slice-zu-Slice-Widersprüche wurden handwerklich sauber aufgelöst
  (`done/`-Immutabilität gewahrt, Korrektur im Folge-Slice mit Zitat).
- geprüft, ohne Befund: Anker-/Link-Auflösung in `spec/lastenheft.md`/`spec/architecture.md`
  (`make doc-check` grün) und Umnummerierungs-Folgen der `architecture.md`-Sektionen (kein
  Repo-Fund verweist auf die verschobenen alten §4/§5/§6).
- geprüft, ohne Befund: die fünf neuen `archive-wave`-Regressionstests fangen ihren jeweiligen
  Bug nachweisbar (gedanklicher Rollback bricht die jeweilige Assertion); kein bestehender Test
  beschädigt.
- notiert, nicht als eigener Fund gewertet: Testabdeckungs-Lücken an den Rändern (Mehrfach-Trenner
  im Titel, kein End-to-End-Nachweis des Doppel-Gedankenstrich-Fixes, keine inhaltliche
  Gegenprobe „Wellen-Modus-Titel ohne Link", ungesicherte „ohne Welle"-Invariante in `Apply()`
  ohne Guard) — vier MEDIUM-Beobachtungen der Test-Abdeckungs-Linse, keine davon blockierend.
- notiert, nicht als eigener Fund gewertet: Spaltenbenennung „Kennung" (§2) vs. „ID" (§4) in
  `spec/architecture.md` — inkonsistent, aber kosmetisch (LOW).
- notiert, nicht als eigener Fund gewertet: `**Archiviert:**`-Platzhalter `<manuell
  auszufuellen>` durchgängig unausgefüllt in 153 Stub-Dateien — bereits in `slice-147`/`148` als
  bewusste, mit `d-check` konsistente Präzedenz benannt.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 4 |
| MEDIUM | 2 |
| LOW | 0 |
| INFO | 0 |

**Finding-Klassen dieses Laufs:** kein-neuer-Fakt-Closure-Anspruch nicht gegen den eigenen Text
geprüft · Glossar-Definition ergänzt Fachbegriff-Kürzel ohne Repo-Beleg · höchste-vergebene-
Kennung mit aktiver-Bestand-Zahl verwechselt · als-vollständig-behauptete-Prüfung war
unvollständig, extern statt intern gefangen · Risiko-Ausgang außerhalb der geschlossenen
Dreier-Menge · Reviewer-Rolle nicht kontext-getrennt ausgelöst

## Verdikt

**Merge-blockierend:** entfällt — die geprüften Slices sind bereits committet (lokal, nicht
gepusht); dieser Report wirkt als Korrektur-Auftrag für Folge-Slices, nicht als Merge-Gate.

**Übergabe:** F-1 bis F-5 gehen als Korrektur-Slices an den Implementer-Kontext (F-3 als
Faktenfehler-Korrektur ohne `done/`-Historie anzufassen, per Präzedenz `slice-151`-Muster). F-6
ist bereits mit `AGENTS.md` §6 (`slice-157`) strukturell behoben — dieser Report ist der erste
Beleg, dass die Korrektur trägt. Die sechs Finding-Klassen gehen in die Closure-Notiz des
Korrektur-Slice und von dort ins Beobachtungs-Register — F-4 selbst benennt die noch fehlende
Register-Verkörperung des eigenen Musters, das dieser Report jetzt zum ersten Mal einträgt.
