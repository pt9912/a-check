# `make doc-check` — Doku-Hygiene des Repos über `d-check`

## Vertrag

Links, Anker, Kennungs-Linkpflicht und Referenzmatrix der Repo-Doku lösen auf;
dazu drei Modul-Zusagen, die über die reine Auflösbarkeit hinausgehen.

| Modul | Zusage | seit |
|---|---|---|
| `links` | jeder Verweis löst auf — und `resolve-from` zusätzlich aus **jeder** Lifecycle-Position, in die die Datei wandern kann | slice-080 |
| `versions` | eine Prosa-Angabe über die Version eines *anderen* eigenen Dokuments stimmt mit dessen Kopf überein (`version-stale`); zweites Muster für Baseline-Pins | slice-133 / slice-173 |
| übrige | Anker, Kennungs-Linkpflicht, Referenzmatrix | — |

Läuft digest-gepinnt, `--network none`, read-only.

## Grenze — was das Grün nicht abdeckt

1. **Symlinks** — bei einem Symlink liest `d-check` den *Zielinhalt*, nicht den
   Linkpfad; ein ins Leere zeigender Symlink bleibt hier grün. Geheilt durch
   `make symlink-check` (`harness/sensors/symlink-check.md`).
2. **Die `current-from`-Trägerdatei** — `versions` nimmt die Datei aus, aus der
   es den Erwartungswert liest (`harness/conventions.md`). Ihre Baseline-Pins
   werden nur deshalb geprüft, weil `.claude/rules/conventions.md` ein Symlink
   darauf ist und *dieser* Pfad geprüft wird. Nebeneffekt, keine Zusage;
   permanent, solange der Erwartungswert dort steht.
   **Wie viele es sind, sagt das Kommando, nicht diese Datei:**
   `grep -c '\.harness/baseline/v[0-9]' harness/conventions.md`.
3. **Zeitdokument-Klassen** — fünf Verzeichnis-Globs sind vom `versions`-Muster
   ausgenommen (`exempt-paths` in `.d-check.yml`): dort ist der genannte Stand
   wahr, weil damals gegen ihn gemessen wurde. Zwei *lebende* Zeiger fallen
   dabei mit heraus (`docs/reviews/README.md`, der Vorlagen-Link in [`MR-018`](../conventions.md#mr-018)) —
   heilbar durch eine Regel je Datei-Rolle statt je Verzeichnis.

   **Das ist ein Geltungsbereich, keine Gate-Senkung** — geprüft an der
   Baseline, nicht angenommen. `v6.5.0` ·
   `regelwerk/grundlagen-harness-dateien.md` §Was ein Kommentar trägt zieht die
   Linie selbst: *„Die Grenze: Sie gilt für einfrierende Artefakte … Der
   Unterschied ist nicht die Wichtigkeit des Ziels, sondern ob der Zeiger
   nachgezogen werden **darf**."* Ein Zeiger in einem Zeitdokument darf es
   nicht; eine Prüfung, die ihn trotzdem einfordert, verlangt einen Regelbruch.
   Die Ausnahme bildet also ab, worüber die Regel spricht — sie senkt keine
   Schwelle, und `AGENTS.md` §3.6 (Gates nur per ADR lockern) greift nicht.

   **Die Stelle, an der der Kurs von einer „Gate-Senkung mit eigener
   Begründungslast" spricht, meint etwas anderes:** ein Ventil im Prüfbereich
   der **Link**-Prüfung, wenn eine Adresse bereits im eingefrorenen Artefakt
   steht. Diesen Fall hat a-check nicht mit einem Ventil gelöst, sondern durch
   **Vermeidung** — die Zitier-Form (`AGENTS.md` §5, slice-176) —, und genau das
   empfiehlt derselbe Absatz: *„die Reparatur ist teurer als die Vermeidung"*.
4. **Digests** — der Erwartungswert kommt versions-förmig aus dem
   `current-from`-Span; ein `sha256:` bricht dort fail-closed ab. Die
   Digest-Gleichheit trägt `make gate-consistency`. Permanent.

**Wie groß der Ausschnitt ist, sagt das Kommando, nicht diese Datei:** die
Schluss-Zeile des Laufs (`d-check: <N> Datei(en) geprüft, <M> Befund(e)`). Sie
sagt etwas über die von `.d-check.yml` erfassten Dateien, nicht über das Repo.

## Sperren

- `Range-Basis nicht auflösbar` — nur bei den Range-gebundenen Modulen, nicht
  hier; `doc-check` läuft ohne Commit-Range.
- fehlendes Image → `docker`-Fehler, kein stiller Durchlauf: der Pin ist ein
  Digest, kein Tag.

## Bindung

Harness-Prozess (Doku-Hygiene; Dogfooding des Stacks) · Bootstrap-Gate ·
Lifecycle-Invariante seit slice-080 (löst `verify-slice-links` ab; die Beobachtung dahinter heißt seit slice-139 [`BEO-PLAN/verweis-auf-wandernden-slice`](../../docs/plan/planning/observations/BEO-PLAN/verweis-auf-wandernden-slice/observation.md), die alte Kennung `SL-002` führt das Register nicht mehr).
