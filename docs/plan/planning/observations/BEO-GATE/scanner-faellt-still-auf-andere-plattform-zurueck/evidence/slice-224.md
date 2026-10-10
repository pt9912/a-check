**Vorgang:** slice-224

**Fund (2026-10-10):** `trivy image --input <OCI-Layout> --platform linux/arm64` gegen das lokal
gebaute Multi-Arch-Archiv meldete in `Metadata.ImageConfig.architecture` `amd64` — dieselbe stille
Rückfall-Form auf einem zweiten Eingangsweg. Gemessen wurde danach je Plattform gegen ein Layout,
dessen Index nur das eine Manifest trägt; dann stand `arm64` im Bericht.
