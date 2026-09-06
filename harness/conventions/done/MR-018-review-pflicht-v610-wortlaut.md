# MR-018 — Review-Pflicht/Rollenwechsel-Absatz in `AGENTS.md` §6 auf `v6.1.0`-Wortlaut

- **Status:** Accepted
- **Datum:** 2026-09-05
- **Geltungsbereich:** [`AGENTS.md`](../../../AGENTS.md) §6 (Minimal Agent Workflow, Absatz nach
  Schritt 8)
- **Ersetzt-Baseline-Regel:** — *(kein Regelwerk-Modul hat sich geändert — `modul-08-agentenrollen.md`
  ist zwischen `v6.0.0` und `v6.1.0` byte-identisch, gemessen in
  [slice-161](../../../docs/plan/planning/done/slice-161-regelwerk-v610-delta-analyse.md) §2. Der
  tatsächliche Wortlaut-Treiber ist `lab/templates/AGENTS.template.md` in seiner `v6.1.0`-Fassung —
  ein **Template**, keine Regelwerk-Regel, und dazu eines, das a-check noch nicht vendored hat; es
  gibt also keinen netzlos auflösbaren Anker, auf den dieses Feld präzise zeigen könnte. Derselbe
  Fall wie [`MR-017`](../../conventions.md#mr-017): eine Provenienz-/Wortlaut-Korrektur, keine
  Regel-Adaption — Rückbau-Kandidat, sobald `v6.1.0` vendored ist und auf das dann vendorte Template
  gezeigt werden kann)*. `modul-08-agentenrollen.md` §Die neun Übergaben bleibt der **konzeptionelle**
  Bezug — Kontext-Trennung ist dort normiert und unverändert; der neue Absatz macht das nur
  **explizit**, ersetzt aber keine dort stehende Regel.
- **Adaption:** `AGENTS.md` §6 trägt nach Schritt 8 einen Absatz, der Schritt 8 explizit als
  Rollenwechsel statt Abschluss deklariert (Handoff Implementer → Reviewer → Verifier, kein
  Self-Review). Der Wortlaut folgt eng dem Zuwachs, den der Kurs (`pt9912/ai-harness-course`) mit
  `v6.1.0` in `lab/templates/AGENTS.template.md` einführt — a-check hatte dieselbe Regel bereits
  taggleich mit slice-159 (2026-09-05) formuliert, nur ausführlicher und ohne diesen Eintrag.
  Zwei a-check-spezifische Sätze bleiben über den Baseline-Wortlaut hinaus: der explizite
  `fork`-Ausschluss (ein `fork`-Subagent erbt den Kontext und zählt darum nicht als getrennter
  Kontext, Modul 8 §Kontext-Trennung) und der Zitat-Anker auf den auslösenden Befund
  ([`BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen`](../../../docs/plan/planning/observations/BEO-HARNESS/behauptete-vollstaendigkeit-extern-gefangen/observation.md)).
- **Begründung:** Zwei getrennte Gründe. Erstens Provenienz: die Regel bestand bereits als
  a-check-eigene Ergänzung ggü. der Baseline, aber ohne `MR`-Eintrag — eine Lücke, die
  [slice-161](../../../docs/plan/planning/done/slice-161-regelwerk-v610-delta-analyse.md) §4.4 beim
  Messen des `v6.1.0`-Sprungs aufdeckte. Zweitens Umfang: die bisherige a-check-Fassung
  duplizierte mechanische Details, die bereits an ihrem eigentlichen Ort stehen: das
  Report-Dateiname-Muster (`docs/reviews/<YYYY-MM-DD>-<slice-oder-diff-ref>.md`) in
  [`review-report.template.md`](../../../.harness/baseline/v6.2.0/templates/docs/reviews/review-report.template.md)
  (verlinkt aus [`docs/reviews/README.md`](../../../docs/reviews/README.md)), die
  HIGH-Verifikationspflicht in [`.harness/skills/reviewer.md`](../../../.harness/skills/reviewer.md) —
  sowie eine
  Chronik-Erklärung, warum die Zeile hier steht — beides widerspricht `AGENTS.md`s eigener
  Ziel-Form ("sie trägt Hard Rules und Pointer … sie dupliziert deren Inhalt nicht, sonst
  entsteht Drift", §1). Der Rückschnitt auf den Baseline-Wortlaut behebt beides gleichzeitig:
  näher am Kurs, kürzer, ohne Informationsverlust (die entfernten Details bleiben an ihrem
  eigentlichen Ort auffindbar).
- **Auflösungs-Trigger:** die nächste Baseline-Migration, die `modul-08-agentenrollen.md` oder
  den Rollenwechsel-Absatz in `AGENTS.template.md` inhaltlich ändert.
- **Ausgelöst durch Baseline-Stand:** `v6.1.0` (Kurs-Welle 118, 2026-09-05); vendored bleibt zum
  Zeitpunkt dieses Eintrags `v6.0.0` — die Übernahme des Wortlauts ist unabhängig vom Re-Vendoring
  (slice-161 §6, Etappe A).
