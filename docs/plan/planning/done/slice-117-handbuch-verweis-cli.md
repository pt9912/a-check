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

- [x] Eine neue `AC-FA-CLI`-Kennung im [Lastenheft](../../../../spec/lastenheft.md) deckt die
      **Usage-Ausgabe**: Kurzbeschreibung, Aufruf-Syntax, Konfigurationsdatei, Handbuch-Verweis.
      Sie sichert die **Anwesenheit** dieser Bestandteile zu, **nicht den Wortlaut** — und trägt
      die vier Bausteine (Happy · Boundary · Negative · Out-of-Scope), weil
      `verify-ac-form` sie als **neue** Anforderung prüft.
- [x] `a-check --help` gibt den Kopf aus; die Spezifikation ist an
      [SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)
      nachgezogen. Beleg: ein Test auf die Bestandteile, nicht auf den Text.
- [x] Das `--print-mk`-Fragment **und** die committete
      [`a-check.mk`](../../../../a-check.mk) tragen die Handbuch-Zeile. Beleg: `make image-test`
      (Fragment-Parität, slice-034).

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

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

- *Die URL-Form-Entscheidung könnte eine eigene ADR verlangen statt einer Plan-Passage* —
  **Ausgang:** gestrichen mit Begründung, und die trägt jetzt doppelt. Erstens ist sie eine
  **Anwendung** von [ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md), nicht eine
  neue Entscheidung. Zweitens hat das Schwester-Tool dieselbe Aufgabe unabhängig gelöst und
  **dieselbe** Form gewählt — tag-frei, mit Verweis auf den Software-Versions-Stempel im
  Handbuch-Kopf. Zwei Repos, dieselbe Zwangslage, dieselbe Antwort: das ist kein Ermessen, das
  eine ADR bräuchte.
- *Ein Test auf Prosa wird brüchig* — **Ausgang:** gestrichen mit Begründung: die geprüften
  **Marken** stehen in der Spezifikation
  ([SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)),
  nicht im Test. Der Test beruft sich darauf, statt sie zu erfinden — genau der Fehler, den
  [slice-120](../done/slice-120-ac-form-abloesen.md) an `verify-ac-form` gemessen hat.
- *Die Handbuch-URL ist die erste Netz-Adresse im erzeugten Fragment* — **Ausgang:** weiter offen
  im **Beobachtungs-Register** als Teil von `BEO-021`-Nachbarschaft: sie wird nie abgerufen (das
  Werkzeug ist hermetisch, das Modul `external` ist nicht aktiv), bleibt aber eine Zusage über ein
  fremdes System, das sich ändern kann.

## 7. Closure-Notiz

**Geliefert:** [AC-FA-CLI-003](../../../../spec/lastenheft.md#ac-fa-cli-003--usage-ausgabe-und-handbuch-verweis) — `a-check --help` trägt eine vollständige Usage-Ausgabe — Kurzbeschreibung,
Aufruf-Syntax mit Pfad-Parameter, Options-Liste, Konfigurations-Hinweis und die Handbuch-URL;
dieselbe URL steht im Kopfkommentar des `--print-mk`-Fragments. **Eine** Konstante im Code speist
beide. Vorher gab `--help` nur Go's Default aus: `Usage of a-check:` plus drei Flags.

**Lerneintrag — Form: geschärfte Regel.** *Wo eine Anforderung nur die **Anwesenheit** zusichert
und nicht den Wortlaut, gehören die geprüften **Marken** in die Spezifikation — nicht in den Test.*
Sonst erfindet der Test seine eigene Erwartung, und die Anforderung deckt sie nicht.
[slice-120](../done/slice-120-ac-form-abloesen.md) hat genau diesen Fall gemessen: `verify-ac-form`
suchte `^**Happy Path:**`, ein Wortlaut, den **niemand** deklariert hatte und den die kanonische
Quelle nicht kennt — der Sensor war seine eigene Autorität. Hier stehen die vier Marken
(`Aufruf:`, `[pfad]`, `.a-check.yml`, `Benutzerhandbuch`) darum in
[SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes),
und drei Tests plus `image-test` berufen sich darauf. *Weil* die Marken deklariert sind, darf die
Prosa dazwischen sich ändern, ohne ein Gate zu brechen — was die Anforderung ausdrücklich erlaubt.

**Drei beobachtbare Closure-Kriterien:**

1. Die neue Anforderung ist die **zwanzigste** und damit der **erste echte Gegenstand** der Regel
   aus [slice-120](../done/slice-120-ac-form-abloesen.md). Sie greift, gegengeprobt: die
   `Boundary`-Marke entfernt ⇒ `section-marker-missing` auf der Überschriftszeile der neuen `AC-*`;
   zurückgebaut ⇒ grün. Der abgelöste Sensor hätte dieselbe Anforderung in Bestandsform beanstandet.
2. Die Zusage „eine Stelle im Code, zwei Ausgaben" ist **beidseitig** geprüft — ein Go-Test und
   eine `image-test`-Stufe vergleichen `--help` und `--print-mk` gegen **dieselbe** Zeichenkette.
   Eine einseitige Probe hätte ein Auseinanderlaufen nicht gesehen.
3. Der Nutzungsfehler bleibt ein Fehler: `--bogus` liefert Exit **2** *und* die Usage. Ohne diese
   Probe wäre eine Usage-Ausgabe, die immer Exit 0 liefert, von einer korrekten nicht zu
   unterscheiden.

**Beobachtbare Architektur-Aussage:** Dieselbe Aufgabe hat das Schwester-Tool in `v0.69.0`
unabhängig gelöst — und **dieselbe** URL-Form gewählt, mit derselben Begründung (das Binary kennt
seinen Release-Kontext nicht) und demselben benannten Preis (wer ein altes Release fährt, liest den
Software-Versions-Stempel im Handbuch-Kopf). Die Entscheidung war damit nicht Geschmack, sondern
von der Zwangslage bestimmt: ein Build ohne eingebackene Version kann keinen versionierten Link
setzen, der nicht den Vorgänger nennt.

**Offene Risiken und ihr Ausgang:** die ersten beiden gestrichen mit Begründung, der dritte weiter
offen im Register.

**Beobachtungs-Register:** keine neue Beobachtung. `BEO-023` (ein Prüfer mit leerer Prüfmenge
bleibt unkalibriert) hat mit dieser Anforderung seinen **ersten Gegenstand** bekommen; der Stand
nennt jetzt den Beleg, dass die Nachfolge-Regel greift.

**Folge-Slices:** keiner offen. Die Roadmap führt danach nur noch die vertagten
[slice-013](../open/slice-013-driving-driven-vertiefung.md) und
[slice-045](../open/slice-045-intern-extern-dateimenge.md).
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt werden die **Spec-Straten** (neue Anforderung,
nachgezogene Spezifikation) und die **Implementierung** (`internal/cli/`). Beide sind eigene
Sub-Areas mit eigener Konventions-Dichte; die Doku-Schicht (Handbuch) hängt an der ersten.

**Vorgelagert — offene Beobachtungen sichten:** keine Treffer in den berührten Sub-Areas —
`BEO-002` (Schichten-Zähler) und `BEO-022` (CR-Texte) liegen anderswo.

Alle berührten Sub-Areas GF.
