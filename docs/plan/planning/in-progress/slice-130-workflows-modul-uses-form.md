# slice-130 — Die Form der `uses:`-Einträge wird geprüft, nicht gelesen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** [`BEO-026`](../observations.md) — bei **3×**: zwei Vorfälle standen beim Schnitt (§3),
der dritte entstand während der Arbeit (§8).

**Berührte Spec-Stellen:** — *(keine)* — die CI-Schicht ist nicht Gegenstand des Lastenhefts.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Die Deklarations-Form der `uses:`-Referenzen in `.github/workflows/` hängt an keinem Lauf. Das
Modul `workflows` ist seit dem Pin `v0.67.0` verfügbar und war nie eingeschaltet.

## 2. Definition of Done

- [x] Modul `workflows` in [`.d-check.yml`](../../../../.d-check.yml) konfiguriert, als
      `make doc-workflows` aufrufbar und im `gates`-Aggregat — mit dem Aktivierungs-Schalter
      `dir:`, ohne den das Modul inert wäre (dieselbe Klasse wie
      [`BEO-014`](../observations.md)).
- [x] Der vom Modul gefundene Rechte-Defekt in
      [`release.yml`](../../../../.github/workflows/release.yml) ist behoben: der aufrufende Job
      führt die Rechte, die das lokale Ziel verlangt.
- [x] Der widersprüchliche Tag-Kommentar ist korrigiert — **gegen die Quelle gemessen**, nicht
      aus der Nachbarzeile geschlossen.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| [`.d-check.yml`](../../../../.d-check.yml) | update | `workflows`-Block mit `dir:` |
| [`Makefile`](../../../../Makefile) | update | `doc-workflows` + Eintrag im `gates`-Aggregat |
| [`AGENTS.md`](../../../../AGENTS.md) §4, [`harness/README.md`](../../../../harness/README.md) | update | neues Target; `doc-targets` erzwingt die Deklaration |
| [`release.yml`](../../../../.github/workflows/release.yml) | update | Rechte-Deklaration und Tag-Kommentar |

**Auszuführende Gates:** `make doc-workflows` (der Gegenstand), `make gates` — tragend
`doc-targets` und `gate-consistency`.

### Der zweite Vorfall ist von mir, und er ist die schärfere Ausprägung

[`BEO-026`](../observations.md) stand bei 1×: beim Schreiben von
[`image-scan.yml`](../../../../.github/workflows/image-scan.yml) entstand `# v5.0.0` an einem
Digest, den `ci.yml` als `# v6.0.2` führte. Der zweite Vorfall liegt in **einer** Datei:

```text
release.yml:75   docker/login-action@650006c6eb7d…  # v4.2.0   (2026-06-21, slice-007)
release.yml:158  docker/login-action@650006c6eb7d…  # v3.6.0   (2026-08-30, slice-127)
```

Derselbe Digest, zwei Versionen, 83 Zeilen auseinander. Die untere Zeile ist beim Docker-Hub-Spiegel
entstanden: der Digest wurde aus der oberen kopiert, der Kommentar **danebengeschrieben**. Das ist
die Klasse aus [`BEO-031`](../observations.md) — kopiert statt gemessen — nur an einem Feld, das
keinen Prüfer hat.

**Gemessen** (`gh api repos/docker/login-action/…`, kein Netz-Zugriff aus dem Gate heraus):

| Tag | Commit |
|---|---|
| `v3.6.0` | `5e57cd11…` |
| `v4.2.0` | `650006c6…` |

Der Digest **ist** v4.2.0; der Kommentar in Zeile 158 ist falsch. Ohne diese Abfrage ließe sich nur
sagen, dass **einer** von beiden falsch ist — nicht welcher.

### Was das Modul deckt — und was ausdrücklich nicht

Das ist die Stelle, an der dieser Slice weniger liefert, als sein Anlass verspricht.

Das Modul prüft die **Form** einer fremden Referenz: voller 40-stelliger SHA
(`uses-pin-missing`) mit Tag-Kommentar dahinter (`uses-pin-untagged`). Es prüft **nicht**, ob der
Kommentar zum SHA **passt** — das wäre Netz, und d-check ist hermetisch.

**Der Widerspruch aus dem Anlass bleibt damit ungedeckt.** Beide Zeilen oben tragen einen
Tag-Kommentar; beide sind formal in Ordnung. Ein Sensor, der sie trennt, existiert nicht — dieser
Slice korrigiert den Fehler, er verhindert den nächsten nicht.

