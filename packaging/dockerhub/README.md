# Docker-Hub-Repository-Metadaten

Quelle für die Darstellung von [`pt9912/a-check`](https://hub.docker.com/r/pt9912/a-check)
auf Docker Hub. Docker Hub hat **drei getrennte Felder**; sie werden hier
unterschiedlich gepflegt:

| Feld auf Docker Hub | Quelle | Limit | Gepflegt durch |
|---|---|---|---|
| **Description** (Kurztext unter dem Repo-Namen) | [`description.txt`](description.txt) | 100 Zeichen | Release-Build bzw. `workflow_dispatch` |
| **Repository overview** (Markdown-Seite) | [`overview.md`](overview.md) | 25.000 Zeichen | derselbe Schritt |
| **Category** | dieses Dokument (siehe unten) | — | **manuell im Web-UI** |

Die ersten beiden setzt `peter-evans/dockerhub-description` bei jedem
Release-Build. Die Action hat **keinen** Input für die Kategorie — deshalb steht
die Entscheidung hier als Text, statt still im UI zu leben.

**Die beiden hochgeladenen Dateien sind ENGLISCH**, dieses Dokument nicht. Die
Hub-Seite ist die Außensicht und folgt darin
[`README.md`](../../README.md) (englisch, mit
[`README.de.md`](../../README.de.md) daneben); `description.txt` und
`overview.md` sind das einzige Repo-Material, das ungelesen auf einer fremden
Plattform steht. Was **hier** steht, ist Betriebswissen für dieses Repo und
bleibt deutsch wie der Rest der Doku.

## Transport: die Darstellung funktioniert — gemessen

**Stand 2026-08-30:** beide Felder sind gesetzt, der Inhalt ist **englisch**.
Beleg ist nicht der Ausgang des Laufs, sondern die **Hub-API**: `description`
trägt den Kurztext aus [`description.txt`](description.txt), `full_description`
die Seite aus [`overview.md`](overview.md) mit ersetztem `__VERSION__`.
Der `workflow_dispatch`-Lauf las die Version korrekt aus
[`version.md`](../../version.md) (`v0.18.0`).

**Damit ist der Token-Scope belegt** — der Punkt, an dem der erste Lauf im
Schwester-Repo scheiterte. **Der Spiegel selbst ist seit demselben Tag belegt:** das
Repository trägt `v0.18.0` und `latest` (von Hand gespiegelt, weil der Tag auf einem Commit
sitzt, der den Spiegel-Schritt noch nicht kannte). **Gemessen:** der **Config**-Digest ist auf
beiden Registries identisch (`sha256:e4f357f0…`), der **Manifest**-Digest nicht
(`356aeaea…` gegen `5bdd40ca…`) — genau die Lage, die
[ADR-0039](../../docs/plan/adr/0039-spiegel-gleichheit-ist-der-config-digest.md) zur
Entscheidung macht.

## Was beim ersten Lauf schiefgehen kann — und woran man es erkennt

Die drei Punkte stammen aus dem Schwester-Repo, wo jeder einen Vorfall gekostet
hat. Sie stehen hier, damit sie nicht zweimal bezahlt werden:

1. **Der Metadaten-`PATCH` braucht einen anderen Token-Scope als der Push.**
   Ein Token mit `read/write` darf pushen, aber die Beschreibung **nicht**
   setzen — die Action verlangt `read/write/delete`. Symptom: das Bild liegt auf
   Docker Hub, die Beschreibung bleibt leer, das Log zeigt `Forbidden`. Der
   Scope lässt sich am **bestehenden** Token ändern; der Wert bleibt derselbe,
   das Secret muss nicht ersetzt werden.
2. **`continue-on-error` setzt `conclusion: success`**, lässt aber `outcome` auf
   `failure`. Ohne einen Folgeschritt, der `outcome` liest, ist der Lauf grün
   und die Seite leer. Der Melder in
   [`hub-description.yml`](../../.github/workflows/hub-description.yml) tut genau
   das — als **Warnung**, nicht als Fehler: das Release soll an der Darstellung
   nicht scheitern, nur nicht schweigen.
3. **`workflow_dispatch` checkt den ORIGIN-Stand aus**, nicht den lokalen. Wer
   die Darstellung ändert, **pusht zuerst** und dispatcht danach; die Gegenprobe
   ist die Hub-API, nicht der Ausgang des Laufs.

## Category

**Gesetzt: „Developer tools"** (`developer-tools`) — sichtbar auf der
Repo-Seite als Marke neben `IMAGE`. Bleibt **manuell**: die Action hat keinen
Input dafür, und kein Lauf setzt sie zurück.

Begründung: a-check prüft Architektur-Regeln in Entwickler-Repositories und
läuft als Gate in CI-Pipelines — die Kategorie beschreibt den Nutzen für den
Suchenden, nicht die Implementierungssprache. „Languages & frameworks" wäre
falsch: Go ist Implementierungsdetail
([ADR-0001](../../docs/plan/adr/0001-go-impl-sprache.md)) und für
niemanden ein Grund, das Image zu ziehen.

Docker Hubs Taxonomie ist eine feste Liste; freie Schlagworte gibt es nicht.

## Warum die Darstellung nicht fail-closed ist

Der **Spiegel** ist fail-closed
([`AC-FA-DIST-002`](../../spec/lastenheft.md#ac-fa-dist-002)): liegt das Bild
nicht auf Docker Hub oder weicht sein Config-Digest ab, ist das Release
fehlgeschlagen. Die **Darstellung** ist es nicht — sie läuft mit
`continue-on-error`. Der Unterschied ist der Gegenstand: das Bild ist die
Zusage, der Beschreibungstext ist Präsentation
([ADR-0039](../../docs/plan/adr/0039-spiegel-gleichheit-ist-der-config-digest.md)).

## `__VERSION__`

[`overview.md`](overview.md) trägt den Platzhalter `__VERSION__` in den
Aufruf-Beispielen; der Release-Build ersetzt ihn durch die Tag-Version, bevor er
die Seite hochlädt. Die Datei im Repo bleibt damit versionsfrei und muss bei
keinem Release angefasst werden — anders als die vier harten Pins, die
[`releasing.md`](../../docs/user/releasing.md) aufzählt.

## Änderungen prüfen

Beide Dateien sind reiner Repo-Inhalt und laufen durch `make doc-check` wie jede
andere Doku. Das Zeichen-Limit der Description prüft der Workflow **fail-fast**,
statt der Action das stille Abschneiden zu überlassen:

```bash
wc -m < packaging/dockerhub/description.txt   # muss <= 100 sein
```

**`-m`, nicht `-c`:** Docker Hub misst **Zeichen**, nicht Bytes. Eine
Byte-Messung wäre bei Nicht-ASCII **strenger** als die Regel und meldete einen
Text rot, den Docker Hub annimmt.
