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

- [ ] Die beiden Versions-Variablen im [`Makefile`](../../../../Makefile) tragen den **gemessenen**
      Wert des jeweils gepinnten Images (§3), nicht den zuletzt geschriebenen.
- [ ] `tools/verify-versions-kohaerent.sh` prüft **zwei** Kohärenz-Regeln (§4) und trägt einen
      Selbsttest, der beide Richtungen jeder Regel abdeckt — auch die, in der der Sensor nichts
      finden **darf**.
- [ ] `make version-coherence` ist im `gates`-Aggregat und in
      [`AGENTS.md`](../../../../AGENTS.md) §4 sowie
      [`harness/README.md`](../../../../harness/README.md) deklariert.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

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
  **Ausgang:** *(beim Abschluss ausfüllen)*
- *Regel 1 könnte legitime Fälle beanstanden — denselben SHA mit absichtlich verschiedenen
  Kommentaren* — **Ausgang:** *(beim Abschluss ausfüllen)*
- *Der Sensor prüft Kohärenz, nicht Wahrheit; zwei gleich falsche Angaben bleiben grün* —
  **Ausgang:** *(beim Abschluss ausfüllen)*

## 7. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice.)_
