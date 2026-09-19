# Review-Report: slice-197 — 2026-09-19

**Review-Art:** Plan-Review (Code-/Doku-Review) — geprüft gegen Slice-Plan, `AGENTS.md` §3/§5,
die Konventionen des Adaptions-Blocks und die vendorte Baseline `v6.6.0`; **nicht** gegen DoD/Spec
(Verifikation, andere Rolle).

**Gegenstand:** Die drei Commits `3f11ab0` (Weg B — die drei `Accepted`-ADRs zitieren Kennung statt
Adresse), `010960e` (`MR-024` + Register-Eintrag) und die Plan-Fortschreibung in
slice-197 §3/§4/§7/§8/§9. Arbeitsstand: `010960e`.

**Skill:** `.harness/skills/reviewer.md` @ Stand `010960e`; Bezugs-Modul `v6.6.0` ·
`regelwerk/modul-10-review-harness.md`
**Modell:** unbekannt (Subagent) · **Datum:** 2026-09-19

> **Zitier-Form** *(dieser Block bleibt stehen — er ist Norm, kein Ausfüll-Hinweis)*. Dieser Report
> friert ein; was er zitiert, bewegt sich weiter. Deshalb **Kennung, nicht Adresse** — `slice-NNN`
> statt seines Lifecycle-Pfads, `make <target>` statt eines Links auf die Sensor-Datei, eine
> Baseline-Stelle als **Tag + Pfad in Inline-Code** statt als Link
> (`v<X.Y.Z>` · `regelwerk/<datei>.md` §<Abschnitt>). Die `pfad`-Felder unten zitieren den
> **geprüften Gegenstand** und halten den Stand dieses Laufs fest; das ist davon nicht betroffen.

**Eingangs-Kontext** (die Verträge, gegen die geprüft wurde):

- slice-197 (Slice-Plan, Stand `010960e`) — §2 Messung, §3 Entscheidung, §7 Risiken, §8 Closure, §9 Modus
- `MR-024` (`harness/conventions/MR-024-historische-kern-drift-deklariert.md`) und `harness/conventions.md` §Aktive Adaptionen
- `AGENTS.md` §3.5, §3.6, §3.7, §4, §5; `harness/README.md` §Sensors; `docs/user/releasing.md` §Freigabe-Checkliste
- `.d-check.yml` (`vcs`-Block) und `d-check.mk` (`doc-immutable`)
- der neue Register-Eintrag `docs/plan/planning/observations/BEO-HARNESS/altbestand-braucht-nachzug/`
- `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Zwei Schritte vor der Modus-Begründung,
  `v6.6.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register,
  `v6.6.0` · `regelwerk/modul-07-carveouts.md` §Werkzeug-Wahl bei Diskrepanz,
  `v6.6.0` · `templates/harness/conventions/MR-NNN-titel.template.md`

---

## Findings

### F-1 — Die zwei Hälften der Deklaration schließen einander aus

