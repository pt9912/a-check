package cli_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/pt9912/a-check/internal/cli"
)

// ackTwoLayers: two port layers that both trigger the port-scope notice
// (no direction declared); ports acknowledges per the argument.
func ackTwoLayers() string {
	return strings.Replace(portScopeCfg, "  domain:\n",
		"  ports2:\n    globs: [\"internal/hexagon/application/order/cancelorder/ports/outbound/**\"]\n    role: port\n  domain:\n", 1)
}

func withAck(cfg, layer, reason string) string {
	return strings.Replace(cfg, "  "+layer+":\n", "  "+layer+":\n    acknowledge:\n      port-scope: \""+reason+"\"\n", 1)
}

func runAck(t *testing.T, cfg string) (code int, out, errs string) {
	t.Helper()
	dir := writeRepo(t, map[string]string{".a-check.yml": cfg, portFile: "package outbound\n"})
	var o, e bytes.Buffer
	code = cli.Run([]string{dir}, &o, &e)
	return code, o.String(), e.String()
}

func TestAcknowledgeShortensNotToSilences(t *testing.T) { // AC-FA-CONF-002 happy: eine Zeile statt der Glob-Zeilen, Befunde und Exit-Code gleich
	c0, o0, e0 := runAck(t, portScopeCfg)
	c1, o1, e1 := runAck(t, withAck(portScopeCfg, "ports", "app-weite Ports, bewusst"))
	if c0 != c1 || o0 != o1 {
		t.Fatalf("die Quittierung betrifft nur die Meldung: Exit %d/%d, stdout %q/%q", c0, c1, o0, o1)
	}
	if !strings.Contains(e0, "ports/outbound/** -> Scope") {
		t.Fatalf("Ausgangslage ohne Quittierung meldet den Glob: %q", e0)
	}
	want := "  Schicht ports: quittiert — app-weite Ports, bewusst (1 Port-Glob(s), Ursache: direction)\n"
	if !strings.Contains(e1, want) {
		t.Fatalf("eine Zeile mit Schicht, Begruendung, Zahl und Ursache fehlt: %q", e1)
	}
	if strings.Contains(e1, "-> Scope") || strings.Contains(e1, "Abhilfe") {
		t.Fatalf("keine Zeile je Port-Glob und keine Abhilfe fuer die quittierte Schicht: %q", e1)
	}
	if !strings.Contains(e1, "Hinweis: 1 Port-Glob(s) erreichen den App-Baum nicht") {
		t.Fatalf("die Kopfzeile zaehlt auch quittierte Globs: %q", e1)
	}
}

func TestAcknowledgeIsPerLayer(t *testing.T) { // AC-FA-CONF-002 boundary: nur die erste Schicht quittiert, die zweite meldet wie bisher
	_, _, e := runAck(t, withAck(ackTwoLayers(), "ports", "bewusst"))
	if !strings.Contains(e, "Schicht ports: quittiert — bewusst") {
		t.Fatalf("erste Schicht quittiert: %q", e)
	}
	if !strings.Contains(e, "cancelorder/ports/outbound/** -> Scope") || !strings.Contains(e, "direction geben") {
		t.Fatalf("zweite Schicht meldet Glob-Zeile und Abhilfe: %q", e)
	}
	if strings.Contains(e, "createorder/ports/outbound/** -> Scope") {
		t.Fatalf("die quittierte Schicht hat keine Glob-Zeile: %q", e)
	}
	if !strings.Contains(e, "Hinweis: 2 Port-Glob(s)") {
		t.Fatalf("Kopfzeile zaehlt alle: %q", e)
	}
	if strings.Index(e, "quittiert") > strings.Index(e, "cancelorder/ports") {
		t.Fatalf("erst die quittierten Schichten, dann die Glob-Zeilen: %q", e)
	}
}

func TestAcknowledgeCapCountsOnlyUnacknowledged(t *testing.T) { // ADR-0045 Punkt 2: der Deckel gilt fuer die Glob-Zeilen nicht quittierter Schichten
	var globs strings.Builder
	for i := 0; i < 12; i++ {
		fmt.Fprintf(&globs, "      - \"internal/hexagon/application/order/createorder/ports/p%02d/**\"\n", i)
	}
	cfg := strings.Replace(portScopeCfg,
		"      - \"internal/hexagon/application/order/createorder/ports/outbound/**\"\n", globs.String(), 1)
	_, _, e := runAck(t, withAck(cfg, "ports", "zwoelf"))
	if !strings.Contains(e, "(12 Port-Glob(s), Ursache: App-Glob ist groeber als der Scope)") {
		t.Fatalf("die Zeile nennt die Zahl der Globs: %q", e)
	}
	if strings.Contains(e, "weitere") || strings.Contains(e, "-> Scope") {
		t.Fatalf("nichts bleibt zu kuerzen, wenn alles quittiert ist: %q", e)
	}
}

func TestAcknowledgeRedundantIsToldWithoutOtherNotice(t *testing.T) { // AC-FA-CONF-002 boundary: ueberfluessig, auch wenn sonst keine Meldung entsteht
	cfg := strings.Replace(portScopeCfg, "  ports:\n", "  ports:\n    direction: outbound\n", 1) // loest den Scope auf
	cfg = withAck(cfg, "ports", "veraltet")
	c, _, e := runAck(t, cfg)
	if c != 0 {
		t.Fatalf("der Hinweis aendert den Exit-Code nicht, got %d", c)
	}
	want := "Hinweis: Schicht ports: Quittierung ohne Anlass — die Schicht löst die Port-Scope-Meldung nicht aus (Begründung: veraltet)\n"
	if !strings.HasSuffix(e, "\n"+want) || strings.Count(e, "Hinweis:") != 1 {
		t.Fatalf("genau die eine Hinweis-Zeile erwartet, got %q", e)
	}
}

func TestAcknowledgeConfigErrorIsExitTwo(t *testing.T) { // AC-FA-CONF-002 negative
	for name, cfg := range map[string]string{
		"Rolle domain": strings.Replace(portScopeCfg, "  domain:\n", "  domain:\n    acknowledge: {port-scope: \"x\"}\n", 1),
		"leer":         withAck(portScopeCfg, "ports", "  "),
	} {
		c, _, e := runAck(t, cfg)
		if c != 2 || !strings.Contains(e, "acknowledge") || !strings.Contains(e, "Schicht") {
			t.Errorf("%s: Exit 2 mit Schicht und Fehler erwartet, got %d %q", name, c, e)
		}
	}
}

func TestAcknowledgeIsDeterministic(t *testing.T) { // AC-QA-01
	cfg := withAck(ackTwoLayers(), "ports", "bewusst")
	_, _, a := runAck(t, cfg)
	_, _, b := runAck(t, cfg)
	if a != b {
		t.Fatalf("zwei Laeufe, zwei Ausgaben:\n%q\n%q", a, b)
	}
}
