# slice-122 — `doc-planning` bekommt seinen Gegenstand

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** [`BEO-014`](../observations.md) — bei **2×**.

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

`make doc-planning` prüft wieder etwas: das Modul `planning` bekommt den Konfigurationsblock, ohne
den es grün meldet, statt zu schweigen.

## 2. Definition of Done

- [ ] [`.d-check.yml`](../../../../.d-check.yml) trägt einen `planning`-Block, der auf a-checks
      **tatsächliche** Roadmap-Form zeigt — die Defaults (`## Aktuelle Welle`,
      `Keine aktive Welle`) treffen sie nicht.
- [ ] Die Wirksamkeit ist in **beide** Richtungen belegt: der unveränderte Bestand ist grün, und
      ein Slice in `in-progress/` bei stehendem Ruhe-Marker liefert `planning-drift`.
- [ ] [`AGENTS.md`](../../../../AGENTS.md) §4 und
      [`harness/README.md`](../../../../harness/README.md) §Sensors nennen den Vertrag, den das
      Target **wirklich** hat; `doc-planning` hängt im `gates`-Aggregat.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| [`.d-check.yml`](../../../../.d-check.yml) | update | der fehlende `planning`-Block |
| [`Makefile`](../../../../Makefile) | update | `doc-planning` ins `gates`-Aggregat |
| [`AGENTS.md`](../../../../AGENTS.md) §4, [`harness/README.md`](../../../../harness/README.md) | update | Vertrag und Bindung |

**Auszuführende Gates:** `make gates` — tragend `doc-planning` (neu wirksam), `doc-targets`
(beide Doku-Tabellen) und `guard-selftest`. Zum Abschluss `make verify`.

### Das Risiko ist vorab gemessen

| Prüfung | Ergebnis |
|---|---|
| `planning:` in [`.d-check.yml`](../../../../.d-check.yml) | **fehlt** — das Modul läuft ohne Gegenstand |
| Defaults gegen a-checks Roadmap | treffen nicht: die Sektion heißt `## Offene Wellen`, der Ruhe-Marker `Nichts in Arbeit.` |
| konfiguriert, unveränderter Bestand | Exit 0, 0 Befunde |
| konfiguriert, Slice in `in-progress/` bei stehendem Ruhe-Marker | **`planning-drift`**, Exit 1 — die Meldung nennt Sektion und Marker |

Die letzte Zeile ist der eigentliche Beleg: **genau der Fall aus `BEO-014`** wird gefangen.

## 4. Trigger

**Start:** eingetreten — `BEO-014` bei 2×, die Lücke ist gemessen.

**Rückführungen:**

- `in-progress` → `open`: falls die Konfiguration über den Bestand Befunde liefert, die eine
  Form-Entscheidung an der Roadmap verlangen statt einer Korrektur.

## 5. Closure-Trigger

Block konfiguriert, Wirksamkeit beidseitig belegt, Deklarationen nachgezogen, Gates grün.

**Was bewusst nicht getan wird:** die **zweite und dritte Fähigkeit** des Moduls (`closure:` und
`waves:`). Die Closure-Struktur prüft seit
[slice-080](../done/slice-080-verify-abloesung-dcheck.md) das Modul `structure` — sie ein zweites
Mal zu verdrahten hieße, zwei Prüfer auf dieselbe Zusage zu setzen, die dann auseinanderlaufen
können. Ein Wellen-Register führt a-check nicht.

## 6. Risiken und offene Punkte

- *Ein zweiter Prüfer auf die Roadmap-Form könnte mit `verify-*` kollidieren* — **Ausgang:** <bei Closure>
- *`doc-planning` im Aggregat macht jeden künftigen Lifecycle-Wechsel gate-pflichtig: wer den
  Slice bewegt und die Roadmap-Zeile vergisst, bekommt ein rotes Gate* — das ist die Absicht,
  aber es ändert den Arbeitsablauf. **Ausgang:** <bei Closure>

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5.)_

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>**

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Gate-/Werkzeug-Schicht**
(`.d-check.yml`, `Makefile`) und mit den zwei Deklarations-Tabellen der **Harness-Einstieg**.

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-014`](../observations.md) ist der Anlass
(2×). [`BEO-023`](../observations.md) (Prüfer mit leerer Prüfmenge) liegt in derselben Schicht und
beschreibt dieselbe Familie — dieser Slice gibt einem Prüfer seinen Gegenstand zurück.

Alle berührten Sub-Areas GF.
