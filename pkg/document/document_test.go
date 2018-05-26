package document

import (
	"errors"
	"testing"
)

type mockCompiler struct {
	compileFn func(Document) (string, error)
}

func (m mockCompiler) Compile(d Document) (string, error) {
	return m.compileFn(d)
}

type mockOrganizer struct {
	organizeFn func(Document, string) error
}

func (m mockOrganizer) Organize(d Document, path string) error {
	return m.organizeFn(d, path)
}

type mockWriter struct {
	writeFn func(string, string) error
}

func (m mockWriter) Write(path, content string) error {
	return m.writeFn(path, content)
}

func TestDocument_Organize(t *testing.T) {
	tests := []struct {
		path      string
		expectErr bool
	}{
		{"/path/doesnt/exist", false},
		{"./path/doesnt/exist", false},
		{".", true}, // not an empty dir
		{"./empty", false},
	}

	for idx, tt := range tests {
		var d Document
		var o mockOrganizer
		o.organizeFn = func(d Document, path string) error {
			return nil
		}

		err := d.Organize(o, tt.path)
		if err == nil && tt.expectErr {
			t.Errorf("[%v] Expected error, got nil", idx)
		} else if err != nil && !tt.expectErr {
			t.Errorf("[%v] Unexpected error, expected=nil, got=%v", idx, err)
		}
	}
}

func TestDocument_Compile(t *testing.T) {
	var d Document

	var w mockWriter
	var retWriterErr error
	w.writeFn = func(path, contents string) error {
		return retWriterErr
	}

	var c mockCompiler
	var retContents string
	var retErr error
	c.compileFn = func(d Document) (string, error) {
		return retContents, retErr
	}

	// Compiler returns error
	retContents = ""
	retErr = errors.New("error compiling")
	if err := d.Compile(c, w, ""); err != retErr {
		t.Errorf("Unexpected error, expected=%v, got=%v", retErr, err)
	}

	// Writer returns error
	retContents = ""
	retErr = nil
	retWriterErr = errors.New("error writing")
	if err := d.Compile(c, w, ""); err != retWriterErr {
		t.Errorf("Unexpected error, expected=%v, got=%v", retWriterErr, err)
	}

	// Happy path
	retContents = `# A Test Book

	## Table of Contents

	- [Testing](#testing)
`
	retErr = nil
	retWriterErr = nil
	expectPath := "/path/to/output"
	w.writeFn = func(path, contents string) error {
		if path != expectPath {
			t.Errorf("Unexpected path, expected=%v, got=%v", expectPath, path)
		} else if contents != retContents {
			t.Errorf("Unexpected contents, expected=%v, got=%v", retContents, contents)
		}
		return nil
	}

	if err := d.Compile(c, w, expectPath); err != nil {
		t.Fatal(err)
	}
}

func TestDocument_ForEachChild(t *testing.T) {
	var doc Document
	var fn WalkSectionFn
	var depth int
	var callCount int

	// No children
	if err := doc.ForEachChild(fn, depth); err != nil {
		t.Errorf("Error when calling with no children, got=%v", err)
	}

	// One child
	doc.Children = []Section{{Folder: "folder1"}}
	depth = 1
	callCount = 0
	fn = func(s Section, d int) error {
		if s.Folder != doc.Children[0].Folder {
			t.Errorf("Unexpected Section provided for one child, expected=%v, got=%v", doc.Children[0], s)
		} else if d != depth {
			t.Errorf("Unexpected depth provided for one child, expected=%v, got=%v", depth, d)
		}
		callCount++
		return nil
	}
	if err := doc.ForEachChild(fn, depth); err != nil {
		t.Errorf("Error when calling with one child, got=%v", err)
	}

	if callCount != 1 {
		t.Errorf("Unexpected call count for one child, expected=1, got=%v", callCount)
	}

	// Multiple children
	doc.Children = []Section{{Folder: "folder1"}, {Folder: "folder2"}, {Folder: "folder3"}}
	depth = 10
	callCount = 0
	fn = func(s Section, d int) error {
		if s.Folder != doc.Children[callCount].Folder {
			t.Errorf("Unexpected Section provided for one child, expected=%v, got=%v", doc.Children[callCount], s)
		} else if d != depth {
			t.Errorf("Unexpected depth provided for one child, expected=%v, got=%v", depth, d)
		}
		callCount++
		return nil
	}
	if err := doc.ForEachChild(fn, depth); err != nil {
		t.Errorf("Error when calling with one child, got=%v", err)
	}

	if callCount != len(doc.Children) {
		t.Errorf("Unexpected call count for multiple children, expected=%v, got=%v", len(doc.Children), callCount)
	}

	// Returns error
	fn = func(s Section, d int) error {
		return errors.New("Test Error")
	}
	if err := doc.ForEachChild(fn, depth); err == nil {
		t.Errorf("Expected error, got nil")
	}
}
