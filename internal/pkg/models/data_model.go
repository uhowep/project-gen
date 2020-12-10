package models

import "reflect"

type DataModel struct {
	ModuleName string
	ModelName  string
	Fields     map[string]*FieldModel
}

type FieldModel struct {
	StructName string
	DbName     string
	Type       reflect.Type
	Addable    bool
	Updatable  bool
	Queryable  bool
	Deletable  bool
	AddHook    func(fields map[string]*FieldModel)
	ReadHook   func(fields map[string]*FieldModel)
}

func NewFieldModel(structName, dbName string, fieldType reflect.Type) *FieldModel {
	return &FieldModel{
		StructName: structName,
		DbName:     dbName,
		Type:       fieldType,
	}
}

func NewDataModel(name string, fields map[string]*FieldModel) *DataModel {
	return &DataModel{
		ModelName: name,
		Fields:    fields,
	}
}
