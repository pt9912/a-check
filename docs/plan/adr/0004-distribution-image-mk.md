# ADR-0004 — Distribution: digest-gepinntes GHCR-Image + `--print-mk`/`a-check.mk`

- **Status:** Accepted
- **Datum:** 2026-06-21
- **Bezug:** [AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk), [AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
- **Schärft:** [SPEC-DIST-001](../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution) — macht die statische/distroless Laufzeitform, den digest-Pin und `--print-mk`/`a-check.mk` verbindlich.
- **Supersedes:** —

## Kontext

Konsumenten-Repos sollen `a-check` als `make a-check`-Gate einbinden — **ohne
Skript-Kopie** (Lastenheft §1,
[AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)).
Das Schwester-Tool [`d-check`](https://github.com/pt9912/d-check) liefert den
include-baren, digest-gepinnten Teil dieses Musters bereits real — dieses Repo
bindet [`d-check.mk`](../../../d-check.mk) ein (Dogfooding). Die
`--print-mk`-**Erzeugung** ist dagegen geteiltes Zielbild:
[`d-check.mk`](../../../d-check.mk) hält selbst fest, dass d-check sie noch
nicht hat und das Fragment bis dahin von Hand gepflegt wird. `a-check`
übernimmt das Distributionsmodell und liefert `--print-mk` zuerst.

## Optionen

1. **GHCR-Image (distroless/static, digest-gepinnt) + `a-check --print-mk`**,
   das ein includebares `a-check.mk` mit dem aktuell gepinnten Image und einem
   `a-check`-Target ausgibt. Trade-off: Docker-Abhängigkeit beim Konsumenten —
   das ist aber Stack-Standard (hermetisch/netzlos,
   [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze))
   und exakt `d-check`s Muster.
2. **Skript-Kopie** (`arch-check.sh` je Repo). Das ist der Status quo, den
   `a-check` ablöst (Lastenheft §1) — die Drift-Quelle selbst. Verworfen.
3. **Binary-Releases** (kein Docker). Trade-off: keine hermetische/distroless
   Hülle, Host-Abhängigkeiten. Im Lastenheft als Out-of-Scope 0.1.0
   ausgewiesen
   ([AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)).

## Entscheidung

**Option 1.** Ein Image, ein Update-Pfad: Konsumenten `include a-check.mk`
(von `--print-mk` erzeugt, digest-gepinnt) und liefern `.a-check.yml` (siehe
[ADR-0003](0003-config-modell-a-check-yml.md)). Die Pin-Hebung ist ein
bewusster Commit
([AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).
Interim folgt das der include-baren [`d-check.mk`](../../../d-check.mk)-Mechanik
dieses Repos (Dogfooding), bis `a-check` sich selbst baut und `--print-mk`
realisiert (slice-003).

## Konsequenzen

- `--print-config` schreibt nichts
  ([AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
  Boundary); `--print-mk` ist nach demselben Prinzip read-only (Ausgabe auf
  stdout) — konsistente Design-Folge über die `--print-config`-Boundary-AK
  hinaus.
- **Trade-off des Pin-Modells:** die Pin-Hebung ist *manuell pro Konsument*
  (kein Auto-Update); ein veralteter Pin bleibt bis zum bewussten Commit
  stehen. Das ist die akzeptierte Kehrseite der Reproduzierbarkeit
  ([AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)) — ein
  zentral via `--print-mk` verteiltes `a-check.mk` hält den Hebungs-Aufwand
  klein (eine Quelle statt N Skript-Kopien).
- Unbekanntes Flag ⇒ Exit-Code 2
  ([AC-FA-CLI-001](../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)).
- **Fitness Function / Gate** (slice-003): Test, dass der `--print-mk`-Output
  ein digest-gepinntes `A_CHECK_IMAGE` plus `a-check`-Target enthält;
  Reproduzierbarkeits-Pin
  ([AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).
- `a-check.mk` selbst entsteht mit slice-003 ([`AGENTS.md`](../../../AGENTS.md)
  §4 Gates-Tabelle); diese ADR legt das *Distributionsmodell* fest, nicht die
  Implementierung.
