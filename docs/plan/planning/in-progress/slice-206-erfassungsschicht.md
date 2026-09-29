# slice-206 — Erfassungsschicht übernehmen (mit MR-014-Abgrenzung)

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt. Er wechselt nur durch `git mv`.

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Folge-Slice aus
[slice-205](../done/wellenlos/slice-205-init-tool-vergleich.md) (Entscheidung
„Übernehmen"); [`MR-014`](../../../../harness/conventions.md#mr-014)
(Abgrenzung nötig).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** — bis zur Priorisierung.

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** geschärfte Regel.

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

- [x] Wrapper, Traeger-Handling und Feldliste sind übernommen; die
      Hook-Verdrahtung ist in `.claude/settings.json` (span-emit auf
      PostToolUse, PostToolUseFailure, SubagentStart, je Exit 0).
- [x] Die [`MR-014`](../../../../harness/conventions.md#mr-014)-Abgrenzung ist als Nachfolge-Eintrag oder Zusatz explizit.
- [x] Die Gegenprobe: ein Testlauf mit dem Traeger erzeugt Erfassungs-Zeilen,
      ein Lauf ohne Traeger bleibt stumm und bricht nicht (Wrapper-Exit 0).
- [x] Unabhängiger Review, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag; Register fortgeschritten;
      jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `.claude/hooks/span-emit.sh` | neu | Der committete Wrapper (aus der Generator-Ausgabe) |
| `harness/erfassung-feldliste.md` | neu | Werkzeug-erzeugt, kanonisch |
| `.claude/settings.json` | update | Hook-Verdrahtung span-emit |
| [`harness/conventions/MR-030-…`](../../../../harness/conventions/MR-030-erfassung-lokal-kein-abfluss.md) | Nachfolge-Eintrag ([`MR-030`](../../../../harness/conventions/MR-030-erfassung-lokal-kein-abfluss.md)) | [`MR-014`](../../../../harness/conventions.md#mr-014)-Abgrenzung explizit |

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
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: Herkunft und
  Integrität sind belegt — Herkunft `/tmp/aih-v6.13.0/.harness/state/bin/`
  (Generator-Ausgabe `v6.13.0`), sha256
  `c6a6a171bef7c9eeb500ddc89ef6bfaf9ca52079454eabf450e94923e0ce9186`, gelegt
  im Rahmen dieses Slices; der Bestand ist gitignored und trägt darum keinen
  Integritäts-Zwang über den Lauf hinaus.

## 7. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** *Ein MR-Nachfolge-Durchgang umfasst
fünf Handgriffe: Nachfolge-Datei · `git mv` des Vorgängers nach
`conventions/done/` · beide Tabellen · Anker-Erhalt · interne Link-Tiefen.*
Verkörpert in [`harness/conventions.md`](../../../../harness/conventions.md)
§Adaptions-Block (`seit slice-206`). Gemessen: der Durchgang verlor drei
 interne Links und eine Tabellen-Zeile an die neue Tiefe — erst
`doc-check` hielt beide Richtungen gegen die Tabellen.

**Was hat funktioniert:** die Gegenprobe in beiden Richtungen vor der
Übergabe — der Lauf mit Traeger erzeugte die Schema-Zeile (inkl. automatischer
Slice-Ableitung aus dem Lifecycle-Verzeichnis), der Lauf ohne Traeger blieb
stumm mit Exit 0; `doc-check` und `gate-consistency` hielten Verdrahtung und
MR-Tabellen scharf.

**Was ging anders als geplant:** der unabhängige Review fand die Sichtung
unvollständig (F-2 — HARNESS und die 2×-Beobachtung
`mr-aufloesungs-trigger-ohne-waechter` fehlten; der Slice selbst ist ihr
drittes Auftreten) und das fehlende Pflichtfeld „Ausgelöst durch
Baseline-Stand" in [`MR-030`](../../../../harness/conventions.md#mr-030) (F-4). Die Probe der done/-Abgrenzung traf ihren
Gegenstand nicht (F-3) und zwei Formgrößen (F-5, F-6). Alle korrigiert, bevor
die Closure geschrieben wurde.

**Steering-Loop-Eintrag:** siehe Lerneintrag oben (MR-Durchgang-Checkliste).
Die Beobachtungs-Klasse `mr-aufloesungs-trigger-ohne-waechter` erreichte
**3×** — Ausgang *geplant* →
[slice-208](../open/slice-208-trigger-audit-mr-eintraege.md) (die „dass"-Hälfte
mechanisiert: eine Closure belegt die Sichtung der aktiven Einträge).

**Beobachtungs-Register (`../observations/`):**
`mr-aufloesungs-trigger-ohne-waechter` → 3. Auflage belegt, Ausgang *geplant*
([evidence/slice-206.md](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/evidence/slice-206.md));
GATE 15 offene Einträge, keiner betrifft diesen Vorgang.

**Folge-Slices:** [slice-208](../open/slice-208-trigger-audit-mr-eintraege.md)
(Trigger-Audit für MR-Einträge mechanisieren) — ist Datei in `open/`.

**Risiken aus §6:** das eine Risiko trägt seinen Ausgang (*entfallen*, mit
Begründung — Herkunft und sha256 des Traegers sind belegt).

**Drei Paarungen:** Anker — verkörpert (MR-Durchgang-Checkliste,
`seit slice-206`) · Folge-Slice — slice-208 existiert als Datei in `open/` ·
Register — mr-aufloesungs-trigger (3×, geplant) · attestierung/probe-liefert
durch slice-204/207 bereits besetzt.

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** GATE (Achsen 1, 2, 3 ✓) und
HARNESS (Achsen 1, 2, 3 ✓) — die Änderungen unter `harness/` (Feldliste,
[`MR-014`](../../../../harness/conventions.md#mr-014)-Eintrag) liegen in der HARNESS-Pfad-Familie.

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde
durchgegangen (2026-09-29): GATE trägt 15 offene Einträge (Zählung über die
`state.md`-Köpfe), keiner betrifft die Sub-Area-Berührung dieses Vorgangs.
HARNESS trägt [`mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md)
bei 2× — und **dieser Slice ist ihr drittes Auftreten**: die
[`MR-014`](../../../../harness/conventions.md#mr-014)-Auflösung
vollzog genau jenen Gegenstand. Der Ausgang ist im Lese-Schritt der Closure
gelegt — *geplant* →
[slice-208](../open/slice-208-trigger-audit-mr-eintraege.md).

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
