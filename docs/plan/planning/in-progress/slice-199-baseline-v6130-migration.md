# slice-199 — Baseline auf `v6.13.0` heben: Migration über 24 Wellen

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Maintainer-Anfrage „auf das neueste Regelwerk umstellen" (2026-09-29).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
[`AC-QA-03`](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/` (direkter
Weg ab `open/`, [`AGENTS.md`](../../../../AGENTS.md) §5).

**Autor:** Claude. **Datum:** 2026-09-29.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der adoptierte Stand ist `v6.13.0` (Kurs-Welle 153 · 2026-09-28). Das
Bundle liegt vendored im Baseline-Verzeichnis unter seinem Tag, `v6.6.0` ist
entfernt, **jede** Nennung des alten Standes außerhalb der eingefrorenen
Artefakte ist behandelt, für **jeden** aktiven `MR`-Eintrag ist der Ausgang
gemessen (fünf Klassen, `modul-02` §Freshness-Audit), die Ziel-Formen sind
gegen ihr Repo-Gegenstück abgeglichen ([`harness/conventions.md`](../../../../harness/conventions.md)
§Baseline: Delta **und** Voll-Abgleich), und die Stichprobe gegen den Bestand
ist gefahren.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Adoption inhaltlicher Neuerungen, die einen Umbau lebender
  a-check-Artefakte verlangen** (neue Pflicht-Felder, umbenannte Sektionen,
  umgestellte Formen). *Ein Folge-Slice übernimmt es* — je Befund eine Kennung,
  vergeben beim Durchgang. Der MR-Durchgang und der Voll-Abgleich
  **klassifizieren** hier, sie bauen nicht um; zusammen wäre der Diff in einer
  Review-Sitzung nicht prüfbar.
- **[slice-189](../open/slice-189-voll-abgleich-spec-straten.md)** (Voll-Abgleich
  der Spec-Straten, Architektur-Klausel). *Bestand bleibt bewusst stehen*: eigener
  Slice, läuft nach diesem Sprung ohnehin gegen den neuen Stand.
- **[slice-190](../open/slice-190-id-schema-deklaration-ueberarbeiten.md) und
  [slice-191](../open/slice-191-chronik-phrasen-sensor.md)**. *Bestand bleibt
  bewusst stehen*: eigener Bestand; falls der Durchgang sie berührt, wird das
  zitiert, nicht abgearbeitet.
- **Der Produkt-Code (`internal/`, `cmd/`).** *Schicht-Abgrenzung*: der Slice
  arbeitet ausschließlich in der Harness-/Doku-Schicht.
- **Der d-check-Werkzeug-Pin.** *Es wäre ein anderer Vorgang*: die vendored
  Baseline ist ein Text-Bestand; das Werkzeug dahinter wird separat gepinnt und
  hängt an keiner Baseline-Migration.

## 2. Delta-Analyse `v6.6.0` → `v6.13.0` (gemessen 2026-09-29)

**Messung:** `git diff v6.6.0 v6.13.0 -- lab/regelwerk lab/templates` gegen den
lokalen Kurs-Checkout: **42 Dateien, +691/−325** — sämtliche 23 Regelwerk-Dateien
und 19 der 20 Vorlagen. Das sind **24 Wellen (130–153)**, gegen 10 Dateien bei
`v6.5.0` → `v6.6.0`. Die Wellen-Themen, aus dem Kurs-CHANGELOG übernommen
(wortgetreue Titel, 130–139 im Lauf zu lesen):

| Welle | Thema |
|---|---|
| 130 | Welle- und Slice-Kennungen sind Namen |
| 131 | Der Rest der Zählraum-Ablösung |
| 132 | Plan vor Code bindet an Akzeptanzkriterien |
| 133 | Linkerhalt bei Operation 3 (Anbinden) |
| 134 | Sensor-Vorlage referenziert, auf allen drei Ebenen |
| 135 | E2E-Gate-Typ + Bewusstes-Brechen-Pflicht für DoD-Testbehauptungen |
| 136 | Append-only-Klausel: welche Artefakt-Klassen sie trägt |
| 137 | Ein Slice, dessen Gegenstand ein anderer übernimmt: der Ausgang aus `open/` und `next/` |
| 138 | Die Referenzmatrix mechanisch vollständig: Welle, Carveout und Roadmap bekommen eine Klasse |
| 139 | Die ADR-Kennung in der Vorlage: Segment, Linkpflicht und kein doppeltes Token |
| 140 | Acht Stellen, an denen das Regelwerk dem Leser eine andere Antwort gibt als seine Quelle |
| 141 | Ein Carveout auf einer Schwelle nennt die ADR, die sie setzt |
| 142 | Fehlannahmen im Spiegel: jeder Punkt trägt seine Behauptung |
| 143 | Einordnung: das Regelwerk ist spec-anchored |
| 144 | Randbedingungen im Lastenheft: eine eigene Reihe `LH-RB-NN` |
| 145 | Hard Rule als vierte Trigger-Klasse |
| 146 | Auch das Workflow-Skelett braucht einen Trigger |
| 147 | Guide-Datei-Wildwuchs: eine Zeilen-Obergrenze, die nur mit sichtbarem Commit steigt |
| 148 | WIP-Limit zählt den Lauf, nicht den Menschen |
| 149 | Regel-Auslagerung: derselbe Schnitt wie `harness/conventions.md`, jetzt für `AGENTS.md` |
| 150 | Der Planungs-Bestand bekommt seinen eigenen Lese-Schritt |
| 151 | Gate-Erweiterung ist nicht automatisch ein ADR-Anlass |
| 152 | Vierter Fehlerklassen-Treffer erschöpft die Prosa-Verkörperung |
| 153 | ADR-Nachzug bei sich änderndem Kontext |

**Gelesen und damit schon eingeordnet (149–153, Volltext):**

- **148 trifft [`AGENTS.md`](../../../../AGENTS.md) §5 am Wortlaut:** a-check
  formuliert das WIP-Limit als Menge in `in-progress/`; die neue Fassung zählt
  **den Lauf** des Rolleninhabers, nicht den Menschen und nicht das Verzeichnis.
- **149 trifft die Ziel-Form des `AGENTS.md`:** §5 startet in der Vorlage als
  **Index-Tabelle** statt Bullet-Liste — ein Voll-Abgleich-Befund auf Vorlage.
- **150 trifft den Planungs-Bestand:** `open/` trägt hier fünf Slices; der neue
  Lese-Schritt verlangt deren Durchgang — **konsolidiert** zuerst.
- **151/153 sind Einordnungs-Regeln** (Gate-Aufnahme, ADR-Nachzug) — sie ändern
  kein Artefakt, sie schärfen Praktik, das Review gegen sie liest.

Die Einordnung der übrigen 19 Wellen leistet der Durchgang in §3 — **je
`MR`-Eintrag gegen die fünf Ausgänge**, nicht gegen diese Tabelle; die Tabelle
ist thematisch, der Durchgang ist normativ.

## 3. Umsetzung (Plan vor Code)

| Schritt | Gegenstand | Beleg/Gegenprobe |
|---|---|---|
| 1 | Bundle aus `git archive v6.13.0` in Wegwerf-Verzeichnis bauen (`tools/build-bundle.sh` des Kurses, der Kurs-Checkout ist sauber, der Arbeitsbaum wird trotzdem nie gelesen) | Dasselbe Verfahren auf `v6.6.0` reproduziert den bisher vendorten Baum **byte-gleich**; `SHA256SUMS` gegenprobt |
| 2 | Neuen Stand vendoren, **alten entfernen** (genau ein Stand), vier Baseline-Symlinks unter `.claude/rules/` umhängen | `make regelwerk-check`, `make symlink-check` |
| 3 | Nennungen des alten Standes in lebenden Dateien **messen** (nicht schätzen), dann bumpt — Ersetzung über das **Pfad-Muster**, nie über die nackte Kennung (Lektion slice-192 §3.6 F-1) | Zähltabelle Dateien/Nennungen, gemessen gegen den Stand vor dem Swap |
| 4 | Stand-Zeile in [`harness/conventions.md`](../../../../harness/conventions.md) §Baseline: Kurs-Welle **153 · 2026-09-28** | aus dem Kopf des vendorten `regelwerk/README.md` |
| 5 | **MR-Durchgang:** neun aktive Einträge ([`MR-012`](../../../../harness/conventions.md#mr-012) … [`MR-024`](../../../../harness/conventions.md#mr-024)) gegen die fünf Ausgänge (`modul-02`); `Ersetzt-Baseline-Regel`-Zeiger wandern nur bei **wortgleichem** Zielabschnitt, gemessen zweistufig (Datei im Delta? Abschnitt?) | Abschnitts-Messung je Eintrag im Slice dokumentiert |
| 6 | **Voll-Abgleich:** die Ziel-Formen mit genau einem Gegenstück, Abschnitt für Abschnitt, maschinelle Vorauswahl durch Wortfolgen-Vergleich; Instanz-Vorlagen und der Vorlagen-Index fallen heraus | Befundliste mit Abgleich-Ergebnis |
| 7 | **Stichprobe gegen den Bestand:** ein Abschnitt aus der Komplementärmenge (kein Delta seit der Adoption), rotierend | Frage pro Regel: steht sie im ausgefüllten Artefakt — oder als deklarierte Abweichung? |
| 8 | Gates und Verifikation: `make regelwerk-check`, `make gates`, `make verify` | Ausgaben in Datei, Exit-Codes getrennt |
| 9 | Unabhängiger Review, Closure, `make slice-mv`, Archivierung | Report unter `docs/reviews/`; Schritte 7–10 des Workflow |

### 3.1 Vendoring — belegt, nicht angenommen

Gebaut aus `git archive v6.13.0` in ein Wegwerf-Verzeichnis, dort
`tools/build-bundle.sh`; der Kurs-Checkout war sauber, der Arbeitsbaum wurde
nie gelesen. **Gegenprobe:** dasselbe Verfahren auf `v6.6.0` reproduziert den
bisher vendorten Baum byte-gleich (`diff -rq`: null Unterschiede) und das
nachgebaute Manifest stimmt mit dem committeten `SHA256SUMS` überein —
**54 Dateien** in beiden Ständen. Der unabhängige Review hat beide Bundles
neu gebaut und beide byte-gleich bestätigt.

### 3.2 Pin-Messung — Instrument und Zähler

**Instrument:** `git grep -c 'v6\.6\.0' <Pfadscope>` am Claim-Commit `2ce5da8`
(Pfadscope: alle getrackten Dateien außer `docs/plan/planning/done`,
`docs/reviews`, `harness/conventions/done`, `.harness/baseline` — die
eingefrorenen Klassen). **Zwei Zähler, verschieden gebaut** (Mess-Regel,
[`AGENTS.md`](../../../../AGENTS.md) §5):

| Zähler | Frage | Ergebnis |
|---|---|---|
| A — Zeilen je Datei (`grep -c`) | Wie viele Zeilen tragen den alten Stand? | **17 Dateien, 49 Zeilen** |
| B — Vorkommen (`grep -o`) | Wie oft steht das Token? | **63 Vorkommen** |

Die Differenz 49↔63: mehrfach genannte Tokens in Zitat-Blöcken und
Pfad+Kennung-Zeilen. **Erratum:** die Commit-Message des Bump-Commits meldet
„17 Dateien/42 Nennungen" — die 17 stimmen, die 42 war ein Summierfehler aus
dem Verlauf; die korrekte Messung steht hier, nachgetragen im Beobachtungs-
Register ([`BEO-HARNESS/messung-ohne-abgelegten-beleg`](../observations/BEO-HARNESS/messung-ohne-abgelegten-beleg/observation.md), 1×).
Ersetzt wurde ausschließlich über das **Pfad-Muster** (`baseline/v6.6.0` →
`baseline/v6.13.0`), nie über die nackte Kennung. **Rest am Lauf-Ende:** 9
Dateien/24 Zeilen nennen `v6.6.0` — jede beabsichtigt: historische Fakten
(CHANGELOG, [`AGENTS.md`](../../../../AGENTS.md) §5 Zitier-Form-Anker,
doc-mentions), Kennung-Zeiger der frozen Einträge, Fakten im Delta dieses
Plans, Provenance-Kommentare in [`.d-check.yml`](../../../../.d-check.yml) und
die beabsichtigten Nennungen der Nachfolge-Einträge.

### 3.3 MR-Durchgang — je Eintrag ein Ausgang

Gemessen zweistufig (Datei im Delta? dann Abschnitt), Quelle-Zeile
ausgenommen; die Zielabschnitts-Diffs wurden gegen die beiden Bundle-Stände
geführt:

| MR | Zielabschnitt im Delta? | Abschnitt | Ausgang |
|---|---|---|---|
| [`MR-012`](../../../../harness/conventions.md#mr-012) | ja (alle 23 geändert) | **Abweichung** (23 242 → 23 083 Z., 4 Hunks) | *bleibt gültig* → **[MR-025](../../../../harness/conventions.md#mr-025)** |
| [`MR-014`](../../../../harness/conventions.md#mr-014) | ja | **wortgleich** (566 Z.) | Zeiger wandert |
| [`MR-015`](../../../../harness/conventions.md#mr-015) | ja | **Abweichung** (10 552 → 13 522 Z.) | *bleibt gültig* → **[MR-026](../../../../harness/conventions.md#mr-026)** |
| [`MR-016`](../../../../harness/conventions.md#mr-016) | ja | **wortgleich** (1 417 Z.) | Zeiger wandert |
| [`MR-019`](../../../../harness/conventions.md#mr-019) | — (kein Zeiger) | — | *bleibt gültig*: Treiber ist die Report-Vorlage; deren Strukturänderung (Tabelle statt F-1-Blöcke) berührt die Opt-in-Entscheidung nicht |
| [`MR-020`](../../../../harness/conventions.md#mr-020) | — (kein Zeiger) | — | *bleibt gültig*: die Referenz auf den generischen Stand gilt unverändert weiter |
| [`MR-022`](../../../../harness/conventions.md#mr-022) | ja | **Abweichung** (10 942 → 6 574 Z.) | *bleibt gültig* → **[MR-027](../../../../harness/conventions.md#mr-027)** (Suffix-Passage Zeile 17 wortgleich) |
| [`MR-023`](../../../../harness/conventions.md#mr-023) | — (kein Zeiger) | — | *bleibt gültig*: die neue Fassung lässt die Kennungs-Form ausdrücklich als Repo-Deklaration |
| [`MR-024`](../../../../harness/conventions.md#mr-024) | — (kein Zeiger) | — | *bleibt gültig*: der selbst-auflösende Trigger (nächstes Release) ist nicht eingetreten |

Erratum: [MR-028](../../../../harness/conventions.md#mr-028)
korrigiert [MR-026](../../../../harness/conventions.md#mr-026)s Schritt-Zuordnung (Review-Fund M3 — Accepted-Einträge
korrigiert man nur über einen Nachfolger).

### 3.4 Voll-Abgleich und Stichprobe

**Voll-Abgleich** (Singleton-Paare, Heading-Struktur gegen die Ziel-Form):
15 Paare abgeglichen — [`AGENTS.md`](../../../../AGENTS.md),
[`harness/README.md`](../../../../harness/README.md),
[`harness/conventions.md`](../../../../harness/conventions.md),
[Lastenheft](../../../../spec/lastenheft.md), [Spezifikation](../../../../spec/spezifikation.md),
[Architektur](../../../../spec/architecture.md), [Projekt-README](../../../../README.md),
[ADR-Index](../../../../docs/plan/adr/README.md),
[Carveouts](../../../../docs/plan/carveouts/README.md),
[Planning-README](../../../../docs/plan/planning/README.md),
[Roadmap](../in-progress/roadmap.md), [Reviewer-Skill](../../../../.harness/skills/reviewer.md),
[Closure-Note-Reviewer-Skill](../../../../.harness/skills/closure-note-reviewer.md),
[`.d-check.yml`](../../../../.d-check.yml), Sensor-Vorlage. **Kein fehlendes
Pflicht-Feld außerhalb der Spec-Straten**; die Abweichungen sind additiv oder
Titel-Wortlaut. Die Spec-Straten tragen die zwei echten Befunde: die neue
`LH-RB`-Reihe (Welle 144) und die Suffix-Form im
Spezifikations-Template — **beide gehören in
[slice-189](../open/slice-189-voll-abgleich-spec-straten.md)**, den
ausdrücklich nicht ausgeführten Abgleich. Die Roadmap ist heading-gleich,
die `.d-check.yml`-Modul-Menge Obermenge des Template-Starters.

**Stichprobe gegen den Bestand** (Komplementärmenge: Abschnitte ohne Delta
zwischen `v6.6.0` und `v6.13.0`): gezogen wurde
`grundlagen-durchsetzungsschicht.md` §„Die Lücke: aspirativ vs. bindend" —
die Frage pro Regel: steht sie im ausgefüllten Artefakt? **Beide Regelfälle
sind verkörpert:** „make/Docker-only" durch den PreToolUse-Command-Guard
([`AGENTS.md`](../../../../AGENTS.md) §3.1), „Gates vor dem Handoff" durch
den Stop-Hook mit [`make record-gates`](../../../../Makefile). **Kein Fund.**

## 4. Definition of Done

- [x] Das Bundle (`regelwerk/` + `templates/`) liegt vendored unter dem Tag
      `v6.13.0` mit `SHA256SUMS`; `v6.6.0` ist entfernt; die Gegenprobe (alter
      Tag byte-gleich reproduziert, Manifest eingeschlossen) ist belegt
      (§3.1); `make regelwerk-check` und `make symlink-check` grün.
- [x] Die Pin-Messung trägt Instrument und Zähler (§3.2), **jede** Nennung des
      alten Standes außerhalb der eingefrorenen Artefakte ist behandelt; die
      vier Baseline-Symlinks zeigen auf den neuen Stand; **der MR-Durchgang**
      ist je Eintrag mit Ausgang belegt (§3.3); `make doc-check` grün.
- [x] **Voll-Abgleich** und **Stichprobe** sind gefahren; die Befunde sind
      benannt (§3.4), Adoption-Aspekte tragen Folge-Slice-Kennungen
      ([slice-200](../open/slice-200-adoption-v6130-agents-und-matrix.md),
      [slice-189](../open/slice-189-voll-abgleich-spec-straten.md)).
- [x] Unabhängiger Review durchgeführt, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md)
      — 0 HIGH · 4 MEDIUM · 3 LOW · 1 INFO; die vier MEDIUM sind im Lauf
      behoben (§7).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag; Beobachtungs-Register
      fortgeschrieben.
- [x] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein öffentlicher Vertrag ist berührt:
[`harness/conventions.md`](../../../../harness/conventions.md) §Baseline
deklariert den adoptierten Stand.

## 5. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei; direkter Weg.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Findet der Voll-Abgleich mehr als drei
  Adoption-Aspekte an Singleton-Artefakten (neue Pflicht-Felder, umbenannte
  Sektionen), wird der Abgleich als eigener Slice abgetrennt — der Pin-Bump
  bleibt hier.
- `in-progress` → `open` (blockiert): Lässt sich der Tag nicht beschaffen oder
  das Bundle nicht netzlos bauen, ist die Beschaffung ein eigener Vorgang —
  dieses Repo baut nicht gegen das Netz.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Massen-Ersetzung trifft historische Aussagen** — die Klasse aus
  [`BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage`](../observations/BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage/observation.md)
  (1×): Ersetzt man mit der nackten Kennung, treffen alle Nennungen, auch die
  historischen. — **Ausgang:** *entfallen*, gestrichen mit Begründung: ersetzt
  wurde ausschließlich über das Pfad-Muster; die Abgleiche nach dem Bump
  zeigen die verbliebenen Nennungen nur noch an beabsichtigten Stellen (§3.2).
  Die Beobachtung bleibt im Register bei 1×.
- **Der MR-Durchgang ist der Wächter für Auflösungs-Trigger** —
  [`BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md)
  steht bei **2×**. Wird hier ein drittes Auftreten benannt, ist die Klasse eine
  Lücke und braucht einen eigenen Folge-Slice. — **Ausgang:** *entfallen*,
  gestrichen mit Begründung: der Durchgang hat die Auflösungs-Trigger aller
  neun Einträge abgefragt; **keiner war eingetreten** ([MR-024](../../../../harness/conventions.md#mr-024)s
  Release-Trigger inklusive) — die Beobachtung bleibt bei 2×.
- **Zwischenstand-Widerspruch:** Die Baseline gilt in der neuen Fassung, während
  die Adoption lebender Artefakte (WIP-Limit-Formulierung, `AGENTS.md` §5-Form,
  Lastenheft-Randbedingungen) Folge-Slices braucht — a-check folgt dann kurzfristig
  einer Regel, die seine Dokumente nicht tragen. — **Ausgang:** *eingetreten* →
  Folge-Slice [slice-200](../open/slice-200-adoption-v6130-agents-und-matrix.md)
  (WIP-Limit, §5-Form, Matrix-Klassen); die Spec-Straten trägt
  [slice-189](../open/slice-189-voll-abgleich-spec-straten.md).
- **Neue Pflicht-Felder an vertraglich gebundenen Artefakten** (Welle 144
  berührt die Lastenheft-Vorlage): eine Pflicht-Form-Änderung am Lastenheft ist
  ein Change Request, kein Doku-Bump. — **Ausgang:** *eingetreten* → Folge-Slice
  [slice-189](../open/slice-189-voll-abgleich-spec-straten.md), wo der
  Lastenheft-Abgleich mit der Trennung Doku-Pflege/CR liegt.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel.** *Beleg-first: Wer eine Mess-Zahl
berichtet, legt Instrument und Lauf-Beleg im selben Zug ab — weil die aus dem
Verlauf summierte Zahl die einzige Quelle wurde und sich als falsch erwies.*
Gemessen an diesem Slice: die Bump-Message meldete „17 Dateien/42 Nennungen";
am Claim-Commit messen die beiden anders gebauten Zähler 49 Zeilen bzw. 63
Vorkommen (die 17 stimmt). Dieselbe Ursache machte den einzigen
Inhaltsfehler eines Accepted-Eintrags teuer: die unbedlockte Schrittzahl in
[MR-026](../../../../harness/conventions.md#mr-026) erzwang das Erratum [MR-028](../../../../harness/conventions.md#mr-028) — die Immutabilität ist genau dafür da.
Beide Fälle hängen am selben Muster und sind gezählt:
[`BEO-HARNESS/messung-ohne-abgelegten-beleg`](../observations/BEO-HARNESS/messung-ohne-abgelegten-beleg/observation.md)
(1×).

**Was hat funktioniert:** das slice-192-Verfahren hält auch bei rund vierfachem
Delta (24 Wellen, 42 Dateien) — netzloser Bundle-Bau mit byte-gleicher
Gegenprobe, zweistufige Zeiger-Messung, Pfad-Muster-Ersatz. Der MR-Durchgang
mit der fünf-Ausgänge-Vokabel ließ sich auf alle neun Einträge anwenden, ohne
einen stillschweigenden Zeiger zu ziehen.

**Was ging anders als geplant:** 19 von 20 Vorlagen änderten sich — die
Plan-Schritte 5–7 (Durchgangs-Listen, Voll-Abgleich, Stichprobe) wuchsen über
die Schätzung; die Belege lagen zunächst nur im Lauf, nicht im Repo, und der
unabhängige Review meldete sie als fehlend (M1–M3). Behoben durch Nachtrag in
§3 und das Erratum [MR-028](../../../../harness/conventions.md#mr-028) — die Beobachtungs-Zeile oben ist der Nachlauf.

**Steering-Loop-Eintrag:** — *(nichts verkörpert; der Lerneintrag ist gezählt,
nicht verkörpert — die bestehenden Mess-Regeln in [`AGENTS.md`](../../../../AGENTS.md) §5
decken die Klasse bereits, was fehlte, war ihre Befolgung im Lauf.)*

**Beobachtungs-Register (`../observations/`):** `BEO-HARNESS/messung-ohne-abgelegten-beleg/`
neu angelegt, Beleg `evidence/slice-199.md` — Zähler 1×.

**Folge-Slices:** [slice-200](../open/slice-200-adoption-v6130-agents-und-matrix.md)
(Adoption `v6.13.0`: WIP-Limit, §5-Form, Matrix-Klassen) — ist eine Datei in
`open/`; die Spec-Straten trägt
[slice-189](../open/slice-189-voll-abgleich-spec-straten.md) (schon da).

**Risiken aus §6:** jedes mit genau einem Ausgang — siehe §6.

**Drei Paarungen:** Anker — kein Steering-Loop-Eintrag verkörpert, kein
`liegt in`-Feld · Folge-Slice — slice-200 und slice-189 existieren als Dateien
in `open/` · Register — die neue Beobachtung trägt ihr `evidence/`-Verzeichnis;
`make verify-observations` deckt die Zitate.

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt sind **HARNESS**
([`AGENTS.md`](../../../../AGENTS.md), [`harness/conventions.md`](../../../../harness/conventions.md),
`harness/` — Achsen 1, 2, 3 ✓) und **GATE** (`.claude/rules/`-Symlinks,
ggf. Makefile-Nennungen — Achsen 1, 2, 3 ✓). Die Sub-Area *Vendored Baseline*
trägt kein Kürzel und keinen Modus (externer, unveränderter Fremdtext);
die Migration ist ihr deklarierter Aktualisierungsweg.

**Vorgelagert — offene Beobachtungen sichten:** das Register wurde durchgegangen
(2026-09-29). Treffer in HARNESS:
[`massen-ersetzung-trifft-die-historische-aussage`](../observations/BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage/observation.md)
(1× — genau dieser Vorgang, Prävention eingeplant) ·
[`mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md)
(2× — der Durchgang ist der Wächter, dritter Fund = Lücke) ·
[`agents-md-hinkt-baseline-dod-item-hinterher`](../observations/BEO-HARNESS/agents-md-hinkt-baseline-dod-item-hinterher/observation.md)
(offen — [`AGENTS.md`](../../../../AGENTS.md) ist Pin-Ziel, Befund wird gezählt) ·
[`zwei-regeln-machen-einander-unmoeglich`](../observations/BEO-HARNESS/zwei-regeln-machen-einander-unmoeglich/observation.md)
(2× — slice-192s Lernsignal, beim Abgleich im Blick). GATE: keine Treffer für
diesen Vorgang.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
