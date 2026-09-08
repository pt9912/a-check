# Planungs-Harness — Slices und Wellen

Zwei Ebenen: der **Slice** ist die Arbeitseinheit, die **Welle** bündelt Slices zu einem
abschließbaren Vorhaben. Der Slice-Zyklus ist eine Zustandsmaschine über Verzeichnisse
(`open/` → `next/` → `in-progress/` → `done/`, [`AGENTS.md`](../../../AGENTS.md) §5); die Welle
liegt eine Ebene darüber und schließt über eine **Prozedur**, nicht über einen Datei-Übergang.

- Slice-Form: die **vendored Ziel-Form** [`slice.template.md`](../../../.harness/baseline/v6.5.0/templates/docs/plan/planning/slice.template.md) — a-check führt keine
  eigene Kopie. **Beim Kopieren:** Zeile `Lerneintrag — Form: <…>` ergänzen, die vier nicht
  geführten Felder streichen ([`AGENTS.md`](../../../AGENTS.md) §5). Höchstens drei **Liefer**-Punkte,
  höchstens zwei Schichten, benannte Lerneintrag-Form. Gezählt wird nur, was mit dem Umfang
  wächst; Gate-Läufe und Closure-Pflichten zählen nicht mit. **Maschinell geprüft von
  `make verify` sind die Zahl und die benannte Lerneintrag-Form**; „höchstens zwei Schichten" ist
  **Review-Sache** und ausdrücklich kein Gate — was eine Schicht ist, ist eine Ermessensfrage über
  Modul-Grenzen, und ein Zähler darüber wäre Schein-Genauigkeit
  (die Regel-Kommentare in [`.d-check.yml`](../../../.d-check.yml) tragen dieselbe Begründung).

## Lifecycle-Bedeutungen

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md` §Lifecycle als State
Machine — der Zustand ist das Verzeichnis, kein Status-Feld. **Wann** gewechselt wird, steht als
Übergangs-Tabelle in [`AGENTS.md`](../../../AGENTS.md) §5; hier steht, **was** ein Verzeichnis
bedeutet.

| Verzeichnis | Bedeutung |
|---|---|
| `open/` | Geplant, noch nicht priorisiert. Keine Garantie auf Umsetzung. |
| `next/` | Als Nächstes priorisiert; das Feld `Verantwortlich:` im Slice-Kopf ist gesetzt. |
| `in-progress/` | Beansprucht: der `git mv` hierher liegt **vor** der Arbeit. |
| `done/` | DoD erfüllt, Closure-Notiz vorhanden, Gates grün. |

**Jeder Wechsel ist ein reiner `git mv`-Commit** ([`AGENTS.md`](../../../AGENTS.md)
§3.3). Beim Übergang nach `done/` ist die Reihenfolge **umgekehrt**: erst der
Inhalt (DoD-Häkchen, Closure-Notiz), dann der reine `git mv` — die Notiz ist die
Bedingung dafür, dass die Datei nach `done/` darf, nicht ihre Folge.
`make slice-mv` fährt den Move und zieht die Verweise **auf** die Datei nach —
Pfade, keine Aussagen: Ein Zustandssatz daneben (etwa der Ruhe-Marker der
Roadmap) bleibt Sache des Laufs.

## Slices vs. Wellen — zwei Ablagen, dieselbe Regel

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md` §Wann Arbeit eine Welle braucht.

- **Slices** tragen ihren Zustand über das **Verzeichnis**.
- Eine **Welle** ebenso: der Welle-Plan liegt **flach** in `planning/`, solange sie läuft, und
  wandert bei Closure per `git mv` nach `done/`, neben seine Ergebnis-Notiz. Den aktiven Durchlauf
  `open/` → `next/` → `in-progress/` durchläuft er nicht; `done/` ist sein einziges
  Lifecycle-Verzeichnis. **Geplante** Wellen haben keine Datei — sie stehen in der Roadmap.
- Der aktive Durchlauf nimmt ausschließlich **Slices** auf; `done/` archiviert zusätzlich
  abgeschlossene **Nicht-Slice-Records**. Aufgelöste Carveouts wandern **nicht** hierher, sondern
  nach [`docs/plan/carveouts/done/`](../carveouts/README.md).

