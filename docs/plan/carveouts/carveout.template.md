# CO-NNN: <Kurztitel des Carveouts>

> **Vorlagen-Hinweis.** Kopieren nach `docs/plan/carveouts/CO-<NNN>-<kurztitel-kebab>.md`,
> Platzhalter ersetzen, diesen Block löschen. Den Index in
> [`README.md`](README.md) mitpflegen. Übersetzt aus
> `.harness/baseline/v3.5.2/templates/docs/plan/carveouts/carveout.template.md`
> ([MR-006](../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert));
> die Verweise zeigen auf a-checks Orte statt auf Kurs-Pfade. Angelegt in slice-065.
>
> **Vor dem Ausfüllen den Trichter durchlaufen** ([README §Trichter](README.md#vor-dem-anlegen-der-trichter)):
> Cluster ⇒ BF-Sub-Area-Markierung, unerreichbarer Trigger ⇒ permanente ADR. Nur eine einzelne
> Diskrepanz mit ernsthaft erreichbarem Trigger gehört hierher.

**Status:** Aktiv | Aufgelöst (`done/CO-NNN-…md`) | Überführt in `<Ziel>`

**Datum angelegt:** YYYY-MM-DD. **Letzte Prüfung:** YYYY-MM-DD.

**Betroffenes Gate:** `<make-target>` — eines der in
[`AGENTS.md`](../../../AGENTS.md) §4 deklarierten Targets, kein erfundenes.

**Geltungsbereich:** <Pfad / Modul / Datei>

**Folge-Slice:** `docs/plan/planning/<zustand>/slice-NNN-….md` (beim Ausfüllen als echten Link setzen) — **Pflicht.** Ohne
Folge-Slice ist der Carveout de facto permanent und gehört über den Trichter in eine ADR.

---

## Begründung

<!--
Warum kann das Gate hier und jetzt nicht hart sein? Technisch, nicht
„noch nicht geschafft".

Tragfähig: „Bibliothek X bietet keine API für die Prüfung",
„externe Abhängigkeit außerhalb unserer Kontrolle".
Nicht tragfähig: „dauert zu lange", „machen wir später".
-->

<…>

## Auflösungs-Trigger

<!--
Eine MESSBARE Bedingung, die ein anderer Mensch ohne Rückfrage als
erreicht beurteilen kann — die Schwelle ist das, was ein Sensor prüft,
der Wellen-Bezug nur der Roadmap-Anker.

Falsch: „wenn Zeit ist", „perspektivisch".
Richtig: „`internal/x`-Coverage ≥ 90 % in `make coverage-gate` ohne
Ausnahme", „sobald d-check das Modul `foo` liefert".
-->

<…>

## Geltungs-Konfiguration

<!--
Wo steht die Ausnahme im Gate-Tool? Die Stelle MUSS einen
`# CO-NNN`-Kommentar tragen — sonst ist sie im `make gates`-Output eine
stille Senkung ohne Begründung.
-->

| Datei | Zeile/Abschnitt | Wert |
|---|---|---|
| <…> | <…> | <…> |

## Verifikation (nach Auflösung)

<!--
Was muss grün sein, damit dieser Carveout verschwinden darf? Die
Auflösung ist ein `git mv` nach `done/` PLUS entfernte Gate-Ausnahme
PLUS grüner Lauf ohne sie — alle drei, sonst wirkt der Carveout
aufgelöst und liegt weiter im aktiven Verzeichnis.
-->

- [ ] Gate-Ausnahme aus `<datei>` entfernt
- [ ] `make gates` grün **ohne** die Ausnahme (Ausgabe in eine Datei, Exit-Code getrennt geprüft)
- [ ] Datei nach `docs/plan/carveouts/done/` verschoben, Status auf `Aufgelöst`
- [ ] Index in [`README.md`](README.md) aktualisiert
