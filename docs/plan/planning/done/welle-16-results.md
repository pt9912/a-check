# Welle welle-16 — Sollform je Datei (`shapes`) — Closure-Notiz

> **Zitier-Form** *(bleibt stehen — Norm, kein Ausfüll-Hinweis).* Dieses
> Artefakt friert ein; was es zitiert, bewegt sich weiter. Deshalb: **Kennung,
> nicht Adresse** — `slice-<Kennung>` statt seines Lifecycle-Pfads, `make <target>`
> statt eines Links auf die Sensor-Datei, eine Baseline-Stelle als
> `v6.13.0` · `regelwerk/<datei>.md` §<Abschnitt> statt als Link.

**Welle:** welle-16-shapes-sollform
**Abschluss:** 2026-10-06
**Verantwortlich:** Claude (Planner); Abnahme beim Maintainer.

## Was wurde geliefert?

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3.

- **Der Change Request „Positivliste von Anweisungen je Datei (Sollform)" ist vollständig
  ausgeliefert** — mit `v0.21.0`, Digest
  `sha256:65165a387c4d974f66ae687bda3d9dcd49742022d5c5cf4f938a27a00d144e58`.
- slice-209: Vertrag — [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012)
  (Lastenheft 0.28.0), [ADR-0041](../../adr/0041-shapes-sollform-je-datei.md) `Accepted` mit
  vier Schärfungen gegenüber den abgenommenen Entscheiden.
- slice-210: Dialekt `kotlin` und `mode: allow-statements` (Befund `shape-unlisted`) — Lexer,
  Zerleger, Dateisuche, Validierung auch im no-scan-Pfad.
- slice-211: `mode: exact` (`shape-differs`) und `unused: fail` (`shape-unused`).
- Spezifikation 0.33.0 → 0.35.0 in drei Präzisierungen aus den Reviews (Backtick-Vorlage,
  Duplikate, Symlinks, Sollform-Identität). Benutzerhandbuch 1.43.
