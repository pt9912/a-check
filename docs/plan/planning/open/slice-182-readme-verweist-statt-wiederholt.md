# slice-182 — `harness/README.md` verweist, statt Regelwerk-Begründung nachzuschreiben

**Welle:** ohne Welle — die Closure-Bedingung wäre die DoD dieses Slice, und
`in-progress/` trägt keinen Welle-Bezug (Baseline-Regelwerk
`modul-06-roadmap.md` §Wann Arbeit eine Welle braucht).

**Bezug:** [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze) (Harness-Integrität), keine aktive ADR.

**Berührte Spec-Stellen:** — (der Slice berührt nur Rang 9 der Source Precedence).

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-09-07.

**Lerneintrag — Form:** geschärfte Regel.

---

## 1. Ziel und Abgrenzung

**Ziel:** [`harness/README.md`](../../../../harness/README.md) trägt für jede
Regel, die im vendorten Regelwerk steht, einen **Verweis** statt einer
nachgeschriebenen Begründung — die Datei ist Schritt 1 des Minimal Agent
Workflow und wird von jedem Lauf ganz gelesen.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Tabellen selbst.** Sie wachsen mit der Zahl der Gates, nicht mit der
  Geschwätzigkeit; ihre Zellen wächtert
  [slice-181](../open/slice-181-tabellenzellen-gewaechtert.md) maschinell. Hier geht es
  um die Prosa **um** sie herum, die kein `table.column` erreicht.
