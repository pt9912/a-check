**Vorgang:** slice-188

**Fund:** F-9 des unabhängigen Reviews. Der Kopf von [`.d-check.yml`](../../../../../../../.d-check.yml)
zählte unter *„Aktiv:"* **fünf** Module auf (`links`, `anchors`, `ids`, `matrix`, `spans`),
während die `modules:`-Zeile **derselben Datei** acht führt (`+ hostpaths`, `+ versions`,
`+ reviews`). Die Aufzählung stand zwölf Zeilen über ihrer eigenen Quelle.

**Dritte Ausprägung derselben Klasse.** Die erste ist der behobene Anlass
(`AGENTS.md` §4, `doc-*`-Targets gegen das `Makefile`), die zweite die `gates`-Zelle in
`harness/README.md`; hier ist es eine **Kommentar**-Aufzählung gegen einen YAML-Schlüssel —
dieselbe Klasse, ein anderes Artefakt, eine andere Sprache.

**Behoben — und zwar mit der Antwort, die `state.md` beim dritten Auftreten vorgesehen hatte:**
Die Aufzählung ist **gestrichen**, nicht korrigiert. An ihrer Stelle steht der Zeiger
(*„Was aktiv ist, sagt die `modules`-Zeile unten — eine Aufzählung daneben driftet gegen sie"*),
und geblieben ist nur, was **nicht** aktiv ist, mit seinem Grund.

**Wie es auffiel:** im unabhängigen Review, nicht im Lauf — und **kein** Gate fing es: die
Kandidaten aus §3.3 desselben Slice sind dieselbe Menge, aber keine Prüfung hält Prosa gegen
den YAML-Schlüssel.
