# slice-061 — Zwei Muster über der Schwelle: `SL-003` und `SL-004`

**Status:** der Zustand ist das **Verzeichnis** dieser Datei
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5) — dieses Feld führt ihn bewusst **nicht** doppelt.
**Deckt:** **F-2** aus
[`2026-07-26-etappe-d-slice-052-053-054.md`](../../../reviews/2026-07-26-etappe-d-slice-052-053-054.md)
(Etiketten-Muster ohne Eintrag) und den in
[slice-060](../done/slice-060-slice-link-invariante.md) festgehaltenen dritten Fehlalarm-Vorfall.
**Bezug:** vierter Fix-Schnitt der Review-Serie.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

Zwei Fehlermuster haben die 3×-Schwelle aus [`AGENTS.md`](../../../../AGENTS.md) §5 gerissen und
haben keinen Eintrag. Die Regel dort ist unmissverständlich: ab dem dritten gleichartigen Vorfall
ist es eine **Harness-Lücke** und verlangt einen Guide oder Sensor — „besser aufpassen" ist keine
Antwort. Ein Muster ohne Eintrag wird beim vierten Mal wieder als Einzelfall behandelt; genau davor
schützt der Kanal.

**Muster A — Commit-Betreff bezeichnet nicht die enthaltene Arbeit.** Drei Commits der
Migrations-Kette tragen substanzielle Arbeit unter einem Betreff, der sie nicht nennt; zwei
`feat`-Commits nennen umgekehrt Substanz, die sie nicht enthalten. Gefunden im Review, dort mit
einem Abgleich aller 36 Commits belegt.

**Muster B — ein neuer Doku-Sensor meldet im ersten Lauf sein eigenes Umfeld.** Dreimal in Folge
hat ein frisch gebauter Sensor Text beanstandet, der *über* seinen Prüfgegenstand spricht, statt
ihn zu sein. Jedes Mal war die Korrektur dieselbe: Zitat-Kontexte ausblenden.

## 2. Betroffene Module

- `docs/plan/steering-loop.md` — zwei neue Einträge.
- [`.claude/commands/slice.md`](../../../../.claude/commands/slice.md) — die beiden Guide-Hälften.

**Eine Schicht** (Planungs-Doku und ihr Skelett). Kein Code, kein Vertrag berührt.

## 3. Auszuführende Gates

`make gates` und `make verify`, Ausgabe je in eine Datei, Exit-Code getrennt geprüft. Kein neuer
Sensor in diesem Slice, also keine Negativ-Probe — die Belege sind die Vorfallszahlen, und die
sind aus der Commit-Historie gemessen, nicht geschätzt.

**Machbarkeits-Beleg für die Sensor-Antwort auf `SL-003`** (gehört zum Eintrag, nicht zum Bau):
die Hypothese *„ein `docs(planning)`/`fix(planning)`-Commit ändert nur `docs/plan/planning/`"*
wurde gegen die Historie geprüft und **diskriminiert sauber** — sie fängt alle drei Vorfälle
(`615e37f`, `f57289d`, `d436da9`) und lässt alle fünf geprüften legitimen `docs(planning)`-Commits
(`5fe94e8`, `327e960`, `5cd09e9`, `38c2610`, `43c8b5c`) durch. Kein Rauschen.

## 4. Was bewusst nicht getan wird

- **Der Sensor zu `SL-003` wird hier nicht gebaut**, obwohl er machbar ist (§3). Grund ist kein
  Aufwand, sondern eine **Reihenfolge**: die Regel „welcher Commit-Typ darf welche Pfade berühren"
  existiert im Repo noch **nicht**. Ein Sensor würde sie also nicht durchsetzen, sondern
  *erfinden* — dieselbe Lage wie bei der Closure-Pflicht vor
  [slice-050](../done/slice-050-verify-schicht.md), die erst ihren Anker in `AGENTS.md` brauchte.
  Erst die Konvention, dann ihr Sensor; Folge-Slice benannt.
- **Für `SL-004` ist kein Sensor vorgesehen**, und das ist keine Auslassung: das Muster betrifft
  **Bauwissen** über Sensoren, nicht wiederkehrende Laufzeit-Fehler. Es gibt keinen Lauf, in dem
  es sich zeigen könnte — der Guide ist hier die vollständige Antwort, nicht die halbe.
