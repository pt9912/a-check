# slice-228 — Port-Scope-Hinweis erscheint trotz gesetzter `direction`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `git mv` (`make slice-mv`).

**Welle:** ohne Welle — die Closure-Bedingung wäre die eigene DoD.

**Bezug:** [ADR-0040](../../adr/0040-portscope-richtungssegment.md) (Richtungssegment im Port-Glob),
[AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) (`port-locality`),
GitHub-Issue #4.

**Berührte Spec-Stellen:** [AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) ·
der Hinweistext ist Ausgabe-Vertrag; ob die Spezifikation ihn zusagt, klärt §3.

**Verantwortlich:** Claude

**Autor:** Claude. **Datum:** 2026-11-05.

**Lerneintrag — Form:** wird bei Closure benannt.

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

## 2. Analyse

Konfiguration des Melders (Issue #4, 2026-11-05): ein App-Glob `src/hexagon/application/**`
für das ganze Gradle-Modul, die Port-Globs liegen darunter
(`…/src/main/kotlin/<pkg>/application/port/{inbound,outbound}/**`, `direction` gesetzt).

**Lesung des Codes (nicht gemessen):** `appTreeContains(scope)` verlangt, dass der Scope ein
Vorfahr eines App-Glob-Präfixes ist. Hier liegen beide Schnitte (Marker wie direction) **unter**
dem App-Präfix — (a) trifft zu, (c) scheitert, weil der App-Glob gröber ist als der Scope. Der
Scope bleibt der Port-Ordner, der Hinweis erscheint. `direction` kann daran nichts ändern: die
beiden Abhilfen des Hinweistexts sind für diese Form beide wirkungslos; wirksam wäre ein feinerer
App-Schnitt (je Use-Case ein Glob). Die Reproduktion entsteht als Test in
`internal/hexagon/core/rules_test.go` (Vorbild: der Test um Zeile 1549) und belegt die Lesung.

## 3. Definition of Done

- [ ] Reproduktion als Test: Konfiguration des Issues (verkürzt) ⇒ Hinweis erscheint; der Test
      ist ohne den Fix aus dem richtigen Grund rot.
- [ ] Der Hinweis nennt nur eine Abhilfe, die den Hinweis beseitigt: unterscheidet *direction
      fehlt* · *direction passt nicht zum letzten Segment* · *Schnitt erreicht den App-Baum
      trotzdem nicht*; ist die Ursache ein Fehler der Ableitung, wird diese behoben.
- [ ] Benutzerhandbuch (`docs/user/benutzerhandbuch.md`, Abschnitt zum Hinweis) und CHANGELOG
      `[Unreleased]` nachgezogen, falls Text oder Verhalten sich ändert.
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben (oder „keine Beobachtung" notiert).
- [ ] Jedes Risiko aus §6 trägt einen Ausgang.

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
nachgereicht — prüfbar am Issue-Thread; erfüllt am 2026-11-05. Offen ist seine Antwort auf die
Frage, ob er die App-Schicht feiner schneiden kann oder `port-locality` gar nicht will; sie
entscheidet, ob der Slice nur den Text ändert oder einen Quittier-Weg braucht (dann ADR).

**Rückführungen:** `in-progress` → `open` (blockiert): die Reproduktion zeigt eine Lücke von
[AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) selbst — dann Change Request statt Fix.

## 6. Risiken und offene Punkte

- **Die Ursache liegt in der Anforderung, nicht im Code:** die Konfiguration des Melders ist
  nach [AC-FA-RULE-010](../../../../spec/lastenheft.md#ac-fa-rule-010--port-lokalität-regel-port-locality) legitim inert (App-Glob gröber als der Scope), und nur der Hinweistext
  ist falsch — die Lesung des Codes spricht dafür, gemessen ist sie nicht. — **Ausgang:**
  *(bei Closure zuzuweisen: eingetreten / entfallen / weiter offen)*
- **Eine Änderung der Auslöse-Logik könnte Hinweise bei Bestands-Konfigurationen verschwinden
  lassen** und damit eine Lücke verdecken. — **Ausgang:** *(bei Closure zuzuweisen:
  eingetreten / entfallen / weiter offen)*

## 7. Closure-Notiz

*(folgt bei Closure)*

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
