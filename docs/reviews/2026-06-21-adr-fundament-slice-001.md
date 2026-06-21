# Review-Report — Fundament-ADRs (slice-001)

- **Review-Art:** Design-/Plan-Review (ADRs, vor Implementierung — kein Diff)
- **Gegenstand:** ADR-0001…0004 + ADR-Index (`docs/plan/adr/`)
- **Datum:** 2026-06-21
- **Modell:** Opus 4.8
- **Methode:** Drei unabhängige Reviewer-Agenten mit frischem Kontext und
  perspektiven-diversen Linsen (Lastenheft-Bezug · Regelwerk-/Konventions-
  Konformität · architektonische Inhalts-Qualität), Synthese durch einen
  vierten Kontext mit **adversarischer Verifikation der HIGH-Befunde** gegen
  die realen Repo-Artefakte (Modul 11).
- **Eingangs-Kontext:** `spec/lastenheft.md`, `AGENTS.md`, `harness/conventions.md`,
  `docs/plan/adr/README.md`, `d-check.mk`, `.d-check.yml`; Regelwerk v1.3.0
  Module 4/8/10 + `grundlagen-konventionen.md`.
- **Skill-Version:** keine — `.harness/skills/reviewer.md` existiert nicht (siehe I4).

Befund-Schema (Modul 10): `kategorie` · `quelle` · `pfad` · `befund` (beobachtbar,
ohne Lösungsvorschlag) · `verifizierbar`. Kategorien: HIGH blockiert · MEDIUM vor
Acceptance klären · LOW nice-to-fix · INFO Hinweis.

## Befunde

### MEDIUM