- `kategorie`: HIGH
- `quelle`: Slice-Plan §7 Risiko 1 (Ausgang *entfallen*, „drei benannte Befunde … kein Freibrief");
  `v6.6.0` · `regelwerk/modul-06-roadmap.md` §Das Beobachtungs-Register (ein Ausgang trägt nur mit
  tragender Begründung)
- `pfad`: `harness/conventions/MR-024-historische-kern-drift-deklariert.md:15–19` (Adaption),
  `:37–41` (Auflösungs-Trigger) · `docs/plan/planning/in-progress/slice-197-adr-kern-drift.md:108–112`
- `befund`: Die Adaption ist als **Kriterium** formuliert („Meldet `make doc-immutable` … an einem
  `Accepted` ADR, deren Änderung ausschließlich ein Pfad-Nachzug ist … gilt der Befund … als
  deklariert") und damit ihrem Wortlaut nach unbeschränkt; die drei Befunde stehen nur im
  Geltungsbereich. Der Auflösungs-Trigger zieht den Eintrag dagegen mit **dem nächsten Release** ab
  — bei der weiten Lesart bliebe die Klasse danach ungedeckt, obwohl derselbe Slice sie als
  wiederkehrend in das Register einträgt; bei der engen Lesart ist §7 Risiko 1 („kein Freibrief")
  richtig, die Kriteriums-Formulierung aber zu weit. Beide Lesarten sind aus dem Text belegbar, und
  §7 entscheidet zwischen ihnen, ohne sie zu messen.
- `verifizierbar`: nein — kein Gate liest `MR`-Einträge; der Gegenstand ist der Wortlaut zweier
  Artefakte.
- `klasse`: „Deklaration: Ausnahme-Umfang und Auflösungs-Trigger schließen einander aus"

### F-2 — Die Wahl des Instruments ist nie gefällt, sie ist eingetreten

- `kategorie`: HIGH
- `quelle`: `AGENTS.md` §3.6 (Prüfregel-Senkung ist ein ADR), §5 *Diskrepanz-Trichter* (verweist auf
  `v6.6.0` · `regelwerk/modul-07-carveouts.md` §Werkzeug-Wahl bei Diskrepanz: BF-Markierung ·
  Carveout · ADR; Ablageorte `conventions.md` §Modus-Deklaration bzw. `docs/plan/carveouts/`),
  `v6.6.0` · `templates/harness/conventions/MR-NNN-titel.template.md` (*„Ein Eintrag, der keine
  benannte Regel ersetzt, ist ein **Fork**, keine Adaption."*)
- `pfad`: `docs/plan/planning/in-progress/slice-197-adr-kern-drift.md:52–70` (§3, Wege A/B/C),
  `harness/conventions/MR-024-historische-kern-drift-deklariert.md:11` (`Ersetzt-Baseline-Regel: —`),
  `docs/user/releasing.md:114` (Item 2)
- `befund`: Der Slice wägt drei Wege zur **Behebung** ab, aber das Instrument für die historische
  Hälfte kommt in keiner Zeile vor — `MR-024` tritt in §3 erst nach der Entscheidung „daneben"; in
  der einzigen früheren Nennung steht „Adaption/`MR`" als Bestandteil des **verworfenen** Wegs A.
  Der Eintrag erklärt anschließend einen roten Exit von Item 2 für unzulässig-nicht-blockierend und
  trägt als `Ersetzt-Baseline-Regel` ein `—`; messbar ist seinerseits nur die Gegenseite
  (`.d-check.yml` `vcs`-Block unverändert, `make doc-immutable` weiter Exit 2). Ob das eine
  Prüfregel-Senkung im Sinne §3.6 und ob der Adaptions-Block der richtige Ort ist, entscheidet damit
  ein Urteil, für das §3.6 selbst ein ADR vorsieht — die Gegenposition (Präzedenz [`MR-019`],
  Sensor-Konfiguration unangetastet) steht im Report, nicht in den Artefakten.
- `verifizierbar`: nein für den Kern (Zuordnungs-Urteil); ja für die Mess-Seite —
  `make doc-immutable RANGE=v0.19.0..HEAD` → Exit 2, `.d-check.yml` `vcs`-Block ohne Änderung.

### F-3 — Die Ausnahme steht nicht dort, wo das Verdikt fällt

- `kategorie`: MEDIUM
- `quelle`: `AGENTS.md` §4 (Gate-Index einmal in `harness/README.md` §Sensors, dort die Bindung);
  `v6.6.0` · `regelwerk/modul-07-carveouts.md` §Ziel-Form: Carveout (eine Gate-Ausnahme muss im
  Gate-Output und in der Bindung sichtbar sein)
- `pfad`: `docs/user/releasing.md:114`, `harness/README.md:102`, `harness/conventions/MR-024-historische-kern-drift-deklariert.md:8–10`
- `befund`: Item 2 lautet unverändert „`make doc-immutable` über die Release-Range, **Exit 0**", und
  die Bindung-Zelle des Gate-Index nennt `AGENTS.md` §3.5 + `slice-029` — beide Stellen, an denen
  das Verdikt entsteht, sagen von `MR-024` nichts. `MR-024` nennt Item 2 seinerseits als
  Geltungsbereich; der Leser der Checkliste liest den Widerspruch (rot vs. „zählt nicht") ohne die
  Deklaration.
- `verifizierbar`: ja — Textsuche über `releasing.md` / `harness/README.md` nach `MR-024`; kein Gate
  führt sie (`make doc-mentions` prüft nur die Sensor-Dateien).

### F-4 — Die Ursachenzuordnung nennt einen Sweep, gemessen sind es zwei

- `kategorie`: MEDIUM
- `quelle`: Slice-Plan §2 („Ausgangsmessung" / Geltungsbereich einer Messung)
- `pfad`: `docs/plan/planning/in-progress/slice-197-adr-kern-drift.md:24–29`,
  `harness/conventions/MR-024-historische-kern-drift-deklariert.md:20–22`,
  `docs/plan/planning/observations/BEO-HARNESS/altbestand-braucht-nachzug/evidence/slice-197.md`
- `befund`: §Ausgangslage nennt als Ursache die **Register-Migration** (`slice-139`) und stellt
  `ADR-0038` als deren Beispiel daneben; gemessen (`git log v0.19.0..HEAD -- <drei Dateien>`) hat
  `f516da2` (`slice-148`, Archivierung des Altbestands) `ADR-0017`/`ADR-0018` geändert und
  `e7b6f16` (`slice-139`) allein `ADR-0038`. `MR-024` und der Beleg führen dagegen „den
  Zeitdokument-Sweep" als **einen** Vorgang, was `ADR-0038` nicht deckt; die Commit-Message von
  `3f11ab0` nennt wiederum `slice-148` allein. Vier Artefakte, drei Fassungen derselben Ursache.
- `verifizierbar`: ja — `git log --oneline v0.19.0..HEAD -- docs/plan/adr/<datei>`;
  `git show --stat` je Commit.

### F-5 — Der Sichtungs-Schritt trifft nicht den Eintrag, der das gewählte Werkzeug betrifft

- `kategorie`: MEDIUM
- `quelle`: `v6.6.0` · `regelwerk/modul-05-planning-harness.md` §Zwei Schritte vor der
  Modus-Begründung (Sichtungs-Schritt ist unbedingt); §Das Beobachtungs-Register (Leser unter der
  Schwelle ist allein die Slice-Planung)
- `pfad`: `docs/plan/planning/in-progress/slice-197-adr-kern-drift.md:144–147` (§8),
  `:168–169` (§9) · `docs/plan/planning/observations/BEO-HARNESS/mr-aufloesungs-trigger-ohne-waechter/observation.md`
- `befund`: §9 meldet für `HARNESS` „zu diesem Gegenstand kein Treffer", §8 „zu dieser Klasse führt
  das Register **keinen** Eintrag" — im Register liegt bei **1×** der Eintrag *Der Auflösungs-Trigger
  eines `MR`-Eintrags hat keinen Wächter*, dessen Gegenstand genau das hier gewählte Instrument ist
  (`MR-024`s Begründung steht und fällt mit seinem Trigger, und niemand liest ihn). Der Vorgang
  dieses Slice ist ein zweites Auftreten dieser Klasse und bleibt durch die beiden Sätze ungezählt.
- `verifizierbar`: nein — die Zuordnung *derselbe Gegenstand?* ist ein Urteil; der Bestand
  (Eintrag, Zähler 1×, Beleg `slice-170`) ist belegbar.

### F-6 — Ein DoD-Haken attestiert den Review, bevor er stattgefunden hat

- `kategorie`: LOW
- `quelle`: Register-Eintrag `BEO-GATE/attestierung-vor-dem-vorgang` (1×) — dessen Beispielzeile
  ist genau diese Form
- `pfad`: `docs/plan/planning/in-progress/slice-197-adr-kern-drift.md:83`
- `befund`: Der Haken `- [x] Unabhängiger Review durchgeführt (Report unter docs/reviews/)` wurde in
  `010960e` gesetzt, während unter `docs/reviews/` kein Report zu `slice-197` lag (er entsteht mit
  diesem Lauf). Die Klasse ist registriert und bei 1×; `slice-197` wäre ihr zweiter Beleg.
- `verifizierbar`: ja — `git show 010960e -- <Plan>` zeigt `- [ ]` → `- [x]`; `ls docs/reviews/`.
- `klasse`: „Attestierung vor dem Vorgang"

### F-7 — Die Begründung trägt die verworfene Alternative im Konjunktiv

- `kategorie`: LOW
- `quelle`: `AGENTS.md` §3.7 (*Falsch:* Konjunktiv über die verworfene Alternative; die Abwägung
  gehört in die ADR)
- `pfad`: `harness/conventions/MR-024-historische-kern-drift-deklariert.md:22–26`
- `befund`: „Die Alternative wäre gewesen, die drei ADRs über `exempt-paths` auszunehmen; das hätte
  sie dauerhaft blind gestellt …" beschreibt eine nicht gewählte Fassung und ihre Folgen, statt den
  geltenden Zustand zu benennen. **Nicht als HIGH gewertet:** die Regel ist auf Kommentare,
  Konfiguration, Skripte und Zustandsfelder gemünzt, und ein `Begründung:`-Feld ist argumentative
  Prosa — die Aussage ist aber dieselbe, die §3.7 als *Falsch* führt, und sie stützt F-2.
- `verifizierbar`: nein — Textgegenstand.

### F-8 — Kennungen im lebenden Dokument ohne Link, im selben Block mit Link

- `kategorie`: LOW
- `quelle`: Haus-Stil der lebenden Harness-Dokumente (`harness/conventions.md` §Baseline verlinkt
  `slice-172`/`slice-173`), `AGENTS.md` §5 (Kennungs-Linkpflicht für Cross-Doc-Kennungen)
- `pfad`: `harness/conventions/MR-024-historische-kern-drift-deklariert.md:21–22, :33`
- `befund`: Der Eintrag nennt `slice-176` und `slice-197` als nackten Inline-Code, während dieselbe
  Datei-Familie Slice-Kennungen sonst verlinkt (und `make doc-check` die Linkpflicht für Kennungen
  führt, hier aber nicht greift). `slice-197` ist zudem zum Prüfzeitpunkt nicht adressierbar — es
  liegt in `in-progress/`, was für die Kennungs-Form spricht und hier unentschieden bleibt.
- `verifizierbar`: ja — `make doc-check` (Exit 0, greift an dieser Stelle nicht).

### INFO-1 — Adressen bleiben in vier weiteren `Accepted`-ADRs stehen, außerhalb des Sensors

- `kategorie`: INFO
- `quelle`: `.d-check.yml` `vcs` (`exclude-sections: [Geschichte]`), `AGENTS.md` §5
- `pfad`: `docs/plan/adr/0014-resolution-roots.md`, `0015-regex-tech-muster.md`,
  `0016-resolution-sprach-parametrisch.md`, `0019-adapterseg-root-subeinheit.md` (Zeilen der
  `Geschichte`-Tabelle)
- `befund`: Dieselben Sweeps haben in vier weiteren `Accepted`-ADRs Adressen nachgezogen — nur in
  den `Geschichte`-Anhängen, die der Sensor als Kern ausschließt; sie erzeugen deshalb keinen
  Befund, tragen aber dieselbe Zitier-Form wie die drei behobenen. Für die zuständige Rolle
  vermerkt, keine Aktion dieses Slice.

## Negativbefunde

- geprüft, ohne Befund: **die drei geänderten ADRs sind adressfrei** — `grep` über `planning/`,
  `observations/`, `BEO-` in `0017`/`0018`/`0038` findet nur noch kennungsförmige Nennungen; die
  verbliebenen Links zeigen auf `docs/plan/adr/` (bewegt sich nicht).
- geprüft, ohne Befund: **kein neuer Befund aus der Änderung selbst** — `make doc-check` Exit 0 über
  590 Dateien (Link-, Anker-, Kennungs- und Matrix-Prüfung; Prosa sieht das Instrument nicht).
- geprüft, ohne Befund: **die Kern-Behauptung „Weg B macht Item 2 nicht grün"** — selbst gemessen in
  drei Ranges (Exit 2, je 3 Befunde), plus die Gegenprobe `3f11ab0..HEAD` → Exit 0 als Beleg der
  Endpunkt-Semantik; die Selbst-Auflösung des Triggers ist damit gestützt (der fehlende Wächter ist
  F-5, nicht die Unwahrheit des Satzes).
- geprüft, ohne Befund: **die Index-Zählung** `harness/conventions.md:163` — zwei verschieden
  gebaute Zähler stimmen überein: 9 Zeilen `^| [MR-` in §Aktive Adaptionen und 9 `.md`-Dateien in
  `harness/conventions/`; davon **fünf** mit Zeiger, **vier** mit `—`.
- geprüft, ohne Befund: **`MR-024` trägt die Pflichtfelder** (Datum, Geltungsbereich,
  Ersetzt-Baseline-Regel, Adaption, Begründung, Auflösungs-Trigger), der Anker `mr-024` löst aus
  beiden Richtungen auf; `make doc-structure` Exit 0.
- geprüft, ohne Befund: **§3.7-Zustandsseite aller neuen Texte** — „Die Gegenmaßnahme ist bereits
  getroffen", „ab jetzt kann dort kein Pfad mehr driften", `state.md` „**Stand:** offen (1×)" nennen
  Zustand und Beleg; keine Chronik, kein `Status:`-Feld als zweite Zustandsquelle.
- geprüft, ohne Befund: **der neue Register-Eintrag** — Kürzel `HARNESS` ist in
  `harness/conventions.md` §Modus-Deklaration deklariert (nicht erfunden), Slug Kebab-Case, drei
  Dateien mit drei Lebensdauern; `state.md` trägt den abgeleiteten Zähler, `evidence/slice-197.md`
  benennt einen Vorgang; die Abgrenzung zu `BEO-HARNESS/massen-ersetzung-trifft-die-historische-aussage`
  ist gezogen, eine Aufzählung in `observations/README.md` verbietet das Register selbst.
- geprüft, ohne Befund: **`make verify-risiko-ausgaenge` Exit 0** (11 Slices, jedes notierte Risiko
  mit Ausgang) und **`make verify` Exit 0** — das betrifft die Verifikations-Schicht und ist hier
  nur als Randbeleg gefahren.
- geprüft, ohne Befund: `MR-024`s Satz „ab jetzt kann dort kein Pfad mehr driften" ist an den drei
  ADRs belegbar (keine Adresse mehr im Kern); die Aussage ist nicht zu weit **für diese drei**.

## Summary

| Kategorie | Anzahl |
|---|---|
| HIGH | 2 |
| MEDIUM | 3 |
| LOW | 3 |
| INFO | 1 |

**Finding-Klassen dieses Laufs:** „Deklaration: Ausnahme-Umfang und Auflösungs-Trigger schließen
einander aus" · „Instrument der Ausnahme nie gefällt (Adaptions-Eintrag statt Diskrepanz-Trichter)"
· „Ausnahme steht nicht dort, wo das Verdikt fällt" · „Ursachenzuordnung ohne Messung" ·
„Sichtungs-Schritt trifft den Eintrag zum gewählten Werkzeug nicht" · „Attestierung vor dem Vorgang"
· „Konjunktiv über die verworfene Alternative" · „Kennung im lebenden Dokument ohne Link"

## Gefahrene Gates (Exit-Code, Geltungsbereich)

| Lauf | Exit | Geltungsbereich |
|---|---|---|
| `make doc-immutable RANGE=v0.19.0..HEAD` | **2** | `docs/plan/adr/[0-9]*.md`, nur der `Accepted`-Kern ohne `Geschichte`, Commit-Range, nicht Arbeitsbaum: 3 Befunde `core-drift-vcs` (`0017`, `0018`, `0038`) — bestätigt §2 |
| `make doc-immutable RANGE=v0.19.0..89fc7dc` | **2** | dieselbe Menge: 3 Befunde → „vorbestehend" bestätigt |
| `make doc-immutable RANGE=89fc7dc..HEAD` | **2** | dieselbe Menge: 3 Befunde — **ohne** Sweep in der Range, d. h. der rote Exit einer reinen Neu-Commit-Range kommt aus `3f11ab0` selbst |
| `make doc-immutable RANGE=f516da2..HEAD` | **2** | dieselbe Menge: 3 Befunde |
| `make doc-immutable RANGE=3f11ab0..HEAD` | 0 | dieselbe Menge: Endpunkt-Semantik des Moduls belegt |
| `make doc-immutable RANGE=v0.18.0..v0.19.0` | 0 | dieselbe Menge: Sweep liegt nach `v0.19.0` |
| `make doc-check` | 0 | 590 Dateien, Link-/Anker-/Kennungs-/Matrix-Modul über das ganze Repo (ohne `.harness/baseline/**`) |
| `make doc-structure` | 0 | Struktur-Invarianten der Doku (Kopffelder, DoD, Closure, Risiko-Form) |
| `make verify-risiko-ausgaenge` | 0 | 11 Slices (11 × `done/` + abschlussbereite `in-progress/`) |
| `make verify` | 0 | Verifikations-Schicht (DoD/Closure/Requirements — 21 Anforderungen, 0 Waisen) |

**Nicht gesehen:** `make gates`/`make ci` als Ganzes wurden **nicht** gefahren (der Plan meldet sie
grün; das ist Verifier-Territorium). Der Arbeitsbaum wurde vor Erstellung dieses Reports gemessen;
der Report selbst verändert `docs/reviews/` und ist damit noch nicht in den Läufen oben enthalten.
`make doc-reviews` greift erst für einen Slice in `done/` — sein Exit 0 ist für diesen Report
**kein** Beleg; gefahren wurde er nur, um zu sehen, dass der neue Report die anderen Module nicht
rot macht.

## Verdikt

**Merge-blockierend:** ja — F-1 und F-2 blockieren (HIGH); F-3, F-4 und F-5 sind MEDIUM und in
diesem Repo vor Acceptance zu klären. Weg B selbst (F-Abnahme der drei ADR-Änderungen) trägt: die
Änderung ist gemessen adressfrei, entscheidungsneutral und ohne neue Befunde; der Streit liegt
allein beim Instrument und beim Umfang der Deklaration.

**Übergabe:** Findings gehen an den Implementer (Rückkante Review → Plan bei Plan-Defekt); die
**Finding-Klassen** gehen zusätzlich in die Slice-Closure §7 und von dort in den Zähler. Dieser
Report ist ein Lauf-Beleg — dieser Diff, dieser Skill, dieses Modell, dieses Verdikt. Er ersetzt
keine Verifikation (Modul 11).
