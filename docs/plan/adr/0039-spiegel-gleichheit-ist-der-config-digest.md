# ADR-0039 — Die Spiegel-Gleichheit ist der Config-Digest; die Pins bleiben GHCR-gebunden

- **Status:** Accepted
- **Datum:** 2026-08-30
- **Autor:** pt9912
- **Bezug:** [AC-FA-DIST-002](../../../spec/lastenheft.md#ac-fa-dist-002) (neu; Lastenheft 0.25.0→0.26.0), [AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) (Reproduzierbarkeit über Digest-Pins), [ADR-0030](0030-kein-digest-im-generierten-fragment.md) (warum das Fragment keinen Digest trägt)
- **Schärft:** [SPEC-DIST-001](../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution) — macht Spiegel, Gleichheits-Größe und Pin-Bindung verbindlich.
- **Supersedes:** —

## Kontext

Das Image soll zusätzlich auf Docker Hub liegen. Die Zusage, die zählt, ist
**dasselbe Bild** — nicht ein zweiter Bau aus derselben Quelle, der zufällig
ähnlich aussieht.

**Die naheliegende Gleichheits-Größe trägt nicht.** Der **Manifest**-Digest ist
**registry-lokal**: er hängt an der Blob-Kompression der jeweiligen Registry.
Ein von GHCR kopierter Digest löst auf Docker Hub **nicht** auf. Wer ihn in die
Doku schreibt, liefert einen Pin, der beim Ziehen fehlschlägt.

**Das trifft a-check härter als ein Repo ohne Digest-Regime.** Hier steht der
Pin an **vier** Stellen — [`a-check.mk`](../../../a-check.mk), beide READMEs und
[`version.md#aktuell`](../../../version.md#aktuell) —, und `make gate-consistency`
hält sie auf **Gleichheit**. Eine zweite Registry bringt einen zweiten
Manifest-Digest und damit die Frage, welcher dort steht.

## Entscheidung

**1. Die Gleichheits-Größe ist der Config-Digest.** Er beschreibt den Inhalt
(Layer, Konfiguration) und ist bei identischem Bild auf beiden Registries
gleich. Die Pipeline liest ihn nach dem Push von **beiden** Refs und vergleicht;
Abweichung heißt: es sind nicht dieselben Bilder.

**2. Fail-closed.** Schlägt der Spiegel-Push fehl oder weicht der Config-Digest
ab, ist das **Release fehlgeschlagen**. Der Spiegel ist eine Zusage, kein
Zusatz. Die Fehlermeldung nennt den bereits veröffentlichten GHCR-Digest — der
Teilstand ist gültig, und wer aufräumt, muss wissen, was schon draußen ist.

**3. Der Spiegel steht NACH dem GHCR-Push.** So liegt die Quelle nie hinter dem
Spiegel zurück.

**4. Die Pin-Stellen dieses Repos bleiben GHCR-gebunden.** `a-check.mk`, beide
READMEs und `version.md#aktuell` nennen weiterhin **einen** Digest, und
`gate-consistency` prüft ihn unverändert. Wer vom Spiegel zieht, nimmt den
Digest **der Registry, aus der er zieht**; die Hub-Seite sagt das ausdrücklich.

**Verworfene Alternative — `version.md` führt beide Digests**, das Meta-Gate
prüft je Registry. Ehrlicher für den Hub-Nutzer, aber jede Release-Prep pflegte
zwei Werte, und das Gate bräuchte **zwei Wahrheiten** statt einer. Der Gewinn
wäre ein fertiger Pin für einen Bezugsweg, den dieses Repo nicht als primären
führt; der Preis wäre eine dauerhaft verdoppelte Pflegestelle — genau die Art
Drift-Quelle, die `gate-consistency` sonst verhindert.

**Verworfene Alternative — Hub nur per Tag.** Bricht mit
[AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) für diesen
Bezugsweg. Ein beweglicher Tag ist genau das, was der Digest verhindern soll.

## Konsequenzen

- **Der Hub-Nutzer bekommt aus der Doku dieses Repos keinen fertigen Pin.** Das
  ist der Preis von Punkt 4, und er steht auf der Hub-Seite, wo dieser Nutzer
  ihn liest — nicht nur hier.
- **Die Darstellung ist nicht fail-closed**, der Spiegel schon. Der Unterschied
  ist der Gegenstand: das Bild ist die Zusage, der Beschreibungstext ist
  Präsentation. Ein Release rot zu machen, weil ein Kurztext nicht gesetzt
  werden konnte, verwechselte beides. Der Fehlschlag wird trotzdem **gemeldet** —
  `continue-on-error` setzt `conclusion: success` und lässt `outcome` auf
  `failure`; ohne einen Folgeschritt, der `outcome` liest, wäre der Lauf grün
  und die Seite leer.
- **Der Metadaten-Upload braucht einen anderen Token-Scope als der Push**
  (`read/write/delete`). Ein Token, das pushen darf, darf damit nicht
  zwangsläufig die Beschreibung setzen — und der Fehlschlag sieht aus wie
  Erfolg, wenn niemand `outcome` liest.
- **Der CVE-Sensor bekommt einen zweiten Ref**, sobald der Spiegel ein Bild
  trägt ([ADR-0037](0037-cve-scan-gegen-das-publizierte-image.md)). Vorher
  nicht: ein Ref ohne Bild machte den Nachtlauf ab dem ersten Tag rot.

## Fitness Function

Die Pipeline selbst: der Vergleich der beiden Config-Digests läuft in jedem
Release und ist fail-closed. Dazu `make gates` — `gate-consistency` belegt, dass
die vier Pin-Stellen **einen** Digest tragen und die Entscheidung aus Punkt 4
damit eingehalten ist.