- **[`AGENTS.md`](../../../../AGENTS.md).** Dieselbe Klasse, andere Datei — der
  Register-Eintrag nennt sie ausdrücklich (*„`AGENTS.md` §1 trägt dieselbe
  Klausel weiterhin im Rumpf"*). Ein Vorgang, der beide anfasst, berührt zwei
  Rang-Stufen der Source Precedence und ist in einer Review-Sitzung nicht mehr
  prüfbar. **Folge-Slice**, wenn dieser hier die Kriterien belegt hat.
- **Ein Sensor.** *„Ist dieser Satz im Regelwerk schon gesagt?"* ist ein Urteil,
  kein Match — dieselbe Grenze, die [`AGENTS.md`](../../../../AGENTS.md) §3.7 für
  sich selbst benennt. Was maschinell geht, ist die Zeichen-Obergrenze in
  Tabellenzellen, und die trägt slice-181.

## 2. Ausgangsmessung (2026-09-07)

**Sichtbare Prosa** je Abschnitt — HTML-Bedienhinweise der Vorlage und
Tabellenzeilen herausgerechnet, gemessen gegen
`v6.5.0` · `templates/harness/README.template.md`:

| Abschnitt | a-check | Ziel-Form | Faktor |
|---|---|---|---|
| §Sensors | 1930 | 346 | **5,6 ×** |
| §Rollen und ihre Übergabe-Artefakte | 1542 | **0** | Abschnitt existiert dort nicht |
| §Safety and scope boundaries | 1055 | 47 | 22 × |
| §Leseordnung | 843 | 464 | 1,8 × |
| §Traceability rules | 753 | 385 | 2,0 × |

Gesamt **21 374** gegen **10 436** Zeichen. Der Faktor allein ist **kein**
Befund — ein ausgefülltes Dokument ist länger als seine Vorlage. Befund ist,
**wo derselbe Text schon woanders steht**; drei Stellen sind belegt:

1. **Doppeltes Zitat in derselben Datei.** `### Nicht-Gates` (500 Zeichen) und
   der Absatz *„Nicht hier, obwohl sie in keinem Aggregat hängen"* (553)
   paraphrasieren beide `modul-13` §Vorhanden ≠ behauptet — der Satz *„ein
   reales Target nicht als Gate zu führen ist keine Harness-Lüge"* steht
   zweimal, 22 Zeilen auseinander.
2. **Die Ziel-Form verbietet genau diese Stelle.** Ihr Bedienhinweis: `kein
   Gate` gehört *„IN DER ZEILE SELBST … **nicht in Prosa daneben**"*, und
   *„WÄCHST DIE SEKTION: … Ob der Überhang schon unter der Tabelle steht oder
   in die Zelle gedrängt wurde, **ist dieselbe Sache**."* An dieselbe Stelle
   setzt sie **68 Zeichen**: eine fette Zeile, keine Überschrift, kein Absatz.
   [slice-177](../done/welle-15/slice-177-sensors-struktur-zwei-tabellen.md) hat
   den Überhang aus den Zellen geholt und ihn darüber wieder aufgebaut.
3. **Dreifacher Satz im Repo.** *„Verifikation grün, Validation rot"* steht in
   `modul-08`:142, in [`MR-016`](../../../../harness/conventions/MR-016-validator-unbesetzt.md):19
   und in `harness/README.md`:189.

Im Rollen-Abschnitt trägt die **Tabelle** (11 Zeilen) die a-check-eigene
Zuordnung *Übergabe → Artefakt* — sie steht nirgends sonst und bleibt. Die 25
Prosa-Zeilen darum sind `modul-08` §Die neun Übergaben, §Rollen-Regeln und
[`MR-016`](../../../../harness/conventions/MR-016-validator-unbesetzt.md); der Schluss-Absatz über die Review-Serie vom 2026-07-26 ist zusätzlich
Chronik in einer gelesenen Datei ([`AGENTS.md`](../../../../AGENTS.md) §3.7).

**Beide Dateien verbieten das selbst:** [`AGENTS.md`](../../../../AGENTS.md) §1
(*„sie dupliziert deren Inhalt nicht; sonst entsteht Drift"*),
[`harness/README.md`](../../../../harness/README.md) §Purpose (*„Diese Datei
dupliziert sie nicht"*).

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `harness/README.md` §Sensors, Nicht-Gates-Einleitung | refactor | auf die Ziel-Form-Zeile; das zweite Zitat entfällt |
| `harness/README.md` §Rollen und ihre Übergabe-Artefakte | refactor | Tabelle bleibt, Prosa auf zwei Zeiger, Chronik-Absatz raus |
| `harness/README.md` §Safety · §Leseordnung · §Traceability | update | je Satz geprüft: im Regelwerk ⇒ Verweis, a-check-eigen ⇒ bleibt |

**Kriterium je Satz** — nicht die Länge, sondern die Herkunft: Steht die
Aussage im vendorten Regelwerk, wird sie ein Verweis. Ist sie a-check-eigen
(eine Zuordnung, eine Schwelle, eine benannte Lücke), bleibt sie.

## 4. Definition of Done

- [ ] §Sensors trägt die Nicht-Gates-Kennzeichnung in der Form der Ziel-Form;
      das doppelte `modul-13`-Zitat ist auf **einen** Verweis reduziert.
- [ ] §Rollen trägt die Zuordnungs-Tabelle unverändert und die Prosa als
      Zeiger auf `modul-08` und [`MR-016`](../../../../harness/conventions/MR-016-validator-unbesetzt.md);
      der Chronik-Absatz ist fort.
- [ ] Die übrigen Abschnitte sind **je Satz** gegen das Regelwerk geprüft, und
      jede Streichung wie jedes Bleiben ist im Review-Report begründet.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Die drei Paarungen (Anker · Folge-Slice ·
Register) trägt die Closure dieses Slice selbst — wellenloser Betrieb. Ein
öffentlicher Vertrag ist nicht berührt: `harness/README.md` ist Rang 9.

**Nebenbefund zur Zähl-Kalibrierung:** `tasks-ignore-pattern` in
[`.d-check.yml`](../../../../.d-check.yml) kennt *„Die drei Paarungen"* und
*„Doku-Update"* nicht, obwohl `modul-05` §Ziel-Form: Slice beide unter den
**konstanten** Posten führt. Beide stehen hier deshalb als feste Zeile statt als
Häkchen — dieselbe Auflösung, die [`AGENTS.md`](../../../../AGENTS.md) §5 für
den Gate-Lauf vorschreibt. Ob das Muster nachzuziehen ist, entscheidet
[slice-169](../in-progress/slice-169-korpus-seitige-kalibrierung.md) oder ein
Folge-Slice, nicht dieser.

## 5. Trigger

**Start** (`next` → `in-progress`): [slice-169](../in-progress/slice-169-korpus-seitige-kalibrierung.md)
liegt in `done/` (WIP-Limit 1).

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt die Satz-für-Satz-Prüfung, dass mehr
  als die drei geplanten Abschnitte betroffen sind, wird nach Abschnitten
  zerlegt statt in einem Lauf durchgezogen.
- `in-progress` → `open` (blockiert): Stellt sich beim Kürzen heraus, dass eine
  Aussage **nur** hier steht und im Regelwerk fehlt, ist das eine benannte
  Spec-Lücke — sie wird eingetragen, und der Slice wartet auf ihre Auflösung.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag geschrieben — und die **gemessene** Gegenprobe: die sichtbare Prosa
je geänderten Abschnitts erneut gezählt, gegen §2 gehalten.

## 7. Risiken und offene Punkte

- **Kürzen entfernt eine Aussage, die nur hier steht** — der Kontext-Trennungs-Absatz
  über die Review-Serie 2026-07-26 trägt eine Selbst-Einschätzung, kein Regelzitat.
  — **Ausgang:** <offen bis Closure>
- **Der Verweis altert anders als der Text.** Ein Zeiger auf `modul-08` §X bricht,
  wenn die nächste Baseline den Abschnitt umbenennt; ein nachgeschriebener Satz
  nicht. — **Ausgang:** <offen bis Closure>
- **Kein Sensor hält das Ergebnis.** Nach dem Kürzen wächst die Prosa beim
  nächsten Slice wieder — genau so ist sie entstanden. — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt ist allein **Harness-Einstieg**
(`HARNESS`; Achsen 1,2,3 laut
[`harness/conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area)).
Die Gate-/Werkzeug-Schicht ist **nicht** berührt: kein Target, kein Skript, kein
Workflow ändert sich.

**Vorgelagert — offene Beobachtungen sichten:** zwei Treffer in `HARNESS`.

- [`baseline-normtext-nachgeschrieben`](../observations/BEO-HARNESS/baseline-normtext-nachgeschrieben/observation.md)
  — bisher **1×** (slice-103). Dieser Slice ist das zweite Vorkommen und hebt
  ihn auf **2×**: Symptom, noch keine Lücke.
- [`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
  — bisher **1×**. Der Absatz über die Review-Serie 2026-07-26 ist das zweite
  Vorkommen, ebenfalls **2×**.

Erreicht einer der beiden mit einem *weiteren* Slice 3×, ist er eine Lücke und
braucht einen eigenen Folge-Slice — für eine inferentielle Regel heißt das einen
Skill oder eine Hard Rule, keinen Sensor ([`AGENTS.md`](../../../../AGENTS.md) §3.7).

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
