// Package config is the configuration adapter (ARC-004): it strictly decodes
// `.a-check.yml` into the core model (SPEC-CONF-001). Unknown keys and type
// errors are fail-closed (the caller maps the error to exit code 2).
package config

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/pt9912/a-check/internal/hexagon/core"
	"github.com/pt9912/a-check/internal/hexagon/port"
	"gopkg.in/yaml.v3"
)

type yamlEdge struct {
	From   string `yaml:"from"`
	To     string `yaml:"to"`
	Reason string `yaml:"reason"`
}

type yamlTech struct {
	Pattern         string    `yaml:"pattern"`
	Adapter         yaml.Node `yaml:"adapter"`          // scalar OR list (AC-FA-RULE-003 0.14.0)
	Match           string    `yaml:"match"`            // "substring" (default) or "regex" (ADR-0015)
	CompositionRoot string    `yaml:"composition_root"` // "allow" (default) or "forbid" (AC-FA-RULE-003 0.14.0)
}

// yamlConstruct is one `constructs` entry (AC-FA-RULE-011, ADR-0027): the same
// shape as yamlTech — `adapter` is the ZONE (scalar or list) — but the pattern
// matches raw source text instead of an import symbol.
type yamlConstruct struct {
	Pattern         string    `yaml:"pattern"`
	Adapter         yaml.Node `yaml:"adapter"`          // scalar OR list: the allowed zone(s)
	Match           string    `yaml:"match"`            // "substring" (default) or "regex"
	CompositionRoot string    `yaml:"composition_root"` // "allow" (default) or "forbid"
}

type yamlMarkers struct {
	IgnoreSymbols []string `yaml:"ignore_symbols"`
}

// yamlResolution is one language's import-resolution config (ADR-0016).
type yamlResolution struct {
	Mode        string   `yaml:"mode"`
	Roots       []string `yaml:"roots"`
	PackageBase string   `yaml:"package_base"`
}

// yamlLayer is the object form of a layers entry (`{globs, role, direction}`,
// AC-FA-RULE-006/008); the glob-list short form is decoded separately (see
// decodeLayer). direction is optional and orthogonal to role.
type yamlLayer struct {
	Globs     []string `yaml:"globs"`
	Role      string   `yaml:"role"`
	Direction string   `yaml:"direction"`
}

type yamlConfig struct {
	Version         int                       `yaml:"version"`
	Languages       map[string][]string       `yaml:"languages"`
	Layers          yaml.Node                 `yaml:"layers"` // rohe Mapping-Node: erhält die Deklarationsreihenfolge (ADR-0013 Tie-Break, slice-038)
	Edges           []yamlEdge                `yaml:"edges"`
	AdapterSink     string                    `yaml:"adapter_sink"`
	Tech            []yamlTech                `yaml:"tech"`
	Constructs      []yamlConstruct           `yaml:"constructs"` // Roh-Text-Monopol (ADR-0027)
	CompositionRoot []string                  `yaml:"composition_root"`
	Allow           []yamlEdge                `yaml:"allow"`
	Markers         *yamlMarkers              `yaml:"markers"`
	Forbidden       map[string][]string       `yaml:"forbidden_constructs"`
	Resolution      map[string]yamlResolution `yaml:"resolution"`
	Exclude         []string                  `yaml:"exclude"` // Scan-Scope (ADR-0018)
}

// Adapter implements port.ConfigPort.
type Adapter struct{}

// New returns a configuration adapter.
func New() port.ConfigPort { return Adapter{} }

