# MR-024 — Historische Kern-Drift-Befunde aus dem Archiv-Sweep sind deklariert

- **Status:** Accepted
- **Datum:** 2026-09-19
- **Geltungsbereich:** `make doc-immutable` (Item **2** der Freigabe-Checkliste in
  [`docs/user/releasing.md`](../../docs/user/releasing.md)) — genauer: Commit-Ranges, die den
  Zeitdokument-Sweep enthalten. Betroffen sind
  [`ADR-0017`](../../docs/plan/adr/0017-relative-resolution-modus.md),
  [`ADR-0018`](../../docs/plan/adr/0018-exclude-scan-scope.md) und
  [`ADR-0038`](../../docs/plan/adr/0038-dependabot-als-hebungskanal.md).
- **Ersetzt-Baseline-Regel:** — *(keine — der Sensor ist a-check-eigen, keine Regel der Baseline;
  der Eintrag deklariert eine Ausnahme im **eigenen** Bestand, wie
  [`MR-019`](../conventions.md#mr-019) und [`MR-020`](../conventions.md#mr-020) es für ihre
  Gegenstände tun.)*
- **Adaption:** Meldet `make doc-immutable` über eine Commit-Range Kern-Drift an einem `Accepted`
  ADR, deren Änderung **ausschließlich** ein Pfad-Nachzug in einem Verweis ist und **keine**
  Entscheidung berührt, gilt der Befund für Ranges, die diesen Nachzug enthalten, als
  **deklariert**. Er zählt nicht gegen Item 2 der Freigabe-Checkliste; die Freigabe nennt ihn mit
  seiner Ursache.
- **Begründung:** Der Anlass ist gemessen und **vorbestehend**. Der Zeitdokument-Sweep hat
  Zeitdokumente in Unterverzeichnisse bewegt und die Verweise **auf** sie repo-weit nachgezogen —
  darunter die Körper dreier `Accepted`-ADRs, die ihre Verifikations-Zeiger als **Adresse**
  zitierten. Entstanden sind drei Kern-Drift-Befunde, die **keine Entscheidung** verändern: In
  [`ADR-0038`](../../docs/plan/adr/0038-dependabot-als-hebungskanal.md) wandert ein Verweis auf ein
  Beobachtungs-Verzeichnis, in 0017/0018 der Pfad eines Slice-Dokuments.

  **Die Gegenmaßnahme ist bereits getroffen, und zwar als Doktrin des Repos:** Seit `slice-176` zitieren **einfrierende Artefakte Kennung statt Adresse** ([`AGENTS.md`](../../AGENTS.md) §5) —
  genau, damit kein Pfad in einem Artefakt altert, das niemand mehr anfassen darf.
  `slice-197` hat die drei ADRs auf diese Form ausgerichtet; **ab jetzt kann dort kein Pfad mehr driften.**

  **Was diese Deklaration nicht ist:** keine Lockerung des Gates. Der Sensor bleibt unverändert
  scharf für jede *künftige* Kern-Änderung — er kann nur nicht zwischen einer Entscheidung und
  einem Pfad-Nachzug unterscheiden, weil beides im selben Diff-Modus erscheint. Die Alternative
  wäre gewesen, die drei ADRs über `exempt-paths` auszunehmen; das hätte sie **dauerhaft** blind
  gestellt, auch für echte Änderungen. Diese Deklaration lässt sie im Blick und benennt nur die
  drei historischen Befunde.
- **Auflösungs-Trigger:** **das nächste Release.** Die Release-Range ist „seit dem letzten
  Release"; sobald dieses Release getaggt ist, enthält die folgende Range (`v0.20.0..HEAD`) den
  Sweep **nicht mehr**, und die drei Befunde können nicht mehr auftreten. Der Eintrag ist damit
  selbst-auflösend und wandert bei der nächsten Freigabe nach
  [`conventions/done/`](../conventions/done/).
