package tools

import (
	"os"
	"path/filepath"
)

func RootPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, nil
}

func BaseDirOs() (string, error) {
	executePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	base_dir := filepath.Dir(executePath)
	return base_dir, nil
}
