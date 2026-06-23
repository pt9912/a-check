package extract

import (
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
