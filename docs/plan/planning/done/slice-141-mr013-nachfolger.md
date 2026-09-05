# slice-141 — [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md): [`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md)s Nachfolge-Eintrag (ADR-Vorlage `v6.0.0`)

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Folge-Slice aus [slice-137 §8](../done/slice-137-adaptions-durchgang-v600.md#8-closure-notiz)
(Etappe B): [`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md)s eigener Auflösungs-Trigger — „die nächste Baseline-Migration" — ist mit
[slice-135](../done/slice-135-regelwerk-v600-delta-analyse.md) eingetreten. [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Konventions-Pflege ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

[`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md) („ADR-Vorlage ist die vendored Fassung `v5.12.0`") löst nach `conventions/done/`, ein
neuer Eintrag [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md) trägt dieselbe Aussage für `v6.0.0` — dieselbe Kette wie [`MR-007`](../../../../harness/conventions/done/MR-007-adr-vorlagen-version.md) → [`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md).
Reine Versions-Referenz-Korrektur: `docs/plan/adr/adr.template.md` ist zwischen `v5.12.0` und
`v6.0.0` **gemessen unverändert** (`git diff v5.12.0 v6.0.0 -- lab/templates/docs/plan/adr/adr.template.md`,
leerer Diff).

## 2. Betroffene Module

`harness/conventions/` (neue Datei für [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md), [`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md) nach `done/`
verschoben, Pfad-Korrektur in der verschobenen Datei) und `harness/conventions.md`s Adaptions-Index
(Zeile aus *Aktive* nach *Aufgelöste Adaptionen*, neue Zeile für [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md), [`MR-007`](../../../../harness/conventions/done/MR-007-adr-vorlagen-version.md)s
„aufgelöst durch"-Zeiger auf den neuen `done/`-Pfad korrigiert). Eine Schicht.

## 3. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor.

## 4. Was bewusst nicht getan wird

- **[`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md) wird nicht inhaltlich verändert**, nur die Pfad-Tiefe nach dem Umzug nach `done/`
  korrigiert (`../conventions.md` → `../../conventions.md`, fünf Stellen) — dieselbe Disziplin wie
  beim `git mv` einer Slice-Datei: der Umzug zieht die Pfad-Berichtigung nach sich, ist aber keine
  inhaltliche Änderung ([`grundlagen-traceability.md` `v6.0.0`](../../../../.harness/baseline/v6.0.0/regelwerk/grundlagen-traceability.md)).
  Das `**Status:** Accepted`-Feld bleibt stehen, obwohl die neue Baseline-Vorlage es nicht mehr
  führt — [`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md) ist unter der alten Form entstanden und bleibt dabei.
- **Keine Überarbeitung der ID-Schema-Deklaration in [`MR-000`](../../../../harness/conventions.md#mr-000).** Genau die bliebe die saubere
  Auflösung ([`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md)s eigener Text), ist aber ein eigenständiger Entwurf, kein Nebenprodukt einer
  Versionskorrektur.

## 5. DoD

- [x] [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md) angelegt, `docs/plan/adr/adr.template.md`-Gleichheit zwischen den Ständen gemessen
      (nicht angenommen) und im Eintrag belegt; [`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md) nach `conventions/done/` verschoben,
      Pfad-Tiefe korrigiert, Inhalt sonst unverändert.
- [x] `conventions.md`s Adaptions-Index konsistent: [`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md) nur noch in *Aufgelöste Adaptionen*
      (Zeiger auf [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md)), [`MR-007`](../../../../harness/conventions/done/MR-007-adr-vorlagen-version.md)s Zeiger auf den neuen `done/`-Pfad korrigiert.
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md) als dritter Eintrag in der [`MR-007`](../../../../harness/conventions/done/MR-007-adr-vorlagen-version.md) → [`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md) → [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md)-Kette, mit
gemessener (nicht angenommener) Belegstelle für die Versionsgleichheit der ADR-Vorlage. Der
Adaptions-Index ist konsistent nachgezogen.

**Lerneintrag — Form: benannte Spec-Lücke.** *Ein `Rückbau-Kandidat`, der zwei Baseline-Migrationen
in Folge überlebt, ohne dass sein eigener Auflösungs-Trigger — die Überarbeitung von [`MR-000`](../../../../harness/conventions.md#mr-000)s
ID-Schema-Deklaration — je gezogen wurde, ist kein Einzelfall mehr, sondern ein Muster: der
naheliegende Trigger (Versionskorrektur bei jeder Migration) ist billiger als der saubere (Schema
überarbeiten), und billig gewinnt, solange niemand zählt.* [`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md) benannte das Muster für sich
selbst bereits vorsorglich; [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md) bestätigt es am zweiten Durchlauf. Noch kein
Beobachtungs-Registereintrag — zwei Vorkommen sind unter der Schwelle —, aber im [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md)-Eintrag
selbst benannt, damit ein dritter Durchlauf (nächste Baseline-Migration ohne Schema-Überarbeitung)
nicht wieder bei null anfängt.

**Zwei beobachtbare Closure-Kriterien:**

1. `git diff v5.12.0 v6.0.0 -- lab/templates/docs/plan/adr/adr.template.md` gegen einen frischen
   Klon von `pt9912/ai-harness-course` bestätigt den leeren Diff, den [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md) behauptet.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.**

- *Die ID-Schema-Überarbeitung in [`MR-000`](../../../../harness/conventions.md#mr-000) bleibt aus, ein dritter Durchlauf ist absehbar* —
  Ausgang: **Beobachtungs-Register**, neu angelegt unter
  [`BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md)
  (1×, `slice-141`; das erste Vorkommen — [`MR-007`](../../../../harness/conventions/done/MR-007-adr-vorlagen-version.md)→[`MR-013`](../../../../harness/conventions/done/MR-013-adr-vorlagen-version.md) — liegt vor der Verzeichnisform und ist im
  [`MR-017`](../../../../harness/conventions/MR-017-adr-vorlagen-version.md)-Eintrag benannt statt nachträglich als Evidence-Datei erfunden).

**Folge-Slices:** keine vergeben.

## 7. Sub-Area-Modus

Berührt: **Harness-Einstieg** (`harness/conventions.md`, `harness/conventions/`) — Greenfield.
