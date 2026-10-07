# Review-Report: slice-217 — 2026-10-07

**Review-Art:** Code-Review — geprüft wird der Diff gegen Slice-Plan, ADR-0043, SPEC-DIST-001 und
die Hard Rules (Modul 10). Unabhängiger Lauf in frischem Kontext, nicht der Autor; keine
DoD-Verifikation (Verifier, Modul 11).

**Gegenstand:** Commit-Range `9534afe..c5533df` — Inhalt in `c5533df` (Dockerfile, Makefile,
`tools/image-multiarch.sh`, `tools/multiarch-check.sh`, `tools/image-test.sh`, Guard-Liste,
Gate-Index, `.gitignore`).

**Skill:** `.harness/skills/reviewer.md` @ `e6917e3` · <!-- d-check:ignore -->
**Modell:** claude-opus-5-5 · **Datum:** 2026-10-07

> **Zitier-Form** *(Norm, kein Ausfüll-Hinweis)*. Dieser Report friert ein; was er zitiert, bewegt
> sich weiter. Deshalb: **Kennung, nicht Adresse** — `slice-<Kennung>` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine Baseline-Stelle als
> **Tag + Pfad in Inline-Code** statt als Link. Ein `pfad`-Feld auf den **geprüften Gegenstand**
> ist davon nicht betroffen — es zitiert den Stand des Laufs und darf ihn festhalten.

**Eingangs-Kontext:**

- slice-217 (§1–§4, §6, §8), slice-218 (§1–§3), welle-18 (Welle-Plan)
- ADR-0043 (`Accepted`), SPEC-DIST-001, AC-FA-DIST-001, AC-QA-03
- `AGENTS.md` §3 (Hard Rules), §6 Schritt 4 (Plan-Änderung vor Code)
- frühere Findings am Bereich: Review slice-216 (F-4: Fitness Function nennt `make image-test`)
- Bestand: `.github/workflows/release.yml`, `.claude/hooks/pretooluse-command-guard.sh` (Selbsttest)

**Eigene Nachmessung** (Mess-Regel 1, Geltungsbereich): dieser Host `x86_64`, Docker 29.8.2, buildx
0.37.1, Builder `a-check-multiarch-cec9f139f45e` (BuildKit 0.33.1), **keine** arm64-Emulation;
eigene Registry `registry:2` auf Port 5187 (danach entfernt), Kopie per `skopeo copy --all`.
Mutationen am OCI-Archiv per Skript im Scratchpad (Index/Manifest/Config neu gehasht, `index.json`
nachgezogen). Gemessen wird `tools/multiarch-check.sh` gegen **BuildKit-förmige** und
**handgebaute** Archive — nicht gegen andere Exporter (`docker save`, `skopeo`-Archive).

