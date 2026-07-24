# slice-043 — Schicht-Abdeckung sichtbar machen (P-Rest-Kandidat 2a)

**Status:** open — **Entwurf zur Abnahme** (spec-first; noch keine Spec-/Code-Änderung).
**Auslöser:** Maintainer-Nachfrage vom 2026-07-24 zu
[slice-042 §8](slice-042-constructs-aufruf-monopol.md) („warum können andere Konsumenten das
nicht gebrauchen?"). Die anschließende Nachmessung über **alle fünf** lokalen
`.a-check.yml`-Konsumenten zeigt eine Fehlerklasse, die mit dem b-cad-P-Rest nichts zu tun hat.
**Bezug:** macht die bewusste fail-open-Grenze aus
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
sichtbar (sie wird **nicht** verschoben); Extraktions-/Auflösungs-Vertrag
[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion);
Konfigurations-Vertrag
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml);
Determinismus [AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus).
[Roadmap](../in-progress/roadmap.md).

> **Hinweis:** Entwurf zur Abnahme. Es werden hier **keine** `AC-*`/`ADR-*`-IDs vergeben
> (Anlege-Prozess:
> [conventions §Anforderungs-Anlege-Prozess](../../../../harness/conventions.md#anforderungs-anlege-prozess)).
> Entscheide §3 **vor** der Umsetzung.

---

## 1. Die Fehlerklasse

Ein Konsument, dessen `languages`-Globs **weiter** reichen als seine `layers`-Globs, hat
gescannte Dateien in **keiner** Schicht. Das ist zulässig und oft gewollt — aber es erzeugt zwei
Symptome, die der Nutzer heute nicht als *eine* Ursache erkennt (Proben aus
[slice-042 §8.1](slice-042-constructs-aufruf-monopol.md)):

| Seite | Verhalten heute | Bewertung |
|---|---|---|
| **Ziel-Seite** — ein Import löst sich auf **keine** Schicht auf | 0 Befunde, still | **fail-open** — dokumentierte Grenze, aber unsichtbar |
| **Quell-Seite** — eine schichtlose Datei importiert eine Schicht | `wrong-direction`, aber mit **leerem** Quell-Schicht-Namen (`wrong-direction:  -> ui`) | korrekt, aber kryptisch |
| Gegenprobe: Zone als Schicht deklariert | regulärer Befund (`core-impurity`) | zwischen stiller Lücke und gefundenem Verstoß steht **nur die Deklaration** |

**Kernpunkt:** Der Konsument kann das heute selbst reparieren — eine Zeile in `layers` (oder in
`exclude`, wenn die Zone nicht dazugehört). Er erfährt nur nirgends, dass die Lücke existiert.
Ein grünes Gate über einem teilweise ungeprüften Baum ist genau die Klasse „stille
Setzung", gegen die dieses Repo sonst fail-closed vorgeht.

## 2. Evidenz (Nachmessung 2026-07-24, alle lokalen Konsumenten)

| Konsument | gescannt, aber in keinem Layer-Glob |
|---|---|
| a-check (Dogfooding) | **0** — Baum vollständig gedeckt |
| b-cad | **0** — Baum vollständig gedeckt |
| d-check | **0** — Baum vollständig gedeckt |
| d-migrate | der gesamte `test/`-Kotlin-Baum (kein `exclude`, `languages: **/*.kt`) |
| m-trace | `apps/api/internal/storage/**`, `apps/api/scripts/coverage-overview/**` |

Der **m-trace-Fall** trägt den Slice: `apps/api/internal/storage` wird von der dortigen Config in
einer `tech`-Regel ausdrücklich als Architektur-Zone geführt (zulässiger `database/sql`-Halter),
hat aber keine Schicht. Ein Import aus `hexagon/domain` dorthin wäre exakt der Verstoß, den das
Gate finden soll — und bliebe still. **Heute latent** (nur `*_test.go` importieren die Zone, und
die sind dort per `exclude` aus dem Scan), also kein akuter Fehlbefund, aber eine unbemerkte
Lücke im laufenden Gate eines realen Konsumenten.

## 3. Zu entscheiden vor der Umsetzung (ADR-Skizze)

| Frage | Optionen | Erste Neigung |
|---|---|---|
| **Gestalt** | (a) immer eine Diagnose-Zeile auf stderr · (b) Opt-in `strict_coverage: true` ⇒ Exit 1 · (c) eigenes Flag (`--print-coverage`) · (d) (a)+(b) | **(d)**: die Diagnose sieht jeder sofort (das ist der Zweck), die Strenge bleibt Konsumenten-Wahl. Ein Default-Exit-1 würde bestehende grüne Gates brechen — Verschärfung ohne ADR wäre unzulässig ([AGENTS §3.6](../../../../AGENTS.md#36-gates-dürfen-nicht-ohne-adr-gelockert-werden) sinngemäß in der Gegenrichtung) |
| **Granularität** | je Datei · je Verzeichnis-Zone · Zähler + Zonen | **Zähler + Zonen** — d-migrates `test/`-Baum sind hunderte Dateien; eine Datei-Liste macht die Meldung unlesbar und die Diagnose wertlos |
| **Zonen-Bildung** | längster gemeinsamer Verzeichnis-Präfix · erste N Segmente · Verzeichnis der Datei | offen — muss **deterministisch und stabil sortiert** sein ([AC-QA-01](../../../../spec/lastenheft.md#ac-qa-01--determinismus), [SPEC-DET-001](../../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag)) |
| **Abgrenzung** | `composition_root`-Dateien sind **per Definition** schichtlos; `exclude`-Dateien sind gar nicht im Scan ([ADR-0018](../../adr/0018-exclude-scan-scope.md)/[ADR-0025](../../adr/0025-exclude-verzeichnis-prune.md)); repo-**externe** Import-Ziele (Fremdbibliotheken) sind kein Befund | beide ersten Klassen **ausnehmen**; die dritte ist der Grund, warum die **Ziel**-Seite nicht einfach fail-closed werden kann — extern und „schichtlos im eigenen Baum" sind auf Symbol-Ebene nicht sicher unterscheidbar |
| **Wo greift die Strenge** | Quell-Seite (gescannte Datei ohne Schicht) · Ziel-Seite (Import ohne Schicht) | **Quell-Seite** — sie ist die *Ursache* und ohne Fehlalarm-Risiko entscheidbar. Die Ziel-Seite ist die Symptom-Seite und mit Externem vermischt; sie wird über die Ursache mit-repariert (Gegenprobe §1) |
| **Verhältnis zu [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)** | — | die Grenze wird **nicht verschoben**, sondern *ausgewiesen* — das ist wörtlich der Vertrag („die Grenze wird ausgewiesen, nicht als Vollständigkeit ausgegeben") |

## 4. Anforderungs-Skizze (Bereich offen — `CONF` oder `CLI`, Entscheid im CR)

- **Happy:** Given ein Baum mit gescannten Dateien außerhalb aller Layer-Globs, when a-check
  läuft, then nennt die Ausgabe Anzahl und Zonen dieser Dateien — deterministisch sortiert,
  zusätzlich zu (nicht anstelle von) den regulären Befunden; der Exit-Code bleibt unverändert.
- **Boundary:** Given ein vollständig gedeckter Baum (a-check/b-cad/d-check), when a-check läuft,
  then **keine** Diagnose-Zeile (kein Rauschen im Normalfall). Given `composition_root`-Dateien,
  then zählen sie **nicht** als ungedeckt. Given `strict_coverage: true` und eine ungedeckte
  Zone, then Exit 1.
- **Negative:** Given `strict_coverage` mit einem Nicht-Bool-Wert, when a-check lädt, then Exit 2
  (fail-closed, Muster der bestehenden Config-Härtung).
- **Out-of-Scope:** keine Auto-Inferenz von Schichten aus Verzeichnisnamen; keine Strenge auf der
  **Ziel**-Seite (Externes ist nicht sicher abgrenzbar, §3); keine Änderung der
  Auflösungs-Semantik ([SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion));
  keine Import-Allowlist (das ist Kandidat 2b, weiter gated —
  [slice-042 §8.2](slice-042-constructs-aufruf-monopol.md)).

## 5. Mitzunehmen: leerer Quell-Schicht-Name

Der in [slice-042 §9](slice-042-constructs-aufruf-monopol.md) notierte Mikro-CR
(`wrong-direction:  -> ui` — Loch an der Stelle des Quell-Schicht-Namens) ist das **Quell-seitige
Symptom derselben Klasse** und gehört hierher: die Meldung soll die Datei als schichtlos
ausweisen statt einen leeren Namen zu rendern. Klein, aber nutzersichtbar; Empfehlung: in diesem
Slice mitnehmen, damit beide Symptome dieselbe Sprache sprechen.

## 6. Betroffene Module und DoD

| Modul | Änderung |
|---|---|
| `internal/adapter/driven/config` | optionaler Schlüssel (`strict_coverage`), fail-closed-Validierung |
| `internal/hexagon/core` | Abdeckungs-Ermittlung + Zonen-Bildung, stabil sortiert; Meldungstext für schichtlose Quelldateien (§5) |
| `internal/cli` | Diagnose-Ausgabe, Exit-Code-Logik ([SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)) |
| Doku | Lastenheft-CR, ADR, [SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema)/[SPEC-DET-001](../../../../spec/spezifikation.md#spec-det-001--determinismus-vertrag), Benutzerhandbuch, [CHANGELOG](../../../../CHANGELOG.md), Roadmap |

**DoD:** Spec-first (Lastenheft-CR → ADR → Spezifikation → Code → Tests); `make gates` **und**
`make ci` grün mit echter Ausgabe; **Regressions-Probe gegen alle fünf lokalen Konsumenten**:
a-check/b-cad/d-check bleiben diagnose-**frei** (sonst ist die Meldung Rauschen), m-trace und
d-migrate melden genau die Zonen aus §2; Gegenprobe: Zone deklarieren ⇒ Diagnose verschwindet und
ein injizierter Verstoß in der Zone wird gefunden; Benutzerhandbuch-Currency. Die Konsumenten
werden **informiert**, nicht umkonfiguriert — die Config-Entscheidung (Schicht vs. `exclude`)
liegt in ihren Repos.
