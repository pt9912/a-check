# Review-Report: slice-216 — 2026-10-07

**Review-Art:** Plan/Design — geprüft wird der Vertrag (Messung, Lastenheft-CR, Spezifikation,
ADR) gegen Slice-Plan, Welle-Plan, aktive ADRs und Hard Rules (Modul 10 §Drei Review-Arten).
Unabhängiger Lauf in frischem Kontext, nicht der Autor; keine DoD-Verifikation.

**Gegenstand:** Commit-Range `683b29c..354a432` — `ef64fb9` (Messung §1b), `b166a0f`
(Lastenheft 0.30.0, Spezifikation 0.39.0), `354a432` (ADR-0043 `Proposed`, ADR-Index).

**Skill:** `.harness/skills/reviewer.md` @ `e6917e3` · <!-- d-check:ignore -->
**Modell:** claude-opus-5-5 · **Datum:** 2026-10-07

> **Zitier-Form** *(Norm, kein Ausfüll-Hinweis)*. Dieser Report friert ein; was er zitiert, bewegt
> sich weiter. Deshalb: **Kennung, nicht Adresse** — `slice-<Kennung>` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine Baseline-Stelle als
> **Tag + Pfad in Inline-Code** statt als Link. Ein `pfad`-Feld auf den **geprüften Gegenstand**
> ist davon nicht betroffen — es zitiert den Stand des Laufs und darf ihn festhalten.

**Eingangs-Kontext:**

- slice-216 (§1, §1b Messung M1–M7, §3, §6), welle-18 (Welle-Plan), slice-217, slice-218
- ADR-0043 (`Proposed`), ADR-0039, ADR-0042 (Muster der Teil-Ablösung), ADR-0037, ADR-0004,
  ADR-0007
- AC-FA-DIST-001, AC-FA-DIST-002, AC-QA-01, AC-QA-03; SPEC-DIST-001; ARC-008, ARC-009
- `AGENTS.md` §3 (Hard Rules), §6 Schritt 7 (CHANGELOG)
- Bestand: `.github/workflows/release.yml`, `tools/image-test.sh`, `tools/image-scan.sh`,
  `Makefile` (`image-test`), `Dockerfile`, `.dockerignore`, `packaging/dockerhub/`,
  `docs/user/releasing.md`, `CHANGELOG.md`

**Eigene Nachmessung** (Mess-Regel 1, Geltungsbereich): dieser Host, Docker 29.8.2, buildx 0.37.1,
Builder `ci` (`docker-container`, BuildKit 0.33.1), klassischer Bildspeicher `overlay2`, **keine**
arm64-Emulation; eigene Dockerfile-Variante im Scratchpad (Cross-Compile wie §1b beschrieben,
Laufzeit-Stufe mit nur dem Versions-Label), Ausgabe `type=oci,rewrite-timestamp=true`,
`--provenance=false --sbom=false`, `SOURCE_DATE_EPOCH` = Commit-Zeit von `354a432`
(`1791363202`). Gemessen:

| Lauf | Cache | Epoch | Index-Digest | mtime `/a-check` |
|---|---|---|---|---|
| a | warm (Cache-Einträge eines **früheren** Baus mit anderer Epoch) | Commit | `75095c0d…` | `1791362882` (**vor** der Commit-Zeit) |
| b | `--no-cache` | Commit | `fe85ce3d…` | `1791363202` |
| c | warm (nach b) | Commit | `fe85ce3d…` | `1791363202` |
| d | warm | Commit − 1000 s | `a85b9d3d…` | `1791362202` |
| e | warm (nach d) | Commit | `fe85ce3d…` | `1791363202` |
| f | `--no-cache` | Commit | `fe85ce3d…` | `1791363202` |

