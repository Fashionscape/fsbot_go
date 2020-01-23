package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GetLocalFolder gets the storage folder for fsbot
func GetLocalFolder() string {
	folderName := fmt.Sprintf("%s/.fsbot", os.Getenv("HOME"))
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(fmt.Sprintf("%s/.fsbot", os.Getenv("HOME")), 0777)
		Check(err)
	}

	return folderName
}

// LocGet returns a file within the storage folder
func LocGet(file string) string {
	storage := GetLocalFolder()
	fileName := fmt.Sprintf("%s/%s", storage, file)

	if storage != "" {
		if fileExists(fileName) {
			return fileName
		} else {
			err := ioutil.WriteFile(fileName, GetExampleConfig(), 0777)
			Check(err)
		}
		return fileName
	}

	fmt.Println(fileName)
	return file
}

func GetExampleConfig() []byte {
	cont, err := ioutil.ReadFile("config.example.json")
	Check(err)
	return cont
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
