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

func TestPythonImport(t *testing.T) { // AC-FA-EXTRACT-001 happy (Python): dotted import
	got := syms(newAdapter().importsFromSource("python", prepSource("python", "import myapp.adapters.db\n")))
	if !has(got, "myapp.adapters.db") {
		t.Fatalf("python import fehlt: %v", got)
	}
}

func TestPythonFromImport(t *testing.T) { // AC-FA-EXTRACT-001 boundary (Python from): Modulpfad nach `from`, Namen nicht expandiert
	got := syms(newAdapter().importsFromSource("python", prepSource("python", "from myapp.adapters import db\n")))
	if !has(got, "myapp.adapters") {
		t.Fatalf("python from-import: Modulpfad fehlt: %v", got)
	}
	if has(got, "db") || has(got, "myapp.adapters.db") {
		t.Fatalf("python from-import: Namen dürfen nicht expandiert werden: %v", got)
	}
}

func TestPythonImportAlias(t *testing.T) { // AC-FA-EXTRACT-001 boundary (Python Alias): `as x` nicht gewertet
	got := syms(newAdapter().importsFromSource("python", prepSource("python", "import myapp.adapters as ad\nfrom myapp import db as d\n")))
	if !has(got, "myapp.adapters") || !has(got, "myapp") {
		t.Fatalf("python alias: Modulpfade fehlen: %v", got)
	}
	if has(got, "ad") || has(got, "d") || has(got, "as") {
		t.Fatalf("python alias darf nicht als Symbol gegriffen werden: %v", got)
	}
}

func TestPythonRelativeNotCounted(t *testing.T) { // AC-FA-EXTRACT-001 Out-of-Scope: dokumentierte Grenze der Python-Extraktion (unabhängig vom gültigen relative-Modus, ADR-0017)
	got := syms(newAdapter().importsFromSource("python", prepSource("python", "from . import x\nfrom ..pkg import y\nfrom .sibling import z\n")))
	if len(got) != 0 {
		t.Fatalf("relative Python-Importe dürfen nicht extrahiert werden (Extraktions-Grenze), got %v", got)
	}
}

func TestPythonImportKeywordPrefix(t *testing.T) { // Mutanten-Boundary (slice-014-Lerneintrag): `importlib.…` ist kein Import-Statement
	got := syms(newAdapter().importsFromSource("python", prepSource("python", "importlib.reload(x)\nimportant = 1\n")))
	if len(got) != 0 {
		t.Fatalf("Keyword-als-Präfix (importlib/important) darf nicht matchen: %v", got)
	}
}

func TestPythonFromKeywordPrefix(t *testing.T) { // Mutanten-Boundary: `from` braucht Whitespace + `import`-Wortgrenze (beidseitig)
	got := syms(newAdapter().importsFromSource("python", prepSource("python", "fromage import x\nfrom a.b\nfrom a.b importx\nfrom a.bimport x\n")))
	if len(got) != 0 {
		t.Fatalf("fromage/`from a.b` ohne import/importx/bimport dürfen nicht matchen: %v", got)
	}
}

func TestPythonFromCommentAndMidline(t *testing.T) { // Review-R1 (Test-Linse): pyFrom-Anker gepinnt — #-Kommentar und Mid-Line matchen nie
	got := syms(newAdapter().importsFromSource("python", prepSource("python", "# from evil import bad\nx = 1; from mid import y\nfrom real.mod import z\n")))
	if has(got, "evil") || has(got, "mid") {
		t.Fatalf("from-Import im #-Kommentar/Mid-Line muss ignoriert werden (Anker): %v", got)
	}
	if !has(got, "real.mod") {
		t.Fatalf("realer from-Import fehlt: %v", got)
	}
}

func TestPythonUnderscoreAndDigits(t *testing.T) { // Review-R1 (Test-Linse): Zeichenklassen gepinnt — _thread/__future__/snake_case/oauth2
	got := syms(newAdapter().importsFromSource("python", prepSource("python", "import _thread\nfrom __future__ import annotations\nimport myapp.data_access\nimport oauth2.client\n")))
	for _, want := range []string{"_thread", "__future__", "myapp.data_access", "oauth2.client"} {
		if !has(got, want) {
			t.Fatalf("Underscore-/Ziffern-Modul %q fehlt: %v", want, got)
		}
	}
}

