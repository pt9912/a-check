# MR-022 — Verfeinerungen tragen `SPEC-*` statt der Suffix-Form (löst [`MR-021`](../conventions.md#mr-021) ab, Zahl korrigiert)

- **Status:** Accepted
- **Datum:** 2026-09-08
- **Geltungsbereich:** [`spec/spezifikation.md`](../../spec/spezifikation.md)
- **Ersetzt-Baseline-Regel:** [`grundlagen-source-precedence.md` §ID-Schema als Klammer](../../.harness/baseline/v6.6.0/regelwerk/grundlagen-source-precedence.md#id-schema-als-klammer)
- **Adaption:** unverändert seit [`MR-011`](../conventions/done/MR-011-verfeinerungs-form.md). Die
  Baseline sieht für die **Verfeinerung** genau einer Anforderung im Technik-Stratum die
  Suffix-Form `<PREFIX>-FA-<NN>.<Buchstabe>` vor. a-check nutzt sie nicht: jede technische
  Festlegung — auch die, die eine Anforderung verfeinert — trägt eine eigene
  `SPEC-<BEREICH>-<NNN>`.
- **Begründung — zwei tragende Gründe, beide gemessen (2026-09-08):**

  1. **Die `SPEC-*`-Anker werden aus `Accepted`-ADRs referenziert, und die sind immutabel**
     ([`AGENTS.md`](../../AGENTS.md) §3.5). Eine Umbenennung erzeugte denselben unauflösbaren
     Widerspruch zwischen `make doc-check` (Anker müssen auflösen) und `make doc-immutable` (ADRs
     sind unantastbar), den der
     [Anforderungs-Anlege-Prozess](../conventions.md#anforderungs-anlege-prozess) für
     `AC-*`-Überschriften bereits beschreibt. **Verschärfend:**
     [`spec/spezifikation.md`](../../spec/spezifikation.md) führt **null** explizite
     `<a id="…">`-Anker; die ADR-Verweise hängen an generierten Slugs und brechen damit schon bei
     einer Umbenennung der Überschrift.
  2. **Ein Umbau träfe alle bestehenden `SPEC-*` samt ihrer Traceability-Verweise, ohne eine
     Aussage zu ändern.** Es sind **sieben** `SPEC-*`-Abschnitte.

  **Zählregel und Geltungsbereich** ([`AGENTS.md`](../../AGENTS.md) §5, `seit slice-171` in
  dieser Form): Gezählt sind Dateien
  `docs/plan/adr/0*.md` (der Index `README.md` ist kein ADR und fällt heraus), deren
  `**Schärft:**`-Feldzeile mindestens einen `SPEC-`-Anker nennt — **34**, davon **30** mit Status
  `Accepted` und vier `Superseded`. Nicht gezählt: Prosa-Erwähnungen des Worts außerhalb der
  Feldzeile, und `SPEC-`-Nennungen an anderer Stelle derselben ADR.

- **Was dieser Eintrag gegenüber [`MR-021`](../conventions.md#mr-021) korrigiert:** Jener nannte
  *„**40** ADR-Dateien tragen ein `Schärft:`-Feld, das auf solche Anker zeigt"*. Die Zahl misst
  etwas anderes als sie aussagt: **40** ist die Zahl der Dateien in `docs/plan/adr/`, die den
  *String* `Schärft:` irgendwo enthalten — der Index eingeschlossen, der nur zwei
  Prosa-Erwähnungen trägt. Über *Felder, die auf einen Anker zeigen* sagt sie nichts.
  **Das ist derselbe Fehler, den [`MR-021`](../conventions.md#mr-021) an
  [`MR-011`](../conventions/done/MR-011-verfeinerungs-form.md) behob** — ein Beleg, der plausibel
  klingt und die gestellte Frage nicht misst.
- **Was dieser Eintrag von [`MR-021`](../conventions.md#mr-021) übernimmt:** dessen Korrektur an
  [`MR-011`](../conventions/done/MR-011-verfeinerungs-form.md) — die Begründung *„die
  Verfeinerungs-Beziehung steht ohnehin explizit im `Schärft:`-Feld"* trifft nicht zu: **null** der
  sieben `SPEC-*`-Abschnitte trägt ein solches Feld, die Verfeinerung steht dort als Prosasatz
  *„Präzisiert [`AC-…`]"*. Das Wort kommt in
  [`spec/spezifikation.md`](../../spec/spezifikation.md) genau **einmal** vor, im Kopf der
  Historie-Sektion, und dort steht ausdrücklich, *„welche ADR eine Festlegung schärft, deklariert
  **die ADR** aufwärts in ihrem `Schärft:`-Feld"*.
  **Die Entscheidung bleibt in allen drei Einträgen dieselbe:** nicht migrieren.
- **Auflösungs-Trigger:** sobald ein Gate die Verfeinerungs-Beziehung aus der **Kennung** ableiten
  soll — dann ist die Suffix-Form billiger als ein Feld. Unverändert seit
  [`MR-011`](../conventions/done/MR-011-verfeinerungs-form.md).
- **Löst ab:** [`MR-021`](../conventions.md#mr-021)
- **Ausgelöst durch Baseline-Stand:** `v6.6.0`
