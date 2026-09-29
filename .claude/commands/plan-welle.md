# Welle planen (Harness)

Argument: $ARGUMENTS

Dieser Command führt die **Planner**-Rolle für *eine* Welle (Modul 6 — Roadmap
Engineering). Der Welle-Zustand ist die **Verzeichnis-Position**: eine offene Welle
liegt **flach** in `docs/plan/planning/`, bei Closure wandert sie per `git mv` nach
`done/` (→ [`close-welle`](close-welle.md)). Der Roadmap-Abschnitt *Offene Wellen* ist
derivativ — ein Zeiger je flacher Welle-Datei; woran gearbeitet wird, sagt das
`Welle:`-Feld der Slices in `in-progress/`.

**Vorab prüfen, ob hier überhaupt eine Welle nötig ist** (`v6.13.0` ·
`regelwerk/modul-06-roadmap.md` §Wann Arbeit eine Welle braucht): Eine Welle braucht
eine beobachtbare Closure-Bedingung, die **mehr** belegt als die DoDs ihrer Slices —
repo-weite Belege (`make ci`/`make verify` über die Einzelnachweise hinaus). Fehlt
dieses Mehr, läuft die Arbeit **wellenlos** als einzelner Slice — in diesem Repo der
Regelfall. Ein Datum triggert nie.

## Kontext lesen

1. [`harness/README.md`](../../harness/README.md) und [`AGENTS.md`](../../AGENTS.md)
   lesen.
2. Den Regelwerk-Abschnitt **on-demand** lesen — hier
   `v6.13.0` · `regelwerk/modul-06-roadmap.md` (dazu Modul 5 für den
   Slice-Lifecycle), nie das ganze Bundle.
3. Die Roadmap
   ([`roadmap.md`](../../docs/plan/planning/in-progress/roadmap.md)) lesen: steht
   die Welle schon unter *Nächste Wellen*? Liegen ihre Slices schon in `open/`?

## Repo-lokale Adaptionen (a-check)

- **Neue Artefakte per `cp` aus den vendored Templates**
  (`.harness/baseline/v6.13.0/templates/…`), dann in-place füllen — keine
  handgeschriebenen Kopien und kein Modellieren auf ein bestehendes Artefakt.
- **Doc-Gate (d-check):** jede `AC-*`-/`ADR-*`-/`MR-*`-Kennung in gescanntem `.md`
  ist ein Anker-Link — ein bloßes Kennungs-Token bricht `make doc-check`
  (`id-unlinked`).
- **Gate-Nachweis + Stop-Hook:** `make gates` endet mit `record-gates`; jede
  Inhaltsänderung nach einem Lauf (inkl. Commit) macht den Stempel ungültig →
  `make gates` erneut laufen.
- **Commit via Message-Datei** (`git commit -F <datei>`), Message mit Kennung
  (`slice-NNN`, `MR-NNN`, `AC-*`); Scope `(planning)` berührt nur
  `docs/plan/planning/`.
- **Lifecycle über `make slice-mv SLICE=… TO=…`** — reiner Move als eigener Commit,
  Verweis-Nachzug durch das Werkzeug; ein `Status:`-Feld gibt es nicht, das
  Verzeichnis ist der Zustand.

## Eröffnung — drei Schritte (Modul 6)

1. **Ziel, Out-of-Scope und Closure-Trigger festlegen** — der Trigger ist eine
   beobachtbare Bedingung, kein Datum (*„ein anderer Mensch kann ohne Rückfrage
   sagen, ob er eingetreten ist"*), und als Start-Trigger kein Ergebnis dieser
   Welle. Erst danach Slices zuordnen.
2. **Beobachtungs-Register sichten**
   ([`observations/`](../../docs/plan/planning/observations/README.md)):
   betrifft eine offene
   Beobachtung eine Sub-Area dieser Welle, gehört sie in die Slice-Planung als
   Risiko; erreicht sie mit dieser Welle 3× (die Zahl der Dateien unter ihrem
   `evidence/`), ist sie eine Lücke und braucht einen eigenen Slice. Keine Treffer
   sind ebenfalls eine Antwort und werden notiert.
3. **Welle-Plan per `cp` anlegen** (`docs/plan/planning/<welle-id>.md`, Ziel-Form
   `welle.template.md`), füllen, Roadmap verdrahten (Zeiger unter *Offene Wellen*),
   `make gates` grün, Commit via `-F`.

**Merke (Modul 6):** Eine Welle endet durch Closure-Kriterien, nicht durch ein Datum
(Welle ≠ Sprint). Wellenlose Arbeit erscheint in der Roadmap nicht — ihr Zustand ist
die Verzeichnis-Position.
