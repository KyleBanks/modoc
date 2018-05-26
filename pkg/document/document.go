package document

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/KyleBanks/modoc/pkg/output"
)

// ErrDirectoryNotEmpty is returned when trying to operate on a non-empty directory.
var ErrDirectoryNotEmpty = errors.New("The directory is not empty")

// WalkSectionFn defines a function that accepts a Section and the depth at
// which it can be found in the Document, where 1 is the root element.
//
// This can serve as a template for recursively traversing a Document.
type WalkSectionFn func(s Section, depth int) error

// Compiler combines the contents of a Document into
// a single string that can be printed or written to disk.
type Compiler interface {
	Compile(Document) (string, error)
}

// Organizer breaks the contents of a Document into
// a folder structure on the filesystem.
type Organizer interface {
	Organize(d Document, path string) error
}

// Document represents the contents of a single document, broken into
// sections and subsections.
type Document struct {
	Title string
	Body  string

	Children []Section
}

// Section represents a single portion of a Document, with its location on the filesystem.
type Section struct {
	Document
	Folder string
}

// Organize explodes the Document onto the filesystem at the path provided.
func (d Document) Organize(o Organizer, path string) error {
	// Ensure the directory does not already contain files
	if files, err := ioutil.ReadDir(path); err != nil && !os.IsNotExist(err) {
		return err
	} else if len(files) > 0 {
		return ErrDirectoryNotEmpty
	}

	return o.Organize(d, path)
}

// Compile generates a single output file containing the compiled document at the path provided.
func (d Document) Compile(c Compiler, w output.Writer, path string) error {
	output, err := c.Compile(d)
	if err != nil {
		return err
	}

	return w.Write(path, output)
}

// ForEachChild runs the provided WalkSectionFn for each child of the Document.
//
// Note: This does not run recursively, only the direct children will be run. If
// you want recursive, execute `ForEachChild` again within the WalkSectionFn.
func (d Document) ForEachChild(fn WalkSectionFn, depth int) error {
	for _, child := range d.Children {
		if err := fn(child, depth); err != nil {
			return err
		}
	}
	return nil
}
