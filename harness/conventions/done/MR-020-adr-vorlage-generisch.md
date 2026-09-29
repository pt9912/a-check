# MR-020 — ADR-Vorlagen-Referenz zeigt generisch auf den vendorten Stand (löst [`MR-017`](../conventions.md#mr-017) permanent ab)

- **Status:** Accepted
- **Datum:** 2026-09-06
- **Geltungsbereich:** [`MR-000`](../conventions.md#mr-000) §ID-Schema-Deklaration, Zeile zu
  `ADR-NNNN`
- **Ersetzt-Baseline-Regel:** — *(keine — wie schon bei [`MR-017`](../conventions/done/MR-017-adr-vorlagen-version.md)/[`MR-013`](../conventions/done/MR-013-adr-vorlagen-version.md)/[`MR-007`](../conventions/done/MR-007-adr-vorlagen-version.md)
  korrigiert dieser Eintrag eine **Repo**-Aussage, keine Baseline-Regel)*.
- **Adaption:** [`MR-000`](../conventions.md#mr-000) deklariert `ADR-NNNN` als vierstellig „gemäß
  Kurs-ADR-Vorlage `v1.3.0`". Ab sofort gilt statt einer fest benannten Versionsnummer: die
  maßgebliche ADR-Vorlage ist **die jeweils aktuell vendorte Fassung** — siehe
  [`harness/conventions.md` §Baseline](../conventions.md#baseline), das ohnehin bei jeder
  Baseline-Migration als einzige Quelle für den aktuellen Stand aktualisiert wird. **Das Schema
  selbst ändert sich nicht** — vierstellig, chronologisch über den ADR-Index.
- **Begründung:** Drei aufeinanderfolgende Baseline-Migrationen haben denselben Ablauf erzeugt —
  [`MR-007`](../conventions/done/MR-007-adr-vorlagen-version.md)→[`MR-013`](../conventions/done/MR-013-adr-vorlagen-version.md)
  (`v5.12.0`), [`MR-013`](../conventions/done/MR-013-adr-vorlagen-version.md)→[`MR-017`](../conventions/done/MR-017-adr-vorlagen-version.md)
  (`v6.0.0`), und mit `v6.1.0`/`v6.2.0` wäre ein dritter Durchlauf
  ([`MR-017`](../conventions/done/MR-017-adr-vorlagen-version.md)→„MR-neu") fällig geworden
  ([slice-163](../../docs/plan/planning/done/slice-163-adaptions-durchgang-v610.md) §4) — genau das
  Muster, das [`MR-017`](../conventions/done/MR-017-adr-vorlagen-version.md) selbst als
  „Beobachtungs-Register-würdig, träte es ein drittes Mal auf" benannt hatte
  ([`BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration`](../../docs/plan/planning/observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md),
  seit `slice-141` bei 1× geführt). Diese Adaption fährt die **saubere** Auflösung, die
  [`MR-017`](../conventions/done/MR-017-adr-vorlagen-version.md) selbst als Alternative genannt
  hatte: die ID-Schema-Deklaration referenziert den Stand generisch statt mit einer Versionsnummer,
  die bei jeder künftigen Migration erneut veraltet. Damit entsteht bei der nächsten
  Baseline-Migration **kein** neuer `MR`-Eintrag mehr für dieses Feld — der generische Verweis
  bleibt gültig, solange `conventions.md` §Baseline selbst mitgepflegt wird (was jede Migration
  ohnehin verlangt).
- **Auflösungs-Trigger:** permanent — der generische Verweis überlebt jede künftige
  Baseline-Migration ohne erneute Anpassung.
- **Löst auf:** [`MR-017`](../conventions/done/MR-017-adr-vorlagen-version.md)
- **Ausgelöst durch Baseline-Stand:** `v6.2.0` (Kurs-Welle 119, 2026-09-05); vendored bleibt zum
  Zeitpunkt dieses Eintrags `v6.0.0`.
