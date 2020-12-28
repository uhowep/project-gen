package datas

import (
	"fmt"
	"os"
	"project-gen/internal/pkg/common"
	"project-gen/internal/pkg/models"
	"reflect"
	"strconv"
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
		fmt.Println(fmt.Sprintf("parse go struct failed because it is non-struct type which is %s", getType.Kind()))
		os.Exit(1)
	}
	// parse to get struct name
	structName := gsd.getStructName(getType)
	// parse to get struct fields
	fieldMaps := make(map[string]*models.FieldModel)
	for i := 0; i < getValue.NumField(); i++ {
		fieldType := getType.Field(i)
		dbName, existTagField := fieldType.Tag.Lookup("field")
		if !existTagField {
			dbName = common.Camel2Case(fieldType.Name)
		}
		field := models.NewFieldModel(fieldType.Name, dbName, fieldType.Type)
		// tags
		genTags := fieldType.Tag.Get("gen")
		tags := strings.Split(genTags, ";")
		for _, tag := range tags {
			kv := strings.Split(tag, ":")
			if len(kv) != 2 {
				continue
			}
			switch kv[0] {
			case "add":
				if val, convertErr := strconv.ParseBool(kv[1]); convertErr == nil && val == true {
					field.Addable = true
				}
			case "update":
				if val, convertErr := strconv.ParseBool(kv[1]); convertErr == nil && val == true {
					field.Updatable = true
				}
			case "delete":
				if val, convertErr := strconv.ParseBool(kv[1]); convertErr == nil && val == true {
					field.Deletable = true
				}
			case "query":
				if val, convertErr := strconv.ParseBool(kv[1]); convertErr == nil && val == true {
					field.Queryable = true
				}
			}
		}
		// fill
		fieldMaps[fieldType.Name] = field
	}
	// return
	return models.NewDataModel(common.ModelTargetPkgName, structName, fieldMaps)
}

// parse to get the struct name used by reflect.Type.String()
func (gsd *GoStructData) getStructName(structType reflect.Type) string {
	typeWholeName := structType.String()
	splitNames := strings.Split(typeWholeName, ".")
	name := splitNames[len(splitNames)-1]
	return name
}
