package common

import (
	"io/ioutil"
	"os"
	"path"
)

// get self project file path with relative path
func GetSelfFilePath(relativePath string) string {
	return path.Join(ProjectSelfPath, relativePath)
}

// get target project file path with relative path
func GetTargetFilePath(relativePath string) string {
	return path.Join(ProjectTargetPath, relativePath)
}

// get file content
func GetFileContent(path string) ([]byte, error) {
	file, openErr := os.OpenFile(path, os.O_RDONLY, 0644)
	if openErr != nil {
		return nil, openErr
	}
	defer file.Close()
	content, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		return nil, readErr
	}
	return content, nil
}