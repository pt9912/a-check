# slice-204 — Review-Haken an Report-Existenz binden (in-progress)

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt. Er wechselt nur durch `git mv`.

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** 5. Auflage von
[`BEO-GATE/attestierung-vor-dem-vorgang`](../observations/BEO-GATE/attestierung-vor-dem-vorgang/observation.md)
(slice-169, slice-197, slice-200, slice-191, slice-205 — Beleg im
evidence-Verzeichnis).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** neuer Sensor.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der DoD-Punkt „Unabhängiger Review durchgeführt" wird in
`in-progress/` an die Existenz des Reports gebunden — der fünfte Vorfall
(slice-205: der Closure-Commit attestierte den Review, der Report entstand
erst danach) zeigte, dass Bedingung 7 (Häkchen in `open/`) diese Gestalt
nicht fängt: sie entsteht im `in-progress/`-Stand.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Änderungen am `reviews`-Modul oder an
  [`MR-019`](../../../../harness/conventions/MR-019-review-dod-opt-in.md).**
  *Es wäre ein anderer Vorgang*: [`MR-019`](../../../../harness/conventions.md#mr-019) (Opt-in) und der done/-Geltungsbereich
  von `doc-reviews` sind deklarierte Entscheidungen; die Bindung hier läuft über
  das `structure`-Modul, nicht über `doc-reviews`.

## 2. Definition of Done

- [x] Der Haken ist an die Report-Existenz gebunden: ein abgehakter
      DoD-Punkt „Unabhängiger Review" in einem `in-progress/`-Slice ohne
      Report unter `docs/reviews/` meldet rot
      (`tools/verify-review-haken.sh`, im `verify`-Aggregat).
- [x] **Gegenprobe in beiden Richtungen** im Selbsttest: mit Report bleibt
      die Stelle grün, ohne Report rot; die `doc-reviews`-Bedeutung
      (done/-Geltungsbereich) bleibt unberührt.
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag; Register fortgeschritten;
      jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

Gebaut über die §6-Schalung — das `structure`-Modul hat keine
Datei-Existenz-Bedingung, und das `reviews`-Modul scannt genau ein `done-dir`
(d-check v0.79.0); ein CR an das Fremdwerkzeug ist Out-of-Scope (§1):

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `tools/verify-review-haken.sh` | neu | Scannt `in-progress/` nach abgehakten „Unabhängiger Review"-Punkten und verlangt die Kennung im Report-Dateinamen; Selbsttest in beiden Richtungen samt Kennungs-Grenze (slice-205 vs. slice-2050) und done/-Abgrenzung |
| `Makefile` | neu | Eigenes Target + Anschluss an das `verify`-Aggregat |
| [`harness/README.md`](../../../../harness/README.md) §Sensors | neu | Gate-Index-Zeile — `make doc-targets` hält beide Richtungen |

## 4. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen:** `in-progress` → `next` (zu groß): entfällt — ein
Sensor-Skript samt Verdrahtung. `in-progress` → `open` (blockiert): bietet das Modul
keine Existenz-Prüfung für Dateien außerhalb des Repos... (entfällt, das
Repo ist die Wurzel).

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 6. Risiken und offene Punkte

- **Die Existenz-Prüfung braucht mehr als ein Pattern** — `forbid-pattern`
  prüft Text, nicht Datei-Existenz; der Weg führt über das `reviews`-Modul
  mit `done-dir`-Äquivalent für in-progress oder eine eigene Schalung.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: die eigene Schalung
  ist in diesem Slice gebaut (`tools/verify-review-haken.sh`, im
  `verify`-Aggregat) — der befürchtete Umweg über ein Folge-Slice trat nicht
  ein.

## 7. Closure-Notiz

**Lerneintrag — Form: neuer Sensor.** [`tools/verify-review-haken.sh`](../../../../tools/verify-review-haken.sh)
(im `verify`-Aggregat; `seit slice-204`) hält den abgehakten Review-DoD an
die Report-Existenz: die attestierte Gestalt wird jetzt im Moment ihres
Entstehens gefangen, nicht erst bei der des nächsten Slice.

**Was hat funktioniert:** die §6-Schalung — der Plan sonderte voraus, dass
`structure` keine Datei-Existenz prüft, und der Umweg über die
Verifikations-Schicht sparte einen CR an das Fremdwerkzeug; `doc-targets`
und `gate-consistency` hielten die Verdrahtung in beiden Richtungen scharf
(Target, `.PHONY`, GATES-Liste, Index-Zeile).

**Was ging anders als geplant:** der unabhängige Review fand einen
merge-blockierenden Deckungsfehler (F-1 — die Grenze behauptete
`next/`-Coverage durch Bedingung 7, die nur `open/` bindet), einen
False-Green-Pfad in der Prüfpipeline (F-2 — `grep -q` unter `pipefail` dreht
SIGPIPE-Treffer um) und eine Probe, die ihren Gegenstand nicht trug (F-3 —
die done/-Fixture lag im Baum und wurde nie gelesen). Alle korrigiert,
bevor die Closure geschrieben wurde.

**Steering-Loop-Eintrag:** siehe Lerneintrag oben. Die
attestierungs-Klasse ([`BEO-GATE/attestierung-vor-dem-vorgang`](../observations/BEO-GATE/attestierung-vor-dem-vorgang/observation.md))
ist damit **verkörpert** (Lese-Schritt dieser Closure). Die
Probe-Klasse ([`BEO-GATE/probe-liefert-den-gegenstand-mit`](../observations/BEO-GATE/probe-liefert-den-gegenstand-mit/observation.md))
erreichte **4×** (F-3) — kein mechanischer Sensor nachgeschaltet: ob eine
Probe ihren Gegenstand trifft, ist Urteil über ihren Aufbau (§3.7); die
Begründung steht im Eintrag.

**Beobachtungs-Register (`../observations/`):**
`attestierung-vor-dem-vorgang` → Ausgang **verkörpert** (`seit slice-204`);
`probe-liefert-den-gegenstand-mit` → 4. Auflage belegt
([evidence/slice-204.md](../observations/BEO-GATE/probe-liefert-den-gegenstand-mit/evidence/slice-204.md)).

**Folge-Slices:** keine.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*entfallen*, mit
Begründung — die eigene Schalung ist gebaut).

**Drei Paarungen:** Anker — verkörpert (Sensor, `seit slice-204`) ·
Folge-Slice — keine genannt · Register — attestierung (verkörpert) ·
probe-liefert (4×, Beleg ergänzt).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** GATE (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde durchgegangen
(2026-09-29):
[`attestierung-vor-dem-vorgang`](../observations/BEO-GATE/attestierung-vor-dem-vorgang/observation.md)
— **dieser Slice ist ihr Ausgang** (4×, die Erweiterung um in-progress).

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
