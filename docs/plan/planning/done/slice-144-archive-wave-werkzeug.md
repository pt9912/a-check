# slice-144 — Etappe A: `archive-wave`-Werkzeug aus `d-check` übernommen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Etappe **A** aus [slice-143 §6](../done/slice-143-archivierung-delta-analyse.md#6-vorschlag-vier-etappen),
per Maintainer-Wort gezogen 2026-09-05 („Etappe A — siehe d-check"). [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Harness-Werkzeug ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

`tools/archive-wave/` aus `d-check` **unverändert** übernehmen — sein eigenes `README.md`
erklärt es als für genau diesen Zweck gebaut: „Portabel: dieses Verzeichnis kann unverändert in
jedes Repo mit demselben `docs/plan/planning/`-Layout kopiert werden." Geprüft, nicht nur zitiert
(§3).

## 2. Betroffene Module

`tools/archive-wave/` (16 Dateien, 2484 Zeilen, eigenes `go.mod`, eigenes `Dockerfile`, eigenes
`Makefile`) neu; `Makefile` (zwei Targets: `archive-wave-test`, `archive-wave`); `AGENTS.md` §4
(zwei Zeilen). Zwei Schichten (Gate-/Werkzeug-Schicht, Harness-Einstieg für die
`AGENTS.md`-Deklaration).

## 3. Portabilität geprüft, nicht angenommen

- **Keine `d-check`-interne Kopplung.** `grep -rn "d-check" *.go` findet nur Kommentare, die
  Portabilität *begründen*, und den `<!-- d-check:ignore -->`-Marker auf der `Hervorgegangen:`-Zeile
  des Stubs — a-check nutzt denselben `d-check` als Schwester-Tool für `make doc-check`, der Marker
  greift also identisch.
- **Keine externen Abhängigkeiten.** `go.mod` nennt nur `module archive-wave` / `go 1.27.0`, kein
  `require`-Block; alle Importe sind Stdlib (`archive/zip`, `path/filepath`, `regexp`, …).
- **Eigene Testsuite grün.** `make archive-wave-test` (Docker, `go test ./...`) — Exit 0.
- **Dry-Run gegen den echten a-check-Bestand** (kein `APPLY=1`, schreibt nichts, `:ro`-Mount):
  `make archive-wave WELLE=welle-01` findet `slice-001` korrekt samt beider Review-Reports und
  meldet einen einzigen geplanten Verweis-Fix (`roadmap.md`). Das Werkzeug arbeitet also korrekt
  gegen a-checks Dateiformen.

## 4. Zwei Funde beim Dry-Run — Präzedenz greift, aber nicht überall

1. **CLI-Kurzform vs. a-checks Wellen-ID.** `WELLE=welle-01-fundament` (a-checks eigene Roadmap-Form)
   findet **null** Slices; erst `WELLE=welle-01` (die Ziffernfolge, die das Werkzeug aus dem
   `**Welle:**`-Feld extrahiert) trifft `slice-001`. `d-check`s eigene Wellen-Kennungen sind rein
   numerisch (`welle-87`); a-checks tragen durchgehend einen beschreibenden Suffix. Nicht im
   Werkzeug behoben (bliebe sonst nicht mehr unverändert-portabel) — in `AGENTS.md` §4 dokumentiert
   und als [`BEO-PLAN/welle-id-kurzform-vs-a-checks-beschreibende-form`](../observations/BEO-PLAN/welle-id-kurzform-vs-a-checks-beschreibende-form/observation.md)
   registriert.
2. **`welle-12` und `welle-13` liefern null Slices — gemessen, nicht vermutet.** `CollectSlices`
   liest das `**Welle:**`-Feld jeder Slice-Datei; `welle-01`…`welle-11`s frühe Slices tragen es
   (15 Dateien, gemessen mit `grep`), `welle-12`/`welle-13`s 39 Slices **keine einzige**. Das
   bestätigt [slice-143 §2](../done/slice-143-archivierung-delta-analyse.md#2-ist-stand--gemessen)
   empirisch: „nur 2 von 13 Wellen mit formaler Closure-Notiz" sagte nichts über die
   `**Welle:**`-Feld-Abdeckung — die ist für genau diese zwei formal geschlossenen Wellen leer,
   und für die informellen frühen Wellen teilweise vorhanden. Direkter Blocker für Etappe C.

## 5. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. `make archive-wave-test` ist **nicht** Teil von `gates`
— eigenes `go.mod`, dieselbe Trennung wie bei `d-check` (Beleg in §3, nicht wiederholt in `gates`).

## 6. Was bewusst nicht getan wird

- **Kein `APPLY=1`-Lauf.** Diese Etappe baut und prüft das Werkzeug gegen ein Test-Fixture (seine
  eigene Suite) und per Dry-Run gegen den echten Bestand — schreibend wird nichts angefasst
  (Etappe C/D aus slice-143 §6).
- **Kein Nachtrag der `**Welle:**`-Felder für `welle-12`/`welle-13`.** Das ist Etappe-C-Arbeit,
  hier nur als Blocker gemessen und benannt (§4.2).
- **Keine Anpassung des Werkzeugs an a-checks Wellen-ID-Form.** Würde die Portabilitäts-Zusage
  brechen (§1); die CLI-Kurzform-Diskrepanz ist dokumentiert (§4.1), nicht behoben.
- **Kein eigener `MR`-Eintrag.** Nichts an diesem Werkzeug widerspricht einer a-check-Adaption oder
  einer Baseline-Regel — reine Werkzeug-Übernahme.

## 7. DoD

- [x] `tools/archive-wave/` unverändert übernommen, Portabilität geprüft statt angenommen (§3):
      keine `d-check`-Kopplung, keine externen Abhängigkeiten, eigene Testsuite grün.
- [x] Zwei Make-Targets (`archive-wave-test`, `archive-wave`) verdrahtet und in `AGENTS.md` §4
      dokumentiert; Dry-Run gegen den echten Bestand bestätigt korrektes Verhalten und deckt zwei
      Funde auf (§4).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 8. Closure-Notiz

**Geliefert:** `tools/archive-wave/` betriebsbereit in a-check, mit gemessener statt behaupteter
Portabilität — eigene Testsuite grün, Dry-Run gegen den echten Bestand erfolgreich. Zwei reale
Funde dabei: eine CLI-Namenskonvention, die a-checks eigene Wellen-ID-Form nicht kennt (dokumentiert,
registriert), und der empirische Beleg, dass `welle-12`/`welle-13` — trotz ihrer formalen
Closure-Notizen — keine einzige `**Welle:**`-Feld-Markierung an ihren Slices tragen.

**Lerneintrag — Form: neuer Sensor (benannt, nicht gebaut).** *Ein Werkzeug, das an einem
strukturierten Feld ansetzt, prüft nur, ob das Feld einen Wert **hat** — nicht, ob der Bestand,
gegen den es später laufen soll, dieses Feld überhaupt konsistent führt. Die Lücke zeigt sich erst
im Dry-Run, nicht beim Kopieren.* Das Werkzeug selbst ist korrekt und gut getestet (`d-check`s
eigene 2484 Zeilen Testsuite); trotzdem hätte ein `APPLY=1`-Lauf gegen `welle-12` ins Leere gegriffen
— fail-closed zwar (kein leeres Archiv), aber ohne Vorwarnung vor dem Versuch. Ein Sensor „jeder
Slice mit `**Welle:**`-Bezug auf eine formal geschlossene Welle trägt das Feld" wäre denkbar, ist
hier aber **nicht** gebaut — ein einzelner gemessener Fall (zwei betroffene Wellen) rechtfertigt
noch keinen Vorratsbau (dieselbe Zurückhaltung wie bei [slice-142](../done/slice-142-claude-rules-symlinks-repariert.md) §5).
*Weil* ein Portabilitäts-Test gegen die eigene Suite nur zeigt, dass das Werkzeug **funktioniert** —
ob der **Zielbestand** zu seinen Annahmen passt, zeigt erst ein Lauf gegen echte Daten, und genau
den verlangt diese Etappe ausdrücklich (§1, „geprüft, nicht nur zitiert").

**Zwei beobachtbare Closure-Kriterien:**

1. `make archive-wave-test` meldet Exit 0; `make archive-wave WELLE=welle-01` (Dry-Run) findet
   `slice-001` und beide zugehörigen Reviews.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.**

- *`welle-12`/`welle-13` haben keine `**Welle:**`-Felder an ihren Slices — Etappe C kann sie ohne
  Nachtrag nicht archivieren* — Ausgang: **Folge-Slice**, Etappe C aus
  [slice-143 §6](../done/slice-143-archivierung-delta-analyse.md#6-vorschlag-vier-etappen); der
  Nachtrag selbst ist eine eigene Entscheidung (39 Dateien, `done/`-Inhaltsänderung).
- *CLI-Kurzform-Diskrepanz kann bei künftiger Handbedienung zu einem stillen Fehlaufruf führen* —
  Ausgang: **Beobachtungs-Register**,
  [`BEO-PLAN/welle-id-kurzform-vs-a-checks-beschreibende-form`](../observations/BEO-PLAN/welle-id-kurzform-vs-a-checks-beschreibende-form/observation.md)
  (1×, `slice-144`) — zusätzlich in `AGENTS.md` §4 direkt an der Zeile dokumentiert, die den Fehler
  sonst wiederholen würde.

**Folge-Slices:** keine vergeben. Etappe C (formale Wellen archivieren) ist in slice-143 §6
vorgeschlagen und braucht eine eigene Kennung bei ihrer Eröffnung.

## 9. Sub-Area-Modus

Berührt: **Gate-/Werkzeug-Schicht** (`tools/archive-wave/`, `Makefile`) — Greenfield.
