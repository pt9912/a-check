# slice-028 — d-check-Matrix-Konvergenz: intra-Spec-Richtung + ADR→Slice-Disziplin

**Status:** done (2026-07-04). **Typ:** Harness-Konvergenz (Doku-Gate), nicht konsumenten-gated.
**Bezug:** [`MR-005`](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung),
schärft [`AGENTS.md` §3.4](../../../../AGENTS.md#34-architektur-sprach-meilensteinfrei-spec-straten-nie-abwärts) maschinell.
[Roadmap](../in-progress/roadmap.md).

## 1. Auslöser

a-checks `matrix`-Modul in [`.d-check.yml`](../../../../.d-check.yml) war die **ältere**
Variante (nur Cross-Klassen `spec-straten → adr`/`spec-straten → slice`). Das
Schwester-Repo `d-check` (Konventions-Vorbild) hat die Referenzmatrix seither
ausgebaut (`DC-FA-MTX-001/002/003`). Konkreter Anlass: ein neu geschriebener ADR
([ADR-0020](../../adr/0020-mehr-wurzel-phantom-guard.md)) zitierte Slices als
Entscheidungs-**Beleg** — das Anti-Muster, das die ausgebaute Matrix maschinell fängt;
a-checks Gate sah es nicht.

## 2. Umfang (umgesetzt)

1. **`.d-check.yml`-`matrix`** um zwei Prüfungen erweitert (Details/Begründung in
   [`MR-005`](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung)):
   - intra-Spec `order` + `direction: no-downward`
     (`lastenheft → spezifikation → architecture`, auch transitiv);
   - `adr → slice`-Regel + `token: 'slice-\d{3}'` — Slice-Kennung im ADR-Körper ist
     `matrix-forbidden`, außer per Marker `<!-- d-check:status-provenance -->`.
2. **Grandfathering** per `exempt-paths` für die immutablen `Accepted`-ADRs (0001–0020);
   neue ADRs ab 0021 tragen den Provenance-Marker.
3. [`MR-005`](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung)
   in `conventions.md` deklariert die Adaption.

## 3. Verifikation

- **Machbarkeit:** a-checks gepinntes `d-check` (v0.35.0) akzeptiert die Schlüssel
  (`order`/`direction`/`token`/`exempt-paths`) und die `adr → slice`-Regel — probeweise
  gegen den Bestand: **16 `matrix-forbidden`** über 10 Bestands-ADRs (0001–0018, 0020),
  alle vom Grandfathering abgedeckt.
- **`make doc-check` grün** (0 Befunde) mit dem Grandfathering; ohne Grandfathering feuert
  die Regel nachweislich (Fitness Function). `make gates` grün.
- Pin-Angleichung an den aktuellen `d-check`-Stand (`v0.35.0 → v0.37.1`) bleibt
  [slice-019](../done/slice-019-dcheck-mk-print-mk-angleichung.md).

## 4. Closure-Notiz

Doku-Gate verschärft (keine Vertrags-/Code-Änderung). Ab jetzt: Spec-Straten dürfen nicht
abwärts (auch intra-Spec), und ADRs argumentieren aus Spec/Verhalten — Slice-Zeiger nur als
deklarierte Provenance. **Lerneintrag:** die Konvergenz mit `d-check` läuft asymmetrisch —
das Vorbild-Repo eilt voraus, a-check zieht seine Harness-Gates nach, sobald ein realer
Anlass (hier: ein ADR mit Slice-als-Beleg) die Lücke sichtbar macht.
