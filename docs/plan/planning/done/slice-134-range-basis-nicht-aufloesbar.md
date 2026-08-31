# slice-134 — Eine Range-Basis, die es im Klon nicht gibt, ist kein Traceability-Befund

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Beide Dependabot-PRs des Kanals aus
[slice-128](../done/slice-128-dependabot-hebungskanal.md)
([ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md)) tragen einen roten `push`-Lauf —
aus einem Grund, der mit dem Geprüften nichts zu tun hat.

**Berührte Spec-Stellen:** — *(keine)* — die CI-Verdrahtung ist nicht Gegenstand des Lastenhefts.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-31.

---

## 1. Ziel

Der Traceability-Schritt misst die Commits, die er messen soll — auch wenn der Branch dazwischen
rebast wurde.

## 2. Definition of Done

- [x] `tools/ci-commit-range.sh` ermittelt die Commit-Range aus den drei Fällen und behandelt eine
      **nicht erreichbare** Basis wie einen neuen Branch; es feuert bei jedem Lauf seinen
      Selbsttest.
- [x] [`ci.yml`](../../../../.github/workflows/ci.yml) ruft das Skript statt der Inline-Weiche —
      die Range-Logik steht damit an **einem** Ort und ist lokal ausführbar.
- [x] `make ci-range-selftest` hängt im `gates`-Aggregat und ist in
      [`AGENTS.md`](../../../../AGENTS.md) §4 sowie
      [`harness/README.md`](../../../../harness/README.md) deklariert.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Der Vorfall, gemessen

Beide PRs des Dependabot-Kanals tragen zwei Läufe: `pull_request` **grün**, `push` **rot** nach
7–14 Sekunden. Der rote bricht im Schritt *Traceability + ADR-Immutabilität (Commit-Range)* ab:

```text
Prüfe Commit-Range: 9dfc858b81…..e33d2734c5…
d-check: error: Range-Basis "9dfc858b81…" nicht auflösbar: reference not found
make: *** [Makefile:158: trace-check] Error 2
```

