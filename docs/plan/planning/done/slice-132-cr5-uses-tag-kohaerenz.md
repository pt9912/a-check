# slice-132 — CR 5 an d-check: ein SHA, ein Tag-Kommentar

**Lifecycle:** Der Zustand dieses Slice ist das Verzeichnis, in dem diese Datei liegt — eines von
`open/`, `next/`, `in-progress/`, `done/`. Er wechselt nur durch `make slice-mv`
([`AGENTS.md`](../../../../AGENTS.md) §3.3/§5).

**Bezug:** [`BEO-026`](../observations.md) (3×, für die Kohärenz verkörpert in
[slice-131](../done/slice-131-versions-kohaerenz.md)) und
[`BEO-022`](../observations.md) — der Prüf-Durchgang aus
[`AGENTS.md`](../../../../AGENTS.md) §5 ist Pflicht, bevor ein CR-Text hinausgeht.

**Berührte Spec-Stellen:** — *(keine)* — ein CR an ein fremdes Werkzeug ändert a-checks Verträge
nicht.

**Verantwortlich:** Implementation (diese Sitzung); Abnahme und Versand beim Maintainer.

**Autor:** Claude (Opus 5), im Auftrag des Maintainers. **Datum:** 2026-08-30.

---

## 1. Ziel

Die erste Regel aus [slice-131](../done/slice-131-versions-kohaerenz.md) liegt dort, wo sie nur
a-check hilft. Sie gehört in das Modul, das dieselben Zeilen ohnehin parst — als Antrag, nicht als
Behauptung.

## 2. Definition of Done

- [x] Der CR-Text steht vollständig in §4 — mit der Abgrenzung gegen das **erklärte**
      Out-of-Scope von `DC-FA-WF-001`, ohne die der Antrag zu Recht abgelehnt würde.
- [x] Der Prüf-Durchgang nach
      [`cr-text-reviewer.md`](../../../../.harness/skills/cr-text-reviewer.md) ist gefahren und
      in §5 protokolliert: je Tatsachen-Behauptung eine Zeile mit dem Handgriff, der sie belegt.
- [x] Beide Zahlen des Schluss-Satzes stehen da — wie viele Behauptungen der Text trägt und wie
      viele belegt sind.

- [x] `make gates` grün — Ausgabe in eine Datei, Exit-Code getrennt geprüft, nie in eine Pipe.
- [x] Closure-Notiz mit benanntem Lerneintrag geschrieben (§7).
- [x] Beobachtungs-Register fortgeschrieben.
- [x] Jedes Risiko aus §6 trägt genau einen Ausgang.

## 3. Was vorher gemessen wurde

Der Prüf-Durchgang verlangt drei Eingänge; alle drei liegen vor:

| Eingang | Handgriff |
|---|---|
| fremdes Lastenheft | `DC-FA-WF-001` gelesen — Bedingungs-Tabelle, „Warum der Pin", Out-of-Scope |
| eigener Bestand | `.github/workflows/`: **8** `uses:`-Referenzen, davon **7** fremde mit 40-hex-SHA und **1** lokale; **3** distinkte SHAs |
| fremder Bestand | d-checks `.github/workflows/`: **3** distinkte SHAs, je **ein** Tag-Kommentar — dort liegt heute **kein** Konflikt |

Der dritte Eingang ist der, der den Antrag ehrlich macht: die Regel fände beim Empfänger **nichts**.
Sie zu beantragen, als läge dort ein Bestand brach, wäre die Ausprägung (B) aus dem Skill.

## 4. Der CR-Text

