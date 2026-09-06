# slice-166 — `MR-017` durch generische ADR-Vorlagen-Referenz ablösen (`MR-020`)

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** [welle-14](welle-14-regelwerk-v610-migration.md).

**Bezug:** [slice-163](../done/slice-163-adaptions-durchgang-v610.md)
§4/§6 (Folge-Slice-Vorschlag), Maintainer-Wort 2026-09-06 ("Danach
[MR-017](../../../../harness/conventions.md#mr-017)/[MR-000](../../../../harness/conventions.md#mr-000)
und Etappe A").

**Berührte Spec-Stellen:** — *(keine)* — Harness-/Konventions-Änderung
ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim
Maintainer.

**Autor:** Claude (Sonnet 5). **Datum:** 2026-09-06.

---

## 1. Ziel

[`MR-017`](../../../../harness/conventions.md#mr-017)s
ausgelösten Auflösungs-Trigger ("die nächste Baseline-Migration") nicht
durch eine weitere Versions-`MR` bedienen (dritter Durchlauf desselben
Musters), sondern
[`MR-000`](../../../../harness/conventions.md#mr-000)s ADR-Vorlagen-Referenz
über eine neue, **permanent** auflösende Adaption
([`MR-020`](../../../../harness/conventions.md#mr-020))
generisch auf den jeweils aktuell vendorten Stand umstellen.

## 2. Definition of Done

- [x] [`MR-020`](../../../../harness/conventions.md#mr-020)
      angelegt (generischer Verweis auf `conventions.md` §Baseline statt
      fester Versionsnummer, `Auflösungs-Trigger: permanent`);
      [`MR-017`](../../../../harness/conventions.md#mr-017)
      nach `harness/conventions/done/` verschoben (reiner `git mv`, Inhalt
      unverändert), Adaptions-Tabellen in `harness/conventions.md`
      nachgezogen.
- [x] Alle Repo-Referenzen auf
      [`MR-017`](../../../../harness/conventions.md#mr-017)s
      alten Pfad (Datei direkt unter `harness/conventions/`) auf den neuen
      `done/`-Pfad nachgezogen (repo-weit, analog zum `slice-mv`-Nachzug bei
      Slices); `BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration`
      auf `verkörpert` gesetzt (der dritte Durchlauf des
      Versions-Bump-Musters ist damit ausgeblieben).
- [x] Unabhängiger Review über getrennten Kontext durchgeführt (Report
      unter `docs/reviews/`); der dabei gefundene Zählfehler (zehn statt
      zwölf betroffene Dateien) vor Abnahme korrigiert.
- [x] `make gates` grün.
- [x] `make verify` grün.
- [x] Jedes Risiko aus §6 trägt einen Ausgang.

## 3. Umsetzung

[`MR-000`](../../../../harness/conventions.md#mr-000) selbst wird **nicht**
verändert (Adaptions-Block-Disziplin: „an einem akzeptierten Eintrag wird
nichts nachträglich inhaltlich geändert"). Die Umformulierung lebt
vollständig in der neuen Adaption
[`MR-020`](../../../../harness/conventions.md#mr-020),
die [`MR-017`](../../../../harness/conventions.md#mr-017)
ablöst — derselbe Mechanismus, mit dem
[`MR-014`](../../../../harness/conventions.md#mr-014)/[`MR-015`](../../../../harness/conventions.md#mr-015)/[`MR-016`](../../../../harness/conventions.md#mr-016)
ihre Vorgänger abgelöst haben.

**Repo-weiter Referenz-Nachzug:** zehn Dateien verwiesen direkt auf den
alten Dateipfad von
[`MR-017`](../../../../harness/conventions.md#mr-017)
(nicht nur über den stabilen `harness/conventions.md#mr-017`-Anker, der
unverändert auflöst) — darunter mehrere bereits geschlossene `done/`-Slices
(`slice-141`, `slice-162`, `slice-163`, `slice-164`, `slice-165`) und
Beobachtungs-Belege. Alle wurden auf den neuen `done/`-Pfad umgeschrieben —
**nur der Verweis**, nicht der historische Inhalt: derselbe Umgang, den
`slice-141` bereits für den analogen
[`MR-013`](../../../../harness/conventions.md#mr-013)→[`MR-017`](../../../../harness/conventions.md#mr-017)-Übergang
vorgemacht hat („Vier `done/`-Slices bekamen ihre Linkziele auf
[`MR-013`](../../../../harness/conventions.md#mr-013)
nachgezogen").

## 4. Trigger

**Start** (`open` → `in-progress`): Maintainer-Freigabe (dieses Gespräch,
2026-09-06), WIP-Limit frei (`in-progress/` leer nach `slice-165`-Closure).

**Rückführungen:**

- `in-progress` → `next`: entfällt — Umfang ist bereits auf drei
  Liefer-Punkte geschnitten und passt.
- `in-progress` → `open`: falls der Maintainer die generische
  Baseline-Referenz ablehnt und eine andere Auflösung verlangt.

## 5. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz geschrieben.

## 6. Risiken und offene Punkte

- *Ein künftiger Verweis auf
  [MR-017](../../../../harness/conventions.md#mr-017)
  (z. B. in einer neuen Beobachtung)
  nennt versehentlich wieder den alten, aktiven Pfad statt `conventions/done/`*
  — **Ausgang:** gestrichen mit Begründung: `doc-check`s
  `id-unlinked`/`target-missing` fängt einen falschen Pfad sofort ab
  (gemessen: alle zehn betroffenen Referenzen mussten für ein grünes
  `make gates` korrigiert werden, s. §3).
- *Die generische Formulierung „die jeweils aktuell vendorte Fassung" in
  [MR-020](../../../../harness/conventions.md#mr-020)
  wird bei einer künftigen Migration nicht konsequent
  eingehalten (jemand vendort einen neuen Stand, vergisst aber
  `conventions.md` §Baseline zu aktualisieren)* — **Ausgang:** weiter
  offen → Beobachtungs-Register (kein bestehender Eintrag trifft genau
  diesen Fall — neu wäre er nur bei tatsächlichem Auftreten anzulegen,
  nicht vorsorglich).

## 7. Closure-Notiz

- **Was hat funktioniert:** der repo-weite `grep` auf den exakten alten
  Pfad-String vor dem `git mv` hat alle zehn Referrer zuverlässig
  gefunden — dieselbe Technik, die `make slice-mv` intern für
  Slice-Dateien automatisiert, hier von Hand für eine `MR`-Datei
  angewendet, weil es dafür kein eigenes Werkzeug gibt.
- **Was ging anders als geplant:** ursprünglich war nur an den
  `harness/conventions.md#mr-017`-Anker gedacht (der bleibt stabil); dass
  daneben zehn Dateien den **Datei-Pfad** direkt verlinken und beim `git
  mv` mitgezogen werden müssen, wurde erst durch den `grep`-Befund
  sichtbar. Das unabhängige Review fand zusätzlich einen eigenen
  Zählfehler: die erste Fassung dieses Slice (und der bereits erstellte
  Commit `c201058`) behaupteten „zwölf Dateien" — die tatsächliche,
  nachgezählte Zahl ist zehn. Im Slice-Text vor Abnahme korrigiert; die
  Commit-Message von `c201058` trägt die falsche Zahl weiter fort (keine
  Amend-Praxis in diesem Repo für bereits erstellte, unveröffentlichte
  Commits) — dieselbe Fehlerklasse wie
  [`BEO-GATE/cr-text-behauptet-statt-gemessen`](../observations/BEO-GATE/cr-text-behauptet-statt-gemessen/observation.md),
  hier bei einer repo-eigenen Zählung statt einer CR an ein fremdes
  Werkzeug.
- **Lerneintrag — Form: benannte Spec-Lücke.** *a-check hat kein Werkzeug
  analog zu `make slice-mv` für Adaptions-Dateien (`harness/conventions/MR-*.md`)
  — die Umbenennungs-Referenz-Pflege lief hier manuell per `grep`+`sed`.
  Bisher dreimal aufgetreten
  ([MR-007](../../../../harness/conventions.md#mr-007)→[MR-013](../../../../harness/conventions.md#mr-013),
  [MR-013](../../../../harness/conventions.md#mr-013)→[MR-017](../../../../harness/conventions.md#mr-017),
  jetzt [MR-017](../../../../harness/conventions.md#mr-017)→[MR-020](../../../../harness/conventions.md#mr-020)),
  aber jedes Mal manuell und ohne
  Selbsttest-Absicherung wie bei `tools/slice-mv.sh`. Kein akuter
  Sensor-Bedarf (die Fehlerklasse ist über `make doc-check` ohnehin
  fail-closed abgesichert — ein vergessener Nachzug wäre sofort rot),
  aber eine benannte Lücke für den Fall einer vierten Wiederholung.*
- **Beobachtungs-Register (`../observations/`):**
  [`BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md)
  auf `verkörpert` gesetzt (`seit slice-166`) — der Eintrag hatte bei 1×
  gestanden (die Verzeichnisform des Registers existierte beim ersten
  Durchlauf
  [MR-007](../../../../harness/conventions.md#mr-007)→[MR-013](../../../../harness/conventions.md#mr-013)
  noch nicht) und ist mit dieser Auflösung erledigt, ohne einen dritten
  Durchlauf zu benötigen.
- **Folge-Slices:** keine.
- **Risiken aus §6:** beide mit Ausgang — siehe §6.
- **Drei Paarungen:** verschoben auf die Closure von `welle-14` (dieser
  Slice trägt ein `**Welle:**`-Feld, [Modul 8](../../../../.harness/baseline/v6.2.0/regelwerk/modul-08-agentenrollen.md#rollen-sequenz-für-eine-welle)).

## 8. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Harness-Einstieg** (`harness/conventions.md`, neue `MR`-Datei,
repo-weiter Referenz-Nachzug in Planungs-/Beobachtungs-Dateien —
Referenz-Pflege zählt nicht als eigene Sub-Area-Berührung, da inhaltlich
nichts an diesen Dateien geändert wird, nur Linkziele), Greenfield,
Schwelle ≥ 2/3 erfüllt.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-HARNESS/` über die
Verzeichnisliste geprüft — 9 `offen` vor diesem Slice (unverändert seit
`slice-165`); `rueckbau-kandidat-ueberlebt-baseline-migration` davon wird
mit diesem Slice `verkörpert` (s. §7), die übrigen acht unverändert,
keiner erreicht 3×.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
