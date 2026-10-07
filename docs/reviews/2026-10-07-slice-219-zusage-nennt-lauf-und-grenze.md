# Review-Report: slice-219 — 2026-10-07

**Review-Art:** Harness-/Doku-Review, unabhängiger Lauf. Der Kontext ist ein frischer Subagent
ohne `fork`, und er hat den Gegenstand nicht verfasst. Geprüft wird der Diff gegen den Slice-Plan
(einschließlich der Plan-Änderung in §1), die Hard Rules (`AGENTS.md` §3, besonders §3.7), die
Mess-Regeln und die Register-Form (`v6.13.0` · `regelwerk/modul-06-roadmap.md` §Das
Beobachtungs-Register, Paarung a). Gegen die DoD wird **nicht** geprüft, das ist Aufgabe des
Verifiers (Modul 11).

**Gegenstand:** slice-219, Commit-Range `94ae6b1..e9e3170` — Inhalt in `e9e3170`
(docs(harness): fünfte Mess-Regel); davor `fa5e423` (Plan-Änderung), `94ae6b1` (Ruhe-Marker).

**Skill:** `.harness/skills/reviewer.md` @ `e9e3170` (sha256 `7e9a9629…26beab36`) ·
**Modell:** `claude-opus-5-5` · **Datum:** 2026-10-07. Der Skill ist selbst Teil des
Gegenstands; angewandt wurde der Stand nach dem Diff, geprüft wurde er wie jeder andere Text.

> **Zitier-Form** *(dieser Block bleibt stehen. Er ist Norm, kein
> Ausfüll-Hinweis)*. Dieser Report friert ein, was er zitiert, bewegt sich
> weiter. Deshalb gilt **Kennung, nicht Adresse**: `slice-<Kennung>` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link. Ein
> `pfad`-Feld auf den **geprüften Gegenstand** ist davon nicht betroffen. Es
> zitiert den Stand des Laufs und darf ihn festhalten.

**Eingangs-Kontext:**

- Slice-Plan slice-219 (§1 Ziel, Plan-Änderung, Abgrenzung; §3; §6; §8)
- `BEO-GATE/zusage-weiter-als-ihre-durchsetzung`: `observation.md`, `state.md`, sechs
  Evidence-Dateien (slice-186, slice-187, slice-216, slice-217, slice-218, slice-220)
- Review-Reports der Belege aus den Archiven entpackt (Scratchpad, Repo unberührt): slice-217
  (F-1, F-2), slice-218 (F-1), slice-220 (F-4); dazu der Report zu slice-215 (Vorgänger-Regel 4,
  frühere Findings am selben Bereich)
- `AGENTS.md` §3, §4, §5 Zeile 15, §6 Schritt 7; `harness/README.md` §Sensors;
  `CHANGELOG.md` `[Unreleased]`