func TestPythonGlobStringKeepsImports(t *testing.T) { // Review-R1 (Code-Linse): /*-Bytefolge im String darf echte Imports nicht fressen (prepSource: kein C-Strip für Python)
	src := "GLOB = \"**/*.py\"\nfrom myapp.adapters import db\nimport myapp.ports.repo\n"
	got := syms(newAdapter().importsFromSource("python", prepSource("python", src)))
	if !has(got, "myapp.adapters") || !has(got, "myapp.ports.repo") {
		t.Fatalf("Imports nach Glob-String verschluckt (C-Strip auf Python): %v", got)
	}
}

func TestPythonHashCommentNotCounted(t *testing.T) { // AC-FA-EXTRACT-001 negative: #-Kommentarzeile matcht die verankerten Muster nie
	got := syms(newAdapter().importsFromSource("python", prepSource("python", "# import evil\n#import evil2\nimport real.mod\n")))
	if has(got, "evil") || has(got, "evil2") {
		t.Fatalf("python import im #-Kommentar muss ignoriert werden: %v", got)
	}
	if !has(got, "real.mod") {
		t.Fatalf("realer python import fehlt: %v", got)
	}
}

func TestPythonMultiImportFirstOnly(t *testing.T) { // AC-FA-EXTRACT-001 Out-of-Scope: `import a, b` -> Erst-Treffer
	got := syms(newAdapter().importsFromSource("python", prepSource("python", "import alpha, beta\n")))
	if !has(got, "alpha") {
		t.Fatalf("Erst-Treffer alpha fehlt: %v", got)
	}
	if has(got, "beta") {
		t.Fatalf("Mehrfach-Import: beta ist dokumentierte Grenze, darf nicht gegriffen werden: %v", got)
	}
}

func TestPythonMultiWhitespace(t *testing.T) { // Mutanten-Boundary: Mehrfach-Whitespace + Einrückung (funktionslokal, import UND from)
	got := syms(newAdapter().importsFromSource("python", prepSource("python", "    import   myapp.util\n    from   myapp.core   import thing\n")))
	if !has(got, "myapp.util") || !has(got, "myapp.core") {
		t.Fatalf("Mehrfach-Whitespace/Einrückung muss matchen: %v", got)
	}
}

func TestCsharpUsing(t *testing.T) { // AC-FA-EXTRACT-001 happy (C#): using-Direktive, dotted Namespace
	got := syms(newAdapter().importsFromSource("csharp", prepSource("csharp", "using MyApp.Adapters.Db;\n")))
	if !has(got, "MyApp.Adapters.Db") {
		t.Fatalf("csharp using fehlt: %v", got)
	}
}

func TestCsharpStaticGlobal(t *testing.T) { // AC-FA-EXTRACT-001 boundary (C#): static/global übersprungen, nicht als Symbol
	got := syms(newAdapter().importsFromSource("csharp", prepSource("csharp", "using static System.Math;\nglobal using MyApp.Core;\nglobal   using   static   Ns.Util ;\n")))
	for _, want := range []string{"System.Math", "MyApp.Core", "Ns.Util"} {
		if !has(got, want) {
			t.Fatalf("Namespace %q fehlt: %v", want, got)
		}
	}
	if has(got, "static") || has(got, "global") {
		t.Fatalf("'static'/'global' dürfen nicht als Symbol gegriffen werden: %v", got)
	}
}

func TestCsharpStaticInNamespace(t *testing.T) { // Mutanten-Boundary (slice-014-Lerneintrag): Keyword als Namespace-Segment bleibt erhalten
	got := syms(newAdapter().importsFromSource("csharp", prepSource("csharp", "using com.static.Foo;\n")))
	if !has(got, "com.static.Foo") || has(got, "static") {
		t.Fatalf("'static' als Segment muss erhalten bleiben, nie als Symbol: %v", got)
	}
}

