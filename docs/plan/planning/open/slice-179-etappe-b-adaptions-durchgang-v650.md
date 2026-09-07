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
  Vorgang. Etappe D ([slice-177](../open/slice-177-sensors-struktur-zwei-tabellen.md))
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

## 3. Umsetzung

*(offen)*

## 4. Definition of Done

- [ ] Alle **sieben** aktiven Adaptionen sind bewertet — je Eintrag: steht die
      ersetzte Regel in `v6.5.0` noch, und ist der Auflösungs-Trigger
      eingetreten? Mit Beleg je Zeile, nicht als Sammelurteil.
- [ ] Jeder Eintrag, dessen Trigger eingetreten ist, ist aufgelöst (Nachfolge-
      Eintrag oder Streichung mit Begründung); jeder andere trägt den Befund.
- [ ] Die `exempt-paths`-Frage ist entschieden: ADR, deklarierte Ausnahme oder
      begründete Nicht-Handlung — und die Entscheidung steht dort, wo sie beim
      nächsten Gate-Streit gelesen wird.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

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
  Eintrag hat keinen Aufhänger und fällt wieder durch* — Ausgang bei Closure;
  genau das ist bei `AGENTS.md` §4 vier Stände lang passiert
  ([`BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt`](../observations/BEO-HARNESS/baseline-regel-nie-erwogen-weil-bestand-sie-verletzt/observation.md), 2×).
- *Die `exempt-paths`-Entscheidung wird zur vierten Repo-Aussage-Korrektur und
  reißt damit eine Schwelle* — Ausgang bei Closure; Klasse
  [`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md) (2×).

## 8. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice;
Lerneintrag — Form: wird dort benannt.)_

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** entsteht mit dem Übergang nach
`in-progress/`.

**Vorgelagert — offene Beobachtungen sichten:** entsteht mit dem Übergang nach
`in-progress/`; zwei Einträge sind in §7 bereits als einschlägig benannt.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
