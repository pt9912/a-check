# slice-179 — Etappe B: Adaptions-Durchgang gegen `v6.5.0`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-15](../welle-15-regelwerk-v650-migration.md)

**Bezug:** [slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md) §3.4 —
Etappe **B** des Schnitts. Vorbild derselben Form: der Durchgang gegen `v6.1.0`
(slice-163, archiviert).

**Berührte Spec-Stellen:** — *(keine)* — Konventions-Bestand ohne
Vertragsberührung.

**Verantwortlich:** — *(noch nicht priorisiert)*

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-07.

---

## 1. Ziel und Abgrenzung

Jede der **sieben** aktiven Adaptionen ist gegen `v6.5.0` bewertet: Steht die
Regel, die sie ersetzt, dort noch? Ist ihr Auflösungs-Trigger eingetreten? Und
die eine Ausnahme, die a-check ohne Eintrag gesetzt hat — `exempt-paths` — ist
eingeordnet.

*(Dieser Plan nutzt bereits die `v6.5.0`-Form §1 **Ziel und Abgrenzung**; die
Kopieranleitung in [`AGENTS.md`](../../../../AGENTS.md) §5 zieht Etappe E nach.)*

**Nicht in diesem Slice**, je Punkt mit Grund:

- **Aufgelöste Einträge in [`conventions/done/`](../../../../harness/conventions/done/)** —
  Bestand bleibt bewusst stehen: sie sind eingefroren und gegen den Stand
  formuliert, der damals galt. Ein Durchgang durch sie prüfte nichts.
- **Neue Adaptionen für Regeln, die a-check bricht** — wäre ein anderer
  Vorgang. Etappe D ([slice-177](../done/slice-177-sensors-struktur-zwei-tabellen.md))
  trägt den bekannten Fall (`AGENTS.md` §4 gegen die Ziel-Form); weitere Funde
  bekommen einen eigenen Slice, keinen Schnellschuss hier.
- **Rückbau eines Eintrags ohne eingetretenen Trigger** — der Durchgang
  *bewertet*; er löst nur auf, wo die Bedingung des Eintrags selbst erfüllt ist.
  Sonst wäre er eine Meinungsänderung mit Migrations-Anlass.

## 2. Ausgangslage (gemessen, 2026-09-07)

