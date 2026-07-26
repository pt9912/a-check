# slice-060 — SL-002: Verweise, die den Lifecycle überleben

**Status:** der Zustand ist das **Verzeichnis** dieser Datei
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5) — dieses Feld führt ihn bewusst **nicht** doppelt.
**Deckt:** die offene *computational* Hälfte von [`SL-002`](../../steering-loop.md) (neun Vorfälle).
**Bezug:** dritter Fix-Schnitt der Review-Serie; [slice-059](../done/slice-059-durchsetzungs-luecken.md)
hat die *inferential* Hälfte geliefert (Schritt 9 des Workflow-Skeletts).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

`SL-002` zählt **neun** Vorfälle: sieben in der Migrations-Kette, zwei weitere in
[slice-058](../done/slice-058-sensor-praezision.md) — die letzten beiden **nach** dem Guide, einer
davon beim Schreiben der Notiz über genau diesen Befund. Damit gilt hier dieselbe Diagnose wie bei
`SL-001` vor [slice-057](../done/slice-057-steering-loop.md): *inferential feedforward* wirkt gegen
Unwissen, nicht gegen Routine. Die Regel ist nach `modul-09` **halb durchgesetzt**, solange die
computational Hälfte fehlt.

**Die Entwurfsfrage, die den Sensor bisher aufgehalten hat, löst sich durch Umdrehen.** Ein Sensor
kann nicht *vorhersagen*, wohin eine Datei verschoben wird — er muss es auch nicht. Alle vier
Lifecycle-Verzeichnisse liegen auf **derselben Ebene**; die prüfbare Eigenschaft ist deshalb keine
Vorhersage, sondern eine **Invariante**:

> Ein relativer Verweis in einer Slice-Datei muss aus **jedem** Lifecycle-Verzeichnis auflösen.

Das deckt mehr ab als der Eintragstitel: nicht nur den `git mv`, sondern auch das **Anlegen** aus
einer Vorlage anderer Verzeichnistiefe — der Fall, der slice-058 zuerst traf und den ein reiner
`git mv`-Wächter nie gesehen hätte.

**Gemessener Bestand:** in `open/` liegen drei Slices, alle sauber; `next/` und `in-progress/` sind
leer. Der Sensor kann sofort scharf gestellt werden — **kein Grandfathering nötig**. Die 22
Dateien in `done/` mit zustandsabhängigen Verweisen bleiben ausgenommen, und zwar sachlich: `done/`
ist Endzustand (`AGENTS.md` §5 kennt fünf Übergänge, keiner führt hinaus), ein künftiger `git mv`
kann ihre Verweise nicht mehr brechen. Ob sie *heute* auflösen, prüft `doc-check` — das ist eine
andere Frage und bereits beantwortet.

## 2. Betroffene Module

- `tools/verify-slice-links.sh` — neu, der Sensor.
- [`Makefile`](../../../../Makefile) — `verify-slice-links`, eingehängt in `verify`.
- [`AGENTS.md`](../../../../AGENTS.md) §4 — Target-Deklaration (`gate-consistency` erzwingt sie).
- [`.claude/hooks/pretooluse-command-guard.sh`](../../../../.claude/hooks/pretooluse-command-guard.sh)
  — das neue Prüf-Target in die `GATES`-Liste; der Drift-Wächter aus slice-059 fordert es ein.
- [`docs/plan/steering-loop.md`](../../steering-loop.md) — `SL-002` bekommt seine Antwort.

**Zwei Schichten:** Gate-/Werkzeug-Schicht und Planungs-Doku.

## 3. Auszuführende Gates

`make gates` und `make verify`, Ausgabe je in eine Datei, Exit-Code getrennt geprüft.

Negativ-Proben:

1. Eine Slice-Datei in `in-progress/` mit einem **präfixlosen** Verweis auf `roadmap.md` muss den
   Sensor **rot** machen — exakt der slice-058-Fall.
2. Dieselbe Datei mit `../in-progress/roadmap.md` muss ihn **grün** lassen.
3. Der Selbsttest muss beide Richtungen dauerhaft mitführen, damit ein totes Muster auffällt.
4. **Drift-Probe von slice-059:** wird das neue Target nicht in die `GATES`-Liste aufgenommen,
   muss `guard-selftest` rot werden. Das ist zugleich der erste echte Einsatz jenes Wächters.

## 4. Was bewusst nicht getan wird

- **Kein Umschreiben der 22 `done/`-Dateien.** Ihre Verweise lösen heute auf und können nicht mehr
  brechen; ein Umbau wäre Geschichts-Politur ohne Erkenntnisgewinn.
- **Keine Prüfung, ob ein Verweisziel inhaltlich das richtige ist.** Der Sensor prüft
  Auflösbarkeit über Verzeichniswechsel, nicht Bedeutung — Review-Sache.
- **Keine Ausweitung auf `roadmap.md` selbst oder auf `docs/reviews/`.** Beide wandern nicht; die
  Invariante hätte dort keinen Gegenstand.

## 5. DoD

