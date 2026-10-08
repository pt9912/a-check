# `make image-scan` — CVE-Scan gegen das publizierte Image

## Vertrag

Trivy, digest-gepinnt, gegen die **publizierten** Images — nicht gegen einen
lokalen Build. `IMAGE_SCAN_REFS` führt **zwei** Referenzen:
`ghcr.io/pt9912/a-check:latest` und den Docker-Hub-Spiegel
`pt9912/a-check:latest`. Jede wird **je Plattform** gescannt
(`IMAGE_SCAN_PLATFORMS`: `linux/amd64`, `linux/arm64`). Vor Vollbericht und
Entscheidungslauf weist ein **eigener Lauf** die **gescannte Architektur** aus Trivys JSON
nach (Funktion `architektur_ok`, gedeckt von `--selftest`): weicht sie von der verlangten
ab oder fehlt sie, ist der Lauf GESCHEITERT. Grund: Trivy fällt bei `--platform` still auf
die vorhandene Plattform zurück — gemessen an einem reinen amd64-Image, „arm64" verlangt,
Exit 0, amd64 gescannt (slice-221). Der Entscheidungslauf bleibt beim Template-Format aus
[`ADR-0037`](../../docs/plan/adr/0037-cve-scan-gegen-das-publizierte-image.md) §Fitness Function. **Netz ist hier der
Zweck**, nicht ein Zugeständnis: eine gepinnte Vuln-DB fände nur die CVEs von
gestern. Genau deshalb hängt das Target in keinem Aggregat — `gates` ist
hermetisch.

Über rot entscheiden **nur behebbare** CRITICAL/HIGH. Der Vollbericht fällt nie.

## Grenze — was das Grün nicht abdeckt

1. **Nicht behebbare CVEs** — sie stehen im Bericht und färben nicht rot. Ein
   Grün heißt „nichts, wogegen ein Upgrade hilft", nicht „keine CVEs".
   Permanent, und die Absicht.
2. **Das lokale Image** — der Sensor kann es nicht scannen ([`ADR-0037`](../../docs/plan/adr/0037-cve-scan-gegen-das-publizierte-image.md),
   §Docker-Socket). Wer eine Hebung vor dem Release prüfen will, nimmt
   `docker save` plus Trivy `--input`. Heilbar nur durch Socket-Zugriff, der
   bewusst nicht gegeben wird.
3. **Der Zeitpunkt** — geprüft wird, was heute publiziert ist. Zwischen zwei
   Läufen kann das Image dieselben CVEs tragen und die DB neue kennen.
4. **Die Namen in Trivys Ausgaben** — das Template liest `.Vulnerabilities`,
   `.Severity`, `.FixedVersion`; der Architektur-Lauf den JSON-Schlüssel
   `"architecture"`. Benennt Trivy die Architektur um, wird der Lauf rot (GESCHEITERT);
   benennt es die Template-Felder um, rendert es nichts und sähe sauber aus. Der
   Digest-Pin hält das still, solange er steht; `--selftest` deckt Zählung und
   Architektur-Abgleich, nicht die Namen.
5. **Drei Läufe, ein Nachweis** — Architektur-Lauf, Vollbericht und Entscheidungslauf
   sind getrennte Trivy-Aufrufe mit derselben Referenz und Plattform. Wird ein Tag
   zwischen ihnen umgehängt, gilt der Nachweis dem ersten Lauf, nicht den beiden
   anderen; bei einem Digest-Pin entfällt das.

## Ausgabe und Ausgänge

| Exit | Bedeutung |
|---|---|
| 0 | keine behebbaren CRITICAL/HIGH |
| 1 | behebbare gefunden |
| 2 | Lauf-Fehler (Netz, Registry, Trivy), gescannte Architektur weicht ab oder fehlt, oder leere Prüfmenge |

**Über `make` sind 1 und 2 nicht unterscheidbar** — make normalisiert auf 2.
Wer den Ausgang braucht, liest die Ausgabe oder ruft `tools/image-scan.sh`
direkt auf. Das ist eine Eigenschaft von `make`, kein Mangel des Skripts.

Selbsttest der Auswertung: `bash tools/image-scan.sh --selftest`, netzlos — Zählung der
Template-Zeilen und Architektur-Abgleich (verlangte gegen gescannte, fehlende Angabe).

## Sperren

- `IMAGE_SCAN_REFS ist leer` / `IMAGE_SCAN_PLATFORMS ist leer` — der Lauf bricht
  **vor jeder Arbeit** (vor dem Anlegen des Caches) ab, Exit 2, mit der Begründung „nichts zu pruefen ist KEIN
  gruener Befundstand" → die Variable setzen oder den Default nicht überschreiben.

*(Kein Netz und ein nicht publiziertes Image führen zu Exit 2, sind aber
Lauf-**Ausgänge**, keine Sperren: Der Lauf hat dann bereits begonnen.)*

## Bindung

[`ADR-0037`](../../docs/plan/adr/0037-cve-scan-gegen-das-publizierte-image.md) · **kein Bestandteil von `gates`** (nicht hermetisch) · Nachtlauf und
`workflow_dispatch` in `.github/workflows/image-scan.yml` · slice-124, je Plattform seit slice-221
([`ADR-0043`](../../docs/plan/adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md)).
