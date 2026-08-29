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

- **Template-Provenienz-Zeiger bleiben auf `v3.5.2`.** **Fünf** Stellen sagen „übersetzt aus /
  Form nach `.harness/baseline/v3.5.2/templates/…`" — `docs/plan/planning/slice.template.md`
  (zweimal), `docs/plan/carveouts/carveout.template.md`, `docs/reviews/README.md`,
  `.harness/skills/reviewer.md` und der Sensors-Kommentar in `harness/README.md`. Sie zeigen auf
  die Form, aus der das jeweilige Artefakt **tatsächlich** entstanden ist. Sie umzuhängen, bevor
  der Form-Vergleich gelaufen ist, würde eine Konformität behaupten, die niemand geprüft hat.
  Etappe C.
- **Historische Aussagen werden nicht umgeschrieben.** Closure-Notizen in `done/` und
  Review-Reports nennen `v3.5.2` als das, was damals galt. Das ist korrekt und bleibt. Ebenso die
  Provenienz-Passagen in den Adaptions-Einträgen — sie gehören in den Durchgang der Etappe B.
- **Eine Ausnahme, mit Grund:** die Beleg-Stelle in
  [slice-093](../open/slice-093-regelwerk-statt-kurs.md) wandert mit. Sie ist kein
  Form-Zeiger und keine historische Aussage, sondern eine Messung am **gerade adoptierten** Stand;
  bewegt sich der Stand und die Zitatstelle nicht, wird aus dem Beleg eine Falschaussage. Der
  Wortlaut ist in beiden Ständen identisch, nur die Zeilennummern verschieben sich (13–14 → 18–19)
  — das ist im Slice jetzt ausgewiesen.
- **Keine Bewertung des Deltas.** Welcher Adaptions-Eintrag durch den neuen Stand welchen der
  fünf Ausgänge nimmt, ist Etappe B — hier wird nichts vorweggenommen.

## 5. DoD

- [x] `.harness/baseline/v5.12.0/{regelwerk,templates}/` + `SHA256SUMS` aus dem Release-ZIP
      materialisiert; `make regelwerk-check` grün und weist den alten Stand als ungeprüft aus —
      Beleg: Target-Ausgabe mit Exit-Code.
- [x] Die Stand-Deklaration nennt an allen drei Stellen `v5.12.0`, `AGENTS.md` §1 verweist nur auf
      Abschnitte, die es in diesem Stand **gibt**, und `conventions.md` §Baseline beschreibt den
      Ist-Stand statt „Etappe A von vier" — Beleg: Diff.
- [x] `make gates` (und bei Abschluss `make verify`) grün — **Ausgabe in eine Datei**, Exit-Code
      getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** 51 Dateien unter `.harness/baseline/v5.12.0/` neben den 42 des Vorgängers, ein
reproduzierbares Integritäts-Manifest, und die Stand-Deklaration an ihren drei Stellen auf
`v5.12.0` — inklusive der Korrektur zweier Aussagen, die schon vor diesem Slice falsch waren: die
Auswahlregel in `AGENTS.md` §1 nannte einen Abschnitt, den es im neuen Stand nicht mehr gibt, und
`conventions.md` §Baseline behauptete seit dem 2026-08-09 „Etappe A von vier", während `welle-12`
die Etappen A–F längst geschlossen hatte.

**Lerneintrag — Form: geschärfte Regel.** *Wählt ein Sensor aus mehreren gleichartigen Kandidaten,
gehört der nicht geprüfte in die Ausgabe — sonst ist „Integrität ok" eine Aussage über einen
Ausschnitt, den niemand benennt.* Heute lagen zum ersten Mal zwei vendored Stände nebeneinander,
34 Tage nachdem `regelwerk-check` genau dafür gebaut wurde, und der Lauf meldete unaufgefordert
*„2 vendored Staende gefunden; geprueft wird der hoechste (v5.12.0) — ungeprueft bleiben: v3.5.2"*.
**Ohne diese Zeile wäre der Lauf ununterscheidbar von einem, der beide Stände geprüft hat**, und
genau das ist der Zustand, in dem eine Migration steckt.
*Präzisierung, damit die Notiz nicht mehr behauptet als sie belegt:* die zweite Hälfte desselben
Fixes — Versions- statt Zeichenordnung (`sort -V`) — hat heute **nichts** gerettet, denn `v5.12.0`
gewinnt auch in Zeichenordnung gegen `v3.5.2`. Sie greift erst beim nächsten Sprung innerhalb
derselben Major-Nummer.

**Zwei beobachtbare Closure-Kriterien:**

1. `make regelwerk-check` **Exit 0**, meldet 51 Dateien für `v5.12.0` gegen `SHA256SUMS` **und**
   die Gegenrichtung Baum → Manifest, und weist `v3.5.2` namentlich als ungeprüft aus.
2. `doc-check` prüft nach dem Vendoring **192** Dateien, nicht 243 — der neue Fremdtext-Baum ist
   vom bestehenden `scan.ignore` gedeckt, ohne dass die Konfiguration angefasst wurde.

**Offene Risiken und ihr Ausgang:**

- *Fünf Form-Zeiger stehen weiter auf dem alten Stand* — Ausgang: **Folge-Slice**, Etappe C
  (§4 nennt sie einzeln).
- *Der alte vendored Baum bleibt liegen* — Ausgang: **Folge-Slice**, Etappe C; er fällt, wenn der
  Form-Review durch ist. Das ist Vorschrift, kein Rückstand.
- *Der adoptierte Stand gilt ab sofort, die Artefakte sind ihm aber noch nicht angeglichen* —
  Ausgang: **Folge-Slices** B, C, D. Damit das kein stiller Zustand ist, hält `conventions.md`
  §Baseline ausdrücklich fest, dass die deklarierten Adaptionen bis zum Durchgang in Kraft bleiben.

**Folge-Slices:** Etappe B, wird unmittelbar geschnitten; C und D danach.

## 7. Sub-Area-Modus

Berührt werden **Vendored Baseline** (`.harness/baseline/`, laut Modus-Tabelle ausdrücklich
**kein Modus** — externer Fremdtext) und die Harness-Deklaration. Alle berührten Sub-Areas mit
Modus sind GF.
