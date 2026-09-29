# Review-Report: slice-189 — 2026-09-29

**Review-Art:** Plan + Code — geprüft wird der Doku-Diff gegen den Slice-Plan,
die Ziel-Formen der vendorten Baseline und die Hard Rules (`AGENTS.md` §3).

**Gegenstand:** `HEAD~2..HEAD` — zwei Commits: `256f3c7` (docs(spec):
Lastenheft-Status Draft→Accepted, Architektur-§8-Historie gestrichen,
Frische-Marker am Kopf) und `a138a88` (docs(planning): §2-Messung,
§3.1–3.3-Klassifizierung, §7-Ausgänge). Ad hoc vorab verifiziert: die Range
trägt exakt diese zwei Commits (`git log --oneline -6`).

**Skill:** `.harness/skills/reviewer.md` @ `a138a88` (HEAD des Laufs) ·
**Modell:** glm-5.3-flash · **Datum:** 2026-09-29.

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- Slice-Plan: `slice-189` (§1 Ziel/Abgrenzung mit Architektur-Klausel, §2
  Ausgangsmessung, §3.1–3.3 Klassifizierung, §4 DoD, §7 Risiken)
- Ziel-Formen: `v6.13.0` · `templates/spec/lastenheft.template.md`,
  `v6.13.0` · `templates/spec/spezifikation.template.md`,
  `v6.13.0` · `templates/spec/architecture.template.md`
- `v6.13.0` · `regelwerk/modul-03-spec.md` §Spec-Stratifizierung und
  §Ziel-Form: Architektur-Sicht; `v6.13.0` · `regelwerk/modul-05-planning-harness.md`
  §Offene Risiken werden bei Closure aufgelöst
- [AGENTS.md](../../AGENTS.md) §3.4, §5, §6 (Schritt 7)
- Adaptionen: [MR-000](../../harness/conventions.md#mr-000),
  [MR-014](../../harness/conventions/MR-014-keine-agenten-telemetrie.md),
  [MR-027](../../harness/conventions/MR-027-verfeinerungs-form-v6130.md)
- früherer Report zum selben Bereich: der zu `slice-187` (aus dessen Archiv;
  dort insbesondere F-1 und F-10)

**Geltungsbereich dieses Laufs** (Mess-Regel 1): geprüft ist der
Commit-Diff `HEAD~2..HEAD` gegen den Plan und die drei Ziel-Formen; die
Klassifizierung ist je Zeile aller drei §3-Tabellen (acht Zeilen plus die
§3.3-Entscheidung) gestichprobt; ausgeführte Sensoren: `make doc-check`,
`make doc-structure`, `make doc-trace`, `make doc-immutable`
(`RANGE=HEAD~2..HEAD`). **Nicht** geprüft: `make gates`/`make verify`
(Verifier-Rolle), der Wiederholbarkeitshinweis der §2-Messung (das Instrument
wurde nicht neu gebaut, seine Zahlen sind gegen slice-187 §2 aus dem Archiv
abgeglichen), und die `AC-*`-Inhalte jenseits des Diff (nur Status-Zeile
geändert, siehe Negativbefund 1).

---

## Findings

### F-1 — Der Status-Flip auf `Accepted` trägt keinen belegten Abnahme-Vorgang; die Begründung mischt Traceability mit Abnahme

- `kategorie`: MEDIUM
- `quelle`: `v6.13.0` · `templates/spec/lastenheft.template.md` („gilt dem
  **Dokument**"; „ab `Accepted` ist jede Änderung eine Vertragsänderung");
  `v6.13.0` · `regelwerk/modul-03-spec.md` §Spec-Stratifizierung
- `pfad`: `spec/lastenheft.md`:5; `slice-189` Plan §3.1, Zeile „Status-Feld"
  (`docs/plan/planning/in-progress/slice-189-voll-abgleich-spec-straten.md`:88)
- `befund`: Die Begründung „das Dokument ist abnahmebindend, alle 21
  Anforderungen geprüft" belegt die Traceability (`make doc-trace`: 21
  Anforderungen, 0 Waisen — selbst nachgemessen), aber „geprüft" ist nicht
  „abgenommen": ein Abnahme-Vorgang des Maintainers ist im Repo nirgends
  dokumentiert, und die letzte Historie-Zeile (0.27.0, 2026-09-05) trug
  ausdrücklich „Status bleibt `Draft`, daher keine Change-Request-Pflicht".
  Dieselbe Tabellenzelle stuft die RB-Ergänzung als „Maintainer-Entscheid"
  ein, den Status-Flip — der denselben Vertragstyp einfriert — aber als
  „Doku-Pflege (hier ausgeführt)". Dass der Maintainer beide Commits selbst
  geführt hat (`pt9912` als Autor), ist ein Träger für den Vorgang, aber
  keine ausgewiesene Abnahme.
