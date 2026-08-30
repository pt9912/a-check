# slice-125 — Go-Toolchain auf 1.27: die neun Befunde beheben

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Risiko-Ausgang aus [slice-124 §6](../done/slice-124-image-scan-cve.md) — der erste Lauf
des CVE-Sensors fand **9 behebbare HIGH**, alle in der `stdlib` `v1.26.4`. Maintainer-Wort:
**Toolchain 1.27**.

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Das Binary wird mit einer Toolchain gebaut, deren `stdlib` die neun gemeldeten CVEs nicht mehr
trägt.

## 2. Definition of Done

- [x] [`Dockerfile`](../../../../Dockerfile) baut mit **1.27.0**: `GO_VERSION` und der
      Basis-Image-Digest sind **gemeinsam** gehoben — ein Tag ohne Digest wäre keine Hebung,
      sondern eine Lockerung.
- [x] Der Beleg ist **gemessen, nicht abgeleitet**: das lokal gebaute Image trägt `stdlib`
      `1.27.0` und **keinen** der neun Befunde. Geprüft mit demselben Trivy-Pin wie
      [`tools/image-scan.sh`](../../../../tools/image-scan.sh), gegen das exportierte Image —
      **ohne** Docker-Socket.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| [`Dockerfile`](../../../../Dockerfile) | update | `GO_VERSION` + Basis-Image-Digest |

**Auszuführende Gates:** `make gates` (tragend `test`, `coverage-gate`, `lint` — alle laufen unter
der neuen Toolchain) und `make image-test`. Zum Abschluss `make verify`.

### Was hier NICHT mitwandert, und warum

**Die `go.mod`-Direktive bleibt bei `1.26`.** Sie nennt die **Sprachversion**, nicht die
Toolchain — sie sagt „dieses Modul braucht mindestens 1.26". Die CVEs sitzen in der `stdlib` des
**gebauten Binaries**, und die kommt aus dem Build-Image. Die Direktive zu heben änderte an den
Befunden nichts und höbe nur die Mindestanforderung ohne Anlass.

### Das Risiko ist vorab gemessen

| Prüfung | Ergebnis |
|---|---|
| Basis-Image `1.27.0` verfügbar | ja — `1.27.0 linux/amd64` |
| Digest aus zwei Quellen | die Tags `1.27` und `1.27.0` liefern **dieselbe** Image-ID und denselben RepoDigest |
| Trivys Fix-Angabe zu allen neun | `1.27.0-rc.3` oder früher — `1.27.0` deckt sie |

## 4. Trigger

**Start:** eingetreten — die Befunde sind gemessen, die Toolchain ist verfügbar.

**Rückführungen:**

- `in-progress` → `open`: falls `make gates` unter der neuen Toolchain Befunde meldet, die eine
  Code-Änderung verlangen. Dann ist die Hebung nicht der Gegenstand, sondern ihr Anlass.

## 5. Closure-Trigger

Toolchain gehoben, Befundfreiheit am gebauten Image gemessen, Gates grün.

**Was bewusst nicht getan wird:** das **Release**. Der Sensor prüft das **publizierte** Image;
bis zum nächsten Release meldet er die neun weiter, und das ist richtig so — er misst, was
Anwender ziehen, nicht was im Baum liegt. Ein Slice kann die Zusage herstellen, veröffentlichen
kann sie nur ein Release.

## 6. Risiken und offene Punkte

- *Ein Major-Sprung von 1.26 auf 1.27 — Sprachverhalten oder Bibliotheken könnten sich ändern* —
  **Ausgang:** entfallen, gestrichen mit Begründung: `make gates` läuft unter der neuen Toolchain
  unverändert grün, Coverage **96,00 %**, `arch-check` 0 Befunde, `make image-test` grün. Keine
  Zeile Code war anzupassen. Die Fläche ist klein — **ein** direktes Modul, kein indirektes.
- *Der Sensor bleibt bis zum Release rot, und ein dauerhaft rotes Abzeichen wird weggeklickt* —
  **Ausgang:** weiter offen im **Beobachtungs-Register**. Der Zustand ist korrekt (das
  publizierte Image trägt die Befunde weiter), aber er ist von einem echten Versäumnis nicht zu
  unterscheiden. Das ist die Kehrseite eines Sensors, der den **publizierten** Stand misst.

