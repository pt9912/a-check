# a-check

A deterministic, network-free architecture gate for hexagonal codebases.
a-check reads a declared architecture from `.a-check.yml` — layers, roles,
allowed edges, technology boundaries — and reports where the real imports
disagree with it. **It repairs nothing and never writes into the checked
repository.**

## Usage

```bash
docker run --rm -v "$PWD:/src:ro" pt9912/a-check:__VERSION__ /src
```

The mounted directory is checked as `/src`; CLI options are appended. The
process runs as non-root, and a **read-only** mount is sufficient.

```bash
docker run --rm pt9912/a-check:__VERSION__ --print-config > .a-check.yml
docker run --rm pt9912/a-check:__VERSION__ --print-mk    > a-check.mk
```

## Exit codes

| Code | Meaning |
|---|---|
| `0` | no findings |
| `1` | findings reported |
| `2` | usage or configuration error (including: invalid `.a-check.yml`) |

## Reproducible runs

For CI, pin to the **digest** rather than to a moving tag. This image is a
**mirror** of `ghcr.io/pt9912/a-check` — the same image, not a second build:
the **config** digest is identical on both registries, and the release pipeline
verifies it after the push.

**The manifest digest, in contrast, is registry-local** — it depends on each
registry's blob compression. So take the digest **from the registry you pull
from**; one copied from GHCR will not resolve here. The digest listed in this
project's own documentation (`a-check.mk`, the READMEs, `version.md`) is the
**GHCR** one.

```bash
docker run --rm -v "$PWD:/src:ro" pt9912/a-check@sha256:<digest> /src
```

`:latest` moves for stable releases only; pre-releases never receive it.

## Rules

Ten rules over the declared architecture: `core-impurity`, `app-impurity`,
`port-impurity`, `lateral-adapter`, `lateral-slice`, `tech-leak`,
`port-direction-mismatch`, `port-locality`, `wrong-direction`,
`construct-leak`. Import extraction covers C++, Go, Rust, Kotlin, Java, Python,
C# and TypeScript — text-heuristic, with the limits reported rather than hidden.

## Documentation

- [README](https://github.com/pt9912/a-check#readme) — overview
- [User handbook](https://github.com/pt9912/a-check/blob/main/docs/user/benutzerhandbuch.md) — task-oriented (German)
- [Releasing](https://github.com/pt9912/a-check/blob/main/docs/user/releasing.md) — versions, digests, checklist

Source and issues: <https://github.com/pt9912/a-check> · Licence: MIT
