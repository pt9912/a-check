# slice-086 — `forbidden_constructs` ohne `port`-Rolle: fail-closed statt still

**Status:** open — der Zustand ist das **Verzeichnis**, nicht dieses Feld
(`open/ → next/ → in-progress/ → done/`, Wechsel nur per `git mv` als eigener Commit,
[`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).
**Deckt:** `CR-4` des Konsumenten-Einsatzes (2026-08-09);
[SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema),
[AC-FA-RULE-004](../../../../spec/lastenheft.md#ac-fa-rule-004--port-disziplin-regel-port-impurity).
**Bezug:** Welle [`welle-13-konsumenten-befunde`](../welle-13-konsumenten-befunde.md).

---

## 0. Trigger

**Beginn:** sofort.

**Rückführungen:**

- `in-progress` → `open`: falls der Entscheid aus §3 zur Ausweitung statt zum Fehlerfall geht —
  das wäre eine Verhaltensänderung mit eigener Regel-Zeile in
  [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) und damit ein
  anderer Slice.

## 1. Auslöser

**Mechanismus: eine Konfiguration wird streng validiert, angenommen — und nie ausgewertet.**

`forbidden_constructs` wird **je Schicht** konfiguriert, aber nur über die Rolle `port`
ausgewertet. Im Code, [`internal/hexagon/core/rules.go`](../../../../internal/hexagon/core/rules.go)
Zeile 34, hängt die gesamte Auswertung an einer Rollen-Bedingung: nur wenn die Schicht die Rolle
`port` trägt, werden die Treffer zu Befunden.

Ein Eintrag für eine `domain`-, `app`- oder Adapter-Schicht durchläuft die strenge Dekodierung,
wird angenommen — und wirkt nie. Gemessen im Konsumenten-Einsatz: **sechs Schichten** mit einem
Include-Muster → **0 Befunde** bei vorhandenem Verstoß.

**Das widerspricht der eigenen Linie.** Dieses Repo lehnt stille Defaults durchgehend ab: ein
unbekannter `languages`-Schlüssel bricht mit Exit 2
([AC-FA-CONF-001](../../../../spec/lastenheft.md#ac-fa-conf-001--konfigurationsdatei-a-checkyml),
0.9.0 — „schließt die stille Nicht-Extraktion (falsch-grün)"), ein leerer `tech.adapter` ebenso
(0.14.0 — „fail-closed statt vormals stillem Never-Leak-Eintrag"). Hier ist derselbe Fall bisher
still.

## 2. Betroffene Module

Zwei Schichten:

1. **Spec** — [`spec/spezifikation.md`](../../../../spec/spezifikation.md)
   ([SPEC-CONF-001](../../../../spec/spezifikation.md#spec-conf-001--konfigurationsschema): der
   Fehlerfall gehört ins Schema) und die Exit-2-Liste.
2. **`internal/adapter/driven/config/`** — die Validierung beim Laden.

## 3. Auszuführende Gates

`make gates`, `make image-test`.

**Der Entscheid, der vor dem Bau fällt** — der CR nennt beide Wege:

| Weg | Folge |
|---|---|
| **Fail-closed** (CR-Vorschlag) — Eintrag für eine Schicht ohne Rolle `port` ergibt **Exit 2**, mit Verweis auf `constructs` als zonen-gebundenes Gegenstück | passt zur bestehenden Linie; **bricht** aber Konfigurationen, die den Eintrag heute wirkungslos tragen |
| **Auswertung ausweiten** — der Block gilt für alle Rollen | Verhaltensänderung statt Fehlerfall; braucht eine eigene Regel-Zeile in [SPEC-RULE-001](../../../../spec/spezifikation.md#spec-rule-001--regel-auswertung) und einen neuen Befund-Namen, weil `port-impurity` dann nicht mehr passt |

**Zu prüfen vor dem Entscheid:** ob im bekannten Konsumenten-Bestand solche Einträge existieren.
Fail-closed auf eine verbreitete Konfiguration wäre ein Breaking Change ohne Vorwarnung; ist der
Bestand leer, ist es eine reine Ehrlichkeits-Korrektur.

**Negativ-Proben:**

| Probe | Erwartung |
|---|---|
| Der Block auf einer `domain`-Schicht | je nach Entscheid Exit 2 **oder** ein Befund — in **keinem** Fall stilles Grün |
| Der Block auf einer `port`-Schicht | unverändert `port-impurity` |
| Konfiguration ohne den Block | byte-identisch |

## 4. Was bewusst nicht getan wird

- **`constructs` anfassen.** Der zonen-gebundene Block
  ([AC-FA-RULE-011](../../../../spec/lastenheft.md#ac-fa-rule-011--konstrukt-monopol-regel-construct-leak))
  ist das funktionierende Gegenstück und bleibt unberührt; er ist im Fehlertext zu nennen.
- **Die Rollen-Inferenz ändern.** Dass die Rolle heuristisch aus dem Schicht-Namen abgeleitet wird,
  ist eine andere Frage.

## 5. DoD

- [ ] Der Entscheid aus §3 ist **getroffen und begründet** in der Closure-Notiz, mit der Messung
      des Konsumenten-Bestands als Grundlage.
- [ ] Ein Eintrag auf einer Nicht-`port`-Schicht führt **nicht** mehr zu stillem Grün. Beleg: die
      ersten beiden Proben aus §3.
- [ ] `make gates` und `make image-test` grün — **Ausgabe in eine Datei**, Exit-Code getrennt
      geprüft, nie in eine Pipe.

## 6. Closure-Notiz

_(beim Abschluss ausfüllen — genau **ein** solcher Abschnitt je Slice,
[`AGENTS.md`](../../../../AGENTS.md) §5; `make verify` prüft das.)_

## 7. Sub-Area-Modus

Alle berührten Sub-Areas GF.
