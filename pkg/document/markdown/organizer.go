package markdown

import (
	"path/filepath"

	"github.com/KyleBanks/modoc/pkg/document"
)

// Organizer handles organizing Documents into a project structure.
type Organizer struct {
	w Writer
}

// NewOrganizer constructs and returns an Organizer.
func NewOrganizer(w Writer) Organizer {
	return Organizer{
		w: w,
	}
}

// Organize generates the project structure of a Document to the specified path.
func (o Organizer) Organize(d document.Document, path string) error {
	if err := o.writeTitle(d, path); err != nil {
		return err
	}
	if err := o.writeBody(d, path); err != nil {
		return err
	}

	for _, child := range d.Children {
		if err := o.Organize(child.Document, filepath.Join(path, child.Folder)); err != nil {
			return err
		}
	}
	return nil
}

func (o Organizer) writeTitle(d document.Document, path string) error {
	return o.w.Write(filepath.Join(path, titleFilename), d.Title)
}

func (o Organizer) writeBody(d document.Document, path string) error {
	return o.w.Write(filepath.Join(path, bodyFilename), d.Body)
}
