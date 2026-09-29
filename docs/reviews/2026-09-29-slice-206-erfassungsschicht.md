# Review-Report: slice-206 — 2026-09-29

**Review-Art:** Plan + Code — geprüft gegen `AGENTS.md` §3/§5, die
Slice-Ziel-Form (`v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice),
die MR-Ziel-Form (`v6.13.0` · `templates/harness/conventions/MR-NNN-titel.template.md`)
und die Generator-Form (`ai-harness-init` v6.13.0, Referenzstand `/tmp/aih-v6.13.0/`).

**Gegenstand:** slice-206, committet auf `main` als `ab93a2c` —
`.claude/hooks/span-emit.sh`, `harness/erfassung-feldliste.md`,
`.claude/settings.json`, `harness/conventions.md`,
`harness/conventions/MR-030-erfassung-lokal-kein-abfluss.md`,
`harness/conventions/done/MR-014-keine-agenten-telemetrie.md`,
Slice-Plan `docs/plan/planning/in-progress/slice-206-erfassungsschicht.md`.

**Hinweis zum Stand:** Der Working Tree trägt zu drei Dateien **uncommittete**
Deltas (Slice-Plan §3-Zeile, `harness/conventions.md` MR-008-Zeile, interne
Links von MR-014), die gegen F-1 und F-3 gerichtet sind. Geprüft und hier
befundet ist der **committete** Stand `ab93a2c`; die `pfad`-Zeilen halten ihn
fest.

**Skill:** `.harness/skills/reviewer.md` @ `a6d19b6` (2026-09-29)
**Modell:** glm-5.3-flash · **Datum:** 2026-09-29

> **Zitier-Form** *(dieser Block bleibt stehen — er ist Norm, kein
> Ausfüll-Hinweis; die `<Platzhalter>` darin sind Formbeispiele)*. Dieser
> Report friert ein; was er zitiert, bewegt sich
> weiter. Deshalb: **Kennung, nicht Adresse** — `slice-<Kennung>` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link
> (`v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt>). Der vendored Baum trägt
> genau einen Tag; der Sprung löscht den alten, und ein Link darauf färbt beim
> nächsten Bump ein Artefakt rot, das niemand mehr anfassen darf. Ein `pfad`-Feld
> auf den **geprüften Gegenstand** ist davon nicht betroffen — es zitiert den
> Stand des Laufs und darf ihn festhalten (`v<X.Y.Z>` ·
> `regelwerk/grundlagen-harness-dateien.md` §harness/README.md als
> Einstiegspunkt — diese Zeile ist selbst ein Beispiel der Form).

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde — ohne
diese Liste ist der Lauf nicht reproduzierbar):

