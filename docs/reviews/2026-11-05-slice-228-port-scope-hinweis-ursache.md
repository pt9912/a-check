# Review-Report: slice-228 — 2026-11-05

**Review-Art:** unabhängiger Lauf (frischer Kontext, kein `fork`) — Code-Review gegen Slice-Plan,
ADR-0040, AC-FA-RULE-010 und `AGENTS.md` §3.

**Gegenstand:** Commit `5975d98` (Range `5aac4b6..5975d98`) — `InertPortScope.Cause`
(`internal/hexagon/core/rules.go`), `writePortScopeNotice`/`causeLabel` (`internal/cli/cli.go`),
Tests, Benutzerhandbuch, CHANGELOG. Slice-Plan: slice-228.

**Skill:** `.harness/skills/reviewer.md` (Stand `5975d98`, unverändert) · **Modell:** claude-sonnet-5-5 · **Datum:** 2026-11-05

**Eingangs-Kontext:** Slice-Plan slice-228 (§1, §3, §6), ADR-0040, AC-FA-RULE-010, SPEC-Abschnitt
„Port-Scope-Diagnose" in `spec/spezifikation.md`, `AGENTS.md` §3, Code `scopeFor`, `portScope`,
`portLocality`, `appTreeContains`, `nestedInAppTree`, `segIndex`.

**Messgrenze:** Es lief kein Go/Docker (Host-Regel, nur Lesen). Alle Befunde sind durch
Code-Lesung hergeleitet, nicht durch Lauf bestätigt; Geltungsbereich der Herleitung: die
Funktionen oben, durchgespielt an konkreten Globs.

---

## Findings

### F-1 — `inertCause` klassifiziert ein grobes App-Glob eine Ebene über dem Port-Ordner als `direction`

