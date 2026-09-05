# slice-146 — `**Welle:**`-Feld für `welle-12`/`welle-13` nachgetragen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Vorbereitung für Etappe C aus [slice-143 §6](../done/slice-143-archivierung-delta-analyse.md#6-vorschlag-vier-etappen),
löst den in [slice-144 §4.2](../done/slice-144-archive-wave-werkzeug.md#4-zwei-funde-beim-dry-run--präzedenz-greift-aber-nicht-überall)
gemessenen Blocker. Präzedenz gelesen, nicht kopiert:
`d-check` [`6354e06`](https://github.com/pt9912) „Welle-Feld für vier historische Slices
nachgetragen" (`slice-191`) — gleiche Disziplin, andere Zahlen. [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Nachtrag eines bereits belegten Fakts, keine neue Aussage.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

39 `done/`-Slices (33 aus `welle-12`, 6 aus `welle-13`) trugen kein `**Welle:**`-Feld, obwohl beide
Wellen ihre Mitgliedschaft bereits **tabellarisch** in der eigenen Ergebnisnotiz führen
([`welle-12-results.md`](../done/welle-12-results.md) §Verifikation: „`slice-046` … `slice-067`" +
„`slice-068` … `slice-078`"; [`welle-13-results.md`](../done/welle-13-results.md) §Verifikation:
„`slice-081` … `slice-086`"). Nachgetragen aus dieser bereits bestehenden Quelle — kein neuer Fakt.

## 2. Betroffene Module

39 Dateien unter `docs/plan/planning/done/`, je eine Zeile `**Welle:** <welle-id>.` eingefügt direkt
nach dem `**Status:**`-Absatz — derselbe Platz und dieselbe Form wie bei a-checks eigenen frühen
Wellen (`slice-001`…`slice-022`, vgl. `**Welle:** welle-01-fundament.`). Eine Schicht
(Planungs-Harness).

## 3. Was genau nachgetragen wurde — und was nicht

- **`welle-12`: 33 Dateien**, `**Welle:** welle-12-regelwerk-migration.` — Beleg:
  [`welle-12-results.md`](../done/welle-12-results.md) §Verifikation, zwei zusammenhängende Bereiche
  (`slice-046`…`slice-067`, `slice-068`…`slice-078`). **Nicht** dazugenommen:
  [`slice-079`](../done/slice-079-gate-consistency-abloesen.md) und
  [`slice-080`](../done/slice-080-verify-abloesung-dcheck.md) — dieselbe Notiz nennt sie
  ausdrücklich **nicht** zur Welle gehörig („aus der Arbeit entstanden, haben aber einen anderen
  Gegenstand … und eigene Trigger").
- **`welle-13`: 6 Dateien**, `**Welle:** welle-13-konsumenten-befunde.` — Beleg:
  [`welle-13-results.md`](../done/welle-13-results.md) §Verifikation, „`slice-081` … `slice-086`".
- **Kein anderes Feld verändert.** Das historische `**Status:**`-Feld (z. T. schon in seinem
  eigenen Wortlaut korrigiert, z. B. `slice-050`/`slice-067`: „der Zustand ist das Verzeichnis …
  dieses Feld führt ihn bewusst nicht doppelt") bleibt unangetastet — derselbe Grundsatz wie beim
  `d-check`-Vorbild: ergänzt wird ein fehlendes Feld, nichts Bestehendes wird korrigiert.
- **Einfügepunkt:** direkt nach dem Ende des `**Status:**`-Absatzes (erste Leerzeile oder nächstes
  `**Feld:**` — beide Formen kommen im Bestand vor), mechanisch bestimmt, nicht von Hand pro Datei.

## 4. Gegenprobe: löst `archive-wave` die Wellen jetzt auf?

Dry-Run (kein `APPLY=1`, `:ro`-Mount, schreibt nichts):

- `make archive-wave WELLE=welle-12` — findet jetzt **33 Slices**, den Welle-Plan-Fall korrekt als
  „keiner — Welle vor der Plan-Datei-Konvention" und listet ~50 Dateien mit geplanten
  Verweis-Fixes.
- `make archive-wave WELLE=welle-13` — findet jetzt **6 Slices** und den eigenen Welle-Plan
  (`welle-13-konsumenten-befunde.md`, bereits in `done/`).

Beide vorher: **0 Slices**, fail-closed abgebrochen (gemessen in
[slice-144 §4.2](../done/slice-144-archive-wave-werkzeug.md)). Der Blocker ist damit **behoben,
nicht nur behauptet**.

## 5. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor.

## 6. Was bewusst nicht getan wird

- **Kein `APPLY=1`-Lauf.** Dieser Slice liefert die Vorbedingung für Etappe C, nicht ihre
  Ausführung — der Dry-Run in §4 bestätigt nur, dass die Vorbedingung jetzt erfüllt ist.
- **Keine weiteren Wellen nachgerüstet.** `welle-00` bis `welle-11` sind nicht Gegenstand dieser
  Migration — ihre Slices tragen das Feld größtenteils bereits (gemessen in
  [slice-144](../done/slice-144-archive-wave-werkzeug.md): 15 von damals grep-treffern), die
  verbleibenden sind Etappe-D-Bestand (Sammel-Archiv-Frage, [slice-143 §5](../done/slice-143-archivierung-delta-analyse.md#5-was-diese-analyse-nicht-geklärt-hat)).

## 7. DoD

- [x] 39 Dateien um `**Welle:**` ergänzt, Quelle je Wellen-Ergebnisnotiz, kein anderes Feld
      berührt (§3).
- [x] Gegenprobe: `archive-wave`-Dry-Run findet jetzt beide Wellen vollständig, vorher 0 Slices
      (§4).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 8. Closure-Notiz

**Geliefert:** der in Etappe A gemessene Blocker ist behoben — 39 `done/`-Slices tragen jetzt ein
`**Welle:**`-Feld, aus der jeweils eigenen Ergebnisnotiz belegt, kein neuer Fakt behauptet. Per
Dry-Run bestätigt: `archive-wave` löst beide Wellen jetzt vollständig auf.

**Lerneintrag — Form: geschärfte Regel.** *Ein fehlendes strukturiertes Feld ist kein fehlender
Fakt, wenn der Fakt bereits woanders im Repo steht — der Nachtrag ist dann eine Formalisierung,
keine neue Behauptung, und braucht darum auch keine neue Abnahme.* `d-check`s eigener Nachtrag
(`slice-191`) traf genau diese Unterscheidung explizit („eine Lücke aus der Zeit vor der
Feld-Konvention, kein echter Zuordnungs-Konflikt") und a-checks Fall liegt sogar eindeutiger: die
Quelle ist keine verstreute Prosa, sondern eine dedizierte, tabellarische Verifikations-Zeile in
der jeweiligen Wellen-Ergebnisnotiz. *Weil* die Alternative — den Nachtrag als „neue Aussage"
behandeln und dafür eine Abnahme verlangen — genau die Sorte Bürokratie wäre, die dieses Regelwerk
an anderer Stelle ausdrücklich vermeidet (Pfad-Berichtigung nach einem `git mv` ist ebenfalls keine
inhaltliche Änderung, obwohl sie den Dateiinhalt anfasst).

**Zwei beobachtbare Closure-Kriterien:**

1. `grep -l "welle-12-regelwerk-migration\." docs/plan/planning/done/*.md | wc -l` → `33`;
   dieselbe Zählung für `welle-13-konsumenten-befunde\.` → `6`.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Keine notiert — der Blocker ist mit der Gegenprobe aus §4
geschlossen, kein offener Rest.

**Folge-Slices:** keine vergeben. Etappe C (`welle-12`/`welle-13` mit `APPLY=1` archivieren) bleibt
aus slice-143 §6 vorgeschlagen und braucht eine eigene Kennung bei ihrer Eröffnung.

## 9. Sub-Area-Modus

Berührt: **Planungs-Harness** (`docs/plan/planning/done/`) — Greenfield.
