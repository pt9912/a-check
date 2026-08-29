# slice-095 — Etappe B: Adaptions-Durchgang gegen `v5.12.0`

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** keine `AC-*`/`ADR-*` — Konventions-Urteil ohne Vertragsberührung.
**Bezug:** Etappe **B** aus [slice-092 §6](../done/slice-092-regelwerk-v5120-delta-analyse.md),
Vorgänger [slice-094](../done/slice-094-baseline-v5120-vendoring.md) (Etappe A).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

`modul-02` §Freshness-Audit des adoptierten Stands verlangt den Durchgang **durch die
Adaptions-Liste, nicht nur durch den Diff**, mit genau **fünf** Ausgängen je Eintrag:
*gegenstandslos* · *bleibt gültig* · *teilweise überholt* · *Bezug entfallen* · *widerspricht*.
Auch `permanent`-Einträge werden mitgeprüft — „permanent" heißt „kein automatischer
Auflösungs-Trigger", nicht „unauflösbar".

Dieser Slice fällt die zehn Urteile. Er **führt sie nicht aus** — warum, steht in §4.

## 2. Betroffene Module

`harness/conventions.md` §Adaptions-Block, urteilend; ausgeführt wird in Etappe C.
Eine Schicht, ein Dokument.

## 3. Die zehn Urteile

Jede Zeile nennt die Baseline-Regel, gegen die geurteilt wurde — Datei und Abschnitt im
vendored Stand `v5.12.0`, netzlos nachschlagbar.

