# slice-189 — Voll-Abgleich: die zwei Spec-Straten

**Welle:** ohne Welle.

**Bezug:** Folge-Slice aus
[slice-187](../in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md) §1
(*es wäre ein anderer Vorgang*).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** [`spec/lastenheft.md`](../../../../spec/lastenheft.md)
und [`spec/spezifikation.md`](../../../../spec/spezifikation.md) sind der
Gegenstand — **welche** Kennungen berührt sind, benennt die Umsetzung.

**Verantwortlich:** —

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die zwei Spec-Straten sind gegen ihre vendorte Ziel-Form abgeglichen,
und **je Kandidat ist getrennt, was Doku-Pflege und was Change Request wäre**.
Letzteres wird **benannt, nicht ausgeführt** — das ist der Kern dieses Slice und
der Grund, warum er von slice-187 abgetrennt wurde.

| Paar | Kandidaten (Instrument: slice-187 §2) |
|---|---|
| [`spec/lastenheft.md`](../../../../spec/lastenheft.md) | 28 |
| [`spec/spezifikation.md`](../../../../spec/spezifikation.md) | 28 |

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Jede Änderung an einer bestehenden `AC-*`.** *Es wäre ein anderer Vorgang:*
  Das Lastenheft ist vertraglich abnahmebindend; eine Änderung ist ein **Change
  Request** mit Versions-Bump und Historie-Zeile
  ([`harness/conventions.md`](../../../../harness/conventions.md)
  §Anforderungs-Anlege-Prozess), keine Entscheidung eines Abgleich-Slice.
- **Die 19 grandfatherten `AC-*`.** *Bestand bleibt bewusst stehen:* Ihre Form
  ist seit slice-054 ausdrücklich ausgenommen, und die Grandfather-Liste wächst
  nicht mit ([`AGENTS.md`](../../../../AGENTS.md) §5).
- **`spec/architecture.md`.** *Schicht-Abgrenzung:* Rang 3, eigenes Paar, in
  slice-185 bereits abgeglichen.
- **Die dreizehn bereits abgeglichenen Paare.** Erledigt (slice-185/186/187/188).

## 2. Ausgangsmessung

Instrument, Parameter und Vorgehen: siehe
[slice-187](../in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md) §2 —
zitiert statt wiederholt, derselbe Aufruf, derselbe Lauf.

**Was die Zahl hier nicht sagt:** Ein Kandidat im Lastenheft kann drei sehr
verschiedene Dinge sein — eine fehlende **Form**-Regel (Doku-Pflege), eine
fehlende **Zusage** (Change Request), oder eine Stelle, an der a-check bewusst
schärfer ist als die Vorlage. Die Sortierung in diese drei ist die eigentliche
Arbeit; 56 Kandidaten sind ihre Reihenfolge.

## 3. Umsetzung

*(entsteht mit der Arbeit)*

## 4. Definition of Done

- [ ] Beide Paare sind abgeglichen; je Paar steht der Ausgang **mit seiner
      Prüf-Ebene** im Plan.
- [ ] Je Kandidat ist getrennt: **Doku-Pflege** (in diesem Slice ausgeführt) ·
      **Change Request** (benannt, mit dem, was er ändern würde) · **bewusst
      schärfer** (mit Begründung).
- [ ] Kein bestehendes `AC-*` ist inhaltlich geändert; `make doc-immutable`
      über die Commit-Range ist grün.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Zwei öffentliche Verträge sind berührt:
Rang 1 und Rang 2 der Source Precedence.

## 5. Trigger

**Start** (`open` → `in-progress`):
[slice-188](../open/slice-188-voll-abgleich-gate-und-skill.md) liegt in `done/` und das
WIP-Limit ist frei. **Reihenfolge ist Absicht:** Der Spec-Abgleich ist der
heikelste; er kommt zuletzt, wenn das Verfahren an vier Slices eingespielt ist.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Trägt eines der beiden Straten allein mehr
  als einen Durchgang, wird das andere abgetrennt.
- `in-progress` → `open` (blockiert): Verlangt ein Kandidat eine Änderung an
  einer bestehenden `AC-*`, ist das ein Change Request — der Slice geht zurück
  und die Entscheidung an den Maintainer.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Der Slice kann in einen Change Request kippen.** Eine Ziel-Form-Regel, die
  eine bestehende `AC-*` berührt, ist keine Doku-Pflege — und die Rückführung ist
  dann der richtige Ausgang, nicht ein stiller Nachtrag.
  — **Ausgang:** <offen bis Closure>
- **`doc-immutable` deckt die Spec-Straten nicht.** Das Modul `vcs` führt
  ausschließlich `docs/plan/adr/[0-9]*.md` (gemessen in slice-187 §3.5); eine
  unbeabsichtigte Änderung an einer `AC-*` meldet **kein** Lauf.
  — **Ausgang:** <offen bis Closure>
- **Ein umbenannter `AC-*`-Anker bricht eine `Accepted`-ADR.** Der Fall ist im
  Repo belegt und die Antwort steht in
  [`harness/conventions.md`](../../../../harness/conventions.md)
  §Anforderungs-Anlege-Prozess (zwei Anker, alter Slug bleibt).
  — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen — berührt ist `SPEC`, erstmals seit slice-185.)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso — **zwei** Quellen:
Register und der Review-Report des Vorgänger-Slice, siehe slice-187 §9.)*
