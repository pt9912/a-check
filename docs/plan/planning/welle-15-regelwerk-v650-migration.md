# Welle welle-15: Regelwerk-Migration `v6.2.0` → `v6.5.0`

**Lifecycle:** Diese Datei entsteht bei der **Eröffnung** der Welle und liegt
flach unter `docs/plan/planning/`; bei Closure wandert sie per `git mv` nach
`done/` (neben ihre `welle-15-results.md`). Der Zustand ist die
Verzeichnis-Position — kein Status-Feld.

**Zielmeilenstein:** kein Meilenstein-Bezug.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Datum:** 2026-09-07.

---

## 1. Welle-Ziel

a-check auf den Kurs-Stand `v6.5.0` heben. Der Sprung überspringt drei
Releases (`v6.3.0`, `v6.3.1`, `v6.4.0`) und ist **um eine Größenordnung
größer** als der vorige: **30 Dateien, +572/−149** gegen `v6.0.0`→`v6.2.0` mit
9 Dateien, `+53/−5` (gemessen 2026-09-07, `git diff --stat v6.2.0 v6.5.0 --
lab/regelwerk lab/templates` im frischen Klon).

Vier Etappen, deren Zuschnitt die Delta-Analyse schärft:

1. **Messen** — was hat sich geändert, was davon berührt a-check.
2. **Vendoring** — der neue Stand unter `.harness/baseline/`, Stand-Deklaration
   an ihren drei Stellen, Symlinks unter `.claude/rules/`. *(Der Zielpfad steht
   hier bewusst nicht als Literal: er existiert erst nach dieser Etappe, und
   `versions` meldete ihn zu Recht als `version-stale` — erster echter Treffer
   des Sensors aus slice-173, drei Minuten nach seiner Inbetriebnahme.)*
3. **Adaptions-Durchgang** — die **sieben** aktiven `MR`-Einträge gegen
   `v6.5.0`, nicht nur die vom Diff berührten.
4. **Neue Ziel-Form `harness/sensors/<target>.md`** — je Sensor eine Datei,
   sobald sein Vertrag mehr als einen Satz braucht. Trifft a-check hart: die
   Tabellenzellen in [`AGENTS.md`](../../../AGENTS.md) §4 und
   [`harness/README.md`](../../../harness/README.md) §Sensors sind genau der
   Überhang, den die Vorlage benennt.

## 2. Trigger (Welle startet)

- `v6.5.0` ist im Kurs-Repo veröffentlicht — bestätigt am 2026-09-07 über
  `git ls-remote --tags`; die Tag-Liste führt `v6.3.0`, `v6.3.1`, `v6.4.0`
  und `v6.5.0` über dem adoptierten `v6.2.0`.
- Der Maintainer hat die Migration angewiesen und die Wellen-Form gewählt
  (dieses Gespräch, 2026-09-07).

## 3. Closure-Trigger (Welle schließt)

Das **Mehr** gegenüber den einzelnen Slice-DoDs: die Stand-Deklaration ist
repo-weit widerspruchsfrei, und der Sensor, der das erzwingt, läuft grün gegen
den **neuen** Stand — das prüft keine einzelne DoD.

- Alle Slices dieser Welle liegen in `done/`.
- Die Stand-Deklaration nennt an allen drei Stellen `v6.5.0`
  ([`conventions.md`](../../../harness/conventions.md#baseline) §Baseline,
  [`AGENTS.md`](../../../AGENTS.md) §1,
  [`harness/README.md`](../../../harness/README.md) §Guides), und
  `make doc-check` meldet **keinen** `version-stale` — der Wächter aus
  slice-173 fährt hier seinen ersten echten Migrations-Lauf.
- `make symlink-check` grün: die vier Baseline-Symlinks unter `.claude/rules/`
  tragen den neuen Stand. Ebenfalls erster echter Einsatz.
- `make gates`, `make verify` und `make ci` je Exit 0 auf dem finalen Stand —
  Ausgabe in eine Datei, Exit-Code getrennt gelesen, nie in eine Pipe.
- Ergebnis-Notiz `done/welle-15-results.md` geschrieben.

## 4. Slices in dieser Welle

| Slice | Titel | Bezug |
|---|---|---|
| slice-174 | Delta-Analyse `v6.2.0` → `v6.5.0` | — (reine Ist-Messung, keine Vertragsberührung) |

Weitere Slices entstehen aus §5 der Delta-Analyse; die Zeile wächst mit. Der
Zustand eines Slice ist sein Lifecycle-Verzeichnis und wird hier **nicht**
gespiegelt.

## 5. Abhängigkeiten

- Wird blockiert von: nichts. `welle-14` ist geschlossen, `in-progress/` war
  bei Eröffnung leer.
- Blockiert: [slice-169](open/slice-169-korpus-seitige-kalibrierung.md) nicht
  formal, aber inhaltlich berührt — die Korpus-seitige Kalibrierung prüft
  Phrasen, die diese Welle in den betroffenen Dokumenten verschiebt. Reihenfolge
  ist eine Planungs-Frage, keine Sperre.

## 6. Out-of-Scope für diese Welle

- **Kein Rückbau bestehender Adaptionen ohne eingetretenen Trigger.** Der
  Durchgang *bewertet* die sieben aktiven Einträge; er löst nur auf, wo die
  Bedingung des Eintrags selbst eingetreten ist.
- **Keine Umstellung auf `harness/sensors/` als Selbstzweck.** Etappe 4 legt
  Dateien nur dort an, wo der Vertrag eines Targets die Zelle sprengt — die
  Vorlage nennt das ausdrücklich als Bedingung, nicht als Pflicht je Target.
- **Kein Nachziehen der übersprungenen Zwischenstände.** Gemessen und vendored
  wird `v6.5.0`; `v6.3.0`/`v6.3.1`/`v6.4.0` erscheinen nur, wo die Delta-Analyse
  eine Änderung einem dieser Increments zuordnen muss.
- **Keine inhaltliche Änderung an Zeitdokumenten**, um den neuen Stand zu
  spiegeln. Ein `done/`-Slice zitiert, was damals galt (slice-173 §2.1).

## 7. Closure-Notiz

*(nach Welle-Abschluss ausfüllen.)*

Ergebnis: `welle-15-results.md`
Zähler: [`observations/`](observations/README.md)
