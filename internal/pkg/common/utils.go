package common

import (
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