> **CR 5 an d-check — `uses-pin-tag-conflict`: ein SHA, ein Tag-Kommentar**
>
> **Antrag.** Eine dritte Bedingung in der Pin-Familie von `DC-FA-WF-001`: *derselbe
> 40-stellige SHA trägt innerhalb der Scan-Menge überall denselben Tag-Kommentar.* Vorschlag für
> den Grund-Code: `uses-pin-tag-conflict`, gemeldet auf **jeder** beteiligten Zeile, weil keine
> von ihnen für sich falsch ist.
>
> **Das ist nicht euer Out-of-Scope.** `DC-FA-WF-001` nimmt die **Gültigkeit** eines SHA
> ausdrücklich aus — „ob der SHA existiert und den Commit bezeichnet, den der Tag-Kommentar
> behauptet, ist eine Netz-Frage und ausdrücklich außerhalb". Dem widerspricht dieser Antrag
> nicht. Er fragt nicht, ob ein Kommentar **wahr** ist, sondern ob zwei Kommentare **einander**
> widersprechen. Das ist eine Aussage der Scan-Menge gegen sich selbst: dieselbe Eingabe, die das
> Modul ohnehin parst, kein Netz, kein git, keine zusätzliche Datei.
>
> **Warum es in die Pin-Familie gehört und nicht daneben.** `uses-pin-untagged` erzwingt bereits
> die **Existenz** des Tag-Kommentars. Damit ist er im Vertrag keine Dekoration, sondern eine
> Zusage — und für eine Zusage ist „vorhanden" die schwächste denkbare Prüfung. Ein Kommentar,
> den zwei Zeilen verschieden schreiben, ist an mindestens einer Stelle falsch; welche, sagt die
> Regel nicht, und das muss sie auch nicht. Sie sagt: hier stimmt etwas nicht, seht nach.
>
> **Der Anlass, gemessen.** a-checks `release.yml` pinnte denselben `docker/login-action`-Digest
> zweimal — einmal `# v4.2.0`, einmal `# v3.6.0`, 83 Zeilen auseinander, beide Male derselbe SHA
> `650006c6…`. Der zweite Kommentar entstand beim Kopieren der ersten Zeile. Über die
> GitHub-API aufgelöst: `v4.2.0` → `650006c6…`, `v3.6.0` → `5e57cd11…`; der Digest ist v4.2.0, der
> zweite Kommentar war falsch. Ein zweiter Fall derselben Art lag zwischen zwei Dateien
> (`# v5.0.0` gegen `# v6.0.2` am selben `actions/checkout`-Digest).
>
> **Wir haben gemessen, dass das Modul ihn heute nicht meldet.** `d-check --enable workflows`
> (Pin `v0.69.0`) lief über a-check, **während beide widersprüchlichen Zeilen im Repo standen**,
> und meldete genau einen Befund — `uses-local-perms-undeclared` in `release.yml:242`. Der
> Tag-Konflikt war nicht darunter. Das ist erwartungsgemäß: beide Zeilen tragen einen
> Tag-Kommentar, `uses-pin-untagged` greift also bei keiner.
>
> **Bei euch fände die Regel heute nichts.** d-checks `.github/workflows/` führt drei distinkte
> SHAs mit je einem Tag-Kommentar; `actions/checkout` steht fünfmal mit identischem Kommentar.
> Wiederholung ist kein Befund — nur Widerspruch. Der Antrag stützt sich also auf **unseren**
> Bestand, nicht auf euren; er ist eine Regressions-Bremse, kein Bestandsräumer. Dass beide Repos
> dieselbe Pin-Konvention fahren und der Fall trotzdem zweimal bei uns entstand, ist genau das
> Argument: die Konvention allein verhindert ihn nicht.
>
> **Aufwand, soweit wir ihn von außen einschätzen können.** Wir haben die Regel bei uns als
> Bash-Sensor gebaut; unsere Fassung braucht 19 Zeilen (Extraktion plus Gruppierung nach SHA) —
> das sagt etwas über unsere Umgebung, nicht über eure Codebasis. In eurer liegen die Referenzen
> bereits geparst vor; was daraus folgt, wisst ihr besser als wir.
>
> **Was wir nicht beantragen.** Keine Aussage darüber, welcher der widersprüchlichen Kommentare
> der richtige ist — das wäre die Gültigkeitsfrage und damit Netz. Und keine Ausweitung auf
> Referenzen ohne Tag-Kommentar; die deckt `uses-pin-untagged` bereits.
>
> **Wenn ihr ablehnt,** bleibt die Regel bei uns als lokaler Sensor stehen und tut dort ihren
> Dienst. Der Antrag ist eine Einordnungsfrage — gehört diese Zusage in die Pin-Familie oder in
> die Harness des Konsumenten? —, keine Blockade.

