# `make doc-workflows` — Deklarations-Form der `uses:`-Referenzen

## Vertrag

Jede `uses:`-Referenz unter `.github/workflows` hält ihre Form: eine **fremde**
nennt einen vollen 40-stelligen SHA mit Tag-Kommentar dahinter; eine **lokale**
(`./…`) braucht keinen Pin, dafür ein existierendes Ziel und einen
Aufrufer-Job, der die verlangten Rechte führt.

## Grenze — was das Grün nicht abdeckt

1. **Die Gültigkeit des Tag-Kommentars** — ob der Tag zum SHA gehört, sagt der
   Lauf nicht. Das wäre eine Registry-Anfrage, und `gates` ist hermetisch.
   Permanent. Der Widerspruch bleibt damit ungedeckt und ist als
   [`BEO-GATE/versionsangabe-neben-digest-ungeprueft`](../../docs/plan/planning/observations/BEO-GATE/versionsangabe-neben-digest-ungeprueft/observation.md)
   registriert; die **Kohärenz** zweier Angaben trägt `make version-coherence`,
   die **Wahrheit** keiner von beiden.

Geprüft wird die **Form**, nicht die Gültigkeit — der Unterschied ist der ganze
Inhalt dieser Grenze.

## Sperren

- eigenes `doc-*`-Target, weil `d-check.mk` für dieses Modul keines erzeugt;
  fehlt es im Makefile, läuft die Prüfung gar nicht.

## Bindung

`DC-FA-WF-001` · im `gates`-Aggregat seit slice-130 (fand dabei den latenten
Release-Bruch in `release.yml`).
