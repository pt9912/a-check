# Review-Report: slice-218 — 2026-10-07

**Review-Art:** Code-Review — geprüft wird der Diff gegen den Slice-Plan (samt Plan-Änderung §1),
ADR-0043, SPEC-DIST-001, AC-FA-DIST-001/-002 und die Hard Rules (Modul 10). Unabhängiger Lauf in
frischem Kontext, nicht der Autor; keine DoD-Verifikation (Verifier, Modul 11).

**Gegenstand:** Commit-Range `332b3f5..bd68b1c` — `489739b` (Plan-Änderung), `bd68b1c` (Inhalt:
`.github/workflows/release.yml`, `tools/image-multiarch.sh`, `tools/image-test.sh`, `Makefile`,
`docs/user/releasing.md`, `docs/user/benutzerhandbuch.md`, `packaging/dockerhub/overview.md`,
`packaging/dockerhub/README.md`, `CHANGELOG.md`).

**Skill:** `.harness/skills/reviewer.md` @ `bd68b1c` · <!-- d-check:ignore -->
**Modell:** claude-opus-5-5 · **Datum:** 2026-10-07

> **Zitier-Form** *(Norm, kein Ausfüll-Hinweis)*. Dieser Report friert ein; was er zitiert, bewegt
> sich weiter. Deshalb: **Kennung, nicht Adresse** — `slice-<Kennung>` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine Baseline-Stelle als
> **Tag + Pfad in Inline-Code** statt als Link. Ein `pfad`-Feld auf den **geprüften Gegenstand**
> ist davon nicht betroffen — es zitiert den Stand des Laufs und darf ihn festhalten.

**Eingangs-Kontext:**

- slice-218 (§1–§8, Plan-Änderung §1 in `489739b`, vor dem Inhalts-Commit)
- ADR-0043 (`Accepted`), ADR-0007 (`Accepted`, Konsequenzen/Fitness Function), ADR-0030 (Verweis
  auf `releasing.md`), SPEC-DIST-001, AC-FA-DIST-001, AC-FA-DIST-002
