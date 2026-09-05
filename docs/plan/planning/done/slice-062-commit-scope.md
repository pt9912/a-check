# slice-062 — Commit-Scope `(planning)`: Konvention, dann Sensor

**Status:** der Zustand ist das **Verzeichnis** dieser Datei
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5) — dieses Feld führt ihn bewusst **nicht** doppelt.
**Welle:** welle-12-regelwerk-migration.
**Deckt:** die offene Sensor-Hälfte von `SL-003` und die dort benannte
Spec-Lücke aus [slice-061](../done/slice-061-steering-loop-eintraege.md).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser — und eine Korrektur an slice-061

slice-061 hielt den Sensor für „fertig und nachweislich rauschfrei", geprüft an **acht** Commits.
Die Prüfung gegen die **gesamte** Historie widerlegt das für die allgemeine Form: die Hypothese
„ein `docs(...)`-Commit ändert keine ausführbaren oder normativen Artefakte" erzeugt **31 Treffer
bei 193** `docs(...)`-Commits — `docs(spec)` ändert legitim `spec/`, `docs(adr)` legitim ADRs. Als
allgemeine Regel ist das Rauschen, kein Sensor.

**Rauschfrei ist sie ausschließlich scope-spezifisch.** Für den Scope `(planning)` gilt über die
ganze Historie: **fünf** Treffer bei **74** Commits, und alle fünf sind echte
Etiketten-Diskrepanzen — keine einzige Fehlmeldung:

| Commit | Betreff (gekürzt) | fremde Pfade |
|---|---|---|
| `615e37f` | slice-049 Closure | `tools/`, `Makefile`, `AGENTS.md`, `.harness/skills/` |
| `f57289d` | slice-050 Roadmap-Link | `.claude/commands/`, `docs/user/` |
| `d436da9` | slice-052 Folge-Slice-Link | `AGENTS.md` |
| `7faa708` | slice-043 in-progress → done | `docs/reviews/` |
| **`f0e7805`** | **SL-003 und SL-004** | `.claude/commands/`, `docs/plan/steering-loop.md` |

**Der fünfte ist der Commit von [slice-061](../done/slice-061-steering-loop-eintraege.md) selbst** —
also der Commit, der `SL-003` anlegt, begeht `SL-003`. Damit steht der Zähler bei **fünf**, und der
Guide aus slice-061 hat seinen ersten Vorfall nicht verhindert: dieselbe Lage wie bei `SL-001` vor
slice-057 und `SL-002` vor slice-060. *Inferential feedforward* wirkt gegen Unwissen, nicht gegen
Routine — zum dritten Mal belegt.

## 2. Betroffene Module

- [`AGENTS.md`](../../../../AGENTS.md) §5 — die Konvention; §4 — das Target.
- `tools/commit-scope-check.sh` + [`Makefile`](../../../../Makefile) — der Sensor.
- [`.claude/hooks/pretooluse-command-guard.sh`](../../../../.claude/hooks/pretooluse-command-guard.sh)
  — neues Prüf-Target in die `GATES`-Liste (der Drift-Wächter fordert es ein).
- [`.github/workflows/ci.yml`](../../../../.github/workflows/ci.yml) — Lauf über die Commit-Range,
  neben `trace-check`.
- `docs/plan/steering-loop.md` — `SL-003` auf fünf Vorfälle, Antwort
  gebaut.

**Zwei Schichten:** Gate-/Werkzeug-Schicht und Harness-/Planungs-Doku.

## 3. Auszuführende Gates

`make gates` und `make verify`, Ausgabe je in eine Datei, Exit-Code getrennt geprüft.

**Negativ-Proben:**

1. Die fünf realen Vorfälle müssen erkannt werden, wenn man die Regel auf sie anwendet — geprüft
   über eine explizite Range statt über eine erfundene Fixture.
2. Die übrigen 69 `(planning)`-Commits müssen durchgehen (kein Rauschen).
3. Ein konstruierter Commit mit Scope `(planning)`, der `Makefile` ändert, muss rot werden; einer,
   der nur `docs/plan/planning/` ändert, grün.
4. Der Drift-Wächter aus slice-059 muss das neue Target einfordern.

**Grandfathering ohne Stichtags-SHA:** ein Commit wird an der Regel gemessen, **die zu seinem
Zeitpunkt galt** — der Sensor prüft, ob `AGENTS.md` im jeweiligen Commit die Konvention bereits
trägt, und überspringt ihn sonst. Das ist präziser als ein Datum und braucht keinen Hash, den es
beim Schreiben des Skripts noch nicht gibt.

## 4. Was bewusst nicht getan wird

