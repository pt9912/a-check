# slice-216 — Multi-Arch-Image: Spec-first (Messung, Lastenheft-CR, ADR, Spezifikation)

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** welle-18 — [Welle-Plan](../welle-18-multi-arch-image.md).

**Bezug:** [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk),
[AC-FA-DIST-002](../../../../spec/lastenheft.md#ac-fa-dist-002),
[AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit);
[ADR-0004](../../adr/0004-distribution-image-mk.md), [ADR-0007](../../adr/0007-latest-tag-politik.md),
[ADR-0039](../../adr/0039-spiegel-gleichheit-ist-der-config-digest.md).

**Berührte Spec-Stellen:** `spezifikation.md`
§[SPEC-DIST-001](../../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `next/` (Maintainer: „ok machen wir so“, 2026-10-07).

**Autor:** Claude. **Datum:** 2026-10-07.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der Vertrag für ein Multi-Arch-Release-Image steht zur Abnahme: gemessen (Basis-Images,
Runner, Build-Werkzeug), im Lastenheft als Änderung an
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
und [AC-FA-DIST-002](../../../../spec/lastenheft.md#ac-fa-dist-002), in einer ADR zur
Build-Strategie und in
[SPEC-DIST-001](../../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution).

**Zu entscheiden (Entwurf, Abnahme beim Maintainer):**

1. **Plattformen:** genau `linux/amd64` und `linux/arm64`.
2. **Pin:** der veröffentlichte Digest ist der des **Image-Index**; er löst auf beiden Plattformen
   auf. Alte Pins (Einzel-Manifest) bleiben gültig.
3. **„Getestet = veröffentlicht":** einmal bauen für beide Plattformen; getestet wird das Gebaute,
   nicht ein zweiter Bau. Wie (Registry-Staging, OCI-Layout, Push nach Test), entscheidet die ADR.
4. **Test auf arm64:** nativer arm64-Runner (Vorschlag) gegen Emulation.
5. **Spiegel-Gleichheit** ([AC-FA-DIST-002](../../../../spec/lastenheft.md#ac-fa-dist-002)): je
   Plattform der Config-Digest, oder der Index unverändert kopiert und dessen Digest verglichen —
   gemessen, welche Größe beim Kopieren stabil bleibt.
6. **Cross-Compile:** die Build-Stufe läuft auf der Plattform des Runners und kompiliert für die
   Ziel-Plattform; die Laufzeit-Stufe führt keinen Befehl aus.

**Plan-Änderung 2026-10-07 (Review F-1 bis F-13, vor dem Fix):** Die Reproduzierbarkeit ist
**keine** Zusage mehr — die Gegenmessung des Reviews widerlegt M4 in der Allgemeinheit (ein
warmer Cache mit älterer mtime ergibt einen anderen Index-Digest); „getestet = veröffentlicht"
hängt an *einmal bauen*, nicht an ihr. Die Ausgabe-Gleichheit beider Plattformen
([AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
Boundary) bekommt einen Träger in jedem Release (ADR und slice-218). Der Slice nimmt außerdem die
CHANGELOG-Zeile für die Vertragsänderung mit; die Folge-Slices nehmen den Image-Test mit
Bild-Referenz (slice-217) und die Hub-Seite (slice-218) in ihren Plan auf.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Code, Dockerfile, Pipeline.** *Schicht-Abgrenzung*: slice-217 und slice-218 setzen den
  abgenommenen Vertrag um; hier wird nur gemessen und geschrieben.
- **Weitere Plattformen, native Binaries, Signaturen.** *Ein anderer Vorgang*: Out-of-Scope der
  Welle (§6 des Welle-Plans).

## 1b. Messung (2026-10-07)

Alle Läufe lokal (Docker 29.8.2, buildx 0.37.1, Builder `docker-container`, Host `linux/amd64`
**ohne** arm64-Emulation), mit einer Dockerfile-Variante im Scratchpad — das Repo-Dockerfile ist
unverändert. Die Variante setzt `--platform=$BUILDPLATFORM` auf die `deps`-Stufe und
`GOOS`/`GOARCH` aus `TARGETOS`/`TARGETARCH` im Kompilier-Schritt; die Laufzeit-Stufe führt keinen
Befehl aus.

| # | Frage | Ergebnis | Geltungsbereich |
|---|---|---|---|
| M1 | Tragen die gepinnten Basis-Images beide Plattformen? | ja — `golang@sha256:0ecdc2a9…` und `distroless/static-debian12@sha256:d093aa3e…` sind **Image-Indizes** mit `linux/amd64` und `linux/arm64/v8` | die zwei Pins im `Dockerfile`; `golangci-lint` nicht geprüft (Gate-Stufe, nicht im Release-Bild) |
| M2 | Baut der Host beide Plattformen ohne Emulation? | ja — Index mit `linux/amd64` und `linux/arm64`; das extrahierte Binary ist je ein statisches ELF für x86-64 bzw. AArch64; Versions-Label auf beiden gesetzt | dieser Host; der Lauf beweist Cross-Compile, nicht Lauffähigkeit auf arm64 (kein Emulator, kein arm64-Rechner) |
| M3 | Was legt buildx ungefragt dazu? | zwei Attestierungs-Manifeste (`unknown/unknown`, Provenance) im Index; mit `--provenance=false --sbom=false` enthält der Index **genau** die zwei Plattformen | Builder `docker-container`; der klassische `docker`-Treiber nicht gemessen |
| M4 | Ist der Bau reproduzierbar? | ohne Vorkehrung **nein**: zweiter Lauf mit Cache gleich, ohne Cache anderer Index-Digest — das Binary ist bitgleich (`e47f8a3a…`), es unterscheiden sich nur die mtime im Laufzeit-Layer und der `created`-Zeitstempel. Mit `SOURCE_DATE_EPOCH` (Commit-Zeit) und `rewrite-timestamp=true` waren zwei Läufe mit und ohne Cache gleich (`256d2a69…`); **die Gegenmessung im Review widerlegt die Allgemeinheit**: ein warmer Cache aus einem Bau mit anderer Epoch ergab einen anderen Digest (`rewrite-timestamp` setzt nur spätere Zeitstempel zurück) | zwei Läufe je Variante auf diesem Host plus die Review-Gegenmessung; ein Bau auf einem Runner nicht gemessen — **keine** Reproduzierbarkeits-Zusage |
| M5 | Bleibt der Index-Digest beim Hochladen und Kopieren? | Hochladen per `skopeo copy --all`: Digest wie gebaut. Kopie Registry → Registry per `docker buildx imagetools create` und per `skopeo copy --all`: **derselbe** Index-Digest, beide Plattformen. Der heutige Spiegel-Weg `docker tag` + `docker push` liefert dagegen **ein Einzel-Manifest** (`linux/amd64`) mit neuem Digest — **arm64 ginge verloren** | zwei lokale `registry:2`; GHCR und Docker Hub nicht gemessen (erst am Tag); der `docker push`-Befund gilt für den klassischen Bildspeicher (`overlay2`), nicht für den containerd-Bildspeicher |
| M6 | Gibt es arm64-Runner für dieses Repo? | das Repo ist **öffentlich**; GitHub bietet dafür `ubuntu-24.04-arm` an | Aussage der Plattform, nicht durch einen Lauf belegt — der erste Lauf ist Teil von slice-218 |
| M7 | Läuft der Image-Test auf arm64 unverändert? | `tools/image-test.sh` ist reines Bash und extrahiert das Binary aus dem Bild; er nimmt heute nur `$(IMAGE):dev` und nimmt Host-Arch = Bild-Arch an — für arm64 braucht er eine Bild-Referenz von außen | Lesart des Skripts, kein Lauf |

**Folgerungen für den Vertrag:** (a) der Pin ist der Index-Digest; (b) „getestet =
veröffentlicht" heißt: **einmal** bauen, das Gebaute (per Digest) auf beiden Plattformen testen,
dann taggen — ein zweiter Bau kann ein anderes Bild ergeben (M4); (c) der Spiegel muss den
Index kopieren, nicht neu pushen (M5), und seine Gleichheits-Größe kann der **Index-Digest**
werden — strenger als der Config-Digest aus
[AC-FA-DIST-002](../../../../spec/lastenheft.md#ac-fa-dist-002), weil er alle Plattformen deckt;
ob GHCR und Docker Hub ihn ebenfalls erhalten, zeigt erst der Tag (Risiko §6).

## 2. Definition of Done

- [ ] Messung: Basis-Images als Index mit beiden Plattformen (Digest), Verfügbarkeit der
      arm64-Runner für dieses Repo, Stabilität von Index- und Config-Digest beim Kopieren zwischen
      Registries — je mit Geltungsbereich (Mess-Regel 1).
- [ ] Lastenheft-CR an [AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)/[AC-FA-DIST-002](../../../../spec/lastenheft.md#ac-fa-dist-002) (drei Pfade plus Out-of-Scope, Versions-Bump, Historie) und Folge-ADR zur Build-Strategie mit Index-Eintrag.
- [ ] [SPEC-DIST-001](../../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution) präzisiert (Plattformen, Index-Digest als Pin, getestetes = veröffentlichtes Bild).
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` grün; Abnahme des Vertrags durch den Maintainer vor dem Closure-Commit.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `spec/lastenheft.md` | update | CR an DIST-001/002 |
| `docs/plan/adr/0043-…md`, `docs/plan/adr/README.md` | neu / update | Build-Strategie |
| `spec/spezifikation.md` | update | Laufzeitform und Distribution |
| `CHANGELOG.md` | update | Vertragsänderung in `[Unreleased]` (Review F-5) |
| `slice-217`, `slice-218`, Welle-Plan | update | Folgepflichten der ADR (Review F-4, F-13) |

## 4. Trigger

**Start** (`next` → `in-progress`): Welle eröffnet, WIP-Limit frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): entfällt voraussichtlich — drei Dokumente.
- `in-progress` → `open` (blockiert): die arm64-Runner sind für dieses Repo nicht verfügbar
  **und** Emulation scheidet aus — dann zurück an den Maintainer.

## 5. Closure-Trigger

DoD vollständig, Vertrag abgenommen, `make gates` Exit 0, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Mechanik, die nur am Tag läuft:** Push, Spiegel und arm64-Test laufen erst in der
  Release-Pipeline; der Vertrag kann Zusagen machen, die vorher kein Lauf prüft (vgl.
  `BEO-GATE/ungelaufene-mechanik-docker-hub-spiegel`, 1×). — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*
- **Zusage weiter als ihre Durchsetzung:** „getestet = veröffentlicht" auf zwei Plattformen ist
  leicht zu schreiben und schwer zu belegen (vgl. `BEO-GATE/zusage-weiter-als-ihre-durchsetzung`,
  2×). — **Ausgang:** *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `SPEC` (1, 2, 3 ✓) · `ADR` (1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-10-07):
`BEO-GATE/ungelaufene-mechanik-docker-hub-spiegel` (1×) und
`BEO-GATE/zusage-weiter-als-ihre-durchsetzung` (2×) als Risiken in §6;
`BEO-GATE/hebungskanal-haengt-an-repo-externen-schaltern` (1×) — die Runner-Verfügbarkeit ist ein
solcher Schalter, darum Teil der Messung. Ein dritter Beleg der 2×-Einträge wäre eine Lücke.
`SPEC`/`ADR`: keine weiteren Treffer, die diese Stellen berühren.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
