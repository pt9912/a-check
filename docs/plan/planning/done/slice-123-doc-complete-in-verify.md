# slice-123 — `doc-complete` prüft vor dem Abschluss, nicht danach

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Wort 2026-08-30. Anlass ist eine gemessene Requirements-Waise:
[AC-FA-CLI-003](../../../../spec/lastenheft.md#ac-fa-cli-003--usage-ausgabe-und-handbuch-verweis)
wird von **keinem** Slice genannt.

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Eine Anforderung ohne referenzierenden Slice fällt beim Abschluss auf, nicht Wochen später: das
Vollständigkeits-Gate hängt an `verify`.

## 2. Definition of Done

- [x] Die bestehende Waise ist behoben — der Slice, der die Anforderung umgesetzt hat, nennt ihre
      Kennung in der Closure-Notiz. Beleg: `make doc-complete` meldet **0 Waisen**, Exit 0.
- [x] `doc-complete` hängt im `verify`-Aggregat; [`AGENTS.md`](../../../../AGENTS.md) §4 und
      [`harness/README.md`](../../../../harness/README.md) §Sensors führen es nicht mehr als
      **advisory**, sondern mit seiner Bindung.
- [x] Die **Ursache** ist als Regel benannt, nicht nur der Einzelfall: wer eine Anforderung
      anlegt, kann ihre Kennung im Plan nicht verlinken (sie existiert noch nicht) — **in der
      Closure-Notiz existiert sie**. Das steht im Workflow-Skelett und in
      [`AGENTS.md`](../../../../AGENTS.md) §5.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `done/slice-117-…md` | update | die Kennung in der Closure-Notiz — behebt die Waise |
| [`Makefile`](../../../../Makefile) | update | `doc-complete` ins `verify`-Aggregat |
| [`AGENTS.md`](../../../../AGENTS.md) §4/§5, [`harness/README.md`](../../../../harness/README.md) | update | Bindung statt „advisory"; die Regel |
| `.claude/commands/slice.md` | update | Schritt 8 nennt den Zeitpunkt |

**Auszuführende Gates:** `make verify` (neu mit `doc-complete`), `make gates` — tragend
`doc-targets` (beide Doku-Tabellen). 

### Der Zielkonflikt, aus dem die Waise entstand

Er ist strukturell und trifft **jeden** Slice, der eine Anforderung neu anlegt:

1. IDs werden **referenziert, nicht erfunden** ([`AGENTS.md`](../../../../AGENTS.md) §5) — beim
   Planen existiert die neue Kennung noch nicht.
2. Jede genannte Kennung ist **linkpflichtig** (`ids`, `link-policy: always`) — ein Link auf eine
   Anforderung, die es nicht gibt, macht `doc-check` rot.
3. Also umschreibt der Plan sie („eine neue `AC-FA-CLI`-Kennung"), und die RTM sieht den Slice
   nicht.

Der Ausweg braucht keine Regel-Änderung, nur den richtigen **Zeitpunkt**: bei der Closure ist die
Anforderung geschrieben. Dort kann und muss die Kennung stehen.

## 4. Trigger

**Start:** eingetreten — die Waise ist gemessen (`make doc-complete`, Exit 1).

**Rückführungen:**

- `in-progress` → `open`: falls `doc-complete` über den Bestand weitere Waisen meldet, die eine
  inhaltliche Entscheidung verlangen statt einer Referenz.

## 5. Closure-Trigger

Waise behoben, Gate gebunden, Regel benannt, Gates grün.

**Was bewusst nicht getan wird:** die **Linkpflicht lockern** oder Planning-Dateien von `ids`
ausnehmen. Beides löste den Zielkonflikt, indem es eine Zusage aufgibt — die Kennungs-Linkpflicht
ist genau die Regel, die verhindert, dass Kennungen ins Leere zeigen.

## 6. Risiken und offene Punkte

- *Die Regel verlagert Arbeit in die Closure, die dort vergessen werden kann* — **Ausgang:**
  gestrichen mit Begründung: genau deshalb steht sie nicht allein als Guide da. `doc-complete`
  hängt seit diesem Slice im `verify`-Aggregat und läuft beim Abschluss jedes Slice — vergessen
  wird ab jetzt gemeldet, nicht überlesen. Ein Guide **plus** Sensor, nicht ein Guide statt eines.
- *`doc-complete` im `verify`-Aggregat macht jede neue Anforderung sofort abschluss-blockierend* —
  **Ausgang:** gestrichen mit Begründung: das ist die Wirkung, die der Slice herstellt, und sie
  ist wohlfeil zu erfüllen — eine Kennung mit Link in der Closure-Notiz. Wer eine Anforderung
  anlegt und in **keinem** Slice nennt, hat sie nicht belegt; das früh zu sagen ist billiger als
  spät.

## 7. Closure-Notiz

**Geliefert:** Die Waise ist behoben (`make doc-complete`: **20 Anforderungen, 0 Waisen**), das
Vollständigkeits-Gate hängt im `verify`-Aggregat, und die Ursache steht als Regel in
[`AGENTS.md`](../../../../AGENTS.md) §5 und im Workflow-Skelett — nicht nur der Einzelfall.

**Lerneintrag — Form: geschärfte Regel.** *Wo zwei Regeln einen Zielkonflikt erzeugen, ist die
Lösung oft kein Nachgeben, sondern ein anderer **Zeitpunkt**.* Hier standen zwei Zusagen
gegeneinander, beide richtig: IDs werden **referenziert statt erfunden** (im Plan existiert die
neue Kennung nicht), und jede genannte Kennung ist **linkpflichtig** (ein Link ins Leere macht
`doc-check` rot). Der Plan konnte die Anforderung also nur umschreiben — und wurde damit für die
Requirements-Matrix unsichtbar. Beide Regeln zu behalten wirkt nach Widerspruch, bis man fragt,
*wann* die Kennung existiert: **bei der Closure**. Dort ist die Anforderung geschrieben, der Link
löst auf, und die Matrix sieht den Slice. *Weil* die Auflösung am Zeitpunkt hängt und nicht am
Regeltext, wäre jede Lockerung — Linkpflicht aufweichen, Planning-Dateien aus `ids` nehmen — teuer
und unnötig gewesen.

**Drei beobachtbare Closure-Kriterien:**

1. `make doc-complete` meldet **0 Waisen** bei 20 Anforderungen; vor diesem Slice war es 1 Waise
   und Exit 1.
2. Die Prüfung läuft ab jetzt **von selbst**: `doc-complete` steht im `verify`-Aggregat, das jeder
   Slice-Abschluss fährt. Vorher war es advisory und lief nie — aufgefallen ist die Waise nur,
   weil jemand das Target von Hand aufrief.
3. Regel und Sensor greifen ineinander:
   [`AGENTS.md`](../../../../AGENTS.md) §5 sagt **wo** die Kennung hingehört (Closure-Notiz) und
   **warum** sie im Plan nicht stehen kann; `doc-complete` prüft, dass sie irgendwo steht. Keiner
   der beiden allein hätte den Fall gefangen.

**Der Anlass war die eigene Arbeit dieser Sitzung.**
[slice-117](../done/slice-117-handbuch-verweis-cli.md) hat
[AC-FA-CLI-003](../../../../spec/lastenheft.md#ac-fa-cli-003--usage-ausgabe-und-handbuch-verweis)
angelegt, umgesetzt, getestet und im Handbuch dokumentiert — und die Kennung **null** mal genannt.
Die Traceability-Matrix führte sie als Waise, während jede inhaltliche Zusage erfüllt war. Das ist
kein Flüchtigkeitsfehler, sondern die vorhersehbare Folge des Zielkonflikts: **jeder** Slice, der
eine Anforderung anlegt, läuft hinein.

**Offene Risiken und ihr Ausgang:** beide gestrichen mit Begründung.

**Beobachtungs-Register:** [`BEO-023`](../observations.md) auf **2×** erhöht (Beleg slice-123) und
um die zweite Form erweitert: ein Prüfer bleibt nicht nur ohne **Gegenstand** unkalibriert,
sondern auch ohne **Aufruf**. `doc-complete` war beides nicht — es hatte einen Gegenstand und
hätte gemeldet; es lief nur nie.

**Folge-Slices:** ein Sensor für [`BEO-006`](../observations.md) (3×, Schwelle überschritten)
bleibt der nächste; dazu [`BEO-024`](../observations.md) und die Entscheidung zu
[`BEO-025`](../observations.md).
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Gate-/Werkzeug-Schicht** (`Makefile`)
und mit den Deklarations-Tabellen und dem Workflow-Skelett der **Harness-Einstieg**.

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-023`](../observations.md) (ein Prüfer ohne
Gegenstand bleibt unkalibriert) beschreibt dieselbe Familie — hier war der Prüfer nicht ohne
Gegenstand, sondern **ohne Aufruf**: `doc-complete` ist advisory und lief nie.

Alle berührten Sub-Areas GF.
