# slice-228 — Port-Scope-Hinweis erscheint trotz gesetzter `direction`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** [ADR-0040](../../adr/0040-portscope-richtungssegment.md) (Richtungssegment im Port-Glob),
[AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) (`port-locality`),
GitHub-Issue #4.

**Berührte Spec-Stellen:** [AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) ·
der Hinweistext ist Ausgabe-Vertrag — [SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes), Abschnitt Port-Scope-Diagnose, Spezifikation 0.40.0.

**Verantwortlich:** Claude

**Autor:** Claude. **Datum:** 2026-11-05.

**Lerneintrag — Form:** benannte Spec-Lücke.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der Hinweis „Port-Glob(s) erreichen den App-Baum nicht — port-locality bleibt dafür
ungeprueft" erscheint für einen Port-Glob mit gesetzter, zum letzten Segment passender
`direction` nur dann, wenn die Abhilfe im Text die Lücke tatsächlich schließt. Gemeldet
(Issue #4, Image v0.23.0): zwei Schichten `…/application/port/{inbound,outbound}/**` mit
`direction` lösen den Hinweis in jedem Lauf aus, obwohl die Abhilfe „der Schicht ihre direction
geben" bereits erfüllt ist.

Der Auslöser ist `InertPortScopes` (`internal/hexagon/core/rules.go`): Port-Ordner im App-Baum,
abgeleiteter Scope (`scopeFor`) außerhalb. Der direction-Schnitt greift nur, wenn der Präfix auf
den `direction`-Wert endet, der Marker-Schnitt den App-Baum verfehlt **und** der Schnitt darüber
ihn erreicht. Welche der drei Bedingungen im gemeldeten Repo scheitert, ist offen — die
App-Globs des Melders fehlen im Issue.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Eine Lockerung der Scope-Ableitung für Globs unterhalb des Port-Ordners** (Issue Punkt 3).
  *Ein anderer Vorgang*: sie ändert, was `port-locality` prüft, und braucht eine eigene ADR
  ([AGENTS.md](../../../../AGENTS.md) §3.5); ob sie nötig ist, zeigt erst die Reproduktion.
- **Änderungen an [AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality).** *Schicht-Abgrenzung*: Anforderungen entstehen nur im
  Lastenheft ([AGENTS.md](../../../../AGENTS.md) §5, Regel 4); trifft die Reproduktion eine Lücke
  der Anforderung, wird sie ein eigener Change Request.
- **Klassische Geschwister-Ports** (`hex/ports/**` neben `hex/services/**`). *Bestand bleibt
  bewusst stehen*: ihre Inertheit ist in [AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) zugesagt und wird bewusst nicht
  gemeldet.
- **Ein Weg, die Meldung zu quittieren** (Wunsch des Melders im Issue). *Ein anderer Vorgang*: eine
  neue Konfigurationsmöglichkeit mit Folgen für Schema und Anforderung, kein Textfix; der Melder legt
  dafür ein eigenes Issue an.

## 2. Analyse

Konfiguration des Melders (Issue #4, 2026-11-05): ein App-Glob `src/hexagon/application/**`
für das ganze Gradle-Modul, die Port-Globs liegen darunter
(`…/src/main/kotlin/<pkg>/application/port/{inbound,outbound}/**`, `direction` gesetzt).

**Befund, gemessen durch Tests:** `appTreeContains(scope)` verlangt, dass der Scope ein Vorfahr
(oder gleich) eines App-Glob-Präfixes ist. Hier liegen beide Schnitte (Marker wie direction)
**unter** dem App-Präfix — (a) trifft zu, (c) scheitert, weil der App-Glob gröber ist als der
Scope. Der Scope bleibt der Port-Ordner, der Hinweis erscheint. `direction` kann daran nichts
ändern: beide Abhilfen des alten Hinweistexts sind für diese Form wirkungslos.
`TestInertPortScopesCoarseAppGlobIsNotADirectionProblem` bildet die Konfiguration des Melders nach.

**Korrektur an der ersten Lesung:** „ein feinerer App-Schnitt (je Use-Case ein Glob)" hilft nur bei
Ports **je Slice**. Bei app-weiten Ports (`…/application/port/**`) wäre der Scope auch dann
`…/application`, alle App-Dateien lägen darin, und die Regel hätte nichts zu melden. Der Hinweis
nennt deshalb „App-Glob **auf oder unterhalb** des Scopes", nicht „feiner schneiden".

## 3. Definition of Done

- [x] Reproduktion als Test: Konfiguration des Issues (verkürzt) ⇒ Hinweis erscheint; der Test
      ist ohne den Fix aus dem richtigen Grund rot.
- [x] Der Hinweis nennt nur eine Abhilfe, die den Hinweis beseitigt: unterscheidet *direction
      fehlt* · *direction passt nicht zum letzten Segment* · *Schnitt erreicht den App-Baum
      trotzdem nicht*; ist die Ursache ein Fehler der Ableitung, wird diese behoben.
- [x] Benutzerhandbuch (`docs/user/benutzerhandbuch.md`, Abschnitt zum Hinweis) und CHANGELOG
      `[Unreleased]` nachgezogen, falls Text oder Verhalten sich ändert.
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein unabhängiger Review ist vorgesehen
([MR-019](../../../../harness/conventions.md#mr-019): Opt-in) — der Slice ändert Kern-Logik und
Ausgabe-Text.

## 4. Plan (vor Code)

| Datei / Komponente | Änderungs-Art | Begründung |
|---|---|---|
| `internal/hexagon/core/rules.go` | ggf. update | `InertPortScopes`/`scopeFor`, falls die Ableitung fehlerhaft ist |
| `internal/hexagon/core/rules_test.go` | update | Reproduktion des Issues |
| `internal/cli/cli.go`, `internal/cli/cli_test.go` | update | Hinweistext je Ursache |
| `docs/user/benutzerhandbuch.md`, `CHANGELOG.md` | update | Ausgabe-Vertrag nachziehen |

## 5. Trigger

**Start** (`open` → `next`): Der Melder von Issue #4 hat die App-Globs seiner Schichten im Issue
nachgereicht — prüfbar am Issue-Thread; erfüllt am 2026-11-05. Seine Antwort auf die Frage nach
dem feineren Schnitt (Use-Cases flach in einem Ordner, `port-locality` für das Modul nicht
entschieden) lautet: Text korrigieren, Quittier-Weg als eigener Wunsch.

**Rückführungen:** `in-progress` → `open` (blockiert): die Reproduktion zeigt eine Lücke von
[AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) selbst — dann Change Request statt Fix.

## 6. Risiken und offene Punkte

- **Die Ursache liegt in der Anforderung, nicht im Code:** die Konfiguration des Melders ist
  nach [AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) legitim inert (App-Glob gröber als der Scope), und nur der Hinweistext
  ist falsch — die Lesung des Codes spricht dafür, gemessen ist sie nicht. — **Ausgang:**
  *entfallen*, gestrichen mit Begründung: gemessen durch den Reproduktions-Test ist die
  Konfiguration des Melders nach der Anforderung inert, die Anforderung trägt und bleibt
  unverändert; falsch war nur der Hinweistext, und der ist korrigiert.
- **Eine Änderung der Auslöse-Logik könnte Hinweise bei Bestands-Konfigurationen verschwinden
  lassen** und damit eine Lücke verdecken. — **Ausgang:** *entfallen*, gestrichen mit Begründung: das Auslöse-Kriterium
  (`InertPortScopes`: Port-Ordner im App-Baum, Scope außerhalb) ist unverändert, nur das Feld
  `Cause` ist hinzugekommen; die Bestands-Tests (aufgelöste Konfiguration still, Geschwister-Ports
  still, Kürzung) sind grün.

## 7. Closure-Notiz

**Lerneintrag — Form: benannte Spec-Lücke.** Die Port-Scope-Diagnose wies eine Abhilfe aus, die
nur für **eine** von zwei Ursachen wirkt, *weil* das Kriterium bewusst der Defekt und nicht die
Ursache ist ([ADR-0040](../../adr/0040-portscope-richtungssegment.md)) — und die Abhilfe-Zeile stillschweigend die gemessene Ursache annahm. Die
Lücke ist in der Spezifikation benannt: [SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)
(Spezifikation 0.40.0) sagt jetzt die Ursache je Zeile zu und hält ausdrücklich fest, dass ein
**Quittieren** der Meldung nicht zugesagt ist — ein Wunsch des Melders, der als eigenes Issue
entschieden wird.

**Geliefert:** `InertPortScope.Cause` (`InertDirection` | `InertCoarseApp`) und `inertCause` im Kern;
Hinweistext je Ursache (`writePortScopeNotice`); drei Kern-Tests und zwei CLI-Tests; Spezifikation
0.40.0, Benutzerhandbuch, CHANGELOG `[Unreleased]`.

**Messung:** Der CLI-Test `TestPortScopeNoticeCoarseAppGlobNamesTheRightRemedy` war vor der
Textänderung rot, mit der Meldung „direction ist gesetzt; diese Abhilfe wirkt hier nicht" — der
gemeldete Fehler, aus dem richtigen Grund. Danach `make test`, `make gates` und `make verify` grün.

**Was hat funktioniert:** das Nachrechnen der Reviewer-Befunde am Code, bevor sie behoben wurden.
F-1 bestätigte sich an der einfachsten Form (`src/app/order/ports/**`): meine erste Klassifikation
hatte denselben Fehler wie der gemeldete, nur eine Ebene flacher, weil der Test nur den tiefen Pfad
des Melders prüfte.

**Was ging anders als geplant:** Meine erste Abhilfe („App-Schicht feiner schneiden") war für die
Struktur des Melders falsch (app-weite Ports); sie steht korrigiert im Hinweis und im Handbuch. Die
Spezifikation sagte die Zeilenform schon zu — der Slice war damit nicht textlos (Review F-2).

**Review:** [Report](../../../reviews/2026-11-05-slice-228-port-scope-hinweis-ursache.md): F-1 (HIGH)
behoben mit Test, F-2 (MEDIUM) Spezifikation nachgezogen, F-3 (LOW) Handbuch und Hinweis präzisiert,
F-4 (INFO) — die DoD nennt drei Ursachen, der Code zwei (*direction fehlt* und *passt nicht* teilen
sich dieselbe Abhilfe); die DoD-Zeile bleibt als Planungs-Wortlaut stehen.

**Steering-Loop-Eintrag:** benannte Spec-Lücke — liegt in
[SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)
(Quittieren der Port-Scope-Meldung nicht zugesagt).

**Beobachtungs-Register (`../observations/`):**
`BEO-KERN/dirvocab-portfor-auseinander` bekommt einen zweiten Beleg (2×): `isPortDirection` führt
`inbound`/`outbound` als drittes Literal neben `portFor` und `dirVocab`. Unter der Schwelle, Stand
`offen`.

**Folge-Slices:** keine — der Quittier-Weg wird erst ein Slice, wenn der Melder sein Issue angelegt
hat und die Entscheidung (ADR) steht.

**Risiken aus §6:** beide tragen den Ausgang *entfallen*, mit Begründung.

**Drei Paarungen:** Anker — der Zielort ist die Spezifikation, `LH-`-Gegenstück entfällt (benannte
Spec-Lücke) · Folge-Slice — keiner genannt · Register — `BEO-KERN/dirvocab-portfor-auseinander`
existiert mit nicht leerem `evidence/`.

**Trigger-Audit der aktiven MR:** [`MR-016`](../../../../harness/conventions.md#mr-016) [`MR-019`](../../../../harness/conventions.md#mr-019) [`MR-025`](../../../../harness/conventions.md#mr-025) [`MR-027`](../../../../harness/conventions.md#mr-027) [`MR-028`](../../../../harness/conventions.md#mr-028) [`MR-029`](../../../../harness/conventions.md#mr-029) [`MR-030`](../../../../harness/conventions.md#mr-030) — 0 offen (geprüft 2026-11-05).

## 8. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `KERN` (Achsen 1, 2, 3 ✓). Der Hinweistext liegt in
`internal/cli`, einer Pfad-Familie, die die Modus-Deklaration in `harness/conventions.md` unter
keiner Sub-Area nennt; der Slice behandelt sie als Ausgabe-Rand von `KERN` und benennt die Lücke
statt sie still zu füllen.

**Vorgelagert — offene Beobachtungen sichten** (2026-11-05):
`BEO-KERN/dirvocab-portfor-auseinander` (1×, offen) — dieselbe `direction`-Zuordnung in zwei
Paketen; der Fall hier betrifft `direction` im Kern, ist aber ein anderer Befund (Hinweis-Auslöser,
nicht Vokabel-Drift). Keine weiteren Treffer.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF.
