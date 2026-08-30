# slice-089 — Welle-Dateien: die Verweis-Invariante ist nicht erfüllbar

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** den Folge-Befund aus [`welle-13-results.md`](../done/welle-13-results.md) §Steering-Loop;
`SL-002`.
**Bezug:** [slice-060](../done/slice-060-slice-link-invariante.md) hat den Sensor gebaut,
[slice-075](../done/slice-075-sensor-messgroesse.md) ihn um Referenz-Definitionen erweitert.

---

## 0. Trigger

**Beginn: sofort.** Der Befund ist gemessen und die nächste Welle-Plan-Datei trifft dieselbe Kante.

**Rückführungen:**

- `in-progress` → `open`: falls der Entscheid auf eine Verzeichnis-Umstellung fällt — das berührt
  die Welle-Lifecycle-Doku in [`planning/README.md`](../README.md) und ist dann ein eigener,
  größerer Schnitt.

## 1. Auslöser

**Mechanismus: der Sensor prüft eine Invariante, die für eine ganze Gattung nicht gilt — und sagt
das nirgends.**

`tools/verify-slice-links.sh` begründet seinen Ansatz
wörtlich: *„Alle vier Lifecycle-Verzeichnisse liegen auf derselben Ebene, also gilt: Ein relativer
Verweis muss aus **jedem** Lifecycle-Verzeichnis auflösen."* Diese Voraussetzung ist bei
Welle-Plan-Dateien **nicht erfüllt**: sie liegen flach unter `docs/plan/planning/` und wandern bei
der Closure nach `done/` — also **eine Ebene tiefer**.

**Gemessen (2026-08-15):** Es gibt **keine** Verweis-Form, die aus beiden Positionen auflöst.

| Form | aus flach | aus `done/` |
|---|---|---|
| `done/slice-081-heuristik-diagnose.md` | ✅ | ❌ |
| `slice-081-heuristik-diagnose.md` | ❌ | ✅ |
| `../done/slice-081-heuristik-diagnose.md` | ❌ | ✅ |
| `../../../spec/lastenheft.md` | ✅ | ❌ |

Jede Zeile trägt genau ein ✅. Das ist keine Eigenheit der Beispiele, sondern folgt aus dem
Tiefenwechsel: ein relativer Pfad, der aus Tiefe *n* auflöst, braucht aus Tiefe *n+1* ein
zusätzliches `../`.

**Der reale Vorfall:** Der `git mv` von
[`welle-13-konsumenten-befunde.md`](../done/welle-13-konsumenten-befunde.md) brach **21** Verweise
auf einen Schlag — gefangen von `doc-check`, also **nach** dem `mv` statt davor. Genau der Zyklus,
den der Sensor abschaffen sollte.

**Die stille Hälfte ist das eigentliche Problem.** Dass Welle-Dateien nicht erfasst sind, steht
**nirgends** — weder im Sensor-Kopf noch in seiner Ausgabe (*„N wandernde(r) Slice(s) …"* klingt
vollständig). Ein Sensor, der eine Gattung nicht sieht und das nicht sagt, ist dieselbe Klasse wie
die False-Greens aus [slice-070](../done/slice-070-grundgesamtheit-messen.md)/[slice-071](../done/slice-071-sensor-scope-vollstaendig.md).

## 2. Betroffene Module

`tools/verify-slice-links.sh` in jedem Fall; bei einem
Struktur-Entscheid zusätzlich [`planning/README.md`](../README.md) (Welle-Lifecycle) und die
bestehende Welle-Datei.

## 3. Auszuführende Gates

`make gates`, `make verify`.

**Der Entscheid, der vor dem Bau fällt.** Die Invariante ist nicht erfüllbar — also ist die Frage
**nicht** „wie prüfe ich sie auch für Welle-Dateien", sondern „was tritt an ihre Stelle".

