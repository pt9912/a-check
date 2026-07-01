package extract

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/pt9912/a-check/internal/hexagon/core"
)

func syms(imps []core.Import) []string {
	var s []string
	for _, i := range imps {
		s = append(s, i.Symbol)
	}
	return s
}

func has(ss []string, want string) bool {
	for _, s := range ss {
		if s == want {
			return true
		}
	}
	return false
}

func TestGoImports(t *testing.T) { // AC-FA-EXTRACT-001 happy + block + alias/underscore
	src := "package x\nimport \"fmt\"\nimport (\n\t\"os\"\n\t_ \"embed\"\n)\n"
	got := syms(newAdapter().importsFromSource("go", stripComments(src)))
	for _, want := range []string{"fmt", "os", "embed"} {
		if !has(got, want) {
			t.Fatalf("missing %q in %v", want, got)
		}
	}
}

func TestRustAliasAndCrate(t *testing.T) { // AC-FA-EXTRACT-001 boundary: use x as y
	got := syms(newAdapter().importsFromSource("rust", stripComments("use tauri as t;\nextern crate serde;\n")))
	if !has(got, "tauri") || !has(got, "serde") {
		t.Fatalf("rust alias/crate not extracted: %v", got)
	}
}

func TestCommentsNotCounted(t *testing.T) { // AC-FA-EXTRACT-001 negative
	src := "// #include \"evil.h\"\n#include \"real.h\"\n/* #include \"blk.h\" */\n"
	got := syms(newAdapter().importsFromSource("cpp", stripComments(src)))
	if has(got, "evil.h") || has(got, "blk.h") {
		t.Fatalf("imports inside comments must be ignored: %v", got)
	}
	if !has(got, "real.h") {
		t.Fatalf("real include missing: %v", got)
	}
}

func TestKotlinImport(t *testing.T) {
	got := syms(newAdapter().importsFromSource("kotlin", stripComments("import a.b.C\n")))
	if !has(got, "a.b.C") {
		t.Fatalf("kotlin import missing: %v", got)
	}
}

func TestJavaImport(t *testing.T) { // AC-FA-EXTRACT-001 happy (Java): dotted import, `;` toleriert
	got := syms(newAdapter().importsFromSource("java", stripComments("package x;\nimport com.foo.Bar;\n")))
	if !has(got, "com.foo.Bar") {
		t.Fatalf("java import missing: %v", got)
	}
}

func TestJavaStaticImport(t *testing.T) { // AC-FA-EXTRACT-001 boundary (Java): import static -> static übersprungen
	got := syms(newAdapter().importsFromSource("java", stripComments("import static com.foo.Bar.baz;\n")))
	if !has(got, "com.foo.Bar.baz") {
		t.Fatalf("java static import nicht als Pfad extrahiert: %v", got)
	}
	if has(got, "static") {
		t.Fatalf("'static' darf nicht als Symbol gegriffen werden: %v", got)
	}
}

func TestJavaCommentNotCounted(t *testing.T) { // AC-FA-EXTRACT-001 negative (sprach-agnostisch, Java)
	got := syms(newAdapter().importsFromSource("java", stripComments("// import com.evil.X;\nimport com.real.Y;\n")))
	if has(got, "com.evil.X") {
		t.Fatalf("java import im Kommentar muss ignoriert werden: %v", got)
	}
	if !has(got, "com.real.Y") {
		t.Fatalf("realer java import fehlt: %v", got)
	}
}

func TestJavaStaticInPath(t *testing.T) { // AC-FA-EXTRACT-001: `static` nur direkt nach `import ` übersprungen, nicht im Pfad
	got := syms(newAdapter().importsFromSource("java", stripComments("import com.static.Foo;\n")))
	if !has(got, "com.static.Foo") {
		t.Fatalf("'static' als Pfad-Segment muss erhalten bleiben: %v", got)
	}
	if has(got, "static") {
		t.Fatalf("'static' darf hier nicht als eigenes Symbol auftauchen: %v", got)
	}
}