- **Keine Korrektur der drei Commit-Nachrichten.** Historie wird nicht umgeschrieben; die
  Zuordnung ist in den Reports und jetzt im Kanal festgehalten.
- **Keine `slice-048`-Korrekturen und keine Status-Felder** — slice-062.

## 5. DoD

- [x] `SL-003` steht im Kanal mit **drei** belegten Vorfällen (Commit-SHAs und je die
      Substanz-Diskrepanz), Klassen-Einordnung, gelieferter Guide-Hälfte und benanntem
      Sensor-Kandidaten samt Diskriminierungs-Beleg aus §3.
- [x] `SL-004` steht im Kanal mit **drei** belegten Vorfällen (slice-050, slice-057, slice-060) und
      der Begründung, warum der Guide hier die vollständige Antwort ist.
- [x] Beide Guide-Hälften stehen im Workflow-Skelett; `make gates` und `make verify` grün —
      Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** zwei Steering-Loop-Einträge mit je drei aus der Historie gemessenen Vorfällen —
`SL-003` (Commit-Betreff bezeichnet nicht die enthaltene Arbeit, mit fünf belegten Commits und
einem Diskriminierungs-Nachweis für den künftigen Sensor) und `SL-004` (ein neuer Doku-Sensor
meldet im ersten Lauf sein eigenes Umfeld). Beide Guide-Hälften stehen im Workflow-Skelett,
Schritt 5 und Schritt 10.

**Lerneintrag — Form: benannte Spec-Lücke.**
> **Ein Sensor darf eine Regel durchsetzen, aber nicht erfinden — und für `SL-003` fehlt die
> Regel.** Der Sensor-Entwurf ist fertig und nachweislich rauschfrei (§3), trotzdem wird er hier
> nicht gebaut: das Repo deklariert **nirgends**, welcher Commit-Typ welche Pfade berühren darf.
> `AGENTS.md` §5 verlangt eine ID pro Commit, die Conventional-Commit-Typen werden gelebt, aber
> ihre Bedeutung ist nirgends normativ. Ein Gate darauf wäre eine stille Setzung — dieselbe
> Klasse, die dieses Repo als Harness-Lüge führt, und ausgerechnet in einem Slice, der
> Buchführungs-Ehrlichkeit adressiert. *Weil* die Konvention fehlt, ist die Lücke die eigentliche
> Lieferung dieses Slice: benannt, belegt, mit fertigem Entwurf für den Tag, an dem sie geschlossen
> ist. Prüfsatz: *bevor ein Sensor gebaut wird, die Stelle zitieren können, die er durchsetzt —
> gibt es sie nicht, ist der nächste Schritt die Regel, nicht das Skript.*

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. Beide Einträge tragen eine **Vorfallszahl** und je Vorfall einen prüfbaren Anker (Commit-SHA
   bzw. Slice-ID); die Pflege-Regel des Kanals verbietet Einträge ohne Zahl, und ohne Anker wäre
   die Zahl nicht nachrechenbar. Der `SL-003`-Sensorentwurf ist gegen acht Commits geprüft: drei
   Treffer, fünf Durchlässe.

**Zur Unterscheidung der beiden Antworten, ausdrücklich:** `SL-004` bekommt **nur** einen Guide,
und das ist kein Rückfall in „besser aufpassen". Bei `SL-001` und `SL-002` war der Guide
nachweislich zu schwach, weil beide Muster in einem *Lauf* auftreten, den ein Sensor beobachten
kann. `SL-004` beschreibt Bauwissen — es gibt keinen Lauf, in dem es sich zeigt, und ein Sensor
über „hat der Autor beim Bau an Zitat-Kontexte gedacht" existiert nicht. Wo kein Beobachtungspunkt
ist, ist der Guide die vollständige Antwort und nicht die halbe.

**Folge-Slices:** slice-062 (Commit-Typ-Konvention deklarieren, dann den `SL-003`-Sensor bauen)
und die verbliebene Doku-Arbeit der Review-Serie (slice-048-Korrekturen, Status-Felder).

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
