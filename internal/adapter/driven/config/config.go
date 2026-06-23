// Package config is the configuration adapter (ARC-004): it strictly decodes
// `.a-check.yml` into the core model (SPEC-CONF-001). Unknown keys and type
// errors are fail-closed (the caller maps the error to exit code 2).
package config

import (
	"bytes"
	"fmt"
	"os"
	"sort"

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
	Pattern string `yaml:"pattern"`
	Adapter string `yaml:"adapter"`
}

type yamlMarkers struct {
	IgnoreSymbols []string `yaml:"ignore_symbols"`
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
	Version         int                  `yaml:"version"`
	Languages       map[string][]string  `yaml:"languages"`
	Layers          map[string]yaml.Node `yaml:"layers"`
	Edges           []yamlEdge           `yaml:"edges"`
	AdapterSink     string               `yaml:"adapter_sink"`
	Tech            []yamlTech           `yaml:"tech"`
	CompositionRoot []string             `yaml:"composition_root"`
	Allow           []yamlEdge           `yaml:"allow"`
	Markers         *yamlMarkers         `yaml:"markers"`
	Forbidden       map[string][]string  `yaml:"forbidden_constructs"`
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
	if len(yc.Languages) == 0 || len(yc.Layers) == 0 || len(yc.Edges) == 0 {
		return core.Model{}, fmt.Errorf("%s: Pflichtblöcke 'languages', 'layers', 'edges' erforderlich", path)
	}

	m := core.Model{
		Languages:       yc.Languages,
		AdapterSink:     yc.AdapterSink,
		CompositionRoot: yc.CompositionRoot,
		Forbidden:       yc.Forbidden,
	}
	for _, name := range sortedKeys(yc.Layers) {
		globs, role, direction, lerr := decodeLayer(yc.Layers[name], name, path)
		if lerr != nil {
			return core.Model{}, lerr
		}
		m.Layers = append(m.Layers, core.Layer{Name: name, Globs: globs, Role: role, Direction: direction})
	}
	for _, e := range yc.Edges {
		m.Edges = append(m.Edges, core.Edge{From: e.From, To: e.To})
	}
	for _, e := range yc.Allow {
		m.Allow = append(m.Allow, core.Edge{From: e.From, To: e.To})
	}
	for _, t := range yc.Tech {
		m.Techs = append(m.Techs, core.Tech{Pattern: t.Pattern, Adapter: t.Adapter})
	}
	if yc.Markers != nil {
		m.IgnoreSymbols = yc.Markers.IgnoreSymbols
	}
	return m, nil
}

// decodeLayer reads a layers entry: a glob list (`name: [globs]`) or an object
// (`{globs, role, direction}`, AC-FA-RULE-006/008). Returns globs, role and the
// (optional) direction.
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
	if !validDirection(yl.Direction) {
		return nil, "", "", fmt.Errorf("%s: Schicht %q: ungültige direction %q (driving|driven)", path, name, yl.Direction)
	}
	return yl.Globs, yl.Role, yl.Direction, nil
}

func knownLayerKey(k string) bool { return k == "globs" || k == "role" || k == "direction" }

func validRole(r string) bool {
	return r == "" || r == "domain" || r == "app" || r == "port" || r == "adapter"
}

func validDirection(d string) bool { return d == "" || d == "driving" || d == "driven" }

func sortedKeys[V any](m map[string]V) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}
