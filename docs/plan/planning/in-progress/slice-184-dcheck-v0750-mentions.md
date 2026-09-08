# slice-184 — `d-check` auf `v0.75.0`, und das neue Modul `mentions` schließt eine benannte Grenze

**Welle:** ohne Welle — reaktiv (neue Werkzeug-Version verfügbar); die
Closure-Bedingung wäre die eigene DoD (Baseline-Regelwerk
`modul-06-roadmap.md` §Wann Arbeit eine Welle braucht).

**Bezug:** [`AC-QA-02`](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)
(Modul-Integrität des Doku-Stacks), [`AC-QA-03`](../../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)
(Digest-Pin). Keine aktive ADR wird berührt.

**Berührte Spec-Stellen:** — (Werkzeug-Pin und Gate-Konfiguration ohne Vertragsberührung).

**Verantwortlich:** Claude — gesetzt beim Übergang nach `in-progress/`.

**Autor:** Claude. **Datum:** 2026-09-08.

**Lerneintrag — Form:** neuer Sensor.

---

## 1. Ziel und Abgrenzung

**Ziel:** Der `d-check`-Pin steht auf `v0.75.0`, und dessen einzige Neuerung —
das Modul **`mentions`** — wächtert die **Gegenrichtung** des Link-Checks: ob
eine *existierende* Datei irgendwo genannt wird.

**Ausdrücklich NICHT in diesem Slice** — je Punkt mit Begründung:

- **Andere Doku-Familien als `harness/sensors/` und die ADRs.** Für beide ist
  der Bedarf **benannt** — die Ziel-Form von `harness/README.md` nennt die Lücke
  wörtlich, und `gate-consistency` trägt für die ADRs eine Eigenbau-Prüfung.
  Für Slice-Pläne, Carveouts oder Beobachtungen gibt es keinen solchen Anlass;
  ein Wächter ohne Anlass ist selbst eine Behauptung.
- **Die `citations`- und `sources`-Module.** Sie waren schon in `v0.74.1`
  verfügbar und sind unkonfiguriert — das ist Bestand, der bewusst stehen
  bleibt, und hat mit diesem Sprung nichts zu tun.
- **Produkt-Code.** Der Slice rührt `internal/` nicht an; er arbeitet in der
  Gate-/Werkzeug-Schicht.

## 2. Ausgangsmessung (2026-09-08)

### 2.1 Was der Sprung `v0.74.1 → v0.75.0` bringt

**Gemessen** über `--print-config` beider Digests, Diff über die volle Ausgabe:
**vier Zeilen**, und alle vier sagen dasselbe — die Modul-Liste bekommt
`mentions`. Kein anderer Konfigurations-Block ändert sich.

Der `--print-mk`-Diff zeigt die Folge: Jedes Target bekommt zusätzlich
`--disable mentions`. Das Modul ist **strikt opt-in**.

**Neuer Digest:** `sha256:18e9cd857f8db3569526d1f9a3cbeba8af51e9f2dd84c17a22444028b977c3da`
(alt: `sha256:e31a372b66dbde26305982424854cfce7c9ab7ce555a94debeee7ee26e6d4641`).

### 2.2 Was `mentions` tut — ausprobiert, weil `--print-config` es nicht sagt

Das Modul hat **keinen** dokumentierten Block in `--print-config`; es kommt dort
genau einmal vor, in der Verfügbar-Liste. Ermittelt gegen eine eigene Fixture:

```
art/zwei.md:1   docs/**/*.md   artifact-unmentioned
                kein Vorkommen als "art/zwei.md" in der Ist-Menge
```

- Es braucht **beide** Listen, sonst bricht es ab:
  `das Modul mentions braucht mentions.artifacts UND mentions.documents
  (DC-FA-MENT-001, fail-closed)`.
- Beide Felder sind **Listen** (`["glob"]`), nicht Einzel-Strings — ein String
  bricht mit `cannot unmarshal !!str into []string`.
- Die Erfolgszeile nennt die Deckung: *„1 von 2 Artefakt(en) erwähnt, über 1
  Dokument(e)"* — also **nicht** nur ein Ja/Nein, sondern eine Quote.

### 2.3 Warum a-check das braucht — die Grenze ist benannt, nicht erfunden

`v6.5.0` · `templates/harness/README.template.md` beschreibt sie wörtlich:

