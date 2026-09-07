# `make regelwerk-check` — Integrität der vendored Baseline (kein Gate)

## Vertrag

**Kein Gate** — das Target urteilt nicht über den Zustand des Repos, es
**misst** die Unversehrtheit von Fremdtext: Alle Dateien unter
`.harness/baseline/<tag>/` stimmen mit dem `SHA256SUMS` desselben Verzeichnisses
überein, und es liegt keine unmanifestierte Datei im Baum. Fail-closed.

Liegen mehrere Stände vendored, prüft der Lauf den **höchsten** und weist die
übrigen **namentlich** als ungeprüft aus — kein stilles Übergehen.

## Grenze — was das Grün nicht abdeckt

1. **Freshness** — ob der adoptierte Stand noch der neueste ist, bleibt
   ausdrücklich ungeprüft: Das wäre eine Netz-Operation. Der Lauf nennt
   stattdessen die Release-Liste und die offene Handlung. Permanent, und der
   Grund, warum das Target in keinem Aggregat hängt.
2. **Ob der vendored Inhalt dem Kurs-Tag entspricht** — das `SHA256SUMS` ist
   beim Vendoring aus dem Baum erzeugt worden; es bezeugt Unversehrtheit *seit
   dem Vendoring*, nicht Herkunft. Die Herkunft belegt der Vendoring-Slice über
   den Release-Digest, einmalig. Permanent.

## Bindung

`kein Gate`; Wartung ·
[`MR-006`](../conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert).
