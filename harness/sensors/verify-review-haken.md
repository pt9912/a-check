# `make verify-review-haken` — Review-DoD-Haken an Report-Existenz binden

## Vertrag

Jeder **abgehakte** DoD-Punkt, dessen Zeile die Phrase „Unabhängiger Review"
trägt, braucht in einem `in-progress/`-Slice einen Report unter
`docs/reviews/`, dessen **Dateiname** die Slice-Kennung trägt
(`slice-<NNN>`; Führende Nullen und die Formen `slice205`/`slice_205` werden
erkannt, `slice-2050` bedient nicht `slice-205`). Die Phrase ist dieselbe
exakte Wortform, die das Modul `reviews` als Opt-in erwartet
([`MR-019`](../conventions/MR-019-review-dod-opt-in.md)) und die
structure (1) als konstanten Posten auszählt.

## Grenze — was das Grün nicht abdeckt

1. **Der Dateinamen-Match ist Form** — ob der Report zum Lauf passt und
   trägt, ist Urteil (Modul 10). Ein Report eines früheren Laufs desselben
   Slice genügt dem Sensor; Bereichs-Namen (`slice135-157`) bedienen nur die
   Endpunkte. Die Phrase wird in Großform geprüft (die deployte Form).
2. **Nur abgehakte Zeilen sind eine Zusage** (Opt-in pro Slice,
   [`MR-019`](../conventions/MR-019-review-dod-opt-in.md)) — ein
   unabgehakter DoD-Punkt läuft grün durch; dafür deckt die
   structure-Bedingung 7 die unchecked-Hälfte in `open/`. `next/` ist durch
   keinen der beiden Wächter gedeckt — deklarierte Lücke (slice-202).
3. **done/ wird nicht gescannt** — dort urteilt das Modul `reviews`
   (`doc-reviews`, done/-Geltungsbereich); die beiden Wächter überlappen
   nicht. Permanent: das Modul scannt genau ein `done-dir` (d-check v0.79.0),
   eine Erweiterung wäre ein CR an das Fremdwerkzeug.

## Bindung

Harness-Prozess ([`AGENTS.md`](../../AGENTS.md) §5) · Antwort auf
[`BEO-GATE/attestierung-vor-dem-vorgang`](../../docs/plan/planning/observations/BEO-GATE/attestierung-vor-dem-vorgang/observation.md)
bei 5× · slice-204 · im `verify`-Aggregat.