Die Basis stammt aus `github.event.before`. **Nachgeschlagen, nicht vermutet:** beide Basen
(`9dfc858b…` für PR #1, `f944e762…` für PR #2) tragen dieselbe Commit-Message wie der jeweilige
Branch-Head — `build(ci) […]: bump …`, mit dem im Kopf verlinkten ADR-Präfix — bei anderem SHA,
und keine von beiden liegt auf `main`. Das ist ein **Force-Push**: Dependabot rebast seinen Branch
auf neueres `main` und schiebt neu.

Der alte Commit bleibt serverseitig auflösbar — die API gibt ihn her —, ist aber von keiner Ref
mehr erreichbar. `actions/checkout` holt ihn deshalb nicht in den Runner-Klon, und `fetch-depth: 0`
ändert daran nichts: die Tiefe betrifft die Historie der geholten Ref, nicht verwaiste Objekte.

## 4. „Nicht auflösbar" ist derselbe Fall wie „neuer Branch"

Die Weiche in [`ci.yml`](../../../../.github/workflows/ci.yml) fängt **einen** der beiden Fälle:

```bash
elif [[ "$PUSH_BEFORE" =~ ^0+$ ]]; then   # neuer Branch
```

| Lage | `github.event.before` | heute | richtig |
|---|---|---|---|
| neuer Branch | all-zeros | Fallback auf `main` | Fallback auf `main` |
| Force-Push | gültig aussehender SHA, im Klon nicht erreichbar | **Exit 2** | Fallback auf `main` |
| normaler Push | erreichbarer SHA | Range benutzen | Range benutzen |

Beide oberen Zeilen bedeuten dasselbe: *keine brauchbare Basis*. Die Antwort darauf steht bereits
im Kommentar der bestehenden Weiche — *„gegen den Default-Branch prüfen statt nur HEAD (sonst
Silent-Green)"* —, sie war nur an das falsche Merkmal gebunden. Ein
`git cat-file -e "$PUSH_BEFORE^{commit}"` fragt nach dem Merkmal, auf das es ankommt:
**Erreichbarkeit im Klon**, nicht Form des Strings.

Die Richtung ist die sichere: der Fallback misst **mehr** Commits, nie weniger. Ein Silent-Green
ist damit ausgeschlossen — genau das, was die ursprüngliche Weiche wollte.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz geschrieben.

**Abgrenzung:** Dieser Slice repariert die **Range-Ermittlung**. Er ändert nichts an den drei
Prüfungen, die auf ihr laufen (`trace-check`, `commit-scope-check`, `doc-immutable`), und nichts an
der Frage, ob ein roter Lauf einen Branch-Schutz auslösen soll — das ist eine
Repository-Einstellung außerhalb des Repos.

## 6. Risiken und offene Punkte

- **Der Selbsttest kann die Semantik von `github.event.before` nicht prüfen.** Er baut sein eigenes
  Repo und misst das Skript; *welchen Wert GitHub in das Feld schreibt*, erfährt nur ein echter
  Lauf — und genau diese Annahme war der Defekt.
  **Ausgang:** *weiter offen* → Beobachtungs-Register, neue Kennung.
- **Der Fallback misst nach einem Force-Push mehr Commits, als der Push gebracht hat.** Ein
  Mensch, der fünf Commits rebast und neu schiebt, wird gegen `main` gemessen.
  **Ausgang:** *entfallen*, **gestrichen mit Begründung**: mehr zu messen ist die gewollte
  Richtung (§4). Weniger zu messen wäre das Silent-Green, gegen das die Weiche gebaut wurde.
- **Die Logik wandert aus dem Workflow in ein Skript** und ist damit an zwei Orten zu lesen.
  **Ausgang:** *entfallen*, **gestrichen mit Begründung**: sie steht danach an **einem** Ort statt
  inline im YAML; der Workflow ruft nur noch. Genau das macht sie lokal ausführbar und den
  Selbsttest überhaupt möglich.

## 7. Closure-Notiz

**Geliefert:** `tools/ci-commit-range.sh` trägt die Weiche samt Selbsttest über vier Fälle;
[`ci.yml`](../../../../.github/workflows/ci.yml) ruft es statt der Inline-Fassung;
`make ci-range-selftest` hängt im `gates`-Aggregat und steht in
[`AGENTS.md`](../../../../AGENTS.md) §4, in
[`harness/README.md`](../../../../harness/README.md) und in der GATES-Liste des
PreToolUse-Command-Guard.

**Lerneintrag — Form: geschärfte Regel.** *Eine Fallunterscheidung, die man nicht ausführen kann,
ist keine Logik, sondern eine Vermutung mit Einrückung.* Die Weiche stand vier Zeilen im YAML und
sah vollständig aus; sie deckte zwei von drei Lagen. Aufgefallen ist das nicht durch Lesen — sie
wurde mehrfach gelesen —, sondern erst durch einen fremden Bot, der die dritte Lage erzeugte.
*Weil* der `run:`-Block eines Workflows nur auf dem Runner läuft, gab es für ihn weder Selbsttest
noch Gate; das Modul `workflows` aus [slice-130](../done/slice-130-workflows-modul-uses-form.md)
prüft die **Deklarations**-Form der `uses:`-Einträge und sieht in die Blöcke ausdrücklich nicht
hinein. Die Regel, die daraus folgt: Ablauf-Logik gehört aus dem Workflow heraus in ein Skript,
sobald sie mehr als einen Fall unterscheidet — nicht wegen der Lesbarkeit, sondern weil sie erst
dort einen Wächter bekommen kann.

**Die Diagnose wäre fast die falsche gewesen.** Der erste Verdacht war „der Commit ist nicht
auffindbar, also hat Dependabot ihn weggedrückt". Die API gab ihn aber her — der Commit
**existiert**. Erst die Frage *wo* er nicht existiert (im Runner-Klon, nicht auf dem Server) traf
den Sachverhalt, und mit ihr das richtige Prädikat: `git cat-file -e` fragt nach Erreichbarkeit,
nicht nach Existenz. Ein Fix auf die erste Diagnose hin hätte die Basis erneut serverseitig geholt
und wäre grün geworden, ohne die Lage zu treffen.

**Drei beobachtbare Closure-Kriterien:**

1. Der Selbsttest fängt den **tatsächlichen** Defekt: gegen die Fassung vor der Reparatur meldet
   er *„unerreichbare Basis (Force-Push) nicht abgefangen"*, Exit 2.
2. Er fängt auch die Gegenrichtung: gegen eine Fassung, die **immer** auf den Default-Branch
   fällt, meldet er *„erreichbare Basis wird nicht benutzt"*, Exit 2. Ohne diesen vierten Fall
   wäre ein Skript, das nie die Range benutzt, von einem richtigen nicht zu unterscheiden.
3. `make gates` grün mit dem neuen Target im Aggregat — `ci-range-selftest ok: vier Faelle`.

**Beobachtungs-Register (`../observations.md`):** neue [`BEO-033`](../observations.md) angelegt
(CI-Schicht, 1×, Beleg `slice-134`) — eine Weiche, deren Verhalten an der Semantik eines **fremden**
Event-Feldes hängt, ist lokal nicht prüfbar; der Selbsttest misst das eigene Skript, nicht die
Annahme darüber. [`BEO-030`](../observations.md) und [`BEO-023`](../observations.md) wurden
gesichtet und **nicht** erhöht (§8).

**Folge-Slices:** keiner. Ob ein roter Branch-Lauf einen Schutz auslösen soll, ist eine
Repository-Einstellung außerhalb des Repos (§5).

**Risiken aus §6:** drei, jedes mit genau einem Ausgang — *weiter offen* → Beobachtungs-Register
([`BEO-033`](../observations.md)); *entfallen*, gestrichen mit Begründung (der Fallback misst mehr,
nie weniger); *entfallen*, gestrichen mit Begründung (die Logik steht danach an **einem** Ort statt
inline).

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** Berührt ist die **CI-Schicht** (Workflow-Verdrahtung und
das Skript dahinter). Schwelle ≥ 2 von 3 Achsen erfüllt: eigene Artefakt-Menge
(`.github/workflows/`, `tools/`), eigene Review-Frage (*misst der Schritt, was er zu messen
behauptet?*), eigener Modus-Eintrag in
[`harness/conventions.md`](../../../../harness/conventions.md). Die Gate-/Werkzeug-Schicht wird
**mitgezogen** (ein Target, zwei Deklarations-Zeilen), ist aber keine zweite Sub-Area: ohne die
CI-Änderung gäbe es das Target nicht.

**Vorgelagert — offene Beobachtungen sichten:** Ein Treffer.
[`BEO-030`](../observations.md) (*die Wirkung des Hebungs-Kanals hängt an zwei Schaltern außerhalb
des Repos*, **1×**) betrifft denselben Kanal, aber eine andere Frage — dort geht es um
Schalter, die **kein Gate** sehen kann, hier um eine Weiche, die das Repo selbst schreibt. Kein
Zähler-Anstieg. [`BEO-023`](../observations.md) (*Prüfer meldet grün ohne Gegenstand*, **2×**)
trifft **nicht**: dieser Prüfer meldet **rot ohne Ursache**, die Gegenrichtung; sie als denselben
Eintrag zu zählen hieße, zwei Fehlerbilder unter einer Kennung zu vermischen.

### Sub-Area: CI-Schicht

- **Modus:** GF — die Workflows entstanden nach der Konvention, nicht vor ihr.
- **Konventionen-Dichte:** mittel. `make ci` ist als CI-Äquivalent deklariert
  ([`AGENTS.md`](../../../../AGENTS.md) §4), und die `uses:`-Form prüft seit
  [slice-130](../done/slice-130-workflows-modul-uses-form.md) das Modul `workflows`. **Ungedeckt
  war bisher die Ablauf-Logik *in* den `run:`-Blöcken** — genau dort sitzt der Defekt.
- **Phase-Reife:** Phase 4 — vier Workflows, gepinnte Referenzen, ein Release-Pfad mit
  fail-closed Spiegel.
- **Evidenz-/Diskrepanz-Risiko:** niedrig für die Reparatur (der Vorfall liegt als Log vor, §3),
  **hoch für die Zusage dahinter**: dass die Weiche nun alle Fälle deckt, ist eine Aussage über
  GitHubs Event-Semantik, die lokal nicht messbar ist — als Risiko in §6 notiert und ins Register
  gegeben.
- **Reconciliation-Aufwand:** keiner. Ein Ort, eine Weiche; kein Bestand wird nachgezogen.