> *„Seine Grenze: Er prüft EINE Richtung — ob das Ziel existiert; **eine Datei
> ohne Index-Zeile** und eine Zeile auf die falsche Datei bleiben still grün."*

Dieselbe Grenze steht seit [slice-181](../done/wellenlos/slice-181-tabellenzellen-gewaechtert.md)
in `harness/sensors/gate-consistency.md`, und dort trägt eine **Eigenbau-Prüfung**
sie für die ADRs (ADR-Index-Vollständigkeit, slice-087).

**Der Bestand ist heute sauber** — 0 von 15 Sensor-Dateien und 0 von 39 ADRs
sind unverlinkt. Die Lücke ist **latent**, nicht offen. Das ist der richtige
Zeitpunkt für einen Wächter: Er kostet nichts und fängt den ersten Fall, statt
ihn zu finden, wenn er schon steht.

## 3. Umsetzung

*(entsteht mit der Arbeit)*

## 4. Definition of Done

- [ ] Pin auf `v0.75.0`: `d-check.mk` **verbatim** aus `--print-mk` des neuen
      Digests, `DCHECK_DIGEST` gesetzt, alle Versions-Nennungen nachgezogen.
- [ ] `mentions` konfiguriert für `harness/sensors/**` gegen die zwei
      Index-Tabellen; Mutations-Probe in **beide** Richtungen — eine Datei ohne
      Index-Zeile meldet rot, der unveränderte Bestand grün.
- [ ] Entschieden und **gemessen**, ob `mentions` die ADR-Index-Eigenbau-Prüfung
      in `gate-consistency` ablöst — Parität in beiden Richtungen, wie slice-079
      sie für `doc-targets` gemessen hat.
- [ ] Unabhängiger Review durchgeführt (Report unter [`docs/reviews/`](../../../reviews/README.md)).
- [ ] Closure-Notiz mit Steering-Loop-Lerneintrag.
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §7 trägt einen Ausgang.

`make gates` und `make verify` grün. Ein öffentlicher Vertrag ist berührt:
`d-check.mk` und `.d-check.yml` sind Teil des Harness-Stacks, und `AGENTS.md` §4
nennt die Targets.

## 5. Trigger

**Start** (`open` → `in-progress`): WIP-Limit frei. Der Anlass — `v0.75.0`
verfügbar — ist eingetreten.

**Rückführungen — vorab benannt:**

- `in-progress` → `next` (zu groß): Zeigt die Paritäts-Messung, dass die
  Ablösung der Eigenbau-Prüfung einen eigenen Umbau braucht, wird sie
  abgetrennt und der Slice trägt nur Pin und Konfiguration.
- `in-progress` → `open` (blockiert): Bricht `v0.75.0` ein bestehendes Gate,
  bleibt der Pin auf `v0.74.1`, und der Bruch wird als CR-Text an `d-check`
  formuliert ([`AGENTS.md`](../../../../AGENTS.md) §5).

## 6. Closure-Trigger

DoD vollständig, `make gates` und `make verify` grün, Closure-Notiz mit
Lerneintrag. Danach Archivierung als wellenloser Slice
([`AGENTS.md`](../../../../AGENTS.md) §6).

## 7. Risiken und offene Punkte

- **`mentions` ist undokumentiert.** Sein Verhalten ist ausprobiert, nicht
  gelesen — was es bei Sonderfällen tut (Inline-Code statt Link, Teilpfade,
  Groß-/Kleinschreibung), ist ungeprüft. — **Ausgang:** <offen bis Closure>
- **Der Sprung kann ein bestehendes Gate brechen**, ohne dass der Diff der
  Konfiguration es zeigt: Ein Modul kann sein *Verhalten* geändert haben, wo
  seine *Optionen* gleich blieben. — **Ausgang:** <offen bis Closure>
- **Die Ablösung der Eigenbau-Prüfung könnte den Geltungsbereich verkleinern** —
  dieselbe Falle, die slice-079 für `doc-targets` gemessen und vermieden hat.
  — **Ausgang:** <offen bis Closure>

## 8. Closure-Notiz

*(bei Closure auszufüllen)*

## 9. Sub-Area-Prüfungen und Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** *(beim Übergang nach `in-progress/`
auszufüllen — der Register-Stand beim Anlegen ist ein anderer als beim Beginn
der Arbeit.)*

**Vorgelagert — offene Beobachtungen sichten:** *(ebenso.)*
