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

Die Delta-Analyse ([slice-174](done/slice-174-regelwerk-v650-delta-analyse.md))
schneidet die Etappen; ihre Buchstaben sind hier und in §4/§6 dieselben:

- **A — Vendoring:** der neue Stand unter `.harness/baseline/`,
  Stand-Deklaration an ihren drei Stellen, Symlinks unter `.claude/rules/`,
  Kurs-Wellen-Stempel. *(Der Zielpfad steht bewusst nicht als Literal: er
  existiert erst nach dieser Etappe, und `versions` meldete ihn zu Recht als
  `version-stale` — erster echter Treffer des Sensors aus slice-173, drei
  Minuten nach seiner Inbetriebnahme.)*
- **B — Adaptions-Durchgang:** die **sieben** aktiven `MR`-Einträge gegen den
  neuen Stand, nicht nur die vom Diff berührten; dazu die offene ADR-Frage zu
  den `exempt-paths` aus slice-173.
- **C — Zitier-Form:** einfrierende Artefakte zitieren die Baseline als
  Kennung statt als Adresse.
- **D — Sensors-Struktur:** je Sensor eine Datei, sobald sein Vertrag mehr als
  einen Satz braucht, plus die zweite Tabelle für Nicht-Gates. Trifft a-check
  hart: **16 von 40** Gate-Zeilen in [`AGENTS.md`](../../../AGENTS.md) §4 sind
  länger als 250 Zeichen.
- **E — Slice-Form:** §1 *Ziel und Abgrenzung*, neuer §8-Titel, und die
  Schritt-Hälfte im Minimal Agent Workflow.

Die Messung selbst (slice-174) trägt keinen Buchstaben — sie ist die
Voraussetzung, nicht eine der Etappen.

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
| [slice-175](done/slice-175-etappe-a-vendoring-v650.md) | Etappe A: Baseline vendoren | slice-174 §3.4 |
| [slice-176](open/slice-176-zitier-form-einfrierende-artefakte.md) | Etappe C: Zitier-Form für einfrierende Artefakte | slice-174 §3.2 T-1 |
| [slice-177](open/slice-177-sensors-struktur-zwei-tabellen.md) | Etappe D: Sensors-Struktur, zwei Tabellen | slice-174 §3.2 T-2/T-4 |
| [slice-178](open/slice-178-slice-form-ziel-und-abgrenzung.md) | Etappe E: Slice-Form §1/§8 und Schritt 4 | slice-174 §3.2 T-3 |

**Etappe B** (Adaptions-Durchgang) bekommt ihren Slice, sobald Etappe A den
Stand gehoben hat — sie misst gegen ihn und wäre vorher gegenstandslos. Der
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
