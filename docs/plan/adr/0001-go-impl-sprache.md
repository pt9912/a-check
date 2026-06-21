# ADR-0001 — Go als Implementierungssprache

- **Status:** Accepted
- **Datum:** 2026-06-21
- **Bezug:** [AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus), [AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze), [AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit), [AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)
- **Schärft:** — (die Spezifikation ist sprachneutral; die Sprachwahl koppelt an keine Spec-§ — würde sie geändert, bliebe das Technik-Stratum unverändert. Diese ADR realisiert die in [SPEC-DIST-001](../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution) geforderte — und durch [ADR-0004](0004-distribution-image-mk.md) verbindlich gemachte — statische/distroless Laufzeitform, ohne dass die Spezifikation die Sprache nennt; die ADR ist die Begründungs-Schicht *unter* den Spec-Straten.)
- **Supersedes:** —

## Kontext

`a-check` ist ein sprach-agnostisches CLI, das *fremde* Quellbäume
text-heuristisch prüft (Lastenheft §1) und als **netzloses,
distroless/static** Image verteilt wird ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
[AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk)),
byte-deterministisch ([AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus))
und reproduzierbar ([AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).
[`AGENTS.md`](../../../AGENTS.md) §3.1 setzt Go bereits als Fundament voraus
(„ein statisches, sprach-agnostisches Binary … kein Host-Go"); diese ADR
liefert die Begründung dahinter, statt sie als „ist halt so" zu lassen.

## Optionen

1. **Go** — ein einziges, statisch gelinktes Binary ohne Runtime, passt
   direkt auf ein distroless/static Image; starke Standardbibliothek
   (`regexp`, Datei-Walk), schnelle reproduzierbare Builds, deterministische
   Ausgabe. Trade-off: kein so reichhaltiges Parser-Ökosystem wie Rust —
   für die Text-Heuristik aber nicht nötig (siehe [ADR-0002](0002-text-heuristische-extraktion.md)).
2. **Rust** — ebenfalls statisches Binary, stärkere Korrektheitsgarantien
   (Speichersicherheit). Trade-off: höhere Build-Komplexität/-zeit, steilere
   Kurve, kein Hermetik-/Determinismus-Vorteil gegenüber Go. Der
   Korrektheits-Vorteil greift hier kaum: das Fehlerrisiko eines
   text-heuristischen Tools liegt in den Mustern/der Config (siehe
   [ADR-0002](0002-text-heuristische-extraktion.md),
   [ADR-0003](0003-config-modell-a-check-yml.md)), nicht in Speichersicherheit.
3. **Python** — schnelle Entwicklung, reiche Bibliothek. Trade-off: kein
   statisches Binary, Interpreter-/Runtime-Abhängigkeit bricht die
   distroless/static-Hülle und erschwert Reproduzierbarkeit und
   Determinismus. Verworfen.

## Entscheidung

**Go.** Ein statisch gelinktes Binary ohne Runtime-Abhängigkeiten passt
unmittelbar auf ein distroless/static Image (Hermetik-Anforderung),
unterstützt netzlosen Lauf und byte-deterministische Ausgabe
(Determinismus-Anforderung), und der digest-gepinnte Image-/Distributionspfad
(Reproduzierbarkeit) wird trivial — das konkrete Distributionsmodell dazu
trägt [ADR-0004](0004-distribution-image-mk.md). Die **Sprachwahl** Go ist
zudem konsistent mit dem Schwester-Tool
[`d-check`](https://github.com/pt9912/d-check), das ebenfalls in Go geschrieben
ist und dessen Harness-Form `a-check` adoptiert — beides in-Repo belegt in
[`harness/conventions.md`](../../../harness/conventions.md)
(§Adoptierte Konventions-Quellen).

## Konsequenzen

- Die Go-Toolchain läuft **in Docker** ([`AGENTS.md`](../../../AGENTS.md)
  §3.1 Docker/make-only); kein Host-Go.
- **Fitness Function / Gate** (entsteht mit slice-003, [`AGENTS.md`](../../../AGENTS.md)
  §4): `make lint` (golangci-lint, Suppression-Verbot §3.2), `make test`
  (Determinismus-Test zu [AC-QA-01](../../../spec/lastenheft.md#ac-qa-01--determinismus)),
  distroless/static Build im `make`-Pfad. **Entscheidungsspezifischer Beleg:**
  ein Build-Gate prüft, dass das Artefakt ein *statisch gelinktes* Binary ohne
  dynamische Abhängigkeiten ist (`file`/`ldd`-Prüfung im distroless-Build) —
  der Test, der die Go-/static-Wahl maschinell festnagelt. Eine ADR ohne
  durchsetzendes Gate bliebe Absichtserklärung (Regelwerk Modul 4).
- Build-Reproduzierbarkeit: gepinnte Go-Version + Modul-Pins; Image
  digest-gepinnt ([AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit)).
