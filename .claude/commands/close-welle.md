# Welle schließen (Harness)

Argument: $ARGUMENTS

Dieser Command führt die **Planner**-Rolle für die **Wellen-Closure** (Modul 6) —
sechs Schritte, jeder hinterlässt einen Beleg, keiner ein Datum. Erst wenn alle sechs
Belege vorliegen, ist die Welle *auditierbar* geschlossen.

**Vorab:** liegt keine Welle vor, läuft die Closure eines **wellenlosen Slice** —
dann tragen die Slice-Closure und `make archive-wave SLICE=<kennung> APPLY=1` die
Schritte, und die Anker lauten `seit slice-<NNN>` statt `seit welle-<NN>`.

## Kontext lesen

1. [`AGENTS.md`](../../AGENTS.md) §4–§6 und [`harness/README.md`](../../harness/README.md).
2. `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Wellen-Closure-Prozedur — on-demand,
   nie das Bundle.
3. Die Welle-Plan-Datei (`docs/plan/planning/<welle-id>.md`), insbesondere ihre
   Closure-Kriterien.

## Repo-lokale Adaptionen (a-check)

- **Repo-weiter Verifikations-Beleg:** `make ci` und `make verify` grün — ein
  separater Replay-Lauf ist hier nicht vorgesehen
  ([`MR-028`](../../harness/conventions/MR-028-welle-closure-ohne-replay-v6130-erratum.md)).
- **Archivierung per Werkzeug:** `make archive-wave WELLE=<welle-id> APPLY=1` — sie
  sammelt die Slices der Welle **und** die wellenlosen seit der letzten Closure. Die
  Operation gehört in das Werkzeug, nicht in Handarbeit; dass das Archiv vollständig
  ist, bezeugt der Archivierungs-Commit.
- **Doc-Gate:** Kennungen (`AC-*`, `ADR-*`, `MR-*`, `slice-NNN`) als Anker-Links —
  die Results-Notiz und die Roadmap werden gescannt.
- **Commit via Message-Datei** (`git commit -F <datei>`), Message mit Kennung.
- **Move ≠ Inhalt:** die Welle-Plan-Datei wandert per `git mv` nach `done/` —
  eigener reiner Move-Commit, getrennt vom Inhalt (Hard Rule §3.3).

## Die sechs Schritte (Modul 6)

1. **Trigger prüfen** — alle Slices der Welle in `done/`; `make ci` und
   `make verify` grün; die Closure-Kriterien der Welle-Datei erfüllt. Belege
   **erzeugen** (Gate-Ausgaben), nicht behaupten. Fehlt einer, schließt die Welle
   nicht.
2. **Trigger-Audit** — vier Klassen, alle vier prüfen: **Carveout** (Modul 7) ·
   **bootstrap-aware Gate** (Modul 13, Stufe hochschalten oder Carveout eröffnen) ·
   **ADR** (Modul 4, Re-Evaluierungs-Trigger — bestätigen oder Folge-ADR mit
   `supersedes`) · **Hard Rule** (Auflösungs-Trigger oder permanent). Keine fälligen
   Trigger → je Klasse eine belegte „0 offen"-Feststellung. **Ein Trigger ohne
   Wächter ist eine Absichtserklärung mit Verfallsdatum.**
3. **Lese-Schritt + Closure-Notiz** — das
   [Beobachtungs-Register](../../docs/plan/planning/observations/README.md)
   durchgehen: jeder Eintrag, dessen
   `evidence/` 3 Dateien führt, bekommt seinen Ausgang (verkörpert · geplant ·
   gestrichen) und wird Steering-Loop-Eintrag. Die Results-Notiz
   `done/<welle-id>-results.md` entsteht per `cp` aus
   `.harness/baseline/v6.13.0/templates/docs/plan/planning/welle-results.template.md`
   und wird gefüllt (geliefert · was funktionierte · was anders lief ·
   Steering-Loop-Einträge · Register-Zeiger · Folge-Slices · Verifikation). Danach
   die **drei Paarungen** prüfen: Anker · Folge-Slice · Register.
4. **Zeitdokumente archivieren** — `make archive-wave WELLE=<welle-id> APPLY=1`; die
   Results-Notiz bleibt vollständig und flach, Review-Reports bekommen keinen Stub.
5. **Wave-Self-Close-Commit** — Results-Notiz + Welle-Datei §7 +
   Roadmap-Fortschreibung in einem Commit; **danach** der reine `git mv` der
   Welle-Datei nach `done/` als eigener Commit.
6. **Roadmap fortschreiben** — Zeile unter *Abgeschlossene Wellen* (Zeiger auf die
   Results-Notiz), der Zeiger verlässt *Offene Wellen*; Umplanungen stehen in
   *Historische Trigger-Verschiebungen* — eine Schließung ist keine Umplanung.

**Merke (Modul 6):** Datum ist Output, nie Trigger. Wer die Welle am Kalendertag
schließt, kappt halbfertige Slices.
