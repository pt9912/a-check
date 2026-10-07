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

**Verantwortlich:** —

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

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Code, Dockerfile, Pipeline.** *Schicht-Abgrenzung*: slice-217 und slice-218 setzen den
  abgenommenen Vertrag um; hier wird nur gemessen und geschrieben.
- **Weitere Plattformen, native Binaries, Signaturen.** *Ein anderer Vorgang*: Out-of-Scope der
  Welle (§6 des Welle-Plans).

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
