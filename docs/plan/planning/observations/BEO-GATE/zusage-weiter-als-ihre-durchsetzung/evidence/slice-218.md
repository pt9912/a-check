**Vorgang:** slice-218

**Fund (Review F-1, F-8; Delta D-1):** Der Kopfkommentar von `release.yml` sagte „jeder Schritt
läuft über make oder die Docker-CLI" — das GitHub-Release läuft über `gh`, Hash und Vergleich über
Bash. Die Hub-Seite sagte ohne Einschränkung, der GHCR-Digest löse auf Docker Hub auf — für
ältere, neu hochgeladene Tags stimmt das nicht (gemessen: `v0.22.0` trägt dort `9587aa5a…`,
auf GHCR `12e961f7…`). Ein Schritt „Spiegel-Fehlschlag melden" lief bei jedem Fehlschlag.
Drei Funde, **ein** Vorgang. Behoben.
