# slice-104 — Der externe Zeiger steht einmal, nicht zweimal

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers.
**Berührte Spec-Stellen:** — *(keine)*
**Deckt:** keine `AC-*`/`ADR-*`.
**Bezug:** Maintainer-Vorgabe 2026-08-29, unmittelbar nach
[slice-103](../done/slice-103-conventions-ohne-chronik.md).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

Das Bullet *Extern (Lehrmaterial)* in §Adoptierte Konventions-Quellen soll weg.

**Warum das trägt, obwohl die Pflichtgliederung „Pointer extern" nennt:** der externe Zeiger geht
nicht verloren, er steht eine Sektion höher. §Baseline nennt den adoptierten Stand als
**Release-Tag mit URL** — dieselbe Quelle, dieselbe Version, nur an der Stelle, an der die
Pflichtgliederung sie ohnehin verlangt (*„welche Konvention adoptiert, mit Stand/Version"*). Was
tatsächlich entfällt, ist ein zweiter Verweis auf dieselbe Sache plus der Zeiger auf `kurs/de/` —
Material, das dieses Repo netzlos weder lädt noch laden soll.

Damit ist es dieselbe Bewegung wie in slice-103: **eine Angabe steht einmal, nicht zweimal.**

## 2. Betroffene Module

`harness/conventions.md`, ein Bullet. Eine Schicht.

## 3. Auszuführende Gates

`make gates` — `doc-check` prüft, dass kein Verweis ins Leere zeigt. Zum Abschluss `make verify`.

**Kein neuer Sensor.** Die Probe ist, dass der Stand nach der Kürzung weiterhin genau einmal mit
Tag und URL im Dokument steht — nachzählbar.

## 4. Was bewusst nicht getan wird

- **Keine Adaption deklariert.** Der Abschnitt verliert keinen Inhalt, nur eine Wiederholung; eine
  Abweichung von der Pflichtgliederung entsteht dadurch nicht. Ein `MR`-Eintrag für eine
  Nicht-Abweichung wäre genau der Fork, den die Baseline ausschließt.
- **§Baseline bleibt unberührt.** Sie trägt den externen Zeiger schon; ihn dort zu verstärken wäre
  wieder Dopplung.

## 5. DoD

- [ ] Das Bullet *Extern (Lehrmaterial)* ist entfernt; §Adoptierte Konventions-Quellen trägt die
      vendored Lese-Form und die In-Repo-Form — Beleg: Diff.
- [ ] Der externe Stand steht weiterhin **genau einmal** im Dokument, mit Tag und URL, in
      §Baseline — Beleg: `grep` über die Datei.

Pflicht, aber **kein** Liefer-Punkt: `make gates` und zum Abschluss `make verify` grün — Ausgabe
in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Berührt wird **Harness-Einstieg** — Greenfield.
