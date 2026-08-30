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

    // ── Regel 2: Gate-Lauf darf nicht verschluckt werden (slice-057) ──────────
    // Ein `make <gate> | tail` liefert den Exit-Code von tail; ein rotes Gate
    // verschwindet spurlos. Dasselbe gilt fuer `make <gate> && git commit` —
    // der Commit haengt dann an einem ungeprueften Lauf. Fuenf reale Vorfaelle
    // am 2026-07-25 (Steering-Loop SL-001).
    // Vollstaendig gegen die deklarierten Pruef-Targets (AGENTS.md §4). Der
    // `--selftest` haelt die Liste dagegen aktuell — ohne ihn driftete sie:
    // `doc-immutable` fehlte, obwohl es als CI-durchgesetzt gefuehrt wird, und
    // `make doc-immutable | tail` lief ungehindert durch (Review 2026-07-26,
    // R-057-F1). Nicht enthalten sind Targets, deren Ausgabe bestimmungsgemaess
    // weiterverarbeitet wird — siehe NICHT_PRUEFEND im Selbsttest.
    const GATES = new Set(["gates","verify","ci","lint","test","coverage-gate",
      "arch-check","doc-check","image-test","trace-check","suppression-check",
      "gate-consistency","verify-risiko-ausgaenge",
      "verify-observations","commit-scope-check","guard-selftest","doc-complete","doc-immutable",
      "doc-commits","doc-planning","doc-tracked","doc-targets","doc-structure","doc-workflows",
      "regelwerk-check"]);

    function hasGateMake(seg) {
      const t = seg.trim().split(/\s+/).filter(Boolean).map(stripQuotes);
      const i = t.indexOf("make");
      if (i === -1) return false;
      for (let k = i + 1; k < t.length; k++) {
        if (t[k].startsWith("-") || t[k].includes("=")) continue;
        if (GATES.has(t[k])) return true;
      }
      return false;
    }

    // Zitierte Argumente ausblenden: ein `make <gate>`-Muster INNERHALB eines
    // Strings ist keine Pipeline, sondern Text (Doku, Testfall, Commit-Message).
    // Ohne das feuerte die Regel auf ihren eigenen Selbsttest — real passiert
    // beim Bau, siehe Steering-Loop SL-001.
    // EHRLICHE GRENZE: ein echtes Sub-Shell-Kommando mit demselben Muster in
    // Anfuehrungszeichen entgeht damit. Der Guard ist ein Stolperdraht, keine
    // Sandbox (Regelwerk §Durchsetzungsschicht): er faengt die versehentliche
    // Drift, nicht die umgeleitete.
    //
    // Heredoc-Inhalte sind DATEN, kein Kommando: eine Commit-Message oder ein
    // Dateiinhalt darf ein Gate-Muster zitieren, ohne es auszufuehren. Ohne das
    // blockierte die in slice-064 verschaerfte Regel ihren eigenen Commit, weil
    // dessen Message den abgelehnten Aufruf zitiert (SL-004 — ein Sensor meldet
    // sein eigenes Umfeld, vierter Vorfall). GRENZE wie bei Sub-Shell-Strings:
    // ein Heredoc an einen Interpreter (`bash <<EOF`) entgeht ebenfalls.
    function stripHeredoc(cmd) {
      return cmd.replace(
        /<<-?\s*[\x27"]?([A-Za-z_][A-Za-z0-9_]*)[\x27"]?\r?\n[\s\S]*?\r?\n\1\b/g,
        (m, tag) => "<<" + tag + "\n" + tag);
    }
    const blank = t => "\u0000".repeat(t.length);
    function stripQuoted(cmd) {
      return cmd.replace(/"[^"]*"|'\''[^'\'']*'\''/g, m => m[0] + blank(m.slice(1, -1)) + m[0]);
    }

    function pipeViolation(raw) {
      const cmd = stripQuoted(stripHeredoc(raw));
      // Trennzeichen erhalten, damit die Folge-Beziehung lesbar bleibt.
      const parts = cmd.split(/(\|\||&&|\||;|\r?\n)/);
      for (let i = 0; i < parts.length; i++) {
        if (!hasGateMake(parts[i])) continue;
        const sep = parts[i + 1];
        if (sep === "|") return true;                    // Ausgabe in eine Pipe
        // Commit am Lauf — in JEDER Verkettung, nicht nur nach `&&`.
        // Bis slice-064 hing diese Pruefung an `sep === "&&"`; ein Commit nach
        // `;` oder Zeilenumbruch fiel durch. Genau so ging 1a9f270 mit rotem
        // `make gates` heraus (SL-001, sechster Vorfall) — und zwar in der
        // Schreibweise, die der Guide selbst nahelegt: wer die Umleitung in eine
        // Datei befolgt, verkettet danach mit `;`.
        const rest = parts.slice(i + 2).join("");
        if (/\bgit\s+commit\b/.test(rest)) return true;
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
      } catch { process.stdout.write("block-toolchain"); return; } // unlesbar → fail-closed
      if (scan(cmd, 0)) { process.stdout.write("block-toolchain"); return; }
      process.stdout.write(pipeViolation(cmd) ? "block-pipe" : "ok");
    });
  '
}

