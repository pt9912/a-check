# slice-155 — ADR-Index und Carveouts-README gegen v6.0.0-Form nachgezogen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Systematischer Abgleich aller a-check-Instanzdokumente gegen ihre v6.0.0-Templates
(Maintainer-Auftrag 2026-09-05) — die beiden letzten der acht Funde. [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Planungs-Harness-Doku, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

Zwei kleine Struktur-Lücken schließen, ohne den bestehenden, in beiden Fällen bereits reicheren
a-check-Inhalt zu verlieren.

## 2. `docs/plan/adr/README.md` — fehlende `## Konventionen`-Überschrift

Der Konventions-Absatz stand unbetitelt direkt nach dem Dokument-Titel, während die Ziel-Form
([`adr/README.template.md`](../../../../.harness/baseline/v6.0.0/templates/docs/plan/adr/README.template.md))
ihn unter `## Konventionen` führt. Überschrift ergänzt, Inhalt und Position unverändert — kein
neuer Fakt.

## 3. `docs/plan/carveouts/README.md` — fehlende `## Aktive/Aufgelöste Carveouts`-Container

Die Ziel-Form
([`carveouts/README.template.md`](../../../../.harness/baseline/v6.0.0/templates/docs/plan/carveouts/README.template.md))
führt zwei stehende Container-Abschnitte (Tabelle + „noch keine"), die den aktuellen Bestand
zeigen — a-check hatte stattdessen nur die tiefere Prosa-Begründung „Warum dieses Verzeichnis
heute leer ist". Beide Container ergänzt (`## Aktive Carveouts` mit leerer Tabelle,
`## Aufgelöste Carveouts` mit „noch keine"); die bestehende Begründung bleibt **zusätzlich**
stehen — sie beantwortet *warum*, die Container beantworten *was aktuell*.

## 4. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor.

## 5. DoD

- [x] `## Konventionen`-Überschrift in `adr/README.md` ergänzt, Inhalt unverändert (§2).
- [x] `## Aktive Carveouts`/`## Aufgelöste Carveouts` in `carveouts/README.md` ergänzt, bestehende
      Begründung erhalten (§3).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** letzte zwei der acht Funde aus dem systematischen v6.0.0-Template-Abgleich
geschlossen. Damit sind alle acht Funde bearbeitet — die drei bewusst nicht behobenen (ADR-Struktur
`## Verglichene Alternativen`/`## Re-Evaluierungs-Trigger`/`## Geschichte`, `welle-12/13-results.md`s
fehlender Register-Zeiger) bleiben als geprüfte, begründete a-check-Konvention bzw. historisch
korrekte Abweichung stehen — nicht als offener Rest.

**Lerneintrag — Form: benannte Spec-Lücke.** *„Kein Gegenstück im Template" und „driftet gegen
die Ziel-Form" sind zwei verschiedene Aussagen, und dieser Slice bestätigt beide Richtungen am
selben Dokumentenpaar: die fehlenden Überschriften waren echte Form-Lücken (behoben), während die
in [slice-150](../done/slice-150-roadmap-form-nachgezogen.md) bereits geprüften a-check-Erweiterungen
(z. B. die Beleg-Spalte im Drift-Log) keine sind.* Der systematische Abgleich, den der Maintainer
mit „sind wir mit der Umstellung durch?" anstieß, war nötig — die reaktive Prüfung einzelner
Hinweise (Carveout-Vorlage, Roadmap) hätte diese beiden kleineren Lücken nicht gefunden, weil
niemand gezielt danach gefragt hätte.

**Zwei beobachtbare Closure-Kriterien:**

1. `grep -c '^## Konventionen' docs/plan/adr/README.md` → `1`;
   `grep -c '^## Aktive Carveouts\|^## Aufgelöste Carveouts' docs/plan/carveouts/README.md` → `2`.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Keines notiert — beide Funde sind geschlossen, kein offener
Rest. Der systematische Template-Abgleich selbst (alle 26 Templates) ist mit diesem Slice
abgeschlossen; ob künftige Baseline-Sprünge dieselbe Vollständigkeits-Prüfung brauchen, ist eine
Wiederholungs-Frage für die nächste Migration, kein offenes Risiko dieser.

**Folge-Slices:** keine vergeben.

## 7. Sub-Area-Modus

Berührt: **Planungs-Harness** (`docs/plan/adr/`, `docs/plan/carveouts/`) — Greenfield: Form steht
in der vendored Vorlage, `doc-check` prüft Verweis-Auflösung.