| Weg | Aussage | Kosten |
|---|---|---|
| **(a) Grenze ausweisen** — Sensor-Kopf + Ausgabe sagen, dass Welle-Dateien **nicht** geprüft sind, und warum | ehrlich, sofort, deckt die stille Hälfte | die Verweise brechen weiterhin; `doc-check` fängt sie nach dem `mv` |
| **(b) Gleiche Ebene herstellen** — Welle-Dateien in ein Verzeichnispaar (`wellen/aktiv/` + `wellen/done/`), dann greift die bestehende Invariante unverändert | löst die Ursache | Struktur- und Doku-Änderung; die Welle-Lifecycle-Beschreibung in [`planning/README.md`](../README.md) hängt daran |
| **(c) Repo-relative Verweise** in Welle-Dateien verlangen (`/docs/plan/...`) | tiefenunabhängig | `links_of()` filtert `^/` heute aus; ob `doc-check` sie auflöst, ist **ungeprüft** |
| **(d) Status quo** — der `git mv` zieht die Verweise nach, `doc-check` ist das Netz | keine Arbeit | die Blindstelle bleibt unbenannt; das ist der Zustand, der den Vorfall erzeugt hat |

**Neigung des Autors: (a) jetzt, (b) beim zweiten Vorkommen.** Die Evidenz ist **ein** Vorfall —
Welle-Plan-Dateien gibt es erst seit `welle-13`. Eine Verzeichnis-Umstellung auf dieser Grundlage
wäre dieselbe Vorratshaltung, die [slice-045](../open/slice-045-intern-extern-dateimenge.md) mit „null
reale Fundstellen" vertagt hat. Was **jetzt** falsch ist, ist nicht die fehlende Prüfung, sondern
die **unbenannte** Lücke: Der Sensor meldet „N wandernde Slice(s)" und klingt dabei vollständig.
Der Entscheid gehört trotzdem beim Bau begründet, insbesondere gegen (c) — das wäre die billigste
echte Lösung, falls `doc-check` repo-relative Ziele auflöst.

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| `make verify` auf dem heutigen Repo | grün, **und** die Ausgabe nennt die nicht geprüfte Gattung |
| Eine Welle-Datei mit gebrochenem Verweis (Fixture) | je nach Entscheid rot **oder** ausdrücklich als ungeprüft ausgewiesen — in **keinem** Fall stillschweigend grün |
| Selbsttest | feuert weiterhin in beiden Richtungen für Slices |

## 4. Was bewusst nicht getan wird

- **Die Welle-Datei doppelt ablegen** (flach *und* `done/`). Zwei Wahrheiten sind schlimmer als eine
  benannte Lücke.
- **`doc-check` ersetzen.** Ob ein Verweis **heute** auflöst, beantwortet es zuverlässig; dieser
  Slice betrifft nur die Frage, ob er den **Wechsel** überlebt.
- **Den Sensor für `done/`-Slices scharf schalten.** `done/` ist Endzustand und ausdrücklich
  ausgenommen ([slice-060](../done/slice-060-slice-link-invariante.md)) — daran ändert sich nichts.

## 5. DoD

- [x] Der Entscheid aus §3 ist getroffen und begründet: **(a) Grenze ausweisen**, inklusive der
      Messung zu (c) — `doc-check` löst repo-relative Ziele **auf** (fehlende melden
      `target-missing`), die Notation scheitert aber an GitHub und an null Präzedenz.
- [x] Die Ausgabe von `make verify` behauptet keine Vollständigkeit mehr, die sie nicht hat —
      sie nennt die ausgenommene Gattung **mit Zahl**, auch bei null. Beleg: die drei Proben unten.
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft.

## 6. Closure-Notiz

**Der Slice begann mit der Korrektur einer eigenen Einschätzung.**
[`welle-13-results.md`](../done/welle-13-results.md) hatte den Befund als „Ein-Zeilen-Fix am Glob"
benannt. Die erste Messung widerlegte das: Weil die Welle-Datei beim `mv` die **Verzeichnistiefe**
wechselt, löst **jede** Verweis-Form aus genau einer der beiden Positionen auf — nie aus beiden.
Ein erweiterter Glob hätte einen Sensor erzeugt, der jede Welle-Datei zu Recht rot meldet, **ohne
dass man ihn grün bekommt**. Die Invariante ist nicht unerfüllt, sie ist unerfüllbar.

