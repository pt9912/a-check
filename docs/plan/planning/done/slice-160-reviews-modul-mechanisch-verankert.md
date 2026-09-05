# slice-160 — Review-Report-Deckung mechanisch über das `reviews`-Modul erzwungen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Vorgabe („Fix-Umriss für a-check", drei Schritte) im direkten Anschluss an
[slice-159](../done/wellenlos/slice-159-reviewer-rolle-in-agents-md-verankert.md): Die dort in
`AGENTS.md` §6 verankerte Review-Pflicht ist **Workflow-Text, mechanisch nicht erzwingbar** (§4
dieses Slice-Plans dort). d-check hat mit `v0.74.1` ein neues Modul `reviews` (`DC-FA-RVW-001`, in
d-checks eigener Architektur-Entscheidung zum Review-Report-Deckungs-Modul begründet) geliefert,
das genau diese Lücke schließt: DoD-Zusage ↔ Report-Existenz.
[Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Tooling-/Gate-Konfiguration, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

Die in slice-159 rein prosaisch verankerte Review-Pflicht bekommt einen **Sensor**: ein `done/`-
Slice, dessen DoD die Phrase „unabhängiger Review" trägt, aber keinen passenden Report unter
`docs/reviews/` hat, macht `make gates` rot — statt sich (wie 23× zuvor) unbemerkt durchzuschleichen.

## 2. Was getan wurde

1. **Pin gehoben** `v0.69.0` → `v0.74.1` (`d-check.mk`, per `--print-mk` neu erzeugt;
   `DCHECK_DIGEST` auf den neuen Release-Digest gesetzt). `v0.74.1` statt des vom Maintainer
   genannten `v0.74.0`, weil `v0.74.1` ein Sicherheits-Patch ist (d-check-CHANGELOG:
   `[0.74.1] — 2026-09-04`, behebt `CVE-2026-56855`/`CVE-2026-78662`) — der neuere Tag trägt
   dieselbe `reviews`-Fähigkeit plus die Behebung, kein Grund für den älteren.
2. **Modul konfiguriert** — `reviews` zur `modules:`-Liste in [`.d-check.yml`](../../../../.d-check.yml)
   ergänzt, `reviews:` Block mit `done-dir: docs/plan/planning/done` und
   `reviews-dir: docs/reviews` (Schema 1:1 aus d-checks eigener `.d-check.yml` und
   `DC-FA-RVW-001` übernommen).
3. **`doc-reviews`-Target** in [`Makefile`](../../../../Makefile) ergänzt (analog `doc-workflows`:
   eigenes Target, weil das erzeugte `d-check.mk`-Fragment für ein neu aktiviertes Modul kein
   Isolations-Target führt), in `gates:` aufgenommen. Der jetzt falsche Kommentar „Das EINZIGE
   eigene doc-*-Target" über `doc-workflows` korrigiert (es sind jetzt zwei), und `doc-workflows`
   selbst um `--disable reviews` ergänzt (die `--disable`-Liste zählt die aktive `modules:`-Menge
   auf, ein neu hinzugekommenes Modul muss dort nachgezogen werden — dieselbe Falle wie
   slice-115/slice-130).
4. **`AGENTS.md` §4** — neue Zeile für `make doc-reviews`, neue Zeile für `make doc-usage` (von
   `--print-mk` seit `v0.74.1` mit erzeugt, advisory, war vorher nicht dokumentiert weil es noch
   nicht existierte).

## 3. Opt-in-Charakter — bewusst nicht rückwirkend

Das Modul prüft nur Slices, deren DoD die Phrase „unabhängiger Review" **trägt**. Kein a-check-
Slice führt diese Phrase bisher wörtlich — auch slice-159 selbst nicht, obwohl dort ein echter
Review stattfand (`docs/reviews/2026-09-05-slice135-157-multi-linsen-review.md`). Der Sensor ist
mit diesem Slice **scharf, aber vorerst ohne Kandidaten**: das ist kein Fehlalarm (das Modul ist
fail-closed nur bei leerer Kandidaten- bzw. unlesbarem Verzeichnis, nicht bei null Zusagen), aber
auch keine Wirkung. Ob künftige Slices die Phrase in ihre DoD aufnehmen, ist eine offene Frage —
siehe §6 Risiken.

## 4. Auszuführende Gates

`make gates` (inkl. des neuen `doc-reviews`), zum Abschluss `make verify`. Ausgabe in eine Datei,
Exit-Code getrennt geprüft, nie in eine Pipe.

## 5. DoD

- [x] Pin `v0.74.1` gehoben, `reviews`-Modul in `.d-check.yml` konfiguriert (`done-dir`,
      `reviews-dir`).
- [x] `doc-reviews`-Target ergänzt, in `gates:`/`.PHONY` aufgenommen, `doc-workflows` samt
      Kommentar auf den neuen Modulstand nachgezogen, `AGENTS.md` §4 ergänzt.
- [x] `make gates` und `make verify` grün.

## 6. Risiken und offene Punkte

- **R-1: Der Sensor hat vorerst keine Kandidaten.** Ohne die DoD-Phrase „unabhängiger Review" in
  künftigen Slices bleibt `doc-reviews` dauerhaft grün, ohne je etwas zu prüfen — dieselbe stille
  Wirkungslosigkeit wie bei der reinen Prosa-Regel aus slice-159, nur maschinell statt menschlich.
  **Ausgang:** wandert ins Beobachtungs-Register als neue Beobachtung (kein 3×-Muster, Erstauftreten) —
  siehe §7 dieses Abschnitts unten.

## 7. Closure-Notiz

**Geliefert:** Die Review-Pflicht aus slice-159 ist jetzt doppelt abgesichert — Workflow-Text in
`AGENTS.md` §6 UND ein Gate (`doc-reviews`), das eine unterlassene Review-Deckung *mechanisch*
findet, sobald ein Slice sich per DoD-Phrase dazu bekennt.

**Lerneintrag — Form: neuer Sensor.** `make doc-reviews` (Modul `reviews`, `DC-FA-RVW-001`) — eine
`done/`-Slice-DoD-Zeile mit der Phrase „unabhängiger Review" ohne passenden Report unter
`docs/reviews/` macht `gates` rot. Schließt strukturell dieselbe Lücke, die 23 Slices lang offen
stand (`BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen`) — diesmal nicht durch eine
Checklisten-Zeile, die man übersehen kann, sondern durch ein Gate, das man nicht übersehen kann.

**Zwei beobachtbare Closure-Kriterien:**

1. `grep -c "^  reviews:" .d-check.yml` → mindestens `1`, UND `grep -c "doc-reviews" Makefile` →
   mindestens `2` (Target-Definition + `gates:`-Aufnahme).
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken:** siehe §6 — R-1 trägt seinen Ausgang dort inline (Register-Eintrag angelegt,
`offen`).

**Folge-Slices:** keine vergeben. Die Frage, ob und wie a-check-Slices künftig die DoD-Phrase
„unabhängiger Review" führen, bleibt bewusst offen — das wäre eine Konventions-Entscheidung, keine
Tooling-Frage, und würde diesen Slice über seine drei Liefer-Punkte hinaus dehnen.

## 8. Sub-Area-Modus

Berührt: **Harness-Tooling** (`.d-check.yml`, `Makefile`, `d-check.mk`) — Greenfield: Konventionen
(Isolations-Target-Muster, `gates:`-Aggregat-Disziplin) entstehen mit dem Repo und werden hier nur
fortgeschrieben, kein Bestands-Abgleich nötig.

### Vorgelagert — offene Beobachtungen sichten

[Beobachtungs-Register](../observations/README.md) durchgesehen: keine offene Beobachtung zur
Sub-Area „Harness-Tooling" mit Stand `offen` unterhalb der Schwelle, die dieser Slice berühren
würde. Keine Treffer.
