# Welle welle-18: Multi-Arch-Image (`linux/amd64` + `linux/arm64`)

**Lifecycle:** Diese Datei entsteht bei der **Eröffnung** der Welle und liegt
flach unter `docs/plan/planning/`; bei Closure wandert sie per `git mv` nach
`done/` (neben ihre `welle-18-results.md`). Der Zustand ist die
Verzeichnis-Position — kein Status-Feld.

**Zielmeilenstein:** kein Meilenstein-Bezug.

**Verantwortlich:** Claude (Planner); Abnahme beim Maintainer. **Datum:** 2026-10-07.

---

## 1. Welle-Ziel

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wann Arbeit eine Welle braucht.

Das Release-Image von a-check läuft nativ auf `linux/amd64` **und** `linux/arm64` — damit auf
macOS mit Apple Silicon ohne Emulation. Ein Konsument pinnt **einen** Digest, der auf beiden
Plattformen auflöst, und die Zusagen aus
[AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
und [AC-FA-DIST-002](../../../spec/lastenheft.md#ac-fa-dist-002) (getestetes Bild = veröffentlichtes
Bild, Spiegel = dasselbe Bild) gelten je Plattform.

## 2. Trigger (Welle startet)

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Roadmap-Regeln — beobachtbar, und kein Ergebnis dieser Welle.

- **Maintainer-Anweisung** (dieses Gespräch, 2026-10-07: „ok machen wir so") auf die Frage nach
  einem Multi-Arch-Image für macOS/ARM. Die Welle stand nicht unter *Nächste Wellen*; eine
  Umplanung liegt darum nicht vor.
- welle-17 liegt in `done/`, `in-progress/` ist leer.

## 3. Closure-Trigger (Welle schließt)

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wann Arbeit eine Welle braucht — das *Mehr* ist der Release-Beleg auf beiden Plattformen gegen
das veröffentlichte Image; keine Slice-DoD kann ihn liefern, weil die Pipeline nur am Tag läuft.

- Alle Slices aus §4 liegen in `done/`.
- `make ci` und `make verify` Exit 0 auf dem finalen Stand — Ausgabe in eine Datei, Exit-Code
  getrennt geprüft (Replay-Ersatz laut [MR-028](../../../harness/conventions.md#mr-028)).
- Ein Release-Tag ist veröffentlicht und nach [`docs/user/releasing.md`](../../user/releasing.md)
  re-gepinnt; die Release-Pipeline ist grün **einschließlich** des Tests auf `linux/arm64`.
- **Gegenprobe am veröffentlichten Image:** `docker buildx imagetools inspect` des gepinnten
  Digests nennt genau `linux/amd64` und `linux/arm64`, auf GHCR **und** Docker Hub mit demselben
  Index-Digest; der Scan läuft auf `linux/amd64` lokal und auf `linux/arm64` im nativen Test-Job
  der Release-Pipeline mit identischer Ausgabe (Job-Log und Vergleichs-Schritt) — ein lokaler
  arm64-Lauf ist auf diesem Host nicht möglich, ein Lauf auf einem Mac des Maintainers ist
  willkommen, aber kein Kriterium; das Versions-Label stimmt auf beiden.
- Ergebnis-Notiz `done/welle-18-results.md` geschrieben.

## 4. Slices in dieser Welle

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md`
§Lifecycle als State Machine — der Zustand eines Slice ist sein Verzeichnis.

| Slice | Titel | Bezug |
|---|---|---|
| [slice-216](done/slice-216-multi-arch-spec-first.md) | Spec-first: Messung, Lastenheft-CR, ADR zur Build-Strategie, Spezifikation — zur Abnahme | [AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk), [AC-FA-DIST-002](../../../spec/lastenheft.md#ac-fa-dist-002) |
| [slice-217](next/slice-217-multi-arch-build.md) | Build für beide Plattformen (Dockerfile, Make-Target, lokaler Beleg) | dieselben |
| [slice-218](open/slice-218-multi-arch-release.md) | Release-Pipeline: ein Bau, Test auf beiden Plattformen, Push, Spiegel | dieselben |

**Reihenfolge ist Abhängigkeit:** slice-217 startet erst, wenn slice-216 in `done/` liegt und der
Maintainer den Vertrag abgenommen hat; slice-218 baut auf dem Make-Target aus slice-217.

## 5. Abhängigkeiten

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Roadmap-Struktur: fünf Abschnitte.

- Blockiert: keine.
- Wird blockiert von: keiner Welle. **Repo-extern:** die arm64-Runner von GitHub Actions — ihre
  Verfügbarkeit für dieses Repo ist ein Schalter, den das Repo nicht sieht; slice-216 misst sie.

## 6. Out-of-Scope für diese Welle

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Eröffnung Schritt 1.

- **`linux/arm/v7`, `s390x`, `ppc64le`.** Kein Konsument hat Bedarf; jede weitere Plattform ist ein
  weiterer Test-Lauf in jedem Release.
- **Native macOS- oder Windows-Binaries.** Nicht-Docker-Distribution schließt
  [AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
  aus; Docker Desktop auf macOS führt `linux/arm64` aus — das genügt dem Ziel.
- **Ein Single-Arch-Fallback** (getrennte Tags je Plattform). Ein Pin, beide Plattformen ist das
  Ziel; zwei Pin-Arten wären zwei Quellen für dieselbe Version.
- **Multi-Arch für die Werkzeug-Images** (`archive-wave`, Gate-Stages). Sie laufen nur in der CI
  und lokal beim Entwickler — nicht beim Konsumenten.
- **CVE-Scan des arm64-Bilds.** `make image-scan` prüft die Plattform des scannenden Rechners;
  die Lücke ist in der ADR benannt, ihr Schließen ein eigener Vorgang.
- **Signaturen und Provenance-Attestierungen** (cosign, SLSA). Ein eigener Vorgang; nur dann hier,
  wenn die Build-Werkzeugkette sie unabschaltbar mitbringt — das entscheidet die ADR.

## 7. Closure-Notiz

Ergebnis: `welle-18-results.md` — Geschwister im Ruheort `done/`
Zähler: `../observations/` — das Beobachtungs-Register, eine Ebene über dem Ruheort
