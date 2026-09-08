# MR-021 — Verfeinerungen tragen `SPEC-*` statt der Suffix-Form (löst [`MR-011`](../conventions.md#mr-011) ab, gemessene Begründung)

- **Status:** Accepted
- **Datum:** 2026-09-08
- **Geltungsbereich:** [`spec/spezifikation.md`](../../spec/spezifikation.md)
- **Ersetzt-Baseline-Regel:** [`grundlagen-source-precedence.md` §ID-Schema als Klammer](../../.harness/baseline/v6.5.0/regelwerk/grundlagen-source-precedence.md#id-schema-als-klammer)
- **Adaption:** unverändert gegenüber [`MR-011`](../conventions.md#mr-011). Die Baseline sieht für
  die **Verfeinerung** genau einer Anforderung im Technik-Stratum die Suffix-Form
  `<PREFIX>-FA-<NN>.<Buchstabe>` vor. a-check nutzt sie nicht: jede technische Festlegung — auch
  die, die eine Anforderung verfeinert — trägt eine eigene `SPEC-<BEREICH>-<NNN>`.
- **Begründung — zwei tragende Gründe, beide gemessen (2026-09-08):**

  1. **Die `SPEC-*`-Anker werden aus `Accepted`-ADRs referenziert, und die sind immutabel**
     ([`AGENTS.md`](../../AGENTS.md) §3.5). Eine Umbenennung erzeugte denselben unauflösbaren
     Widerspruch zwischen `make doc-check` (Anker müssen auflösen) und `make doc-immutable` (ADRs
     sind unantastbar), den der
     [Anforderungs-Anlege-Prozess](../conventions.md#anforderungs-anlege-prozess) für
     `AC-*`-Überschriften bereits beschreibt. Gemessen: **40** ADR-Dateien tragen ein
     `Schärft:`-Feld, das auf solche Anker zeigt.
  2. **Ein Umbau träfe alle bestehenden `SPEC-*` samt ihrer Traceability-Verweise, ohne eine
     Aussage zu ändern.** Gemessen: **sieben** `SPEC-*`-Abschnitte in
     [`spec/spezifikation.md`](../../spec/spezifikation.md).

- **Was dieser Eintrag gegenüber [`MR-011`](../conventions.md#mr-011) korrigiert:** Jener begründete
  zusätzlich damit, *„die Verfeinerungs-Beziehung steht in a-check ohnehin explizit im
  `Schärft:`-Feld und nicht in der Kennung"*. Das trifft nicht zu und ist der Grund für diese
  Ablösung. Gemessen: **null** der sieben `SPEC-*`-Abschnitte trägt ein `Schärft:`-Feld; die
  Verfeinerung steht dort als Prosasatz *„Präzisiert [`AC-…`]"*. Das Feld lebt in der
  **Gegenrichtung** — in **40** ADR-Dateien, die auf `SPEC-*` zeigen, und einmal im Lastenheft.
  **Die scheinbare Gegenanzeige stützt den Befund:** Das Wort kommt in
  [`spec/spezifikation.md`](../../spec/spezifikation.md) genau **einmal** vor, im Kopf der
  Historie-Sektion — und dort steht ausdrücklich, *„welche ADR eine Festlegung schärft, deklariert
  **die ADR** aufwärts in ihrem `Schärft:`-Feld"*. Die Spezifikation sagt selbst, dass das Feld
  nicht ihres ist.
  **Die Entscheidung bleibt unberührt:** Nicht migrieren war und ist richtig; falsch war allein
  der dritte Beleg.
- **Auflösungs-Trigger:** sobald ein Gate die Verfeinerungs-Beziehung aus der **Kennung** ableiten
  soll — dann ist die Suffix-Form billiger als ein Feld. Unverändert gegenüber
  [`MR-011`](../conventions.md#mr-011).
- **Löst ab:** [`MR-011`](../conventions.md#mr-011) (das seinerseits
  [`MR-004`](../conventions.md#mr-004) auflöste)
- **Ausgelöst durch Baseline-Stand:** `v6.5.0`