- Slice-Plan `slice-206` (committeter Stand, s. Hinweis oben)
- `AGENTS.md` §3 (Hard Rules) und §5; `.harness/skills/reviewer.md`
- `AC-QA-02` (Hermetik und ehrliche Heuristik-Grenze) aus `spec/lastenheft.md`
- `MR-014` (aufgelöst) und `MR-030` (Nachfolge-Eintrag) samt Ziel-Form
- `BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter` (Beobachtungs-Register)
- Generator-Form unter `/tmp/aih-v6.13.0/` (Wrapper, Feldliste, settings.json)
  und der Traeger-Artikel (sha256 `c6a6a171…e9186`)

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | §3 des committeten Plans nennt `harness/conventions/MR-014-…` als Änderungs-Art «Nachfolge-Eintrag» und linkt den Pfad vor dem Umzug; der tatsächlich erzeugte Nachfolge-Eintrag MR-030 taucht in §3 nicht auf — Plan-Tabelle und Implementierung sind auseinander. | `v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice | docs/plan/planning/in-progress/slice-206-erfassungsschicht.md:65 | ja — Lesen des committeten Stands (`git show ab93a2c:…`); kein Gate prüft Planinhalt gegen Diff | Plan benennt falsches Artefakt |
| F-2 | HIGH | Der vorgelagerte Sichtungs-Schritt meldet «keine Treffer in GATE», obwohl der Slice HARNESS als berührt führt und dort `BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter` bei 2× steht — dessen Gegenstand (Auflösungs-Trigger eines `MR`-Eintrags ohne Wächter) ist genau die hier vollzogene Auflösung von MR-014; als drittes Auftreten der Klasse würde der Eintrag die Schwelle erreichen und eine Lücke mit eigenem Folge-Slice begründen. | `v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Zwei Schritte vor der Modus-Begründung | docs/plan/planning/in-progress/slice-206-erfassungsschicht.md:101 | nein — inferentiell; das Register liest kein Gate | Sichtungs-Schritt lässt berührte Sub-Area aus |
| F-3 | MEDIUM | Der committete Stand trägt die Pfad-Berichtigung des MR-014-Umzugs nicht: die MR-008-Zeile der aufgelösten Tabelle linkt `conventions/MR-014-…` vor dem Umzug, und die internen Links von MR-014 (Vorgänger, Ersetzt-Baseline-Regel, AC-QA-03) stehen noch in der Tiefe von `harness/conventions/` — `make doc-check` (im `gates`-Aggregat) läuft darauf rot, bis die Berichtigung als eigener Commit nachgezogen ist. | `v6.13.0` · `templates/harness/conventions/MR-NNN-titel.template.md` (Umzug zieht Pfad-Berichtigung als eigener Commit nach sich) | harness/conventions.md:196 | ja — `make doc-check` | Umzug ohne Pfad-Berichtigung |
| F-4 | MEDIUM | MR-030 trägt «Löst auf: MR-014», aber nicht das damit Pflicht gewordene Feld «Ausgelöst durch Baseline-Stand»; es ist weder gesetzt noch ist ausdrücklich benannt, dass kein Baseline-Stand die Ablösung ausgelöst hat. | `v6.13.0` · `templates/harness/conventions/MR-NNN-titel.template.md` | harness/conventions/MR-030-erfassung-lokal-kein-abfluss.md:20 | nein — Form-Frage ohne Sensor; Review-Sache | Pflichtfeld der Ziel-Form nicht gesetzt |
| F-5 | LOW | Der Plan-Kopf deklariert «Lerneintrag — Form: neuer Sensor», der Slice legt aber keinen Sensor an; die übernommene Feldliste benennt ihre Grenzen ausdrücklich als «Grenzen, die kein Sensor hält». | `v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Closure- und Lerneintrag-Regeln | docs/plan/planning/in-progress/slice-206-erfassungsschicht.md:20 | ja — Closure-Prüfung (`closure-note-reviewer`) | Formfeld ohne Gegenstück im Planinhalt |
| F-6 | INFO | Die Begründung von MR-030 stützt die Auflösung allein auf die Maintainer-Freigabe; dass zugleich der eigene Auflösungs-Trigger des Vorgängers («sobald Agenten-Läufe im Repo selbst abrechenbar werden») mit der Token-Erfassung der Feldliste eingetreten ist, wird nicht benannt. | `v6.13.0` · `templates/harness/conventions/MR-NNN-titel.template.md` | harness/conventions/MR-030-erfassung-lokal-kein-abfluss.md:16 | nein — Urteil über die Begründung, kein Sensor | Auflösungs-Begründung ohne Vorgänger-Trigger |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Wrapper-Logik (`.claude/hooks/span-emit.sh` ab `set -uo pipefail` gegen `/tmp/aih-v6.13.0/.claude/hooks/span-emit.sh`) | geprüft, ohne Befund — byte-identisch; nur der Header ist adaptiert (Fremd-Kennungen `LH-*`/`ADR-*` entfernt, Anker auf `slice-206`, Feldliste und MR-030 gesetzt) |
| Wrapper-Verhalten (mit Traeger, ohne Traeger via `CLAUDE_PROJECT_DIR` auf leere Wurzel, ohne `CLAUDE_PROJECT_DIR` aus dem Dateiort) | geprüft, ohne Befund — je Exit 0, stdout und stderr stumm; Traeger ist gitignored, sha256 stimmt mit der genannten und mit dem Generator-Artikel überein |
| MR-Durchgang Tabellen (`harness/conventions.md` aktiv + aufgelöst, Anker, Immutabilität von MR-014, Wortlaut MR-030 gegen Freigabe-Lesart) | geprüft, ohne Befund — MR-030-Zeile aktiv mit Anker, MR-014-Zeile aufgelöst mit mitgezogenem Anker `#mr-014`, MR-014 ist 100 %-Rename (inhalte unverändert), Wortlaut deckt «lokal ja, draußen nein»; die Pfad- und Form-Befunde stehen als F-3/F-4/F-6 |
| Feldliste (`harness/erfassung-feldliste.md` gegen Generator-Form, Fremd-Kennungen, Gegenprobe-Zeile) | geprüft, ohne Befund — byte-identisch, keine `LH-*`/`ADR-00*`; die Gegenprobe-Zeile `.harness/state/spans/gegenprobe_2026_09_29.jsonl` validiert auf das Schema (15 Pflichtfelder als leere Aussage, 3 Optional-Felder besetzt) |
| `.claude/settings.json` (JSON-Valide, drei neue Events, Bestands-Hooks, Timeout) | geprüft, ohne Befund — JSON valide; PostToolUse, PostToolUseFailure, SubagentStart mit leerem Matcher und Timeout 5; PreToolUse-Guard und Stop unverändert (Diff gegen `ab93a2c^` ist rein additiv) |
| §3.7 Kommentar-Regeln (Wrapper-Header, Feldliste, MR-030) | geprüft, ohne Befund — Zusage-, Kopplungs- und Rang-Zeiger-Form; keine Chronik, kein abwesender Text |
| Risiko-Ausgang §6 (Platzhalter «bei Closure») | geprüft, ohne Befund — vorgesehener Zustand vor Closure; `make verify-risiko-ausgaenge` greift erst mit ausgefüllter §7 und verlangt dann einen Ausgang aus der geschlossenen Menge |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 2 |
| LOW | 1 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Plan benennt falsches Artefakt ·
Sichtungs-Schritt lässt berührte Sub-Area aus · Umzug ohne Pfad-Berichtigung ·
Pflichtfeld der Ziel-Form nicht gesetzt · Formfeld ohne Gegenstück im
Planinhalt · Auflösungs-Begründung ohne Vorgänger-Trigger

## Verdikt

**Merge-blockierend:** ja — die zwei HIGH-Findings stehen auf, und F-3 hält
zugleich `make gates` im committeten Stand rot (`doc-check` gegen die
dangling Links). Die uncommitteten Working-Tree-Deltas richten F-1 und F-3
bereits — sie sind zu committen, nicht der Befund zu verwerfen. F-2 ist vor
der Closure zu entscheiden: Benennt die Closure `BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter`
als drittes Auftreten, folgt daraus nach Regelwerk ein eigener Folge-Slice.

**Übergabe:** Findings gehen an den Implementer (Rückkante
Review → Plan bei Plan-Defekt, F-1); die **Finding-Klassen** gehen zusätzlich
in die Slice-Closure §7 und von dort in den Zähler. Dieser Report selbst
ist ein **Lauf-Beleg** (Audit: dieser Diff, dieser Skill, dieses Modell,
dieses Verdikt) — er wird über Läufe hinweg nicht wieder gelesen, und
muss es nicht. Der Report ersetzt keine
Verifikation — DoD-/Spec-Konformität prüft der Verifier separat
(Modul 11; anderes Prüf-Artefakt, anderer Eingabe-Kontext).
