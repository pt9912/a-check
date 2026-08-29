# slice-094 — Etappe A: Baseline `v5.12.0` committet vendoren

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** keine `AC-*`/`ADR-*` — Harness-/Konventions-Änderung ohne Vertragsberührung, wie
[slice-047](../done/slice-047-baseline-vendoring.md).
**Bezug:** Etappe **A** aus [slice-092 §6](../done/slice-092-regelwerk-v5120-delta-analyse.md), am
2026-08-29 per Maintainer-Wort gezogen („Leg los A-D").
[`MR-006`](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
bleibt in Kraft und behält seinen Gegenstand. [Roadmap](../in-progress/roadmap.md).

---

## 1. Was Etappe A tut — und was nicht

Sie materialisiert den neuen Stand **neben** dem alten und zieht die Stand-Deklaration nach.
**Nichts weiter:** kein Adaptions-Durchgang (B), keine Form (C), keine neue Mechanik (D).

**Das alte Verzeichnis bleibt stehen.** Das ist keine Bequemlichkeit, sondern Vorschrift der
Baseline selbst (`modul-02` §Freshness-Audit): die neue Referenz-Form ist die Vergleichsgrundlage
für den Form-Review, und *„das alte Verzeichnis fällt erst, wenn der Review durch ist"*. Weil der
Vendoring-Pfad `<tag>`-gescopt ist, liegen beide Stände nebeneinander und
`diff -r .harness/baseline/v3.5.2/templates .harness/baseline/v5.12.0/templates` ist netzlos
möglich — genau das braucht Etappe C.

`make regelwerk-check` ist auf diesen Zustand vorbereitet: es wählt den höchsten Stand per
`sort -V` und **weist den ungeprüften Rest namentlich aus**, statt ihn stillschweigend zu
übergehen (slice-049, Review-Fund R-049-F4).

## 2. Betroffene Module

- `.harness/baseline/v5.12.0/{regelwerk,templates}/` + `SHA256SUMS` — neuer vendored Baum.
- Stand-Deklaration an ihren **drei** Stellen: `harness/conventions.md` §Baseline und
  §Adoptierte Konventions-Quellen, `AGENTS.md` §1, `harness/README.md` §Guides.

Eine Schicht (Harness-Deklaration), plus ein Verzeichnis Fremdtext.

## 3. Auszuführende Gates

`make regelwerk-check` (kein Gate, Wartung: Integrität fail-closed **und** die Gegenrichtung
Baum → Manifest), dann `make gates` — tragend ist `doc-check`, weil die Deklarations-Verweise
umziehen. Zum Abschluss `make verify`.

**Kein neuer Sensor**, also keine Negativ-Probe. Der Integritätsnachweis ist reproduzierbar
(`sha256sum -c`), nicht behauptet.

## 4. Was bewusst nicht getan wird

- **Template-Provenienz-Zeiger bleiben auf `v3.5.2`.** Vier Stellen sagen „übersetzt aus /
  Form nach `.harness/baseline/v3.5.2/templates/…`" — in `docs/plan/planning/slice.template.md`,
  `docs/reviews/README.md`, `.harness/skills/reviewer.md` und dem Sensors-Kommentar in
  `harness/README.md`. Sie zeigen auf die Form, aus der das jeweilige Artefakt **tatsächlich**
  entstanden ist. Sie umzuhängen, bevor der Form-Vergleich gelaufen ist, würde eine
  Konformität behaupten, die niemand geprüft hat. Etappe C.
- **Historische Aussagen werden nicht umgeschrieben.** Closure-Notizen in `done/` und
  Review-Reports nennen `v3.5.2` als das, was damals galt. Das ist korrekt und bleibt.
- **Keine Bewertung des Deltas.** Welcher Adaptions-Eintrag durch den neuen Stand welchen der
  fünf Ausgänge nimmt, ist Etappe B — hier wird nichts vorweggenommen.

## 5. DoD

- [ ] `.harness/baseline/v5.12.0/{regelwerk,templates}/` + `SHA256SUMS` aus dem Release-ZIP
      materialisiert; `make regelwerk-check` grün und weist den alten Stand als ungeprüft aus —
      Beleg: Target-Ausgabe mit Exit-Code.
- [ ] Die Stand-Deklaration nennt an allen drei Stellen `v5.12.0`, `AGENTS.md` §1 verweist nur auf
      Abschnitte, die es in diesem Stand **gibt**, und `conventions.md` §Baseline beschreibt den
      Ist-Stand statt „Etappe A von vier" — Beleg: Diff.
- [ ] `make gates` (und bei Abschluss `make verify`) grün — **Ausgabe in eine Datei**, Exit-Code
      getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Berührt werden **Vendored Baseline** (`.harness/baseline/`, laut Modus-Tabelle ausdrücklich
**kein Modus** — externer Fremdtext) und die Harness-Deklaration. Alle berührten Sub-Areas mit
Modus sind GF.
