# slice-065 — Etappe F (2/3): Carveout-Ort und Diskrepanz-Trichter

**Status:** der Zustand ist das **Verzeichnis** dieser Datei
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5) — dieses Feld führt ihn bewusst **nicht** doppelt.
**Deckt:** **B-14** (Carveout-Ort fehlt) und **B-10** (Diskrepanz-Trichter ungenutzt) aus
[slice-048 §5](../done/slice-048-modul-delta-lesen.md).
**Bezug:** Etappe F, zweiter von drei Schnitten; 1/3 war
slice-057.
[Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**B-14 — ein deklariertes Werkzeug ohne Ort.** `CO-NNN` ist seit
[MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration)
als ID-Reihe deklariert („Carveouts: `CO-NNN` (bisher ungenutzt)"), aber es gibt weder Verzeichnis
noch Vorlage. Dieselbe Klasse wie **B-18** (`next/` war deklariert und existierte nicht) — und
dort war die Antwort, den Ort anzulegen, nicht die Deklaration zurückzunehmen.

**B-10 — der Trichter ist ungenutzt.** `modul-07` trennt drei Werkzeuge über **zwei sequenzielle
Fragen** (Granularität *vor* Temporalität): Cluster ⇒ BF-Sub-Area-Markierung; einzelne Diskrepanz
mit erreichbarem Trigger ⇒ Carveout; Trigger nie erreichbar ⇒ permanente ADR. a-check kennt
faktisch nur ADRs und `MR-*`.

**Bestandsmessung — es gibt heute null Carveouts, und das ist kein Versäumnis.** Jede bestehende
Ausnahme fällt begründet in ein anderes Werkzeug:

| Ausnahme | Trichter-Ergebnis |
|---|---|
| `.golangci.yml` §exclusions (`cyclop` für table-driven Tests u. a.) | **permanent** — Design-Aussage, kein Übergang; getragen von [ADR-0005](../../adr/0005-lint-profil.md) |
| `.d-check.yml` `scan.ignore: .harness/baseline/**` | **permanent** — Fremdtext, getragen von [MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert) |
| `.d-check.yml` `exempt-paths` für ADR 0001–0020 | **permanent** — Immutabilität, getragen von [MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung) |
| `SLICE_FORM_FROM=52`, AC-Grandfathering, `commit-scope-check`-Stichtag | **kein Carveout** — `modul-07` schließt bootstrap-aware Gates ausdrücklich aus: sie stufen das *Gate*, sie lösen keine Diskrepanz |
| `regelwerk-check`-Integrität nicht im `gates`-Aggregat | **kein Carveout** — kein Gate wurde ausgenommen; `modul-13` führt das Target selbst als Muster „vorhanden, aber nicht als Gate behauptet" |

Der Ort entsteht also **leer** — wie `next/` in slice-053. Ein Ort ohne Inhalt ist ehrlich; eine
ID-Reihe ohne Ort ist eine stille Setzung.

## 2. Betroffene Module

- `docs/plan/carveouts/README.md` — neu: Zweck, **Trichter**, Dateikonvention, Pflichtfelder,
  Auflösungs-Mechanik, aktueller Bestand (B-14, B-10).
- `docs/plan/carveouts/carveout.template.md` — neu, aus der vendored Vorlage **übersetzt**
  (Verweise auf a-checks Orte statt auf Kurs-Pfade).
- [`AGENTS.md`](../../../../AGENTS.md) §5 — ein Satz plus Verweis, damit der Trichter im
  Agenten-Einstieg auffindbar ist.

**Zwei Schichten:** Planungs-Doku und Harness-Briefing.

## 3. Auszuführende Gates

`make gates` und `make verify`, Ausgabe je in eine Datei, Exit-Code getrennt geprüft. Kein neuer
Sensor — der Carveout-Mechanismus wird von `modul-07` über den **Audit-Slice** auditiert, und der
hängt am Wellen-Closure (**B-13**), das erst 3/3 anlegt. Ausdrücklich benannt statt
stillschweigend ausgelassen.

## 4. Was bewusst nicht getan wird

- **Kein erfundener Carveout.** Der Ort bleibt leer, weil die Messung leer ist. Einen Bestandsfall
  zum Carveout umzudeuten, nur damit das Verzeichnis nicht leer aussieht, wäre die Umkehrung des
  Trichters.
- **Kein Audit-Slice `SL-CO-AUDIT-<welle>`.** Er setzt eine auditierbar geschlossene Welle voraus
  (B-13), die a-check bis heute nicht hat — Gegenstand von Etappe F 3/3.
- **Keine Rücknahme der `CO-NNN`-Deklaration** aus [MR-000](../../../../harness/conventions.md#mr-000--baseline-aussage-inkl-id-schema-deklaration): die Disziplin des Adaptions-Blocks
  verbietet inhaltliche Änderungen an akzeptierten Einträgen, und die Reihe ist ab jetzt nutzbar.

## 5. DoD

- [x] `docs/plan/carveouts/` existiert mit Index-README und übersetzter Vorlage; die sechs
      Pflicht-Header-Felder und die Auflösung per `git mv` sind benannt (B-14).
- [x] Der Trichter steht repo-lokal mit beiden Fragen in der richtigen Reihenfolge und zeigt je
      Ergebnis auf a-checks realen Ort (Carveout-Verzeichnis · Modus-Tabelle in
      `conventions.md` · ADR-Verzeichnis) (B-10).
- [x] Der aktuelle Bestand ist als **leer** ausgewiesen, mit der Einordnung jeder bestehenden
      Ausnahme; `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt
      geprüft.

## 6. Closure-Notiz

**Geliefert:** `docs/plan/carveouts/` mit Index-README und übersetzter Vorlage; der
Diskrepanz-Trichter steht repo-lokal — im README ausführlich, in
[`AGENTS.md`](../../../../AGENTS.md) §5 als auffindbarer Dreizeiler. Der Bestand ist gegen den
Trichter geprüft und als **leer** ausgewiesen, mit Einordnung jeder bestehenden Ausnahme.

**Lerneintrag — Form: geschärfte Regel.**
> **Ein leerer Ort ist ein Ergebnis, kein Versäumnis — solange die Leere gemessen und nicht
> behauptet ist.** Die naheliegende Sorge bei B-14 war, Bürokratie für einen Fall aufzubauen, den
> es nicht gibt. Die Messung dreht das um: fünf bestehende Ausnahmen wurden durch den Trichter
> geschickt, und **jede** landete begründet woanders — drei permanent ([ADR-0005](../../adr/0005-lint-profil.md), [MR-005](../../../../harness/conventions.md#mr-005--referenzmatrix-intra-spec-richtung--adrslice-disziplin-d-check-angleichung), [MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)),
> zwei sind bootstrap-aware Gates, die `modul-07` ausdrücklich ausschließt. *Weil* dieser Weg
> dokumentiert ist, ist das leere Verzeichnis eine Aussage („geprüft, nichts gefunden") statt einer
> Lücke. Prüfsatz für jeden künftigen Mechanismus ohne Bestand: *erst den Bestand durch das
> Werkzeug schicken, dann entscheiden, ob das Werkzeug gebraucht wird — die Zuordnung ist das
> Ergebnis, nicht die Zahl.*

**Zwei beobachtbare Closure-Kriterien:**

1. `make gates` und `make verify` grün auf demselben Stand (je Exit 0) — belegt.
2. `docs/plan/carveouts/` enthält README und Vorlage; das README nennt **beide** Trichter-Fragen
   in der Reihenfolge Granularität → Temporalität, führt je Ergebnis den realen Zielort und listet
   alle fünf geprüften Bestandsausnahmen mit ihrer Einordnung.

**Was der Ort noch nicht hat:** den Audit-Slice `SL-CO-AUDIT-<welle>` aus `modul-07`. Er setzt
eine auditierbar geschlossene Welle voraus (**B-13**), die a-check nicht hat — Gegenstand von
Etappe F 3/3. Im README steht das als offener Punkt, statt einen Audit zu behaupten, den niemand
fährt. Solange es null Carveouts gibt, hat er ohnehin keinen Gegenstand.

**Folge-Slices:** slice-066 (Etappe F 3/3 — Wellen-Closure B-13 und Rollen-Übergaben B-9); dort
wird der Carveout-Audit-Punkt Teil der Prozedur.

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
