# MR-027 — Verfeinerungen tragen `SPEC-*` statt der Suffix-Form, Zeiger auf `v6.13.0` (löst [`MR-022`](../conventions.md#mr-022) ab)

- **Status:** Accepted
- **Datum:** 2026-09-29
- **Geltungsbereich:** [`spec/spezifikation.md`](../../spec/spezifikation.md)
- **Ersetzt-Baseline-Regel:** [`grundlagen-source-precedence.md` §ID-Schema als Klammer](../../.harness/baseline/v6.13.0/regelwerk/grundlagen-source-precedence.md#id-schema-als-klammer)
- **Adaption:** unverändert seit [`MR-022`](done/MR-022-verfeinerungs-form.md), das die
  Adaption seit [`MR-011`](../conventions/done/MR-011-verfeinerungs-form.md) unverändert führt. Die
  Baseline sieht für die **Verfeinerung** genau einer Anforderung im Technik-Stratum die
  Suffix-Form `<PREFIX>-FA-<NN>.<Buchstabe>` vor. a-check nutzt sie nicht: jede technische
  Festlegung — auch die, die eine Anforderung verfeinert — trägt eine eigene
  `SPEC-<BEREICH>-<NNN>`.
- **Begründung:** dieselbe wie bei [`MR-022`](done/MR-022-verfeinerungs-form.md), beide
  Gründe gemessen und unverändert. Dieser Eintrag existiert, weil der Zielabschnitt im Sprung
  `v6.6.0` → `v6.13.0` **nicht wortgleich** blieb — der Zeiger wandert darum nicht still,
  sondern über diesen Nachfolger.
- **Gemessene Abweichung des Zielabschnitts** (`v6.6.0` → `v6.13.0`, 10 942 → 6 574 Zeichen, ohne
  die Quelle-Zeile): der Abschnitt ist stark gekürzt und umgebaut — die Vergabe-Passage
  (Nummer → Kennung) neu geschrieben, der Unterabschnitt heißt seitdem *„Vergabe: woher die
  nächste Kennung kommt"* (zuvor *… nächste Nummer kommt*), und die Kennungs-Tabelle trägt die
  neue `RB`-Reihe (Welle 144). **Die Suffix-Passage (Zeile 17 beider Fassungen: „Suffix:
  `LH-FA-03` ist Vertrag, `LH-FA-03.a` ist dessen Verfeinerung") ist wortgleich geblieben** —
  der Ausgang ist *bleibt gültig*.
- **Auflösungs-Trigger:** sobald ein Gate die Verfeinerungs-Beziehung aus der **Kennung** ableiten
  soll — dann ist die Suffix-Form billiger als ein Feld. Unverändert seit
  [`MR-011`](../conventions/done/MR-011-verfeinerungs-form.md).
- **Löst ab:** [`MR-022`](../conventions.md#mr-022)
- **Ausgelöst durch Baseline-Stand:** `v6.13.0`
