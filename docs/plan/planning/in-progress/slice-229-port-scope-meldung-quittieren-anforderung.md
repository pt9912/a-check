# slice-229 — Quittieren der Port-Scope-Meldung: Anforderung und Entscheidung

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** GitHub-Issue #7 (Wunsch), GitHub-Issue #4 (Anlass),
[ADR-0040](../../adr/0040-portscope-richtungssegment.md) (Port-Scope-Diagnose),
[AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (Konfigurationsdatei, strikt dekodiert),
[AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) (`port-locality`).

**Berührte Spec-Stellen:** [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) (eine neue AC im Bereich CONF entsteht daneben) ·
[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) und
[SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) werden vom Folge-Slice nachgezogen, nicht hier.

**Verantwortlich:** Claude

**Autor:** Claude. **Datum:** 2026-11-05.

**Lerneintrag — Form:** benannte Spec-Lücke (§7).

---

## 1. Ziel und Abgrenzung

**Ziel:** Der Wunsch aus Issue #7 — eine Schicht quittiert die Port-Scope-Meldung gezielt — ist als
**Anforderung** im Lastenheft und als **Entscheidung** in einer ADR festgeschrieben, bevor Code
entsteht. Die Maintainer-Entscheidung (2026-11-05): annehmen, eng gefasst. Der Rahmen:

1. **Nur die Port-Scope-Diagnose**, **je Schicht**, mit **Pflicht-Begründung**, in der Konfiguration.
2. **Die Meldung bleibt sichtbar, nur kürzer:** die Zeile nennt die **Schicht** und die Begründung,
   nicht nur eine Zahl — sonst verdeckt die Quittierung, was die Meldung zeigt.
3. **Eine überflüssige Quittierung fällt auf:** löst die Schicht die Meldung gar nicht aus, wird das
   gemeldet, damit eine veraltete Quittierung nicht unbemerkt stehen bleibt.
4. **Exit-Code und Regel unberührt:** `port-locality` prüft weiter genau so viel wie vorher.

Offen, und von der ADR zu entscheiden: der Name und die Form des Schlüssels (er soll sich später
auf andere Diagnosen erweitern lassen, ohne dass das hier zugesagt wird), die Zeilenform der
kurzen Meldung, und wie sich ein Konfigurations-Fehler (Schicht ohne Port-Rolle, leere Begründung)
verhält — Exit 2 wie bei jedem strikten Schema-Verstoß, oder Hinweis.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Die Umsetzung** (Schema-Dekodierung, Kern, Hinweis-Ausgabe, Handbuch). Ein Folge-Slice
  übernimmt sie: [slice-230](../open/slice-230-port-scope-meldung-quittieren-umsetzung.md). Der Schnitt
  nach Lieferwert: hier entsteht die Zusage, dort die Lieferung; Schema, Kern und Ausgabe sind mehr
  als zwei Schichten.
- **Quittieren der anderen drei Diagnosen** (Abdeckung, Grenze, Auflösung). *Bestand bleibt bewusst
  stehen*: bisher meldet ein einziger Konsument das Rauschen einer einzigen Diagnose; ein
  verallgemeinerter Entwurf ohne zweiten Beleg wäre geraten. Der Schlüssel wird nur so benannt,
  dass die Erweiterung möglich bleibt.
- **Eine Änderung von `port-locality` selbst** und von der Scope-Ableitung. *Ein anderer Vorgang*:
  die Quittierung ändert, was gemeldet wird, nicht was geprüft wird.
- **Ein stummes Quittieren und ein Schalter für alle Hinweise.** *Ein anderer Vorgang*, und vom
  Melder selbst ausgeschlossen: ein Schweigen verdeckte genau das, was der Hinweis sichtbar macht.

## 2. Definition of Done

- [x] Lastenheft: eine neue AC im Bereich CONF für das Quittieren (drei Akzeptanzkriterien
      Happy/Boundary/Negative im Given/When/Then-Stil plus Out-of-Scope), Versions-Bump und
      Historie-Zeile (Change Request, [AGENTS.md](../../../../AGENTS.md) §5 Regel 4).
- [x] Eine ADR mit der Form des Schlüssels, der Zeilenform und dem Fehlerverhalten; `Schärft:`
      auf die Spezifikation, **nicht** auf das Lastenheft; ADR-Index nachgezogen.
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein unabhängiger Review ist vorgesehen
([MR-019](../../../../harness/conventions.md#mr-019): Opt-in) — der Slice legt eine
Vertragsänderung fest.

## 3. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `spec/lastenheft.md` | update | neue AC im Bereich CONF, Versions-Bump, Historie |
| `docs/plan/adr/NNNN-….md`, `docs/plan/adr/README.md` | neu / update | Entscheidung zur Form; Index |

## 4. Trigger

**Start** (`open` → `next`): Die Maintainer-Entscheidung zu Issue #7 liegt vor — erfüllt am
2026-11-05 (annehmen, eng gefasst). Ein Release hängt nicht daran: der Hinweistext sagt seit slice-228, dass
die Meldung für app-weite Ports erwartbar ist. Ein **zweiter** Konsument mit demselben Rauschen
zöge den Entwurf auf alle Diagnosen, bevor der Slice startet.

**Rückführungen:** `in-progress` → `open` (blockiert): die ADR scheitert an einem Schema-Zwang —
eine ältere Version lehnt den neuen Schlüssel mit Exit 2 ab; dann ist die Frage der
Versions-Kopplung vor der ADR zu klären.

## 5. Closure-Trigger

DoD vollständig, Gates grün, Closure-Notiz mit Lerneintrag.

## 6. Risiken und offene Punkte

- **Das strikte Schema lehnt den neuen Schlüssel in älteren Versionen mit Exit 2 ab:** ein
  Konsument, der die Konfiguration anpasst und ein altes Image fährt, bricht. — **Ausgang:**
  *eingetreten*, Folge-Slice [slice-230](../open/slice-230-port-scope-meldung-quittieren-umsetzung.md): die [ADR-0045](../../adr/0045-port-scope-meldung-quittieren.md) nimmt es an (Punkt 5, keine Übergangsform), und das Handbuch sagt die Versions-Kopplung dort, wo der Schlüssel steht.
- **Eine Quittierung deckt auch einen später entstandenen echten Fehler zu:** sie gilt für alle
  Port-Globs der Schicht. — **Ausgang:** *eingetreten*, Folge-Slice [slice-230](../open/slice-230-port-scope-meldung-quittieren-umsetzung.md): die Zeile nennt die Zahl der Port-Globs, `port-locality` bleibt unverändert; die Grenze (kein Soll-Wert) steht in der [AC-FA-CONF-002](../../../../spec/lastenheft.md#ac-fa-conf-002--port-scope-meldung-je-schicht-quittieren) unter Out-of-Scope und in der [ADR-0045](../../adr/0045-port-scope-meldung-quittieren.md) unter Konsequenzen.

## 7. Closure-Notiz

**Lerneintrag — Form: benannte Spec-Lücke.** Das Lastenheft sagt das Quittieren der Port-Scope-Meldung
jetzt zu ([AC-FA-CONF-002](../../../../spec/lastenheft.md#ac-fa-conf-002--port-scope-meldung-je-schicht-quittieren), Lastenheft 0.31.0), die Spezifikation noch nicht:
[SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) hält
bis zur Umsetzung fest, dass das Quittieren nicht zugesagt ist, und
[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) kennt den
Schlüssel `acknowledge` nicht. Die Lücke ist gewollt und benannt; [slice-230](../open/slice-230-port-scope-meldung-quittieren-umsetzung.md) schließt sie.

**Geliefert:** [AC-FA-CONF-002](../../../../spec/lastenheft.md#ac-fa-conf-002--port-scope-meldung-je-schicht-quittieren) mit Happy, zwei Boundary, Negative, Determinismus und Out-of-Scope, Lastenheft
0.31.0 mit Historie-Zeile und Verweis in [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml); [ADR-0045](../../adr/0045-port-scope-meldung-quittieren.md) (Schlüssel `acknowledge` am
Schicht-Objekt, Zeilenform, Reihenfolge, Deckel, Exit-2-Fälle, keine Übergangsform); ADR-Index; die
Versionsangabe in `releasing.md`.

**Review:** [Report](../../../reviews/2026-11-05-slice-229-port-scope-quittieren-anforderung.md):
keine HIGH, vier MEDIUM — alle gegen den Code nachgerechnet und in AC und ADR behoben, bevor etwas
gepusht war: F-1 Reihenfolge, Deckel, Kopfzeile und Restzahl festgelegt, die „ohne Anlass"-Zeile
erscheint auch ohne sonstige Meldung; F-2 „effektive Rolle" statt `role: port`; F-3 die Aussage über
ältere Versionen ist Konsequenz und kein Akzeptanzkriterium mehr (nicht testbar); F-4 Schicht-Zuordnung
bei gemischten Ursachen und teilweise auslösenden Globs sowie die Entdoppelung je Glob. LOW: F-5
Randfälle des Exit 2 und F-7 getrenntes Boundary-Kriterium übernommen, F-6 als benannte Grenze
geführt. INFO: F-8 (Spezifikation hinkt bis slice-230 nach) ist die benannte Lücke oben; F-9 ([ADR-0045](../../adr/0045-port-scope-meldung-quittieren.md)
war vor dem Review `Accepted`) — es war noch nicht gepusht, die Korrektur geschah am Entwurf, ohne
Supersedes.

**Was hat funktioniert:** der Review las den Ausgabecode (`writePortScopeNotice`, `InertPortScopes`,
der Decoder) gegen die Entwürfe; so fielen die Stellen auf, die die Texte nicht nannten.

**Was ging anders als geplant:** Die Entwürfe waren aus der Beschreibung des Wunsches geschrieben und
nannten, was der Bestand längst entscheidet (Sortierung nach Glob, frühe Rückkehr bei null Treffern,
Entdoppelung je Glob, Rollen-Inferenz). Die ADR stand schon `Accepted`, als der Review sie fand.

**Steering-Loop-Eintrag:** benannte Spec-Lücke — liegt in
[SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes) und
[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema) (Quittieren
nicht nachgezogen, bis [slice-230](../open/slice-230-port-scope-meldung-quittieren-umsetzung.md) liegt).

**Beobachtungs-Register (`../observations/`):**
`BEO-SPEC/anforderung-legt-ausgabe-ohne-ausgabecode-fest` ist neu (1×): eine Anforderung erweiterte
eine vorhandene Meldung, ohne den Ausgabecode gelesen zu haben. Unter der Schwelle, Stand `offen`.

**Folge-Slices:** [slice-230](../open/slice-230-port-scope-meldung-quittieren-umsetzung.md) (Umsetzung, Spezifikation, Handbuch, CHANGELOG).

**Risiken aus §6:** beide tragen den Ausgang *eingetreten* mit Folge-Slice [slice-230](../open/slice-230-port-scope-meldung-quittieren-umsetzung.md).

**Drei Paarungen:** Anker — der Zielort ist die Spezifikation, `LH-`-Gegenstück entfällt (benannte
Spec-Lücke) · Folge-Slice — [slice-230](../open/slice-230-port-scope-meldung-quittieren-umsetzung.md) existiert als Datei in `open/` · Register —
`BEO-SPEC/anforderung-legt-ausgabe-ohne-ausgabecode-fest` existiert mit nicht leerem `evidence/`.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-11-05).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `SPEC` (Achsen 1, 2, 3 ✓) und `ADR` (Achsen 1, 2, 3 ✓).

**Vorgelagert — offene Beobachtungen sichten** (2026-11-05): `BEO-KERN/dirvocab-portfor-auseinander`
(2×, offen) — dieselbe `direction`-Zuordnung steht an drei Stellen; berührt diesen Slice nur,
wenn die ADR ein Richtungs-Vokabular in den Schlüssel aufnimmt. Sonst keine Treffer zu Diagnose-
Rauschen oder Quittierung.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
