# slice-171 — Zwei falsche Aussagen im Konventions-Bestand korrigieren

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (der Closure-Trigger wäre die eigene DoD — kein
repo-weites Mehr).

**Bezug:** Ergebnis des Durchgangs durch alle aktiven Adaptionen am
2026-09-06 (Maintainer-geführt, Eintrag für Eintrag). Der Durchgang selbst
hat [`MR-018`](../../../../harness/conventions.md#mr-018) aufgelöst
([slice-170](../done/wellenlos/slice-170-mr018-aufloesen.md)); die beiden
hier notierten Funde blieben liegen.

**Berührte Spec-Stellen:** — *(keine)* — Harness-Konventionen ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

---

## 1. Ziel

Zwei Aussagen im Konventions-Bestand stimmen nicht mehr mit dem Repo
überein. Beide lösen auf, beide sind grün, keine wird von einem Gate
gesehen — sie fallen nur auf, wenn jemand sie liest.

## 2. Analyse (vor der Umsetzung)

### 2.1 [`MR-011`](../../../../harness/conventions.md#mr-011) begründet sich mit einem Feld, das die genannte Beziehung nicht trägt

Der Eintrag rechtfertigt, warum a-check die Suffix-Form der Baseline nicht
übernimmt, unter anderem so: *„die Verfeinerungs-Beziehung steht in a-check
ohnehin explizit im `Schärft:`-Feld und nicht in der Kennung."*

**Gemessen (2026-09-06):**

| Frage | Ergebnis |
|---|---|
| Existiert `**Schärft:**` als Feld? | ja — **41** Vorkommen in [`docs/plan/adr/`](../../adr/README.md), eines in [`spec/lastenheft.md`](../../../../spec/lastenheft.md) |
| Welche Richtung trägt es dort? | ADR → `SPEC-*` (bzw. im Lastenheft `AC-*` → `AC-*`) |
| Trägt ein `SPEC-*`-Eintrag es? | **nein** — 7 `SPEC-*`-Überschriften in [`spec/spezifikation.md`](../../../../spec/spezifikation.md), **0** `**Schärft:**`-Felder |
| Wie steht die Verfeinerung dort wirklich? | als Prosasatz „Präzisiert [`AC-…`]" |

Die **Entscheidung** bleibt unberührt: der Maintainer hat am 2026-09-06
entschieden, **nicht** auf die Suffix-Form zu migrieren — ein Umbau träfe
alle bestehenden `SPEC-*` samt ihrer Traceability-Verweise, ohne eine
Aussage zu ändern. Falsch ist nur die **Begründung**, und die trägt der
Eintrag als Beleg.

**Der tragfähige Grund steht daneben und ist stärker:** die `SPEC-*`-Anker
werden aus `Accepted`-ADRs referenziert, und die sind immutabel
([`AGENTS.md`](../../../../AGENTS.md) §3.5). Eine Umbenennung erzeugte
denselben unauflösbaren Widerspruch zwischen `make doc-check` und
`make doc-immutable`, den der
[Anforderungs-Anlege-Prozess](../../../../harness/conventions.md#anforderungs-anlege-prozess)
für `AC-*`-Überschriften bereits beschreibt.

**Weg:** ein akzeptierter Eintrag wird nicht überschrieben
([`conventions.md`](../../../../harness/conventions.md) §Adaptions-Block).
Also ein **neuer** `MR`-Eintrag, der [`MR-011`](../../../../harness/conventions.md#mr-011) ablöst und die Begründung auf
den gemessenen Stand stellt; [`MR-011`](../../../../harness/conventions.md#mr-011) wandert nach `conventions/done/`.
Präzedenz: [`MR-017`](../../../../harness/conventions.md#mr-017) →
[`MR-020`](../../../../harness/conventions.md#mr-020)
([slice-166](../done/slice-166-mr020-adr-vorlage-generisch.md)). Die
Kennung des neuen Eintrags wird beim Schreiben vergeben, nicht hier — IDs
werden referenziert, nicht erfunden
([`AGENTS.md`](../../../../AGENTS.md) §5).

### 2.2 [`harness/README.md`](../../../../harness/README.md) deklariert eine aufgelöste Adaption

Die Rollen-Tabelle führt beide Validator-Kanten als *„unverkörpert,
deklariert als [`MR-009`](../../../../harness/conventions.md#mr-009)"*. [`MR-009`](../../../../harness/conventions.md#mr-009)
ist seit [slice-163](../done/slice-163-adaptions-durchgang-v610.md)
**aufgelöst** und liegt in `conventions/done/`; aktiv ist
[`MR-016`](../../../../harness/conventions.md#mr-016).

Kein Gate fängt das: der Link löst auf, weil die Doppel-Anker-Disziplin den
alten Slug am Leben hält. Genau dafür ist sie da — sie schützt den Verweis,
nicht seine Richtigkeit.

**Im selben Abschnitt zu prüfen:** der Absatz *Kontext-Trennung, real
angewandt* beschreibt die Review-Serie vom 2026-07-26 als Selbst-Review und
schließt mit *„ein unabhängiger Lauf bleibt eine eigene Übergabe"*. Seit
`welle-14` ist der unabhängige Lauf die Regel (slice-161…slice-170, je ein
Report aus getrenntem Kontext). Der Satz beschreibt seine Serie korrekt,
liest sich aber als Gegenwart — zu entscheiden ist, ob er bleibt,
umformuliert wird oder um den neuen Stand ergänzt.

## 3. Umsetzung

*(offen — entsteht mit der Umsetzung)*

## 4. Definition of Done

- [ ] [`MR-011`](../../../../harness/conventions.md#mr-011) durch einen neuen `MR`-Eintrag abgelöst, dessen Begründung
      den gemessenen Stand trägt; [`MR-011`](../../../../harness/conventions.md#mr-011) in `conventions/done/`, beide
      Tabellen in `conventions.md` nachgezogen, Anker `mr-011` erhalten.
- [ ] `harness/README.md` §Rollen nennt [`MR-016`](../../../../harness/conventions.md#mr-016) statt [`MR-009`](../../../../harness/conventions.md#mr-009); der Absatz
      zur Kontext-Trennung ist entschieden (bleibt / ergänzt / umformuliert).
- [ ] Unabhängiger Review durchgeführt (Report unter `docs/reviews/`).
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe und WIP-Limit frei.

**Rückführungen:** wächst der Umfang über die zwei Korrekturen hinaus —
etwa weil der Durchgang weitere überholte Aussagen im selben Abschnitt
findet — zurück nach `next/` zur Zerlegung.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz
geschrieben.

## 7. Risiken und offene Punkte

- *Der neue Eintrag wiederholt den Fehler von [`MR-011`](../../../../harness/conventions.md#mr-011) in anderer Form —
  eine Begründung, die plausibel klingt und nicht gemessen ist* — Ausgang
  bei Closure.
- *Die Ablösung von [`MR-011`](../../../../harness/conventions.md#mr-011) entfernt einen der fünf `v6.0.0`-Anker im Feld
  `Ersetzt-Baseline-Regel`; vier bleiben, und die Querschnitts-Entscheidung
  über sie ist damit weiterhin offen* — Ausgang bei Closure; die Klasse
  liegt im Register
  ([`BEO-HARNESS/zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md)).

## 8. Closure-Notiz

*(wird beim Übergang nach `done/` geschrieben; Lerneintrag — Form: wird
dort benannt.)*

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Harness-Einstieg** (`harness/`, `AGENTS.md`), Greenfield, Schwelle ≥ 2/3
erfüllt.

**Vorgelagert — offene Beobachtungen sichten:** entsteht mit dem Übergang
nach `in-progress/` (der Register-Stand ist beim Anlegen ein anderer als
beim Beginn der Arbeit). Einschlägig sind absehbar
[`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
und
[`BEO-HARNESS/zwei-baseline-staende-nach-migrationsende`](../observations/BEO-HARNESS/zwei-baseline-staende-nach-migrationsende/observation.md).

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
