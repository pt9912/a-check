# slice-113 — `steering-loop.md` fällt

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-29.
**Berührte Spec-Stellen:** — *(keine)*
**Deckt:** keine `AC-*`.
**Bezug:** Maintainer-Entscheidung 2026-08-29; entblockt durch
[`ADR-0035`](../../adr/0035-grenz-diagnose-gegen-globs.md), die
[`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md) ablöst.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Ziel

`docs/plan/steering-loop.md` ist seit slice-101 der **zweite** Zähler für dieselbe Sache — heute
ist derselbe Vorfall in beide gelaufen (`SL-004` auf fünf erhöht, `BEO-005` mit denselben Belegen
`slice-099, slice-100`). Ein Vorfall, zwei Bücher, zwei Stände. Die Ziel-Form kennt keine solche
Datei: der Zähler **ist** das Beobachtungs-Register.

**Der Slice hieß ursprünglich „Wanderung". Er heißt jetzt „Löschung", und der Grund ist gemessen.**

*Migration scheidet aus:* `modul-06` bindet den Beleg an die **Slice-Closure**, in der die
Beobachtung festgehalten wird — deshalb ist `slice-NNN` die Beleg-Form. Die sechs `SL-*` entstanden
**vor** dem Register und belegen ihre Vorfälle mit Commit-Hashes (`SL-001`, `SL-002`) und Wellen
(`SL-006`). Sie in Slice-Belege umzurechnen hieße, nachträglich zuzuordnen, welche Closure sie
*hätte* festhalten sollen — genau die *„Nacherzählung, kein Beleg"*, die
[`planning/README.md`](../README.md) §Ab wann das gilt schon für zwölf Wellen ausgeschlossen hat.

*Archiv scheidet aus:* Wer soll damit etwas anfangen? Gemessen tragen **fünf von sechs** Regeln
ihre Herkunft bereits am **Verkörperungs-Ort** — `SL-001` im Guard, `SL-002` in
[`AGENTS.md`](../../../../AGENTS.md) und `verify-slice-links.sh`, `SL-003` in
`commit-scope-check.sh`, `SL-004` in `verify-closure-notes.sh`, `SL-005` in
`gate-consistency.sh`. Wer wissen will, warum der Guard Pipes ablehnt, liest es im Guard. `SL-006`
steht nirgends, **weil nichts gebaut wurde** — sein Eintrag sagt selbst, `doc-check` fange jeden
Fall. Ein Archiv wäre eine dritte Kopie neben Verkörperung und Welle-Notiz und kostete jeden Lauf
Kontext.

Was verlorengeht, ist die **Analyse**. Die hält `git` — so will es Hard Rule 3.7, heute mit
slice-108 adoptiert: *„die Abwägung gehört in die ADR, die Historie in `git`."*

## 2. Definition of Done

- [x] `docs/plan/steering-loop.md` ist entfernt; die drei lebenden Verweise
      ([`AGENTS.md`](../../../../AGENTS.md), `.claude/commands/slice.md`,
      `tools/verify-slice-links.sh`) zeigen aufs Beobachtungs-Register.
- [x] Die historischen Zeiger in `done/`, `docs/reviews/` und
      [`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md) sind entlinkt, ihre **Aussage**
      unverändert — wie in slice-112.
- [x] `make doc-immutable STAGED=1` grün: der Eingriff an
      [`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md) ist zulässig, **gemessen** statt
      angenommen — sie trägt seit [`ADR-0035`](../../adr/0035-grenz-diagnose-gegen-globs.md) nicht
      mehr `Accepted`.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `docs/plan/steering-loop.md` | entfällt | zweiter Zähler |
| `AGENTS.md`, `.claude/commands/slice.md`, `tools/verify-slice-links.sh` | update | lebende Verweise |
| `done/`, `docs/reviews/`, [`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md) | update | tote Zeiger entlinken |

**Auszuführende Gates:** `make doc-immutable STAGED=1` **vor** dem Commit, dann `make gates`
(tragend `doc-check`), zum Abschluss `make verify`.

## 4. Trigger

**Start:** eingetreten — [`ADR-0035`](../../adr/0035-grenz-diagnose-gegen-globs.md) liegt vor und
zitiert `SL-004` nicht.

**Rückführungen:** `in-progress` → `open`, falls `doc-immutable` den Eingriff an der abgelösten ADR
doch ablehnt — dann bleibt die Datei, bis dafür ein eigener Weg gefunden ist.

## 5. Closure-Trigger

Datei weg, kein toter Zeiger, `doc-immutable` und `gates` grün.

