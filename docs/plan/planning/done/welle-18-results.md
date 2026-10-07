# Welle welle-18 — Multi-Arch-Image (`linux/amd64` + `linux/arm64`) — Closure-Notiz

> **Zitier-Form** *(bleibt stehen — Norm, kein Ausfüll-Hinweis).* Dieses
> Artefakt friert ein; was es zitiert, bewegt sich weiter. Deshalb: **Kennung,
> nicht Adresse** — `slice-<Kennung>` statt seines Lifecycle-Pfads, `make <target>`
> statt eines Links auf die Sensor-Datei, eine Baseline-Stelle als
> `v6.13.0` · `regelwerk/<datei>.md` §<Abschnitt> statt als Link.

**Welle:** welle-18-multi-arch-image
**Abschluss:** 2026-10-07
**Verantwortlich:** Claude (Planner); Abnahme beim Maintainer.

## Was wurde geliefert?

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3.

- **Das Release-Image läuft nativ auf `linux/amd64` und `linux/arm64`** — mit `v0.23.0`, ein Pin
  für beide Plattformen: Index-Digest
  `sha256:97cb6d4eb52a0c9fb8f352baeea4f028691fdffbe534499141668dd9329c3f44`, auf GHCR und Docker
  Hub derselbe.
- slice-216: Vertrag — Messung M1–M7, Lastenheft 0.30.0
  ([AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
  [AC-FA-DIST-002](../../../../spec/lastenheft.md#ac-fa-dist-002)), Spezifikation 0.39.0,
  [ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md) `Accepted` (löst die
  Gleichheits-Größe von [ADR-0039](../../adr/0039-spiegel-gleichheit-ist-der-config-digest.md)
  ab).
- slice-217: Cross-Compile im Dockerfile, `make image-multiarch` (Index als OCI-Archiv, geprüft
  von `tools/multiarch-check.sh`), `make image-test IMAGE_REF=… VERSION=…`.
- slice-218: Release-Pipeline — ein Bau, Upload ohne Tag, nativer Test je Plattform,
  Scan-Vergleich, Bild-Tag erst danach, Spiegel als Index-Kopie; `releasing.md`, Handbuch 1.46,
  Hub-Seite.

## Was hat funktioniert?

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3.

- **Messen vor dem Entscheiden.** Der lokale Versuch mit zwei Registries zeigte, dass der
  bisherige Spiegel-Weg arm64 verloren hätte — und dass eine Index-Kopie den Digest behält. Beides
  hat den Vertrag geformt und hielt am Tag: GHCR und Docker Hub tragen denselben Index-Digest.
- **Die Pipeline lief im ersten echten Lauf durch**, alle fünf Jobs grün, einschließlich des
  nativen arm64-Tests — die Schritte waren vorher einzeln gegen lokale Registries geprobt.
- **Gegenproben als Archive, die genau eine Eigenschaft brechen** — die teuerste (ein Bau ohne
  Cross-Compile, amd64-Binary im arm64-Bild) war rot an der Stelle, die den Fehler benennt.

## Was ging anders als geplant?

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3.

- **Sieben Review-Läufe** (je Slice ein Review und ein bis zwei Delta-Reviews; gezählt an den
  Reports: drei Erst-Reviews, vier Delta-Abschnitte). Die wiederkehrende
  Klasse: Zusagen in Vertrag, Kommentar oder Doku, die der Prüfer nicht hielt — Reproduzierbarkeit
  (widerlegt durch eine Gegenmessung), Plattform-Vergleich ohne Träger, ein Prüfer, der in zwei
  Fassungen mehr versprach, als er prüfte, ein Kopfkommentar über die Pipeline. Fünf Vorgänge
  insgesamt im Register; Ausgang *geplant* (slice-219).
- **Umbauten verloren Prüfungen still** — die Index-Beschriftung in `multiarch-check`, die fünf
  OCI-Labels in der Pipeline. Beide fing der Review, beide sind zurück; die Klasse steht neu im
  Register.
- **Eine eigene Messung reichte weiter als ihr Geltungsbereich** (M4, Reproduzierbarkeit): vierter
  Beleg nach der Verkörperung als Mess-Regel 1.

## Steering-Loop-Einträge

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3 — hier stehen nur Beobachtungen, die im Register 3× erreicht
haben.

- **`BEO-GATE/zusage-weiter-als-ihre-durchsetzung`** (5×) — Ausgang *geplant*: slice-219
  verkörpert die Schreibregel „eine Zusage über eine Prüfung nennt den Lauf, der sie hält, und
  seine Grenze am Satz". Kein Sensor: ob ein Satz den Geltungsbereich seines Prüfers trifft, ist
  ein Urteil über zwei Formulierungen.
- **`BEO-PLAN/review-geltungsbereich-zu-eng`** (4×, verkörpert als Mess-Regel 1) — vierter Beleg
  nach der Verkörperung: die Prosa-Form gilt als ausgeschöpft; **kein Sensor möglich** (Urteil über
  die Reichweite einer Folgerung), die Gegenmessung im unabhängigen Review bleibt die Antwort.
  Begründung im `state.md` des Eintrags.

## Beobachtungs-Register (Zeiger)

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Das Beobachtungs-Register — der Zähler wird nicht hier gepflegt.

Der Zähler steht im [Beobachtungs-Register](../observations/README.md). Neu oder fortgeschrieben
in dieser Welle: `BEO-GATE/zusage-weiter-als-ihre-durchsetzung` (5×, geplant) ·
`BEO-PLAN/review-geltungsbereich-zu-eng` (4×, verkörpert) ·
`BEO-GATE/pipefail-bricht-pruefer-stumm-ab` (1×) · `BEO-GATE/umbau-verliert-pruefung-still` (2×).

**Lese-Schritt:** außer den zwei Einträgen oben steht kein offener Eintrag bei 3× oder mehr.

**Bestand in `open/`, Lese-Schritt:** slice-013 und slice-045 haben diese Closure unverändert
überstanden; keine gemeinsame Ursache, je **bestätigt** — beide trigger-gebunden, die Trigger sind
nicht gefeuert (Nachmessung 2026-10-07: kein Port→Port-Import in pgwire-recorder, keine
Konsumenten-Anfrage). slice-219 entstand in dieser Closure.

## Folge-Slices

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3 — derivativ.

- slice-219 (Regel „Zusage nennt Lauf und Grenze", wellenlos) — in `open/`.

## Verifikation

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 1 — Replay-Ersatz nach
[MR-028](../../../../harness/conventions.md#mr-028).

- Alle drei Slices in `done/`; `make ci` Exit 0 und `make verify` Exit 0 auf dem Stand nach dem
  Re-Pin (lokal, Ausgabe in Dateien, Exit-Codes getrennt gelesen); `make preflight` Exit 0 vor dem
  Push (41 Commits Range).
- CI auf `main` grün: Lauf 37629916361. Release-Pipeline grün: Lauf 37630094148 (Tag `v0.23.0`) —
  build, test (linux/amd64), test (linux/arm64), publish, Hub-Darstellung.
- **Belege aus dem Lauf:** `multiarch-check: ok — … mit genau: linux/amd64 linux/arm64`;
  „hochgeladen ohne Tag … (gleich dem geprüften Archiv)"; Test-Job arm64: „Binary passt zur
  Host-Plattform aarch64", „Versions-Label 0.23.0"; Test-Job amd64: dasselbe mit „x86_64"; beide
  Scan-Hashes `b8ae01d2…` — „Scan-Ausgabe auf beiden Plattformen gleich"; „Spiegel geprueft:
  Index-Digest identisch".
- **Gegenprobe am veröffentlichten Image (eigener Lauf):** `docker buildx imagetools inspect`
  nennt für `ghcr.io/pt9912/a-check:v0.23.0`, `:latest` und `docker.io/pt9912/a-check:v0.23.0`
  denselben Digest `97cb6d4e…` mit genau `linux/amd64` und `linux/arm64`; der gezogene Digest
  läuft lokal (`amd64`, Label `0.23.0`). Ein arm64-Lauf ist auf diesem Host nicht möglich; den trägt
  der native Test-Job. Re-Pin danach mit `make gates` Exit 0 (`gate-consistency`: Pins konsistent).
- `make doc-immutable RANGE=v0.22.0..HEAD`: 0 Befunde.
- Coverage gesamt 96,1 % (Schwelle 90 %). Carveouts: keine.
- **Trigger-Audit, vier Klassen:** Carveout 0 offen (Verzeichnis führt nur die README) ·
  bootstrap-aware Gate 0 (keines im Bestand) · ADR 0 fällig (die Trigger von
  [ADR-0043](../../adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md) — Runner nicht
  verfügbar, weitere Plattform, Index-Digest nicht erhalten, Attestierungen gefordert — sind
  nicht eingetreten; der dritte ist mit diesem Release widerlegt) · Hard Rule 0 (keine mit
  Auflösungs-Trigger) — dazu die sieben aktiven MR-Einträge: 0 offen.
- **Drei Paarungen:** Anker — die Steering-Loop-Einträge der Slices (`Makefile:image-multiarch`,
  `.github/workflows/release.yml`) existieren; `seit slice-179` steht in der Regel-Datei ·
  Folge-Slice — slice-219 existiert in `open/` · Register — alle vier genannten Pfade existieren
  als Verzeichnis mit nicht leerem `evidence/` (`make verify-observations` im `verify`-Lauf grün).
