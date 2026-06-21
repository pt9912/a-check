# Review-Report — Benutzerhandbuch (`docs/user/benutzerhandbuch.md`)

- **Review-Art:** Doku-Review (Benutzerhandbuch) gegen Standard + reales Produkt
- **Gegenstand:** `docs/user/benutzerhandbuch.md` (Handbuch-Version 1.0)
- **Datum:** 2026-06-21 · **Modell:** Opus 4.8 · **Skill:** `.harness/skills/reviewer.md`
- **Methode:** Drei unabhängige Reviewer-Agenten (frischer Kontext, perspektiven-divers:
  Standard-Konformität · Faktentreue gegen Code/Spec · Nutzer-Tauglichkeit), Synthese
  mit adversarischer Verifikation der HIGH-Befunde gegen die realen Artefakte (Modul 11).

## Adversarisch verifiziert

| Gemeldet | Verifikation | Ergebnis |
|---|---|---|
| C-H1: `make a-check` (§3.3) bricht im Vorab-Stand (Image `ghcr.io/pt9912/a-check:0.1.0` nicht veröffentlicht) | `a-check.mk` pinnt diesen Tag; README sagt „GHCR-Release folgt"; Handbuch nennt für den make-Pfad keine Override-Brücke | **bestätigt — HIGH** |
| A-HIGH: Tag `a-check:dev` in README nicht belegt | `Makefile` `build` → `-t a-check:dev` (Reviewer B Negativbefund 9) — Tag ist korrekt, nur Quellenangabe unpräzise | **herabgestuft auf LOW** |

## Befunde (konsolidiert)

### HIGH

- **H1 — make-Pfad ohne Vorab-Brücke** (`benutzerhandbuch.md` §3.3): Das per
  `--print-mk` erzeugte `a-check.mk` pinnt auf das (noch nicht veröffentlichte)
  GHCR-Image; `make a-check` schlägt im Vorab-Stand fehl. Das Handbuch erklärt
  nur für `docker run` die `<a-check-image>`-Ersetzung, nicht für den make-Pfad.

### MEDIUM

- **M1 — undokumentierte Config-Schlüssel** (§4 + Glossar): `allow` und
  `forbidden_constructs` fehlen im Schema-Beispiel/Glossar, obwohl `--print-config`
  (das §3.2 erzeugen lässt) `forbidden_constructs` ausgibt und §3.4
  („verbotenes Konstrukt") darauf verweist — toter Verweis. (A-M, B-LOW1, C-M3, C-M4)
- **M2 — Exit-0 verschweigt stderr** (§2): „Exit 0 = nichts auf der Ausgabe"
  verschweigt, dass `report.go` immer `gesamt: 0 Befund(e)` auf stderr schreibt.
- **M3 — fehlender Image-Fehlerfall** (§6): der häufigste Erstnutzer-Fehler
  (Image nicht abrufbar) fehlt in der Fehlerbehebung.
- **M4 — fehlende Autor/Team-Angabe** (Kopf): Standard §9 verlangt „Autor oder Team".

### LOW

- L1 — `<a-check-image>` → `a-check:dev`-Gleichsetzung im Vorab-Stand nicht explizit (Tag korrekt).
- L2 — „mit allen … Defaults" (§4) überzeichnet: Optionalblöcke entfallen, haben keinen Default-Wert.
- L3 — Glossar deckt zentrale Begriffe nicht (Schicht, Kante/`edges`, `adapter_sink`, „Befund").
- L4 — `markers.ignore_symbols` wirkt nur auf Importe (nicht auf `forbidden_constructs`); im Handbuch nicht abgegrenzt.
- L5 — `adapter_sink`-Default-Semantik (fehlt ⇒ kein Adapter darf einen anderen importieren) nicht erwähnt.
- L6 — §3.4/§3.5 folgen nicht der Schritt-Form von §3.1–§3.3 (Tabelle/Fließtext) — für eine Regel-Übersicht vertretbar.

### INFO

- Keine Screenshots (für ein CLI korrekt, Standard §6). Rollen/Rechte (Standard §7) sauber als nicht-anwendbar begründet. Pre-Release-Formulierungen in §3.3-Hinweisen ehrlich.

## Negativbefunde (geprüft, ohne Befund)

- **Faktentreue (Reviewer B, hoch):** CLI/Flags/Exit-Codes, Befund-Format
  `pfad:zeile: regel: meldung`, die fünf Regelbeschreibungen, die `.a-check.yml`-
  Pflicht/Optional-Aussage, das YAML-Beispiel (strict-decodierbar), `--print-mk`/
  `a-check.mk`-Einbindung und die Read-only/netzlos/distroless/Digest-Aussagen
  stimmen alle mit `cli.go`/`config.go`/`rules.go`/`report.go`/`Dockerfile`/
  `Makefile`/`a-check.mk` und Spezifikation/Lastenheft überein.
- **Standard-Kern:** aufgabenbasiert (§2), direkte Sprache (§4), Versionsangaben (§9),
  Überschriftenstruktur/keine Farb-only-Info (§11), keine sensiblen Daten (§10),
  interne Links lösen auf (§13).

## Kategorie-Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 1 (H1) |
| MEDIUM | 4 (M1–M4) |
| LOW | 6 |
| INFO | 3 |
| verworfen/herabgestuft | 1 (A-HIGH → LOW) |

## Verdikt

Faktentreue und Standard-Kern sind hoch; **ein HIGH blockiert die
Praxistauglichkeit für die Pre-Release-Zielgruppe**: der dokumentierte
make-Gate-Pfad (§3.3) läuft ohne Override-Brücke ins Leere. Vor Freigabe sind
H1 und die vier MEDIUM (undokumentierte Config-Schlüssel, Exit-0-stderr,
Image-Fehlerfall, Autor-Angabe) zu beheben; die LOW betreffen Glossar-/
Präzisions-Politur.

## Disposition (Implementer, 2026-06-21)

Alle Befunde behoben (Handbuch auf 1.1). Beleg: `make doc-check` grün.

| Finding | Aktion |
|---|---|
| H1 | §1/§3.3/§6: Vorab-Brücke `make a-check A_CHECK_IMAGE=a-check:dev` dokumentiert; `<a-check-image>` → `a-check:dev` explizit gemacht. |
| M1 | §4 um `allow` + `forbidden_constructs` (Beispiel + Optionalliste) erweitert; §3.4 `port-impurity` an `forbidden_constructs` gebunden; Glossar ergänzt. |
| M2 | §2: Exit-0 weist nun die stderr-Zusammenfassung `gesamt: 0 Befund(e)` aus; stdout/stderr eingeführt. |
| M3 | §6: Fehlerfall „Image nicht gefunden" ergänzt. |
| M4 | Kopf: Autor (pt9912, Maintainer) ergänzt. |
| L1–L5 | „Defaults"-Wording korrigiert; Glossar (Schicht, Kante/`edges`, `adapter_sink`, `forbidden_constructs`, Befund); `ignore_symbols`-Reichweite (nur Importe); `adapter_sink`-Default-Semantik. |
| L6 | bewusst beibehalten: Regel-Übersicht (§3.4) als Tabelle — der Standard erlaubt Tabellen für Übersichten. |
