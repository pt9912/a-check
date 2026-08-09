# slice-074 — `doc-targets` prüft, was es zu prüfen behauptet

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** Befund vom 2026-08-09 bei der Messung zu
[slice-073](../open/slice-073-dcheck-statt-eigenbau.md); Bezug zu
[MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert).
**Bezug:** Roadmap-Zeile *Aktuelle Welle* in der [Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**Mechanismus: ein als Gate dokumentiertes Fremdtool-Modul ist nicht konfiguriert und prüft
deshalb nichts.** Neu gegenüber der Gruppe A: die Ursache liegt nicht im Eigenbau, sondern in
einer fehlenden Konfiguration — der Sensor ist fremd, korrekt und funktionsfähig, er bekommt nur
nie gesagt, worauf er schauen soll.

[`AGENTS.md`](../../../../AGENTS.md) §4 führt:

> `make doc-targets` — Deklarations-Konsistenz Doku ↔ Build-Targets (Modul `targets`,
> DC-FA-TGT-001; neu in v0.51.1)

Gemessen am 2026-08-09: In [`.d-check.yml`](../../../../.d-check.yml) gibt es **keinen**
`targets:`-Block. `d-check.mk:66` aktiviert das Modul per `--enable targets`, es findet keine
Konfiguration und meldet **`0 Befund(e)`, Exit 0** — auch dann, wenn `AGENTS.md` ein Target
behauptet, das es nicht gibt. Die Gegenprobe im selben Lauf:

```
Phantom-Target in AGENTS.md eingetragen:
  make doc-targets       EXIT=0   "0 Befund(e)"      <- prueft nichts
  make gate-consistency  EXIT=2   "dokumentiert 'make phantom-target'"
```

Mit Konfiguration prüft dasselbe Modul beide Richtungen korrekt:

```
targets:
  makefiles: [Makefile, d-check.mk]
  doc-tables: [AGENTS.md]
  authority: AGENTS.md
  exempt-targets: [help, build, compile, hooks, arch-graph]

Phantom in AGENTS.md         -> gate-phantom,       EXIT 2
undokumentiertes Make-Target -> gate-undocumented,  EXIT 2
```

`makefiles` muss **beide** Fragmente nennen: mit `[Makefile]` allein meldet das Modul
`doc-tracked`, `doc-targets` und `doc-help` als `gate-phantom` — sie stehen in
[`d-check.mk`](../../../../d-check.mk). `gate-consistency` liest seit jeher beide.

**Warum das zählt.** Ein Gate, das in `AGENTS.md` §4 steht und nichts prüft, ist eine
Harness-Lüge — dieselbe Klasse wie `F-1`, `F-5` und `F-12` aus dem
[Review-Report](../../../reviews/2026-08-09-welle-12-unabhaengig.md), nur an einer Stelle, die
kein Review betrachtet hat: der Konfiguration eines Fremdwerkzeugs.

**Und die Ablösung war die erklärte Absicht.** d-checks `slice-063` (Modul `targets`, Release
**v0.38.0**) führt als DoD-Punkt ausdrücklich einen *„Paritäts-Mutations-Beleg vs.
`gate-consistency.sh`"* — das Modul wurde gegen genau dieses Skript gebaut und geprüft. a-check
hat das Target in `AGENTS.md` §4 aufgenommen, die Konfiguration aber nie nachgezogen und fährt
seither den Eigenbau weiter; der aktuelle Pin steht bei `v0.51.1`, dreizehn Minor-Versionen
später. Dieser Slice holt keine neue Fähigkeit ins Repo, er nimmt eine seit `v0.38.0`
bereitstehende endlich in Betrieb.

## 2. Betroffene Module

Zwei Schichten:

1. **[`.d-check.yml`](../../../../.d-check.yml)** — der fehlende `targets:`-Block.
2. **[`AGENTS.md`](../../../../AGENTS.md)** §4 — die Zeile muss den tatsächlichen Status nennen.
   Das Muster dafür existiert: `make regelwerk-check` ist dort ausdrücklich als **„kein Gate —
   Wartung"** gekennzeichnet.

## 3. Auszuführende Gates

`make doc-targets`, `make gate-consistency`, `make gates`.

**Negativ-Proben** — beide bereits gefahren, im Slice zu wiederholen:

| Probe | Erwartung |
|---|---|
| Phantom-Target in einer `AGENTS.md`-Tabellenzeile | `gate-phantom`, Exit ≠ 0 |
| Rezept-Target im `Makefile` ohne `AGENTS.md`-Zeile | `gate-undocumented`, Exit ≠ 0 |
| unveränderter Bestand | Exit 0 — die Konfiguration darf keine Falsch-Positive erzeugen |

Die dritte Zeile ist nicht Beiwerk: mit `makefiles: [Makefile]` allein wären es drei.

## 4. Was bewusst nicht getan wird

- **Die Ablösung von `gate-consistency` (1)+(2).** Dass beide Prüfer dasselbe melden, ist in
  [slice-073](../open/slice-073-dcheck-statt-eigenbau.md) gemessen; die Entscheidung, den Eigenbau
  zu entfernen, gehört dorthin und in dessen Folge-Slice. Bis dahin läuft die Prüfung **doppelt** —
  bewusst, und hier benannt statt stillschweigend.
- **`doc-targets` in `gates` hängen.** Solange der Eigenbau dieselbe Invariante im Aggregat prüft,
  wäre das ein zweiter Lauf ohne zusätzliche Aussage. Die Einhängung ist Teil der Ablösung, nicht
  dieses Slice.
- **Der CR an d-check**, dass ein aktiviertes, aber unkonfiguriertes Modul still `0 Befund(e)`
  meldet statt fail-closed abzubrechen. Das ist ein Fremdrepo-Befund; er gehört in den CR-Text von
  [slice-073](../open/slice-073-dcheck-statt-eigenbau.md), nicht in eine lokale Korrektur.
- **Die übrigen d-check-Module auf Konfigurations-Lücken prüfen.** Naheliegend, aber ohne Messung
  spekulativ — und wenn, dann als eigener Schnitt mit eigener Zählung.

## 5. DoD

- [x] `.d-check.yml` trägt einen `targets:`-Block, und `make doc-targets` meldet beide Richtungen.
      Beleg: Phantom-Target in `AGENTS.md` → Exit 2 · `phantom-target gate-phantom`;
      undokumentiertes Rezept-Target → Exit 2 · `neues-gate gate-undocumented`. Vor diesem Slice
      lieferten **beide** Fixtures Exit 0 mit „0 Befund(e)".
- [x] Die `AGENTS.md`-§4-Zeile nennt den tatsächlichen Status des Targets (im Aggregat oder
      nicht). Beleg: `make gate-consistency` grün, das Doku ↔ Makefile prüft.
- [x] `make gates` grün und **ohne neue Falsch-Positive** — **Ausgabe in eine Datei**, Exit-Code
      getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** [`.d-check.yml`](../../../../.d-check.yml) trägt einen `targets:`-Block mit
`makefiles: [Makefile, d-check.mk]`; die `AGENTS.md`-§4-Zeile weist das Target als **nicht im
`gates`-Aggregat** aus und nennt den Grund. `make doc-targets` prüft damit, was es zu prüfen
behauptet.

**Lerneintrag — Form: geschärfte Regel.** Als Prüfsatz: *Ein Gate einzubinden heißt nicht, es zu
konfigurieren — und ein unkonfiguriertes Modul meldet grün, nicht „unklar".*

**Die Ursache** ist eine Lücke zwischen zwei Arbeitsschritten, die niemandem gehörte: das Target
kam mit dem d-check-Pin-Bump ins Repo und wurde ordnungsgemäß in `AGENTS.md` §4 dokumentiert — der
Konfigblock, ohne den es wirkungslos ist, gehörte weder zum Pin-Bump noch zur Doku-Zeile. Kein
Sensor konnte es fangen: `gate-consistency` prüft, ob ein dokumentiertes Target **existiert**,
nicht ob es etwas **tut**. Und d-check selbst bricht bei aktiviertem, aber unkonfiguriertem Modul
nicht fail-closed ab, sondern meldet `0 Befund(e)`.

**Besonders belastend:** das Modul war laut d-checks `slice-063` (v0.38.0) ausdrücklich als
Ablösung von `gate-consistency.sh` gebaut, **mit Paritäts-Mutations-Beleg gegen genau dieses
Skript**. Es stand dreizehn Minor-Versionen bereit und lief die ganze Zeit ins Leere.

**Zwei beobachtbare Closure-Kriterien:**

1. Ein Phantom-Target in einer `AGENTS.md`-Tabellenzeile macht `make doc-targets` rot
   (`gate-phantom`), ein undokumentiertes Rezept-Target ebenso (`gate-undocumented`) — beide
   Fixtures lieferten vor diesem Slice Exit 0.
2. Der unveränderte Bestand bleibt bei allen drei betroffenen Läufen grün (`doc-targets`,
   `gate-consistency`, `doc-check`) — die Konfiguration tauscht keine Falsch-Negative gegen
   Falsch-Positive.

**Folge-Slices:** [slice-073](../open/slice-073-dcheck-statt-eigenbau.md) entscheidet, ob
`gate-consistency` (1)+(2) entfällt und `doc-targets` ins Aggregat wandert. Bis dahin läuft die
Prüfung **doppelt** — bewusst und in `AGENTS.md` benannt.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
