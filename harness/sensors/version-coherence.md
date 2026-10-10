# `make version-coherence` — Kohärenz doppelt deklarierter Versions-Angaben und lesbare `FROM`-Zeilen

## Vertrag

Drei Regeln, alle hermetisch: Derselbe `uses:`-SHA trägt unter
`.github/workflows/` überall denselben Tag-Kommentar. Eine
Versions-Variable, die `Makefile` **und** `Dockerfile` führen, hat an beiden
Orten denselben Wert. Und jede `FROM`-Zeile mit `@sha256:` in `Dockerfile` und
`tools/archive-wave/Dockerfile` trägt einen **wörtlichen Tag** — kein `${…}`, kein
Digest allein ([`ADR-0044`](../../docs/plan/adr/0044-dependabot-docker-woertliche-tags.md)):
Dependabot macht aus einer Zeile nur dann eine Abhängigkeit, wenn der Tag im Text steht.

## Grenze — was das Grün nicht abdeckt

1. **Die Wahrheit einer Angabe** — geprüft wird **Divergenz, nicht Unwahrheit**.
   Zwei übereinstimmend *falsche* Angaben bleiben grün. Permanent: Die Registry
   zu fragen wäre Netz, und `gates` ist hermetisch.
2. **Welche Seite recht hat** — der Sensor erklärt **keine** zur führenden. Er
   meldet, dass zwei Orte auseinanderlaufen, und überlässt das Urteil dem
   Leser. Permanent, und Absicht.
3. **Dependabots Verhalten** — Regel (3) prüft die **Form** der Zeile, nicht ob Dependabot
   sie fortschreibt. Dass er den Index-Digest eines Multi-Arch-Tags richtig mitführt, zeigt
   erst ein echter PR. Die Zeilen-Auswahl ist `FROM <…>@sha256:…`; ein `FROM` ohne Digest
   (Stage-Verweis wie `FROM deps AS test`) gehört nicht zur Prüfmenge.
4. **Einfach deklarierte Angaben** — eine Version, die nur an einem Ort steht,
   hat keinen Partner und fällt aus der Prüfmenge. Für Baseline-Pins schließt
   `versions` in `doc-check` diese Lücke, für andere niemand.

## Bindung

slice-131 · Antwort auf
[`BEO-GATE/versionsangabe-neben-digest-ungeprueft`](../../docs/plan/planning/observations/BEO-GATE/versionsangabe-neben-digest-ungeprueft/observation.md)
bei 3× · Regel (3) seit slice-226 ([`ADR-0044`](../../docs/plan/adr/0044-dependabot-docker-woertliche-tags.md)) · im `gates`-Aggregat.
