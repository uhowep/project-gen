package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"project-gen/internal/pkg/common"
)

type PathConfig struct {
	ModelPath string `json:"model_path"`
	DaoPath   string `json:"dao_path"`
	ApiPath   string `json:"api_path"`
}

func NewPathConfig(filePath string) *PathConfig {
	if filePath == "" {
		fmt.Println("get target path config failed with empty filePath parameter")
		os.Exit(1)
	}
	// load path config from json file
	filePath = common.GetSelfFilePath(filePath)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(fmt.Sprintf("get %s content failed with %s", filePath, err.Error()))
		os.Exit(1)
	}
	// parse
	var pathConfig PathConfig
	if unmarshalErr := json.Unmarshal(content, &pathConfig); unmarshalErr != nil {
		fmt.Println(fmt.Sprintf("parse %s file to json failed with %s", filePath, unmarshalErr.Error()))
		os.Exit(1)
	}
	// return
	return &pathConfig
}
