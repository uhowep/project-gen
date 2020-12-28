package models

import (
	"project-gen/internal/pkg/common"
)

type DaoModel struct {
	DaoName string
	VarName string
	Tables  []*TableModel
}

func NewDaoModel(daoName string, tables []*TableModel) *DaoModel {
	varName := common.Lcfirst(daoName)
	if varName == "" {
		varName = daoName
	}
	return &DaoModel{
		DaoName: daoName,
		VarName: varName,
		Tables:  tables,
	}
}
