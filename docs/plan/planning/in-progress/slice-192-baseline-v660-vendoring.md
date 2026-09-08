# slice-192 — Baseline auf `v6.6.0` heben: Vendoring, Pins, MR-Zeiger

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** Maintainer-Meldung „es gibt eine neue Regelwerks-Version `v6.6.0`".
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
[`AC-QA-03`](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit).

**Berührte Spec-Stellen:** — · Der Slice berührt kein Spec-Stratum.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der adoptierte Stand ist `v6.6.0`. Das Bundle liegt vendored im
Baseline-Verzeichnis unter seinem Tag, `v6.5.0` ist entfernt, alle **46**
Pin-Vorkommen in
den **14** lebenden Dateien zeigen auf den neuen Stand, und für jeden aktiven
`MR`-Eintrag ist **gemessen**, ob sein `Ersetzt-Baseline-Regel`-Zeiger mitwandern
darf.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Adoption der einen inhaltlichen Neuerung** — *„Der Gate-Index steht
  einmal"* (§2). *Ein Folge-Slice übernimmt es* mit Kennung
  ([slice-193](../open/slice-193-gate-index-steht-einmal.md)): Sie streicht eine
  41-zeilige Tabelle aus [`AGENTS.md`](../../../../AGENTS.md) §4 und zieht drei
  Konfigurationen um. Das ist ein Umbau, kein Pin-Bump; zusammen wäre der Diff
  in einer Review-Sitzung nicht prüfbar.
- **Der Voll-Abgleich der 15 Ziel-Form-Paare.** *Ein Folge-Slice übernimmt es:*
  [slice-188](../open/slice-188-voll-abgleich-gate-und-skill.md) und
  [slice-189](../open/slice-189-voll-abgleich-spec-straten.md) tragen ihn
  ohnehin und laufen nach diesem Slice gegen `v6.6.0` statt zweimal.
  [`harness/conventions.md`](../../../../harness/conventions.md) §Baseline
  verlangt beim Sprung **Delta *und* Voll-Abgleich**; das Delta steht in §2, der
  Voll-Abgleich hat seine Adresse.
- **Zwei vendored Stände nebeneinander stehen lassen.** *Bestand bleibt nicht
  stehen:* §Baseline sagt, genau **ein** Stand liegt vendored; mehrere sind nur
  während einer Migration zulässig. Der alte geht mit diesem Slice.

## 2. Delta-Analyse `v6.5.0` → `v6.6.0` (2026-09-08)

**Gemessen** gegen den lokalen Kurs-Checkout, `git diff v6.5.0 v6.6.0 --
lab/regelwerk lab/templates`: **10 Dateien, +87/−32**. Kurs-Welle 129,
2026-09-08. Thema: *„Der Gate-Index steht einmal"*.

