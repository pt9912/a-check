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

- [x] [`.d-check.yml`](../../../../.d-check.yml) trägt einen `planning`-Block, der auf a-checks
      **tatsächliche** Roadmap-Form zeigt — die Defaults (`## Aktuelle Welle`,
      `Keine aktive Welle`) treffen sie nicht.
- [x] Die Wirksamkeit ist in **beide** Richtungen belegt: der unveränderte Bestand ist grün, und
      ein Slice in `in-progress/` bei stehendem Ruhe-Marker liefert `planning-drift`.
- [x] [`AGENTS.md`](../../../../AGENTS.md) §4 und
      [`harness/README.md`](../../../../harness/README.md) §Sensors nennen den Vertrag, den das
      Target **wirklich** hat; `doc-planning` hängt im `gates`-Aggregat.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

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

- *Ein zweiter Prüfer auf die Roadmap-Form könnte mit `verify-*` kollidieren* — **Ausgang:**
  gestrichen mit Begründung: es gibt keinen. Die Roadmap-Form prüft heute **niemand** sonst; die
  `verify-*`-Reste (`verify-risiko-ausgaenge`, `verify-observations`) sehen Slice-Dateien und das
  Register, nicht die Roadmap. Die zweite Modul-Fähigkeit (`closure:`) bleibt deshalb
  unkonfiguriert — **dort** gäbe es die Kollision, mit `doc-structure`.
- *`doc-planning` im Aggregat macht jeden Lifecycle-Wechsel gate-pflichtig* — **Ausgang:**
  eingetreten und beabsichtigt. Wer einen Slice bewegt und die Roadmap-Zeile vergisst, bekommt ab
  jetzt ein rotes Gate statt eines stillen Widerspruchs. Das trifft genau die Lücke, die
  `make slice-mv` offenlässt: das Werkzeug zieht **Pfade** nach, keine Semantik
  ([slice-118](../done/slice-118-lifecycle-wechsel-werkzeug.md) §7). Werkzeug und Gate greifen
  jetzt ineinander.

## 7. Closure-Notiz

**Geliefert:** Das Modul `planning` ist konfiguriert und `make doc-planning` hängt im
`gates`-Aggregat. Es prüft jetzt, was sein Name verspricht: liegt ein Slice in `in-progress/`,
benennt ihn die Roadmap-Sektion, statt den Ruhe-Marker zu tragen.

**Lerneintrag — Form: geschärfte Regel.** *Fail-closed schützt erst ab der ersten
Konfigurationszeile — ein Modul **ohne** Block ist nicht „fail-closed inaktiv", sondern still.*
Das ist gemessen, und die Messung hat meine Erwartung korrigiert: Ich hatte angenommen, die
Defaults (`## Aktuelle Welle`, `Keine aktive Welle`) würden an a-checks Roadmap **falsch**
prüfen. Sie prüfen nicht falsch — sie melden **fail-closed**: *„kanonische Überschrift fehlt …
Aktiv-Status nicht bestimmbar"*. Das Modul ist also sorgfältig gebaut; es kann nur nichts
melden, solange niemand ihm sagt, wo es hinsehen soll. *Weil* die Fail-closed-Zusage erst
**innerhalb** der Konfiguration greift, ist „das Modul ist fail-closed" keine Auskunft über ein
Repo, das keinen Block führt — und genau diese Verwechslung hielt `BEO-014` zwei Runden offen.

**Drei beobachtbare Closure-Kriterien:**

1. Wirksamkeit in beide Richtungen, mit der **echten** Konfiguration gefahren: unveränderter
   Bestand ⇒ Exit 0; Ruhe-Marker bei belegtem `in-progress/` ⇒ `planning-drift`, und die Meldung
   nennt Sektion **und** Marker.
2. Die **Grenze** steht im Konfigurations-Kommentar, weil sie ohne Probe nicht sichtbar ist:
   geprüft wird die Äquivalenz *„Slice(s) vorhanden ⟺ Ruhe-Marker steht nicht"*, **nicht** ob
   jeder einzelne Slice genannt ist. Zwei Slices bei einer Sektion, die nur einen nennt, bleiben
   grün — gemessen —, und das WIP-Limit fängt das Modul ebenfalls nicht.
3. `closure:` und `waves:` bleiben unkonfiguriert, und der Kommentar sagt warum: die
   Closure-Struktur prüft seit [slice-080](../done/slice-080-verify-abloesung-dcheck.md) das
   Modul `structure`. Zwei Prüfer auf dieselbe Zusage laufen auseinander — dieselbe Klasse wie
   `BEO-024` aus [slice-121](../done/slice-121-port-richtung-inbound-outbound.md), hier
   vermieden statt eingegangen.

**Offene Risiken und ihr Ausgang:** der erste gestrichen mit Begründung, der zweite eingetreten
und beabsichtigt.

**Beobachtungs-Register:** [`BEO-014`](../observations.md) ist **verkörpert** — der Zähler bleibt
bei 2×, sein Stand nennt jetzt den Ort. Damit ist die Beobachtung für `planning` aufgelöst; die
allgemeinere Familie führt [`BEO-023`](../observations.md) weiter (ein Prüfer ohne Gegenstand
bleibt unkalibriert).

**Ein zweiter Register-Eintrag fiel dabei an, ungesucht.** [`BEO-006`](../observations.md)
geht auf **2×**: der Risiko-Ausgang von
[slice-121](../done/slice-121-port-richtung-inbound-outbound.md) traf die geschlossene
Dreier-Menge nicht (*„eingetreten und beabsichtigt"* ist keiner der drei) — und fiel erst
auf, als **dieser** Slice `make verify` fuhr. slice-121 lag da schon in `done/`. Genau die
Reihenfolge-Falle, die `BEO-006` seit slice-099 beschreibt, hier zum ersten Mal real
eingetreten. Korrigiert: es war nie ein Risiko, sondern die Absicht — also **gestrichen
mit Begründung**.

**Folge-Slices:** keiner zwingend. Offen bleibt [`BEO-024`](../observations.md) — die Zuordnung
in zwei Paketen aus [slice-121](../done/slice-121-port-richtung-inbound-outbound.md); sie ist
strukturell lösbar, indem eine Seite die andere liest.
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Gate-/Werkzeug-Schicht**
(`.d-check.yml`, `Makefile`) und mit den zwei Deklarations-Tabellen der **Harness-Einstieg**.

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-014`](../observations.md) ist der Anlass
(2×). [`BEO-023`](../observations.md) (Prüfer mit leerer Prüfmenge) liegt in derselben Schicht und
beschreibt dieselbe Familie — dieser Slice gibt einem Prüfer seinen Gegenstand zurück.

Alle berührten Sub-Areas GF.