func TestCsharpAlias(t *testing.T) { // AC-FA-EXTRACT-001 boundary (C# Alias): Ziel (rechte Seite) gewertet, Alias nie
	got := syms(newAdapter().importsFromSource("csharp", prepSource("csharp", "using Db = MyApp.Adapters.Db;\nusing X=Ns.Y;\n")))
	if !has(got, "MyApp.Adapters.Db") || !has(got, "Ns.Y") {
		t.Fatalf("Alias-Ziele fehlen: %v", got)
	}
	if has(got, "Db") || has(got, "X") {
		t.Fatalf("Alias-Name darf nicht als Symbol gegriffen werden: %v", got)
	}
}

func TestCsharpUsingStatementsNotCounted(t *testing.T) { // AC-FA-EXTRACT-001 negative (C#): Ressourcen-Statements sind keine Direktiven
	src := "using var f = File.Open(p);\nusing (var g = File.Open(q))\nusing FileStream h = File.Open(r);\nusing(var k = m())\n"
	got := syms(newAdapter().importsFromSource("csharp", prepSource("csharp", src)))
	if len(got) != 0 {
		t.Fatalf("using-Statements dürfen nie als Import gewertet werden: %v", got)
	}
}

func TestCsharpSemicolonRequired(t *testing.T) { // Mutanten-Boundary: Pflicht-; direkt nach dem Namen (Kern-Anker gegen Statements)
	got := syms(newAdapter().importsFromSource("csharp", prepSource("csharp", "using System.Text\nusings Ns.X;\nusingFoo.Y;\n")))
	if len(got) != 0 {
		t.Fatalf("fehlendes ';'/Keyword-Präfix dürfen nicht matchen: %v", got)
	}
}

func TestCsharpCommentNotCounted(t *testing.T) { // AC-FA-EXTRACT-001 negative: C-Strip greift für C# (anders als Python)
	got := syms(newAdapter().importsFromSource("csharp", prepSource("csharp", "// using Evil.X;\n/* using Blk.Y; */\nusing Real.Z;\n")))
	if has(got, "Evil.X") || has(got, "Blk.Y") {
		t.Fatalf("using im Kommentar muss ignoriert werden: %v", got)
	}
	if !has(got, "Real.Z") {
		t.Fatalf("reales using fehlt: %v", got)
	}
}

func TestCsharpGenericAliasNotCounted(t *testing.T) { // AC-FA-EXTRACT-001 Out-of-Scope: Typ-Alias auf generischen Typ / namespace-Deklaration
	got := syms(newAdapter().importsFromSource("csharp", prepSource("csharp", "using L = List<int>;\nnamespace MyApp.Adapters;\n")))
	if len(got) != 0 {
		t.Fatalf("Typ-Alias auf Generic/namespace-Deklaration dürfen nicht matchen: %v", got)
	}
}

func TestCsharpUnderscoreDigitsIndent(t *testing.T) { // Mutanten-Boundary: Zeichenklassen + Einrückung
	got := syms(newAdapter().importsFromSource("csharp", prepSource("csharp", "    using _Internal.V2;\n")))
	if !has(got, "_Internal.V2") {
		t.Fatalf("Underscore-/Ziffern-Namespace (eingerückt) fehlt: %v", got)
	}
}

func TestCsharpStringNotCounted(t *testing.T) { // Review-R1 (Test-Linse): ^-Anker gepinnt — using im String-Literal/mid-line matcht nie
	got := syms(newAdapter().importsFromSource("csharp", prepSource("csharp", "var s = \"using Sneaky.Ns;\";\n        Log(\"using Evil.X;\");\n")))
	if len(got) != 0 {
		t.Fatalf("using im String/mid-line darf nicht matchen (^-Anker): %v", got)
	}
}

func TestCsharpStaticKeywordPrefix(t *testing.T) { // Review-R1 (Test-Linse): static-Gruppe braucht \s+ — 'staticData' ist kein 'static'-Keyword
	got := syms(newAdapter().importsFromSource("csharp", prepSource("csharp", "using staticData.X;\n")))
	if !has(got, "staticData.X") || has(got, "Data.X") {
		t.Fatalf("'static' als Segment-Präfix muss erhalten bleiben: %v", got)
	}
}

