# slice-126 — Lint-Pin auf `v2.13.2`

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Maintainer-Wort 2026-08-30. Schließt an
[slice-125](../done/slice-125-go-toolchain-1-27.md) an — der neue Lint-Stand ist mit derselben
Toolchain gebaut, auf die dieser Slice gehoben hat.

**Berührte Spec-Stellen:** — *(keine)*

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Das Lint-Gate läuft auf `v2.13.2` — demselben Stand, der mit der eben gehobenen Toolchain gebaut
wurde.

## 2. Definition of Done

- [x] [`Dockerfile`](../../../../Dockerfile) pinnt `v2.13.2`: Version **und** Digest gemeinsam
      gehoben, wie bei jedem Basis-Image dieses Repos.
- [x] `make lint` ist grün — **oder** die neuen Befunde sind behoben. Kommt eine Regel hinzu, die
      den Bestand trifft, ist ihre Behandlung Teil dieses Slice: eine Deaktivierung wäre eine
      Schwellen-Senkung und bräuchte nach [`AGENTS.md`](../../../../AGENTS.md) §3.6 eine ADR.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| [`Dockerfile`](../../../../Dockerfile) | update | Lint-Version + Digest |

**Auszuführende Gates:** `make gates` — tragend `lint` (der Gegenstand) und `suppression-check`
(das Suppression-Verbot gilt unverändert). Zum Abschluss `make verify`.

### Das Risiko ist vorab gemessen

| Prüfung | Ergebnis |
|---|---|
| `v2.13.2` verfügbar | ja, Digest ermittelt |
| womit gebaut | `go1.27.0` — dieselbe Toolchain wie seit [slice-125](../done/slice-125-go-toolchain-1-27.md) |
| Sprung | `v2.12.2` → `v2.13.2`, eine Minor-Stufe |

**Ungemessen bleibt der Befundstand.** Ob der neue Lint-Stand über den Bestand Befunde meldet,
sagt erst der Lauf — eine Minor-Stufe bringt bei diesem Werkzeug regelmäßig neue Prüfungen. Genau
dafür ist §2 zweigeteilt: grün **oder** behoben.

## 4. Trigger

**Start:** eingetreten — Maintainer-Wort, der Stand ist verfügbar.

**Rückführungen:**

- `in-progress` → `next`: falls die neuen Befunde eine Umgestaltung verlangen, die für sich ein
  Slice ist. Dann liefert dieser hier den Pin nicht.

## 5. Closure-Trigger

Pin gehoben, Lint grün, Gates grün.

**Was bewusst nicht getan wird:** eine neue Regel **deaktivieren**, um grün zu werden. Das wäre
eine Schwellen-Senkung und verlangt eine ADR ([`AGENTS.md`](../../../../AGENTS.md) §3.6); eine
Inline-Suppression ist ohnehin verboten (§3.2). Behoben wird der Code, nicht der Prüfer.

## 6. Risiken und offene Punkte

- *Eine Minor-Stufe bringt neue Prüfungen; der Bestand könnte sie brechen* — **Ausgang:**
  entfallen, gestrichen mit Begründung: `make lint` läuft über den unveränderten Bestand grün.
  Keine neue Regel greift, keine Behebung war nötig, und die zweite Hälfte der DoD („oder behoben")
  blieb ungenutzt. Das ist der günstige Ausgang, nicht der erwartete — er war vor dem Lauf nicht
  bestimmbar.

## 7. Closure-Notiz

**Geliefert:** Der Lint-Pin steht auf `v2.13.2`, Version und Digest gemeinsam gehoben. `make lint`
und `make gates` laufen über den unveränderten Bestand grün; `suppression-check` bestätigt, dass
keine Inline-Suppression entstanden ist.

**Lerneintrag — Form: geschärfte Regel.** *Ein Slice darf eine Ungewissheit **tragen**, statt sie
wegzuplanen — die DoD nennt dann beide Ausgänge, und der Plan sagt, warum die Messung fehlt.*
Bei den letzten drei Pin-Hebungen dieses Repos stand jeweils eine Zeile *„das Risiko ist vorab
gemessen, nicht abgeschätzt"* im Plan, und sie trug: Digest, Fragment-Diff, Sensor-Läufe waren
vorher bekannt. Hier ging das **nicht** — ob ein Lint-Stand über einen fremden Bestand meldet,
sagt erst sein Lauf; jede Vorab-Aussage wäre eine Behauptung gewesen, genau die Klasse, die
[`BEO-022`](../../../../docs/plan/planning/observations/BEO-GATE/cr-text-behauptet-statt-gemessen/observation.md) zählt. Die ehrliche Form war deshalb nicht „gemessen", sondern
eine **zweigeteilte DoD** (`grün` **oder** `behoben`) plus der ausdrückliche Satz, was ungemessen
bleibt. *Weil* eine Ungewissheit im Plan sichtbar ist, verschiebt sie sich nicht stillschweigend
in die Ausführung.

**Zwei beobachtbare Closure-Kriterien:**

1. Die Hebung ist **vollständig**: Version **und** Digest. Ein gehobener Tag mit stehendem Digest
   zöge weiterhin das alte Bild — dieselbe Klasse Fehler, die
   [slice-125](../done/slice-125-go-toolchain-1-27.md) beim Basis-Image vermieden hat.
2. Der Ausgang steht in der DoD, die ihn vorgesehen hat: die erste Hälfte trat ein, die zweite
   blieb ungenutzt. Ein Plan, der nur den günstigen Fall kennt, hätte den anderen zur Überraschung
   gemacht.

**Offene Risiken und ihr Ausgang:** das eine entfallen, gestrichen mit Begründung.

**Beobachtungs-Register:** keine neue Beobachtung. Der Lauf hat nichts gezeigt, was über die
Hebung hinausweist — und das ausdrücklich festzuhalten ist die Antwort, nicht ein leeres Feld.

**Folge-Slices:** keiner aus diesem Slice. Offen bleibt die Kette aus
[slice-124](../done/slice-124-image-scan-cve.md) und
[slice-125](../done/slice-125-go-toolchain-1-27.md): das **Release**, das die Toolchain-Hebung
veröffentlicht, und danach der **Hebungs-Kanal** (Dependabot).
## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** berührt wird die **Build-Schicht** (`Dockerfile`), wie in
[slice-125](../done/slice-125-go-toolchain-1-27.md).

**Vorgelagert — offene Beobachtungen sichten:** keine Treffer in der berührten Sub-Area.

Alle berührten Sub-Areas GF.
