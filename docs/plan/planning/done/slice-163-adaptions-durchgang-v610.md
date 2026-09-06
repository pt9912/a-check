# slice-163 — Regelwerk-Migration `v6.0.0` → `v6.1.0`: Adaptions-Durchgang (Etappe B)

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-14](welle-14-regelwerk-v610-migration.md).

**Bezug:** [slice-161](../done/slice-161-regelwerk-v610-delta-analyse.md)
(Delta-Analyse), [slice-162](../done/slice-162-review-pflicht-absatz-v610-wortlaut.md)
(fand den diff-only-blinden-Fleck), Maintainer-Entscheidung 2026-09-06
("Etappe B: Adaptions-Durchgang zuerst").

**Berührte Spec-Stellen:** — *(keine)* — reine Ist-Messung, keine
Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim
Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

> **Analyse zur Abnahme.** Wie slice-161: keine Kennungen vergeben, keine
> Artefakte geändert. Der Folge-Slice-Vorschlag (§6) gehört vor der
> Umsetzung abgenommen.

**Geltungsbereich der Quelle:** `lab/regelwerk/` und `lab/templates/` aus
einem frischen Klon von `pt9912/ai-harness-course`, Tags `v6.0.0` und
`v6.1.0` (identisches Verfahren wie slice-161).

---

## 1. Anlass

`slice-162` deckte auf, dass slice-161s diff-only-Vorgehen (nur die sechs
vom `git diff v6.0.0 v6.1.0` berührten Dateien geprüft) einen echten Fund
verfehlt hätte, wäre er nicht durch eine Rückfrage aufgefallen
(`AGENTS.md` §6 ohne `MR`-Deklaration — Belege, nicht der Diff selbst,
zeigten den Bruch). `welle-14` benennt daraufhin zwei zusätzliche Belege
für dasselbe Muster —
[`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
(2×) und
[`rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md)
(1×) — und die Maintainer-Entscheidung: vor Etappe A (Vendoring) einen
vollständigen Adaptions-Durchgang fahren, alle **18** MR-Dateien unter
[`harness/conventions/`](../../../../harness/conventions/) (7 aktiv, 11
aufgelöst) einzeln gegen den `v6.1.0`-Stand, nicht nur die vom Diff
berührten.

## 2. Verfahren

Für jede der 7 **aktiven** MR-Dateien: das Feld `Ersetzt-Baseline-Regel`
gelesen, die referenzierte Datei mit
`git diff v6.0.0 v6.1.0 -- <datei>` aus dem frischen Klon geprüft (nicht
nur "war sie im 6-Datei-Diff aus slice-161 §2 enthalten", sondern jede
Datei einzeln gegen den vollen Tag-Vergleich). Für die 11 **aufgelösten**
Dateien unter `conventions/done/`: geprüft, ob sie auf eine der drei laut
slice-161 §2 geänderten Regelwerk-Dateien (`modul-07`, `modul-10`,
`modul-13`) zeigen — eine aufgelöste Adaption wird nicht neu bewertet
(„nichts nachträglich inhaltlich geändert", `harness/conventions.md`
§Adaptions-Block), aber eine Überschneidung wäre meldenswert gewesen.

## 3. Ergebnis je aktiver MR

