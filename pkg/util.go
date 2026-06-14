package pkg

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func GetFileNameFromURL(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "downloaded_file"
	}
	return filepath.Base(u.Path)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func HandleFileCreation(rawURL string, targetFolder string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fileName := GetFileNameFromURL(rawURL)
	basePath := filepath.Join(currentDir, targetFolder, fileName)

	
	if !fileExists(basePath) {
		return basePath, nil
	}

	
	ext := filepath.Ext(fileName)
	nameWithoutExt := strings.TrimSuffix(fileName, ext)

	
	counter := 1
	for {
		
		newFileName := fmt.Sprintf("%s_%d%s", nameWithoutExt, counter, ext)
		newFullPath := filepath.Join(currentDir, targetFolder, newFileName)

		if !fileExists(newFullPath) {
			return newFullPath, nil
		}
		counter++
	}
}
