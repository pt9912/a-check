# slice-020 — Python-Sprach-Backend (welle-06-sprach-backends)

**Status:** done (2026-07-02). Abnahme erteilt (Entscheide A–D gemäß Empfehlung); Umsetzung +
`make gates`/`make ci` + Multi-Linsen-Review (4 Linsen, R1) + Delta-Re-Review erledigt; Synthese
[`docs/reviews/2026-07-02-slice-020-python-backend.md`](../../../reviews/2026-07-02-slice-020-python-backend.md).
**Welle:** welle-06-sprach-backends (zweites Backend-Inkrement nach
[slice-014](../done/slice-014-java-backend.md)).
**Bezug:** erweitert [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
um Python; innerhalb [ADR-0002](../../adr/0002-text-heuristische-extraktion.md)
(text-heuristisch, **kein** neuer ADR); schärft
[SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion).
Die Symbol→Schicht-Auflösung ist **bereits geliefert**: Python ist fester-Wurzel-dotted
([ADR-0014](../../adr/0014-resolution-roots.md)/[ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md),
[slice-015](../done/slice-015-resolution-roots.md)) — mit dem Backend wird `python` als
`languages`-Schlüssel zulässig und damit `resolution`-fähig; die Sprach-Validierung
([AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
[slice-017](../done/slice-017-unbekannte-sprache-exit2.md)) bezieht die Menge aus der
Backend-Registry, **kein** Schema-Delta.
[Roadmap welle-06](../in-progress/roadmap.md). **Trigger:** Polyglot-Bestand
(Python-Repos), Maintainer-Priorität 2 (nach dem Go/C++-Kern).

> **Hinweis:** Entwurf zur Abnahme. Die in §3 als Code-Fence gesetzten AC-Texte sind
> unverbindlich — gültig erst nach Freigabe in [`spec/lastenheft.md`](../../../../spec/lastenheft.md).
> DoD §5 offen; Entscheidungen §6 **vor** der Umsetzung zu treffen.

---

## 1. Ziel

Ein **Python**-Backend für die Import-Extraktion, analog Java/Kotlin — damit Python-Repos
ihre Hexagon-Architektur über a-check + `.a-check.yml` prüfen können, ohne dass die
Engine sich ändert. Reine Extraktions-Erweiterung; keine neue Regel, kein neues
Modell-Konzept, kein neuer Auflösungs-Modus.

## 2. Problem

a-check v0.4.0 wählt das Extraktions-Backend über `languages` ∈ {`cpp`, `go`, `rust`,
`kotlin`, `java`} (`extract.go`-Registry). `languages: {python: …}` bricht heute — korrekt,
[slice-017](../done/slice-017-unbekannte-sprache-exit2.md) — mit Exit 2. Python hat
**zwei** absolute Import-Formen:

- `import a.b.c` (auch `import a.b as x`) — dotted, fast deckungsgleich mit Kotlin;
- `from a.b import c` (auch `from a.b import c as d`) — **neu**: das Modul steht nach
  `from`, nicht nach `import`; ein zweites Muster ist nötig.

**Relative Importe** (`from . import x`, `from ..pkg import y`) lösen gegen den *Ort der
importierenden Datei* auf — das ist das Signal des **reservierten `relative`-Modus**
([slice-015 §4](../done/slice-015-resolution-roots.md), Folge-ADR) und gehört nicht in
diesen Slice (§6 Entscheid B).

## 3. Entwurf (zur Abnahme)

### 3.1 Anforderungs-Erweiterung — AC-FA-EXTRACT-001 (Python)

```text
AC-FA-EXTRACT-001 (erweitert um Python): die Backend-Liste wird um Python ergänzt —
C++ (#include), Go (import), Rust (use/extern crate), Kotlin (import), Java (import,
inkl. import static), Python (import und from … import). Beide Python-Formen liefern
den gepunkteten Modulpfad; ein Alias (as x) und die hinter `from … import` stehenden
Namen werden nicht als Symbol gewertet.

Neue/ergaenzte Akzeptanzkriterien (zu den bestehenden Happy/Boundary/Negative):
- Happy (Python import): Given `import myapp.adapters.db`, when das Python-Backend
  laeuft, then liefert es das Symbol `myapp.adapters.db`.
- Boundary (Python from): Given `from myapp.adapters import db`, when das Python-Backend
  laeuft, then liefert es `myapp.adapters` (den Modulpfad nach `from`; die importierten
  Namen werden nicht expandiert).
- Boundary (Alias): Given `import myapp.adapters as ad`, when das Python-Backend laeuft,
  then liefert es `myapp.adapters` (das `as ad` wird nicht gewertet).
- Negative bleibt sprach-agnostisch (import-aehnliche Zeile in Kommentar/String wird
  nicht gewertet — bestehende Heuristik-Grenze AC-QA-02); eine `#`-Kommentarzeile
  matcht die verankerten Muster nie.

Out-of-Scope: relative Importe (`from .`/`from ..`) — anderes Aufloesungs-Signal
(reservierter relative-Modus), werden nicht extrahiert (dokumentierte Heuristik-Grenze
AC-QA-02); Mehrfach-Import in einem Statement (`import a, b`) wird nur als Erst-Treffer
(`a`) gegriffen; `__init__`-Re-Export-Semantik; Toolchain-Backends (importlib/AST);
Import-aehnliche Zeilen in Docstrings (bestehende String-Grenze).
```

### 3.2 Versions-Bump

Lastenheft + Spezifikation **0.10.0 → 0.11.0** (neue Sprach-Unterstützung, MINOR).
`fünf → sechs Sprachen` über die Doku (README, Benutzerhandbuch, **und**
`spec/architecture.md` — [ARC-003](../../../../spec/architecture.md) zählt die Zielsprachen konkret auf; Review-R1 fing,
dass die Entwurfs-Formulierung „nur falls sie zählt" genau dort zum Übersehen einlud).

### 3.3 Auflösung: kein Schema-Delta (Rezept dokumentieren)

Python ist gepunktet, hat aber — anders als JVM — **kein** natürliches
Reverse-Domain-Präfix. Die `.`→`/`-Normalisierung ist per
[ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md) an gesetztes `package_base`
gebunden (gepunktete Sprache signalisiert sich darüber). Das **Rezept** für die üblichen
Python-Layouts ist damit heute ausdrückbar:

```yaml
# src-Layout: src/myapp/{domain,ports,adapters}/…, Importe `myapp.…`
resolution:
  python: {mode: fixed-root, roots: ["src/myapp"], package_base: "myapp"}
# flaches Single-Package-Layout: myapp/… an der Scan-Wurzel
#   python: {mode: fixed-root, roots: ["myapp"], package_base: "myapp"}
```

(`package_base` strippt das Top-Package, `roots` fügt seinen Verzeichnispfad wieder an —
Komposition statt neuem Schalter.) Das Rezept wird im Benutzerhandbuch als
Python-Beispiel dokumentiert. Ein eigenes `dotted: true`-Flag (Schema-Delta + Folge-ADR
zu [ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md)) ist **vertagt** (§6
Entscheid D).

## 4. Umsetzungsplan

1. `internal/adapter/driven/extract/extract.go`: Felder `pyImp`
   (`^\s*import\s+([A-Za-z_][A-Za-z0-9_.]*)`) und `pyFrom`
   (`^\s*from\s+([A-Za-z_][A-Za-z0-9_.]*)\s+import\b`) in `newAdapter`;
   Registry-Eintrag `"python": … lineMatches(src, a.pyImp, a.pyFrom)`. Die
   Sprach-Validierung ([slice-017](../done/slice-017-unbekannte-sprache-exit2.md)) und
   die `resolution`-Zulässigkeit ziehen automatisch nach (Registry = einzige Quelle).
2. Tests (`extract_test.go`) — inkl. der **Mutanten-Boundary-Tests** aus dem
   [slice-014-Lerneintrag](../done/slice-014-java-backend.md#7-closure-notiz)
   (Regex-Backends brauchen Tests gegen die spezifischen Mutanten): Happy dotted;
   `from … import` → Modulpfad; Alias (`as`); `importlib.reload(x)` (Keyword-als-Präfix,
   kein Match); `from`-Zeile ohne `import` (kein Match); `# import os`
   (Kommentar-Grenze); `import a, b` (Erst-Treffer-Grenze); relative Importe (kein
   Match); Mehrfach-Whitespace.
3. Resolution-Integrationstest (`rules_test.go` oder bestehende Testform): Python-Datei
   mit `import myapp.adapters.db` + `resolution`-Rezept (§3.3) → löst auf die
   Adapter-Schicht auf (z. B. `core-impurity` aus einer Domänen-Datei) — belegt, dass
   Backend + gelieferter `fixed-root`-Modus zusammen greifen
   ([AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)
   Happy-Auflösung, jetzt für Python).
4. Spec: `spec/lastenheft.md`
   ([AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
   + Bump 0.11.0 + Historie), `spec/spezifikation.md`
   ([SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion):
   Python-Muster + Backend-Menge `{cpp, go, rust, kotlin, java, python}` + Bump).
5. „fünf → sechs Sprachen"-Sweep: README (Sprach-Aufzählung), Benutzerhandbuch (§1/§4
   `languages`-Enum + Python-Beispiel inkl. `resolution`-Rezept §3.3 + Historie-Zeile),
   `spec/architecture.md` ([ARC-003](../../../../spec/architecture.md)-Sprachliste — Review-R1-Nachtrag, slice-014-Präzedenz).
   **Nicht** [ADR-0002](../../adr/0002-text-heuristische-extraktion.md): `Accepted` ⇒
   immutable (AGENTS §3.5).
6. `make gates`; **Multi-Linsen-Review** (Code · Vertrag/Spec · Test · Regelwerk)
   schriftlich nach `docs/reviews/`; Verifikation gegen DoD/Spec; Commits; Roadmap-/
   CHANGELOG-Nachtrag.

## 5. Definition of Done

- [x] [AC-FA-EXTRACT-001](../../../../spec/lastenheft.md#ac-fa-extract-001--sprach-backends-für-die-import-extraktion)
      um Python erweitert (Happy `import` + Boundary `from`/Alias + Out-of-Scope inkl.
      Subpaket-Form-Grenze), Bump 0.11.0 + Historie; [SPEC-EXTRACT-001](../../../../spec/spezifikation.md#spec-extract-001--import-extraktion)
      nachgezogen (Muster + Backend-Menge + Python-nicht-C-gestrippt-Präzisierung).
- [x] `extract.go`: `pyImp`/`pyFrom` + Registry-Eintrag + `prepSource` (Review-R1: kein
      C-Strip für Python); `make arch-check` (Dogfooding) grün (0 Befunde).
- [x] Tests: Happy/`from`/Alias + Mutanten-Boundary (§4.2, um Review-R1-Mutanten erweitert:
      Anker/Zeichenklassen/Einrückung/Glob-String) + Resolution-Integration (§4.3, pinnt
      `.`→`/` über Mehrsegment-Import).
- [x] „fünf → sechs Sprachen"-Sweep vollständig (README, Benutzerhandbuch inkl.
      `resolution`-Rezept, `spec/architecture.md` [ARC-003](../../../../spec/architecture.md) — **ohne**
      [ADR-0002](../../adr/0002-text-heuristische-extraktion.md), immutable).
- [x] `make gates` + `make ci` grün; Multi-Linsen-Review (4 Linsen) + Delta-Re-Review bestanden
      ([Synthese](../../../reviews/2026-07-02-slice-020-python-backend.md)); Verifikation gegen
      DoD/Spec (empirische Fixture-Gegenprobe).
- [x] Closure: **reiner** `git mv` nach `done/` (AGENTS §3.3, getrennt von Inhalts-Edits);
      **2 beobachtbare Kriterien** + **Lerneintrag** (§7).

## 6. Offen / Entscheidungen zur Abnahme

> **Abnahme (2026-07-02):** Entscheide A–D gemäß Empfehlung bestätigt.

- **Entscheid A — `from`-Import-Symbol:** Modulpfad nach `from` (`from a.b import c` →
  `a.b`; Empfehlung) vs. Modul+Namen expandieren (`a.b.c`, bei Listen mehrere Symbole).
  *Empfehlung: nur Modulpfad — eine Zeile ⇒ höchstens ein Symbol je Muster (bestehende
  Heuristik-Linie), deterministisch, und die Schicht-Auflösung braucht nur den Modulpfad.
  Grenzfall `from myapp import adapters` (Modul löst auf kein Layer-Glob) wird als
  Heuristik-Grenze dokumentiert, nicht expandiert.*
- **Entscheid B — relative Importe:** nicht extrahieren + als Grenze dokumentieren
  (Empfehlung) vs. extrahieren-und-unauflösbar-lassen. *Empfehlung: nicht extrahieren —
  relative Importe sind das Auflösungs-Signal des reservierten `relative`-Modus
  (Folge-ADR); ein extrahiertes, nie auflösbares Symbol wäre Rauschen ohne Regel-Wert.
  Ehrlich ausgewiesen in Out-of-Scope + Benutzerhandbuch
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)-Linie).*
- **Entscheid C — ADR?** Kein neuer ADR — Extraktion innerhalb
  [ADR-0002](../../adr/0002-text-heuristische-extraktion.md), Auflösung innerhalb
  [ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md). *Empfehlung: bestätigt —
  Python fügt keine neue Architektur-Entscheidung hinzu.*
- **Entscheid D — Auflösung ohne Schema-Delta:** `package_base`-Rezept (§3.3, Empfehlung)
  vs. neues `dotted: true`-Flag (bräuchte Folge-ADR zu
  [ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md) + Schema-Delta in
  [AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml)).
  *Empfehlung: Rezept — kein Schema-Delta auf Verdacht; das `dotted`-Flag wird erst
  gebaut, wenn ein realer Python-Pilot ein Layout zeigt, das das Rezept nicht trägt
  (z. B. flaches Multi-Package-Layout ohne gemeinsames Top-Package).*
- **Risiko/Notiz — synthetische Verifikation:** wie bei Java gibt es noch keinen
  benannten Python-Pilot; die DoD-Verifikation läuft gegen eigene Fixtures. Die Sprache
  bleibt gated-geliefert: Aktivierung, wenn ihr Repo real pilotiert.
- **Risiko/Notiz — `stripComments` ist C-orientiert** (`//`, `/* */`): für Python
  harmlos — `#`-Kommentarzeilen matchen die verankerten Muster nie, und ein gestripptes
  `//` (Floor-Division) kann kein `^\s*import`/`^\s*from` erzeugen. Docstrings mit
  import-ähnlichen Zeilen sind die **bestehende** String-Grenze
  ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)),
  kein Python-Spezifikum. Ein Python-eigenes `#`-Stripping ist **nicht** nötig
  (kein Verhaltens-Delta) und bleibt out-of-scope.