**Weg (c) wurde gemessen und trotzdem verworfen — die Messung allein hätte in die Irre geführt.**
`doc-check` löst repo-relative Ziele (`/spec/…`) **vollwertig** auf; die erste Probe (existierende
Ziele, keine Beanstandung) hätte man als „funktioniert" lesen können. Erst die **Gegenprobe** mit
fehlenden Zielen zeigte, dass wirklich geprüft und nicht bloß übersprungen wird:

```text
docs/plan/planning/welle-99.md:4  /spec/gibtsnicht.md                     target-missing
docs/plan/planning/welle-99.md:5  /docs/plan/planning/done/fehlt.md       target-missing
```

Verworfen wurde die Notation dennoch, aus zwei Gründen außerhalb von d-check: GitHub löst einen
führenden `/` gegen die **Site**-Wurzel auf, nicht gegen das Repo (die Links wären im Browser
kaputt), und im ganzen Repo **inklusive** der vendored Baseline kommt sie an **null** Stellen vor.
Eine Insel-Notation für eine Datei-Gattung wäre teurer als die Lücke, die sie schließt.

**Weg (b) fiel an einer begründeten Anordnung.** „Gleiche Ebene herstellen" klingt sauber, hätte
aber die Plan-Datei aus `done/` herausgenommen — wo sie laut Prozedur **neben ihrer
Ergebnis-Notiz** liegt. Entweder wandert die Notiz mit (dann sind Wellen von Slices getrennt) oder
sie bleibt (dann sind Plan und Ergebnis getrennt). Beides ist schlechter als die benannte Lücke,
und die Evidenz ist **ein** Vorfall.

**Die drei Proben:**

```text
(1) make verify, Ist-Zustand          -> ok, "Welle-Plan-Dateien AUSGENOMMEN (aktuell 0 flach)"
(2) flache welle-99-probe.md mit kaputtem Verweis:
      verify-slice-links               -> EXIT=0, zaehlt sie sichtbar: "aktuell 1 flach"
      make doc-check                   -> EXIT=2, "welle-99-probe.md:3 done/gibtsnicht.md target-missing"
(3) Selbsttest                         -> feuert weiterhin in beiden Richtungen fuer Slices
```

Probe (2) ist die eigentliche Aussage des Slice: Der Sensor bleibt grün — **aber nicht still**. Die
Datei erscheint in seiner Zählung, und das Netz darunter greift.

**Beobachtbare Architektur-Aussage: ein Sensor darf eine Lücke haben, aber nicht verschweigen.**
Die Ausgabe nennt die ausgenommene Gattung **immer**, auch bei null aktiven Welle-Dateien. Eine
Grenze, die nur bei Gelegenheit sichtbar wird, ist keine — wer wissen will, was der Sensor abdeckt,
liest seine Ausgabe, nicht seinen Quelltext. Damit steht neben der bestehenden Ausnahme (`done/`
ist Endzustand) jetzt die zweite, gleichrangig formuliert.

**Lerneintrag — Form: benannte Spec-Lücke.** Als Prüfsatz: *Bevor ein Sensor auf eine neue
Datei-Gattung erweitert wird, ist zu prüfen, ob seine Invariante für sie überhaupt **erfüllbar**
ist — sonst entsteht ein Gate, das niemand grün bekommt.* Die Invariante von
[slice-060](../done/slice-060-slice-link-invariante.md) trägt ihre Voraussetzung im Kommentar
(„alle vier Lifecycle-Verzeichnisse liegen auf derselben Ebene"); dass sie eine **Voraussetzung**
ist und keine Beschreibung, fiel erst auf, als eine Gattung ohne sie auftauchte. **Zu prüfen wäre**,
ob die übrigen `verify-*`-Sensoren ähnliche stillschweigende Voraussetzungen tragen —
`verify-closure-notes` und `verify-ac-form` sind die Kandidaten.

**Mitgenommen:** [`planning/README.md`](../README.md) behauptete noch, die Closure-Prozedur sei
„deklariert, aber noch nicht belegt". Zwei Durchläufe später ist das falsch; korrigiert, mit
Zeigern auf beide Ergebnis-Notizen.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
