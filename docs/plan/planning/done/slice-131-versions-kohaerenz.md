# slice-131 — Dieselbe Angabe an zwei Orten wird abgeglichen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** [`BEO-026`](../observations.md) — bei **3×**, Schwelle überschritten. „Besser aufpassen"
ist keine Antwort ([`AGENTS.md`](../../../../AGENTS.md) §5).

**Berührte Spec-Stellen:** — *(keine)* — die Build-Konfiguration ist nicht Gegenstand des
Lastenhefts.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Wo dieselbe Identität zweimal deklariert ist, fällt ein Auseinanderlaufen auf — statt jahrelang
richtig auszusehen.

## 2. Definition of Done

- [x] Die beiden Versions-Variablen im [`Makefile`](../../../../Makefile) tragen den **gemessenen**
      Wert des jeweils gepinnten Images (§3), nicht den zuletzt geschriebenen.
- [x] `tools/verify-versions-kohaerent.sh` prüft **zwei** Kohärenz-Regeln (§4) und trägt einen
      Selbsttest, der beide Richtungen jeder Regel abdeckt — auch die, in der der Sensor nichts
      finden **darf**.
- [x] `make version-coherence` ist im `gates`-Aggregat und in
      [`AGENTS.md`](../../../../AGENTS.md) §4 sowie
      [`harness/README.md`](../../../../harness/README.md) deklariert.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Der dritte Vorfall, gemessen

Beide Werte stammen aus dem **gepinnten** Image, nicht aus einer Tabelle:

| Variable | `Makefile` | `Dockerfile`-`ARG` | im gepinnten Image |
|---|---|---|---|
| `GO_VERSION` | `1.26.4` | `1.27.0` | **`go1.27.0`** |
| `GOLANGCI_LINT_VERSION` | `v2.12.2` | `v2.13.2` | **`2.13.2`** |

Das Makefile sticht den `ARG` per `--build-arg`, und der Digest daneben sticht den Tag. Wirksam ist
also der Digest; **beide** Makefile-Zahlen sind falsch, und **keine** von ihnen hatte je eine
Wirkung. Das ist die unangenehme Hälfte: eine folgenlose Zahl wird nicht dadurch harmlos, dass sie
folgenlos ist — sie wird beim nächsten Lesen geglaubt.

## 4. Die Klasse ist enger, als sie aussah — und damit netzlos prüfbar

`BEO-026` hieß zuerst „der Tag-Kommentar wird nicht gegen den Digest geprüft". Das stimmt, und es
ist eine **Netz**-Frage. Aber es beschreibt keinen der drei Vorfälle:

| # | Vorfall | Was auseinanderlief |
|---|---|---|
| 1 | `image-scan.yml` `# v5.0.0` ↔ `ci.yml` `# v6.0.2` | zwei Kommentare, **ein** Digest |
| 2 | `release.yml` `# v4.2.0` ↔ `# v3.6.0` | zwei Kommentare, **ein** Digest |
| 3 | `Makefile` ↔ `Dockerfile` | zwei Zahlen, **ein** Pin |

**Jedes Mal war es dieselbe Angabe an zwei Orten, ohne Abgleich** — nie die Frage, ob eine von
ihnen der Wahrheit entspricht. Diese Klasse braucht kein Netz: sie ist eine Aussage über den
Bestand gegen sich selbst.

Der Sensor trägt darum zwei Regeln:

1. **`uses:`-Kohärenz** — derselbe 40-stellige SHA unter `.github/workflows/` trägt überall
   denselben Tag-Kommentar. Deckt Vorfall 1 und 2.
2. **Pin-Deklarations-Kohärenz** — eine Versions-Variable, die
   [`Makefile`](../../../../Makefile) **und** [`Dockerfile`](../../../../Dockerfile) führen, hat
   an beiden Orten denselben Wert. Deckt Vorfall 3.

**Was er nicht prüft, und was das offen lässt:** ob der Tag-Kommentar den Digest korrekt benennt.
Zwei *übereinstimmend falsche* Angaben bleiben grün. Der Sensor macht Divergenz sichtbar, nicht
Unwahrheit — die verlangte Netz, und `gates` ist hermetisch ([ADR-0037](../../adr/0037-cve-scan-gegen-das-publizierte-image.md)
zieht diese Grenze für den einen Fall, in dem Netz der Zweck ist).

## 5. Abgrenzung

- **Kein d-check-Modul.** Regel 1 wäre ein Kandidat für das Modul `workflows`
  ([slice-130](../done/slice-130-workflows-modul-uses-form.md)); Regel 2 nicht, sie liest
  `Makefile` und `Dockerfile`. Ein CR für Regel 1 bleibt möglich und ist **kein** Teil dieses
  Slice — er ginge durch den Prüf-Durchgang aus [`AGENTS.md`](../../../../AGENTS.md) §5.
- **Kein Versions-Bump.** Die Werte werden an den **gepinnten** Stand angeglichen, nicht gehoben;
  Heben ist der Kanal aus [slice-128](../done/slice-128-dependabot-hebungskanal.md).
- **Kein zweiter Ort für die Wahrheit.** Der Sensor erklärt keine der beiden Deklarationen zur
  führenden — er verlangt nur, dass sie sich nicht widersprechen.

