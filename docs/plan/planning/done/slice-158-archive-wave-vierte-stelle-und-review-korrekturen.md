# slice-158 — archive-wave: vierte Stub-Stelle behoben (Review-Korrektur)

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** [Review-Report 2026-09-05](../../../reviews/2026-09-05-slice135-157-multi-linsen-review.md),
Finding F aus der Code-Korrektheit-Linse (`tools/archive-wave/`) — ausgelöst durch die
Maintainer-Nachfrage „Warum ist der reviews-Ordner leer?" und die Entscheidung, jetzt echte,
kontext-getrennte Reviews nachzuholen. [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Werkzeug-Korrektur, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

Die Code-Korrektheit-Linse des Multi-Linsen-Reviews fand: die Behauptung in
[slice-152](../done/wellenlos/slice-152-archive-wave-titel-link-und-sweep.md)/[slice-153](../done/wellenlos/slice-153-archive-wave-welle-feld-und-titel-form.md)
— „Titel-Link-Bug an allen drei Stub-Erzeugungsfunktionen behoben" — war selbst falsch. Es gibt
eine **vierte** Stub-Erzeugungsstelle (`Apply()`s Welle-Plan-Titel, via `WelleStub`), die
denselben Bug unverändert trug. Genau das Muster, das dieser gesamte Review-Nachholvorgang
adressiert: eine als vollständig behauptete Prüfung, unvollständig, diesmal von einem
unabhängigen Kontext gefangen statt vom Maintainer.

## 2. Befund verifiziert, nicht übernommen

Vor der Korrektur selbst geprüft: `Apply()` (`archive.go`, Welle-Plan-Zweig) liest den Welle-Plan-
Titel per `readTitle()` und reicht ihn **ohne** `RewriteFieldForMove` direkt an `WelleStub()`
weiter — bestätigt per Quellcode-Lesen. Im aktuellen Bestand nicht sichtbar geworden (der einzige
committete Welle-Plan-Stub, `welle-13-konsumenten-befunde.md`, hat einen linkfreien Titel), aber
ein reales, latentes Risiko für jede künftige Welle mit einem Link im Plan-Titel.

## 3. Behoben

- `archive.go`: Welle-Plan-Titel läuft jetzt ebenfalls durch `RewriteFieldForMove`.
- Invarianten-Kommentar über der Paketdeklaration korrigiert: **vier** Stub-Erzeugungsstellen
  statt drei, die falsche Zählung selbst benannt statt stillschweigend überschrieben.
- Regressionstest `TestFixture_WellePlanTitelMitLink` ergänzt.

**Zwei weitere, unabhängig vom Reviewer gefundene MEDIUM-Funde derselben Linse, ebenfalls
behoben** (beide latent, im heutigen Bestand nicht manifestiert, aber real):

- `collect.go`: `welleFieldRE` (CollectSlices, numerische Wellen-Zugehörigkeit) und
  `welleFieldStartRE` (ReadWelleField, voller Rohwert) lasen bei einem Feldwert in der Zeile
  **nach** `**Welle:**` unterschiedlich — `\s*` schließt in Go `\n` ein, `[ \t]*` nicht.
  `welleFieldRE` auf `[ \t]*` vereinheitlicht (keine Verhaltensänderung am realen Bestand, der den
  Wert immer einzeilig schreibt); Regressionstest `TestCollectSlices_ReadWelleField_Konsistenz`.
- `archive.go`: `Apply()` hatte anders als `ApplySlice()` keinen `ohne Welle`-Fallback bei leer
  gelesenem Feld — defensiv ergänzt (im heutigen Bestand unerreichbar, da `CollectSlices` nur
  Dateien mit nicht-leerem Feld liefert, aber `Plan` ist ein exportierter Typ ohne Guard).

**Ein drittes, geprüftes und NICHT übernommenes Finding:** die Linse stufte den Kommentar über
`SliceStubStandalone` als HIGH ein („behauptet fälschlich 'keine neue Norm'"). Geprüft: der
Kommentar war tatsächlich ungenau — die Funktion führt EIN Feld statt der von der Ziel-Form
verlangten ZWEI (`Archiviert mit:`/`Geschlossen:` verschmolzen zu `Archiviert:`). Der
**Feldwert-Unterschied** war korrekt beschrieben, die **Feld-Struktur-Abweichung** nicht. Kommentar
präzisiert; das Laufzeitverhalten selbst bleibt unverändert (entspricht `d-check`s eigener,
identischer Präzedenz für wellenlose Stubs).

## 4. Was bewusst nicht behoben wird

- **Wellen-Modus meldet keine toten Review-Verweise** (Reviewer-Finding, MEDIUM) — echte
  Inkonsistenz zum Slice-Modus, aber eine Funktionserweiterung, keine Korrektur eines
  Fehlverhaltens; eigener Folge-Slice, kein Bestandteil dieser Review-Korrektur.
- Die restlichen LOW/INFO-Funde der Linse (En-Dash-Variante, Mehrfach-`# `-Zeile, leere Slice-ID,
  unbeachteter Zip-Close-Fehler, `<manuell auszufuellen>`-Platzhalter) — im heutigen Bestand ohne
  Auswirkung, jeweils einzeln benannt in der Reviewer-Ausgabe, kein eigener Korrektur-Bedarf jetzt.

## 5. Auszuführende Gates

`make gates`, `make archive-wave-test`, zum Abschluss `make verify`. Kein neuer Sensor.

## 6. DoD

- [x] Reviewer-Befund selbst verifiziert (Quellcode gelesen, nicht nur übernommen) vor der
      Korrektur (§2).
- [x] Vierte Stub-Stelle + zwei weitere MEDIUM-Funde behoben, je mit Regressionstest; ein
      HIGH-Finding geprüft und nur teilweise (Kommentar, nicht Verhalten) übernommen (§3).
- [x] `make gates`, `make archive-wave-test` und `make verify` grün — Ausgabe in eine Datei,
      Exit-Code getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

**Geliefert:** die vierte, real latente Stub-Erzeugungsstelle behoben, zwei weitere
MEDIUM-Inkonsistenzen im selben Werkzeug behoben — alle drei ausschließlich durch das
kontext-getrennte Review gefunden, keine davon durch die eigene Prüfung in slice-152/153.

**Lerneintrag — Form: geschärfte Regel.** *„Drei Stellen, alle behoben" war selbst eine
unvollständig geprüfte Behauptung — genau das Muster, das dieser ganze Review-Nachholvorgang
untersucht.* Das ist kein Einzelfall mehr, sondern der **vierte** Beleg für
[`BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen`](../observations/BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen/observation.md)
(nach `slice-142`, `slice-135`/`139`/`143`, `slice-150`/`151`) — diesmal nicht vom Maintainer,
sondern vom ersten tatsächlich kontext-getrennten Review dieser Session gefangen. Das ist der
Beleg, dass die Korrektur (Reviews kontext-getrennt fahren) wirkt: derselbe Fehlertyp, aber
diesmal von der dafür vorgesehenen Rolle gefunden, nicht von außen.

**Zwei beobachtbare Closure-Kriterien:**

1. `grep -c "VIER, nicht drei" tools/archive-wave/archive.go` → `1`;
   `go test`-Äquivalent `make archive-wave-test` grün mit den drei neuen/geänderten Tests.
2. `make gates`, `make archive-wave-test` und `make verify` je **Exit 0** auf dem Stand dieses
   Slice, Ausgabe in Dateien, Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Eines notiert: *Wellen-Modus meldet keine toten
Review-Verweise (§4).* **Ausgang: weiter offen**,
[Beobachtungs-Register](../observations/README.md) — kein Sensor-Kollisions-Fund, sondern eine
Funktionslücke; wird nicht in dieser Review-Korrektur behoben. Ein vierter Beleg für
`BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen` ist oben bereits eingetragen
(evidence/slice-158.md).

**Folge-Slices:** keine vergeben. Wellen-Modus-Dangling-Review-Meldung bleibt benannt, nicht
beauftragt.

## 8. Sub-Area-Modus

Berührt: **Gate-/Werkzeug-Schicht** (`tools/archive-wave/`) — Greenfield: alle drei Funde
getestet (`archive-wave-test`).
