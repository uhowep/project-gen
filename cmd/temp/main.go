package main

import (
	"flag"
	"fmt"
	"os"
	"project-gen/internal/generator"
	"project-gen/internal/pkg/common"
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
	// project structure
	stList := structures.NewStructures().Parse2StructureList()
	if len(stList.Paths) == 0 {
		fmt.Println(fmt.Sprintf("StructureList is empty"))
		os.Exit(1)
		return
	}
	// generate
	gen := generator.NewGenerator()
	if genErr := gen.GenerateProjectStructure(stList); genErr != nil {
		fmt.Println(fmt.Sprintf("generate structure occur error %s which used %v", genErr.Error(), stList))
		os.Exit(1)
		return
	}
	// return
	return
}
