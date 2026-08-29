# slice-106 — Roadmap: `Offene Wellen` statt `Aktuelle Welle`

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers.
**Berührte Spec-Stellen:** — *(keine)*
**Deckt:** keine `AC-*`/`ADR-*`.
**Bezug:** Befund B/§5 aus [slice-105](../done/slice-105-form-review-nachholen.md).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

Die Roadmap führt `## Aktuelle Welle`; die Sollform heißt `## Offene Wellen`. Das ist kein Rename:
`modul-06` hängt an dem Abschnitt **zwei unabhängige Aussagen**.

- **Die Liste ist derivativ** — ein Zeiger je *flacher Welle-Datei*, Bijektion in beide
  Richtungen. a-check hat keine flachen Welle-Dateien; die Liste ist hier also leer, aber sie ist
  eine Liste und keine Prosa.
- **Der Ruhe-Marker** steht **zusätzlich** zur Liste, wenn `in-progress/` keinen Slice trägt —
  nicht an ihrer Stelle. Sein Wortlaut steht in `modul-06`, **absichtlich nicht** in der Vorlage:
  ein Doku-Sensor matcht ihn sonst als Substring des Regeltexts und meldete Ruhe bei
  beanspruchtem Slice.

Der Plural ist Absicht. `modul-06` warnt ausdrücklich, ein Wächter, der den Abschnitt gegen
*genau eine* Datei hält, melde legitime Zustände als Drift.

Dazu fehlen der Datei die **sechs Regelwerk-Zeiger** und die **Format-Regel** im Kopf
([slice-105](../done/slice-105-form-review-nachholen.md) §4 A).

## 2. Betroffene Module

`docs/plan/planning/in-progress/roadmap.md` und der Fremdverweis in
`docs/plan/planning/README.md` Schritt 5. Eine Schicht.

## 3. Was dabei zusätzlich entfällt

Der heutige Abschnitt trägt eine Liste *Erledigt ohne Welle* mit drei Slice-Einträgen und je einer
Zusammenfassung. Das ist ein **Statusbericht über Slices** — und die Roadmap sagt in ihrer eigenen
Präambel, dass sie keiner ist: *„Was ein einzelner Slice geliefert und gelernt hat, steht in seiner
Closure-Notiz unter `done/`."* Bis slice-067 taten das zwei gewachsene Status-Blöcke; die Liste ist
derselbe Wildwuchs, eine Ebene kleiner. Sie entfällt — nach derselben Regel, aus der slice-103 die
Chronik aus `conventions.md` genommen hat.

## 4. Auszuführende Gates

`make gates` — `doc-check` prüft die neuen Anker und den nachgezogenen Fremdverweis; zusätzlich
`make doc-planning` (Roadmap ↔ `in-progress/`). Zum Abschluss `make verify`.

**Kein neuer Sensor** — aber eine offene Frage, die dieser Slice **benennt statt löst**: ob
`doc-planning` die Marker-Hälfte überhaupt hält. Das Modul läuft ohne Konfigurationsblock in
`.d-check.yml`, und genau diese Konstellation ließ das Modul `targets` laut slice-074 dreizehn
Minor-Versionen ins Leere laufen.

## 5. Was bewusst nicht getan wird

- **Keine `Trigger`-Spalte in der Meilenstein-Tabelle.** Sie fehlt zwar gegenüber der Sollform,
  aber schon gegenüber `v3.5.2` — das ist eine **ältere** Abweichung, nicht Teil dieser Migration.
  Sie zu füllen hieße, Trigger für drei erreichte Meilensteine nachträglich zu erfinden.
- **Keine flachen Welle-Dateien angelegt.** Die Liste ist leer, weil es keine offenen Wellen gibt
  — nicht, weil eine Datei fehlt.

## 6. DoD

- [x] Der Abschnitt heißt `## Offene Wellen` und trägt **beide** Aussagen: eine (leere) Liste und
      den Ruhe-Marker im Wortlaut aus `modul-06` — Beleg: Diff.
- [x] Kopf trägt die Format-Regel, jede der sechs Sektionen ihren Regelwerk-Zeiger; der
      Statusbericht *Erledigt ohne Welle* ist entfallen — Beleg: Diff und Zeiger-Zählung 6/6.
- [x] Der Fremdverweis in `docs/plan/planning/README.md` nennt den neuen Abschnittsnamen;
      `doc-check` und `doc-planning` bleiben grün — Beleg: Target-Ausgabe.

Pflicht, aber **kein** Liefer-Punkt: `make gates` und zum Abschluss `make verify` grün — Ausgabe
in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

**Geliefert:** `## Offene Wellen` mit beiden Aussagen, sechs Regelwerk-Zeiger, die Format-Regel im
Kopf — und der Statusbericht *Erledigt ohne Welle* ist entfallen.

**Lerneintrag — Form: benannte Spec-Lücke.** *Die Marker-Hälfte des Abschnitts ist ungewächtert,
und das ist beim Bauen zufällig bewiesen worden.* Während dieser Slice in `in-progress/` lag,
stand im Abschnitt bereits der Ruhe-Marker — also die Aussage „kein Slice beansprucht", bei
beanspruchtem Slice. `modul-06` nennt das ausdrücklich einen Defekt (*„ein stehengebliebener
Marker bei beanspruchtem Slice"*), und `make doc-planning` lief darüber **grün**. Damit ist
gemessen, was §4 nur vermutet hatte: das Modul `planning` hält die Marker-Hälfte nicht — es läuft
ohne Konfigurationsblock in `.d-check.yml`, dieselbe Konstellation, in der das Modul `targets`
laut slice-074 dreizehn Minor-Versionen ins Leere lief. *Weil* ein aktiviertes Modul ohne
Konfiguration nicht schweigt, sondern **grün meldet** — und grün von „geprüft und in Ordnung"
nicht zu unterscheiden ist.

**Zwei beobachtbare Closure-Kriterien:**

1. Sechs Regelwerk-Zeiger in der Datei (gezählt), Abschnitt heißt `## Offene Wellen`, Liste und
   Marker stehen **beide** — nicht einer statt des anderen.
2. `doc-check` 221 Dateien 0 Befunde und `doc-planning` Exit 0, **auch** im widersprüchlichen
   Zwischenzustand — das ist der Beleg für den Lerneintrag, nicht sein Gegenbeweis.

**Offene Risiken und ihr Ausgang:**

- *Die Marker-Hälfte ist ungewächtert* — Ausgang: **weiter offen**, als `BEO-014` im
  Beobachtungs-Register, mit dem Zwischenzustand oben als Beleg.
- *Die Wahrheit des Markers kippt mit dem `git mv`* — er ist während `in-progress/` falsch und
  danach richtig. Ausgang: **weiter offen**, gedeckt durch `BEO-006`: ein Marker-Wächter muss am
  selben Punkt laufen wie der Closure-Notiz-Wächter, nämlich **nach** dem `mv`.
- *Die `Trigger`-Spalte der Meilenstein-Tabelle fehlt* — Ausgang: **weiter offen**, als `BEO-013`;
  sie fehlt schon gegenüber `v3.5.2`, ist also keine Migrationsfolge, und sie zu füllen hieße,
  Trigger für drei erreichte Meilensteine zu erfinden.

**Folge-Slices:** die Slice-Vorlage (Befund §5 aus slice-105), danach die übrigen
Regelwerk-Zeiger.

## 8. Sub-Area-Modus

Berührt wird **Planungs-Harness** — Greenfield.