**M1 — d-check `--print-mk` als Vorbild überzeichnet**
- quelle: Faktencheck (Inhalts-Qualität)
- pfad: `docs/plan/adr/0004-distribution-image-mk.md` (Kontext/Entscheidung)
- befund: ADR-0004 sagt, d-check „liefert exakt dieses Muster" und das Repo
  „dogfoodet es bereits". `d-check.mk:5-7` dokumentiert jedoch ausdrücklich,
  dass d-check **noch kein** `--print-mk` hat (Zielbild, „von Hand gepflegt").
  Das include-bare `d-check.mk` existiert real und ist digest-gepinnt; das
  `--print-mk`-Erzeugungs-Muster ist hingegen geteiltes Zielbild, nicht
  belegter Stand.
- verifizierbar: ja (`grep` auf `d-check.mk`)

**M2 — strict-decode „konsistent mit d-check's `.d-check.yml`" unbelegt**
- quelle: Faktencheck (Inhalts-Qualität)
- pfad: `docs/plan/adr/0003-config-modell-a-check-yml.md` (Option 1)
- befund: Die real vorhandene `.d-check.yml` enthält weder ein `strict`-/
  `fail-closed`-Feld noch eine entsprechende Aussage (grep: 0 Treffer). Die
  Attribution von strict-decode/fail-closed an d-check ist aus den
  Repo-Artefakten nicht belegbar. (Reviewer meldete HIGH; nach Verifikation
  MEDIUM — nicht *falsch*, aber unbelegt; der YAML-Config-Ansatz selbst ist
  korrekt als Stack-Muster.)
- verifizierbar: ja (`grep -i strict .d-check.yml`)

**M3 — ADR deklariert sich als „Technik-Quelle" (Stratum-Rolle)**
- quelle: AGENTS.md §3.4 / Regelwerk §Spec-Straten
- pfad: `docs/plan/adr/0001-go-impl-sprache.md` (Schärft-Feld-Begründung; analog 0002/0003/0004)
- befund: Die Formulierung „bis dahin ist diese ADR die Technik-Quelle" weist
  der ADR eine Stratum-Rolle zu, die formal dem Technik-Stratum
  (`spec/spezifikation.md`) vorbehalten ist. Die ADR ist laut Regelwerk die
  Begründungs-Schicht *unter* den Spec-Straten, kein Stratum-Ersatz. Klärung:
  zulässige Bootstrap-Brücke (Phase < 4) oder Stratum-Vermischung?
- verifizierbar: ja (Wortlaut)

**M4 — Pin-Staleness als Konsequenz unerwähnt**
- quelle: Vollständigkeit der Entscheidung
- pfad: `docs/plan/adr/0004-distribution-image-mk.md` (Konsequenzen)
- befund: Das Modell „digest-gepinnt + Pin-Hebung als Commit" trägt eine
  Konsequenz, die nicht benannt ist: manueller Pin-Hebungs-Aufwand pro
  Konsument (kein Auto-Update) — dieselbe Drift-Klasse, die das Tool reduziert.
- verifizierbar: nein

**M5 — ADR-Index gibt Bezug verkürzt/divergent wieder**
- quelle: Lastenheft-/Index-Konsistenz
- pfad: `docs/plan/adr/README.md` (Bezug-Spalte)
- befund: Die Bezug-Spalte führt je ADR genau eine AC-ID, während die
  ADR-Köpfe mehrere tragen (ADR-0001: AC-QA-01/-02/-03 + AC-FA-DIST-001;
  ADR-0004: AC-FA-DIST-001 + AC-QA-03 + AC-QA-02). Der Index gibt den Bezug
  damit unvollständig wieder (Spaltenschema deklariert keinen „Primär-Bezug").
- verifizierbar: ja

### LOW

**L1 — AC-QA-01 begründend, aber nicht im Kopf-Bezug** · `0002` (Konsequenzen vs. Bezug-Feld) · verifizierbar: ja
ADR-0002 zitiert AC-QA-01 (Determinismus) in den Konsequenzen, führt sie aber nicht im `Bezug`-Feld.

**L2 — AC-FA-RULE-* tragend, aber nicht im Bezug** · `0002`/`0003` · verifizierbar: ja
Beide stützen die Entscheidung auf „die fünf Regeln (AC-FA-RULE-*)", ohne die Familie im Bezug zu führen.

**L3 — Rust-Trade-off nicht abgewogen** · `0001` (Optionen/Entscheidung) · verifizierbar: nein
Der genannte Rust-Vorteil „stärkere Korrektheitsgarantien" verschwindet in der Entscheidung folgenlos.

**L4 — Slice-Provenance außerhalb der Historie-Zone** · `0001-0004` + Index · verifizierbar: ja
`slice-002`/`slice-003`-Tokens stehen im Body (als Provenance/Verifikations-Zeiger, semantisch erlaubt), nicht in einer `## Historie`-Zone; ein künftiger token-basierter Abwärts-Check (`check-references`) würde sie fangen.

**L5 — Bezug AC-FA-CLI-001 ist Nebenbezug** · `0003` · verifizierbar: nein
Der tragende Vertrag der Config-Entscheidung ist AC-FA-CONF-001; AC-FA-CLI-001 regelt nur die allgemeine Exit-Code-Semantik.

**L6 — Themen-Überlappung ohne Querverweis** · `0001` ↔ `0004` · verifizierbar: ja
ADR-0001 behandelt den digest-gepinnten Distributionspfad mit (Kern von ADR-0004), ohne dorthin zu verweisen.

### INFO

**I1 — distroless/static korrekt gebunden** · `0001` · verifizierbar: ja
distroless/static ist durchweg an AC-QA-02 (a-check) gebunden; „Konsistent mit d-check (Go)" bezieht sich eng auf Go. (Reviewer meldete HIGH-Konflation; nach Verifikation entkräftet — Rest ist Lese-Ambiguität.)

**I2 — read-only-Zuschreibung an `--print-mk` ist Verallgemeinerung** · `0004` · verifizierbar: ja
Die AC-FA-DIST-001-Boundary nennt explizit nur `--print-config` als read-only; ADR-0004 dehnt das (plausibel) auf `--print-mk` aus.

**I3 — Fitness-Function-Anker uneinheitlich konkret** · `0001` vs. `0002/0003/0004` · verifizierbar: ja
ADR-0002/0003/0004 tragen je einen konkreten slice-003-prüfbaren Anker; ADR-0001s Anker ist eher Sammel-Gate (lint/test/Build) als entscheidungsspezifischer Test.

**I4 — keine Reviewer-Skill-Datei** · `.harness/` (leer) · verifizierbar: ja
`.harness/skills/reviewer.md` existiert nicht; ein Reviewer-Agent driftet ohne Skill zwischen Sessions (Modul 10). Steering-Loop-Kandidat.

## Negativbefunde (geprüft, ohne Befund)

- **MADR-Format & Dateinamen:** alle vier ADRs mit vollständigen Kopf-Feldern (Status/Datum/Bezug/Schärft/Supersedes) und Body-Blöcken (Kontext/Optionen-mit-Trade-offs/Entscheidung/Konsequenzen); `0001`–`0004` vierstellig, kebab. Konform.
- **Erfundene Anforderungen:** keine — alle Aussagen (Hermetik, Determinismus, fail-closed, digest-Pin, opt-in-Re-Eval) aus AC-Text/§1 herleitbar; ADRs begründen Lösung, nicht Anforderung.
- **Out-of-Scope-Treue:** Toolchain-Backends (0002), Includes (0003), Binary-Releases (0004) decken sich exakt mit den Out-of-Scope-Zeilen des Lastenhefts.
- **Schärft-Verbot:** kein `Schärft:` zeigt aufs Lastenheft (alle `—`, konsistent mit „spezifikation.md erst slice-002").
- **Immutability & Status:** alle `Proposed` (editierbar), konsistent mit §3.5 und offenem Schärft-Ziel.
- **ID-Linkpflicht & Anker:** AC-/ADR-Erwähnungen im Body als Markdown-Link mit auflösenden Ankern; Familien-Glob `AC-FA-RULE-*` korrekt unverlinkt; Index-Zeilen verlinkt (bestätigt durch grünes `make doc-check`).
- **Referenz-Richtung (SDP):** keine verbotene Entscheidungsgrundlage über einen Slice; Slice-Nennungen durchweg Provenance/Verifikations-Zeiger.
- **Hard Rule §3.1 (Docker/make-only):** durchgängig eingehalten (Go-Toolchain in Docker, kein Host-Go).

## Kategorie-Summary

| Kategorie | Anzahl | IDs |
|---|---|---|
| HIGH | 0 | — (zwei gemeldete HIGH nach Verifikation auf M2/I1 herabgestuft) |
| MEDIUM | 5 | M1, M2, M3, M4, M5 |
| LOW | 6 | L1–L6 |
| INFO | 4 | I1–I4 |

## Verdikt

Kein blockierender (HIGH) Befund nach adversarischer Verifikation. Das ADR-Set
ist MADR-/Konventions-/Linkpflicht-konform und gate-grün. Vor Acceptance
(slice-002) sollten die fünf MEDIUM geklärt werden — vorrangig die
d-check-Faktentreue (M1, M2), die Stratum-Formulierung (M3) und die
Index-Bezug-Konsistenz (M5). Da die ADRs `Proposed` sind, ist jetzt der
billigste Zeitpunkt (Modul 10: frühes Finding = billiges Finding).

## Disposition (Implementer, 2026-06-21)

Behebung als getrennter Implementer-Schritt nach dem Review. Beleg:
`make doc-check` grün — 11 Dateien, 0 Befunde.

| Finding | Aktion |
|---|---|
| M1 | ADR-0004: „exakt dieses Muster" entschärft — include-barer `d-check.mk`-Teil als real, `--print-mk`-Erzeugung als geteiltes Zielbild (d-check hat sie noch nicht) benannt. |
| M2 | ADR-0003: strict-decode nicht mehr d-check zugeschrieben; nur der YAML-Config-Ansatz als Stack-gemeinsam, strict-decode als `a-check`s bewusste Wahl. |
| M3 | ADR-0001 Schärft: „diese ADR die Technik-Quelle" entfernt; ADR als Begründungs-Schicht *unter* den Straten, `spec/spezifikation.md` als Technik-Stratum benannt. |
| M4 | ADR-0004: Pin-Staleness (manuelle Pin-Hebung pro Konsument) als Konsequenz/Trade-off ergänzt. |
| M5 | ADR-Index: Bezug-Spalte je ADR vollständig (alle Kopf-ACs verlinkt). |
| L1 | ADR-0002 Bezug: `AC-QA-01` ergänzt. |
| L2 | ADR-0002/0003: `AC-FA-RULE-*` als *nachgelagerte Konsumenten* geklärt, begründende AC (EXTRACT-001/CONF-001) benannt. |
| L3 | ADR-0001: Rust-Korrektheitsvorteil abgewogen (Risiko liegt in Mustern/Config, nicht Speichersicherheit). |
| L4 | Bewusst beibehalten: Slice-Nennungen sind Verifikations-Zeiger/Provenance (Regelwerk-Matrix erlaubt ADR→Slice als Kontext); kein struktureller Umzug nötig. |
| L5 | Bewusst beibehalten: `AC-FA-CONF-001` bleibt primär (zuerst gelistet); `AC-FA-CLI-001` als Exit-Code-Vertrag im Bezug belassen. |
| L6 | ADR-0001: Querverweis auf ADR-0004 (Distributionsmodell) ergänzt. |
| I1 | ADR-0001: „Konsistent mit d-check" auf die *Sprachwahl* Go eingegrenzt. |
| I2 | ADR-0004: read-only von `--print-config` (Boundary-AK) und `--print-mk` (konsistente Design-Folge) getrennt. |
| I3 | ADR-0001: entscheidungsspezifischer Fitness-Anker (statisch-gelinktes-Binary-Build-Check) ergänzt. |
| I4 | `.harness/skills/reviewer.md` angelegt (Reviewer-Skill, Modul 10). |