Was das Modul stattdessen deckt, ist die **lokale** Referenz: existiert das Ziel, und führt der
aufrufende Job die Rechte, die es verlangt. Dort hat es beim ersten Lauf etwas gefunden (§4).

## 4. Der erste Lauf findet einen latenten Release-Bruch

Probeweise gefahren, bevor dieser Slice geschnitten war — ein Befund:

```text
.github/workflows/release.yml:242	.github/workflows/hub-description.yml	uses-local-perms-undeclared
```

Der Sachverhalt, drei Zeilen aus zwei Dateien:

| Ort | Deklaration |
|---|---|
| `release.yml` Kopf | `permissions: {}` |
| `release.yml` Job `hub-description` | *(kein eigenes `permissions:`)* |
| `hub-description.yml` Job `sync` | `permissions: contents: read` |

Ein aufgerufener Workflow bekommt nur, was der aufrufende **Job** selbst führt. Der Job erbt den
Workflow-Kopf `{}` und kann `contents: read` nicht weitergeben. GitHub bricht diesen Fall **vor dem
ersten Job** ab — ohne Log, ohne Job, ohne Hinweis auf die schuldige Zeile.

**Das ist kein hypothetischer Fall, sondern der beobachtete.** Beim Release v0.18.0 musste die
Hub-Darstellung von Hand über `workflow_dispatch` gestartet werden; dort gilt der eigene
Workflow-Kopf statt des ererbten. Die Ursache stand seit
[slice-127](../done/slice-127-dockerhub-spiegel.md) im Repo und ist keinem Gate aufgefallen, weil
kein Gate dorthin sah.

## 5. Abgrenzung

- **Kein Versions-Bump.** `docker/login-action` steht bei v4.2.0, verfügbar ist v4.6.0. Das Heben
  ist der Kanal aus [slice-128](../done/slice-128-dependabot-hebungskanal.md), nicht die Hand.
- **Keine Lauffähigkeits-Aussage.** Das Modul deckt **eine** Deklarations-Klasse; ein grüner Lauf
  heißt „diese Klasse liegt nicht vor", nicht „der Workflow läuft".
- **Kein `doc-workflows` im erzeugten Fragment.** [`d-check.mk`](../../../../d-check.mk) ist
  verbatim aus `--print-mk` und führt für dieses Modul **kein** Target, obwohl der Pin es kennt
  (es steht in den `--disable`-Listen). Das Target entsteht darum im eigenen
  [`Makefile`](../../../../Makefile); das Fragment bleibt unberührt.

## 6. Risiken und offene Punkte

- *Die Rechte-Korrektur könnte den Release-Job selbst berühren und einen laufenden Pfad brechen* —
  **Ausgang:** gestrichen mit Begründung: die Änderung **gibt** dem aufrufenden Job zwei Zeilen
  `permissions:`, sie nimmt keine. Der `release`-Job und der Spiegel-Schritt sind unberührt; der
  einzige Pfad, dessen Rechte sich ändern, ist der, der heute nicht läuft.
- *Das Modul könnte im Bestand weitere Befunde erzeugen, die den `gates`-Lauf rot machen* —
  **Ausgang:** gestrichen mit Begründung: gemessen, nicht vermutet — **genau ein** Befund im
  ganzen Bestand (§4), behoben, danach `doc-workflows` 0 Befunde und `make gates` **EXIT=0**.
  Die vier Workflow-Dateien tragen acht `uses:`-Referenzen — **sieben** fremde, alle formgerecht
  gepinnt, und **eine** lokale, die genau den Befund trug.
- *Der ungedeckte Digest↔Tag-Widerspruch bleibt ohne Sensor und wiederholt sich* —
  **Ausgang:** **eingetreten**, noch in diesem Slice, an einer dritten Stelle (§8) ⇒ Folge-Slice
  benannt und [`BEO-026`](../observations.md) auf **3×** geschärft.

## 7. Closure-Notiz

**Geliefert:** Das Modul `workflows` ist konfiguriert, als `make doc-workflows` aufrufbar und im
`gates`-Aggregat. Es hat im ersten Lauf einen latenten Release-Bruch gefunden (§4), der behoben
ist; der widersprüchliche Tag-Kommentar in
[`release.yml`](../../../../.github/workflows/release.yml) trägt jetzt den **gemessenen** Wert.

