# MR-023 — ID-Schema-Deklaration um die Beobachtungs-Kennung ergänzt

- **Status:** Accepted
- **Datum:** 2026-09-08
- **Geltungsbereich:** [`MR-000`](../conventions.md#mr-000) §ID-Schema-Deklaration, fehlende Zeile
  zur Beobachtungs-Kennung
- **Ersetzt-Baseline-Regel:** — *(keine — wie bei [`MR-020`](../conventions.md#mr-020) korrigiert
  dieser Eintrag eine **Repo**-Aussage, keine Baseline-Regel: die Ziel-Form verlangt die Kennung
  in der Deklaration, a-check führt sie dort nicht.)*
- **Adaption:** Die Deklaration in [`MR-000`](../conventions.md#mr-000) nennt Anforderungen, ADRs,
  Adaptionen, Carveouts und Slices, aber **keine Beobachtungs-Kennung**. Die vendorte Ziel-Form
  `harness/conventions.template.md` führt sie in ihrer Liste. Ab sofort gilt: Die
  Beobachtungs-Kennung ist **der Pfad** `BEO-<KUERZEL>/<slug>` unter
  [`docs/plan/planning/observations/`](../../docs/plan/planning/observations/README.md) — keine
  fortlaufende Nummer. `<KUERZEL>` wird in [§Modus-Deklaration pro Sub-Area](../conventions.md#modus-deklaration-pro-sub-area)
  **nachgeschlagen, nicht erfunden**; `<slug>` ist lowercase Kebab-Case.
- **Begründung:** Die Form ist seit [slice-138](../../docs/plan/planning/done/wellenlos/slice-138-sub-area-kuerzel.md)/[slice-139](../../docs/plan/planning/done/wellenlos/slice-139-beobachtungsregister-migration.md)
  gelebte Praxis und an zwei Orten beschrieben ([`AGENTS.md`](../../AGENTS.md) §5 und
  [`docs/plan/planning/README.md`](../../docs/plan/planning/README.md)) — nur nicht dort, wo ein
  Leser das ID-Schema nachschlägt. Das ist dieselbe Klasse, die
  [slice-187](../../docs/plan/planning/in-progress/slice-187-voll-abgleich-erstdurchgang-rest.md)
  §3.1 an drei `AGENTS.md`-Stellen gefunden hat: **die Praxis war da, der Satz fehlte.**

  **Warum ein `MR` und kein Absatz daneben:** §Disziplin nennt für eine Korrektur an einem
  akzeptierten Eintrag **zwei** Instrumente — neuer `MR` oder ausdrückliche Aufhebung.
  [`MR-020`](../conventions.md#mr-020) hat genau diesen Defekt (eine veraltete Zeile in derselben
  Deklaration) so behoben und ist der Präzedenzfall. slice-187 hatte zunächst ein drittes,
  nirgends vorgesehenes Werkzeug gewählt (einen freistehenden Absatz oberhalb des Eintrags); der
  unabhängige Review hat das als Befund geführt.

  **Generisch statt aufzählend**, aus demselben Grund wie bei
  [`MR-020`](../conventions.md#mr-020): Der Eintrag nennt die **Form** und verweist für das Kürzel
  auf §Modus-Deklaration, die ohnehin mit jeder neuen Sub-Area gepflegt wird. Ein hier
  eingefrorenes Kürzel-Verzeichnis wäre bei der nächsten Sub-Area veraltet und erzwänge den
  nächsten Nachfolge-Eintrag.
- **Auflösungs-Trigger:** permanent für die Form. Der Eintrag wird **entbehrlich**, sobald die
  ID-Schema-Deklaration als Ganzes überarbeitet ist statt durch Nachfolge-Einträge geflickt —
  dieser Rückbau ist als [slice-190](../../docs/plan/planning/open/slice-190-id-schema-deklaration-ueberarbeiten.md)
  geplant und ist der Ausgang von
  [`BEO-HARNESS/adaption-korrigiert-repo-aussage`](../../docs/plan/planning/observations/BEO-HARNESS/adaption-korrigiert-repo-aussage/observation.md)
  bei 3×.
