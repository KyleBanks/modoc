package markdown

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/KyleBanks/modoc/pkg/document"
)

const (
	titleFilename = "TITLE"
	bodyFilename  = "README.md"
)

// NewDocument generates a Document from the provided source path.
//
// The path provided should be to the root of the source tree.
func NewDocument(path string) (*document.Document, error) {
	var d document.Document

	if err := populate(&d, path); err != nil {
		return nil, err
	}

	if err := findChildren(&d, path); err != nil {
		return nil, err
	}

	return &d, nil
}

func populate(d *document.Document, path string) error {
	t, err := title(path)
	if err != nil {
		return err
	}
	b, err := body(path)
	if err != nil {
		return err
	}

	d.Title = t
	d.Body = b
	return nil
}

func findChildren(d *document.Document, path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		if !f.IsDir() || strings.HasPrefix(f.Name(), ".") || strings.HasPrefix(f.Name(), "_") {
			continue
		}

		child := document.Section{Folder: filepath.Join(path, f.Name())}
		if err := populate(&child.Document, child.Folder); err != nil {
			return err
		}
		if err := findChildren(&child.Document, child.Folder); err != nil {
			return err
		}

		d.Children = append(d.Children, child)
	}

	return nil
}

func title(path string) (string, error) {
	title, err := ioutil.ReadFile(filepath.Join(path, titleFilename))
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(title)), nil
}

func body(path string) (string, error) {
	body, err := ioutil.ReadFile(filepath.Join(path, bodyFilename))
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
		}
	}

	return string(body), nil
}
