# Review-Report: slice-221 — 2026-10-08

**Review-Art:** Code — geprüft gegen den Slice-Plan (inkl. Messung §1b), ADR-0037, ADR-0043,
die Hard Rules in `AGENTS.md` §3 und die Mess-Regeln (`harness/rules/mess-regeln.md`).

**Gegenstand:** Commit-Range `d6ddf9d..2236547` — Inhalt in `2236547`
(`tools/image-scan.sh`, `harness/sensors/image-scan.md`, Gate-Index-Zelle in `harness/README.md`,
`CHANGELOG.md`); `69d4938` verlinkt nur Kennungen im Slice-Plan.

**Skill:** `.harness/skills/reviewer.md` @ `5b349a7` · <!-- d-check:ignore -->
**Modell:** claude-opus-5-5 · **Datum:** 2026-10-08

> **Zitier-Form** *(dieser Block bleibt stehen — er ist Norm, kein
> Ausfüll-Hinweis)*. Dieser Report friert ein; was er zitiert, bewegt sich
> weiter. Deshalb: **Kennung, nicht Adresse** — `slice-<Kennung>` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link. Ein
> `pfad`-Feld auf den **geprüften Gegenstand** zitiert den Stand des Laufs und
> darf ihn festhalten.

**Eingangs-Kontext:**

- slice-221 (Slice-Plan in `in-progress/`, Stand `69d4938`)
- ADR-0037 (CVE-Scan gegen das publizierte Image; Accepted) — insbesondere Punkt 3 und §Fitness Function
- ADR-0043 (Multi-Arch; §Konsequenzen, die benannte arm64-Lücke)
- `AGENTS.md` §3 (Hard Rules), §3.7
- `harness/rules/mess-regeln.md` (Regeln 1–5), `.github/workflows/image-scan.yml` (Leser der Marker)

**Eigene Läufe dieses Reviews** (Arbeitsdateien ausschließlich im Scratchpad, nicht im Repo):

