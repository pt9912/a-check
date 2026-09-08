# slice-180 — `slice-mv` lernt die dritte Verweis-Form

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (der Closure-Trigger wäre die eigene DoD — kein
repo-weites Mehr).

**Bezug:** Lese-Schritt der Closure von
[welle-15](../done/welle-15/welle-15-regelwerk-v650-migration.md) —
[`BEO-PLAN/verweis-auf-wandernden-slice`](../observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md)
steht bei **7×**, und die Verkörperung hat eine Lücke, die allein in dieser
Welle **zehnmal** aufgetreten ist.

**Berührte Spec-Stellen:** — *(keine)* — Werkzeug ohne Vertragsberührung.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-07.

---

## 1. Ziel und Abgrenzung

`make slice-mv` zieht auch die dritte im Bestand vorkommende Verweis-Form nach:
`<lifecycle-verzeichnis>/slice-NNN-….md` **ohne** `../`-Präfix, wie sie eine
flach unter `docs/plan/planning/` liegende Welle-Datei schreibt.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Verweise *in* der wandernden Datei** — es wäre ein anderer Vorgang und
  ist bereits gedeckt: `doc-check`s `links.resolve-from` trägt die
  Gegenrichtung.
- **Ein Sensor auf die Verweis-Form selbst** — Bestand bleibt bewusst stehen:
  `doc-check` meldet den toten Verweis bereits nach dem `mv`; was fehlt, ist
  nicht die Meldung, sondern das Nachziehen.

## 2. Analyse (vor der Umsetzung)

**Gemessen über welle-15:** Bei **jedem** der zehn Lifecycle-Wechsel ihrer sechs
Slices ließ das Werkzeug den Verweis in der flach liegenden Welle-Datei stehen —
`[slice-NNN](open/…)` bzw. `(in-progress/…)`. Jedes Mal meldete `make doc-check`
`target-missing`, jedes Mal war die Korrektur ein `sed`.

**Warum die Form vorher nicht vorkam:** `welle-14` lag bei den Übergängen ihrer
Slices bereits in `done/` — ihre Verweise trugen `../done/…`. Wellenlose Slices
haben gar keine Welle-Datei, die auf sie zeigt. Erst eine **offene** Welle-Datei,
die auf Slices in Lifecycle-Verzeichnissen zeigt, erzeugt die dritte Form; sie
ist damit dieselbe Klasse wie
[`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md):
kalibriert an den Vorkommen, die es beim Bauen gab.

**Zu klären vor dem Umbau:** ob `tools/slice-mv.sh` die Form über eine dritte
Regel ergänzt oder ob eine Verallgemeinerung beide bestehenden ablöst — die
zweite Variante ist kürzer und riskanter.

**Beantwortet, und die Antwort stand im Werkzeug selbst.** Der Selbsttest
führt seit slice-118 eine **Gegenprobe**, die genau diese Form ausdrücklich
schützt:

```
echo 'open/slice-999-x.md'      # keine der zwei Formen
...
if ! grep -qx 'open/slice-999-x.md' "$f"; then
    echo "... praefixlose Nicht-Form mitgeaendert"
```

Eine Verallgemeinerung (`s|$3/$2|$4/$2|g`) bricht sie sofort — gemessen, siehe
§3, Probe (c). Der nackte Pfad im Fließtext ist eben **keine** Verweis-Form: Er
kann alles sein, und ein Werkzeug, das ihn umschreibt, ist von einem
kaputten nicht zu unterscheiden.

**Also: dritte Regel, aber am Kontext festgemacht statt am Pfad.** Die Form
wird an ihren zwei Ankern erkannt — als Markdown-Link-Ziel `](open/slice-…)`
oder in Inline-Code `` `open/slice-…` ``. Beide Anker schließen die ersten
zwei Formen von selbst aus: dort steht vor dem Verzeichnis ein `/` (aus `../`
bzw. dem repo-relativen Pfad). Und beide lassen den nackten Pfad in Ruhe.

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `tools/slice-mv.sh`, `rewrite_file()` | update | zwei `sed`-Regeln für die zwei Anker der dritten Form |
| `tools/slice-mv.sh`, `self_test()` | update | zwei Positiv-Fälle, eine Negativ-Gegenprobe geschärft |
| `tools/slice-mv.sh`, Usage-Block | update | nennt jetzt alle drei Formen |

**Vier Mutations-Proben auf Kopien** — jede in der Richtung, die sie belegen
soll:

| Mutation | Erwartet | Gemessen |
|---|---|---|
| unverändert | grün | Selbsttest läuft durch, 0 Fehler |
| Link-Regel entfernt | rot | *„Form 3 als Link nicht ersetzt"* |
| Inline-Code-Regel entfernt | rot | *„Form 3 in Inline-Code nicht ersetzt"* |
| Regel **zu weit** (`s\|$3/$2\|$4/$2\|g`) | rot | *„nackter Pfad mitgeaendert"* |

Die letzte Zeile ist die wichtigere Hälfte: Ohne sie wäre eine Ersetzung, die
jedes `open/` trifft, von einer kontext-gebundenen nicht zu unterscheiden.

**End-to-End auf einem Klon** — weil das Repo gerade keine offene Welle führt
und die Form ohne sie nicht vorkommt: flache Datei `welle-99-probe.md` mit
allen drei Formen angelegt, `slice-mv slice-181 in-progress` gefahren. Alle
drei umgeschrieben, das Ziel existiert.
**Eine Beobachtung aus der Probe, die nicht zum Erfolg zählt:** Der
Geschwister-Verweis `../in-progress/…` löst aus einer **flach** liegenden Datei
nicht auf — `../` führt aus `planning/` heraus. Das Werkzeug hat ihn korrekt
ersetzt; die Form gehört nur nicht in eine flache Datei. Die Probe war an
dieser Stelle unrealistisch, nicht das Werkzeug falsch.

## 4. Definition of Done

- [ ] `make slice-mv` zieht die dritte Form nach; ein Lifecycle-Wechsel eines
      Slice mit `**Welle:**`-Feld lässt `make doc-check` grün, **ohne**
      Handgriff danach.
- [ ] Der Selbsttest des Werkzeugs deckt alle **drei** Formen ab, je Richtung
      eine Mutations-Probe.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe und WIP-Limit frei.
Dringlichkeit steigt mit der nächsten offenen Welle — ohne sie kommt die Form
nicht vor.

**Rückführungen:** zeigt sich, dass eine Verallgemeinerung die zwei bestehenden
Formen bricht, zurück nach `next/` und getrennt schneiden.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz geschrieben.
Danach Archivierung als wellenloser Slice ([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- *Die dritte Regel deckt eine vierte nicht, die beim nächsten Betriebsmodus
  entsteht — dieselbe Kalibrierungs-Lücke eine Ebene höher* — **Ausgang:**
  *weiter offen* → Beobachtungs-Register,
  [`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md).
  Der Slice verkleinert das Risiko, statt es zu schließen: Die dritte Regel ist
  am **Kontext** festgemacht (Link-Ziel, Inline-Code) statt an einer
  Pfad-Schreibweise, und beide Anker sind sprach-unabhängig. Eine vierte Form
  bräuchte einen vierten Anker — aber sie bräuchte auch einen Betriebsmodus,
  den es heute nicht gibt.

