# slice-084 — Heuristik-Grenzen dort, wo Konsumenten lesen

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** Konsumenten-Befund vom 2026-08-09 (Mehrfach-Direktiven);
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).
**Bezug:** [slice-081](../open/slice-081-heuristik-diagnose.md) — die Laufzeit-Diagnose löst dasselbe
Problem grundsätzlicher; dieser Slice ist die Sofortmaßnahme.

---

## 0. Trigger

**Beginn:** sofort. Reine Doku-Currency, kein Vertrag berührt, kein Entscheid offen.

**Verwerfen**, falls [slice-081](../open/slice-081-heuristik-diagnose.md) zuerst gebaut wird **und** seine
Diagnose die betroffenen Formen vollständig meldet — dann wäre der Handbuch-Nachtrag ein zweiter,
schlechter gepflegter Ort für dieselbe Aussage. Bis dahin steht der Bedarf.

## 1. Auslöser

**Mechanismus: eine Grenze ist dokumentiert — nur nicht dort, wo sie gebraucht wird.**

Ein Konsument meldete am 2026-08-09, Mehrfach-Direktiven auf einer Zeile (`import a, b`,
`using A; using B;`) würden nur einmal gegriffen und das sei „nirgends ausgewiesen". Nachgemessen:
**es ist ausgewiesen** — im Lastenheft, im Out-of-Scope-Absatz von
[AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion), sogar
zweimal (generisch und Python-spezifisch).

Der Befund trägt trotzdem, nur anders gelagert. Verglichen mit der **relativen Python-Import**-Grenze:

| Ort | relativer Import | Mehrfach-Direktive |
|---|---|---|
| [`spec/lastenheft.md`](../../../../spec/lastenheft.md) | ✅ | ✅ |
| [`docs/user/benutzerhandbuch.md`](../../../user/benutzerhandbuch.md) | **8 Treffer** | **0** |
| Quelltext-Kommentar | ✅ | **0** |

Der Unterschied ist nicht *dokumentiert vs. undokumentiert*, sondern **an drei Orten vs. an einem,
den Konsumenten nicht lesen**. Das Lastenheft ist vertraglich abnahmebindend — es ist keine
Bedienungsanleitung.

**Betroffen sind alle Backends außer C++**, wo Präprozessor-Direktiven ohnehin zeilenweise sind.
Die Regexe in
[`internal/adapter/driven/extract/extract.go`](../../../../internal/adapter/driven/extract/extract.go)
sind durchgehend `^\s*…`-verankert und werden mit `FindStringSubmatch` (Erst-Treffer) ausgewertet.

## 2. Betroffene Module

Eine Schicht: **Benutzer-Doku** —
[`docs/user/benutzerhandbuch.md`](../../../user/benutzerhandbuch.md), plus ein Quelltext-Kommentar
in [`extract.go`](../../../../internal/adapter/driven/extract/extract.go) an der Stelle, die die
Grenze erzeugt.

**Kein Lastenheft, keine ADR** — der Vertrag ist unverändert und korrekt. Es wird nichts
entschieden, nur umgezogen.

## 3. Auszuführende Gates

`make doc-check`, `make gates`.

**Negativ-Probe** — hier ist der Abgleich die Probe, wie in
[slice-077](../done/slice-077-status-aussagen.md):

| Probe | Erwartung |
|---|---|
| Grenzen im Handbuch-Abschnitt zu Heuristik-Grenzen | decken die im Lastenheft geführten ab, die einen Konsumenten betreffen |
| `extract.go` an der Erst-Treffer-Stelle | trägt einen Kommentar, der die Grenze benennt |

Eine Fixture gibt es nicht und wäre gelogen: der Fehler ist ein fehlender Ort, kein Verhalten.

## 4. Was bewusst nicht getan wird

- **Alle zwanzig Out-of-Scope-Punkte ins Handbuch kopieren.** Das Lastenheft bliebe die Quelle;
  ein vollständiges Duplikat driftet garantiert. Aufzunehmen ist, was einen Konsumenten beim
  **Konfigurieren** trifft — nicht, was ein Abnahme-Kriterium abgrenzt.
