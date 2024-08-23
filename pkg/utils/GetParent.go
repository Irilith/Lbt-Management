package utils

import (
	"path/filepath"
	"strings"
)

func SplitFolder(path string) (string, string) {
	if !strings.Contains(path, "/") {
		return path, path
	}
	splitPath := strings.Split(path, "/")
	if len(splitPath) == 0 {
		return splitPath[0], ""
	}
	if strings.Contains(path, ".") {
		return splitPath[len(splitPath)-2], splitPath[len(splitPath)-1]
	} else {
		splitPath := strings.Split(path, "/")
		return splitPath[len(splitPath)-1], splitPath[len(splitPath)-0]
	}
}

func GetParent(path string) string {
	dirPath := filepath.Dir(path)
	base := filepath.Base(dirPath)
	return base
}
