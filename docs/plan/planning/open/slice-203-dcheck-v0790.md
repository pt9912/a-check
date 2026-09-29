# slice-203 — d-check auf `v0.79.0` heben

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt. Er wechselt nur durch `git mv`.

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Maintainer-Wunsch „d-check aktualisieren" (2026-09-29). Pin
`v0.75.0` → `v0.79.0` — vier Releases. [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
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

- **Neue Module opt-in schalten** (Modul `file` mit `max-lines` — Welle 147) und
  `structure.open-tasks-require-marker` (Attestierungs-Hebel). *Ein Folge-Slice übernimmt
  es*: beide verdienen eine eigene Entscheidung über Schwellen und Scopes, nicht einen
  Anhängsel-Edit am Pin.
- **Inhaltliche Anpassungen an `.d-check.yml`**, die über die Regenierung
  hinausgehen. *Es wäre ein anderer Vorgang*: der Pin-Bump ändert die Konfiguration
  nicht — er lässt sie gegen die neue Fassung laufen.

## 2. Definition of Done

- [ ] `DCHECK_IMAGE`/`DCHECK_DIGEST` tragen `v0.79.0` bzw. den Release-Digest;
      `d-check.mk` ist mit `--print-mk` der neuen Fassung regeneriert; die
      a-check-Pin-Anpassung (Kommentar-Block, Digest-Zeile) ist nachgeführt.
- [ ] Gates und Verifikation laufen gegen den neuen Pin; **jeder** neue Befund
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
  — **Ausgang:** bei Closure.
- **`vcs`/`commits` lösen Ranges immer auf** — Läufe, die still grün waren,
  können Exit 2 melden. — **Ausgang:** bei Closure.

## 7. Closure-Notiz

*(wird vor dem `git mv` nach `done/` gefüllt)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** GATE (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde durchgegangen
(2026-09-29): keine Treffer in GATE für diesen Vorgang.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
