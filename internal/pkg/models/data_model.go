package models

import "reflect"

type DataModel struct {
	ModelName string
	Fields    map[string]*FieldModel
}

type FieldModel struct {
	Name            string
	Type            reflect.Type
	Addable         bool
	Editable        bool
	SimpleQueryable bool
	ListQueryable   bool
	Deletable       bool
	DefaultValue    interface{}
}

func NewFieldModel(name string, fieldType reflect.Type) *FieldModel {
	return &FieldModel{
		Name: name,
		Type: fieldType,
	}
}

func NewDataModel(name string, fields map[string]*FieldModel) *DataModel {
	return &DataModel{
		ModelName: name,
		Fields:    fields,
	}
}
