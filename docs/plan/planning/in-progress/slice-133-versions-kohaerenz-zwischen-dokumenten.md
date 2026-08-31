# slice-133 — Eine Aussage über ein anderes Dokument wird gegen dessen Kopf gehalten

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Prosa-Drift in [`docs/user/releasing.md`](../../../../docs/user/releasing.md)
§Versionsquelle, gefunden bei der Release-Prep zur nächsten Version.
[`BEO-023`](../observations.md) (*Prüfer meldet grün ohne Gegenstand*) ist die Risiko-Klasse, gegen
die dieser Slice sich selbst absichern muss — nicht sein Anlass.

**Berührte Spec-Stellen:** — *(keine)* — die Doku-Konfiguration ist nicht Gegenstand des
Lastenhefts.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-31.

---

## 1. Ziel

Eine Versions-Angabe, die ein Dokument über ein **anderes** eigenes Dokument macht, fällt auf,
wenn sie nicht mehr stimmt — statt so lange richtig auszusehen, bis jemand zufällig nachschlägt.

## 2. Definition of Done

- [ ] Das Modul `versions` ist in [`.d-check.yml`](../../../../.d-check.yml) konfiguriert und läuft
      über `modules:` in `make doc-check` — damit **ohne neues Target** im `gates`-Aggregat.
- [ ] Die Angabe in [`releasing.md`](../../../../docs/user/releasing.md) §Versionsquelle nennt den
      **gemessenen** Stand des Lastenhefts, nicht den zuletzt geschriebenen.
- [ ] [`AGENTS.md`](../../../../AGENTS.md) §4 und
      [`harness/README.md`](../../../../harness/README.md) nennen die neue Fähigkeit **samt ihrer
      Grenze** (kein Digest).

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Der Befund, und warum ihn nichts gefangen hat

[`releasing.md`](../../../../docs/user/releasing.md) §Versionsquelle sagte *„Das Lastenheft steht
bei **0.17.0**"*. Der Kopf von [`spec/lastenheft.md`](../../../../spec/lastenheft.md) trägt
`**Version:** 0.26.0` — **neun** Minor-Stände Abstand.

