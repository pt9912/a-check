# Welle welle-17 — Benannte `shapes`-Dialekte `gomod` und `json` — Closure-Notiz

> **Zitier-Form** *(bleibt stehen — Norm, kein Ausfüll-Hinweis).* Dieses
> Artefakt friert ein; was es zitiert, bewegt sich weiter. Deshalb: **Kennung,
> nicht Adresse** — `slice-<Kennung>` statt seines Lifecycle-Pfads, `make <target>`
> statt eines Links auf die Sensor-Datei, eine Baseline-Stelle als
> `v6.13.0` · `regelwerk/<datei>.md` §<Abschnitt> statt als Link.

**Welle:** welle-17-shapes-generischer-dialekt
**Abschluss:** 2026-10-07
**Verantwortlich:** Claude (Planner); Abnahme beim Maintainer.

## Was wurde geliefert?

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3.

- **Die Sollform je Datei prüft `go.mod` und JSON-Dateien wie `package.json`** — mit `v0.22.0`,
  Digest `sha256:12e961f799e6d50d25cf68f1a0b230cf222f51c2360bf91933cd7489174a26a9`.
- slice-212: Vertrag — Messung an realen Manifesten,
  [AC-FA-RULE-012](../../../../spec/lastenheft.md#ac-fa-rule-012) (Lastenheft 0.29.0),
  [ADR-0042](../../adr/0042-shapes-benannte-dialekte-gomod-json.md) `Accepted`, Spezifikation
  0.36.0; die einzeilige, umkehrbare Meldung aller `shape-*`-Befunde als benannte Spec-Lücke.
- slice-213: Dialekt `gomod` und die einzeilige Meldung mit Art-Präfix bei `shape-unused`;
  Spezifikation 0.37.0.
- slice-214: Dialekt `json`; Spezifikation 0.38.0 verlangt gültiges JSON nach RFC 8259.
  Benutzerhandbuch 1.45.

## Was hat funktioniert?

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3.

- **Lexik aus der Primärquelle statt aus dem Gedächtnis.** Für `gomod` fand der Review im Code
  keinen Durchlass und keinen Fehlalarm — anders als bei `kotlin` in welle-16, wo vier
  Lexik-Lücken aus erinnerten Sprachregeln kamen.
- **Messung an realen Dateien in jedem Slice:** 8 `go.mod` der Konsumenten vor dem Vertrag, 153
  `go.mod` und 233 `package.json` gegen die Implementierung, ohne einen falschen Exit 2.
- **Den Vertrag heben statt die Zusage kürzen.** Als die Doku für `json` mehr versprach als der
  Code hielt, ging die Gültigkeit nach RFC 8259 als Plan-Änderung in die Spezifikation; ein Prüfer
  schloss sechs Findings mit einer Regel.

## Was ging anders als geplant?

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3.

- **Der Start-Trigger wurde ersetzt:** statt „zweiter Konsument mit Nicht-Kotlin-Manifest" das
  Maintainer-Wort (Drift-Log der Roadmap). Die Messung an realen Konsumenten-Dateien trug den
  fehlenden Beleg teilweise.
- **Der `json`-Vertrag zählte Fehlerfälle auf**, statt Gültigkeit zu verlangen; die erste
  Implementierung war fail-safe, aber nicht fail-closed. Erst der Review fand das.
- **Testkommentare sagten wieder mehr zu als ihre Assertion** — in jedem der zwei
  Implementierungs-Slices, auch im Nachlauf. Die Klasse steht bei 4×; ihre Verkörperung
  (slice-215) ist noch nicht geschrieben.
- **Delta-Reviews fanden neue Wortlaut-Findings** in den Fixes selbst (Plan-Kopf, Zuordnung einer
  Fehlerbedingung zur RFC-Regel). Fixes brauchen denselben Prüf-Durchgang wie der Erst-Code.

## Steering-Loop-Einträge

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3 — hier stehen nur Beobachtungen, die im Register 3× erreicht
haben.

- **`BEO-SPEC/ausgabe-einzeilig-nicht-zugesichert`** (3×) — verkörpert als benannte Spec-Lücke in
  der Spezifikation 0.36.0 ([SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung), einzeilige und umkehrbare Meldung), seit slice-212.
- **`BEO-GATE/testbeschreibung-weiter-als-assertion`** (4×) — Ausgang *geplant*: slice-215
  verkörpert die Regel als vierte Mess-Regel. Kein Sensor möglich (Urteil über Text); die
  Prosa-Form ist nicht ausgeschöpft, sondern noch nicht geschrieben.

## Beobachtungs-Register (Zeiger)

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Das Beobachtungs-Register — der Zähler wird nicht hier gepflegt.

Der Zähler steht im [Beobachtungs-Register](../observations/README.md). Neu oder fortgeschrieben
in dieser Welle: `BEO-SPEC/ausgabe-einzeilig-nicht-zugesichert` (3×, verkörpert) ·
`BEO-GATE/testbeschreibung-weiter-als-assertion` (4×, geplant) ·
`BEO-SPEC/vertrag-auf-einen-konsumenten-belegt` (2×) ·
`BEO-SPEC/fehlerfaelle-statt-grammatik-der-quelle` (1×) ·
`BEO-ADAPT/rekursiver-pruefer-ohne-tiefengrenze` (1×).

**Lese-Schritt:** kein offener Eintrag steht bei 3× oder mehr ohne Ausgang.

**Bestand in `open/`, Lese-Schritt:** slice-013 und slice-045 haben diese Closure unverändert
überstanden; keine gemeinsame Ursache, je **bestätigt** — beide trigger-gebunden, kein Trigger
gefeuert. slice-215 entstand in dieser Welle (slice-213) und ist **bestätigt**: Er trägt den
Ausgang *geplant* der 4×-Beobachtung.

## Folge-Slices

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 3 — derivativ.

- slice-215 (Mess-Regel „Testkommentar gegen Assertion", wellenlos) — in `open/`.

## Verifikation

Regeln dieser Sektion: Baseline-Regelwerk `modul-06-roadmap.md`
§Wellen-Closure-Prozedur, Schritt 1 — Replay-Ersatz nach
[MR-028](../../../../harness/conventions.md#mr-028).

- Alle drei Slices in `done/`; `make ci` Exit 0 und `make verify` Exit 0 auf dem Stand nach dem
  Re-Pin (lokal, Ausgabe in Dateien, Exit-Codes getrennt gelesen); `make preflight` Exit 0 vor dem
  Push (44 Commits Range).
- CI auf `main` grün: Lauf 37585746599. Release-Pipeline grün: Lauf 37585872239 (Tag `v0.22.0`).
- Digest gegen die Registry gegengeprüft: `docker pull` des Digests gelingt, `RepoDigests` nennt
  `sha256:12e961f799e6d50d25cf68f1a0b230cf222f51c2360bf91933cd7489174a26a9`, OCI-Label
  `org.opencontainers.image.version` = `0.22.0`. Re-Pin danach mit `make gates` Exit 0
  (`gate-consistency`: Pins konsistent).
- `make doc-immutable RANGE=v0.21.0..HEAD`: 0 Befunde.
- **Gegenprobe an realen Manifesten gegen das veröffentlichte Image:** Kopien des `go.mod` von
  d-check und des `package.json` von m-trace, je eine `shapes`-Konfiguration
  (`allow-statements`, `unused: fail`, Einträge in Quellform). Grün (Exit 0, 0 Befunde). Mit
  `github.com/evil/lib v1.0.0` im direkten `require`-Block rot (Exit 1) mit dem **gesehenen** Befund
  `go.mod:5: shape-unlisted: require (github.com/go-git/go-billy/v5 v5.9.1;github.com/go-git/go-git/v5 v5.19.2;gopkg.in/yaml.v3 v3.0.1;github.com/evil/lib v1.0.0)`;
  mit `"evil-lib": "1.0.0"` in `devDependencies` rot (Exit 1) mit
  `package.json:18: shape-unlisted: "devDependencies":{…,"vitest":"^4.1.5","evil-lib":"1.0.0"}` —
  je dazu `shape-unused` für den nicht mehr passenden Eintrag. Nach dem Zurücksetzen wieder grün.
- Coverage gesamt 96,1 % (Schwelle 90 %). Carveouts: keine.
- **Trigger-Audit, vier Klassen:** Carveout 0 offen (Verzeichnis führt nur die README) ·
  bootstrap-aware Gate 0 (keines im Bestand) · ADR 0 fällig (die Trigger von
  [ADR-0042](../../adr/0042-shapes-benannte-dialekte-gomod-json.md) — weiteres Format, feinere
  JSON-Einheit — sind nicht eingetreten; [ADR-0041](../../adr/0041-shapes-sollform-je-datei.md)
  Punkt 10 ist mit dieser Welle eingelöst) · Hard Rule 0 (keine mit Auflösungs-Trigger) — dazu die
  sieben aktiven MR-Einträge: 0 offen.
- **Drei Paarungen:** Anker — der verkörperte Eintrag liegt in einer versionierten Spec
  (benannte Spec-Lücke), kein `liegt in`-Feld zu paaren · Folge-Slice — slice-215 existiert in
  `open/` · Register — alle fünf genannten Pfade existieren als Verzeichnis mit nicht leerem
  `evidence/` (`make verify-observations` im `verify`-Lauf grün).
