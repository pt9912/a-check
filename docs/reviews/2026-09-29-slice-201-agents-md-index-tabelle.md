# Review-Report: slice-201 — 2026-09-29

**Review-Art:** Code-Review (Doku-Refactoring) — geprüft gegen den Slice-Plan
und die Ziel-Form (`v6.13.0` · `templates/AGENTS.template.md` §5, Welle 149)
sowie `harness/conventions.md` §Baseline (Modul 10 §Drei Review-Arten).

**Gegenstand:** Commit-Range `86c92ac..0106785` (`HEAD~2..HEAD`), zwei
Commits: `6288479` „§5 als Index-Tabelle, 14 Regeln ausgelagert" und
`0106785` „slice-201 Closure — 14/14 Wortfolgen identisch". `86c92ac` (reiner
`git mv` open → in-progress) ist die unberührte Basis und trägt den
Altstand von `AGENTS.md` §5, gegen den gemessen wurde.

**Skill:** `.harness/skills/reviewer.md` @ `0106785` ·
**Modell:** glm-5.3-flash · **Datum:** 2026-09-29

> **Zitier-Form** *(dieser Block bleibt stehen — er ist Norm, kein
> Ausfüll-Hinweis; die `<Platzhalter>` darin sind Formbeispiele)*. Dieser
> Report friert ein; was er zitiert, bewegt sich
> weiter. Deshalb: **Kennung, nicht Adresse** — `slice-<Kennung>` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link
> (`v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt>). Der vendored Baum trägt
> genau einen Tag; der Sprung löscht den alten, und ein Link darauf färbt beim
> nächsten Bump ein Artefakt rot, das niemand mehr anfassen darf. Ein
> `pfad`-Feld auf den **geprüften Gegenstand** ist davon nicht betroffen — es
> zitiert den Stand des Laufs und darf ihn festhalten.

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde — ohne diese
Liste ist der Lauf nicht reproduzierbar):

- `slice-201` (Slice-Plan, Stand `0106785`: `docs/plan/planning/in-progress/slice-201-agents-md-index-tabelle.md`)
- `v6.13.0` · `templates/AGENTS.template.md` §5 (Index-Tabelle, Welle 149)
- `harness/conventions.md` §Baseline (adoptierter Stand `v6.13.0`)
- `AGENTS.md` §3 (Hard Rules), Altstand §5 aus `git show 86c92ac:AGENTS.md`
- `.d-check.yml` (`structure`-, `citations`-Bezüge der neuen Tabelle)
- `.harness/skills/reviewer.md` §Mess-Regeln (Bezug auf `AGENTS.md` §5)

**Geltungsbereich der eigenen Messungen dieses Laufs** (Mess-Regel 1):
alle Zahlen unten messen **Wortfolgen außerhalb von Markdown-Link-Zielen**
(Whitespace normiert) und die **Anker-/Zeilen-Mengen** der genannten
Dateien — nichts darüber. Link-Ziele sind nicht Wortfolge, sondern separat
geprüft: auflösend per `make doc-check`, semantisch per Stichprobe
(`../../` aus `harness/rules/*.md` → Repo-Root). `make gates`/`make verify`
hat dieser Lauf **nicht** selbst ausgeführt (Verifier-Sache); selbst
ausgeführt sind `make doc-check` und `make doc-structure` (beide Exit 0,
623 Dateien, 0 Befunde).

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | MEDIUM | §8 des Slice-Plans deklariert für die Beobachtung `BEO-HARNESS/agents-md-hinkt-baseline-dod-item-hinterher` „Befund wird gezählt"; die Closure-Notiz §7 meldet „keine Beobachtung angefallen", und `evidence/` enthält nur `slice-164.md`. Eine der beiden Aussagen ist damit falsch: Entweder zählt der Vorgang als Auftreten — dann fehlt der Evidence-Beleg und der abgeleitete Zähler steht zu niedrig — oder er zählt nicht — dann bleibt die ungekorrigierte §8-Aussage im Archiv stehen. | `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register (Zähler = Zahl der gültigen Evidence-Dateien; „Eingetragen wird bei der Slice-Closure") | `docs/plan/planning/in-progress/slice-201-agents-md-index-tabelle.md`:130–132 gegen :115 | ja — `ls docs/plan/planning/observations/BEO-HARNESS/agents-md-hinkt-baseline-dod-item-hinterher/evidence/` (nur `slice-164.md`); `state.md`: „Stand: offen" | Register-Sichtung und Closure-Meldung zum selben Eintrag widersprüchlich |
| F-2 | LOW | Der Wortfolgen-Vergleich-DoD wurde im Closure-Commit selbst umformuliert: vorab „die Menge der Aussagen von §5 (Regeln, Begründungen, Durchsetzungs-Zusagen) ist … unverändert — gemessen, nicht angenommen", danach „14/14 ausgelagerte Regeln Wortfolgen-identisch (normierte Zähler …)". Die Prüfbasis wandert damit mit dem Lauf; die engere Fassung benennt ihren Geltungsbereich allerdings sauber und ist durch die eigene Probe dieses Laufs gedeckt (drei Kurzregeln wortidentisch, 14/14 ausgelagert identisch — 17/17). | `AGENTS.md` §6 Schritt 4 (Plan-Änderung gehört vor den Code); `slice-201` §2 | `docs/plan/planning/in-progress/slice-201-agents-md-index-tabelle.md`:49–52 (Diff `0106785`, DoD-Hunk) | ja — `git show 0106785` zeigt die DoD-Umformulierung | DoD-Text im Abschluss-Commit geändert statt vorab |
| F-3 | LOW | Tabellen-Zeile 9 komprimiert „drei Pfade (Happy/Boundary/Negative) plus Out-of-Scope" zu „drei Pflicht-Bausteine" — die Zählung der Kurzform weicht von Volltext und `structure`-Hint („Happy, Boundary, Negative **und** Out-of-Scope") ab; Zeile 12 verwischt die Zweier-Schwelle („Vorfall → Eintrag" statt „Ab dem **zweiten** gleichartigen Vorfall"). Der Zeiger resolved korrekt; die Kurzform allein gelesen, zählt bzw. schwellt anders als ihre Quelle. | `v6.13.0` · `templates/AGENTS.template.md` §5 (Kurzform-Zeiger-Logik); `slice-201` §2 (Form-Ziel) | `AGENTS.md`:223 (Zeile 9), `AGENTS.md`:226 (Zeile 12) | ja — `grep -n 'drei Pflicht-Bausteine' AGENTS.md` gegen `harness/rules/ac-form.md`; Hint in `.d-check.yml` (`structure`, AC-Form) | Kurzform-Zeiger trägt andere Zählung/Schwelle als ihr Volltext |
| F-4 | INFO | Die geschärfte Regel der Closure („Wortfolgen-Probe vor dem Commit — und die Linktiefe zweimal prüfen") trägt keinen Zielort und keinen Register-Eintrag; sie lebt ausschließlich in der Closure-Notiz, die mit dem Slice archiviert wird (Volltext ins ZIP, Stub bleibt). Gleiche Konstellation wie in `slice-167` — dort ebenfalls ohne Verkörperung akzeptiert —, deshalb INFO und kein MEDIUM. | `v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Closure- und Lerneintrag-Regeln | `docs/plan/planning/in-progress/slice-201-agents-md-index-tabelle.md`:97–101, :112–113 | ja — `grep -rn 'Linktiefe' --include='*.md' .` liefert nur die Closure-Notiz | Geschärfte Regel ohne Zielort nur in Closure-Notiz |
| F-5 | INFO | Die Form-Regel der Index-Tabelle (Kurzregeln vollständig in der Tabelle, Volltext nach `harness/rules/<name>.md`) steht im `<!-- -->`-Kommentar des neuen §5 und im vendorten Template-Kommentar — im lesenden Bestand trägt sie kein Fließtext-Stück. Die Tabelle exemplifiziert die Form selbst und folgt der Ziel-Form wörtlich, darum nur INFO; beim nächsten §5-Umbau ist der Kommentar die einzige repo-eigene Adresse der Regel. | HIGH-Klasse „Norm nur im Template-Kommentar" (Skill) — hier mit zweitem, repo-eigenem Kommentarstand | `AGENTS.md`:207–211 | ja — `sed -n '205,212p' AGENTS.md` | Form-Norm nur im Kommentarblock |
| F-6 | INFO | Zwei Altlasten wurden wortgetreu aus Alt-§5 konserviert und sind **nicht** Neubefund dieses Slices: (a) MR-019 nennt als Geltungsbereich „§5 (Slice-Form, DoD-Posten-Liste)" — die DoD-Posten-Liste stand schon vor `slice-201` nicht mehr in §5; (b) `harness/rules/steering-loop.md` trägt den Link-Text `docs/plan/steering-loop.md` auf das Ziel `docs/plan/planning/observations/README.md` — schon im Altstand so. Die Verbatim-Pflicht des Slices schloss die Korrektur bewusst aus. | `harness/conventions/MR-019-review-dod-opt-in.md` (Geltungsbereich); Verbatim-Zusage in `slice-201` §3 | `harness/rules/steering-loop.md`:7; `harness/conventions/MR-019-review-dod-opt-in.md`:5 | ja — `git show 86c92ac:AGENTS.md` (Altstand identisch) | Bestands-Drift unverändert konserviert |

## Proben und ihre Geltungsbereiche

- **Substanz-Probe (zweiter Zähler, anders gebaut):** Alt-§5 (`git show
  86c92ac:AGENTS.md`) wurde per `awk` in 17 Bullets zerlegt, die 14
  Auslagerungs-Dateien per Kopf-/Anker-Strip in reinen Volltext, beide
  Whitespace-normiert, Link-Ziele auf einen Platzhalter normiert und
  paarweise verglichen: **14/14 identisch**. Die drei im Rest verbleibenden
  Bullets (Alt-Bullets 4–6) sind mit Tabellen-Zeilen 4–6 wortidentisch.
  Geltungsbereich: Wortfolgen und Aussagen-Mengen; Link-Ziele ausgenommen
  (separat geprüft, s. u.), HTML-Kommentare und neue H1-/Anker-Zeilen der
  Auslagerungs-Dateien ausgenommen (die sind Zugewinn, nicht Bestand).
- **Zähler-Gegenprobe (dritte Zählung, anderes Kriterium):** 17 Alt-Bullets
  = 17 Tabellen-Zeilen (`grep -cE '^\| [0-9]+ '`) = 14 Herkunfts-Anker
  „Regel N" in `harness/rules/*.md` (N = 1, 2, 3, 7–17, je genau einmal)
  + 3 Kurzregeln mit `—` in der Datei-Spalte. Alle drei Zähler weichen
  nicht ab.
- **Link-Probe:** `make doc-check` Exit 0 — 623 Dateien, 0 Befunde
  (bestätigt die „0 Befunde über 623 Dateien"-Behauptung aus §6 der
  Closure). Semantische Stichprobe: `../../` aus `harness/rules/*.md`
  löst auf Repo-Root auf (`../../AGENTS.md`, `../../harness/conventions.md`
  samt Anker `#modus-deklaration-pro-sub-area` — Anker per `doc-check`
  geprüft, `../../.harness/baseline/v6.13.0/...`, `../../.harness/skills/...`).
- **Umfang-Probe:** `git show 6288479 -- AGENTS.md` trägt genau einen Hunk
  (`@@ -204,151 +204,31 @@`) — nur §5 berührt; §4 und §6 byte-identisch.
  Dateibestand des Commits: `AGENTS.md` + 14 neue `harness/rules/*.md` —
  exakt Plan §3; der Closure-Commit berührt nur die Slice-Datei. Out-of-Scope
  (§1) wurde nicht stillschweigend geweitet.
- **Form-Probe:** Spaltenköpfe `| # | Regel | Datei |` wie Ziel-Form;
  Kurzregeln in Zeilen 4–6 tragen `—`; jede ausgelagerte Zeile nennt in
  Regel-Spalte und Datei-Spalte dieselbe `harness/rules/<name>.md`; jede
  Auslagerungs-Datei trägt den Herkunfts-Anker „Ausgelagert aus
  `AGENTS.md` §5, Regel N — seit slice-201".
- **Nicht ausgeführt:** `make gates`/`make verify` (Docker-Volläufe) —
  die Grün-Meldung des Closures ist hier nicht unabhängig wiederholt
  worden; `make doc-check` und `make doc-structure` sind es
  (beide Exit 0).

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Substanz-Vergleich Alt-§5 ↔ `harness/rules/*.md` + Kurzregeln | geprüft, ohne Befund (17/17 Wortfolgen identisch) |
| Substanz-Vergleich Alt-§5 ↔ Tabellen-Kurzformen | geprüft, ohne Befund — mit Ausnahme F-3 (Zählung/Schwelle in Zeilen 9 und 12) |
| `make doc-check` / `make doc-structure` | geprüft, ohne Befund (Exit 0, 623 Dateien, 0 Befunde je Lauf) |
| Diff-Umfang gegen Plan §1/§3 (Out-of-Scope-Disziplin) | geprüft, ohne Befund |
| Referenzen auf `AGENTS.md` §5 im Bestand (`harness/README.md` Zeilen 94/98/101, `.d-check.yml`-Kommentare, `harness/sensors/verify-risiko-ausgaenge.md`, `docs/plan/planning/README.md`, `harness/conventions.md` MR-019/023) | geprüft, ohne Befund — alle verweisen auf §5 als Ganzes; keine zitiert Zeilennummern oder Absatz-Muster (Begründung des gestrichenen Risikos in §6 trägt) |
| CHANGELOG `[Unreleased]` | geprüft, ohne Befund — Regelwortlaut unverändert, Praxis seit `slice-186`: §5-Reformierungen ohne eigenen Eintrag |
| `.harness/skills/reviewer.md` §Mess-Regeln-Bezug („die Zusage selbst steht in `AGENTS.md` §5") | geprüft, ohne Befund — die Zusage steht als Zeiger-Zeile 15 weiter in §5; der Volltext ist eine Ebene weiter (F-4-Verwandtschaft, kein eigener Befund) |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 1 |
| LOW | 2 |
| INFO | 3 |

**Finding-Klassen dieses Laufs:** Register-Sichtung und Closure-Meldung zum
selben Eintrag widersprüchlich · DoD-Text im Abschluss-Commit geändert statt
vorab · Kurzform-Zeiger trägt andere Zählung/Schwelle als ihr Volltext ·
Geschärfte Regel ohne Zielort nur in Closure-Notiz · Form-Norm nur im
Kommentarblock · Bestands-Drift unverändert konserviert

## Verdikt

**Merge-blockierend:** ja — für den **Übergang nach `done/`**, nicht für die
Konversion. F-1 (MEDIUM) betrifft allein die Register-Buchführung des
Slice-Dokuments und ist vor dem `git mv` zu klären: entweder
`evidence/slice-201.md` nachtragen (dann zählt das Auftreten) oder §8 auf
die Closure-Aussage hin korrigieren. Die geprüften Konversions-Artefakte
selbst — `AGENTS.md` §5 und die 14 Dateien unter `harness/rules/` — sind
ohne Befund: Wortfolgen 17/17 identisch, Links auflösen semantisch richtig,
Form entspricht der Ziel-Form, Umfang exakt Plan §3.

**Übergabe:** Findings gehen an den Implementer (F-1 vor dem `git mv`, F-2
und F-3 nach Wahl des Implementers als Nachtrag oder Folge-Slice); die
**Finding-Klassen** gehen zusätzlich in die Slice-Closure §7 und von dort in
den Zähler. Dieser Report selbst ist ein **Lauf-Beleg** — er wird über Läufe
hinweg nicht wieder gelesen, und muss es nicht.
