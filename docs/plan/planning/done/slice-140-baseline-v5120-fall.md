# slice-140 — Alten Baseline-Baum `v5.12.0` fallen lassen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Folge-Slice aus [slice-139 §8](../done/slice-139-beobachtungsregister-migration.md#8-closure-notiz)
— Etappe A ([slice-136](../done/slice-136-baseline-v600-vendoring.md) §1) hatte den alten Baum
ausdrücklich bis zum Abschluss des Form-Reviews stehen lassen (`modul-02` §Freshness-Audit); der
Review ist mit slice-137 (Etappe B, Adaptions-Durchgang) und slice-138/139 (Etappe C, Kürzel +
Migration) abgeschlossen. [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** — *(keine)* — Harness-Wartung ohne Vertragsberührung.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

`.harness/baseline/v5.12.0/` fällt; `.harness/baseline/v6.0.0/` bleibt als einziger vendored Stand
— genau der Zustand, für den `make regelwerk-check` gebaut ist (ein Stand ohne Hinweis auf
ungeprüfte Nachbarn).

## 2. Betroffene Module

- `.harness/baseline/v5.12.0/` — 51 Dateien entfernt.
- `harness/conventions.md` §Baseline — der Übergangssatz („aktuell liegt zusätzlich `v5.12.0`
  vendored, bis der Form-Review …") entfällt.

Eine Schicht (Vendored Baseline + ihre Deklaration in Harness-Einstieg).

## 3. Warum jetzt sicher

- **Keine Live-Links zeigen auf den Baum.** Die verbleibenden `v5.12.0`-Erwähnungen außerhalb von
  `done/` und dem Baum selbst (`harness/conventions.md`, sechs `MR-<NNN>`-Dateien) sind
  ausnahmslos Versionsnummern in Backticks — keine `[Text](.harness/baseline/v5.12.0/…)`-Links,
  gemessen mit `grep -rn`. `make doc-check` prüft nur echte Links; das Entfernen des Baums bricht
  keinen.
- **Die sechs `MR`-Dateien bleiben unangetastet** (Append-only, [`AGENTS.md`](../../../../AGENTS.md)
  §3.5 sinngemäß) — ihre Erwähnungen sind historische Sachaussagen über den Stand, unter dem sie
  entstanden, keine Zeiger auf den Baum.
- **`make regelwerk-check` wählt ohnehin nur den höchsten Stand.** Der Fall des alten Baums ändert
  am geprüften Ergebnis nichts — er entfernt nur den bislang namentlich ausgewiesenen ungeprüften
  Rest.

## 4. Auszuführende Gates

`make regelwerk-check` (kein Gate, Wartung: muss nach dem Fall **einen** Stand melden, keinen
Hinweis mehr auf ungeprüfte Nachbarn), dann `make gates`. Zum Abschluss `make verify`.

## 5. Was bewusst nicht getan wird

- **Keine Bewertung der sechs `MR`-Dateien.** Ihr Inhalt ist unverändert korrekt; dieser Slice
  bewegt nur den Baum, nicht die Adaptionen.
- **Kein [`MR-013`](../../../../harness/conventions/MR-013-adr-vorlagen-version.md)-Nachfolger.** Eigener, bereits benannter Folge-Slice (slice-137 §8).

## 6. DoD

- [x] `.harness/baseline/v5.12.0/` entfernt; `.harness/baseline/` enthält nur noch `v6.0.0/`.
- [x] `harness/conventions.md` §Baseline beschreibt den Ist-Zustand (ein vendorter Stand) statt der
      Übergangsphase.
- [x] `make regelwerk-check` meldet **einen** Stand ohne Hinweis auf ungeprüfte Nachbarn;
      `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie
      in eine Pipe.

## 7. Closure-Notiz

**Geliefert:** der Fall des alten Baseline-Baums — 51 Dateien entfernt, `harness/conventions.md`
§Baseline auf den Ist-Zustand gekürzt. `make regelwerk-check` bestätigt einen einzelnen vendorten
Stand.

**Lerneintrag — Form: benannte Spec-Lücke.** *Der Freshness-Audit-Mechanismus der Baseline
(`modul-02`) verlangt, den alten Baum „bis der Review durch ist" stehen zu lassen, definiert aber
nicht, wer den Abschluss dieses Reviews FESTSTELLT — hier war es implizit „die letzte Etappe des
eigenen Migrationsplans ist committet", eine repo-lokale Setzung, keine Baseline-Vorschrift.* Für
a-check trägt genau dieser Slice die Feststellung; ein anderes Repo mit einem formalen
Review-Sign-off (menschliche Freigabe, ein eigener Gate) bräuchte hier einen anderen Auslöser. Die
Baseline benennt das Ende der Übergangsphase nicht als Ereignis mit Träger — das ist keine Lücke,
die a-check schließen müsste (der eigene Auslöser trägt genug), aber sie ist benannt, damit ein
künftiger Bump nicht dieselbe Frage stillschweigend neu beantwortet.

**Zwei beobachtbare Closure-Kriterien:**

1. `find .harness/baseline -mindepth 1 -maxdepth 1 -type d` listet ausschließlich `v6.0.0`.
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Keine notiert — der Fall ist mechanisch und in §3 begründet,
ohne offenen Rest.

**Folge-Slices:** keine vergeben.

## 8. Sub-Area-Modus

Berührt: **Vendored Baseline** (`.harness/baseline/`, kein Modus, [`MR-006`](../../../../harness/conventions.md#mr-006--baseline-committet-vendored-statt-per-url-referenziert))
und **Harness-Einstieg** (`harness/conventions.md`) — Greenfield.
