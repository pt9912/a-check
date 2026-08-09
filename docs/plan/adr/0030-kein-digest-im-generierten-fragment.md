# ADR-0030 — `--print-mk` gibt keinen Digest aus, sondern einen Platzhalter

- **Status:** Accepted
- **Datum:** 2026-08-09
- **Autor:** pt9912
- **Bezug:** [AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit), [AC-FA-DIST-001](../../../spec/lastenheft.md#ac-fa-dist-001--distribution-image---print-mk-a-checkmk), [SPEC-DIST-001](../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution), [ADR-0004](0004-distribution-image-mk.md), [ADR-0007](0007-latest-tag-politik.md)
- **Schärft:** [SPEC-DIST-001](../../../spec/spezifikation.md#spec-dist-001--laufzeitform-und-distribution) — den Inhalt des erzeugten Makefile-Fragments.
- **Supersedes:** —

## Kontext

Das `--print-mk`-Fragment trug bisher einen konkreten Digest, einkompiliert als Konstante. **Dieser
Digest ist strukturell immer der des Vorgänger-Release.**

Der Grund ist keine Nachlässigkeit, sondern ein Fixpunkt-Problem: Der Digest ist der Hash des
Image-Inhalts. Ein Image, das seinen eigenen Digest enthält, müsste einen SHA-256-Fixpunkt
realisieren. Der Release-Ablauf spiegelt das — der Re-Pin (`releasing.md` Schritt 6) erfolgt
zwangsläufig **nach** dem Push, kann den Digest also nur ins *nächste* Image tragen.

Gemessen am 2026-08-09:

| Tag | einkompiliert | echter Digest |
|---|---|---|
| `v0.16.0` | `6425c93a…` | `aef28cfe…` |
| `v0.15.0` | `f1b8ff5e…` | `6425c93a…` |

Unabhängig bestätigt am lokalen Image-Cache: `ghcr.io/pt9912/a-check:v0.15.0` trägt
`sha256:6425c93a…` — den Wert, den `cli.go` zum Tag `v0.16.0` führte.

**Der Schaden ist real eingetreten.** Ein Konsument folgte dem dokumentierten Bump-Weg („Image
ziehen, `--print-mk` laufen lassen") und pinnte dadurch `v0.15.0` statt `v0.16.0`. Er merkte es
nicht — der Digest war ja gültig. Erst als der `constructs`-Block fehlte, den sein Skelett
brauchte, wurde die Abweichung sichtbar: als vermeintlicher Konfigurationsfehler.

Zwei Doku-Stellen behaupteten dabei das Unmögliche (`releasing.md` §5 und Freigabe-Item 6 „`--print-mk`
gibt den neuen Digest", **mit Beleg-Slot**). Kein Gate fing es: `gate-consistency` prüft nur den
Repo-Stand, `image-test` die Fragment-Parität gegen das **lokal gebaute** Image, wo die Konstante
per Konstruktion stimmt.

## Entscheidung

**Das erzeugte Fragment enthält keinen konkreten Digest, sondern einen Platzhalter mit
Bezugsquelle.** Die committete `a-check.mk` im Repo trägt weiterhin den echten Digest des aktuellen
Release — sie ist das *gepinnte Artefakt*, das erzeugte Fragment nur die *Vorlage* dafür.

Der Platzhalter ist so gewählt, dass eine unveränderte Übernahme **sichtbar abbricht** statt still
ein fremdes Release zu ziehen.

## Alternativen

**Beweglichen Versions-Tag ausgeben** (`ghcr.io/pt9912/a-check:v0.16.0`). Verworfen aus zwei
Gründen. Erstens ist die Version im Binary **nicht bekannt**: `ARG VERSION` steht in der
runtime-Stage des Dockerfile, der Build-Schritt hat kein `-ldflags -X`. Der Weg verlangt also eine
Build-Änderung. Zweitens bricht er die **Fragment-Parität**: ein lokaler Build trägt
`VERSION=0.0.0-dev`, die committete `a-check.mk` müsste dieselbe Zeichenfolge tragen und wäre als
Pin unbrauchbar. Beides ist lösbar, aber der Preis steht in keinem Verhältnis — und ein
beweglicher Tag gibt die Digest-Härte auf, die [ADR-0007](0007-latest-tag-politik.md) für
`:latest` bewusst ablehnt.

**Digest zur Laufzeit auflösen.** Netzlos unmöglich; der Manifest-Digest ist Registry-Metadatum und
liegt außerhalb des Container-Rootfs. Mit Netz wäre es machbar (Version einbacken, Registry
fragen), verletzt aber die Hermetik des Ausgabepfads und macht ihn nichtdeterministisch. Die
Information existiert genau eine Ebene höher: der **Host** kennt den Digest über
`docker inspect --format '{{index .RepoDigests 0}}'` — dasselbe, was `release.yml` bereits nutzt.
Das ist der Weg, den die Doku dem Konsumenten nennen kann, ohne dass das Binary ihn gehen muss.

**Nur die Doku korrigieren.** Der Fragment-Inhalt bliebe falsch, nur ehrlich kommentiert. Das
Akzeptanzkriterium des Konsumenten schließt es aus: *„Die Ausgabe enthält keinen Digest, der auf
ein anderes Release zeigt als das laufende Binary."* Ein korrekt kommentierter falscher Digest
erfüllt das nicht — und der Fehlpin entstand gerade dadurch, dass ein gültig **aussehender** Wert
übernommen wurde.

## Konsequenzen

**Positiv.** Die Ausgabe kann nicht mehr auf ein fremdes Release zeigen. Die unveränderte Übernahme
scheitert sichtbar statt still. Der Platzhalter erzwingt genau den bewussten Commit, den
[AC-QA-03](../../../spec/lastenheft.md#ac-qa-03--reproduzierbarkeit) ohnehin verlangt — die Regel
wird durchgesetzt statt nur behauptet. Und: keine Build-Änderung, die Fragment-Parität bleibt
trivial prüfbar.

**Negativ.** Das erzeugte Fragment ist **nicht sofort lauffähig**. Wer `include a-check.mk` fährt,
muss zuerst den Digest einsetzen. Das ist der akzeptierte Preis: das Fragment war vorher sofort
lauffähig und dabei falsch.

**Für Bestands-Konsumenten ändert sich nichts** — ihre bereits gepinnte `a-check.mk` bleibt gültig.
Betroffen ist nur, wer das Fragment **neu erzeugt**.

## Geschichte

| Datum | Änderung |
|---|---|
| 2026-08-09 | Entwurf und Annahme. Auslöser: Maintainer-Befund plus `CR-5` eines realen Konsumenten-Einsatzes mit belegtem Fehlpin (slice-083). |
