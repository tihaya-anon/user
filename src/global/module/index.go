package module

import (
	"fmt"
	"os"
	"path/filepath"
)

var moduleRoot string

func init() {
	moduleRoot, _ = findGoModDir()
}

func findGoModDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		goModPath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("go.mod not found")
}

func GetSrc() string {
	return moduleRoot
}

func GetRoot() string {
	return filepath.Dir(moduleRoot)
}