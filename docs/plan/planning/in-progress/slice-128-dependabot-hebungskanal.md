# slice-128 — Dependabot als Hebungs-Kanal

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Folge-Slice aus [slice-124](../done/slice-124-image-scan-cve.md) — der Sensor meldet,
dieser Kanal hebt. Muster aus dem Schwester-Repo (`.github/dependabot.yml`).

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** — *(bis zur Priorisierung)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Was `make image-scan` meldet, bekommt einen Weg nach oben, der nicht daran hängt, dass jemand
hinsieht.

## 2. Definition of Done

- [ ] `.github/dependabot.yml` deckt beide Ökosysteme dieses Repos — Module und Actions — mit
      `commit-message.prefix`, der eine Traceability-Kennung trägt (siehe §3).
- [ ] Eine ADR begründet **zwei** Entscheidungen: warum der Kanal existiert (der Sensor meldet,
      hebt aber nicht) und warum **kein** Automerge. Sie steht im ADR-Index.
- [ ] Die **Betriebs-Kopplung** ist dokumentiert, nicht nur die Datei: ohne die Repository-Schalter
      (`dependabot_security_updates`, Alerts) öffnet ein CVE **ohne** neues Upstream-Release
      keinen PR — der Kanal erreicht die Fundklasse des Sensors dann nur zur Hälfte.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `.github/dependabot.yml` | neu | der Kanal |
| `docs/plan/adr/00NN-*.md` + [Index](../../adr/README.md) | neu | die zwei Entscheidungen |
| [`docs/user/releasing.md`](../../../../docs/user/releasing.md) | update | die Repository-Schalter als Vorbedingung |

**Auszuführende Gates:** `make gates`, dazu `make trace-check` und `make commit-scope-check` über
die Commit-Range — sie sind der Gegenstand der Kopplung unten. Zum Abschluss `make verify`.

### Die Kopplung, und sie steht VOR der Konfiguration

**`make trace-check` verlangt von JEDER Commit-Message eine Kennung** — `AC-*`, `ADR-*`, `MR-*`
oder `slice-*`, geprüft in der PR-CI über die Commit-Range
([`ci.yml`](../../../../.github/workflows/ci.yml), [`AGENTS.md`](../../../../AGENTS.md) §5).
**Dependabots Botschaften tragen keine.** Ein PR wäre damit rot, bevor ihn jemand ansieht.

Die Lösung macht die neue Commit-Klasse gate-**konform**, statt sie auszunehmen: jeder
`commit-message.prefix` trägt die Kennung der ADR, die den Kanal erlaubt. `commits.exempt-pattern`
in [`.d-check.yml`](../../../../.d-check.yml) bleibt **unverändert** — eine Ausnahme hätte die
Zusage aufgeweicht, statt sie zu erfüllen.

**Zu messen ist das vor der Konfiguration, nicht danach:** eine Botschaft ohne Kennung gegen das
echte Gate halten, dieselbe mit Präfix daneben. Erst dann steht fest, dass der Präfix trägt.

### Der Zuschnitt ist kleiner als im Schwester-Repo — und das ändert die Form

| | Schwester-Repo | a-check |
|---|---|---|
| direkte Requires | 2 | **1** |
| indirekte Requires | 20 | **0** |
| externe Actions | 7 | **2** |

Zwei Folgen: `allow: dependency-type: all` ist dort die Bedingung, unter der der Eintrag überhaupt
etwas ausrichtet (dreizehn von vierzehn Befunden lagen **indirekt**) — hier wäre es wirkungslos,
weil es keine indirekten gibt. Und die `groups`-Bündelung, die dort Rauschen verhindert, hat hier
kaum etwas zu bündeln. Beides gehört geprüft, nicht kopiert.

### Was bewusst NICHT in den Kanal gehört

**`docker`.** Ein Digest-Hoist ist ein **bewusster Commit**, kein Auto-Update — dieselbe Linie wie
[ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md) und wie die beiden Hebungen in
[slice-125](../done/slice-125-go-toolchain-1-27.md) und
[slice-126](../done/slice-126-lint-pin-v2-13-2.md), die je mit gelesener Ausgabe von Hand liefen.
Der Trivy-Pin in [`tools/image-scan.sh`](../../../../tools/image-scan.sh) und der
`d-check`-Pin in [`d-check.mk`](../../../../d-check.mk) fallen ohnehin heraus: sie leben in keinem
Manifest, das der Kanal liest.

**Automerge.** Der Kanal öffnet PRs; geprüft wird mit `make ci`, und der Merge bleibt ein Akt.

## 4. Trigger

**Start:** eingetreten — [slice-124](../done/slice-124-image-scan-cve.md) hat den Sensor geliefert,
und sein Erstlauf hat neun Befunde gemessen. Ein Kanal ohne Sensor hätte nichts Gemessenes
gehoben; diese Bedingung ist erfüllt.

**Rückführungen:**

- `in-progress` → `open`: falls die Messung der Kopplung zeigt, dass der Präfix **nicht** trägt.
  Dann ist die Frage nicht die Konfiguration, sondern das Verhältnis von `trace-check` zu
  maschinell erzeugten Commits — und das ist ein eigener Entscheid.

## 5. Closure-Trigger

Kanal konfiguriert, Kopplung gemessen, ADR steht, Betriebs-Vorbedingung dokumentiert, Gates grün.

**Was bewusst nicht getan wird:** die **Repository-Schalter setzen**. Sie leben in der
GitHub-Oberfläche, nicht im Repo; dieser Slice kann sie nennen und ihre Wirkung erklären, setzen
muss sie der Maintainer. Ein Slice, der eine Zusage über einen Schalter behauptet, den er nicht
sieht, wäre eine Harness-Lüge.

## 6. Risiken und offene Punkte

- *Der Präfix könnte an `commit-scope-check` scheitern statt an `trace-check`* — die Scope-Regel
  gilt nur für `(planning)`, aber gemessen ist das nicht. **Ausgang:** <bei Closure>
- *Ein wöchentlicher PR-Strom ohne Merge-Disziplin wird zum Rauschen, das man wegklickt* —
  dieselbe Klasse wie ein dauerhaft rotes Abzeichen. **Ausgang:** <bei Closure>
- *Der Kanal deckt die Fundklasse des Sensors nur, wenn die Repository-Schalter an sind — und das
  ist außerhalb dieses Repos nicht prüfbar* — **Ausgang:** <bei Closure>

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5.)_

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>**

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden die **CI-Schicht** (`.github/`), die
**Spec-Straten** (ADR) und die **Benutzer-Doku** (`releasing.md`).

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-026`](../observations.md) (die Form der
Workflow-`uses:`-Einträge ist ungeprüft) liegt in derselben Schicht und bleibt offen — dieser
Slice fügt keinen Workflow hinzu, sondern eine Konfiguration, die künftige `uses:`-Hebungen
**erzeugt**. Damit wird die Beobachtung praktisch relevanter, nicht kleiner.

Alle berührten Sub-Areas GF.
