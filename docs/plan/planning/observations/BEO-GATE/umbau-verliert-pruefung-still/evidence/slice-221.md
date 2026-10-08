**Vorgang:** slice-221

**Fund (Review F-1, F-2):** Der Umbau von `tools/image-scan.sh` stellte den Entscheidungslauf von
Trivys Template auf JSON um. Die Anker-Probe des Selbsttests („Marker nur am Zeilenende") ging
dabei still verloren — die Ersatz-Probe hielt die Eigenschaft über das Escaping, nicht über den
Anker; mit entferntem Anker blieb `--selftest` grün. Zugleich wich der Lauf von den Fixtures in
[ADR-0037](../../../../../../../docs/plan/adr/0037-cve-scan-gegen-das-publizierte-image.md) §Fitness Function ab. Behoben durch Zurücksetzen auf die Fassung vor dem Umbau und
Ergänzen statt Ersetzen; Mutation „Anker entfernt" ist wieder rot.
