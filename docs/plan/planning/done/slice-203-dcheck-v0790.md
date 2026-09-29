# slice-203 — d-check auf `v0.79.0` heben

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt. Er wechselt nur durch `git mv`.

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Maintainer-Wunsch „d-check aktualisieren" (2026-09-29). Pin
`v0.75.0` → `v0.79.0` — sieben Tags. [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
[`AC-QA-03`](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der d-check-Pin steht auf `v0.79.0` (Release-Digest `b4b8756b…`); `d-check.mk`
ist mit `--print-mk` der neuen Fassung regeneriert; die Verhaltensänderungen der vier
Releases (`links` prüft Referenz-Definitionen und mehrzeilige Adressen, `vcs`/`commits`
lösen Ranges immer auf) sind gegen den Bestand gefahren und ihre Befunde klassifiziert.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Neue Module opt-in schalten** (Modul `file` mit `max-lines` — im d-check-CHANGELOG als slice-236 attribuiert; die Motivation trägt das Baseline-Delta Welle 147, „Guide-Datei-Wildwuchs") und
  `structure.open-tasks-require-marker` (Attestierungs-Hebel). *Ein Folge-Slice übernimmt
  es*: beide verdienen eine eigene Entscheidung über Schwellen und Scopes, nicht einen
  Anhängsel-Edit am Pin.
- **Inhaltliche Anpassungen an `.d-check.yml`**, die über die Regenierung
  hinausgehen. *Es wäre ein anderer Vorgang*: der Pin-Bump ändert die Konfiguration
  nicht — er lässt sie gegen die neue Fassung laufen.

## 2. Definition of Done

- [x] `DCHECK_IMAGE`/`DCHECK_DIGEST` tragen `v0.79.0` bzw. den Release-Digest;
      `d-check.mk` ist mit `--print-mk` der neuen Fassung regeneriert; die
      a-check-Pin-Anpassung (Kommentar-Block, Digest-Zeile) ist nachgeführt.
- [x] Gates und Verifikation laufen gegen den neuen Pin; **jeder** neue Befund
      gegenüber `v0.75.0` ist klassifiziert (beheben oder als Grenze benennen).
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md);
      Closure-Notiz mit Lerneintrag; Register fortgeschritten; jedes Risiko
      trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `d-check.mk` | update | `--print-mk` der neuen Fassung; a-check-Pin-Anpassung nachgeführt |

## 4. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen — vorab benennen:**

- `in-progress` → `next` (zu groß): meldet die neue Fassung mehr als drei
  Klassifizierungs-Befunde, wird die Behebung aufgeteilt.
- `in-progress` → `open` (blockiert): ist das Image nicht beschaffbar, ist die
  Beschaffung ein eigener Vorgang.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 6. Risiken und offene Punkte

- **`links` prüft Referenz-Definitionen neu** — tote Ziele hinter
  `[label]: ziel`-Definitionen, die bisher unsichtbar waren, werden gemeldet.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: `make doc-check` mit
  der neuen Fassung meldet **0 Befunde** über 609 Dateien — kein totes
  Referenz-Ziel und keine mehrzeilige Adresse im Bestand.
- **`vcs`/`commits` lösen Ranges immer auf** — Läufe, die still grün waren,
  können Exit 2 melden. — **Ausgang:** *entfallen*, gestrichen mit Begründung:
  `make gates` lief die volle Sequenz gegen den neuen Pin — die Range-Auflösung
  (doc-immutable, doc-commits, trace-check) blieb grün; kein stiller
  Null-Lauf hat einen echten Mangel verdeckt.

## 7. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** *Fremd-Provenanz mit fremder
Vorgangs-Kennung zitieren.* Das Modul `file` stammt aus dem d-check-CHANGELOG
(dort `slice-236`); meine Vorversion attribuierte es auf „Welle 147" — eine
Nummer aus dem Baseline-Delta, die im geprüften d-check-Bestand nicht
existiert (Review-F-2). Zwei Zählräume, zwei Kennungssysteme: die Kennung
stammt aus dem Zählraum des geprüften Artefakts, nicht aus dem eigenen.

**Was hat funktioniert:** der Pin-Bump ist ein Zwei-Zeilen-Vorgang geblieben —
`DCHECK_DIGEST` heben, `d-check.mk` mit `--print-mk` der neuen Fassung
regenerieren, die a-check-Pin-Anpassung nachführen. Der Diff ist minimal und
lesbar.

**Was ging anders als geplant:** nichts — beide §6-Risiken (strenge
Range-Auflösung, Referenz-Definitions-Prüfung) sind gemessen und entfallen.

**Beobachtungs-Register (`../observations/`):** keine Beobachtung angefallen.

**Folge-Slices:** — *(keine; die opt-in-Neuheiten des neuen Standes — Modul
`file`, `structure.open-tasks-require-marker` — warten auf eigene Entscheide
und sind in slice-199s Folge-Chain nicht verplant.)*

**Risiken aus §6:** jedes mit genau einem Ausgang — siehe §6.

**Drei Paarungen:** Anker — kein Steering-Loop-Eintrag · Folge-Slice — keine ·
Register — keine Beobachtung angefallen.

**Beobachtbare Architektur-Aussage:** die Link-Hygiene des Bestands ist
kontinuierlich stark genug, dass vier d-check-Releases mit verschärfter
Erkennung (Referenz-Definitionen, mehrzeilige Adressen, immer aufgelöste
Ranges) **null** Nacharbeit erzeugen — gemessen: `make doc-check` 0 Befunde
über 609 Dateien, `make gates` grün, `make verify` grün gegen `v0.79.0`.

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** GATE (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde durchgegangen
(2026-09-29): keine Treffer in GATE für diesen Vorgang.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
