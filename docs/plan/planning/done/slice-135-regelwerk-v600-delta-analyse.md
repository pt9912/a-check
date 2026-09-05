# slice-135 — Regelwerk-Migration `v5.12.0` → `v6.0.0`: Delta-Analyse

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Anfrage 2026-09-05 („können wir schon auf das aktuelle Regelwerk (v6.0.0)
umsteigen — d-check ist schon migriert"). Präzedenz: [slice-092](../done/slice-092-regelwerk-v5120-delta-analyse.md),
dieselbe Analyse-Form für den Sprung `v3.5.2` → `v5.12.0`. Referenz-Migration im Schwester-Repo:
`d-check` `welle-88` (`slice-193`…`slice-195`), **nicht ungeprüft übernommen** — jeder Fund unten
ist gegen den a-check-Bestand selbst gemessen, nicht aus dem d-check-Commit zitiert
([`.harness/skills/cr-text-reviewer.md`](../../../../.harness/skills/cr-text-reviewer.md)-Disziplin
sinngemäß auf einen Fremdrepo-Vergleich angewandt).

**Berührte Spec-Stellen:** — *(keine)* — reine Ist-Messung, keine Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

> **Analyse zur Abnahme.** Wie bei slice-092: keine Kennungen vergeben, keine Artefakte geändert.
> Der Etappen-Schnitt (§6) gehört vor der Umsetzung abgenommen. Die Analyse **ist** die Lieferung.

**Geltungsbereich der Quelle:** gemessen und gelesen wurde ausschließlich `lab/regelwerk/` und
`lab/templates/` (das vendored Artefakt). Der Kurs unter `kurs/de/` ist nicht Gegenstand.

---

## 1. Ist-Stand

| | |
|---|---|
| a-check gepinnt auf | **`v5.12.0`** (Kurs-Welle 98, 2026-08-26) |
| aktuelles Release | **`v6.0.0`** (Kurs-Welle 116, 2026-09-03) — zwei Tage alt |
| dazwischen | **ein** Major-Sprung (v5 → v6), acht Kurs-Wellen |

Der Pin steht an drei deklarierenden Stellen — `harness/conventions.md` §Baseline (kanonisch),
`AGENTS.md` §1, `harness/README.md` §Guides — plus in **jedem** Pfad-Verweis auf
`.harness/baseline/v5.12.0/…` (108 Zeilen außerhalb des vendored Baums selbst, gemessen mit
`grep -rn`) und in [`MR-013`](../../../../harness/conventions/MR-013-adr-vorlagen-version.md)
(ADR-Vorlagen-Version). `.d-check.yml`s `scan.ignore` ist bereits versionsagnostisch
(`.harness/baseline/**`) — anders als bei slice-092 keine Fundstelle hier.

## 2. Umfang des Sprungs (gemessen)

`git diff --name-status v5.12.0 v6.0.0 -- lab/regelwerk lab/templates` (Roh-Git-Tags, unabhängig
vom vendored Extrakt geklont): **26 Dateien mit echtem Inhalt**, `+579 / −110`.

| Bereich | Dateien | Zeilen |
|---|---|---|
| `lab/regelwerk/` | 14 geändert | siehe unten |
| `lab/templates/` | 12 geändert (davon 3 neu, 1 gelöscht) | siehe unten |

Zum Vergleich: der letzte Sprung (`v3.5.2` → `v5.12.0`) war `+2932/−1012` über 29 Dateien allein in
`regelwerk/`. Dieser Sprung ist **deutlich kleiner** — knapp ein Fünftel des Volumens.

**Methodischer Befund, der die Vorsicht aus slice-092 §2 bestätigt und schärft.** Ein `diff -rq`
zwischen a-checks bereits vendiertem `v5.12.0`-Baum und dem frisch aus dem Release-ZIP
extrahierten `v6.0.0`-Baum meldet **alle 26 Dateien unter `regelwerk/`** als „verschieden" — zwölf
davon (`grundlagen-bootstrap`, `grundlagen-klassifikation`, `grundlagen-referenz-richtung`,
`modul-01/03/04/07/09/11/12/15/16`) tragen aber laut `git diff` auf den rohen Tags **keine
Zeilenänderung**. Der einzige Unterschied ist die pro-Tag umgeschriebene Link-Adresse
(`tools/rewrite-doc-links.py`, dieselbe Beobachtung wie in slice-092 §2). Wer nur `diff -rq` gegen
den vendored Baum fährt, überschätzt den inhaltlichen Umfang um **plus zwölf Dateien**. Für §4
unten zählt ausschließlich der `git diff`-Befund gegen die rohen Tags.

**Provenienz des Release-Assets — unabhängig geprüft, nicht angenommen.** `lab-regelwerk.zip`
(213.635 Bytes) aus `https://github.com/pt9912/ai-harness-course/releases/download/v6.0.0/`
frisch geladen; sein `sha256sum` stimmt exakt mit dem von der GitHub-Release-API gemeldeten
Asset-Digest überein. Der extrahierte Baum ist zusätzlich **byte-identisch**
(`diff -rq`, keine Abweichung) mit dem bereits vendierten `.harness/baseline/v6.0.0/` im
Schwester-Repo `d-check` — zwei unabhängige Wege zum selben Ergebnis, keiner ungeprüft vom anderen
übernommen.

**Datei-Diff `lab/templates/docs/plan/planning/`:** `observations.template.md` **entfernt**,
`observation.template.md` (neu, Singular — Verzeichnis-Form) und `archiv-stub-slice.template.md` /
`archiv-stub-welle.template.md` (beide neu) hinzugekommen.

## 3. Was sich **nicht** gedreht hat (Entwarnung — geprüft, nicht angenommen)

- **Docker-Hermetik ist bereits erfüllt.** `modul-02` und `modul-14` bekommen umfangreiche neue
  Abschnitte zur Mount-Disziplin (§4.4 unten) — geprüft gegen a-checks eigenen [`Dockerfile`](../../../../Dockerfile)
  und [`Makefile`](../../../../Makefile): jede Stage nutzt `COPY . .`, kein Stage mountet den
  Arbeitsbaum; `make lint`/`test`/`coverage` rufen `docker build --target <stage>` mit
  `NO_CACHE_FILTER_*`-Variablen — genau der erste der beiden in der neuen Baseline verlangten
  Griffe. Keine Handlung nötig.
- **Vendoring-Modell im Prinzip unverändert:** weiterhin `.harness/baseline/<tag>/{regelwerk,templates}/`
  aus dem self-contained Bundle, `SHA256SUMS`-Integrität via `make regelwerk-check`
  ([`MR-006`](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
  behält seinen Gegenstand).
- **WIP-Limit = 1, Slice-Lerneintrag-Formen, Review-Harness-Vokabular (`HIGH`/`MEDIUM`/`LOW`/`INFO`):**
  keine Änderung in den gelesenen Modulen.
- **Die Verzeichnis-Form der Adaptionen** (`harness/conventions/MR-<NNN>-*.md` + Index in
  `conventions.md`), 2026-08-29 mit slice-095 als Etappe C des letzten Sprungs eingeführt, ist
  bereits die von `v6.0.0` verlangte Default-Form — kein Rückstand hier, anders als beim
  vorletzten Sprung.

## 4. Die echten Brocken

### 4.1 Beobachtungs-Register: Tabelle → Verzeichnis (`modul-06`, größte Einzeländerung, +140/−30 Zeilen)

Kennung wird zum **Pfad** `BEO-<KUERZEL>/<slug>/` mit drei Dateien (`observation.md` unveränderlich,
`state.md` veränderlich, `evidence/<vorgangs-id>.md` je Beleg eine Datei); der Zähler wird
**abgeleitet** (Anzahl Evidence-Dateien) statt **geführt** (Tabellenspalte).

a-check führt [`docs/plan/planning/observations.md`](../../../../docs/plan/planning/observations/README.md) aktuell als **Tabelle**
mit **33 Einträgen** (`BEO-001`…`BEO-033`, Stand vor diesem Slice — der selbst einen 34. schreibt,
§8), gepflegtem `Zähler`-Feld und Freitext-Belegliste — das
ist eine **echte Datenmigration**, kein Formwechsel. `d-checks` slice-194/195 haben genau das als
eigene, dritte Etappe abgetrennt, weil sie sonst mit der Re-Vendoring-Etappe zusammen die
Drei-Liefer-Punkte-Grenze der eigenen Slice-Form sprengt (Modul 5) — dieselbe Größenordnung gilt
hier: 33 Einträge sind 33 neue Verzeichnisse mit je bis zu drei Dateien.

**Kopplung, nicht optional:** Sobald das neue Schema greift, ist die Kennung `BEO-<KUERZEL>/<slug>`
— das **verlangt** ein Kürzel je Sub-Area. a-check führt aktuell **keine** Kürzel-Spalte in
[`conventions.md` §Modus-Deklaration](../../../../harness/conventions.md#modus-deklaration-pro-sub-area)
(alle a-check-Kennungen — `ADR-NNNN`, `slice-NNN`, `MR-NNN` — sind bereichssegment-frei; die
Baseline befreit davon nur Repos, bei denen *keine* Kennungsklasse ein Segment führt — und stellt
im selben Atemzug fest, dass die Beobachtungs-Kennung selbst jetzt eine ist). Etappe D kann also
nicht laufen, ohne zuerst acht Kürzel zu vergeben (eines je Sub-Area-Zeile der Tabelle).

### 4.2 Zeitdokumente-Archivierung (`modul-05`/`06`/`08`/`10`, neuer Wellen-Closure-Schritt)

Fünf Closure-Schritte werden sechs: Schritt 4 „Zeitdokumente archivieren" — geschlossene Slices,
Welle-Pläne und Review-Reports wandern nach `done/<welle-id>/archiv.zip`, ein gekürzter Stub bleibt.
Die Baseline nennt das **ausdrücklich optional**: „Kein Zwang zum Nachrüsten — und kein Verbot […]
ein Repo bleibt ohne das konform." a-check hat aktuell **keine offene Welle-Datei**
(siehe [Roadmap](../in-progress/roadmap.md) §Offene Wellen) — der Trigger für diesen Schritt tritt
also nicht akut ein. Einzige Wachsamkeits-Pflicht bei künftiger Einführung: „wer archiviert, zieht
den Geltungsbereich der vorhandenen Sensoren mit" — ein Sensor, der auf `done/*.md` keilt statt auf
`done/**/*.md`, sähe archivierte Stubs nicht mehr. **Kandidat für Vertagung**, kein akuter Brocken.

### 4.3 Gate-vs-Beleg-Rollentrennung (`modul-13`, neuer Abschnitt)

Klarstellung: ein Beleg-Lauf darf nicht vom Gate-Target abhängen; `|| true` gehört an Beleg-Läufe,
nie an Gate-Läufe. **Nicht vertieft geprüft** — `grep -rn '|| true' tools/*.sh Makefile` findet
22 Stellen, alle auf den ersten Blick Shell-Idiome gegen `grep -c`-Leerfund oder tolerierte
Traversierungsfehler, keine offensichtliche Gate-Exit-Unterdrückung. Eine vollständige Prüfung
gegen die neue Zwei-Rollen-Regel steht aus (§5).

### 4.4 Docker-Hermetik-Vorschrift (`modul-02` + `modul-14`, größter Text-Zuwachs, +80/+12 Zeilen)

Siehe §3 — a-check ist bereits konform. Kein Handlungsbedarf, aber die neue Baseline macht die
bisher unausgesprochene Praxis jetzt zu einer **benannten** Regel; ein Verweis darauf in
`harness/conventions.md` wäre eine Bestätigung, kein Nachzug.

### 4.5 `versions`-Modul: Baseline-Pin-Kohärenz bislang nicht abgedeckt (`templates/.d-check.yml`)

Die neue Baseline zeigt jetzt ein Beispielmuster für einen `versions`-Pin auf
`.harness/baseline/<tag>/…` selbst. a-check betreibt `versions:` bereits (seit
[slice-133](../done/slice-133-versions-kohaerenz-zwischen-dokumenten.md)) — aber nur mit *einem*
Pattern für die Lastenheft-Versionskohärenz. Die drei Baseline-Pin-Stellen (`conventions.md`,
`AGENTS.md`, `harness/README.md`) plus
[`MR-013`](../../../../harness/conventions/MR-013-adr-vorlagen-version.md) sind bislang **nur
durch Disziplin**, nicht durch einen Sensor synchron gehalten. Ein zweiter `pin-pattern`-Eintrag
wäre die naheliegende, kleine Erweiterung.

### 4.6 `vcs`-Modul: MR-Immutabilität bislang unbelegt (`templates/.d-check.yml`)

a-checks `vcs:`-Konfiguration deckt ausschließlich ADRs (`docs/plan/adr/[0-9]*.md`,
`immutable-when: '^- **Status:** Accepted'`). Für `harness/conventions/MR-*.md` gibt es **kein**
Pendant — die Zusage „an einem akzeptierten Eintrag wird nichts nachträglich inhaltlich geändert"
(`conventions.md` §Adaptions-Block) ist reine Prosa, kein Gate. Die neue Baseline schlägt genau das
vor, mit einem an MR-Einträge angepassten Muster (`immutable-when: '^- **Datum:**'`, weil MR-Dateien
kein Status-Feld tragen — ihr Zustand ist die Verzeichnis-Position). Realer, klein umsetzbarer
Nachzug.

### 4.7 Kürzel-Spalte in der Modus-Deklaration (`grundlagen-harness-dateien.md`, `conventions.template.md`)

Direkte Konsequenz aus §4.1 — siehe dort. Kein eigenständiger Brocken.

## 5. Was diese Analyse **nicht** geklärt hat

- **Gelesen wurden zwölf Quellen** vollständig: `modul-00`, `modul-02` (Docker-Abschnitt), `modul-05`,
  `modul-06`, `modul-08`, `modul-10`, `modul-13`, `modul-14`, `grundlagen-begriffe`,
  `grundlagen-durchsetzungsschicht`, `grundlagen-harness-dateien`, `grundlagen-source-precedence`,
  `grundlagen-traceability`, `README.md` (Index), sowie die geänderten `templates/`-Dateien im Diff.
  **Nur im Diff-Stat vermessen, nicht als Volltext gelesen:** die zwölf Dateien ohne
  Zeilenänderung (§2) — für sie genügt der Nachweis „keine inhaltliche Änderung", ein Volltext-Lesen
  ist damit ausdrücklich nicht nötig.
- **`|| true`-Bestand gegen die neue Gate/Beleg-Regel (§4.3)** ist nur überflogen, nicht Zeile für
  Zeile gegen `modul-13` §Gate und Beleg geprüft.
- **Das Urteil je der 16 aktiven `MR`-Einträge** (welcher Ausgang: bleibt gültig / teilweise
  überholt / widerspricht / …) ist **nicht** gefällt — das ist die Arbeit einer Etappe B, nicht der
  Analyse.
- **Ob `docs/plan/planning/reconciliation.md` und die Brownfield-Mechanik betroffen sind** — a-check
  führt keine BF-Sub-Area, daher nicht geprüft; falls sich das ändert, wäre das nachzuholen.

## 6. Vorschlag: drei Etappen

Kleiner als beim letzten Sprung (vier Etappen bei slice-092), weil die Verzeichnis-Form der
Adaptionen bereits steht (§3) und die Docker-Hermetik bereits erfüllt ist (§3/§4.4) — zwei der vier
damaligen Etappen entfallen ersatzlos.

| Etappe | Inhalt | Warum getrennt |
|---|---|---|
| **A — Re-Vendoring** | `.harness/baseline/v6.0.0/` **neben** `v5.12.0` anlegen (SHA256SUMS bereits vorliegend, siehe §2), Pin an den drei Stellen + [`MR-013`](../../../../harness/conventions/MR-013-adr-vorlagen-version.md), `conventions.md` §Baseline auf Ist-Stand | mechanisch, sofort prüfbar, Voraussetzung für den netzlosen Vergleich beider Formen |
| **B — Adaptions-Durchgang + Kürzel-Vergabe** | die 16 aktiven `MR`-Einträge einzeln bewerten; acht Sub-Area-Kürzel vergeben (Voraussetzung für C) | berührt die Konventions-Identität; Lesearbeit, kein Diff |
| **C — Beobachtungs-Register-Migration** | ~33 Tabellenzeilen → ~33 Verzeichnisse (`observation.md`/`state.md`/`evidence/`), Wellen-Closure-Prozedur (Modul 6/8/10) auf sechs Schritte nachziehen, optional `versions`/`vcs`-Nachzüge aus §4.5/§4.6 | eigene Datenmigration mit eigenem Nachweis (Register-Paarung bleibt grün), unabhängig von B abschließbar |

**Zeitdokumente-Archivierung (§4.2) bewusst außerhalb dieser drei Etappen** — sie hat keinen
akuten Trigger (keine offene Welle) und ist von der Baseline selbst als optional deklariert; sie
wird erst fällig, wenn a-check das nächste Mal eine Welle schließt, und bräuchte dann einen
eigenen Slice.

## 7. DoD

- [x] Sprung-Umfang gemessen und die Herkunft des Release-Assets **unabhängig** belegt (§2) —
      Zip-Digest gegen GitHub-API, Baum-Identität gegen `d-check`.
- [x] Die Brocken benannt und je Brocken der betroffene a-check-Bestand **selbst geprüft**, nicht
      aus dem `d-check`-Commit übernommen (§4); Entwarnungen ebenso verifiziert (§3); Lücken der
      Analyse ausgewiesen statt Vollständigkeit behauptet (§5).
- [x] Etappen-Vorschlag steht (§6); `make gates` und `make verify` grün — Ausgabe in eine Datei,
      Exit-Code getrennt geprüft, nie in eine Pipe.

## 8. Closure-Notiz

**Geliefert:** die Messung des Sprungs `v5.12.0` → `v6.0.0`, unabhängig von der `d-check`-Migration
verifiziert (Zip-Digest, Baum-Identität) statt aus deren Commit-Message übernommen; sieben benannte
Brocken mit je einer betroffenen a-check-Stelle, vier verifizierte Entwarnungen (davon eine echte
Überraschung: die Docker-Hermetik-Vorschrift, die `modul-02`/`modul-14` neu und ausführlich
einführen, erfüllt a-check bereits) und ein Drei-Etappen-Vorschlag, kleiner als beim letzten Sprung,
weil zwei der vier damaligen Etappen (Verzeichnis-Form der Adaptionen, Docker-Hermetik) bereits
erledigt sind.

**Lerneintrag — Form: geschärfte Regel.** *Ein Migrations-Vorbild aus einem Schwester-Repo, das
denselben Fremdtext vendored, ist eine Abkürzung für die Beschaffung (Provenienz-Nachweis per
Baum-Identität statt eigenem ZIP-Download nötig) — niemals für das Urteil, welcher Fund den eigenen
Bestand betrifft.* Konkret: `d-checks` Commit-Message nennt eine lange Liste eigener
Konventions-Adaptionen und Sensor-Anpassungen (eine davon eigens für die Wellenlose-Archivierung
vergeben, vier `.claude/rules`-Aliase, Tombstones für einen entfernten `v5.18.0`-Baum); **keine**
davon hat ein Gegenstück in a-check, weil a-checks `MR`-Bestand, Sensor-
Konfiguration und Baseline-Historie eigenständig gewachsen sind. Jeder Brocken in §4 ist deshalb
gegen a-checks eigene Dateien gemessen (`grep`, `diff -rq`, Lastenheft-/`.d-check.yml`-Lektüre),
nicht aus der `d-check`-Analogie übernommen — am deutlichsten bei §4.1 (Beobachtungs-Register): der
`d-check`-Commit nennt „Beobachtungs-Register-Neugestaltung […] NICHT in diesem Slice umgesetzt",
was für a-check nichts über den eigenen Migrationsaufwand aussagt, bis die eigenen 33 Zeilen gezählt
sind. *Weil* zwei Repos, die dieselbe Baseline vendoren, an verschiedenen Punkten ihrer eigenen
Adaptions-Geschichte stehen, und eine Analogie, die das übergeht, den Migrationsaufwand systematisch
falsch schätzt.

**Zwei beobachtbare Closure-Kriterien:**

1. Die Zahlen in §2 sind mit `git diff --stat v5.12.0 v6.0.0 -- lab/regelwerk lab/templates`
   (gegen einen frischen Klon von `pt9912/ai-harness-course`) nachrechenbar, und §5 nennt die
   nicht vollständig gelesenen Module namentlich statt Vollständigkeit zu behaupten.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.**

- *Das Urteil je der 16 `MR`-Einträge steht aus* — Ausgang: **Folge-Slice**, Etappe B aus §6.
- *Die Beobachtungs-Register-Migration (33 Einträge) ist ungeschätzt im Detailaufwand* — Ausgang:
  **Folge-Slice**, Etappe C aus §6; der Umfang ist in §4.1 benannt, nicht vermessen.
- *`|| true`-Bestand gegen die neue Gate/Beleg-Regel ist nicht vollständig geprüft* — Ausgang:
  **Beobachtungs-Register**, [`BEO-034`](../../../../docs/plan/planning/observations/BEO-GATE/gate-beleg-trennung-nur-ueberflogen/observation.md) — unter der Schwelle, wird bei der
  nächsten Berührung von `modul-13` oder `tools/*.sh` erneut aufgegriffen.
- *Ob der Sprung jetzt in Etappen kommt, ist Maintainer-Entscheidung* — Ausgang: **gestrichen mit
  Begründung**: kein Risiko im Sinn der Dreier-Menge (nichts kann hier „eintreten"), sondern die
  Kernfrage, die diese Analyse laut ihrem eigenen Kopf-Vermerk („Analyse zur Abnahme") beantwortet
  bekommt, statt sie vorwegzunehmen — dieselbe Fehlklassifikation wie bei
  [`BEO-025`](../../../../docs/plan/planning/observations/BEO-PLAN/risiko-ausgang-fuer-gewollte-wirkung/observation.md).

**Folge-Slices:** keine vergeben. Drei Etappen A–C sind in §6 vorgeschlagen und brauchen die
Abnahme, bevor sie IDs bekommen.

## 9. Sub-Area-Modus

Berührt wird ausschließlich **Planungs-Harness** (`docs/plan/planning/`) — Greenfield, wie bei
slice-092.
