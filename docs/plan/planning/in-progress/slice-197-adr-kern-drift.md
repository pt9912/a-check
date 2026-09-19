# slice-197 — Drei `Accepted`-ADRs tragen einen geänderten Kern

**Welle:** ohne Welle.

**Bezug:** Item **2** der Freigabe-Checkliste in
[`docs/user/releasing.md`](../../../../docs/user/releasing.md) — „`make doc-immutable` über die
Release-Range, Exit 0". Übernommen von [slice-196](../done/wellenlos/slice-196-changelog-unreleased-nachtragen.md) §7.

**Berührte Spec-Stellen:** — (Gegenstand sind drei ADRs, kein Spec-Stratum).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude, im Auftrag des Maintainers. **Datum:** 2026-09-19.

**Lerneintrag — Form:** benannte Spec-Lücke.

---

## 1. Ziel und Abgrenzung

**Ziel:** Ein Release lässt sich freigeben, ohne dass Item 2 der Checkliste rot ist — die drei
`Accepted`-ADRs, deren Kern sich über die Release-Range geändert hat, haben eine Entscheidung.

**Ausgangslage (gemessen):** `make doc-immutable RANGE=v0.19.0..HEAD` meldet **Exit 2** für
[`ADR-0017`](../../adr/0017-relative-resolution-modus.md), [`ADR-0018`](../../adr/0018-exclude-scan-scope.md) und [`ADR-0038`](../../adr/0038-dependabot-als-hebungskanal.md) (`core-drift-vcs`: „Core einer immutablen Datei hat sich über
die Commit-Range geändert"). Der Befund ist **vorbestehend** — er tritt ohne die Commits dieser
Sitzung ebenso auf (`RANGE=v0.19.0..89fc7dc`, Exit 2). Ursache ist die Register-Migration
(`slice-139`), die Pfade in ADR-Körpern nachgezogen hat; bei [`ADR-0038`](../../adr/0038-dependabot-als-hebungskanal.md) ist die geänderte Zeile ein
Verweis auf ein Beobachtungs-Verzeichnis, das vorher flach lag.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Das Werkzeug ändern** (`vcs`-Modul, `head-allow`). *Schicht-Abgrenzung:* Der Sensor misst
  richtig — ein `Accepted`-ADR-Kern, der sich ändert, *ist* die Klasse, gegen die §3.5 steht. Ihn
  weicher zu stellen, wäre eine Gate-Lockerung und damit ADR-pflichtig ([`AGENTS.md`](../../../../AGENTS.md) §3.6).
- **Die vierzehn weiteren Slices des Bereichs.** *Es wäre ein anderer Vorgang:* hier stehen genau
  die drei Befunde, die die Release-Range meldet.

## 2. Ausgangsmessung

| Messung | Bau | Ergebnis |
|---|---|---|
| Befunde über die Release-Range | `make doc-immutable RANGE=v0.19.0..HEAD` | **3** ([`ADR-0017`](../../adr/0017-relative-resolution-modus.md), [`ADR-0018`](../../adr/0018-exclude-scan-scope.md), [`ADR-0038`](../../adr/0038-dependabot-als-hebungskanal.md)), Exit 2 |
| dieselbe Messung ohne diese Sitzung | `RANGE=v0.19.0..89fc7dc` | **3**, Exit 2 — **vorbestehend** |
| Art der Änderung | `git diff v0.19.0..HEAD -- docs/plan/adr/` | Pfad-Nachzüge in Verweisen, keine Entscheidungsänderung |

**Geltungsbereich:** die Commit-Range, nicht der Arbeitsbaum — `doc-immutable` ist ein
`vcs`-Modul und liest `git`.

## 3. Umsetzung

Die Entscheidung steht **vor** der Umsetzung und ist der eigentliche Inhalt dieses Slice:

| Weg | Was er bedeutet | Preis |
|---|---|---|
| **A — die Nachzüge als gewollt ausweisen** (Adaption/`MR` oder ein deklarierter Ausnahme-Satz) | Der Kern hat sich geändert, und das Repo sagt, dass ein **Pfad**-Nachzug keine Entscheidung ändert | Eine Ausnahme, die für *jeden* künftigen Fall dieser Art gilt — und die den Sensor stumpf macht, wenn sie zu breit ist |
| **B — die Verweise aus dem ADR-Körper herausziehen** (der ADR nennt den Gegenstand, nicht den Pfad) | Der Kern trägt keine Pfade mehr, die wandern können | Drei `Accepted`-ADRs würden erneut geändert — dieselbe Klasse, diesmal **gewollt** und dokumentiert |
| **C — Folge-ADR mit `Supersedes`** | Der Formalfehler wird nach §3.5 geheilt | Drei ADRs für drei Verweise — unverhältnismäßig, und sie sagen inhaltlich nichts Neues |

**Gewählt: Weg B** (Maintainer, 2026-09-19) — die Verweise wandern aus dem ADR-Körper heraus.
Umgesetzt: die drei ADRs zitieren ihre Zeitdokumente jetzt als **Kennung** statt als Adresse, genau
die Form, die [`AGENTS.md`](../../../../AGENTS.md) §5 seit `slice-176` für einfrierende Artefakte
verlangt.

**Und eine Messung, die den Plan in einem Punkt widerlegt:** B macht Item 2 **nicht** grün. Der
Sensor prüft die **Commit-Range**, und der verursachende Nachzug liegt **in** ihr — jede heutige
Änderung ist selbst wieder eine Kern-Änderung derselben Range. `make doc-immutable
RANGE=v0.19.0..HEAD` meldet nach B weiterhin Exit 2. B beseitigt die Drift-**Quelle** für die
Zukunft, nicht die Vergangenheit; für den historischen Befund tritt
[`MR-024`](../../../../harness/conventions.md#mr-024) daneben.

**Auszuführende Gates:** `make doc-immutable RANGE=v0.19.0..HEAD` (das Item selbst), `make gates`,
zum Abschluss `make verify`.

## 4. Definition of Done

- [x] Die drei Befunde sind entschieden und die Entscheidung ist begründet abgelegt — Weg **B**,
      ausgeführt; die historische Hälfte deklariert [`MR-024`](../../../../harness/conventions.md#mr-024).
- [x] Die verbleibende Ausnahme ist **deklariert und benannt**: drei historische Kern-Drift-Befunde,
      [`MR-024`](../../../../harness/conventions.md#mr-024) — mit einem Auflösungs-Trigger, der sich
      mit dem nächsten Release selbst einlöst.

- [x] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [x] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün.

## 5. Trigger

**Start** (`open` → `in-progress`): das WIP-Limit ist frei.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Erweist sich Weg **B** als richtig, sind die drei ADRs drei
  Vorgänge mit je eigener Prüfung — dann wird geteilt.
- `in-progress` → `open` (blockiert): Fällt die Wahl auf **C**, ist das eine Architect-Entscheidung
  und keine Implementierung.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit Lerneintrag — und
`make doc-immutable` über die Release-Range als eigener Beleg.

## 7. Risiken und offene Punkte

- **Die Ausnahme wird zu breit.** Ein deklarierter Freibrief für „Pfad-Nachzüge" macht den Sensor
  stumpf für echte Kern-Änderungen. — **Ausgang:** *entfallen*, gestrichen mit Begründung: Die
  Deklaration nennt **drei benannte Befunde** und einen Grund, sie ist kein Freibrief; die drei
  ADRs bleiben über `exempt-paths` **im** Geltungsbereich des Sensors, statt ausgenommen zu werden
  — der Weg, der sie dauerhaft blind gestellt hätte.
- **Weg B erzeugt denselben Befund erneut.** Drei `Accepted`-ADRs zu ändern, um einen
  Immutabilitäts-Befund zu beheben, ist eine Änderung an `Accepted`-ADRs. — **Ausgang:**
  *entfallen*, gestrichen mit Begründung: Der Befund nach der Änderung ist **derselbe** historische
  — kein neuer. Die Änderung ist einmalig geschehen, berührt keine Entscheidung und entfernt die
  Adresse, die ihn erzeugen konnte. Die Messung danach (`Exit 2`, historische Hälfte) steht in §3
  und in der Closure.
- **Der Befund ist älter als die Range.** Er entstand im Archiv-Sweep und wäre bei jedem Release
  seitdem aufgetreten; ob es weitere gibt, sagt nur eine Messung über mehr als die Release-Range. —
  **Ausgang:** *weiter offen* → **Beobachtungs-Register**:
  [`BEO-HARNESS/altbestand-braucht-nachzug`](../observations/BEO-HARNESS/altbestand-braucht-nachzug/observation.md)
  — die Klasse dahinter ist „eine neue Regel gilt für den Bestand, den sie beim Entstehen sieht";
  ob es **weitere** Befunde gibt, sagt nur eine Messung über mehr als die Release-Range, und die
  ist ein eigener Vorgang.

## 8. Closure-Notiz

**Lerneintrag — Form: benannte Spec-Lücke.** *Die Hausregel gegen genau diese Klasse existierte
längst — sie galt nur für Artefakte, die nach ihr entstanden sind.* Seit `slice-176` zitieren
einfrierende Artefakte **Kennung statt Adresse**, und genau deshalb altert kein Pfad in einem
`Accepted` ADR mehr. Die drei betroffenen ADRs stammen von **davor**; ihre Adressen hat der
Archiv-Sweep nachgezogen, und der Immutabilitäts-Sensor meldet das seither bei **jedem** Release.
**Weil** eine Regel nur für den Bestand gilt, den sie beim Entstehen sieht, braucht jede neue
Zitier-Regel einen **Nachzug des Altbestands** — hier nachgeholt.

**Und eine Lehre über die eigene Planung.** §3 dieses Plans bot drei Wege an und stellte B als den
dar, der den Sensor scharf lässt. Die Messung nach der Umsetzung zeigte: B lässt ihn für die
**Zukunft** scharf, aber Item 2 wird davon nicht grün — der Sensor prüft die **Commit-Range**, und
die Vergangenheit bleibt in ihr. Der Plan hatte den Sensor für ein Zustands-Werkzeug gehalten; er
ist ein **Verlaufs**-Werkzeug. Der Fehler ist mit einer Messung gefunden worden, nicht mit einem
Argument.

**Steering-Loop-Eintrag:** gezählt, nicht verkörpert. Zu dieser Klasse führt das Register **keinen**
Eintrag (siehe §9); ob sie einen braucht, entscheidet das nächste Auftreten — der
[`MR-024`](../../../../harness/conventions.md#mr-024)-Trigger löst sich mit dem nächsten Release
selbst ein.

**Beobachtungs-Register ([`../observations/`](../observations/README.md)):** ein Verzeichnis
**neu angelegt** — [`BEO-HARNESS/altbestand-braucht-nachzug`](../observations/BEO-HARNESS/altbestand-braucht-nachzug/observation.md),
Beleg `evidence/slice-197.md`, Zähler **1×**. Der Vorgang selbst ist ein Nachzug; die **Klasse**
dahinter ist neu und wird mit ihm benannt. `BEO-HARNESS/chronik-in-gelesenen-dateien`
bleibt bei 3× mit Ausgang *geplant*; [`MR-024`](../../../../harness/conventions.md#mr-024) ist **kein** Beleg dafür (der Eintrag ist ein
Adaptions-Eintrag, kein Chronik-Fall).

**Folge-Slices:** keine.

**Risiken aus §7:** alle drei *entfallen* oder sind entschieden — siehe dort.

**Drei Paarungen:** Anker — kein `liegt in`-Feld gesetzt, weil nichts verkörpert wurde · Folge-Slice
— keiner genannt · Register — keine Beobachtung genannt, keine zu decken.

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt sind `ADR` (die drei Entscheidungen samt Index) und,
je nach Weg, `HARNESS` (der Adaptions-Block) oder `GATE` (der Sensor). Alle GF.

**Vorgelagert — offene Beobachtungen sichten:** Register über `ADR`, `HARNESS`, `GATE` gelesen — zu
diesem Gegenstand kein Treffer; die Klasse *Chronik/Drift in ADRs* ist nicht geführt.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF. Kein Block je Sub-Area nötig.
