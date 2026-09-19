# slice-195 — Handbuch-Vokabel für die Adapter-Rolle

**Welle:** ohne Welle.

**Bezug:** Zweiter Fund der externen Meldung `hexslice-architecture` (2026-09-19), im Repo
zeilengenau nachgeprüft. Folgt dem Zuschnitt von
[`ADR-0036`](../../adr/0036-port-richtung-inbound-outbound.md).

**Berührte Spec-Stellen:** — (der Slice berührt kein Spec-Stratum; er zieht ein
Benutzer-Dokument auf eine `Accepted`-Entscheidung nach).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude, im Auftrag des Maintainers. **Datum:** 2026-09-19.

**Lerneintrag — Form:** benannte Spec-Lücke.

---

## 1. Ziel und Abgrenzung

**Ziel:** Das Benutzerhandbuch führt für die Adapter-Rolle **ein** Vokabular — das von
[`ADR-0036`](../../adr/0036-port-richtung-inbound-outbound.md) entschiedene (`driving`/`driven`) —,
und die Handbuch-Familie hat einen deklarierten Ort im Beobachtungs-Register.

**Ausgangslage (gemessen):** `docs/user/benutzerhandbuch.md:365` zeigt
`internal/adapters/{inbound,outbound}/…`; derselbe Text sagt ab `:726` („Ein Port *treibt* nichts
… ein Adapter ist nicht *eingehend*") und die Regel-Tabelle ab `:257` das Gegenteil. `inbound`
und `outbound` sind seit [`ADR-0036`](../../adr/0036-port-richtung-inbound-outbound.md) das
**Port**-Vokabular. Das Dokument widerspricht sich damit selbst für dieselbe Rolle.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Ein Nachzug der Lab-Config des Konsumenten** — *ein anderer Vorgang in einem anderen Repo:*
  dort steht ein Kommentar, der die Kopplung Glob-Präfix ↔ Richtungssegment festhält; er ist
  Gegenstand von [slice-194](../done/wellenlos/slice-194-portscope-richtungssegment.md), nicht dieses Slice.
- **Eine MR-Adaption für die Kürzel-Vergabe** — *es wäre ein anderes Werkzeug:* der
  Adaptions-Block trägt Abweichungen **gegenüber der Baseline**
  ([`harness/conventions.md`](../../../../harness/conventions.md) §Adaptions-Block). Eine neue
  Zeile in der **Modus-Deklaration** ist keine Abweichung, sondern die Deklaration selbst (so
  entstanden alle acht bestehenden Zeilen).
- **Ein Nachzug aller Benutzer-Dokumente auf das Vokabular** — *Bestand bleibt bewusst stehen:*
  gemessen wird zuerst der Bestand; was der Lauf **zusätzlich** findet, gehört hierher, was er
  nicht findet, erzeugt keine Aufgabe.
- **Die Heuristik-Grenze des Handbuchs** — *Schicht-Abgrenzung:* die Grenze der Extraktion
  ([`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze))
  ist Gegenstand anderer Slices; hier geht es um ein Vokabular.

## 2. Ausgangsmessung — zwei Zähler, ein Ergebnis

**Erhoben, nicht gelesen** (Mess-Regel 3: eine Menge wird zweimal verschieden gezählt):

| Zähler | Bau | Ergebnis |
|---|---|---|
| 1 | `grep` auf die alte Wendung `adapters/{inbound,outbound}` und ihre Einzelformen | **1** Treffer (`benutzerhandbuch.md:365`) |
| 2 | `awk` über Zeilen, die `dapter` **und** `inbound\|outbound` führen | **6** Rohzeilen — davon **eine** fehlerhaft (dieselbe wie bei 1); die fünf übrigen führen das Vokabular richtig |

Zähler 1 liefert die **rohe** Menge (eine Zeile), Zähler 2 eine größere, die erst durch die
Klassifikation „führt das Vokabular richtig" auf dieselbe eine zusammengeht — die fünf richtigen
Zeilen sind die Gegenprobe, nicht das Ergebnis.

**Geltungsbereich:** Prosa der drei Dokumente unter `docs/user/`. `benutzerhandbuch-standard.md`
und `releasing.md` führen das Vokabular **überhaupt nicht** (gemessen, nicht angenommen) — es gibt
dort also nichts nachzuziehen. Andere Strata wurden nicht gemessen. **Nicht dasselbe Muster:** `adapters/outbound` als Schicht-**Name**
(etwa in Test-Fixtures und in [`ADR-0028`](../../adr/0028-ziel-glob-schattenwurf.md)) ist kein
zweites Vorkommen — dort ist `outbound` ein Name, nicht der `direction`-Wert der Rolle.

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`docs/user/benutzerhandbuch.md`](../../../../docs/user/benutzerhandbuch.md) | update | die eine gemessene Stelle auf `driving`/`driven`; Änderungshistorie-Zeile |
| [`harness/conventions.md`](../../../../harness/conventions.md) | update | Zeile für die **Benutzer-Doku** in §Modus-Deklaration: Pfad-Familie `docs/user/`, Kürzel **`USER`** (vergeben, nicht nachgeschlagen — die Tabelle führte die Familie nicht), Achsen, Modus GF, Graduation `n/a` |
| [`CHANGELOG.md`](../../../../CHANGELOG.md) | update | Benutzer-Doku-Korrektur |

**Der Pflicht-Blick, den dieser Slice auslöst.** Jede Änderung an
[`harness/conventions.md`](../../../../harness/conventions.md) verlangt den **breiteren** Blick
([`AGENTS.md`](../../../../AGENTS.md) §1): Source Precedence, ID-Schema, Adaptionen. Die neue
Zeile darf kein Kürzel vergeben, das schon existiert (`SPEC`, `ADR`, `KERN`, `ADAPT`, `PLAN`,
`GATE`, `REVIEW`, `HARNESS`) oder den Zählraum der Beobachtungs-Kennung doppelt belegt — geprüft
wird gegen die Tabelle, nicht erinnert.

**Warum die Zeile zum Slice gehört.** Ohne sie hat dieser Fund **keinen** Register-Ort: alle
`evidence/`-Dateien tragen einen Vorgang, und die Kennung ist der Pfad `BEO-<KUERZEL>/<slug>` —
mit `<KUERZEL>` aus genau dieser Tabelle. Die Zeile ist damit die Bedingung dafür, dass der
Beleg geschrieben werden kann, nicht ein Nachzug danach.

**Auszuführende Gates:** `make gates` (tragend `doc-check`, `doc-structure`, `doc-mentions`,
`doc-planning`), zum Abschluss `make verify`.

## 4. Definition of Done

- [x] Das Benutzerhandbuch führt für die Adapter-Rolle ein Vokabular; die Messung (§2) fand
      **eine** Stelle, und die ist erledigt.
- [x] Die Benutzer-Doku hat eine Zeile in §Modus-Deklaration der
      [`harness/conventions.md`](../../../../harness/conventions.md) mit deklariertem Kürzel.

- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei. **Keine** Abhängigkeit von
[slice-194](../done/wellenlos/slice-194-portscope-richtungssegment.md) — dieser Slice zitiert dort nichts und
ändert dort nichts; die beiden berühren verschiedene Sub-Areas.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Erweist sich die Kürzel-Vergabe als Adaptions-pflichtig
  (also als Abweichung gegenüber der Baseline statt als Deklaration), sind es **zwei** Slices:
  die Doku-Korrektur hier, die Adaption daneben.
- `in-progress` → `open` (blockiert): Kollidiert der Kürzel-Vorschlag mit einem bereits
  vergebenen, ist die Zeile keine Deklaration mehr, sondern eine Umbenennung — und die ist ein
  eigener Vorgang mit eigenem Register-Nachzug.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag — und der
Register-Beleg liegt unter dem Kürzel, das dieser Slice vergibt. Die drei Paarungen trägt im
wellenlosen Repo die Slice-Closure selbst (§8).

## 7. Risiken und offene Punkte

- **Der Nachzug trifft nur die zitierte Zeile.** Ein zweites Vorkommen bleibt stehen, und das
  Dokument widerspricht sich danach **leiser** als vorher. Dagegen steht die Messung in §2.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: Zwei verschieden gebaute Zähler nennen
  **eine** Stelle (§2), und die beiden anderen Dokumente unter `docs/user/` führen das Vokabular
  nicht — es gibt keine zweite Stelle, die stehenbleiben könnte.
- **Das Kürzel ist eine Setzung, die keiner nachschlagen kann.** Es entsteht in diesem Slice;
  „nachgeschlagen, nicht erfunden" gilt für seinen **Gebrauch**, nicht für seine Vergabe.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: `USER` ist gegen die Tabelle geprüft
  (frei) und steht ab jetzt dort, wo jeder Lauf ihn nachschlägt; sein **erster** Gebrauch ist der
  Register-Pfad `BEO-USER/…` dieses Slice.
- **Die Zeile könnte die Modus-Deklaration als Ganzes berühren.** Wird die Handbuch-Familie
  nicht als eigene Sub-Area anerkannt, ist die Antwort eine Änderung an einer Deklaration, die
  jeder Lauf liest. — **Ausgang:** *entfallen*, gestrichen mit Begründung: Die Modus-Deklaration
  ist **keine** Adaption — sie deklariert, sie weicht nicht ab, und die Zeile entsteht im selben
  Bau wie die acht bestehenden. Ein `MR`-Eintrag wäre das falsche Werkzeug gewesen.

## 8. Closure-Notiz

**Lerneintrag — Form: benannte Spec-Lücke.** *Ein Fund kann keinen Ort haben — und dann zählt er
nicht.* Der Vokabel-Widerspruch war **gemessen** und trotzdem nirgends ablegbar: Die Kennung des
Registers ist der Pfad `BEO-<KUERZEL>/<slug>`, und für `docs/user/` führte die Modus-Deklaration
**keine** Zeile — kein Kürzel, kein Pfad, kein Beleg. **Weil** eine Ablage ohne Adresse nichts
aufnimmt, war die Deklarations-Zeile die **Bedingung** dieses Slice, nicht sein Nachzug.

**Der zweite Lerneintrag ist der kleinere und allgemeinere:** *Zwei Vokabulare für dieselbe Rolle
sind kein Tippfehler, sondern eine Fassung, die nicht nachgezogen wurde.* Das Handbuch trug
`inbound`/`outbound` an der Adapter-Rolle über den Umbau hinaus, während §4 und die Regel-Tabelle
**derselben Datei** längst das Gegenteil sagten. Gefunden hat das ein **Leser**, nicht ein Lauf —
und der Konsument, nicht das Repo.

**Steering-Loop-Eintrag:** gezählt, nicht verkörpert.
[`BEO-USER/handbuch-vokabel-der-adapter-rolle`](../observations/BEO-USER/handbuch-vokabel-der-adapter-rolle/observation.md)
ist **neu angelegt** (1×); die Schwelle ist nicht erreicht, ein Ausgang nicht fällig.

**Beobachtungs-Register ([`../observations/`](../observations/README.md)):** ein Verzeichnis
**neu angelegt** — `BEO-USER/handbuch-vokabel-der-adapter-rolle/`, Beleg `evidence/slice-195.md`,
Zähler **1×**. Es ist zugleich der erste Eintrag unter dem Kürzel `USER`, das dieser Slice in der
Modus-Deklaration vergibt.

**Folge-Slices:** keine.

**Risiken aus §7:** alle drei *entfallen*, gestrichen mit Begründung — siehe dort.

**Drei Paarungen:** Anker — kein `liegt in`-Feld gesetzt, weil nichts verkörpert wurde; nichts zu
paaren · Folge-Slice — keiner genannt · Register getragen (der genannte Pfad existiert und trägt
einen Beleg).

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt sind die **Benutzer-Doku** (`docs/user/`, als
Sub-Area bislang undeklariert — genau das behebt dieser Slice) und `HARNESS`
([`harness/conventions.md`](../../../../harness/conventions.md) ist namentlich Teil dieser
Sub-Area). Die Handbuch-Familie erfüllt die Schwelle über zwei Achsen: eigene Pfad-/Datei-Familie
(`docs/user/`) und eine sinnvolle eigene Diskrepanz-Zeile — dieser Fund **ist** sie.

**Vorgelagert — offene Beobachtungen sichten:** Register über die berührten Kürzel gelesen:
`HARNESS` führt
[`chronik-in-gelesenen-dateien`](../observations/BEO-HARNESS/chronik-in-gelesenen-dateien/observation.md)
(3×, Ausgang *geplant* → [slice-191](../open/slice-191-chronik-phrasen-sensor.md)) und
[`hard-rule-37-ohne-sensor`](../observations/BEO-HARNESS/hard-rule-37-ohne-sensor/observation.md)
(2×) — beide betreffen andere Gegenstände. Für die Benutzer-Doku gibt es **keinen** Treffer, weil
es für sie **kein** Kürzel gibt: der Befund, der diesen Slice trägt, ist zugleich der Grund,
warum er im Register bisher nicht stehen kann.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF. Kein Block je Sub-Area nötig.
