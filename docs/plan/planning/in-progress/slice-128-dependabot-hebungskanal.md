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

- [x] `.github/dependabot.yml` deckt beide Ökosysteme dieses Repos — Module und Actions — mit
      `commit-message.prefix`, der eine Traceability-Kennung trägt (siehe §3).
- [x] Eine ADR begründet **zwei** Entscheidungen: warum der Kanal existiert (der Sensor meldet,
      hebt aber nicht) und warum **kein** Automerge. Sie steht im ADR-Index.
- [x] Die **Betriebs-Kopplung** ist dokumentiert, nicht nur die Datei: ohne die Repository-Schalter
      (`dependabot_security_updates`, Alerts) öffnet ein CVE **ohne** neues Upstream-Release
      keinen PR — der Kanal erreicht die Fundklasse des Sensors dann nur zur Hälfte.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

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

- *Der Präfix könnte an `commit-scope-check` scheitern statt an `trace-check`* — **Ausgang:**
  gestrichen mit Begründung: **gemessen**, beide Formen sind grün. Die Scope-Regel gilt nur für
  `(planning)`, und `build(deps)`/`build(ci)` fallen nicht darunter — das war zu erwarten, aber
  ungeprüft.
- *Ein wöchentlicher PR-Strom ohne Merge-Disziplin wird zum Rauschen* — **Ausgang:** weiter offen
  im **Beobachtungs-Register** als Teil von `BEO-030`. Das Limit von fünf offenen PRs je Ökosystem
  ist eine Gegenmaßnahme, keine Lösung; ob der Strom trägt, sagt erst der Betrieb.
- *Der Kanal deckt die Fundklasse des Sensors nur bei gesetzten Repository-Schaltern, und das ist
  außerhalb dieses Repos nicht prüfbar* — **Ausgang:** weiter offen im **Beobachtungs-Register**
  (`BEO-030`), dokumentiert in [`releasing.md`](../../../../docs/user/releasing.md)
  §Vorbedingungen.

## 7. Closure-Notiz

**Geliefert:** `.github/dependabot.yml` für beide Ökosysteme dieses Repos,
[ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md) für die zwei Entscheidungen, und die
Betriebs-Vorbedingungen als eigener Abschnitt in
[`releasing.md`](../../../../docs/user/releasing.md).

**Lerneintrag — Form: geschärfte Regel.** *Eine neue Commit-Klasse wird gate-**konform** gemacht,
nicht ausgenommen — die Ausnahme gilt weiter, wenn der Anlass längst weg ist.* Der bequeme Weg lag
offen: `commits.exempt-pattern` in [`.d-check.yml`](../../../../.d-check.yml) um Dependabots
Präfix erweitern. Er wäre kürzer gewesen und hätte genau die richtige Klasse getroffen. Er nimmt
aber eine Zusage **zurück**, statt sie zu erfüllen: die Traceability-Pflicht sagt, dass jeder
Commit auf eine Entscheidung zeigt — und ein Dependabot-Commit tut das, nämlich auf die ADR, die
den Kanal erlaubt. *Weil* der Präfix diese ADR nennt, ist er kein Trick, sondern die wahre Aussage.
Die Ausnahme hätte behauptet, es gäbe keine Entscheidung dahinter.

**Drei beobachtbare Closure-Kriterien:**

1. Die Kopplung ist in **beide** Richtungen gegen das echte Gate gemessen, vor der Konfiguration:
   `build(deps): bump …` ⇒ **Exit 2**, `commit-untraceable`; dieselbe Zeile mit der Kennung von [ADR-0038](../../adr/0038-dependabot-als-hebungskanal.md) ⇒
   **Exit 0**. Ohne die rote Hälfte wäre ein Präfix, der nichts bewirkt, von einem wirksamen nicht
   zu unterscheiden.
2. `commits.exempt-pattern` ist **unverändert** — nachprüfbar am Diff dieses Slice.
3. Der Zuschnitt ist gemessen und die Konfiguration folgt ihm: **kein**
   `allow: dependency-type: all` (es gibt kein indirektes Require) und **kein** `groups` (bei einem
   Require gibt es nichts zu bündeln). Beides steht im Schwester-Repo und wäre hier eine Zeile ohne
   Gegenstand — kopiert statt geprüft.

**Was der Kanal nicht kann, und wo das steht:** ohne die Repository-Schalter öffnet ein CVE **ohne**
neues Upstream-Release keinen PR. Die Schalter leben in der GitHub-Oberfläche; dieser Slice kann sie
nennen und ihre Wirkung erklären, setzen muss sie der Maintainer. Ein Slice, der eine Zusage über
einen Schalter behauptet, den er nicht sieht, wäre eine Harness-Lüge — deshalb steht die
Vorbedingung in [`releasing.md`](../../../../docs/user/releasing.md) und nicht als Häkchen hier.

**Offene Risiken und ihr Ausgang:** der erste gestrichen mit Begründung (gemessen), die anderen
beiden weiter offen im Register.

**Beobachtungs-Register:** `BEO-030` neu angelegt (CI-Schicht, 1×, Beleg slice-128): die Wirkung
dieses Kanals hängt an zwei Schaltern außerhalb des Repos, und kein Gate kann sagen, ob sie stehen.

**Folge-Slices:** [slice-127](../open/slice-127-dockerhub-spiegel.md) (Docker-Hub-Spiegel) — sein
erster Trigger, das Release, ist mit **v0.18.0** eingetreten; offen bleibt der zweite (Hub-Repo und
Token-Scope). Dazu [`BEO-026`](../observations.md): der Kanal **erzeugt** künftige
`uses:`-Hebungen, deren Form heute kein Gate prüft — die Lücke wird dadurch praktisch relevanter,
nicht kleiner.
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden die **CI-Schicht** (`.github/`), die
**Spec-Straten** (ADR) und die **Benutzer-Doku** (`releasing.md`).

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-026`](../observations.md) (die Form der
Workflow-`uses:`-Einträge ist ungeprüft) liegt in derselben Schicht und bleibt offen — dieser
Slice fügt keinen Workflow hinzu, sondern eine Konfiguration, die künftige `uses:`-Hebungen
**erzeugt**. Damit wird die Beobachtung praktisch relevanter, nicht kleiner.

Alle berührten Sub-Areas GF.