| Eintrag | Ausgang | Baseline-Regel, an der es hängt |
|---|---|---|
| [MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration) | **bleibt gültig** | `grundlagen-harness-dateien.md` §Konventionsspeicher führt [`MR-000`](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration) weiterhin ausdrücklich als Adoptions-Erklärung — Sonderrolle, kein Adaptions-Eintrag |
| [MR-001](../../../../harness/conventions.md#mr-001--source-precedence-mit-eigener-spezifikations-schicht) | **gegenstandslos** | `grundlagen-source-precedence.md` führt `spec/spezifikation.md` **selbst** als Rang 2 von neun. Die Abweichung „drei statt zwei Spec-Ränge" existiert nicht mehr |
| [MR-002](../../../../harness/conventions.md#mr-002--id-schema-mit-bereichskürzeln-ab-initialer-fassung) | **gegenstandslos** | `grundlagen-source-precedence.md` §ID-Schema als Klammer: *„Die Zählteile stehen hier ohne Bereichssegment; `LH-FA-03` und `LH-FA-IDX-003` sind **beide wohlgeformt**."* Das Vertrags-Präfix ist zudem frei wählbar |
| [MR-003](../../../../harness/conventions.md#mr-003--source-precedence-ohne-docsuser-rang) | **bleibt gültig** | inhaltlich seit 2026-06-21 aufgelöst; die neue Fassung ändert nur die **Form** der Auflösung (Nachfolge-Eintrag statt Inline-Vermerk, `conventions/done/`) — Etappe C |
| [MR-004](../../../../harness/conventions.md#mr-004--spezifikation-und-architektur-strata-und-id-schemata) | **teilweise überholt** | Straten-Zuordnung und `SPEC-*`/`ARC-*` als **Struktur**-IDs sind jetzt Default (§ID-Schema als Klammer). Repo-spezifisch bleibt, dass a-check die neue Verfeinerungs-Form `<PREFIX>-FA-<NN>.<Buchstabe>` **nicht** führt |
| [MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung) | **teilweise überholt** | die neue Datei `grundlagen-referenz-richtung.md` regelt die Abwärts-Disziplin inkl. `ADR-`/`slice-` im Spec-Körper selbst. Repo-spezifisch bleibt allein die **Kodierung** in `.d-check.yml` (Token-Erkennung, Provenance-Marker, Grandfathering 0001–0020) |
| [MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert) | **gegenstandslos** | `modul-02` Schritt 2 macht das Vendoring zum Bootstrap-**Default**. Der Eintrag sagte schon selbst, er weiche nicht ab, sondern *stelle den Default her* — nach der neuen Regel *„ein Eintrag, der keine benannte Regel ersetzt, ist ein Fork, keine Adaption"* war er nie eine |
| [MR-007](../../../../harness/conventions.md#mr-007--adr-vorlagen-version-v352-statt-v130) | **teilweise überholt** | sein eigener Trigger lautet wörtlich „bis zur nächsten Baseline-Migration, die ihn ihrerseits ablöst" — die ist jetzt. Der Nachfolger nennt `v5.12.0` |
| [MR-008](../../../../harness/conventions.md#mr-008--kein-replay-keine-agenten-telemetrie) | **teilweise überholt** | `modul-12` grenzt sich **neu** selbst ein: die Regeln sprechen vom *„nicht-deterministischen Kern"*. Der Satz steht in `v3.5.2` nicht. Damit ist ein Teil der Abweichung Default geworden; der Rest (keine Agenten-Telemetrie, Ersatz-Closure-Kriterium `make ci`) bleibt |
| [MR-009](../../../../harness/conventions.md#mr-009--validator-rolle-unbesetzt-zwei-übergaben-ohne-artefakt) | **teilweise überholt** | `modul-08` sagt **neu** selbst: *„Die beiden Validator-Kanten laufen nicht in jeder Sequenz"*, ihr Beleg sei repo-extern, Artefaktklasse *keins*. Auch dieser Satz fehlt in `v3.5.2`. Was bleibt: dass die Rolle hier gar nicht besetzt ist |

**Verteilung:** 3× gegenstandslos · 2× bleibt gültig · 5× teilweise überholt · **0× Bezug
entfallen** · **0× widerspricht**.

**Das Ergebnis mit der größten Tragweite ist die Null.** *Widerspricht* ist laut `modul-02` der
einzige Ausgang, an dem „das Delta die Antwort **nicht** vorgibt, denn hier hat das Repo eine
Wahl". Dieser Durchgang trifft ihn **kein einziges Mal**: jedes der zehn Urteile folgt aus dem
Delta, keines aus einer Präferenz. Die Migration verlangt damit keine Konventions-Entscheidung —
sie verlangt Ausführung.

**Die Hälfte des Bestands ist eingeholt.** Drei Einträge verlieren ihren Gegenstand ganz, fünf
teilweise. Das ist kein Qualitätsurteil über die Adaptionen: [`MR-001`](../../../../harness/conventions.md#mr-001--source-precedence-mit-eigener-spezifikations-schicht), [`MR-002`](../../../../harness/conventions.md#mr-002--id-schema-mit-bereichskürzeln-ab-initialer-fassung) und [`MR-006`](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
haben beschrieben, was die Baseline zwei Major-Versionen später selbst übernommen hat.

## 4. Der Fund, der den Etappen-Schnitt korrigiert

[slice-092 §6](../done/slice-092-regelwerk-v5120-delta-analyse.md) sah für Etappe B vor, das neue
Pflichtfeld `Ersetzt-Baseline-Regel` **in die zehn Einträge nachzutragen**. Das geht nicht, und
zwar aus einer Regel derselben Baseline: *„Einträge werden nie überschrieben"* — Rückbau ist ein
**neuer Eintrag, kein Edit**, dieselbe Append-only-Disziplin wie bei `Accepted`-ADRs.

Ein Pflichtfeld nachträglich in einen akzeptierten Eintrag zu schreiben, wäre genau das verbotene
Edit. Das Feld gehört in die **Nachfolge-Einträge**, die Etappe C schreibt. Was dieser Slice
liefert, ist die Recherche dafür: Spalte 3 der Tabelle oben **ist** der Inhalt des Feldes, je
Eintrag genau eine Regel.

Der Etappen-Schnitt aus slice-092 war an dieser Stelle also falsch — nicht zu groß, sondern mit
einer Anweisung versehen, die die Baseline verbietet. Korrigiert wird er hier, nicht dort:
slice-092 liegt in `done/` und wird nicht überschrieben.

## 5. Auszuführende Gates

`make gates` — der Slice ändert nur ein Planning-Dokument, aber der Gate-Nachweis hängt am
Inhalts-Hash des Arbeitsbaums, nicht am Umfang der Änderung. Zum Abschluss `make verify`.

**Kein neuer Sensor**, also keine Negativ-Probe: geliefert werden Urteile, kein Zustand, den ein
Gate hält.

## 6. Was bewusst nicht getan wird

- **Kein Rückbau, keine Nachfolge-Einträge, keine Formänderung.** Alles das ist Etappe C. Ein
  Rückbau in der heutigen Inline-Form und unmittelbar danach der Umzug in die Verzeichnis-Form
  wäre dieselbe Arbeit zweimal.
- **Die Adaptionen bleiben bis dahin in Kraft** — so steht es seit Etappe A ausdrücklich in
  `harness/conventions.md` §Baseline. Ein Urteil ist noch keine Aufhebung.
- **Kein Urteil über die Kurs-Klausel.** Ob das Repo dem vendored Regelwerk statt dem Kurs folgt,
  ist eine **neue** Adaption und nicht Teil dieses Bestands — sie hat ihren eigenen Slice
  ([slice-093](../open/slice-093-regelwerk-statt-kurs.md)).

## 7. DoD

- [x] Alle zehn Einträge tragen ein Urteil aus der geschlossenen Fünfer-Menge, je mit **genau
      einer** benannten Baseline-Regel (Datei + Abschnitt im vendored Stand) — Beleg: §3.
- [x] Die beiden Urteile, die an **neuen** Sätzen der Baseline hängen, sind als solche belegt
      (der Satz fehlt in `v3.5.2`), und die Nicht-Nachtragbarkeit des Pflichtfelds ist aus der
      Append-only-Regel begründet — Beleg: §3 und §4.
- [x] `make gates` (und bei Abschluss `make verify`) grün — **Ausgabe in eine Datei**, Exit-Code
      getrennt geprüft, nie in eine Pipe.

## 8. Closure-Notiz

**Geliefert:** zehn Urteile aus der geschlossenen Fünfer-Menge, je mit genau einer benannten
Baseline-Regel — und damit die Recherche, aus der Etappe C ihre Nachfolge-Einträge schreibt.

**Lerneintrag — Form: geschärfte Regel.** *Ein Eintrag im Adaptions-Block muss eine **benannte**
Baseline-Regel ersetzen; tut er das nicht, gehört er nicht hinein.* Der neue Stand formuliert den
Prüfsatz selbst — *„Ein Eintrag, der keine benannte Regel ersetzt, ist ein Fork, keine Adaption"* —
und dieser Durchgang zeigt, dass er greift: [`MR-006`](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert) stand seit dem 2026-07-25 als Adaption,
*weil* er beschrieb, was das Repo tut, statt zu benennen, wovon es abweicht. Sein eigener Text
sagte es die ganze Zeit („keine inhaltliche Abweichung … im Gegenteil: sie **stellt ihn her**") —
nur fehlte die Regel, an der das auffällt. Der Prüfsatz gilt ab sofort auch nach vorn: er ist die
erste Frage an die Adaption, die
[slice-093](../open/slice-093-regelwerk-statt-kurs.md) vorschlägt.

**Zwei beobachtbare Closure-Kriterien:**

1. Jede der zehn Zeilen nennt Datei **und** Abschnitt im vendored Stand; alle sind netzlos unter
   `.harness/baseline/v5.12.0/regelwerk/` nachschlagbar. Die zwei Urteile, die an neuen Sätzen
   hängen, sind mit `grep -c` gegen `v3.5.2` belegt — je **0** Treffer.
2. `make gates` und `make verify` je **Exit 0**; `doc-check` 193 Dateien, 0 Befunde.

**Offene Risiken und ihr Ausgang:**

- *Die Urteile sind gefällt, aber nicht ausgeführt* — Ausgang: **Folge-Slice**, Etappe C.
- *Das Pflichtfeld fehlt weiterhin in allen zehn Einträgen* — Ausgang: **Folge-Slice**, Etappe C;
  es kann per Append-only-Regel nur in Nachfolge-Einträgen entstehen (§4).
- *`slice-092` §6 trägt weiterhin die unausführbare Anweisung* — Ausgang: **gestrichen mit
  Begründung**. Der Slice liegt in `done/` und wird nicht überschrieben; die Korrektur steht in §4
  dieses Slice und ist von dort auf slice-092 verlinkt. Eine Planungsaussage nachträglich
  zurechtzurücken wäre dieselbe Klasse Edit, die dieser Slice gerade als verboten belegt hat.

**Folge-Slices:** Etappe C, wird unmittelbar geschnitten; D danach.

## 9. Sub-Area-Modus

Berührt wird der Konventionsspeicher. Er liegt unter **keiner** Zeile der Modus-Deklaration pro
Sub-Area — dieselbe Lücke, die [slice-091 §7](../done/slice-091-claude-md-auf-verweis-reduzieren.md)
benannt hat. Alle berührten Sub-Areas mit Modus sind GF.