- [MR-024](../../../../harness/conventions.md#mr-024) aufgelöst: Sein Auflösungs-Trigger war
  dieses Release.

## Was hat funktioniert?

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3.

- **Spec-first mit Review vor dem Code.** Drei Funde, mit denen sich eine Abhängigkeit als
  Kommentar hätte verstecken lassen, fielen im Vertrag an — dort kostet ein Fund einen Satz.
- **Mutations-Gegenprobe für jeden Fix.** Sieben Fixes wurden bewusst gebrochen und gingen aus
  dem richtigen Grund rot (Dollar-Präfix, Regex-Verankerung, Backtick-Vorlage, `./`-Normalisierung,
  Hardlink, Erst-Treffer-Markierung, Symlink-Präfix). So trennte die Welle „Test grün" von „Test
  belegt die Zusage".
- **Vermutungen als Sonden übergeben.** Der Reviewer bekam meine Annahmen als Prüfaufträge, nicht
  als Feststellungen; die eine falsche („die `files`-Suche folgt keinem Symlink") fiel so auf, bevor
  das Release sie auslieferte.
- **Plan-Änderung vor dem Code**, dreimal: wenn ein Review eine Vertragslücke fand, ging sie erst
  in §1 des Slice-Plans, dann in die Spezifikation, dann in den Code.

## Was ging anders als geplant?

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3.

- **Elf Review-Läufe statt drei.** Kein Slice kam mit einem Lauf aus; die Befunde trafen zwei
  Klassen wiederholt — eine aus dem Gedächtnis aufgezählte Lexik der Zielsprache (vier Fälle über
  slice-209 und slice-210) und eine Hermetik-Zusage, die am Pfad-Text statt am Dateisystem geprüft
  wurde (drei Stellen in slice-211, die letzte in Code aus slice-210). Beide stehen im Register.
- **Jeder Fix schloss nur die gemeldete Stelle.** Die Symlink-Lücke zeigte sich nacheinander an der
  Sollform-Datei, an einem Verzeichnis in ihrem Pfad und am Präfix der `files`-Suche; Spezifikation
  0.35.0 entstand dadurch in drei Schritten.
- **Ein Glossar-Satz zählte „zehn Regeln"** — viertes Auftreten einer schon verkörperten Klasse;
  beantwortet mit einem Zeiger statt einer Zahl und einer Begründung, warum kein Sensor möglich ist.

## Steering-Loop-Einträge

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3 — hier stehen nur Beobachtungen, die im Register 3× erreicht
haben.

- **Keiner.** Kein offener Register-Eintrag erreichte in dieser Welle 3×. Der vierte Beleg zu
  `BEO-HARNESS/aggregat-aufzaehlung-hinkt-dem-makefile-hinterher` traf einen bereits
  **verkörperten** Eintrag; sein Ausgang bleibt, die Begründung ohne Sensor steht in seinem
  `state.md` (slice-210).

## Beobachtungs-Register (Zeiger)

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Das Beobachtungs-Register — der Zähler wird nicht hier gepflegt.

Der Zähler steht im [Beobachtungs-Register](../observations/README.md). Neu oder fortgeschrieben
in dieser Welle: `BEO-GATE/marke-enger-als-bestandsform` (1×) ·
`BEO-SPEC/lexik-vertrag-ohne-sprach-gegenprobe` (2×) ·
`BEO-SPEC/vertrag-auf-einen-konsumenten-belegt` (1×) ·
`BEO-ADAPT/hermetik-lexikalisch-statt-am-dateisystem` (1×) ·
`BEO-GATE/testbeschreibung-weiter-als-assertion` (2×) ·
`BEO-SPEC/ausgabe-einzeilig-nicht-zugesichert` (2×).

**Bestand in `open/`, Lese-Schritt:** slice-013 und slice-045 haben diese Closure unverändert
überstanden. Gruppierung: keine gemeinsame Ursache. Je **bestätigt** — beide sind
trigger-gebunden (Richtungs-Namen eines Konsumenten bzw. eine der drei Bedingungen aus slice-045
§0), und keiner der Trigger ist gefeuert; ihre Zeilen in *Nächste Wellen* der Roadmap tragen sie.

## Folge-Slices

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3 — derivativ.

- Keine. Der generische Dialekt steht als gated Zeile unter *Nächste Wellen* der Roadmap, ohne
  Slice-Datei — sein Trigger (ein zweiter Konsument mit Nicht-Kotlin-Manifest) ist nicht gefeuert.

## Verifikation

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 1 — Replay-Ersatz nach
[MR-028](../../../../harness/conventions.md#mr-028).

- Alle drei Slices in `done/`; `make ci` Exit 0 und `make verify` Exit 0 auf dem Stand nach dem
  Re-Pin (lokal, Ausgabe in Dateien, Exit-Codes getrennt gelesen); `make preflight` Exit 0 vor dem
  Push (57 Commits Range).
- CI auf `main` grün: Lauf 37503705022 (Commit `c02807c`). Release-Pipeline grün: Lauf
  37503905727 (Tag `v0.21.0`) — SemVer, `make ci`, OCI-Label, GHCR, Docker-Hub-Spiegel,
  GitHub-Release.
- Digest gegen die Registry gegengeprüft: `docker buildx imagetools inspect` nennt
  `sha256:65165a387c4d974f66ae687bda3d9dcd49742022d5c5cf4f938a27a00d144e58`, OCI-Label
  `org.opencontainers.image.version` = `0.21.0`. Re-Pin danach mit `make gates` Exit 0
  (`gate-consistency`: Pins konsistent).
- `make doc-immutable RANGE=v0.20.0..HEAD`: 0 Befunde — keine deklarierten Ausnahmen mehr nötig.
- **Gegenprobe am Leitfall gegen das veröffentlichte Image:** eine `shapes`-Konfiguration in
  Adopter-Form für `domain/build.gradle.kts` lief grün (Exit 0); nach Einfügen von
  `implementation("com.evil:lib:1.0")` in den `dependencies`-Block rot (Exit 1) mit dem
  **gesehenen** Befund
  `domain/build.gradle.kts:9: shape-unlisted: dependencies{testImplementation(kotlin("test"));implementation("com.evil:lib:1.0")}`
  und — weil `unused: fail` gesetzt war — `shape-unused` für den nun nicht mehr passenden Block.
- Coverage gesamt 96,4 % (Schwelle 90 %). Carveouts: keine.
- **Trigger-Audit, vier Klassen:** Carveout 0 offen (Verzeichnis führt nur die README) ·
  bootstrap-aware Gate 0 (keines im Bestand) · ADR 0 fällig ([ADR-0021](../../adr/0021-commits-modul-trace-check.md) hängt am d-check-Pin-Bump,
  der nicht stattfand; die Trigger von [ADR-0041](../../adr/0041-shapes-sollform-je-datei.md) sind nicht eingetreten) · Hard Rule 0 (keine mit
  Auflösungs-Trigger) — dazu die MR-Einträge: [MR-024](../../../../harness/conventions.md#mr-024) fällig und aufgelöst, die übrigen sieben
  0 offen.
- **Drei Paarungen:** Anker — kein Steering-Loop-Eintrag mit `liegt in` in dieser Welle, nichts
  zu paaren · Folge-Slice — keiner genannt · Register — alle sechs genannten Pfade existieren als
  Verzeichnis mit nicht leerem `evidence/` (`make verify-observations` im `verify`-Lauf grün).
