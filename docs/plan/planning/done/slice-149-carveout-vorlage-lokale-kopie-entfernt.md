# slice-149 — Carveout-Vorlage: lokale Kopie entfernt, driftet nicht mehr

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Hinweis 2026-09-05 („`docs/plan/carveouts/carveout.template.md` — die
Templates liegen jetzt in `.harness/baseline/v6.0.0/templates`"), während Etappe D
([slice-148](../done/slice-148-archivierung-etappe-d.md)) lief. Präzedenz:
[`slice-112`](../done/wellenlos/slice-112-lokale-slice-vorlage-entfernen.md) (dieselbe Entscheidung
für `slice.template.md`). [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Vorlagen-Pflege, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

Der Hinweis prüft sich als real: `docs/plan/carveouts/carveout.template.md` ist eine lokale
Kopie der vendored Vorlage. Prüfen, ob sie noch mit `v6.0.0` übereinstimmt — und falls nicht, den
Fund benennen und die Konsequenz ziehen, die `slice-112` für `slice.template.md` bereits gezogen
hat.

## 2. Fund: die eigene Migration hat den Zeiger bewegt, nicht den Inhalt

`git log --follow` zeigt drei berührende Commits. Der letzte —
[`slice-139`](../done/slice-139-beobachtungsregister-migration.md), aus **dieser**
Sitzung — änderte an der Datei **ausschließlich** den Provenienz-Zeiger im Kopf-Kommentar von
`v5.12.0` auf `v6.0.0` (`git show`, ein Ein-Zeilen-Diff). Die Commit-Message dieses Vorgängers
behauptete dabei: *„vier weitere Template-Provenienz-Zeiger auf `v6.0.0` nachgezogen — der
Formvergleich, auf den Etappe A wartete, ist damit abgeschlossen."* Ein `diff` gegen die aktuelle
Baseline-Datei zeigt: **der Formvergleich fand für diese Datei nicht statt.**

Strukturelle Differenzen (`diff` als Beleg, nicht Behauptung):

- Der Abschnitt `## Geschichte` (Audit-Tabelle: Datum · Ereignis · Verweis) fehlt in a-checks
  Kopie vollständig.
- Jeder Abschnitt der Baseline trägt eine `Regeln dieser Sektion: Baseline-Regelwerk
  modul-07-carveouts.md §Ziel-Form: Carveout`-Zeile; a-checks Kopie hat sie nie übernommen (auch
  nicht vor `v6.0.0`).
- Feldwerte weichen ab: `Status` führt bei a-check zusätzlich `Überführt in <Ziel>` (dritter Wert,
  aus dem Trichter in [README §Trichter](../../carveouts/README.md#vor-dem-anlegen-der-trichter));
  `Betroffenes Gate` und `Folge-Slice` sind bei a-check strenger gefasst (Bindung an
  [`AGENTS.md`](../../../../AGENTS.md) §4 bzw. Pflicht-Link statt Beispiel-Pfad).

Das ist selbst ein **Harness-Lügen-Muster** im Sinne von `modul-13`: eine Prüfung wurde als
abgeschlossen behauptet, ohne dass sie für dieses Artefakt lief.

## 3. Entscheidung: entfernen statt nachziehen — Präzedenz, keine neue Abwägung

[`slice-112`](../done/wellenlos/slice-112-lokale-slice-vorlage-entfernen.md) traf exakt diese
Entscheidung für `slice.template.md` aus demselben Grund (Drift zwischen Kopie und vendored Stand)
und benannte ausdrücklich: *„Die Carveout-Vorlage bleibt [damals] … ihr Fall wäre eine eigene
Entscheidung mit eigener Begründung, nicht ein Analogieschluss."* Mit diesem Fund liegt die eigene
Begründung jetzt vor — dieselbe Klasse von Drift, die `slice.template.md` schon einmal zeigte.

**Was nicht entfällt:** a-checks legitime Kürzung um `## Geschichte`. Der Abschnitt trägt
Audit-Einträge einer Wellen-Closure-Prozedur — a-check schließt bis heute keine Welle auditierbar
([README §Audit](../../carveouts/README.md#audit), Fund **B-13**, offen). Das ist keine Drift,
sondern ein zutreffender aktueller Zustand; er gehört als **Anpassungs-Hinweis** dokumentiert, nicht
stillschweigend in eine wiederhergestellte lokale Kopie zurückgeholt.

## 4. Umsetzung

- `docs/plan/carveouts/carveout.template.md` entfernt (`git rm`).
- [`docs/plan/carveouts/README.md`](../../carveouts/README.md) §Form eines Carveouts zeigt jetzt
  auf die vendored Ziel-Form direkt, mit einem „Beim Kopieren anzupassen"-Absatz (Pfade
  course-relativ → a-checks eigene, `## Geschichte` streichen mit Begründung, `Betroffenes Gate` an
  `AGENTS.md` §4 binden) — dieselbe Form wie `AGENTS.md` §5 sie für die Slice-Vorlage schon führt.
- Kein Sensor referenzierte die lokale Datei (`grep` über `.d-check.yml`/`Makefile`/`tools/*.sh`,
  leer) — nichts weiter anzupassen.

## 5. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor.

## 6. DoD

- [x] Drift real gemessen (`git log --follow` + `diff` gegen die aktuelle Baseline-Datei), nicht
      angenommen (§2).
- [x] Lokale Kopie entfernt, Verweis auf vendored Ziel-Form mit Anpassungs-Hinweis nachgezogen (§4).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 7. Closure-Notiz

**Geliefert:** `docs/plan/carveouts/carveout.template.md` entfernt — dieselbe Konsequenz, die
`slice-112` für `slice.template.md` zog, jetzt mit eigener, gemessener Begründung statt
Analogieschluss.

**Lerneintrag — Form: geschärfte Regel.** *Ein Provenienz-Zeiger, der auf die neue Baseline-Version
bumpt, ist kein Beleg dafür, dass der Inhalt dahinter geprüft wurde — beides zusammen zu behaupten
(„Formvergleich abgeschlossen"), ohne einen `diff` gegen die Baseline-Datei geführt zu haben, ist
eine Harness-Lüge, egal wie beiläufig sie entsteht.* Der Fund kam nicht aus einem eigenen Sensor,
sondern aus einer externen Nachfrage — dieselbe Lücke, die `.harness/skills/cr-text-reviewer.md`
für CR-Texte an fremde Werkzeuge benennt (Klasse *eigener Bestand*, Ausprägung „gar nicht
gemessen"), hier erstmals an einer Template-Provenienz-Zeile statt einem CR-Text. Für **lokale
Kopien vendorter Vorlagen ohne eigenen MR-Adaptions-Eintrag** ist die konsequente Antwort dieselbe
wie bei `slice.template.md`: keine Kopie führen, die driften könnte.

**Zwei beobachtbare Closure-Kriterien:**

1. `test -f docs/plan/carveouts/carveout.template.md` → nicht mehr vorhanden.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Keines notiert — der Fund ist mit der Entfernung geschlossen,
kein offener Rest.

**Folge-Slices:** keine vergeben. Ob andere lokale Template-Kopien (`docs/reviews/review-report.template.md`
u. Ä.) denselben Fund tragen, ist mit diesem Slice **nicht** geprüft — eine gezielte Nachfrage, kein
gemessener Befund; wird nicht spekulativ behauptet.

## 8. Sub-Area-Modus

Berührt: **Planungs-Harness** (`docs/plan/carveouts/`) — Greenfield: Form und Provenienz stehen in
der vendored Vorlage, `doc-check` prüft Verweis-Auflösung.
