package structures

import (
	"encoding/json"
	"fmt"
	"os"
	"project-gen/internal/pkg/common"
	"project-gen/internal/pkg/models"
)

type Structures struct {
	templateContent []byte
}

func NewStructures() *Structures {
	templatePath := common.GetSelfFilePath("assets/project-structure-tpl.json")
	content, err := common.GetFileContent(templatePath)
	if err != nil {
		fmt.Println("get %s content failed with %s", templatePath, err.Error())
		os.Exit(1)
	}
	// return
	return &Structures{
		templateContent: content,
	}
}

func (st *Structures) Parse2StructureList() *models.StructureList {
	paths := make([]string, 0)
	if unmarshalErr := json.Unmarshal(st.templateContent, &paths); unmarshalErr != nil {
		fmt.Println(fmt.Sprintf("parse structure failed because unmarshal file occur error: %s", unmarshalErr))
		os.Exit(1)
	}
	return models.NewStructureList(paths)
}
