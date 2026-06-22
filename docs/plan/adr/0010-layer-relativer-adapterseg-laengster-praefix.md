# ADR-0010 — Layer-relativer `adapterSeg` + längster-Präfix-Auflösung

- **Status:** Accepted
- **Datum:** 2026-06-22
- **Autor:** pt9912
- **Bezug:** [AC-FA-RULE-006](../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung) (löst die `adapterSeg`-Namens-Generalisierung aus dem Out-of-Scope ein; Lastenheft 0.3.0→0.4.0), [AC-FA-RULE-002](../../../spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter) (`lateral-adapter`)
- **Schärft:** [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) — macht die layer-relative Adapter-Sub-Einheit und die längster-Präfix-Auflösung verbindlich.
- **Supersedes:** —

## Kontext

[ADR-0009](0009-rollen-basierter-regel-dispatch.md) (welle-10a) generalisierte den
Regel-Dispatch über Rollen, verschob aber zwei heuristik-gebundene Stellen bewusst:

1. `lateral()` unterschied zwei Adapter über `adapterSeg`, das das **Literal**
   `"adapters"` im Pfad suchte. Für fremd benannte `role: adapter`-Schichten
   (`src/geometry`, `io`) lieferte es `""` → die Intra-Schicht-Unterscheidung griff
   nicht (dokumentiertes Falsch-Negativ, in welle-10a als Regression-Pin
   festgeschrieben).
2. `targetLayer` löste einen Import über den **ersten** passenden Glob-Präfix
   (Substring) auf — bei überlappenden/verschachtelten Schichten gewann nicht der
   spezifischste.

## Entscheidung

1. **`adapterSeg` layer-relativ:** die Adapter-Sub-Einheit ist das erste
   Pfad-Segment **nach dem Glob-Präfix der Schicht** (Helfer `globPrefix`).
   Greift für beliebige Layer-Namen; für `adapters/**` **streng identisch** zum
   bisherigen Verhalten.
2. **`targetLayer` längster-Präfix:** der spezifischste (längste) matchende
   Glob-Präfix gewinnt statt des ersten.
3. **Segment-bewusstes Matching (`segIndex`):** ein Präfix matcht nur an
   **Pfad-Segment-Grenzen** (Beginn oder nach `/`, Ende vor `/` oder Pfadende) —
   `io` matcht nicht in `audio`. Das macht (1) streng rückwärtskompatibel und
   schließt Substring-Falschmatches in `adapterSeg` **und** `targetLayer`.

## Konsequenzen

- `lateral-adapter` ist **vollständig** namensunabhängig (Cross-Layer **und**
  Intra-Layer) — die Out-of-Scope-Klausel der Schicht-Rollen-Anforderung ist eingelöst.
- **Verhaltensänderung:** ein fremd benannter Intra-Adapter, der eine andere
  Sub-Einheit importiert, feuert jetzt `lateral-adapter` (der welle-10a-Regression-Pin
  ist umgekehrt).
- **Rückwärtskompat:** a-checks Dogfooding bleibt grün — die driven-Adapter
  (`internal/adapter/driven/{config,extract,report}`) bekommen jetzt zwar
  Sub-Einheiten, importieren sich aber nicht gegenseitig ⇒ kein `lateral`.
- Verschachtelte Schichten lösen korrekt auf (spezifischster Präfix).
- Lastenheft 0.3.0 → **0.4.0**.

## Fitness Function

- `make arch-check` (Dogfooding, [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  unverändert 0 Befunde.
- `make test`: `TestForeignAdapterIntraLateral` (fremde Namen intra ⇒ `lateral`),
  `TestTargetLayerLongestPrefix` (verschachtelte Schichten).

## Re-Evaluierungs-Trigger

- `LayerOf` (eigene Schicht einer Datei) bleibt Erst-Treffer (anderer Mechanismus,
  Glob-Regex-Match); bei verschachtelten Schichten kann das von `targetLayer`s
  längster-Präfix abweichen — Angleichung bei Bedarf.
- welle-10b weiter: `app`-Rolle, `driving`/`driven`-Ports.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-06-22 | Proposed — welle-10b (b1); löst `adapterSeg`-Namens-Generalisierung + längster-Präfix-Auflösung. |
| 2026-06-22 | Proposed → Accepted (Sign-off Auftraggeber; Multi-Linsen-Review + `segIndex`-Delta-Review bestanden, segment-bewusstes Matching ergänzt). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
