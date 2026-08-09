# slice-085 — Diagnose: Schicht ohne auflösende Importe

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `CR-2` des Konsumenten-Einsatzes (2026-08-09);
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung).
**Bezug:** Welle [`welle-13-konsumenten-befunde`](../done/welle-13-konsumenten-befunde.md); Vorbild ist
die Abdeckungs-Diagnose aus [slice-043](../done/slice-043-schicht-abdeckung-sichtbar.md) /
[ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md).

---

## 0. Trigger

**Beginn:** sofort.

**Rückführungen:**

- `in-progress` → `open`: falls sich zeigt, dass „löst auf keine Schicht auf" ohne die
  Auflösungs-Konfiguration des Konsumenten nicht von „zeigt legitim nach außen" zu trennen ist.
  Dann wäre die Diagnose selbst fail-open und braucht erst einen Entscheid.

## 1. Auslöser

**Mechanismus: alle Dateien in Schichten, alle Symbole extrahiert — und trotzdem prüft niemand
etwas.** Das ist die gefährlichste Konfiguration, die dieses Werkzeug erlaubt: **vollständig grün,
vollständig blind.**

Gemessen im Konsumenten-Einsatz (2026-08-09): ein Mono-Scan über sechs Sprach-Skelette mit
sprach-präfixierten Schicht-Globs (`go/internal/ui/**`) meldete **0 Befunde**, obwohl ein
`service → ui`-Verstoß eingebaut war. Go-Importe tragen den Modulpfad; das Glob-Literalpräfix kommt
darin nicht als Segment-Run vor, also gilt jedes Ziel als **extern** — der fail-open-Pfad aus
[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) greift, wie
vorgesehen, und schweigt.

**Die bestehende Diagnose greift hier nicht.** Der Hinweis aus
[slice-043](../done/slice-043-schicht-abdeckung-sichtbar.md) meldet Dateien **ohne Schicht**; hier
lagen alle Dateien korrekt in Schichten. Die Lücke sitzt eine Ebene tiefer: bei den **Zielen** der
Importe, nicht bei den Quellen.

**Der Konsument erkannte es nur, weil er einen Verstoß einbaute.** Genau das soll das Werkzeug
abnehmen.

**Abgrenzung zu [slice-081](../done/slice-081-heuristik-diagnose.md):** dort wird eine Zeile gar
nicht **extrahiert**; hier wird sie extrahiert und löst auf **keine Schicht auf**. Zwei
verschiedene Blindstellen, zwei Diagnosen — eine gemeinsame würde beide Ursachen verwischen.

## 2. Betroffene Module

Zwei Schichten:

1. **Spec** — [`spec/lastenheft.md`](../../../../spec/lastenheft.md) und
   [`spec/spezifikation.md`](../../../../spec/spezifikation.md); ADR nach dem Muster von
   [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) (advisory, kein Gate).
2. **`internal/hexagon/core/`** + [`internal/cli/`](../../../../internal/cli/cli.go) — die Zählung
   je Schicht und ihre Ausgabe, neben dem bestehenden Abdeckungs-Hinweis.

## 3. Auszuführende Gates

`make gates`, `make image-test`.

**Anforderung aus dem CR, wörtlich zu erfüllen:**

> „Schicht `<name>`: N Datei(en), 0 von M Import-Symbolen lösen auf eine Schicht auf" — wenn in
> einer Schicht **kein einziges** extrahiertes Symbol auf irgendeine Schicht auflöst, **obwohl
> Symbole extrahiert wurden**.

**Negativ-Proben — die zweite und dritte sind die eigentliche Arbeit:**

| Probe | Erwartung |
|---|---|
| Schicht mit Symbolen, von denen **keines** auflöst | Hinweis mit Name, Datei- und Symbolzahl |
| Schicht, in der **mindestens eines** auflöst | **kein** Hinweis |
| Schicht **ohne** extrahierte Symbole (reines Typen-Paket) | **kein** Hinweis — das ist legitim |
| Exit-Code | unverändert |

Die dritte Zeile trennt diese Diagnose von einer Zählung, die nur „0" sieht: ein Paket ohne
Importe ist kein Befund, sondern normal.

## 4. Was bewusst nicht getan wird

- **Gatend machen.** Wie
  [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md) für die Abdeckungs-Diagnose
  entschieden hat: advisory. Ein Repo, das legitim nach außen zeigt, darf nicht rot werden.
- **Die Auflösung reparieren.** Dass ein Glob-Literalpräfix nicht im Modulpfad vorkommt, ist eine
  Konfigurationsfrage des Konsumenten (`resolution`), keine Werkzeug-Grenze. Die Diagnose sagt ihm,
  **dass** er sie hat.
- **`CR-1`** — nicht extrahierte Schreibweisen sind
  [slice-081](../done/slice-081-heuristik-diagnose.md).

## 5. DoD