| Datei | Änderung | Trifft a-check |
|---|---|---|
| `templates/AGENTS.template.md` | §4 **verliert die Tabelle**; nur noch Regel + Zeiger auf `harness/README.md` §Sensors | **ja, groß** — §4 trägt 41 Zeilen, 7094 Zeichen |
| `templates/harness/README.template.md` | *„DIES IST DER EINZIGE GATE-INDEX"*; Target-Zelle = **nackter Name** (Argument in die Nachbarspalte, Link erlaubt) | **ja** für den Index-Anspruch; **nein** für die Zellen-Form — gemessen: null Zellen mit Argument in der Code-Span, 16 verlinkte je Tabelle |
| `templates/.d-check.yml` | `targets`-Beispiel mit `doc-tables: [harness/README.md]` und `authority: harness/README.md`; dazu *„Aktivieren heißt zwei Schritte"* | **ja** — a-check führt `doc-tables: [AGENTS.md, harness/README.md]`, `authority: AGENTS.md` |
| `templates/Makefile` | Kopfkommentar nennt nur noch `harness/README.md` §Sensors | ja, klein |
| `regelwerk/grundlagen-harness-dateien.md` | +17: der Gate-Index steht einmal, **mit Begründung** (beide Dateien liegen in jedem Lauf-Kontext, ein zweiter Index wird pro Lauf zweimal bezahlt); Target-Zelle nackt | ja — die Begründung ist die tragende |
| `regelwerk/modul-13-quality-gates.md` | +20: die Hard Rule hat eine **maschinelle Hälfte** (Deklarations-Sensor, beide Richtungen); Ausnahmeliste **namentlich, nie Glob** | teils erfüllt — `make doc-targets` läuft, `exempt-targets` ist namentlich |
| `regelwerk/modul-09-implementierung.md` | Ziel-Form-Satz: „Gate-**Regel** und Zeiger" statt „Gate-Tabelle" | Folge von oben |
| `regelwerk/modul-02-harness-bootstrap.md` | Bootstrap-Tabelle: `AGENTS.md` §4 fällt als Phasen-Träger weg | nein — Bootstrap ist durch |
| `regelwerk/modul-15-observability.md` | Drift-Regeln reden vom Gate-Index statt von `AGENTS.md` | nein — [`MR-014`](../../../../harness/conventions.md#mr-014) nimmt das Modul aus |
| `regelwerk/README.md` | Stand-Zeile | ja, mechanisch |

**Ein einziger inhaltlicher Punkt trifft a-check**, und er trifft hart: der
Gate-Index. Alles andere ist Pin-Arbeit. **Die Ironie gehört benannt:**
[slice-187](../done/wellenlos/slice-187-voll-abgleich-erstdurchgang-rest.md)
hat gerade den Satz *„Diese Tabelle listet auf; die Bindung steht in
`harness/README.md` §Sensors"* nach §4 übernommen — aus der `v6.5.0`-Ziel-Form,
in der er seit `v5.12.0` stand. `v6.6.0` löscht ihn zusammen mit der Tabelle.
Die Übernahme war **richtig gegen den adoptierten Stand** und ist mit dem
nächsten überholt; das ist kein Fehler des Abgleichs, sondern der Normalfall
eines Repos, das einer Baseline folgt.

## 3. Umsetzung

*(entsteht mit der Arbeit)*

## 4. Definition of Done

- [ ] Das Bundle (`regelwerk/` + `templates/`) liegt vendored unter dem Tag
      `v6.6.0` mit `SHA256SUMS`; `v6.5.0` ist entfernt; `make regelwerk-check`
      grün.
- [ ] Alle **46** Pin-Vorkommen in den **14** lebenden Dateien und die vier
      Baseline-Symlinks unter `.claude/rules/` zeigen auf `v6.6.0`;
      `make symlink-check` und `make doc-check` grün.
- [ ] Für **jeden** aktiven `MR`-Eintrag mit Baseline-Zeiger ist **gemessen**,
      ob der referenzierte Abschnitt im neuen Stand wortgleich ist — nicht
      angenommen. Wo nicht, trägt die Stelle die Abweichung sichtbar.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein öffentlicher Vertrag ist berührt:
`harness/conventions.md` §Baseline deklariert den adoptierten Stand.

## 5. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt die Zeiger-Messung, dass mehr als zwei
  `MR`-Einträge eine Abweichung tragen, wird der Adaptions-Durchgang abgetrennt.
- `in-progress` → `open` (blockiert): Lässt sich das Bundle nicht netzlos
  erzeugen, ist die Beschaffung ein eigener Vorgang — dieses Repo baut nicht
  gegen das Netz.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Zwischen diesem Slice und [slice-193](../open/slice-193-gate-index-steht-einmal.md)
  widerspricht [`AGENTS.md`](../../../../AGENTS.md) §4 dem adoptierten Stand.**
  Der Gate-Index steht dann an zwei Orten, obwohl die Baseline einen verlangt.
  — **Ausgang:** <offen bis Closure>
- **Der Pin-Bump ist mechanisch und trifft 46 Stellen.** Ein übersehener Pin
  zeigt ins Leere, sobald `v6.5.0` weg ist; `make doc-check` fängt tote Links,
  aber ein Pin **im Fließtext** ohne Link bleibt stehen.
  — **Ausgang:** <offen bis Closure>
- **Die MR-Zeiger dürfen nur wandern, wenn der Zielabschnitt wortgleich ist.**
  Fünf aktive Einträge tragen einen; `modul-15` ist im Delta geändert, und
  [`MR-014`](../../../../harness/conventions.md#mr-014) zeigt genau dorthin.
  — **Ausgang:** <offen bis Closure>
- **Das Bundle entsteht aus einem fremden Arbeitsbaum.** Der lokale
  Kurs-Checkout trägt eine uncommittete Änderung; ein Build daraus wäre nicht
  der Tag. — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen — berührt sind `HARNESS` und die **Vendored Baseline**, die
ausdrücklich **keinen Modus** trägt.)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso — **zwei** Quellen:
Register nach **allen** berührten Kürzeln und der Review-Report des
Vorgänger-Slice.)*
