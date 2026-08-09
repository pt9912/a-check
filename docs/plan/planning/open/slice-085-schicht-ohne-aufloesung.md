# slice-085 — Diagnose: Schicht ohne auflösende Importe

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `CR-2` des Konsumenten-Einsatzes (2026-08-09);
[AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze),
[SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung).
**Bezug:** Welle [`welle-13-konsumenten-befunde`](../welle-13-konsumenten-befunde.md); Vorbild ist
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

**Abgrenzung zu [slice-081](../in-progress/slice-081-heuristik-diagnose.md):** dort wird eine Zeile gar
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
  [slice-081](../in-progress/slice-081-heuristik-diagnose.md).

## 5. DoD

- [ ] Spec-first: die Diagnose steht als Vertrag im Lastenheft, geschärft durch eine ADR, bevor
      Code entsteht. Beleg: Lastenheft-Versionszeile + ADR mit `Status: Accepted`.
- [ ] Der Hinweis erscheint genau dann, wenn eine Schicht Symbole trägt und keines auflöst. Beleg:
      die ersten drei Proben aus §3.
- [ ] `make gates` und `make image-test` grün — **Ausgabe in eine Datei**, Exit-Code getrennt
      geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
