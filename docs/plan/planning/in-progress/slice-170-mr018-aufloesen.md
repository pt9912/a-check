# slice-170 — `MR-018` auflösen: der Rückbau-Trigger ist eingetreten

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese
Datei liegt — eines von `open/`, `next/`, `in-progress/`, `done/`. Er
wechselt nur durch `make slice-mv` ([`AGENTS.md`](../../../../AGENTS.md)
§3.3/§5).

**Welle:** ohne Welle (der Closure-Trigger wäre die eigene DoD — kein
repo-weites Mehr).

**Bezug:** [`MR-018`](../../../../harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md)
nennt seinen eigenen Rückbau-Trigger; dieser Slice vollzieht ihn.
Beobachtung:
[`BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration`](../observations/BEO-HARNESS/rueckbau-kandidat-ueberlebt-baseline-migration/observation.md).

**Berührte Spec-Stellen:** — *(keine)* — Harness-Konventionen ohne
Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim
Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:**
2026-09-06.

---

## 1. Ziel

[`MR-018`](../../../../harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md) auflösen. Die Adaption war eine **Provenienz-Korrektur**: der
Wortlaut des Rollenwechsel-Absatzes in [`AGENTS.md`](../../../../AGENTS.md)
§6 stammte aus einem Kurs-Template, das a-check nicht vendored hatte —
also gab es keinen netzlos auflösbaren Anker, auf den das Feld
`Ersetzt-Baseline-Regel` hätte zeigen können. Dieser Grund ist entfallen.

## 2. Analyse (vor der Umsetzung)

**Gemessen, nicht angenommen** — `diff` der beiden vendorten Stände:

```text
diff .harness/baseline/v6.0.0/templates/AGENTS.template.md \
     .harness/baseline/v6.2.0/templates/AGENTS.template.md
236a237,243
> Dieser Workflow deckt ausschließlich die Implementer-Rolle ab. Schritt 8
> ist der Rollenwechsel, kein Abschluss: Bericht → Handoff an Reviewer …
```

Der Absatz, um den es geht, liegt seit
[`slice-167`](../done/slice-167-etappe-a-vendoring-v620.md) **netzlos im
Repo**. Damit sind beide im Eintrag genannten Trigger eingetreten:

| Trigger aus dem Eintrag | Zustand |
|---|---|
| „Rückbau-Kandidat, sobald der Stand vendored ist und auf das dann vendorte Template gezeigt werden kann" | eingetreten (slice-167) |
| Auflösungs-Trigger: „die nächste Baseline-Migration, die `modul-08` oder den Rollenwechsel-Absatz in `AGENTS.template.md` inhaltlich ändert" | eingetreten — die Migration hat den Absatz **hinzugefügt** |

