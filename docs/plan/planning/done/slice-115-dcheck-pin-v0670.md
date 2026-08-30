# slice-115 — d-check-Pin auf `v0.67.0`

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-29.
**Berührte Spec-Stellen:** — *(keine)*
**Deckt:** die zweite Trigger-Hälfte von
[slice-080](../in-progress/slice-080-verify-abloesung-dcheck.md).
**Bezug:** Maintainer-Auftrag 2026-08-29. [Roadmap](../in-progress/roadmap.md).

---

## 1. Ziel

Der Pin steht auf `v0.51.1`, d-check ist bei `v0.67.0` — **sechzehn Releases**. `slice-080` wartet
auf zwei Bedingungen; die erste ist heute gemessen (`structure` und `links.resolve-from` sind
vorhanden), die zweite ist dieser Slice.

**Das Risiko ist vorab gemessen, nicht abgeschätzt.** `v0.67.0` läuft über den heutigen Bestand mit
**228 Dateien, 0 Befunden** — identisch zu `v0.51.1`; ebenso die Module `targets` und `planning`.
Sechzehn Releases haben an dieser Doku nichts gedreht.

Das Fragment [`d-check.mk`](../../../../d-check.mk) ist **erzeugt** (`--print-mk`), nicht
handgeschrieben. Der Vergleich mit der neuen Fassung zeigt genau eine Änderung an der Target-Menge:
**`doc-structure`** kommt hinzu — das Modul, auf das `slice-080` wartet.

## 2. Definition of Done

- [x] `d-check.mk` ist aus `v0.67.0 --print-mk` neu erzeugt, die **einzige** a-check-Anpassung
      (Digest sticht Tag) ist wieder angebracht, und der Digest ist aus **zwei** Quellen bestätigt.
- [x] `doc-structure` ist in [`AGENTS.md`](../../../../AGENTS.md) §4 **und**
      [`harness/README.md`](../../../../harness/README.md) §Sensors deklariert und in der
      GATES-Liste des PreToolUse-Guard — `doc-targets` und `guard-selftest` belegen beides.
- [x] `make gates` bleibt über den **unveränderten** Bestand grün: derselbe Befundstand mit einem
      Werkzeug, das sechzehn Releases jünger ist.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `d-check.mk` | neu erzeugt | Fragment + Pin + Digest |
| [`AGENTS.md`](../../../../AGENTS.md) §4, `harness/README.md` §Sensors | update | `doc-structure` deklarieren |
| `.claude/hooks/pretooluse-command-guard.sh` | update | GATES-Liste |

**Auszuführende Gates:** `make gates` — tragend `doc-check` (neues Werkzeug), `doc-targets`
(Deklarations-Konsistenz über **beide** Tabellen) und `guard-selftest` (Drift-Wächter der
GATES-Liste). Zum Abschluss `make verify`.

## 4. Trigger

**Start:** eingetreten — das Release liegt vor, das Risiko ist gemessen.

**Rückführungen:** `in-progress` → `open`, falls das neue Werkzeug über den unveränderten Bestand
Befunde meldet, die eine Konventions-Entscheidung verlangen statt einer Korrektur.

## 5. Closure-Trigger

Pin gehoben, Digest zweifach belegt, `doc-structure` deklariert, Gates grün.

**Was bewusst nicht getan wird:** `doc-structure` wandert **nicht** ins `gates`-Aggregat und wird
**nicht** konfiguriert. Ob es die vier `verify-*` ablösen kann, ist Gegenstand von
[slice-080](../in-progress/slice-080-verify-abloesung-dcheck.md) und braucht dessen Messung. Ein Target
einzubinden, ohne es zu konfigurieren, ist genau der Fehler, den `slice-074` gemessen hat: das
Modul `targets` lief so **dreizehn Minor-Versionen** ins Leere. Ebenso bleiben `workflows`,
`citations` und `sources` unangetastet — je eine eigene Sichtung.

## 6. Risiken und offene Punkte

- *`sources` ist jetzt erreichbar und deckt die Asset-Integrität der vendored Baseline, die
  slice-047 offengelassen hat* — **Ausgang:** weiter offen im **Beobachtungs-Register**.
- *Sechzehn Releases können Verhalten geändert haben, das dieser Bestand nicht auslöst* — der
  Null-Befund beweist Gleichstand nur für die **heutige** Doku. **Ausgang:** gestrichen mit
  Begründung: eine Aussage über ungeschriebene Dokumente wäre nicht belegbar, und der nächste
  Gate-Lauf misst sie ohnehin.
