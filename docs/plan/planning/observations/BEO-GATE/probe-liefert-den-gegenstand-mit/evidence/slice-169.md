**Vorgang:** slice-169

**Fund:** Die Mutations-Probe des Kalibrierungs-Selbsttests mutierte das **Muster** in
`.d-check.yml` und zeigte, dass der Lauf danach rot wird. Die Zusage des Sensors galt aber der
**Kandidatenmenge des `reviews`-Moduls** — und gegen die lief keine Probe. Der Sensor zählte eine
Obermenge (5 statt 2) und blieb grün, als beide echten Kandidaten entwertet wurden.

**Wie es auffiel:** unabhängiger Review, F-1.

**Nachgetragen mit slice-180**, weil das Muster dort zum zweiten Mal auftrat und beim Abschluss
von slice-169 noch keinen Eintrag hatte. Der Vorgang ist abgeschlossen und liegt in `done/`; der
Beleg ist damit form- und lagerichtig, nur später geschrieben.
