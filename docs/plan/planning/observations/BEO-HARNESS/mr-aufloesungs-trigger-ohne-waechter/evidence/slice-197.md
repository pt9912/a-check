**Vorgang:** slice-197

**Fund:** Der unabhängige Review hat gemeldet, dass der Sichtungs-Schritt dieses Slice „kein
Treffer" notierte, obwohl **dieser Eintrag** genau das gewählte Instrument betrifft. Der Slice
führt einen `MR`-Eintrag ein ([`MR-024`](../../../../../../../harness/conventions.md#mr-024)) und stützt ihn auf dessen **Auflösungs-Trigger** — der
Gegenstand dieses Eintrags ist, dass dieses Pflichtfeld keinen Wächter hat.

**Zweites Auftreten derselben Klasse.** Das erste ist `slice-170`; dort ist der Vorgang zuerst
aufgefallen. Hier ist es die **Anwendung**: Wer ein Instrument wählt, dessen Wächter fehlt, muss
das benennen — sonst behauptet er eine Auflösung, die niemand einlöst.

**Wie es auffiel:** im unabhängigen Review, nicht im Lauf. Der Plan hat den Eintrag beim Sichten
gelesen und nicht als Treffer erkannt; die Konsequenz steht jetzt in §9 des Plans und in der
Closure-Notiz.

**Und was das für [`MR-024`](../../../../../../../harness/conventions.md#mr-024) heißt:** Dessen Trigger („das nächste Release") ist **prüfbar ohne
Wächter**, weil er sich an einem Datum im Release-Register ablesen lässt — aber niemand prüft ihn.
Das bleibt die offene Hälfte dieses Eintrags.