- *Eine einfrierende Welle-Ergebnisnotiz zitiert einen Slice als **Adresse**
  statt als Kennung, und `slice-mv` muss sie deshalb anfassen.* Gefunden beim
  Übergang dieses Slice:
  [`welle-15-results.md`](../done/welle-15-results.md) entstand **nach** der
  Zitier-Form-Regel ([`AGENTS.md`](../../../../AGENTS.md) §5, `seit slice-176`)
  und führt sieben nackte `slice-NNN`-Kennungen — an zwei Stellen aber einen
  Link auf den Lifecycle-Pfad. — **Ausgang:** *weiter offen* →
  Beobachtungs-Register,
  [`BEO-PLAN/slice-mv-fasst-einfrierendes-artefakt-an`](../observations/BEO-PLAN/slice-mv-fasst-einfrierendes-artefakt-an/observation.md).
  Er gehört dorthin, weil er dieselbe Ursache hat: Das Werkzeug kennt die
  Verweis-Form, aber nicht die Klasse des Dokuments. **Nicht in diesem Slice
  behoben** — die Zitier-Form auf ein eingefrorenes Artefakt anzuwenden ist ein
  eigener Vorgang, und §1 grenzt diesen hier auf das Werkzeug ab.

## 8. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice;
Lerneintrag — Form: wird dort benannt.)_

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Gate-/Werkzeug-Schicht** `GATE` (`tools/slice-mv.sh`), Achsen 1,2,3 laut
[`harness/conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area).
Der Harness-Einstieg ist **nicht** berührt: `AGENTS.md` §4 und
[`harness/README.md`](../../../../harness/README.md) beschreiben `slice-mv`
bereits als „samt der Verweise auf ihn" — die Aussage bleibt wahr und ändert
sich nicht.

**Vorgelagert — offene Beobachtungen sichten:** gesichtet am 2026-09-08 gegen
den gemergten Stand. Drei Treffer in `PLAN`/`GATE`:

- [`BEO-PLAN/verweis-auf-wandernden-slice`](../observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md)
  — **7×**, der Eintrag, den dieser Slice bedient. Er steht auf *verkörpert mit
  benannter Lücke*; der Slice schließt die Lücke.
- [`BEO-PLAN/slice-mv-fasst-einfrierendes-artefakt-an`](../observations/BEO-PLAN/slice-mv-fasst-einfrierendes-artefakt-an/observation.md)
  — **2×**, dasselbe Werkzeug, andere Richtung: dort zieht es Verweise nach, wo
  es nicht darf. **Erreicht mit diesem Slice die Schwelle nicht** — der
  `slice-mv` beim Übergang nach `in-progress/` fasste zwar erneut ein
  einfrierendes Artefakt an ([`welle-15-results.md`](../done/welle-15-results.md)),
  aber dort **durfte** er: der Verweis ist ein Markdown-Link, der sonst bräche.
  Das ist ein anderer Fall als die zwei Review-Reports und **kein** dritter
  Beleg. Was dort stattdessen auffiel, steht in §7 als offener Punkt.
- [`BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise`](../observations/BEO-GATE/muster-trifft-nur-die-haeufige-schreibweise/observation.md)
  — der Eintrag, dessen Klasse §2 dem eigenen Befund zuordnet. Sein Zähler wird
  von diesem Slice **nicht** erhöht: Der Fund ist derselbe Vorgang, den
  `verweis-auf-wandernden-slice` schon zählt.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
