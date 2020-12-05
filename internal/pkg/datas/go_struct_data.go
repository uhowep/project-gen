package datas

import (
	"fmt"
	"os"
	"project-gen/internal/pkg/models"
	"reflect"
	"strings"
)

type GoStructData struct {
	GoStruct interface{}
}

func NewGoStructData(goStruct interface{}) *GoStructData {
	return &GoStructData{
		GoStruct: goStruct,
	}
}

func (gsd *GoStructData) Parse2DataModel() *models.DataModel {
	getType := reflect.TypeOf(gsd.GoStruct)
	getValue := reflect.ValueOf(gsd.GoStruct)
	if getType.Kind() != reflect.Struct {
		fmt.Println(fmt.Sprintf("parse go struct failed because it is non-struct type which is %s", reflect.TypeOf(gsd.GoStruct).String()))
		os.Exit(1)
	}
	// parse to get struct name
	structName := gsd.getStructName(getType)
	// parse to get struct fields
	fieldMaps := make(map[string]*models.FieldModel)
	for i := 0; i < getValue.NumField(); i++ {
		fieldName := getType.Field(i).Name
		field := models.NewFieldModel(fieldName, getType.Field(i).Type)
		fieldMaps[fieldName] = field
	}
	// return
	return models.NewDataModel(structName, fieldMaps)
}

// parse to get the struct name used by reflect.Type.String()
func (gsd *GoStructData) getStructName(structType reflect.Type) string {
	typeWholeName := structType.String()
	splitNames := strings.Split(typeWholeName, ".")
	name := splitNames[len(splitNames)-1]
	return name
}
