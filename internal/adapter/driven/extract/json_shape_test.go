package extract

import (
	"strings"
	"testing"
)

func eqJSON(t *testing.T, src string, want ...string) {
	t.Helper()
	st, _, err := normalizeJSON(src)
	if err != nil {
		t.Fatalf("normalizeJSON(%q): %v", src, err)
	}
	if got := texts(st); strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Fatalf("normalizeJSON(%q)\n got: %q\nwant: %q", src, got, want)
	}
}

// Ein package.json in realer Form: je Mitglied des Wurzel-Objekts eine Anweisung,
// Leerraum außerhalb von Zeichenketten entfällt (SPEC-EXTRACT-001, json Schritt 2/3).
func TestJSONRealForm(t *testing.T) {
	src := "{\n  \"name\": \"svc\",\n  \"private\": true,\n  \"scripts\": {\n    \"build\": \"tsc -p .\"\n  },\n  \"dependencies\": {\n    \"react\": \"^18.2.0\",\n    \"zod\": \"3.22.4\"\n  }\n}\n"
	want := []string{`"name":"svc"`, `"private":true`, `"scripts":{"build":"tsc -p ."}`, `"dependencies":{"react":"^18.2.0","zod":"3.22.4"}`}
	eqJSON(t, src, want...)
	compact := `{"name":"svc","private":true,"scripts":{"build":"tsc -p ."},"dependencies":{"react":"^18.2.0","zod":"3.22.4"}}`
	eqJSON(t, compact, want...)
	eqJSON(t, "\r\n{ \"name\" :\t\"svc\" ,\r\"private\":true,\"scripts\":{\"build\":\"tsc -p .\"},\"dependencies\":{\"react\":\"^18.2.0\",\"zod\":\"3.22.4\"}}", want...)
}

// Zeichenketten byte-genau: Kommas, Klammern und maskierte Anführungszeichen darin
// trennen kein Mitglied; verschachtelte Kommas ebenso nicht.
func TestJSONStringsAndNesting(t *testing.T) {
	eqJSON(t, `{"a":"x, } { \" y","b":[1,{"c":[2,3]}],"d":null}`, `"a":"x, } { \" y"`, `"b":[1,{"c":[2,3]}]`, `"d":null`)
	eqJSON(t, `{"e":-1.5e+10,"f":false}`, `"e":-1.5e+10`, `"f":false`)
	eqJSON(t, `{}`)
	eqJSON(t, "\uFEFF{\"a\":1}", `"a":1`)
}

// Alle fail-closed-Fälle der Spezifikation: Exit 2, nie still grün.
func TestJSONUnsplittable(t *testing.T) {
	for _, src := range []string{
		"{\"a\":1} // c",          // Kommentar: `/` außerhalb einer Zeichenkette
		"{\"a\":1 # c\n}",         // `#`
		"{'a':1}",                 // JSON5-Anführungszeichen
		"{\"a\":\f1}",             // Seitenvorschub ist kein JSON-Leerraum
		"{\"a\":nu ll}",           // Leerraum zwischen zwei Literal-Zeichen
		"{\"a\":1 2}",             // … zwischen zwei Zahl-Zeichen
		"{\"a\" \"b\":1}",         // … zwischen zwei Zeichenketten
		"{\"a\":\"x\ny\"}",        // Zeilenende in der Zeichenkette
		"{\"a\":\"x\ty\"}",        // Tabulator in der Zeichenkette
		"{\"a\":\"offen}",         // nicht geschlossene Zeichenkette
		"[1,2]",                   // Wurzel ist kein Objekt
		"\"x\"",                   // Wurzel ist kein Objekt
		"",                        // keine Wurzel
		"{\"a\":1}{\"b\":2}",      // Inhalt nach der Wurzel
		"{\"a\":[1}",              // falsch gepaarte Klammer
		"{\"a\":1",                // Wurzel nicht geschlossen
		"{,}",                     // leeres Mitglied
		"{\"a\":1,}",              // abschließendes Komma
		"{1:2}",                   // Mitglied beginnt nicht mit einer Zeichenkette
		"{\"a\"}",                 // Mitglied ohne `:`
	} {
		if _, _, err := normalizeJSON(src); err == nil {
			t.Errorf("%q: Fehler erwartet", src)
		}
	}
}

// Original-Zeile = Zeile des ersten Zeichens des Mitglieds; Zeilenzahl zählt LF,
// CRLF und einzelnes CR.
func TestJSONLines(t *testing.T) {
	st, lines, err := normalizeJSON("{\n\"a\":1,\r\n\r\"b\":\n2\n}\n")
	if err != nil || len(st) != 2 || st[0].Line != 2 || st[1].Line != 4 || lines != 6 {
		t.Fatalf("st=%+v lines=%d err=%v", st, lines, err)
	}
}

// literal-Einträge in Quellform: ein Mitglied, vor dem Zerlegen in { }
// eingeschlossen; ein Eintrag, der selbst schon { … } ist, ergibt kein Mitglied.
func TestJSONLiteralSource(t *testing.T) {
	st, _, err := normalizeJSON(literalSource("json", `"dependencies": { "react": "^18.2.0" }`))
	if err != nil || len(st) != 1 || st[0].Text != `"dependencies":{"react":"^18.2.0"}` {
		t.Fatalf("st=%+v err=%v", st, err)
	}
	if _, _, err := normalizeJSON(literalSource("json", `{"a":1}`)); err == nil {
		t.Fatal(`Eintrag {"a":1}: Fehler erwartet`)
	}
	if literalSource("gomod", "x") != "x" || literalSource("kotlin", "x") != "x" {
		t.Fatal("nur json wird eingeschlossen")
	}
}
