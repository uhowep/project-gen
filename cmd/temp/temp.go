package main

import (
	"flag"
	"fmt"
	"os"
	"project-gen/assets/gomodels"
	"project-gen/internal/generator"
	"project-gen/internal/pkg/common"
	"project-gen/internal/pkg/datas"
	"project-gen/internal/pkg/models"
	"project-gen/internal/pkg/structures"
)

var (
	projectSelf   string
	projectTarget string
)

func init() {
	flag.StringVar(&projectSelf, "i", "./", "self project path")
	flag.StringVar(&projectTarget, "o", "./", "output project target path. (default is `./`)")
	flag.Parse()
}

func main() {
	common.ProjectSelfPath = projectSelf
	common.ProjectTargetPath = projectTarget
	// init path config
	pathConfig := models.NewPathConfig("assets/path_config.json")
	// project structure
	stList := structures.NewStructures().Parse2StructureList()
	if len(stList.Paths) == 0 {
		fmt.Println(fmt.Sprintf("StructureList is empty"))
		os.Exit(1)
		return
	}
	// generate
	gen := generator.NewGenerator()
	// generate project struct
	if genErr := gen.GenerateProjectStructure(stList); genErr != nil {
		fmt.Println(fmt.Sprintf("generate structure occur error: %s which used %v", genErr.Error(), stList))
		os.Exit(1)
		return
	}
	// generate model
	ta := datas.NewGoStructData(gomodels.TestA{
		MyFieldA1: "",
		MyFieldA2: "",
		MyFieldA3: "",
	})
	tb := datas.NewGoStructData(gomodels.TestB{
		MyFieldB1: "",
		MyFieldB2: "",
		MyFieldB3: "",
	})
	dataModels := []*models.DataModel{ta.Parse2DataModel(), tb.Parse2DataModel()}
	if genErr := gen.GenerateModels(pathConfig.ModelPath, dataModels); genErr != nil {
		fmt.Println(fmt.Sprintf("generate model occur error: %s", genErr.Error()))
		os.Exit(1)
		return
	}
	// return
	return
}