| Lauf | Ergebnis |
|---|---|
| `bash tools/image-scan.sh --selftest` | Exit 0, 11× ok, „Fehlschlaege: 0" |
| `bash tools/image-scan.sh` (Default, gegen das Publizierte) | Exit 0; 4× „OK — …" mit `gescannt: amd64` / `gescannt: arm64` für beide Referenzen |
| `IMAGE_SCAN_REFS=ghcr.io/pt9912/a-check:v0.22.0` | Exit 2; „Trivy hat die Architektur 'amd64' gescannt, verlangt war 'arm64' — GESCHEITERT (kein Nachweis fuer linux/arm64)." |
| dasselbe mit entferntem Abgleich (Mutation `if false`) | Exit 0; „OK … (linux/arm64, gescannt: amd64)" — die Gegenprobe ist also rot **wegen** des Abgleichs |
| `IMAGE_SCAN_PLATFORMS=" "` | Exit 2, „IMAGE_SCAN_PLATFORMS ist leer — …" |
| Befund-Pfad (Mutation: Entscheidungslauf `--severity UNKNOWN --ignore-unfixed`, eine Referenz, arm64) | Exit 1; Zeile `DLA-4792-1 tzdata`; Marker „1 behebbare CRITICAL/HIGH-Befunde." |
| Zweiter Zähler (Regel 3): `"Severity":`-Zeilen vs. `zaehle()` auf fünf JSON-Ausgaben | deckungsgleich (1/1/0/0/0); Trivy-Template-Zählung vs. JSON-Zählung bei `--severity UNKNOWN --ignore-unfixed`: 1 = 1 |
| Mutation Anker entfernt (`grep -cE '"VulnerabilityID":'`) | **Selbsttest grün** („Fehlschlaege: 0") — siehe F-1 |
| Mutation `grep -c VulnerabilityID` (ohne Anführungszeichen) | Selbsttest rot: „FAIL Schluessel nur im Text erwartet 0, war: 1" |
| `make doc-structure`, `make doc-check` | je Exit 0, 0 Befunde |

Geltungsbereich: Trivy `0.74.0` am gepinnten Digest, die Images `ghcr.io/pt9912/a-check:latest`,
`pt9912/a-check:latest`, `ghcr.io/pt9912/a-check:v0.22.0` am 2026-10-08. `--ignore-unfixed` ist im
JSON-Lauf **nicht** direkt belegt: das Image trägt keinen unbehebbaren Befund, an dem die Filterung
sichtbar würde; belegt ist nur, dass `--severity` im JSON-Lauf filtert und der einzige Befund
(`tzdata`, `Status: fixed`) bei `--ignore-unfixed` erhalten bleibt.

---

## Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| F-1 | HIGH | Der Selbsttest-Kommentar nennt `Schluessel nur im Text` „die eigentliche Probe des Ankers"; mit entferntem Anker (`^[[:space:]]*`) bleibt der Selbsttest grün (0 Fehlschläge) — die Probe hält die Eigenschaft über das Escaping von `\"` im JSON, nicht über den Anker. Damit ist auch die in ADR-0037 §Fitness Function beschriebene Anker-Probe („Marker nur am Zeilenende, darf **nicht** zählen"), die vor dem Umbau rot wurde, beim Umbau still entfallen. | Mess-Regel 4; ADR-0037 §Fitness Function; `BEO-GATE/umbau-verliert-pruefung-still` | `tools/image-scan.sh:128-131` | ja — Mutation `grep -cE '"VulnerabilityID":'` + `--selftest` | Testkommentar weiter als seine Assertion |
| F-2 | MEDIUM | ADR-0037 §Fitness Function (Accepted, immutabel) beschreibt den Selbsttest als sieben Fixtures über das Template-Format („Marker nur am Zeilenende", „leer gerenderte Feldnamen", „rendert das Template nichts"); der Diff stellt den Entscheidungslauf auf JSON um und ersetzt diese Fixtures. Der Fitness-Function-Anker der ADR beschreibt damit einen Selbsttest, den es nicht mehr gibt; weder Slice-Plan §3 noch eine Folge-ADR nennen den Formatwechsel. | ADR-0037 §Fitness Function; `AGENTS.md` §3.5; `v6.13.0` · `regelwerk/modul-08-agentenrollen.md` §Rollen-Regeln | `tools/image-scan.sh:81-83`; ADR-0037 Z. 71-77 | ja — Abgleich ADR-Text gegen `--selftest`-Ausgabe | Fitness-Function-Anker driftet beim Umbau vom Code |
| F-3 | MEDIUM | Kopf-Kommentar (GRENZE) und Sensor-Datei Grenze 4 sagen „`--selftest` deckt die Auswertung, nicht die Schlüssel"; der Abgleich `[ "${got}" != "${want}" ]` und die Ableitung `want="${plat#*/}"` liegen außerhalb der getesteten Funktionen — mit entferntem Abgleich bleibt der Selbsttest grün. Gehalten wird der Abgleich nur von der einmaligen Live-Gegenprobe; der Satz nennt diese Grenze nicht. | Mess-Regel 5 | `tools/image-scan.sh:44-45`, `:187`, `:213-217`; `harness/sensors/image-scan.md` Grenze 4 | ja — Mutation `if false` + `--selftest` | Zusage weiter als ihre Durchsetzung |
| F-4 | LOW | Der eingeschobene Satz zur Architektur-Abweichung steht zwischen „… ruft `tools/image-scan.sh` direkt auf." und „Das ist eine Eigenschaft von `make`, kein Mangel des Skripts." — das „Das" bezieht sich jetzt auf die Exit-2-Zuordnung der Architektur-Abweichung, die keine Eigenschaft von `make` ist. | `AGENTS.md` §3.7 (Teilersetzung) | `harness/sensors/image-scan.md` §Ausgabe und Ausgänge | nein — Leseurteil | Teilersetzung verschiebt den Bezug eines Folgesatzes |
| F-5 | LOW | Die Exit-Tabelle der Sensor-Datei führt für 2 weiterhin nur „Lauf-Fehler (Netz, Registry, Trivy)"; Architektur-Abweichung/fehlende Angabe und leere Prüfmenge enden ebenfalls mit 2 und stehen nur in Prosa bzw. unter §Sperren. | Maintainability | `harness/sensors/image-scan.md` §Ausgabe und Ausgänge | nein — Leseurteil | Ausgangs-Tabelle unvollständig gegen das Skript |
| F-6 | LOW | Die Bindung-Zelle der Gate-Index-Zeile nennt weiter nur ADR-0037 und slice-124; die Sensor-Datei bindet die Plattform-Prüfung zusätzlich an ADR-0043 und slice-221. | Gate-Index-Bindung (`harness/README.md` §Sensors) | `harness/README.md` Zeile `make image-scan` | nein — Leseurteil | Bindung-Spalte nicht nachgezogen |
| F-7 | INFO | `IMAGE_SCAN_PLATFORMS`-Sperre steht nach `mkdir -p "$CACHE"`; die Sensor-Datei sagt „bricht vor jeder Arbeit ab". Folgenlos für den Befundstand. | Mess-Regel 5 (Wortlaut) | `tools/image-scan.sh:168`, `:177-180` | ja — Lesen der Reihenfolge | Sperre nach erster Seitenwirkung |
| F-8 | INFO | `befunde()` paart `VulnerabilityID`- und `PkgName`-Zeilen positionsweise (`paste - -`) und ist ungetestet; fehlt einem Befund `PkgName`, verschiebt sich nur die lesbare Liste — Zählung und Exit hängen nicht daran. | Maintainability | `tools/image-scan.sh:103-107` | nein | Anzeige-Funktion ungetestet |

## Negativbefunde

| Bereich | Ergebnis |
|---|---|
| Vollbericht je Referenz/Plattform: läuft weiterhin, fällt nie, Scheitern → „GESCHEITERT", `errored=1` | geprüft, ohne Befund |
| Entscheidungslauf: `--severity CRITICAL,HIGH --ignore-unfixed` wirkt im JSON-Lauf (Severity-Filter gemessen, siehe Geltungsbereich oben) | geprüft, ohne Befund |
| Zählung `"VulnerabilityID"`: zählt nur Befunde — eine Schlüsselzeile je Befund im eingerückten JSON, kein anderer Abschnitt trägt den Schlüssel (`--scanners vuln`, ohne `--show-suppressed`); zweiter Zähler deckungsgleich | geprüft, ohne Befund |
| Architektur-Abgleich falsch grün? Einziges `"architecture"` im JSON steht unter `Metadata.ImageConfig` (Z. 41, vor `Results` Z. 223); in Strings ist `"` als `\"` escaped und matcht das Muster nicht; Abweichung und fehlende Angabe → Exit 2 | geprüft, ohne Befund |
| Leere Prüfmenge (`IMAGE_SCAN_REFS`, `IMAGE_SCAN_PLATFORMS`) → Exit 2 | geprüft, ohne Befund |
| Exit-Code-Semantik 0/1/2, Vorrang „gescheitert" vor „Befund" | geprüft, ohne Befund |
| Marker für `.github/workflows/image-scan.yml`: „GESCHEITERT" in allen drei Fehlerpfaden inkl. Architektur, „behebbare CRITICAL/HIGH-Befunde" im Befundpfad (live belegt) | geprüft, ohne Befund |
| Hard Rules §3.1 (nur Docker/bash/grep/sed, kein Fremd-Interpreter), §3.2, §3.6 (keine Schwelle gesenkt) | geprüft, ohne Befund |
| Kommentar-Klassen §3.7 im neuen Kopf-Block PLATTFORMEN und in den Funktions-Kommentaren | geprüft, ohne Befund (außer F-1/F-3) |
| Gate-Index-Zelle: Länge < 250, `make doc-structure` grün; „Grenzen: siehe Datei" erfüllt Mess-Regel 5 | geprüft, ohne Befund |
| CHANGELOG `[Unreleased]`: Eintrag vorhanden, Aussage deckt sich mit dem Lauf | geprüft, ohne Befund |
| Behauptete Belege des Implementers (4× sauber, Gegenprobe v0.22.0 rot mit genannter Meldung, ohne Abgleich grün) | nachgefahren, reproduziert |

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 |
| MEDIUM | 2 |
| LOW | 3 |
| INFO | 2 |

**Finding-Klassen dieses Laufs:** Testkommentar weiter als seine Assertion · Fitness-Function-Anker driftet beim Umbau vom Code · Zusage weiter als ihre Durchsetzung · Teilersetzung verschiebt den Bezug eines Folgesatzes · Ausgangs-Tabelle unvollständig gegen das Skript · Bindung-Spalte nicht nachgezogen · Sperre nach erster Seitenwirkung · Anzeige-Funktion ungetestet

## Verdikt

**Merge-blockierend:** ja — F-1 (HIGH) und F-2/F-3 (MEDIUM). Die Kernleistung des Slice
(Plattform-Schleife, fail-closed Architektur-Abgleich, erhaltene Zählung und Marker) ist
nachgefahren und trägt; blockierend sind die Belege **über** sie: eine Selbsttest-Zusage, die
eine Mutation überlebt, und ein ADR-Fitness-Function-Anker, der nach dem Formatwechsel einen
nicht mehr existierenden Selbsttest beschreibt. F-1 ist für `BEO-GATE/umbau-verliert-pruefung-still`
ein Kandidat für den dritten Beleg — die Zuordnung trifft die Closure.

**Übergabe:** Findings an den Implementer; F-2 berührt eine Accepted-ADR und geht, falls der
Implementer widerspricht, den Konflikt-Pfad über den Architect (`v6.13.0` ·
`regelwerk/modul-08-agentenrollen.md` §Konflikt-Pfad als Rollen-Sequenz). Die Finding-Klassen
gehen in die Closure §7. DoD-Konformität prüft der Verifier separat.

---

## Delta-Review

**Gegenstand:** Commit-Range `df194a5..7d721e7` — `d01651b` (Plan-Änderung in slice-221 §1b, vor
dem Fix) und `7d721e7` (`tools/image-scan.sh` von `2236547~1` neu aufgebaut, Sensor-Datei,
Gate-Index-Zelle). **Review-Art:** Code. **Modell:** claude-opus-5-5 · **Datum:** 2026-10-08.
Geprüft gegen dieselben Verträge wie oben; das Skript wurde gegen `2236547~1` gediffed, nicht nur
gegen den ersten Wurf.

**Eigene Läufe** (Arbeitsdateien nur im Scratchpad):

| Lauf | Ergebnis |
|---|---|
| `bash tools/image-scan.sh --selftest` | Exit 0, 12× ok (sieben Zähl-Fixtures aus ADR-0037 unverändert, fünf Architektur-Fälle), „Fehlschlaege: 0" |
| Mutation Anker (`'^FINDING '` → `'FINDING '`) | rot: „FAIL Marker nur am Zeilenende erwartet 0, war: 1" |
| Mutation „Abgleich immer ok" (`if true`) | rot: „FAIL arm64 verlangt, amd64 gescannt erwartet abweichung, war: ok" (+ „Angabe fehlt", „erste Angabe gilt") |
| Mutation `want="$1"` | rot: „FAIL arm64 verlangt, arm64 gescannt erwartet ok, war: abweichung" (+ amd64) |
| Mutation „letzte statt erste Angabe" (`sed -n '$p'`) | rot: „FAIL erste Angabe gilt erwartet abweichung, war: ok" |
| Mutation „gescannt = verlangt" (Extraktor umgangen) | rot: drei FAIL wie bei „Abgleich immer ok" |
| Mutation `[ -n "${got}" ]` entfernt | **grün** — siehe D-3 |
| `bash tools/image-scan.sh` gegen das Publizierte | Exit 0, 4× „OK … gescannt: amd64/arm64" |
| `IMAGE_SCAN_REFS=ghcr.io/pt9912/a-check:v0.22.0` | Exit 2; amd64 OK; arm64: „Trivy hat die Architektur 'amd64' gescannt, verlangt war 'arm64' — GESCHEITERT (kein Nachweis fuer linux/arm64)." — für arm64 läuft **kein** Vollbericht mehr |
| Befund-Pfad (Mutation `--severity UNKNOWN`, eine Referenz, arm64) | Exit 1; Zeile `UNKNOWN tzdata … -> fix … DLA-4792-1`; Marker „1 behebbare CRITICAL/HIGH-Befunde." |
| `IMAGE_SCAN_PLATFORMS=" "` mit `TRIVY_CACHE` auf ein Scratch-Verzeichnis | Exit 2, Meldung „… ist leer …"; das Cache-Verzeichnis wird **nicht** angelegt |

### Stand der Findings aus dem ersten Lauf

| ID | Stand | Beleg |
|---|---|---|
| F-1 | behoben | Entscheidungslauf und Zähl-Fixtures wieder auf dem Stand von `2236547~1`; Anker-Mutation rot (oben) |
| F-2 | behoben | Template-Entscheidung und die sieben Fixtures aus ADR-0037 §Fitness Function stehen wieder; JSON nur im zusätzlichen Architektur-Lauf; Plan-Änderung in §1b vor dem Fix (`d01651b`) |
| F-3 | behoben | Abgleich in `architektur_ok`, im Selbsttest; Mutationen „immer ok", `want="$1"`, Extraktor umgangen je rot. Der Aufruf in der Schleife ist netzlos nicht testbar und wird von der Live-Gegenprobe gehalten |
| F-4 | behoben | eingeschobener Satz entfernt, „Das" bezieht sich wieder auf die make-Normalisierung |
| F-5 | behoben | Exit-2-Zeile nennt Architektur-Abweichung/fehlende Angabe und leere Prüfmenge |
| F-6 | behoben | Bindung-Zelle: ADR-0037, ADR-0043; slice-124, slice-221 |
| F-7 | behoben | Plattform-Sperre vor `mkdir -p "$CACHE"`; gemessen: kein Verzeichnis angelegt |
| F-8 | entfallen | `befunde()` gibt es nicht mehr; die Anzeige kommt aus den `FINDING`-Zeilen wie vor dem Umbau |

### Neue Findings

| ID | Kategorie | Befund | Quelle | Pfad | Verifizierbar | Klasse |
|---|---|---|---|---|---|---|
| D-1 | LOW | Der Selbsttest-Kommentar endet mit „ohne den Abgleich waere ein reines amd64-Bild als arm64 grün" — ein Konjunktiv über den abwesenden Zustand neben einer gültigen Abgrenzung („der Ausgang von architektur_ok, nicht nur der Extraktor"). Der Kommentar trägt damit eine Klasse, aber zusätzlich den Satzteil, den §3.7 als Falsch-Form nennt. Dasselbe Muster stand im ersten Wurf („Ohne sie waere der Anker …") und ist im ersten Lauf unbenannt geblieben. | `AGENTS.md` §3.7 | `tools/image-scan.sh:146-147` | nein — Leseurteil | Konjunktiv über verworfene Alternative im Kommentar |
| D-2 | INFO | Die OK-Zeile druckt `gescannt: ${plat#*/}`, also die verlangte Architektur; gemessen ist sie im Architektur-Lauf derselben Iteration, nicht im Entscheidungslauf, dessen Ergebnis die Zeile meldet. Die Sensor-Datei benennt das als Grenze 5 („Drei Läufe, ein Nachweis"; bei Tags wie dem Default `:latest` gilt sie). | Mess-Regel 5 | `tools/image-scan.sh:235` | nein | Ausgabe nennt abgeleiteten Wert als gemessen |
| D-3 | INFO | Die Bedingung `[ -n "${got}" ]` in `architektur_ok` überlebt ihre Entfernung im Selbsttest: bei nicht leerem `want` scheitert die Gleichheit ohnehin. Wirksam wäre sie nur bei einer Plattform ohne Architektur-Segment (`want` leer). Die Zusage „fehlende Angabe ist KEIN Erfolg" ist durch den Fall „Angabe fehlt" gedeckt. | Mess-Regel 4 (Randfall) | `tools/image-scan.sh:110` | ja — Mutation + `--selftest` | Redundante Bedingung ohne eigene Probe |

### Negativbefunde (Delta)

| Bereich | Ergebnis |
|---|---|
| Zähl-Pfad, Template, sieben Fixtures gegen `2236547~1` | unverändert, ohne Befund |
| Architektur-Lauf vor Vollbericht: Scheitern oder Abweichung → `errored=1`, `continue`, Marker „GESCHEITERT" | geprüft, ohne Befund |
| Exit-Semantik 0/1/2, Marker für `.github/workflows/image-scan.yml` | live geprüft, ohne Befund |
| Plan-Änderung `d01651b`: steht vor dem Fix-Commit, nennt die Änderung vor dem Code (`AGENTS.md` §6 Schritt 4); Commit-Scope `(planning)` nur `docs/plan/planning/` | geprüft, ohne Befund |
| Sensor-Datei: Grenze 4 (Namen in beiden Ausgaben) und neue Grenze 5 nennen Lauf und Grenze (Mess-Regel 5) | geprüft, ohne Befund |
| Gate-Index-Zelle (Vertrag unverändert, Bindung ergänzt) | geprüft, ohne Befund |
| CHANGELOG-Zeile gegen den neuen Aufbau | trifft weiter zu, ohne Befund |
| Hard Rules §3.1, §3.2, §3.5 (ADR-0037 nicht berührt), §3.6 | geprüft, ohne Befund |

### Summary (Delta)

| Kategorie | Anzahl |
|---|---|
| HIGH | 0 |
| MEDIUM | 0 |
| LOW | 1 |
| INFO | 2 |

**Finding-Klassen des Delta-Laufs:** Konjunktiv über verworfene Alternative im Kommentar ·
Ausgabe nennt abgeleiteten Wert als gemessen · Redundante Bedingung ohne eigene Probe

### Verdikt (Delta)

**Merge-blockierend:** nein. F-1 bis F-7 sind behoben, F-8 ist entfallen; jede Behebung habe ich
mit einer roten Mutation oder einem Live-Lauf nachgefahren. Offen bleiben ein LOW (D-1) und zwei
INFO; der Implementer nimmt sie an oder begründet, warum nicht. Die Closure-Zuordnung von F-1 zu
`BEO-GATE/umbau-verliert-pruefung-still` (der dritte Beleg) bleibt bestehen: Der Befund ist im
Vorgang slice-221 aufgetreten, auch wenn er behoben ist.
