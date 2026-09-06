# slice-168 — Kalibrierungs-Selbsttest für phrasen-/muster-basierte Prüfer

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (Trigger: 3×-Schwelle im Beobachtungs-Register,
zugewiesen beim Lese-Schritt der `welle-14`-Closure — kein Mehr über die
eigene DoD hinaus).

**Bezug:** [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
(3×: slice-120, slice-123, slice-165).

**Berührte Spec-Stellen:** — *(keine)* — Harness-/Werkzeug-Änderung ohne
Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim
Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

---

## 1. Ziel

Einen wiederkehrenden Fehler verhindern: ein phrasen- oder
muster-basierter Prüfer meldet grün, weil seine Prüfmenge leer ist (kein
Text trifft das gesuchte Muster) — nicht weil tatsächlich geprüft wurde.

## 2. Analyse (vor der Umsetzung)

Drei bisherige Fälle, unterschiedlicher Natur:

- `verify-ac-form` suchte `^**Happy Path:**`, einen Wortlaut, den das
  Repo null Mal führte (slice-120).
- `doc-complete` war advisory und lief nie, obwohl es einen Gegenstand
  gehabt hätte (slice-123) — **strukturell bereits behoben**: seit
  slice-123 im `verify`-Aggregat, kein „ohne Aufruf"-Fall mehr möglich.
- `make doc-reviews`s Trigger-Phrase „unabhängiger Review" wurde von
  a-checks eigener, tatsächlich verwendeter DoD-Formulierung
  („Unabhängiges Plan-Review …") nie getroffen — vier Slices liefen vier
  Mal durch, ohne dass der Prüfer je etwas sah (slice-165).

**Beantwortung der drei Analyse-Fragen:**

1. **Betrifft die Lücke primär `.d-check.yml`-Konfigurationen oder auch
   a-checks eigene `tools/*.sh`-Prüfer?** Nur noch Erstere. `verify-ac-form`
   existiert nicht mehr (durch `d-check`-Modul `structure` abgelöst,
   slice-080); a-checks verbleibende Bash-Prüfer
   (`verify-risiko-ausgaenge.sh`, `verify-observations.sh`,
   `commit-scope-check.sh`) tragen bereits eigene `self_test()`-Funktionen
   mit Gut-/Schlecht-Fixtures (Muster seit slice-102). Die **einzige**
   ungedeckte Klasse sind phrasen-/muster-basierte `d-check`-Modul-Configs.
2. **Ist a-checks `self_test()`-Muster auf `.d-check.yml`-Configs
   übertragbar?** Ja, mit einer Anpassung: a-checks eigene Prüfer testen
   ihre **eigene Bash-Logik** gegen Text-Fixtures im selben Prozess.
   `d-check` ist ein externes, digest-gepinntes Werkzeug — a-check kann
   dessen Code nicht selbst testen, aber es kann **denselben gepinnten
   Digest** gegen eigene Fixtures fahren und das **Ergebnis** prüfen. Das
   ist kein Unit-Test von `d-check`, sondern ein Kalibrierungs-Test der
   eigenen Konfiguration — dieselbe Sache, die in `slice-165` bereits
   einmal von Hand in einem Scratch-Repo gemacht wurde, jetzt als
   wiederholbares Skript.
3. **Repo-weite Lösung oder Workflow-Checkliste?** Eine **automatisierte**
   Lösung ist für die beiden real betroffenen Muster (`reviews`-Modul,
   `structure` `tasks-ignore-pattern`) machbar und wurde umgesetzt (§3) —
   für deren **Werkzeug-Seite**. Ungedeckt bleiben weitere heute schon
   konfigurierte phrasen-basierte Muster (`versions.pin-pattern`,
   `commits.exempt-pattern`, `vcs.immutable-when`,
   `matrix.exclude-sections`); das ist keine künftige, sondern eine
   **bestehende** Lücke und geht an
   [`slice-169`](../open/slice-169-korpus-seitige-kalibrierung.md).

## 3. Umsetzung

**Neues Werkzeug:** [`tools/dcheck-phrase-selftest.sh`](../../../../tools/dcheck-phrase-selftest.sh) —
baut für jedes der zwei betroffenen Muster ein isoliertes Fixture-Repo
(minimale `.d-check.yml`, ein Fixture-Dokument) und fährt den gepinnten
`d-check`-Digest zweimal: **Positiv** (die heute empfohlene Formulierung
„unabhängiger Review" muss auslösen) und **Negativ** (eine andere,
harmlose Formulierung darf **nicht** auslösen) — dieselbe
Zwei-Richtungen-Disziplin wie in `tools/verify-risiko-ausgaenge.sh`s
`self_test()`. Vier Kontrollen insgesamt (2 Muster × 2 Richtungen).

**Getestet, nicht nur behauptet:** die Negativ-Kontrolle wurde durch eine
absichtliche Fehl-Formulierung im Skript selbst verifiziert — sie schlägt
korrekt fehl, wenn die Trigger-Phrase auf die falsche Wortform
(„unabhängig**es**" statt „unabhängig**er**") gesetzt wird (§7).

**Makefile:** neues Target `dcheck-phrase-selftest`, im `gates`-Aggregat
zwischen `suppression-check` und `guard-selftest`. Dokumentiert in
`AGENTS.md` §4 und `harness/README.md` §Sensors.

**Nicht umgesetzt — die Korpus-Seite.** Der Selbsttest prüft, ob
*`d-check`* auf die gewählte Phrase reagiert. Ausgefallen ist zweimal die
Gegenrichtung: a-checks **eigener Korpus** traf das Muster nicht mehr
(slice-120, slice-165). Das Skript liest den echten Korpus nie an, es
hardcodet die richtige Fixture-Zeile. Gemessen: die Kandidatenmenge des
`reviews`-Moduls ist heute **nicht leer** (slice-164, -165, -166, -167) —
der Ausfall ist also nicht live, aber jederzeit wieder erreichbar; drei
geschlossene Slices mit vorhandenem Report (slice-161, -162, -163) tragen
bereits eine nicht-auslösende Wortform. Ebenfalls offen: das
`tasks-ignore-pattern` steht hier als **Kopie** neben dem Original in
`.d-check.yml`, ohne Kopplung. Beides geht an
[`slice-169`](../open/slice-169-korpus-seitige-kalibrierung.md).

## 4. Definition of Done

- [x] Analyse-Fragen aus §2 beantwortet, Umsetzungsweg entschieden
      (automatisiertes Gate, nicht nur Checkliste).
- [x] `tools/dcheck-phrase-selftest.sh` umgesetzt, ins Makefile und in den
      `gates`-Aggregat eingehängt, in `AGENTS.md`/`harness/README.md`
      dokumentiert.
- [x] `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf` auf `verkörpert`
      gesetzt, mit auflösbarem Zielort und Herkunfts-Anker.
- [x] Unabhängiger Review durchgeführt (Report unter `docs/reviews/`).
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe (dieses Gespräch,
2026-09-06), WIP-Limit frei.

**Rückführungen:** keine vorgesehen — Umfang ist auf einen Liefer-Punkt
geschnitten und passt.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz
geschrieben.

## 7. Risiken und offene Punkte

- *Eine repo-weite, automatisierte Lösung ist für extern konfigurierte
  `d-check`-Module nicht möglich (das Werkzeug liegt außerhalb von
  a-checks Kontrolle)* — **Ausgang:** gestrichen mit Begründung: für die
  beiden real betroffenen Muster ist eine Fixture-basierte Kalibrierung sehr wohl
  möglich (§3) — der gepinnte Digest ist deterministisch, Fixtures lassen
  sich reproduzierbar dagegen fahren; die anfängliche Sorge war unbegründet.
- *Das neue Skript testet nur die zwei heute bekannten Muster — ein
  drittes, künftiges phrasen-basiertes Modul bleibt ungedeckt, bis jemand
  eine Fixture dafür schreibt* — **Ausgang:** eingetreten → Folge-Slice
  [`slice-169`](../open/slice-169-korpus-seitige-kalibrierung.md). Der
  unabhängige Review hat das Risiko als bereits realisiert nachgewiesen:
  `versions.pin-pattern`, `commits.exempt-pattern`, `vcs.immutable-when`
  und `matrix.exclude-sections` sind heute konfiguriert und ungedeckt —
  kein künftiger Fall, sondern ein bestehender.

## 8. Closure-Notiz

- **Was hat funktioniert:** das Skript wurde gegen sich selbst getestet
  (eine absichtlich falsche Trigger-Phrase eingesetzt, geprüft dass der
  Selbsttest das fängt, dann zurückgesetzt) — dieselbe Disziplin, die
  `tools/verify-risiko-ausgaenge.sh`s eigener `self_test()` für sich
  selbst hat, hier einmalig von Hand nachvollzogen statt als Teil des
  Skripts (das Skript testet `d-check`-Konfiguration, nicht sich selbst).
- **Was ging anders als geplant:** die ursprüngliche Sorge aus §7 (eine
  automatisierte Lösung sei für externe Module evtl. nicht möglich)
  bestätigte sich nicht — der Kniff war, nur den **einen betroffenen**
  `.d-check.yml`-Modul-Block als Fixture-Konfiguration zu verwenden statt
  a-checks volle, auf den eigenen Baum zugeschnittene Konfiguration zu
  kopieren (die in einem leeren Fixture-Repo mit Konfigurationsfehlern
  anderer Module abgebrochen wäre).
- **Was der unabhängige Review korrigiert hat:** die Zusage war zu breit
  formuliert. Gedeckt ist die **Werkzeug-Seite** der beiden Muster (feuert
  `d-check` noch auf die gewählte Phrase?), nicht die **Korpus-Seite**
  (benutzt a-check die Phrase noch?) — und genau die ist zweimal
  ausgefallen. Der Review hat außerdem die Zusage stärker belegt als der
  Slice selbst: **vier** Mutationen, eine je Kontrolle, alle vier korrekt
  rot (der Slice hatte eine geprüft). Report:
  [`2026-09-06-slice-168-…`](../../../reviews/2026-09-06-slice-168-pruefer-kalibrierungs-selbsttest.md).
- **Lerneintrag — Form: neuer Sensor.** *`make dcheck-phrase-selftest`
  (`tools/dcheck-phrase-selftest.sh`) kalibriert die **Werkzeug-Seite**
  phrasen-basierter `d-check`-Modul-Konfigurationen gegen den gepinnten
  Digest — vier Kontrollen (`reviews`-Trigger-Phrase, `structure`
  `tasks-ignore-pattern`, je Positiv/Negativ). Liegt in
  `Makefile:dcheck-phrase-selftest`, im `gates`-Aggregat. Auslöser:
  `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf` (slice-120, slice-123,
  slice-165 — 3×). Die Korpus-Seite derselben Beobachtung trägt der Sensor
  **nicht**; sie geht an
  [`slice-169`](../open/slice-169-korpus-seitige-kalibrierung.md).*
- **Beobachtungs-Register (`../observations/`):**
  `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf` auf `geplant` gesetzt mit
  der Kennung [`slice-169`](../open/slice-169-korpus-seitige-kalibrierung.md)
  — die Regel ist beschlossen, aber erst zur Hälfte geschrieben; was
  bereits steht, ist im Eintrag mit Zielort benannt.
- **Folge-Slices:** [`slice-169`](../open/slice-169-korpus-seitige-kalibrierung.md)
  (Korpus-seitige Kontrolle der Kandidatenmenge, Kopplung des
  `tasks-ignore-pattern` an `.d-check.yml`).
- **Risiken aus §7:** beide mit Ausgang — siehe §7.
- **Drei Paarungen:** entfällt — dieser Slice ist wellenlos (`**Welle:**
  ohne Welle`).

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Gate-/Werkzeug-Schicht** (`tools/dcheck-phrase-selftest.sh`, `Makefile`,
`.d-check.yml`-Kenntnis), Greenfield, Schwelle ≥ 2/3 erfüllt.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-GATE/` über die
Verzeichnisliste geprüft — 11 `offen` vor diesem Slice (unverändert seit
`slice-167`); `pruefer-ohne-gegenstand-oder-aufruf` davon wird mit diesem
Slice `verkörpert` (s. §8), die übrigen zehn unverändert, keiner erreicht
3×.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
