# Review-Report: slice-220 — 2026-10-07

**Review-Art:** Code (Doku-Diff) — geprüft gegen den Slice-Plan (§1 Ziel und Abgrenzung, §1
Quellen), die Hard Rules und die Mess-Regeln; jede Tatsachenbehauptung über fremde Werkzeuge
zusätzlich gegen deren Primärquelle, frisch gelesen (nicht dem Slice-Zitat vertraut).
Unabhängiger Lauf in frischem Kontext (kein `fork`).

**Gegenstand:** Commit-Range `d374830..3a107ce` (ein Commit, `3a107ce`): `docs/user/benutzerhandbuch.md`
(neuer Abschnitt §1 „macOS: Docker Desktop und Colima", Version 1.47, Historie-Zeile) und
`CHANGELOG.md` (`[Unreleased]`).

**Skill:** `.harness/skills/reviewer.md` @ `e6917e3` · <!-- d-check:ignore -->
**Modell:** claude-opus-5-5 · **Datum:** 2026-10-07

> **Zitier-Form** *(dieser Block bleibt stehen — er ist Norm, kein
> Ausfüll-Hinweis; die `<Platzhalter>` darin sind Formbeispiele)*. Dieser
> Report friert ein; was er zitiert, bewegt sich
> weiter. Deshalb: **Kennung, nicht Adresse** — `slice-<Kennung>` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link
> (`v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt>). Der vendored Baum trägt
> genau einen Tag; der Sprung löscht den alten, und ein Link darauf färbt beim
> nächsten Bump ein Artefakt rot, das niemand mehr anfassen darf. Ein `pfad`-Feld
> auf den **geprüften Gegenstand** ist davon nicht betroffen — es zitiert den
> Stand des Laufs und darf ihn festhalten (`v<X.Y.Z>` ·
> `regelwerk/grundlagen-harness-dateien.md` §harness/README.md als
> Einstiegspunkt — diese Zeile ist selbst ein Beispiel der Form).

**Eingangs-Kontext:**

- slice-220 (Slice-Plan, Stand `in-progress/`)
- ADR-0043 (Multi-Arch, ein Bau, getestet, dann getaggt) — als Bezug der Release-Aussage
- AC-FA-DIST-001 (Plattformen `linux/amd64` und `linux/arm64`)
- `AGENTS.md` §3 (Hard Rules), §6 Schritt 7 (CHANGELOG); `harness/rules/mess-regeln.md`
- `.github/workflows/release.yml` (Stand `3a107ce`), `a-check.mk`
- **Fremdquellen, gelesen 2026-10-07 (Branch `main` bzw. Live-Doku):**
  Colima `embedded/defaults/colima.yaml`, `cmd/start.go`, `cmd/delete.go`, `app/app.go`
  (`getStatus`/`Status`), `config/files.go`, `environment/vm/lima/yaml.go`;
  Docker-Doku *Docker Desktop settings* (`docs.docker.com/desktop/settings-and-maintenance/settings/`)
  und *Bind mounts* (`docs.docker.com/engine/storage/bind-mounts/`).

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | Der Stolperstein sagt für beide Laufzeiten zu, ein Pfad außerhalb der Freigaben ergebe im Container ein **leeres** Verzeichnis, und nennt Docker Desktop im selben Absatz als Freigabe-Fall; die Docker-Desktop-Doku sagt im Satz direkt nach dem im Slice-Plan zitierten: „If your project is outside this directory then it must be added to the list, otherwise you may get **Mounts denied** or cannot start service errors at runtime." Dieselbe Pauschalaussage steht in der CHANGELOG-Zeile („ein Repo außerhalb erscheint im Container leer") und in der Historie-Zeile („leeres `/src`"). Für Colima ist sie haltbar (`-v` legt eine fehlende Quelle im Daemon-Host — der VM — als leeres Verzeichnis an, Docker-Doku *Bind mounts*). | nachweislich falsche Tatsachenbehauptung (Reviewer-Skill HIGH-Liste), hier gegen die vom Slice selbst zitierte Quelle; slice-220 §1 Ziel trägt denselben Satz | `docs/user/benutzerhandbuch.md:68-71`; `CHANGELOG.md:12-13`; `docs/user/benutzerhandbuch.md:1141` | nein — kein Gate liest Fremd-Doku; Gegenprobe ist die Quelle bzw. ein `docker run -v /opt/x:/src` unter Docker Desktop | Verhaltensaussage über ein fremdes Werkzeug von einem Fall auf alle verallgemeinert |
| F-2 | MEDIUM | Die Colima-Freigabe-Anweisung „mit `colima start --mount <pfad>`" wirkt nicht wie beschrieben: bei laufender Instanz bricht `colima start` mit „already running, ignoring" ab, ohne die Mounts anzuwenden (`cmd/start.go`), und eine nicht leere Mount-Liste ersetzt den Standard-Mount von `$HOME`, statt ihn zu ergänzen (`environment/vm/lima/yaml.go`: `~` nur bei `len(conf.Mounts) == 0`). Dasselbe gilt für den Weg über `mounts:` in `colima.yaml` (wirksam erst nach Neustart). Beides sagt der Text nicht. | Maintainability / Geltungsbereich einer Zusage (`harness/rules/mess-regeln.md` Regel 1); Fremdquelle Colima | `docs/user/benutzerhandbuch.md:71-72` | nein — Gegenprobe ist die Colima-Quelle bzw. ein Lauf auf macOS, den das Projekt nicht führt | Bedienanweisung für ein fremdes Werkzeug ohne Vorbedingung und Nebenwirkung |
| F-3 | LOW | Der Text sagt „die Aussagen unten stützen sich auf die Dokumentation von Docker Desktop und Colima"; der Slice-Plan belegt mit Zitat nur zwei davon (Docker-Freigaben; Colima `arch`/`$HOME`). `colima status` zeigt die Architektur, `--mount`, `colima delete` + `colima start`, der Konfigurationsort `~/.colima/default/colima.yaml` und das leere Verzeichnis unter Colima sind unbelegt im Slice — vom Reviewer gegen die Colima-Quelle bzw. die Docker-Doku *Bind mounts* nachgeprüft und **zutreffend** (Konfigurationsort: Default ohne `XDG_CONFIG_HOME` bzw. ohne vorhandenes `~/.config/colima`, `config/files.go`). | `harness/rules/mess-regeln.md` Regel 1 (Geltungsbereich eines Belegs) | `docs/user/benutzerhandbuch.md:55-56`; slice-220 §1 Quellen | nein | Quellenliste deckt die Aussagen nicht, für die sie zitiert wird |
| F-4 | LOW | „seine Ausgabe muss der auf `linux/amd64` gleichen" weitet den Gegenstand: `release.yml` vergleicht einen SHA-256 über Exit-Code, stdout und stderr **eines Scans** innerhalb von `make image-test`, nicht die Ausgabe des Image-Tests. „vor dem Versions-Tag" ist für Nutzer mehrdeutig: der Lauf wird **durch** den Git-Versions-Tag ausgelöst; zurückgehalten wird der Bild-Tag auf GHCR (der Workflow nennt ihn selbst „Versions-Tag"). Die Kernaussage — nativer `ubuntu-24.04-arm`-Runner, Gleichheit Pflicht vor `publish` — stimmt mit `release.yml` (auch am Tag `v0.23.0`) überein. | `harness/rules/mess-regeln.md` Regel 1; `.github/workflows/release.yml` (Jobs `test-arm64`, `publish`) | `docs/user/benutzerhandbuch.md:53-55` | ja — Abgleich gegen `release.yml` | Zusage benennt einen weiteren Gegenstand als den geprüften |
| F-5 | LOW | Der Einstiegssatz ist absolut („auf Apple Silicon läuft sie als `arm64`, und `docker run` zieht … `linux/arm64` — ohne Emulation"); der Colima-Absatz zwei Absätze tiefer nennt den Gegenfall (VM mit `--arch x86_64` → `linux/amd64` unter Emulation). Die Einschränkung steht, aber nicht an der Aussage. | Maintainability | `docs/user/benutzerhandbuch.md:43-45` vs. `:64-65` | nein | Absolute Aussage, deren Ausnahme erst später steht |
| F-6 | INFO | Der bestehende „Hinweis zum Image" (Zeile 35, nicht im Diff) nennt bereits „Docker Desktop auf macOS mit Apple Silicon zieht `linux/arm64`"; der neue Abschnitt sagt es ein zweites Mal, ohne Querverweis zwischen beiden. | Maintainability | `docs/user/benutzerhandbuch.md:35`, `:43-45` | nein | — |

**Adversarische Gegenprüfung F-1** (vor Übernahme, Skill-Pflicht): Gesucht wurde ein Docker-Desktop-
Fall, in dem ein nicht freigegebener Pfad leer statt mit Fehler erscheint. Die Settings-Seite nennt
nur den Fehler; die Bind-Mount-Seite beschreibt das Anlegen fehlender Quellen „on the Docker
host" — unter Docker Desktop ist das die VM, die Freigabe-Prüfung greift aber vorher auf dem
Mac-Pfad. Kein Beleg für die Leer-Lesart unter Docker Desktop gefunden; das Finding bleibt HIGH.
Für Colima wurde die Gegenrichtung geprüft und bestätigt (kein Freigabe-Filter, `-v` legt an).

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Docker Desktop: Default-Freigaben `/Users`, `/Volumes`, `/private`, `/tmp`, `/var/folders` | geprüft gegen Live-Doku, wortgleich — ohne Befund |
| Docker Desktop: Menüpfad *Settings → Resources → File sharing* | geprüft: „File sharing" ist Unterabschnitt des Tabs *Resources* (neben Advanced, Proxies, Network) — ohne Befund |
| Colima: `arch: host` Default, „cannot be changed after virtual machine is created" | geprüft gegen `embedded/defaults/colima.yaml` — ohne Befund |
| Colima: nur `$HOME` gemountet | geprüft (`colima.yaml` „$HOME is mounted as writable", `lima/yaml.go`) — ohne Befund (Lima-interne Zusatz-Mounts außer Betracht) |
| Colima: `colima status` zeigt Architektur | geprüft: `Status()` druckt `arch: <guest-arch>` — ohne Befund |
| Colima: Flag `--mount` | geprüft: `--mount`/`-V`, Standard schreibgeschützt (`:w` für schreibbar) — Existenz ohne Befund; Wirkung siehe F-2 |
| Colima: `colima delete` + `colima start` legt neu an | geprüft: `delete` „deletes everything and a startup afterwards is like the initial startup" — ohne Befund |
| „kein natives macOS-Programm" | geprüft: kein `darwin`-Build in `Makefile`/`Dockerfile`, GitHub-Release ohne Binär-Assets — ohne Befund |
| `a-check.mk` läuft unverändert, `$(DOCKER)` Standard `docker` | geprüft gegen `a-check.mk` (`DOCKER ?= docker`) — ohne Befund |
| Hard Rules §3.1–§3.6 | geprüft: kein Code, kein Spec-Stratum, keine ADR, kein Gate berührt; Commit-Message trägt `AC-FA-DIST-001` und `slice-220` — ohne Befund |
| §3.7 (Kommentar/Zustandsfeld) | nicht anwendbar — kein Kommentar, kein Zustandsfeld im Diff |
| Plan-Abgrenzung (§1 Out-of-Scope) | geprüft: keine weitere Laufzeit (Podman, Rancher, OrbStack), kein Code, kein Fragment — ohne Befund |
| Form: Handbuch-Version 1.46 → 1.47, Historie-Zeile 1.47, CHANGELOG `[Unreleased]` `### Added` | geprüft — Form ohne Befund; Inhalt der Historie- und CHANGELOG-Zeile siehe F-1 |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 1 |
| LOW | 3 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** Verhaltensaussage über ein fremdes Werkzeug von einem Fall auf alle verallgemeinert · Bedienanweisung für ein fremdes Werkzeug ohne Vorbedingung und Nebenwirkung · Quellenliste deckt die Aussagen nicht, für die sie zitiert wird · Zusage benennt einen weiteren Gegenstand als den geprüften · Absolute Aussage, deren Ausnahme erst später steht

## Verdikt

**Merge-blockierend:** ja — F-1 (HIGH) und F-2 (MEDIUM). F-1 ist eine Tatsachenbehauptung, die
der Satz direkt neben dem im Slice-Plan zitierten Quelltext widerlegt; sie steht an drei Stellen
(Handbuch, Historie, CHANGELOG) und im Slice-Ziel §1.

**Übergabe:** Findings gehen an den Implementer; da F-1 auch slice-220 §1 *Ziel* trägt, ist es
zugleich ein Plan-Defekt (Rückkante Review → Plan). Die Finding-Klassen gehen in die
Slice-Closure §7. Dieser Report ersetzt keine Verifikation (DoD, `make gates`/`make verify`).