Kein Gate sah dorthin, und zwar aus einem benennbaren Grund: `gate-consistency` hält die
Versions-Nummer von `version.md#aktuell` gegen den CHANGELOG (slice-018), also **a-checks eigene
Release-Achse**. Die Lastenheft-Achse ist eine **zweite**, ausdrücklich davon verschiedene — der
Satz sagt das selbst (*„eigene Doku-Achse, ≠ der Release-Tag"*) — und hatte keinen Wächter.

## 4. `versions` statt `citations` — gemessen, nicht angenommen

Zwei Module des Schwester-Tools kommen in Frage. Beide wurden gegen denselben Fall gehalten:

| Modul | Mechanik | Verhalten bei Drift |
|---|---|---|
| `citations` | `<!-- d-check:cite <pfad>:<von>-<bis> -->` markiert ein wortgleiches Zitat | `citation-mismatch` — **aber**: Zitate unter 16 Zeichen bleiben ungeprüft, und die Zeilen-Referenz `:3-3` bricht, sobald sich der Lastenheft-Kopf verschiebt |
| `versions` | `pin-pattern` fängt die Angabe, `current-from` liefert den Erwartungswert aus einem `Datei#Anker`-Span | `version-stale` — nennt den **erwarteten** Wert, braucht keine Zeilennummern |

Die 16-Zeichen-Schwelle ist der Grund für die Wahl, nicht ein Detail: eine Direktive über
`„0.17.0"` (sechs Zeichen) sieht wie ein Wächter aus und **schweigt** — gemessen im selben Lauf,
in dem die Langform meldete. Das ist [`BEO-023`](../observations.md) in seiner ersten Form
(*ohne Gegenstand*), erzeugt durch die eigene Konfiguration.

**Die Grenze von `versions`, ebenfalls gemessen:** der Erwartungswert kommt **versions-förmig** aus
dem Span, unabhängig davon, was die Capture-Gruppe fängt. Ein `sha256:`-Digest als `pin-pattern`
liefert `version-stale` gegen die Versionsnummer daneben; ein Span **ohne** Versionsnummer bricht
mit `versions.current-from: keine Version im adressierten Span`, **Exit 2**. Die Digest-Gleichheit
der harten Pins kann das Modul also **nicht** übernehmen — sie bleibt bei
`tools/gate-consistency.sh`.

## 5. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz geschrieben.

**Abgrenzung:** Dieser Slice legt **ein** Muster-Quellen-Paar an. Weitere Prosa-Angaben über
fremde oder eigene Versionen (der `d-check`-Pin, die Baseline `v5.12.0`) bleiben ungedeckt; sie
sind je eine eigene Messung, kein Anhängsel.

## 6. Risiken und offene Punkte

- **Das Muster ist an den Wortlaut des Satzes gebunden.** Wer ihn umformuliert, leert die
  Prüfmenge — und `versions` meldet dann grün, weil es nichts findet.
  **Ausgang:** *weiter offen* → [`BEO-023`](../observations.md), Klasse *ohne Gegenstand*. Der
  Zähler bleibt bei **2×**: ein dritter **Vorfall** ist nicht eingetreten, dieser Slice hat die
  Leermenge im Gegenteil ausgeschlossen (§4, Lauf A und B). Einen Zähler ohne Vorfall zu erhöhen
  wäre genau die Schein-Genauigkeit, die [`AGENTS.md`](../../../../AGENTS.md) §5 verbietet.
- **`current-from` hängt am Heading-Slug `#lastenheft--a-check`.** Wird die H1 des Lastenhefts
  umbenannt, zeigt der Anker ins Leere.
  **Ausgang:** *entfallen* — d-check bricht diesen Fall **fail-closed** ab (Exit 2, „keine Version
  im adressierten Span" bzw. `anchor-missing` über das Modul `anchors`). Ein Bruch, der auffällt,
  ist kein offenes Risiko.
- **Das Modul kann die Digest-Gleichheit nicht tragen**, der Eigenbau bleibt also stehen.
  **Ausgang:** *entfallen* — gemessen (§4) und an beiden Orten benannt, an denen jemand das
  Gegenteil vermuten würde ([`AGENTS.md`](../../../../AGENTS.md) §4 und der Kommentar über dem
  `versions:`-Block).

## 7. Closure-Notiz

*(wird vor dem `git mv` nach `done/` ausgefüllt)*

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** Berührt ist **Gate-/Werkzeug-Schicht** (die
`.d-check.yml`-Konfiguration und die Sensor-Deklarationen). Die Schwelle ≥ 2 von 3 Achsen ist
erfüllt: eigener Konventions-Ort ([`harness/conventions.md`](../../../../harness/conventions.md)
§Modus-Deklaration), eigene Artefakt-Menge (`.d-check.yml`, `Makefile`, `tools/`), eigene
Review-Frage (*deckt der Sensor, was er zu decken behauptet?*). Die berührte Doku
(`releasing.md`, `AGENTS.md`, `harness/README.md`) ist **Mitzieh-Fläche**, keine zweite Sub-Area:
ohne die Konfiguration gäbe es dort nichts zu schreiben.

**Vorgelagert — offene Beobachtungen sichten:** Zwei Treffer.
[`BEO-023`](../observations.md) (*Prüfer meldet grün ohne kalibriert zu sein*, **2×**) trifft
diesen Slice unmittelbar und ist in §4/§6 verarbeitet — beide Richtungen sind am **echten**
Bestand gemessen, nicht nur an einer Fixture. [`BEO-028`](../observations.md) (*CHANGELOG
`[Unreleased]` wird nicht je Slice gepflegt*, **1×**) ist bei der Release-Prep, aus der dieser
Slice stammt, ein **zweites** Mal eingetreten: 21 Commits seit `v0.18.0` ohne Eintrag. Der Zähler
wird bei der Closure erhöht.

### Sub-Area: Gate-/Werkzeug-Schicht

- **Modus:** GF — die Konventionen stehen vor dem Code, `.d-check.yml` ist deklarativ.
- **Konventionen-Dichte:** hoch. Die Sensor-Deklaration ist in
  [`AGENTS.md`](../../../../AGENTS.md) §4 und
  [`harness/README.md`](../../../../harness/README.md) §Sensors doppelt verankert, und
  `make gate-consistency` erzwingt, dass die Modulliste des netzlosen `doc-check` intakt bleibt.
- **Phase-Reife:** Phase 4 — die Schicht trägt zwölf Gates, zwei Aggregate und einen Meta-Sensor;
  neue Sensoren entstehen hier als Konfiguration, nicht als Erstaufbau.
- **Evidenz-/Diskrepanz-Risiko:** niedrig, **weil gemessen**. Das Muster hat genau **einen**
  Gegenstand im Bestand (`releasing.md:38`), und der `current-from`-Span trägt genau **ein**
  versions-fähiges Token. Ohne diese beiden Zählungen wäre das Risiko hoch — das ist der
  Unterschied, den [`BEO-023`](../observations.md) benennt.
- **Reconciliation-Aufwand:** keiner. Kein Bestand wird nachgezogen; die einzige Diskrepanz, die
  der Sensor findet, ist die, die ihn ausgelöst hat, und sie wird in diesem Slice behoben.
