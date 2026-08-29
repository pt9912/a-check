# slice-097 — Etappe C2: die Adaptions-Urteile ausführen

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** keine `AC-*`/`ADR-*` — Ausführung der Urteile aus
[slice-095](../done/slice-095-adaptions-durchgang-v5120.md).
**Bezug:** Etappe **C** aus [slice-092 §6](../done/slice-092-regelwerk-v5120-delta-analyse.md),
zweite Hälfte. Vorgänger [slice-096](../done/slice-096-adaptions-verzeichnis-form.md).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

Etappe B hat zehn Urteile gefällt, Etappe C1 die Form gewechselt. Was fehlt, ist die Ausführung:
drei Rückbauten, fünf Ablösungen durch engere Nachfolger — und in jedem Nachfolger das
Pflichtfeld `Ersetzt-Baseline-Regel`, das in die Bestandseinträge nicht nachgetragen werden
durfte.

## 2. Betroffene Module

`harness/conventions.md` (Index) und `harness/conventions/` samt `done/`.
Eine Schicht, ein Konventionsspeicher.

## 3. Was entsteht

**Sieben** neue Einträge, nicht acht — weil die Regel *„Wer mehrere Regeln ersetzen will,
schreibt mehrere Einträge"* in beide Richtungen wirkt:

| Neu | Art | Löst auf | Ersetzt-Baseline-Regel |
|---|---|---|---|
| [`MR-010`](../../../../harness/conventions.md#mr-010) | Rückbau, direkt nach `done/` | [`MR-001`](../../../../harness/conventions.md#mr-001), [`MR-002`](../../../../harness/conventions.md#mr-002), [`MR-006`](../../../../harness/conventions.md#mr-006) | — *(löst auf, setzt nichts)* |
| [`MR-011`](../../../../harness/conventions.md#mr-011) | engerer Nachfolger | [`MR-004`](../../../../harness/conventions.md#mr-004) | `grundlagen-source-precedence.md` §ID-Schema als Klammer |
| [`MR-012`](../../../../harness/conventions.md#mr-012) | engerer Nachfolger | [`MR-005`](../../../../harness/conventions.md#mr-005) | `grundlagen-referenz-richtung.md` §Referenz-Richtung (SDP) |
| [`MR-013`](../../../../harness/conventions.md#mr-013) | engerer Nachfolger | [`MR-007`](../../../../harness/conventions.md#mr-007) | — *(korrigiert eine Repo-Aussage, keine Baseline-Regel)* |
| [`MR-014`](../../../../harness/conventions.md#mr-014) | engerer Nachfolger | [`MR-008`](../../../../harness/conventions.md#mr-008) | `modul-15-observability.md` §Kernidee |
| [`MR-015`](../../../../harness/conventions.md#mr-015) | engerer Nachfolger | [`MR-008`](../../../../harness/conventions.md#mr-008) | `modul-06-roadmap.md` §Wellen-Closure-Prozedur |
| [`MR-016`](../../../../harness/conventions.md#mr-016) | engerer Nachfolger | [`MR-009`](../../../../harness/conventions.md#mr-009) | `modul-08-agentenrollen.md` §Die neun Übergaben |

**[`MR-008`](../../../../harness/conventions.md#mr-008) bekommt zwei Nachfolger.** Sein Rest zerfällt in zwei verschiedene Baseline-Regeln:
die Telemetrie-Erwartung (`modul-15`) und das Welle-Closure-Kriterium „Replay-Lauf grün"
(`modul-06`). Ein Eintrag darf genau eine Regel ersetzen.

**Der Rückbau ist *ein* Eintrag und geht direkt nach `done/`.** `modul-02` verlangt für den
Rückbau einen *neuen Eintrag statt eines Edits* — das ist erfüllt. Ihn in die Tabelle der
**aktiven** Adaptionen zu stellen wäre dagegen falsch: er erklärt, dass das Repo **nicht** mehr
abweicht, und jeder Agentenlauf läse eine Nicht-Abweichung. Genau diese Kosten hat die
Verzeichnis-Form gerade beseitigt. Das ist eine Auslegung, keine Vorschrift — sie steht hier,
damit sie überstimmbar ist.

## 4. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Tragend ist `doc-check`.

**Die Negativ-Probe ist gelaufen, bevor der erste Eintrag entstand.** Offen war, ob ein Anker in
den vendored Baum überhaupt geprüft wird — der Baum liegt in `scan.ignore`. Gemessen: ein Link
mit **falschem** Anker dorthin ergibt `anchor-missing`, Exit 2. Das Pflichtfeld ist damit
maschinell gedeckt und nicht dekorativ; `scan.ignore` nimmt den Baum als *Quelle* aus, nicht als
*Ziel*.

## 5. Was bewusst nicht getan wird

- **Kein Bestandseintrag wird angefasst.** Die neun bestehenden Dateien bleiben Wort für Wort,
  wie sie sind; sie wandern nur nach `done/`, und das per `git mv` in eigenem Commit.
- **[`MR-013`](../../../../harness/conventions.md#mr-013) wird nicht wegdiskutiert.** Unter dem strengen Fork-Test der neuen Baseline ist er
  keine Adaption — er korrigiert eine Aussage in [`MR-000`](../../../../harness/conventions.md#mr-000), keine Baseline-Regel. Er entsteht
  hier trotzdem, weil die Alternative wäre, eine heute falsche Versionsangabe stehenzulassen.
  Als **Kandidat für Rückbau** ist er in seinem eigenen Text benannt; sauber auflösen lässt er
  sich erst, wenn die ID-Schema-Deklaration in [`MR-000`](../../../../harness/conventions.md#mr-000) selbst überarbeitet wird — was die neue
  Vorlage nahelegt, aber dieser Slice nicht leistet.
- **Keine Roadmap-Zeile, kein Beobachtungs-Register.** Etappe D.

## 6. DoD

- [x] Sieben Nachfolge-Einträge liegen unter `harness/conventions/` (der Rückbau unter `done/`),
      je mit den Pflichtfeldern inklusive `Löst auf` und `Ausgelöst durch Baseline-Stand` —
      Beleg: Verzeichnis-Zustand und Diff.
- [x] Die acht abgelösten Bestandseinträge liegen in `conventions/done/`, inhaltlich unverändert;
      der Index führt sie in der Tabelle *Aufgelöste Adaptionen* mit ihrem Nachfolger, und beide
      Anker jeder Zeile sind erhalten — Beleg: `doc-check` mit 0 Befunden bei **136**
      Alt-Slug-Verweisen im Bestand (einer weniger als vor C2: die Index-Zeile zu `ADR-NNNN`
      adressiert [`MR-000`](../../../../harness/conventions.md#mr-000) jetzt über die kurze Kennung).
- [x] `make gates` (und bei Abschluss `make verify`) grün — **Ausgabe in eine Datei**, Exit-Code
      getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

**Geliefert:** sieben Nachfolge-Einträge, acht abgelöste Einträge in `conventions/done/`, und ein
Index, der **sechs** aktive Adaptionen führt statt vorher acht Volltexte.

**Lerneintrag — Form: geschärfte Regel.** *`scan.ignore` nimmt einen Baum als **Quelle** aus, nicht
als **Ziel**: Anker dorthin werden geprüft.* Das neue Pflichtfeld verlangt einen Link **mit
Abschnitts-Anker** in die vendored Baseline — und die liegt in `scan.ignore`. Ob so ein Anker
überhaupt validiert wird, war offen; ein falscher hätte identisch ausgesehen wie ein richtiger.
Die Negativ-Probe **vor** dem ersten Eintrag hat es entschieden: ein Link mit erfundenem Anker
dorthin ergibt `anchor-missing`, Exit 2. *Weil* die Antwort das Feld von einer Behauptung in einen
gegateten Nachweis verwandelt, war die Probe die halbe Minute wert — und sie ist der Grund, warum
die sechs `Ersetzt-Baseline-Regel`-Angaben mehr sind als Dekoration.

**Zwei beobachtbare Closure-Kriterien:**

1. `doc-check` prüft 211 Dateien mit **0** Befunden, bei 136 Verweisen auf alte
   Überschriften-Slugs — der Formwechsel aus C1 hält auch nach dem Umzug von acht weiteren
   Dateien.
2. `harness/conventions/` enthält **6** Dateien, `harness/conventions/done/` **10**. Vor C1 waren
   es zehn Volltexte in einer Datei, die jeder Agentenlauf mitlas.

**Offene Risiken und ihr Ausgang:**

- *Die Platzierung des Rückbau-Eintrags in `done/` ist Auslegung, keine Vorschrift* — Ausgang:
  **weiter offen**, gehört ins Beobachtungs-Register. Das gibt es noch nicht; es entsteht in
  Etappe D. Bis dahin steht die Auslegung samt Begründung in §3 dieses Slice.
- *Der Nachfolger der ADR-Vorlagen-Version ist unter dem Fork-Test ein Rückbau-Kandidat* —
  Ausgang: **weiter offen**, ebenfalls fürs Register. Er trägt einen eigenen Auflösungs-Trigger
  (die Überarbeitung der ID-Schema-Deklaration), aber ein Trigger, den niemand abfragt, ist kein
  Wächter — genau dafür ist das Register da.
- *Die acht verschobenen Einträge brauchten Link-Anpassungen um eine Verzeichnisebene* —
  Ausgang: **gestrichen mit Begründung**; erledigt in einem eigenen Inhalts-Commit nach der
  reinen Umbenennung, wie §3.3 es verlangt.

**Folge-Slices:** Etappe C3 (Slice-Kopffelder und die Metrik von `verify-slice-form`), danach D —
in der das Beobachtungs-Register entsteht, das zwei der drei Risiken oben aufnimmt.

## 8. Sub-Area-Modus

Berührt wird der Konventionsspeicher — weiterhin ohne eigene Zeile in der Modus-Deklaration pro
Sub-Area ([slice-091 §7](../done/slice-091-claude-md-auf-verweis-reduzieren.md)). Alle berührten
Sub-Areas mit Modus sind GF.