func TestJavaStaticMultiWhitespace(t *testing.T) { // AC-FA-EXTRACT-001: `import static` mit Mehrfach-Whitespace
	got := syms(newAdapter().importsFromSource("java", stripComments("import   static   com.x;\n")))
	if !has(got, "com.x") || has(got, "static") {
		t.Fatalf("import static mit Mehrfach-Whitespace: erwarte com.x ohne 'static', got %v", got)
	}
}

func TestJavaWildcard(t *testing.T) { // AC-FA-EXTRACT-001 Out-of-Scope: Wildcard heuristisch gegriffen (Trailing-Dot-Symbol)
	got := syms(newAdapter().importsFromSource("java", stripComments("import com.foo.*;\n")))
	if !has(got, "com.foo.") {
		t.Fatalf("Wildcard heuristisch: erwarte Symbol 'com.foo.' (Trailing-Dot, nicht expandiert), got %v", got)
	}
}

func TestBackendRegistrySet(t *testing.T) { // slice-017: Registry ist die Single Source — genau {cpp,go,rust,kotlin,java}
	got := make([]string, 0)
	for n := range newAdapter().backends {
		got = append(got, n)
	}
	sort.Strings(got)
	if strings.Join(got, ",") != "cpp,go,java,kotlin,rust" {
		t.Fatalf("Backend-Registry = %v, erwarte cpp/go/java/kotlin/rust", got)
	}
}

func TestCheckLanguagesUnknown(t *testing.T) { // slice-017: unbekannte Sprache -> Fehler; exaktes Meldungsformat gepinnt
	err := newAdapter().checkLanguages(map[string][]string{"python": {"**/*.py"}})
	if err == nil {
		t.Fatal("erwarte Fehler für unbekannte Sprache")
	}
	if err.Error() != `unbekannte Sprache "python" (cpp|go|java|kotlin|rust)` {
		t.Fatalf("Meldungsformat driftet (Name/Menge/Klammerung/Reihenfolge): %q", err.Error())
	}
}

func TestCheckLanguagesCaseSensitive(t *testing.T) { // slice-017: Sprach-Keys sind case-sensitiv — "Go" != "go"
	err := newAdapter().checkLanguages(map[string][]string{"Go": {"**/*.go"}})
	if err == nil || !strings.Contains(err.Error(), `"Go"`) {
		t.Fatalf("Case-Variante 'Go' muss brechen (Registry-Lookup ist case-sensitiv), got %v", err)
	}
}

func TestCheckLanguagesMixedUnsupported(t *testing.T) { // slice-017: Mono-Repo go+unsupported -> Fehler (go rettet nicht), positions-unabhängig
	// typescript sortiert NACH go — die unsupported bricht, obwohl go zuerst geprüft wird.
	err := newAdapter().checkLanguages(map[string][]string{"go": {"**/*.go"}, "typescript": {"**/*.ts"}})
	if err == nil || !strings.Contains(err.Error(), "typescript") || !strings.Contains(err.Error(), "unbekannte Sprache") {
		t.Fatalf("gemischt (unsupported nach go): typescript muss brechen, got %v", err)
	}
	// csharp sortiert VOR go — auch die zuerst-sortierte unsupported bricht.
	err = newAdapter().checkLanguages(map[string][]string{"csharp": {"**/*.cs"}, "go": {"**/*.go"}})
	if err == nil || !strings.Contains(err.Error(), "csharp") {
		t.Fatalf("gemischt (unsupported vor go): csharp muss brechen, got %v", err)
	}
}

func TestCheckLanguagesSupported(t *testing.T) { // slice-017: nur unterstützte Sprachen (Mono-Repo go+cpp) -> kein Fehler
	if err := newAdapter().checkLanguages(map[string][]string{"go": {"**/*.go"}, "cpp": {"**/*.cpp"}}); err != nil {
		t.Fatalf("go+cpp (beide unterstützt) müssen akzeptiert werden, got %v", err)
	}
}

func TestExtractSetsLanguage(t *testing.T) { // ADR-0016 (F5): Extract markiert jede Datei mit ihrer Sprache (fürs Threading)
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "x.go"), []byte("package x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	files, err := newAdapter().Extract(dir, core.Model{Languages: map[string][]string{"go": {"**/*.go"}}})
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 || files[0].Language != "go" {
		t.Fatalf("Extract muss Language='go' setzen (Threading-Quelle), got %+v", files)
	}
}
