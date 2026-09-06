# slice-168 — Kalibrierungs-Selbsttest für phrasen-/muster-basierte Prüfer

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (Trigger: 3×-Schwelle im Beobachtungs-Register,
zugewiesen beim Lese-Schritt der `welle-14`-Closure).

**Bezug:** [`BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf`](../observations/BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf/observation.md)
(3×: slice-120, slice-123, slice-165).

**Berührte Spec-Stellen:** — *(noch offen — Ziel dieses Slice ist zunächst
eine Analyse, welche Spec-Stelle oder welches Werkzeug betroffen ist)*.

**Verantwortlich:** — *(noch nicht priorisiert)*.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

---

## 1. Ziel

Einen wiederkehrenden Fehler verhindern: ein phrasen- oder
muster-basierter Prüfer (`.d-check.yml`-Modul-Konfiguration oder
`tools/*.sh`) meldet grün, weil seine Prüfmenge leer ist (kein Text
trifft das gesuchte Muster) oder weil er nie aufgerufen wird — nicht weil
tatsächlich geprüft wurde. Drei bisherige Fälle, alle unterschiedlicher
Natur:

- `verify-ac-form` suchte `^**Happy Path:**`, einen Wortlaut, den das
  Repo null Mal führte (slice-120).
- `doc-complete` war advisory und lief nie, obwohl es einen Gegenstand
  gehabt hätte (slice-123).
- `make doc-reviews`s Trigger-Phrase „unabhängiger Review" wurde von
  a-checks eigener, tatsächlich verwendeter DoD-Formulierung
  („Unabhängiges Plan-Review …") nie getroffen — vier Slices liefen vier
  Mal durch, ohne dass der Prüfer je etwas sah (slice-165).

**Nicht Ziel dieses Slice:** die drei bereits gefundenen Einzelfälle
erneut zu beheben (das ist bereits geschehen). Ziel ist ein **genereller**
Schutz, der eine vierte Wiederholung verhindert.

## 2. Analyse-Fragen (vor jeder Umsetzung zu klären)

- Betrifft die Lücke primär `.d-check.yml`-Modul-Konfigurationen
  (phrasen-basierte Muster wie `tasks-ignore-pattern`, Trigger-Phrasen)
  oder auch a-checks eigene `tools/*.sh`-Prüfer?
- a-checks eigene Bash-Prüfer (`tools/verify-risiko-ausgaenge.sh`) tragen
  bereits einen `self_test()` mit Gut-/Schlecht-Fixtures — ist das ein
  Muster, das sich auf `.d-check.yml`-Modul-Konfigurationen übertragen
  lässt (z. B. eine Pflicht-Fixture-Datei pro konfiguriertem Modul, gegen
  die `make doc-check`/`doc-structure`/`doc-reviews` selbst getestet
  werden)? Oder ist das strukturell nicht möglich, weil `d-check` ein
  externes, nicht in a-check eingebettetes Werkzeug ist?
- Ist eine **repo-weite** Lösung realistisch, oder bleibt es bei der
  bisherigen Praxis (empirischer Test vor jeder neuen Phrasen-Einführung,
  wie in `slice-165`/`slice-166`/`slice-167` demonstriert) — dann wäre
  der „Sensor" aus dem Steering-Loop-Eintrag eine **Checkliste im
  Workflow**, kein automatisiertes Gate?

## 3. Definition of Done

- [ ] Analyse-Fragen aus §2 beantwortet, Umsetzungsweg entschieden
      (automatisiertes Gate vs. Workflow-Checkliste vs. beides).
- [ ] Gewählte Lösung umgesetzt.
- [ ] `BEO-GATE/pruefer-ohne-gegenstand-oder-aufruf` auf `verkörpert`
      gesetzt, mit auflösbarem Zielort und Herkunfts-Anker.
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

## 4. Risiken und offene Punkte

- *Eine repo-weite, automatisierte Lösung ist für extern konfigurierte
  `d-check`-Module nicht möglich (das Werkzeug liegt außerhalb von
  a-checks Kontrolle)* — Ausgang noch offen, Teil der Analyse in §2.

## 5. Closure-Notiz

*(folgt bei Abschluss)*
