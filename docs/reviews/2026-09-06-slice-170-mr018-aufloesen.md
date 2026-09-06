# Review-Report: slice-170 — 2026-09-06

**Review-Art:** Plan — geprüft gegen den Slice-Plan, den aufgelösten
Adaptions-Eintrag und die Baseline (Modul 10 §Drei Review-Arten). Die
Kern-Behauptung ist ein **Trigger-Urteil** („die Bedingung ist eingetreten"),
und das ist nur gegen die Artefakte prüfbar, nicht gegen den Plan.

**Gegenstand:** Commits `e168f41`, `f171af7`, `ceff549` und `5b65f20` —
letzterer entstand *während* des Reviews und hat den Arbeitsbaum-Stand
ersetzt; alle Befunde sind gegen `5b65f20` gemessen.

**Skill:** `.harness/skills/reviewer.md` @ Stand `5649d84` (unverändert
seit Anlage) · <!-- d-check:ignore -->
**Modell:** claude-sonnet-5, **getrennter Kontext** (eigener Subagent, kein
`fork`) · **Datum:** 2026-09-06

**Eingangs-Kontext:**

- `docs/plan/planning/done/slice-170-mr018-aufloesen.md`
- `harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md`,
  `harness/conventions.md`
- `.harness/baseline/v6.2.0/regelwerk/` (`modul-05`, `modul-06`, `modul-08`,
  `grundlagen-harness-dateien.md`), beide vendorten `AGENTS.template.md`
- die beiden neuen Register-Verzeichnisse unter `observations/BEO-HARNESS/`

**Eigene Läufe:** `make gates`, `make verify`, `make doc-check`,
`make doc-immutable RANGE=…`, `make commit-scope-check RANGE=…`,
`make trace-check` (vier Ranges), `git show --stat -M`, `git log --follow`,
Schleife über das Feld `Ersetzt-Baseline-Regel` aller 20 `MR`-Dateien.

---

## Findings

### F-1 — Falsche Zählaussage in einer ab Merge unveränderlichen `evidence/`-Datei

- `kategorie`: HIGH
- `quelle`: Schleife über `Ersetzt-Baseline-Regel` in `harness/conventions/MR-*.md`
  und `harness/conventions/done/MR-*.md`
- `pfad`: `…/zwei-baseline-staende-nach-migrationsende/evidence/slice-170.md`;
  wortgleich `slice-170-mr018-aufloesen.md` §7
- `befund`: Behauptet war „Die Auflösung von `MR-018` entfernt einen von
  **sechs** `v6.0.0`-Ankern; fünf bleiben stehen." Gemessen tragen genau
  **fünf** `MR`-Dateien einen `v6.0.0`-Anker in diesem Feld — `MR-011`,
  `MR-012`, `MR-014`, `MR-015`, `MR-016` —, **vor wie nach** diesem Slice
  dieselben fünf. `MR-018`s Feld trägt seit `541581f` (2026-09-05) ein „—"
  und war nie Teil der Menge. Es wurde nichts entfernt, und die Ausgangszahl
  war nicht sechs. Die Zahl stammt ungeprüft aus `slice-167` §3 bzw.
  `welle-14-results.md`, die `MR-018` mitzählen.
- `verifizierbar`: ja — die Schleife ist reproduzierbar; kein Sensor prüft
  Zählaussagen, `gates`/`verify` waren grün.
- `klasse`: Zahl aus einem Vorgänger-Dokument übernommen statt gemessen

### F-2 — Der zweite Trigger ist umgedeutet, nicht gemessen

- `kategorie`: MEDIUM
- `quelle`: `diff` der beiden vendorten `modul-08` und `AGENTS.template.md`
- `pfad`: `slice-170-mr018-aufloesen.md` §2, Trigger-Tabelle
- `befund`: Der Auflösungs-Trigger verlangt eine Baseline-Migration, die den
  Gegenstand **inhaltlich ändert**. Upstream hat zwischen `v6.1.0` und
  `v6.2.0` daran nichts geändert; a-check hat den Zuwachs nur **nachvollzogen**.
  Eingetreten ist die im selben Eintrag getrennt notierte
  Rückbau-Kandidat-Bedingung. Der Plan verbuchte beide als eingetreten.
- `verifizierbar`: ja.
- `klasse`: zwei Bedingungen zu einer verschmolzen

### F-3 — Verweis-Form widerspricht der Baseline, ohne dass die Regel erwogen wird

- `kategorie`: MEDIUM
- `quelle`: `.harness/baseline/v6.2.0/regelwerk/grundlagen-harness-dateien.md`
  Zeilen 259–268
- `pfad`: die in `f171af7` nachgezogenen Links, die neuen Links in Plan und
  `evidence/`
- `befund`: Die Baseline schreibt für Verweise aus `AGENTS.md`, einem Slice
  oder einer ADR auf eine Adaption die Index-Form
  `harness/conventions.md#mr-<NNN>` vor, ausdrücklich **nicht** den Pfad auf
  die Eintrags-Datei — mit der Begründung, ein Pfad-Link breche „genau in
  dem Moment, in dem die Adaption sich auflöst". Genau das trat ein: acht
  Fremddateien mussten nachgezogen werden. Der Slice zieht die Pfad-Links
  weiter und legt neue an, statt umzustellen; die Regel wird nirgends
  genannt.
- `verifizierbar`: ja — `doc-check` prüft Auflösbarkeit, nicht Form.
- `klasse`: Baseline-Regel unerwogen, weil der Bestand sie schon verletzt

### F-4 — Die Lücke wird upstream verortet, obwohl sie im Repo näher liegt

- `kategorie`: MEDIUM
- `quelle`: `modul-06-roadmap.md` Tabelle *Träger im Repo ohne Wellen*;
  `grep -n "Trigger-Audit" AGENTS.md harness/README.md docs/plan/planning/README.md`
- `pfad`: `…/mr-aufloesungs-trigger-ohne-waechter/observation.md`;
  `slice-170` §8 Lerneintrag
- `befund`: Richtig gemessen ist, dass Closure-Schritt 2 nur drei
  Artefaktklassen nennt. Falsch war der Zusatz „und die Slice-Closure kennt
  den Schritt in dieser Form ebenfalls nicht": `modul-06` weist den
  Trigger-Audit im wellenlosen Betrieb ausdrücklich der Slice-Closure zu.
  Ungenannt blieb die nähere Lücke — a-check hat den Schritt für **keine**
  der drei Klassen verkörpert (null Treffer in den drei Harness-Dateien).
- `verifizierbar`: ja.
- `klasse`: fremde Spec-Lücke benannt, eigene übersehen

### F-5 — Nachgezogene Links lassen die umgebende Prosa falsch stehen

- `kategorie`: LOW
- `pfad`: `done/slice-167-…md` (Tabellenzelle), `done/slice-162-…md`
- `befund`: In `slice-167` zeigt `MR-018` nach `conventions/done/`, während
  dieselbe Zelle ihn als „(aktiv)" führt und von „MR-017 (aufgelöst)"
  abgrenzt. In `slice-162` war der Link*text* der alte, nicht mehr
  existierende Pfad bei neuem Ziel.

### F-6 — Als Zitat markierte Stelle war keines

- `kategorie`: LOW
- `pfad`: `…/zwei-baseline-staende-nach-migrationsende/evidence/slice-170.md`
- `befund`: zitiert „keinen absehbaren Trigger"; die Quelle schreibt „kein
  absehbarer Trigger".

### F-7 — `make trace-check` ist im lokalen Klon nicht befragbar (vorbestehend)

- `kategorie`: LOW
- `pfad`: `Makefile:205`
- `befund`: bricht mit `Range-Basis-Vorfahren nicht lesbar: object not
  found` ab, für jeden geprüften Range. Der Klon ist nicht shallow, die
  Objekte lösen für `git` auf; `.git/objects/pack/` enthält ein
  nicht-kanonisch benanntes `loose-<sha>.pack`. Nicht durch diesen Slice
  verursacht.

### F-8 — Fremde Beobachtung trägt einen nicht deklarierten Sub-Area-Namen

- `kategorie`: INFO
- `pfad`: `…/BEO-HARNESS/sensor-ohne-dod-phrase-wirkungslos/observation.md`
- `befund`: „Sub-Area: Harness-Tooling" — ein Name, den die
  Modus-Deklaration nicht führt. Vorbestehend.

### F-9 — Ruhe-Marker durch erklärende Prosa ersetzt

- `kategorie`: INFO
- `pfad`: `in-progress/roadmap.md`
- `befund`: Der Klammersatz sagte zusätzlich etwas über `in-progress/` aus —
  eine zweite Aussage über denselben Zustand, neben dem Marker, der beim
  Leeren zurückmuss.

## Negativbefunde

Zwanzig Kontrollen ohne Befund, darunter: der `diff` der vendorten
`AGENTS.template.md` (Rückbau-Bedingung **eingetreten**) · Wort-für-Wort-
Vergleich `AGENTS.md` §6 gegen das Template (Präfix identisch, zwei Sätze
angehängt, keine Streichung) · die Begründung „kein Nachfolge-Eintrag"
trägt, und die `MR-003`-Präzedenz wird nicht überdehnt · Anker `mr-018`
erhalten, alle drei Verweise darauf lösen auf · kein Link zeigt mehr auf den
alten Pfad · Rename bei **100 %**, `git log --follow` reicht darüber hinweg ·
§3.3 eingehalten (die verschobene Datei mit 0 Zeilen Änderung im
Rename-Commit, der Link-Tiefen-Fix als eigener) · `doc-immutable` Exit 0,
keine ADR berührt, `MR-018`s Felder unangetastet · Slice-Form zum Review-Zeitpunkt: **2**
zählende Liefer-Punkte, eine Schicht, Kopffelder vollständig (mit der
F-3-Einarbeitung sind es **3** — die Größen-Regel bleibt gewahrt, ohne
Reserve) · Closure-Notiz
mit Ursache statt Floskel · beide Risiko-Ausgänge aus der geschlossenen
Dreier-Menge (die zuerst vorgefundene Fassung „eingetreten →
Maintainer-Entscheidung" wäre **unzulässig** gewesen) · Register-Form beider
neuer Einträge vollständig, Kürzel `HARNESS` deklariert, Zähler abgeleitet ·
die Abgrenzung zur Geschwister-Beobachtung trägt (sie steht auf
`verkörpert`; ein Beleg dort hieße „die Verkörperung wirkt nicht") ·
WIP-Limit = 1 · `commit-scope-check` Exit 0 · die DoD-Zeile trägt die
Trigger-Phrase in der Form, die der Selbsttest als Positiv-Kontrolle
nachweist · `gates`/`verify` je zweimal Exit 0, Arbeitsbaum danach leer.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 3 |
| LOW | 3 |
| INFO | 2 |

**Finding-Klassen dieses Laufs:** Zahl aus einem Vorgänger-Dokument
übernommen statt gemessen · zwei Bedingungen zu einer verschmolzen ·
Baseline-Regel unerwogen, weil der Bestand sie schon verletzt · fremde
Spec-Lücke benannt, eigene übersehen

## Verdikt

**Merge-blockierend:** war **ja** (F-1) — eine gegen die Artefakte
widerlegbare Zahlenaussage stand in einer `evidence/`-Datei, die ab Merge
unveränderlich ist. Vor Abschluss korrigiert; der Stand war noch nicht
gepusht, ein Nachfolge-Eintrag also nicht nötig.

**Eingearbeitet:** F-1 (Zahl gemessen und die Herkunft der falschen Zahl
benannt), F-2 (Trigger-Tabelle trennt die zwei Bedingungen; die Auflösung
trägt allein aus der ersten), F-4 (Beobachtung und Lerneintrag zeigen jetzt
auf die eigene Lücke), F-5/F-6/F-9 (Text). F-7 als eigene Beobachtung
registriert (`BEO-GATE/trace-check-lokal-nicht-befragbar`). **F-3 ebenfalls
eingearbeitet** — auf ausdrückliche Maintainer-Entscheidung wurde die
Verweis-Form repo-weit umgestellt: **129** Pfad-Links in 19 Dateien tragen
jetzt die Index-Form `harness/conventions.md#mr-<NNN>`; ausgenommen bleiben
`conventions.md` selbst (es *ist* der Index) und die akzeptierten
Eintrags-Dateien. Die Klasse dahinter ist als
`BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`
registriert. **Offen:** nur F-8 — vorbestehend und fremd.

**Übergabe:** Dieser Report ist Lauf-Beleg, keine Verifikation —
DoD-/Spec-Konformität prüft der Verifier separat (Modul 11).
