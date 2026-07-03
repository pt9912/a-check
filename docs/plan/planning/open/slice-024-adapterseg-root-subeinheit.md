# slice-024 — Root-Sub-Einheit für `adapterSeg` (entsperrt die b-cad-Richtungs-Config)

**Status:** open — **Entwurf zur Abnahme** (2026-07-03). Entscheide §6 **vor** der Umsetzung.
**Bezug:** Change Request an [AC-FA-RULE-002](../../../../spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter)
(Sub-Einheiten-Definition); **neuer Folge-ADR [ADR-0019](../../adr/0019-adapterseg-root-subeinheit.md)**
(Proposed) zu [ADR-0010](../../adr/0010-layer-relativer-adapterseg-laengster-praefix.md); schärft
[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung).
[Roadmap](../in-progress/roadmap.md). **Trigger:** b-cad-Pilot (Meilenstein M3) — Maintainer-Urteil
„Richtungen bewusst setzen" ist ohne diesen Fix nur um den Preis von 40 Falsch-Positiven umsetzbar.

> **Hinweis:** Der AC-Text in §3 (Code-Fence) war der Entwurf zur Abnahme (Vor-D-Stand);
> maßgeblich ist die freigegebene Fassung in [`spec/lastenheft.md`](../../../../spec/lastenheft.md)
> (0.15.0, inkl. Blatt-Klassifikation aus Entscheid D).
> [ADR-0019](../../adr/0019-adapterseg-root-subeinheit.md) bleibt `Proposed`, bis Entscheid D
> abgenommen ist.

---

## 1. Auslöser (Pilot-Evidenz, verifiziert am v0.8.0-Image)

Die b-cad-Vollrichtungs-Config (pro-Adapter-Layer mit `direction`, nach den dortigen Umbauten
slice-028/029) liefert **40 × `lateral-adapter`** — ausnahmslos Same-Directory-Includes
(`io/dxf_reader.cpp → "adapters/io/dxf_reader.h"`), **null echte Verstöße**, null Richtungs-,
Services- oder Kern-Befunde. Ursache: das Sub-Einheiten-Segment aus
[ADR-0010](../../adr/0010-layer-relativer-adapterseg-laengster-praefix.md) degeneriert für
Dateien **direkt im Layer-Root** zum **Dateinamen**. Pro-Adapter-Layer sind aber die
Voraussetzung jeder `direction`-Modellierung
([AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch):
eine Richtung je Schicht) — die Falsch-Positiv-Klasse blockiert also genau den vom Maintainer
beauftragten Ausbau. Dateinamen sind keine Architektur-Einheiten; Sub-Einheiten sind — wie
Schichten — Verzeichnisse.

## 2. Betroffene Artefakte (vor der Implementierung benannt)

- **Slice-ID:** slice-024.
- **ADR:** [ADR-0019](../../adr/0019-adapterseg-root-subeinheit.md) (neu; `Proposed` →
  `Accepted` erst nach Abnahme **inkl. Umsetzungs-Entscheid D**, §6; erweitert
  [ADR-0010](../../adr/0010-layer-relativer-adapterseg-laengster-praefix.md), `Supersedes: —`).
- **AC:** [AC-FA-RULE-002](../../../../spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter)
  (Sub-Einheiten-Grenzfall Root); Bezug ohne Änderung:
  [AC-FA-RULE-006](../../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)
  (`adapterSeg`-Generalisierung, wie [ADR-0010](../../adr/0010-layer-relativer-adapterseg-laengster-praefix.md)).
- **Spec:** [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)
  (`lateral-adapter`-Zeile).
- **Module:** `internal/hexagon/core` (`adapterSeg`), Tests in `rules_test.go` + ein
  CLI-Integrationsfall; Benutzerhandbuch-Halbsatz (Sub-Einheiten sind Verzeichnisse).
- **Version:** Lastenheft + Spezifikation **0.14.0 → 0.15.0**.
- **Gates:** `make gates` → `make ci`; Multi-Linsen-Review + ggf. Delta → `docs/reviews/`.

