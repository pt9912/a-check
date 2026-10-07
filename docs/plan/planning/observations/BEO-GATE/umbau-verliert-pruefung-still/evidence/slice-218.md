**Vorgang:** slice-218

**Fund (Review F-4):** Der Umbau von `release.yml` auf Multi-Arch ersetzte den Schritt „Verify OCI
labels"; die Versions-Prüfung zog in den Image-Test um, die Prüfung von source, description,
licenses, title und vendor auf „nicht leer" fiel ersatzlos weg — ohne Nennung in Plan oder ADR.
Behoben in `tools/multiarch-check.sh`, Gegenprobe `vendor` leer → rot.
