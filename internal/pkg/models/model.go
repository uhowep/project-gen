package models

import "reflect"

type DataModel struct {
	ModelName string
	Fields map[string]FieldModel
}

type FieldModel struct {
	Name string
	Type reflect.Type
	Addable bool
	Editable bool
	Queryable bool
	Deletable bool
	DefaultValue interface{}
}