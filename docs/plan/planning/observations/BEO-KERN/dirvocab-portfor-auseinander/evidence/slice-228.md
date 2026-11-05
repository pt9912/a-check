**Vorgang:** slice-228
**Fund:** `isPortDirection` (`internal/hexagon/core/rules.go`) führt `inbound`/`outbound` als drittes Literal neben `portFor` (Kern) und `dirVocab` (Config-Adapter). Der Review-Fund F-1 entstand genau dort: die Ursachen-Klassifikation kannte zunächst keine Richtungs-Vokabel und hielt jedes Segment für eine. Kein Test hält die drei Stellen gegeneinander.
