# slice-150 — Roadmap-Form gegen die v6.0.0-Ziel-Form nachgezogen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Vorgabe 2026-09-05 („roadmap.md noch anpassen, siehe
`.harness/baseline/v6.0.0/templates/docs/plan/planning/roadmap.template.md"), unmittelbar nach
[slice-149](wellenlos/slice-149-carveout-vorlage-lokale-kopie-entfernt.md) — dieselbe Fund-Klasse
(lokales Artefakt vs. vendored Ziel-Form), diesmal an `roadmap.md` selbst statt an einer Vorlage.
[Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Planungs-Harness-Pflege, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

`docs/plan/planning/in-progress/roadmap.md` strukturell gegen
[`.harness/baseline/v6.0.0/templates/docs/plan/planning/roadmap.template.md`](../../../../.harness/baseline/v6.0.0/templates/docs/plan/planning/roadmap.template.md)
halten — anders als bei `carveout.template.md` ([slice-149](wellenlos/slice-149-carveout-vorlage-lokale-kopie-entfernt.md))
führt a-check hier **keine** lokale Kopie der Vorlage, sondern das **Instanz-Dokument** selbst; ein
`diff` ist darum kein Werkzeug — der Vergleich ist strukturell (Abschnitte, Tabellen-Spalten),
nicht textuell.

## 2. Vier reale Befunde, gemessen per Abschnitts-für-Abschnitt-Vergleich

| # | Fund | Art |
|---|---|---|
| 1 | Kopf-Feld `**Status:** Aktiv. **Letzte Änderung:** 2026-08-30.` — kein Baseline-Feld, und **bereits stale**: die Roadmap wurde seither mehrfach editiert (Etappen A–D, `slice-149`) | echte Drift |
| 2 | Meilensteine-Tabelle trug **drei** Spalten (`Meilenstein · Welle(n) · Status`); die Ziel-Form verlangt **vier** (`Meilenstein · Welle(n) · Trigger · Status`) | Struktur-Lücke |
| 3 | Der Absatz **„Format-Regel"** stand **zweimal** — einmal im Kopf (Zeile 5–8), wortgleich noch einmal nach dem Drift-Log (Zeile 129–133 vor dieser Änderung) | Dublette |
| 4 | Der Block **„Hinweis zur Slice-Buchführung"** hing verwaist nach dem Drift-Log — ein Dokument-weiter Hinweis an der falschen Stelle, nicht Teil der fünf/sechs Baseline-Abschnitte | Fehlplatzierung |

**Nicht angetastet, geprüft und für gültig befunden:**

- Die Zusatz-Tabelle **„Offene Slices ohne Welle"** unter *Offene Wellen* — kein Baseline-Gegenstück,
  aber deckt einen echten a-check-Fall (Slices, die auf ein Fremdrepo warten, siehe direkt
  darunter stehende Begründung) und verletzt keine der beiden Sensor-Zusagen (`heading`/`marker`
  aus `.d-check.yml` Modul `planning`).
- Die **fünfte Spalte `Beleg`** im Drift-Log (Baseline hat nur drei: `Datum · Was wurde
  geändert? · Warum?`) — eine Erweiterung um belegbare Nachweise, konsistent mit dem sonst überall
  in diesem Repo verlangten Beleg-Zwang, keine Drift.
- Das lange `BEDIENHINWEIS`-Kommentar der Baseline zum Zwei-Hälften-Wächter der *Offene
  Wellen*-Sektion — inhaltlich bereits in `.d-check.yml` (Kommentar über dem `planning`-Modul)
  geführt; eine zweite Kopie hätte zwei Stellen, die driften können.

## 3. Umsetzung

- Kopf-Feld entfernt; der „Hinweis zur Slice-Buchführung" zieht an seine Stelle (Dokument-Kontext
  gehört an den Anfang, nicht ans Ende).
- Meilensteine-Tabelle auf vier Spalten erweitert. **Kein neuer Fakt behauptet** — die
  Trigger-Spalte übernimmt Text, der vorher als Parenthese im `Meilenstein`-Feld stand — die
  [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Referenz
  aus M2 wandert samt Link in die `Trigger`-Spalte, die `Meilenstein`-Zelle bleibt „M2:
  Dogfooding" — reine Umsortierung bestehenden Texts in die richtige Spalte.
- Dubletten-Absatz entfernt (Format-Regel steht jetzt nur noch einmal, im Kopf).

## 4. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor; das Modul `planning` prüft
`heading`/`marker` unverändert weiter.

## 5. DoD

- [x] Vier Befunde per Abschnitts-Vergleich gemessen, nicht angenommen (§2); drei bewusst
      unangetastete Abweichungen einzeln begründet (§2).
- [x] Alle vier behoben, kein neuer Fakt behauptet — nur bestehender Text neu einsortiert (§3).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** `roadmap.md` ist strukturell wieder deckungsgleich mit der `v6.0.0`-Ziel-Form —
vier Befunde behoben (ein stales Kopf-Feld, eine fehlende Tabellenspalte, ein doppelter Absatz, ein
verwaister Hinweis-Block), drei bewusste a-check-Erweiterungen belassen und einzeln begründet statt
stillschweigend übernommen oder gestrichen.

**Lerneintrag — Form: geschärfte Regel.** *Ein manuell gepflegtes Zustandsfeld wie „Letzte
Änderung: <Datum>" ist in einem Repo, dessen eigene Disziplin durchgängig „Zustand ist das
Verzeichnis/die Struktur, kein Feld" lautet ([`AGENTS.md`](../../../../AGENTS.md) §3.7, sinngemäß
für Zustandsfelder), ein Fremdkörper — und dieser Slice fand ihn genau dort, wo die Disziplin ihn
vorhergesagt hätte: bereits stale, weil kein Schritt der letzten vier Slices (`slice-147`…`149`)
ihn anfasste, obwohl jeder von ihnen die Datei änderte.* Dieselbe Lehre wie
[slice-149](wellenlos/slice-149-carveout-vorlage-lokale-kopie-entfernt.md), an einem anderen
Artefakt-Typ: **manuell nachzuführende Metadaten sind ein Drift-Risiko unabhängig davon, ob sie in
einer Vorlagen-Kopie oder in einem Instanz-Dokument stehen.**

**Zwei beobachtbare Closure-Kriterien:**

1. `grep -c '^\*\*Format-Regel:\*\*' docs/plan/planning/in-progress/roadmap.md` → `1` (vorher `2`);
   die Meilensteine-Tabellenkopfzeile führt vier Spalten (`grep -c '|' <<< '<Kopfzeile>'` → `5`
   Pipe-Zeichen für vier Zellen).
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Keines notiert — der Fund ist mit der Korrektur geschlossen,
kein offener Rest.

**Folge-Slices:** keine vergeben. Ob weitere Instanz-Dokumente (nicht nur Vorlagen-Kopien) gegen
ihre jeweilige Baseline-Ziel-Form driften, ist mit diesem Slice **nicht** geprüft — eine gezielte
Nachfrage traf `roadmap.md`, kein systematischer Durchgang; wird nicht spekulativ behauptet.

## 7. Sub-Area-Modus

Berührt: **Planungs-Harness** (`docs/plan/planning/in-progress/roadmap.md`) — Greenfield: Form
steht in der vendored Vorlage, `doc-check` prüft Verweis-Auflösung, Modul `planning` prüft
`heading`/`marker`.
