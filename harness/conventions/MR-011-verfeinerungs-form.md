# MR-011 — Verfeinerungen tragen `SPEC-*` statt der Suffix-Form (schärft [`MR-004`](../conventions.md#mr-004))

- **Status:** Accepted
- **Datum:** 2026-08-29
- **Geltungsbereich:** [`spec/spezifikation.md`](../../spec/spezifikation.md)
- **Ersetzt-Baseline-Regel:** [`grundlagen-source-precedence.md` §ID-Schema als Klammer](../../.harness/baseline/v6.0.0/regelwerk/grundlagen-source-precedence.md#id-schema-als-klammer)
- **Adaption:** Die Baseline sieht für die **Verfeinerung** genau einer Anforderung im
  Technik-Stratum die Suffix-Form `<PREFIX>-FA-<NN>.<Buchstabe>` vor. a-check nutzt sie nicht:
  jede technische Festlegung — auch die, die eine Anforderung verfeinert — trägt eine eigene
  `SPEC-<BEREICH>-<NNN>`.
- **Begründung:** Die Straten-Zuordnung selbst und `SPEC-*`/`ARC-*` als **Struktur**-IDs sind mit
  `v5.12.0` Default geworden; was [`MR-004`](../conventions.md#mr-004) darüber hinaus setzte, ist damit
  gegenstandslos. Übrig bleibt allein die Kennungs-Form der Verfeinerung. Ein Umbau träfe **alle**
  bestehenden `SPEC-*` und damit die Traceability-Verweise im Code, ohne eine Aussage zu ändern:
  die Verfeinerungs-Beziehung steht in a-check ohnehin explizit im `Schärft:`-Feld und nicht in der
  Kennung.
- **Auflösungs-Trigger:** sobald ein Gate die Verfeinerungs-Beziehung aus der **Kennung** ableiten
  soll — dann ist die Suffix-Form billiger als ein Feld.
- **Löst auf:** [`MR-004`](../conventions.md#mr-004)
- **Ausgelöst durch Baseline-Stand:** `v5.12.0`
