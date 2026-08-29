# slice-093 — Das vendored Regelwerk ist die maßgebliche Fassung, nicht der Kurs

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** keine `AC-*`/`ADR-*` — eine Adaption ggü. der Baseline, deklariert im Adaptions-Block
von [`harness/conventions.md`](../../../../harness/conventions.md).
**Bezug:** Maintainer-Ansage 2026-08-29 („Wir benutzen das Regelwerk, nicht den Kurs!") während
[slice-092](../done/slice-092-regelwerk-v5120-delta-analyse.md).
[Roadmap](../in-progress/roadmap.md).

---

## 0. Trigger — **negativ eingetreten am 2026-08-29**

**Beginn war:** Bestätigung des Maintainers, dass die Ansage als **deklarierte Adaption** landen
soll (und nicht als bloße Arbeitsanweisung für die laufende Sitzung). Der Unterschied ist
beobachtbar: eine Adaption ändert zwei Sätze im Repo und bekommt einen Eintrag, eine
Arbeitsanweisung ändert nichts.

**Die Antwort war „nein".** Der Trigger ist damit entschieden, nicht offen — der Slice wird
geschlossen, ohne umgesetzt zu werden. Die DoD-Haken bleiben leer, weil nichts von dem entstanden
ist, was sie fordern; die Closure-Notiz sagt, warum.

Der Slice wartet **nicht** auf die Entscheidung über die `v5.12.0`-Migration — er gilt unter dem
heute adoptierten Stand genauso (§4).

## 1. Auslöser

Die Ansage steht gegen den Wortlaut an zwei Stellen im Repo:

| Stelle | Wortlaut |
|---|---|
| `AGENTS.md:33-36` | „Das vendored Regelwerk … trägt keine eigene Normativität: bei Konflikt gilt der Kurs" |
| `harness/conventions.md:50-53` | „**Extern (Lehrmaterial, maßgeblich für den Inhalt)** … bei Konflikt gilt der Kurs" |

**Beide sind korrekte Wiedergaben einer echten, adoptierten Baseline-Regel.** Sie steht als
Rumpftext im vendored Regelwerk selbst
(`.harness/baseline/v5.12.0/regelwerk/README.md`, Zeilen 18–19): *„Es trägt keine eigene
Normativität: maßgeblich für den Inhalt bleibt der Kurs."* Netzlos nachprüfbar, kein Zitat aus
zweiter Hand. Der Wortlaut ist im vorherigen Stand identisch (dort Zeilen 13–14) — die Klausel hat
den Sprung `v3.5.2` → `v5.12.0` unverändert überstanden, dieser Slice hängt also nicht an ihm.

*(Dieselbe Aussage steht zusätzlich im `conventions.template.md` der Baseline — dort aber in einem
HTML-Kommentar, also als Bedienhinweis. Die Fundstelle der **Norm** ist der Regelwerk-Rumpf oben;
der Kommentar ist ihre Dopplung, nicht ihr Zuhause.)*

**Warum die Abweichung trägt: die Regel ist in diesem Repo nicht ausführbar.** Sie verlangt im
Konfliktfall einen Blick in den Kurs. Der Kurs ist **nicht vendored**, das Repo ist per Hard Rule
netzlos ([`AGENTS.md`](../../../../AGENTS.md) §3.1, Scans mit `--network none`), und die stehende
Anweisung lautet, ihn nicht zu lesen. Eine Konfliktregel, die kein Agent hier auswerten kann, ist
nicht bloß unbenutzt — sie lädt zu genau dem Griff nach außen ein, den die Vendoring-Entscheidung
abgeschafft hat.

## 2. Betroffene Module

- `harness/conventions.md` — neuer Eintrag im Adaptions-Block (nächste freie Nummer), plus der
  Wortlaut in §Adoptierte Konventions-Quellen.
- `AGENTS.md` §1 — derselbe Satz.

Eine Schicht (Harness-Konventionen), zwei Dateien.

## 3. Auszuführende Gates

`make gates` — tragend sind `doc-check` (die Linkpflicht auf Adaptions-Kennungen ist über
[`.d-check.yml`](../../../../.d-check.yml) gegatet, ein neuer Eintrag zieht neue Verweise nach
sich) und `gate-consistency`. Zum Abschluss `make verify`.

**Kein neuer Sensor**, also keine Negativ-Probe: Gegenstand ist eine Deklaration, kein Zustand,
den ein Gate künftig hält.

## 4. Was bewusst nicht getan wird

- **Keine Vorwegnahme des `v5.12.0`-Modells.** Der Eintrag entsteht in der heute geltenden
  Inline-Form, nicht als eigene Datei unter `harness/conventions/`. Die Verzeichnis-Form gehört in
  Etappe C aus [slice-092 §6](../done/slice-092-regelwerk-v5120-delta-analyse.md) — falls die
  Migration kommt. **Vorbereitet** wird nur der Inhalt: die zu ersetzende Baseline-Regel ist oben
  benannt, damit Etappe B das neue Pflichtfeld nur noch übernimmt statt es zu recherchieren.
- **Der Kurs wird nicht bewertet.** Ob der Extrakt den Kurs irgendwo inhaltlich verfehlt, ist hier
  **nicht** gemessen und kann es auch nicht werden — der Kurs liegt nicht vor. Genau diese
  Unprüfbarkeit ist das Argument (§1), nicht ein Befund über die Qualität des Extrakts.
- **Kein Carveout, keine ADR.** Der Diskrepanz-Trichter ([`AGENTS.md`](../../../../AGENTS.md) §5)
  greift für Ausnahmen von *repo-eigenen* Regeln und Gates. Hier weicht das Repo von einem
  *Baseline-Default* ab — dafür ist der Adaptions-Block das benannte Instrument.

## 5. DoD

- [ ] Der Adaptions-Eintrag steht mit Geltungsbereich, Adaption, Begründung und
      Auflösungs-Trigger; die ersetzte Baseline-Regel ist mit Datei und Zeilen benannt — Beleg:
      Diff.
- [ ] Beide Fundstellen aus §1 tragen den neuen Wortlaut und verweisen auf den Eintrag; keine
      dritte Stelle behauptet weiterhin den Kurs-Vorrang — Beleg: `grep` über die Repo-Doku.
- [ ] `make gates` (und bei Abschluss `make verify`) grün — **Ausgabe in eine Datei**, Exit-Code
      getrennt geprüft, nie in eine Pipe.

## 6. Closure-Notiz

**Geliefert:** eine **Entscheidung**, kein Artefakt. Die Ansage „Wir benutzen das Regelwerk, nicht
den Kurs" bleibt Arbeitsanweisung; sie wird **nicht** als Adaption deklariert. `AGENTS.md` §1 und
`harness/conventions.md` behalten ihren Wortlaut, und die Klausel im vendored Regelwerk bleibt,
was sie ist: eine Provenienz-Aussage über die Ableitung des Extrakts, keine Leseanweisung an den
Agenten.

**Lerneintrag — Form: geschärfte Regel.** *Eine Arbeitsanweisung ist keine Konventions-Änderung,
und die Messung entscheidet das nicht.* Die Analyse in §1 war korrekt und ist es weiterhin: die
Klausel steht als Rumpftext im adoptierten Regelwerk, und sie verlangt im Konfliktfall etwas, das
dieses Repo netzlos nicht tun kann. Daraus folgt aber **nicht**, dass sie deklariert abweichend
gehört. *Weil* der Schluss von „diese Regel ist hier nicht ausführbar" auf „also weichen wir
erklärtermaßen ab" einen Schritt überspringt, den nur der Maintainer gehen kann: ob eine inert
gebliebene Klausel überhaupt stört. Sie stört nicht — ein Tie-Break, der nie gezogen wird, kostet
nichts. **Der Prüfsatz:** ein Befund begründet, dass eine Frage gestellt wird, nie schon ihre
Antwort.

**Zwei beobachtbare Closure-Kriterien:**

1. Im Repo hat sich **nichts** geändert: `AGENTS.md` §1 und `harness/conventions.md`
   §Adoptierte Konventions-Quellen tragen denselben Wortlaut wie vor diesem Slice, und der
   Adaptions-Block führt keinen neuen Eintrag. Das ist die Lieferung, nicht ihr Ausbleiben.
2. Der Slice liegt in `done/` statt gelöscht zu sein — eine still entfernte Datei wäre von einer,
   die es nie gab, nicht zu unterscheiden. Dieselbe Regel, die das Beobachtungs-Register für
   gestrichene Einträge setzt.

**Offene Risiken und ihr Ausgang:**

- *Der `git mv` dieses Slice hat drei Verweise in bereits geschlossenen Slices gebrochen* — sie
  zeigten auf `../open/`. `verify-slice-links` sah das nicht: es nimmt `done/` als Endzustand aus
  und richtet seine Invariante auf **wandernde** Slices, nicht auf Verweise **auf** sie. Gefangen
  hat es `doc-check`. Ausgang: **weiter offen**, als `BEO-008` im Beobachtungs-Register; die drei
  Verweise sind nachgezogen.

**Folge-Slices:** keine.

## 7. Sub-Area-Modus

Berührt werden die Harness-Konventionen und das Briefing. Beide liegen unter **keiner** Zeile der
Modus-Deklaration pro Sub-Area — dieselbe Lücke, die
[slice-091 §7](../done/slice-091-claude-md-auf-verweis-reduzieren.md) benannt hat. Alle berührten
Sub-Areas mit Modus sind GF.

*(Nachtrag beim Abschluss: die Lücke ist seit slice-101 geschlossen — die Sub-Area
**Harness-Einstieg** ist deklariert und trägt im Beobachtungs-Register die Kennung `BEO-001`.
Der Satz oben beschreibt den Stand, als dieser Plan geschrieben wurde.)*