func TestBackendRegistrySet(t *testing.T) { // slice-017: Registry ist die Single Source — genau {cpp,csharp,go,java,kotlin,python,rust,typescript}
	got := make([]string, 0)
	for n := range newAdapter().backends {
		got = append(got, n)
	}
	sort.Strings(got)
	if strings.Join(got, ",") != "cpp,csharp,go,java,kotlin,python,rust,typescript" {
		t.Fatalf("Backend-Registry = %v, erwarte cpp/csharp/go/java/kotlin/python/rust/typescript", got)
	}
}

func TestCheckLanguagesUnknown(t *testing.T) { // slice-017: unbekannte Sprache -> Fehler; exaktes Meldungsformat gepinnt
	err := newAdapter().checkLanguages(map[string][]string{"ruby": {"**/*.rb"}})
	if err == nil {
		t.Fatal("erwarte Fehler für unbekannte Sprache")
	}
	if err.Error() != `unbekannte Sprache "ruby" (cpp|csharp|go|java|kotlin|python|rust|typescript)` {
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
	// ruby sortiert NACH go — die unsupported bricht, obwohl go zuerst geprüft wird.
	// (typescript ist seit slice-022 ein Backend und taugt nicht mehr als Fixture —
	// derselbe Präzedenzfall wie csharp in slice-021.)
	err := newAdapter().checkLanguages(map[string][]string{"go": {"**/*.go"}, "ruby": {"**/*.rb"}})
	if err == nil || !strings.Contains(err.Error(), "ruby") || !strings.Contains(err.Error(), "unbekannte Sprache") {
		t.Fatalf("gemischt (unsupported nach go): ruby muss brechen, got %v", err)
	}
	// fsharp sortiert VOR go — auch die zuerst-sortierte unsupported bricht.
	// (csharp ist seit slice-021 ein Backend und taugt nicht mehr als Fixture.)
	err = newAdapter().checkLanguages(map[string][]string{"fsharp": {"**/*.fs"}, "go": {"**/*.go"}})
	if err == nil || !strings.Contains(err.Error(), "fsharp") {
		t.Fatalf("gemischt (unsupported vor go): fsharp muss brechen, got %v", err)
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

func TestExtractExcludeSkipsFiles(t *testing.T) { // ADR-0018 (slice-023): exclude nimmt Dateien VOR der Extraktion aus — sie liefern gar nichts
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "core"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "core", "x.go"), []byte("package core\nimport \"os\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// Die Test-Datei enthält Import UND forbidden-construct-Muster — beides darf
	// nie auftauchen, weil die Datei nicht gelesen wird (Scan-Scope, ADR-0018).
	if err := os.WriteFile(filepath.Join(dir, "core", "x_test.go"), []byte("package core\nimport \"os\"\nimpl leak\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	m := core.Model{
		Languages: map[string][]string{"go": {"**/*.go"}},
		Exclude:   []string{"**/*_test.go"},
		Forbidden: map[string][]string{"": {"impl "}},
	}
	files, err := newAdapter().Extract(dir, m)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 || files[0].Path != "core/x.go" {
		t.Fatalf("die exclude-Datei darf im Ergebnis gar nicht existieren, got %+v", files)
	}
	// Ohne exclude: byte-identisches Alt-Verhalten — beide Dateien im Scan.
	m.Exclude = nil
	files, err = newAdapter().Extract(dir, m)
	if err != nil || len(files) != 2 {
		t.Fatalf("ohne exclude müssen beide Dateien gescannt werden, got %v / %+v", err, files)
	}
}

func TestDirActionPrunePredicate(t *testing.T) { // ADR-0025 (slice-035): die reine Prune-Entscheidung — root-frei, deterministisch, diskriminierend
	// Der load-bearing Beweis: nur rekursive Teilbaum-Muster (…/** bzw. **)
	// prunen; Nicht-Teilbaum-Formen (…/*, …/, Datei-Globs) dürfen NICHT prunen
	// (sonst stiller Coverage-Verlust, AC-QA-02). Deckt zugleich .git-Skip und
	// den Wurzel-Guard (rel == ".") ab — beides über Extract kaum beobachtbar.
	cases := []struct {
		name    string
		rel     string
		dirName string
		exclude []string
		want    error
	}{
		// prune: rekursive Teilbaum-Muster
		{".security/** prunet .security", ".security", ".security", []string{".security/**"}, filepath.SkipDir},
		{"dist/** prunet dist", "dist", "dist", []string{"dist/**"}, filepath.SkipDir},
		{"**/node_modules/** prunet a/node_modules", "a/node_modules", "node_modules", []string{"**/node_modules/**"}, filepath.SkipDir},
		{"**/node_modules/** prunet tiefes b/c/node_modules", "b/c/node_modules", "node_modules", []string{"**/node_modules/**"}, filepath.SkipDir},
		{"** prunet jedes Verzeichnis", "irgendwo", "irgendwo", []string{"**"}, filepath.SkipDir},
		// KEIN prune: Nicht-Teilbaum-Formen (F-1: sonst über-prune → False-Green)
		{"src/* prunet src NICHT (Single-Segment)", "src", "src", []string{"src/*"}, nil},
		{"trailing-slash sub/ prunet NICHT", "sub", "sub", []string{"sub/"}, nil},
		{"Datei-Glob **/*_test.go prunet core NICHT", "core", "core", []string{"**/*_test.go"}, nil},
		{"dir/**/*.go ist nicht teilbaum-deckend → kein prune", "dir", "dir", []string{"dir/**/*.go"}, nil},
		{"node_modules/** (root-verankert) prunet pkg/node_modules NICHT", "pkg/node_modules", "node_modules", []string{"node_modules/**"}, nil},
		{"Präfix-Kollision: a/** prunet ab NICHT", "ab", "ab", []string{"a/**"}, nil},
		// Sonderfälle
		{".git wird immer geskippt", ".git", ".git", nil, filepath.SkipDir},
		{"Scan-Wurzel wird nie geprunet (auch nicht von **)", ".", "root", []string{"**"}, nil},
		{"leere exclude-Liste: kein prune", "any", "any", nil, nil},
	}
	for _, c := range cases {
		if got := dirAction(c.rel, c.dirName, c.exclude); got != c.want {
			t.Errorf("%s: dirAction(%q,%q,%v) = %v, want %v", c.name, c.rel, c.dirName, c.exclude, got, c.want)
		}
	}
}

func TestExtractExcludeNoOverPruneOnSingleStar(t *testing.T) { // ADR-0025 / F-1: ein Single-Star-Glob (src/*) darf den Teilbaum NICHT prunen
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "src", "app"), 0o755); err != nil {
		t.Fatal(err)
	}
	// src/loose.go matcht src/* (Datei-Ausschluss); src/app/x.go matcht src/*
	// NICHT (^src/[^/]*$) und darf daher NICHT durch einen Verzeichnis-Prune
	// verloren gehen — sonst stille Coverage-Lücke (der gefixte F-1-Bug).
	if err := os.WriteFile(filepath.Join(dir, "src", "loose.go"), []byte("package src\nimport \"os\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "src", "app", "x.go"), []byte("package app\nimport \"fmt\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// Zusätzlich eine Nicht-Sprach-Datei: deckt den fileImports-keep=false-Zweig (lang == "").
	if err := os.WriteFile(filepath.Join(dir, "src", "app", "README.md"), []byte("# doc\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	m := core.Model{
		Languages: map[string][]string{"go": {"**/*.go"}},
		Exclude:   []string{"src/*"},
	}
	files, err := newAdapter().Extract(dir, m)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 || files[0].Path != "src/app/x.go" {
		t.Fatalf("src/* darf src/app/x.go nicht mit-prunen (loose.go datei-ausgeschlossen, x.go bleibt), got %+v", files)
	}
}

func TestExtractExcludePrunesNestedSubtree(t *testing.T) { // ADR-0025: End-to-End des realen Falls — ".security/**" nimmt den ganzen Teilbaum aus, Geschwister bleiben
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, ".security", "cache"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".security", "cache", "x.go"), []byte("package cache\nimport \"os\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "core.go"), []byte("package core\nimport \"fmt\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	m := core.Model{
		Languages: map[string][]string{"go": {"**/*.go"}},
		Exclude:   []string{".security/**"},
	}
	files, err := newAdapter().Extract(dir, m)
	if err != nil {
		t.Fatal(err)
	}
	// Für einen …/**-Glob sind Datei-Filter und Prune output-äquivalent (invariant);
	// die diskriminierende Prune-Entscheidung beweist TestDirActionPrunePredicate.
	if len(files) != 1 || files[0].Path != "core.go" {
		t.Fatalf("verschachtelter ausgeschlossener Teilbaum muss vollständig ausfallen, Geschwister bleiben, got %+v", files)
	}
}

func TestExtractExcludeDoesNotPruneOnFileGlob(t *testing.T) { // ADR-0025: ein Datei-Glob (**/*_test.go) prunet das enthaltende Verzeichnis NICHT
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "core"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "core", "x.go"), []byte("package core\nimport \"os\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "core", "x_test.go"), []byte("package core\nimport \"testing\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	m := core.Model{
		Languages: map[string][]string{"go": {"**/*.go"}},
		Exclude:   []string{"**/*_test.go"},
	}
	files, err := newAdapter().Extract(dir, m)
	if err != nil {
		t.Fatal(err)
	}
	// Das Verzeichnis "core" wird betreten (x.go gescannt), nur die Testdatei fällt datei-genau aus.
	if len(files) != 1 || files[0].Path != "core/x.go" {
		t.Fatalf("ein Datei-Glob darf das Verzeichnis nicht prunen (x.go bleibt, x_test.go fällt), got %+v", files)
	}
}

// --- slice-022: TypeScript-Backend (AC-FA-EXTRACT-001) ---

func tsSyms(src string) []string {
	return syms(newAdapter().importsFromSource("typescript", prepSource("typescript", src)))
}

func TestTypescriptImportForms(t *testing.T) { // Happy: Clauses, beide Quote-Arten, Semikolon optional (ASI)
	src := "import { Db } from '../adapters/db';\n" +
		"import Repo from \"../adapters/repo\"\n" + // Default-Clause, Double-Quotes, semikolonfrei
		"import * as ns from './ports/ns';\n" + // Namespace-Clause
		"import A, { B as C } from './ports/mixed';\n" + // gemischte Clause
		"import type { T } from './ports/t';\n" + // Typ-Import (Entscheid B: gewertet)
		"import { type U, f } from './ports/u';\n" // Inline-type-Modifier
	got := tsSyms(src)
	for _, want := range []string{"../adapters/db", "../adapters/repo", "./ports/ns", "./ports/mixed", "./ports/t", "./ports/u"} {
		if !has(got, want) {
			t.Fatalf("TS-Importform fehlt: %q in %v", want, got)
		}
	}
}

func TestTypescriptSideEffectAndRequire(t *testing.T) { // Boundary: Seiteneffekt-Import + Interop-Form
	got := tsSyms("import './polyfill';\nimport fs = require('fs');\nimport pg = require(\"pg\")\n")
	for _, want := range []string{"./polyfill", "fs", "pg"} {
		if !has(got, want) {
			t.Fatalf("Seiteneffekt/require-Form fehlt: %q in %v", want, got)
		}
	}
}

func TestTypescriptReexports(t *testing.T) { // Boundary (Entscheid C): Re-Exports sind echte Abhängigkeits-Kanten (Barrel-Dateien)
	src := "export * from './core/model';\n" +
		"export { X } from \"./x\"\n" +
		"export * as ns from './y';\n" +
		"export type { T } from './z';\n"
	got := tsSyms(src)
	for _, want := range []string{"./core/model", "./x", "./y", "./z"} {
		if !has(got, want) {
			t.Fatalf("Re-Export-Form fehlt: %q in %v", want, got)
		}
	}
}

func TestTypescriptContinuationLine(t *testing.T) { // Entscheid G (BLOCKER-Fix): Prettier-umbrochener Import — Schlusszeile } from '…'
	src := "import {\n  makeDb,\n  closeDb,\n} from '../adapters/db';\n" +
		"export {\n  A,\n} from \"../adapters/a\"\n" +
		"}\n" // nacktes } ohne from: kein Match
	got := tsSyms(src)
	if !has(got, "../adapters/db") || !has(got, "../adapters/a") {
		t.Fatalf("Fortsetzungszeile '} from …' muss den Specifier liefern, got %v", got)
	}
	if len(got) != 2 {
		t.Fatalf("nur die beiden Fortsetzungs-Specifier erwartet, got %v", got)
	}
}

func TestTypescriptExpressionNoMatch(t *testing.T) { // Negative: Ausdrucks-Zeilen, Keyword-Präfix, Kommentare, Triple-Slash
	src := "const m = await import('./lazy');\n" + // dynamisches import im Ausdruck
		"import('./x').then(m => m.run());\n" + // zeilen-anführendes Ausdrucks-import (schärfster tsSide-Mutant)
		"const x = require('pg');\n" + // require im Ausdruck
		"export const q = knex.from('users');\n" + // blockiert vom from(-Anker (kein Quote nach from)
		"export const x = from './adapters/db5';\n" + // pinnt die Mittelteil-Klasse: '=' darf nie überspannt werden (Review-R1 T-1)
		"export obj.prop from './adapters/db7';\n" + // pinnt die Mittelteil-Klasse: '.' darf nie überspannt werden (Review-R1 T-1)
		"export {};\n" + // leerer Export ohne from
		"declare module './mod' {\n" + // Ambient-Deklaration
		"importX from './evil';\n" + // Keyword-Präfix
		"exportX from './evil5';\n" + // Keyword-Präfix (export-Seite)
		"import { X } from `./y`\n" + // Backtick-Specifier: Template-Literal-Grenze, nie gegriffen
		"const s = 'import { X } from \"./evil2\"';\n" + // Import-String mitten im Ausdruck
		"// import { Y } from './evil3';\n" + // Zeilen-Kommentar (C-Strip)
		"/* import { Z } from './evil4'; */\n" + // Block-Kommentar (C-Strip)
		"/// <reference path=\"./types.d.ts\" />\n" // Triple-Slash-Direktive (fällt dem C-Strip zu)
	if got := tsSyms(src); len(got) != 0 {
		t.Fatalf("Ausdrucks-/Kommentar-Zeilen dürfen nie Symbole liefern, got %v", got)
	}
}

func TestTypescriptJsSpecifier(t *testing.T) { // Happy: NodeNext-.js-Specifier auf .ts-Datei — Symbol wörtlich erhalten
	got := tsSyms("import { Repo } from \"../adapters/db.js\"\n")
	if !has(got, "../adapters/db.js") {
		t.Fatalf(".js-Specifier muss wörtlich geliefert werden, got %v", got)
	}
}

func TestKotlinDeclarations(t *testing.T) { // slice-031/ADR-0023: Top-Level-Deklarationen erkannt, unabhängig vom Dateinamen
	src := stripComments("package com.ex.conn\n" +
		"fun Pool.asJdbc(): Connection = TODO()\n" +
		"class HikariPool\n" +
		"interface Marker\n" +
		"object Registry\n" +
		"typealias Alias = String\n" +
		"const val MAX = 1\n" +
		"internal fun helper() {}\n")
	got := newAdapter().declarations(src)
	want := []string{"Alias", "HikariPool", "MAX", "Marker", "Registry", "asJdbc", "helper"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("Kotlin-Deklarationen = %v, erwarte %v", got, want)
	}
}

func TestKotlinDeclarationsCommentNotCounted(t *testing.T) { // slice-031: Deklaration im Kommentar zählt nicht (Kommentar-Strip)
	src := prepSource("kotlin", "// fun ghost() {}\n/* class Ghost */\nfun real() {}\n")
	if got := newAdapter().declarations(src); len(got) != 1 || got[0] != "real" {
		t.Fatalf("nur die echte Top-Level-fun 'real' erwartet, got %v", got)
	}
}

func TestKotlinDeclarationsIndentedMemberNotCounted(t *testing.T) { // slice-031: eingerückte Member zählen nicht (nur Spalte-0-Top-Level)
	src := stripComments("class Outer {\n    fun member() {}\n    val field = 1\n}\n")
	if got := newAdapter().declarations(src); len(got) != 1 || got[0] != "Outer" {
		t.Fatalf("nur die Top-Level-Klasse 'Outer' erwartet (Member eingerückt), got %v", got)
	}
}

func TestNonKotlinNoDeclarations(t *testing.T) { // slice-031/ADR-0023: nicht deklarations-bewusste Backends liefern ein leeres Set (No-op)
	if got := newAdapter().declarationsFor("go", "func Foo() {}\ntype Bar struct{}\n"); got != nil {
		t.Fatalf("Go-Backend darf keine Deklarationen liefern (No-op), got %v", got)
	}
}

// --- ADR-0034: Zeichenketten-Literale im Kommentar-Strip --------------------

// Der Fall, der v0.17.0 ausgeliefert wurde: ein Glob-String oeffnete einen
// Phantom-Blockkommentar, der alles danach verschluckte.
func TestGlobStringDoesNotSwallowImports(t *testing.T) {
	src := "package core\n\nvar globs = []string{\"/**\"}\n\nimport _ \"fix/internal/adapters\"\n"
	if got := stripComments(src); !strings.Contains(got, "fix/internal/adapters") {
		t.Fatalf("Import nach einem /**-String verschluckt:\n%q", got)
	}
}

// Die //-Variante: eine URL im Import-Specifier.
func TestURLInStringDoesNotStartLineComment(t *testing.T) {
	src := "import { serve } from \"https://deno.land/std/http/server.ts\";\n"
	if got := stripComments(src); !strings.Contains(got, "deno.land/std/http/server.ts") {
		t.Fatalf("URL-Import an // abgeschnitten:\n%q", got)
	}
}

// Backtick-Literale sind mehrzeilig und escapefrei (Go-Raw-Strings, TS-Templates).
func TestBacktickLiteralIsOpaqueAndMultiline(t *testing.T) {
	src := "var p = `a /* b\nc */ d`\nimport _ \"pkg/x\"\n"
	got := stripComments(src)
	for _, want := range []string{"a /* b", "c */ d", "pkg/x"} {
		if !strings.Contains(got, want) {
			t.Fatalf("Backtick-Literal nicht verbatim uebernommen (%q fehlt):\n%q", want, got)
		}
	}
}

// Die Gegenrichtung — ohne sie waere ein Strip, der GAR NICHTS entfernt,
// von einem korrekten nicht zu unterscheiden.
func TestRealCommentsAreStillStripped(t *testing.T) {
	src := "package core\n/*\nimport _ \"blockkommentar/x\"\n*/\n// import _ \"zeilenkommentar/y\"\nimport _ \"echt/z\"\n"
	got := stripComments(src)
	for _, gone := range []string{"blockkommentar/x", "zeilenkommentar/y"} {
		if strings.Contains(got, gone) {
			t.Fatalf("echter Kommentar nicht entfernt (%q noch da):\n%q", gone, got)
		}
	}
	if !strings.Contains(got, "echt/z") {
		t.Fatalf("echter Import mitentfernt:\n%q", got)
	}
}

// Zeilennummern muessen stabil bleiben — Befunde tragen sie.
func TestStripPreservesLineCount(t *testing.T) {
	src := "a\nvar s = \"/*\"\n/* weg\nauch weg */\nb\n"
	if got, want := strings.Count(stripComments(src), "\n"), strings.Count(src, "\n"); got != want {
		t.Fatalf("Zeilenzahl veraendert: %d statt %d", got, want)
	}
}

// Ein unbalanciertes Anfuehrungszeichen darf hoechstens SEINE Zeile kosten,
// nicht den Rest der Datei (ADR-0034, Regel 2).
func TestUnbalancedQuoteEndsAtNewline(t *testing.T) {
	src := "var s = \"offen\nimport _ \"pkg/danach\"\n"
	if got := stripComments(src); !strings.Contains(got, "pkg/danach") {
		t.Fatalf("unbalanciertes Anfuehrungszeichen frass ueber die Zeile hinaus:\n%q", got)
	}
}