- **Keine Ausweitung auf andere Scopes.** `(spec)`, `(adr)`, `(harness)`, `(review)` bleiben
  ungeregelt: die Messung zeigt 31 Treffer bei 193 Commits, die gelebte Praxis ist dort heterogen
  (Mehrfach-Scopes wie `docs(review),test(core,cli)`), und eine Regel, die den Bestand massenhaft
  bricht, wird abgeschaltet statt befolgt. Wenn ein zweiter Scope auffällig wird, ist die Messung
  zu wiederholen — nicht die Regel zu raten.
- **Keine Korrektur der fünf Commit-Nachrichten.** Historie wird nicht umgeschrieben.
- **Kein Pre-Commit-Hook.** Das Target ist der Sensor; ein Hook wäre ein zweiter Ort für dieselbe
  Regel und driftet.

## 5. DoD

- [x] [`AGENTS.md`](../../../../AGENTS.md) §5 deklariert: ein Commit mit Scope `(planning)` berührt
      ausschließlich `docs/plan/planning/`; mit Begründung und der Messung, warum die Regel dort
      und nur dort gilt.
- [x] `make commit-scope-check` setzt sie über eine Commit-Range durch, misst jeden Commit an der
      zu seinem Zeitpunkt geltenden Fassung und trägt einen Selbsttest; belegt durch die vier
      Proben aus §3.
- [x] Target in [`AGENTS.md`](../../../../AGENTS.md) §4, in der `GATES`-Liste und in der CI;
      `SL-003` trägt fünf Vorfälle und die gebaute Antwort.
      `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft.

## 6. Closure-Notiz

**Geliefert:** die Commit-Scope-Konvention in [`AGENTS.md`](../../../../AGENTS.md) §5 und ihr
Sensor `make commit-scope-check` — in der `GATES`-Liste, in §4 deklariert und in der CI über die
Commit-Range neben `trace-check`. `SL-003` trägt fünf Vorfälle und ist
damit in beiden Quadranten durchgesetzt.

**Lerneintrag — Form: geschärfte Regel.**
> **Eine Stichprobe belegt Rauschfreiheit nur für die Menge, aus der sie gezogen ist.** slice-061
> nannte den Sensorentwurf „nachweislich rauschfrei" — geprüft an **acht** Commits, alle aus der
> Migrations-Kette. Über die volle Historie fiel die allgemeine Form sofort: 31 Treffer bei 193
> `docs(...)`-Commits, *weil* `docs(spec)` legitim `spec/` und `docs(adr)` legitim ADRs ändert —
> Fälle, die in der Kette schlicht nicht vorkamen. Die Behauptung war nicht falsch gemessen,
> sondern falsch **verallgemeinert**. Prüfsatz: *bevor eine Sensor-Hypothese „rauschfrei" heißt,
> die Grundgesamtheit benennen, gegen die sie geprüft wurde — und prüfen, ob der Bestand außerhalb
> davon andere Fälle enthält.* Getragen hat die Regel am Ende nur scope-spezifisch, und genau so
> ist sie deklariert.

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. Drei Läufe mit unterschiedlichem Ergebnis, je belegt: ein Commit mit Scope `(planning)`, der
   `Makefile` ändert, ergibt **EXIT=1** mit Nennung der fremden Pfade; ein `(planning)`-Commit, der
   nur `docs/plan/planning/` ändert, **EXIT=0**; über die **gesamte** Historie **EXIT=0** mit
   `74 vor Einfuehrung der Regel (grandfathered)` — der Bestand wird nicht rot, ohne dass ein
   Stichtags-Hash nötig wäre.

**Selbstanwendung, unbequem:** der fünfte `SL-003`-Vorfall ist `f0e7805` — der Commit, mit dem
slice-061 den Eintrag `SL-003` **anlegte**. Der Guide aus jenem Slice hat seinen ersten Vorfall
nicht verhindert. Das ist zum dritten Mal dieselbe Beobachtung (`SL-001` vor slice-057, `SL-002`
vor slice-060) und stützt die Modul-09-Regel deutlicher als jedes Zitat: *inferential feedforward*
wirkt gegen Unwissen, nicht gegen Routine. Wäre der Sensor in slice-061 gebaut worden, hätte er
den eigenen Commit gefangen.

**Folge-Slices:** die verbliebene Doku-Arbeit der Review-Serie (slice-048-Korrekturen,
Status-Felder, INFO-Kategorie, Zahlenpaar). Aus diesem Slice selbst: keine — die Ausweitung auf
weitere Scopes ist ausdrücklich an eine neue Messung gebunden, nicht an einen Folge-Slice.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
