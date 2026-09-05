# slice-154 — Spec: fehlende v6.0.0-Pflichtabschnitte nachgezogen

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** Systematischer Abgleich aller a-check-Instanzdokumente gegen ihre v6.0.0-Templates
(Maintainer-Auftrag 2026-09-05). Zwei der acht Funde betreffen die kanonischen Spec-Straten:
`spec/lastenheft.template.md` verlangt §5/§6, die a-check bislang übersprang (Nummerierung
sprang 4→7); `spec/architecture.template.md` verlangt §3 (Externe Abhängigkeiten) und §5
(Fehlermodelle), die a-check ebenfalls nicht führte. [Roadmap](../in-progress/roadmap.md).

**Berührte Spec-Stellen:** `spec/lastenheft.md` §5/§6 (neu), `spec/architecture.md` §4/§6 (neu,
Folge-Sektionen umnummeriert).

**Verantwortlich:** Implementation (diese Sitzung); Abnahme beim Maintainer.

**Autor:** Claude (Sonnet 5), im Auftrag des Maintainers. **Datum:** 2026-09-05.

---

## 1. Ziel

Beide fehlenden Abschnitts-Paare nachziehen — **ohne** neue Fakten zu behaupten: beide Straten
konsolidieren, was an anderer Stelle im jeweiligen Dokument bereits belegt steht, an der von der
Ziel-Form vorgesehenen Stelle.

## 2. Vorab geprüft: braucht das einen Change Request?

`spec/lastenheft.md` §7 (Historie) bindet die CR-Pflicht **an den Status**: „Ab Status `Accepted`
ist jede Änderung … eine Vertragsänderung." Beide Straten führen `**Status:** Draft` — die
CR-Pflicht greift noch nicht. Normale Spec-Bearbeitung, kein CR.

## 3. `spec/lastenheft.md` — §5 Globale Out-of-Scope-Punkte, §6 Glossar

- **§5:** Der bislang in §1 stehende Absatz „Out of Scope (Produkt)" (drei Aussagen: keine
  compile-time-Modulgrenzen-Ersetzung, Heuristik statt Parser, kein Auto-Fix) an die von der
  Ziel-Form vorgesehene Stelle verschoben; §1 zeigt jetzt per Anker-Verweis dorthin. Unterschieden
  von den **per-AC** Out-of-Scope-Zeilen in §3/§4 (die bleiben unverändert, begrenzen nur ihren
  jeweiligen AC).
