# slice-197 — Drei `Accepted`-ADRs tragen einen geänderten Kern

**Welle:** ohne Welle.

**Bezug:** Item **2** der Freigabe-Checkliste in
[`docs/user/releasing.md`](../../../../docs/user/releasing.md) — „`make doc-immutable` über die
Release-Range, Exit 0". Übernommen von [slice-196](../in-progress/slice-196-changelog-unreleased-nachtragen.md) §7.

**Berührte Spec-Stellen:** — (Gegenstand sind drei ADRs, kein Spec-Stratum).

**Verantwortlich:** — (bis zur Priorisierung).

**Autor:** Claude, im Auftrag des Maintainers. **Datum:** 2026-09-19.

**Lerneintrag — Form:** wird bei Closure benannt.

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

**Zu wählen beim Bau**, mit dem Maintainer; die Entscheidung gehört in eine ADR oder in den
Adaptions-Block, nicht in eine Commit-Message.

**Auszuführende Gates:** `make doc-immutable RANGE=v0.19.0..HEAD` (das Item selbst), `make gates`,
zum Abschluss `make verify`.

## 4. Definition of Done

- [ ] Die drei Befunde sind entschieden und die Entscheidung ist begründet abgelegt.
- [ ] `make doc-immutable` über die Release-Range ist **Exit 0** — oder die verbleibende Ausnahme
      ist deklariert und benannt.

- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

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
  stumpf für echte Kern-Änderungen. — **Ausgang:** <offen bis Closure>
- **Weg B erzeugt denselben Befund erneut.** Drei `Accepted`-ADRs zu ändern, um einen
  Immutabilitäts-Befund zu beheben, ist eine Änderung an `Accepted`-ADRs. — **Ausgang:** <offen bis
  Closure>
- **Der Befund ist älter als die Range.** Er entstand in `slice-139` und wäre bei jedem Release
  seitdem aufgetreten; ob es weitere gibt, sagt nur eine Messung über mehr als die Release-Range. —
  **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt sind `ADR` (die drei Entscheidungen samt Index) und,
je nach Weg, `HARNESS` (der Adaptions-Block) oder `GATE` (der Sensor). Alle GF.

**Vorgelagert — offene Beobachtungen sichten:** Register über `ADR`, `HARNESS`, `GATE` gelesen — zu
diesem Gegenstand kein Treffer; die Klasse *Chronik/Drift in ADRs* ist nicht geführt.

**Modus-Begründungsblock:** alle berührten Sub-Areas GF. Kein Block je Sub-Area nötig.