**Wortgleichheit der fünf `Ersetzt-Baseline-Regel`-Ziele** ist in
[slice-175](../done/slice-175-etappe-a-vendoring-v650.md) §2.2 bereits gemessen:
vier unverändert, [`MR-015`](../../../../harness/conventions.md#mr-015)s Ziel `regelwerk/modul-06-roadmap.md` mit `+3/−1` —
die von ihm ersetzte Replay-Zusage aber unberührt. **Das ist die Datei-Ebene.**
Dieser Slice prüft die **Aussagen**-Ebene: ob die ersetzte Regel inhaltlich noch
dort steht und ob sie noch dasselbe verlangt.

**Ein Trigger ist vorab geprüft und *nicht* eingetreten:**
[`MR-019`](../../../../harness/conventions.md#mr-019) löst auf bei *„der
nächsten Baseline-Migration, wenn sie das Template erneut ändert und diese
Adaption dadurch gegenstandslos wird"*. `slice.template.md` hat sich in `v6.5.0`
geändert (`+36/−10`), die **Review-DoD-Zeile** darin jedoch nicht — gemessen mit
`git diff -w v6.2.0 v6.5.0` auf die Zeile. Der Eintrag bleibt gegenstandsvoll.

**Offen und ohne Eintrag:** `exempt-paths` in
[`.d-check.yml`](../../../../.d-check.yml). Der Kurs nennt ein solches Ventil
eine *„Gate-Senkung mit eigener Begründungslast"*;
[`AGENTS.md`](../../../../AGENTS.md) §3.6 verlangt dafür eine ADR, und keine der
39 nennt `exempt-paths` oder `version-stale`
([slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md) §3.4).

## 3. Durchgang

### 3.1 Die sieben aktiven Einträge gegen `v6.5.0`

Zwei Ebenen, getrennt geprüft. Die **Datei**-Ebene (ist der Zielabschnitt
wortgleich?) liegt aus
[slice-175](../done/slice-175-etappe-a-vendoring-v650.md) §2.2 vor. Diese Tabelle
prüft die **Aussagen**-Ebene: Steht die *ersetzte Regel* in `v6.5.0` noch, und
verlangt sie noch dasselbe?

| Eintrag | Ersetzte Regel | Steht sie noch? | Trigger eingetreten? |
|---|---|---|---|
| [`MR-011`](../../../../harness/conventions.md#mr-011) | Suffix-Form `<PREFIX>-FA-<NN>.<Buchstabe>` für Verfeinerungen | **ja** — Wortlaut unverändert in `grundlagen-source-precedence.md` §ID-Schema als Klammer | nein (Trigger: eigene Erfahrung, dass ein Feld teurer ist) |
| [`MR-012`](../../../../harness/conventions.md#mr-012) | Aufwärts-Richtung für **jede** Kante im bindenden Text | **ja** — die Decken-Regel steht unverändert | nein (löst erst auf, wenn keine der 20 ADRs mehr `Accepted` ist) |
| [`MR-014`](../../../../harness/conventions.md#mr-014) | Span-Telemetrie je Tool-Call, Token-Bilanz je Rolle | **ja** — `modul-15` verlangt sie unverändert | nein (permanent: a-check ruft kein Modell auf) |
| [`MR-015`](../../../../harness/conventions.md#mr-015) | Replay-Lauf als Teil des Welle-Closure-Triggers | **ja** — Zeile 33 und 319 in `modul-06-roadmap.md`, beide unverändert | nein (bräuchte ein Golden Set) |
| [`MR-016`](../../../../harness/conventions.md#mr-016) | neun Rollen-Übergaben, jede mit Artefakt | **ja** — §Die neun Übergaben unverändert | nein (bräuchte einen repo-externen Abnehmer) |
| [`MR-019`](../../../../harness/conventions.md#mr-019) | Review-Report als unbedingter DoD-Checkbox-Punkt | **ja** — die Zeile in `slice.template.md` ist unverändert, obwohl die Datei sich um `+36/−10` geändert hat | **nein**, und das war der einzige Kandidat (§2) |
| [`MR-020`](../../../../harness/conventions.md#mr-020) | — *(korrigiert eine Repo-Aussage, kein Baseline-Regel-Ersatz)* | n/a | nein (permanent) |

**Alle sieben bleiben.** Kein Eintrag ist gegenstandslos geworden, keiner
löst auf. Der Durchgang hat damit nichts zu ändern — und das ist ein Ergebnis,
kein Leerlauf: Es ist die erste Migration, bei der das vorab gemessen wurde
statt angenommen.

**Ein Befund über die Methode:** [`MR-015`](../../../../harness/conventions.md#mr-015)s Zielabschnitt hat sich geändert
(`+3/−1`, [slice-175](../done/slice-175-etappe-a-vendoring-v650.md) §2.2), seine
ersetzte Regel nicht. Wer nur die Datei vergleicht, meldet hier falsch — in
beide Richtungen: Eine unveränderte Datei kann die Regel woanders verloren
haben, eine geänderte sie behalten. Die zwei Ebenen sind nicht ineinander
überführbar.

### 3.2 `exempt-paths` — Geltungsbereich, keine Gate-Senkung

Die Frage kam aus
[slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md) §3.4: Der Kurs
nenne ein Ausnahme-Ventil eine *„Gate-Senkung mit eigener Begründungslast"*,
`AGENTS.md` §3.6 verlange dafür eine ADR, und keine der 39 nenne `exempt-paths`.

**Die Prämisse trägt nicht.** Am Ort nachgelesen (`v6.5.0` ·
`regelwerk/grundlagen-harness-dateien.md`) meint der Satz ein Ventil im
Prüfbereich der **Link**-Prüfung, wenn eine Adresse bereits im eingefrorenen
Artefakt steht — und derselbe Absatz nennt die Alternative: *„die Reparatur ist
teurer als die Vermeidung."* Genau die hat
[slice-176](../done/slice-176-zitier-form-einfrierende-artefakte.md) gewählt.

Für `versions` gilt etwas anderes, und der Kurs sagt es zwei Absätze vorher:
*„Die Grenze: Sie gilt für einfrierende Artefakte … Der Unterschied ist nicht die
Wichtigkeit des Ziels, sondern ob der Zeiger nachgezogen werden **darf**."* Ein
Zeiger im Zeitdokument darf nicht nachgezogen werden; eine Prüfung, die ihn
trotzdem einfordert, verlangt einen Regelbruch. `exempt-paths` bildet ab,
worüber die Regel spricht — **Geltungsbereich, keine Schwellensenkung**.
`AGENTS.md` §3.6 greift nicht, eine ADR entsteht nicht.

**Aufgeschrieben ist die Einordnung dort, wo sie beim nächsten Gate-Streit
gelesen wird:** `harness/sensors/doc-check.md` §Grenze, Punkt 3 — neben der
Ausnahme selbst, nicht in einem Slice, den niemand mehr öffnet.

## 4. Definition of Done

- [x] Alle **sieben** aktiven Adaptionen sind bewertet — je Eintrag: steht die
      ersetzte Regel in `v6.5.0` noch, und ist der Auflösungs-Trigger
      eingetreten? Mit Beleg je Zeile, nicht als Sammelurteil.
- [x] Jeder Eintrag, dessen Trigger eingetreten ist, ist aufgelöst (Nachfolge-
      Eintrag oder Streichung mit Begründung); jeder andere trägt den Befund.
- [x] Die `exempt-paths`-Frage ist entschieden: ADR, deklarierte Ausnahme oder
      begründete Nicht-Handlung — und die Entscheidung steht dort, wo sie beim
      nächsten Gate-Streit gelesen wird.
- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Etappe A liegt in `done/`
([slice-175](../done/slice-175-etappe-a-vendoring-v650.md)) — erfüllt seit
2026-09-07; Maintainer-Freigabe und WIP-Limit frei.

**Rückführungen:** findet der Durchgang mehr als zwei Einträge mit
eingetretenem Trigger, wird die Auflösung ein eigener Slice und dieser bleibt
die Bewertung — zurück nach `next/` zur Zerlegung. Erscheint ein weiteres
Release, zurück nach `open/`.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz geschrieben.
Der Slice trägt ein `**Welle:**`-Feld und archiviert **mit seiner Welle**.

## 7. Risiken und offene Punkte

- *Der Durchgang prüft die Einträge, die es gibt — eine Baseline-Regel ohne
  Eintrag hat keinen Aufhänger und fällt wieder durch* — **Ausgang:** weiter
  offen → Beobachtungs-Register,
  [`BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md)
  (2×). Das Risiko ist **nicht** eingetreten und auch nicht entfallen: Dieser
  Durchgang hat wieder nur die sieben Einträge geprüft, und eine Regel ohne
  Eintrag hätte er wieder nicht gesehen. Dass diesmal keine durchfiel, weiß
  niemand — es hat sie nur niemand gesucht. Der Eintrag bleibt der einzige Ort,
  an dem das steht.
- *Die `exempt-paths`-Entscheidung wird zur vierten Repo-Aussage-Korrektur und
  reißt damit eine Schwelle* — **Ausgang:** entfallen, gestrichen mit
  Begründung. Es entsteht **kein** `MR`-Eintrag: Die Prüfung ergab, dass
  `exempt-paths` ein Geltungsbereich ist und keine Gate-Senkung (§3.2), also
  gibt es nichts zu adaptieren und nichts zu korrigieren. Klasse
  [`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
  (2×) bleibt bei zwei.

## 8. Closure-Notiz

**Lerneintrag — Form: geschärfte Regel** (Datei-Ebene und Aussagen-Ebene eines
Adaptions-Zeigers sind nicht ineinander überführbar).

- **Was hat funktioniert:** Der Durchgang hat **nichts geändert**, und das ist
  das Ergebnis. Sieben Einträge, sieben Befunde, kein Trigger eingetreten — zum
  ersten Mal vorab gemessen statt angenommen. Ein Adaptions-Durchgang, der
  nichts findet, ist kein Leerlauf; er ist der Beleg, dass der Bestand den
  Sprung überstanden hat.

- **Was ging anders als geplant:** Die `exempt-paths`-Frage, die
  [slice-174](../done/slice-174-regelwerk-v650-delta-analyse.md) §3.4 als
  offene ADR-Pflicht notiert hatte, **löste sich beim Nachlesen auf**. Das
  Zitat *„Gate-Senkung mit eigener Begründungslast"* stammt aus einem Absatz
  über die **Link**-Prüfung; zwei Absätze vorher zieht derselbe Text die Linie,
  die `exempt-paths` bei `versions` rechtfertigt — *„ob der Zeiger nachgezogen
  werden **darf**"*. Ein Zitat ohne seinen Absatz trug hier eine Pflicht, die
  es nicht gibt.

- **Steering-Loop-Eintrag — geschärfte Regel:** Der `Ersetzt-Baseline-Regel`-Zeiger
  einer Adaption wird auf **zwei** Ebenen geprüft, und sie sind nicht ineinander
  überführbar: ob der Zielabschnitt wortgleich ist (Datei), und ob die *ersetzte
  Regel* dort noch dasselbe verlangt (Aussage). Eine unveränderte Datei kann die
  Regel woanders verloren haben; eine geänderte sie behalten —
  [`MR-015`](../../../../harness/conventions.md#mr-015) ist der zweite Fall
  (`+3/−1` am Abschnitt, Replay-Zusage unberührt). — liegt in
  `docs/plan/planning/done/slice-179-…md` §3.1 als Muster für den nächsten
  Durchgang.

- **Beobachtungs-Register (`../observations/`):** keine Beobachtung angefallen,
  kein Beleg. Die zwei in §7 genannten Einträge bleiben bei 2× — der erste, weil
  das Risiko weder eintrat noch entfiel, der zweite, weil die befürchtete
  Korrektur gar nicht entstand.

- **Folge-Slices:** keine.

- **Risiken aus §7:** zwei, jedes mit genau einem Ausgang — einmal *weiter
  offen* → Register, einmal *gestrichen mit Begründung*.

- **Drei Paarungen:** trägt die Welle-Closure — der Slice hat ein
  `**Welle:**`-Feld.

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Harness-Einstieg** (`harness/conventions/`, `harness/sensors/`), Achsen 1,2,3,
deklariert in
[`conventions.md`](../../../../harness/conventions.md#modus-deklaration-pro-sub-area).
Die **Vendored Baseline** ist Gegenstand der Messung, aber nicht geändert; sie
führt ohnehin keinen Modus.

**Vorgelagert — offene Beobachtungen sichten:** Register am 2026-09-07
durchgegangen. Zwei einschlägig, beide in §7 mit Ausgang:
[`baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md)
(2×) und
[`adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
(2×). Keiner erreicht mit diesem Slice die Schwelle.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
