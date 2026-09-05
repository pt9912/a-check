# slice-151 — „Offene Slices ohne Welle" aus der Roadmap entfernt

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Korrektur 2026-09-05 an [slice-150](../done/slice-150-roadmap-form-nachgezogen.md) §2
— die Tabelle „Offene Slices ohne Welle" existiert im Baseline-Template nicht.
[Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Planungs-Harness-Pflege, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

Die eigene Fehleinschätzung aus [slice-150 §2](../done/slice-150-roadmap-form-nachgezogen.md#2-vier-reale-befunde-gemessen-per-abschnitts-für-abschnitt-vergleich)
korrigieren: dort hatte ich die Tabelle „Offene Slices ohne Welle" geprüft und als „kein
Baseline-Gegenstück, aber deckt einen echten a-check-Fall" **behalten**. Der Maintainer wies zu
Recht zurück — die Prüfung war unvollständig.

## 2. Was die vorige Prüfung übersah

[`modul-06-roadmap.md`](../../../../.harness/baseline/v6.0.0/regelwerk/modul-06-roadmap.md)
§Wann Arbeit eine Welle braucht sagt es **wörtlich**: *„Wellenlose Arbeit erscheint nicht in der
Roadmap — weder beim Start noch beim Abschluss. Ihr Zustand ist die Verzeichnis-Position (Modul
5); `ls docs/plan/planning/in-progress/` beantwortet ‚was läuft gerade' autoritativ und ohne
Pflegeaufwand … Eine Reihenfolge einzelner Slices kennt der Harness nicht."*

Die entfernte Tabelle widersprach dem in **zwei** Punkten:

1. Sie listete `slice-013`/`slice-045` **in der Roadmap**, obwohl beide wellenlos sind — genau das,
   was die Baseline ausschließt.
2. Ihre eigene Bildunterschrift — „ihre Trigger stehen in ihrem eigenen `§0`" — zitierte
   unbeabsichtigt die Begründung, warum sie selbst **nicht** hätte existieren sollen: Wenn der
   Trigger schon im Slice steht, braucht die Roadmap keine zweite Kopie.

**Zusätzlich redundant:** Beide Slices standen bereits vollständig in *Nächste Wellen* — mit Trigger,
Link und zugeordneter künftiger Welle. Die entfernte Tabelle duplizierte das, in einer schwächeren
Form (kein Trigger-Text, nur „siehe *Nächste Wellen*").

## 3. Umsetzung

Die Tabelle samt einleitendem Satz und der Fremdrepo-Begründung aus `## Offene Wellen` entfernt.
Der Abschnitt endet jetzt direkt nach dem Ruhe-Marker `Nichts in Arbeit.` — deckungsgleich mit der
Baseline-Ziel-Form. Kein Informationsverlust: `slice-013`/`slice-045` bleiben unverändert Zeilen in
*Nächste Wellen*, ihre §0-Abschnitte in den Slice-Dateien selbst unverändert.

Keine anderen Verweise auf die Tabelle oder ihren Anker im lebendigen Bestand (`grep` repo-weit,
einzige Fundstelle war [slice-150](../done/slice-150-roadmap-form-nachgezogen.md) selbst — frozen
`done/`-Prosa, nicht angetastet).

## 4. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor.

## 5. DoD

- [x] Baseline-Zitat wörtlich gegen den entfernten Block gehalten, nicht nur „driftet vermutlich"
      behauptet (§2).
- [x] Block entfernt, Informationsgehalt in *Nächste Wellen* bereits vorhanden verifiziert, keine
      anderen Verweise betroffen (§3).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** die Roadmap ist jetzt auch in diesem letzten Punkt deckungsgleich mit der
`v6.0.0`-Ziel-Form — wellenlose Arbeit erscheint nicht mehr in ihr.

**Lerneintrag — Form: geschärfte Regel.** *„Kein Baseline-Gegenstück, deckt aber einen echten
Fall" ist keine hinreichende Begründung, um eine Abweichung zu behalten — sie prüft nur, ob das
Vorbild dagegen ist, nicht, ob eine ANDERE, ausdrücklich formulierte Regel dagegen ist.* Der
Formvergleich in [slice-150](../done/slice-150-roadmap-form-nachgezogen.md) hielt die Roadmap
gegen das **Template**, aber zog `modul-06-roadmap.md` selbst nicht mehr heran — genau die Quelle,
deren Wortlaut die entfernte Tabelle unbeabsichtigt selbst zitierte. *Weil* ein Template die Regel
nur **verkörpert**, nicht **ersetzt**, prüft eine Formkonformitäts-Prüfung künftig gegen **beide**:
das Template für die Struktur, das Regelwerk-Kapitel für ausdrücklich formulierte Verbote, die im
Template nur implizit (durch Abwesenheit) stehen.

**Zwei beobachtbare Closure-Kriterien:**

1. `grep -c "Offene Slices ohne Welle" docs/plan/planning/in-progress/roadmap.md` → `0`.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Keines notiert — der Fund ist mit der Entfernung geschlossen.

**Folge-Slices:** keine vergeben.

## 7. Sub-Area-Modus

Berührt: **Planungs-Harness** (`docs/plan/planning/in-progress/roadmap.md`) — Greenfield: Form
steht in der vendored Vorlage und im Regelwerk-Kapitel, `doc-check` prüft Verweis-Auflösung.