- **§6:** Fünf Begriffe, die im Dokument bereits feststehend, aber nirgends gesammelt definiert
  sind — Definition je Begriff aus dem AC übernommen, der ihn einführt: Schicht-Rolle
  ([AC-FA-RULE-006](../../../../spec/lastenheft.md#ac-fa-rule-006--schicht-rollen-generische-regel-anwendung)),
  Sub-Einheit ([AC-FA-RULE-002](../../../../spec/lastenheft.md#ac-fa-rule-002--keine-lateralen-adapter-kanten-regel-lateral-adapter)),
  Composition Root ([AC-FA-RULE-003](../../../../spec/lastenheft.md#ac-fa-rule-003--tech-kapselung-regel-tech-leak)),
  driving/driven ([AC-FA-RULE-008](../../../../spec/lastenheft.md#ac-fa-rule-008--richtungs-dimension-regel-port-direction-mismatch)),
  Heuristik-Grenze ([AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)).
  Kein neuer Begriff erfunden — nur bereits verwendete gesammelt.
- Versions-Bump `0.26.0` → `0.27.0`, Historie-Zeile ergänzt.

## 4. `spec/architecture.md` — §4 Externe Abhängigkeiten, §6 Fehlermodelle und Resilienz

- **Neues §4:** `a-check` hat **keine** Laufzeit-Abhängigkeit (netzlos, `--network none`,
  [AC-QA-02](../../../../spec/lastenheft.md#ac-qa-02--hermetik-und-ehrliche-heuristik-grenze)) — die
  einzigen externen Berührungspunkte liegen auf dem **Distributions**-Weg:
  [ARC-008](../../../../spec/architecture.md#4-externe-abhängigkeiten) GHCR (primäre Registry) und
  [ARC-009](../../../../spec/architecture.md#4-externe-abhängigkeiten) Docker Hub (Spiegel,
  [AC-FA-DIST-002](../../../../spec/lastenheft.md#ac-fa-dist-002--docker-hub-spiegel)). Zwei neue
  Kennungen, fortlaufend nach dem bestehenden Höchststand
  [ARC-007](../../../../spec/architecture.md#2-komponenten).
- **Neues §6:** Fehlerquelle-je-Behandlung-Schicht-Tabelle, aus dem bereits in
  [AC-FA-CLI-001](../../../../spec/lastenheft.md#ac-fa-cli-001--aufruf-scan-wurzel-und-exit-codes)
  belegten Exit-Code-Vertrag (0/1/2) abgeleitet — kein neuer Fakt, nur die Schicht-Zuordnung
  ergänzt, die die Sicht (nicht der Vertrag) trägt.
- **Folge-Sektionen umnummeriert** (altes §4 Sequenz → §5, altes §5 Geltung der Constraints → §7,
  altes §6 Historie → §8). Vorab geprüft: nur `spec/architecture.md#2-komponenten` wird von außen
  zitiert ([`docs/plan/adr/0008-...md`](../../adr/0008-ports-duerfen-domaenen-typen-referenzieren.md),
  zweimal) — §2 bleibt unverändert, kein Verweis bricht.
- Versions-Bump `0.3.0` → `0.4.0`, Historie-Zeile ergänzt.

## 5. Nebenfund: stale Versions-Referenz

`docs/user/releasing.md` nannte „Das Lastenheft steht bei 0.26.0" — vom `versions`-Modul als
`version-stale` gemeldet, sobald §3 den Bump auf `0.27.0` setzte. Nachgezogen; kein eigener
Liefer-Punkt, reine Konsequenz von §3.

## 6. Auszuführende Gates

`make gates`, zum Abschluss `make verify`. Kein neuer Sensor.

## 7. DoD

- [x] CR-Pflicht vorab geprüft (Status `Draft` an beiden Straten, keine Pflicht) statt
      angenommen (§2).
- [x] Beide Abschnitts-Paare nachgezogen, ausschließlich aus bereits belegten Aussagen
      konsolidiert, keine neue Aussage erfunden (§3, §4).
- [x] `make gates` und `make verify` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft,
      nie in eine Pipe.

## 8. Closure-Notiz

**Geliefert:** beide Spec-Straten führen jetzt alle sechs bzw. acht von der `v6.0.0`-Ziel-Form
verlangten Abschnitte. Kein neuer Fakt behauptet — beide Nachträge konsolidieren, was im
jeweiligen Dokument bereits stand.

**Lerneintrag — Form: benannte Spec-Lücke.** *Die Nummerierungs-Lücke selbst war der Hinweis* —
`spec/lastenheft.md` sprang von §4 direkt zu §7, ohne dass irgendein Sensor das je meldete
(kein Gate prüft Abschnitts-Vollständigkeit gegen die Ziel-Form, nur `doc-structure`s
Slice-/Closure-Regeln). Diese Lücke ist damit eine **benannte, aber noch nicht gedeckte**
Sensor-Fähigkeit: ein Modul, das ein Spec-Stratum gegen seine Ziel-Form-Abschnittsliste hält,
existiert nicht. Ob es eines braucht (und wie teuer ein solcher Vergleich für ein Draft-Dokument
wäre, das sich noch häufig ändert), ist hier **nicht** entschieden — nur die konkrete Lücke
geschlossen, nicht die Klasse.

**Zwei beobachtbare Closure-Kriterien:**

1. `grep -c '^## ' spec/lastenheft.md` → mindestens 7 (§1…§7); `grep -c '^## ' spec/architecture.md`
   → 8 (§1…§8).
2. `make gates` und `make verify` je **Exit 0** auf dem Stand dieses Slice, Ausgabe in Dateien,
   Exit-Codes getrennt geprüft.

**Offene Risiken und ihr Ausgang.** Keines notiert — beide Nachträge sind vollständig, kein
offener Rest an diesem konkreten Fund. Die im Lerneintrag benannte Sensor-Lücke (Spec-Form gegen
Ziel-Form) ist **nicht** als Risiko notiert — sie ist eine mögliche künftige Fähigkeit, kein aus
diesem Slice folgendes offenes Risiko.

**Folge-Slices:** keine vergeben.

## 9. Sub-Area-Modus

Berührt: **Spec-Straten** (`spec/lastenheft.md`, `spec/architecture.md`) — Greenfield: Form steht
in der vendored Vorlage; keine bestehende Prüfung hält die Abschnittsliste gegen sie (benannte
Lücke, §Closure-Notiz).
