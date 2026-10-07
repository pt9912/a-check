# ADR-0043: Multi-Arch-Image — ein Cross-Compile-Bau, je Plattform nativ getestet, dann getaggt; der Spiegel kopiert den Index

**Status:** Accepted

**Datum:** 2026-10-07

**Autor:** pt9912 (Anweisung und Abnahme), ausgeführt im Auftrag

**Bezug:** [AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[AC-FA-DIST-002](../../../spec/lastenheft.md#ac-fa-dist-002),
[AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit);
[ADR-0039](0039-spiegel-gleichheit-ist-der-config-digest.md) — **löst deren Gleichheits-Größe
ab**: Punkt 1, das Wort „Config-Digest" in Punkt 2 und die Stellen in §Konsequenzen und §Fitness
Function, die ihn tragen, lesen sich als **Index-Digest**; dazu **den Pin-Satz in Punkt 4** („wer
vom Spiegel zieht, nimmt den Digest der Registry, aus der er zieht"). Fail-closed (Punkt 2),
Reihenfolge (Punkt 3) und die GHCR-Bindung der Pin-Stellen (übriger Punkt 4) gelten weiter;
ADR-0039 selbst bleibt unverändert `Accepted`.
[ADR-0004](0004-distribution-image-mk.md), [ADR-0007](0007-latest-tag-politik.md).

**Schärft:** [SPEC-DIST-001](../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution)
(Plattformen, Index-Digest als Pin, getestetes = veröffentlichtes Bild, Spiegel).

**Regeln:** Baseline-Regelwerk `modul-04-adrs.md` §Ziel-Form: ADR (MADR).

---

## Kontext

Das Release-Image gibt es nur für `linux/amd64`. Auf macOS mit Apple Silicon läuft es unter
Emulation — langsam, und für einen Konsumenten nicht offensichtlich. Das Lastenheft sagt jetzt zu:
**ein** Pin für `linux/amd64` und `linux/arm64`, und veröffentlicht wird das Bild, das auf jeder
Plattform getestet wurde.

Gemessen (lokal, Host `linux/amd64` ohne Emulator, zwei lokale Registries):

- Die gepinnten Basis-Images (`golang`, `distroless/static`) sind Image-Indizes mit beiden
  Plattformen. a-check ist reines Go ohne cgo: ein Cross-Compile auf dem Bau-Rechner liefert ein
  statisches AArch64-Binary, und die Laufzeit-Stufe führt keinen Befehl aus — **kein Emulator**
  nötig.
- `buildx` legt ungefragt Attestierungs-Manifeste (`unknown/unknown`) in den Index.
- Der Bau ist **nicht** reproduzierbar: das Binary ist bitgleich, aber mtime und `created`
  ändern den Digest. `SOURCE_DATE_EPOCH` mit `rewrite-timestamp=true` glich zwei Läufe an, eine
  Gegenmessung mit einem warmen Cache aus einem Bau mit anderer Epoch ergab trotzdem einen anderen
  Digest — `rewrite-timestamp` setzt nur spätere Zeitstempel zurück.
- Ein unverändert kopierter Index (`imagetools create`, `skopeo copy --all`) behält seinen
  Digest auf der Ziel-Registry. Der heutige Spiegel-Weg `docker tag` + `docker push` schreibt aus
  dem lokalen Bildspeicher ein **Einzel-Manifest** der Host-Plattform — **arm64 ginge verloren**,
  und der Digest wäre ein anderer. Die Begründung von ADR-0039 Punkt 1 („der Manifest-Digest ist
  registry-lokal") gilt für diesen Neu-Push, nicht für die Kopie.

Die heutige Pipeline testet ein lokal gebautes Bild und pusht genau dieses. Mit zwei Plattformen
gibt es kein lokales Bild mehr, das beide trägt; „getestet = veröffentlicht" braucht eine neue
Form.

## Entscheidung

Wir wählen **einen Cross-Compile-Bau beider Plattformen, den Test genau dieses Bilds je
Plattform auf einem nativen Runner, und erst danach den Bild-Tag** — im Einzelnen:

1. **Plattformen:** genau `linux/amd64` und `linux/arm64`.
2. **Cross-Compile:** die Stufe, die kompiliert, läuft auf der Plattform des Bau-Rechners
   (`--platform=$BUILDPLATFORM`) und setzt `GOOS`/`GOARCH` aus `TARGETOS`/`TARGETARCH`; die
   Laufzeit-Stufe bleibt ohne `RUN`. Ein Emulator ist für den Bau nicht nötig.
3. **Der Index trägt nur die zwei Plattformen:** Attestierungen und SBOM sind abgeschaltet
   (`--provenance=false --sbom=false`); sie sind ein eigener Vorgang.
4. **Keine Reproduzierbarkeits-Zusage:** ein neuer Bau desselben Commits darf einen anderen
   Digest ergeben. „Getestet = veröffentlicht" stützt sich allein auf Punkt 5 — es gibt genau
   einen Bau, und dessen Digest wird getestet und getaggt.
5. **Ein Bau, getestet, dann getaggt:** Die Pipeline baut den Index **einmal** und lädt ihn
   **ohne Tag** (nur per Digest) nach GHCR. Je ein Test-Job auf einem nativen `linux/amd64`- und
   einem nativen `linux/arm64`-Runner zieht **diesen Digest** und fährt den Image-Test samt Versions-Label; beide legen die Ausgabe eines
   festen Scans ab, und ein Vergleichs-Schritt verlangt sie byte-identisch samt Exit-Code. Erst wenn
   Tests und Vergleich grün sind, setzt ein letzter Schritt den Bild-Tag `vX.Y.Z` (und bei stabilen
   Releases `latest`) per `imagetools create` auf denselben Digest — kein zweiter Bau, kein neuer
   Digest. Ein roter Lauf lässt den Index ungetaggt zurück; er trägt keine Version.
6. **Der Pin ist der Index-Digest.** `a-check.mk`, beide READMEs und `version.md#aktuell` nennen
   ihn; `gate-consistency` prüft ihn unverändert auf Gleichheit.
7. **Spiegel:** der Index wird per `imagetools create` unverändert nach Docker Hub kopiert. Die
   Gleichheits-Größe ist der **Index-Digest** auf beiden Registries, dazu die Plattform-Menge des
   Spiegels gleich der der Quelle; fail-closed wie bisher (ADR-0039 Punkte 2 und 3). Damit löst der
   GHCR-Digest auch auf dem Spiegel auf.
8. **Lokal:** ein Make-Target baut denselben Index als OCI-Archiv; lokal getestet wird die
   Plattform des Hosts. Die übrigen Gate-Stufen (Lint, Test, Coverage) bleiben einplattformig — sie
   prüfen Quelltext, nicht das Plattform-Bild.

## Verglichene Alternativen

| Option | Pro | Contra |
|---|---|---|
| A — nichts tun (nur `linux/amd64`) | kein Aufwand; macOS läuft unter Emulation | Emulation langsam und für den Konsumenten unsichtbar; der Bedarf ist benannt |
| B — Bau **und** Test unter QEMU auf einem amd64-Runner | ein Runner-Typ, ein Job | der Bau braucht den Emulator gar nicht (Cross-Compile); ein Test unter Emulation prüft den Emulator mit, nicht die Plattform |
| C — je Plattform ein eigener nativer Bau, danach Indizes zusammenführen | jeder Bau läuft nativ | **zwei** Bauten; ohne Reproduzierbarkeit zwei Bilder, die zufällig gleich heißen; die Zusammenführung ist ein dritter Schritt mit eigenem Fehlerbild |
| D — ein Bau, das OCI-Archiv als Workflow-Artefakt an die Test-Jobs, dort per Digest aus einer job-lokalen Registry getestet, Upload nach GHCR erst danach | nichts ungetaggt in GHCR | je Test-Job eine lokale Registry und ein Kopier-Werkzeug mehr; getestet wird eine lokale Kopie, und der Upload ist ein weiterer Schritt, dessen Digest erneut verglichen werden muss — der Weg zum Konsumenten (Ziehen aus GHCR) bleibt ungetestet |
| **E — ein Cross-Compile-Bau, ungetaggt nach GHCR, nativer Test je Plattform per Digest, dann Tag (gewählt)** | getestet wird der Index-Digest, der getaggt wird, auf demselben Weg, auf dem ein Konsument ihn zieht; kein Emulator | ein roter Lauf hinterlässt einen ungetaggten Index in GHCR; der arm64-Runner ist ein Schalter außerhalb des Repos; auch hier landet das Plattform-Bild beim Test im lokalen Bildspeicher des Runners |

## Konsequenzen

- Positiv: ein Pin für beide Plattformen; macOS mit Apple Silicon läuft nativ.
- Positiv: der Spiegel ist strenger als bisher — der Index-Digest deckt jede Plattform, der
  Config-Digest deckte genau eine.
- Negativ: die Pipeline wird mehrjobig (Bau, zwei Tests, Tag, Spiegel); ein Fehler zeigt sich erst
  am Tag.
- Negativ: ungetaggte Index-Versionen bleiben nach roten Läufen in GHCR liegen (sichtbar in der
  Paketliste, von keinem Pin referenziert).
- Negativ: `make image-scan` prüft bei einem Index die Plattform des scannenden Rechners
  (`linux/amd64`); das arm64-Bild bleibt ungescannt. Beide stehen auf derselben distroless-Version
  und demselben Quellstand — die Lücke ist benannt, nicht geschlossen.
- Folgepflicht: Dockerfile und Make-Target; Release-Pipeline mit Test-Jobs, Tag-Schritt und neuem
  Spiegel; `releasing.md` nennt den arm64-Runner unter den Vorbedingungen; der Image-Test nimmt
  eine Bild-Referenz von außen; die Hub-Seite sagt, dass der GHCR-Digest dort auflöst.

## Fitness Function (falls maschinell prüfbar)

| Tooling | Regel | Make-Target |
|---|---|---|
| Release-Pipeline | der getaggte Digest ist der getestete; der Index nennt genau zwei Plattformen; die Scan-Ausgabe beider Plattformen ist byte-identisch | — (Pipeline-Schritt, kein lokales Target) |
| `tools/image-test.sh` | Image-Test gegen eine übergebene Bild-Referenz, je Plattform nativ | `make image-test` |
| `tools/gate-consistency.sh` | ein Digest an allen Pin-Stellen | `make gate-consistency` |

## Re-Evaluierungs-Trigger

- Die arm64-Runner stehen diesem Repo nicht mehr zur Verfügung — dann Option B für den Test; das
  ist eine Änderung der Spezifikation (sie verlangt einen Rechner der Plattform) und eine Folge-ADR.
- Ein Konsument braucht eine weitere Plattform.
- GHCR oder Docker Hub erhalten den Index-Digest beim Kopieren nicht — dann fällt die
  Gleichheits-Größe zurück auf den Config-Digest je Plattform.
- Attestierungen werden gefordert (Signatur, Provenance) — dann ändert sich die Plattform-Menge des
  Index (zusätzliche `unknown/unknown`-Einträge), und Punkt 3 ist neu zu entscheiden.

## Geschichte

| Datum | Ereignis | Verweis |
|---|---|---|
| 2026-10-07 | Proposed | Maintainer-Anweisung; Messung lokal (Cross-Compile, Reproduzierbarkeit, Index-Kopie) |
| 2026-10-07 | Accepted | Maintainer-Abnahme („ok“) nach unabhängigem Review samt Delta-Review |

Nach `Accepted` wird diese Datei **nicht mehr inhaltlich überschrieben**.
Spätere Korrekturen oder Schärfungen entstehen als neue ADR mit
`Supersedes ADR-NNNN` (Baseline-Regelwerk `modul-04-adrs.md`
§Hard Rule für Accepted-ADRs).