**Warum kein Nachfolge-Eintrag.** Was nach dem Rückbau von der Abweichung
bliebe, sind zwei a-check-eigene Sätze (`fork`-Ausschluss, `BEO`-Anker).
Die **schärfen** `modul-08` §Kontext-Trennung und ersetzen dort nichts;
eine ausgefüllte Ziel-Form, die mehr sagt als ihre Vorlage, ist keine
Adaption — sonst bräuchte jede gefüllte Vorlage einen Eintrag. Präzedenz
für eine Auflösung **durch Ereignis** statt durch Nachfolger ist
[`MR-003`](../../../../harness/conventions.md#mr-003).

**Was der Fund über den Prozess sagt.** `welle-14-results.md` schreibt,
der Eintrag löse sich „erst mit der jeweils nächsten Baseline-Migration" —
notiert **nach** slice-167, also nachdem diese Migration den Trigger
gezogen hatte. Die Ursache ist strukturell: der **Trigger-Audit** der
Wellen-Closure (`modul-06`, Schritt 2) prüft *Carveout · bootstrap-aware
Gate · ADR*; `MR`-Einträge stehen nicht darin. Ein `MR` mit
Auflösungs-Trigger hat damit keinen Wächter und wird nur gefunden, wenn
jemand ihn von Hand liest — hier: der Maintainer, im Durchgang durch alle
aktiven Einträge.

## 3. Umsetzung

1. [`MR-018`](../../../../harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md) per `git mv` nach
   [`harness/conventions/done/`](../../../../harness/conventions/done/) —
   Datei-Inhalt **unverändert** (Adaptions-Block-Disziplin: an einem
   akzeptierten Eintrag wird nichts nachträglich geändert).
2. [`harness/conventions.md`](../../../../harness/conventions.md): Zeile
   verlässt *Aktive Adaptionen*, tritt in *Aufgelöste Adaptionen* mit dem
   Auflösungs-Ereignis; der Anker `mr-018` zieht mit um und hält alle
   bestehenden Verweise gültig.
3. Verweise auf den alten Dateipfad repo-weit nachgezogen. **Nicht**
   angefasst: die Review-Reports — sie nennen den Pfad als Inline-Code in
   ihrer Eingangs-Kontext-Liste, nicht als Link, und sind Lauf-Belege eines
   vergangenen Standes.

## 4. Definition of Done

- [ ] [`MR-018`](../../../../harness/conventions/done/MR-018-review-pflicht-v610-wortlaut.md) liegt in `conventions/done/`, Inhalt unverändert; beide
      Tabellen in `conventions.md` nachgezogen, Anker `mr-018` erhalten.
- [ ] Jeder **verlinkende** Verweis auf den alten Pfad zeigt auf den neuen;
      `make doc-check` belegt es.
- [ ] Unabhängiger Review durchgeführt (Report unter `docs/reviews/`).
- [ ] `make gates` grün.
- [ ] `make verify` grün.
- [ ] Jedes Risiko trägt einen Ausgang.

## 5. Trigger

**Start** (`open` → `in-progress`): Maintainer-Anweisung im Durchgang durch
die aktiven Adaptionen (2026-09-06), WIP-Limit frei.

**Rückführungen:** keine vorgesehen — zwei Liefer-Punkte, eine Schicht.

## 6. Closure-Trigger

DoD vollständig, `make gates`/`make verify` grün, Closure-Notiz
geschrieben.

## 7. Risiken und offene Punkte

- *Der Anker `mr-018` wird beim Umzug in die zweite Tabelle vergessen —
  dann brechen die Verweise aus [`MR-019`](../../../../harness/conventions/MR-019-review-dod-opt-in.md) und `slice-162`, die auf
  `conventions.md#mr-018` zeigen* — Ausgang bei Closure.
- *Die Auflösung entfernt einen von sechs `v6.0.0`-Ankern, löst
  `.harness/baseline/v6.0.0/` aber nicht ab — fünf bleiben stehen; wer die
  Zahl für den Fortschritt hält, verwechselt einen Schritt mit dem Ziel* —
  Ausgang bei Closure.

## 8. Closure-Notiz

*(wird beim Übergang nach `done/` geschrieben; Lerneintrag — Form: wird
dort benannt.)*

## 9. Sub-Area-Modus

**Vorgelagert — Sub-Area-Wahl prüfen:** eine Sub-Area berührt —
**Harness-Einstieg** (`harness/`), Greenfield, Schwelle ≥ 2/3 erfüllt.

**Vorgelagert — offene Beobachtungen sichten:** `BEO-HARNESS/` über die
Verzeichnisliste geprüft.
`rueckbau-kandidat-ueberlebt-baseline-migration` ist einschlägig und
bekommt mit diesem Slice sein **zweites** Auftreten (bisher 1×,
slice-141); die Schwelle 3× ist damit nicht erreicht. Kein weiterer
Eintrag unter `BEO-HARNESS/` betrifft diesen Slice.

**Alle berührten Sub-Areas GF** — kein Begründungsblock nötig.
