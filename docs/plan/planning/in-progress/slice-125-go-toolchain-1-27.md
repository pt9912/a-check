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

- [ ] [`Dockerfile`](../../../../Dockerfile) baut mit **1.27.0**: `GO_VERSION` und der
      Basis-Image-Digest sind **gemeinsam** gehoben — ein Tag ohne Digest wäre keine Hebung,
      sondern eine Lockerung.
- [ ] Der Beleg ist **gemessen, nicht abgeleitet**: das lokal gebaute Image trägt `stdlib`
      `1.27.0` und **keinen** der neun Befunde. Geprüft mit demselben Trivy-Pin wie
      [`tools/image-scan.sh`](../../../../tools/image-scan.sh), gegen das exportierte Image —
      **ohne** Docker-Socket.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

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
  **Ausgang:** <bei Closure>
- *Der Sensor bleibt bis zum Release rot, und ein dauerhaft rotes Abzeichen wird weggeklickt* —
  **Ausgang:** <bei Closure>

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5.)_

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>**

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Build-Schicht** (`Dockerfile`). Kein
Code, keine Doku-Verträge.

**Vorgelagert — offene Beobachtungen sichten:** keine Treffer in der berührten Sub-Area;
[`BEO-026`](../observations.md) (Workflow-`uses:`-Form) liegt in der CI-Schicht.

Alle berührten Sub-Areas GF.
