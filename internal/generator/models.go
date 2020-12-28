package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"project-gen/internal/pkg/common"
	"project-gen/internal/pkg/models"
	"text/template"
)

func (gen *Generator) GenerateModels(targetPath string, dataModels []*models.DataModel) error {
	//todo: set path to absolute path outside and drop the code of common.GetTargetFilePath
	targetPath = common.GetTargetFilePath(targetPath)
	// get template
	templateFilePath := common.GetSelfFilePath("assets/templates/model/model.tpl")
	templateBytes, templateErr := ioutil.ReadFile(templateFilePath)
	if templateErr != nil {
		return fmt.Errorf("get model template content %s for generate model occur error: %s", templateFilePath, templateErr.Error())
	}
	modelTpl, parseErr := template.New("model.tpl").Parse(string(templateBytes))
	if parseErr != nil {
		return fmt.Errorf("parse model template %s for generate model occur error: %s", templateFilePath, parseErr.Error())
	}
	// generate
	for _, dataModel := range dataModels {
		// create file
		fileName := fmt.Sprintf("%s.go", common.Camel2Case(dataModel.ModelName))
		filePath := filepath.Join(targetPath, fileName)
		file, createErr := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if createErr != nil {
			return fmt.Errorf("create model file for generate model %s occur error: %s", filePath, createErr.Error())
		}
		// write content
		if generateErr := modelTpl.Execute(file, dataModel); generateErr != nil {
			return fmt.Errorf("generate model file %s for generate model occur error: %s", filePath, generateErr.Error())
		}
		file.Close()
		fmt.Println(fmt.Sprintf("generate %s success for generate model", filePath))
	}
	// return
	return nil
}

