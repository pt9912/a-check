# slice-159 — Reviewer-Rolle in `AGENTS.md` §6 verankert

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Auftrag „Deshalb sollten wir jetzt ein paar Reviews nachholen, sonst können
wir diese Muster nicht brechen" — die strukturelle Korrektur zum in
[slice-158](../done/wellenlos/slice-158-archive-wave-vierte-stelle-und-review-korrekturen.md) §7
mit Kennung versehenen Register-Eintrag
[`BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen`](../observations/BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen/observation.md)
(4×, Schwelle überschritten). [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Workflow-Änderung, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

Die Reviewer-Rolle so in `AGENTS.md` §6 verankern, dass ein Agent, der §6 als vollständige
Checkliste behandelt (genau der Fehler, der zu 23 ungeprüften Slices führte), sie nicht mehr
übersehen kann.

## 2. Warum `docs/reviews/README.md` allein nicht reichte

Die Review-Pflicht stand bereits seit slice-043 in `docs/reviews/README.md`: „Ab slice-043
entsteht die Synthese wieder vor dem Merge, und der Slice-DoD führt sie als eigenen Punkt." Das
war korrekte, aber **wirkungslose** Dokumentation — sie lag an einem Ort, den die eigentliche
Ausführungs-Checkliste (`AGENTS.md` §6) nie referenzierte. 23 Slices (`135`…`157`) liefen ohne
Review, nicht weil die Regel fehlte, sondern weil sie am falschen Ort stand, um auf den
tatsächlichen Arbeitsablauf zu wirken (Modul-0-Prinzip:
[`harness/README.md`](../../../../harness/README.md)).

## 3. Was ergänzt wurde

Neuer Absatz in `AGENTS.md` §6, zwischen der 8-Schritte-Liste und dem `make verify`-Hinweis
(Reihenfolge nach Modul 1: Code → Review → Verifikation → Closure): vor dem Abschluss jedes
Slice mit Code- oder Vertragsänderung ein Review gegen Plan/ADR/Konventionen — **über einen echten
getrennten Kontext**, ausdrücklich mit der Klarstellung, dass ein `fork`-Subagent (erbt den
eigenen Kontext) dafür **nicht** zählt. Synthese als `docs/reviews/<YYYY-MM-DD>-*.md` nach
Template. Der Absatz benennt seinen eigenen Auslöser (den 23-Slice-Ausfall,
[`BEO-HARNESS`](../observations/BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen/observation.md))
explizit — nicht als Chronik, sondern als Beleg, warum die Zeile **hier** und nicht nur in
`docs/reviews/README.md` steht.

## 4. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor — die Regel ist Workflow-Text,
mechanisch nicht erzwingbar (dieselbe Grenze wie bei der Selbst-Archivierungsroutine,
[slice-157](../done/wellenlos/slice-157-selbst-archivierung-ab-jetzt.md)).

## 5. DoD

- [x] Reviewer-Schritt in `AGENTS.md` §6 verankert, mit Begründung warum an dieser Stelle statt
      nur in `docs/reviews/README.md` (§2, §3).
- [x] Kontext-Trennungs-Anforderung explizit gemacht (`fork` zählt nicht) — schließt genau die
      Lücke, die diese Sitzung die ganze Zeit über hatte.
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** die Reviewer-Rolle ist jetzt an der Stelle verankert, die tatsächlich als
Checkliste dient — nicht nur an der Stelle, die sie beschreibt.
[`BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen`](../observations/BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen/observation.md)
erreicht mit vier Belegen die Schwelle und ist mit diesem Slice verkörpert (`state.md`, Anker
`seit slice-159`).

**Lerneintrag — Form: geschärfte Regel.** *Eine Regel, die an einem Ort steht, der nicht die
tatsächliche Ausführungs-Checkliste eines Agenten ist, wirkt nicht — unabhängig davon, wie lange
sie schon korrekt dokumentiert ist.* `docs/reviews/README.md` hatte die Review-Pflicht seit
slice-043 richtig stehen; sie wirkte trotzdem nicht, weil `AGENTS.md` §6 — die Datei, die ein
Agent tatsächlich als „was ist zu tun" liest — sie nie zitierte. Dieselbe Lehre wie bei der
Selbst-Archivierung ([slice-157](../done/wellenlos/slice-157-selbst-archivierung-ab-jetzt.md)):
ein Vorschlag in einer Analyse oder eine Regel in einem Nachbar-Dokument ist Historie, keine
Handlungsanweisung — nur was in der aktiven Checkliste steht, wirkt auf den nächsten Lauf.

**Zwei beobachtbare Closure-Kriterien:**

1. `grep -c "Vor dem Abschluss, bei jedem Slice mit Code- oder Vertragsänderung" AGENTS.md` → `1`.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Keines notiert — die Verkörperung ist mit diesem Slice
abgeschlossen. Ob die neue Zeile tatsächlich befolgt wird (mechanisch nicht erzwingbar), zeigt
sich erst an künftigen Slice-Closures — das ist dasselbe, bereits in
[slice-157](../done/wellenlos/slice-157-selbst-archivierung-ab-jetzt.md) registrierte
Praxis-Kosten-Risiko der Selbst-Archivierung, kein neues.

**Folge-Slices:** keine vergeben. Die in
[slice-158](../done/wellenlos/slice-158-archive-wave-vierte-stelle-und-review-korrekturen.md) §4
benannte Wellen-Modus-Dangling-Review-Lücke, die MR-Zahl-Korrektur (`slice-135`/`136`) und die
Spec-Präzisierungen (`spec/architecture.md`, `spec/lastenheft.md` aus dem Review-Report) bleiben
offen — eigene Folge-Slices, nicht Teil dieser Verankerung.

## 7. Sub-Area-Modus

Berührt: **Harness-Einstieg** (`AGENTS.md`) — Greenfield: Briefing und Konventionen entstehen vor
der Regel, die sie beschreiben.
