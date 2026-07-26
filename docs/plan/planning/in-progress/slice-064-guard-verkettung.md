# slice-064 — SL-001: der Commit hing an `;` statt an `&&`

**Status:** der Zustand ist das **Verzeichnis** dieser Datei
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5) — dieses Feld führt ihn bewusst **nicht** doppelt.
**Deckt:** den sechsten Vorfall von [`SL-001`](../../steering-loop.md), belegt in `1a9f270`.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser — ein Vorfall aus dem eigenen Lauf

Beim Abschluss von [slice-063](../done/slice-063-doku-korrekturen.md) ging ein Commit mit
**rotem** `make gates` heraus (`1a9f270`, Exit 2, zwei `id-unlinked`-Befunde). Der Ablauf war:

```sh
make gates > $SP/g63.log 2>&1; echo "gates EXIT=$?"
…
git add -A && git commit -q -F - <<'EOF'
```

Der Lauf war korrekt in eine Datei umgeleitet und der Exit-Code ausgegeben — genau wie
[`SL-001`](../../steering-loop.md) es vorschreibt. Gelesen wurde er nicht, und der Commit folgte
im selben Tool-Call.

**Die Guard-Regel 2 sieht das nicht.** `pipeViolation()` prüft den Rest auf `git commit`
**ausschließlich**, wenn der Separator nach dem Gate-Lauf `&&` ist:

```js
if (sep === "&&") {
  const rest = parts.slice(i + 2).join("");
  if (/\bgit\s+commit\b/.test(rest)) return true;
}
```

Bei `;` oder Zeilenumbruch fällt die Prüfung durch. Damit ist die *computational* Hälfte von
`SL-001` seit slice-057 lückenhaft, und die Lücke trifft genau die Schreibweise, die der Guide
selbst nahelegt: wer die Umleitung in eine Datei befolgt, verkettet danach mit `;` statt mit `&&`.

## 2. Betroffene Module

- [`.claude/hooks/pretooluse-command-guard.sh`](../../../../.claude/hooks/pretooluse-command-guard.sh)
  — Regel 2, Separator-Bedingung.
- [`docs/plan/steering-loop.md`](../../steering-loop.md) — `SL-001` auf sechs Vorfälle.

**Zwei Schichten:** Durchsetzungsschicht und Planungs-Doku.

## 3. Auszuführende Gates

`make gates` (enthält `guard-selftest`) und `make verify`, Ausgabe je in eine Datei, Exit-Code
getrennt geprüft — **und diesmal gelesen, bevor committet wird**.

Negativ-Proben:

1. `make gates > datei 2>&1; git commit -m x` muss **abgelehnt** werden (der reale Fall).
2. Dasselbe mit Zeilenumbruch statt `;` muss abgelehnt werden.
3. `make gates > datei 2>&1; echo "EXIT=$?"` **ohne** Commit muss weiter durchgehen — sonst wäre
   die vorgeschriebene Form selbst blockiert.
4. Live-Probe im echten Tool-Call, nicht nur im Selbsttest.

## 4. Was bewusst nicht getan wird

- **Keine Ausnahme für geprüfte Verzweigungen.** `if make gates; then git commit; fi` wäre
  inhaltlich korrekt, wird aber ebenfalls abgelehnt: den Exit-Code wirklich zu lesen ist von außen
  nicht von „ausgeben und ignorieren" zu unterscheiden, und ein Guard, der das zu erraten
  versucht, wird unzuverlässig. Der Commit gehört in einen eigenen Aufruf — das kostet eine Zeile
  und ist die Praxis, die diese Serie ohnehin durchgehend verwendet hat.
- **Keine Ausweitung auf andere Folgekommandos** (`git push`, `git tag`). Sie sind nicht
  aufgetreten; eine Regel ohne Vorfall ist geraten, nicht gemessen.

## 5. DoD

- [x] Regel 2 lehnt einen Gate-Lauf mit nachfolgendem `git commit` in **jeder** Verkettung ab
      (`;`, Zeilenumbruch, `&&`); die vorgeschriebene Form ohne Commit bleibt erlaubt.
