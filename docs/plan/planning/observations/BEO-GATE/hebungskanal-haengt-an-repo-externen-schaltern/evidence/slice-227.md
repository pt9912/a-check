**Vorgang:** slice-226

**Fund (2026-10-10):** Der `docker`-Eintrag, den [ADR-0038](../../../../../adr/0038-dependabot-als-hebungskanal.md) ausgeschlossen hatte, wäre ohne Umbau
**gültig und wirkungslos** gewesen: Dependabots `FROM`-Parser verwirft
`golang:${GO_VERSION}@sha256:…` ohne Meldung (kein wörtlicher Tag, keine Abhängigkeit). Die Wirkung des
Kanals hängt damit an einer Eigenschaft von Dependabots Parser, die kein Gate des Repos sieht —
dieselbe Form wie die zwei Schalter außerhalb des Repos. Gemessen mit der Regex aus der Quelle
(dependabot-core, Stand `dc25739e`); verkörpert ist nur die Form der Zeilen
(`make version-coherence`, Regel 3). Das Fortschreiben des Index-Digests und die PR-Menge bleiben
bis zum ersten echten PR ungemessen.