## 7. Closure-Notiz

**Geliefert:** Die Toolchain steht auf `1.27.0`, `GO_VERSION` und Basis-Image-Digest gemeinsam
gehoben. Der Beleg ist gemessen: das neu gebaute Image trägt `stdlib=v1.27.0` und **0** behebbare
CRITICAL/HIGH — gegenüber `v1.26.4` mit **9** im publizierten Stand.

**Lerneintrag — Form: geschärfte Regel.** *Wer einem Sensor eine Grenze gibt, gibt sie auch dem
Beleg, der ihn bestätigen soll.*
[ADR-0037](../../adr/0037-cve-scan-gegen-das-publizierte-image.md) verbietet den Docker-Socket —
richtig, ein Werkzeug soll keinen Host-Root-Pfad bekommen, den es nicht braucht — und nennt die
Folge selbst: *„ein lokal gebautes Bild ist über dieses Skript nicht scanbar"*. Beim Schreiben der
ADR las sich das wie eine Randnotiz. **Einen Slice später war es der Kernweg:** dieser Slice
behebt Befunde, und sein lokal gebautes Image liegt in keiner Registry — `make image-scan` konnte
die Wirkung also gar nicht zeigen. Der Beleg brauchte einen anderen Weg (`docker save` plus
Trivys `--input`), der dieselbe Grenze respektiert statt sie aufzuweichen. *Weil* eine Grenze im
Sensor eine Grenze in jedem künftigen Beleg ist, gehört bei ihrer Formulierung die Frage dazu:
**womit zeigt man dann, dass eine Behebung wirkt?**

**Drei beobachtbare Closure-Kriterien:**

1. Vorher/nachher am selben Werkzeug und demselben Pin: **9 → 0** behebbare CRITICAL/HIGH,
   `stdlib` `v1.26.4` → `v1.27.0`. Beide Zahlen stammen aus Trivy `0.74.0` mit identischem
   Digest — ohne denselben Pin wäre der Vergleich wertlos.
2. Die Hebung ist **vollständig**: Tag **und** Digest. Ein gehobener Tag mit stehendem Digest
   zöge weiterhin das alte Bild und sähe dabei aktuell aus — dieselbe Klasse Fehler wie ein
   Digest, der den Vorgänger nennt ([ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md)).
3. Der Bestand blieb unberührt: keine Code-Änderung, `make gates` und `make image-test` grün.

**Ein Falschtreffer des eigenen Guard, beim Schreiben dieses Slice:** der PreToolUse-Guard hat das
Heredoc abgewiesen, mit dem der Plan entstehen sollte — er las die Zeichenfolge im **Text** als
Toolchain-Aufruf. Der Guard ist damit an derselben Stelle streng, an der er es sein soll (er
prüft Kommandos, nicht Absichten), aber er trifft auch Text, der über die Toolchain **schreibt**.
Ausgewichen mit einem Datei-Schreib-Werkzeug statt einer Shell-Umleitung.

**Offene Risiken und ihr Ausgang:** der erste entfallen (gestrichen mit Begründung), der zweite
weiter offen im Register.

**Beobachtungs-Register:** `BEO-027` neu angelegt (Durchsetzungsschicht, 1×, Beleg slice-125): der
Command-Guard trifft Text, der die Toolchain **nennt**, nicht nur Kommandos, die sie **aufrufen**.

**Folge-Slices:** das **Release**, das die Hebung veröffentlicht — bis dahin meldet der Sensor die
neun weiter, und das ist richtig. Danach der **Hebungs-Kanal** (Dependabot), der künftige Fälle
dieser Art meldet, ohne dass jemand hinsieht.
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Build-Schicht** (`Dockerfile`). Kein
Code, keine Doku-Verträge.

**Vorgelagert — offene Beobachtungen sichten:** keine Treffer in der berührten Sub-Area;
[`BEO-026`](../observations.md) (Workflow-`uses:`-Form) liegt in der CI-Schicht.

Alle berührten Sub-Areas GF.
