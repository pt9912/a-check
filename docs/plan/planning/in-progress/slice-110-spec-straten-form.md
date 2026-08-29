# slice-110 — Spec-Straten: Urteil und Regelwerk-Zeiger

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-29.
**Berührte Spec-Stellen:** die Straten selbst — `spec/lastenheft.md`, `spec/architecture.md`,
`spec/spezifikation.md`; **keine** `AC-*`/`SPEC-*`/`ARC-*`-Kennung wird angelegt oder geändert.
**Deckt:** den Spec-Anteil aus `BEO-011`.
**Bezug:** [slice-105](../done/slice-105-form-review-nachholen.md) §4.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Ziel

Die drei ungeprüften Zeilen aus `BEO-011` lesen und urteilen — und den mechanischen Teil
nachtragen. **Vertragsinhalt wird nicht erfunden:** was Aussagen über Umfang, Begriffe oder
Fehlerverhalten verlangt, gehört dem Maintainer und wird benannt, nicht geschrieben.

Gelesen ergibt sich:

| Stratum | Befund | Urteil |
|---|---|---|
| `lastenheft.md` | §5 *Globale Out-of-Scope-Punkte* und §6 *Glossar* fehlen — die Nummerierung springt von 4 auf **7** | **echt**; Inhalt ist Vertragsaussage |
| `architecture.md` | *Externe Abhängigkeiten* und *Fehlermodelle und Resilienz* fehlen; die übrigen sind Namensvarianten | **echt**; Inhalt ist Architektur-Aussage |
| `spezifikation.md` | gliedert nach **Vertrags-Kennung** (`SPEC-CONF-001`, `SPEC-EXTRACT-001`, …) statt nach den sieben Themen der Ziel-Form | **bewusst anders** — aber **undeklariert** |

Der dritte Befund ist der schwerste: eine Strukturabweichung in einer kanonischen Quelle vom
**Rang 2**, für die kein Adaptions-Eintrag existiert. Nach der Fork-Regel ist eine Abweichung ohne
benannte ersetzte Regel keine Adaption, sondern ein Fork.

## 2. Definition of Done

- [ ] Die drei Straten tragen ihre Regelwerk-Zeiger (2 · 3 · 6) — Beleg: Zählung.
- [ ] `BEO-011` ist aufgelöst und nach `Gestrichene Einträge` verschoben; die vier Zeilen sind
      gelesen — Beleg: Register.
- [ ] Was Maintainer-Inhalt braucht, steht als eigene Beobachtung mit Beleg im Register, nicht als
      erfundener Text in der Spec — Beleg: Register und Diff.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `spec/lastenheft.md` | update | zwei Regelwerk-Zeiger |
| `spec/architecture.md` | update | drei Regelwerk-Zeiger |
| `spec/spezifikation.md` | update | sechs Regelwerk-Zeiger |
| `docs/plan/planning/observations.md` | update | `BEO-011` auflösen, zwei neue Zeilen |

**Auszuführende Gates:** `make gates` — tragend sind `doc-check` und **`matrix`**, weil die
Spec-Straten der Referenz-Richtung unterliegen: kein Stratum darf abwärts verweisen. Die neuen
Zeiger nennen Baseline-Abschnitte, keine ADRs oder Slices — genau deshalb sind sie zulässig.

## 4. Trigger

**Start:** unmittelbar.

**Rückführungen:** `in-progress` → `open`, falls ein Zeiger die Referenzmatrix rot macht — dann
ist die Zeiger-Form für Spec-Straten selbst eine Konventionsfrage.

## 5. Closure-Trigger

Zeiger gezählt, `BEO-011` aufgelöst, offene Inhalte benannt, Gates grün.

**Was bewusst nicht getan wird:** **Keine Sektion wird angelegt.** *Globale Out-of-Scope-Punkte*,
*Glossar*, *Externe Abhängigkeiten* und *Fehlermodelle und Resilienz* verlangen Aussagen über den
Vertrag und die Architektur — sie zu erfinden wäre schlimmer als ihr Fehlen, weil eine erfundene
Vertragsaussage abgenommen aussieht. Ebenso wird die Gliederung der Spezifikation **nicht**
umgebaut: ob sie nach Kennung oder nach Thema gliedert, ist eine Konventions-Entscheidung.

## 6. Risiken und offene Punkte

- *Vier Sektionen fehlen inhaltlich und brauchen Maintainer-Inhalt* — **Ausgang:** weiter offen,
  `BEO-017` im Beobachtungs-Register.
- *Die Gliederung der Spezifikation weicht undeklariert ab* — **Ausgang:** weiter offen,
  `BEO-018`. Entweder Adaptions-Eintrag mit benannter ersetzter Regel oder Angleichung; beides ist
  eine Entscheidung, keine Messung.
- *Die Nummerierung des Lastenhefts springt von 4 auf 7* — **Ausgang:** entfallen, gestrichen mit
  Begründung: das ist keine eigene Beobachtung, sondern der sichtbare Teil von `BEO-017`.

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird **Spec-Straten** (`spec/`), in der
Modus-Deklaration geführt, alle drei Achsen erfüllt.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-011` betrifft diese Sub-Area und wird mit
diesem Slice aufgelöst; sonst keine Treffer.

Alle berührten Sub-Areas GF.
