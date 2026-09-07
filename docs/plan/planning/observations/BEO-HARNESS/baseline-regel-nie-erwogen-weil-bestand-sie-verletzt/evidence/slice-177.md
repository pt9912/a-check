**Vorgang:** slice-177 (Planung; Fund beim Maintainer-Einwand am 2026-09-07)

**Fund:** Die Ziel-Form `v6.2.0` · `templates/AGENTS.template.md`
§4 sagt zur Gate-Tabelle: *„Diese Tabelle **listet auf**; definiert wird hier nichts. Die
**Bindung** eines Targets — welche Anforderung oder Entscheidung es durchsetzt — steht in
`harness/README.md` §Sensors."* Ihre Beispielzellen sind `<…>`.

a-checks [`AGENTS.md`](../../../../../../../AGENTS.md) §4 führt dort Bindungen, Begründungen,
Grenzen und Vorfallsgeschichten: **16 von 40** Zellen über 250 Zeichen, Spitze `doc-check` mit
**1871**. Dieselben Verträge stehen zusätzlich in `harness/README.md` §Sensors, dort kürzer — bei
`doc-check` 639 Zeichen. Zwei Fassungen derselben Zusage.

**Gemessen, seit wann die Regel gilt:** `git show <tag>:lab/templates/AGENTS.template.md` im
Kurs-Klon — der Absatz steht **wortgleich in `v5.12.0`, `v6.0.0`, `v6.2.0` und `v6.5.0`**. a-check
hat ihn nie befolgt, und **keine** aktive Adaption deckt die Abweichung: [`MR-019`](../../../../../../../harness/conventions.md#mr-019) betrifft
`AGENTS.md` §5, nicht §4.

**Warum es zwei Adaptions-Durchgänge überlebt hat** — und das ist der Kern dieses Eintrags:
`slice-137` (gegen `v6.0.0`) und `slice-163` (gegen `v6.1.0`) prüften die **vorhandenen**
`MR`-Einträge gegen die neue Baseline. Eine Regel **ohne** Eintrag hat dort keinen Aufhänger. Die
Beobachtung sagt genau das voraus: *„der Adaptions-Durchgang tut das für Regeln, zu denen es
bereits einen Eintrag gibt — für Regeln ohne Eintrag tut es niemand."*

**Zweite Instanz** (slice-170 · dieser Fund). Aufgefallen ist sie wie in der Beobachtung
beschrieben — nicht beim Brechen, sondern beim **Lesen**: der Maintainer fragte, warum wir uns
nicht an die Vorlage halten. Kein Sensor hätte es gemeldet; es ist eine Form, die keiner prüft.
