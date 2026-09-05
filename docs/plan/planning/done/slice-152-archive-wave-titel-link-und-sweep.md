# slice-152 — Titel-Link-Bug in Apply/ApplySlice behoben, weiterer Sweep

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Hinweis 2026-09-05 („In `docs/plan/planning/done` liegen immer noch Slices.
Siehe wie es d-check gelöst hat"), direkt nach [slice-151](../done/slice-151-offene-slices-ohne-welle-entfernt.md).
[Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Bestandspflege + Werkzeug-Fix, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

Prüfen, wie d-check mit verbleibenden flachen `done/`-Slices umgeht, und den eigenen Bestand
entsprechend nachziehen.

## 2. Präzedenz gelesen, nicht kopiert

d-check hält `docs/plan/planning/done/` **nicht** dauerhaft leer — aktueller Bestand dort: zwei
flache Slices (die beiden jüngsten). Archivierung läuft dort ebenfalls als periodischer,
von Hand ausgelöster Sweep (`tools/archive-wave`, kein Gate), nicht als automatischer Schritt jeder
Slice-Closure.

## 3. Erster Sweep-Versuch — zwei reale Funde

Alle 17 flachen Slices (`slice-135`…`151`) mit `-slice=<id> APPLY=1` archiviert. `make gates` blieb
grün, `make verify` brach an **zwei** Stellen:

1. **`tools/verify-risiko-ausgaenge.sh`** — fail-closed bei `count == 0`: „null Dateien sind
   Bestandsverlust, nicht ‚nichts zu prüfen'". Der non-rekursive Glob `done/slice-*.md` fand nach
   dem Sweep nichts mehr.
2. **`.d-check.yml`-Modul `structure`** — vier Regeln melden „Regel trifft keine Datei … das Gate
   liefe leer", aus demselben Grund.

Beide Sensoren setzen voraus, dass `done/` **mindestens eine** flache Slice-Datei führt — eine
Annahme, die der neue, gründlichere Sweep erstmals verletzte. Konsistent mit d-checks eigenem
Bestand (nie exakt null): Sweep zurückgesetzt (`git restore`/`git clean` auf den Stand vor dieser
Änderung, `tools/archive-wave/`-Fixes ausgenommen), auf **15** Slices (`135`…`149`) begrenzt,
`150`/`151` bleiben flach als Puffer — dieselbe Größenordnung wie bei d-check.

## 4. Zweiter, unabhängiger Fund: derselbe Titel-Link-Bug an zwei weiteren Stellen

Der erneute Lauf brach an `slice-141` — Titel `[MR-017](...): [MR-013](...)s Nachfolge-Eintrag`.
Derselbe Bug wie in [slice-148](../done/wellenlos/slice-148-archivierung-etappe-d.md) (`ApplyReview`
schrieb einen im Titel eingebetteten Link nicht depth-korrekt um) stand unverändert in **zwei**
weiteren, strukturell identischen Funktionen: `Apply` (Wellen-Modus) und `ApplySlice`
(wellenlos-Modus). Der frühere Fix zog nicht automatisch nach, weil die drei Funktionen denselben
Code dreifach statt einmal tragen.

Behoben: `title` läuft jetzt an beiden Stellen durch `RewriteFieldForMove`, wie zuvor schon das
Welle-Feld. Je ein Regressionstest (`TestFixture_SliceTitelMitLink`,
`TestRunSlice_Apply_TitelMitLink`) ergänzt — Fixture defekt gefunden (fehlendes `docs/reviews/`
im zweiten Test-Root, behoben), danach `make archive-wave-test` grün.

Arbeitsbaum erneut zurückgesetzt (Tool-Fixes ausgenommen), die 15 Slices mit dem reparierten
Werkzeug neu archiviert.

## 5. Steering-Loop: dritter Beleg, Register-Schwelle erreicht

[`BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig`](../observations/BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig/observation.md)
erreicht mit diesem Slice **3×** (`slice-147`, `slice-148`, `slice-152`). Nach Register-Regel
([`modul-06-roadmap.md`](../../../../.harness/baseline/v6.0.0/regelwerk/modul-06-roadmap.md)
§Das Beobachtungs-Register) wird der Eintrag zur **verkörperten Regel**: ein Invarianten-Kommentar
über der Paketdeklaration in `tools/archive-wave/archive.go` hält jetzt ausdrücklich fest, dass
jede Stub-Erzeugung Titel **und** Feld durch `RewriteFieldForMove` schicken muss — mit Verweis auf
die drei Regressionstests, die das je Modus einzeln prüfen. `state.md` trägt den Ausgang mit
Herkunfts-Anker `seit slice-152`.

## 6. Ergebnis

- 15 weitere Slices archiviert (`done/wellenlos/`), zwei (`slice-150`, `slice-151`) bleiben flach
  als Puffer — `docs/plan/planning/done/` hat damit nie null flache Slice-Dateien geführt.
- Titel-Link-Bug an allen drei Stub-Erzeugungsfunktionen (`Apply`, `ApplySlice`, `ApplyReview`)
  behoben, mit Regressionstests je Modus.
- Neuer Invarianten-Kommentar in `tools/archive-wave/archive.go`, der die drei Stellen künftig
  als eine geprüfte Zusage statt dreier unabhängiger Kopien führt.

## 7. Auszuführende Gates

`make gates`, `make archive-wave-test`, zum Abschluss `make verify`. Kein neuer Sensor —
`verify-risiko-ausgaenge` und die `structure`-Regeln bleiben unverändert; ihre Zero-Match-Annahme
wird durch den belassenen Puffer eingehalten, nicht durch eine Sensor-Änderung umgangen.

## 8. DoD

- [x] d-checks Praxis geprüft (nie exakt null flache Slices), nicht angenommen (§2).
- [x] Beide realen Funde (Zero-Match-Sensoren, Titel-Link-Bug an zwei weiteren Stellen) behoben
      bzw. durch einen begründeten Puffer umgangen, mit Regressionstest bzw. Gegenprobe (§3, §4).
- [x] `make gates`, `make archive-wave-test` und `make verify` grün — Ausgabe in eine Datei,
      Exit-Code getrennt geprüft, nie in eine Pipe.

## 9. Closure-Notiz

**Geliefert:** 15 weitere Alt-Slices archiviert, ein Puffer von zwei flachen Slices bewusst
belassen (Präzedenz d-check), der Titel-Link-Bug an allen drei betroffenen Stellen behoben statt
nur an der zuletzt gefundenen.

**Lerneintrag — Form: neuer Sensor.** *Ein Fix, der eine von drei strukturell identischen
Code-Kopien trifft, ist kein vollständiger Fix — er behebt das Symptom an der Stelle, die gerade
gebrochen ist, nicht die Klasse.* Der neue Invarianten-Kommentar plus drei Regressionstests
(einer je Kopie) machen die Zusage jetzt an allen drei Stellen explizit statt implizit — kein
Refactor auf eine gemeinsame Funktion (das Risiko, drei bereits getestete Pfade in einem Schritt
umzubauen, stand außer Verhältnis zum Nutzen einer strukturellen Garantie, die drei benannte Tests
ebenso zuverlässig tragen). *Weil* Modul 6 bei 3× ausdrücklich eine verkörperte Regel statt einer
weiteren Registerzeile verlangt, trägt dieser Slice sie: Anker `seit slice-152` in
[`BEO-GATE`](../observations/BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig/observation.md).

**Zwei beobachtbare Closure-Kriterien:**

1. `ls docs/plan/planning/done/*.md | grep -v -- '-results.md$' | wc -l` → `2` (`slice-150`,
   `slice-151`); `ls docs/plan/planning/done/wellenlos/*.md | wc -l` → `92` (`77` aus
   [slice-148](../done/wellenlos/slice-148-archivierung-etappe-d.md) `+ 15` aus diesem Slice).
2. `make gates`, `make archive-wave-test` und `make verify` je **Exit 0** auf dem Stand dieses
   Slice, Ausgabe in Dateien, Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Eines notiert:
[`BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig`](../observations/BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig/observation.md)
— **verkörpert** (dritter Beleg, Schwelle erreicht, §5). Der Zero-Match-Sensor-Fund selbst
(`verify-risiko-ausgaenge`/`structure` setzen `done/` ≠ ∅ voraus) trägt **keinen** eigenen
Ausgang — er ist durch den belassenen Puffer umgangen, nicht behoben; ob `done/` je wieder auf
null Slices zulaufen soll (und die Sensoren dafür angepasst werden müssten), ist eine offene
Repo-Entscheidung, hier nicht getroffen.

**Folge-Slices:** keine vergeben.

## 10. Sub-Area-Modus

Berührt zwei Sub-Areas:

- **Gate-/Werkzeug-Schicht** (`tools/archive-wave/`) — Greenfield: beide Funde getestet
  (`archive-wave-test`), Invariante als Kommentar verkörpert.
- **Planungs-Harness** (`docs/plan/planning/done/`) — Greenfield: Form steht in der vendored
  Vorlage, `doc-structure` prüft sie.
