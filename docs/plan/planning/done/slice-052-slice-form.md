# slice-052 — Etappe D (1/3): Slice-Form als Vorlage und Sensor

**Status:** in-progress — erster Schnitt der **Etappe D (Form)** aus
[slice-048 §5](../done/slice-048-modul-delta-lesen.md).
**Deckt:** Fund **B-1** (Größen-Regel nirgends abgebildet) und **B-5** (Lerneintrag-Form nicht
benannt). **Nicht hier:** B-6/B-18/B-12 (Lifecycle und Roadmap-Form) und B-15 (AC-Form) — 2/3 und
3/3. [Roadmap](../in-progress/roadmap.md).

Dieser Slice ist zugleich die **erste Anwendung** der Vorlage, die er anlegt.

---

## 1. Auslöser

Die Baseline setzt „≤ 3 DoD-Punkte, höchstens zwei Schichten" als *harte* Schnitt-Regel und
verlangt den Lerneintrag in **einer von drei benannten Formen**. Gemessen in slice-048: a-check
lag bei 4–7 DoD-Punkten, und die Regeln standen nirgends im Repo — es gab **keine Slice-Vorlage**.
Eine Regel ohne Ort ist keine Regel, sondern eine Erinnerung.

## 2. Betroffene Module

- `docs/plan/planning/slice.template.md` — neu, aus der vendored Vorlage **übersetzt**: Verweise
  zeigen auf `.harness/baseline/…` und auf a-checks eigene Regeln statt auf Kurs-URLs.
- `tools/verify-slice-form.sh` + [`Makefile`](../../../../Makefile) — der Sensor, an `verify`
  gehängt.
- [`AGENTS.md`](../../../../AGENTS.md) §4/§5 — Target-Zeile und Verweis auf die Vorlage.

Zwei Schichten (Planungs-Doku, Harness-Targets).

### Korrektur an slice-048

Fund **B-2** behauptete, die `§8 Sub-Area-Modus-Begründung` fehle in fast allen Slices. Beim Lesen
der Vorlage zeigt sich: **§8 ist nur Pflicht, wenn mindestens eine berührte Sub-Area Brownfield
oder Hybrid ist**; bei reinem Greenfield genügt ein Satz. Da a-check durchgehend GF ist, war das
Fehlen des Blocks *nicht* der Verstoß, als den slice-048 ihn führte. Der echte Kern von B-2 bleibt
bestehen und liegt in Etappe C: der **pauschale Repo-Modus** (`*` → Greenfield) ist gegen die
Baseline falsch geschnitten. Die Vorlage bildet §7 darum in der Kurzform ab.

## 3. Auszuführende Gates

`make verify` (neu: `verify-slice-form`) und `make gates`. Negativ-Probe für den neuen Sensor:
ein vierter DoD-Punkt und eine Closure ohne Form-Angabe müssen ihn rot machen; ein alter Slice
mit denselben Mängeln darf ihn **nicht** rot machen — sonst wäre das Grandfathering wirkungslos.

## 4. Was bewusst nicht getan wird

- **Keine Rückwirkung.** Die 51 Slices vor diesem sind grandfathered. Ein rückwirkendes
  Umschreiben wäre Geschichts-Politur; dieselbe Mechanik benutzt die Baseline für ihre
  Referenz-Richtungs-Regel („prüft nur ab Einführung neu"). Der Stichtag ist als Reifestufe
  **mit Trigger** im Skript-Kopf dokumentiert — eine Stufung ohne Trigger wäre laut Modul 7 ein
  „Bootstrap-Schlupfloch".
- **Kein Zähler für „höchstens zwei Schichten".** Was eine Schicht ist, ist eine Ermessensfrage;
  ein Zähler darüber wäre Schein-Genauigkeit. Bleibt Review-Sache — und steht so im Skript.

## 5. DoD

- [x] `docs/plan/planning/slice.template.md` existiert mit Größen-Regel, den drei
      Lerneintrag-Formen und §7 in der GF-Kurzform; in [`AGENTS.md`](../../../../AGENTS.md) §5
      verankert (B-1, B-5).
- [x] `make verify-slice-form` prüft ab slice-052 die DoD-Zahl und die benannte Lerneintrag-Form,
      mit Selbsttest **und** Negativ-Probe in beide Richtungen (greift ab Stichtag, schweigt
      darunter).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft.

## 6. Closure-Notiz

**Geliefert:** `docs/plan/planning/slice.template.md` als erste Slice-Vorlage des Repos und
`make verify-slice-form` als bootstrap-aware Sensor darauf. Dieser Slice ist die erste Anwendung
beider.

**Lerneintrag — Form: neuer Sensor.**
> Die Größen-Regel war seit slice-048 *bekannt* und wurde trotzdem im selben Atemzug gebrochen —
> slice-047 hatte sieben DoD-Punkte, slice-046 sechs. Bekanntheit ist keine Durchsetzung. Der
> Sensor macht aus der Regel eine Beobachtung: ein vierter Haken bricht `make verify`, *weil* die
> Zahl gezählt und nicht eingeschätzt wird. Zugleich zeigt die Gegenprobe, dass der Stichtag
> trägt — derselbe Mangel in slice-040 bleibt still, also politur-frei.

**Zwei beobachtbare Closure-Kriterien:**

1. `make verify` und `make gates` grün auf demselben Stand (je Exit 0) — belegt.
2. Die Negativ-Probe ist **beidseitig** dokumentiert: ab Stichtag rot (slice-052 mit viertem
   Punkt), darunter still (slice-040 mit vier Punkten) — ein Grandfathering, das nur in eine
   Richtung geprüft ist, ist ein unbelegter Freibrief.

**Folge-Slices:** [slice-053](slice-053-lifecycle-und-roadmap-form.md) (Etappe D, 2/3).

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