// Load reads and strictly decodes the config at path.
func (Adapter) Load(path string) (core.Model, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return core.Model{}, err
	}
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true) // fail-closed on unknown keys (SPEC-CONF-001)
	var yc yamlConfig
	if err := dec.Decode(&yc); err != nil {
		return core.Model{}, fmt.Errorf("%s: %w", path, err)
	}
	if yc.Version != 1 {
		return core.Model{}, fmt.Errorf("%s: 'version: 1' erforderlich", path)
	}
	if len(yc.Languages) == 0 || yc.Layers.Kind == 0 || len(yc.Edges) == 0 {
		return core.Model{}, fmt.Errorf("%s: Pflichtblöcke 'languages', 'layers', 'edges' erforderlich", path)
	}

	m := core.Model{
		Languages:       yc.Languages,
		AdapterSink:     yc.AdapterSink,
		CompositionRoot: yc.CompositionRoot,
		Forbidden:       yc.Forbidden,
	}
	layers, lerr := decodeLayers(yc.Layers, path)
	if lerr != nil {
		return core.Model{}, lerr
	}
	m.Layers = layers
	for _, e := range yc.Edges {
		m.Edges = append(m.Edges, core.Edge{From: e.From, To: e.To})
	}
	for _, e := range yc.Allow {
		m.Allow = append(m.Allow, core.Edge{From: e.From, To: e.To})
	}
	if berr := decodeOptionalBlocks(&m, yc, path); berr != nil {
		return core.Model{}, berr
	}
	return m, nil
}

// decodeOptionalBlocks decodes and validates everything past the mandatory
// blocks. It lives apart from Load for one reason: Load's branch count sits at
// the lint profile's ceiling (ADR-0005), and each new fail-closed validation adds
// one. Splitting on the mandatory/optional seam keeps the next one from forcing a
// suppression — which the Hard Rules forbid anyway (AGENTS §3.2).
func decodeOptionalBlocks(m *core.Model, yc yamlConfig, path string) error {
	if ferr := validateForbidden(yc.Forbidden, m.Layers, path); ferr != nil {
		return ferr
	}
	techs, terr := decodeTechs(yc.Tech, path)
	if terr != nil {
		return terr
	}
	m.Techs = techs
	constructs, cerr := decodeConstructs(yc.Constructs, path)
	if cerr != nil {
		return cerr
	}
	m.Constructs = constructs
	if eerr := validateExclude(yc.Exclude, path); eerr != nil {
		return eerr
	}
	m.Exclude = yc.Exclude
	if yc.Markers != nil {
		m.IgnoreSymbols = yc.Markers.IgnoreSymbols
	}
	res, rerr := decodeResolution(yc.Resolution, yc.Languages, path)
	if rerr != nil {
		return rerr
	}
	m.Resolution = res
	return nil
}

// forbiddenHint names the sibling block in every message. constructs is the
// COUNTERPART, not a replacement (ADR-0033): it scopes by zone and scan-wide,
// forbidden_constructs by layer. Saying "use constructs instead" would send a
// consumer who wants a per-layer blacklist down a dead end, because expressing
// that with constructs means enumerating every OTHER zone and maintaining that
// list forever.
const forbiddenHint = "; das zonen-gebundene Gegenstueck ist constructs (nicht deckungsgleich: es " +
	"erlaubt ein Muster NUR in seinen Zonen, scan-weit)"

