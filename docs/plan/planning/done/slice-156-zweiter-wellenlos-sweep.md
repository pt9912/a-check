# slice-156 — Zweiter Wellenlos-Sweep: Puffer auf zwei Slices zurückgeschnitten

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Hinweis 2026-09-05 („In `docs/plan/planning/done` liegen noch Wellen und
Slices. Reviews-Ordner ist leer") — derselbe Sweep-Bedarf wie in
[slice-153](../done/wellenlos/slice-153-archive-wave-welle-feld-und-titel-form.md), der Puffer war seither
auf sechs Slices (`150`…`155`) gewachsen, weil die Etappen dazwischen (`154`, `155`) neue flache
Slices erzeugten, ohne erneut zu fegen. [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Bestandspflege, kein Vertrags-Fakt.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

`slice-150`…`153` archivieren, `154`/`155` als Puffer belassen — wieder auf die von d-check
beobachtete Größenordnung (ein bis zwei flache Slices).

## 2. Vorab geprüft

Repo-weiter `grep` nach `#anchor`-Links in `slice-150`…`153`: **ein** Treffer, `slice-151` zitiert
`slice-150` per Abschnitts-Anker — beide werden in diesem Sweep archiviert, die Zeile verschwindet
mit `slice-151`s eigenem Stub (self-resolving, dieselbe Lehre wie
[slice-148](../done/wellenlos/slice-148-archivierung-etappe-d.md) §4). Kein externer Anker-Bruch.

## 3. Ausführung

`archive-wave SLICE=<id> APPLY=1` für `slice-150`, `151`, `152`, `153`, in dieser Reihenfolge.
Keine Fehlschläge, keine unbehebbaren toten Review-Verweise. Stichprobe: `slice-150`s Stub trägt
`**Welle:** ohne Welle` und keinen doppelten Gedankenstrich — beide in
[slice-153](../done/wellenlos/slice-153-archive-wave-welle-feld-und-titel-form.md) behobenen Funde halten.

## 4. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor.

## 5. DoD

- [x] Anker-Bruch-Risiko vorab per `grep` geprüft, nicht angenommen (§2).
- [x] Vier Slices archiviert, Stub-Form stichprobenartig verifiziert (§3).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** `docs/plan/planning/done/` trägt jetzt wieder zwei flache Slices (`154`, `155`)
statt sechs — derselbe Puffer-Umfang wie bei d-check.

**Lerneintrag — Form: geschärfte Regel.** *Ein einmaliger Sweep hält den Puffer nur so lange
klein, wie danach kein weiterer wellenloser Slice entsteht — jede der vier Folge-Etappen nach
`slice-153` (`154`, `155`, plus dieser hier) hat selbst wieder einen flachen Slice erzeugt, ohne
dass ein Mechanismus das automatisch nachzieht.* Das ist exakt die in
[slice-143](../done/wellenlos/slice-143-archivierung-delta-analyse.md) §6 benannte,
bislang nicht beauftragte „Etappe D+1"-Selbst-Archivierungsroutine: ohne sie bleibt Fegen ein
wiederkehrender manueller Vorgang, der bei jeder erneuten Nachfrage neu auffällt. Der Bedarf ist
mit diesem Slice ein zweites Mal beobachtet — noch nicht als eigene Beobachtung registriert, weil
er kein Sensor-Kollisions-Fund ist (das deckt bereits
[`BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig`](../observations/BEO-GATE/archiv-sensor-vorpruefung-unvollstaendig/observation.md)),
sondern ein Workflow-Bedarf.

**Zwei beobachtbare Closure-Kriterien:**

1. `ls docs/plan/planning/done/*.md | grep -v -- '-results.md$' | wc -l` → `2`
   (`slice-154`, `slice-155`); `ls docs/plan/planning/done/wellenlos/*.md | wc -l` → `96`
   (`92` aus [slice-153](../done/wellenlos/slice-153-archive-wave-welle-feld-und-titel-form.md) `+ 4`
   aus diesem Slice).
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Keines notiert — der Sweep ist mechanisch und vollständig
belegt. Der im Lerneintrag benannte Workflow-Bedarf (Selbst-Archivierung ab Etappe D+1) bleibt
unbeauftragt und ist kein Risiko *dieses* Slice, sondern der bislang unbeauftragte Vorschlag aus
[slice-143](../done/wellenlos/slice-143-archivierung-delta-analyse.md).

**Folge-Slices:** keine vergeben.

## 7. Sub-Area-Modus

Berührt: **Planungs-Harness** (`docs/plan/planning/done/`) — Greenfield: Form steht in der
vendored Vorlage, `doc-structure` prüft sie.
