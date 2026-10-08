# `make image-scan` — CVE-Scan gegen das publizierte Image

## Vertrag

Trivy, digest-gepinnt, gegen die **publizierten** Images — nicht gegen einen
lokalen Build. `IMAGE_SCAN_REFS` führt **zwei** Referenzen:
`ghcr.io/pt9912/a-check:latest` und den Docker-Hub-Spiegel
`pt9912/a-check:latest`. Jede wird **je Plattform** gescannt
(`IMAGE_SCAN_PLATFORMS`: `linux/amd64`, `linux/arm64`), und der Entscheidungslauf
weist die **gescannte Architektur** aus Trivys JSON nach: weicht sie von der
verlangten ab oder fehlt sie, ist der Lauf GESCHEITERT. Grund: Trivy fällt bei
`--platform` still auf die vorhandene Plattform zurück — gemessen an einem reinen
amd64-Image, „arm64" verlangt, Exit 0, amd64 gescannt (slice-221). **Netz ist hier der
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
4. **Die Schlüssel in Trivys JSON** — gezählt wird `"VulnerabilityID"`, die
   Architektur steht unter `"architecture"`. Benennt Trivy die Architektur um, wird
   der Lauf rot (GESCHEITERT); benennt es `"VulnerabilityID"` um, zählt er null
   Befunde und sähe sauber aus. Der Digest-Pin hält das still, solange er steht;
   `--selftest` deckt die Auswertung, nicht die Schlüssel.
5. **Der Vollbericht** (Tabelle) läuft mit `--platform`, weist seine Architektur aber
   nicht selbst nach — das tut der Entscheidungslauf derselben Plattform.

## Ausgabe und Ausgänge

| Exit | Bedeutung |
|---|---|
| 0 | keine behebbaren CRITICAL/HIGH |
| 1 | behebbare gefunden |
| 2 | Lauf-Fehler (Netz, Registry, Trivy) |

**Über `make` sind 1 und 2 nicht unterscheidbar** — make normalisiert auf 2.
Wer den Ausgang braucht, liest die Ausgabe oder ruft `tools/image-scan.sh`
direkt auf. Eine Abweichung der gescannten Architektur ist Exit 2 (Lauf-Ausgang
„GESCHEITERT"), nicht Exit 1. Das ist eine Eigenschaft von `make`, kein Mangel des Skripts.

Selbsttest der Auswertung: `bash tools/image-scan.sh --selftest`, netzlos.

## Sperren

- `IMAGE_SCAN_REFS ist leer` / `IMAGE_SCAN_PLATFORMS ist leer` — der Lauf bricht
  **vor jeder Arbeit** ab, Exit 2, mit der Begründung „nichts zu pruefen ist KEIN
  gruener Befundstand" → die Variable setzen oder den Default nicht überschreiben.

*(Kein Netz und ein nicht publiziertes Image führen zu Exit 2, sind aber
Lauf-**Ausgänge**, keine Sperren: Der Lauf hat dann bereits begonnen.)*

## Bindung

[`ADR-0037`](../../docs/plan/adr/0037-cve-scan-gegen-das-publizierte-image.md) · **kein Bestandteil von `gates`** (nicht hermetisch) · Nachtlauf und
`workflow_dispatch` in `.github/workflows/image-scan.yml` · slice-124, je Plattform seit slice-221
([`ADR-0043`](../../docs/plan/adr/0043-multi-arch-ein-bau-getestet-dann-getaggt.md)).