- [x] Der reale Fall aus `1a9f270` liegt als Selbsttest-Fixture vor und ist zusätzlich im echten
      Tool-Call belegt.
- [x] [`SL-001`](../../steering-loop.md) trägt den sechsten Vorfall samt geschlossener Lücke;
      `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft.

## 6. Closure-Notiz

**Geliefert:** Regel 2 des PreToolUse-Guard lehnt einen Gate-Lauf mit nachfolgendem `git commit`
jetzt in **jeder** Verkettung ab, nicht nur nach `&&`. Der reale Fall liegt als
Selbsttest-Fixture vor und ist zusätzlich im echten Tool-Call belegt. Die Ablehnungs-Meldung nennt
die Verkettungen ausdrücklich und verlangt den Commit als **zweiten** Aufruf.
[`SL-001`](../../steering-loop.md) trägt den sechsten Vorfall samt geschlossener Lücke.

**Lerneintrag — Form: geschärfte Regel.**
> **Ein Guard, der nur die auffällige Schreibweise kennt, treibt den Fehler in die unauffällige.**
> Regel 2 fing `make <gate> && git commit` — die Form, in der der Zusammenhang *sichtbar* ist. Der
> reale Rückfall kam in der Form, die der eigene Guide nahelegt: Umleitung in eine Datei, dann
> `;`, dann der Commit. Der Fehler wurde nicht seltener, sondern unsichtbarer, *weil* die
> vorgeschriebene Schreibweise ihn nicht mehr an `&&` band. Prüfsatz für jede
> Durchsetzungs-Regel: *welche Schreibweise empfiehlt der zugehörige Guide — und deckt die Regel
> genau diese ab?* Ein Sensor, der die empfohlene Form nicht prüft, prüft die falsche.

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt, und der
   Exit-Code diesmal **vor** dem Commit gelesen.
2. Der Aufruf `make gates > datei 2>&1; git commit -m x` wird im echten Tool-Call **abgelehnt**;
   `make gates > datei 2>&1; echo "EXIT=$?"` ohne Commit läuft weiter durch. Beide Fälle
   zusätzlich als dauerhafte Selbsttest-Fixture.

**Zur Herkunft des Vorfalls, unbequem:** er stammt aus dem Abschluss von slice-063 — also aus
einem Slice dieser Serie, der Belegtreue zum Gegenstand hatte. Zwischen der Regel („den Exit-Code
getrennt prüfen") und ihrer Befolgung lag der Umstand, dass ich den ausgegebenen Code nicht
gelesen habe. Genau dagegen ist ein *computational*-Sensor gebaut, und genau dort hatte er ein
Loch. Das ist der vierte Beleg in dieser Serie für dieselbe Modul-09-Aussage — und der erste, bei
dem nicht der Guide zu schwach war, sondern der Sensor zu eng.

**Zweiter Fund im selben Slice — der Sensor blockierte seinen eigenen Commit.** Die verschärfte
Regel lehnte die Commit-Message dieses Slice ab, weil sie den abgelehnten Aufruf **zitiert**: ein
Heredoc ist für den Guard bis dahin Teil des Kommandos gewesen. Behoben wurde der **Sensor**, nicht
der Text — `stripHeredoc()` blendet Heredoc-Inhalte vor der Auswertung aus, analog zu
`stripQuoted()`. Das ist der **vierte** Vorfall von [`SL-004`](../../steering-loop.md) und der
erste außerhalb eines Doku-Sensors; die Regel „Zitat-Kontexte gehören von Anfang an ausgeblendet"
gilt damit für die Durchsetzungsschicht genauso. Die Grenze bleibt dieselbe wie bei
Sub-Shell-Strings: ein Heredoc an einen Interpreter (`bash <<EOF`) entgeht ebenfalls —
Stolperdraht, keine Sandbox.

**Folge-Slices:** keine.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
