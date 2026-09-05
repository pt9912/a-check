# slice-148 — Etappe D: Altbestand (93 Slices) + 6 Reviews archiviert

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Etappe **D** aus [slice-143 §6](../done/slice-143-archivierung-delta-analyse.md#6-vorschlag-vier-etappen),
per Maintainer-Wort gezogen 2026-09-05, letzte der vier Etappen nach
[slice-144](../done/slice-144-archive-wave-werkzeug.md) (A),
[slice-145](../done/slice-145-archivierung-sensor-geltungsbereich.md) (B),
[slice-147](../done/slice-147-archivierung-etappe-c.md) (C). [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Planungs-Harness-Pflege, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

Den in [slice-143 §5](../done/slice-143-archivierung-delta-analyse.md#5-was-diese-analyse-nicht-geklärt-hat)
offen gelassenen Rest archivieren: die zwölf frühen Prosa-Wellen (`welle-00`…`welle-11`, ohne
formale `-results.md`) und alle übrigen wellenlosen Slices seit `welle-13`, plus die
eigenständigen Review-Reports ohne Slice-Partner.

## 2. Die offene Entscheidung aus §5 — beantwortet durch Präzedenz, nicht Annahme

„Sammel-Archiv oder einzeln wellenlos" war ausdrücklich offen gelassen. `d-check`s eigener
Bestand (`docs/plan/planning/done/wellenlos/`, 52 Slices, je ein eigenes `<id>-archiv.zip`)
kennt **kein** Sammel-Archiv — jeder wellenlose Slice bekommt sein eigenes. `archive-wave` selbst
hat dafür keinen Batch-Modus (`-slice=` nimmt genau eine Kennung). Entscheidung: **einzeln**,
wortgleich zu `d-check`s Praxis — kein neuer Mechanismus erfunden, keiner gebraucht.

## 3. Scope-Abgrenzung: 93 von 106, nicht mehr

**Ausgenommen: die 13 Slices dieser laufenden Migration selbst (`slice-135`…`147`).** Sie tragen
zwar auch kein `**Welle:**`-Feld, sind aber die **aktive Erzählung** der gerade laufenden Arbeit —
dicht wechselseitig verlinkt (`slice-146` zitiert `slice-143`/`144`/`145`, `slice-147` zitiert
`slice-143`, …). Archivierung ist für Zeitdokumente gedacht, „aus deren Volltext kein lesender
Knoten mehr zitiert" (Modul 6) — das trifft hier nicht zu. Sie bleiben liegen; ein künftiger Pass
(oder die ab §7 beschriebene Routine) holt sie nach, sobald sie selbst zu Altbestand geworden sind.

**Zwei Archivierungs-Modi, nach vorhandenem `**Welle:**`-Feld unterschieden:**

- **Wave-Modus (`-welle=`), 10 Aufrufe, 16 Slices:** Slices, die trotz fehlender `-results.md`
  bereits ein `**Welle:**`-Feld tragen (`welle-01`…`welle-10`, je 1–4 Mitglieder, numerisch über
  `welleIDInFieldRE` gefunden). `Apply()` behandelt eine Welle ohne Plan-Datei bereits korrekt
  (Kommentar in `archive.go`: „`welle-60..66`, vor der Plan-Datei-Konvention") — kein Sonderfall,
  keine Tool-Änderung nötig.
- **Slice-Modus (`-slice=`), 77 Aufrufe:** alles ohne `**Welle:**`-Feld, `slice-017` als Kanarienvogel
  zuerst, dann 76 im Batch (Docker-Aufruf pro Slice, sequenziell, `run_in_background` wegen
  Laufzeit > 120 s).
- **Review-Modus (`-review=`), 6 Aufrufe:** eigenständige Reviews ohne `slice-NNN` im Dateinamen
  (`benutzerhandbuch`, `ports-domaenen-typen-adr-0008` + `-rereview`, `nachtrag-etappe-c-f2`,
  `welle-12-unabhaengig`, `v0170-go-kern`) — die restlichen 19 alten Reviews tragen ein `slice-NNN`
  im Namen und wurden automatisch mit ihrem Slice eingesammelt (`CollectReviews`).

## 4. Vorab behoben: ein Anker-Bruch, der nicht erst am Gate auffiel

Vor dem ersten Wave-Lauf geprüft (nicht angenommen): trägt irgendein **lebendes** Dokument einen
`#anchor`-Link in einen Slice, der archiviert wird? `RewriteRepo` zieht nur den **Pfad** nach, das
`#fragment` bleibt unverändert (`rewrite.go`-Kommentar: „Ein Fragment wird abgetrennt und separat
zurückgegeben"). Ein Stub trägt aber keine `##`-Überschriften mehr — ein bestehender Anker wie
`#7-closure-notiz-nach-done` liefe ins Leere.

Repo-weiter `grep` nach `slice-NNN.md#…`-Links fand **neun** Treffer, alle in
[`roadmap.md`](../in-progress/roadmap.md) §Abgeschlossene Wellen (Closure-Belege für
`welle-01`…`welle-11`: `slice-001`, `002`, `003`, `004`, `005`, `006`, `008`, `022`, `023`). Alle
übrigen Treffer zitierten **von** einem Alt-Slice **zu** einem anderen Alt-Slice — beide werden
archiviert, die Zitat-Zeile verschwindet mit dem Stub des Ziters selbst (self-resolving). Vor dem
ersten `APPLY=1`-Lauf neun Anker-Links in `roadmap.md` auf bare File-Links umgestellt (das `§N`
bleibt als Prosa-Hinweis stehen, ist aber kein Link-Fragment mehr).

## 5. Ausführung — zwei Tool-Funde, real gemessen

Kanarienvögel zuerst (`welle-01`, `slice-017`, `benutzerhandbuch`-Review), je mit `make gates`
geprüft, bevor der Rest im Batch lief. Nach dem vollen Batch (93 Slices + 6 Reviews) brach
`make gates` an zwei Stellen — keine davon vorab erkannt:

1. **`docs/reviews/README.md`** zitierte drei Reviews (`slice-042`/`043`/`044`) per Markdown-Link
   — die drei wurden mit ihrem jeweiligen Slice eingesammelt und **ohne** Stub gelöscht (Kanon:
   Reviews im Slice-/Wellen-Modus haben keine Identität jenseits ihres Slice). Behoben: die drei
   Links zeigen jetzt auf den jeweiligen Slice-Stub unter `done/wellenlos/`, mit einem Satz, der
   erklärt, wo der Review-Volltext jetzt liegt.
2. **`ApplyReview`-Bug (Code, `tools/archive-wave/archive.go`).** Ein Review-Titel kann selbst
   einen Markdown-Link tragen (gemessen am [ADR-0008](../../adr/0008-ports-duerfen-domaenen-typen-referenzieren.md)-Review,
   dessen Überschrift die ADR per eingebettetem Link zitierte). `ExtractFullHeading` übernahm ihn
   unverändert in den Stub
   — der Link blieb von der ALTEN Position aus geschrieben und brach an der neuen, tieferen
   `ReviewArchiveDir` (`docs/reviews/archiv/` liegt eine Ebene tiefer als `docs/reviews/`).
   Behoben: `ApplyReview` schickt den Titel jetzt durch `RewriteFieldForMove` (dieselbe Funktion,
   die das Welle-Feld eines Slice-Stubs schon depth-korrekt schreibt), Test
   `TestRunReview_Apply_TitelMitLink` ergänzt. Die zwei betroffenen Stubs
   (`2026-06-22-ports-domaenen-typen-adr-0008[-rereview].md`) aus ihrem bereits geschriebenen
   `archiv.zip` wiederhergestellt und mit dem reparierten Werkzeug neu archiviert — kein manuelles
   Nachbessern des Stub-Texts.

Beide Funde bestätigen den Lerneintrag aus [slice-147 §9](../done/slice-147-archivierung-etappe-c.md#9-closure-notiz):
eine feste Fundstellen-Liste fängt nur, was sie schon kennt. Weiterer Beleg an
[`BEO-GATE`](../observations/BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig/observation.md)
(§9).

## 6. Ergebnis

- **12 Wellen-Verzeichnisse** unter `done/`: `welle-01`…`welle-10` (neu, 16 Slices), `welle-12`/
  `welle-13` (aus Etappe C).
- **`done/wellenlos/`**: 77 Slices, je mit eigenem `<id>-archiv.zip`.
- **`docs/reviews/archiv/`**: 6 eigenständige Reviews mit Stub + `archiv.zip`.
- **19 weitere Reviews** ohne eigenen Stub, mit ihrem Slice eingesammelt (Kanon: keine Identität
  jenseits des Slice).
- `docs/plan/planning/done/` trägt danach nur noch die 13 Slices der laufenden Migration
  (`slice-135`…`147`) plus `welle-12-results.md`/`welle-13-results.md` flach.
- `docs/reviews/` trägt danach nur noch `README.md`.

## 7. Was bewusst nicht getan wird

- **`slice-135`…`147` bleiben unarchiviert** — §3 begründet warum; kein Folge-Slice vergeben, weil
  kein fester Trigger benennbar ist (die Migration ist noch nicht als „alt" zu betrachten).
- **Keine Selbst-Archivierungs-Routine für künftige Closures verdrahtet.** `slice-143 §6` sieht
  „ab Etappe D+1" vor, dass jede neue wellenlose Slice-Closure sich selbst archiviert — das ist
  eine Workflow-Änderung am 8-Schritt-Ablauf ([Modul 9](../../../../.harness/baseline/v6.0.0/regelwerk/modul-09-implementierung.md)),
  kein Bestandteil dieses Slice. Bleibt benannt, nicht umgesetzt.
- **Kein Nachtrag des `<manuell auszufuellen>`-Platzhalters** — dieselbe Präzedenz wie in
  [slice-147 §7](../done/slice-147-archivierung-etappe-c.md#7-was-bewusst-nicht-getan-wird).

## 8. DoD

- [x] Scope-Frage aus slice-143 §5 (Sammel- vs. Einzel-Archiv) per Präzedenz beantwortet, nicht
      angenommen (§2).
- [x] 93 Slices (16 Wave- + 77 Slice-Modus) und 6 eigenständige Reviews real archiviert, zwei
      Tool-/Doku-Funde behoben mit Gegenprobe (§4, §5).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 9. Closure-Notiz

**Geliefert:** Etappe D — letzte der vier aus [slice-143 §6](../done/slice-143-archivierung-delta-analyse.md#6-vorschlag-vier-etappen)
— ist umgesetzt. 93 Alt-Slices und 6 eigenständige Reviews archiviert, `docs/plan/planning/done/`
und `docs/reviews/` tragen jetzt nur noch aktive bzw. jüngste Vorgänge.

**Lerneintrag — Form: geschärfte Regel.** *Eine Sensor-Vorprüfung, die nur bekannte Modul-Klassen
gegen die Stub-Form hält (§ids, §ignore-refs aus slice-147), übersieht Fundstellen, die **keine**
Sensor-Konfiguration sind, sondern Prosa in einer lebendigen Datei (`roadmap.md`s Anker-Links) oder
ein Code-Pfad im Werkzeug selbst, der bei der letzten Etappe nicht getroffen wurde (`ApplyReview`
vs. `ApplySlice`/`SliceStub`, die ihre Felder schon depth-korrekt schrieben).* Der Anker-Fund
(§4) wurde **vor** dem Gate-Bruch gefangen, weil er aus derselben Disziplin folgte, die
[slice-146](../done/slice-146-welle-feld-nachtrag.md) einführte: gemessen, nicht angenommen, per
`grep` **vor** dem ersten Schreibzugriff. Der `ApplyReview`-Fund (§5.2) wurde **nicht** vorab
gefangen — *weil* er in Code liegt, den kein `grep` über Markdown-Dateien sieht. **Zwei
verschiedene Fund-Klassen brauchen zwei verschiedene Vorab-Techniken**: Prosa-Kollisionen sind
`grep`-bar, Code-Pfade nur durch Lesen des Codes selbst oder durch den Lauf, der sie trifft.

**Zwei beobachtbare Closure-Kriterien:**

1. `ls docs/plan/planning/done/wellenlos/*.md | wc -l` → `77`;
   `ls docs/plan/planning/done/welle-*/  -d | wc -l` → `12`;
   `ls docs/reviews/archiv/*.md | wc -l` → `6`.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Eines notiert, ein Ausgang:
[`BEO-GATE`](../observations/BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig/observation.md) —
**weiter offen**, zweiter Beleg (Register-Zähler jetzt 2×): die Vorprüfungs-Methode selbst ist
noch immer eine feste, von Hand geführte Liste. Ein dritter Fund dieser Klasse macht sie zu einer
Harness-Lücke (§Steering-Loop, [`AGENTS.md`](../../../../AGENTS.md)).

**Folge-Slices:** keine vergeben. Die „Etappe D+1"-Selbst-Archivierungs-Routine aus
[slice-143 §6](../done/slice-143-archivierung-delta-analyse.md#6-vorschlag-vier-etappen) bleibt
benannt, aber nicht beauftragt — sie braucht eine eigene Kennung bei ihrer Eröffnung.

## 10. Sub-Area-Modus

Berührt zwei Sub-Areas:

- **Gate-/Werkzeug-Schicht** (`tools/archive-wave/`) — Greenfield: der `ApplyReview`-Fund ist
  getestet (`TestRunReview_Apply_TitelMitLink`), `make archive-wave-test` grün.
- **Planungs-Harness** (`docs/plan/planning/done/`, `docs/reviews/`) — Greenfield: Form und
  Größenregel stehen in der Vorlage, `doc-structure` prüft sie; die Lifecycle-Invariante
  (`links.resolve-from`) prüft `doc-check`.
