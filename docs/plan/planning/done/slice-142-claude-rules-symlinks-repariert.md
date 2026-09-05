# slice-142 — `.claude/rules/`-Symlinks nach dem `v5.12.0`-Fall repariert

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Fund 2026-09-05 nach [slice-140](../done/slice-140-baseline-v5120-fall.md):
vier Symlinks unter `.claude/rules/` zeigten weiterhin auf `.harness/baseline/v5.12.0/…` — einen
Baum, den slice-140 entfernt hatte. `make gates` hatte den Bruch nicht gemeldet.
[Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Harness-Wartung ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

Vier defekte Symlinks reparieren, und benennen, warum die gesamte v5.12.0→v6.0.0-Migration
(slice-135…141) sie nicht fand.

## 2. Betroffene Module

`.claude/rules/{modul-01-entwicklungszyklus,modul-05-planning-harness,modul-06-roadmap,modul-08-agentenrollen}.md`
— vier Symlinks, Ziel-Segment `v5.12.0` → `v6.0.0`.

## 3. Warum die Migration sie nicht fand — zwei unabhängige Lücken, dieselbe Wirkung

`.claude/rules/` entstand mit [`afca439`](../../../../CHANGELOG.md) ([`MR-006`](../../../../harness/conventions/done/MR-006-baseline-vendored.md)-Beleg, „vier Symlinks auf
die vendored Baseline v5.12.0 … damit die per-Session geladenen Regelwerk-Abschnitte aus der
committeten Quelle kommen statt aus einer zweiten Kopie") — ein **Claude-Code-natives** Verzeichnis,
das bei jeder Sitzung automatisch in den Kontext geladen wird. Kein Dokument dieses Repos beschreibt
das als „Live-Referenz auf die Baseline", die ein Baseline-Bump nachziehen muss.

1. **Grep fand sie nicht.** Die Migration hat ihre eigene Vollständigkeit wiederholt über
   `grep -rln "v5\.12\.0"` gemessen ([slice-135 §1](../done/slice-135-regelwerk-v600-delta-analyse.md#1-ist-stand),
   [slice-140 §3](../done/slice-140-baseline-v5120-fall.md#3-warum-jetzt-sicher)). `grep` folgt
   Symlinks standardmäßig zur **Zieldatei** — ein Symlink, dessen *Pfad* die Versionsnummer trägt,
   meldet sich damit nie über den Dateiinhalt, sondern nur über `find -type l` oder `readlink`,
   wonach nirgends gesucht wurde.
2. **`make gates` prüft keine Symlinks.** Bestätigt durch Gegenprobe: mit den defekten Symlinks
   (`git stash` auf `.claude/rules/`) meldete `make gates` **Exit 0** — `d-check`s `links`-Modul
   prüft Markdown-Links (`[Text](Ziel)`), keine Dateisystem-Symlinks; kein anderes Ziel deckt sie.

Beide Lücken träfen jeden künftigen Baseline-Bump gleichermaßen — die zweite ist die
grundsätzlichere, weil sie auch unabhängig von einer bestimmten Such-Disziplin gilt.

## 4. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor — siehe §6 (Registereintrag statt
Sofortbau).

## 5. Was bewusst nicht getan wird

- **Kein neuer `d-check`-Sensor für Symlink-Ziele.** Ein einziger Vorfall; ein Sensor jetzt wäre
  Vorratsbau vor Bedarf. Der Fund wandert ins Beobachtungs-Register (§6).
- **Keine Erweiterung des Migrations-Verfahrens** (etwa `find -type l` als festen Schritt). Dieselbe
  Begründung — ein Vorkommen ist kein Muster.

## 6. DoD

- [x] Alle vier Symlinks lösen auf `v6.0.0` auf (`find . -type l` bestätigt keine dangling Links
      im Repo).
- [x] Gegenprobe geführt: `make gates` mit den defekten Symlinks Exit 0 — belegt die Lücke, statt
      sie zu behaupten.
- [x] `BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft` angelegt; `make gates` und `make verify`
      grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

**Geliefert:** vier reparierte Symlinks, plus die **gemessene** Bestätigung, dass `make gates` sie
nicht erwischt hätte (nicht nur vermutet) — Gegenprobe mit den defekten Originalen vor der Reparatur.

**Lerneintrag — Form: neuer Sensor (benannt, nicht gebaut) + geschärfte Regel.** *Eine
Migrations-Vollständigkeits-Prüfung, die auf Dateiinhalt sucht (`grep`), ist blind für Referenzen,
die im Dateisystem selbst liegen, nicht im Text — ein Symlink-Ziel ist so eine Referenz, und die
Blindheit ist strukturell, nicht durch ein präziseres Suchmuster behebbar.* Die gesamte Migration
(slice-135…141) durchsuchte wiederholt nach `v5.12.0` im Text, kein einziger Durchlauf nach
Symlinks. Der Fund kam vom Maintainer, nicht aus einem eigenen Schritt. Der zweite Teil der Lehre
ist schärfer: **kein Gate hätte es gefangen**, auch mit korrektem Suchmuster nicht — `d-check`s
Geltungsbereich ist Markdown-Linkpflicht, nicht Dateisystem-Integrität. Ein Sensor „jeder Symlink im
Repo löst auf" wäre klein und allgemein (kein Bezug zur Baseline nötig), ist aber noch nicht gebaut
— ein einzelnes Vorkommen rechtfertigt das nicht, siehe §5. *Weil* eine Suchmethode ihre eigene
Reichweite nicht kennt: `grep` fand nichts und meldete Vollständigkeit, wo tatsächlich eine
Referenzklasse komplett außerhalb seiner Sichtbarkeit lag.

**Zwei beobachtbare Closure-Kriterien:**

1. `find . -type l -not -path './.git/*' | xargs -I{} test -e {}` — kein Symlink im Repo ist
   dangling (Exit 0 je Datei).
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.**

- *Kein Sensor gegen dangling Symlinks* — Ausgang: **Beobachtungs-Register**, neu angelegt unter
  [`BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft`](../observations/BEO-GATE/symlink-ziel-nach-baseline-bump-ungeprueft/observation.md)
  (1×, `slice-142`).

**Folge-Slices:** keine vergeben.

## 8. Sub-Area-Modus

Berührt: **Gate-/Werkzeug-Schicht** (`.claude/rules/` — Teil der Durchsetzungsschicht, per
[slice-138](../done/slice-138-sub-area-kuerzel.md) unter `Gate-/Werkzeug-Schicht` gefasst) —
Greenfield.
