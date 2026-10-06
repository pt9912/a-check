# Welle welle-16: Sollform je Datei (`shapes:`)

**Lifecycle:** Diese Datei entsteht bei der **Eröffnung** der Welle und liegt
flach unter `docs/plan/planning/`; bei Closure wandert sie per `git mv` nach
`done/` (neben ihre `welle-16-results.md`). Der Zustand ist die
Verzeichnis-Position — kein Status-Feld.

**Zielmeilenstein:** kein Meilenstein-Bezug.

**Verantwortlich:** Claude (Planner); Abnahme beim Maintainer.
**Datum:** 2026-10-06.

---

## 1. Welle-Ziel

Ein Adopter kann für eine **benannte Datei** — Leitfall `build.gradle.kts` des
Fachkern-Moduls — festlegen, dass sie **nur** ausdrücklich erlaubte Anweisungen
enthält; jede andere ist ein Befund (fail-safe statt Verbotsliste). Gespiegelt
an den Gegenprobe-Fällen des CR (siehe §2): Leerraum, Kommentare und
Zeilenumbrüche ändern das Urteil nicht; eine zusätzliche Anweisung — auch in
einem erlaubten Block, auch mit Zeilenumbruch vor `{`, auch „im Kommentar
begonnen" — ist rot; ein verbotenes Muster nur in Kommentar oder Zeichenkette
ist grün; nicht zerlegbare Dateien sind Exit 2, nie still grün.

Die Regel steht **neben** dem Konstrukt-Monopol
([AC-FA-RULE-011](../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak)),
nicht in ihm: dort ist eine Form in einer Zone erlaubt, hier ist in einer Datei
nur die Liste erlaubt.

## 2. Trigger (Welle startet)

- Der Change Request „Positivliste von Anweisungen je Datei (Sollform)" liegt
  vor — eingereicht vom Maintainer in diesem Gespräch (2026-10-06), Anlass: ein
  Adopter will das Fachkern-Modul gegen Fremdabhängigkeiten und Plugins sichern.
- Maintainer hat die Welle angewiesen (dieses Gespräch, 2026-10-06).

## 3. Closure-Trigger (Welle schließt)

Das *Mehr* gegenüber den Slice-DoDs ist der **Release-Beleg**: Der CR verlangt
die Veröffentlichung „wie bisher mit Digest", und die steht in keiner einzelnen
DoD.

- Alle Slices aus §4 liegen in `done/`.
- `make ci` Exit 0 auf dem finalen Stand — Ausgabe in eine Datei, Exit-Code
  getrennt geprüft, nie in eine Pipe (Replay-Ersatz laut
  [MR-028](../../../harness/conventions.md#mr-028)).
- Ein Release-Tag mit der neuen Regelart ist veröffentlicht, das Image
  digest-gepinnt in [`a-check.mk`](../../../a-check.mk) — der Tag ist eine
  Maintainer-Aktion nach [`docs/user/releasing.md`](../../user/releasing.md).
- **Gegenprobe am Leitfall:** eine `shapes`-Konfiguration für eine
  `build.gradle.kts` in Adopter-Form lädt gegen das veröffentlichte Image, ist
  grün, und wird rot, sobald eine Fremdabhängigkeit eingefügt wird — der
  Befund wird gesehen, nicht nur der Exit-Code.
- Ergebnis-Notiz `done/welle-16-results.md` geschrieben.

## 4. Slices in dieser Welle

| Slice | Titel | Bezug |
|---|---|---|
| [slice-209](done/slice-209-shapes-spec-first.md) | Spec-first: Anforderung, ADR, Spezifikation für `shapes:` — zur Abnahme | [AC-FA-RULE-011](../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak) (Abgrenzung), [AC-FA-CONF-001](../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml) |
| [slice-210](open/slice-210-shapes-kotlin-allow-statements.md) | Dialekt `kotlin` + `mode: allow-statements` | neue Regel-Anforderung aus slice-209 |
| [slice-211](open/slice-211-shapes-exact-und-unused.md) | `mode: exact` (`expect:`) + Opt-in-Befund für nie treffende Einträge | neue Regel-Anforderung aus slice-209 |

**Reihenfolge ist Abhängigkeit:** slice-210 und slice-211 starten erst, wenn
slice-209 in `done/` liegt und der Maintainer die Anforderung abgenommen hat —
Spec führt, Code folgt (Sub-Area `SPEC` ist Greenfield).

## 5. Abhängigkeiten

- Blockiert: keine. Der generische Dialekt (konfigurierbare Kommentar-,
  Zeichenketten- und Trennzeichen für `go.mod`, `package.json`, …) baut auf
  dieser Welle auf, steht aber als eigene Zeile unter *Nächste Wellen* der
  [Roadmap](in-progress/roadmap.md) mit eigenem Trigger.
- Wird blockiert von: keiner Welle. Keine offene Welle, kein Slice in
  `in-progress/`.

## 6. Out-of-Scope für diese Welle

- **Generischer Dialekt.** Belegt ist ein Konsument mit Kotlin-Bedarf; ein
  zweiter mit Nicht-Kotlin-Manifest ist der Trigger (Roadmap, *Nächste
  Wellen*). Ohne ihn würde ein Konfigurations-Raum für Kommentar-/
  String-/Trennzeichen geraten statt gemessen.
- **Abhängigkeitsgraphen, Build-Werkzeug-Aufrufe, Gradle-Semantik.** Laut CR
  selbst ausgeschlossen: a-check bleibt hermetisch
  ([AC-QA-02](../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze));
  was ein erlaubtes Plugin selbst einträgt, gehört in eine Prüfung im Build.
- **`ignore: [version-literals]`.** Abnahme-Entscheid in slice-209:
  gestrichen — ein unscharf umrissener Ausschluss ist die Lücke, durch die eine
  Fremdabhängigkeit grün durchgeht; den Zweck deckt ein voll verankerter
  Regex-Eintrag.
- **Ein Warn-Level.** a-check kennt keines; der Hinweis `shape-unused` aus dem
  CR wird deshalb Opt-in-**Befund** (slice-211), kein neuer Schweregrad.
- **Die fail-closed Import-Allowlist je Schicht** — der verwandte, gated Faden
  aus dem Out-of-Scope von
  [AC-FA-RULE-011](../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak).
  Die ADR aus slice-209 grenzt beide ab; gebaut wird er hier nicht.

## 7. Closure-Notiz

Ergebnis: <folgt bei Closure — Zeiger auf `welle-16-results.md`>
Zähler: <folgt bei Closure — Zeiger auf `../observations/`>
