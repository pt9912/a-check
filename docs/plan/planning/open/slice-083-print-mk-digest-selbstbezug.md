# slice-083 — `--print-mk` nennt den Digest des Vorgängers

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** Maintainer-Befund vom 2026-08-09, bestätigt durch einen **realen Fehlpin** im Einsatz;
[AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit),
[AC-FA-DIST-001](../../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk).
**Bezug:** [ADR-0004](../../adr/0004-distribution-image-mk.md),
[ADR-0007](../../adr/0007-latest-tag-politik.md); Geschwister
[slice-082](../done/slice-082-print-mk-docker-indirektion.md).

---

## 0. Trigger

**Beginn:** nachdem [slice-082](../done/slice-082-print-mk-docker-indirektion.md) in `done/` liegt — beide
ändern dieselben zwei Dateien, und die Fragment-Parität erzwingt gemeinsames Wandern.

**Rückführungen:**

- `in-progress` → `open`: falls der Entscheid aus §3 eine Lastenheft-Änderung verlangt, die der
  Maintainer nicht trägt. Dann bleibt der Defekt bestehen und die **Doku** muss ehrlich werden —
  das wäre ein anderer, kleinerer Slice.

## 1. Auslöser

**Mechanismus: eine Ausgabe sieht autoritativ aus und kann es strukturell nicht sein.**

Der Digest steht als Konstante in
[`internal/cli/cli.go:123`](../../../../internal/cli/cli.go). Er entsteht aber erst **beim Push** —
das Binary ist vorher gebaut. Jedes publizierte Image trägt deshalb den Digest seines
**Vorgängers**. Gemessen am 2026-08-09:

| Tag | einkompiliert | echter Digest |
|---|---|---|
| `v0.16.0` | `6425c93a…` | `aef28cfe…` |
| `v0.15.0` | `f1b8ff5e…` | `6425c93a…` |

Unabhängig bestätigt am lokalen Image-Cache: `ghcr.io/pt9912/a-check:v0.15.0` trägt
`sha256:6425c93a…` — genau den Wert, den `cli.go` zum Tag `v0.16.0` führte.

**Das ist kein Bug zum Wegfixen.** Ein Image, das seinen eigenen Digest enthält, wäre ein
SHA-256-Fixpunkt. Keine Umsortierung des Release-Prozesses löst das;
[`docs/user/releasing.md`](../../../user/releasing.md) Schritt 6 (Re-Pin nach Publish) ist bereits
die einzig mögliche Reihenfolge.

