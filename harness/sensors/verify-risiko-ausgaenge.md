# `make verify-risiko-ausgaenge` — jedes notierte Risiko trägt einen Ausgang

## Vertrag

Jedes in §6 eines Slice-Plans **notierte** Risiko trägt genau einen Ausgang aus
der geschlossenen Dreier-Menge (`AGENTS.md` §5). Geprüft in `done/` **und** in
`in-progress/`, sobald dort die Closure-Notiz **ausgefüllt** ist — der Auslöser
ist ihr Zustand, nicht das Verzeichnis.

## Grenze — was das Grün nicht abdeckt

1. **Die Existenz eines Risiko-Blocks** — ein Slice ganz ohne §6 fällt nicht
   auf. Geprüft wird, dass *notierte* Risiken einen Ausgang tragen, nicht dass
   welche notiert sind. Permanent; die Existenz zu erzwingen hieße, einen
   leeren Block als Erfüllung zu akzeptieren.
2. **Ob der Ausgang trägt** — ob das Risiko wirklich nicht mehr eintreten kann,
   ob die genannte Folge-Slice-ID es auffängt, ist ein Urteil. Der Lauf prüft
   die **Form**: dass ein Ausgang dasteht und welcher der drei es ist.
   Permanent (Baseline `modul-05`).
3. **Die Zeichenfolge statt der Aussage** — das Muster sucht `Folge-Slice`,
   `Carveout`, `gestrichen mit Begründung` oder `Beobachtungs-Register` hinter
   `Ausgang:`. Ein Ausgang, der eines dieser Wörter in der **Verneinung** führt
   („kein Folge-Slice"), passiert. Gefunden im Review zu slice-176; heilbar nur
   durch ein Muster, das die Verneinung erkennt — also durch Sprachverständnis.

## Sperren

- Slices mit dem Vorlagen-Platzhalter „beim Abschluss ausfüllen" in der
  Closure-Notiz gelten als **in Arbeit** und werden **sichtbar** übersprungen;
  die Schluss-Zeile nennt ihre Zahl.

## Bindung

Harness-Prozess (`AGENTS.md` §5; Baseline `modul-05` §Offene Risiken werden bei
Closure aufgelöst) · slice-102, `in-progress/`-Erweiterung slice-129 · im
`verify`-Aggregat. Bleibt lokal, weil die Prüfung §6 mit §7 vergleicht und
`structure` abschnitts-**lokal** ist.
