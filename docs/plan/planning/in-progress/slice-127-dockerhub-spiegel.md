# slice-127 — Docker-Hub-Spiegel

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Absicht 2026-08-30 („noch nicht. Aber das werden wir angehen"). Muster aus
dem Schwester-Repo (`packaging/dockerhub/`, `hub-description.yml`, Spiegel-Schritt im Release).

**Berührte Spec-Stellen:**
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk) ·
[AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)

**Verantwortlich:** — *(bis zur Priorisierung)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Das Image liegt zusätzlich auf Docker Hub — als **Spiegel** desselben Bildes, mit gepflegter
Darstellung und einer Pin-Regel, die beide Registries trägt.

## 2. Definition of Done

- [ ] Die Spec-Kette trägt den Spiegel: eine neue `AC-FA-DIST`-Kennung im
      [Lastenheft](../../../../spec/lastenheft.md) (Zusage: **dasselbe** Bild, nicht ein zweiter
      Bau) und eine ADR für die **Digest-Entscheidung** aus §3 samt Index-Eintrag.
- [ ] Der Release spiegelt **fail-closed** und prüft die Gleichheit nach dem Push; die Darstellung
      (`packaging/dockerhub/`, Kurztext + Overview) wird gesetzt, aber **nicht** fail-closed — der
      Spiegel ist die Zusage, der Beschreibungstext ist Präsentation.
- [ ] Die Betriebs-Vorbedingungen stehen in
      [`docs/user/releasing.md`](../../../../docs/user/releasing.md): Token-Scope, Repository-Name,
      und was ein grüner Lauf **nicht** sagt.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

**Spec-first** ([`AGENTS.md`](../../../../AGENTS.md) §5): Lastenheft → ADR → Spezifikation →
Workflow → Doku.

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| [`spec/lastenheft.md`](../../../../spec/lastenheft.md) | update | neue `AC-FA-DIST`-Kennung, Versions-Bump, Historie |
| `docs/plan/adr/00NN-*.md` + [Index](../../adr/README.md) | neu | die Digest-Entscheidung unten |
| `packaging/dockerhub/` | neu | `description.txt`, `overview.md`, `README.md` (Betriebswissen) |
| `.github/workflows/release.yml` | update | Spiegel-Push + Gleichheits-Prüfung |
| `.github/workflows/hub-description.yml` | neu | Darstellung, zwei Eingänge |
| [`docs/user/releasing.md`](../../../../docs/user/releasing.md) | update | Vorbedingungen |
| [`tools/image-scan.sh`](../../../../tools/image-scan.sh) | update | der Hub-Ref gehört in `IMAGE_SCAN_REFS` — **erst wenn er ein Bild trägt** |

**Auszuführende Gates:** `make gates` (tragend `doc-check`, `gate-consistency` für die
Pin-Konsistenz), `make image-test`. Zum Abschluss `make verify`.

### Die eine Entscheidung, die dieser Slice treffen muss

