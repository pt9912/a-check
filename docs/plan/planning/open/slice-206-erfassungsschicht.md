# slice-206 — Erfassungsschicht übernehmen (mit MR-014-Abgrenzung)

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt. Er wechselt nur durch `git mv`.

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Folge-Slice aus
[slice-205](../in-progress/slice-205-init-tool-vergleich.md) (Entscheidung
„Übernehmen"); [`MR-014`](../../../../harness/conventions.md#mr-014)
(Abgrenzung nötig).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** — bis zur Priorisierung.

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** neuer Sensor.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die Erfassungsschicht des Generators ist übernommen —
`.claude/hooks/span-emit.sh` (der committete Wrapper), der Traeger (aus
`.harness/state/bin/`), `harness/erfassung-feldliste.md` (werkzeug-erzeugt)
und die Hook-Verdrahtung in `.claude/settings.json` (PostToolUse,
PostToolUseFailure, SubagentStart → span-emit; Exit 0 in jedem Zweig).

**Die [`MR-014`](../../../../harness/conventions.md#mr-014)-Abgrenzung:** [`MR-014`](../../../../harness/conventions.md#mr-014)
verbietet in seinem geltenden Wortlaut Tool-Call-Spans **uneingeschränkt** —
genau solche erzeugt die Erfassungsschicht. Die Übernahme ist darum nur mit
einem **Nachfolge-Eintrag** zu [`MR-014`](../../../../harness/conventions.md#mr-014) zu führen, der die
Lokal/Draußen-Abgrenzung explizit macht — **Maintainer-Vorbehalt**: die
Freigabe der Erfassung selbst ist die des Maintainers.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Das Auswerten der erfassten Daten.** *Bestand bleibt bewusst stehen*: erst
  sammeln, dann auswerten — ein Auswertungs-Slice braucht echte Daten.

## 2. Definition of Done

- [ ] Wrapper, Traeger-Handling und Feldliste sind übernommen; die
      Hook-Verdrahtung ist in `.claude/settings.json` (span-emit auf
      PostToolUse, PostToolUseFailure, SubagentStart, je Exit 0).
- [ ] Die [`MR-014`](../../../../harness/conventions.md#mr-014)-Abgrenzung ist als Nachfolge-Eintrag oder Zusatz explizit.
- [ ] Die Gegenprobe: ein Testlauf mit dem Traeger erzeugt Erfassungs-Zeilen,
      ein Lauf ohne Traeger bleibt stumm und bricht nicht (Wrapper-Exit 0).
- [ ] Unabhängiger Review, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag; Register fortgeschritten;
      jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `.claude/hooks/span-emit.sh` | neu | Der committete Wrapper (aus der Generator-Ausgabe) |
| `harness/erfassung-feldliste.md` | neu | Werkzeug-erzeugt, kanonisch |
| `.claude/settings.json` | update | Hook-Verdrahtung span-emit |
| [`harness/conventions/MR-014-…`](../../../../harness/conventions/MR-014-keine-agenten-telemetrie.md) | Nachfolge-Eintrag | [`MR-014`](../../../../harness/conventions.md#mr-014)-Abgrenzung explizit |

## 4. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei **und** der
Maintainer hat die Erfassung freigegeben ([`MR-014`](../../../../harness/conventions.md#mr-014)-Vorbehalt).

**Rückführungen — vorab benennen:**

- `in-progress` → `next` (zu groß): entfällt — Übernahme einer fertigen
  Schicht.
- `in-progress` → `open` (blockiert): ohne Maintainer-Freigabe bleibt der
  Slice offen — der Vorbehalt blockiert.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 6. Risiken und offene Punkte

- **Der Traeger ist ein Binärartikel** — Herkunft und Integrität sind zu
  belegen (Herkunft: `.harness/state/bin/` des Generators).
  — **Ausgang:** bei Closure.

## 7. Closure-Notiz

*(wird vor dem `git mv` nach `done/` gefüllt)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** GATE (Achsen 1, 2, 3 ✓) und
HARNESS (Achsen 1, 2, 3 ✓) — die Änderungen unter `harness/` (Feldliste,
MR-014-Eintrag) liegen in der HARNESS-Pfad-Familie.

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde durchgegangen
(2026-09-29): keine Treffer in GATE für diesen Vorgang.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
