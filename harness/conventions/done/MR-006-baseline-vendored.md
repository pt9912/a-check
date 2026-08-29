# MR-006 — Baseline committet vendored statt per URL referenziert

- **Datum:** 2026-07-25
- **Geltungsbereich:** [`AGENTS.md`](../../../AGENTS.md) §1, [§Baseline](../../conventions.md#baseline),
  [`harness/README.md` §Guides](../../README.md#guides-feedforward-quellen)
- **Adaption:** Provenienz/Konkretisierung, **keine** inhaltliche Abweichung vom
  `v3.5.2`-Default — im Gegenteil: sie **stellt ihn her**. Regelwerk *und* Templates der
  Baseline liegen **committet vendored** unter
  `.harness/baseline/v3.5.2/{regelwerk,templates}/`, materialisiert aus dem
  self-contained `lab-regelwerk.zip` des Releases, mit `SHA256SUMS` als
  Integritätsmanifest. Bis dahin referenzierte dieses Repo die Baseline **nur per URL**
  (Modell `v1.3.0`).
- **Begründung:** Weil `regelwerk/` und `templates/` **parallel** vendored liegen, lösen die
  `../templates/…`-Verweise des Regelwerks („so sieht das Artefakt aus") **netzlos lokal**
  auf — ein URL-Verweis kann das nicht. Der Nachschlag wird damit reproduzierbar über den
  `<tag>` und unabhängig von Netz und Login; die Kontext-Hygiene bleibt gewahrt, weil pro
  Abschnitt eine Datei geladen wird statt des ganzen Bundles. Real-Vorbild in der Flotte:
  `ai-harness-init` (dortige `MR-007`, vendored `v3.5.1`). <!-- d-check:ignore (fremde MR-Kennung des Repos ai-harness-init, nicht a-checks) -->
- **Abgrenzung:** Der vendored Baum ist **externer, unveränderter Fremdtext** mit eigenen
  Platzhaltern und Template-Kennungen — er ist **nicht** a-checks Doku. Darum nimmt
  [`.d-check.yml`](../../../.d-check.yml) ihn per `scan.ignore` aus dem Doku-Gate: sonst prüfte
  `doc-check` fremde Platzhalter gegen a-checks Kennungs- und Linkregeln.
- **Nummern-Hinweis:** Das `conventions.template.md` der Baseline führt diese Adaption als
  `MR-003`. <!-- d-check:ignore (MR-Kennung des Baseline-Templates, nicht a-checks) --> Diese Nummer ist hier belegt
  ([`MR-003`](../../conventions.md#mr-003), aufgelöst 2026-06-21), darum
  die nächste freie — dieselbe Praxis wie in `ai-harness-init` (dort `MR-007`). <!-- d-check:ignore (fremde MR-Kennung des Repos ai-harness-init, nicht a-checks) --> Ob die
  Nummern-Identität mit dem Template hergestellt wird, entscheidet Etappe C der Migration.
- **Auflösungs-Trigger:** permanent (Provenienz/Baseline-Konformität).
