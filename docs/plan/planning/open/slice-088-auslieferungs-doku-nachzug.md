# slice-088 — Auslieferungs-Doku nachziehen (Release-Blocker für v0.17.0)

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** den Doku-Nachzug von [slice-082](../done/slice-082-print-mk-docker-indirektion.md),
[slice-083](../done/slice-083-print-mk-digest-selbstbezug.md),
[slice-084](../done/slice-084-handbuch-heuristik-grenzen.md),
[slice-085](../done/slice-085-schicht-ohne-aufloesung.md) und
[slice-086](../done/slice-086-forbidden-constructs-fail-closed.md).
**Bezug:** blockiert das Release `v0.17.0`; verwandt mit
[`SL-005`](../../steering-loop.md#sl-005--eine-neue-datei-wird-nicht-in-ihren-handgepflegten-index-eingetragen)
(handgepflegter Ort ohne Sensor).

---

## 0. Trigger

**Beginn: sofort.** Der Slice ist ein **Release-Blocker**: `v0.17.0` würde eine Anleitung
ausliefern, die den Konsumenten in genau den Fehler führt, den zwei Slices dieser Welle gerade
behoben haben.

**Rückführungen:**

- `in-progress` → `open`: falls sich beim Schreiben zeigt, dass der Pin-Weg selbst noch offen ist
  (etwa weil die Release-Prozedur den Digest anders liefern soll) — dann ist das ein
  Vertrags-Entscheid und keine Doku-Arbeit.

## 1. Auslöser

**Mechanismus: der Doku-Nachzug ist angeordnet, aber von keinem Sensor gedeckt.** Schritt 7 des
Workflow-Skeletts ([`.claude/commands/slice.md`](../../../../.claude/commands/slice.md)) verlangt
ihn *„falls ein öffentlicher Vertrag berührt ist (Lastenheft, Spezifikation, Benutzerhandbuch,
ADR-Index, CHANGELOG)"*. Der ADR-Index hat seit
[slice-087](../done/slice-087-index-vollstaendigkeit.md) einen Sensor — die übrigen vier Orte
nicht. **Fünf Slices haben ihn übersprungen, alle bei grünem `make gates`.**

**Der schwerste Fall — die Anleitung führt in den Fehler, den sie beschreiben soll.**
[Benutzerhandbuch §3.3](../../../user/benutzerhandbuch.md) beschreibt:

```bash
docker run --rm <a-check-image> --print-mk > a-check.mk   # Schritt 1
make a-check                                              # Schritt 3
```

Seit [slice-083](../done/slice-083-print-mk-digest-selbstbezug.md) fehlt dazwischen der **zwingende**
Schritt „Digest eintragen"; das Fragment trägt bewusst einen Platzhalter und bricht ab. Der Hinweis
darunter ist zusätzlich **falsch**: *„Das Fragment pinnt das veröffentlichte Image über
`A_CHECK_IMAGE` (`@sha256:`-Digest)."* Genau dieser Anleitung ist der Konsument gefolgt, dessen
Fehlpin die Welle ausgelöst hat.

**Der Bestand, gemessen am 2026-08-09:**

| Ort | Fehlt | Aus |
|---|---|---|
| Handbuch §3.3 | Pin-Schritt + falscher „pinnt"-Satz | [slice-083](../done/slice-083-print-mk-digest-selbstbezug.md) |
| Handbuch | `DOCKER`-Indirektion **komplett** (null Treffer), inkl. der Reihenfolge-Falle `?=` vor/nach `include` | [slice-082](../done/slice-082-print-mk-docker-indirektion.md) |
| Handbuch §10 | vier Historien-Zeilen; die Historie endet bei `1.35` (2026-07-25), obwohl der Inhalt geändert wurde | 081/084/085/086 |
| Handbuch §4 | Exit-2-Fälle von `forbidden_constructs` nur im Glossar und als Zeilenkommentar — bei `constructs` stehen sie im Fließtext | [slice-086](../done/slice-086-forbidden-constructs-fail-closed.md) |
| [`CHANGELOG.md`](../../../../CHANGELOG.md) | drei Einträge (082, 083, 084) | 082/083/084 |
| [`version.md`](../../../../version.md) | nennt `internal/cli/cli.go` noch als harte Pin-Stelle | [slice-083](../done/slice-083-print-mk-digest-selbstbezug.md) |

**Das ist eine Klasse, keine sechs Einzelfälle:** jedes Mal wurde Code oder Vertrag geändert und der
begleitende Doku-Ort nicht mitgeführt. Der `cli.go`-Drift trat sogar **zweimal** auf — in
`harness/README.md` (dort in [slice-087](../done/slice-087-index-vollstaendigkeit.md) beiläufig
korrigiert) und in [`version.md`](../../../../version.md).

## 2. Betroffene Module

Reine Doku, kein Code, kein Vertrag:
[`docs/user/benutzerhandbuch.md`](../../../user/benutzerhandbuch.md),
[`CHANGELOG.md`](../../../../CHANGELOG.md), [`version.md`](../../../../version.md).

## 3. Auszuführende Gates

`make gates`.

**Belege je Behebung:**

| Probe | Erwartung |
|---|---|
| §3.3 Schritt für Schritt nachvollzogen | die Anleitung führt zu einem **laufenden** `make a-check`, ohne dass der Leser raten muss |
| `grep -c DOCKER` im Handbuch | **> 0**, inklusive der Reihenfolge-Regel (`?=` greift nur vor dem `include`) |
| Historie §10 | vier neue Zeilen, lückenlos an `1.35` anschließend |
| `[Unreleased]` im CHANGELOG | nennt **alle** Änderungen der Welle, nicht drei von sechs |
| `make gates` | grün (`gate-consistency` prüft die Pin-Konsistenz mit) |

**Ausdrücklich kein Sensor in diesem Slice.** Die naheliegende Antwort auf „fünf Slices haben den
Nachzug übersprungen" wäre ein Vollständigkeits-Sensor über CHANGELOG und Handbuch-Historie. Der
gehört aber erst geschnitten, wenn die Regel steht, *wann* ein Eintrag Pflicht ist — sonst erfindet
der Sensor sie. Dieselbe Reihenfolge wie bei
[`SL-003`](../../steering-loop.md#sl-003--commit-betreff-bezeichnet-nicht-die-enthaltene-arbeit):
erst die Konvention, dann ihr Sensor. Als Folge-Slice benannt, nicht mitgebaut.

## 4. Was bewusst nicht getan wird

- **Das Release selbst.** Tag, Push und Digest-Pin sind Maintainer-Sache
  ([`releasing.md`](../../../user/releasing.md)); dieser Slice macht `v0.17.0` nur *release-fähig*.
- **Die Versions-Nummer festlegen.** `v0.17.0` ist die begründete Empfehlung (SemVer `0.x`, zwei
  Features plus Breaking Change, sechzehn Minor-Releases ohne Patch-Präzedenz) — der Entscheid
  bleibt beim Maintainer, ebenso die Frage, ob ein Breaking Change der Anlass für `1.0.0` ist.
- **Den Handbuch-Aufbau ändern.** Nachgezogen wird, was fehlt; die Gliederung bleibt.

## 5. DoD

- [ ] Die Anleitung in §3.3 ist **durchgespielt** und führt zu einem laufenden Gate — inklusive
      Digest-Schritt und `DOCKER`-Indirektion mit ihrer Reihenfolge-Regel.
- [ ] Handbuch-Historie, `CHANGELOG` und `version.md` nennen den vollständigen Stand der Welle.
      Beleg: die fünf Proben aus §3.
- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
