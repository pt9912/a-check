# Nachtrag zum Review-Report Etappe C — F-2 zurückgezogen — 2026-07-26

**Review-Art:** Nachtrag zu [`2026-07-26-etappe-c-slice-055-056.md`](2026-07-26-etappe-c-slice-055-056.md).
Eigene Datei statt Überschreibung, nach
[`docs/reviews/README.md`](README.md) und Regelwerk Modul 10 („ein Report pro Lauf; Folgeläufe als
neue Datei" — Auditierbarkeit). Der **Befundtext** des Ursprungs-Reports bleibt unverändert als
Beleg seines Zeitpunkts; er erhält lediglich einen Rückzugs-Vermerk mit Verweis hierher, damit ein
Leser des Reports den Rückzug nicht übersehen kann, sowie eine korrigierte Summary.

**Skill:** `.harness/skills/reviewer.md` @ `20ee992` (2026-07-25) · <!-- d-check:ignore (Adopter-spezifischer Skill-Pfad, existiert im Ziel-Repo ggf. nicht) -->
**Modell:** `claude-opus-5[1m]` · **Datum:** 2026-07-26

**Anlass:** Beim Vorbereiten der Fix-Slices wurde die Slice-Vorlage gelesen, die im
Ursprungs-Report nicht im Eingangs-Kontext stand. Sie widerlegt das Finding.

---

## F-2 wird zurückgezogen — die Praxis ist vorlagenkonform

**Ursprüngliche Behauptung** (Etappe C, F-2, MEDIUM): Die zweite Hälfte von Fund B-2 sei nicht
geliefert; fünf Slices der Kette trügen wortgleich „Alle berührten Sub-Areas GF." ohne die
Sub-Areas zu nennen und ohne eine der vier geforderten Achsen.

**Warum sie nicht trägt.** Die vendored Baseline-Vorlage
`.harness/baseline/v3.5.2/templates/docs/plan/planning/slice.template.md` §8 legt fest:

> **Status:** Pflicht-Sektion bei mindestens einer berührten Sub-Area in BF oder Hybrid. Bei
> reinem GF genügt der Hinweis *„alle berührten Sub-Areas GF (siehe Kurs Modul 5 §Worked
> Mini-Example)"*.

Die vier Pflichtkriterien sind damit an **BF/Hybrid** gebunden. Da alle berührten Sub-Areas dieses
Repos auf Greenfield stehen (belegt in der Modus-Tabelle, geprüft im Ursprungs-Report), ist der
Ein-Satz-Block **die vorgesehene Form** — nicht eine Verkürzung. a-checks eigene Vorlage
(`docs/plan/planning/slice.template.md`, angelegt in slice-052) gibt die Regel korrekt wieder, und
[slice-052 §3](../plan/planning/done/welle-12/slice-052-slice-form.md) hat den Punkt bereits ausdrücklich
aufgeklärt: B-2 habe insoweit zu weit gegriffen.

**Verifikationsweg, der das Finding zunächst stützte und warum er zu kurz war.** Regelwerk
`modul-05` §Ziel-Form Sub-Area-Modus-Begründung formuliert ohne Modus-Bedingung („Pro berührter
Sub-Area vier Pflichtkriterien (vier, nicht erweitern)" … „Ein Block pro berührter Sub-Area") und
führt „Modus (GF/BF/Hybrid)" selbst als Feld — daraus schien zu folgen, dass der Block auch für GF
gilt. Das Modul ist jedoch ein **didaktik-freier Extrakt ohne eigene Normativität**
([`AGENTS.md`](../../AGENTS.md) §1); die Ziel-Form liefert die Template-Datei, und sie ist
eindeutig. Der Extrakt ist an dieser Stelle knapper als die Vorlage, auf die er verweist.

**Ein zwischenzeitlich erwogenes HIGH ist damit ebenfalls gefallen:** die Vermutung, slice-052 habe
eine Baseline-Forderung ohne `MR-*`-Deklaration abgeschwächt (Hard Rule
[`AGENTS.md`](../../AGENTS.md) §3.6, „jede Schwellen-Senkung … ist ein ADR"). Es gibt keine
Abschwächung — die Vorlage ist getreu übersetzt. Eine `MR-*`-Deklaration ist folglich **nicht**
erforderlich.

## Auswirkung auf die Bilanz der Review-Serie

| | vorher | jetzt |
|---|---|---|
| Etappe C | 2 MEDIUM | **1 MEDIUM** (nur noch F-1, B-17 entfällt stillschweigend) |
| Serie gesamt | 4 HIGH · 10 MEDIUM · 8 LOW · 1 INFO (23) | **4 HIGH · 9 MEDIUM · 8 LOW · 1 INFO (22)** |

Für die Fix-Planung entfällt der Punkt „zweite Hälfte von B-2 liefern oder begründet streichen"
ersatzlos.

## Negativbefunde

- geprüft, ohne Befund: **a-checks Vorlage gegen die vendored Vorlage** — die GF-Ausnahme ist in
  beiden inhaltsgleich; die Übersetzung ändert nur die Verweisziele (vendored Pfade statt
  Kurs-URLs), wie in slice-052 §2 angekündigt.
- geprüft, ohne Befund: **keine undeklarierte Adaption** — `harness/conventions.md` enthält
  folgerichtig keinen `MR`-Eintrag zum Sub-Area-Block; es gibt nichts zu deklarieren.

## Lehre für den Reviewer-Skill

Der Eingangs-Kontext des Skills nennt Diff, Lastenheft, ADRs, Hard Rules und frühere Findings —
**nicht** die Vorlage, gegen die ein Form-Befund zu prüfen wäre. Genau daran ist dieses Finding
entstanden: geprüft wurde gegen die Beschreibung eines Funds (slice-048 B-2) statt gegen die
normative Ziel-Form. Kandidat für die Skill-Pflege (Modul 10 §Pflege): *bei Form-Befunden gehört
die einschlägige Vorlage in den Kontext-Eingang, und der Extrakt gilt nie gegen sie.*

## Verdikt

**F-2 der Etappe C ist zurückgezogen.** Der Ursprungs-Report bleibt im Übrigen unverändert gültig;
sein Verdikt („merge-blockierend: nein") ändert sich nicht.
