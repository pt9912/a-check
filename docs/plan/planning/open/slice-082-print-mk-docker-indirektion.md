# slice-082 — `--print-mk`: `$(DOCKER)` statt wörtlichem `docker`

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** Konsumenten-Befund vom 2026-08-09;
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk).
**Bezug:** [ADR-0004](../../adr/0004-distribution-image-mk.md); Geschwister
[slice-083](../open/slice-083-print-mk-digest-selbstbezug.md).

---

## 0. Trigger

**Beginn:** sofort. Kein Vertrag ist berührt, kein Entscheid offen.

**Reihenfolge:** **vor** [slice-083](../open/slice-083-print-mk-digest-selbstbezug.md). Beide ändern
[`internal/cli/cli.go`](../../../../internal/cli/cli.go) **und**
[`a-check.mk`](../../../../a-check.mk); die Fragment-Parität in
[`tools/image-test.sh`](../../../../tools/image-test.sh) erzwingt, dass beide Dateien im selben
Commit wandern. Nacheinander, nicht parallel.

**Rückführungen:**

- `in-progress` → `open`: falls sich zeigt, dass `DOCKER ?=` im Fragment mit der `include`-Position
  im Konsumenten-Makefile kollidiert und keine reihenfolge-unabhängige Form existiert.

## 1. Auslöser

**Mechanismus: das gelieferte Fragment ignoriert eine Indirektion, die das konsumierende Repo
bereits führt.**

[`internal/cli/cli.go:133`](../../../../internal/cli/cli.go) und `:136` — und damit das
ausgelieferte [`a-check.mk`](../../../../a-check.mk) — rufen `docker` **wörtlich**:

```make
a-check:
	docker run --rm --network none -v "$(CURDIR)":/src:ro $(A_CHECK_IMAGE) /src
```

Ein Konsument mit `DOCKER ?= podman` fährt seine eigenen Targets über podman und `make a-check`
über docker. Im Zweifel schlägt das nicht einmal fehl, sondern läuft gegen eine andere Runtime als
der Rest des Repos.

**Gemessen am 2026-08-09:** die Indirektion fehlt in dieser Werkzeugfamilie durchgängig —
`a-check.mk` (2 Stellen), `d-check.mk` (11), a-checks eigenes
[`Makefile`](../../../../Makefile) (4). Und `Makefile:24` zeigt die Inkonsistenz im eigenen Haus:
`DOCKER_BUILD := docker build …` existiert, `docker run` läuft daneben wörtlich. Die Idee ist da,
sie wurde nur nie auf `run` ausgedehnt.

**Kein Vertragsbruch.**
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
verlangt „ein `include`-bares Fragment mit digest-gepinntem `A_CHECK_IMAGE` und einem
`a-check`-Scan-Target" — `$(DOCKER)` statt `docker` verletzt davon nichts. Darum **kein
spec-first**.

## 2. Betroffene Module

Eine Schicht: **Distribution** —
[`internal/cli/cli.go`](../../../../internal/cli/cli.go) (die Fragment-Konstante) und
[`a-check.mk`](../../../../a-check.mk) (das committete Fragment). Ob a-checks eigenes
[`Makefile`](../../../../Makefile) mitzieht, ist Teil des Slice.

## 3. Auszuführende Gates

`make image-test` (Fragment-Parität), `make gates`.

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| Fragment mit `DOCKER=echo` aufgerufen | das Scan-Target ruft `echo run …` statt `docker run …` |
| Fragment ohne gesetztes `DOCKER` | unverändert `docker run …` — der Default bleibt |
| `cli.go` geändert, `a-check.mk` nicht | `make image-test` **rot** (Fragment-Parität) |

Die dritte Zeile ist keine Probe des Fixes, sondern der **Absicherung**: sie belegt, dass eine
halbe Änderung nicht durchkommt.

**Die Reihenfolge-Falle gehört dokumentiert.** `DOCKER ?= docker` im Fragment greift nur, wenn der
Konsument seine eigene Definition **vor** dem `include` setzt oder hart zuweist (`DOCKER = podman`).
Steht sein `DOCKER ?= podman` **nach** dem `include`, hat das Fragment die Variable bereits belegt.
Ohne diesen Hinweis tauscht der Slice ein stilles Problem gegen ein leiseres.

## 4. Was bewusst nicht getan wird

- **Der Digest-Selbstbezug.** Eigener Slice
  ([slice-083](../open/slice-083-print-mk-digest-selbstbezug.md)), weil er [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) berührt und
  spec-first ist.
- **`d-check.mk` patchen.** Fremdlieferung; ein Befund dort gehört gemeldet, nicht lokal
  korrigiert.

## 5. DoD

- [ ] Das `--print-mk`-Fragment nutzt `$(DOCKER)` mit `DOCKER ?= docker`, und die
      Reihenfolge-Bedingung ist im Fragment selbst kommentiert. Beleg: die ersten beiden Proben
      aus §3.
- [ ] `a-check.mk` und die Konstante in `cli.go` sind gemeinsam gewandert. Beleg:
      `make image-test` grün (Fragment-Parität).
- [ ] `make gates` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
