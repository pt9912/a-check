# ADR-0026 — HexSlice Vertical-Slice-Regeln: `lateral-slice` + `port-locality`

- **Status:** Accepted
- **Datum:** 2026-07-24
- **Autor:** pt9912
- **Bezug:** [AC-FA-RULE-009](../../../spec/lastenheft.md#ac-fa-rule-009--slice-isolation-regel-lateral-slice), [AC-FA-RULE-010](../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality), [AC-FA-RULE-006](../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
- **Schärft:** [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) — zwei neue Regeln über das bestehende Rollenmodell.
- **Supersedes:** —

## Kontext

Das Rollenmodell ([AC-FA-RULE-006](../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung))
erzwingt die **hexagonale** Achse (Domain-/App-Reinheit, Kanten, `tech`-Kapselung, Adapter-Lateralität),
aber **nicht** die **Vertical-Slice**-Achse, die HexSlice von gewöhnlicher Hexagonal unterscheidet:
(1) Use-Case-Slices koppeln nicht an fremde Slice-Interna; (2) Ports leben „so lokal wie möglich, so
gemeinsam wie nötig". Ein realer HexSlice-Go-Konsument (`hexslice-architecture/lab/examples/go`) liefert
die Evidenz — beide Verstöße bleiben heute unentdeckt.

**Prerequisite (verifiziert):** die Ziel-Auflösung klassifiziert einen Import nur, wenn das Ziel-Glob ein
**sauberes literales Präfix** trägt (Segment-Run-Match, modul-präfix-tolerant). Globs mit Wildcard in der
Mitte (`…/**/ports/**`) oder Datei-Endung (`…/*.go`) lösen als Import-**Ziel** nicht auf — im Beispiel
sind `app→app`/`app→port` damit heute `extern`, die Slice-Kanten also gar nicht erzwungen.

## Entscheidung

Zwei neue, **kategorische** Regeln über das Rollenmodell, beide **pfad-heuristisch** auf dem vorhandenen
modul-präfix-toleranten Segment-Match:

1. **`lateral-slice` (Slice-Isolation, Glob-pro-Slice):** die **Slice-Identität** einer `app`-Datei ist
   das **längste `app`-Rollen-Glob-Literalpräfix**, das ihren Pfad matcht. Eine `app`-Datei, die ein
   `app`-Ziel **derselben Schicht** (`tl == f.Layer`) mit **anderer** Slice-Identität importiert, ist ein
   Befund. Zwei Slices = zwei Globs **derselben** `app`-Schicht; ein einziges breites Glob lässt die Regel
   inert (opt-in). **Getrennte `app`-Layer** (z. B. `services`/`services_geo`) sind **edge-regiert**, nicht
   `lateral-slice` — sonst bräche jeder klassisch-hexagonale Sub-Service mit deklarierter Kante.
2. **`port-locality` (pfad-abgeleiteter Scope):** der **Scope** eines Ports ist das matchende
   **`port`-Rollen-Glob-Literalpräfix minus seinem letzten Segment** (dem Port-Ordner-Marker, typisch
   `ports`). Die Regel greift **nur für im App-Baum geschachtelte Ports** (`appTreeContains(portScope)`:
   der Scope ist Vorfahr eines `app`-Glob-Präfixes). Eine `app`-Datei, deren Pfad diesen Scope **nicht
   enthält**, verletzt die Lokalität. **Geschwister-Ports** (klassisch, `…/ports` neben `…/services`) →
   inert. Nur `app`-Importeure; ein Adapter, der einen Port implementiert, ist nicht erfasst.
3. Beide **kategorisch** (nicht über `edges`/`allow` aufhebbar), `composition_root` befreit. Erst-Treffer:
   `… → lateral-adapter → lateral-slice → tech-leak → port-direction-mismatch → port-locality →
   wrong-direction` (`port-locality` **vor** `wrong-direction`, damit eine erlaubte `app→port`-Kante die
   Lokalitäts-Verletzung nicht maskiert).
4. **Kein Config-Schema.** Beide leiten aus vorhandenen `layers`-Globs + Pfaden ab. Die saubere-Präfix-Glob-
   **Disziplin** (per-Slice, per-Port-Ordner) ist Voraussetzung und wird als
   [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Grenze
   ausgewiesen, nicht still.

## Konsequenzen

- **Opt-in / rückwärtskompatibel:** ohne per-Slice-`app`-Globs bzw. ohne Port-Layer feuert keine der Regeln;
  bestehende Configs (auch a-checks Eigen-`.a-check.yml`, ohne `app`-Slices/nested Ports) ändern sich nicht
  → `make arch-check` bleibt 0.
- **Konsumenten-Disziplin:** wer die Vertical-Slice-Achse gaten will, schreibt `app`-/`port`-Globs mit
  sauberen literalen Präfixen (per-Slice `…/createorder/**`, per-Port-Ordner `…/ports/**`). Die
  Beispiel-Config wird entsprechend korrigiert mitgeliefert.
- **Geteilter Slice-Code** (nicht über Ports) wird von `lateral-slice` gemeldet — bewusst; ein `app_sink`
  (analog `adapter_sink`) bleibt ein künftiges Inkrement.

## Verworfene Alternativen

- **`layerOfCand` auf Voll-Glob-Match vertiefen** (damit `…/**/ports/**`/`…/*.go` auch als Import-Ziel
  auflösen, Original-Config unverändert): verworfen — der verankerte Voll-Glob-Regex bräche die
  **Modul-Präfix-Toleranz** des Segment-Matches (`hexslice/example/internal/…`), ein größerer, riskanter
  Resolver-Umbau mit eigenem ADR. Config-Disziplin ist die kleinere, ehrlichere Wahl.
- **Segment-Sub-Einheit für Slices** (wie `adapterSeg`, erstes Segment nach einem breiten `app/**`):
  verworfen — die Granularität hinge an der Glob-Tiefe (bei `application/**` wäre die „Slice" die
  Business-Area statt der Use-Case).
- **Explizite `scope:`-Deklaration je Port-Sub-Layer:** verworfen zugunsten der pfad-abgeleiteten,
  zero-config-Ableitung (deckt die drei HexSlice-Ebenen automatisch).

## Fitness Function

- `make test`: AC-Tests je Regel (Happy/Boundary/Negative, kategorisch mit `allow`-Kante) + Integrations-
  Fixture aus dem realen Go-Beispiel (positiv 0 Befunde; injizierte Cross-Slice-/Cross-Scope-Verstöße
  gefangen).
- `make arch-check` (Dogfooding): unverändert **0** (keine `app`-Slices/nested Ports in der Eigen-Config).
- **Konsumenten-Regressionsprobe** (Review-Härtung): die klassisch-hexagonalen Bestandskonsumenten
  b-cad (sibling-Ports + `services`/`services_geo`-Sub-Layer), d-check und d-migrate bleiben **0** —
  die Same-Layer-Einschränkung (`lateral-slice`) und der `appTreeContains`-Guard (`port-locality`)
  verhindern die 27 Falsch-Positive der ersten Fassung (Regressionstests im `core`-Paket).

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-07-24 | Proposed → Accepted (Sign-off Auftraggeber; Wahl Option B+C der HexSlice-Gap-Analyse, Design-Entscheide Glob-pro-Slice + pfad-abgeleiteter Scope + Config-Disziplin per Frage bestätigt). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
