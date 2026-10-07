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

For CI, pin to the **digest** rather than to a moving tag. The image is built for
**`linux/amd64` and `linux/arm64`** (macOS with Apple Silicon runs it natively), and the
digest is that of the **image index** — one pin for both platforms.

This repository is a **mirror** of `ghcr.io/pt9912/a-check` — the same image, not a second
build: the index is copied unchanged, and the release pipeline verifies that its digest is
**identical** on both registries. The digest listed in this project's own documentation
(`a-check.mk`, the READMEs, `version.md`) therefore resolves here as well — for releases
built for both platforms (from v0.23.0 on). For older tags, take the digest from this
registry: they were pushed anew and carry a different manifest digest here.

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