- `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register

**Sonden:**

- `make doc-check`: Exit 0, „725 Datei(en) geprüft, 0 Befund(e)" — Geltungsbereich: Links und
  Anker; Prosa-Aussagen sieht das Instrument nicht.
- `make verify-observations`: Exit 0, „83 Beobachtung(en) mit nicht leerem evidence/" —
  Geltungsbereich: Deckung Register ↔ Zitate; Zielort und Herkunfts-Anker (Paarung a) prüft es
  nicht, die habe ich gelesen.
- **Zählstellen, zwei verschieden gebaute Zähler** (Mess-Regel 3), je über das Repo ohne
  `docs/plan/planning/done/`, `.harness/baseline/` und `docs/reviews/`:
  Zähler A — Zahlwort (`vier|fünf|drei|4|5`) direkt vor `Mess-Regeln|Urteilsregeln`;
  Zähler B — jede Zeile mit `mess-regel`, gefiltert auf ein Zahl- oder Mengenwort
  (`vier|fünf|drei|alle|übrig|2–5`) in derselben Zeile; Zähler C — `alle vier|den drei/vier
  übrigen|Regel 4|Regeln 1–4` über `harness/`, `AGENTS.md`, Skills, Register, Planning-Lifecycle
  ohne `done/`, `README.md`, `CHANGELOG.md`, `docs/user/`. A findet `AGENTS.md:229`,
  `mess-regeln.md:1`/`:5`, `reviewer.md:114`; B zusätzlich die beiden `state.md`; C zusätzlich
  `reviewer.md:221` („Alle vier Belege" — zählt Belege von Regel 4, nicht Regeln; bleibt wahr) und
  `kandidaten-klassifikation-groeber-als-der-kandidat/state.md:34` („die zwei älteren
  Mess-Regeln" — relativ, bleibt wahr). Alle Zählstellen der Regel-Anzahl sind nachgezogen.
- Herkunfts-Anker: `seit slice-219` steht in `harness/rules/mess-regeln.md:29` (und `:3`) und
  `.harness/skills/reviewer.md:226`.

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | MEDIUM | Regel 5 greift laut Wortlaut, wenn jemand schreibt, „dass etwas geprüft, verglichen oder durchgesetzt wird". Zwei der für Ausprägung 3 angeführten Belege sind keine solchen Sätze: slice-216 („derselbe Commit ergibt denselben Index-Digest", „byte-identische Ausgabe auf beiden Plattformen") sind Eigenschafts-Zusagen eines Vertrags, die keine Prüfung behaupten; slice-218 („jeder Schritt läuft über make oder die Docker-CLI") beschreibt die Zusammensetzung der Pipeline. Auch `observation.md` definiert die Klasse als „sagt zu, dass eine Regel **geprüft** wird". Die Herleitung sagt „Drei Ausprägungen, alle belegt" und „die Regel verlegt die Frage vor den Review" — für diese Belege feuert die Regel beim Schreiben nach ihrem Wortlaut nicht. | Mess-Regel 4/5 auf die Regel selbst angewandt; Evidence slice-216, slice-218 | `harness/rules/mess-regeln.md:30-33`; `.harness/skills/reviewer.md:237-240`, `:246-248` | nein — Urteil über Wortlaut gegen Beleg | Regel reicht weniger weit als ihre Belege |
| F-2 | MEDIUM | Regel 5 verlangt die Grenze „**am selben Satz**", und Ausprägung 2 benennt „die Grenze steht woanders als die Zusage" als Fehlerform. `AGENTS.md` §4 schickt dagegen ausdrücklich die Grenze eines Targets, die nicht in die Zelle passt, nach `harness/sensors/`, und vier Gate-Index-Zeilen tragen genau diese Form („Grenzen: siehe Datei" bei `make doc-check`, `make gate-consistency`, `make verify-review-haken`, `make verify-trigger-audit`). Ob ein Zeiger am Satz die Regel erfüllt, sagt weder die Regel noch die Herleitung; je nach Lesart sind diese Zeilen regelkonform oder Exemplare von Ausprägung 2. | `AGENTS.md` §4; Regel 5 | `harness/rules/mess-regeln.md:32`; `.harness/skills/reviewer.md:234-236`; `AGENTS.md:173-174`; `harness/README.md:71`, `:76`, `:97`, `:98` | ja — `grep -n "siehe Datei\|siehe Sensor-Datei" harness/README.md` | Zwei Regeln gleichen Rangs ohne Abgrenzung |
| F-3 | LOW | „slice-220: „die Ausgabe des Image-Tests"" steht in Zitatzeichen, ist aber kein Wortlaut des Belegs: Der Handbuch-Satz lautete „seine Ausgabe muss der auf `linux/amd64` gleichen" (Evidence slice-220, Report slice-220 F-4); „die Ausgabe des Image-Tests" ist die Formulierung des Reviewers. Kleiner: „„jeder Schritt über make oder die Docker-CLI"" lässt das „läuft" des Originals aus. | Reviewer-Skill §Mess-Regeln (Belegtreue) | `.harness/skills/reviewer.md:233-234`, `:239-240` | ja — `grep` des Zitats in Evidence und Archiv-Report | Paraphrase in Zitatform |
| F-4 | LOW | „slice-186: … geprüft wird nur, ob der Ruhe-Marker fehlt" ist enger als der Beleg: geprüft wird die Äquivalenz in beide Richtungen („Slice vorhanden ⟺ Ruhe-Marker fehlt", Evidence slice-186; Gate-Index `make doc-planning`: „ist das Verzeichnis leer, steht er"). | Mess-Regel 1 (Geltungsbereich) | `.harness/skills/reviewer.md:231-232` | ja — Evidence slice-186 gegen Skill | Paraphrase enger als ihr Beleg |
| F-5 | LOW | Zusage und Herleitung weichen für den Fall „kein Lauf" ab: Die Regel sagt „Gibt es den Lauf **noch nicht**, sagt der Satz das" (zeitlich, Lauf kommt); die Herleitung sagt „Findet sich kein Lauf, ist der Satz eine Absicht und sagt es, oder er fällt". Der dauerhafte Fall ohne Lauf — wie §3.7 oder die Mess-Regeln selbst („Kein Sensor") — steht nur in der Herleitung, die der Schreibende laut Skill-Kopf nicht liest. | Skill §Mess-Regeln Kopf („Zwei Adressaten, zwei Orte") | `harness/rules/mess-regeln.md:33`; `.harness/skills/reviewer.md:229`, `:242-243` | nein — Lesart | Zusage und Herleitung weichen im Wortlaut ab |
| F-6 | LOW | `state.md` trägt „Sechs Belege." in der `Stand:`-Zeile — einen geführten Zähler neben dem `evidence/`-Verzeichnis. Das Regelwerk sagt „Es gibt kein Feld, in das man ihn schreibt"; ein weiterer Beleg nach der Verkörperung (Prosa-Form ausgeschöpft?) ließe die Zeile still falsch werden. Derselbe Slice nimmt aus dem Nachbar-Zustand eine Zahl genau aus diesem Grund heraus. Vorbestand mit derselben Form: `testbeschreibung-weiter-als-assertion/state.md` („Vier Belege"). Ausgang *verkörpert*, Zielort, Herkunfts-Anker und Paarung (a) sind sonst korrekt. | `v6.13.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register („Der Zähler wird abgeleitet, nicht geführt") | `docs/plan/planning/observations/BEO-GATE/zusage-weiter-als-ihre-durchsetzung/state.md:1` | ja — `ls evidence/ \| wc -l` gegen die Zeile | Zähler im Register geführt statt abgeleitet |
| F-7 | LOW | §6 des Slice-Plans sagt „drei der fünf Fälle fing der Review"; die Herleitung sagt „Fünf der sechs Fälle fand der unabhängige Review". Gegen die Evidence stimmt die zweite (slice-186 F-7, slice-216, slice-217, slice-218, slice-220 per Review; slice-187 beim Abgleich selbst). Die Plan-Datei wurde in der Range bearbeitet (`fa5e423`), die Zahl blieb stehen. | Mess-Regel 3 | `docs/plan/planning/in-progress/slice-219-zusage-nennt-lauf-und-grenze.md:81` | ja — Evidence-Dateien zählen | Zahl neben einer Aufzählung nicht nachgezogen |
| F-8 | LOW | Titel und `AGENTS.md` §5 Zeile 15 sagen, die Mess-Regeln binden, „wer einen Beleg schreibt". Regel 5 bindet, wer eine Zusage „in Doku, Kommentar, Vertrag oder Gate-Index" schreibt — einen weiteren Adressatenkreis. Wer einen Gate-Index-Satz oder Handbuch-Satz schreibt, erkennt sich im Zeiger nicht als Adressat; der Zeiger ist damit enger als die Regel, auf die er zeigt. Die Plan-Änderung begründet die Verortung („eine Zusage über eine Prüfung ist eine Aussage über einen Beleg"), der Titel trägt diese Lesart nicht. | `AGENTS.md` §5 Zeile 15; Regel 5 | `AGENTS.md:229`; `harness/rules/mess-regeln.md:1`, `:5`, `:30-31` | nein — Urteil | Zeiger enger als die Regel, auf die er zeigt |
| F-9 | INFO | Regel 5 überschneidet sich mit Regel 1: Ausprägung 1 („nennt einen weiteren Gegenstand als den, den der Prüfer sieht") ist Regel 1s „sagt, ob er den Gegenstand deckt", übertragen von einer einmaligen Messung auf eine stehende Prüf-Zusage. Gegenstand (Zusage statt Messung) und Pflicht (Lauf benennen) unterscheiden sich — eine Schärfung, kein Duplikat. Die Herleitung von Regel 4 nennt ihre Verwandtschaft zu Regel 2 ausdrücklich, die von Regel 5 nennt Regel 1 nicht. | Skill §Mess-Regeln | `.harness/skills/reviewer.md:226-250` | nein | — (Hinweis) |
| F-10 | INFO | Die Plan-Änderung nennt „mit drei verschieden gebauten Suchen gefunden", ohne die Suchen oder ihren Geltungsbereich zu nennen; das Ergebnis ist durch die Sonden oben bestätigt. Einen CHANGELOG-Eintrag oder seine Begründung („Mess-Regeln binden den Harness-Lauf, nicht den Konsumenten", wie in slice-215) führt der Plan nicht; `AGENTS.md` §6 Schritt 7 nennt „eine Regel" unter den öffentlichen Verträgen. | Mess-Regel 1; `AGENTS.md` §6 Schritt 7 | `docs/plan/planning/in-progress/slice-219-zusage-nennt-lauf-und-grenze.md:33-39` | nein | Begründung für fehlenden CHANGELOG-Eintrag nicht notiert |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| (1) Tatsachen der Herleitung gegen die sechs Evidence-Dateien und Archiv-Reports | geprüft: „Register-Eintrag bei 6×" (sechs Dateien), „Fünf der sechs Fälle fand der unabhängige Review" (stimmt — slice-187 fand der Abgleich selbst), slice-216 (Plattform-Gleichheit ohne Träger, Reproduzierbarkeit durch Gegenmessung widerlegt), slice-217 (Array-Feld, dritte Plattform; Gate-Index nannte Pipeline-Verwendung vor der Verdrahtung in slice-218 — Report slice-217 F-1), slice-187 (Zitat „analog zur ADR-Immutabilität", `paths`-Liste) treffen. Befunde nur F-1, F-3, F-4 |
| (2) Wortlaut der Regel: anwendbar, sagt selbst nicht mehr zu als sie hält | Regel und Herleitung sagen „kein Sensor" und behaupten keinen Lauf — Regel 5 auf sich selbst angewandt ohne Befund. Befunde zur Reichweite F-1, F-2, F-5 |
| (3) Zählstellen | drei verschieden gebaute Zähler, alle Zählstellen der Regel-Anzahl nachgezogen; ohne Befund |
| (4) `state.md` Form und §3.7 | Ausgang *verkörpert*, Form `liegt in <Zielort>`, Herkunfts-Anker `seit slice-219`; Paarung (a): beide Zielorte tragen den Anker. Keine Chronik — die alte Erzählung („Beide bisherigen Fälle sind so behoben worden", „Beim dritten Mal …") ist entfernt. Nachbar-`state.md` ohne Zahl, bleibt wahr. Befund nur F-6 |
| (5) Überschneidung Regel 1/Regel 5 | Schärfung, kein Duplikat; F-9 |
| (6) Außerhalb von §1 | Diff berührt genau die Dateien der Plan-Tabelle §3 nach Plan-Änderung (`fa5e423` liegt vor `e9e3170`, `AGENTS.md` §6 Schritt 4); kein Sensor, kein Nachzug alter Zusagen, kein Code. Ohne Befund |
| Hard Rules §3.1–§3.6 | kein Code, kein Move mit Inhalt, keine ADR berührt, kein Gate gelockert; Commit-Scope `(planning)` in `fa5e423`/`94ae6b1` nur unter `docs/plan/planning/`; jede Message nennt slice-219. Ohne Befund |
| §3.7 in den neuen Texten | Regel, Herleitung und `state.md` im Indikativ über den Zustand; Herkunfts-Angaben in einem Feld. Ohne Befund |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 2 |
| LOW | 6 |
| INFO | 2 |

**Wiederkehrende Klassen:** *Paraphrase in Zitatform* (F-3) und *Paraphrase enger als ihr Beleg*
(F-4) traten beide schon im Review zu slice-215 auf (F-3, D-2) — am selben Abschnitt, an der
Herleitung der Vorgänger-Regel. *Zahl neben einer Aufzählung nicht nachgezogen* (F-7) trägt die
Klasse aus dem Review zu slice-218 (F-9). Gezählt wird bei der Closure, nicht hier.

## Verdikt

**Abnahme-blockierend:** ja, wegen F-1 und F-2 (MEDIUM): Sie sind vor der Closure zu klären. Beide
betreffen die Reichweite des Regel-Wortlauts, nicht seine Herleitung. Bei F-1 passen Regel und zwei
ihrer Belege nicht zusammen. F-2 betrifft das Verhältnis zu `AGENTS.md` §4. Wenn der Implementer
F-2 als Konflikt zweier Regeln widerspricht, ist das eine Architect-Frage; in dem Fall läuft der
Konflikt-Pfad über den Architect. Die LOW-Findings sind Belegtreue und Zahlen; die Tatsachen der
Herleitung tragen sonst.

**Übergabe:** Findings an den Implementer; Finding-Klassen zusätzlich in die Closure §7 von
slice-219. Dieser Report ist Lauf-Beleg und ersetzt keine Verifikation gegen die DoD.