// validateForbidden fails closed on every forbidden_constructs entry that can
// never report (ADR-0033, SPEC-CONF-001). Until now the block was passed through
// unchecked and had four silent exits, all ending in exit 0 — the same
// false-green class that languages, tech.adapter and constructs have long
// rejected loudly.
//
// The role binding is NOT an oversight to be widened here: AC-FA-RULE-004 is
// literally "Port-Disziplin" and names construct-freedom as a PORT property, so
// evaluating only role port FULFILLS the contract. Widening it would be a
// Lastenheft change (MR-001), not a validation.
//
// Layer keys are checked in sorted order so a config with two errors always
// reports the same one first (SPEC-DET-001).
func validateForbidden(fc map[string][]string, layers []core.Layer, path string) error {
	names := make([]string, 0, len(fc))
	for n := range fc {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, name := range names {
		l, ok := layerNamed(name, layers)
		if !ok {
			return fmt.Errorf("%s: forbidden_constructs[%q]: unbekannte Schicht — kein Eintrag in layers%s", path, name, forbiddenHint)
		}
		if role := core.EffectiveRole(l); role != "port" {
			return fmt.Errorf("%s: forbidden_constructs[%q]: Schicht hat die Rolle %q, ausgewertet wird nur %q (Port-Disziplin) — der Eintrag wuerde nie melden%s", path, name, roleLabel(role), "port", forbiddenHint)
		}
		if len(fc[name]) == 0 {
			return fmt.Errorf("%s: forbidden_constructs[%q]: leere Musterliste unzulaessig (sie wuerde nie melden)", path, name)
		}
		for _, p := range fc[name] {
			if p == "" {
				return fmt.Errorf("%s: forbidden_constructs[%q]: leeres Muster unzulaessig (es wuerde nie melden)", path, name)
			}
		}
	}
	return nil
}

// layerNamed finds a decoded layer by name.
func layerNamed(name string, layers []core.Layer) (core.Layer, bool) {
	for _, l := range layers {
		if l.Name == name {
			return l, true
		}
	}
	return core.Layer{}, false
}

// roleLabel spells a layer role for a message. A layer that resolves to NO role
// (neither explicit nor inferable from its name) would render as an empty pair of
// quotes, which reads like a bug rather than the actual situation.
func roleLabel(role string) string {
	if role == "" {
		return "keine (weder role: gesetzt noch aus dem Namen ableitbar)"
	}
	return role
}

// decodeTechs builds the tech list: adapter scalar/list decode plus NewTech
// validation (match, composition_root, empty list — AC-FA-RULE-003 0.14.0).
func decodeTechs(entries []yamlTech, path string) ([]core.Tech, error) {
	var out []core.Tech
	for _, t := range entries {
		adapters, aerr := decodeZoneList(t.Adapter, "tech-Muster", t.Pattern, path)
		if aerr != nil {
			return nil, aerr
		}
		tech, terr := core.NewTech(t.Pattern, adapters, t.Match, t.CompositionRoot)
		if terr != nil {
			return nil, fmt.Errorf("%s: %w", path, terr)
		}
		out = append(out, tech)
	}
	return out, nil
}

// decodeConstructs builds the raw-text construct list (AC-FA-RULE-011): same
// zone decode as tech, then core.NewConstruct validates pattern/zone/match/
// composition_root fail-closed (exit 2). A missing block yields nil — no rule.
func decodeConstructs(entries []yamlConstruct, path string) ([]core.Construct, error) {
	var out []core.Construct
	for _, c := range entries {
		zones, zerr := decodeZoneList(c.Adapter, "constructs-Muster", c.Pattern, path)
		if zerr != nil {
			return nil, zerr
		}
		con, cerr := core.NewConstruct(c.Pattern, zones, c.Match, c.CompositionRoot)
		if cerr != nil {
			return nil, fmt.Errorf("%s: %w", path, cerr)
		}
		out = append(out, con)
	}
	return out, nil
}

// decodeZoneList reads a tech/constructs entry's `adapter` as a scalar or a
// string list (AC-FA-RULE-003 0.14.0, AC-FA-RULE-011). A non-empty scalar becomes
// a one-element list — byte-identical to the pre-list behavior. An ABSENT or
// EMPTY adapter (also `null`) fails closed (exit 2): in the pre-0.14.0 code it
// was a silent never-leak dead entry (`strings.Contains(path, "")` is always
// true, so the pattern never reported) — a false-green trap, not a behavior worth
// preserving (Review-R1 B1; same ethos line as the empty resolution root).
// A YAML alias is dereferenced first (yaml.Node fields keep Kind==AliasNode).
// kind names the block in the message ("tech-Muster"/"constructs-Muster"); both
// share the mechanic, so they share the decode.
func decodeZoneList(node yaml.Node, kind, pattern, path string) ([]string, error) {
	if node.Kind == yaml.AliasNode && node.Alias != nil {
		node = *node.Alias
	}
	switch node.Kind {
	case 0:
		return nil, fmt.Errorf("%s: %s %q: adapter fehlt (Pfad oder Pfad-Liste erforderlich)", path, kind, pattern)
	case yaml.ScalarNode:
		var s string
		if err := node.Decode(&s); err != nil {
			return nil, fmt.Errorf("%s: %s %q: adapter: %w", path, kind, pattern, err)
		}
		if s == "" {
			return nil, fmt.Errorf("%s: %s %q: leerer adapter unzulässig (war ein stiller Never-Leak-Eintrag)", path, kind, pattern)
		}
		return []string{s}, nil
	case yaml.SequenceNode:
		var list []string
		if err := node.Decode(&list); err != nil {
			return nil, fmt.Errorf("%s: %s %q: adapter: %w", path, kind, pattern, err)
		}
		if len(list) == 0 {
			return nil, fmt.Errorf("%s: %s %q: leere adapter-Liste unzulässig", path, kind, pattern)
		}
		for _, a := range list {
			if a == "" {
				return nil, fmt.Errorf("%s: %s %q: leerer adapter-Listen-Eintrag unzulässig", path, kind, pattern)
			}
		}
		return list, nil
	default:
		return nil, fmt.Errorf("%s: %s %q: adapter muss Pfad oder Pfad-Liste sein", path, kind, pattern)
	}
}

// validateExclude rejects empty exclude globs (the invalid case of the
// otherwise total glob engine — a silent match-nothing entry, ADR-0018).
func validateExclude(globs []string, path string) error {
	for _, g := range globs {
		if g == "" {
			return fmt.Errorf("%s: exclude: leerer Glob unzulässig", path)
		}
	}
	return nil
}

// decodeResolution maps the resolution block to the core model (ADR-0016),
// validating each entry (see resolutionEntry). A blank block yields nil.
func decodeResolution(res map[string]yamlResolution, langs map[string][]string, path string) (map[string]core.ResolutionConfig, error) {
	if len(res) == 0 {
		return nil, nil
	}
	out := make(map[string]core.ResolutionConfig, len(res))
	for _, lang := range sortedKeys(res) {
		cfg, err := resolutionEntry(lang, res[lang], langs, path)
		if err != nil {
			return nil, err
		}
		out[lang] = cfg
	}
	return out, nil
}

// resolutionEntry validates + normalizes one resolution entry (ADR-0016): the
// key must be a declared language (no silent no-op for typos); path duldet kein
// roots/package_base; fixed-root braucht mindestens eines; relative (ADR-0017)
// duldet ebenfalls kein roots/package_base (fail-closed statt still ignoriert);
// namespace ist reserviert (→ exit 2). A blank mode defaults to path.
func resolutionEntry(lang string, r yamlResolution, langs map[string][]string, path string) (core.ResolutionConfig, error) {
	if _, ok := langs[lang]; !ok {
		return core.ResolutionConfig{}, fmt.Errorf("%s: resolution[%q]: keine unter languages deklarierte Sprache", path, lang)
	}
	mode := r.Mode
	if mode == "" {
		mode = "path"
	}
	switch mode {
	case "path":
		if len(r.Roots) > 0 || r.PackageBase != "" {
			return core.ResolutionConfig{}, fmt.Errorf("%s: resolution[%q]: mode path duldet kein roots/package_base", path, lang)
		}
	case "fixed-root":
		if len(r.Roots) == 0 && r.PackageBase == "" {
			return core.ResolutionConfig{}, fmt.Errorf("%s: resolution[%q]: mode fixed-root braucht roots und/oder package_base", path, lang)
		}
	case "relative":
		if len(r.Roots) > 0 || r.PackageBase != "" {
			return core.ResolutionConfig{}, fmt.Errorf("%s: resolution[%q]: mode relative duldet kein roots/package_base", path, lang)
		}
	case "namespace":
		return core.ResolutionConfig{}, fmt.Errorf("%s: resolution[%q].mode %q ist reserviert (Folge-ADR, noch nicht implementiert)", path, lang, mode)
	default:
		return core.ResolutionConfig{}, fmt.Errorf("%s: resolution[%q].mode %q ungültig (path|fixed-root|relative)", path, lang, mode)
	}
	roots, err := cleanRoots(r.Roots, lang, path)
	if err != nil {
		return core.ResolutionConfig{}, err
	}
	return core.ResolutionConfig{Mode: mode, Roots: roots, PackageBase: r.PackageBase}, nil
}

// cleanRoots trims trailing slashes and rejects empty roots (config footgun,
// sonst „src/" → „src//x" bzw. „" → „/x", die still nichts matchen).
func cleanRoots(roots []string, lang, path string) ([]string, error) {
	if len(roots) == 0 {
		return nil, nil
	}
	out := make([]string, 0, len(roots))
	for _, r := range roots {
		t := strings.TrimRight(r, "/")
		if t == "" {
			return nil, fmt.Errorf("%s: resolution[%q]: leerer root", path, lang)
		}
		out = append(out, t)
	}
	return out, nil
}

// decodeLayer reads a layers entry: a glob list (`name: [globs]`) or an object
// (`{globs, role, direction}`, AC-FA-RULE-006/008). Returns globs, role and the
// (optional) direction.
// decodeLayers builds the ordered layer list from the raw `layers` mapping node,
// preserving DECLARATION ORDER (ADR-0013 tie-break: "bei Gleichstand die zuerst
// deklarierte"; slice-038). The prior map[string]yaml.Node decode sorted keys
// alphabetically, silently violating that contract. Decoding into a raw yaml.Node
// keeps document order in Content — but forfeits two fail-closed checks the Go-map
// decode gave for free, so both are reconstructed here:
//   - Guard 1 (Kind): a non-mapping `layers:` (sequence/scalar) would misparse in
//     the pairwise loop; require MappingNode (SPEC-CONF-001 strict decode).
//   - Guard 2 (duplicates): yaml.v3 only flags duplicate keys when decoding into a
//     map/struct (uniqueKeys); the Node path bypasses it, so a repeated layer name
//     would yield two silent core.Layer entries — reject it (AC-QA-02 fail-closed).
func decodeLayers(node yaml.Node, path string) ([]core.Layer, error) {
	if node.Kind != yaml.MappingNode {
		return nil, fmt.Errorf("%s: 'layers' erwartet eine Mapping (Schicht → Globs/Objekt)", path)
	}
	if len(node.Content) == 0 {
		return nil, fmt.Errorf("%s: Pflichtblock 'layers' ist leer", path)
	}
	var layers []core.Layer
	seen := make(map[string]bool, len(node.Content)/2)
	for i := 0; i+1 < len(node.Content); i += 2 {
		name := node.Content[i].Value
		if seen[name] {
			return nil, fmt.Errorf("%s: Schicht %q mehrfach deklariert", path, name)
		}
		seen[name] = true
		globs, role, direction, lerr := decodeLayer(*node.Content[i+1], name, path)
		if lerr != nil {
			return nil, lerr
		}
		layers = append(layers, core.Layer{Name: name, Globs: globs, Role: role, Direction: direction})
	}
	return layers, nil
}

func decodeLayer(node yaml.Node, name, path string) ([]string, string, string, error) {
	switch node.Kind {
	case yaml.SequenceNode:
		var globs []string
		if err := node.Decode(&globs); err != nil {
			return nil, "", "", fmt.Errorf("%s: Schicht %q: %w", path, name, err)
		}
		return globs, "", "", nil
	case yaml.MappingNode:
		return decodeLayerObject(node, name, path)
	default:
		return nil, "", "", fmt.Errorf("%s: Schicht %q: erwarte Glob-Liste oder {globs, role, direction}", path, name)
	}
}

// decodeLayerObject decodes the strict object form {globs, role, direction}. It
// is strict by hand — KnownFields(true) on the decoder is NOT inherited by
// yaml.Node.Decode, so unknown keys and invalid enums are rejected explicitly
// (SPEC-CONF-001).
func decodeLayerObject(node yaml.Node, name, path string) ([]string, string, string, error) {
	for i := 0; i+1 < len(node.Content); i += 2 {
		if k := node.Content[i].Value; !knownLayerKey(k) {
			return nil, "", "", fmt.Errorf("%s: Schicht %q: unbekannter Schlüssel %q", path, name, k)
		}
	}
	var yl yamlLayer
	if err := node.Decode(&yl); err != nil {
		return nil, "", "", fmt.Errorf("%s: Schicht %q: %w", path, name, err)
	}
	if !validRole(yl.Role) {
		return nil, "", "", fmt.Errorf("%s: Schicht %q: ungültige role %q (domain|app|port|adapter)", path, name, yl.Role)
	}
	// Die Richtung wird gegen das Vokabular der EFFEKTIVEN Rolle geprueft:
	// fehlt `role`, gilt die Namens-Inferenz (AC-FA-RULE-006). Sonst kaeme
	// `ports: {globs: [...], direction: driving}` ohne explizite Rolle durch —
	// genau die Schreibweise, die ADR-0036 abschafft.
	effRole := yl.Role
	if effRole == "" {
		effRole = core.InferRole(name)
	}
	if !validDirectionFor(effRole, yl.Direction) {
		return nil, "", "", fmt.Errorf("%s: Schicht %q: ungültige direction %q an role %q (%s)",
			path, name, yl.Direction, effRole, dirVocabList(effRole))
	}
	return yl.Globs, yl.Role, yl.Direction, nil
}

func knownLayerKey(k string) bool { return k == "globs" || k == "role" || k == "direction" }

// dirVocab nennt die fuer eine Rolle gueltige Richtungs-Menge (AC-FA-RULE-008,
// ADR-0036). DIE EINE TABELLE der Config-Seite: sie speist die Validierung UND
// die Fehlermeldung. Als Funktion statt als Paket-Variable, weil das Lint-Profil
// globale Zustandstraeger verbietet (ADR-0005) — und eine Suppression waere
// nach AGENTS §3.2 ohnehin keine Option.
//
// Ein Port TREIBT nichts, er wird benutzt: seine Richtung sagt, wohin die
// Schnittstelle zeigt (inbound|outbound). Ein Adapter ist nicht EINGEHEND,
// er treibt oder wird getrieben (driving|driven).
func dirVocab(role string) []string {
	switch role {
	case "port":
		return []string{"inbound", "outbound"}
	case "adapter":
		return []string{"driving", "driven"}
	default:
		return nil
	}
}

// validDirectionFor prueft die Richtung gegen das Vokabular IHRER Rolle. Die
// leere Richtung ist immer gueltig (die Dimension ist opt-in). Eine Richtung an
// einer Rolle ohne Vokabular ist ein Fehler: sie waere wirkungslos und sagte
// etwas, das keine Regel liest (ADR-0036).
func validDirectionFor(role, d string) bool {
	if d == "" {
		return true
	}
	for _, v := range dirVocab(role) {
		if d == v {
			return true
		}
	}
	return false
}

// dirVocabList rendert die gueltige Menge fuer die Fehlermeldung. Sie MUSS
// rollen-spezifisch sein — eine feste Menge zu nennen schickte den Leser auf
// die falsche Haelfte (ADR-0036 §Konsequenzen).
func dirVocabList(role string) string {
	v := dirVocab(role)
	if v == nil {
		return "keine — nur an role: port oder role: adapter"
	}
	return strings.Join(v, "|")
}

func validRole(r string) bool {
	return r == "" || r == "domain" || r == "app" || r == "port" || r == "adapter"
}


func sortedKeys[V any](m map[string]V) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}