| Probe | Ergebnis (Meldung) |
|---|---|
| `make image-multiarch VERSION=9.9.9-rev` | Exit 0, Index `46d6b58a…`, beide Plattformen, ELF 62/183 |
| Prüfung mit falscher Version | rot — „linux/amd64: Versions-Label '9.9.9-rev', erwartet '1.2.3'" |
| Index nur amd64 | rot — „Plattformen 'linux/amd64'" |
| Index amd64 doppelt + arm64 | rot — „Plattformen 'linux/amd64 linux/amd64 linux/arm64'" |
| Attestierungs-Eintrag ergänzt | rot — „der Index führt 'unknown/unknown'" |
| arm64-Eintrag zeigt auf das amd64-Manifest | rot — „linux/arm64: Config sagt linux/amd64" |
| arm64-Manifest mit amd64-Binary-Layer | rot — „linux/arm64: Binary hat ELF e_machine '62', erwartet '183'" |
| Feldreihenfolge `digest` vor `mediaType` | rot — „Plattformen 'linux/amd64'" (Split greift nicht, fail-closed) |
| Index eingerückt (Whitespace) | grün — korrekt, `blob()` entfernt Whitespace |
| Layer unkomprimiert / Binary nicht im letzten Layer | rot, Exit 2 — Meldung von `tar`, nicht vom Prüfer |
| **arm64-Eintrag mit Array-Feld (`urls` bzw. `platform.os.features`) + dritter Eintrag `linux/s390x`** | **grün** — „ok — … mit genau: linux/amd64 linux/arm64" (siehe F-2) |
| Dockerfile ohne `GOOS`/`GOARCH` (Kopie im Scratchpad) | rot — „linux/arm64: Binary hat ELF e_machine '62', erwartet '183'" |
| `make image-test-ref` mit Index-Digest aus der Registry | Exit 0 |
| `make image-test-ref` mit arm64-Manifest-Digest | rot — „ELF e_machine 183, Host x86_64 erwartet 62" |
| `make image-test-ref` ohne `IMAGE_REF` | Exit 2 — „IMAGE_REF=<tag\|digest> fehlt" |
| Guard-Liste ohne die zwei neuen Targets (Kopie, Selbsttest aus der Repo-Wurzel) | rot — „Pruef-Target 'image-multiarch' fehlt in der GATES-Liste" (und `image-test-ref`) |
| `make guard-selftest`, `make doc-targets`, `make gates`, `make doc-check` | Exit 0 (Ausgabe in Datei, Exit-Code getrennt) |
| `make doc-structure` (Teil von `make verify`) | rot, Exit 2 — ein Befund `section-cell-oversized` an `harness/README.md:93` (F-5) |
| `docker buildx imagetools inspect moby/buildkit:v0.33.1` | Index-Digest `cec9f139…` = Pin im Makefile |

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | Gate-Index und Makefile-Kommentar sagen im Indikativ, `make image-test-ref` sei der Weg, auf dem die Release-Pipeline den einmal gebauten Digest je Plattform testet. Auf diesem Stand ruft `.github/workflows/release.yml` nur `make ci` (Bau + Test des lokalen `:dev`-Bilds); weder `image-test-ref` noch ein Digest-Test kommen dort vor — die Verdrahtung ist Gegenstand von slice-218. Adversarisch geprüft: `grep image-test-ref\|IMAGE_REF` in `.github/` ohne Treffer. | nachweislich falsche Tatsachenbehauptung (Skill HIGH); ADR-0043 Punkt 5 | `harness/README.md:92`; `Makefile:227-228` | ja — `grep -rn image-test-ref .github/` | Künftiger Zustand im Indikativ als bestehender behauptet |
| F-2 | MEDIUM | `multiarch-check.sh` schneidet die Manifest-Liste am **ersten** `]` ab. Trägt ein Eintrag ein Array (`urls`, `platform.os.features`), fallen alle folgenden Einträge aus der Prüfung — gemessen: ein Index mit `linux/amd64`, `linux/arm64` und `linux/s390x` meldet „ok — mit genau: linux/amd64 linux/arm64". Die Schlussprüfung (Zeile 87) zählt dieselbe abgeschnittene Liste und ist damit keine unabhängige zweite Zählung. Kopfkommentar (1) sagt „keine weitere Plattform", Zeile 11-13 „fail-closed"; die tragenden Annahmen (kein Array in einem Eintrag, `mediaType` als erstes Feld) stehen nicht im Kommentar, genannt ist nur „kompakt", das `blob()` ohnehin herstellt. BuildKit erzeugt für `linux/*` heute kein solches Feld — der Pfad ist erreichbar, aber nicht vom aktuellen Erzeuger. | SPEC-DIST-001 („nur diese zwei Plattform-Bilder"); Mess-Regel „Testkommentar sagt nicht mehr zu, als seine Assertion prüft" | `tools/multiarch-check.sh:4-5`, `:11-13`, `:44`, `:86-87` | ja — Mutations-Probe oben (Array-Feld + dritter Eintrag) | Prüfer-Zusage weiter als seine Prüfung |
| F-3 | MEDIUM | §4 des Slice-Plans benennt vorab als Rückführung `in-progress → next`: „das Make-Target braucht eine neue Werkzeugkette im Repo (Builder-Instanz, Emulation) — dann teilen". Geliefert ist ein eigener `docker-container`-Builder mit eigenem, digest-gepinntem BuildKit-Image (`BUILDKIT_IMAGE`, `buildx create`) — eine Builder-Instanz. Weder der Plan noch die Commit-Message sagen, ob die Bedingung als eingetreten gilt oder warum nicht (die Klammer lässt offen, ob *eines* der beiden genügt). Eine Plan-Änderung vor dem Code (`AGENTS.md` §6 Schritt 4) liegt in der Range nicht vor. | slice-217 §4; `AGENTS.md` §6 Schritt 4 | `docs/plan/planning/in-progress/slice-217-multi-arch-build.md:67-68`; `tools/image-multiarch.sh:7-11`, `:28-31`; `Makefile:25` | nein — Urteil über die Lesart der Bedingung | Vorab benannte Rückführungs-Bedingung eingetreten, nicht behandelt |
| F-4 | MEDIUM | ADR-0043 §Fitness Function bindet „Image-Test gegen eine übergebene Bild-Referenz, je Plattform nativ" an `make image-test`. Geliefert ist ein **neues** Target `make image-test-ref`; `make image-test` bleibt an `build` und `${IMAGE}:dev` gebunden. Die Abweichung vom Anker der `Accepted`-ADR ist weder im Slice-Plan noch in der Commit-Message benannt (Review slice-216 F-4 hatte genau diese Zeile adressiert). | ADR-0043 §Fitness Function; `v6.13.0` · `regelwerk/modul-08-agentenrollen.md` §Rollen-Regeln („nicht stillschweigend einer ADR widersprechen") | `docs/plan/adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md:117`; `Makefile:223`, `:229` | ja — Fitness-Tabelle gegen `make help` | Fitness-Function-Anker und geliefertes Target weichen unbenannt ab |
| F-5 | MEDIUM | Die neue Gate-Index-Zeile `make image-multiarch` überschreitet die Zellengrenze der Vertrags-Spalte (279 Zeichen, Grenze 250); `make doc-structure` — Teil von `make verify`, nicht von `make gates` — ist auf diesem Stand rot: „section-cell-oversized … was darueber hinausgeht, gehoert nach harness/sensors/<target>.md". Die Übergabe meldet `make gates` und `make ci` grün, `make verify` nicht. | Konvention Gate-Index (Zellengrenze, `.d-check.yml` Modul `structure`) | `harness/README.md:93` | ja — `make doc-structure` (Exit 2, ein Befund) | Gate-Index-Zelle über der Zellengrenze |
| F-6 | LOW | `tools/image-test.sh` liest `IMAGE_REF` aus der Umgebung auch unter `make image-test`. GNU make exportiert Kommandozeilen-Variablen in die Rezept-Umgebung (gemessen mit einem Mini-Makefile: `make x IMAGE_REF=foo` → `env=foo`). `make image-test IMAGE_REF=x` bzw. `make ci IMAGE_REF=x` oder ein exportiertes `IMAGE_REF` baut `:dev` und testet dann `x`; der Gate-Index sagt für `image-test` „gegen das gebaute Image". Die Ausgabe nennt das geprüfte Bild (Zeile 47), die Zeile „OK" nicht. | Gate-Index `make image-test`; ADR-0043 („getestet = veröffentlicht") | `tools/image-test.sh:26`; `Makefile:223-224` | ja — `make image-test IMAGE_REF=<anderes Bild>` und die Bild-Zeile der Ausgabe lesen | Umgebungsvariable koppelt zwei Targets still |
| F-7 | LOW | §3 des Slice-Plans nennt vier Dateien; geliefert sind zusätzlich `tools/image-multiarch.sh`, `tools/multiarch-check.sh`, die Guard-Liste und `.gitignore`. Die Guard-Zeile ist Folge des Selbsttests, die zwei Skripte sind eigene Liefer-Substanz (131 Zeilen), die der Plan-Tisch nicht führt. | slice-217 §3 | `docs/plan/planning/in-progress/slice-217-multi-arch-build.md:55-59` | ja — `git show --stat c5533df` gegen §3 | Plan-Tabelle führt die gelieferten Träger unvollständig |
| F-8 | LOW | Der Kopf von `tools/image-test.sh` sagt weiter „gegen das lokal gebaute Runtime-Image"; seit diesem Diff prüft das Skript wahlweise ein übergebenes Bild (Zeile 22-23 sagt es richtig). Der Kopfsatz beschreibt nicht mehr vollständig, was da ist. | `AGENTS.md` §3.7 | `tools/image-test.sh:2-3` | nein — Lesart | Kopfkommentar nach Erweiterung nicht nachgezogen |
| F-9 | INFO | Die Fehlerpfade „Layer nicht gzip" und „Binary nicht im letzten Layer" enden rot mit Exit 2 und der Meldung von `tar`, nicht mit einem `multiarch-check: FAIL`. Fail-closed, aber die Diagnose nennt die verletzte Annahme nicht. | — | `tools/multiarch-check.sh:70-75` | ja — Probe oben | — (Hinweis) |
| F-10 | INFO | Der Pin `BUILDKIT_IMAGE` trägt Tag **und** Digest; `make version-coherence` prüft nur `uses:`-SHAs und Makefile/Dockerfile-Doppel, nicht dieses Paar. Heute stimmt es (nachgemessen). Klasse von `BEO-GATE/versionsangabe-neben-digest-ungeprueft`; Adressat ist die Closure. Ebenso bleibt der Builder-Container nach dem Lauf stehen — kein Aufräumen, keine Aussage dazu. | AC-QA-03 | `Makefile:25`; `tools/image-multiarch.sh:28-31` | ja — `docker buildx imagetools inspect` des Tags | — (Hinweis) |
| F-11 | INFO | ADR-0043 Punkt 5 und SPEC-DIST-001 verlangen den Image-Test „einschließlich des Versions-Labels"; `tools/image-test.sh` prüft das Label nicht. slice-218 §2 führt „Versions-Label je Plattform" in `release.yml` — ob der Träger der Image-Test oder ein Pipeline-Schritt wird, entscheidet dort der Implementer. | ADR-0043 Punkt 5 | `tools/image-test.sh` | nein | — (Hinweis an slice-218) |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Dockerfile — Wirkung von `--platform=$BUILDPLATFORM` auf `compile`/`lint`/`test`/`coverage` | geprüft: bei einplattformigem `docker build` ist `BUILDPLATFORM` = `TARGETPLATFORM`; `lint` kopiert nur den plattformneutralen Modul-Cache aus `deps`; `# syntax=docker/dockerfile:1.7` setzt BuildKit ohnehin voraus. `make gates` Exit 0. Ohne Befund |
| Dockerfile — Laufzeit-Stufe ohne `RUN`, `GOOS`/`GOARCH` aus `TARGETOS`/`TARGETARCH` (ADR-0043 Punkt 2) | geprüft; Gegenprobe ohne `GOOS`/`GOARCH` rot mit ELF-Meldung. Ohne Befund |
| `--provenance=false --sbom=false` (ADR-0043 Punkt 3) | geprüft; Attestierungs-Probe rot. Ohne Befund |
| `multiarch-check.sh` — Config-Architektur, Versions-Label, ELF je Plattform, Dubletten, Fehlzeiger | geprüft per Mutation (Tabelle oben), jeweils rot aus dem richtigen Grund; Befund nur F-2 |
| `image-test.sh` — Zusage „Plattform des Hosts geprüft" | geprüft: ELF-Prüfung vor dem Nativ-Lauf, rot gegen das arm64-Manifest auf `x86_64`, grün gegen den Index (Docker zieht die Host-Plattform). Ohne Befund |
| Guard-Liste und `make guard-selftest` | geprüft; Mutation (Liste ohne die zwei Targets) rot mit Namen. Ohne Befund |
| Gate-Index — beide Zeilen existieren als Regel (`make doc-targets` Exit 0), Bindung ADR-0043/slice-217, „Nicht hier"-Absatz nachgezogen | geprüft; Befunde F-1 (Wahrheit der Release-Pipeline-Aussage) und F-5 (Zellengrenze) |
| §1 Ausschlüsse — Release-Pipeline, Push, Spiegel, Gate-Stages für arm64 | geprüft: `release.yml`, `packaging/`, Gate-Stages im Diff unberührt. Ohne Befund |
| Kommentare `AGENTS.md` §3.7 — Dockerfile, Makefile, beide neuen Skripte | geprüft: die Konjunktive (`od -N2` am Pipe-Ende, `head -1`) zeigen nach vorn auf einen Fallstrick für den Ändernden (Klasse Abgrenzung), nicht zurück auf eine verworfene Alternative. Befund nur F-8 |
| Hard Rules §3.1–§3.6 (Docker/make-only, keine Suppression, kein Move, keine ADR überschrieben, kein Gate gelockert) | geprüft, ohne Befund |
| `BUILDKIT_IMAGE` digest-gepinnt (AC-QA-03), Pin = Index-Digest des Tags | nachgemessen, ohne Befund |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 4 |
| LOW | 3 |
| INFO | 3 |

**Finding-Klassen dieses Laufs:** Künftiger Zustand im Indikativ als bestehender behauptet ·
Prüfer-Zusage weiter als seine Prüfung · Vorab benannte Rückführungs-Bedingung eingetreten, nicht
behandelt · Fitness-Function-Anker und geliefertes Target weichen unbenannt ab · Gate-Index-Zelle über der Zellengrenze · Umgebungsvariable
koppelt zwei Targets still · Plan-Tabelle führt die gelieferten Träger unvollständig ·
Kopfkommentar nach Erweiterung nicht nachgezogen

## Verdikt

**Abnahme-blockierend:** ja — F-1 (HIGH) und F-2 bis F-5 (MEDIUM) sind vor der Closure zu klären.
F-3 ist eine Plan-Frage (Lesart der Rückführungs-Bedingung) und geht an den Planner; F-4 betrifft
den Anker einer `Accepted`-ADR — widerspricht der Implementer, läuft der Konflikt-Pfad über den
Architect. Die übrigen Behauptungen der Übergabe (Gegenproben) sind nachgefahren und bestätigt.

**Übergabe:** Findings an den Implementer; die Finding-Klassen zusätzlich in die Closure §7 von
slice-217. Dieser Report ist Lauf-Beleg und ersetzt keine Verifikation gegen die DoD.
