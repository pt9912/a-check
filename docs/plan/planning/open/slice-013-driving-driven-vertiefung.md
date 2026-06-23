# slice-013 — Driving/Driven-Vertiefung (Backlog, Trigger-getrieben)

**Status:** open (Backlog — feuert auf Konsumenten-Bedarf, **nicht** terminiert).
**Bezug:** Carry-forward aus [slice-012 §7](../done/slice-012-driving-driven-layerof.md)
(welle-10b/b2b); löst die in [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md)
bewusst *out-of-scope* gestellten Richtungs-Inkremente ein; verfeinert
[AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--driving-driven-port-richtung-regel-port-direction-mismatch).
[Roadmap welle-10](../in-progress/roadmap.md).

> **Backlog-Stub.** Kein Entwurf zur Abnahme — nur der **getrackte** Carry-forward, damit
> die offene Arbeit nicht in Prosa verschwindet. Wird zum vollen Slice ausgearbeitet,
> sobald der Trigger (§1) feuert.

## 1. Auslöser (Gate)

slice-012 lieferte `direction` (`driving`/`driven`) + die Regel `port-direction-mismatch`
**inert** — ohne `direction` ändert sich nichts. Bevor weitere Richtungs-Arbeit Sinn hat,
muss der Bedarf real sein:

- **Bedarfs-Gate:** mindestens ein Konsument (b-cad/d-check/d-migrate) aktiviert getrennte
  `driving`/`driven`-**Adapter- und -Port**-Schichten in seiner `.a-check.yml`. Ohne diese
  Aktivierung bleibt dieser Slice schlafend.

## 2. Geplanter Umfang (a-check-seitig)

Beides ist in [ADR-0012](../../adr/0012-driving-driven-richtung-orthogonale-dimension.md)
als out-of-scope vermerkt:

1. **Port→Port-Richtungsregeln** — Richtungs-Abgleich nicht nur `adapter→port`, sondern
   auch zwischen Ports untereinander (z. B. darf ein `driving`-Port keinen `driven`-Port
   direkt sprechen). Verhältnis zu `wrong-direction`/`edges` zu klären.
2. **Auto-Inferenz der Richtung** aus Pfad/Namen (`driving`/`driven` im Pfad) statt
   expliziter Deklaration — abzuwägen gegen die „explizit statt geraten"-Linie von
   slice-012.

## 3. Konsumenten-Pilot (a-check-fremd)

Die eigentliche **Aktivierung** geschieht in den Konsumenten-Repos (deren `.a-check.yml` +
Schicht-Schnitt), nicht in a-check. Sie überlappt mit dem offenen Meilenstein **M3
„Pilot-Einbindung"** (welle-05). Dieser Slice trägt nur den a-check-seitigen Anteil; der
Pilot ist der Trigger.

## 4. Vor der Umsetzung zu klären

- Bedarf bestätigt? (Gate, §1) — sonst nicht ziehen.
- Port→Port: kanten-basiert (über `edges`/`allow`) oder kategorisch wie
  `port-direction-mismatch`?
- Auto-Inferenz: überhaupt gewollt, oder widerspricht sie der expliziten Deklaration
  (Determinismus/Überraschungsfreiheit)?