## 3. Entwurf (zur Abnahme) — AC-FA-RULE-002-Ergänzung

```text
AC-FA-RULE-002 (Sub-Einheiten-Grenzfall präzisiert): Innerhalb einer Schicht
werden Adapter-Sub-Einheiten relativ zum Schicht-Glob-Präfix unterschieden;
eine Datei, deren Pfad-Rest nach dem Präfix KEIN weiteres Verzeichnis enthält
(Datei direkt im Layer-Root), gehört zur Root-Sub-Einheit ''. Importe zwischen
Root-Dateien derselben Schicht sind damit keine lateralen Kanten; Root ↔
Unterverzeichnis und verschiedene Unterverzeichnisse bleiben lateral,
Cross-Layer-Adapter-Importe bleiben kategorisch.

Neue/ergaenzte Akzeptanzkriterien:
- Happy (Root↔Root): Given zwei Dateien direkt im Root desselben
  Adapter-Layers (x.cpp importiert x.h), when a-check laeuft, then KEIN
  lateral-adapter (Root-Sub-Einheit).
- Boundary (Root↔Unterverzeichnis): Given eine Root-Datei, die eine Datei
  eines Unterverzeichnisses derselben Schicht importiert (oder umgekehrt),
  when a-check laeuft, then lateral-adapter (verschiedene Sub-Einheiten).
- Negative (Bestand): Given Importe zwischen zwei Unterverzeichnissen
  derselben Schicht oder zwischen zwei Adapter-Layern, when a-check laeuft,
  then lateral-adapter wie bisher (kategorisch; adapter_sink-Ausnahme
  unveraendert).

Out-of-Scope: Datei-granulare Sub-Einheiten als Opt-in (Re-Eval ADR-0019);
jede Aenderung an Cross-Layer-lateral oder adapter_sink.
```

## 4. Umsetzungsplan (Reihenfolge: ADR → Lastenheft → Spec → Code → Tests)

1. [ADR-0019](../../adr/0019-adapterseg-root-subeinheit.md) `Proposed → Accepted` — erst mit
   der Abnahme **inkl. Entscheid D** (§6); [ADR-Index](../../adr/README.md) ist ergänzt.
2. Lastenheft: §3-Text einarbeiten, Bump **0.15.0** + Historie-Zeile;
   [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung)-`lateral-adapter`-Zeile
   + Bump + Historie.
3. `adapterSeg` (`rules.go`): enthält `rest` nach dem Präfix keinen `/`, Rückgabe `""`
   (statt des Dateinamens) — ein Zweig; Doc-Kommentar mit
   [ADR-0019](../../adr/0019-adapterseg-root-subeinheit.md)-Referenz.
   **Achtung Wechselwirkung:** `adapterSeg` liefert `""` bisher auch für „Präfix matcht
   nicht" — beide `""`-Fälle vergleichen gleich; prüfen, ob der Nicht-Match-Fall (`bestEnd < 0`)
   von Root unterscheidbar bleiben muss (heute: Nicht-Match ⇒ `""` ⇒ gleich ⇒ kein lateral —
   Verhalten dokumentieren, Test pinnen).
4. Tests (Mutanten-Boundary): Root↔Root kein lateral (pinnt den Fix; bricht am Alt-Code);
   Root↔Subdir lateral; Subdir↔Subdir lateral (Bestand); Cross-Layer lateral (Bestand);
   uniform über `relative`-/`fixed-root`-Kandidaten; CLI-Integration mit
   b-cad-förmigem Fixture (flacher Adapter, `.cpp→.h`).