## 7. Closure-Notiz

**Abschluss (2026-07-02).** slice-020 (welle-06 — zweites Sprach-Backend) umgesetzt und gate-belegt.

- **Gate-Beleg:** `make gates` grün — `lint` 0 issues, alle Test-Pakete `ok`, `coverage-gate`
  95,90 % (≥ 90 %), `arch-check` **0** (Dogfooding), `doc-check` 0,
  gate-consistency/guard-selftest/record-gates ok; `make ci` (inkl. `image-test`) grün.
- **Code:** `pyImp`/`pyFrom`-Regexes + Registry-Eintrag `"python"`; `prepSource` nimmt Python vom
  C-Kommentar-Stripping aus (Review-R1-MAJOR: eine `/*`-Bytefolge in einem Python-String — z. B. das
  Glob `"**/*.py"` — fraß sonst alle Folge-Imports, falsch-grün; empirisch reproduziert und
  gegengeprobt). Auflösung unverändert über den gelieferten `fixed-root`-Modus
  ([ADR-0016](../../adr/0016-resolution-sprach-parametrisch.md)) per `package_base`-Rezept.
- **Review (4 Linsen + Delta):** Code · Vertrag/Spec · Test · Regelwerk, danach Delta-Re-Review des
  Fix-Diffs — 5 × MAJOR, 5 × MINOR gefunden und gefixt (u. a. [ARC-003](../../../../spec/architecture.md)-Sweep-Lücke, staler
  Handbuch-Header, ungepinnte `.`→`/`-Konvertierung). Kein BLOCKER.
  Doc: [Review](../../../reviews/2026-07-02-slice-020-python-backend.md).
