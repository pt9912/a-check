# ADR-0011 — Domain/Application-Trennung: Rolle `app` + strenge `domain`

- **Status:** Accepted
- **Datum:** 2026-06-22
- **Autor:** pt9912
- **Bezug:** [AC-FA-RULE-007](../../../spec/lastenheft.md#ac-fa-rule-007--rolle-app-und-strenge-domain) (neu; Lastenheft 0.4.0→0.5.0), [AC-FA-RULE-006](../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung) (Rollen-Mechanismus), [AC-FA-RULE-001](../../../spec/lastenheft.md#ac-fa-rule-001--kern-reinheit-regel-core-impurity)
- **Schärft:** [SPEC-RULE-001](../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) — macht die Rolle `app` und die strenge `domain`-Reinheit verbindlich.
- **Supersedes:** —

## Kontext

[ADR-0009](0009-rollen-basierter-regel-dispatch.md) (welle-10a) führte Rollen
{`domain`, `port`, `adapter`} ein und ließ die Domain/Application-Trennung als
Re-Evaluierungs-Trigger offen. Zwei Lücken:

1. Keine `app`-Rolle: wer Use-Cases von der puren Domäne trennt, kann die
   Application-Schicht nicht modellieren — sie ist rein gegen Technik, darf aber
   Ports nutzen. Es blieb nur `port` (zu eng) oder rollenlos (gar keine Reinheit).
2. `domain → port` war nur kanten-geregelt — eine deklarierte Kante konnte die
   Trennung „Domäne kennt keine Ports" aushebeln.

## Entscheidung

1. **Neue Rolle `app`** (kein Adapter/Tech, darf `domain`+`port`) → Befund
   `app-impurity`. Namens-Inferenz `application`/`app` → `app`; explizite `role: app`
   möglich. Reinheit kategorisch, Richtung (`app → domain`, `app → port`)
   kanten-geregelt.
2. **`domain` verschärft:** die innerste Schicht importiert nur `domain` (+ stdlib);
   ein Import auf eine `app`-/`port`-/`adapter`-Schicht **oder** ein `tech`-Muster
   ist `core-impurity`, **kategorisch** (Kante hebt nicht auf). „Domäne kennt keine
   Ports" wird harte Invariante statt Kanten-Konvention. Rollenlose Ziel-Schichten
   bleiben kanten-geregelt.

## Konsequenzen

- Vier-Schichten-Hexagon (`domain ← app ← port ← adapter`) voll prüfbar; b-cad/
  d-migrate können ihre `application`-Schicht modellieren.
- **Verhaltensänderung:** `domain → port` / `domain → app` feuern jetzt
  `core-impurity` (vorher höchstens `wrong-direction`, per Kante aufhebbar).
- **Sechster Befund-Name** `app-impurity` — Output-Konsumenten (CI-Parser, Doku)
  nachziehen; bestehende Namen stabil.
- **Rückwärtskompat:** ohne `app`-Layer und ohne `domain → port`-Import bleibt alles
  grün — a-checks Dogfooding unverändert (der Kern importiert keine Ports).
- **Migration:** wer Ports aus einer `role: domain`-Schicht importierte, wird rot —
  Port-Nutzung in eine `role: app`-Schicht heben.
- Lastenheft 0.4.0 → **0.5.0**.

## Fitness Function

- `make arch-check` (Dogfooding, [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)):
  unverändert 0 Befunde.
- `make test`: `app` happy / `app → adapter` / `app → tech` ⇒ `app-impurity`
  (kategorisch); `domain → port` / `domain → app` ⇒ `core-impurity` (kategorisch);
  Namens-Inferenz `application`/`app`; klassische Config unverändert grün.

## Re-Evaluierungs-Trigger

- `driving`/`driven`-Port-Subtypen (welle-10b/b2b).
- `LayerOf` (eigene Schicht einer Datei) bleibt Erst-Treffer — Angleichung an
  `targetLayer`s längster-Präfix bei Bedarf.

## Geschichte

| Datum | Ereignis |
|---|---|
| 2026-06-22 | Proposed — welle-10b (b2a); Rolle `app` + strenge `domain` (Domain/Application-Trennung). |
| 2026-06-22 | Proposed → Accepted (Sign-off Auftraggeber: A = neue Anforderung, B = volle `domain`-Schärfung; Multi-Linsen-Review + Delta bestanden). Ab jetzt immutable; Ablösung nur via Folge-ADR mit `Supersedes`. |