Das Binary ist in allen Läufen bitgleich (amd64 `e47f8a3a…` — derselbe Wert wie §1b M4). Die
zwölf Basis-Layer aus `distroless` tragen in der `history` `1970-01-01T00:00:00Z`, nur die vier
eigenen Schritte die Commit-Zeit. Kopie per `skopeo copy --all` Registry → Registry (zwei lokale
`registry:2`): Index-Digest unverändert, beide Plattformen — **bestätigt M5** in diesem
Geltungsbereich. `make doc-check doc-structure gate-consistency`: Exit 0, Ausgabe in Datei.

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | Die Spezifikation sagt zu, alle Zeitstempel im Bild stünden auf der Commit-Zeit und derselbe Commit ergebe denselben Index-Digest; ADR-0043 Punkt 4 und §1b M4 („mit und ohne Cache identisch") tragen dasselbe. Gegenmessung: derselbe Commit mit denselben Flags ergab `75095c0d…` (warmer Cache aus einem früheren Bau, Datei-mtime vor der Commit-Zeit bleibt stehen — `rewrite-timestamp` klemmt nur spätere Zeiten) gegen `fe85ce3d…` kalt; die Basis-Layer tragen 1970, nicht die Commit-Zeit. | nachweislich falsche Tatsachenbehauptung (Skill HIGH); SPEC-DIST-001; ADR-0043 Punkt 4 | `spec/spezifikation.md:704-705`; `docs/plan/adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md:63-65`, `:97`; `docs/plan/planning/in-progress/slice-216-multi-arch-spec-first.md:69` | ja — zwei `buildx`-Läufe derselben Revision, der erste auf einem Cache aus einem Bau mit anderer `SOURCE_DATE_EPOCH`; Index-Digests vergleichen (Tabelle oben) | Messung mit zu engem Geltungsbereich als Vertragszusage übernommen |
| F-2 | MEDIUM | Der Punkt „Reproduzierbar" in SPEC-DIST-001 hat keinen Träger im Lastenheft: weder AC-FA-DIST-001 noch AC-QA-03 (dort: Digest-Pin und Platzhalter im Fragment) sagen einen bit-reproduzierbaren Bau zu; die Spezifikation führt damit eine neue Zusage ein, statt eine bestehende zu präzisieren. | Source Precedence (Spezifikation präzisiert das Lastenheft); Prüffrage 4 | `spec/spezifikation.md:704-705`; `spec/lastenheft.md:761-766` | nein — Urteil über Rang-Zuordnung | Spezifikation sagt mehr zu als das Lastenheft |
| F-3 | MEDIUM | AC-FA-DIST-001 Boundary (Plattformen) sagt byte-identische Ausgabe und gleichen Exit-Code auf `linux/amd64` und `linux/arm64` für den veröffentlichten Digest zu. Weder ADR-0043 (Entscheidung, Fitness Function) noch slice-217/slice-218 sehen einen Vergleich **zwischen** den Plattformen vor; `tools/image-test.sh` vergleicht nur nativ gegen Container innerhalb einer Plattform. Einziger Beleg ist die einmalige Gegenprobe der welle-18-Closure, nicht jedes Release. | AC-FA-DIST-001; ADR-0043 §Fitness Function | `spec/lastenheft.md:703`; `docs/plan/adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md:106-112`; `tools/image-test.sh:40-137` | ja — `grep` nach einem plattformübergreifenden Vergleich in Plan und Skripten: kein Treffer | Zusage weiter als ihre Durchsetzung |
| F-4 | MEDIUM | Zwei Folgepflichten aus ADR-0043 haben keinen Träger in §3 von slice-217 oder slice-218: (a) der Image-Test gegen eine Bild-Referenz von außen — heute hängt `make image-test` an `build` und `tools/image-test.sh` fest an `${IMAGE}:dev`, die Fitness Function nennt aber `make image-test` als Test des übergebenen Digests; (b) die Hub-Seite — `packaging/dockerhub/overview.md` sagt „one copied from GHCR will not resolve here" und `packaging/dockerhub/README.md` nennt den Config-Digest als Gleichheits-Größe, beides widerspricht AC-FA-DIST-002 0.30.0 nach dem ersten Multi-Arch-Release. | ADR-0043 §Konsequenzen (Folgepflicht); AC-FA-DIST-002 („Die Hub-Seite sagt das") | `docs/plan/adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md:102-104`, `:111`; `Makefile:217-218`; `tools/image-test.sh:22`; `packaging/dockerhub/overview.md:36-43`; `packaging/dockerhub/README.md:83-84` | ja — §3 beider Folge-Slices gegen die Folgepflicht-Liste abgleichen | Folgepflicht ohne Träger-Slice |
| F-5 | MEDIUM | Der Slice ändert Lastenheft und Spezifikation, trägt aber keinen Eintrag in `CHANGELOG.md` `[Unreleased]`; §3 des Slice-Plans nennt die Datei nicht. Der vorige Spec-first-Slice derselben Form (slice-212, `f639a2f`) trug ihn. | `AGENTS.md` §6 Schritt 7 | `CHANGELOG.md:7`; `docs/plan/planning/in-progress/slice-216-multi-arch-spec-first.md` §3 | nein — kein Gate deckt es (`AGENTS.md` §6 sagt das selbst) | Vertragsänderung ohne CHANGELOG-Zeile |
| F-6 | MEDIUM | Das Contra von Alternative D („Laden … schreibt eine Plattform neu — getestet würde eine Kopie, nicht der Index") ist unbelegt und trifft auf die gewählte Option E ebenso zu: auch der Test-Job in E zieht per Digest eine Plattform in den lokalen Bildspeicher. §1b M5 belegt zudem, dass ein OCI-Archiv per `skopeo copy --all` mit unverändertem Index-Digest hochgeladen wird — die Identität wäre auch in D per Digest prüfbar. | Skill MEDIUM (unbelegte Tatsachenbehauptung); `modul-04` Ziel-Form ADR (Alternativen) | `docs/plan/adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md:89` | nein — die Abwägung ist Urteil; die Archiv-Hälfte wäre messbar, ist es nicht | Alternative mit unbelegtem Contra verworfen |
| F-7 | LOW | Die Teil-Ablösung erklärt ADR-0039 Punkte 2, 3 und den übrigen Punkt 4 für weiter gültig; Punkt 2 nennt aber wörtlich den Config-Digest als Abbruchgrund, und ADR-0039 §Konsequenzen („Hub-Nutzer bekommt keinen fertigen Pin") und §Fitness Function (Vergleich der Config-Digests) sind von der Ablösung ebenso betroffen, ohne genannt zu sein. | ADR-0039; Muster ADR-0042 | `docs/plan/adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md:12-15` | nein — Lesart | Teil-Ablösung zählt die betroffenen Stellen unvollständig |
| F-8 | LOW | Die Spezifikation bindet den Test an „einen Rechner **dieser** Plattform"; der erste Re-Evaluierungs-Trigger der ADR nennt als Rückfall Option B (Emulation). Der Rückfall bräuchte damit eine Spezifikations-Änderung, was der Trigger nicht sagt. | SPEC-DIST-001; ADR-0043 §Re-Evaluierungs-Trigger | `spec/spezifikation.md:706-708`; `docs/plan/adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md:116` | nein — Lesart | Re-Evaluierungs-Pfad ohne den Spec-Rang |
| F-9 | LOW | Die CVE-Prüfung des veröffentlichten Bilds (ADR-0037, `tools/image-scan.sh`, Refs per Tag) prüft bei einem Index die Plattform des scannenden Rechners; ADR-0043 §Konsequenzen nennt nicht, ob das arm64-Bild gescannt wird. | ADR-0037; ADR-0043 §Konsequenzen | `tools/image-scan.sh:63`; `docs/plan/adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md:92-104` | ja — `make image-scan` gegen den ersten Multi-Arch-Index, Plattform im Trivy-Bericht lesen | Folge an einem bestehenden Sensor nicht benannt |
| F-10 | LOW | AC-FA-DIST-001 Negative (Plattform-Test) sagt „wird **kein** Versions-Tag gesetzt"; der Git-Tag `v*` existiert zu diesem Zeitpunkt bereits (er löst die Pipeline aus). Gemeint ist der Bild-Tag, der Satz sagt es nicht. | AC-FA-DIST-001 | `spec/lastenheft.md:704` | nein — Wording | Mehrdeutiger Begriff in einem Akzeptanzkriterium |
| F-11 | LOW | Das Out-of-Scope von AC-FA-DIST-002 schließt weiterhin aus, „Pin-Stellen dieses Repos auf den Hub-Digest umzustellen"; nach der neuen Beschreibung sind GHCR- und Hub-Digest derselbe Wert, der Ausschluss hat keinen Gegenstand mehr. | AC-FA-DIST-002 | `spec/lastenheft.md:744` | nein — Wording | Out-of-Scope-Punkt nach Vertragsänderung gegenstandslos |
| F-12 | INFO | §1b M5 nennt als Geltungsbereich die Registries, nicht den Bildspeicher. Das Ergebnis „`docker tag` + `docker push` liefert ein Einzel-Manifest" hängt am klassischen Speicher (`overlay2` auf diesem Host); die Folgerung (c) gilt trotzdem, weil der Test-/Spiegel-Rechner per Digest nur eine Plattform lokal hält. | Mess-Regel 1 (Geltungsbereich) | `docs/plan/planning/in-progress/slice-216-multi-arch-spec-first.md:70` | ja — derselbe Lauf mit containerd-Bildspeicher | Geltungsbereich nennt nicht alle tragenden Bedingungen |
| F-13 | INFO | Die Gegenprobe im Closure-Trigger von welle-18 verlangt einen Scan mit `--platform linux/arm64`; dieser Host hat keine arm64-Emulation (§1b M2), und unter Emulation prüfte die Gegenprobe laut ADR-0043 Option B den Emulator mit. Der Welle-Plan liegt außerhalb der Range; Adressat ist der Planner. | welle-18 §3; ADR-0043 Option B | `docs/plan/planning/welle-18-multi-arch-image.md:47-50` | nein | — (Hinweis an den Planner) |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| §1b M1, M2, M3, M6, M7 — Geltungsbereich je Zeile genannt, Folgerung (a) folgt aus M1/M2/M5 | geprüft, ohne Befund |
| §1b M5 — Index-Digest bei Kopie Registry → Registry | nachgemessen (`skopeo copy --all`, zwei `registry:2`), bestätigt; GHCR/Docker Hub bleiben außerhalb des Geltungsbereichs, im Slice als Risiko und in der ADR als Re-Evaluierungs-Trigger geführt |
| Zusage „derselbe Digest löst auf dem Spiegel auf" (AC-FA-DIST-002) | geprüft: gehalten durch die fail-closed-Prüfung des Index-Digests — weicht Docker Hub ab, scheitert das Release statt still zu lügen; ohne Befund |
| AC-Form AC-FA-DIST-001/-002 (drei Pfade + Out-of-Scope, Versions-Bump 0.30.0, Historie-Zeile) | geprüft, ohne Befund; `make doc-structure` Exit 0 |
| Spezifikation — Abwärts-Referenzen (ADR/Slice/Welle/Commit) in den neuen Zeilen und der Historie-Zeile 0.39.0 | geprüft, ohne Befund (`AGENTS.md` §3.4) |
| ADR-0043 — fünf Alternativen, Kontext mit Messung, Konsequenzen positiv/negativ, Re-Evaluierungs-Trigger, `Schärft:` SPEC-DIST-001, keine Slice-/Welle-Kennung im Text, Index-Zeile | geprüft; Befunde F-6, F-7, F-8, F-9, sonst ohne Befund |
| Teil-Ablösung ADR-0039 nach dem Muster ADR-0042 (Bezug-Feld, ADR-0039 bleibt `Accepted`, Index-Titel nennt die Ablösung) | geprüft; Befund F-7, Form sonst ohne Befund |
| ARC-008, ARC-009 (`spec/architecture.md`), AC-QA-01, AC-QA-03 | geprüft, kein Widerspruch zur neuen Fassung |
| `README.md`, `docs/user/benutzerhandbuch.md` — Aussagen zu Plattform oder Spiegel-Digest | geprüft, keine Fundstelle |
| `docs/user/releasing.md` — Lastenheft-Stand nachgezogen; Checkliste Zeile 6 (`docker inspect`) bleibt bis slice-218 Bestand, dort eingeplant | geprüft, ohne Befund |
| Hard Rules §3.1–§3.6 im Diff (kein Code, keine Suppression, kein Move, ADR nicht `Accepted` überschrieben, kein Gate gelockert) | geprüft, ohne Befund |
| `make doc-check`, `make doc-structure`, `make gate-consistency` | Exit 0 (Ausgabe in Datei, Exit-Code getrennt) |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 5 |
| LOW | 5 |
| INFO | 2 |

**Finding-Klassen dieses Laufs:** Messung mit zu engem Geltungsbereich als Vertragszusage
übernommen · Spezifikation sagt mehr zu als das Lastenheft · Zusage weiter als ihre Durchsetzung ·
Folgepflicht ohne Träger-Slice · Vertragsänderung ohne CHANGELOG-Zeile · Alternative mit
unbelegtem Contra verworfen · Teil-Ablösung zählt die betroffenen Stellen unvollständig ·
Re-Evaluierungs-Pfad ohne den Spec-Rang · Folge an einem bestehenden Sensor nicht benannt ·
Mehrdeutiger Begriff in einem Akzeptanzkriterium · Out-of-Scope-Punkt nach Vertragsänderung
gegenstandslos · Geltungsbereich nennt nicht alle tragenden Bedingungen

## Verdikt

**Abnahme-blockierend:** ja — F-1 (HIGH) und F-2 bis F-6 (MEDIUM) sind vor der Abnahme des
Vertrags durch den Maintainer zu klären. F-1 und F-3 treffen die Klasse, die der Slice in §6 selbst
als Risiko führt (*Zusage weiter als ihre Durchsetzung*); die Zuordnung zum Register geschieht bei
der Closure, nicht hier.

**Übergabe:** Findings an den Implementer; die Finding-Klassen zusätzlich in die Closure §7 von
slice-216. Dieser Report ist Lauf-Beleg und ersetzt keine Verifikation gegen die DoD.

---

## Delta-Review

**Gegenstand:** `6c08942..73efcaf` — `6c08942` (Plan-Änderung), `60b3964` (Lastenheft,
Spezifikation, CHANGELOG), `c6256a0` (ADR-0043, weiter `Proposed`), `73efcaf` (Messung M4/M5,
slice-217, slice-218, welle-18). Derselbe Reviewer-Kontext wie oben, gleicher Skill-Stand, gleiches
Modell; der Delta-Gegenstand stammt vom Implementer.
`make doc-check doc-structure gate-consistency`: Exit 0 (Ausgabe in Datei, Exit-Code getrennt).

### Zuordnung der Findings

| ID | Stand nach Delta | Beleg |
|---|---|---|
| F-1 | behoben — der Punkt „Reproduzierbar" ist aus SPEC-DIST-001 gestrichen; die Spezifikation sagt ausdrücklich, dass ein Neubau nicht denselben Digest zusagt; ADR-0043 Punkt 4 lautet „Keine Reproduzierbarkeits-Zusage", Kontext und §1b M4 nennen die Gegenmessung samt Mechanismus; die positive Konsequenz „lässt sich nachbauen" ist entfernt | `spec/spezifikation.md` SPEC-DIST-001; ADR-0043 Z. 41-44, Punkt 4; slice-216 §1b M4 |
| F-2 | behoben — mit F-1 entfällt die Zusage ohne Lastenheft-Träger | wie F-1 |
| F-3 | behoben — Ausgabe-Vergleich beider Plattformen in SPEC-DIST-001, ADR-0043 Punkt 5 (Bild-Tag erst nach Vergleich), Fitness-Zeile, DoD von slice-218 | siehe D-1 zur Reichweite |
| F-4 | behoben — (a) `tools/image-test.sh` mit Bild-Referenz in DoD und §3 von slice-217; (b) `packaging/dockerhub/overview.md` und `README.md` in DoD und §3 von slice-218 | slice-217 §2/§3; slice-218 §2/§3 |
| F-5 | behoben — `CHANGELOG.md` `[Unreleased]` trägt die Vertragsänderung, als „noch nicht implementiert" gekennzeichnet; §3 des Slice nennt die Datei | `CHANGELOG.md` Z. 9-19 |
| F-6 | behoben — D neu beschrieben (job-lokale Registry, Upload danach), Contra nennt den ungetesteten Weg zum Konsumenten; E-Contra nennt den lokalen Bildspeicher auch für E | ADR-0043 §Verglichene Alternativen |
| F-7 | behoben — Bezug nennt Punkt 1, das Wort in Punkt 2, §Konsequenzen und §Fitness Function | ADR-0043 Z. 12-17 |
| F-8 | behoben — Trigger nennt Spezifikations-Änderung und Folge-ADR | ADR-0043 §Re-Evaluierungs-Trigger |
| F-9 | behoben — als negative Konsequenz benannt und im Out-of-Scope von welle-18 | ADR-0043 §Konsequenzen; welle-18 §6 |
| F-10 | behoben — „Versions-Tag `vX.Y.Z` in der Registry" bzw. „Bild-Tag" in Lastenheft, Spezifikation und ADR | `spec/lastenheft.md` Z. 704 |
| F-11 | behoben — Out-of-Scope-Punkt gestrichen | `spec/lastenheft.md` Z. 744 |
| F-12 | behoben — Geltungsbereich von M5 nennt `overlay2` gegen containerd-Bildspeicher | slice-216 §1b M5 |
| F-13 | behoben — Gegenprobe nimmt den arm64-Lauf aus dem nativen Test-Job der Pipeline (Job-Log, Vergleichs-Schritt) und benennt, dass ein lokaler arm64-Lauf auf diesem Host nicht möglich ist | welle-18 §3 |

### Neue Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| D-1 | INFO | Die Spezifikation und ADR-0043 setzen den Vergleich an **einen** Scan einer festen Fixture; AC-FA-DIST-001 Boundary sagt Gleichheit für „denselben Scan" allgemein zu. Die Durchsetzung ist damit eine Stichprobe der Zusage — benannt in der Spezifikation („einer festen Fixture"), keine Lücke im Wortlaut, aber ihre Reichweite. | AC-FA-DIST-001; SPEC-DIST-001 | `spec/spezifikation.md` SPEC-DIST-001 (Getestet = veröffentlicht); `spec/lastenheft.md:703` | nein — Urteil über Reichweite | — (Hinweis) |
| D-2 | INFO | Lastenheft 0.30.0 und Spezifikation 0.39.0 ändern Inhalt ohne Versions-Bump; die Historie-Zeile der Spezifikation ist nachgezogen, die des Lastenhefts nennt den gestrichenen Out-of-Scope-Punkt nicht. Beide Stände sind im selben Slice unveröffentlicht (`Draft`), und slice-212 (`5855129`) hielt es im Nachlauf ebenso. | Lastenheft §Historie | `spec/lastenheft.md` Historie 0.30.0 | nein | — (Hinweis) |
| D-3 | INFO | Liefer-Punkt 3 von slice-218 bündelt jetzt vier Dokument-Familien (`releasing.md`, Benutzerhandbuch, Hub-Seite, CHANGELOG); die Größen-Regel zählt Punkte, nicht Dateien. Die Rückführung „zu groß" ist in §4 des Slice vorab benannt; Adressat ist der Planner beim Übergang nach `next/`. | `v6.13.0` · `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice | `docs/plan/planning/open/slice-218-multi-arch-release.md` §2 | nein | — (Hinweis an den Planner) |

### Negativbefunde (Delta)

| Bereich | Ergebnis |
|---|---|
| Spezifikation — Abwärts-Referenzen in den geänderten Zeilen und der Historie-Zeile 0.39.0 | geprüft, ohne Befund (`AGENTS.md` §3.4) |
| ADR-0043 — Rest-Verweise auf eine Reproduzierbarkeits-Zusage (`grep` nach `SOURCE_DATE_EPOCH`, „reproduzierbar") | geprüft: nur noch Kontext (Messung) und Option C (Contra), beide zutreffend; ohne Befund |
| Lastenheft — AC-Form nach den Änderungen | geprüft, ohne Befund; `make doc-structure` Exit 0 |
| slice-216 Plan-Änderung vor dem Fix (`AGENTS.md` §6 Schritt 4) | geprüft: Plan-Änderung `6c08942` liegt vor den Fix-Commits; ohne Befund |
| slice-217 — Größen-Regel nach Aufnahme des Image-Tests | geprüft: drei Liefer-Punkte; ohne Befund |
| Hard Rules §3.1–§3.6 im Delta | geprüft, ohne Befund |

### Summary (Delta)

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 0 |
| LOW | 0 |
| INFO | 3 |

Alle 13 Findings des ersten Laufs sind behoben. Neue Finding-Klassen: keine (nur Hinweise).

### Verdikt (Delta)

**Abnahme-blockierend:** nein. Der Vertrag kann dem Maintainer zur Abnahme vorgelegt werden;
ADR-0043 bleibt bis dahin `Proposed`. D-1 bis D-3 sind Hinweise ohne erwartete Aktion in diesem
Slice. Die Finding-Klassen des ersten Laufs gehen weiter in die Closure §7 von slice-216.
