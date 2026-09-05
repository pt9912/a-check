# Slice-Provenienz aus Gedächtnis statt aus `git log` behauptet

**Sub-Area:** Planungs-Harness

Eine Provenienz-Aussage über eine andere Datei ("Absatz X besteht seit
slice-NNN") wurde aus dem Gesprächskontext/Gedächtnis übernommen statt mit
`git log -p --follow -S"<Textfragment>" -- <Datei>` verifiziert — dieselbe
Fehlerklasse wie beim `cr-text-behauptet-statt-gemessen`-Muster
([`BEO-GATE/cr-text-behauptet-statt-gemessen`](../../BEO-GATE/cr-text-behauptet-statt-gemessen/observation.md)),
hier aber nicht bei einer CR-Behauptung an ein fremdes Werkzeug, sondern bei
einer repo-eigenen historischen Aussage in einem Slice-Plan.
