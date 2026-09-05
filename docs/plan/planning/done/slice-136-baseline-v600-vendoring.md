# slice-136 — Etappe A: Baseline `v6.0.0` committet vendoren

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Etappe **A** aus [slice-135 §6](../done/slice-135-regelwerk-v600-delta-analyse.md#6-vorschlag-drei-etappen),
per Maintainer-Wort gezogen 2026-09-05 („Leg los mit Etappe A"). Präzedenz:
[slice-094](../done/slice-094-baseline-v5120-vendoring.md), dieselbe Etappe A für den vorigen
Sprung. [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Harness-/Konventions-Änderung ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Was Etappe A tut — und was nicht

Sie materialisiert den neuen Stand **neben** dem alten und zieht die drei deklarierenden Stellen
nach. **Nichts weiter:** kein Adaptions-Durchgang (Etappe B), keine Beobachtungs-Register-Migration
(Etappe C).

**Das alte Verzeichnis bleibt stehen** — Vorschrift der Baseline selbst (`modul-02`
§Freshness-Audit): die neue Referenz-Form ist die Vergleichsgrundlage für den Form-Review, und das
alte Verzeichnis fällt erst, wenn der Review durch ist. `make regelwerk-check` ist darauf
vorbereitet: es wählt den höchsten Stand per `sort -V` und weist den ungeprüften Rest namentlich
aus (slice-049, bestätigt in slice-094).

**Template-Provenienz-Zeiger bleiben unverändert auf `v5.12.0`** — dieselbe Disziplin wie in
slice-094 §4: sie zeigen auf die Form, aus der das jeweilige a-check-Artefakt **tatsächlich**
entstanden ist, und mehrere der geänderten Templates (`slice.template.md`,
`welle-results.template.md`, `harness/conventions/MR-NNN-titel.template.md`) tragen inhaltlich
bereits die neue Beobachtungs-Register-Form (Etappe C), die a-check noch nicht adoptiert hat. Sie
jetzt umzuhängen würde eine Konformität behaupten, die niemand geprüft hat. Betroffen:

- `AGENTS.md` §5 „Slice-Form" — Pointer auf `slice.template.md`.
- `harness/README.md` §Sensors-Kommentar — Pointer auf `harness/README.template.md`.
- `docs/plan/carveouts/carveout.template.md`, `docs/reviews/README.md`,
  `.harness/skills/reviewer.md` — je ein Provenienz-Zeiger, unverändert (jeweils inhaltlich
  ungeändert zwischen `v5.12.0`/`v6.0.0`, aber die Bewertung selbst ist Etappe-B-Lesearbeit, kein
  Diff dieser Etappe).
- Die `Ersetzt-Baseline-Regel`-Anker in den sechs aktiven `MR-<NNN>`-Dateien und in `conventions.md`
  §Aktive Adaptionen — Gegenstand des Adaptions-Durchgangs (Etappe B).

## 2. Betroffene Module

- `.harness/baseline/v6.0.0/{regelwerk,templates}/` + `SHA256SUMS` — neuer vendored Baum, 53
  Dateien (gegenüber 51 bei `v5.12.0` — drei neue Template-Dateien, eine gelöschte).
- Stand-Deklaration an ihren **drei** Stellen: `harness/conventions.md` §Baseline und
  §Adoptierte Konventions-Quellen, `AGENTS.md` §1, `harness/README.md` §Guides (Zeile 46, nicht
  der Sensors-Kommentar — siehe §1).

Eine Schicht (Harness-Deklaration), plus ein Verzeichnis Fremdtext.

## 3. Auszuführende Gates

`make regelwerk-check` (kein Gate, Wartung: Integrität fail-closed **und** die Gegenrichtung
Baum → Manifest — meldet unaufgefordert die zwei vendored Stände und den ungeprüften), dann
`make gates`, zum Abschluss `make verify`.

**Kein neuer Sensor**, also keine Negativ-Probe. Der Integritätsnachweis ist reproduzierbar
(`sha256sum -c`), zusätzlich unabhängig von der Herkunft belegt (slice-135 §2: Zip-Digest gegen
GitHub-API, Baum-Identität gegen `d-check`).

## 4. Was bewusst nicht getan wird

- **Kein Adaptions-Urteil.** Welcher der 16 aktiven `MR`-Einträge durch `v6.0.0` welchen Ausgang
  nimmt (bleibt gültig / teilweise überholt / widerspricht / …), ist Etappe B — hier wird nichts
  vorweggenommen.
- **Keine Beobachtungs-Register-Migration.** Die 33 Tabellenzeilen bleiben Tabellenzeilen — Etappe
  C.
- **Historische Aussagen werden nicht umgeschrieben.** Closure-Notizen in `done/` und
  Review-Reports nennen `v5.12.0` als das, was damals galt. Das ist korrekt und bleibt.

## 5. DoD

- [x] `.harness/baseline/v6.0.0/{regelwerk,templates}/` + `SHA256SUMS` aus dem Release-ZIP
      materialisiert (Provenienz bereits in slice-135 §2 unabhängig belegt); `make regelwerk-check`
      grün, weist `v5.12.0` als ungeprüften Rest namentlich aus.
- [x] Die Stand-Deklaration nennt an allen drei Stellen `v6.0.0`; Template-Provenienz-Zeiger bleiben
      bewusst auf `v5.12.0` (§1) — Beleg: Diff.
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** 53 Dateien unter `.harness/baseline/v6.0.0/` neben den 51 des Vorgängers, ein
reproduzierbares Integritäts-Manifest, und die Stand-Deklaration an ihren drei Stellen auf
`v6.0.0` — bei bewusst unverändert gelassenen Template-Provenienz-Zeigern, weil ein Teil der
geänderten Templates bereits die neue, noch nicht adoptierte Beobachtungs-Register-Form trägt.

**Lerneintrag — Form: geschärfte Regel.** *Ein Template-Provenienz-Zeiger und ein
Stand-deklarierender Index-Pointer sehen in der Grep-Trefferliste identisch aus („referenziert
`.harness/baseline/<tag>/…`"), sind aber zwei verschiedene Zusagen — der eine sagt „hier ist der
aktuell adoptierte Stand", der andere „von hier stammt dieses konkrete Artefakt tatsächlich" — und
nur der erste darf bei einem Re-Vendoring blind mitwandern.* Konkret: `AGENTS.md` §5 (Pointer auf
`slice.template.md`) sieht identisch aus wie `AGENTS.md` §1 (Pointer auf `regelwerk/README.md`),
aber `slice.template.md` änderte sich inhaltlich in Richtung einer Form, die a-check noch nicht
führt (`../observations/` statt `../observations.md`) — ein blinder Bump hätte Agenten ab sofort
eine Slice-DoD-Zeile kopieren lassen, die auf einen nicht existierenden Pfad zeigt. slice-094 hatte
diese Unterscheidung bereits für den vorigen Sprung getroffen (§4, „5 Stellen"); dieser Slice
bestätigt sie am unabhängig gemessenen Delta von slice-135 erneut, statt sie als gegeben zu
übernehmen — dieselbe Disziplin wie bei jedem übernommenen Vorbild (slice-135s eigener
Lerneintrag). *Weil* die beiden Pointer-Arten ohne diese Unterscheidung ununterscheidbar sind und
ein Re-Vendoring sie sonst pauschal gleich behandelt.

**Zwei beobachtbare Closure-Kriterien:**

1. `bash tools/regelwerk-check.sh` meldet „Integritaet ok — 53 Datei(en)" für `v6.0.0` und weist
   `v5.12.0` als ungeprüft aus.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.**

- *Die sechs Template-Provenienz-Zeiger sind gegen `v6.0.0` noch nicht bewertet* — Ausgang:
  **Folge-Slice**, Etappe B aus slice-135 §6.
- *`.harness/baseline/v5.12.0/` bleibt bis zum Form-Review vendored* — Ausgang: **gestrichen mit
  Begründung**: kein Risiko im Sinn der Dreier-Menge (nichts kann hier „eintreten"), sondern eine
  von der Baseline selbst vorgeschriebene, beabsichtigte Übergangsphase (§1) — dieselbe
  Fehlklassifikation wie bei [`BEO-025`](../observations.md), die Etappe B/C von selbst beendet.

**Folge-Slices:** keine vergeben. Etappe B und C sind in slice-135 §6 vorgeschlagen und brauchen
eigene Kennungen bei ihrer eigenen Eröffnung.

## 7. Sub-Area-Modus

Berührt wird ausschließlich **Vendored Baseline** (`.harness/baseline/`, kein Modus — externer
Fremdtext, [`MR-006`](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert))
und **Harness-Einstieg** (`AGENTS.md`, `harness/`) — Greenfield.
