# Commit/PR-Traceability: jede Message nennt eine Kennung

*Ausgelagert aus [`AGENTS.md`](../../AGENTS.md) §5, Regel 1 — seit slice-201. Der Wortlaut ist unverändert; die Links sind der neuen Tiefe nachgeführt.*

Commits/PRs müssen mindestens eine `AC-*`- oder `ADR-*`-ID nennen
  (auch `MR-*`/`slice-NNN` gelten). Durchgesetzt durch `make trace-check`
  (d-check-Modul `commits`, [ADR-0021](../../docs/plan/adr/0021-commits-modul-trace-check.md))
  — lokal über `HEAD~1..HEAD`, in der CI über den Commit-Range
  ([`.github/workflows/ci.yml`](../../.github/workflows/ci.yml)). IDs werden nur
  beim Spec-/ADR-Schreiben nach dem deklarierten Schema vergeben (siehe
  [`harness/conventions.md`](../../harness/conventions.md)) — nie ad hoc im
  Commit/PR; Agenten referenzieren IDs, sie erfinden keine.
  **Struktur-IDs zählen nicht:** `SPEC-<NNN>` adressiert *innerhalb* der
  Spezifikation und gehört nicht in die Commit-Message — die `id-patterns` in
  [`.d-check.yml`](../../.d-check.yml) führen sie folgerichtig nicht.