5. Benutzerhandbuch: Halbsatz im Glossar/§4 („Sub-Einheiten sind Unterverzeichnisse;
   Dateien im Schicht-Root bilden eine gemeinsame Root-Einheit"), Historie-Zeile.
6. `make gates`/`make ci`; Multi-Linsen-Review + ggf. Delta; **Verifikation:** die
   verifizierte b-cad-V4-Config gegen das lokal gebaute Image ⇒ **0 Befunde**; Closure.

## 5. Definition of Done

- [ ] [ADR-0019](../../adr/0019-adapterseg-root-subeinheit.md) **Accepted** + Index.
- [ ] Lastenheft + Spezifikation **0.15.0** mit Historie-Zeilen.
- [ ] `adapterSeg`-Root-Zweig; Dogfooding grün (0 Befunde).
- [ ] Tests §4.4 inkl. Alt-Code-brechendem Root↔Root-Pin + CLI-Integration.
- [ ] Benutzerhandbuch nachgezogen.
- [ ] `make gates` + `make ci` grün; Review bestanden; b-cad-V4-Gegenprobe 0 Befunde.
- [ ] **Maintainer-Abnahme der Entscheide A–C (§6).**
- [ ] Closure: reiner `git mv` nach `done/`; 2 beobachtbare Kriterien + Lerneintrag.

## 6. Offen / Entscheidungen zur Abnahme

> **Abnahme (2026-07-03):** Entscheide A–C gemäß Empfehlung bestätigt (Maintainer-Wort).
> **Entscheid D (Umsetzungs-Fund, ausstehend):** Das Dogfooding deckte auf, dass die reine
> Root-Regel **Go-Paket-Blätter** falsch behandelt — `report_test.go` (externes Testpaket)
> importiert sein eigenes Paket `…/driven/report`: der Kandidat endet auf dem
> Paket-**Verzeichnis**, die Root-Deutung machte daraus einen neuen Falsch-Positiv und
> hätte zugleich Cross-Paket-Lateral in Go-Repos geblendet. Umgesetzt per dokumentierter
> Empfehlung: **Blatt-Klassifikation** — datei-förmiges Blatt (`.`) → Root `''`,
> verzeichnis-förmiges Blatt → ist die Sub-Einheit; endungslose Datei-Specifier
> (TS `./b`) als dokumentierte [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Grenze
> + [ADR-0019](../../adr/0019-adapterseg-root-subeinheit.md)-Re-Eval-Trigger.
> ADR bleibt `Proposed` bis D abgenommen ist (Lerneintrag slice-023 angewandt:
> Alt-Code-Semantik erst am Dogfooding verifiziert, Abweichung deklariert statt still).

- **Entscheid A — Root-Sub-Einheit `''` als Default (kein Opt-in):** Die Datei-Ebene war nie
  eine Architektur-Aussage; ein Schalter würde den Degenerat-Fall konservieren. *Empfehlung:
  Default, Opt-in-Datei-Granularität nur als Re-Eval-Trigger im ADR.*
- **Entscheid B — Lockerungs-Charakter explizit:** Für Sammel-Layer mit flachen Dateien im
  Root entfällt die bisherige (dateinamen-basierte) Trennung — das ist eine bewusste
  **Gate-Lockerung per ADR** ([AGENTS §3.6](../../../../AGENTS.md#36-gates-dürfen-nicht-ohne-adr-gelockert-werden)-konform),
  begründet als Falsch-Positiv-Tilgung (b-cad: 40/40 falsch, 0 echt). *Empfehlung: bestätigen.*
- **Entscheid C — Nicht-Match-`""` bleibt wie bisher:** `adapterSeg` liefert `""` auch, wenn
  der Layer-Präfix im Symbol nicht vorkommt (Alt-Verhalten: kein lateral) — der Root-Fall
  fällt semantisch zusammen; keine Sonder-Markierung. *Empfehlung: bestätigen + per Test
  pinnen (Alt-Code-Verifikation dokumentiert — Lerneintrag slice-023 angewandt).*
- **Notiz — MeshSource-Naht (b-cad):** Die Kante `ui_command → ui_view` (Interface-Naht)
  braucht unabhängig von diesem Slice `adapter_sink` (nimmt sie aus dem kategorischen
  `lateral`) **plus** deklarierte Kante (für `wrong-direction`) — in der V4-Config
  dokumentiert, kein a-check-Delta.

## 7. Closure-Notiz (nach `done/`)

_(folgt bei Closure.)_
