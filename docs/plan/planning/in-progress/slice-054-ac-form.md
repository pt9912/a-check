# slice-054 — Etappe D (3/3): AC-Form für neue Anforderungen

**Status:** in-progress — letzter Schnitt der **Etappe D (Form)** aus
[slice-048 §5](../done/slice-048-modul-delta-lesen.md).
**Deckt:** Fund **B-15** (AC-Form ohne Happy/Boundary/Negative-Pfade).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

Die Baseline verlangt je funktionalem Akzeptanzkriterium **drei Pfade im Given/When/Then-Stil** —
Happy · Boundary · Negative — plus einen Out-of-Scope-Block. Gemessen in slice-048 über alle 19
`AC-*`: **keines** trägt die Drei-Pfad-Gliederung; die Form ist Prosa unter „**Beschreibung:**"
mit eingebetteten Rand- und Negativsätzen. Immerhin **16 von 19** tragen Out-of-Scope.

Der Negativpfad ist der, den die Baseline als teuerste Auslassung benennt: *„ein Satz »das System
darf nicht …« spart später drei Reviews."* Genau dieser Satz fehlt in a-checks ACs nicht
inhaltlich — er steckt in der Prosa —, aber er ist **nicht als Pfad auffindbar**. Wer prüfen will,
ob eine Anforderung ihren Fehlerfall definiert, muss den Fließtext lesen statt eine Zeile.

## 2. Entscheidung: Form für Neues, Bestand grandfathered

Die 19 bestehenden `AC-*` werden **nicht** umgeschrieben. Begründung:

- Sie sind **vertraglich abnahmebindend**. Eine Umformatierung wäre eine Änderung am
  Vertrags-Stratum ohne fachlichen Anlass — mit Lastenheft-Bump, Spec-Nachzug und Re-Review von
  19 Anforderungen, die inhaltlich unverändert bleiben.
- Der Inhalt ist **da**; es fehlt die Gliederung. Ein Umbau träfe die Form, nicht die Substanz —
  und riskierte, beim Zerlegen von Prosa in drei Pfade eine Aussage zu verschieben.
- Die Baseline selbst grandfathert in derselben Lage: *„vor Einführung dieser Konvention
  entstandene Grenzfälle werden grandfathered, nicht durch eine superseding ADR nachgezogen. Der
  Gate prüft nur ab Einführung neu."*

Also: die Form gilt für **neue** `AC-*` ab slice-054, und der Sensor führt die 19 bestehenden
namentlich als grandfathered. Namentlich, nicht per Datums-Heuristik — eine Liste, die beim
Hinzufügen einer neuen `AC-*` **nicht** mitwächst, ist selbst der Beleg dafür, dass die Regel ab
jetzt greift.

## 3. Betroffene Module

- [`AGENTS.md`](../../../../AGENTS.md) §5 — Form-Regel für neue `AC-*`.
- `tools/verify-ac-form.sh` + [`Makefile`](../../../../Makefile) — Sensor an `verify`.

Zwei Schichten (Agenten-Briefing, Harness-Targets). Das Lastenheft selbst wird **nicht** angefasst.

## 4. Was bewusst nicht getan wird

- **Kein Umschreiben der 19 bestehenden ACs** (§2).
- **Keine Prüfung der Given/When/Then-*Wortwahl*.** Der Sensor prüft, dass die drei Pfade als
  benannte Zeilen existieren — ob der Satz danach wirklich ein Randfall ist, ist semantisch und
  bleibt Review-Sache. Ein Regex über „Given" würde Form-Erfüllung mit Inhalt verwechseln.

## 5. DoD

- [x] [`AGENTS.md`](../../../../AGENTS.md) §5 deklariert die Drei-Pfad-Form plus Out-of-Scope für
      **neue** `AC-*` und benennt das Grandfathering des Bestands (B-15).
- [x] `make verify-ac-form` prüft jede `AC-*` außerhalb der Grandfather-Liste auf Happy ·
      Boundary · Negative · Out-of-Scope; Selbsttest und Negativ-Probe in beide Richtungen.
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft.

## 6. Closure-Notiz

**Geliefert:** die Drei-Pfad-Form für neue `AC-*` in [`AGENTS.md`](../../../../AGENTS.md) §5 und
`make verify-ac-form` als Sensor darauf, mit den 19 bestehenden `AC-*` namentlich grandfathered.
Das Lastenheft selbst blieb unangetastet. **Etappe D ist damit vollständig.**

**Lerneintrag — Form: geschärfte Regel.**
> **Eine Grandfather-Liste, die nicht mitwächst, ist selbst der Sensor.** Die Alternative wäre
> eine Datums- oder Nummern-Heuristik gewesen („alles ab AC-…-012") — die hätte gleich ausgesehen
> und wäre still gescheitert, sobald eine neue `AC-*` in eine bestehende Familie einsortiert wird
> (a-checks Kennungen sind nach Thema gruppiert, nicht chronologisch — eine zwoelfte Regel-Kennung
> waere neu, laege numerisch aber mitten im Bestand). Prüfsatz für künftige Bootstrap-Stufen: *trägt die
> Abgrenzungs-Regel die Neuheit selbst, oder leitet sie sie aus etwas ab, das auch anders
> wachsen kann?*

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. Negativ-Probe beidseitig: eine fingierte Proben-Kennung ohne Pfade macht den Sensor rot,
   dieselbe Lücke in allen 19 grandfatherten `AC-*` lässt ihn still.

**Folge-Slices:** keine aus diesem Slice. Offen aus Etappe B: **C** (`MR-*`-Bereinigung,
inklusive des echten Kerns von B-2) und **F** (Betriebsmodell).

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