emit_block_pipe() {
  cat <<'JSON'
{
  "decision": "block",
  "reason": "Gate-Lauf nicht verschlucken (AGENTS.md §6, slice-057): `make <gate> | …` liefert den Exit-Code des letzten Pipe-Glieds, nicht den von make — ein rotes Gate verschwindet spurlos. Ebenso ein `git commit` nach einem Gate-Lauf im SELBEN Aufruf — gleich mit welcher Verkettung (`&&`, `;`, Zeilenumbruch): der Commit haengt dann an einem Lauf, dessen Exit-Code niemand gelesen hat. Richtig: `make <gate> > /tmp/gates.log 2>&1; echo \"EXIT=$?\"` als eigener Aufruf, den Exit-Code lesen, und den Commit erst danach in einem ZWEITEN Aufruf."
}
JSON
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
  assert block-toolchain '{"tool_input":{"command":"go build ./..."}}'                    "Host-go"
  assert block-toolchain '{"tool_input":{"command":"sudo apt-get install -y x"}}'         "sudo+apt-get"
  assert block-toolchain '{"tool_input":{"command":"env FOO=bar pip3 install x"}}'        "env+Zuweisung+pip3"
  assert block-toolchain '{"tool_input":{"command":"bash -lc \"npm install\""}}'          "Sub-Shell -lc npm"
  assert block-toolchain '{"tool_input":{"command":"/usr/local/bin/golangci-lint run"}}'  "absoluter Pfad golangci-lint"
  assert ok    '{"tool_input":{"command":"make help"}}'                         "make erlaubt"
  assert ok    '{"tool_input":{"command":"git commit -m \"erwaehnt pip und npm\""}}' "Toolname nur im Arg-String"
  assert ok    '{"tool_input":{"command":"docker run --rm img npm test"}}'      "npm als docker-Argument"
  assert ok    '{"tool_input":{"command":"grep -rn \"go \" ."}}'                "go nur im grep-Muster"
  # Regel 2 (slice-057): Gate-Lauf nicht verschlucken.
  assert block-pipe '{"tool_input":{"command":"make gates | tail -20"}}'              "make gates in eine Pipe"
  assert block-pipe '{"tool_input":{"command":"make verify | grep ok"}}'              "make verify in eine Pipe"
  assert block-pipe '{"tool_input":{"command":"make gates && git commit -m x"}}'      "Commit an ungepruefenten Lauf gekettet"
  assert ok    '{"tool_input":{"command":"make gates > /tmp/g.log 2>&1"}}'       "Umleitung in eine Datei ist richtig"
  assert ok    '{"tool_input":{"command":"make help | grep verify"}}'           "Nicht-Gate-Target darf gepiped werden"
  assert ok    '{"tool_input":{"command":"grep -E x /tmp/g.log | tail -3"}}'    "Pipe ohne make"
  assert ok    '{"tool_input":{"command":"make gates && echo fertig"}}'         "Verkettung ohne Commit"
  assert ok    '{"tool_input":{"command":"echo \"make gates | tail\" >> doku.md"}}'   "Muster nur im Argument-String"
  assert block-pipe '{"tool_input":{"command":"make doc-immutable | tail -1"}}'  "doc-immutable ist CI-durchgesetzt"
  # slice-064: der reale Fall aus 1a9f270 — Commit nach `;` statt nach `&&`.
  assert block-pipe '{"tool_input":{"command":"make gates > /tmp/g.log 2>&1; git commit -m x"}}'   "Commit nach ; am Gate-Lauf"
  assert block-pipe '{"tool_input":{"command":"make gates > /tmp/g.log 2>&1\ngit commit -m x"}}'   "Commit nach Zeilenumbruch"
  assert ok    '{"tool_input":{"command":"make gates > /tmp/g.log 2>&1; echo \"EXIT=$?\""}}'  "vorgeschriebene Form ohne Commit"
  # SL-004: ein Heredoc, das das Muster ZITIERT, ist kein Aufruf. Diese Fixture
  # trifft das Muster beinahe und prueft es damit wirklich (Lehre slice-058).
  assert ok    '{"tool_input":{"command":"git commit -F - <<EOF\nfix: make gates > log 2>&1; git commit -m x wird abgelehnt\nEOF"}}'  "Gate-Muster nur im Heredoc zitiert"
  assert ok    '{"tool_input":{"command":"make doc-repair | git apply -"}}'     "doc-repair liefert einen Patch auf stdout"

  # ── Drift-Waechter fuer die GATES-Liste (slice-059, R-057-F1) ───────────────
  # Ohne ihn ist die Liste eine Momentaufnahme: ein neues Pruef-Target waere
  # ungeschuetzt, und niemand saehe es. `gate-consistency` gleicht Doku gegen
  # Makefile ab — diese Liste gegen nichts. Jetzt gegen beide Make-Fragmente.
  #
  # NICHT_PRUEFEND: Targets, deren Ausgabe bestimmungsgemaess weiterverarbeitet
  # wird oder die keinen Pruef-Exit-Code tragen. Sie DUERFEN in eine Pipe —
  # `make doc-repair | git apply -` ist der vorgesehene Aufruf, und ein Guard,
  # der ihn blockiert, wuerde umgangen statt befolgt.
  #
  # `image-scan` steht hier, weil sein Ausgang bestimmungsgemaess WEITERVERARBEITET
  # wird: make normalisiert jeden Fehlschlag auf Exit 2, also liest der Workflow den
  # Ausgang aus dem LOG (ADR-0037). Ein Pipe-Verbot schuetzte einen Exit-Code, den
  # niemand auswerten kann (slice-124).
  #
  # `slice-mv` steht hier, weil es NICHT prueft, sondern BEWEGT: sein Exit-Code sagt
  # "Bewegung geglueckt", nicht "Bestand in Ordnung". Ein Pipe-Verbot darauf schuetzte
  # keinen Befund — es gibt keinen (slice-118).
  NICHT_PRUEFEND="help doc-help doc-doctor doc-repair doc-trace compile build arch-graph a-check a-check-graph record-gates hooks slice-mv image-scan"
  gates_liste="$(sed -n '/const GATES = new Set(\[/,/\]);/p' "$0" | grep -oE '"[a-z][a-z0-9-]*"' | tr -d '"' | tr '\n' ' ')"
  alle_targets="$(grep -hoE '^[a-z][a-z0-9-]*:' Makefile d-check.mk 2>/dev/null | tr -d ':' | sort -u)"
  for t in $alle_targets; do
    case " $NICHT_PRUEFEND " in *" $t "*) continue ;; esac
    case " $gates_liste " in
      *" $t "*) ;;
      *) echo "guard-selftest FAIL: Pruef-Target '$t' fehlt in der GATES-Liste (Regel 2 greift dort nicht)" >&2
         echo "  -> aufnehmen, oder in NICHT_PRUEFEND begruenden, falls seine Ausgabe gepiped werden soll" >&2
         fail=1 ;;
    esac
  done

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
case "$verdict" in
  block-toolchain) emit_block ;;
  block-pipe)      emit_block_pipe ;;
esac
# Pass-Fall: keine Ausgabe — die normale Permission-Prüfung übernimmt.
