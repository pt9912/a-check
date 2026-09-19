**Vorgang:** slice-198

**Fund:** Beim Release `v0.20.0` (2026-09-19) war `make ci` lokal grün und der CI-Lauf auf `main`
rot (Lauf `35456322975`, Schritt „Traceability + ADR-Immutabilität", Exit 2). Ursache der roten
Schritte war eine ADR-Kern-Änderung **in der gepushten Range** — aber die Ursache dafür, dass der
Pre-Flight sie **nicht** sah, ist die Mengen-Differenz:

| Lauf | Was er fährt |
|---|---|
| `make ci` | `gates` + `image-test` |
| `.github/workflows/ci.yml` | `make ci` **plus** `trace-check`, `commit-scope-check`, `doc-immutable` über die Commit-Range |

`doc-immutable` ist **nicht** in `gates` — es ist im hermetischen Default absichtlich inert
(`--enable vcs`) und läuft nur als eigener CI-Schritt. Der lokale Lauf war also wahr und deckte die
CI nicht.

**Wie es auffiel:** an der roten CI-Ampe des ersten Release seit Monaten. Kein Gate nennt die
Differenz; beide Mengen sind für sich deklariert.

**Behoben** mit `make preflight` (slice-198), das `ci` plus die drei Range-Schritte über
`origin/main..HEAD` fährt. Die Probe: gegen die Range, die die CI rot machte
(`89fc7dc..ed7a3d8`), ist es **rot** mit denselben drei `core-drift-vcs`-Meldungen; gegen die
aktuelle Range grün. Kein Ersatz für `ci` in der Release-Pipeline — dort wäre „die Push-Range" die
Release-Range.
