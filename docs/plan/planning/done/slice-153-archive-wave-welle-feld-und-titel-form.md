# slice-153 — archive-wave: `ohne Welle` statt leer, kein Doppel-Gedankenstrich

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Systematischer Abgleich aller a-check-Instanzdokumente gegen ihre v6.0.0-Templates
(Maintainer-Auftrag 2026-09-05, „Sind wir mit der Umstellung durch?"). Zwei von acht Funden
betreffen die von `tools/archive-wave` erzeugten Stubs gegen
[`archiv-stub-slice.template.md`](../../../../.harness/baseline/v6.0.0/templates/docs/plan/planning/archiv-stub-slice.template.md).
[Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Werkzeug-Fix + Bestandspflege, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

Zwei reale, verifizierte Funde beheben und alle betroffenen, bereits archivierten Stubs
nachziehen — nicht nur die künftig erzeugten.

## 2. Die zwei Funde

1. **`**Welle:**` leer statt `ohne Welle`.** Die Ziel-Form verlangt: „Ein Slice ohne
   Wellen-Zugehörigkeit behält `ohne Welle`." `ReadWelleField` liefert für einen Slice **ohne**
   `**Welle:**`-Feld (a-checks tatsächliches Muster — anders als d-checks eigenes „— wellenlos."-Feld)
   einen leeren String, der unverändert in den Stub floss. Betraf ausschließlich `ApplySlice`
   (Wellenlos-Modus) — der Wellen-Modus hat immer einen echten Feldwert, da er nur Slices mit
   passendem `**Welle:**` sammelt.
2. **Doppelter Gedankenstrich im Titel** (`# slice-135 — — Titel`). `ExtractTitle`s Regex
   (`h1RE`) war auf d-checks Überschriftenform „# Slice slice-190: Titel" (Doppelpunkt-Trenner)
   zugeschnitten; a-checks eigene Form „# slice-135 — Titel" (Gedankenstrich-Trenner) ließ den
   Trenner selbst im erfassten Titel — `SliceStub`/`SliceStubStandalone` hängen ihrerseits einen
   eigenen „ — "-Trenner davor.

Beide real gemessen (nicht angenommen): Fund 1 an allen 92 wellenlosen Stubs, Fund 2 an praktisch
allen Stubs beider Modi (a-checks Titel-Konvention ist durchgängig Gedankenstrich-basiert).

## 3. Behoben

- `h1RE` erweitert um einen optionalen Gedankenstrich-Trenner nach dem optionalen Doppelpunkt —
  wirkt für **beide** Überschriftenformen korrekt (Regressionstest ergänzt in `TestExtractTitle`).
- `ApplySlice` setzt `field = "ohne Welle"`, wenn `ReadWelleField` leer zurückgibt
  (Regressionstest `TestRunSlice_Apply_OhneWelleFeld`).

## 4. Nachzug des Bestands — nicht nur künftig, auch rückwirkend

Ein Patch der Stub-Texte wäre eine zweite, ungeprüfte Implementierung derselben Logik gewesen.
Stattdessen: **vollständige Neu-Archivierung** mit dem reparierten Werkzeug.

1. Alle 12 Wellen-Verzeichnisse und alle 92 wellenlosen Stubs per `unzip` aus ihrem jeweiligen
   `archiv.zip` in den Original-Zustand zurückgeholt (Byte-Identität mit dem committeten Original
   stichprobenartig gegen `git show` verifiziert, nicht angenommen).
2. Stub-Verzeichnisse/-Zips gelöscht.
3. `archive-wave WELLE=<id> APPLY=1` für alle 12 Wellen, dann `archive-wave SLICE=<id> APPLY=1`
   für alle 92 wellenlosen Slices — mit dem reparierten Werkzeug, in derselben Reihenfolge wie
   die ursprüngliche Archivierung.
4. Stichprobe nach dem Lauf: `grep -rl '^# slice-[0-9]* — —'` → 0 Treffer;
   `grep -rl '^\*\*Welle:\*\* $'` in `wellenlos/` → 0 Treffer.

## 5. Auszuführende Gates

`make gates`, `make archive-wave-test`, zum Abschluss `make verify`. Kein neuer Sensor.

## 6. DoD

- [x] Beide Funde real gemessen (nicht angenommen), mit Regressionstest je Fund (§2, §3).
- [x] Vollständiger Bestand (12 Wellen + 92 wellenlose Slices) mit dem reparierten Werkzeug neu
      archiviert, keine manuelle Text-Patch-Logik (§4).
- [x] `make gates`, `make archive-wave-test` und `make verify` grün — Ausgabe in eine Datei,
      Exit-Code getrennt geprüft, nie in eine Pipe.

## 7. Closure-Notiz

**Geliefert:** beide Stub-Form-Funde behoben, kompletter Bestand (104 archivierte Vorgänge)
mit dem reparierten Werkzeug neu erzeugt statt manuell gepatcht.

**Lerneintrag — Form: geschärfte Regel.** *Ein aus einem Schwester-Repo importiertes Werkzeug
trägt dessen Formkonventionen fest verdrahtet — der Doppelpunkt-Titel-Trenner war d-checks
Konvention, nie explizit für a-checks Gedankenstrich-Form geprüft, bis der reale Bestand es
zeigte.* Dieselbe Lehre wie beim Titel-Link-Bug ([slice-152](../done/slice-152-archive-wave-titel-link-und-sweep.md)):
ein importiertes Werkzeug ist erst dann wirklich geprüft, wenn es gegen den **eigenen**, nicht nur
den Herkunfts-Bestand gelaufen ist — Herkunfts-Tests allein reichen nicht, weil sie die Herkunfts-
Konvention testen, nicht die eigene.

**Zwei beobachtbare Closure-Kriterien:**

1. `grep -rl '^# slice-[0-9]* — —' docs/plan/planning/done/` → leer;
   `grep -rl '^\*\*Welle:\*\* $' docs/plan/planning/done/wellenlos/*.md` → leer.
2. `make gates`, `make archive-wave-test` und `make verify` je **Exit 0**, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Keines notiert — beide Funde sind mit Fix + vollständigem
Nachzug geschlossen, kein offener Rest. Der bereits verkörperte
[`BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig`](../observations/BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig/observation.md)
(seit `slice-152`) bleibt unverändert — diese beiden Funde sind Stub-**Form**-Fehler, keine
Sensor-Kollision der bereits verkörperten Klasse.

**Folge-Slices:** keine vergeben.

## 8. Sub-Area-Modus

Berührt zwei Sub-Areas:

- **Gate-/Werkzeug-Schicht** (`tools/archive-wave/`) — Greenfield: beide Funde getestet
  (`archive-wave-test`).
- **Planungs-Harness** (`docs/plan/planning/done/`) — Greenfield: Form steht in der vendored
  Vorlage, `doc-structure` prüft sie.