Flach neben den Lifecycle-Verzeichnissen liegt das **Beobachtungs-Register**
([`observations/`](observations/README.md), Verzeichnisform seit slice-139): je Beobachtung ein
Verzeichnis `BEO-<KUERZEL>/<slug>/`, der Steering-Loop-Zähler abgeleitet aus der Zahl der
Evidence-Dateien, fortgeschrieben bei **jeder** Slice-Closure, unabhängig von Wellen. Ein
`reconciliation.md` führt a-check nicht — es gehört zum Brownfield-Bootstrap, den dieses Repo
nicht hatte.

## Beim Kopieren der Slice-Ziel-Form

Umgezogen aus [`AGENTS.md`](../../../AGENTS.md) §5 mit slice-186 — eine
Bedienungsanleitung für die Vorlage gehört dorthin, wo die Ablage beschrieben
ist, nicht ins Briefing. Die Regel selbst ist unverändert.

- **Slice-Form:** neue Slices entstehen aus der **vendored Ziel-Form**
  [`.harness/baseline/v6.5.0/templates/docs/plan/planning/slice.template.md`](../../../.harness/baseline/v6.5.0/templates/docs/plan/planning/slice.template.md) — a-check führt keine eigene Kopie, sie würde gegen die Baseline driften.
  **Beim Kopieren anzupassen** — sieben Punkte, jeder gegen den Bestand gemessen (slice-178):

  1. Die Zeile `Lerneintrag — Form: <…>` **ergänzen** — die Ziel-Form kennt sie nicht als Feld,
     `make verify` verlangt sie.
  2. **Ein** Feld streichen: das *Reconciliation-Register* — a-check hat keinen
     Brownfield-Bootstrap, `reconciliation.md` existiert nicht. **Alles andere bleibt**, auch was
     frühere Fassungen dieser Liste zum Streichen empfahlen: `**Welle:**` führen **174** Dateien,
     den **Herkunfts-Anker** (`— liegt in <Zielort>` plus `seit slice-<NNN>` dort) führt das Repo
     an **58** Stellen, und die *drei Paarungen* stehen in jeder Closure-Notiz.
  3. **Wer die drei Paarungen trägt, entscheidet der Repo-Zustand — nicht das `Welle:`-Feld des
     Slice.** Baseline `modul-06-roadmap.md` §Wann Arbeit eine Welle braucht: *„Wellenlos ist eine
     Eigenschaft des **Repos** … das Kopf-Feld `**Welle:**` sagt nur, ob dieser Slice in ein Bündel
     gehört; daraus folgt für die Vorgänge unten nichts."* Liegt eine **offene Welle** vor, trägt
     ihre Closure die Paarungen — **auch für Slices ohne Wellen-Zugehörigkeit**. Erst ohne
     Wellen-Betrieb trägt die Slice-Closure sie selbst.
  4. **Die Abschnitts-Nummern der Ziel-Form gelten nicht unverändert.** a-check schiebt zwischen
     Ziel und DoD einen **Analyse-Abschnitt** ein (er hält die Messung, auf der die DoD steht) und
     stellt die DoD hinter den Umsetzungs-Abschnitt. Wieviel sich dadurch verschiebt, hängt vom
     Slice ab — der Bestand führt den Sub-Area-Abschnitt als §7 bis §10. **Nicht die Nummer
     kopieren, sondern die Reihenfolge der Ziel-Form einhalten und den Analyse-Abschnitt
     einfügen, wo er gebraucht wird.**
  5. **§1 heißt *Ziel und Abgrenzung*** — die Ziel-Form nennt ihn anders. Was er trägt
     (Begründungs-Pflicht je Punkt, die vier Ausschluss-Klassen als Suchraster, keine
     Mindestzahl), steht in `modul-05` §Ziel-Form: Slice.
  6. **Der Sub-Area-Abschnitt heißt *Sub-Area-Prüfungen und Modus-Begründung*** — der Titel trägt beide Hälften, weil
     nur die zweite bedingt ist. Die zwei *Vorgelagert*-Blöcke (Sub-Area-Wahl prüfen · offene
     Beobachtungen sichten) laufen in **jedem** Slice-Plan, unabhängig von Modus und Slice-Typ;
     **der Abschnitt entfällt nie**.
  7. Die Ziel-Form führt eine **Review-DoD-Zeile**
     ([`MR-019`](../../../harness/conventions.md#mr-019)) — ihr Wortlaut wird beim Kopieren auf die exakte
     Trigger-Phrase „unabhängiger Review" umgeschrieben, statt den Baseline-Wortlaut unverändert
     zu übernehmen; sonst prüft `make doc-reviews` sie nie (empirisch geprüft, slice-165 §3/§4).

  **Größen-Regel, Zählregel und die drei Lerneintrag-Formen stehen in `modul-05`
  §Ziel-Form: Slice.** Repo-eigen ist ihre **Durchsetzung**: `make verify` prüft beides ab
  slice-052, ältere Slices sind grandfathered. Der Gate-Lauf steht als feste Zeile unter dem
  DoD — als Checkbox ist er ab slice-098 ein Befund, weil er zu den konstanten Posten zählt.
  Ab demselben Stichtag trägt der Kopf `Verantwortlich:`, `Autor:` und die berührten
  Spec-Stellen; `—` ist eine gültige Antwort, Schweigen nicht.
## Aktueller Stand

Regeln dieser Sektion: Baseline-Regelwerk `modul-05-planning-harness.md` §Lifecycle als State
Machine — kein Snapshot; der Stand ergibt sich aus den Verzeichnissen.

**Hier steht bewusst keine Tabelle.** Ein Snapshot des Stands driftet gegen die Verzeichnisse,
sobald ein `git mv` läuft, und niemand merkt es — der Stand ist `ls` über `open/`, `next/`,
`in-progress/`, `done/`.

## Roadmap

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md` §Roadmap-Struktur.

Sequenzierungs-Autorität ist [`in-progress/roadmap.md`](in-progress/roadmap.md) — auch nach einer
Wellen-Closure.

## Wellen-Closure-Prozedur

Quelle: `modul-06` §Wellen-Closure. **Sechs Schritte, jeder mit einem Beleg — keiner mit einem
Datum.** Erst wenn alle sechs Belege vorliegen, ist eine Welle *auditierbar* geschlossen.

Schritt 4 — **Zeitdokumente archivieren** — ist in a-check **verkörpert**: `make archive-wave`
bewegt den Volltext ins Archiv und lässt einen Stub zurück; für einen wellenlosen Slice ist es
ein Pflichtschritt der Closure ([`AGENTS.md`](../../../AGENTS.md) §6). Der Bestand trägt das:
`done/welle-*/archiv.zip` gibt es für **13** geschlossene Wellen, und `done/wellenlos/` führt je
Slice ein eigenes.

**1 — Trigger prüfen.** Alle Slices der Welle liegen in `done/`, und der Lauf ist grün. Das ist
die beobachtbare Bedingung, nicht der Kalendertag.
*Beleg:* `make ci` mit Exit 0 (Ausgabe in eine Datei, Exit-Code getrennt geprüft).
**Repo-spezifische Ersetzung:** die Baseline verlangt hier zusätzlich einen *Replay-Lauf*. a-check
führt keinen — das ist als bewusste Abweichung deklariert
([MR-008](../../../harness/conventions.md#mr-008--kein-replay-keine-agenten-telemetrie)), und das
Kriterium ist dort ausdrücklich durch „`make ci` grün" ersetzt. Ohne diese Zeile bliebe jede
Wellen-Closure unerfüllbar, ohne dass jemand sagen könnte warum.

**2 — Carveout-Audit.** Jeder offene Carveout wird geprüft: aufgelöst, verlängert (mit
Folge-Slice) oder als permanent akzeptiert und über den Trichter in eine ADR überführt. Eine Welle
darf **mit** dokumentiertem Carveout schließen — nie mit einem stillen roten Gate.
*Beleg:* je Carveout ein aktuelles `Letzte Prüfung:`-Datum in
[`docs/plan/carveouts/`](../carveouts/README.md). Solange dort **null** Carveouts liegen, ist der
Schritt mit dem Verweis auf die leere Bestandstabelle erfüllt — das ist eine Aussage, keine
Auslassung.

**3 — Welle schließen.** Eine Ergebnis-Notiz `done/welle-NN-results.md` schreiben: *geliefert · was
funktionierte · was anders lief · **Steering-Loop-Einträge** · Folge-Slices · die Verifikation aus
Schritt 1*. Ohne Lerneintrag ist die Welle nicht „fertig", nur „weg". Zugleich wandert die
Welle-Plan-Datei per `git mv` von flach nach `done/`, neben ihre Ergebnis-Notiz — der Zustand ist
die Verzeichnis-Position, kein `Status`-Feld, wie beim Slice.
*Beleg:* die Notiz existiert und nennt je Punkt etwas Prüfbares.
**Dieser `git mv` bricht *jeden* relativen Verweis der Plan-Datei** — anders als beim Slice wechselt
er die **Verzeichnistiefe** (flach → `done/`), und ein Pfad aus Tiefe *n* braucht aus *n+1* ein
zusätzliches `../`. Bei `welle-13` waren es **21** Verweise auf einen Schlag. Die
Lifecycle-Invariante (`links.resolve-from`, seit slice-080 im Modul `links`) kann das **nicht**
abfangen: ihre Bedingung („ein Verweis löst aus jedem Lifecycle-Verzeichnis
auf") setzt gleiche Ebenen voraus und ist hier nachweislich unerfüllbar — die Lücke ist deshalb
ausdrücklich ausgewiesen ([slice-089](done/wellenlos/slice-089-welle-datei-verweis-invariante.md)).
**Also: die Verweise im selben Commit nachziehen und `make gates` laufen lassen**; `doc-check` ist
das Netz, aber erst *nach* dem `mv`.
**Zum Steering-Loop:** die Einträge werden aus `docs/plan/steering-loop.md`
**gezogen**, nicht dorthin verschoben. Das Register bleibt der laufende Zähl-Ort; ein Kanal, der
erst bei der nächsten Closure entsteht, würde zwischen zwei Wellen nichts zählen — genau der
Fehler, den slice-057 vermieden hat.

**4 — Zeitdokumente archivieren.** Die Slice-Dateien der Welle, ihr eigener Plan und die
Review-Reports dieser Slices wandern in ein unveränderliches `done/<welle-id>/archiv.zip`; am Ort
bleibt ein gekürzter Stub, die Ergebnis-Notiz bleibt vollständig und flach.
*Beleg:* `make archive-wave WELLE=<id> APPLY=1` als eigener Commit, danach `make gates` und
`make verify` auf dem archivierten Stand erneut grün. Eingesammelt wird **nach der Welle, nicht
nach dem Verzeichnis** — auch die wellenlosen Slices, die seit der letzten Closure geschlossen
wurden.

**5 — Wave-Self-Close-Commit.** Ein einzelner, beobachtbarer Commit markiert den Abschluss.
*Beleg:* der Commit-Hash. Der Audit sieht *einen* Punkt, an dem die Welle schloss, statt eines
verstreuten Verschwindens über mehrere Commits.

**6 — Roadmap fortschreiben.** Die Welle wandert aus *Offene Wellen* in die Tabelle
*Abgeschlossene Wellen* (mit Zeiger auf ihre Ergebnis-Notiz); die erste Zeile aus *Nächste Wellen*
rückt unter *Offene Wellen* nach, sofern ihr Trigger gefeuert hat. Hat ein Trigger dabei eine Umplanung ausgelöst, bekommt
*Historische Trigger-Verschiebungen* ihren Eintrag.
*Beleg:* der Roadmap-Diff.

### Ab wann das gilt

**Ab der nächsten Welle.** Die zwölf bestehenden Wellen (`welle-00` … `welle-11`) existieren
ausschließlich als Prosa-Überschriften in der Roadmap; sie bekommen **keine** rückwirkenden
Ergebnis-Notizen. Eine nachträglich rekonstruierte Wellen-Closure wäre eine Nacherzählung, kein
Beleg — dieselbe Entscheidung wie bei der Review-Lücke slice-027…041
([`docs/reviews/README.md`](../../reviews/README.md)). Auch Welle-Plan-Dateien entstehen erst
künftig: wo es nie eine gab, gibt es nichts nach `done/` zu verschieben.

Die Prozedur ist **belegt**: `welle-12` ist ihr erster Durchlauf
([`done/welle-12-results.md`](done/welle-12-results.md)), `welle-13` der zweite und erste **mit**
einer Plan-Datei ([`done/welle-13-results.md`](done/welle-13-results.md)) — womit Schritt 3
vollständig gefahren ist statt nur zur Hälfte.
