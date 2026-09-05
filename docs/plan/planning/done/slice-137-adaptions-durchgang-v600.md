# slice-137 — Etappe B: Adaptions-Durchgang gegen `v6.0.0`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Etappe **B** aus [slice-135 §6](../done/slice-135-regelwerk-v600-delta-analyse.md#6-vorschlag-drei-etappen),
Vorgänger [slice-136](../done/slice-136-baseline-v600-vendoring.md) (Etappe A). Präzedenz:
[slice-095](../done/slice-095-adaptions-durchgang-v5120.md), dieselbe Etappe B für den vorigen
Sprung. [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Konventions-Urteil ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Auslöser und Korrektur des Ausgangswerts

`modul-02` §Freshness-Audit verlangt den Durchgang durch die Adaptions-**Liste**, nicht nur durch
den Diff, mit fünf Ausgängen je Eintrag: *gegenstandslos* · *bleibt gültig* · *teilweise überholt*
· *Bezug entfallen* · *widerspricht*. Dieser Mechanismus selbst ist zwischen `v5.12.0` und `v6.0.0`
**unverändert** — `modul-02`s einzige Änderung in diesem Sprung betrifft die
Docker-Hermetik-Regel eines generierten Fragments, nicht das Freshness-Audit.

**Korrektur einer eigenen Zahl.** [slice-135 §4.1](../done/slice-135-regelwerk-v600-delta-analyse.md#4-die-echten-brocken)
und [slice-136](../done/slice-136-baseline-v600-vendoring.md) sprechen wiederholt von „16 aktiven
`MR`-Einträgen". Das ist falsch: `harness/conventions.md` §Aktive Adaptionen führt **sechs**
([`MR-011`](../../../../harness/conventions/MR-011-verfeinerungs-form.md)…[`MR-016`](../../../../harness/conventions/MR-016-validator-unbesetzt.md));
16 ist die **höchste je vergebene** Kennung — zehn davon
([`MR-001`](../../../../harness/conventions/done/MR-001-spezifikations-schicht.md)…[`MR-010`](../../../../harness/conventions/done/MR-010-rueckbau-drei-adaptionen.md))
liegen bereits in `conventions/done/`. Beide Slices liegen in `done/` und werden nicht editiert
(dieselbe Append-only-Disziplin wie bei den Adaptionen selbst, [`AGENTS.md`](../../../../AGENTS.md)
§3.5 sinngemäß); die Korrektur steht hier, wo sie auffiel.

## 2. Betroffene Module

`harness/conventions.md` §Adaptions-Block, urteilend; ausgeführt wird — falls überhaupt nötig —
in einer eigenen Etappe C (siehe §4).

## 3. Die sechs Urteile

Jede Zeile nennt die Baseline-Regel, gegen die geurteilt wurde — Datei und Abschnitt, mit
Zeilen-Diff zwischen `v5.12.0` und `v6.0.0` **direkt nachgerechnet** (`diff` gegen die betroffene
Sektion allein, nicht gegen die ganze Datei — eine Datei kann an anderer Stelle wachsen, ohne dass
die zitierte Sektion sich bewegt).

| Eintrag | Ausgang | Baseline-Regel, an der es hängt |
|---|---|---|
| [MR-011](../../../../harness/conventions/MR-011-verfeinerungs-form.md) | **bleibt gültig** | `grundlagen-source-precedence.md` §ID-Schema als Klammer (Zeilen 252–330): Sektion **byte-identisch** zwischen beiden Ständen |
| [MR-012](../../../../harness/conventions/MR-012-referenzmatrix-grandfathering.md) | **bleibt gültig** | `grundlagen-referenz-richtung.md`: die ganze Datei ist **byte-identisch** bis auf den `Quelle:`-Tag-Kommentar (Link-Rewrite) |
| [MR-013](../../../../harness/conventions/MR-013-adr-vorlagen-version.md) | **teilweise überholt** | trägt keine `Ersetzt-Baseline-Regel` (Rückbau-Kandidat, siehe eigener Text); sein eigener Auflösungs-Trigger — „die Überarbeitung der ID-Schema-Deklaration **oder die nächste Baseline-Migration**, je nachdem was zuerst eintritt" — ist mit diesem Sprung **eingetreten**. `docs/plan/adr/adr.template.md` selbst ist zwischen den Ständen **unverändert** (nicht im 26-Dateien-Delta aus slice-135 §2) — der Nachfolge-Eintrag ist eine reine Versions-Referenz-Korrektur, keine inhaltliche Änderung |
| [MR-014](../../../../harness/conventions/MR-014-keine-agenten-telemetrie.md) | **bleibt gültig** | `modul-15-observability.md`: die ganze Datei ist **byte-identisch** bis auf den `Quelle:`-Tag-Kommentar |
| [MR-015](../../../../harness/conventions/MR-015-welle-closure-ohne-replay.md) | **bleibt gültig** | `modul-06-roadmap.md` §Wellen-Closure-Prozedur Schritt 1 („Trigger prüfen … `make gates` und der Replay-Lauf sind grün") — Zeile **wortgleich**, nur um 40 Zeilen verschoben (die Sektion davor wuchs um den neuen Archivierungs-Schritt); die inhaltlich große Änderung dieses Moduls (Beobachtungs-Register-Neugestaltung, sechster Closure-Schritt) berührt eine **andere** Teil-Sektion und diesen Eintrags Gegenstand nicht |
| [MR-016](../../../../harness/conventions/MR-016-validator-unbesetzt.md) | **bleibt gültig** | `modul-08-agentenrollen.md` §Die neun Übergaben und ihre Artefakte (Zeilen 99–126): Sektion **byte-identisch**; die geänderte Sektion dieses Moduls ist §Rollen-Sequenz für eine Welle (Schritt-Umnummerierung 5→6), eine andere |

**Verteilung:** 0× gegenstandslos · **5× bleibt gültig** · 1× teilweise überholt · 0× Bezug
entfallen · 0× widerspricht.

**Deutlich ruhiger als beim letzten Durchgang** ([slice-095](../done/slice-095-adaptions-durchgang-v5120.md):
3 gegenstandslos, 2 bleibt gültig, 5 teilweise überholt bei zehn Einträgen). Kein Zufall: `v6.0.0`
ist additiv (neue Konzepte — Beobachtungs-Register-Verzeichnisform, Zeitdokumente-Archivierung,
Docker-Hermetik-Vorschrift, Gate/Beleg-Rollentrennung), keine der sechs a-check-Adaptionen liegt in
einem der neu bearbeiteten Abschnitte. `v5.12.0` hatte dagegen mehrfach denselben Boden wie a-checks
eigene Adaptionen neu betreten (Struktur-IDs als Default, `nicht-deterministischer Kern` in
`modul-12`, Validator-Kanten-Klarstellung in `modul-08`) — das war die Konvergenz, die diese
Adaptionen teilweise eingeholt hat. Dieser Sprung tut das an keiner Stelle.

## 4. Ein zweiter Fund: die Kürzel-Frage aus slice-135 löst sich **nicht** in dieser Etappe

[slice-135 §4.1](../done/slice-135-regelwerk-v600-delta-analyse.md#4-die-echten-brocken)
hatte offengelassen, ob a-check schon heute „eine Kennungsklasse mit Bereichssegment" führt (was
die neue Kürzel-Spalte in der Modus-Deklaration sofort verpflichtend machen würde) — mit Verweis
auf `AC-FA-<BEREICH>-<NNN>`. Gelesen gegen [`grundlagen-source-precedence.md` §Vergabe: woher die
nächste Nummer kommt](../../../../.harness/baseline/v6.0.0/regelwerk/grundlagen-source-precedence.md#vergabe-woher-die-nächste-nummer-kommt)
(neu in `v6.0.0`, aber bereits identifiziert als Ziel dieser Prüfung) löst sich die Frage eindeutig
**gegen** die in slice-135 vermutete Kopplung:

- Das `BEREICH`-Segment in `AC-FA-<BEREICH>-<NNN>` ist eine **Struktur-ID** — sie lebt in *einer*
  Datei (`spec/lastenheft.md`), Kollisionen sind im selben Diff sichtbar. Die Baseline sagt
  darüber ausdrücklich: „Sie brauchen deshalb **kein** Bereichssegment" — der `BEREICH`-Teil dort
  ist eine **inhaltliche** Kategorisierung, nicht das kollisionsvermeidende Kürzel, um das es in der
  Modus-Deklaration geht.
- Das kollisionsvermeidende Kürzel ist ausschließlich für Artefakte **mit je eigener Datei** gedacht
  (ADR, Slice, Carveout, `MR`) — und dort **optional**: „Ein Repo mit einem schreibenden Menschen
  braucht kein Segment — dichte Nummern sind dort billiger und lesbarer." a-check ist genau dieser
  Fall (Baseline-Geltungsbereich: „ein schreibender Mensch plus Agenten je Repository").
- Die zitierte Zwangsbedingung „seit die Kennung einer Beobachtung der Pfad `BEO-<KUERZEL>/<slug>`
  **ist**" hängt exakt an der **Umsetzung** der neuen Beobachtungs-Register-Form — die ist Etappe C,
  nicht B. Vor Etappe C führt a-check **keine** Kennungsklasse mit Bereichssegment, und die
  Kürzel-Spalte bleibt gestrichen.

**Der Etappen-Schnitt aus slice-135 §6 war an dieser Stelle also zu weit gefasst** — „Etappe B:
Adaptions-Durchgang **+ Kürzel-Vergabe**" benannte eine Abhängigkeit, die erst mit Etappe C selbst
entsteht, nicht vor ihr. Korrigiert wird das hier, nicht in slice-135 (liegt in `done/`, wird nicht
überschrieben): **Kürzel-Vergabe ist Teil von Etappe C**, ein eigener erster Schritt dort, keine
Vorbedingung dieser Etappe.

## 5. Auszuführende Gates

`make gates` — ändert nur ein Planning-Dokument, aber der Gate-Nachweis hängt am
Inhalts-Hash des Arbeitsbaums. Zum Abschluss `make verify`.

**Kein neuer Sensor**, also keine Negativ-Probe: geliefert werden Urteile, kein Zustand, den ein
Gate hält.

## 6. Was bewusst nicht getan wird

- **Kein Rückbau, keine Nachfolge-Einträge für [`MR-013`](../../../../harness/conventions/MR-013-adr-vorlagen-version.md).** Ein Eintrag wird nie überschrieben; der
  Nachfolger dafür ist ein eigener Folge-Slice, keine Pflicht dieser Etappe — er ist klein
  (reine Versions-Referenz, siehe §3) und kann mit Etappe C zusammenfallen oder eigenständig laufen.
- **Kein Rückbau der fünf `bleibt-gültig`-Einträge.** Sie ändern sich nicht — es gibt nichts
  aufzulösen.
- **Keine Kürzel-Vergabe** (§4) — Etappe C.
- **Die Adaptionen bleiben bis zur Ausführung in Kraft** — Etappe A hat das in `harness/conventions.md`
  §Baseline bereits so deklariert (der doppelte Vendoring-Zustand während der Migration).

## 7. DoD

- [x] Alle sechs aktiven Adaptionen einzeln beurteilt, je gegen die **konkrete zitierte Sektion**
      (nicht die ganze Datei) diff-geprüft (§3).
- [x] Die Kürzel-Frage aus slice-135 §4.1 geklärt und der zu weit gefasste Etappen-Schnitt korrigiert
      (§4); die „16"-Falschzahl aus slice-135/136 richtiggestellt, ohne die `done/`-Dateien zu
      editieren (§1).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 8. Closure-Notiz

**Geliefert:** sechs Urteile (fünf *bleibt gültig*, ein *teilweise überholt*) mit je einer
diff-geprüften Sektion als Beleg — deutlich ruhiger als der vorige Durchgang, weil `v6.0.0` additiv
neue Konzepte einführt statt bestehende a-check-Adaptionen einzuholen. Dazu zwei Korrekturen an
eigenen früheren Aussagen dieser Migration: eine Zählungenauigkeit (16 statt 6 aktive `MR`) und ein
zu weit gefasster Etappen-Schnitt (Kürzel-Vergabe gehört nicht vor, sondern in Etappe C).

**Lerneintrag — Form: geschärfte Regel.** *„Die Datei hat sich geändert" und „die zitierte Sektion
hat sich geändert" sind zwei verschiedene Aussagen, und nur die zweite entscheidet über eine
Adaption — ein Diff gegen die ganze Datei überschätzt systematisch, welche Adaptionen betroffen
sind.* Konkret: `modul-06-roadmap.md` wuchs in diesem Sprung um 140 Zeilen — ein naiver
„Datei geändert ⇒ Adaption prüfen, Ausgang vermutlich *teilweise überholt*"-Reflex hätte
[MR-015](../../../../harness/conventions/MR-015-welle-closure-ohne-replay.md) fälschlich als
betroffen eingestuft. Erst der Diff **der zitierten Sektion allein** zeigte: Schritt 1, an dem
der Eintrag hängt, ist wortgleich, nur um 40 Zeilen verschoben, weil ein *anderer* Abschnitt derselben
Datei wuchs. Dieselbe Falle stellte sich bei `modul-08-agentenrollen.md` für
[MR-016](../../../../harness/conventions/MR-016-validator-unbesetzt.md). *Weil* die Baseline-Dateien
mehrere Regeln pro Datei bündeln und eine Adaption nur an *einer* hängt, ist „Datei im Diff" ein
notwendiges, aber kein hinreichendes Kriterium — das hinreichende ist die Sektion.

**Zwei beobachtbare Closure-Kriterien:**

1. Jede Zeile der Tabelle in §3 ist mit dem genannten `diff`-Befehl gegen die beiden vendorten
   Bäume (`.harness/baseline/v5.12.0/` und `.harness/baseline/v6.0.0/`) nachrechenbar.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.**

- *[`MR-013`](../../../../harness/conventions/MR-013-adr-vorlagen-version.md) braucht einen Nachfolge-Eintrag (reine Versions-Referenz), noch ohne Kennung* —
  Ausgang: **Folge-Slice** — kann mit Etappe C zusammenfallen, da beide den Adaptions-Speicher
  anfassen, oder eigenständig laufen; welche Reihenfolge, entscheidet die nächste Slice-Planung.
- *Die fünf `bleibt-gültig`-Urteile sind nicht durch einen Sensor, sondern durch manuellen `diff`
  belegt* — Ausgang: **gestrichen mit Begründung**: kein Risiko im Sinn der Dreier-Menge, sondern
  dieselbe Beleglage wie bei der Präzedenz [slice-095](../done/slice-095-adaptions-durchgang-v5120.md),
  die ebenfalls manuell diffte; ein Sensor für „ändert sich eine zitierte Baseline-Sektion" wäre
  baubar, aber ohne einen dritten Beleg dafür, dass genau **das** wiederholt auffällt, verfrühtes
  Werkzeug.

**Folge-Slices:** keine vergeben. Der [`MR-013`](../../../../harness/conventions/MR-013-adr-vorlagen-version.md)-Nachfolger und Etappe C sind benannt (§8, slice-135 §6),
brauchen aber eigene Kennungen bei ihrer eigenen Eröffnung.

## 9. Sub-Area-Modus

Berührt wird ausschließlich **Harness-Einstieg** (`AGENTS.md`, `harness/conventions.md` — die
Adaptions-Urteile, keine Schreibung) und **Planungs-Harness** (`docs/plan/planning/`) — Greenfield.
