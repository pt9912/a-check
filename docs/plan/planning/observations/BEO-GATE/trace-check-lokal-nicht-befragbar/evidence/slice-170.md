**Vorgang:** slice-170
**Fund:** Vom unabhängigen Review als Nebenbefund gemeldet und nachgemessen: `make trace-check`
bricht mit `Range-Basis-Vorfahren nicht lesbar: object not found` ab — mit `RANGE=HEAD~2..HEAD`
ebenso wie mit dem Default. `git rev-parse --is-shallow-repository` → `false`,
`git cat-file -t HEAD~2` → `commit`. `ls .git/objects/pack/` zeigt neben
`pack-2e8999b9….pack` ein `loose-ead221f3….pack`. Nicht durch diesen Slice verursacht; die
Traceability der vier Commits wurde ersatzweise von Hand geprüft (alle nennen `slice-170`).
