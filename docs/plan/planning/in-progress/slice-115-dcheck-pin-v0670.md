# slice-115 — d-check-Pin auf `v0.67.0`

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-29.
**Berührte Spec-Stellen:** — *(keine)*
**Deckt:** die zweite Trigger-Hälfte von
[slice-080](../open/slice-080-verify-abloesung-dcheck.md).
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

- [ ] `d-check.mk` ist aus `v0.67.0 --print-mk` neu erzeugt, die **einzige** a-check-Anpassung
      (Digest sticht Tag) ist wieder angebracht, und der Digest ist aus **zwei** Quellen bestätigt.
- [ ] `doc-structure` ist in [`AGENTS.md`](../../../../AGENTS.md) §4 **und**
      [`harness/README.md`](../../../../harness/README.md) §Sensors deklariert und in der
      GATES-Liste des PreToolUse-Guard — `doc-targets` und `guard-selftest` belegen beides.
- [ ] `make gates` bleibt über den **unveränderten** Bestand grün: derselbe Befundstand mit einem
      Werkzeug, das sechzehn Releases jünger ist.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

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
[slice-080](../open/slice-080-verify-abloesung-dcheck.md) und braucht dessen Messung. Ein Target
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

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Gate-/Werkzeug-Schicht** (`d-check.mk`,
`.claude/`) und mit den zwei Deklarations-Tabellen der **Harness-Einstieg**.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-014` (`doc-planning` ohne
Konfigurationsblock) betrifft dieselbe Schicht und bleibt offen — dieser Slice konfiguriert kein
Modul.

Alle berührten Sub-Areas GF.