**Der Defekt ist die Autorität der Ausgabe**, nicht der Wert. Und er hat Schaden angerichtet: der
in [`d-check.mk:14`](../../../../d-check.mk) dokumentierte Bump-Weg („Image ziehen, `--print-mk`
laufen lassen") hat im realen Einsatz einen **Fehlpin** erzeugt — der Konsument pinnte dauerhaft
eine Version zurück, ohne es zu merken, weil er ja einen gültigen Digest bekam.

**Zwei Doku-Stellen behaupten das Unmögliche:**
[`releasing.md`](../../../user/releasing.md) Schritt 5 („gibt … den **aktuell** digest-gepinnten
`A_CHECK_IMAGE` aus") und Freigabe-Item 6 („`--print-mk` gibt den neuen Digest", **mit
Beleg-Slot**). Item 6 ist nicht erfüllbar.

**Kein Gate fängt es:** `gate-consistency` prüft nur den Repo-Stand;
[`tools/image-test.sh`](../../../../tools/image-test.sh) prüft Fragment-Parität gegen das **lokal
gebaute** Image, wo die Konstante per Konstruktion stimmt. Das publizierte Image wird nie geprüft.

## 2. Betroffene Module

Zwei Schichten:

1. **Spec/Doku** — [`spec/lastenheft.md`](../../../../spec/lastenheft.md)
   ([AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) verlangt heute
   ausdrücklich einen `@sha256:`-Digest im gelieferten Fragment), ADR,
   [`docs/user/releasing.md`](../../../user/releasing.md) §5 und Freigabe-Item 6.
2. **Distribution** — [`internal/cli/cli.go`](../../../../internal/cli/cli.go) und
   [`a-check.mk`](../../../../a-check.mk).

## 3. Auszuführende Gates

`make image-test`, `make gates`.

**Der Entscheid, der vor dem Bau fällt.** Drei Ausgänge, keiner vorweggenommen:

| Ausgang | Kosten |
|---|---|
| **Tag statt Digest** — `…:v0.16.0`; die Version ist beim Build bekannt (`ARG VERSION`, OCI-Label) | selbstkonsistent und immer korrekt, **verliert die Digest-Härte** aus [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) |
| **Digest zur Laufzeit auflösen** — der Host kennt ihn (`docker inspect --format '{{index .RepoDigests 0}}'`, wie `release.yml:130`); ein `a-check-pin`-Target im Fragment gibt ihn aus | netzlos, deterministisch; verschiebt eine Handlung zum Konsumenten |
| **Nur die Doku korrigieren** — Ausgabe bleibt, aber deklariert sich als „Stand beim Build" | billigster Schnitt, der Fehlpin-Weg bleibt aber gangbar |

Die ersten beiden berühren
[AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) im Wortlaut und sind damit
**spec-first**: Lastenheft-CR → ADR → Spezifikation → Code.

**Nachtrag 2026-08-09 — `CR-5` schärft das Akzeptanzkriterium.** Der Konsument hat den Befund als
formalen Change Request nachgereicht und formuliert die Bedingung schärfer, als der Slice sie
hatte:

> Die Ausgabe enthält **keinen Digest, der auf ein anderes Release zeigt** als das laufende Binary.

Das schließt Ausgang 3 („nur die Doku korrigieren") **aus**: ein korrekt kommentierter, aber
falscher Digest erfüllt es nicht. Übrig bleiben der bewegliche Tag und die Laufzeit-Auflösung — der
CR nennt als dritte Variante einen **Platzhalter** (`@sha256:<release-digest>`), der beim Einsetzen
auffällt.

Dazu die Schadenshöhe, die der CR beziffert: der Fehlpin führte zu `v0.15.0` statt `v0.16.0` — und
damit fehlte dem Konsumenten der `constructs`-Block, den sein Skelett brauchte. Der Defekt kostet
also nicht nur einen falschen Pin, sondern eine **fehlende Fähigkeit**, deren Abwesenheit als
Konfigurationsfehler erscheint.

**Negativ-Proben** — je nach Ausgang, aber diese gelten immer:

| Probe | Erwartung |
|---|---|
| Fragment-Ausgabe des laufenden Binaries | enthält **keinen** Digest eines anderen Release (`CR-5`-AK) |
| `releasing.md` Freigabe-Item 6 | nennt einen **erfüllbaren** Beleg — oder das Item entfällt |
| Fragment-Ausgabe | behauptet keine Aktualität, die sie nicht prüfen kann |

## 4. Was bewusst nicht getan wird

- **Den Fixpunkt „lösen".** Er ist nicht lösbar; jeder Ausgang oben umgeht ihn, keiner hebt ihn auf.
- **Den Release-Prozess umsortieren.** Schritt 6 ist bereits die einzig mögliche Reihenfolge.
- **`$(DOCKER)`** — [slice-082](../done/slice-082-print-mk-docker-indirektion.md), läuft vorher.

## 5. DoD

- [ ] Der Entscheid aus §3 ist **getroffen und begründet** in der Closure-Notiz; berührt er
      [AC-QA-03](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit), liegt die Lastenheft-Änderung mit ADR vor, **bevor** Code entsteht.
- [ ] Keine Doku-Stelle behauptet mehr, `--print-mk` nenne den aktuellen Digest. Beleg:
      [`releasing.md`](../../../user/releasing.md) §5 und Freigabe-Item 6, gegen den gewählten
      Ausgang gelesen.
- [ ] `make gates` und `make image-test` grün — **Ausgabe in eine Datei**, Exit-Code getrennt
      geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
