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

type yamlConfig struct {
	Version         int                 `yaml:"version"`
	Languages       map[string][]string `yaml:"languages"`
	Layers          map[string][]string `yaml:"layers"`
	Edges           []yamlEdge          `yaml:"edges"`
	AdapterSink     string              `yaml:"adapter_sink"`
	Tech            []yamlTech          `yaml:"tech"`
	CompositionRoot []string            `yaml:"composition_root"`
	Allow           []yamlEdge          `yaml:"allow"`
	Markers         *yamlMarkers        `yaml:"markers"`
	Forbidden       map[string][]string `yaml:"forbidden_constructs"`
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
		m.Layers = append(m.Layers, core.Layer{Name: name, Globs: yc.Layers[name]})
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

func sortedKeys(m map[string][]string) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}
