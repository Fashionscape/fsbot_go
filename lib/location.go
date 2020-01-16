package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// GetLocalFolder gets the storage folder for fsbot
func GetLocalFolder() string {
	if runtime.GOOS == "darwin" {
		return fmt.Sprintf("%s/.fsbot", os.Getenv("HOME"))
	}

	return ""
}

// LocGet returns a file within the storage folder
func LocGet(file string) string {
	storage := GetLocalFolder()

	if storage != "" {
		if fileExists(file) {
			return file
		}

		return fmt.Sprintf("%s/%s", storage, file)
	}

	return file
}

// IsImage checks if the given filename is an image format
func IsImage(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case "png":
		return true
	case "jpg":
		return true
	case "jpeg":
		return true
	default:
		return false
	}
}

func fileExists(file string) bool {
	info, err := os.Stat(file)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
