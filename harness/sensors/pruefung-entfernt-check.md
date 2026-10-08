# `make pruefung-entfernt-check` — eine entfernte Prüfung braucht eine Begründung

## Vertrag

Entfernt ein Commit unter `tools/` oder `.github/workflows/` eine **Fehlerpunkt-Zeile** —
`fail "`, `::error::`, `exit 1`/`exit 2`, `probe "`, `assert `, ohne Kommentarzeilen — und steht
dieselbe Zeile (nach Trim) im selben Commit nicht in **derselben Datei** wieder da (gezählt als
Multimenge), muss seine Message eine Zeile
`Entfernte-Pruefungen: <Grund>` mit nicht leerem Grund tragen. Sonst rot, mit den entfernten
Zeilen in der Ausgabe. Regel: [`harness/rules/entfernte-pruefungen.md`](../rules/entfernte-pruefungen.md).

Drei Aufrufer, eine Wahrheit: der `commit-msg`-Hook (`MSGFILE=`, gegen den Index, **vor** dem
Commit), `make preflight` und der CI-Workflow (`RANGE=`). Ohne Argument prüft der Lauf
`HEAD~1..HEAD`. Vor jedem Lauf fährt das Target den Selbsttest (`--selftest`, netzlos bis auf ein
Wegwerf-Repo im Temp-Verzeichnis, das ohne die git-Umgebung des Aufrufers läuft: im Hook setzt
git bei `commit -a` einen absoluten `GIT_INDEX_FILE`, und der Selbsttest berührt ihn nicht).

## Grenze — was das Grün nicht abdeckt

1. **Ob die Prüfung verloren ist.** Der Sensor zeigt die entfernten Zeilen und verlangt einen
   Grund; ob die Prüfung ersetzt oder verloren ist, entscheidet der Autor, und der Review hält es
   gegen den Diff.
2. **Andere Formen und Orte.** Nur die fünf Muster, nur `tools/` und `.github/workflows/`: ein
   `return 1` ohne Meldung, ein Go-Testfall, eine Prüfung in `.claude/hooks/` sieht er nicht.
3. **Umformuliert oder in eine andere Datei verschoben zählt als entfernt**; ein nacktes `exit 1`,
   das in derselben Datei an anderer Stelle neu entsteht, gleicht ein entferntes aus. Gemessen im
   Bestand (Fassung vor der Datei-Bindung): 12 von 94 Werkzeug-Commits hätten die Zeile gebraucht,
   darunter alle drei Belege.
4. **Altbestand.** Ein Commit, dessen `AGENTS.md` den Anker `Entfernte-Pruefungen:` noch nicht
   trägt, wird übersprungen (Grandfathering wie bei `commit-scope-check`).
5. **Der Hook ist opt-in pro Klon** (`make hooks`); die klon-unabhängige Kontrolle ist der
   CI-Range-Schritt.

## Ausgänge

| Exit | Bedeutung |
|---|---|
| 0 | kein Commit entfernt eine Fehlerpunkt-Zeile ohne Begründung |
| 1 | mindestens einer — die Zeilen stehen in der Ausgabe |
| 2 | Range nicht auflösbar, `git diff --cached` gescheitert (Hook), oder der Selbsttest schlägt fehl |

## Bindung

[`AGENTS.md`](../../AGENTS.md) §5 Regel 18 · Antwort auf
[`BEO-GATE/umbau-verliert-pruefung-still`](../../docs/plan/planning/observations/BEO-GATE/umbau-verliert-pruefung-still/observation.md)
bei 3× · slice-222.
