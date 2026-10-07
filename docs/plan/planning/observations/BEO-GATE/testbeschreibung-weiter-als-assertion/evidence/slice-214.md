**Vorgang:** slice-214

**Fund (Review F-2, F-7; Delta-Review D-3):** `TestJSONUnsplittable` nannte sich „alle
fail-closed-Fälle der Spezifikation" und „Exit 2", prüfte drei Fälle nicht und nur den Fehler,
nicht den Exit-Code; der Nachfolge-Kommentar von `TestJSONValidity` sagte „jede Form, die die
Quelle verbietet" — drei Mutationen (`1e+-2`, `--1`, fehlendes `\b`) überstanden die Suite. Drei
Funde, **ein** Vorgang. Behoben: Kommentare sagen nur zu, was die Fälle prüfen; neue Fälle, je
mit Mutations-Gegenprobe rot.
