# Die neue Kennung steht in der Closure-Notiz

*Ausgelagert aus [`AGENTS.md`](../../AGENTS.md) §5, Regel 3 — seit slice-201. Der Wortlaut ist unverändert; die Links sind der neuen Tiefe nachgeführt.*

**Wer eine Anforderung anlegt, nennt ihre Kennung in der Closure-Notiz.** Im **Plan** kann er es
  nicht: IDs werden referenziert statt erfunden, die neue Kennung existiert dort noch nicht, und
  jede genannte ist linkpflichtig — ein Link ins Leere macht `doc-check` rot. Also umschreibt der
  Plan sie („eine neue `AC-FA-CLI`-Kennung"), und die Requirements-Matrix sieht den Slice **nicht**.
  Bei der Closure ist die Anforderung geschrieben; dort steht die Kennung mit Link. Durchgesetzt
  durch `make doc-complete` im `verify`-Aggregat — eine Anforderung ohne
  referenzierenden Slice ist abschluss-blockierend.