- [x] `make verify-slice-links` prüft für `open/`, `next/` und `in-progress/`, dass jeder relative
      Verweis aus **allen vier** Lifecycle-Verzeichnissen auflöst; `done/` ist mit Begründung
      ausgenommen. Selbsttest führt beide Richtungen (präfixlos rot, `../`-Form grün).
- [x] Das Target hängt an `make verify`, steht in [`AGENTS.md`](../../../../AGENTS.md) §4 und in
      der `GATES`-Liste des Guard; belegt durch grünen `guard-selftest` **und** eine Drift-Probe,
      die bei fehlendem Eintrag rot wird.
- [x] [`SL-002`](../../steering-loop.md) trägt die gebaute Antwort mit Beleg; `make gates` und
      `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** `make verify-slice-links` — die computational Hälfte von `SL-002`. Der Sensor prüft
für die wandernden Verzeichnisse, dass jeder relative Verweis aus **jedem** Lifecycle-Verzeichnis
auflöst, hängt an `make verify`, steht in [`AGENTS.md`](../../../../AGENTS.md) §4 und in der
`GATES`-Liste des Guard. [`SL-002`](../../steering-loop.md) ist damit in **beiden** Quadranten
durchgesetzt und trägt den Beleg.

**Lerneintrag — Form: geschärfte Regel.**
> **Eine Prüfung, die eine Vorhersage zu brauchen scheint, lässt sich oft als Invariante
> formulieren — und wird dadurch erst baubar.** `SL-002` blieb über zwei Slices offen, *weil* die
> Aufgabe als „melde vor dem `git mv`, welche Verweise brechen" gelesen wurde: ein Sensor müsste
> dafür wissen, wohin verschoben wird, und `make verify` weiß das nicht. Die Umkehrung löst es
> ohne jede Zusatzinformation — alle vier Lifecycle-Verzeichnisse liegen auf derselben Ebene, also
> ist die prüfbare Eigenschaft *„der Verweis löst aus jedem von ihnen auf"*. Der Sensor wurde
> dadurch nicht nur möglich, sondern **stärker** als die ursprüngliche Formulierung: er erfasst
> auch das Anlegen aus einer Vorlage anderer Verzeichnistiefe, also die Vorfälle acht und neun,
> die ein reiner `git mv`-Wächter nie gesehen hätte. Prüfsatz: *bevor ein Sensor an fehlendem
> Kontext scheitert, prüfen, ob die Eigenschaft ohne diesen Kontext formulierbar ist — eine
> Invariante braucht keinen Zeitpunkt.*

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt; `verify` führt
   jetzt fünf Sensoren statt vier.
2. Beide Richtungen des Sensors mit je einem Lauf belegt: eine Slice-Datei mit präfixlosem
   `[Roadmap](roadmap.md)` ergibt **EXIT=1** mit benanntem Zielverzeichnis; dieselbe Datei mit
   `../in-progress/roadmap.md` ergibt **EXIT=0**. Beide Fälle liegen zusätzlich als dauerhafte
   Selbsttest-Fixture im Skript.

**Der Sensor hat beim ersten scharfen Lauf sich selbst gemeldet — als Fehlalarm.** Das
Slice-Dokument *zitiert* die falsche Form in Backticks („Falsch ist `[Roadmap](roadmap.md)`"), und
die Link-Extraktion sah darin einen Verweis. Behoben wurde der **Sensor**, nicht das Dokument:
Code-Blöcke und Inline-Code werden jetzt vor der Extraktion entfernt, denn Text *über* einen
Verweis ist kein Verweis. Das ist derselbe Fehlalarm-Typ wie bei
[slice-050](../done/slice-050-verify-schicht.md) (Kursiv-Regex traf substanziellen Text) und
[slice-057](../done/slice-057-steering-loop.md) (Muster im Argument-String) — offenbar der
Standardfehler beim Bau eines Doku-Sensors in diesem Repo, und beim dritten Mal wäre ein eigener
Steering-Loop-Eintrag fällig. Die Fixture dazu liegt dauerhaft im Selbsttest und trifft das Muster
**beinahe**, kann es also wirklich prüfen (Lehre aus slice-058).

**Der Drift-Wächter aus [slice-059](../done/slice-059-durchsetzungs-luecken.md) hat seinen ersten
echten Einsatz bestanden — unaufgefordert.** Nach dem Eintragen von `verify-slice-links` ins
`Makefile`, aber vor dem Eintrag in die `GATES`-Liste, wurde `make guard-selftest` **rot**:
`Pruef-Target 'verify-slice-links' fehlt in der GATES-Liste`. Der Wächter wurde einen Slice zuvor
gegen genau diese Drift gebaut, und der Fall trat beim nächsten Anlass ein, ohne dass jemand daran
denken musste. Das ist der Unterschied zwischen einer Liste, die gepflegt werden *soll*, und
einer, deren Pflege eingefordert wird.

**Folge-Slices:** slice-061 (Belegtreue der Planungs-Doku: `SL-003` für das Etiketten-Muster,
slice-048-Korrekturen, Status-Felder) — die aus slice-060 verschobene Doku-Arbeit. Aus diesem
Slice selbst: keine.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
