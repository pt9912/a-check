# slice-157 — Selbst-Archivierung wird Teil des Workflows (Etappe D+1)

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Auftrag 2026-09-05 („Leg los mit der Selbst-Archivierungsroutine") — setzt
den in [slice-143](../done/wellenlos/slice-143-archivierung-delta-analyse.md) §6
vorgeschlagenen, bislang unbeauftragten Punkt „Ab Etappe D+1 (künftig): jede neue Slice-Closure
archiviert sich selbst" um. [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Workflow-Änderung, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

`AGENTS.md` §6 (Minimal Agent Workflow) um den Selbst-Archivierungs-Schritt ergänzen — so, dass
zwei manuelle Sweeps ([slice-153](../done/wellenlos/slice-153-archive-wave-welle-feld-und-titel-form.md),
[slice-156](../done/slice-156-zweiter-wellenlos-sweep.md)) künftig entfallen, weil der Puffer nie
über den gerade aktiven Slice hinaus wächst.

## 2. Wo genau, mit welcher Begründung

Baseline-Regelwerk
[`modul-06-roadmap.md`](../../../../.harness/baseline/v6.0.0/regelwerk/modul-06-roadmap.md)
§Wann Arbeit eine Welle braucht, Tabelle *Träger im Repo ohne Wellen*, Zeile „Zeitdokumente
archivieren": Träger ist die **Slice-Closure selbst**, Zeitpunkt „nach den Paarungen" — bei
a-check ohne formale Wellen-Closure-Prozedur entspricht das dem Zeitpunkt **nach grünem
`make verify`**, dem a-check-eigenen Äquivalent der Paarungs-Prüfung (DoD/Closure-Fragen).

Neuer Absatz in `AGENTS.md` §6, direkt nach dem bestehenden `make verify`-Hinweis: bei einem
**wellenlosen** Slice (kein `**Welle:**`-Feld, keine aktive Welle) sofort
`make archive-wave SLICE=<id> APPLY=1`, danach `make gates`/`make verify` auf dem archivierten
Stand erneut grün, als **eigener Commit** direkt im Anschluss an den Closure-Commit — dieselbe
Zwei-Commit-Disziplin, die [`AGENTS.md`](../../../../AGENTS.md) §3.3 für jeden `git mv` mit
Content-Rewrite ohnehin verlangt (die Archivierung IST genau das: Datei wandert, Inhalt wird zum
Stub).

**Bewusst nicht geändert:** Ein Slice mit echtem `**Welle:**`-Feld archiviert weiterhin **mit
seiner Welle** bei deren Closure (`WELLE=<id>`), nicht einzeln — a-check hat aktuell keine aktive
Welle, dieser Fall bleibt unverändert dem Wellen-Closure-Ablauf überlassen.

Die `archive-wave`-Zeile in `AGENTS.md` §4 nachgezogen: „kein Gate" bleibt korrekt (kein
`make gates`/`make ci`-Bestandteil), aber „von Hand ausgelöster Vorgang" war ab jetzt falsch —
für wellenlose Slices ist es ein **Pflichtschritt** des Workflows.

## 3. Gegenprobe: die Regel an sich selbst anwenden

Dieser Slice ist selbst wellenlos. Nach seiner Closure (`git mv` nach `done/`, `make gates`/
`make verify` grün) wird er **nach der von ihm selbst eingeführten Regel** sofort archiviert — im
selben Arbeitsgang, als zweiter Commit. Das ist die schärfste verfügbare Gegenprobe: die neue
Regel funktioniert nur, wenn sie am ersten Slice, der ihr unterliegt, tatsächlich greift.

## 4. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor — die Regel ist Workflow-Text, kein
Gate-Code.

## 5. DoD

- [x] `AGENTS.md` §6 um den Selbst-Archivierungs-Schritt ergänzt, mit Baseline-Zitat statt
      Behauptung (§2).
- [x] `AGENTS.md` §4 (`archive-wave`-Zeile) nachgezogen — „von Hand ausgelöst" durch „Pflichtschritt
      bei wellenlosem Abschluss" ersetzt (§2).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** die in [slice-143](../done/wellenlos/slice-143-archivierung-delta-analyse.md) §6
vorgeschlagene, seit Etappe D unbeauftragte Selbst-Archivierungsroutine ist jetzt Teil von
`AGENTS.md` §6 — kein Backlog mehr, der erst ein späterer Sweep aufräumt.

**Lerneintrag — Form: geschärfte Regel.** *Ein Workflow-Schritt, der nur in einer Analyse
vorgeschlagen, aber nie in die Briefing-Datei selbst geschrieben wird, bleibt folgenlos — genau
das belegen [slice-153](../done/wellenlos/slice-153-archive-wave-welle-feld-und-titel-form.md)
und [slice-156](../done/slice-156-zweiter-wellenlos-sweep.md): derselbe Bedarf wurde zweimal
beobachtet und zweimal manuell nachgeholt, weil `AGENTS.md` §6 ihn nie verlangte.* Die Lehre ist
nicht „archiviere öfter", sondern *wo eine Regel steht, entscheidet, ob sie befolgt wird* — ein
Vorschlag in einer `done/`-Analyse ist Historie, keine Handlungsanweisung; nur was im aktiven
Briefing steht, wirkt auf den nächsten Lauf (Modul-0-Prinzip,
[`harness/README.md`](../../../../harness/README.md)).

**Zwei beobachtbare Closure-Kriterien:**

1. `grep -c "Danach, wenn der Slice wellenlos ist" AGENTS.md` → `1`.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Eines notiert: *Der neue Schritt verdoppelt den
Abschluss-Aufwand jedes wellenlosen Slice (zweiter Gate-/Verify-Lauf, zweiter Commit).* **Ausgang:
weiter offen**, [Beobachtungs-Register](../observations/README.md) — kein Sensor-Kollisions-Fund
(die bestehende [`BEO-GATE`](../observations/BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig/observation.md)
deckt das nicht), sondern ein Praxis-Kosten-Risiko: ob der Mehraufwand sich lohnt, zeigt sich erst
an mehreren künftigen Slice-Closures, nicht an diesem einen.

**Folge-Slices:** keine vergeben.

## 7. Sub-Area-Modus

Berührt: **Harness-Einstieg** (`AGENTS.md`) — Greenfield: Briefing und Konventionen entstehen vor
der Regel, die sie beschreiben.
