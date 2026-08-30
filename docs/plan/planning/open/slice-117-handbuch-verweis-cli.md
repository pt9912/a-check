# slice-117 — Handbuch-Verweis in `--print-mk` und `--help`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `git mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Auftrag 2026-08-30.
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
(das erzeugte Fragment),
[AC-FA-CLI-001](../../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)
(der Aufruf).

**Berührte Spec-Stellen:**
[SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) ·
[SPEC-DIST-001](../../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution)

**Verantwortlich:** — *(bis zur Priorisierung)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Wer `a-check` in der Hand hat, findet das Benutzerhandbuch, ohne das Repo zu kennen: eine
Verweis-Zeile im erzeugten Fragment und ein Aufruf-Kopf in `--help`.

## 2. Definition of Done

- [ ] Eine neue `AC-FA-CLI`-Kennung im [Lastenheft](../../../../spec/lastenheft.md) deckt die
      **Usage-Ausgabe**: Kurzbeschreibung, Aufruf-Syntax, Konfigurationsdatei, Handbuch-Verweis.
      Sie sichert die **Anwesenheit** dieser Bestandteile zu, **nicht den Wortlaut** — und trägt
      die vier Bausteine (Happy · Boundary · Negative · Out-of-Scope), weil
      `verify-ac-form` sie als **neue** Anforderung prüft.
- [ ] `a-check --help` gibt den Kopf aus; die Spezifikation ist an
      [SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)
      nachgezogen. Beleg: ein Test auf die Bestandteile, nicht auf den Text.
- [ ] Das `--print-mk`-Fragment **und** die committete
      [`a-check.mk`](../../../../a-check.mk) tragen die Handbuch-Zeile. Beleg: `make image-test`
      (Fragment-Parität, slice-034).

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

**Spec-first** ([`AGENTS.md`](../../../../AGENTS.md) §5): Lastenheft → Spezifikation → Code →
Tests. Die Reihenfolge ist hier nicht formal, sondern trägt: der Test kann erst wissen, was er
prüft, wenn die Anforderung sagt, was zugesichert ist (Anwesenheit) und was nicht (Wortlaut).

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| [`spec/lastenheft.md`](../../../../spec/lastenheft.md) | neu | die Anforderung; IDs entstehen **nur hier** |
| [`spec/spezifikation.md`](../../../../spec/spezifikation.md) §[SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) | update | die technische Fassung |
| `internal/cli/cli.go` | update | `fs.Usage` setzen; Handbuch-Zeile im `mkFragment` |
| [`a-check.mk`](../../../../a-check.mk) | update | committete Fassung, sonst bricht die Parität |
| `internal/cli/cli_test.go` | update | je Bestandteil eine Prüfung |
| [`docs/user/benutzerhandbuch.md`](../../../../docs/user/benutzerhandbuch.md) | update | Handbuch-Version + Änderungszeile |

**Auszuführende Gates:** `make gates` (tragend `test`, `coverage-gate`, `doc-check`) und
`make image-test` — Letzteres ist hier **nicht optional**: es ist das einzige Gate, das die
committete `a-check.mk` gegen `--print-mk` misst. Zum Abschluss `make verify`.

### Die URL-Form ist die eine Entscheidung dieses Slice

Der Link zeigt **tag-frei** auf `main`:
`https://github.com/pt9912/a-check/blob/main/docs/user/benutzerhandbuch.md`

**Nicht, weil es bequem ist, sondern weil das Binary seinen eigenen Release-Kontext nicht kennt.**
Der Build fährt `-ldflags="-s -w"` **ohne** `-X` — es gibt keine eingebackene Version. Ein
tag-stabiler Link müsste zur Build-Zeit entstehen und nähme, unverändert eingebacken, immer den
Stand des **Vorgänger**-Release: gültig aussehend und falsch. Das ist wortgleich die Lage, die
[ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md) für den Image-Digest entschieden
hat, und dieselbe, aus der `internal/cli/cli.go` bis slice-083 als vierte Pin-Stelle entfernt
wurde.

**Der Preis, benannt:** wer ein altes Release fährt, landet auf dem aktuellen Handbuch. Das ist
vertretbar, weil das Handbuch seine Software-Version selbst als Verweis auf
[`version.md#aktuell`](../../../../version.md#aktuell) führt statt als literale Nummer — der Leser
sieht dort, wogegen er liest.

## 4. Trigger

**Start:** eingetreten — Maintainer-Auftrag, keine Vorbedingung.

**Rückführungen:**

- `in-progress` → `next`: falls sich zeigt, dass der Usage-Kopf eine eigene ADR verlangt (§6) —
  dann sind es zwei Slices, nicht einer.

## 5. Closure-Trigger

Anforderung steht, `--help` und Fragment tragen den Verweis, `make image-test` und `make gates`
grün.

**Was bewusst nicht getan wird:** `verify-ac-form` **ablösen**. Dieser Slice legt zwar das
zwanzigste `AC-*` an und macht damit den alternativen Trigger aus
[slice-116 §4](../done/slice-116-nullmengen-haerte-cr.md) wahr — die Ablösung selbst ist
ein eigener Slice, sonst trägt dieser vier Liefer-Punkte statt drei.

## 6. Risiken und offene Punkte

- *Die URL-Form-Entscheidung könnte eine eigene ADR verlangen statt einer Plan-Passage* — sie ist
  eine **Anwendung** von [ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md), nicht
  eine neue Entscheidung; §3 nennt die Ableitung. **Ausgang:** <bei Closure>
- *Ein Test auf Prosa wird brüchig* — geprüft wird die **Anwesenheit** der Bestandteile (URL,
  Aufruf-Zeile, Konfigurations-Hinweis), nicht ihr Wortlaut; die Anforderung sagt das
  ausdrücklich. **Ausgang:** <bei Closure>
- *Die Handbuch-URL ist die erste Netz-Adresse im erzeugten Fragment* — sie wird nie abgerufen
  (das Fragment ist Text, das Modul `external` ist nicht aktiv), aber sie ist eine Zusage über
  ein fremdes System. **Ausgang:** <bei Closure>

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5.)_

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>**

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden die **Spec-Straten** (neue Anforderung,
nachgezogene Spezifikation) und die **Implementierung** (`internal/cli/`). Beide sind eigene
Sub-Areas mit eigener Konventions-Dichte; die Doku-Schicht (Handbuch) hängt an der ersten.

**Vorgelagert — offene Beobachtungen sichten:** keine Treffer in den berührten Sub-Areas —
`BEO-002` (Schichten-Zähler) und `BEO-022` (CR-Texte) liegen anderswo.

Alle berührten Sub-Areas GF.