## 5. Prüf-Durchgang (Skill `cr-text-reviewer`)

Je Satz, der eine **Tatsache über ein System** behauptet — eine Zeile, mit dem Handgriff:

```
Antrag        :: "DC-FA-WF-001 nimmt die Gültigkeit eines SHA ausdrücklich aus" :: fremder Vertrag
              :: Out-of-Scope-Absatz von DC-FA-WF-001 gelesen und wörtlich zitiert
Antrag        :: "uses-pin-untagged erzwingt bereits die Existenz des Tag-Kommentars" :: fremder Vertrag
              :: Bedingungs-Tabelle von DC-FA-WF-001, Zeile 2
Anlass        :: "release.yml pinnte denselben Digest zweimal, v4.2.0 und v3.6.0" :: eigener Bestand
              :: git blame auf release.yml:75 und :158 vor der Korrektur (slice-130)
Anlass        :: "v4.2.0 → 650006c6…, v3.6.0 → 5e57cd11…" :: fremder Bestand
              :: gh api repos/docker/login-action/git/ref/tags/<tag>, beide Tags aufgelöst
Anlass        :: "ein zweiter Fall lag zwischen zwei Dateien (v5.0.0 gegen v6.0.2)" :: eigener Bestand
              :: BEO-026, Beleg slice-124
Messung       :: "das Modul meldet ihn heute nicht — ein Befund, und der war es nicht" :: fremder Vertrag
              :: d-check v0.69.0 mit --enable workflows über a-check gefahren, WÄHREND beide
                 Zeilen im Repo standen; Ausgabe: genau `uses-local-perms-undeclared`
Empfänger     :: "d-check führt drei distinkte SHAs mit je einem Tag-Kommentar" :: fremder Bestand
              :: grep über /Development/d-check/.github/workflows/, nach SHA gruppiert
Empfänger     :: "actions/checkout steht fünfmal mit identischem Kommentar" :: fremder Bestand
              :: derselbe grep, uniq -c
Aufwand       :: "unsere Fassung braucht 19 Zeilen" :: eigener Bestand
              :: awk über die beiden Funktionen in tools/verify-versions-kohaerent.sh
```

**Neun Tatsachen-Behauptungen, neun belegt.**

Zwei Sätze sind **ohne** Handgriff geblieben und bleiben stehen, weil sie keine Tatsache über ein
System behaupten: *„für eine Zusage ist ‚vorhanden' die schwächste denkbare Prüfung"* (ein
Argument) und *„der Antrag ist eine Einordnungsfrage, keine Blockade"* (eine Haltung).