| MR | Ersetzt-Baseline-Regel (Datei) | `git diff v6.0.0 v6.1.0` auf diese Datei | Befund |
|---|---|---|---|
| [MR-011](../../../../harness/conventions/MR-011-verfeinerungs-form.md) | `grundlagen-source-precedence.md` | 0 Zeilen | unberührt |
| [MR-012](../../../../harness/conventions/MR-012-referenzmatrix-grandfathering.md) | `grundlagen-referenz-richtung.md` | 0 Zeilen | unberührt |
| [MR-014](../../../../harness/conventions/MR-014-keine-agenten-telemetrie.md) | `modul-15-observability.md` | 0 Zeilen | unberührt |
| [MR-015](../../../../harness/conventions/MR-015-welle-closure-ohne-replay.md) | `modul-06-roadmap.md` | 0 Zeilen | unberührt |
| [MR-016](../../../../harness/conventions/MR-016-validator-unbesetzt.md) | `modul-08-agentenrollen.md` | 0 Zeilen | unberührt |
| [MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md) | — *(Repo-Aussage-Korrektur, kein Baseline-Regel-Ersatz)* | `adr.template.md`: 0 Zeilen | **Auflösungs-Trigger ausgelöst** — §4 |
| [MR-018](../../../../harness/conventions/MR-018-review-pflicht-v610-wortlaut.md) | — *(dito, Template statt Regelwerk)* | — bereits als Rückbau-Kandidat ausgewiesen | unverändert: bleibt bis Etappe A vendored ist (slice-161 §4.4) |

Fünf der sieben aktiven Adaptionen
([MR-011](../../../../harness/conventions/MR-011-verfeinerungs-form.md)/[MR-012](../../../../harness/conventions/MR-012-referenzmatrix-grandfathering.md)/[MR-014](../../../../harness/conventions/MR-014-keine-agenten-telemetrie.md)/[MR-015](../../../../harness/conventions/MR-015-welle-closure-ohne-replay.md)/[MR-016](../../../../harness/conventions/MR-016-validator-unbesetzt.md))
referenzieren
Regelwerk-Dateien, die zwischen `v6.0.0` und `v6.1.0` **byte-identisch**
sind (jede einzeln mit `git diff <tag1> <tag2> -- <pfad>` gemessen, nicht
nur aus dem 6-Datei-Diff-Stat übernommen) — keiner der drei tatsächlich
geänderten Regelwerk-Module (`modul-07`, `modul-10`, `modul-13`) wird von
einer aktiven Adaption referenziert. Kein Handlungsbedarf für diese fünf.

## 4. Der eine echte Brocken: MR-017s Auflösungs-Trigger ist eingetreten

[`MR-017`](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md)
nennt als Auflösungs-Trigger *"die Überarbeitung der ID-Schema-Deklaration
— oder die nächste Baseline-Migration, je nachdem was zuerst eintritt"*.
Mit `welle-14` ist Zweiteres eingetreten. Inhaltlich hat sich nichts
verschoben (`adr.template.md` bleibt byte-identisch, s. Tabelle) — der
Trigger feuert rein durch das Ereignis "Migration", nicht durch einen
Wortlaut-Bruch.

