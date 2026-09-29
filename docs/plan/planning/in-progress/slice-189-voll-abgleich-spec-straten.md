# slice-189 — Voll-Abgleich: die Spec-Straten und die offene Architektur-Klausel

**Welle:** ohne Welle.

**Bezug:** Folge-Slice aus
[slice-187](../done/wellenlos/slice-187-voll-abgleich-erstdurchgang-rest.md) §1
(*es wäre ein anderer Vorgang*).
[`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze).

**Berührte Spec-Stellen:** [`spec/lastenheft.md`](../../../../spec/lastenheft.md)
und [`spec/spezifikation.md`](../../../../spec/spezifikation.md) sind der
Gegenstand — **welche** Kennungen berührt sind, benennt die Umsetzung.

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** wird bei Closure benannt.

---

## 1. Ziel und Abgrenzung

**Ziel:** Die zwei Spec-Straten sind gegen ihre vendorte Ziel-Form abgeglichen,
und **je Kandidat ist getrennt, was Doku-Pflege und was Change Request wäre**.
Letzteres wird **benannt, nicht ausgeführt** — das ist der Kern dieses Slice und
der Grund, warum er von slice-187 abgetrennt wurde.

| Paar | Kandidaten (Instrument: slice-187 §2) |
|---|---|
| [`spec/lastenheft.md`](../../../../spec/lastenheft.md) | 28 |
| [`spec/spezifikation.md`](../../../../spec/spezifikation.md) | 28 |
| [`spec/architecture.md`](../../../../spec/architecture.md) | 20 — **nur die eine offene Klausel**, siehe unten |

**Und ein liegengebliebener Punkt, der hierher gehört.** slice-185 hat
`spec/architecture.md` abgeglichen und **eine von drei** Klauseln der dortigen
Hard Rule als **verletzt** offen gelassen: *„keine Historie — `Letzte Änderung`
oben ist ein Frische-Marker, kein Protokoll"*. §8 der Datei trägt weiterhin eine
Versions-Tabelle. Adressiert war der Punkt an slice-186; dort und in slice-187
kommt er **nicht** vor, und slice-187 §3.3 hat ihn nicht bemerkt. Gefunden hat
es der unabhängige Review zu slice-187 (F-1).

Er gehört hierher und nirgends sonst: Ob a-check den Historie-Abschnitt streicht
und auf `git` verweist oder die Abweichung als `MR` deklariert, ist **eine
Entscheidung über ein Spec-Stratum** — genau der Gegenstand dieses Slice. Die
übrigen zwei Klauseln sind eingehalten und in slice-185 belegt; sie werden hier
nicht erneut geprüft.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Jede Änderung an einer bestehenden `AC-*`.** *Es wäre ein anderer Vorgang:*
  Das Lastenheft ist vertraglich abnahmebindend; eine Änderung ist ein **Change
  Request** mit Versions-Bump und Historie-Zeile
  ([`harness/conventions.md`](../../../../harness/conventions.md)
  §Anforderungs-Anlege-Prozess), keine Entscheidung eines Abgleich-Slice.
- **Die 19 grandfatherten `AC-*`.** *Bestand bleibt bewusst stehen:* Ihre Form
  ist seit slice-054 ausdrücklich ausgenommen, und die Grandfather-Liste wächst
  nicht mit ([`AGENTS.md`](../../../../AGENTS.md) §5).
- **Die zwei eingehaltenen Klauseln von `spec/architecture.md`.** *Bestand
  bleibt bewusst stehen:* Sie sind in slice-185 gemessen (je null Treffer) und
  als *bewusst abweichend* geschlossen — die Regel steht in
  [`AGENTS.md`](../../../../AGENTS.md) §3.4, sie zusätzlich in die Datei zu
  schreiben wäre die Doppelung, die dieses Repo gerade abbaut. **Die dritte
  Klausel ist ausdrücklich Gegenstand** (§1).
- **Die dreizehn bereits abgeglichenen Paare.** Erledigt (slice-185/186/187/188).

## 2. Ausgangsmessung

Neu gemessen gegen `v6.13.0` (slice-192-Lektion: der alte Stand zählt nicht):
Heading-Abgleich je Stratum (Vorlage vs. Artefakt, `grep -E "^#{1,3} "`), dazu
die Inhalte der Vorlagen-Sektionen gegen die Artefakt-Inhalte.

**Was die Zahl hier nicht sagt:** Ein Kandidat im Lastenheft kann drei sehr
verschiedene Dinge sein — eine fehlende **Form**-Regel (Doku-Pflege), eine
fehlende **Zusage** (Change Request), oder eine Stelle, an der a-check bewusst
schärfer ist als die Vorlage. Die Sortierung in diese drei ist die eigentliche
Arbeit; die 56 Kandidaten der Ausgangsmessung (slice-187, Stand `v6.5.0`) sind
neu sortiert.

## 3. Umsetzung

### 3.1 Lastenheft (Rang 1)

| Vorlagen-Element | Befund | Ausgang |
|---|---|---|
| `LH-FA-NN`/`LH-QA-NN`-Schema | a-check nutzt `AC-FA-*/AC-QA-*` | *bewusst schärfer*: das Schema ist in [`MR-000`](../../../../harness/conventions.md#mr-000) deklariert |
| **`LH-RB-NN`-Randbedingungen-Reihe** (Welle 144) | fehlt als eigene Reihe — a-checks Randbedingungen (Docker/make-only, 19 Grandfathers) sind teils formalisiert (die Hermetik als [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)), teils nur als Regeln getragen | **Change Request (benannt, nicht ausgeführt)**: eine `AC-RB-*`-Reihe würde bestehende Grenzen formalisieren — vertragliche Ergänzung, Maintainer-Entscheid |
| `Status:`-Feld (Draft \| In Review \| Accepted) | trägt „Draft" — richtig: die Abnahme ist ein Maintainer-Entscheid, der Historie-Stand 0.27.0 trug ausdrücklich „Status bleibt Draft". **Der im Lauf vollzogene Flip auf „Accepted" ist zurückgenommen** | **Change Request (benannt)**: der Status-Flip ist die Maintainer-Entscheid, nicht Doku-Pflege |
| §7 Historie | vorhanden ✓ | — |

### 3.2 Spezifikation (Rang 2)

| Vorlagen-Element | Befund | Ausgang |
|---|---|---|
| Kategorie-Gliederung (Algorithmen/Datenstrukturen/Defaults/Fehler-Codes/Metriken/Externe Verträge) | a-check gliedert nach Vertrags-Kennungen ([`SPEC-CONF-001`](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema), [`SPEC-EXTRACT-001`](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion), …) | *bewusst schärfer*: die Vertrags-Orientierung ist die deklarierte Form ([`MR-000`](../../../../harness/conventions.md#mr-000)/[`MR-027`](../../../../harness/conventions/MR-027-verfeinerungs-form-v6130.md)); die Inhalte liegen in den `SPEC-*`-Sektionen |
| Suffix-Form `LH-FA-01.a` (Verfeinerung) | a-check nutzt `SPEC-<BEREICH>-<NNN>` | *bewusst schärfer*: [`MR-027`](../../../../harness/conventions/MR-027-verfeinerungs-form-v6130.md) |
| Metriken- und Tracing-Felder | a-check hat keine Metriken ([`MR-014`](../../../../harness/conventions/MR-014-keine-agenten-telemetrie.md)) | *bewusst schärfer* |
| §7 Historie | vorhanden ✓ | — |

### 3.2a Repo → Vorlage (Review F-2: die Gegrichtung)

Die Prüfung des Erstlaufs lief nur Vorlage → Repo. Die Gegrichtung
(repo-seitige Abweichungen von der Ziel-Form) nachgemessen:

| Artefakt | Abweichung | Ausgang |
|---|---|---|
| `spezifikation.md` Kopf-Version 0.32.0 | die Ziel-Form sieht „keine Version" vor | **Change Request (benannt)**: die Version ist repo-weit in Commits/CHANGELOG referenziert — ihre Streichung ist ein Maintainer-Entscheid |
| `spezifikation.md` Kopf-Status „Draft" | die Ziel-Form sieht „keinen eigenen Status" vor | **Change Request (benannt)**: dieselbe Entscheidung |
| `spezifikation.md` Historie mit Version-Spalte | Ziel-Form: Historie ohne Version-Spalte | **Change Request (benannt)**: folgt der Kopf-Entscheidung |
| `architecture.md` Status „Draft" | die Ziel-Form trägt „Aktiv" | **Change Request (benannt)**: dieselbe Entscheidung |

### 3.3 Architektur (Rang 3) — die dritte Klausel

**Entscheidung (§1: kein dritter Weg):** der Historie-Abschnitt (§8,
Versions-Tabelle mit vier Einträgen) ist **gestrichen**; die Änderungsgeschichte
liegt in `git` (`git log -- spec/architecture.md`), der Frische-Marker
(„Letzte Änderung") steht am Kopf. Damit ist die dritte Klausel erfüllt — ihr
Anker ist nicht `AGENTS.md` §3.4, sondern `v6.13.0` ·
`regelwerk/modul-03-spec.md` §Ziel-Form: Architektur-Sicht (Review F-3: die
Klausel steht dort und im Template-Hard-Rule-Block, nicht in §3.4) — die zwei
übrigen Klauseln bleiben wie in slice-185 belegt.

## 4. Definition of Done

- [ ] Beide Spec-Straten sind abgeglichen; je Paar steht der Ausgang **mit
      seiner Prüf-Ebene** im Plan, und **je Kandidat** ist getrennt:
      **Doku-Pflege** (hier ausgeführt) · **Change Request** (benannt, mit dem,
      was er ändern würde) · **bewusst schärfer** (mit Begründung).
- [ ] Die offene dritte Architektur-Klausel trägt eine **Entscheidung**:
      Historie-Abschnitt gestrichen und auf `git` verwiesen, oder Abweichung als
      `MR` deklariert. Kein dritter Weg, kein Weiterreichen.
- [ ] Kein bestehendes `AC-*` ist inhaltlich geändert; `make doc-immutable`
      über die Commit-Range ist grün.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Zwei öffentliche Verträge sind berührt:
Rang 1 und Rang 2 der Source Precedence.

## 5. Trigger

**Start** (`open` → `in-progress`):
[slice-188](../done/wellenlos/slice-188-voll-abgleich-gate-und-skill.md) liegt in `done/` und das
WIP-Limit ist frei. **Reihenfolge ist Absicht:** Der Spec-Abgleich ist der
heikelste; er kommt zuletzt, wenn das Verfahren an vier Slices eingespielt ist.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Trägt eines der beiden Straten allein mehr
  als einen Durchgang, wird das andere abgetrennt.
- `in-progress` → `open` (blockiert): Verlangt ein Kandidat eine Änderung an
  einer bestehenden `AC-*`, ist das ein Change Request — der Slice geht zurück
  und die Entscheidung an den Maintainer.

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag.

## 7. Risiken und offene Punkte

- **Der Slice kann in einen Change Request kippen.** Eine Ziel-Form-Regel, die
  eine bestehende `AC-*` berührt, ist keine Doku-Pflege — und die Rückführung ist
  dann der richtige Ausgang, nicht ein stiller Nachtrag.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: kein Kandidat
  erforderte eine AC-Änderung; die CR-Kandidaten (die `AC-RB-*`-Reihe) sind
  benannt und an den Maintainer gereicht — der Slice hat sie nicht ausgeführt.
- **`doc-immutable` deckt die Spec-Straten nicht.** Das Modul `vcs` führt
  ausschließlich `docs/plan/adr/[0-9]*.md` (gemessen in slice-187 §3.5); eine
  unbeabsichtigte Änderung an einer `AC-*` meldet **kein** Lauf.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: alle Stratum-Änderungen
  dieses Slices sind absichtlich, im Review geprüft und im Diff sichtbar; die
  allgemeine Lücke (vcs deckt nur ADRs) bleibt benannt — slice-187 §3.5.
- **Der Historie-Abschnitt könnte an anderen Stellen zitiert sein.** Ein
  Streichen bricht dann Anker; die Antwort ist dieselbe wie unten.
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: nach dem Streichen
  meldet `make doc-check` 0 Befunde — keine Stelle zitiert §8.
- **Ein umbenannter `AC-*`-Anker bricht eine `Accepted`-ADR.** Der Fall ist im
  Repo belegt und die Antwort steht in
  [`harness/conventions.md`](../../../../harness/conventions.md)
  §Anforderungs-Anlege-Prozess (zwei Anker, alter Slug bleibt).
  — **Ausgang:** *entfallen*, gestrichen mit Begründung: es wurde kein
  `AC-*`-Anker umbenannt — die Konversion änderte Status-Wert und Historie,
  nicht die Anker.

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen — berührt ist `SPEC`, erstmals seit slice-185.)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso — **zwei** Quellen:
Register und der Review-Report des Vorgänger-Slice, siehe slice-187 §9.)*
