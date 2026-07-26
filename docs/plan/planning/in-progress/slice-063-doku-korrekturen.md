# slice-063 — Vier Falschaussagen in slice-048 und ein Feld, das lügt

**Status:** der Zustand ist das **Verzeichnis** dieser Datei
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5) — dieses Feld führt ihn bewusst **nicht** doppelt.
**Deckt:** **F-3**, **F-4**, **F-5** aus
[`2026-07-26-etappe-b-slice-048.md`](../../../reviews/2026-07-26-etappe-b-slice-048.md), **F-3**
aus [`2026-07-26-etappe-e-slice-050-051.md`](../../../reviews/2026-07-26-etappe-e-slice-050-051.md)
und **F-1** aus [`2026-07-26-etappe-c-slice-055-056.md`](../../../reviews/2026-07-26-etappe-c-slice-055-056.md).
**Bezug:** letzter Fix-Schnitt der Review-Serie; reine Doku, kein Sensor.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

Fünf Befunde ohne Sensor-Anteil. Vier betreffen
[slice-048](../done/slice-048-modul-delta-lesen.md), das als Analyse-Dokument weiterhin gelesen
wird und darum keine falschen Aussagen tragen darf:

| Fund | Stelle | Sache |
|---|---|---|
| **R-050-F3** | §3.1 B-3, §4 Negativbefund | behauptet, die maschinelle Referenz-Richtungs-Prüfung fehle. Falsch: das `matrix`-Modul leistet sie (`.d-check.yml`, gedeckt durch [`MR-005`](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung)). Der Autor hat das selbst erkannt — die Korrektur steht bis heute **nur in einer Commit-Message**. |
| **R-048-F3** | Kopf, §1, §7 | „20 Dateien, 2867 Zeilen": die 20 aufgezählten Dateien tragen **2778** Zeilen; 2867 erreicht man nur mit der nicht aufgezählten Index-`README.md` (89). |
| **R-048-F4** | §3.4 B-19 | als **INFO** geführt („Hinweis, keine Aktion erwartet"), formuliert aber einen Handlungsauftrag und hat real zu [`MR-008`](../../../../harness/conventions.md#mr-008--kein-replay-keine-agenten-telemetrie) geführt. |
| **R-055-F1** | §5 Etappen-Zuordnung | **B-17** ist der Etappe C zugewiesen, dort aber weder behandelt noch als entfallen begründet; er war bereits bei seiner Entstehung gegenstandslos, weil [MR-003](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang) seit 2026-06-21 eine `Aufgelöst`-Markierung trägt. |

Der fünfte betrifft den Bestand: **R-048-F5** — ein Slice in `done/` trägt im Kopf
„**Status:** in-progress". Nach [`AGENTS.md`](../../../../AGENTS.md) §5 *ist* der Zustand das
Verzeichnis; das Feld ist damit eine zweite Quelle, die dem Verzeichnis widerspricht.
**Gemessen: 14 Dateien** — der Review nannte fünf, weil er nur gegen `main` zählte; die Kette hat
neun weitere hinzugefügt.

## 2. Betroffene Module

- [`docs/plan/planning/done/slice-048-modul-delta-lesen.md`](../done/slice-048-modul-delta-lesen.md)
  — vier Korrekturen.
- 14 Dateien in `docs/plan/planning/done/` — Status-Feld.

**Eine Schicht** (Planungs-Doku). Kein Code, kein Vertrag, kein Sensor.

## 3. Auszuführende Gates

`make gates` und `make verify`, Ausgabe je in eine Datei, Exit-Code getrennt geprüft. Kein neuer
Sensor, also keine Negativ-Probe — die Belege sind die korrigierten Aussagen selbst, jede gegen
ein Repo-Artefakt nachgerechnet (`.d-check.yml` §matrix, Zeilenzählung der vendored Baseline,
[MR-003](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang)-Auflösungsblock).

## 4. Was bewusst nicht getan wird

- **Keine Umschreibung von Befund-Kategorien.** B-19 bleibt als **INFO** stehen; ergänzt wird ein
  Vermerk, dass Kategorie und abgeleitete Handlung nicht zusammenpassten. Ein Analyse-Dokument ist
  Beleg seines Zeitpunkts — eine nachträglich „richtig" gesetzte Kategorie wäre Geschichtsklitterung.
  Korrigiert werden nur **Tatsachenaussagen**, die heute falsch sind.
- **Jeder Eingriff wird gekennzeichnet.** Präzedenz: [slice-050](../done/slice-050-verify-schicht.md)
  („eine still nachgetragene Historie wäre schlimmer als die Lücke").
- **Kein Sensor für das Status-Feld.** Ausdrücklich so gewollt: die Feld-Form ist eine
  Vorlagen-Frage, und die Vorlage führt seit slice-052 keinen Zustandswert mehr. Ein Gate darauf
  würde einen Bestand prüfen, der nicht mehr wächst.

## 5. DoD

- [x] Die vier Stellen in slice-048 sind korrigiert bzw. vermerkt, jede mit Kennzeichnung des
      Nachtrags und mit dem Artefakt, gegen das nachgerechnet wurde (R-050-F3, R-048-F3, R-048-F4,
      R-055-F1).
- [x] Kein Dokument in `done/` behauptet im Status-Feld einen Zustand; die 14 betroffenen Dateien
      tragen die zustandsfreie Form der Vorlage (R-048-F5).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** vier gekennzeichnete Nachträge in
[slice-048](../done/slice-048-modul-delta-lesen.md) — zwei Tatsachen-Korrekturen
(`check-references`, Zahlenpaar), ein Kategorie-Vermerk zu B-19 und der Verbleib von B-17 — sowie
die zustandsfreie Status-Zeile in **14** Dateien unter `done/`.

**Lerneintrag — Form: geschärfte Regel.**
> **Bei einem Nachtrag an abgeschlossener Doku entscheidet die Art der Aussage, nicht ihr Alter:
> Tatsachen werden korrigiert, Einschätzungen bleiben stehen.** Beide Fälle lagen hier
> nebeneinander. Die `check-references`-Behauptung und das Zahlenpaar sind **falsch** — sie
> beschreiben das Repo unzutreffend, und wer sie liest, zieht falsche Schlüsse; sie werden
> berichtigt. Die INFO-Einstufung von B-19 ist dagegen ein **Urteil ihres Zeitpunkts**: sie war zu
> mild, aber sie zu überschreiben hieße, die Analyse besser aussehen zu lassen, als sie war, und
> nähme dem Dokument seinen Beleg-Charakter. Prüfsatz: *bevor ein Nachtrag etwas ändert, fragen —
> ist es falsch oder war es nur unklug? Falsches wird korrigiert, Unkluges bekommt einen Vermerk.*

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. `grep -l '^\*\*Status:\*\* in-progress' docs/plan/planning/done/*.md` liefert **keinen**
   Treffer mehr (vorher 14); jeder der vier slice-048-Nachträge nennt das Artefakt, gegen das
   nachgerechnet wurde (`.d-check.yml` §matrix, Zeilenzählung der vendored Baseline, `MR-008`,
   der `Aufgelöst`-Block in `MR-003`).

**Eigene Fehlmessung, korrigiert:** der Review-Fund nannte **fünf** betroffene Status-Felder, weil
er nur gegen `main` zählte — real sind es **14**, die Kette hat neun weitere hinzugefügt. Dieselbe
Klasse wie der Stichproben-Fehler aus [slice-062](../done/slice-062-commit-scope.md): eine Messung
gilt für die Menge, aus der sie gezogen ist. Zweimal in zwei Slices; beim dritten Mal gehört das in
den Steering-Loop.

**Der Sensor, der hier fehlt und keiner sein muss:** ein Gate auf die Status-Zeile würde einen
Bestand prüfen, der nicht mehr wächst — die Vorlage führt seit slice-052 keinen Zustandswert, und
alle sechs Slices dieser Serie sind ohne angelegt worden. Wo die Ursache beseitigt ist, ist ein
Wächter über die Folgen Zierrat.

**Folge-Slices:** keine. Damit sind alle 22 Findings der Review-Serie abgearbeitet oder mit
Begründung geschlossen.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