**Was bewusst nicht getan wird:** Die sechs Zähler wandern **nicht** ins Register — siehe §1. Die
zwei Ergänzungen, die heute an `SL-004` entstanden sind, verschwinden mit der Datei; ihr Zähler ist
`BEO-005`, der die Belege bereits trägt. Und die Welle-Notizen bleiben unberührt: sie tragen die
Vorfälle ihrer Welle an ihrem vorgesehenen Ort.

## 6. Risiken und offene Punkte

- *Die `SL-*`-Kennungen verlieren ihren Anker; 154 Zitate im Bestand zeigen ins Leere* —
  **Ausgang:** weiter offen im **Beobachtungs-Register**, falls jemand die Auffindbarkeit
  tatsächlich vermisst. Die Kennungen bleiben im Text lesbar, und fünf von sechs sind am
  Verkörperungs-Ort auffindbar.
- *`SL-006` ist die einzige Beobachtung ohne Verkörperung und ohne Register-Zeile* —
  **Ausgang:** weiter offen im **Beobachtungs-Register**: sie beginnt bei ihrem nächsten Auftreten
  neu, mit dann korrektem Slice-Beleg. Das Register ist die Vergabestelle.
- *Die Analyse der sechs Einträge ist nach der Löschung nur noch in `git`* — **Ausgang:**
  gestrichen mit Begründung: genau das verlangt Hard Rule 3.7, und die Verkörperungs-Orte tragen
  das Warum.

## 7. Closure-Notiz

**Geliefert:** `docs/plan/steering-loop.md` ist weg — 302 Zeilen und der zweite Zähler. 28 Dateien
angepasst: drei lebende Verweise zeigen aufs Register, der Rest ist entlinkt, die Aussagen stehen
unverändert.

**Lerneintrag — Form: geschärfte Regel.** *Das „Warum" einer Regel gehört an die Regel, nicht in
eine Datei daneben.* Die Frage *„wer soll mit einem Archiv etwas anfangen?"* hat die Messung
erzwungen, die den Slice umgedreht hat: **fünf von sechs** Einträgen stehen längst am
Verkörperungs-Ort — `SL-001` im Guard, `SL-003` in `commit-scope-check.sh`, `SL-004` in
`verify-closure-notes.sh` und so fort. Niemand hatte das geplant; es ist passiert, **weil** wer
einen Guard oder Sensor baut, das Warum daneben schreibt. Die separate Datei war damit von Anfang
an die redundante Hälfte, und `SL-006` beweist die Regel von der anderen Seite: dort steht nichts,
weil nichts gebaut wurde. **Der Prüfsatz:** bevor ein erklärendes Artefakt archiviert wird, ist zu
messen, ob seine Erklärung nicht ohnehin dort steht, wo sie wirkt.

**Zwei beobachtbare Closure-Kriterien:**

1. `make doc-immutable STAGED=1` läuft mit **Exit 0** über den Eingriff an
   [`ADR-0031`](../../adr/0031-heuristik-grenzen-diagnose.md) — zulässig, weil sie seit
   [`ADR-0035`](../../adr/0035-grenz-diagnose-gegen-globs.md) nicht mehr `Accepted` trägt und das
   Immutabilitäts-Prädikat damit nicht greift. **Gemessen, nicht angenommen** — genau der Punkt,
   an dem ich die Immutabilität zuvor als Sackgasse gelesen hatte.
2. `doc-check` prüft **228** Dateien mit **0** Befunden, nachdem eine Datei gelöscht wurde, auf die
   27 andere zeigten.

**Offene Risiken und ihr Ausgang:** siehe §6 — zwei bleiben offen und gehen ins
**Beobachtungs-Register**, sobald sie eintreten; einer ist gestrichen mit Begründung.

**Beobachtungs-Register:** keine Beobachtung angefallen. Die zwei offenen Risiken sind
**bedingt** — sie treten ein, wenn jemand die Auffindbarkeit vermisst oder `SL-006`s Muster
wiederkehrt. Eine Kennung entsteht beim Erstauftreten, nicht auf Vorrat; das war die Lehre aus
slice-114.

**Folge-Slices:** keine. Offen bleibt der d-check-Pin-Bump, dessen erste Trigger-Hälfte gemessen
ist.

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird **Planungs-Harness**; die drei lebenden
Verweise liegen in **Harness-Einstieg** und **Gate-/Werkzeug-Schicht**, tragen aber je eine Zeile.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-005` ist der Zähler, der heute doppelt geführt
wurde, und bleibt bestehen; `BEO-008` (Verweise aus `done/` auf wandernde Slices) hat dieser Slice
selbst schon einmal ausgelöst.

Alle berührten Sub-Areas GF.
