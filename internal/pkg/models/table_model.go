package models

import (
	"fmt"
	"strings"
)

type TableModel struct {
	OriginDataModel *DataModel
	AddTplVar       AddTpl
	UpdateTplVar    UpdateTpl
	DeleteTplVar    DeleteTpl
	QueryTplVar     QueryTpl
}

func NewTableModel(originDataModel *DataModel) *TableModel {
	dao := &TableModel{
		OriginDataModel: originDataModel,
		AddTplVar: AddTpl{
			ModuleName: originDataModel.ModuleName,
			ModelName:  originDataModel.ModelName,
		},
		UpdateTplVar: UpdateTpl{
			ModuleName: originDataModel.ModuleName,
			ModelName:  originDataModel.ModelName,
		},
		DeleteTplVar: DeleteTpl{
			ModuleName: originDataModel.ModuleName,
			ModelName:  originDataModel.ModelName,
		},
		QueryTplVar: QueryTpl{
			ModuleName: originDataModel.ModuleName,
			ModelName:  originDataModel.ModelName,
		},
	}
	// parse
	for fieldName, field := range originDataModel.Fields {
		structFieldName := fmt.Sprintf("%s.%s", originDataModel.ModelName, fieldName)
		dbFieldName := field.DbName
		if field.Updatable {
			dao.UpdateTplVar.WhereArgs = append(dao.UpdateTplVar.WhereArgs, structFieldName)
			dao.UpdateTplVar.WhereConditions = append(dao.UpdateTplVar.WhereConditions, dbFieldName)
		}
		if field.Deletable {
			dao.DeleteTplVar.WhereArgs = append(dao.DeleteTplVar.WhereArgs, structFieldName)
			dao.DeleteTplVar.WhereConditions = append(dao.DeleteTplVar.WhereConditions, dbFieldName)
		}
		if field.Queryable {
			dao.QueryTplVar.WhereArgs = append(dao.QueryTplVar.WhereArgs, structFieldName)
			dao.QueryTplVar.WhereConditions = append(dao.QueryTplVar.WhereConditions, dbFieldName)
		}
	}
	// return
	return dao
}

type AddTpl struct {
	ModuleName string
	ModelName  string
}

type UpdateTpl struct {
	ModuleName      string
	ModelName       string
	WhereConditions []string
	WhereArgs       []string
}

type DeleteTpl struct {
	ModuleName      string
	ModelName       string
	WhereConditions []string
	WhereArgs       []string
}

type QueryTpl struct {
	ModuleName      string
	ModelName       string
	WhereConditions []string
	WhereArgs       []string
}

func (update UpdateTpl) GetWhereConditionStr() string {
	for ind, condition := range update.WhereConditions {
		update.WhereConditions[ind] = fmt.Sprintf("%s=?", condition)
	}
	return strings.Join(update.WhereConditions, " and ")
}

func (update UpdateTpl) GetWhereArgStr() string {
	return strings.Join(update.WhereArgs, ", ")
}

func (delete DeleteTpl) GetWhereConditionStr() string {
	for ind, condition := range delete.WhereConditions {
		delete.WhereConditions[ind] = fmt.Sprintf("%s=?", condition)
	}
	return strings.Join(delete.WhereConditions, " and ")
}

func (delete DeleteTpl) GetWhereArgStr() string {
	return strings.Join(delete.WhereArgs, ", ")
}

func (query QueryTpl) GetWhereConditionStr() string {
	for ind, condition := range query.WhereConditions {
		query.WhereConditions[ind] = fmt.Sprintf("%s=?", condition)
	}
	return strings.Join(query.WhereConditions, " and ")
}

func (query QueryTpl) GetWhereArgStr() string {
	return strings.Join(query.WhereArgs, ", ")
}
