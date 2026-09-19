# slice-195 — Handbuch-Vokabel für die Adapter-Rolle

**Welle:** ohne Welle.

**Bezug:** Zweiter Fund der externen Meldung `hexslice-architecture` (2026-09-19), im Repo
zeilengenau nachgeprüft. Folgt dem Zuschnitt von
[`ADR-0036`](../../adr/0036-port-richtung-inbound-outbound.md).

**Berührte Spec-Stellen:** — (der Slice berührt kein Spec-Stratum; er zieht ein
Benutzer-Dokument auf eine `Accepted`-Entscheidung nach).

**Verantwortlich:** — (bis zur Priorisierung).

**Autor:** Claude, im Auftrag des Maintainers. **Datum:** 2026-09-19.

**Lerneintrag — Form:** wird bei Closure benannt (eine der drei Formen der Ziel-Form).

---

## 1. Ziel und Abgrenzung

**Ziel:** Das Benutzerhandbuch führt für die Adapter-Rolle **ein** Vokabular — das von
[`ADR-0036`](../../adr/0036-port-richtung-inbound-outbound.md) entschiedene (`driving`/`driven`) —,
und die Handbuch-Familie hat einen deklarierten Ort im Beobachtungs-Register.

**Ausgangslage (gemessen):** `docs/user/benutzerhandbuch.md:350` zeigt
`internal/adapters/{inbound,outbound}/…`; derselbe Text sagt ab `:726` („Ein Port *treibt* nichts
… ein Adapter ist nicht *eingehend*") und die Regel-Tabelle ab `:242` das Gegenteil. `inbound`
und `outbound` sind seit [`ADR-0036`](../../adr/0036-port-richtung-inbound-outbound.md) das
**Port**-Vokabular. Das Dokument widerspricht sich damit selbst für dieselbe Rolle.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Ein Nachzug der Lab-Config des Konsumenten** — *ein anderer Vorgang in einem anderen Repo:*
  dort steht ein Kommentar, der die Kopplung Glob-Präfix ↔ Richtungssegment festhält; er ist
  Gegenstand von [slice-194](../in-progress/slice-194-portscope-richtungssegment.md), nicht dieses Slice.
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

## 2. Ausgangsmessung (beim Übergang nach `in-progress/` zu wiederholen)

Gelesen, nicht gefahren: `docs/user/benutzerhandbuch.md:350` (Beispielstruktur §3.7) gegen
`:242` (Regel-Tabelle) und `:726-737` (§4 „Richtung"). Vor dem Umbau läuft ein Zähler über das
Dokument, der **beide** Vokabulare auffindet — sonst behebt der Slice die zitierte Zeile und
lässt die zweite stehen.

**Geltungsbereich:** Prosa des Handbuchs. Andere Dokumente unter `docs/user/` sind **nicht**
gemessen; findet der Lauf dort dasselbe Muster, ist die Entscheidung, ob es mitgeht, Teil dieses
Slice (§4).

## 3. Umsetzung

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| [`docs/user/benutzerhandbuch.md`](../../../../docs/user/benutzerhandbuch.md) | update | `:350` auf `driving`/`driven`; zweites Vorkommen aus der Messung (§2) mit; Änderungshistorie-Zeile |
| [`harness/conventions.md`](../../../../harness/conventions.md) | update | Zeile für die **Benutzer-Doku** in §Modus-Deklaration: Pfad-Familie `docs/user/`, Kürzel (Vorschlag `USER`), Achsen, Modus GF, Graduation `n/a` |
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

- [ ] Das Benutzerhandbuch führt für die Adapter-Rolle ein Vokabular; das zweite Vorkommen aus
      der Messung (§2) ist mit erledigt.
- [ ] Die Benutzer-Doku hat eine Zeile in §Modus-Deklaration der
      [`harness/conventions.md`](../../../../harness/conventions.md) mit deklariertem Kürzel.

- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei. **Keine** Abhängigkeit von
[slice-194](../in-progress/slice-194-portscope-richtungssegment.md) — dieser Slice zitiert dort nichts und
ändert dort nichts; die beiden berühren verschiedene Sub-Areas.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Erweist sich die Kürzel-Vergabe als Adaptions-pflichtig
  (also als Abweichung gegenüber der Baseline statt als Deklaration), sind es **zwei** Slices:
  die Doku-Korrektur hier, die Adaption daneben.
- `in-progress` → `open` (blockiert): Kollidiert der Kürzel-Vorschlag mit einem bereits
  vergebenen, ist die Zeile keine Deklaration mehr, sondern eine Umbenennung — und die ist ein
  eigener Vorgang mit eigenem Register-Nachzug.

## 6. Risiken und offene Punkte

- **Der Nachzug trifft nur die zitierte Zeile.** Ein zweites Vorkommen bleibt stehen, und das
  Dokument widerspricht sich danach **leiser** als vorher. Dagegen steht die Messung in §2.
  — **Ausgang:** <offen bis Closure>
- **Das Kürzel ist eine Setzung, die keiner nachschlagen kann.** Es entsteht in diesem Slice;
  „nachgeschlagen, nicht erfunden" gilt für seinen **Gebrauch**, nicht für seine Vergabe.
  — **Ausgang:** <offen bis Closure>
- **Die Zeile könnte die Modus-Deklaration als Ganzes berühren.** Wird die Handbuch-Familie
  nicht als eigene Sub-Area anerkannt, ist die Antwort eine Änderung an einer Deklaration, die
  jeder Lauf liest. — **Ausgang:** <offen bis Closure>

## 7. Closure-Notiz

*(bei Closure auszufüllen)*

## 8. Sub-Area-Prüfungen und Modus-Begründung

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
