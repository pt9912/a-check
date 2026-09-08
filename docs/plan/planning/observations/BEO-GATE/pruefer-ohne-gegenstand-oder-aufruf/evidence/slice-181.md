**Vorgang:** slice-181 (Planung; Fund beim Maintainer-Hinweis am 2026-09-07)

**Fund:** Eine **dritte Ausprägung**, die der Eintrag bisher nicht kennt. Die ersten beiden waren
ein Prüfer ohne **Gegenstand** (leere Prüfmenge) und einer ohne **Aufruf** (advisory, lief nie).
Diese ist ein ganzes Modul-**Feld ohne Konfiguration**: `structure.table.column` mit
`cell-max-chars`/`cell-min-chars` liegt seit dem Pin `v0.74.1` im Werkzeug, ist im Startgerüst
dokumentiert (`d-check --print-config`) und wurde nie eingeschaltet.

**Gemessen:** Probeweise eingesetzt meldet die Konfiguration **8** überlange Zellen — sieben in
`harness/README.md` §Sensors (bis **356** Zeichen), eine in `harness/conventions.md`
§Aktive Adaptionen (**333**).

**Warum das schwerer wiegt als die zwei bekannten Formen:**
[slice-177](../../../../done/welle-15/slice-177-sensors-struktur-zwei-tabellen.md) hat genau diese
Frage — welche Zelle ist zum Absatz geworden? — **von Hand** an 16 Zellen beantwortet und den
Schnitt selbst begründet. Der unabhängige Review nannte die Begründung eine Selbstrechtfertigung
statt einer Messung (F-2) und ließ den Slice die Arbeit verdreifachen. Der Prüfer, der die Frage
mechanisch beantwortet, lag die ganze Zeit im gepinnten Werkzeug.

**Die Lehre ist nicht „mehr konfigurieren".** Sie ist: Wer eine Frage von Hand beantwortet, die
nach einem Sensor klingt, sieht im Werkzeug nach, **bevor** er zählt. Bei einem digest-gepinnten
Fremdwerkzeug heißt das `--print-config` lesen — es dokumentiert jedes Feld, auch die nicht
genutzten.

**Ausgang:** trägt [slice-181](../../../../done/wellenlos/slice-181-tabellenzellen-gewaechtert.md).
