# slice-120 — `verify-ac-form` ablösen: der letzte Eigenbau-Sensor

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** [slice-116 §4](../done/slice-116-nullmengen-haerte-cr.md) — der dort gestellte Trigger,
beide Hälften eingetreten. Schließt die Ablösung aus
[slice-080](../done/slice-080-verify-abloesung-dcheck.md) ab.

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Der vierte und letzte Eigenbau-Sensor aus [slice-080](../done/slice-080-verify-abloesung-dcheck.md)
entfällt: `verify-ac-form` geht ins Modul `structure` über.

## 2. Definition of Done

- [ ] Der Pin steht auf `v0.69.0`, das Fragment ist aus `--print-mk` **neu erzeugt** (nicht
      zeilenweise getauscht, slice-115), der Digest aus **zwei** Quellen belegt.
- [ ] Eine fünfte `structure`-Regel deckt die AC-Form mit `exempt-expect-count: 19`; die 19
      grandfatherten Anforderungen bleiben ausgenommen, ein zwanzigstes `AC-*` wird geprüft.
- [ ] `tools/verify-ac-form.sh` ist entfernt, seine Einhängung in
      [`Makefile`](../../../../Makefile), [`AGENTS.md`](../../../../AGENTS.md) §4 und der
      GATES-Liste des Guard nachgezogen. Beleg: Paritäts-Mutations-Probe in **beide** Richtungen.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| [`d-check.mk`](../../../../d-check.mk) | neu erzeugt | Pin `v0.69.0` |
| [`.d-check.yml`](../../../../.d-check.yml) | update | fünfte `structure`-Regel |
| `tools/verify-ac-form.sh` | entfällt | abgelöst |
| [`Makefile`](../../../../Makefile), [`AGENTS.md`](../../../../AGENTS.md) §4, `.claude/hooks/` | update | Einhängung und Deklaration |

**Auszuführende Gates:** `make gates` — tragend `doc-check`, `doc-targets`, `guard-selftest`. Zum
Abschluss `make verify`.

### Das Risiko ist vorab gemessen, nicht abgeschätzt

Beide Trigger-Hälften aus [slice-116 §4](../done/slice-116-nullmengen-haerte-cr.md) sind erfüllt
und je einzeln belegt:

| Prüfung | Ergebnis |
|---|---|
| `exempt-expect-count` in `--print-config` | vorhanden, mit beiden Config-Rändern dokumentiert |
| Digest aus zwei Quellen | Release-Notes == `docker inspect` |
| Fragment-Diff `--print-mk` | nur Pin + eine **Handbuch-Zeile** im Kopfkommentar; keine neuen Targets, Disable-Listen unverändert |
| fünf `doc-*`-Sensoren gegen `v0.69.0` | alle Exit 0 über den unveränderten Bestand |
| **19 deklariert, 19 ausgenommen** | **Exit 0, kein Befund** |
| **18 deklariert, 19 ausgenommen** | `section-exempt-mismatch`, Meldung nennt **beide** Zahlen |

Die letzten zwei Zeilen sind der eigentliche Beleg: die legitime Leermenge ist grün, und die
Deklaration ist keine Erlaubnis, sondern eine Behauptung, die widerlegt werden kann.

## 4. Trigger

**Start:** eingetreten — beide Hälften aus [slice-116 §4](../done/slice-116-nullmengen-haerte-cr.md).

**Rückführungen:**

- `in-progress` → `open`: falls die Paritäts-Probe in einer Richtung nicht deckt.

## 5. Closure-Trigger

Pin gehoben, Regel konfiguriert, Skript entfernt, Parität in beide Richtungen belegt, Gates grün.

**Was bewusst nicht getan wird:** die **Zahl 19 automatisch pflegen**. Sie ist eine Deklaration und
gehört von Hand nachgezogen, wenn die Grandfather-Liste sich ändert — was sie laut
[`AGENTS.md`](../../../../AGENTS.md) §5 nicht tut („die Liste wächst nicht mit"). Der Empfänger
nennt die Gefahr selbst: *wer die Zahl mitzieht, ohne die Aufzählung zu prüfen, hat einen Wächter,
der nur noch sich selbst bestätigt.*

## 6. Risiken und offene Punkte

- *Die Zahl `19` steht ab jetzt an zwei Orten — in der Aufzählung und als Deklaration* —
  **Ausgang:** <bei Closure>
- *Der Befund bricht laut Änderungsprotokoll des Empfängers die **Datei** ab und verdeckt den Rest*
  — das ist bei einer Datei mit **einer** Regel folgenlos, bei mehreren nicht.
  **Ausgang:** <bei Closure>

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5.)_

**Lerneintrag — Form: <geschärfte Regel | neuer Sensor | benannte Spec-Lücke>**

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Gate-/Werkzeug-Schicht**
(`.d-check.yml`, `tools/`, `Makefile`, `.claude/`) und mit der Deklarations-Tabelle der
**Harness-Einstieg**.

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-014`](../observations.md) (aktiviertes Modul
ohne Konfigurationsblock meldet grün) betrifft dieselbe Schicht und bleibt für `planning` offen;
dieser Slice konfiguriert eine Regel, er aktiviert kein neues Modul.

Alle berührten Sub-Areas GF.
