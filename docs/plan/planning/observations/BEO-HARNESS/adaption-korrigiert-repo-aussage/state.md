**Stand:** geplant → [slice-190](../../../open/slice-190-id-schema-deklaration-ueberarbeiten.md) (3×)

**Der Trigger, den dieser Eintrag selbst nannte, ist eingetreten:** Mit
[`MR-023`](../../../../../../harness/conventions.md#mr-023) repariert zum **zweiten** Mal ein
Nachfolge-Eintrag dieselbe Zeile-Familie — die ID-Schema-Deklaration in
[`MR-000`](../../../../../../harness/conventions.md#mr-000) —, und beide Male steht
`Ersetzt-Baseline-Regel: —`. Weiterflicken erzeugt beim nächsten Mal den dritten Eintrag.

**Der Ausgang ist die Überarbeitung, nicht ein weiterer Eintrag:** Die Deklaration wird als Ganzes
abgelöst, generisch statt aufzählend, sodass [`MR-020`](../../../../../../harness/conventions.md#mr-020) und [`MR-023`](../../../../../../harness/conventions.md#mr-023) entbehrlich werden. Geplant als
[slice-190](../../../open/slice-190-id-schema-deklaration-ueberarbeiten.md); dort steht auch,
warum das nicht in einer Closure nebenbei geht — die abzulösenden Kennungen stehen repo-weit in
Commit-Messages.

**Kein Sensor**, und das ist benannt: Ob ein Eintrag eine Baseline-Regel oder eine eigene Aussage
korrigiert, ist ein Urteil über seinen Gegenstand ([`AGENTS.md`](../../../../../../AGENTS.md)
§3.7). Das Pflichtfeld `Ersetzt-Baseline-Regel` macht es **sichtbar** — das ist seine ganze
Leistung, und sie genügt, solange jemand hinsieht.