- `kategorie`: HIGH
- `quelle`: Slice-Plan §1/§3 (DoD-Punkt „nennt nur eine Abhilfe, die den Hinweis beseitigt"), AC-FA-RULE-010, Doc-Kommentar zu `InertDirection`/`InertCoarseApp`
- `pfad`: `internal/hexagon/core/rules.go:376-381` (`inertCause`), Zusage in `:366-367`
- `befund`: `inertCause` fragt nur, ob `parentOf(parentOf(p))` im App-Baum liegt, und prüft weder
  `Direction` der Schicht noch, ob `lastSegment(p)` überhaupt ein Richtungssegment ist. Der Kommentar
  sagt zu, die Klassen unterschieden sich durch die Frage „würde das Abschneiden des Richtungssegments
  den App-Baum erreichen?" — das fragt der Code nicht. Durchgespielt: App-Glob `src/app/**`, Port-Glob
  `src/app/order/ports/**`, keine `direction`. `scopeFor` liefert `src/app/order` (Marker-Schnitt);
  `appTreeContains` ist falsch (App-Präfix `src/app` ist Vorfahr, nicht Nachfahr des Scopes), also
  wird gemeldet. `inertCause`: `parentOf(parentOf("src/app/order/ports")) = "src/app"` liegt im
  App-Baum, Ergebnis `InertDirection`. Der Hinweis nennt „Ursache: direction" und empfiehlt „der
  Schicht ihre direction geben — sie muss dem letzten Segment des Globs entsprechen"; das letzte
  Segment ist `ports`, keine `direction` und keine Port-Ordner-Kürzung hilft. Das ist derselbe Defekt
  wie in Issue 4, für die einfachste Form des groben App-Globs (genau eine Ebene zwischen App-Präfix
  und Port-Ordner). Kein Test deckt diese Form: der Coarse-Test nutzt einen tiefen Port-Pfad, bei dem
  `parentOf(parentOf(p))` zufällig unterhalb des App-Präfixes liegt.
- `verifizierbar`: ja — Test mit obiger Konfiguration (`InertPortScopes`, Erwartung `InertCoarseApp`).
- `klasse`: Klassifikator prüft Ersatz-Bedingung statt der zugesagten Frage

### F-2 — Zeilenform der Port-Scope-Diagnose in der Spezifikation nicht nachgezogen

- `kategorie`: MEDIUM
- `quelle`: Slice-Plan Kopf („Hinweistext ist Ausgabe-Vertrag; ob die Spezifikation ihn zusagt, klärt §3"), SPEC-Abschnitt Port-Scope-Diagnose
- `pfad`: `spec/spezifikation.md:597-599` gegen `internal/cli/cli.go:178`
- `befund`: Die Spezifikation sagt je Glob eine Zeile `<glob> -> Scope <verzeichnis>` zu; die Ausgabe
  trägt jetzt ` (Ursache: …)` am Zeilenende und zwei bedingte Abhilfe-Zeilen. Der Commit fasst
  `spec/spezifikation.md` nicht an, §3 des Slice-Plans nennt die Klärung nicht als DoD-Punkt. Der
  bestehende Test (`Contains`) bleibt dadurch grün; ein Lauf, der die Zeilenform hält, existiert nicht.
- `verifizierbar`: nein — kein Gate vergleicht Spec-Prosa mit Ausgabeformat.
- `klasse`: Ausgabe-Vertrag geändert, Spec-Stratum nicht mitgezogen

### F-3 — Handbuch und Hinweistext: Aussage zu „app-weiten Ports" und „unterhalb des Scopes" ist ungenau

- `kategorie`: LOW
- `quelle`: Prüfpunkt (4) des Auftrags; `portLocality`, `portScope`, `appTreeContains`
- `pfad`: `docs/user/benutzerhandbuch.md` (Abschnitt Port-Scope-Hinweis, Bullet „Ursache `App-Glob …`"); `internal/cli/cli.go:189`
- `befund`: (a) „Liegen die Ports app-weit (`…/application/port/**`), hat `port-locality` für sie
  keine Aussage, und der Hinweis ist zu erwarten" stimmt nur, wenn der App-Glob gröber ist als der
  Scope (Gradle-Modul-Fall). Liegt `…/application/port/**` unter App-Glob `…/application/**`, ist
  der Scope `…/application` gleich dem App-Präfix, `appTreeContains` wahr, `portLocality` ohne
  Befund und kein Hinweis. (b) „Wirksam ist ein App-Glob unterhalb des Scopes": `appTreeContains`
  verlangt Gleichheit oder Nachfahr (`segIndex(p, scope) >= 0`); ein App-Glob **auf** dem Scope
  genügt. (c) Die Aussage „keine Aussage" für app-weite Ports trifft zu im Sinn „kein Befund
  möglich"; bei App-Glob auf dem Scope entscheidet die Regel zwar, meldet aber nie (jede App-Datei
  liegt im Scope). Die Aussagen stehen im Kontext des groben Globs und sind dort richtig; die
  Klammer `…/application/port/**` macht sie allgemeiner lesbar.
- `verifizierbar`: ja — Durchspielen gegen `appTreeContains`/`portLocality`, oder Test.
- `klasse`: Handbuch-Aussage weiter als der Code-Fall, auf den sie sich stützt

### F-4 — DoD-Punkt nennt drei Ursachen, der Code unterscheidet zwei

- `kategorie`: INFO
- `quelle`: Slice-Plan §3 zweiter Punkt
- `pfad`: `internal/hexagon/core/rules.go:361-372`
- `befund`: §3 verlangt „direction fehlt · direction passt nicht zum letzten Segment · Schnitt
  erreicht den App-Baum trotzdem nicht". Umgesetzt sind zwei Klassen; die ersten beiden fallen in
  `InertDirection`, der Abhilfetext deckt sie gemeinsam („sie muss dem letzten Segment des Globs
  entsprechen"). Ob das den Punkt erfüllt, ist Frage des Verifiers, kein Review-Befund.
- `verifizierbar`: nein
- `klasse`: Plan-Granularität gegen Umsetzung (Verweis an Verifier)

## Negativbefunde

- (2) Kürzung bei `noticeLimit`: geprüft, ohne Befund. Die Flags `direction`/`coarse` werden vor dem
  `continue` für **alle** Einträge gesetzt, auch für gekürzte; „… und N weitere" erscheint weiter
  vor den Abhilfe-Zeilen; der Kopf zählt `len(scopes)`. `TestPortScopeNoticeCapNamesRemainder` deckt
  den Rest-Hinweis; ein Test, der eine Ursache **nur** jenseits der Kappung hat, fehlt (nicht
  nachweisbar falsch, nur ungedeckt).
- (1) Randfälle: leerer Präfix — durch `p == ""` in `InertPortScopes` ausgeschlossen, bevor
  `inertCause` läuft; Präfix ohne `/` ergibt `inner == ""` und damit `InertCoarseApp`, dort aber
  ohnehin Scope `""` und nicht gemeldet. Mehrere App-Globs und Modul-Präfix-Toleranz: `inertCause`
  und `scopeFor` nutzen dasselbe Prädikat `appTreeContains`/`segIndex`, kein Abweichen gefunden.
  `direction` gesetzt und passend ist im gemeldeten Fall konsistent mit `InertCoarseApp` (passende
  `direction` hätte `scopeFor` schon abgezogen, wenn der Schnitt darüber den Baum erreichte). Nicht
  passende `direction` mit erreichbarem Schnitt: `InertDirection`, Abhilfetext trifft zu. Ausnahme
  siehe F-1.
- (3) Kommentar-Regel §3.7 für `writePortScopeNotice`: geprüft, ohne Befund — Indikativ, nennt
  Zusage und Abgrenzung (Geschwister-Ports), keine Chronik, keine verworfene Alternative.
  `causeLabel`- und Konstanten-Kommentare ebenso (Zusage-Gehalt des `inertCause`-Kommentars: F-1).
- (5) Out-of-Scope §1: geprüft, ohne Befund. `scopeFor`/`portScope`/`portLocality` und
  `spec/lastenheft.md` unverändert; Geschwister-Ports bleiben ungemeldet; kein Quittier-Weg.
- `AGENTS.md` §3.2 Suppressions: geprüft, ohne Befund (keine `//nolint` im Diff).
- Determinismus (AC-QA-01): geprüft, ohne Befund — Reihenfolge der Ausgabe stammt aus der
  sortierten Liste; Flags sind reihenfolgeunabhängig.
- CHANGELOG `[Unreleased]`: geprüft, ohne Befund (Eintrag vorhanden, nennt Ausgabe-Änderung und
  unveränderten Exit-Code).
- Commit-Traceability: Message nennt slice-228, AC-FA-RULE-010, ADR-0040; ADR-0040 nicht geändert.

## Kategorie-Summary

HIGH 1 · MEDIUM 1 · LOW 1 · INFO 1

## Verdikt

Nicht freigabefähig: F-1 (HIGH) lässt den Fehler aus Issue 4 für die Ein-Ebenen-Form des groben
App-Globs bestehen und widerspricht der im Code-Kommentar gegebenen Zusage. F-2 vor Merge zu klären.
Kein Finding wird wegen eines Implementer-Widerspruchs herabzustufen sein; ab HIGH mit
Rollen-Widerspruch läuft der Konflikt-Pfad über den Architect.
