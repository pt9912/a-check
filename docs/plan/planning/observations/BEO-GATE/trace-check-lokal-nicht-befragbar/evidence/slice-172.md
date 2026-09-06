**Vorgang:** slice-172
**Fund:** `make trace-check` bricht unverändert mit
`d-check: error: Range-Basis-Vorfahren nicht lesbar: object not found` ab (Exit 2), gemessen am
2026-09-06 im selben Klon. Zweites Auftreten in einem anderen Vorgang — die Ursachen-Frage aus
`state.md` (liest `d-check` nur kanonisch benannte Packs, oder ist die Ablage dieses Klons
unüblich?) ist damit weiterhin unbeantwortet, aber nicht mehr einmalig.

**Neu an diesem Auftreten:** gefunden hat es der **unabhängige Review**, nicht der Implementer —
der Slice selbst hat `gates`, `verify` und `ci` gefahren, in denen `trace-check` nicht hängt. Das
bestätigt die Kern-Aussage des Eintrags empirisch: der Ausfall ist nicht still, aber er wird nur
gesehen, wenn jemand das Target eigens aufruft. Vier Commits dieses Slice sind lokal nie gegen die
Traceability-Regel geprüft worden; die CI holt das gegen einen frischen Klon nach.
