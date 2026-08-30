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

- [x] Der Pin steht auf `v0.69.0`, das Fragment ist aus `--print-mk` **neu erzeugt** (nicht
      zeilenweise getauscht, slice-115), der Digest aus **zwei** Quellen belegt.
- [x] Eine fünfte `structure`-Regel deckt die AC-Form mit `exempt-expect-count: 19`; die 19
      grandfatherten Anforderungen bleiben ausgenommen, ein zwanzigstes `AC-*` wird geprüft.
- [x] `tools/verify-ac-form.sh` ist entfernt, seine Einhängung in
      [`Makefile`](../../../../Makefile), [`AGENTS.md`](../../../../AGENTS.md) §4 und der
      GATES-Liste des Guard nachgezogen. Beleg: Paritäts-Mutations-Probe in **beide** Richtungen.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

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
  **Ausgang:** gestrichen mit Begründung: das ist nicht die Redundanz, sondern die **Zusage**. Eine
  Aufzählung allein kann veralten, ohne dass es auffällt; die Zahl macht genau das laut. Der
  Empfänger nennt die Kehrseite selbst — wer die Zahl mitzieht, ohne die Aufzählung zu prüfen, hat
  einen Wächter, der nur noch sich selbst bestätigt —, und die steht als Grenze im
  Konfigurations-Kommentar.
- *Der Befund bricht die **Datei** ab und verdeckt den Rest* — **Ausgang:** weiter offen im
  **Beobachtungs-Register** als Teil von `BEO-023`: heute folgenlos, weil genau **eine** Regel auf
  `spec/lastenheft.md` zeigt. Kommt eine zweite dazu, verdeckt ein `section-exempt-mismatch` sie.

## 7. Closure-Notiz

**Geliefert:** Der Pin steht auf `v0.69.0` (Digest aus zwei Quellen, Fragment neu erzeugt), eine
fünfte `structure`-Regel deckt die AC-Form mit `exempt-expect-count: 19`, und
`tools/verify-ac-form.sh` ist entfallen. Damit ist die Ablösung aus
[slice-080](../done/slice-080-verify-abloesung-dcheck.md) **vollständig**: alle vier Eigenbau-Sensoren
sind in Modulen aufgegangen, zusammen **653 Zeilen** Shell.

**Lerneintrag — Form: geschärfte Regel.** *Ein Prüfer mit leerer Prüfmenge meldet grün und ist
damit nicht „unbenutzt", sondern **unkalibriert** — seine Zusage ist ungeprüft, solange kein
Gegenstand entsteht.* Die Paritäts-Probe hat einen Sensor gefunden, der seit slice-054 grün im
Aggregat stand und **jede** neue Anforderung falsch beanstandet hätte. Er suchte
`^**Happy Path:**` — Zeilenanfang **ohne** Listen-Marker, mit einem Wort, das im Repo **null**
mal vorkommt und das die kanonische Quelle nicht kennt
([`harness/conventions.md`](../../../../harness/conventions.md) §Anforderungs-Anlege-Prozess sagt
*„Happy/Boundary/Negative"*). Alle 19 bestehenden schreiben `- **Happy:**` als Listenpunkt.

*Weil* seine Prüfmenge leer war, konnte das nicht auffallen: er meldete `0 neue AC(s) geprueft, 19
grandfathered` — und das ist **korrekt**. Die Leermenge auszuweisen (slice-070, Fund F-12) schützt
also vor dem stillen Verschwinden des Prüfbestands, **nicht** vor einer falschen Zusage über einen
Gegenstand, den es nie gab. Das sind zwei verschiedene Ehrlichkeiten.

Aufgedeckt hat es nicht das Lesen, sondern die **Probe mit einem echten zwanzigsten `AC-*`**. Ohne
sie wäre der Fehler entweder mit dem Skript verschwunden (folgenlos) oder ins Modul kopiert worden
(verewigt) — und die Ablösung hätte in beiden Fällen „Parität" gemeldet.

**Drei beobachtbare Closure-Kriterien:**

1. Die Proben-Matrix trennt Sensor und Modul in **vier** Formen, nicht in einer:

   | Probe | Form der neuen AC | Sensor | Modul |
   |---|---|---|---|
   | A | ohne die vier Bausteine | rot | rot |
   | B | `- **Happy:**` — Bestandsform aller 19 | **rot** | grün |
   | B′ | `- **Happy Path:**` | **rot** | grün |
   | B″ | `**Happy Path:**` ohne Listen-Marker | grün | **rot** |
   | C | unverändert, 0 neue | grün | grün |

   Der Sensor ist **nur** in einer Form grün, die keine bestehende Anforderung hat.
2. Das ist **keine Lockerung** ([`AGENTS.md`](../../../../AGENTS.md) §3.6, also ohne ADR
   zulässig): ein Gate, das jede gültige Eingabe rot meldet, ist nicht streng, sondern kaputt. Die
   neue Regel misst gegen die **kanonische Quelle** und gegen den Bestand — der Sensor gegen
   keines von beidem.
3. Die Deklaration ist widerlegbar: `18` statt `19` liefert `section-exempt-mismatch` mit **beiden**
   Zahlen. Eine Erlaubnis (`exempt-may-empty`, die erste Antragsform) hätte das nicht geleistet.

**Offene Risiken und ihr Ausgang:** der erste gestrichen mit Begründung, der zweite weiter offen im
Register.

**Beobachtungs-Register:** `BEO-023` neu angelegt (Gate-/Werkzeug-Schicht, 1×, Beleg slice-120):
ein Prüfer mit leerer Prüfmenge bleibt unkalibriert. Verwandt mit `BEO-014` (aktiviertes Modul
**ohne Konfiguration** meldet grün), aber verschieden: dort fehlt der Gegenstand der
Konfiguration, hier der Gegenstand der Prüfung.

**Folge-Slices:** [slice-117](../done/slice-117-handbuch-verweis-cli.md) — und der bekommt durch
dieses Release eine Präzedenz: `v0.69.0` löst dieselbe Aufgabe für das Schwester-Tool und wählt
**dieselbe** URL-Form (Hauptzweig, ohne Versionsangabe, mit Verweis auf den Software-Versions-
Stempel im Handbuch-Kopf) — unabhängig gefunden, bevor slice-117 geschnitten war.
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Gate-/Werkzeug-Schicht**
(`.d-check.yml`, `tools/`, `Makefile`, `.claude/`) und mit der Deklarations-Tabelle der
**Harness-Einstieg**.

**Vorgelagert — offene Beobachtungen sichten:** [`BEO-014`](../observations.md) (aktiviertes Modul
ohne Konfigurationsblock meldet grün) betrifft dieselbe Schicht und bleibt für `planning` offen;
dieser Slice konfiguriert eine Regel, er aktiviert kein neues Modul.

Alle berührten Sub-Areas GF.