**Lerneintrag — Form: geschärfte Regel.** *Wer einen Sensor wegen eines Vorfalls einschaltet,
prüft zuerst, ob er genau **diesen** Vorfall fängt — sonst kauft er Deckung für eine andere Frage
und hält den Anlass für erledigt.* Der Anlass war ein Digest mit zwei Tag-Kommentaren. Das Modul
prüft, **dass** ein Tag-Kommentar dasteht, nicht **welcher**; beide Zeilen des Anlasses sind
formgerecht. Hätte ich das nicht vor dem Schnitt nachgelesen, stünde `BEO-026` heute als
„verkörpert" im Register — *weil* ein grünes Target wie eine Antwort aussieht, auch wenn es eine
andere Frage beantwortet. Dieselbe Klasse wie [`BEO-023`](../observations.md), nur eine Ebene
höher: dort war die Prüfmenge leer, hier ist es die Prüf**frage**.

**Der Sensor hat trotzdem mehr geliefert als der Anlass verlangte.** Er deckt die *lokale*
Referenz, an die niemand gedacht hatte, und fand dort den Grund, warum die Hub-Darstellung bei
v0.18.0 von Hand gestartet werden musste (§4). Das ist die beobachtbare Architektur-Aussage: die
Rechte-Kette eines aufgerufenen Workflows ist eine **Deklaration**, keine Laufzeit-Eigenschaft —
und Deklarationen sind hermetisch prüfbar.

**Drei beobachtbare Closure-Kriterien:**

1. `make doc-workflows` meldete am Bestand **einen** Befund (`uses-local-perms-undeclared`,
   `release.yml:242`), nach der Korrektur **null**.
2. Der falsche Tag-Kommentar ist gegen die Quelle geprüft, nicht gegen die Nachbarzeile:
   `v4.2.0` → `650006c6…`, `v3.6.0` → `5e57cd11…`. Der Digest **ist** v4.2.0.
3. Zwei Bestands-Sensoren haben die Erweiterung selbst beanstandet, bevor sie grün wurde:
   `gate-consistency` (`doc-workflows` fehlte in `.PHONY`) und `guard-selftest` (es fehlte in der
   `GATES`-Liste des Command-Guard, sein Lauf wäre also pipebar geblieben). Beide Male war der
   Befund berechtigt — der Sensor für neue Targets funktioniert.

**Was die Disable-Liste betrifft:** das Target zählt die **sechs Module der `modules:`-Liste** auf,
nicht alle Module des Pins. Die Fragment-Targets tun Letzteres, weil sie generisch erzeugt sind und
die Konfiguration nicht kennen; von Hand wäre es eine geschlossene Liste gegen eine offene Menge —
die Falle aus [slice-115](../done/slice-115-dcheck-pin-v0670.md). Die sechs Namen stehen im selben
Repo und fallen beim Ändern auf.

**Offene Risiken und ihr Ausgang:** zwei gestrichen mit Begründung, eines **eingetreten** (§8).

**Beobachtungs-Register:** [`BEO-026`](../observations.md) ist von **1×** auf **3×** geschärft und
neu gefasst — die Klasse ist nicht „Workflow-`uses:`", sondern **jede Versions-Angabe neben einem
Digest**. Sie bleibt **offen**: die Form deckt dieser Slice, die Gültigkeit nicht.

**Folge-Slice:** die dritte Ausprägung aus §8 — die widersprüchlichen Versions-Variablen zwischen
[`Makefile`](../../../../Makefile) und [`Dockerfile`](../../../../Dockerfile).

## 8. Der dritte Vorfall entstand während dieses Slice

Beim Lesen des Makefile-Kopfs, aus einem anderen Grund:

| Ort | Angabe |
|---|---|
| [`Makefile`](../../../../Makefile) | `GO_VERSION ?= 1.26.4`, `GOLANGCI_LINT_VERSION ?= v2.12.2` |
| [`Dockerfile`](../../../../Dockerfile) `ARG` | `GO_VERSION=1.27.0`, `GOLANGCI_LINT_VERSION=v2.13.2` |

Das Makefile sticht den `ARG` per `--build-arg`. Was tatsächlich läuft, sagt aber **keine** der
beiden Zahlen, denn daneben steht ein Digest, und der sticht den Tag. **Gemessen** im gepinnten
Basis-Image: `go1.27.0`.

Die Hebung ist also im `ARG` angekommen und im Makefile nicht — folgenlos für den Build, aber die
Zahl dort behauptet seither etwas Falsches. Das ist dieselbe Klasse wie der Tag-Kommentar, nur
schärfer: dort widersprach eine Zahl einem Digest, hier widersprechen sich **zwei** Zahlen, und
gültig ist keine.

**Nicht in diesem Slice.** Die Größen-Regel lässt drei Liefer-Punkte, und die drei sind vergeben;
ein vierter würde den Slice dehnen statt zerlegen ([`AGENTS.md`](../../../../AGENTS.md) §5).