- *Die zweite Trigger-Hälfte von `slice-080` fällt damit* — **Ausgang:** Folge-Slice; `slice-080`
  wird startbar, bleibt aber an seiner externen Vorbedingung (CR-Einreichung) hängen.

## 7. Closure-Notiz

**Geliefert:** Der Pin steht auf `v0.67.0`, der Digest ist aus zwei Quellen belegt, das Fragment
ist neu erzeugt, und `doc-structure` ist in beiden Doku-Tabellen und der GATES-Liste deklariert.

**Lerneintrag — Form: geschärfte Regel.** *Ein Pin-Bump ist keine Zeilen-Änderung, solange das
Werkzeug seine Module über eine **geschlossene Verbots-Liste gegen eine offene Menge** auswählt.*
Jedes Einzelmodul-Target im Fragment hat die Form `--enable X --disable <alle anderen, einzeln
aufgezählt>`. Diese Aufzählung ist beim Erzeugen vollständig und veraltet mit **jedem** Modul, das
stromaufwärts hinzukommt. Hier waren es zwei (`structure`, `workflows`); ein reiner Tausch der
beiden Pin-Zeilen hätte sie in **fünf** Targets mitlaufen lassen, die per Vertrag genau ein Modul
fahren. Der Fehler wäre **grün** eingezogen: er wirkt in Richtung *mehr* Prüfung und fällt erst
auf, wenn ein zugeschaltetes Modul zufällig etwas findet — dann aber unter einem Target-Namen, der
etwas anderes verspricht. Kein Gate deckt das ab: `gate-consistency` prüft die Wohlgeformtheit von
Tag und Digest, nicht die Vollständigkeit der Disable-Listen. *Weil* das so ist, ist die im Kopf
von [`d-check.mk`](../../../../d-check.mk) genannte Reihenfolge — erst `--print-mk`, dann den
Digest setzen — nicht Bequemlichkeit, sondern die einzige Form, die diese Drift ausschließt.

**Zwei beobachtbare Closure-Kriterien:**

1. `diff` des committeten Fragments gegen `d-check:v0.67.0 --print-mk` ist auf die
   a-check-Anpassung reduziert: fünf Rezept-Zeilen tragen `--disable structure --disable
   workflows`, ein Target (`doc-structure`) kommt hinzu, sonst nichts.
2. Alle sechs `doc-*`-Sensoren laufen über den **unveränderten** Bestand grün, jeder mit
   **229 Dateien, 0 Befunden** — `doc-check`, `doc-targets`, `doc-planning`, `doc-tracked`,
   `doc-immutable STAGED=1` und das neue `doc-structure`.

**Offene Risiken und ihr Ausgang:**

- *`sources` ist jetzt erreichbar und deckt die Asset-Integrität der vendored Baseline* —
  **Ausgang:** weiter offen im **Beobachtungs-Register** als `BEO-021`.
- *Sechzehn Releases können Verhalten geändert haben, das dieser Bestand nicht auslöst* —
  **Ausgang:** gestrichen mit Begründung: der Null-Befund gilt für die heutige Doku, und eine
  Aussage über ungeschriebene Dokumente wäre nicht belegbar. Der nächste Gate-Lauf misst sie.
- *Die zweite Trigger-Hälfte von [slice-080](../in-progress/slice-080-verify-abloesung-dcheck.md) fällt
  damit* — **Ausgang:** Folge-Slice; beide Hälften sind jetzt erfüllt.

**Beobachtungs-Register:** `BEO-021` neu angelegt; `BEO-014` auf **2×** erhöht — `doc-structure`
meldet ohne `structure`-Block grün statt zu schweigen, dieselbe Form, die `targets` nach slice-074
dreizehn Minor-Versionen lang zeigte. Das ist hier **kein** Wiederholungsfehler: die Doku sagt an
beiden Stellen ausdrücklich, dass kein Konfigurationsblock existiert und das Target nicht im
Aggregat steht.

**Folge-Slices:** [slice-080](../in-progress/slice-080-verify-abloesung-dcheck.md) ist an seiner
Werkzeug-Bedingung entblockt und hängt nur noch an der externen Vorbedingung (CR-Einreichung).

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Gate-/Werkzeug-Schicht** (`d-check.mk`,
`.claude/`) und mit den zwei Deklarations-Tabellen der **Harness-Einstieg**.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-014` (`doc-planning` ohne
Konfigurationsblock) betrifft dieselbe Schicht und bleibt offen — dieser Slice konfiguriert kein
Modul.

Alle berührten Sub-Areas GF.
