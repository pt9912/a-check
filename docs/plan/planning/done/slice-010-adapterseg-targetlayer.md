# slice-010 — Layer-relativer `adapterSeg` + längster-Präfix (welle-10b/b1)

**Status:** done.
**Welle:** welle-10-regel-engine-generalisierung (Inkrement **b1**).
**Bezug:** Re-Evaluierungs-Trigger aus [ADR-0009](../../adr/0009-rollen-basierter-regel-dispatch.md); löst die `adapterSeg`-Namens-Generalisierung aus dem [AC-FA-RULE-006](../../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)-Out-of-Scope ein; Entscheidung [ADR-0010](../../adr/0010-layer-relativer-adapterseg-laengster-praefix.md).

---

## 1. Ziel

Drei Korrektheits-Cleanups aus den welle-10a-/b1-Reviews, alle rückwärtskompatibel:

1. **`adapterSeg` layer-relativ:** Adapter-Sub-Einheit = erstes Segment nach dem (längsten) Schicht-Glob-Präfix → `lateral-adapter` *vollständig* namensunabhängig (auch intra).
2. **`targetLayer` längster-Präfix:** spezifischster Glob gewinnt (verschachtelte Schichten; bei Gleichstand der zuerst deklarierte).
3. **Segment-bewusstes Matching (`segIndex`):** Präfix nur an Pfad-Segment-Grenzen → streng rückwärtskompatibel für `adapters/**` + Substring-Falschmatch (`io` in `audio`) geschlossen.

## 2. Umsetzung

- `rules.go`: Helfer `globPrefix`/`segIndex`/`layerByName`; `targetLayer` längster-Präfix; `adapterSeg(s, layer)` layer-relativ (längster Glob); `lateral` reicht die Schicht durch; `roleOf`-DRY.
- `spec/lastenheft.md` 0.3.0→0.4.0 (Out-of-Scope eingelöst), [ADR-0010](../../adr/0010-layer-relativer-adapterseg-laengster-praefix.md), Spezifikation 0.4.0.
- Tests: `TestForeignAdapterIntraLateral` (kehrt den welle-10a-Regression-Pin um), `TestSameAdapterSubunitNoLateral`, `TestTargetLayerLongestPrefix` (inkl. modul-qualifiziert, Pfadende, Reihenfolge, `audio`, Tie).

## 3. Definition of Done

- [x] Out-of-Scope der Schicht-Rollen-Anforderung eingelöst; Lastenheft 0.4.0 + Historie.
- [x] Folge-ADR `Accepted` + ADR-Index; Spezifikation 0.4.0.
- [x] Engine: `adapterSeg` layer-relativ + längster-Präfix + segment-bewusst; `make arch-check` (Dogfooding) unverändert grün.
- [x] Tests: fremde Intra-Adapter ⇒ `lateral`, gleiche Sub-Einheit ⇒ kein `lateral`, verschachtelte/modul-qualifizierte Auflösung, Tie-Break.
- [x] Multi-Linsen-Review (2 Linsen) + `segIndex`-Delta-Review bestanden.

## 4. Closure-Notiz (nach `done/`)

**Belege:** `make gates` grün (`arch-check` 0 Befunde — Dogfooding unverändert; die driven-Adapter bekommen jetzt Sub-Einheiten, importieren sich aber nicht). Der welle-10a-Regression-Pin ist sauber **umgekehrt** (kein-`lateral` → `lateral`) — beweist den kontrollierten a→b-Übergang. Das `segIndex`-Delta-Review pinnte die Zusage „streng identisch für `adapters/**`" über 11 Vektoren empirisch (0 Mismatches), keine OOB-/Terminierungs-Fehler.

**Lerneintrag:**

- *Segment-bewusstes Matching schlägt Substring:* `strings.Contains`-Präfixe matchen `io` in `audio`; `segIndex` (an Grenzen geankert) macht die Auflösung korrekt **und** die Rückwärtskompat streng.
- *Getracter Review fängt, was Gates nicht fangen:* der Reviewer belegte die „identisch"-Zusage über 11 Vektoren — Gates messen Zeilen, nicht Äquivalenz.

**Folge (welle-10b):** `app`-Rolle, `driving`/`driven`-Ports; `LayerOf` längster-Präfix (Asymmetrie zu `targetLayer`).
