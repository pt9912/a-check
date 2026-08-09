# slice-081 — Laufzeit-Diagnose für nicht extrahierte Import-Formen

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** Konsumenten-Befund vom 2026-08-09 (realer Einsatz in einem Fremd-Repo);
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
**Bezug:** Vorbild ist die Abdeckungs-Diagnose aus
[slice-043](../done/slice-043-schicht-abdeckung-sichtbar.md) /
[ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md).

---

## 0. Trigger

**Beginn:** sofort. Der Befund stammt aus einem realen Einsatz und wartet auf nichts.

**Rückführungen:**

- `in-progress` → `open`: falls die Messung zeigt, dass „nicht extrahiert" ohne AST nicht
  zuverlässig von „keine Import-Zeile" zu trennen ist — dann wäre die Diagnose selbst eine
  Heuristik über einer Heuristik und braucht erst einen Entscheid.

## 1. Auslöser

**Mechanismus: das Werkzeug kennt seine Blindstellen, sagt sie aber nicht am Repo.**

[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
verspricht, die Heuristik-Grenzen **offenzulegen statt zu verschweigen**. Eingelöst wird das heute
in der Doku — der Out-of-Scope-Absatz von
[AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion) listet
über zwanzig Grenzen. Am **Repo** sagt a-check davon nichts.

Die eine Hälfte existiert bereits: seit
[slice-043](../done/slice-043-schicht-abdeckung-sichtbar.md) meldet ein Scan auf stderr, welche
Dateien in **keinem** `layers`-Glob liegen. Es fehlt das Gegenstück:

> N Datei(en) enthalten Import-Schreibweisen, die dieses Backend nicht extrahiert.

**Belegt aus dem Einsatz** (Fremd-Repo, 2026-08-09): elternrelative C++-Includes, relative
Python-Importe und Mehrfach-Direktiven mussten je Sprach-Skelett von Hand nachgestellt werden, um
zu erkennen, wo das Werkzeug blind ist. Ein grünes Gate über teilweise nicht extrahiertem Code
sieht aus wie ein grünes Gate über geprüftem Code — dieselbe Klasse, die
[slice-043](../done/slice-043-schicht-abdeckung-sichtbar.md) für die Schicht-Seite geschlossen hat.

**Diese Diagnose subsumiert den zweiten Konsumenten-Befund** (Mehrfach-Direktiven, siehe
[slice-084](../done/slice-084-handbuch-heuristik-grenzen.md)): statt jede Grenze einzeln im Handbuch zu
suchen, meldet das Werkzeug, was es in **diesem** Baum nicht gegriffen hat.

## 2. Betroffene Module

Zwei Schichten:

1. **Spec** — [`spec/lastenheft.md`](../../../../spec/lastenheft.md) (neue `AC-*` oder Erweiterung
   von [AC-FA-CLI-001](../../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)),
   [`spec/spezifikation.md`](../../../../spec/spezifikation.md), ADR nach dem Muster von
   [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md).
2. **`internal/adapter/driven/extract/`** + [`internal/cli/`](../../../../internal/cli/cli.go) —
   die Zählung und ihre Ausgabe.

## 3. Auszuführende Gates

`make gates`, `make image-test` (die Diagnose ist Teil des CLI-Vertrags).

**Der Entscheid, der vor dem Bau fällt: was zählt als „nicht extrahiert"?** Der Slice nimmt ihn
nicht vorweg. Drei Kandidaten, die sich in Präzision und Rauschen unterscheiden:

| Kandidat | Aussage | Risiko |
|---|---|---|
| Zeilen, die ein **Import-Schlüsselwort** tragen, aber von keinem Regex gegriffen werden | präzise am Fund | Kommentare/Strings mit `import` erzeugen Rauschen |
| Dateien **ohne einen einzigen** extrahierten Import | billig | eine Datei ohne Importe ist normal, kein Befund |
| Nur die **benannten** Grenzen (relative Importe, Mehrfach-Direktiven) gezielt erkennen | rauschfrei | wächst mit jeder neuen Grenze mit, also Wartungsposten |

**Advisory, nicht gatend** — wie
[ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) es für die Abdeckungs-Diagnose
entschieden hat: der Exit-Code bleibt unberührt, vollständige Extraktion erzeugt **keine** Ausgabe.
Eine Diagnose, die bei jedem Lauf spricht, wird weggeschaltet.

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| Python-Datei mit `from . import x` | gemeldet |
| Python-Datei mit `import a, b` | gemeldet |
| Baum ohne solche Formen | **keine Ausgabe**, Exit unverändert |

## 4. Was bewusst nicht getan wird

- **Die Grenzen selbst schließen.** Ein AST-Backend ist in
  [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
  ausdrücklich Out-of-Scope. Dieser Slice macht die Grenze *sichtbar*, er verschiebt sie nicht.
- **Gatend machen.** Ein `strict_extraction` analog `strict_coverage` ist wie dort vertagt — die
  advisory Form reicht für den belegten Bedarf.
- **Die zwei `--print-mk`-Defekte.** Eigene Slices
  ([slice-082](../done/slice-082-print-mk-docker-indirektion.md),
  [slice-083](../open/slice-083-print-mk-digest-selbstbezug.md)).

## 5. DoD

- [ ] Spec-first: die Diagnose steht als Vertrag im Lastenheft, geschärft durch eine ADR, bevor
      Code entsteht. Beleg: Lastenheft-Versionszeile + ADR mit `Status: Accepted`.
- [ ] Ein Scan meldet nicht extrahierte Import-Formen auf stderr, ohne den Exit-Code zu ändern;
      ein sauberer Baum bleibt still. Beleg: die drei Proben aus §3.
- [ ] `make gates` und `make image-test` grün — **Ausgabe in eine Datei**, Exit-Code getrennt
      geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
