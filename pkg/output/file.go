package output

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const defaultFilePermissions = 0644

// FileSystem writes to the filesystem.
type FileSystem struct{}

// Write outputs the contents provided to a particular path, creating all necessary directories
// in the path.
func (FileSystem) Write(path string, contents string) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	return ioutil.WriteFile(path, []byte(contents), defaultFilePermissions)
}
