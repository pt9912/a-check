#!/usr/bin/env bash
# pretooluse-command-guard — Tool-Call-Gate der Durchsetzungsschicht (dritter
# Bindepunkt; Regelwerk §Durchsetzungsschicht). Lehnt Host-Paketmanager und die
# Host-Go-Toolchain VOR der Ausführung ab — a-check ist Docker/make-only
# (AGENTS.md §3.1; Begründung: Toolchain-Reproduzierbarkeit AC-QA-03 + Hermetik
# AC-QA-02). Stack-Vorbild d-check (.claude/hooks/pretooluse-command-guard.sh).
#
# Geprüft wird die Befehlsposition jedes Kommando-Segments (Trennung an
# ; && || | $( ` ( und Zeilenenden) — `git commit -m "… pip …"` oder
# `docker run img npm test` bleiben erlaubt; `/usr/bin/pip` und `sudo pip`
# werden erkannt. Sub-Shell-Strings (`bash -c "…"`, auch Flag-Bündel -lc/-ec/-cx)
# werden rekursiv geprüft (Tiefe ≤ 3, dann fail-closed). Bewusst NICHT geprüft:
# andere Interpreter (`python -c`, `node -e`, `find -exec`) — der Guard ist ein
# Stolperdraht gegen versehentliche Host-Toolchain-Nutzung, keine Sandbox
# (dokumentierte Restlücke, analog zum Stop-Hook).
#
# Pass-Fall: KEINE Ausgabe — "approve" würde das Permission-System überspringen;
# ohne Ausgabe läuft die normale Permission-Entscheidung weiter.
# Selbsttest: `pretooluse-command-guard.sh --selftest` (eingehängt in `make gates`).
set -euo pipefail

# Fail-closed: ohne node ist keine zuverlässige Prüfung möglich → blockieren
# (Normalmodus) bzw. Selbsttest rot.
if ! command -v node >/dev/null 2>&1; then
  echo "pretooluse-command-guard: node nicht gefunden — fail-closed." >&2
  exit 2
fi

guard_verdict() {
  # $1 = vollständiges Hook-JSON; gibt "block" oder "ok" auf stdout.
  printf '%s' "$1" | node -e '
    const BLOCKED = new Set(["apt","apt-get","aptitude","dpkg","brew","pip",
      "pip3","pipx","npm","pnpm","yarn","npx","corepack","cargo","rustup",
      "rustc","gem","bundle","conda","poetry","go","gofmt","golangci-lint",
      "staticcheck"]); // Host-Go/-Toolchains: AGENTS §3.1
    const PREFIXES = new Set(["sudo","env","command","exec","nice","time",
      "xargs","nohup","eval"]);
    const SHELLS = new Set(["bash","sh","zsh","dash","ksh"]);
    const stripQuotes = t => t.replace(/^["'\'']+|["'\'']+$/g, "");

    function scan(cmd, depth) {
      if (depth > 3) return true; // zu tief verschachtelt → fail-closed
      const segments = cmd.split(/(?:;|&&|\|\||\||\$\(|`|\(|\r?\n)/);
      for (const seg of segments) {
        const tokens = seg.trim().split(/\s+/).filter(Boolean).map(stripQuotes);
        let i = 0;
        while (i < tokens.length &&
               (/^[A-Za-z_][A-Za-z0-9_]*=/.test(tokens[i]) || PREFIXES.has(tokens[i]))) i++;
        if (i >= tokens.length) continue;
        const head = tokens[i].replace(/^.*\//, ""); // /usr/bin/pip → pip
        if (BLOCKED.has(head)) return true;
        if (SHELLS.has(head)) {
          // -c auch in Flag-Bündeln (-lc, -ec, -cx, …): bei sh/bash ist c das
          // einzige Single-Letter-Flag mit Kommando-String-Semantik.
          const cIdx = tokens.findIndex((t, k) => k > i && /^-[a-z]*c[a-z]*$/.test(t));
          if (cIdx !== -1 && cIdx + 1 < tokens.length &&
              scan(tokens.slice(cIdx + 1).join(" "), depth + 1)) return true;
        }
      }
      return false;
    }

    let s = "";
    process.stdin.on("data", d => s += d);
    process.stdin.on("end", () => {
      let cmd = "";
      try {
        const j = JSON.parse(s);
        cmd = String((j.tool_input && j.tool_input.command) || "");
      } catch { process.stdout.write("block"); return; } // unlesbar → fail-closed
      process.stdout.write(scan(cmd, 0) ? "block" : "ok");
    });
  '
}

emit_block() {
  cat <<'JSON'
{
  "decision": "block",
  "reason": "a-check ist Docker/make-only (AGENTS.md §3.1): Host-Paketmanager und die Host-Go-Toolchain (go/golangci-lint/pip/npm/cargo/apt/brew/…) sind verboten. Nutze die make-Targets (make lint/test/build/arch-check/gates); die Go-Toolchain läuft in Docker."
}
JSON
}

# ── Selbsttest ───────────────────────────────────────────────────────────────
if [ "${1:-}" = "--selftest" ]; then
  fail=0
  assert() { # $1 erwartet (block|ok)  $2 json  $3 beschreibung
    local got; got="$(guard_verdict "$2")"
    if [ "$got" != "$1" ]; then
      echo "guard-selftest FAIL: erwartet '$1', bekam '$got' — $3" >&2
      fail=1
    fi
  }
  assert block '{"tool_input":{"command":"go build ./..."}}'                    "Host-go"
  assert block '{"tool_input":{"command":"sudo apt-get install -y x"}}'         "sudo+apt-get"
  assert block '{"tool_input":{"command":"env FOO=bar pip3 install x"}}'        "env+Zuweisung+pip3"
  assert block '{"tool_input":{"command":"bash -lc \"npm install\""}}'          "Sub-Shell -lc npm"
  assert block '{"tool_input":{"command":"/usr/local/bin/golangci-lint run"}}'  "absoluter Pfad golangci-lint"
  assert ok    '{"tool_input":{"command":"make help"}}'                         "make erlaubt"
  assert ok    '{"tool_input":{"command":"git commit -m \"erwaehnt pip und npm\""}}' "Toolname nur im Arg-String"
  assert ok    '{"tool_input":{"command":"docker run --rm img npm test"}}'      "npm als docker-Argument"
  assert ok    '{"tool_input":{"command":"grep -rn \"go \" ."}}'                "go nur im grep-Muster"
  if [ "$fail" -ne 0 ]; then
    echo "guard-selftest: FEHLGESCHLAGEN" >&2
    exit 1
  fi
  echo "guard-selftest ok: Denylist greift (Host-Toolchain blockiert; make/git/docker erlaubt)."
  exit 0
fi

# ── Normalmodus ──────────────────────────────────────────────────────────────
input="$(cat)"
verdict="$(guard_verdict "$input")"
if [ "$verdict" = "block" ]; then
  emit_block
fi
# Pass-Fall: keine Ausgabe — die normale Permission-Prüfung übernimmt.