**Der naheliegende Move — eine neue `MR`-Kennung anlegen, die denselben
Text mit `v6.1.0` statt `v6.0.0` wiederholt — wäre der dritte Durchlauf
desselben Musters:**
[MR-007](../../../../harness/conventions/done/MR-007-adr-vorlagen-version.md)→[MR-013](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md)
(erster Durchlauf, vor der Verzeichnisform des Registers),
[MR-013](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md)→[MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md)
(zweiter, belegt in
[`BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md),
`evidence/slice-141.md`).
[MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md)
selbst benennt diesen genauen Fall: *"ein Muster, das für sich genommen
ein Beobachtungs-Register-Eintrag wäre, träte es ein drittes Mal auf"*.
Einen Eintrag zu verkörpern, der die eigene Wiederholung schon
vorhergesagt hat, wäre die Harness-Lüge, gegen die `AGENTS.md` §3.7 und
`modul-13` stehen — ein Sensor, den man kennt, aber nicht anwendet.

**Vorschlag statt einer weiteren Versions-`MR`:** die **saubere**
Auflösung fahren, die
[MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md)
als Alternative nennt —
[MR-000](../../../../harness/conventions.md#mr-000)s
ID-Schema-Deklaration so umformulieren, dass die ADR-Vorlagen-Zeile nicht
mehr eine Versionsnummer trägt, die bei jeder Baseline-Migration erneut
veraltet, sondern generisch auf *"die jeweils aktuell vendorte Fassung,
siehe [`harness/conventions.md`
§Baseline](../../../../harness/conventions.md#baseline)"* verweist — der
Baseline-Stand ist dort bereits die einzige Quelle, die bei jeder
Migration ohnehin aktualisiert wird (Liefer-Punkt 2 von Etappe A,
slice-161 §6).
[MR-000](../../../../harness/conventions.md#mr-000)
selbst wird dabei **nicht** verändert (§Adaptions-Block: "an einem
akzeptierten Eintrag wird nichts nachträglich inhaltlich geändert") —
die Umformulierung braucht einen neuen `MR`-Eintrag (Kennung wird bei
dessen eigener Anlage vergeben, nicht hier), der
[MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md)
mit permanentem Auflösungs-Trigger ablöst, analog zur Form, in der
[MR-014](../../../../harness/conventions/MR-014-keine-agenten-telemetrie.md)/[MR-015](../../../../harness/conventions/MR-015-welle-closure-ohne-replay.md)/[MR-016](../../../../harness/conventions/MR-016-validator-unbesetzt.md)
ihre Vorgänger abgelöst haben. Das ist ein eigener Folge-Slice (§6) — die
Formulierung einer neuen Adaption ist eine inhaltliche Entscheidung
(Architect-Rolle, [Modul 8](../../../../.harness/baseline/v6.0.0/regelwerk/modul-08-agentenrollen.md#rollen-sequenz-für-eine-welle)),
keine reine Analyse.

## 5. Ergebnis der aufgelösten (`done/`) MR-Dateien

Keine der elf aufgelösten Dateien
([MR-001](../../../../harness/conventions/done/MR-001-spezifikations-schicht.md)
… [MR-010](../../../../harness/conventions/done/MR-010-rueckbau-drei-adaptionen.md),
[MR-013](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md))
referenziert `modul-07`, `modul-10` oder `modul-13` in ihrem
`Ersetzt-Baseline-Regel`-Feld (`grep` über alle elf Dateien, ein
Zufallstreffer auf "modul-13" in
[MR-009](../../../../harness/conventions/done/MR-009-validator-unbesetzt.md)s
Fließtext betrifft eine andere Aussage, keinen Baseline-Bezug). Aufgelöste Einträge werden nicht neu
bewertet — hier nur bestätigt, dass keine Überschneidung mit dem
`v6.1.0`-Diff besteht, die eine Neubewertung nahelegen würde.

## 6. Vorschlag: ein Folge-Slice, dann Etappe A

- **Folge-Slice** (noch keine ID): einen neuen `MR`-Eintrag formulieren
  (Kennung wird bei Anlage vergeben) — löst
  [MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md)
  permanent ab, ADR-Vorlagen-Referenz zeigt generisch auf
  `conventions.md` §Baseline statt auf eine feste Versionsnummer. Klein
  (ein Eintrag, ein Feld in
  [MR-000](../../../../harness/conventions.md#mr-000) unverändert,
  Adaptions-Block-Tabelle aktualisiert) — passt in die
  Drei-Liefer-Punkte-Grenze.
- **Danach Etappe A** (Vendoring, slice-161 §6) — unverändert durch diesen
  Durchgang: keiner der drei Liefer-Punkte dort hängt an
  [MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md)
  oder dem hier vorgeschlagenen Folge-Eintrag.

## 7. Risiken und offene Punkte

- *Der vorgeschlagene Folge-Slice wird nicht gezogen,
  [MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md)
  bleibt aktiv liegen* — **Ausgang:** weiter offen → Beobachtungs-Register
  ([`BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md)
  bleibt `offen`); ein drittes `MR-<NNN>`-Duplikat bei der übernächsten
  Migration wäre dann der tatsächliche dritte Durchlauf.
