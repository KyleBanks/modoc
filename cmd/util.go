package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func cleanPath(path string) (string, error) {
	return filepath.Abs(path)
}

func fail(msg string, err error) {
	if len(msg) == 0 {
		fmt.Println(err)
	} else if err == nil {
		fmt.Println(msg)
	} else {
		fmt.Println(msg, "\n", err)
	}
	os.Exit(1)
}
