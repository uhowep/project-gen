package models

import "reflect"

type DataModel struct {
	ModuleName string
	ModelName  string
	Fields     map[string]*FieldModel
}

type FieldModel struct {
	FieldName string
	DbName    string
	Type      reflect.Type
	Addable   bool
	Updatable bool
	Queryable bool
	Deletable bool
	AddHook   func(fields map[string]*FieldModel)
	ReadHook  func(fields map[string]*FieldModel)
}

func NewFieldModel(fieldName, dbName string, fieldType reflect.Type) *FieldModel {
	return &FieldModel{
		FieldName: fieldName,
		DbName:    dbName,
		Type:      fieldType,
	}
}

func NewDataModel(moduleName, modelName string, fields map[string]*FieldModel) *DataModel {
	return &DataModel{
		ModuleName: moduleName,
		ModelName:  modelName,
		Fields:     fields,
	}
}
