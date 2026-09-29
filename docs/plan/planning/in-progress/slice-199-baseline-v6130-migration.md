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

## 4. Definition of Done

- [ ] Das Bundle (`regelwerk/` + `templates/`) liegt vendored unter dem Tag
      `v6.13.0` mit `SHA256SUMS`; `v6.6.0` ist entfernt; `make regelwerk-check`
      grün; die Gegenprobe (alter Tag byte-gleich reproduziert) ist im Slice
      belegt.
- [ ] **Jede** Nennung des alten Standes außerhalb der eingefrorenen Artefakte
      ist behandelt — Messung dokumentiert —, und die vier Baseline-Symlinks
      zeigen auf den neuen Stand; `make symlink-check` und `make doc-check`
      grün.
- [ ] **MR-Durchgang:** für **jeden** aktiven `MR`-Eintrag ist der Ausgang aus
      den fünf Klassen benannt und der Zeiger-Messung unterzogen — nicht
      angenommen.
- [ ] **Voll-Abgleich** und **Stichprobe** sind gefahren; die Befunde sind
      benannt, Adoption-Aspekte tragen Folge-Slice-Kennungen oder eine
      Begründung, warum sie hier nicht entstehen.
- [ ] Unabhängiger Review durchgeführt, Report unter [`docs/reviews/`](../../../../docs/reviews/README.md).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag; Beobachtungs-Register
      fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

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
  historischen. — **Ausgang:** *weiter offen* → Register; Prävention ist
  eingeplant (Schritt 3: Pfad-Muster statt Kennung), die Klasse bleibt lebendig
  für den Review.
- **Der MR-Durchgang ist der Wächter für Auflösungs-Trigger** —
  [`BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter`](../observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md)
  steht bei **2×**. Wird hier ein drittes Auftreten benannt, ist die Klasse eine
  Lücke und braucht einen eigenen Folge-Slice. — **Ausgang:** bei Closure.
- **Zwischenstand-Widerspruch:** Die Baseline gilt in der neuen Fassung, während
  die Adoption lebender Artefakte (WIP-Limit-Formulierung, `AGENTS.md` §5-Form,
  Lastenheft-Randbedingungen) Folge-Slices braucht — a-check folgt dann kurzfristig
  einer Regel, die seine Dokumente nicht tragen. — **Ausgang:** *eingetreten* →
  Folge-Slices aus Schritt 5/6, benannt und befristet wie bei slice-192; die
  Alternative — alles in einem Slice — war in Schritt `next` geblieben.
- **Neue Pflicht-Felder an vertraglich gebundenen Artefakten** (Welle 144
  berührt die Lastenheft-Vorlage): eine Pflicht-Form-Änderung am Lastenheft ist
  ein Change Request, kein Doku-Bump. — **Ausgang:** bei Closure.

## 8. Closure-Notiz

*(wird vor dem `git mv` nach `done/` gefüllt)*

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
