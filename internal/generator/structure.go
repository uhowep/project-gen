package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"project-gen/internal/pkg/common"
	"project-gen/internal/pkg/models"
)

type Generator struct {

}

func NewGenerator() *Generator {
	return &Generator{}
}

func (gen *Generator) GenerateProjectStructure(stList *models.StructureList) error {
	// generate
	for _, path := range stList.Paths {
		dir, fileStr := filepath.Split(path)
		// generate directory
		if dir != "" && (dir[0] == '.' || dir[0] == '/') {
			return fmt.Errorf("the path %s cannot start `.` or start with `/`", path)
		} else if dir != "" {
			dir = common.GetTargetFilePath(dir)
			if mkdirErr := os.MkdirAll(dir, 0755); mkdirErr != nil {
				return fmt.Errorf("mkdir occur error %s", mkdirErr.Error())
			}
		}
		// generate file
		if fileStr != "" {
			wholeFileStr := common.GetTargetFilePath(path)
			file, createErr := os.OpenFile(wholeFileStr, os.O_CREATE, 0644)
			if createErr != nil {
				return fmt.Errorf("create file %s occur error %s", fileStr, createErr.Error())
			}
			file.Close()
		}
	}
	// return
	return nil
}