**Ein Satz wurde beim Durchgang umgeschrieben.** Der Entwurf sagte, die Regel sei „in eurer
Codebasis billig". Das ist eine Tatsache über ein fremdes System, für die kein Handgriff existiert
— a-check kennt d-checks Modul-Interna nicht. Er steht jetzt als Aussage über die **eigene**
Fassung da, mit der Grenze im Satz („das sagt etwas über unsere Umgebung, nicht über eure
Codebasis"). Das ist genau die Ausprägung (B) des Skills: die eigene Menge messen und über die
fremde aussagen.

## 6. Risiken und offene Punkte

- *Der Empfänger lehnt ab, und die Regel bleibt dauerhaft doppelt gepflegt* —
  **Ausgang:** gestrichen mit Begründung: „doppelt gepflegt" trifft nicht zu. Bei Ablehnung bleibt
  **eine** Fassung stehen — die eigene, die seit [slice-131](../done/slice-131-versions-kohaerenz.md)
  ohnehin läuft. Der Antrag kostet nichts, wenn er scheitert, und der Text sagt das dem Empfänger
  auch, statt Druck zu erzeugen.
- *Der Text behauptet etwas über d-checks Interna, das a-check nicht wissen kann* —
  **Ausgang:** **eingetreten** im Entwurf, im Prüf-Durchgang gefangen und behoben (§5, letzter
  Absatz): ein Satz nannte die Regel „in eurer Codebasis billig" und steht jetzt als Aussage über
  die eigene Fassung da, mit der Grenze im Satz. Kein Folge-Slice — der Skill hat genau das
  geleistet, wofür er da ist; [`BEO-022`](../observations.md) bleibt bei 3× und verkörpert.

## 7. Closure-Notiz

**Geliefert:** CR 5 steht vollständig in §4, der Prüf-Durchgang in §5 — neun Tatsachen-Behauptungen,
neun belegt, mit dem Handgriff je Zeile. Der Text ist übergabefertig; der Versand liegt beim
Maintainer.

**Lerneintrag — Form: geschärfte Regel.** *Ein Antrag, der beim Empfänger nichts fände, wird
stärker, wenn er das sagt — nicht schwächer.* Die Messung an d-checks eigenen Workflows ergab
**null** Konflikte (drei SHAs, je ein Tag). Der Reflex ist, diesen Absatz wegzulassen: er nimmt dem
Antrag die Dringlichkeit. Er trägt sie aber gerade — *weil* beide Repos dieselbe Pin-Konvention
fahren und der Fall trotzdem zweimal bei uns entstand, ist belegt, dass die Konvention allein ihn
nicht verhindert. Ein verschwiegener Nullbefund wäre zudem die Ausprägung (B) aus dem Skill in
ihrer bequemsten Form: nicht falsch messen, sondern nicht erwähnen.

**Der Skill hat gegriffen, und zwar an der Stelle, für die er gebaut ist.** Der Entwurf nannte die
Regel „in eurer Codebasis billig" — eine Tatsache über ein fremdes System, für die a-check keinen
Handgriff hat. Sie klang wie Entgegenkommen und war eine Behauptung über fremde Interna. Der Satz
steht jetzt als Aussage über die **eigene** Fassung da, mit der Grenze ausgesprochen. Das ist der
vierte belegte Fall der Klasse *„gemessen wird die eigene Menge, ausgesagt wird über die fremde"* —
und der erste, den der Skill vor dem Versand gefangen hat statt der Empfänger danach.

**Drei beobachtbare Closure-Kriterien:**

1. Der Prüf-Durchgang nennt **beide** Zahlen (neun Behauptungen, neun belegt). Eine Zahl ohne die
   andere ist laut Skill keine Auskunft.
2. Die Abgrenzung gegen das erklärte Out-of-Scope steht **zitiert** im Antrag, nicht paraphrasiert
   — ohne sie wäre CR 5 als Netz-Frage zu Recht abgelehnt worden.
3. Die stärkste Zeile des Antrags ist eine Messung am **fremden** System: `d-check --enable
   workflows` lief über a-check, während beide widersprüchlichen Zeilen im Repo standen, und
   meldete den Konflikt nicht.

**Was offen bleibt:** ob der Empfänger den Antrag annimmt. Das ist keine offene Frage dieses Slice
— die Regel läuft lokal, und der Text sagt das ausdrücklich.

**Offene Risiken und ihr Ausgang:** eines gestrichen mit Begründung, eines eingetreten und im
Durchgang behoben.

**Beobachtungs-Register:** [`BEO-022`](../observations.md) bleibt bei **3×** und verkörpert; dieser
Slice belegt den Skill zum ersten Mal an einem CR, den er **vor** dem Versand korrigiert hat, und
der Stand nennt das jetzt. [`BEO-026`](../observations.md) bleibt unverändert — der Antrag ändert
nichts an der Deckung, solange er nicht angenommen ist.

**Folge-Slices:** keiner. Wird CR 5 angenommen, ist die Ablösung der lokalen Regel 1 ein eigener,
dann erst schneidbarer Slice.
