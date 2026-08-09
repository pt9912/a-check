# slice-073 — d-check statt Eigenbau: Analyse und CR (zur Abnahme)

**Status:** open — **Analyse zur Abnahme** (Messung + CR-Formulierung; **keine** Ablösung im
Code — die ist Folge-Slice). Der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** Maintainer-Frage vom 2026-08-09 — *„Warum können wir d-check nicht anstatt der
`verify-*.sh` verwenden?"*; Bezug zu
[MR-006](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert)
und [ADR-0021](../../adr/0021-commits-modul-trace-check.md).
**Bezug:** Roadmap-Zeile *Aktuelle Welle* in der [Roadmap](../in-progress/roadmap.md).

---

## 1. Auslöser

**897 Zeilen Shell** prüfen in `tools/` Invarianten über Markdown — während das Repo mit d-check
ein gepinntes, hermetisches Doku-Gate mit **19 Regelmodulen** einbindet.

Das ist a-checks eigene Doktrin, angewandt auf a-check: der Maßstab dieses Projekts ist
**Skript-Reduktion über die Konsumenten-Flotte**. Ein Shell-Skript im Konsumenten-Repo, das eine
**generische** Invariante prüft, ist ein CR-Kandidat für das Werkzeug — nicht dauerhaft
Konsumenten-Sache. Hier ist a-check der Konsument, und `tools/` ist sein P-Rest.

