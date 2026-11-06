package config

import (
	"strings"
	"testing"
)

const ackHead = "version: 1\nlanguages:\n  go: [\"**/*.go\"]\nlayers:\n  app:\n    globs: [\"src/app/**\"]\n    role: app\n"
const ackTail = "edges:\n  - {from: app, to: ports}\n"

func ackCfg(layer string) string { return ackHead + layer + ackTail }

func TestAcknowledgeDecodes(t *testing.T) { // AC-FA-CONF-002 happy
	m, err := New().Load(write(t, ackCfg("  ports:\n    globs: [\"src/app/ports/**\"]\n    role: port\n    acknowledge:\n      port-scope: \"app-weite Ports\"\n")))
	if err != nil {
		t.Fatalf("gueltige Quittierung abgelehnt: %v", err)
	}
	if got := m.Layers[1].AckPortScope; got != "app-weite Ports" {
		t.Fatalf("Begruendung nicht dekodiert, got %q", got)
	}
	if m.Layers[0].AckPortScope != "" {
		t.Fatal("eine Schicht ohne acknowledge traegt keine Quittierung")
	}
}

func TestAcknowledgeOnInferredPortRole(t *testing.T) { // AC-FA-CONF-002: die effektive Rolle zaehlt, nicht nur das explizite role:
	_, err := New().Load(write(t, ackCfg("  ports:\n    globs: [\"src/app/ports/**\"]\n    acknowledge: {port-scope: \"x\"}\n")))
	if err != nil {
		t.Fatalf("Namens-Inferenz macht ports zur Port-Schicht, got %v", err)
	}
}

func TestAcknowledgeViolationsFailClosed(t *testing.T) { // AC-FA-CONF-002 negative: jeder Fall Exit-2-Fehler mit Schicht und Fehler
	cases := map[string]struct{ layer, want string }{
		"Rolle nicht port":      {"  extra:\n    globs: [\"src/x/**\"]\n    role: adapter\n    acknowledge: {port-scope: \"x\"}\n", "nur an einer Schicht mit der Rolle port"},
		"keine Rolle ableitbar": {"  zzz:\n    globs: [\"src/x/**\"]\n    acknowledge: {port-scope: \"x\"}\n", "nur an einer Schicht mit der Rolle port"},
		"anderer Schluessel":    {"  ports:\n    globs: [\"src/app/ports/**\"]\n    role: port\n    acknowledge: {coverage: \"x\"}\n", "unbekannter Schlüssel"},
		"keine Map":             {"  ports:\n    globs: [\"src/app/ports/**\"]\n    role: port\n    acknowledge: \"port-scope\"\n", "nicht leere Map"},
		"Liste":                 {"  ports:\n    globs: [\"src/app/ports/**\"]\n    role: port\n    acknowledge: [port-scope]\n", "nicht leere Map"},
		"null":                  {"  ports:\n    globs: [\"src/app/ports/**\"]\n    role: port\n    acknowledge:\n", "nicht leere Map"},
		"leere Map":             {"  ports:\n    globs: [\"src/app/ports/**\"]\n    role: port\n    acknowledge: {}\n", "nicht leere Map"},
		"doppelter Schluessel":  {"  ports:\n    globs: [\"src/app/ports/**\"]\n    role: port\n    acknowledge:\n      port-scope: \"a\"\n      port-scope: \"b\"\n", ""},
		"Begruendung null":      {"  ports:\n    globs: [\"src/app/ports/**\"]\n    role: port\n    acknowledge: {port-scope: null}\n", "String"},
		"Begruendung Zahl":      {"  ports:\n    globs: [\"src/app/ports/**\"]\n    role: port\n    acknowledge: {port-scope: 5}\n", "String"},
		"Begruendung Liste":     {"  ports:\n    globs: [\"src/app/ports/**\"]\n    role: port\n    acknowledge: {port-scope: [a]}\n", "String"},
		"Begruendung leer":      {"  ports:\n    globs: [\"src/app/ports/**\"]\n    role: port\n    acknowledge: {port-scope: \"\"}\n", "leer"},
		"Begruendung Leerraum":  {"  ports:\n    globs: [\"src/app/ports/**\"]\n    role: port\n    acknowledge: {port-scope: \"  \\t \"}\n", "leer"},
	}
	for name, c := range cases {
		_, err := New().Load(write(t, ackCfg(c.layer)))
		if err == nil {
			t.Errorf("%s: erwartet Fehler (fail-closed)", name)
			continue
		}
		if !strings.Contains(err.Error(), "Schicht") || !strings.Contains(err.Error(), c.want) {
			t.Errorf("%s: Meldung nennt Schicht und Fehler (%q erwartet), got %v", name, c.want, err)
		}
	}
}