- **Die Grenze schließen.** `FindAllString` statt `FindStringSubmatch` wäre eine
  Verhaltensänderung mit Vertragsbezug — das ist [slice-081](../open/slice-081-heuristik-diagnose.md) oder
  ein eigener Schnitt, nicht dieser.

## 5. DoD

- [x] Das Benutzerhandbuch führt die Mehrfach-Direktiven-Grenze dort, wo es die übrigen
      Heuristik-Grenzen führt. Beleg: neuer Punkt 4 unter *„a-check findet nichts, obwohl Verstöße
      erwartet werden"* (§6 Fehlerbehebung) — mit einer Tabelle der vier Formen, die beim
      Konfigurieren auffallen.
- [x] `extract.go` trägt an der Erst-Treffer-Stelle einen Kommentar, der die Grenze benennt —
      wie es für den relativen Python-Import bereits der Fall ist. Beleg: Kommentarblock über
      `lineMatches`.
- [x] `make gates` grün — **Ausgabe in eine Datei**, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** Die Mehrfach-Direktiven-Grenze steht jetzt an **drei** Orten statt an einem —
[Lastenheft](../../../../spec/lastenheft.md) (unverändert),
[Benutzerhandbuch](../../../user/benutzerhandbuch.md) §6 und als Kommentar über `lineMatches` in
[`extract.go`](../../../../internal/adapter/driven/extract/extract.go). Damit trägt sie dieselbe
Abdeckung wie die relative-Python-Import-Grenze, an der sich der Konsumenten-Befund maß.

**Der Ort war die eigentliche Arbeit.** Naheliegend wäre ein Absatz bei der Python-Konfiguration
gewesen — dort steht die verwandte Grenze. Gewählt ist stattdessen der Fehlerbehebungs-Abschnitt
*„a-check findet nichts, obwohl Verstöße erwartet werden"*: er listete drei Ursachen (Globs,
Datei-Endung, `exclude`), und **keine davon war die Extraktion**. Genau dort sucht jemand, dem eine
erwartete Kante fehlt — und genau dort fand er bisher nichts.

**Lerneintrag — Form: geschärfte Regel.** Als Prüfsatz: *Eine Grenze ist dort zu dokumentieren, wo
jemand nach ihrem Symptom sucht — nicht dort, wo sie technisch entsteht.*

**Die Ursache** der Lücke ist die Trennung der Doku-Straten, die das Repo bewusst führt: das
Lastenheft ist **vertraglich abnahmebindend** und listet Out-of-Scope vollständig — über zwanzig
Grenzen in einem Absatz. Das Handbuch ist **Bedienungsanleitung** und wählt aus. Die Auswahl folgte
bisher der Sprache (was betrifft Python? was C#?), nicht dem **Symptom** (was tue ich, wenn ein
Befund fehlt?). Für den relativen Import fielen beide Achsen zusammen — er ist sprachspezifisch
*und* symptomatisch. Die Mehrfach-Direktive ist sprachübergreifend und fiel deshalb durch das
Raster.

**Bewusst nicht getan:** die übrigen ~20 Out-of-Scope-Punkte nachziehen. Aufgenommen sind vier
Formen, die beim **Konfigurieren** auffallen; ein vollständiges Duplikat würde gegen das Lastenheft
driften. Der Absatz verweist stattdessen auf die vollständige Liste.

**Zwei beobachtbare Closure-Kriterien:**

1. Die Grenze ist in `spec/lastenheft.md`, `docs/user/benutzerhandbuch.md` **und** `extract.go`
   auffindbar — vor diesem Slice nur im Lastenheft.
2. Der Handbuch-Eintrag steht unter dem Symptom („findet nichts, obwohl Verstöße erwartet werden"),
   nicht unter der Sprache — nachprüfbar an der Abschnitts-Überschrift.

**Folge-Slices:** keine. **Verwerfungs-Bedingung nicht eingetreten:**
[slice-081](../open/slice-081-heuristik-diagnose.md) ist noch nicht gebaut, der Handbuch-Eintrag
also weiterhin die einzige Stelle, an der ein Konsument die Grenze findet. Liefert 081 später eine
Laufzeit-Diagnose, bleibt dieser Eintrag trotzdem sinnvoll — er erklärt *warum*, die Diagnose sagt
nur *dass*.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