- *Ob die vorgeschlagene Formulierung des neuen `MR`-Eintrags so trägt,
  ist Maintainer-Entscheidung* — **Ausgang:** gestrichen mit Begründung: kein
  Risiko im Sinn der Dreier-Menge, sondern die Kernfrage, die dieser
  Analyse laut Kopf-Vermerk zur Abnahme vorgelegt wird.

## 8. DoD

- [x] Alle 18 MR-Dateien (7 aktiv, 11 aufgelöst) einzeln gegen den vollen
      `v6.0.0`↔`v6.1.0`-Tag-Vergleich geprüft — jede Referenz-Datei
      einzeln mit `git diff <tag1> <tag2> -- <pfad>` gemessen (§2/§3),
      nicht aus dem 6-Datei-Diff-Stat aus slice-161 übernommen.
- [x] Der eine echte Fund
      ([MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md)-Trigger)
      benannt, gegen sein eigenes
      Vorhersage-Zitat gehalten, Folge-Slice vorgeschlagen statt eine
      dritte Musterwiederholung stillschweigend zu fahren (§4/§6).
- [x] Unabhängiges Plan-Review über getrennten Kontext durchgeführt; der
      dabei gefundene Zählfehler in §10 (7 statt 8 offene
      `BEO-HARNESS`-Einträge) vor Abnahme korrigiert.
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Beobachtungs-Register gesichtet (§10); kein neuer Eintrag — der
      relevante Fund ist bereits als `rueckbau-kandidat-ueberlebt-baseline-migration`
      registriert.
- [x] Jedes Risiko aus §7 trägt einen Ausgang.

## 9. Closure-Notiz

- **Was hat funktioniert:** das dateiweise `git diff <tag> <tag> --
  <pfad>`-Verfahren statt der Diff-Stat-Liste aus slice-161 zu fahren, hat
  den blinden Fleck strukturell geschlossen, den slice-162 nur zufällig
  fing — jede der sieben aktiven Referenzen wurde einzeln gemessen, nicht
  aus der Sechs-Datei-Liste abgeleitet. Das unabhängige Review (getrennter
  Kontext) fand zusätzlich einen Zählfehler in der eigenen ersten
  Fassung von §10 (7 statt 8 offene `BEO-HARNESS`-Einträge, Ursache: eine
  vorbestehende, bereits in
  [F-8](../../../reviews/2026-09-05-slice-161-delta-analyse.md) benannte
  Sub-Area-Feld-Inkonsistenz) — vor Abnahme korrigiert.
- **Was ging anders als geplant:** erwartet war ein Fund *innerhalb* des
  MR-Wortlauts (ein Baseline-Abschnitt, der sich unter einer aktiven
  Adaption verschoben hat); tatsächlich lag der einzige Fund in einem
  **Trigger**, der nicht am Wortlaut, sondern am Migrations-*Ereignis*
  selbst hängt
  ([MR-017](../../../../harness/conventions/done/MR-017-adr-vorlagen-version.md))
  — eine Klasse von Fund, die ein reiner
  Wortlaut-Diff nie zeigen kann, weil es keinen Wortlaut-Unterschied gibt.
- **Lerneintrag — Form: geschärfte Regel.** *Ein Adaptions-Durchgang nach
  einer Baseline-Migration prüft nicht nur, ob referenzierte
  Baseline-Abschnitte sich inhaltlich verschoben haben, sondern auch, ob
  ein Auflösungs-Trigger der Form "die nächste Baseline-Migration" durch
  das Migrations-Ereignis selbst ausgelöst wurde — unabhängig davon, ob
  der referenzierte Wortlaut sich geändert hat. Ein `git diff`-Verfahren
  wie in slice-161 §2 findet die erste Klasse, nicht die zweite; nur ein
  gezielter Blick auf jedes `Auflösungs-Trigger`-Feld findet beide.*
