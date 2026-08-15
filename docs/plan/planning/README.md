# Planungs-Harness — Slices und Wellen

Zwei Ebenen: der **Slice** ist die Arbeitseinheit, die **Welle** bündelt Slices zu einem
abschließbaren Vorhaben. Der Slice-Zyklus ist eine Zustandsmaschine über Verzeichnisse
(`open/` → `next/` → `in-progress/` → `done/`, [`AGENTS.md`](../../../AGENTS.md) §5); die Welle
liegt eine Ebene darüber und schließt über eine **Prozedur**, nicht über einen Datei-Übergang.

- Slice-Form: [`slice.template.md`](slice.template.md) — höchstens drei DoD-Punkte, höchstens
  zwei Schichten, benannte Lerneintrag-Form. **Maschinell geprüft von `make verify` sind die
  DoD-Zahl und die benannte Lerneintrag-Form**; „höchstens zwei Schichten" ist **Review-Sache**
  und ausdrücklich kein Gate — was eine Schicht ist, ist eine Ermessensfrage über Modul-Grenzen,
  und ein Zähler darüber wäre Schein-Genauigkeit
  ([`tools/verify-slice-form.sh`](../../../tools/verify-slice-form.sh) trägt dieselbe Begründung).
  Bis slice-076 las sich diese Zeile, als fiele alle drei Punkte ein Gate.
- Sequenzierungs-Autorität ist die [Roadmap](in-progress/roadmap.md); sie bleibt es auch nach
  einer Wellen-Closure.

## Wellen-Closure-Prozedur

Angelegt in slice-066 (Fund **B-13**), Quelle `modul-06` §Wellen-Closure. **Fünf Schritte, jeder
mit einem Beleg — keiner mit einem Datum.** Erst wenn alle fünf Belege vorliegen, ist eine Welle
*auditierbar* geschlossen.

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
zusätzliches `../`. Bei `welle-13` waren es **21** Verweise auf einen Schlag. `verify-slice-links`
kann das **nicht** abfangen: seine Invariante („ein Verweis löst aus jedem Lifecycle-Verzeichnis
auf") setzt gleiche Ebenen voraus und ist hier nachweislich unerfüllbar — der Sensor weist die
Lücke deshalb ausdrücklich aus ([slice-089](in-progress/slice-089-welle-datei-verweis-invariante.md)).
**Also: die Verweise im selben Commit nachziehen und `make gates` laufen lassen**; `doc-check` ist
das Netz, aber erst *nach* dem `mv`.
**Zum Steering-Loop:** die Einträge werden aus [`docs/plan/steering-loop.md`](../steering-loop.md)
**gezogen**, nicht dorthin verschoben. Das Register bleibt der laufende Zähl-Ort; ein Kanal, der
erst bei der nächsten Closure entsteht, würde zwischen zwei Wellen nichts zählen — genau der
Fehler, den slice-057 vermieden hat.

**4 — Wave-Self-Close-Commit.** Ein einzelner, beobachtbarer Commit markiert den Abschluss.
*Beleg:* der Commit-Hash. Der Audit sieht *einen* Punkt, an dem die Welle schloss, statt eines
verstreuten Verschwindens über mehrere Commits.

**5 — Roadmap fortschreiben.** Die Welle wandert aus *Aktuelle Welle* in die Tabelle
*Abgeschlossene Wellen* (mit Zeiger auf ihre Ergebnis-Notiz); die erste Zeile aus *Nächste Wellen*
wird zur neuen aktuellen. Hat ein Trigger dabei eine Umplanung ausgelöst, bekommt
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