**Der Manifest-Digest ist registry-lokal.** Er hängt an der Blob-Kompression der jeweiligen
Registry; ein von GHCR kopierter Digest löst auf Docker Hub **nicht** auf. Das trifft a-check
härter als das Schwester-Repo, denn hier steht das ganze Pin-Regime auf **einem** Digest an
**neun** Stellen — [`a-check.mk`](../../../../a-check.mk), beide READMEs (je dreimal) und
[`version.md`](../../../../version.md) (zweimal) —, und `gate-consistency` hält sie auf
**Gleichheit** ([AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).

Drei Wege, und die ADR muss sich für einen entscheiden:

| Weg | Was er kostet |
|---|---|
| **A — Pins bleiben GHCR-gebunden**, die Hub-Seite nennt den Unterschied und verweist auf „den Digest der Registry, aus der du ziehst" | Der Hub-Nutzer bekommt aus der Doku dieses Repos keinen fertigen Pin. Dafür bleibt das Pin-Regime unverändert und `gate-consistency` ungerührt. |
| **B — `version.md` führt beide Digests**, das Meta-Gate prüft je Registry | Ehrlicher für den Hub-Nutzer, aber jede Release-Prep pflegt zwei Werte, und das Gate braucht eine zweite Wahrheit statt einer. |
| **C — Hub nur per Tag** | Bricht mit [`AC-QA-03`](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) für die Hub-Seite; ein beweglicher Tag ist genau das, was der Digest verhindern soll. |

**Empfehlung: A.** Sie hält die Zusage, die zählt (*dasselbe Bild*) über die **Config**-Digest-
Gleichheit, die der Release prüft, und lässt das Pin-Regime unangetastet. Der Preis ist benannt
und steht auf der Hub-Seite, wo der Hub-Nutzer ihn liest.

### Was das Schwester-Repo teuer gelernt hat

Drei Punkte, die beim Nachbau durchrutschen und dort je einen Vorfall gekostet haben:

1. **Der Metadaten-`PATCH` braucht einen anderen Token-Scope als der Push.** Das Bild lag auf
   Docker Hub, die Beschreibung blieb leer — `Forbidden`, weil der Token `read/write` statt
   `read/write/delete` trug.
2. **`continue-on-error` setzt `conclusion: success`**, lässt aber `outcome` auf `failure`. Ohne
   einen Folgeschritt, der `outcome` liest, ist der Lauf grün und die Seite leer.
3. **`workflow_dispatch` checkt den Origin-Stand aus.** Ein Lauf setzte die Seite aus einem alten
   Commit, obwohl der neue Text lokal schon geschrieben war. Erst pushen, dann dispatchen; die
   Gegenprobe ist die Hub-API, nicht der Ausgang des Laufs.

Dazu die Byte-Falle: Docker Hub misst die Description in **Zeichen**, nicht Bytes — `wc -m`, nicht
`-c`. Bei Nicht-ASCII wäre eine Byte-Messung strenger als die Regel und meldete Text rot, den die
Plattform annimmt.

## 4. Trigger

**Start** (`open` → `in-progress`), zweiteilig:

1. **Das Release-Vorhaben aus der Sicherheits-Kette ist durch** —
   [slice-125](../done/slice-125-go-toolchain-1-27.md) wartet auf ein Release, und zwei
   Release-Umbauten gleichzeitig sind einer zu viel.
2. **Die Betriebs-Vorbedingungen stehen**: ein Docker-Hub-Repository, `DOCKERHUB_USERNAME` und ein
   `DOCKERHUB_TOKEN` mit Scope `read/write/delete`. Ohne sie ist der Slice nicht abschließbar,
   sondern nur schreibbar.

**Rückführungen:**

- `in-progress` → `open`: falls die Digest-Entscheidung eine Änderung an
  [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) verlangt. Das wäre ein
  eigener CR am Lastenheft und nicht Teil dieses Slice.

## 5. Closure-Trigger

Spec-Kette steht, Spiegel läuft fail-closed, Darstellung gesetzt, Vorbedingungen dokumentiert,
Gates grün.

**Was bewusst nicht getan wird:** den **Hub-Ref in den CVE-Sensor** aufnehmen, solange er kein
Bild trägt — ein Ref ohne Bild machte den Nachtlauf ab dem ersten Tag rot. Er kommt dazu, wenn
der erste Spiegel-Push durch ist. Ebenso wenig wird die **Kategorie** automatisiert: die Action
hat dafür keinen Input, sie bleibt eine Entscheidung im Web-UI und gehört als Text ins
`packaging/`-README, statt still dort zu leben.

## 6. Risiken und offene Punkte

- *Die Digest-Entscheidung könnte [`AC-QA-03`](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) berühren statt nur zu ergänzen* —
  **Ausgang:** <bei Closure>
- *Ein zweiter Registry-Push macht das Release länger und fehleranfälliger; er steht nach dem
  GHCR-Push, damit die Quelle nie hinter dem Spiegel zurückliegt* — **Ausgang:** <bei Closure>
- *Die Hub-Darstellung ist das einzige Repo-Material, das ungelesen auf einer fremden Plattform
  steht* — **Ausgang:** <bei Closure>

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5.)_

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>**

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden die **Spec-Straten** (neue Anforderung,
ADR), die **CI-Schicht** (`release.yml`, neuer Workflow), eine **neue Sub-Area** `packaging/` und
die **Benutzer-Doku** (`releasing.md`).

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-026`](../observations.md) (die Form der
Workflow-`uses:`-Einträge ist ungeprüft) trifft diesen Slice unmittelbar — er fügt einen Workflow
mit einer **fremden** Action hinzu (`peter-evans/dockerhub-description`), und deren Pin prüft
heute kein Gate.

Alle berührten Sub-Areas GF.
