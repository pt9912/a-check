# Review-Report: slice-178 — 2026-09-07

**Review-Art:** Plan-/Doku-Review (unabhängiger Lauf) — geprüft wird die
Kopieranleitung in `AGENTS.md` §5 als **Aussage über den eigenen Bestand**,
Punkt für Punkt gegen das Repo gehalten, dazu `AGENTS.md` §6 Schritt 4 gegen
die Baseline und der Slice selbst als Muster seiner eigenen Regel.
Kontext-Trennung eingehalten: dieser Lauf hat den Gegenstand nicht verfasst und
den Verfasser-Kontext nicht geerbt.

**Gegenstand:** slice-178, Commit-Range `HEAD~1..HEAD` (`f9cc84b`,
`feat(harness): slice-178 Etappe E`) — drei Dateien, `+146/−30`. Der
Lifecycle-`git mv` liegt davor (`ccb5aea`) und ist nicht Gegenstand.

**Skill:** `.harness/skills/reviewer.md` @ `3fae6d3` (slice-176)
**Modell:** `claude-opus-5[1m]` · **Datum:** 2026-09-07

> **Zitier-Form** *(dieser Block bleibt stehen — er ist Norm, kein
> Ausfüll-Hinweis)*. Dieser Report friert ein; was er zitiert, bewegt sich
> weiter. Deshalb **Kennung, nicht Adresse** — `slice-NNN` statt seines
> Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link
> (`v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt>). Der vendored Baum trägt
> genau einen Tag; der nächste Sprung löscht den alten, und ein Link darauf
> färbt ein Artefakt rot, das niemand mehr anfassen darf. Das `pfad`-Feld auf
> den **geprüften Gegenstand** ist davon nicht betroffen — es hält den Stand
> des Laufs fest und darf das.

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- slice-178 (Plan, `in-progress/`), slice-174 §3.2 T-3 (der Auftrag),
  slice-175/176/177/179 (die Geschwister-Etappen als Bestands-Vergleich),
  slice-165 §3/§4 (die Messung hinter Punkt 6)
- `v6.5.0` · `templates/docs/plan/planning/slice.template.md` (§1, §2, §7, §8),
  `regelwerk/modul-05-planning-harness.md` §Ziel-Form: Slice ·
  §Offene Risiken werden bei Closure aufgelöst · §Zwei Schritte vor der
  Modus-Begründung, `regelwerk/modul-06-roadmap.md` §Wann Arbeit eine Welle
  braucht · §Das Beobachtungs-Register, `regelwerk/modul-09-implementierung.md`
  §Minimal Agent Workflow, `regelwerk/grundlagen-traceability.md`
  §Herkunfts-Anker für Steering-Loop-Regeln
- `AGENTS.md` §3 (Hard Rules, insb. §3.3 und §3.7), §4 (Gates), §5, §6;
  `harness/conventions.md` §Modus-Deklaration pro Sub-Area; `MR-019`
- Beobachtungs-Register: alle 14 Einträge unter `BEO-HARNESS/` gesichtet
- Gemessener Bestand: 174 Slice-Dateien mit `**Welle:**`-Feld, 20
  nicht-archivierte Slice-Pläne (Abschnitts-Nummern und -Titel), 58 Vorkommen
  `seit slice-<NNN>` außerhalb der vendored Baseline
- ADRs: keine berührt

---

## Findings

### HIGH

**H-1 · `AGENTS.md` §5 kündigt fünf Punkte an und listet sechs**

- **quelle:** nachweislich falsche Tatsachenbehauptung, gegen dasselbe Artefakt
  verifiziert (Reviewer-Skill §Klassifikation)
- **pfad:** `AGENTS.md:292` („**Beim Kopieren anzupassen** — fünf Punkte,
  gemessen am Bestand (slice-178):") gegen `AGENTS.md:294–321` (Punkte 1–6)
- **befund:** Die Einleitung nennt fünf Punkte, die Liste trägt sechs
  nummerierte. Die Umstellung von Fließtext auf eine nummerierte Liste wurde
  in der Closure-Notiz damit begründet, dass ein Absatz „keine Zeile hat, die
  man gegen den Bestand halten kann" — die erste Zeile der neuen Liste hält der
  eigenen Zählung nicht stand. Commit-Message und §3 des Plans sagen
  konsistent „sechs".
- **verifizierbar:** nein — kein Sensor zählt Listenpunkte gegen ihre Ansage.

**H-2 · Punkt 2 erklärt den Herkunfts-Anker zu einem Feld, „das a-check nicht
führt" — der Slice selbst führt es vier Absätze weiter unten**

- **quelle:** `v6.5.0` · `regelwerk/grundlagen-traceability.md`
  §Herkunfts-Anker für Steering-Loop-Regeln; `v6.5.0` ·
  `templates/docs/plan/planning/slice.template.md` §7 (Feld `liegt in`)
- **pfad:** `AGENTS.md:296–301` („**Zwei** Felder streichen, die a-check nicht
  führt: Reconciliation-Register … und *Herkunfts-Anker*") gegen
  `docs/plan/planning/done/slice-178-slice-form-ziel-und-abgrenzung.md:153`
- **befund:** Der Herkunfts-Anker ist im Slice-Template das Feld
  `— liegt in <Zielort>` in §7 plus der `seit slice-<NNN>`-Marker am Zielort.
  slice-178 schreibt in seiner eigenen Closure-Notiz
  „— liegt in `AGENTS.md §5`"; slice-176 und slice-177 führen dasselbe Feld,
  slice-174, slice-175 und slice-179 führen ausdrücklich seine Negation
  („Kein `liegt in`-Feld: mit diesem Slice wurde nichts verkörpert") — beides ist
  Gebrauch, nicht Streichung. Repo-weit stehen 58 `seit slice-<NNN>`-Anker
  außerhalb der vendored Baseline, davon 20 in `state.md`-Dateien des
  Beobachtungs-Registers und 2 in `AGENTS.md` selbst. Wer Punkt 2 befolgt,
  streicht ein Feld, das die letzten sechs Slices durchgehend bedienen. Der
  Punkt ist zugleich der einzige der sechs, der die alte Vier-Felder-Aussage
  ungeprüft weiterträgt: §2.2 des Plans misst `Welle:` und die *drei Paarungen*
  nach, den Herkunfts-Anker nicht.
- **verifizierbar:** nein — kein Sensor hält die Kopieranleitung gegen den
  Slice-Bestand; genau das ist der Steering-Loop-Eintrag dieses Slice.

**H-3 · Punkt 2 macht den Träger der drei Paarungen vom Kopf-Feld eines Slice
abhängig; die Baseline sagt ausdrücklich das Gegenteil**

- **quelle:** `v6.5.0` · `regelwerk/modul-06-roadmap.md` §Wann Arbeit eine
  Welle braucht („**Achse zuerst:** *Wellenlos* ist eine Eigenschaft des
  **Repos** … Das Kopf-Feld `**Welle:**` eines Slice-Plans sagt nur, ob dieser
  Slice in ein Bündel gehört; **daraus folgt für die Vorgänge unten nichts**.
  Ein Repo mit Wellen hat eine Welle-Closure, und die liest und prüft alles,
  was seit der letzten Welle in `done/` liegt — **auch Slices ohne
  Wellen-Zugehörigkeit**."); `v6.5.0` ·
  `templates/docs/plan/planning/slice.template.md` §2, letztes DoD-Item („im
  Repo **ohne** Wellen-Betrieb hier geprüft, im Repo **mit** Wellen von der
  nächsten Welle-Closure")
- **pfad:** `AGENTS.md:297–299` („die *drei Paarungen* auch: bei wellenlosen
  Slices trägt die Closure sie, bei Slices mit `**Welle:**`-Feld die
  Welle-Closure"); angewandt in
  `docs/plan/planning/done/slice-178-slice-form-ziel-und-abgrenzung.md:168`
- **befund:** Die Regel wird als *pro Slice* entscheidbar formuliert. Baseline
  und Ziel-Form binden sie an das Repo: a-check hat mit `welle-15` eine offene
  Welle, also trägt die Welle-Closure die Paarungen für **jeden** seit der
  letzten Closure geschlossenen Slice — auch für die 100+ mit
  `**Welle:** ohne Welle`. Nach dem Wortlaut von Punkt 2 müssten genau die ihre
  Paarungen selbst tragen. Das Ergebnis ist für slice-178 zufällig richtig (er
  trägt `welle-15`), das Kriterium ist es nicht. Zusätzlich widerspricht der
  Satz `AGENTS.md` §6, das im selben Dokument **zwei** Bedingungen führt („kein
  `**Welle:**`-Feld, keine aktive Welle in `in-progress/`"). Eine Abweichung
  von der Baseline-Regel ist an keiner Stelle als `MR-*` deklariert — nach dem
  Diskrepanz-Trichter (`AGENTS.md` §5) wäre sie deklarationspflichtig.
- **verifizierbar:** nein — `make verify` prüft die Existenz der
  Closure-Abschnitte, nicht die Zuordnung ihres Trägers.

### MEDIUM

**M-1 · „Acht Slices im Bestand widerlegen das" — die Grundgesamtheit ist
weder genannt noch reproduzierbar**

- **quelle:** unbelegte/fehlkalibrierte Tatsachenbehauptung;
  `BEO-GATE/cr-text-behauptet-statt-gemessen` („hast du *das* gemessen,
  worüber du redest?")
- **pfad:** `AGENTS.md:299–300`; Plan §2.2 Nr. 1; Closure-Notiz §8;
  Commit-Message
- **befund:** Repo-weit tragen **174** Dateien unter `docs/plan/planning/` ein
  `**Welle:**`-Feld, nicht acht. Auch der genannte Split stimmt in keinem
  plausiblen Ausschnitt: `welle-15` steht sechsmal (fünf in `done/` plus
  slice-178 selbst), `ohne Welle` in `open/` zweimal — behauptet sind
  „fünf … drei". Acht ergibt sich nur aus einem ungenannten Fenster (die
  jüngsten Slices plus `open/`). Die **Schlussfolgerung** — `Welle:` bleibt —
  ist richtig und sogar deutlich stärker belegt als behauptet; die **Messung**,
  auf die sich der Slice ausdrücklich beruft („gemessen statt übernommen"),
  trägt sie nicht.
- **verifizierbar:** ja — `grep -rl '^\*\*Welle:\*\*' docs/plan/planning/ | wc -l`.

**M-2 · Punkt 3 nennt eine Ursache, die die behauptete Verschiebung nicht
erzeugt**

- **quelle:** `v6.5.0` · `templates/docs/plan/planning/slice.template.md`
  (§1 Ziel · §2 DoD · §3 Plan · … · §8 Sub-Area)
- **pfad:** `AGENTS.md:302–306`
- **befund:** *Ein* eingeschobener Abschnitt verschiebt um **eine** Position —
  das trifft `Sub-Area §8 → §9` und erklärt `DoD §2 → §4` nicht. Die zweite
  Position entsteht dadurch, dass a-check die DoD **hinter** den
  Umsetzungs-/Plan-Abschnitt stellt statt davor; das steht nirgends. Wer Punkt 3
  wörtlich befolgt, setzt die DoD auf §3 und bekommt eine Nummerierung, die die
  Anleitung im selben Satz für falsch erklärt — derselbe Fehlertyp, den §2.2
  Nr. 2 an der Vorgänger-Fassung beanstandet. Zweitens decken die beiden
  genannten Namen („*Ausgangslage*" / „*Analyse (vor der Umsetzung)*") den
  Bestand nicht: slice-178 hat seinen eigenen Abschnitt **in diesem Commit** von
  „Analyse (vor der Umsetzung)" auf „Analyse" umbenannt, slice-174 führt
  „Ausgangsmessung (vor der Analyse)". Drittens gilt die Ziel-Nummerierung nur
  für die jüngste Kohorte: über die 20 nicht-archivierten Pläne steht der
  Sub-Area-Abschnitt auf §7 (2×), §8 (4×), §9 (9×) und §10 (2×).
- **verifizierbar:** nein — kein Sensor prüft Abschnitts-Nummern gegen die
  Ziel-Form.

**M-3 · Punkt 5 schreibt einen §9-Titel vor, den kein Slice im Repo trägt —
auch slice-178 nicht**

- **quelle:** `v6.5.0` · `templates/docs/plan/planning/slice.template.md` §8
- **pfad:** `AGENTS.md:314–317` gegen
  `docs/plan/planning/done/slice-178-slice-form-ziel-und-abgrenzung.md:171`
  (`## 9. Sub-Area-Modus`)
- **befund:** Über alle 20 nicht-archivierten Pläne lautet der Titel ausnahmslos
  „Sub-Area-Modus". §2.2 Nr. 3 des Plans formuliert im Präteritum („§9 **trug**
  den alten Titel *Sub-Area-Modus* … die Verkürzung, gegen die die neue Fassung
  ausdrücklich argumentiert") und liest sich damit wie ein behobener Befund; der
  Diff benennt nur `AGENTS.md`. Der Ausschluss in §1 deckt das nicht: er nimmt
  „die bestehenden Slices in `open/`" aus, nicht den Slice, der die Regel
  schreibt — dessen §2 in demselben Commit umbenannt wurde, sein §9 aber nicht.
  Die Anleitung beschreibt damit eine Form, deren erster Beleg noch aussteht.
- **verifizierbar:** nein — `doc-structure` prüft die Kopffelder und die
  Closure-Struktur, nicht die Abschnitts-Titel.

**M-4 · Das Risiko in §7 trägt zwei Ausgänge zugleich und landet in keinem**

- **quelle:** `v6.5.0` · `regelwerk/modul-05-planning-harness.md`
  §Offene Risiken werden bei Closure aufgelöst (*eingetreten* → Carveout oder
  Folge-Slice mit ID · *entfallen* → gestrichen mit Begründung · *weiter offen*
  → Beobachtungs-Register); `v6.5.0` · `regelwerk/modul-06-roadmap.md`
  §Das Beobachtungs-Register (Zähler = Zahl der Evidence-Dateien)
- **pfad:**
  `docs/plan/planning/done/slice-178-slice-form-ziel-und-abgrenzung.md:123–131`
  und `:156–161`
- **befund:** Die Zeile deklariert „**Ausgang:** weiter offen → Register" und
  sagt vier Zeilen später „Das Risiko ist **eingetreten**". Die drei Ausgänge
  sind eine geschlossene Menge; *eingetreten* verlangt Carveout oder
  Folge-Slice mit ID, §8 nennt „Folge-Slices: keine". Der gewählte Ausgang
  *weiter offen* wiederum bedeutet, dass das Risiko **ins Register wandert** —
  §8 legt aber weder ein Verzeichnis an noch eine Evidence-Datei („keine neue
  Beobachtung … bleibt bei 2×"). Damit hinterlässt der Ausgang im Register keine
  Spur, und der Zähler, der die 3×-Schwelle trägt, bewegt sich nicht.
  `make verify-risiko-ausgaenge` bleibt grün, weil es die **Form** des Ausgangs
  prüft, nicht seinen Vollzug.
- **verifizierbar:** teilweise — die Form ja (`make verify-risiko-ausgaenge`,
  grün), der Widerspruch und der ausbleibende Registereintrag nein.

**M-5 · Die §9-Sichtung übergeht die beiden Registereinträge, die genau diesen
Abschnitt benennen**

- **quelle:** `v6.5.0` · `regelwerk/modul-05-planning-harness.md`
  §Zwei Schritte vor der Modus-Begründung (Sichtungs-Schritt, Sub-Area-bezogen)
- **pfad:**
  `docs/plan/planning/done/slice-178-slice-form-ziel-und-abgrenzung.md:180–183`
- **befund:** Gesichtet und als „einschlägig" benannt ist genau ein Eintrag
  (`BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`, 2×).
  Unter derselben Sub-Area *Harness-Einstieg* stehen offen: (a)
  `BEO-HARNESS/agents-md-hinkt-baseline-dod-item-hinterher` — Titel und Text
  benennen wörtlich „`AGENTS.md` §5 führt die eigene Liste dieser Posten in
  eigenen Worten" und die daraus folgende Drift, also den Gegenstand dieses
  Slice; (b) `BEO-HARNESS/chronik-in-gelesenen-dateien`, dessen `state.md`
  „`AGENTS.md` §5 … trägt je eine gemessene Reststelle" sagt. Beide betreffen
  denselben Abschnitt, den der Slice umschreibt. Ohne Nennung ist „gesehen und
  verworfen" von „übersehen" nicht unterscheidbar — und (a) stünde mit einem
  Beleg aus diesem Slice bei 2×.
- **verifizierbar:** nein — `make verify-observations` prüft die Deckung
  zitierter Pfade, nicht die Vollständigkeit der Sichtung.

**M-6 · Der Commit trägt Chronik in eine Datei, die jeder Lauf liest**

- **quelle:** `AGENTS.md` §3.7 („**Falsch:** abwesenden Text beschreiben …
  **Richtig:** die geltende Zusage nennen; die vorige hält `git`"; Durchsetzung
  ausdrücklich „keine — die Regel ist inferentiell … Sie hängt am Review");
  `BEO-HARNESS/chronik-in-gelesenen-dateien` (offen)
- **pfad:** `AGENTS.md:299–300` („Bis slice-178 stand hier „vier Felder
  streichen" einschließlich `Welle:`; acht Slices im Bestand widerlegen das.")
- **befund:** Der Satz beschreibt abwesenden Text — die Form, die §3.7 als
  Negativbeispiel führt. Er ist nicht der erste seiner Art in `AGENTS.md` (§5
  trägt seit slice-077 dieselbe Bauart bei der WIP-Grenze), aber er ist ein
  **neuer** und trifft genau die Stelle, die der offene Registereintrag als
  „gemessene Reststelle" benennt. Der Slice hat den Abschnitt umgeschrieben und
  die Chronik dabei vermehrt statt reduziert; die dazugehörige Beobachtung
  bleibt ungezählt (vgl. M-5).
- **verifizierbar:** nein — §3.7 ist ausdrücklich sensorlos.

**M-7 · Die DoD-Zeile „Unabhängiger Review durchgeführt" war beim Commit
unwahr**

- **quelle:** Harness-Lüge-Klasse *behauptete Vollständigkeit*
  (`BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen`); Abgrenzung:
  die DoD-Erfüllung selbst gehört dem Verifier, geprüft ist hier nur der
  Wahrheitswert der Aussage zum Commit-Zeitpunkt
- **pfad:**
  `docs/plan/planning/done/slice-178-slice-form-ziel-und-abgrenzung.md:101`
- **befund:** Das Häkchen wurde im selben Commit gesetzt wie die Umsetzung;
  unter `docs/reviews/` existierte zu diesem Zeitpunkt kein Report für
  slice-178 (dieser hier ist der erste). `make doc-reviews` kann das nicht
  sehen: das Modul wertet `done/` aus, der Slice liegt in `in-progress/` — der
  Sensor lief im Gate-Lauf mit und meldete 0 Befunde, ohne Gegenstand zu haben.
  Dieselbe Konstruktion trägt auch die Häkchen für `make gates`/`make verify`,
  die ich reproduzieren konnte (beide Exit 0); für den Review gab es nichts zu
  reproduzieren.
- **verifizierbar:** teilweise — `make doc-reviews` bestätigt es erst nach dem
  `git mv` nach `done/`.

**M-8 · Der dritte Ausschluss in §1 ist eine Etikettierung, keine Begründung —
und das Etikett ist falsch**

- **quelle:** `v6.5.0` · `templates/docs/plan/planning/slice.template.md` §1
  („je Punkt eine Begründung, nicht nur eine Nennung: ein Ausschluss ohne Grund
  ist eine Behauptung"), vier Klassen als Suchraster
- **pfad:**
  `docs/plan/planning/done/slice-178-slice-form-ziel-und-abgrenzung.md:40–41`
- **befund:** „**Andere Etappen von welle-15** — Schicht-Abgrenzung." nennt eine
  Klasse statt eines Grundes; die ersten beiden Ausschlüsse zeigen im selben
  Abschnitt, wie es gemeint ist. Zudem trifft die Klasse nicht: Etappe B, C und
  D berühren dieselbe Sub-Area *Harness-Einstieg* wie dieser Slice — es ist ein
  **anderer Vorgang**, keine andere Schicht. §1 ist die erste Anwendung der
  Regel, die dieser Slice einführt.
- **verifizierbar:** nein.

### LOW

**L-1 · Punkt 6 liest sich unbedingt, `MR-019` hält die Review-DoD-Zeile
opt-in**

- **quelle:** `MR-019` („der Review-Report bleibt **Opt-in** über
  `make doc-reviews`/`DC-FA-RVW-001`, ausgelöst durch die exakte Phrase")
- **pfad:** `AGENTS.md:318–321`
- **befund:** Der Punkt steht in einer Liste „Beim Kopieren anzupassen" und
  formuliert ohne Vorbehalt („ihr Wortlaut wird beim Kopieren … umgeschrieben").
  Der Wortlaut ist gegenüber der Vorgänger-Fassung unverändert; die Umstellung
  auf nummerierte Schritte verstärkt aber seine Normativität. Dass die Zeile
  weggelassen werden **darf** — der Kern von `MR-019` — steht nur im verlinkten
  Eintrag, nicht im Punkt.
- **verifizierbar:** nein.

**L-2 · DoD-Punkt 1 und die Regel, die er abnimmt, nummerieren denselben
Abschnitt verschieden**

- **quelle:** interne Konsistenz
- **pfad:**
  `docs/plan/planning/done/slice-178-slice-form-ziel-und-abgrenzung.md:95–97`
  („dem neuen §8-Titel") gegen `AGENTS.md:314` („**§9 heißt** …")
- **befund:** Beide Zahlen sind unter Punkt 3 erklärbar (§8 in der Ziel-Form,
  §9 in a-checks Fassung), aber der DoD-Punkt prüft `AGENTS.md` und nennt dort
  die Ziel-Form-Nummer. Für einen Leser, der nur die DoD sieht, zeigt das
  Häkchen auf einen Abschnitt, den `AGENTS.md` anders nummeriert.
- **verifizierbar:** nein.

### INFO

**I-1 · `make trace-check` in dieser Umgebung nicht ausführbar**

- **pfad:** `Makefile:207`
- **befund:** `make trace-check RANGE=HEAD~1..HEAD` bricht mit
  `d-check: error: Range-Basis-Vorfahren nicht lesbar: object not found`
  (Exit 2) ab; derselbe Abbruch für einen unbeteiligten älteren Range
  (`HEAD~3..HEAD~2`). Der Befund ist also umgebungsbedingt (read-only
  gemountetes Repo im Container), keine Eigenschaft dieses Commits. Die
  Commit-Message nennt `slice-178` und `welle-15`, die Regel aus §5 ist
  inhaltlich erfüllt.
- **verifizierbar:** ja — in einem Klon mit vollem git-Zugriff.

**I-2 · `Verantwortlich: — (noch nicht priorisiert)` bei einem Slice in
`in-progress/`**

- **pfad:**
  `docs/plan/planning/done/slice-178-slice-form-ziel-und-abgrenzung.md:16`
- **befund:** Nach `v6.5.0` · `regelwerk/modul-05-planning-harness.md`
  §Lifecycle als State Machine hält das Feld den Rolleninhaber, der die Arbeit
  *hält*; `—` gilt „bis zur Priorisierung". Der Slice wird bearbeitet. Der
  Befund ist repo-weit (slice-174 bis slice-179 tragen alle dieselbe Zeile) und
  daher kein Fehler dieses Slice; die Kopieranleitung, die dieser Slice
  überarbeitet, erwähnt das Feld nicht. Zuständig: eine spätere
  Planner-Entscheidung, kein Nacharbeits-Punkt hier.
- **verifizierbar:** nein — die Ziel-Form sagt ausdrücklich „Kein Sensor prüft
  das Feld".

---

## Negativbefunde (geprüft, ohne Befund)

- **`AGENTS.md` §6 Schritt 4 gegen `v6.5.0` ·
  `regelwerk/modul-09-implementierung.md` §Minimal Agent Workflow:** der Absatz
  ist inhaltlich deckungsgleich — Plan-Ausgabe nennt Out-of-Scope · Dokument-
  Hälfte ist §1 *Ziel und Abgrenzung* · der Lauf schreibt fort statt neu zu
  erfinden · darf nicht stillschweigend weiten · Mitnahme = Plan-Änderung vor
  dem Code. Kein Zusatz, keine Auslassung, keine Verschärfung. Die Baseline
  stellt den Satz als eigenen Absatz unter die acht Schritte, `AGENTS.md` zieht
  ihn in Schritt 4 hinein — dieselbe Aussage, engerer Ort.
- **`harness/README.md` §Minimal agent workflow:** trägt den Verweis statt der
  Wiederholung und bleibt damit auf der eigenen Zusage („Diese Datei dupliziert
  sie nicht"). Kein zweiter Normtext.
- **Punkt 4 gegen `v6.5.0` ·
  `templates/docs/plan/planning/slice.template.md` §1:** Titel, Begründungs-
  pflicht je Punkt, „keine Mindestzahl", die vier Klassen als **Suchraster** und
  die Auflage, dass eine genannte Folge-Slice-Kennung den Punkt annehmen muss —
  alles korrekt und ohne Zusatz übernommen.
- **Punkt 5 gegen `templates/…/slice.template.md` §8:** Titel, „der Abschnitt
  entfällt nie", die zwei *Vorgelagert*-Blöcke in jedem Slice-Plan unabhängig
  von Modus und Slice-Typ — korrekt.
- **Punkt 1 (`Lerneintrag — Form:`):** die Ziel-Form kennt das Feld nicht,
  `make verify` verlangt es — die Aussage stimmt in beide Richtungen.
- **Punkt 2, erste Hälfte (Reconciliation-Register):** korrekt.
  `docs/plan/planning/reconciliation.md` existiert nicht, und die Ziel-Form
  nimmt das Item für Repos ohne Brownfield-Bootstrap selbst aus.
- **Größen-Regel:** zwei Liefer-Punkte (die beiden `AGENTS.md`-Zusagen; die
  `harness/README.md`-Zeile ist deren Verweis-Hälfte), eine Sub-Area
  *Harness-Einstieg*. Gate-Läufe, Review-Report, Closure-Notiz, Register und
  Risiko-Ausgänge zählen nicht mit — konform.
- **Hard Rule §3.3:** Lifecycle-`git mv` (`ccb5aea`) und Inhaltsänderung
  (`f9cc84b`) sind getrennte Commits.
- **Hard Rules §3.1/§3.2/§3.4/§3.5/§3.6:** keine Toolchain, kein Code, keine
  Suppression, kein Spec-Stratum und keine ADR berührt, keine Schwelle gesenkt.
- **Commit-Scope:** `feat(harness)` — die `(planning)`-Regel greift nicht;
  `make commit-scope-check RANGE=HEAD~1..HEAD` Exit 0.
- **Kennungs-Linkpflicht und Anker:** `make doc-check` im Gate-Lauf über 495
  Dateien ohne Befund; die neuen `slice-NNN`-Nennungen in `AGENTS.md` lösen
  keine Linkpflicht aus (gleiche Form wie slice-077/098/165 im Bestand).
- **Sub-Area-Wahl (§9, erster Vorgelagert-Block):** *Harness-Einstieg* ist in
  `harness/conventions.md` mit Kürzel `HARNESS` und Achsen 1,2,3 deklariert;
  die Abgrenzung zum *Planungs-Harness* („Gegenstand der Anleitung, aber nicht
  geändert") trifft zu — kein Slice-Plan außer dem eigenen wurde angefasst.
- **Gate-Läufe reproduziert:** `make gates` Exit 0, `make verify` Exit 0 (0
  Waisen, 21 Anforderungen), `make commit-scope-check` Exit 0. Die
  entsprechenden DoD-Häkchen tragen.
- **Roadmap-Kopplung:** `welle-15` führt slice-178 als Etappe E, die Roadmap
  benennt ihn unter *Offene Wellen* statt des Ruhe-Markers —
  `make doc-planning` grün, konsistent.

---

## Kategorie-Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 3 |
| MEDIUM | 8 |
| LOW | 2 |
| INFO | 2 |

## Verdikt

**Nicht abnahmefähig ohne Nacharbeit.**

Die beiden Zusagen aus dem Auftrag — §1 *Ziel und Abgrenzung* mit
Begründungspflicht und die Schritt-Hälfte in §6 — sind sauber und
baseline-treu umgesetzt; §6 Schritt 4 ist wörtlich deckungsgleich mit
`modul-09`. Der Mehrwert des Slice, die Kopieranleitung als *Aussage über den
eigenen Bestand* zu lesen, ist richtig erkannt und im Steering-Loop-Eintrag
richtig verallgemeinert.

Blockierend ist, dass genau diese Prüfung unvollständig blieb. Der
Steering-Loop-Eintrag verspricht „die sechs Punkte, jeder gegen den Bestand
gemessen"; gemessen sind zwei (H-2, M-1). Punkt 2 trägt die alte, ungeprüfte
Aussage über den Herkunfts-Anker weiter, obwohl der Slice das Feld in seiner
eigenen Closure-Notiz benutzt (H-2), und ergänzt sie um ein Träger-Kriterium,
das die Baseline ausdrücklich verwirft (H-3). Die Zahl, die den einen
tatsächlich gemessenen Punkt belegt, stimmt nicht (M-1). Und die erste Zeile
der Liste, deren Zweck es ist, zeilenweise gegen den Bestand haltbar zu sein,
zählt ihre eigenen Zeilen falsch (H-1).

Der Slice als Muster seiner Regel ist halb eingelöst: §1 trägt Titel und
Begründungen (ein Punkt davon nur als Etikett, M-8), §9 trägt den alten Titel
weiter, den derselbe Slice als „Verkürzung" beanstandet (M-3).

Für die Verifikation getrennt zu halten: `make gates` und `make verify` sind
reproduzierbar grün; keine der Findings ist ein Gate-Bruch. Alle drei HIGH
liegen in Sätzen, für die es im Repo keinen Sensor gibt — das ist der Befund,
nicht die Entschuldigung.