- `verifizierbar`: je zur Hälfte — `make doc-trace` bestätigt „21 geprüft";
  die Abnahme ist gegen kein Repo-Artefakt belegbar (das ist die Klärung).
- `klasse`: unbelegte Tatsachenbehauptung

### F-2 — Der Abgleich klassifiziert nur Vorlage→Repo; repo-seitige Ziel-Form-Abweichungen beider Straten bleiben unklassifiziert

- `kategorie`: MEDIUM
- `quelle`: Reviewer-Skill §Mess-Regeln, *Geltungsbereich einer Messung*;
  `v6.13.0` · `regelwerk/modul-03-spec.md` §Ziel-Form: Spezifikation
  („Einen eigenen Status trägt sie nicht"; „Ihre Historie führt Datum und
  Änderung — keine Version")
- `pfad`: `spec/spezifikation.md`:3, 5, 561; `spec/architecture.md`:5;
  `slice-189` Plan §2
- `befund`: [spezifikation.md](../../spec/spezifikation.md) trägt Kopf-Version
  (0.32.0), Kopf-Status („Draft") und eine Version-Spalte in der
  Historie-Tabelle — die Ziel-Form setzt im Kopf weder Version noch Status an
  und verlangt für die Historie `| Datum | Änderung |`;
  [architecture.md](../../spec/architecture.md) trägt `Status: Draft` gegen die
  Ziel-Form „Status: Aktiv". §3.1/§3.2 klassifizieren ausschließlich
  Vorlagen-Elemente; diese repo-seitigen Abweichungen — für das
  Vorlage→Repo-Instrument unsichtbar, Grenze 3 in slice-187 §2 — stehen in
  keiner Zeile, obwohl der DoD je-Kandidat-Klassifizierung für beide Straten
  fordert und der Plan „Voll-Abgleich" sagt.
- `verifizierbar`: ja — Kopf- und Tabellenvergleich gegen
  `v6.13.0` · `templates/spec/spezifikation.template.md` und
  `v6.13.0` · `templates/spec/architecture.template.md` (Zeile 9).
- `klasse`: Geltungsbereich einer Messung nur in eine Richtung genannt
  (Bestand: slice-187-Review F-10)

### F-3 — Die „keine Historie"-Klausel wird „Hard Rule §3.4" zugeordnet; sie steht dort nicht

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §3.4; `v6.13.0` ·
  `regelwerk/modul-03-spec.md` §Ziel-Form: Architektur-Sicht („Keine Historie,
  nur `**Letzte Änderung:**` im Kopf"); `v6.13.0` ·
  `templates/spec/architecture.template.md` (Hard-Rule-Block, „**keine
  Historie**")
- `pfad`: `slice-189` Plan §3.3
  (`docs/plan/planning/in-progress/slice-189-voll-abgleich-spec-straten.md`:105);
  Commit-Message `256f3c7` („verstoss gegen Hard Rule 3.4")
- `befund`: `AGENTS.md` §3.4 trägt zwei Aussagen (sprach-/meilensteinfrei;
  kein Spec-Stratum referenziert abwärts) — die Historie-Klausel steht dort
  nicht, und über die Git-Historie von `AGENTS.md` nie. slice-185 hat das
  selbst so vermerkt: „nicht durch §3.4 gedeckt — dort steht nichts über
  Historie-Abschnitte". Die dritte Klausel steht im Hard-Rule-Block der
  Architektur-Ziel-Form und in modul-03; die Streich-Entscheidung selbst trägt
  in der Substanz, aber der genannte Anker verifiziert die Norm nicht — in
  einem Plan, der mit der Closure einfriert.
- `verifizierbar`: ja — `grep -n "Historie" AGENTS.md` (nur die §3.7-Begründung
  „die Historie in `git`") und `git log -p -- AGENTS.md | grep -c "keine
  Historie"` (0).
- `klasse`: falsche Fundstellen-Attribution

### F-4 — Die Lastenheft-Änderung trägt keinen CHANGELOG-Eintrag in `[Unreleased]`

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §6, Schritt 7 — „Ein Slice, der einen öffentlichen
  Vertrag ändert — Lastenheft, Spezifikation, Benutzerhandbuch oder eine
  Regel —, trägt seinen Eintrag in `[Unreleased]` **in sich**"; „Kein Gate
  deckt das"
- `pfad`: `CHANGELOG.md`:7; Commit `256f3c7`
- `befund`: `256f3c7` ändert das [Lastenheft](../../spec/lastenheft.md) (Rang 1
  der Source Precedence); der `[Unreleased]`-Block des
  [CHANGELOG](../../CHANGELOG.md) führt nur den slice-199-Eintrag (Baseline-
  Hebung) — der Status-Flip auf `Accepted` samt seiner kontraktlichen Folge
  (ab jetzt ist jede Änderung am Dokument eine Vertragsänderung) ist nicht
  eingetragen.
- `verifizierbar`: ja — `grep -n "slice-189" CHANGELOG.md` (0 Treffer).
- `klasse`: öffentlicher Vertrag ohne CHANGELOG-Eintrag

### F-5 — Die LH-RB-Befundzelle nennt die Hermetik als nicht formalisiert; sie ist als AC-QA-02 formalisiert

- `kategorie`: LOW
- `quelle`: `AC-QA-02`
  ([lastenheft.md](../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
  „läuft **netzlos** (`--network none`) im distroless/static-Image")
- `pfad`: `slice-189` Plan §3.1, Zeile „`LH-RB-NN`-Randbedingungen-Reihe"
  (`docs/plan/planning/in-progress/slice-189-voll-abgleich-spec-straten.md`:87)
- `befund`: Die Befundzelle sagt, a-checks Randbedingungen „(Docker/make-only,
  19 Grandfathers, Hermetik)" seien „als Anforderungen nicht formalisiert".
  Docker/make-only (Hard Rule §3.1) und die 19 Grandfathers
  (`harness/rules/ac-form.md`:9) sind keine Lastenheft-Anforderungen —
  die Hermetik dagegen ist als AC-QA-02 formalisiert. Die Einstufung
  „Change Request (benannt, nicht ausgeführt)" trägt an den beiden ersten
  Beispielen; das dritte verifiziert nicht.
- `verifizierbar`: ja — `grep -in "grandfather" AGENTS.md` (0 Treffer),
  `grep -in "grandfather" harness/rules/ac-form.md` (Zeile 9, „19"),
  AC-QA-02-Wortlaut im Lastenheft §4.
- `klasse`: Befund-Begründung nennt falsches Gegenstück

### F-6 — Der Plan trägt zwei Abschnitte `## 3. Umsetzung`; der Platzhalter widerspricht der gefüllten Klassifizierung

- `kategorie`: LOW
- `quelle`: Maintainability (Plan-Struktur; Zustand ist das Feld, nicht die
  Dopplung)
- `pfad`: `docs/plan/planning/in-progress/slice-189-voll-abgleich-spec-straten.md`:80
  und :108–110
- `befund`: Commit `a138a88` fügt die Klassifizierung als zweiten Abschnitt
  `## 3. Umsetzung` (Z. 80) ein; der gleichlautende Platzhalter-Abschnitt
  (Z. 108, „*(entsteht mit der Arbeit)*") bleibt stehen — der Plan hat zwei
  Abschnitte Nummer 3, und der zweite behauptet über die Klassifizierung
  denselben Zustand („entsteht") als sei sie ungeschrieben.
  `make doc-structure` meldet 0 Befunde — der Defekt ist unsensor-covered.
- `verifizierbar`: ja — `grep -n "^## 3. Umsetzung"` auf die Plan-Datei (zwei
  Treffer); `make doc-structure` (Exit 0, 0 Befunde).
- `klasse`: zwei Quellen für denselben Zustand

---

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Lastenheft-Inhalt (Diff `256f3c7`) | geprüft, ohne Befund — der Diff ändert genau eine Zeile (Status); kein bestehendes `AC-*` ist inhaltlich geändert; `make doc-immutable RANGE=HEAD~2..HEAD` grün (Exit 0, 0 Befunde) |
| Architektur-Klauseln 1 + 2 (§3.4) | geprüft, ohne Befund — `spec/architecture.md` trägt keine Wellen-, Slice-, Commit-Hash-, Closure- oder ADR-Referenz (grep; die zwei Treffer sind die negierten Regel-Nennungen selbst) |
| Zitate auf den gestrichenen §8 | geprüft, ohne Befund — `make doc-check` grün (627 Dateien, 0 Befunde); grep nach verwaisten §8-/Anker-Referenzen leer; der §7-Risiko-Ausgang 3 ist damit belegt |
| Frische-Marker (Klausel 3) | geprüft, ohne Befund — „Letzte Änderung" steht am Kopf (`spec/architecture.md`:7) mit git-Verweis, Form wie §3.3 entscheidet; die Versions-Tabelle (vier Einträge) ist entfernt |
| Kandidaten-Zahlen §1/§2 | geprüft, ohne Befund — 28/28/20 stehen so in slice-187 §2 (Archiv-Tabelle „Paar / Kandidaten / Stand" gelesen); „56" = 28 + 28 (die zwei Spec-Straten); die Neu-Messung gegen `v6.13.0` ist als slice-192-Lektion benannt |
| „21 Anforderungen geprüft" | geprüft, ohne Befund — `make doc-trace`: 21 Anforderungen, 0 Waisen, alle `ok` |
| Klassifizierungs-Stichprobe (8 Zeilen + §3.3) | geprüft, ohne Befund bei 6 von 8 — Schema→MR-000 (existiert, deklariert `AC-FA-*`/`AC-QA-*`), Suffix-Form→MR-027 (existiert, löst MR-022 ab, sagt exakt das Behauptete), Metriken→MR-014 (existiert; null Metriken-Treffer in `spezifikation.md`), Kategorie-Gliederung (Defaults: 13 Treffer; Exit-Codes in den `SPEC-*`-Sektionen vorhanden), LH-RB-Reihe (Template trägt `LH-RB-01`, Lastenheft keine RB-Zeile — Ausgang CR korrekt), §7-Historie-Zeilen beider Straten vorhanden; Befund-Einwände in F-1 und F-5 |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 4 |
| LOW | 2 |
| INFO | 0 |

**Finding-Klassen dieses Laufs:** unbelegte Tatsachenbehauptung ·
Geltungsbereich einer Messung nur in eine Richtung genannt ·
falsche Fundstellen-Attribution · öffentlicher Vertrag ohne
CHANGELOG-Eintrag · Befund-Begründung nennt falsches Gegenstück ·
zwei Quellen für denselben Zustand

## Verdikt

**Merge-blockierend:** nein — die Commits liegen bereits auf `main`, und die
Substanz des Ranges trägt: die §8-Streichung samt Frische-Marker setzt die
dritte Klausel korrekt um (Negativbefunde 2–4), kein `AC-*` ist inhaltlich
geändert, und die Klassifizierung ist an der Stichprobe mit einer Ausnahme
(F-5) belegt.

**Closure-blockierend für slice-189: ja** — vor dem Übergang nach `done/`
sind zu klären: der Abnahme-Beleg für den Status-Flip oder die Zurücknahme auf
`Draft`/`In Review` (F-1), die Nachklassifizierung der repo-seitigen
Abweichungen oder die ausdrückliche Benennung der Messrichtung als Geltungs-
bereich (F-2), die Fundstelle in §3.3 (F-3), der CHANGELOG-Eintrag (F-4) und
der Doppel-Abschnitt (F-6). Die Finding-Klassen gehen in die Slice-Closure §7
und von dort in den Zähler.
