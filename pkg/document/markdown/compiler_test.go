package markdown

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/KyleBanks/modoc/pkg/document"
)

// update allows you to set the --update flag when running tests
// to update the golden files.
var update = flag.Bool("update", false, "update .golden files")

func TestCompiler_Compile(t *testing.T) {
	var d document.Document = document.Sample

	tests := []struct {
		c      Compiler
		golden string
	}{
		{Compiler{}, "compile_none.golden"},
		{Compiler{true, false, false}, "compile_header.golden"},
		{Compiler{false, true, false}, "compile_toc.golden"},
		{Compiler{false, false, true}, "compile_body.golden"},
		{Compiler{true, true, false}, "compile_header_toc.golden"},
		{Compiler{true, false, true}, "compile_header_body.golden"},
		{Compiler{false, true, true}, "compile_toc_body.golden"},
		{Compiler{true, true, true}, "compile_all.golden"},
	}

	for idx, tt := range tests {
		out, err := tt.c.Compile(d)
		if err != nil {
			t.Errorf("[%v] Unexected error, expected=nil, got=%v", idx, err)
		}

		golden := filepath.Join("testdata", tt.golden)
		if *update {
			updateGoldenFile(t, golden, out)
		}
		expect := loadGoldenFile(t, golden)

		if out != expect {
			t.Errorf("[%v] Unexpected compiled output, expected=%v, got=%v", idx, expect, out)
		}
	}
}

func updateGoldenFile(t *testing.T, path, content string) {
	if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func loadGoldenFile(t *testing.T, path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(bytes)
}
