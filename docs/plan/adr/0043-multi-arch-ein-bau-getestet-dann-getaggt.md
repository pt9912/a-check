# ADR-0043: Multi-Arch-Image — ein Cross-Compile-Bau, je Plattform nativ getestet, dann getaggt; der Spiegel kopiert den Index

**Status:** Proposed

**Datum:** 2026-10-07

**Autor:** pt9912 (Anweisung und Abnahme), ausgeführt im Auftrag

**Bezug:** [AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[AC-FA-DIST-002](../../../spec/lastenheft.md#ac-fa-dist-002),
[AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit);
[ADR-0039](0039-spiegel-gleichheit-ist-der-config-digest.md) — **löst deren Punkt 1 ab** (die
Gleichheits-Größe ist der Index-Digest statt des Config-Digests) und **den Pin-Satz in Punkt 4**
(„wer vom Spiegel zieht, nimmt den Digest der Registry, aus der er zieht"); ihre Punkte 2, 3 und
der übrige Punkt 4 gelten weiter, ADR-0039 selbst bleibt unverändert `Accepted`.
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
- Der Bau ist ohne Vorkehrung **nicht** reproduzierbar: das Binary ist bitgleich, aber mtime und
  `created` ändern den Digest. Mit `SOURCE_DATE_EPOCH` und `rewrite-timestamp=true` ist der
  Index-Digest über Läufe mit und ohne Cache identisch.
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
Plattform auf einem nativen Runner, und erst danach den Versions-Tag** — im Einzelnen:

1. **Plattformen:** genau `linux/amd64` und `linux/arm64`.
2. **Cross-Compile:** die Stufe, die kompiliert, läuft auf der Plattform des Bau-Rechners
   (`--platform=$BUILDPLATFORM`) und setzt `GOOS`/`GOARCH` aus `TARGETOS`/`TARGETARCH`; die
   Laufzeit-Stufe bleibt ohne `RUN`. Ein Emulator ist für den Bau nicht nötig.
3. **Der Index trägt nur die zwei Plattformen:** Attestierungen und SBOM sind abgeschaltet
   (`--provenance=false --sbom=false`); sie sind ein eigener Vorgang.
4. **Reproduzierbar:** `SOURCE_DATE_EPOCH` ist die Commit-Zeit des getaggten Commits,
   `rewrite-timestamp=true` setzt die Zeitstempel im Bild darauf. Derselbe Commit ergibt denselben
   Index-Digest.
5. **Ein Bau, getestet, dann getaggt:** Die Pipeline baut den Index **einmal** und lädt ihn
   **ohne Tag** (nur per Digest) nach GHCR. Je ein Test-Job auf einem nativen `linux/amd64`- und
   einem nativen `linux/arm64`-Runner zieht **diesen Digest** und fährt den Image-Test samt
   Versions-Label. Erst wenn beide grün sind, setzt ein letzter Schritt `vX.Y.Z` (und bei stabilen
   Releases `latest`) per `imagetools create` auf denselben Digest — kein zweiter Bau, kein neuer
   Digest. Ein roter Test lässt den Index ungetaggt zurück; er trägt keine Version.
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
| D — ein Bau, Test per OCI-Archiv als Workflow-Artefakt statt Registry | nichts ungetaggt in GHCR | das Laden eines Multi-Plattform-Archivs in den klassischen Bildspeicher holt eine Plattform und schreibt sie neu — getestet würde eine Kopie, nicht der Index |
| **E — ein Cross-Compile-Bau, ungetaggt nach GHCR, nativer Test je Plattform per Digest, dann Tag (gewählt)** | getestet ist bitgenau, was getaggt wird; kein Emulator; reproduzierbar | ein roter Lauf hinterlässt einen ungetaggten Index in GHCR; der arm64-Runner ist ein Schalter außerhalb des Repos |

## Konsequenzen

- Positiv: ein Pin für beide Plattformen; macOS mit Apple Silicon läuft nativ.
- Positiv: der Spiegel ist strenger als bisher — der Index-Digest deckt jede Plattform, der
  Config-Digest deckte genau eine.
- Positiv: derselbe Commit ergibt denselben Digest; ein Release lässt sich nachbauen und vergleichen.
- Negativ: die Pipeline wird mehrjobig (Bau, zwei Tests, Tag, Spiegel); ein Fehler zeigt sich erst
  am Tag.
- Negativ: ungetaggte Index-Versionen bleiben nach roten Läufen in GHCR liegen (sichtbar in der
  Paketliste, von keinem Pin referenziert).
- Folgepflicht: Dockerfile und Make-Target; Release-Pipeline mit Test-Jobs, Tag-Schritt und neuem
  Spiegel; `releasing.md` nennt den arm64-Runner unter den Vorbedingungen; der Image-Test nimmt
  eine Bild-Referenz von außen; die Hub-Seite sagt, dass der GHCR-Digest dort auflöst.

## Fitness Function (falls maschinell prüfbar)

| Tooling | Regel | Make-Target |
|---|---|---|
| Release-Pipeline | der getaggte Digest ist der getestete; der Index nennt genau zwei Plattformen | — (Pipeline-Schritt, kein lokales Target) |
| `tools/image-test.sh` | Image-Test gegen eine übergebene Bild-Referenz, je Plattform nativ | `make image-test` |
| `tools/gate-consistency.sh` | ein Digest an allen Pin-Stellen | `make gate-consistency` |

## Re-Evaluierungs-Trigger

- Die arm64-Runner stehen diesem Repo nicht mehr zur Verfügung — dann Option B für den Test.
- Ein Konsument braucht eine weitere Plattform.
- GHCR oder Docker Hub erhalten den Index-Digest beim Kopieren nicht — dann fällt die
  Gleichheits-Größe zurück auf den Config-Digest je Plattform.
- Attestierungen werden gefordert (Signatur, Provenance) — dann ändert sich die Plattform-Menge des
  Index (zusätzliche `unknown/unknown`-Einträge), und Punkt 3 ist neu zu entscheiden.

## Geschichte

| Datum | Ereignis | Verweis |
|---|---|---|
| 2026-10-07 | Proposed | Maintainer-Anweisung; Messung lokal (Cross-Compile, Reproduzierbarkeit, Index-Kopie) |

Nach `Accepted` wird diese Datei **nicht mehr inhaltlich überschrieben**.
Spätere Korrekturen oder Schärfungen entstehen als neue ADR mit
`Supersedes ADR-NNNN` (Baseline-Regelwerk `modul-04-adrs.md`
§Hard Rule für Accepted-ADRs).
