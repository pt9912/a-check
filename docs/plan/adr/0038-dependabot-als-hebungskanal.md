# ADR-0038 — Dependabot als Hebungs-Kanal: gate-konform statt ausgenommen, ohne Automerge

- **Status:** Accepted
- **Datum:** 2026-08-30
- **Autor:** pt9912
- **Bezug:** [`AGENTS.md`](../../../AGENTS.md) §5 (Traceability-Pflicht je Commit), §4 (`make trace-check`, `make image-scan`); [ADR-0037](0037-cve-scan-gegen-das-publizierte-image.md) (der Sensor, der meldet); [ADR-0030](0030-kein-digest-im-generierten-fragment.md) (Digest-Hebung als bewusster Commit)
- **Schärft:** — (Betriebs-/Prozess-Entscheidung; koppelt an keine Spec-§. Sie realisiert einen Kanal, den die Spec-Straten nicht beschreiben — dieselbe Verortung wie [ADR-0005](0005-lint-profil.md) und [ADR-0037](0037-cve-scan-gegen-das-publizierte-image.md).)
- **Supersedes:** —

## Kontext

[ADR-0037](0037-cve-scan-gegen-das-publizierte-image.md) hat einen Sensor gebracht, der bekannte
Schwachstellen im publizierten Image meldet. Sein Erstlauf fand neun behebbare HIGH. **Gehoben hat
sie ein Mensch** — mit gelesener Ausgabe, in zwei Slices, an einem Tag.

Das trägt einmal. Es trägt nicht als Verfahren: ein Sensor, dessen Befunde nur dann einen Weg nach
oben finden, wenn jemand hinsieht, verlagert die Zusage vom Werkzeug auf die Aufmerksamkeit.

**Die Kopplung, und sie steht vor der Konfiguration.** `make trace-check` verlangt von **jeder**
Commit-Message eine Kennung — `AC-*`, `ADR-*`, `MR-*` oder `slice-*` —, geprüft in der PR-CI über
die Commit-Range ([`AGENTS.md`](../../../AGENTS.md) §5). Dependabots Botschaften tragen keine. Ein
PR wäre rot, bevor ihn jemand ansieht.

**Gemessen gegen das echte Gate, nicht angenommen:**

| Botschaft | `make trace-check` |
|---|---|
| `build(deps): bump gopkg.in/yaml.v3 from 3.0.1 to 3.0.2` | **Exit 2**, `commit-untraceable` |
| `build(deps) [ADR-0038]: bump gopkg.in/yaml.v3 …` | **Exit 0** |
| `build(ci) [ADR-0038]: bump actions/checkout …` | **Exit 0** |

## Entscheidung

**1. Die neue Commit-Klasse wird gate-KONFORM gemacht, nicht ausgenommen.** Jeder
`commit-message.prefix` trägt die Kennung **dieses** Entscheids. Das ist kein Trick: die ADR, die
den Kanal erlaubt, ist der ehrliche Anker für die Commits, die daraus entstehen.
`commits.exempt-pattern` in [`.d-check.yml`](../../../.d-check.yml) bleibt **unverändert** — eine
Ausnahme hätte die Zusage aufgeweicht, statt sie zu erfüllen, und sie hätte für **jede** künftige
maschinelle Commit-Klasse gegolten.

**2. Kein Automerge.** Der Kanal öffnet PRs; geprüft wird mit `make ci`, und der Merge bleibt ein
Akt. Dieselbe Linie wie bei jeder Digest-Hebung dieses Repos
([ADR-0030](0030-kein-digest-im-generierten-fragment.md)): ein Pin-Wechsel ist ein **bewusster**
Commit mit gelesener Ausgabe.

**3. Zwei Ökosysteme, nicht drei.** Module und GitHub Actions — beide leben in einem Manifest, das
der Kanal liest. **`docker` bleibt draußen**: ein Digest-Hoist ist ein bewusster Commit, und die
beiden Hebungen dieses Tages (Toolchain, Lint) liefen genau so. Der Trivy-Pin in
`tools/image-scan.sh` und der Schwester-Tool-Pin in `d-check.mk` fallen ohnehin heraus — sie
leben in keinem Manifest.

**Verworfene Alternative — `commits.exempt-pattern` um Dependabots Präfix erweitern.** Sie wäre
kürzer und träfe genau die richtige Klasse. Sie nimmt aber eine Zusage zurück, statt sie zu
erfüllen: die Traceability-Pflicht sagt, dass **jeder** Commit auf eine Entscheidung zeigt. Ein
Dependabot-Commit tut das — auf diese ADR. Die Ausnahme hätte behauptet, er tue es nicht.

## Konsequenzen

- **Der Kanal erreicht die Fundklasse des Sensors nur zur Hälfte, wenn die Repository-Schalter aus
  sind.** Ohne `dependabot_security_updates` und die Dependabot-Alerts öffnet ein CVE **ohne**
  neues Upstream-Release keinen PR. Die Schalter leben in der GitHub-Oberfläche, nicht im Repo;
  diese ADR kann sie nennen, setzen muss sie der Maintainer
  ([`docs/user/releasing.md`](../../user/releasing.md) §Vorbedingungen).
- **Die Form aus dem Repo bleibt Pflicht.** Ein PR, der einen Tag statt eines SHA setzte, wäre ein
  Rückschritt gegenüber der digest-gepinnten Praxis. Heute prüft das **kein** Gate
  ([`BEO-026`](../../../docs/plan/planning/observations/BEO-GATE/versionsangabe-neben-digest-ungeprueft/observation.md)) — der Kanal macht diese Lücke praktisch relevanter,
  weil er künftige `uses:`-Hebungen **erzeugt**.
- **Der Zuschnitt ist klein und bleibt es vorerst:** ein direktes Modul, kein indirektes, zwei
  externe Actions. Deshalb steht hier **kein** `allow: dependency-type: all` — im Schwester-Repo
  ist es die Bedingung, unter der der Eintrag überhaupt wirkt (dreizehn von vierzehn Befunden
  lagen indirekt); hier wäre es eine Zeile ohne Gegenstand.
- **Ein wöchentlicher PR-Strom ohne Merge-Disziplin wird zum Rauschen.** Dieselbe Klasse wie ein
  dauerhaft rotes Abzeichen: was man regelmäßig wegklickt, prüft nichts mehr. Das Limit von fünf
  offenen PRs je Ökosystem ist die Gegenmaßnahme, nicht die Lösung.

## Fitness Function

`make trace-check MSGFILE=<datei>` gegen beide Formen: eine Dependabot-Botschaft **ohne** Kennung
muss `commit-untraceable` melden, dieselbe **mit** dem Präfix muss grün sein. Beide Richtungen
sind oben gemessen; ohne die rote Hälfte wäre ein Präfix, der nichts bewirkt, von einem wirksamen
nicht zu unterscheiden.
