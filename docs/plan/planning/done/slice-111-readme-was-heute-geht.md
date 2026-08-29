# slice-111 — `README.md`: „What can I do today?"

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.
**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-29.
**Berührte Spec-Stellen:** — *(keine)*
**Deckt:** den letzten offenen Posten aus
[slice-105](../done/slice-105-form-review-nachholen.md) §4.
**Bezug:** [slice-110](../done/slice-110-spec-straten-form.md).
[Roadmap](../in-progress/roadmap.md).

---

## 1. Ziel

Der Projekt-Überblick trägt vier der fünf Sektionen der Ziel-Form. Es fehlt genau eine —
*Was kann ich heute tun?* — und sie ist die einzige, die in der Vorlage einen Regelwerk-Zeiger
trägt: *„ehrlicher Ist-Stand — was **jetzt** läuft, nicht was geplant ist. Keine Erfolgsmeldung
ohne lauffähigen Beleg."*

**Beide Vergleichsverfahren aus slice-105 haben hier danebengelegen** — sie meldeten *drei*
fehlende Sektionen. Der Grund ist banal und lehrreich: `README.md` ist **englisch**, die Vorlage
deutsch. Weder exakter Titel-Vergleich noch Kernbegriff-Vergleich überbrücken das.

## 2. Definition of Done

- [x] `README.md` trägt *What can I do today?* mit dem Regelwerk-Zeiger — Beleg: Diff.
- [x] Die Liste nennt nur, was **läuft**, je mit einem nachprüfbaren Beleg (Befehl, Release,
      Konsument) — Beleg: jede Zeile ist ohne Rückfrage prüfbar.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Plan (vor Code)

| Datei | Änderungs-Art | Begründung |
|---|---|---|
| `README.md` | update | eine Sektion samt Zeiger, direkt nach *What is a-check?* |

**Auszuführende Gates:** `make gates` (tragend `doc-check`, weil die neue Sektion auf Spec-IDs und
Konsumenten verweist), zum Abschluss `make verify`. **Kein neuer Sensor.**

## 4. Trigger

**Start:** unmittelbar.

**Rückführungen:** `in-progress` → `open`, falls eine Zeile der Liste nicht ohne Rückfrage
belegbar ist — dann ist sie eine Absichtserklärung und gehört nicht hinein.

## 5. Closure-Trigger

Sektion steht, jede Zeile belegt, Gates grün.

**Was bewusst nicht getan wird:** `README.md` wird **nicht** ins Deutsche übersetzt. Die Sprache
ist eine Konsumenten-Entscheidung — der Überblick richtet sich an fremde Repos, die Harness-Doku
an dieses. Die Ziel-Form gibt Sektionen vor, keine Sprache.

## 6. Risiken und offene Punkte

- *Die Sprachdifferenz bleibt ein blinder Fleck jedes Form-Vergleichs* — **Ausgang:** weiter
  offen, `BEO-019` im Beobachtungs-Register.
- *„Ist-Stand" veraltet mit dem nächsten Release* — **Ausgang:** gestrichen mit Begründung: die
  Liste nennt Fähigkeiten mit Beleg, keine Versionsnummern; sie veraltet erst, wenn eine Fähigkeit
  wegfällt, und das ist ein eigener Vorgang.

## 7. Closure-Notiz

**Geliefert:** `README.md` trägt *What can I do today?* — fünf Fähigkeiten, jede mit einem Beleg,
den ein Fremder ohne Rückfrage prüfen kann, plus einen Zeiger dorthin, wo das **Nicht**-Vorhandene
steht.

**Lerneintrag — Form: benannte Spec-Lücke.** *Form-Vergleiche gegen die Vorlagen sind sprachblind.*
Beide Verfahren aus slice-105 meldeten für diese Datei drei fehlende Sektionen; fehlte **eine**.
`README.md` ist englisch, die Ziel-Form deutsch — *Core idea* und *Kerngedanke* teilen kein Wort,
und ein Titel-Vergleich sieht das nicht. *Weil* beide Verfahren auf Wortgleichheit beruhen, das
eine streng, das andere unscharf: gegen einen Sprachwechsel hilft weder Strenge noch Unschärfe.
**Der Prüfsatz:** vor jedem Form-Vergleich prüfen, ob Artefakt und Vorlage dieselbe Sprache
sprechen; wenn nicht, ist die Zuordnung Handarbeit.

**Zwei beobachtbare Closure-Kriterien:**

1. Jede der fünf Zeilen nennt einen Beleg, der ohne Rückfrage prüfbar ist — Digest in
   `version.md`, das Fragment `a-check.mk`, zwei namentlich genannte Konsumenten. Keine nennt
   Geplantes.
2. `doc-check` 226 Dateien 0 Befunde — die neuen Verweise lösen auf.

**Offene Risiken und ihr Ausgang:**

- *Die Sprachdifferenz bleibt ein blinder Fleck jedes Form-Vergleichs* — **Ausgang:** weiter
  offen, `BEO-019` im Beobachtungs-Register.
- *`README.md` liegt unter keiner deklarierten Sub-Area* — **Ausgang:** gestrichen mit Begründung:
  eine Sub-Area für eine einzelne Datei zu erfinden wäre teurer als die Lücke; sie ist hier
  benannt.
- *„Ist-Stand" veraltet* — **Ausgang:** gestrichen mit Begründung; die Liste nennt Fähigkeiten mit
  Beleg, keine Versionsnummern.

**Beobachtungs-Register:** `BEO-019` neu angelegt.

**Folge-Slices:** keine — der Form-Review aus slice-105 ist damit abgearbeitet. Was offen bleibt,
steht im Register.

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `README.md` liegt unter **keiner** deklarierten Sub-Area —
es ist Projekt-Überblick, nicht Harness-Einstieg und nicht Spec. Die Lücke ist hier klein genug,
um sie zu benennen statt eine Sub-Area für eine Datei zu erfinden.

**Vorgelagert — offene Beobachtungen sichten:** keine Treffer für diese Datei.

Alle berührten Sub-Areas GF.