Die Frage wurde nie gestellt. [slice-050](../done/slice-050-verify-schicht.md) hat die
Doppelung nur für die *Referenzmatrix* geprüft („ein zweiter Prüfer wäre Doppelung") und die
Struktur-Sensoren nicht betrachtet.

**Erste Messung (Handbuch `d-check` §6 Regelmodule, Stand v0.51.1):**

| Eigenbau | Zeilen | d-check-Pendant | Befund |
|---|---|---|---|
| `gate-consistency` Checks (1)+(2) | ~45 von 308 | Modul **`targets`** (`gate-phantom`/`gate-undocumented`) | **vollständige Doppelung** |
| `verify-slice-links` | 146 | `links` prüft nur den Ist-Ort | Verallgemeinerung fehlt |
| `verify-closure-notes` | 146 | — | kein Struktur-Modul |
| `verify-slice-form` | 166 | — | kein Struktur-Modul |
| `verify-ac-form` | 131 | — | kein Struktur-Modul |

**Eine dritte Achse kam am selben Tag dazu:** das CLI-Werkzeug `mq` (Markdown → Node-Baum,
jq-artige Abfragen; als Skill `.claude/skills/markdown-mq/` hinterlegt). Alle fünf Eigenbau-Prüfer
bauen Markdown-Struktur mit **zeilenbasierten** Mitteln nach — `sed` schneidet Code-Fences heraus,
`grep -oE` extrahiert Links, Checkboxen werden über Abschnitts-Grenzen hinweg gezählt. Genau
daraus entstanden `F-4` und `R-068-F4`.

Gemessen an einer Fixture mit Inline-Link, Referenz-Link und einem Link im Code-Block:

| Weg | Inline | Referenz (`F-4`) | im Code-Block |
|---|---|---|---|
| heutiges `links_of()` (`sed`+`grep`) | gefunden | **übersehen** | korrekt ignoriert (per `sed`-Hack) |
| `mq '.link'` | gefunden | **übersehen** | korrekt ignoriert (strukturell) |
| `mq '.link'` + `.link_ref` + `.definition` | gefunden | **gefunden** | korrekt ignoriert (strukturell) |

`mq` löst `F-4` also **nicht** durch einen Selektor — die Auflösung `link_ref` → `definition` →
Ziel bleibt eigene Logik. Was es liefert, ist die *Struktur*: Code-Blöcke fallen ohne
`sed`-Vorfilter weg, und ein Referenz-Link ist ein eigener Node-Typ statt eines nicht getroffenen
Regex-Falls.

**Der Haken ist die Hermetik.** `mq` ist Node-basiert und ein Host-Werkzeug;
[`AGENTS.md`](../../../../AGENTS.md) §3.1 verlangt Docker/make-only. Als Analyse-Werkzeug für
Menschen und Agenten ist es unbedenklich, **in einem Gate** wäre es eine Host-Abhängigkeit — dort
nur über ein Image. Diese Unterscheidung gehört in die Messung, nicht in eine Fußnote.

Der harte Fund ist die erste Zeile der Tabelle oben. `targets` ist konfiguriert als

```yaml
targets:
  makefiles: [Makefile]     # Regelnamen-Quelle
  doc-tables: [AGENTS.md]   # Richtung 1 => gate-phantom
  authority: AGENTS.md      # Richtung 2 => gate-undocumented
  exempt-targets: []        # Utility ohne Doku-Pflicht
```

— das ist Zeile für Zeile, was `gate-consistency.sh` in `doc_targets()`,
`check_documented_exist()` und `UTILITY_TARGETS` selbst gebaut hat. `make doc-targets` **existiert
und ist in [`AGENTS.md`](../../../../AGENTS.md) §4 dokumentiert**, läuft aber in keinem Aggregat;
`gates` fährt stattdessen den Eigenbau.

**Der CR-Weg ist zweifach eingefahren — und trägt einen Namen.** In d-check existiert **kein**
offener CR für ein Struktur-Modul (`open/` und `next/` sind leer, kein Eintrag in Lastenheft oder
Spezifikation). Aber genau dieses Muster ist dort zweimal gelaufen:

| d-check-Modul | abgelöstes Skript | Anlass laut Lastenheft |
|---|---|---|
| `vcs` (`DC-FA-VCS-001`) | `adr-immutable-check.sh` | *„vollständig mechanisieren (Copy-Drift über die Repo-Familie) … die verteilte git-Garantie für Konsumenten ohne Skript-Kopie (**„verteilen statt kopieren"**)"* |
| `targets` (`DC-FA-TGT-001`) | `gate-consistency.sh` | DoD-Punkt *„Paritäts-Mutations-Beleg vs. `gate-consistency.sh`"* |

Beide Male: ein Shell-Skript aus der Repo-Familie wird zum Modul, mit **Paritätsbeleg gegen das
abgelöste Skript**. Der CR-Text dieses Slice hat damit zwei Präzedenzfälle, an deren Form er sich
zu halten hat — einschließlich der Erwartung, dass ein Paritäts-Mutations-Beleg dazugehört und
nicht nachgereicht wird.

## 2. Betroffene Module

Dieser Slice ändert **keinen** Sensor. Er liefert zwei Dokumente:

1. eine **Abdeckungs-Messung** je Eigenbau-Prüfung gegen die 19 d-check-Module,
2. einen **CR-Text** für d-check zu dem, was generisch ist.

Berührt werden damit nur `docs/plan/planning/` und (für den CR) `docs/`.

## 3. Auszuführende Gates

`make doc-check`, `make verify` — beide unverändert grün, weil nichts am Code geändert wird.

**Die Messung ist die Beleg-Arbeit dieses Slice**, nicht ein Gate-Lauf. Sie ist erst gültig, wenn
je Eigenbau-Prüfung eine der vier Aussagen belegt ist:

| Aussage | Beleg |
|---|---|
| **abgelöst** — ein d-check-Modul prüft dasselbe | Lauf beider Prüfer gegen dieselbe Negativ-Fixture, gleiche Befund-Menge |
| **CR-fähig** — generisch, aber kein Modul deckt es | benannte Abstraktion, die nicht a-check-spezifisch ist |
| **bleibt lokal, strukturell** — repo-spezifisch, aber die Markdown-Struktur trägt die Prüfung | der Node-Typ, der die zeilenbasierte Heuristik ersetzt, **plus** der Weg, wie er hermetisch verfügbar wird |
| **bleibt lokal, zeilenbasiert** — weder generisch noch strukturell fassbar | die Repo-Eigenheit, an der es hängt, ist benannt |

Ohne die erste Zeile bliebe der Slice eine Vermutung — genau der Fehler, an dem der
Vorgänger-Entwurf `4b029e4` gescheitert ist.

## 4. Was bewusst nicht getan wird

- **Die Ablösung selbst.** `doc-targets` in `gates` zu hängen und die 45 Zeilen zu entfernen ist
  ein **Folge-Slice** mit eigener Negativ-Probe. Analyse und Eingriff im selben Schnitt wäre
  genau das Dehnen, das `R-068-F5` am Vorgänger-Entwurf bemängelt hat.
- **Den CR an d-check stellen.** Dieser Slice **formuliert** ihn; ihn einzureichen ist ein Akt
  gegenüber einem Fremdrepo und gehört dem Maintainer.
- **`suppression-check` und `regelwerk-check`.** Sie prüfen Go-Quellen bzw. Datei-Hashes — nicht
  d-checks Domäne. Ohne Messung sind sie hier nicht einmal Kandidat.
- **Ein Urteil über die verbleibenden `gate-consistency`-Checks.** Pin-Konsistenz (4),
  `.d-check.yml`-Modulliste (3) und `.PHONY`-Vollständigkeit (5, aus
  [slice-068](../done/slice-068-phony-vollstaendig.md)) werden **mitgemessen**, aber dieser Slice
  behauptet über sie nichts vorab.

## 5. DoD

- [ ] Abdeckungs-Messung liegt vor: je Eigenbau-Prüfung eine der vier Aussagen aus §3, **mit dem
      dort geforderten Beleg**. Beleg: das Mess-Dokument, mit Fixture-Läufen für jede
      „abgelöst"-Zeile.
- [ ] CR-Text für d-check liegt vor, mit einer Abstraktion, die **nicht** a-check-spezifisch ist —
      formuliert als Modul-Vertrag (was es prüft, welche Grund-Codes, welche Konfiguration), nicht
      als Wunschliste, **und mit dem Paritäts-Mutations-Beleg als benanntem DoD-Punkt**, wie ihn
      `vcs` und `targets` beide tragen. Beleg: der Text, gegen die Form der 19 bestehenden Module
      und gegen die zwei Präzedenz-CRs aus §1 gehalten.
- [ ] `make doc-check` und `make verify` grün — **Ausgabe in eine Datei**, Exit-Code getrennt
      geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
