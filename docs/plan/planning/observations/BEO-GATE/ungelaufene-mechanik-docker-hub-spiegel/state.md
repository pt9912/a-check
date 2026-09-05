**Stand:** offen

**behoben und belegt** (2026-08-30): Muster um `tr -d` ergänzt, Config-Digest auf beiden Registries identisch (`sha256:e4f357f0…`), Manifest-Digest verschieden — die ADR-Entscheidung ist damit live bestätigt. Der Riegel hätte den Fehler fail-closed gefangen, aber den Spiegel blockiert. **Lehre:** ein Vergleich zweier **leerer** Werte meldet Gleichheit — jede Gleichheitsprüfung braucht den Riegel auf Nicht-Leere daneben
