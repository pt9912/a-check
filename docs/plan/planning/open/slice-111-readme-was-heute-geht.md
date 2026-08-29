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

- [ ] `README.md` trägt *What can I do today?* mit dem Regelwerk-Zeiger — Beleg: Diff.
- [ ] Die Liste nennt nur, was **läuft**, je mit einem nachprüfbaren Beleg (Befehl, Release,
      Konsument) — Beleg: jede Zeile ist ohne Rückfrage prüfbar.

- [ ] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [ ] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [ ] Beobachtungs-Register fortgeschrieben.
- [ ] Jedes Risiko aus §6 trägt genau einen Ausgang.

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

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 8. Sub-Area-Modus-Begründung

**Vorgelagert — Sub-Area-Wahl prüfen:** `README.md` liegt unter **keiner** deklarierten Sub-Area —
es ist Projekt-Überblick, nicht Harness-Einstieg und nicht Spec. Die Lücke ist hier klein genug,
um sie zu benennen statt eine Sub-Area für eine Datei zu erfinden.

**Vorgelagert — offene Beobachtungen sichten:** keine Treffer für diese Datei.

Alle berührten Sub-Areas GF.