- `AGENTS.md` §3 (Hard Rules), §6 Schritt 4 und 7
- frühere Findings am Bereich: Review slice-217 (F-10/F-11 an slice-218 übergeben; F-11
  „Versions-Label im Image-Test")
- Bestand: Vorgänger-Fassung von `release.yml` (`332b3f5`), `.github/workflows/hub-description.yml`

**Eigene Nachmessung** (Mess-Regel 1, Geltungsbereich): Host `x86_64`, Docker 29.8.2 (klassischer
Bildspeicher, overlay2), buildx 0.37.1, BuildKit-Pin aus dem `Makefile` in einem **eigenen**
Builder mit Host-Netz; zwei eigene `registry:2` (Ports 5201/5202), danach entfernt. **Keine**
arm64-Ausführung (kein Emulator auf dem Host) — der arm64-Test-Job und der Plattform-Vergleich
sind damit **nicht** gelaufen; gemessen ist die Mechanik je Schritt, nicht die Pipeline. GitHub-
Actions-Semantik (Job-Outputs, `needs`, `permissions`, Runner-Label) ist gelesen, nicht gelaufen.

| Probe | Ergebnis (Meldung) |
|---|---|
| Bau beider Plattformen mit **beiden** Ausgaben (OCI-Archiv + `push-by-digest`), Befehl wie im Skript | Exit 0; `multiarch-check` ok, Index `57097e51…` |
| Registry-Seite nach dem Upload | Tag-Liste: `NAME_UNKNOWN` (kein Tag); `HEAD manifests/<idx>` → 200, `Docker-Content-Digest` = Archiv-Digest |
| Metadaten-Datei `containerimage.digest` | = Archiv-Digest (Datei ist eingerückt; das `tr -d` im Skript greift) |
| Riegel-Mutation: Metadaten mit fremdem Digest, Schluss von `image-multiarch.sh` isoliert | rot — „hochgeladen sha256:00097e51…, geprüft sha256:57097e51…: nicht dasselbe Bild" |
| `imagetools create -t …:v9.9.9-rev -t …:latest <repo>@<idx>` + `inspect --format '{{.Manifest.Digest}}'` | Exit 0, gelesener Digest = `idx` |
| Kopie in die **zweite** Registry (Grenze wie GHCR → Docker Hub) | Exit 0, Index-Digest gleich, `inspect` zeigt `linux/amd64` + `linux/arm64` |
| `imagetools inspect` auf nicht existierendes Tag | Exit 1 — „not found" (der ERR-Trap im Spiegel-Schritt feuert damit) |
| ERR-Trap bei `X="$(false)"` unter `set -e` · explizites `exit 1` | Trap feuert · Trap feuert **nicht** (die Ungleich-Meldung nennt den GHCR-Digest selbst) |
| `make image-test IMAGE_REF=<repo>@<idx> VERSION=9.9.9-rev SCAN_OUT_DIR=…` | Exit 0; `scan.exit` = `1`, stdout eine Befundzeile mit `/src`-relativem Pfad, stderr Zusammenfassung — nichts Plattform-Abhängiges |
| `make image-test IMAGE_REF=<repo>@<idx>` **ohne** `VERSION` | rot, Exit 2 — „Versions-Label … ist '9.9.9-rev', erwartet '0.0.0-dev'" (F-3) |
| `make doc-workflows`, `make version-coherence`, `make doc-check`, `make doc-targets` | Exit 0 (Ausgabe in Datei, Exit-Code getrennt) |
| Labels beider Plattform-Configs | alle sechs OCI-Labels gesetzt (heute; F-4) |

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | Der Kopfkommentar sagt „jeder Schritt läuft über `make` oder die Docker-CLI". Der Schritt „Create GitHub Release" läuft über `gh`, der Vergleichs-Schritt und die Test-Schritte über `sha256sum`/`cut`/Bash-Vergleiche. Gegen die Datei selbst verifiziert. | Skill HIGH „nachweislich falsche Tatsachenbehauptung"; `AGENTS.md` §3.1/§3.7 | `.github/workflows/release.yml:23-24` | ja — `grep -n 'gh release\|sha256sum' release.yml` | Allquantor im Kommentar weiter als der Bestand |
| F-2 | HIGH | Der Kommentar über dem Spiegel trägt einen Konjunktiv über die verworfene Alternative: „ein Neu-Push aus dem lokalen Bildspeicher schriebe nur die Host-Plattform" — Begründung aus ADR-0043 §Kontext, nicht Zusage/Kopplung/Abgrenzung am Ort. Dieselbe Form im neu umgebrochenen Kommentar über `hub-description` („… wäre eine Falschaussage"; Wortlaut aus dem Bestand übernommen). | `AGENTS.md` §3.7; Skill HIGH „Kommentar trägt keine der Kommentar-Klassen" | `.github/workflows/release.yml:226-227`, `:281-283` | nein — inferentiell, kein Gate | Kommentar beschreibt die verworfene Alternative |
| F-3 | MEDIUM | `make image-test` reicht `EXPECT_VERSION=$(VERSION)` immer durch; `VERSION` hat den Default `0.0.0-dev`. Damit endet der im CHANGELOG und im Gate-Index beschriebene Aufruf `make image-test IMAGE_REF=…` gegen **jedes** veröffentlichte Bild rot (gemessen). Der Makefile-Kommentar nennt die Pflicht, die `##`-Hilfe, der CHANGELOG-Eintrag und die Gate-Index-Zeile `make image-test` nicht. | AC-FA-DIST-001 (Image-Test gegen Bild-Referenz, ADR-0043 §Fitness Function); `AGENTS.md` §6 Schritt 7 | `Makefile:227-228`; `CHANGELOG.md:15-17`; `tools/image-test.sh:50-52` | ja — `make image-test IMAGE_REF=<release-ref>` ohne `VERSION` | Neue Pflicht-Eingabe still an einen bestehenden Aufruf gekoppelt |
| F-4 | MEDIUM | Der Vorgänger-Schritt „Verify OCI labels" prüfte zusätzlich `…image.source`, `…description`, `…licenses`, `…title`, `…vendor` auf nicht leer. Die neue Pipeline prüft nur das Versions-Label (`multiarch-check`, `image-test`); die übrigen fünf Prüfungen entfallen, ohne dass Plan §1, ADR-0043 oder `releasing.md` das nennen. Ob das unter `AGENTS.md` §3.6 („Prüfregel") fällt, ist eine Architect-Frage — der Schritt steht nicht im Gate-Index. | `AGENTS.md` §3.6 (offen); Slice-Plan §1 | `.github/workflows/release.yml` (Vorgänger `332b3f5`, Schritt „Verify OCI labels") | ja — `git show 332b3f5:.github/workflows/release.yml` gegen den Diff | Bestehende Prüfung beim Umbau still entfallen |
| F-5 | LOW | Der Scan-Hash läuft über die **Verkettung** `scan.exit`+`scan.stdout`+`scan.stderr` ohne Trenner. Ein Byte, das auf einer Plattform von stdout nach stderr wandert (etwa die Zusammenfassungszeile), ergibt denselben Hash; Spec und ADR sagen „byte-identische Ausgabe" je Strom. Bei gleichem Quellstand praktisch ausgeschlossen. | SPEC-DIST-001 („byte-identische Ausgabe und denselben Exit-Code") | `.github/workflows/release.yml:133`, `:165` | ja — zwei Dateisätze mit verschobener Grenze hashen | Vergleichs-Größe gröber als die Zusage |
| F-6 | LOW | Der Bild-Tag-Schritt setzt `vX.Y.Z` per `imagetools create` ohne Existenz-Prüfung. Ein „Re-run all jobs" (naheliegend nach einem roten Spiegel bei fehlenden Hub-Secrets) baut neu — ADR-0043 Punkt 4: anderer Digest — und überschreibt den bereits gesetzten GHCR-Tag; die Rücklese-Prüfung (`GOT = DIGEST`) ist dann grün. `releasing.md` §Incident-Klausel sagt „ein veröffentlichter Versions-Tag wird nie überschrieben", §Ablauf nennt den Wiederanlauf nicht. Gleiches Verhalten im Vorgänger (`docker push`). | ADR-0043 Punkt 4/5; `releasing.md` §Incident-Klausel | `.github/workflows/release.yml:206-214`; `docs/user/releasing.md:51-90` | ja — Workflow zweimal auf denselben Tag | Wiederanlauf-Pfad unbenannt |
| F-7 | LOW | Die Umnummerierung in `releasing.md` (Re-Pin jetzt Schritt 8, Schritt 6 ist der Spiegel) lässt den Verweis „der Re-Pin (`releasing.md` Schritt 6)" in der `Accepted`-ADR-0030 auf den falschen Schritt zeigen; die ADR ist immutabel. | `AGENTS.md` §3.5; `harness/rules/zitier-form-einfrierende.md` | `docs/user/releasing.md:54-90`; `docs/plan/adr/0030-kein-digest-im-generierten-fragment.md:17` | ja — `grep -n 'Schritt 6' docs/plan/adr/0030-*` | Positions-Verweis in einfrierendem Artefakt driftet |
| F-8 | LOW | `overview.md` sagt unbedingt „The digest listed in this project's own documentation … therefore resolves here as well". Für die vor diesem Release gespiegelten Tags (neu hochgeladen, Manifest-Digest verschieden — `packaging/dockerhub/README.md` misst es selbst) und im Fenster zwischen Tag und Re-Pin (Schritt 8; die Doku nennt dann noch den Vorgänger-Digest) gilt das nicht. | AC-FA-DIST-002 („löst derselbe Digest auch auf dem Spiegel auf") | `packaging/dockerhub/overview.md:37-40` | ja — Vorgänger-Digest auf Docker Hub auflösen | Zusage ohne zeitliche Grenze |
| F-9 | LOW | §Vorbedingungen sagt „**Zwei** Dinge dieses Betriebs leben nicht im Repo"; die Tabelle führt mit der neuen arm64-Zeile **fünf**. Die Zahl war schon vorher falsch (vier), der Diff erweitert die Tabelle, ohne sie nachzuziehen. Dazu: Der Kommentar „Gleich bis auf `runs-on`" über den Test-Jobs — sie unterscheiden sich auch in `name` und Ausgabe-Text (Anzeige, nicht Mechanik). | Mess-Regel „zweimal verschieden zählen" | `docs/user/releasing.md:98`; `.github/workflows/release.yml:103-104` | ja — Zeilen zählen | Zahl neben einer Aufzählung nicht nachgezogen |
| F-10 | INFO | Die Plattform-Menge auf dem Spiegel (ADR-0043 Punkt 7, SPEC-DIST-001, AC-FA-DIST-002 Happy „nennt beide Plattformen") wird nicht eigens geprüft; sie folgt aus der Digest-Gleichheit (inhaltsadressiert) mit dem Index, dessen Archiv `multiarch-check` geprüft und dessen Upload-Digest der Riegel gleichgesetzt hat. Das `imagetools inspect` danach gibt aus, prüft nichts. Die Kette trägt; benannt ist sie nirgends. | ADR-0043 Punkt 7 | `.github/workflows/release.yml:245-251` | nein | Geprüfte Größe nur per Implikation gedeckt |
| F-11 | INFO | `image-multiarch.sh` liest den Index-Digest mit einem **anderen** Extraktor (erster `sha256:` in `index.json`) als `multiarch-check.sh` (Schlüssel `"digest"`); der Kommentar „hier liegt also einer vor" gilt für das Vorhandensein, nicht für denselben Wert. Bei BuildKit-förmigem `index.json` gleich (gemessen). Der Workflow liest `.build/a-check-multiarch.oci.tar.digest` fest verdrahtet — gekoppelt an den Default von `MULTIARCH_OUT`, unkommentiert. | Mess-Regel „zweimal verschieden zählen" | `tools/image-multiarch.sh:57-60`; `.github/workflows/release.yml:98` | ja | Zwei Extraktoren für denselben Wert |
| F-12 | INFO | Schlägt der Docker-Hub-**Login** fehl (Secrets fehlen), läuft der Spiegel-Schritt nicht, und die Meldung „BEREITS VEROEFFENTLICHT: …@<digest>" erscheint nicht; der Digest steht dann nur im Job-Summary des Tag-Schritts. Struktur unverändert aus dem Vorgänger. | AC-FA-DIST-002 Negative | `.github/workflows/release.yml:230-238` | ja | Fehlermeldung deckt nicht jeden Spiegel-Fehlpfad |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Job-Graph: `build → test-amd64/test-arm64 → publish → hub-description`; Outputs `version`/`is_stable`/`digest`/`scan` je Job deklariert und gelesen; `hub-description` liest `needs.build.outputs` und führt `build` in `needs` | gelesen, ohne Befund |
| `permissions` je Job: `build` packages write (Upload), Tests packages read, `publish` contents write (Release) + packages write (Tag), `hub-description` contents read; Top-Level `{}` | ohne Befund |
| Login-Weitergabe an den `docker-container`-Builder (Client-Credentials) und `imagetools create` über zwei Registries mit beiden Logins | Mechanik lokal ohne Auth nachgefahren (Kopie zwischen zwei Registries, Digest gleich); Auth-Weitergabe gelesen, nicht gelaufen |
| Verschluckte Fehler: `set -euo pipefail` in jedem `run:`; leerer Digest per Regex, leere Scan-Hashes per Nicht-Leer-Riegel; fehlende Scan-Dateien brechen `cat` | ohne Befund |
| Scan-Ausgabe auf Plattform-Abhängiges | gemessen (amd64): `/src`-relative Pfade, keine Plattform-, Zeit- oder Versionsangabe. arm64 nicht gemessen |
| `ubuntu-24.04-arm` als Label; Docker und `make` auf beiden Runnern; `releasing.md` §Vorbedingungen nennt den Runner (ADR-0043 Folgepflicht) | gelesen, ohne Befund |
| `uses:`-Form (`make doc-workflows`, `make version-coherence`) — keine neue Action, alle SHA-gepinnt mit Tag-Kommentar | Exit 0, ohne Befund |
| ADR-0043 Punkte 1–8 gegen Pipeline und Skripte: ein Bau, Upload ohne Tag, nativer Test je Plattform samt Versions-Label, Scan-Vergleich, Tag erst danach auf denselben Digest, Spiegel per Kopie mit Index-Digest-Gleichheit, fail-closed | gedeckt; Befunde nur F-5, F-10 |
| Handbuch 1.46: Plattformen, Index-Digest als Pin, `imagetools inspect --format '{{.Manifest.Digest}}'` | Befehl nachgefahren, liefert den Index-Digest; ohne Befund |
| `releasing.md` Schritte 1–8 gegen die Jobs; Checkliste Item 6 | stimmt mit der Pipeline überein; Befunde F-6, F-7, F-9 |
| Plan-Änderung vor Code (`AGENTS.md` §6 Schritt 4): `489739b` vor `bd68b1c`, §3 führt die Skripte und das `Makefile` | ohne Befund |
| Hard Rules §3.1–§3.5 (Docker/make-only, keine Suppression, kein Move mit Inhalt, kein Spec-Stratum berührt, keine ADR geändert) | ohne Befund; §3.6 offen in F-4 |
| CHANGELOG `[Unreleased]` trägt die Änderung des öffentlichen Vertrags (`AGENTS.md` §6 Schritt 7) | vorhanden; Befund nur F-3 |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 2 |
| LOW | 5 |
| INFO | 3 |

**Finding-Klassen dieses Laufs:** Allquantor im Kommentar weiter als der Bestand · Kommentar
beschreibt die verworfene Alternative · Neue Pflicht-Eingabe still an einen bestehenden Aufruf
gekoppelt · Bestehende Prüfung beim Umbau still entfallen · Vergleichs-Größe gröber als die Zusage ·
Wiederanlauf-Pfad unbenannt · Positions-Verweis in einfrierendem Artefakt driftet · Zusage ohne
zeitliche Grenze · Zahl neben einer Aufzählung nicht nachgezogen · Geprüfte Größe nur per
Implikation gedeckt · Zwei Extraktoren für denselben Wert · Fehlermeldung deckt nicht jeden
Spiegel-Fehlpfad

## Verdikt

**Abnahme-blockierend:** ja, nach der HIGH-Liste des Skills — F-1 und F-2 (beide Kommentare, beide
gegen die Datei verifiziert) und F-3/F-4 (MEDIUM) sind vor der Closure zu klären. F-4 enthält eine
Architect-Frage (§3.6: ist der entfallene Label-Schritt eine „Prüfregel"?); widerspricht der
Implementer dort, läuft der Konflikt-Pfad über den Architect.

**Zur Hauptfrage — würde die Pipeline am Tag funktionieren?** Jede Einzel-Mechanik, die sich ohne
GitHub und ohne arm64 nachfahren lässt, ist nachgefahren und trägt: Upload ohne Tag mit
Registry-Digest gleich Archiv-Digest, Riegel rot bei Abweichung, Tag-Setzen und Kopie über zwei
Registries mit gleichem Index-Digest und beiden Plattformen, `{{.Manifest.Digest}}`-Rücklese,
Image-Test gegen den Digest samt Label, Scan-Ausgabe ohne Plattform-Anteil. Ungeprüft bleiben
Auth-Weitergabe gegen GHCR/Docker Hub, der arm64-Runner und der Vergleich zweier echter
Plattform-Läufe — das ist das in §6 benannte Risiko „Erst am Tag geprüft", kein Befund.

**Übergabe:** Findings an den Implementer; die Finding-Klassen zusätzlich in die Closure §7 von
slice-218. Dieser Report ist Lauf-Beleg und ersetzt keine Verifikation gegen die DoD.
