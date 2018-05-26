package output

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFileSystem_Write(t *testing.T) {
	// Happy path
	{
		dir, err := ioutil.TempDir("", "modoc_test")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(dir)

		path := filepath.Join(dir, "section", "subsection", "subsection2")
		contents := `
	This is
	Multiline
	# Test Contents
	`

		var f FileSystem
		if err := f.Write(path, contents); err != nil {
			t.Fatal(err)
		}

		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}

		if string(bytes) != contents {
			t.Errorf("Unexpected contents written, expected=%v, got=%v", contents, string(bytes))
		}
	}

	// Error cases
	tests := []string{
		"/do/not/have/permissions",
		// Directories aren't valid
		"/",
		"../",
		".",
	}

	for idx, path := range tests {
		var f FileSystem
		if err := f.Write(path, "content"); err == nil {
			t.Errorf("[%v] Expected error for bad path, got nil", idx)
		}
	}
}