- **Beobachtungs-Register (`../observations/`):** kein neuer Eintrag —
  [`BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md)
  trägt den Fund bereits (1×); dieser Slice bestätigt ihn erneut, ohne den
  dritten Durchlauf tatsächlich zu erzeugen — kein neues Auftreten im Sinn
  der Zählregel (§Ein Vorgang zählt einmal).
- **Folge-Slices:** noch keine ID vergeben — vorgeschlagen: einen neuen
  `MR`-Eintrag formulieren (§6), danach Etappe A (Vendoring, slice-161
  §6).
- **Risiken aus §7:** beide mit Ausgang — siehe §7.
- **Drei Paarungen:** verschoben auf die Closure von `welle-14` (dieser
  Slice trägt ein `**Welle:**`-Feld, [Modul 8](../../../../.harness/baseline/v6.0.0/regelwerk/modul-08-agentenrollen.md#rollen-sequenz-für-eine-welle)).

## 10. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Harness-Einstieg** (`AGENTS.md`, `harness/conventions.md`,
`harness/conventions/*.md`), Greenfield, Schwelle ≥ 2/3 erfüllt
(`harness/conventions.md` §Modus-Deklaration pro Sub-Area).

**Vorgelagert — offene Beobachtungen sichten:** Register für
**Harness-Einstieg** durchgegangen — direkt über die Verzeichnisliste unter
`BEO-HARNESS/` geprüft (`state.md` je Verzeichnis gelesen), nicht per
`grep`-Musterabgleich auf den Sub-Area-Namen: **8** `offen` (2 weitere
`verkörpert`, korrekt ausgeschlossen). **Korrektur ggü. slice-161 §10:**
jene Sichtung fand nur 7 — das achte Verzeichnis,
[`sensor-ohne-dod-phrase-wirkungslos`](../observations/BEO-HARNESS/sensor-ohne-dod-phrase-wirkungslos/observation.md)
(von slice-160 angelegt, vor slice-161), trägt im eigenen Feld
`**Sub-Area:**` den Wert *„Harness-Tooling"* — eine Bezeichnung, die
`harness/conventions.md` §Modus-Deklaration nirgends führt, obwohl das
Verzeichnis selbst unter dem `BEO-HARNESS`-Präfix liegt. Bereits als
[F-8](../../../reviews/2026-09-05-slice-161-delta-analyse.md) im
Plan-Review von slice-161 benannt (INFO, „Hinweis für einen künftigen
Register-Aufräum-Slice", dort ohne Handlungsbedarf) — slice-161 §10 selbst
zog daraus aber nicht die Konsequenz für die eigene Sichtungs-Zählung: ein
Namensabgleich auf den Sub-Area-Freitext übersieht dieses Verzeichnis, ein
Abgleich auf den Verzeichnis-Pfad nicht. Das unabhängige Review dieses
Slice (s. §9) fand den Zählfehler unabhängig erneut. Die acht:
[`adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md),
[`baseline-normtext-nachgeschrieben`](../observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/observation.md),
[`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md),
[`hard-rule-37-ohne-sensor`](../observations/BEO-HARNESS/hard-rule-37-ohne-sensor/observation.md),
[`rueckbau-eintrag-ablage-auslegung`](../observations/BEO-HARNESS/rueckbau-eintrag-ablage-auslegung/observation.md),
[`rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md)
(dieser Slice bestätigt den Fund erneut, s. §9),
[`selbst-archivierung-verdoppelt-abschluss-aufwand`](../observations/BEO-HARNESS/selbst-archivierung-verdoppelt-abschluss-aufwand/observation.md),
[`sensor-ohne-dod-phrase-wirkungslos`](../observations/BEO-HARNESS/sensor-ohne-dod-phrase-wirkungslos/observation.md)
— alle bei 1×, keiner erreicht mit diesem Slice 3×. Die
Sub-Area-Feld-Diskrepanz selbst (`Harness-Tooling` statt eines
deklarierten Kürzel-Namens) ist kein Gegenstand dieses Slice — sie
berührt nicht `v6.1.0`, sondern eine bereits vorher bestehende
Repo-interne Unschärfe.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