## 6. Risiken und offene Punkte

- *Die Angleichung der Makefile-Werte könnte den Build verändern* —
  **Ausgang:** gestrichen mit Begründung: **gemessen am Lauf**, nicht überlegt. Der `gates`-Lauf
  nach der Änderung meldet **23× `CACHED`** — dieselben Layer wie vorher, weil der Digest
  unverändert ist. Geändert hat sich nur, was in der Referenz **danebensteht**:
  `golang:1.26.4@sha256:0ecdc2a9…` heißt jetzt `golang:1.27.0@sha256:0ecdc2a9…`.
- *Regel 1 könnte legitime Fälle beanstanden — denselben SHA mit absichtlich verschiedenen
  Kommentaren* — **Ausgang:** gestrichen mit Begründung: der Fall existiert nicht. Ein SHA bezeichnet **einen**
  Commit; zwei Namen dafür sind keine Absicht, sondern der Befund. Der Selbsttest hält die
  Gegenrichtung fest (derselbe SHA mehrfach mit **gleichem** Tag schweigt), damit die Regel nicht
  auf Wiederholung, sondern auf Widerspruch anspricht.
- *Der Sensor prüft Kohärenz, nicht Wahrheit; zwei gleich falsche Angaben bleiben grün* —
  **Ausgang:** weiter offen, [`BEO-026`](../observations.md) im Register — der Zähler bleibt bei
  3×, sein Stand nennt jetzt die gedeckte und die offene Hälfte.

## 7. Closure-Notiz

**Geliefert:** `make version-coherence` hängt im `gates`-Aggregat und prüft zwei Kohärenz-Regeln
(§4); die beiden Makefile-Werte tragen den gemessenen Stand des gepinnten Images. Der Sensor fand
beim ersten Lauf genau die zwei Divergenzen, wegen derer er gebaut wurde, und schweigt danach.

**Lerneintrag — Form: geschärfte Regel.** *Eine Beobachtung, die als Netz-Frage formuliert ist,
verhindert den Sensor, den sie eigentlich verlangt — man prüfe zuerst, was die Vorfälle
gemeinsam haben, nicht was ihre Überschrift sagt.* [`BEO-026`](../observations.md) hieß drei Slices
lang „der Tag-Kommentar wird nicht gegen den Digest geprüft". Das stimmt und ist unlösbar, *weil*
es die Registry braucht — und genau darum lag die Beobachtung so lange still: sie beschrieb ihre
eigene Unmöglichkeit. Beim Nachzählen war **keiner** der drei Vorfälle diese Frage. Jeder war
dieselbe Angabe an zwei Orten, verschieden geschrieben (§4) — eine Aussage über den Bestand gegen
sich selbst, hermetisch prüfbar, in einem Skript von unter 140 Zeilen.

**Die beobachtbare Architektur-Aussage** steht in der Referenz selbst: `golang:1.26.4@sha256:0ecdc…`
war eine Zeile, die sich **in sich** widersprach — der Tag nannte eine Version, der Digest enthielt
eine andere, und die Ausführung ignorierte den Tag. Solche Zeilen sind nicht falsch im Sinne von
kaputt, sie sind *unwiderlegbar*: nichts an ihnen bricht, also fällt nichts auf. Ein Sensor ist
dort die einzige Instanz, die überhaupt hinsieht.

**Drei beobachtbare Closure-Kriterien:**

1. Der Sensor meldete am Bestand **zwei** Divergenzen (`GO_VERSION`, `GOLANGCI_LINT_VERSION`),
   nach der Angleichung **null** — bei 3 gepinnten SHAs und 3 doppelt deklarierten Variablen.
2. Die Angleichung ist **folgenlos für den Build**, gemessen: 23× `CACHED`, derselbe Digest,
   dieselben Layer.
3. Der Selbsttest deckt **beide** Richtungen **beider** Regeln — auch die, in der der Sensor
   schweigen muss. Ein Prüfer, der immer feuert, ist von einem korrekten nicht zu unterscheiden;
   dieselbe Überlegung wie in [slice-129](../done/slice-129-risiko-ausgaenge-vor-dem-mv.md).

**Was der Sensor nicht kann, und warum das hier steht:** er erklärt keine der beiden Deklarationen
zur führenden. Läuft etwas auseinander, sagt er *dass*, nicht *welche* stimmt — die Antwort stand
in diesem Slice im gepinnten Image und war eine Messung, keine Ableitung. Ein Sensor, der hier
selbst entschiede, würde raten.

**Offene Risiken und ihr Ausgang:** zwei gestrichen mit Begründung, eines weiter offen.

**Beobachtungs-Register:** [`BEO-026`](../observations.md) ist **verkörpert für die Kohärenz**;
der Zähler bleibt bei 3×, und der Stand nennt beide Hälften — die gedeckte und die offene.

**Folge-Slices:** keiner zwingend. Regel 1 wäre ein CR-Kandidat für das `workflows`-Modul
([slice-130](../done/slice-130-workflows-modul-uses-form.md)) — sie ist dort besser aufgehoben als
hier, weil sie fremde Repos genauso trifft; nötig ist er nicht, der Sensor läuft.
