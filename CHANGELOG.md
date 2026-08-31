# Changelog

Alle nennenswerten Änderungen an diesem Projekt werden in dieser Datei
dokumentiert. Das Format folgt [Keep a Changelog](https://keepachangelog.com/de/1.1.0/),
die Versionierung folgt [SemVer](https://semver.org/lang/de/).

## [Unreleased]

## [0.19.0] - 2026-08-31

### Added

- **Das Release-Image liegt zusätzlich auf Docker Hub** (`AC-FA-DIST-002`, Lastenheft 0.26.0,
  [ADR-0039](docs/plan/adr/0039-spiegel-gleichheit-ist-der-config-digest.md); slice-127). Zugesagt
  ist **dasselbe Bild**, nicht ein zweiter Bau: der Spiegel-Schritt läuft **nach** dem GHCR-Push und
  vergleicht den **Config**-Digest beider Registries. Der Manifest-Digest sieht aus wie *die*
  Identität eines Image, ist aber registry-lokal (Blob-Kompression) und auf beiden Seiten
  verschieden — auch bei identischem Inhalt; ein von GHCR kopierter Manifest-Digest löst auf Docker
  Hub **nicht** auf, und wer ihn als Pin weitergibt, merkt es erst beim Konsumenten. Der Schritt ist
  **fail-closed**: schlägt der Spiegel fehl, ist das Release fehlgeschlagen, und die Fehlermeldung
  nennt den bereits veröffentlichten GHCR-Digest — der Teilstand ist gültig, und wer aufräumt, muss
  wissen, was schon draußen ist. Das Pin-Regime bleibt unangetastet: vier Stellen, ein Digest. Wer
  vom Spiegel zieht, nimmt den Digest **seiner** Registry.

### Fixed

- **Der Release startet die Hub-Darstellung wieder selbst** (slice-130). Der Job `hub-description`
  in [`release.yml`](.github/workflows/release.yml) ruft einen lokalen Workflow, der
  `contents: read` verlangt, führte aber kein eigenes `permissions:` und erbte den Workflow-Kopf
  `permissions: {}`. Ein aufgerufener Workflow bekommt nur, was der aufrufende **Job** selbst
  deklariert; GitHub bricht diesen Fall **vor dem ersten Job** ab — ohne Log, ohne Job, ohne Hinweis
  auf die schuldige Zeile. Beim Release v0.18.0 musste die Hub-Darstellung darum von Hand über
  `workflow_dispatch` gestartet werden. Die Ursache stand seit slice-127 im Repo und fiel keinem
  Gate auf, weil keines dorthin sah.

- **Der Config-Digest-Vergleich des Spiegels greift** (slice-127). `docker manifest inspect` gibt
  mehrzeiliges, eingerücktes JSON aus; ein einzeiliges `sed`-Muster traf darauf nicht und lieferte
  leer — und ein Vergleich zweier leerer Werte meldet fälschlich Gleichheit. Der Riegel dahinter
  fängt das (fail-closed), der Spiegel wäre aber blockiert statt geprüft gewesen. Ebenfalls
  korrigiert: die Hub-Referenz im CVE-Sensor.

### Harness (nicht anwender-sichtbar)

- **`make doc-workflows`** prüft die Deklarations-Form der `uses:`-Referenzen unter
  `.github/workflows/` und hängt im `gates`-Aggregat (slice-130). Geprüft wird die **Form**, nicht
  die **Gültigkeit**: dass ein Tag-Kommentar dasteht, nicht welcher — das wäre Netz. Der Sensor fand
  im ersten Lauf den oben behobenen Release-Bruch.

- **`make version-coherence`** hält **doppelt deklarierte** Versions-Angaben gegeneinander
  (slice-131): derselbe `uses:`-SHA trägt überall denselben Tag-Kommentar, und eine Versions-Variable,
  die `Makefile` **und** `Dockerfile` führen, hat an beiden Orten denselben Wert. Geprüft wird
  **Divergenz, nicht Unwahrheit** — zwei übereinstimmend falsche Angaben bleiben grün.

- **`verify-risiko-ausgaenge` prüft `in-progress/` mit**, sobald die Closure-Notiz dort ausgefüllt
  ist (slice-129). Der Auslöser ist der **Zustand** der Notiz, nicht das Verzeichnis: der alte
  Sensor sah nur `done/` und zwang damit zu einer Inhaltsänderung **nach** dem `git mv` — also zu
  genau dem Commit, den AGENTS.md §3.3 verbietet.

- **Dependabot als Hebungs-Kanal** für beide Ökosysteme dieses Repos
  ([ADR-0038](docs/plan/adr/0038-dependabot-als-hebungskanal.md); slice-128). Der Commit-Präfix
  nennt die ADR und ist damit gate-**konform**, statt per `commits.exempt-pattern` von der
  Traceability-Pflicht ausgenommen zu werden.

- **CR 5 an d-check** (ein SHA, ein Tag-Kommentar) steht übergabefertig im Slice; der Versand liegt
  beim Maintainer (slice-132).

- **Die Regelwerk-Abschnitte unter `.claude/rules/` sind Symlinks** auf die vendored Baseline
  (`MR-006`) statt einer zweiten Kopie, die dagegen driften könnte.

- **`make doc-check` hält eine Doku-Aussage gegen den Kopf des Dokuments, über das sie redet**
  (Modul `versions`; slice-133). Konkreter Anlass: `releasing.md` nannte das Lastenheft bei
  `0.17.0`, während es bei `0.26.0` stand — neun Minor-Stände, die kein Gate sah, *weil*
  `gate-consistency` nur a-checks eigene Release-Achse gegen den CHANGELOG hält und die
  Lastenheft-Achse ausdrücklich eine zweite ist. **Grenze, gemessen:** einen `sha256:`-Digest kann
  das Modul nicht tragen — der Erwartungswert kommt versions-förmig aus dem `current-from`-Span;
  die Digest-Gleichheit der harten Pins bleibt bei `tools/gate-consistency.sh`.

- **Die Commit-Range der CI fällt bei unerreichbarer Basis auf den Default-Branch**
  (`make ci-range-selftest`, `tools/ci-commit-range.sh`; slice-134). Die Weiche fing bisher nur den
  *neuen Branch* (all-zeros); ein **Force-Push** liefert einen gültig aussehenden SHA, den der
  Runner-Klon nicht kennt — `actions/checkout` holt verwaiste Objekte nicht, auch nicht mit
  `fetch-depth: 0`. Dependabot rebast bei jedem Lauf, und die Range-Prüfungen brachen darum mit
  *„Range-Basis nicht auflösbar"* ab. Die Logik liegt jetzt in einem Skript statt inline im
  Workflow — **nicht** der Lesbarkeit wegen, sondern weil ein `run:`-Block weder Selbsttest noch
  Gate hat.

## [0.18.0] - 2026-08-30

### Changed — BREAKING

- **Die Richtungs-Dimension trägt je Rolle ihr eigenes Vokabular** (`AC-FA-RULE-008` neu gefasst,
  Lastenheft 0.25.0, [ADR-0036](docs/plan/adr/0036-port-richtung-inbound-outbound.md) supersedes
  ADR-0012; slice-121). An `role: port` gilt `direction: inbound|outbound`, an `role: adapter`
  `driving|driven`. Ein Port *treibt* nichts, er wird benutzt; ein Adapter ist nicht *eingehend*,
  er treibt oder wird getrieben. Bisher galt `driving`/`driven` für **beide**, obwohl die
  Beschreibung die Äquivalenz („`driving` = primär/inbound") selbst nannte und nur die eine Hälfte
  als Wert zuließ. `port-direction-mismatch` prüft dadurch eine **Paarung** (`driving`↔`inbound`,
  `driven`↔`outbound`) statt einer String-Gleichheit; die Regel-Aussage ist unverändert.
  **Migration:** eine Zeile je Port-Schicht. Die falsche Vokabel an einer Rolle ist **Exit 2** mit
  einer Meldung, die Schicht, Wert und die für **diese** Rolle gültige Menge nennt — kein still
  akzeptiertes Alias, weil a-check kein Warn-Level kennt und die alte Schreibweise sonst unbemerkt
  bliebe. **Ebenfalls neu abgewiesen:** eine Richtung an einer Schicht **ohne** Port-/Adapter-Rolle
  (auch aus der Namens-Inferenz); sie lud bis 0.24.0 stillschweigend, war aber wirkungslos.

### Added

- **`a-check --help` gibt eine vollständige Usage-Ausgabe** (`AC-FA-CLI-003`, Lastenheft 0.24.0;
  slice-117): Kurzbeschreibung, Aufruf-Syntax mit Pfad-Parameter, Options-Liste,
  Konfigurations-Hinweis auf `.a-check.yml` und die **URL des Benutzerhandbuchs**. Dieselbe URL
  trägt der Kopfkommentar des per `--print-mk` erzeugten Fragments — es reist in ein **fremdes**
  Repo, und sein Kopf ist der einzige Ort, an dem ein Zeiger auf die Dokumentation dauerhaft
  mitfährt. **Zugesichert ist die Anwesenheit der Bestandteile, nicht ihr Wortlaut.** Die URL zeigt
  **tag-frei** auf den Hauptzweig: das Binary trägt keine eingebackene Version und könnte einen
  versionierten Link nur mit dem Stand des Vorgänger-Release füllen — dieselbe Mechanik, die
  [ADR-0030](docs/plan/adr/0030-kein-digest-im-generierten-fragment.md) für den Image-Digest
  entschied. Vorher gab `--help` nur Go's Default aus: `Usage of a-check:` plus drei Flags.

- **CVE-Scan gegen das publizierte Image** (`make image-scan`,
  [ADR-0037](docs/plan/adr/0037-cve-scan-gegen-das-publizierte-image.md); slice-124). Trivy
  digest-gepinnt gegen `ghcr.io/pt9912/a-check:latest`, mit netzlos prüfbarer Auswertung
  (`--selftest`) und einem zeitgesteuerten Workflow. **a-checks erstes Netz-Gate** — und darum
  ausdrücklich **nicht** im `gates`-Aggregat: der Scan braucht Netz für die Vuln-Datenbank, und
  das ist hier der **Zweck**, nicht ein Zugeständnis; `gates` bleibt hermetisch. Über rot
  entscheiden **nur behebbare** CRITICAL/HIGH — ein Nachtlauf, der an unbehebbaren
  Basis-Image-CVEs rot wird, ist in zwei Wochen ein weggeklicktes Abzeichen. Der Erstlauf fand
  **neun** behebbare HIGH im publizierten Image, während `make gates` über denselben Baum grün war.

### Fixed

- **Neun bekannte HIGH-Schwachstellen der Go-`stdlib` behoben** (slice-125). Die Toolchain steht
  auf `1.27.0` (Tag **und** Basis-Image-Digest gemeinsam gehoben); das neu gebaute Image trägt
  `stdlib v1.27.0` und **null** behebbare CRITICAL/HIGH — gegenüber `v1.26.4` mit neun. Gemessen
  mit demselben Trivy-Pin, gegen das exportierte Image (der Scanner liest sonst aus der Registry
  und kann lokale Bilder nicht sehen).

### Changed

- **Lint-Pin auf `v2.13.2`** (slice-126), Version und Digest gemeinsam gehoben; `make lint` läuft
  über den unveränderten Bestand grün.

- **doc-check-Tooling (`d-check.mk`): Pin des Schwester-Tools von `v0.51.1` auf `v0.69.0`**
  (slice-115, slice-080, slice-120). Das Fragment ist bei jedem Schritt **neu erzeugt**, nicht
  umgepinnt: die Einzelmodul-Targets führen eine geschlossene Verbots-Liste gegen eine **offene**
  Modul-Menge, und jedes stromaufwärts hinzugekommene Modul liefe sonst still in fünf Targets mit,
  die per Vertrag genau eines fahren. Dev-Tooling, netzlos, read-only.

### Harness (nicht anwender-sichtbar)

Diese Einträge berühren keinen Vertrag des Werkzeugs, ändern aber, woran seine Entwicklung
gemessen wird — sie stehen hier, damit ein Leser der Historie sie nicht im `git log` suchen muss.

- **Die vier Eigenbau-Sensoren aus slice-080 sind vollständig durch `d-check`-Module abgelöst**
  (slice-080, slice-120): `verify-slice-form`, `verify-slice-links`, `verify-ac-form` und die
  strukturelle Hälfte von `verify-closure-notes` — zusammen **653 Zeilen** Shell, jede mit
  Paritäts-Mutations-Beleg in beide Richtungen vor dem Entfernen. Was lokal blieb, heißt jetzt
  nach dem, was es tut (`verify-risiko-ausgaenge`).
- **`make slice-mv`** (slice-118): der Lifecycle-Wechsel zieht die Verweise **auf** die bewegte
  Datei selbst nach — repo-weit, in beiden vorkommenden Formen.
- **`doc-planning` und `doc-complete` prüfen wieder etwas** (slice-122, slice-123): das erste lief
  ohne Konfigurationsblock und meldete grün, das zweite war advisory und lief nie. Beide hängen
  jetzt in einem Aggregat.
- **Skill `cr-text-reviewer`** (slice-119): Belegbarkeit eines CR-Textes, bevor er ein fremdes
  Repo erreicht.

## [0.17.0] - 2026-08-09

### Changed

- **BREAKING: `--print-mk` gibt einen Platzhalter statt eines Digests aus (`ADR-0030`,
  `SPEC-DIST-001` 0.27.0, slice-083).** Das erzeugte Fragment trug bisher einen konkreten
  `A_CHECK_IMAGE`-Digest — zwangsläufig den des **Vorgänger**-Release, weil das Binary den Digest
  des Image, in dem es läuft, nicht kennen kann (er entsteht erst beim Push). Der Wert sah
  autoritativ aus und war falsch; **bei einem Konsumenten hat das einen realen Fehlpin verursacht.**
  Jetzt steht dort `ghcr.io/pt9912/a-check@sha256:SETZE-HIER-DEN-RELEASE-DIGEST-EIN` — bewusst
  **kein gültiger Image-Verweis**, sodass eine unveränderte Übernahme sichtbar abbricht, statt still
  ein fremdes Release zu ziehen. **Migration:** nach `--print-mk` den Digest des Release eintragen,
  aus dem das Fragment stammt (Release-Notes oder
  `docker image inspect --format '{{index .RepoDigests 0}}' <image>:<tag>`). Die **committete**
  `a-check.mk` dieses Repos trägt weiterhin den echten Digest; die Fragment-Paritäts-Prüfung in
  `tools/image-test.sh` vergleicht seit diesem Release alles **außer** der Pin-Zeile.
- **`--print-mk` ruft die Container-Runtime über `$(DOCKER)` (slice-082).** Das Fragment setzte
  wörtlich `docker` ein; ein Repo mit einer `DOCKER`-Indirektion (podman, nerdctl, ein Wrapper)
  fuhr damit die Hälfte seiner Targets anders als die andere. Jetzt steht `DOCKER ?= docker` im
  Kopf und alle Rezepte rufen `$(DOCKER)`. **Reihenfolge beachten:** `?=` greift nur, wenn `DOCKER`
  noch nicht belegt ist — eine eigene Runtime gehört **vor** das `include` oder wird hart mit `=`
  gesetzt.
- **BREAKING: `forbidden_constructs` bricht fail-closed statt still zu wirken (`ADR-0033`,
  `SPEC-CONF-001` 0.30.0, slice-086).** Der Block wurde bisher **ungeprüft** durchgereicht und hatte
  **vier** stille Ausgänge, die alle mit Exit 0 endeten: (a) die genannte Schicht existiert **nicht**
  in `layers` (Tippfehler), (b) ihre effektive Rolle ist **nicht `port`**, (c) ein **leeres Muster**,
  (d) eine **leere Musterliste**. Alle vier sind jetzt **Exit 2** mit einer Meldung, die den Grund
  nennt; Schicht-Schlüssel werden sortiert geprüft, damit bei mehreren Fehlern stets derselbe zuerst
  erscheint. Das schließt dieselbe falsch-grüne Klasse, die `languages`, `tech.adapter` und
  `constructs` längst laut abweisen. **Auslöser:** ein realer Einsatz mit sechs Schichten und einem
  Include-Muster → **0 Befunde** bei vorhandenem Verstoß. **Migration:** Einträge auf
  Nicht-`port`-Schichten wirkten nie — sie sind zu entfernen; `constructs` ist das **Gegenstück**
  (zonen-gebunden, scan-weit), **kein** Ersatz für eine Schicht-Blacklist. Gemessen: von sieben
  lokalen Konsumenten-Konfigurationen nutzt **keine** den Block. Die Bindung an `role: port` bleibt —
  sie löst `AC-FA-RULE-004` („Port-Disziplin") ein; eine Ausweitung wäre eine Lastenheft-Änderung
  und bleibt einem eigenen Slice vorbehalten.

### Documentation

- **Heuristik-Grenzen im Benutzerhandbuch, wo Konsumenten lesen (slice-084).** §6 („a-check findet
  nichts, obwohl Verstöße erwartet werden") führt jetzt die Import-Formen auf, die die
  text-heuristische Extraktion **nicht** greift: mehrere Direktiven auf einer Zeile (`import a, b`,
  `using A; using B;` — nur die erste zählt), relative Python-Importe, kompaktes TypeScript ohne
  Whitespace, import-ähnliche Zeilen in Strings/Docstrings. Die vollständige Liste stand bisher nur
  im Lastenheft; hier stehen die, die beim Konfigurieren auffallen. Zwei davon meldet a-check seit
  diesem Release zusätzlich zur Laufzeit (Grenz-Diagnose).

### Added

- **Auflösungs-Diagnose für die blinde Konfiguration (`ADR-0032`, `SPEC-CLI-001` 0.29.0,
  slice-085):** löst im **gesamten Scan** kein einziges extrahiertes Symbol auf eine Schicht auf,
  obwohl Symbole extrahiert wurden, nennt a-check je Schicht mit Symbolen eine Zeile —
  `Schicht <name>: N Datei(en), 0 von M Import-Symbolen lösen auf eine Schicht auf`. Das deckt die
  gefährlichste Konfiguration ab, die das Werkzeug zulässt: **alle Dateien in Schichten, alle
  Symbole extrahiert, und trotzdem wird keine Kante beurteilt**, weil jedes Ziel über den
  fail-open-Pfad als repo-extern gilt. Gemessener Auslöser: ein Mono-Scan mit sprach-präfixierten
  Globs (`go/internal/ui/**`) meldete 0 Befunde, obwohl ein Verstoß eingebaut war — Go-Importe
  tragen den Modulpfad, in dem das Präfix nicht vorkommt. **Die Auslösung ist repo-weit, nicht je
  Schicht:** eine einzelne Schicht ohne auflösende Symbole ist legitim (ein abhängigkeitsfreier
  Kern importiert nur die Standardbibliothek) und von kaputter Auflösung nicht zu unterscheiden.
  Daraus folgt ausdrücklich, dass der **Teil**ausfall — eine von mehreren Schichten falsch
  aufgelöst — **still bleibt**. Schichten ohne Symbole werden nicht genannt, `composition_root`
  zählt nicht; **Exit-Code unberührt**, Meldung auch bei null Befunden, stabil nach Schichtnamen
  sortiert.

- **Grenz-Diagnose für nicht beurteilbare Import-Zeilen (`ADR-0031`, `SPEC-CLI-001` 0.28.0,
  slice-081):** ein Scan weist jetzt auf stderr — nach der Abdeckungs-Diagnose — die Import-Zeilen
  als `pfad:zeile: form` aus, deren **Schreibweise** per Konstruktion zu keiner prüfbaren Kante
  führt. Zwei Klassen: (1) **nicht extrahiert** — relativer Python-Import (`from . import x`) und
  zweite Direktive auf derselben Zeile (`import a, b`, `using A; using B;`); (2) **extrahiert,
  aber strukturell unauflösbar** — ein `./`/`../`-Pfad unter einem `resolution`-Modus, der nicht
  `relative` ist (unter `relative` bleibt dieselbe Zeile still). Bisher stand diese bewusste
  Grenze (`AC-QA-02`) nur in der Doku; am geprüften Baum sagte a-check nichts, und ein Konsument
  musste je Sprach-Skelett von Hand einen Verstoß einbauen, um seine blinden Stellen zu finden.
  **Der Exit-Code bleibt unberührt** (advisory), und die Meldung erscheint **gerade auch bei null
  Befunden** — dort trennt sie „sauber" von „nicht angesehen". Ein Symbol, das nur im konkreten
  Baum kein Ziel findet, wird **nicht** gemeldet: das ist von repo-externem Code nicht
  unterscheidbar. Stabil nach (Pfad, Zeile) sortiert, ab zehn Zeilen gekürzt **mit ausgewiesener
  Restzahl**; ein Baum ohne solche Zeilen erzeugt **keine** Ausgabe. Die Befundmenge selbst bleibt
  unverändert — die Diagnose fasst die Regel-Auswertung nicht an.

## [0.16.0] - 2026-07-25

### Added

- **Abdeckungs-Diagnose für schichtlose Dateien (`ADR-0029`, `SPEC-CLI-001` 0.26.0, slice-043):** ein
  Scan nennt jetzt auf stderr — nach der Zusammenfassung — die gescannten Dateien, die in **keinem**
  `layers`-Glob liegen und damit **keiner Schicht-Regel** unterliegen. Bisher war diese bewusste
  fail-open-Grenze (`AC-QA-02`) unsichtbar: ein grünes Gate über einem teilweise ungeprüften Baum
  sah aus wie eines über einem geprüften. **Der Exit-Code bleibt unberührt** (advisory);
  vollständige Abdeckung erzeugt **keine** Ausgabe — die Meldung ist Signal, nicht Rauschen
  (verifiziert: sechs von sieben lokalen Konsumenten-Konfigurationen bleiben diagnose-frei).
  `composition_root`-Dateien zählen nicht, `exclude`-Dateien sind nie im Scan; gezählt wird nur die
  **Quell**-Seite (ein Import-**Ziel** ohne Schicht ist von repo-externem Code nicht unterscheidbar).
  Pfade stabil sortiert, ab zehn Dateien gekürzt **mit ausgewiesener Restzahl**. Eine Opt-in-Strenge
  (`strict_coverage` ⇒ Exit 1) ist ausdrücklich **vertagt**. Abhilfe beim Konsumenten: Schicht in
  `layers` deklarieren oder Datei in `exclude` aufnehmen.

- **Roh-Text-Konstrukt-Monopol `constructs` / Regel `construct-leak` (`AC-FA-RULE-011`, `ADR-0027`,
  `SPEC-CONF-001`/`SPEC-RULE-001`/`SPEC-EXTRACT-001` 0.24.0, slice-042):** ein neuer optionaler
  `constructs`-Block hebt die bewährte `tech`-Scoping-Mechanik — Zone als Pfad **oder** Pfad-Liste,
  `match: substring|regex` (RE2), `composition_root: allow|forbid` — von **extrahierten
  Import-Symbolen** auf **Roh-Quelltext**: das Muster darf nur in seiner Zone vorkommen, jedes
  Vorkommen außerhalb ist ein Befund `construct-leak` (Exit 1). Damit sind Invarianten gatebar, die an
  einem Konstrukt **ohne Import-Zeile** hängen — der Auslöser ist das `dlopen`/`dlsym`/`dlclose`-
  **Aufruf**-Monopol eines C++-Konsumenten, das bisher als `grep`-Skript im Konsumenten-Repo lebte.
  Die Prüfung ist **scan-weit** (auch Dateien in keinem `layers`-Glob; `exclude` greift davor) und
  läuft auf der **kommentar-bereinigten** Quelle — ein Treffer nur im Kommentar meldet nicht, eine
  ausgewiesene Divergenz zur `grep`-Referenz. Fail-closed beim Laden (leeres/fehlendes
  `pattern`/`adapter`, unbekanntes `match`/`composition_root`, ungültige Regex → Exit 2); **opt-in**:
  ohne Block ändert sich nichts. Verifiziert per Paritätsprobe gegen die `grep`-Referenz und per
  Fitness-Probe am realen Konsumenten-Baum (injizierter Aufruf gefangen, Zone grün); die vier lokalen
  Bestandskonsumenten bleiben bei 0 Befunden.

### Fixed

- **Kein leerer Quell-Schicht-Name mehr im `wrong-direction`-Befund (slice-043):** importierte eine
  gescannte Datei **ohne** Schicht eine deklarierte Schicht, rendete die Meldung ein Loch
  (`wrong-direction:  -> ui`). Sie weist die Quelle jetzt als `(ohne Schicht)` aus — das
  quell-seitige Symptom derselben Abdeckungs-Lücke, die die neue Diagnose meldet.

- **Kein Fehlbefund mehr bei unauflösbarem Ziel-Glob (`ADR-0028`, `SPEC-RULE-001` 0.25.0, slice-044):**
  ein Layer-Glob mit **Wildcard in der Mitte** (`…/application/**/ports/**`) kann als Import-**Ziel**
  nicht auflösen; der Kandidat fiel bisher auf das nächst-passende, **umschließende** Glob zurück.
  Folge war ein **Falsch-Positiv** — ein Adapter, der über eine **deklarierte** `adapter → ports`-Kante
  einen verschachtelten Port importiert, wurde als `wrong-direction: adapter -> application` gemeldet.
  Besonders teuer, weil die naheliegende Reparatur (die falsche Kante deklarieren) danach **echte**
  Verstöße dauerhaft verdeckt hätte. Jetzt wird die Zuordnung **zurückgezogen**: das Ziel gilt als
  **extern** (kein Befund), nie als die umschließende Schicht. Der Rückzug ist eng gefasst — literaler
  **Kopf** *und* **Tail-Marker** des anderen Globs müssen im Kandidaten vorkommen und der Kopf
  mindestens so spezifisch sein wie der gewählte Präfix; ohne Tail-Marker bleibt alles wie bisher.
  Die Datei→Schicht-Zuordnung (Quell-Seite) ist unberührt: Port-Dateien unter solchen Globs bleiben
  voll geprüft. **Inert für den Bestand** — b-cad, d-check, d-migrate, m-trace, belief-agent und das
  HexSlice-Beispiel bleiben unverändert bei 0 Befunden (keiner nutzt Innen-Wildcards).

### Changed

- **Befund-Sortierung ist jetzt eine Totalordnung (`AC-QA-01`, `SPEC-DET-001`, slice-042):** neben
  Pfad, Zeile und Regelname entscheidet zuletzt die **Meldung**. Eine Datei-Zeile kann mehrere Befunde
  **derselben** Regel tragen (zwei `constructs`- oder zwei `forbidden_constructs`-Muster); ohne den
  letzten Schlüssel hing deren Reihenfolge an der internen Eingabe-Ordnung. Die Lücke bestand latent
  bereits vor `constructs`.

- **`--print-graph`-Legende lesbarer (slice-041):** der Legenden-Knoten listet die kategorischen Regeln
  und Stil-Hinweise jetzt **je eine pro Zeile** (statt einer langen, mitten im Wort umbrechenden Zeile)
  und wickelt den Text in ein `<div style='text-align:left'>` für **linksbündige** Ausrichtung, wo der
  Renderer inline-Styles zulässt (sonst harmlos zentriert). Reine Präsentation — gleiche Inhalte,
  deterministisch, kein Vertrags-/Spec-Bump.

## [0.15.0] - 2026-07-24

### Added

- **HexSlice Vertical-Slice-Regeln `lateral-slice` + `port-locality` (`AC-FA-RULE-009`/`AC-FA-RULE-010`,
  `ADR-0026`, `SPEC-RULE-001` 0.22.0, slice-039):** zwei neue kategorische Regeln gaten die
  **Vertical-Slice-Achse** von HexSlice über das bestehende Rollenmodell. `lateral-slice` meldet, wenn
  eine `app`-Datei eine **fremde Use-Case-Slice derselben `app`-Schicht** importiert (Slices = per-Glob-
  Untereinheiten *einer* Schicht; getrennte `app`-Layer bleiben edge-regiert). `port-locality` meldet,
  wenn eine `app`-Datei einen **im Application-Baum geschachtelten** Port außerhalb dessen pfad-
  abgeleiteten Scope-Verzeichnisses importiert (use-case-lokal ⊂ business-area ⊂ app-weit; nur
  `app`-Importeure, Adapter-Implementierung nicht erfasst). Beide sind **opt-in** über saubere
  Präfix-Globs und lassen klassisch-hexagonale Configs (Geschwister-Ports, Sub-Layer mit Kante)
  unberührt — verifiziert gegen b-cad/d-check/d-migrate (0 Befunde). Benutzerhandbuch §3.7 zeigt die
  Einrichtung als Arbeitsanleitung.
- **`--print-graph`-Legende nennt alle fünf kategorischen Regeln (`AC-FA-CLI-002`, `SPEC-CLI-002` 0.23.0,
  slice-040):** die Mermaid-Legende führt zusätzlich `lateral-slice` und `port-locality` (reine
  Legenden-Notiz, keine gezeichnete Kante).

### Fixed

- **Layer-Tie-Break folgt der Deklarationsreihenfolge (`ADR-0013`-Konformität, slice-038):** bei
  Literal-Präfix-**Gleichstand** zweier `layers`-Globs entschied faktisch die alphabetische Reihenfolge
  statt der dokumentierten „zuerst deklarierten" Schicht — `layers` wurde als Map dekodiert und
  alphabetisch sortiert. Der Decode erhält jetzt die **Dokumentreihenfolge** (`decodeLayers` über eine
  rohe `yaml.Node`); zwei fail-closed-Prüfungen, die der Map-Decode gratis lieferte (Duplikat-Schlüssel,
  Nicht-Mapping-`layers`), sind explizit rekonstruiert. Kein neuer Vertrag/ADR/Lastenheft-Bump.

## [0.14.0] - 2026-07-23

### Fixed

- **`exclude` beschneidet den Verzeichnis-Walk (`ADR-0025`, `SPEC-EXTRACT-001` 0.21.0, slice-035):**
  `exclude` filterte bisher nur einzelne Dateien — der Scan-Walk stieg trotzdem in jeden Ordner ab.
  Ein ausgeschlossener, aber **unlesbarer** Teilbaum (z. B. ein Trivy-Cache unter `.security/**`)
  brach den Scan mit Exit 2 ab, obwohl `.security/**` in `exclude` stand. Jetzt wird ein Verzeichnis,
  dessen ganzer Teilbaum von einem rekursiven Muster (`**` oder `<präfix>/**`) gedeckt ist, **gar nicht
  erst betreten** (Prune vor dem Lesen des Ordnerinhalts) — der Abbruch entfällt, große Fremdcode-
  Teilbäume (`**/node_modules/**`, `**/dist/**`) werden übersprungen statt durchlaufen. Der Prune ist
  **beweisbar output-äquivalent** zum Datei-Ausschluss: nur teilbaum-deckende Muster prunen (ein
  Teil-Muster wie `src/*` prunt **nicht**, sonst gingen nicht ausgeschlossene Dateien still verloren);
  ein **nicht** ausgeschlossener unlesbarer Ordner bricht weiterhin fail-closed ab. Realisiert die
  Verzeichnis-Absicht von `ADR-0018`; kein Lastenheft-Bump (Schärfung des Wie).

## [0.13.0] - 2026-07-09

### Added

- **`--print-graph`: Architektur-Graph als Mermaid (`AC-FA-CLI-002` 0.19.0, `ADR-0024`,
  `SPEC-CLI-002`, slice-032):** `a-check --print-graph [pfad]` rendert die in `.a-check.yml`
  **deklarierte** Architektur als Mermaid-`flowchart` auf stdout — ein Knoten je Schicht (nach
  effektiver Rolle gefärbt), eine Kante je `edges`-Eintrag, eine gestrichelte Kante je `allow`-Eintrag,
  `composition_root`/`adapter_sink` als isolierte Notizknoten, implizite Regeln als Legende.
  **Read-only, kein Scan, deterministisch** (byte-identische Ausgabe; stabile interne Knoten-IDs +
  Escaping-Vertrag, sodass Mermaid-heikle Layer-Namen die Syntax nicht brechen; unbekannte
  `edges`/`allow`-Endpunkte werden als Dangling-Knoten sichtbar). Ladezeitiger Config-Fehler (inkl.
  unbekannter Sprache), unbekanntes Flag oder ein Restargument nach dem Pfad → Exit 2. Umsetzung: neuer
  `graph`-Präsentationsadapter (`ARC-007`) hinter dem driven Port `GraphPort`; `ExtractionPort` um einen
  validation-only `Validate`-Einstieg (kein Datei-Walk) erweitert; `core.EffectiveRole` als von
  Regel-Engine und Renderer geteilter Rollen-Resolver.
- **`a-check-graph`-`make`-Target im `a-check.mk`-Fragment (`AC-FA-DIST-001` 0.20.0, slice-033):**
  `--print-mk` liefert neben `a-check` (Scan-Gate) jetzt ein `a-check-graph`-Target, das
  `--print-graph` über dasselbe digest-gepinnte `A_CHECK_IMAGE` + netzlosen read-only-Mount ausführt
  (`make a-check-graph > architektur.mmd`) — Convenience für Konsumenten, die bereits `include a-check.mk`
  fahren; kein zweiter Digest.

## [0.12.0] - 2026-07-06

### Fixed

- **Split-Package-Auflösung (`AC-FA-CONF-001`/`AC-FA-EXTRACT-001` 0.17.0→0.18.0, slice-031,
  ADR-0023):** bei `mode: fixed-root` mit ≥ 2 `roots` löst ein importiertes **Top-Level-Symbol**,
  dessen Datei ≠ Symbolname ist (Kotlin-Extension-Funktion, zweite Klasse je Datei), jetzt über die
  **reale Top-Level-Deklaration** auf sein Modul auf — statt an einem **Split-Package** (dasselbe
  Paket real über zwei Schicht-Module) mit **Exit 2** den ganzen Scan abzubrechen. Evidenz-Rangfolge
  *deklariert > nur-Paketverzeichnis > keine*: genau ein deklarierender Root ⇒ eindeutig; real
  **deklariert** in ≥ 2 verschiedenen Schichten ⇒ Exit 2 (fail-closed); kein Deklarations-Treffer ⇒
  extern (fail-open, ebenso ein Wildcard-/Paket-Import über eine Schicht-Grenze); ein **eindeutiges**
  Paket-Verzeichnis löst rückwärtskompatibel. Das **Kotlin**-Backend liefert dafür zusätzlich die
  Top-Level-Deklarationen (übrige Backends no-op). ADR-0023 Supersedes ADR-0022. Anlass:
  d-migrate-Pilot — der `asJdbc`-Exit-2 real reproduziert und getilgt; schärft AC-QA-02.

## [0.11.0] - 2026-07-05

### Fixed

- **KMP-/Multi-Modul-Auflösung (`AC-FA-CONF-001` 0.16.0→0.17.0, slice-027, ADR-0022):** `mode:
  fixed-root` mit ≥ 2 `roots` löst den internen FQN jetzt **datei-mengen-bewusst** gegen die real
  gescannten Dateien auf (endungs-agnostisch, package==directory) statt je Root einen Phantom-
  Kandidaten am Wurzel-Präfix zu bilden. Damit **lädt und löst** die legitime disjunkte KMP-Multi-
  Modul-Config (geteiltes `package_base`, disjunkte Sub-Namespaces) korrekt — die verbotene
  `domain → application`-Kante wird gemeldet (vorher: still falsch-grün **oder** Reject des
  Ladezeit-Guards). Der statische Guard aus 0.16.0 entfällt (ADR-0022 Supersedes ADR-0020); echte
  Mehrdeutigkeit (gleicher FQN real in ≥ 2 Roots, **verschiedene** Schichten) bricht nach dem Scan
  mit Exit 2, same-layer `expect`/`actual` löst sauber. Anlass: belief-agent-KMP-Bericht; schärft
  AC-QA-02. `--print-config` dokumentiert die Multi-Modul-Resolution.

### Added

- **Release-Register `version.md` + Pin-Konsistenz-Gate (slice-018):** neues `version.md`
  (Repo-Wurzel, `#aktuell`-Anker) als *eine* Wahrheit für die aktuelle Release-Koordinate
  (Version, Datum, voller `@sha256:`-Digest). `make gate-consistency` prüft jetzt zusätzlich die
  Pin-Konsistenz: der Digest ist in `a-check.mk`, `internal/cli/cli.go`, dem README-`docker run`-
  Beispiel und `version.md#aktuell` identisch, die Version stimmt mit dem aktuellsten CHANGELOG-
  Release, und `d-check.mk` trägt eine wohlgeformte Tag/Digest-Deklaration. Jede harte Pin-Datei
  muss genau **einen** Digest tragen (fail-closed gegen einen Decoy-Zweitdigest). Ein Selbsttest
  beweist die Fitness-Function offline für alle Dimensionen — ein gedrifteter Digest/eine falsche
  Version macht `make gates` rot. README-/Handbuch-/
  `releasing.md`-Prosa verlinkt auf `version.md#aktuell` statt literaler Nummern (Opt 1). Schließt
  die stille Pin-Drift, die den stale README-`v0.2.0`-Pin am 2026-07-01 nur per Zufalls-Audit
  auffallen ließ; schärft AC-QA-03 (Reproduzierbarkeit). Netzlose Grenze (AC-QA-02): die
  Tag→Digest-Auflösung wird beim Re-Pin (online) verifiziert. d-checks tag-basierte Module
  `versions`/`pins` passen nicht auf a-checks Digest-Pins (verworfen: Opt 2).

### Changed

- **Traceability-Gate (`make trace-check`):** von `tools/trace-check.sh` (98 Zeilen bash)
  auf das d-check-Modul `commits` umgestellt (ADR-0021, slice-030) — eine Skript-Kopie
  weniger; Verhalten verifiziert gleich (ID-Pflicht, Merge-/Revert-Ausnahme, fail-closed bei
  unauflösbarer Range). Der commit-msg-Hook läuft jetzt über das digest-gepinnte Image.
  Dev-Tooling; die Traceability-**Regel** (AGENTS §5) unverändert.
- **doc-check-Tooling (`d-check.mk`):** Pin des Schwester-Tools `d-check` von **v0.35.0**
  auf **v0.37.1** gehoben (`@sha256:3bbdb19b…`, via `DCHECK_DIGEST`); `d-check.mk` verbatim
  aus `v0.37.1 --print-mk` regeneriert (alle 10 `doc-*`-Targets, in AGENTS §4 als advisory
  gelistet). Dev-Tooling, netzlos, read-only; a-checks aktive `.d-check.yml`-Module
  unverändert (slice-019).

## [0.10.0] - 2026-07-04

Lastenheft/Spezifikation 0.15.0 → 0.16.0 (slice-026; ADR-0020, belief-agent-KMP-Evidenz)
— fail-closed-Guard gegen mehrdeutige Mehr-Wurzel-Auflösung; bringt den Guard ins Image.

### Added

- **`AC-FA-CONF-001` (0.15.0→0.16.0):** fail-closed-Guard gegen **mehrdeutige
  Mehr-Wurzel-Auflösung** — `mode: fixed-root` mit ≥ 2 `roots`, von denen zwei Roots je
  eine andere Schicht **erzwingen** (die Schicht, in die ein Root allein am Wurzel-Präfix
  auflöst — längster passender Glob-Präfix, exakt wie die Import-Auflösung), bricht mit
  Exit 2 statt still falsch-grün. Anlass: belief-agent-Bericht — KMP (`commonMain`/`jvmMain`
  teilen `package_base`) bei flachen Source-Set-Globs fing die illegale `core → adapter`-
  Kante nicht (Phantom-Kandidaten Root × Paketpfad; die Zuordnung entschied der längste
  Präfix statt das Symbol). Meldung nennt Roots + Schichten + Rezept (paket-spezifische
  Globs tiefer als die Roots). Verschachtelte Schichten, unter denen beide Roots dieselbe
  Schicht erzwingen, laden unverändert; die **asymmetrische** Phantom-Form (nur ein Root
  erzwingt eine Schicht) bleibt dokumentierte Grenze (`AC-QA-02`). Stufe 2
  (datei-mengen-bewusste Auflösung) gated als slice-027.

### Changed

- **Selbst-Pin** (`--print-mk`/`a-check.mk`/`cli.go`-`aCheckImage`) auf den **v0.10.0**-Digest
  `@sha256:0932cb1d…` gehoben (AC-QA-03, ADR-0004/ADR-0007).

## [0.9.0] - 2026-07-04

Lastenheft/Spezifikation 0.14.0 → 0.15.0 (slice-024; ADR-0019, b-cad-Pilot-Evidenz)
— dieses Release bringt die Root-Sub-Einheit ins Image und entsperrt den
b-cad-Pilot-Schnitt (M3).

### Changed

- **`AC-FA-RULE-002` (0.14.0→0.15.0):** `lateral-adapter`-Sub-Einheiten präzisiert —
  **Blatt-Klassifikation**: ein datei-förmiges Blatt (`.`) direkt im Layer-**Root**
  gehört zur Root-Sub-Einheit (Sub-Einheiten sind **Verzeichnisse**, keine
  Dateinamen), ein verzeichnis-förmiges Blatt (Go-Paket-Pfad) **ist** die
  Sub-Einheit; Root↔Root same-layer meldet nicht mehr, Root↔Unterverzeichnis,
  Cross-Paket und Cross-Layer unverändert kategorisch. Bewusste Gate-Lockerung per
  ADR-0019: die b-cad-Vollrichtungs-Config erzeugte 40 Falsch-Positive der Klasse
  `x.cpp → x.h` bei 0 echten Verstößen; pro-Adapter-Layer (Voraussetzung der
  `direction`-Modellierung) sind damit falsch-positiv-frei. Endungslose
  Datei-Specifier (TS `./b`) bleiben dokumentierte Grenze (`AC-QA-02`).
- **Selbst-Pin** (`--print-mk`/`a-check.mk`/`cli.go`-`aCheckImage`) auf den **v0.9.0**-Digest
  `@sha256:0378211f…` gehoben (AC-QA-03, ADR-0004/ADR-0007).

## [0.8.0] - 2026-07-03

Spezifikation 0.13.0 → 0.14.0 (folgt dem Lastenheft-CR 0.14.0, **CR d-check-Pilot**,
slice-023; ADR-0018): die drei Deltas, auf die d-check seine `arch-check`-Ablösung
als Vorbedingung gestellt hat — dieses Release entsperrt den dortigen Umbau.

### Added

- **`AC-FA-RULE-003`/`AC-FA-CONF-001` (0.13.0→0.14.0):** `tech.adapter` auch als
  Pfad-**Liste** — das Symbol ist in **jedem** gelisteten Adapter erlaubt (leere
  Liste/leerer Eintrag → Exit 2; der nicht-leere Skalar bleibt byte-identisch); die
  `tech-leak`-Meldung nennt alle gelisteten Adapter in Deklarationsreihenfolge.
  **Fail-closed-Härtung:** ein leerer oder **fehlender** `tech.adapter` bricht jetzt
  mit Exit 2 — vor 0.14.0 war er ein **stiller Never-Leak-Eintrag** (das Muster
  meldete nie; `AC-QA-02`-Ethos wie beim leeren `resolution`-Root).
- **`AC-FA-RULE-003`/`AC-FA-CONF-001` (0.14.0):** `composition_root: allow|forbid`
  je `tech`-Eintrag (Default `allow` = bisheriges Verhalten): `forbid` schaltet nur
  die `tech-leak`-Ausnahme der Composition Root für diesen Eintrag ab — die
  Schicht-Regel-Ausnahme des Verdrahtungspunkts bleibt; anderer Wert → Exit 2.
- **`AC-FA-CONF-001` (0.14.0):** optionaler **`exclude`**-Block (ADR-0018) —
  Top-Level-Datei-Globs relativ zur Scan-Wurzel; Treffer fallen **vor** der
  Extraktion vollständig vom Scan (auch `forbidden_constructs`); leerer Glob →
  Exit 2; ohne Block byte-identisch. Evidenz: d-check (`**/*_test.go`) + m-trace
  (`node_modules/` je Workspace, `dist/`, `*.d.ts`-Suffix-Falle).

### Changed

- **Selbst-Pin** (`--print-mk`/`a-check.mk`/`cli.go`-`aCheckImage`) auf den **v0.8.0**-Digest
  `@sha256:a1c9c4d6…` gehoben (AC-QA-03, ADR-0004/ADR-0007).

## [0.7.0] - 2026-07-03

Lastenheft/Spezifikation 0.12.0 → 0.13.0: **TypeScript-Backend** plus der bislang
reservierte **`relative`-Auflösungs-Modus** (ADR-0017, erweitert ADR-0016); slice-022.
Bringt beides ins veröffentlichte Image.

### Added

- **`AC-FA-EXTRACT-001` (Lastenheft/Spezifikation 0.12.0→0.13.0):** achtes Sprach-Backend
  **TypeScript** (`languages`-Schlüssel `typescript`) — ES-Module-Formen → Modul-Specifier
  in `'…'`/`"…"` (Semikolon optional/ASI): `import … from` (inkl. `import type`),
  Seiteneffekt-Import, Re-Exports `export … from`, `import X = require(…)` sowie die
  Fortsetzungszeile `} from '…'` mehrzeilig (Prettier-)umbrochener Imports; der Mittelteil
  ist auf Import-Clause-Zeichen beschränkt — Ausdrucks-Zeilen (`knex.from('users')`,
  dynamisches `import()`/`require()`) matchen nie (`AC-QA-02`-Grenze).
- **`AC-FA-CONF-001` (0.12.0→0.13.0):** `resolution.mode: relative` (ADR-0017) — Specifier
  `.`/`..`/`./…`/`../…` lösen lexikalisch gegen das Verzeichnis der importierenden Datei auf
  (Quellpfad-Threading bis `targetLayer`); Bare-Imports und Wurzel-Escapes liefern eine
  **leere** Kandidatenmenge (kein Geister-Match); `roots`/`package_base` bei `relative` →
  Exit 2; nur noch `namespace` reserviert.

### Changed

- **Selbst-Pin** (`--print-mk`/`a-check.mk`/`cli.go`-`aCheckImage`) auf den **v0.7.0**-Digest
  `@sha256:41eb368e…` gehoben (AC-QA-03, ADR-0004/ADR-0007).
- **`lateral-adapter`** prüft Sub-Einheit und `adapter_sink` auf dem gemäß `resolution`
  **normalisierten Ziel-Kandidaten** statt am Roh-Symbol (Review-Fix slice-022; `path`-Modus
  verhaltens-identisch — unter `resolution`-Modi `adapter_sink` als Pfad-Fragment schreiben).

## [0.6.0] - 2026-07-02

Sprach-Backend seit `v0.5.0`; Lastenheft/Spezifikation 0.11.0 → 0.12.0.
Bringt das **C#-Backend** (`using`-Direktiven) ins veröffentlichte Image; die
Schicht-Auflösung nutzt den bestehenden `resolution`-Block (`fixed-root`-Rezept).

### Added

- **`AC-FA-EXTRACT-001` (Lastenheft/Spezifikation 0.11.0→0.12.0):** siebtes Sprach-Backend
  **C#** (`languages`-Schlüssel `csharp`) — `using`-**Direktiven** → gepunkteter Namespace
  (`global`/`static` übersprungen, Alias-Form liefert ihr **Ziel**); das Pflicht-`;` direkt nach
  dem Namen schließt `using`-**Statements** (`using var …`, `using (…)`) kategorisch aus.
  Schicht-Auflösung über den `fixed-root`-Modus unter der .NET-Konvention Namespace==Verzeichnis
  (`AC-QA-02`-Grenze; frei deklarierte Namespaces bleiben unaufgelöst — der reservierte
  `namespace`-Modus/Index bleibt Exit 2, eigener Folge-Slice). slice-021.

### Changed

- **Selbst-Pin** (`--print-mk`/`a-check.mk`/`cli.go`-`aCheckImage`) auf den **v0.6.0**-Digest
  `@sha256:b349a150…` gehoben (AC-QA-03, ADR-0004/ADR-0007).

## [0.5.0] - 2026-07-02

Sprach-Backend + Import-Auflösung seit `v0.4.0`; Lastenheft/Spezifikation 0.8.0 → 0.11.0.
Bringt das **Python-Backend** und den sprach-parametrischen **`resolution`-Block**
(fixed-root/dotted, Mono-Repo-tauglich) ins veröffentlichte Image; unbekannte
`languages`-Schlüssel brechen jetzt mit Exit 2 statt still falsch-grün.

### Added

- **`AC-FA-EXTRACT-001` (Lastenheft/Spezifikation 0.10.0→0.11.0):** sechstes Sprach-Backend
  **Python** — `import a.b.c` (inkl. Alias) und `from a.b import c` → gepunkteter Modulpfad;
  Schicht-Auflösung über den bereits gelieferten `fixed-root`-Modus (Rezept `package_base` =
  Top-Package + `roots`, dokumentiert im Benutzerhandbuch §4). Relative Importe (`from .`) werden
  nicht extrahiert — Signal des reservierten `relative`-Modus, dokumentierte Grenze (`AC-QA-02`).
  Python wird **nicht** C-Kommentar-gestrippt (`prepSource`, Review-Befund: eine `/*`-Bytefolge in
  einem Python-String-Literal wie `"**/*.py"` hätte sonst echte Folge-Imports verschluckt —
  falsch-grün). Benutzerhandbuch-Currency: `resolution`-Block (slice-015) dort nachdokumentiert;
  „unbekannte Sprache wird ignoriert" auf das Exit-2-Verhalten (0.9.0) richtiggestellt. slice-020.
- **`AC-FA-CONF-001` (Lastenheft/Spezifikation 0.9.0→0.10.0):** optionaler `resolution`-Block — Map
  **Sprache → `{mode, roots, package_base}`** (`mode ∈ {path, fixed-root}`; `relative`/`namespace`
  reserviert → Exit 2). Löst gepunktete (JVM/Python) und `src`-gewurzelte (C++) Importe **pro Sprache**
  auf ihre Schicht auf (Mono-Repo-tauglich) — Sprach-Threading via `FileImports.Language`; `.`→`/` an
  `package_base` gebunden (Pfad-Sprachen behalten `.`-Endungen); Grenze Paket==Verzeichnis. Default
  (ohne Block) unverändert, rückwärtskompatibel. Prerequisit fürs Python-Backend. ADR-0016 (erweitert
  ADR-0014); slice-015.
- **`AC-FA-CONF-001` (Lastenheft/Spezifikation 0.8.0→0.9.0):** ein `languages`-Schlüssel außerhalb der
  unterstützten Backends (`cpp`/`go`/`rust`/`kotlin`/`java`) bricht mit **Exit 2** ab statt still nichts
  zu extrahieren (falsch-grün); Backend-Registry als Single Source. slice-017.

### Changed

- **doc-check-Pin** (Schwester-Tool `d-check`) von v0.24.0 auf **v0.35.0** digest-gepinnt (`@sha256:9d7b23ac…`);
  Gate-Tooling, netzlos, a-checks aktive Module unverändert.
- **Selbst-Pin** (`--print-mk`/`a-check.mk`/`cli.go`-`aCheckImage`) auf den **v0.5.0**-Digest
  `@sha256:81951e61…` gehoben (AC-QA-03, ADR-0004/ADR-0007).

## [0.4.0] - 2026-07-01

Regel-Engine-Schärfung + Sprach-Backend seit `v0.3.0`; Lastenheft/Spezifikation 0.6.0 → 0.8.0.
Bringt `match: regex` (b-cad-Regel E) und das Java-Backend ins veröffentlichte Image.

### Added

- **`AC-FA-RULE-003`/`AC-FA-CONF-001` (Lastenheft/Spezifikation 0.7.0→0.8.0):** `tech`-Muster
  optional als **RE2-Regex** — `match: substring|regex` je Eintrag (Default `substring`,
  rückwärtskompatibel/byte-identisch ohne `match`). Macht ein nur als Muster fassbares Framework
  wie Qt (`Q[A-Za-z]`) ausdrückbar (schließt die letzte Lücke zum b-cad-`arch-check.sh`-Ersatz,
  Regel E). Mehrfach-Treffer lösen in **Deklarationsreihenfolge** (Erst-Treffer) — die Spec-Aussage
  „längster Präfix" galt für `tech` nie und ist richtiggestellt. Unbekanntes `match`/nicht
  kompilierbare Regex → Exit 2. ADR-0015; welle-05/-06 (b-cad-Pilot); slice-016.
- **`AC-FA-EXTRACT-001` (Lastenheft 0.6.0→0.7.0):** fünftes Sprach-Backend **Java**
  (`languages`-Schlüssel `java`; `import …;` inkl. `import static …;` — das `static`
  übersprungen, `;` ignoriert, Wildcard heuristisch). Text-heuristisch wie die übrigen
  Backends, innerhalb ADR-0002 (kein neuer ADR); getrieben vom Konsumenten-Bedarf
  (belief-agent). welle-06; slice-014.

### Changed

- **doc-check-Pin** (Schwester-Tool `d-check`) von **v0.24.0** auf **v0.35.0**
  (`@sha256:9d7b23ac…`) gehoben — Gate-Tooling, netzlos; a-checks aktive Module unverändert.
- **Selbst-Pin** (`--print-mk`/`a-check.mk`/`cli.go`-`aCheckImage`) auf den **v0.4.0**-Digest
  `@sha256:b0d6e33c…` gehoben (AC-QA-03, ADR-0004/ADR-0007).

## [0.3.0] - 2026-06-23

Dritte Welle: `welle-10b/b2b` — Driving/Driven-Port-Richtung + `LayerOf`-Angleichung an
`targetLayer`. Lastenheft/Spezifikation 0.5.0 → 0.6.0; **sieben** Regeln.

### Added

- **`AC-FA-RULE-008` (Lastenheft 0.5.0→0.6.0):** Driving/Driven-Port-Richtung —
  optionale Schicht-`direction` (`driving`/`driven`), **orthogonal** zur Rolle; neue
  Regel `port-direction-mismatch` (ein Adapter spricht nur Ports **seiner** Richtung),
  **kategorisch** (über `edges`/`allow` nicht aufhebbar). Ohne `direction` keine Prüfung
  (rückwärtskompatibel). `layers`-Objektform um `direction` (und das in 0.5.0 fehlende
  `app`) erweitert. ADR-0012; slice-012.

### Changed

- **`LayerOf` (ADR-0013):** die Schicht-Zuordnung einer Datei nimmt den spezifischsten/
  längsten **literalen** Glob-Präfix (`litPrefixLen`, Angleichung an `targetLayer`) statt
  des Erst-Treffers — Verhaltensänderung nur bei verschachtelten Schicht-Globs. slice-012.
- `--print-mk`/`a-check.mk` und der `aCheckImage`-Default sind auf den
  v0.2.0-Release **digest-gepinnt** (`ghcr.io/pt9912/a-check@sha256:4132a7af…`) —
  Pin-Hebung nach dem Release (AC-QA-03, ADR-0004/ADR-0007).

## [0.2.0] - 2026-06-22

Zweite Welle: das Regel-Modell dispatcht über Layer-**Rollen** statt -Namen und ist
auf vier Schichten (`domain`/`app`/`port`/`adapter`) ausgebaut; Ports dürfen
Domänentypen referenzieren. Lastenheft 0.1.0 → 0.5.0.

### Added

- **`AC-FA-RULE-006` (Lastenheft 0.2.0→0.4.0):** Schicht-**Rollen** — die
  Reinheits-Regeln dispatchen über eine Layer-Rolle (`domain`/`port`/`adapter`, aus
  `role:` oder Namens-Inferenz) statt über die Namen `core`/`ports`/`adapters`; fremd
  benannte Schichten sind voll prüfbar. `layers`-Eintrag als Glob-Liste **oder**
  `{globs, role}`; `lateral-adapter` cross-layer + kategorisch. ADR-0009; b1 (ADR-0010)
  macht `adapterSeg`/`targetLayer` vollständig namensunabhängig (längster,
  segment-bewusster Präfix). welle-10a/b1.
- **`AC-FA-RULE-007` (Lastenheft 0.4.0→0.5.0):** neue Schicht-Rolle `app`
  (Application-/Use-Case-Schicht) — darf `domain`+`port` referenzieren, aber keinen
  Adapter/Tech: neuer Befund `app-impurity`. Zugleich `domain` verschärft (Import auf
  `app`/`port`/`adapter`/Tech ⇒ `core-impurity`, kategorisch — „Domäne kennt keine
  Ports"); `role`-Schema um `app`. ADR-0011. **Breaking für geprüfte Repos:** eine
  `role: domain`-Schicht, die einen `port`/`app`-Layer importiert, wird jetzt rot
  (vorher per deklarierter Kante grün) — Migration: Port-/Use-Case-Nutzung in eine
  `role: app`-Schicht heben. a-checks Eigen-Dogfooding bleibt unverändert grün;
  Multi-Linsen-Review.
- Benutzerhandbuch 1.6: die Schicht-`role` dokumentiert (Objektform `{globs, role}`,
  Rollen, Namens-Inferenz, Vorrang, Vier-Schichten-`app`-Modell).

### Changed

- **`AC-FA-RULE-004` (Lastenheft 0.1.0→0.2.0):** Ports dürfen jetzt Domänen-/
  Kern-Typen referenzieren (`ports → core` per deklarierter Kante); `port-impurity`
  feuert nur noch bei Adapter-/Tech-Import, nicht mehr bei Kern-Import. Motiviert
  durch die Vier-Repo-Evidenz (b-cad/d-migrate-Ports referenzieren die Domäne);
  ADR-0008 (Accepted). a-check selbst auf eine echte `ports`-Schicht umgebaut
  (`internal/hexagon/{core,port}`, `internal/adapter/driven/*`), Dogfooding grün
  (AC-QA-02); Multi-Linsen-Review (`docs/reviews/2026-06-22-…`).
- `--print-mk`/`a-check.mk` und der `aCheckImage`-Default sind auf den
  v0.1.0-Release **digest-gepinnt**
  (`ghcr.io/pt9912/a-check@sha256:13459f44…`) statt auf die Tag-Form — Pin-Hebung
  nach dem ersten Release (AC-QA-03, ADR-0004/ADR-0007).

## [0.1.0] - 2026-06-21

Erstes Release: a-check als sprach-agnostischer Hexagonal-Architektur-Checker
(text-heuristisch, netzlos, distroless/static) inkl. Harness, Quality-Gates,
Durchsetzungsschicht und CI-/Release-Pipeline. Distribution als digest-gepinntes
GHCR-Image + `--print-mk`/`a-check.mk`.

### Added

- Bootstrap — Harness-Gerüst (AGENTS.md, harness/-Trias, Lastenheft 0.1.0)
  und das Doku-Gate `make doc-check` via Schwester-Tool d-check
  (digest-gepinnt, netzlos, read-only).
- slice-001 — Fundament-ADRs ADR-0001..0004 (Go als Implementierungssprache;
  text-heuristische Extraktion; Config-Modell `.a-check.yml`; Distribution
  inkl. `--print-mk`/`a-check.mk`); Status Accepted.
- slice-002 — Technik-Stratum `spec/spezifikation.md`
  (SPEC-CONF/EXTRACT/RULE/CLI/DET/DIST-001) und Sicht-Stratum
  `spec/architecture.md` (ARC-001..006); Spec-Strata in `harness/conventions.md`
  (MR-004) deklariert.
- slice-003 — Go-Implementierung (Hexagon: `internal/core`/`adapters`/`cli`,
  `cmd/a-check`): fünf Regeln AC-FA-RULE-001..005, text-heuristische
  Extraktion C++/Go/Rust/Kotlin (AC-FA-EXTRACT-001), strict-decode
  `.a-check.yml` (AC-FA-CONF-001), CLI/Exit-Codes (AC-FA-CLI-001),
  `--print-config`/`--print-mk` (AC-FA-DIST-001), Determinismus (AC-QA-01).
  Multi-Stage-Dockerfile (static/distroless, digest-gepinnte Bases, AC-QA-02/03).
- slice-003 — Quality-Gates `make lint`/`test`/`coverage-gate`/`arch-check`/`gates`
  (Dockerfile-Stages, Muster d-check/u-boot); `a-check.mk` via `--print-mk`.
  Lint-Profil golangci-lint v2 (ADR-0005); Coverage-Gate 90 % (ADR-0006, Ist 92,6 %).
  Dogfooding: a-check prüft seine eigene Hexagon-Architektur (AC-QA-02), 0 Befunde.
- slice-004 — Durchsetzungsschicht: Meta-Gates `make gate-consistency`
  (dokumentierte Targets ↔ Makefile + `.d-check.yml`-Module; Schutz gegen
  Harness-Lügen, schützt die doc-check-Beweisaussage AC-QA-02) und
  `make record-gates` (inhaltsbasierter Working-Tree-Hash-Nachweis) plus
  `.claude`-Stop-Hook als Handoff-Gate (fail-closed, loop-guarded, bootstrap-aware).
- slice-005 — Durchsetzungsschicht vollständig: PreToolUse-Command-Guard
  (`.claude/hooks/pretooluse-command-guard.sh`) lehnt Host-Toolchain/-Paketmanager
  (go/golangci-lint/pip/npm/cargo/apt/brew/…) vor der Ausführung fail-closed ab
  (Tool-Call-Gate, AGENTS §3.1); Selbsttest `make guard-selftest` (in `make gates`).
- slice-006 — CI: PR-/Push-Workflow `.github/workflows/ci.yml` (SHA-gepinnt,
  `permissions: {}`, Tags ausgenommen) fährt `make ci` (= `gates` + `make image-test`:
  AC-FA-DIST-001 `--print-mk`/`--print-config`/unbekanntes Flag + nativ==Container-
  Determinismus, AC-QA-02) und `make trace-check` (AC-/ADR-/MR-/slice-ID je Commit,
  AGENTS §5). Dockerfile-OCI-Labels (`org.opencontainers.image.*`) + `VERSION`-Build-Arg.
- slice-007 — Release-Pipeline `.github/workflows/release.yml` (auf `v*`-Tags,
  SHA-gepinnt): SemVer-Validate → `make ci VERSION=` → GHCR-Login → Tag (`:latest`
  nur stabil, ADR-0007) → OCI-Label-Verify → Push → GitHub-Release mit Digest-Pin.
  `:latest`-Tag-Politik in ADR-0007 (Accepted); `releasing.md` auf die reale
  Pipeline aktualisiert.
- slice-008 — lokaler `commit-msg`-Hook (`.githooks/commit-msg` + `make hooks`):
  ruft `trace-check` vor dem Commit (AGENTS §5), opt-in pro Klon; dieselbe
  Wahrheit wie CI/`make trace-check`.