- [x] Spec-first: die Diagnose steht als Vertrag, geschärft durch eine ADR, bevor der Code
      committet wird. Beleg: [ADR-0032](../../adr/0032-aufloesungs-diagnose-repoweit.md) mit
      `Status: Accepted`, [`spec/spezifikation.md`](../../../../spec/spezifikation.md) 0.29.0
      ([SPEC-CLI-001](../../../../spec/spezifikation.md#spec-cli-001--aufruf-scan-wurzel-und-exit-codes)).
      **Kein Lastenheft-Bump** — Präzedenz [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md):
      eine advisory stderr-Zeile ohne Exit-Code-Wechsel schärft das *Wie* der bestehenden Ausgabe.
- [x] Der Hinweis erscheint **genau dann, wenn im ganzen Scan** nichts auflöst — nicht je Schicht.
      Die Bedingung wurde beim Bau geändert, auf Messung; Belege in der Closure-Notiz.
- [x] `make gates` und `make image-test` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft.

## 6. Closure-Notiz

**Der Rückführungs-Trigger aus §0 feuerte — und wurde nicht durch Rückführung beantwortet, sondern
durch einen Entscheid.** Er lautete: *„falls sich zeigt, dass ‚löst auf keine Schicht auf' ohne die
Auflösungs-Konfiguration des Konsumenten nicht von ‚zeigt legitim nach außen' zu trennen ist."*
Genau das zeigte sich. Die Trennung gelingt nicht **je Schicht**, wohl aber **repo-weit**; damit
war der Entscheid fällig, nicht der Weg zurück nach `open/`.

**Die im CR wörtlich verlangte Regel wurde zuerst gebaut und gemessen.** Das war der wichtigste
Schritt dieses Slice. Sie feuerte auf dem eigenen Baum:

```text
gesamt: 0 Befund(e)
Hinweis: Schicht core: 3 Datei(en), 0 von 8 Import-Symbolen lösen auf eine Schicht auf
```

Getroffen hat es `internal/hexagon/core` — die Schicht, die laut
[`ARC-001`](../../../../spec/architecture.md) **abhängigkeitsfrei** sein *muss* und darum nur die
Standardbibliothek importiert. Ein reiner Domänenkern ist per Konstruktion das erste Ziel dieser
Regel: dass dort nichts auflöst, ist der Architektur-Erfolg, nicht der Verdachtsfall. Hätte ich die
Regel aus dem CR übernommen und danach getestet, wäre sie mit einem grünen Test in die
Spezifikation gewandert und hätte jeden sauber geschichteten Konsumenten angebellt.

**Der Entscheid ([ADR-0032](../../adr/0032-aufloesungs-diagnose-repoweit.md)):** Auslösung
repo-weit — im ganzen Scan entsteht keine Kante —, Ausgabe unverändert im CR-Wortlaut je Schicht.
Der Konsument bekommt exakt die Meldung, die er angefragt hat; anders ist nur, **wann** sie
erscheint.

**Beobachtbare Architektur-Aussage: die dritte Diagnose sitzt auf derselben Achse wie die zwei
davor, und die Achse ist jetzt sichtbar.** [ADR-0029](../../adr/0029-abdeckungs-diagnose-advisory.md)
meldet die **Datei** ohne Schicht, [ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md) die
**Zeile** ohne Kante, [ADR-0032](../../adr/0032-aufloesungs-diagnose-repoweit.md) das **Ziel** ohne
Schicht — grob nach fein, in fester stderr-Reihenfolge. Alle drei sind advisory, alle drei lassen
den Exit-Code unberührt, und alle drei halten an derselben Grenze: was von repo-externem Code nicht
unterscheidbar ist, wird nicht behauptet. Wer eine vierte Diagnose baut, hat damit ein Muster statt
einer Einzelfallentscheidung.

**Die vier Proben aus §3, gegen `a-check:dev`:**

| Probe | Ergebnis |
|---|---|
| Konsumenten-Muster (`go/internal/…`-Globs, Go-Modulpfade) | beide Schichten gemeldet, Exit **0** |
| eine Schicht löst auf, die andere nicht | **still** — die Kern-Entscheidung |
| Schichten ganz ohne Import-Symbole | **still** |
| eigener Baum (`make arch-check`) | **still**, `gesamt: 0 Befund(e)` |

**Lerneintrag — Form: geschärfte Regel.** Als Prüfsatz: *Eine Anforderung, die eine Regel wörtlich
vorgibt, wird gegen den eigenen Baum gemessen, bevor sie in die Spezifikation geht — nicht danach.*
Der CR war präzise, gut begründet und aus echtem Schmerz geboren; er hatte trotzdem eine
Falsch-Positiv-Klasse, die sein Autor nicht sehen konnte, weil sein Repo keinen abhängigkeitsfreien
Kern hat. Das Dogfooding ist hier nicht Qualitätssicherung *nach* dem Bau gewesen, sondern das
Instrument, das den Entwurf korrigiert hat. Die Kosten waren gering — eine Implementierung, ein
`make arch-check` —, der vermiedene Schaden groß: eine `Accepted`-ADR ist immutabel.

**Ein Befund reproduzierte sich sofort:** [ADR-0031](../../adr/0031-heuristik-grenzen-diagnose.md)
fehlte im ADR-Index — derselbe Fund, den
[slice-081](../done/slice-081-heuristik-diagnose.md) gestern für
[ADR-0030](../../adr/0030-kein-digest-im-generierten-fragment.md) diagnostiziert und als Folge-Slice
benannt hat. Beim nächsten ADR ist er mir wieder passiert, obwohl ich ihn kannte. Ein Fehler, den
Wissen allein nicht verhindert, braucht einen Sensor; die Priorität des Folge-Slice steigt damit
von „irgendwann" auf „vor der nächsten ADR".

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
