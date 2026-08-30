# slice-130 — Die Form der `uses:`-Einträge wird geprüft, nicht gelesen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** [`BEO-026`](../observations.md) — bei **2×** nach dem zweiten Vorfall (§3).

**Berührte Spec-Stellen:** — *(keine)* — die CI-Schicht ist nicht Gegenstand des Lastenhefts.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Die Deklarations-Form der `uses:`-Referenzen in `.github/workflows/` hängt an keinem Lauf. Das
Modul `workflows` ist seit dem Pin `v0.67.0` verfügbar und war nie eingeschaltet.

## 2. Definition of Done

- [ ] Modul `workflows` in [`.d-check.yml`](../../../../.d-check.yml) konfiguriert, als
      `make doc-workflows` aufrufbar und im `gates`-Aggregat — mit dem Aktivierungs-Schalter
      `dir:`, ohne den das Modul inert wäre (dieselbe Klasse wie
      [`BEO-014`](../observations.md)).
- [ ] Der vom Modul gefundene Rechte-Defekt in
      [`release.yml`](../../../../.github/workflows/release.yml) ist behoben: der aufrufende Job
      führt die Rechte, die das lokale Ziel verlangt.
- [ ] Der widersprüchliche Tag-Kommentar ist korrigiert — **gegen die Quelle gemessen**, nicht
      aus der Nachbarzeile geschlossen.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

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
  **Ausgang:** *(beim Abschluss ausfüllen)*
- *Das Modul könnte im Bestand weitere Befunde erzeugen, die den `gates`-Lauf rot machen* —
  **Ausgang:** *(beim Abschluss ausfüllen)*
- *Der ungedeckte Digest↔Tag-Widerspruch bleibt ohne Sensor und wiederholt sich* —
  **Ausgang:** *(beim Abschluss ausfüllen)*

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice.)_