- **Verifikation (gegen DoD/Spec):** alle DoD-Haken erfüllt; die Python-ACs
  (Happy `import`/Boundary `from`/Boundary Alias) sind durch Tests gepinnt; Lastenheft/Spezifikation
  0.11.0 konsistent; Benutzerhandbuch 1.13 dokumentiert Rezept + Grenzen (inkl. des
  `resolution`-Nachtrags aus [slice-015](../done/slice-015-resolution-roots.md) und des
  Exit-2-Currency-Fixes zu [slice-017](../done/slice-017-unbekannte-sprache-exit2.md)).
- **2 beobachtbare Kriterien:** (1) `TestPythonFixedRootResolution` + Container-Gegenprobe — Python-Domäne
  mit `GLOB = "**/*.py"`, `from myapp.adapters import db` und `import myapp.adapters.db` unter dem Rezept
  ⇒ Exit 1 mit genau 2 × `core-impurity` (Zeile 2 + 3; pinnt Glob-String-Robustheit **und**
  Mehrsegment-`.`→`/`). (2) `TestCheckLanguagesUnknown` — `languages: {ruby: …}` ⇒ Exit 2 mit
  `unbekannte Sprache "ruby" (cpp|go|java|kotlin|python|rust)`; das gepinnte Meldungsformat beweist
  `python` in der Backend-Menge.
- **Lerneintrag (geschärfte Regel):** Ein neues Sprach-Backend reviewt auch die **geteilte
  Vorverarbeitung**, nicht nur seine Regexes — das C-förmige `stripComments` lief unbedingt über alle
  Sprachen und wurde erst durch das sechste Backend (andere Kommentar-/String-Syntax) zum
  Falsch-grün-Kanal; Kommentar-Syntax ist eine **Sprach-Eigenschaft** (jetzt `prepSource`). Zweitens:
  Sweep-Listen benennen die zählenden Stellen **explizit** („Architektur nur falls sie zählt" lud zum
  Übersehen von [ARC-003](../../../../spec/architecture.md) ein). Der slice-014-Lerneintrag (Mutanten-Boundary-Tests) bestätigte sich
  erneut: vier Linsen fanden fünf weitere überlebende Mutanten-Klassen.
- welle-06 bleibt **offen** (Kandidaten C#/TypeScript, je mit eigenem Auflösungs-Modus per Folge-ADR);
  slice-020 ist ihr zweites Inkrement.
