# MR-029 — ID-Schema-Deklaration als Ganzes (löst [`MR-020`](../conventions.md#mr-020) und [`MR-023`](../conventions.md#mr-023) ab)

- **Status:** Accepted
- **Datum:** 2026-09-29
- **Geltungsbereich:** [`MR-000`](../conventions.md#mr-000) §ID-Schema-Deklaration
- **Ersetzt-Baseline-Regel:** — *(keine — wie [`MR-020`](done/MR-020-adr-vorlage-generisch.md)/[`MR-023`](done/MR-023-id-schema-beobachtungs-kennung.md) korrigiert dieser Eintrag eine **Repo**-Aussage, keine Baseline-Regel; er ist generisch formuliert und überlebt Baseline-Migrationen.)*
- **Adaption (die vollständige, aktuelle Deklaration):**
  - **Funktionale Anforderungen:** `AC-FA-<BEREICH>-<NNN>`; das Bereichskürzel
    wird im Lastenheft §3 deklariert (bereits
    [`MR-002`](done/MR-002-id-schema-bereichskuerzel.md)).
  - **Nichtfunktionale Anforderungen:** `AC-QA-<NN>`.
  - **ADRs:** `ADR-NNNN`, vierstellig, chronologisch über den ADR-Index; die
    maßgebliche ADR-Vorlage ist **die jeweils aktuell vendorte Fassung**
    (bereits [`MR-020`](done/MR-020-adr-vorlage-generisch.md)).
  - **Konventions-Adaptionen:** `MR-NNN`.
  - **Carveouts:** `CO-<NNN>` (bisher ungenutzt).
  - **Slices:** `slice-<NNN>`, numerisch — a-checks deklarierte Form; der
    Kurs lässt die Form ausdrücklich als Repo-Deklaration
    ([`MR-000`](../conventions.md#mr-000) §ID-Schema-Deklaration).
  - **Beobachtungen:** Pfad-Form `BEO-<KUERZEL>/<slug>`; das Kürzel ist das
    Sub-Area-Kürzel aus der Modus-Deklaration (bereits
    [`MR-023`](done/MR-023-id-schema-beobachtungs-kennung.md)).
- **Begründung:** Die Deklaration war über Korrektur-Einträge gepflegt
  ([`MR-020`](done/MR-020-adr-vorlage-generisch.md),
  [`MR-023`](done/MR-023-id-schema-beobachtungs-kennung.md)) — jeder
  davon korrigierte eine Repo-Aussage statt einer Baseline-Regel und fiel damit
  unter den Fork-Test ( [`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../../docs/plan/planning/observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md),
  3×: slice-097, slice-162, slice-187). Dieser Eintrag fasst die Deklaration als
  Ganzes, **generisch** formuliert: Schwellen- und versionsfreie Formen, deren
  Gültigkeit von [`harness/conventions.md`](../conventions.md) §Baseline
  mitgepflegt wird — eine künftige Baseline-Migration erzeugt für dieses Feld
  keinen Korrektur-Eintrag mehr.
- **Auflösungs-Trigger:** permanent — die Formen sind generisch deklariert; eine
  inhaltliche Änderung einer Form wäre ein neuer Eintrag.
- **Löst ab:** [`MR-020`](../conventions.md#mr-020), [`MR-023`](../conventions.md#mr-023)
- **Ausgelöst durch:** Maintainer-Wunsch „ID-Schema-Deklaration überarbeiten"
  (slice-190) und [`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../../docs/plan/planning/observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
  (3×: slice-097, slice-162, slice-187).
