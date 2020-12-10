package models

import (
	"fmt"
	"strings"
)

type DaoModel struct {
	OriginDataModel *DataModel
	AddTplVar       *AddTpl
	UpdateTplVar    *UpdateTpl
	DeleteTplVar    *DeleteTpl
	QueryTplVar     *QueryTpl
}

func NewDaoModel(originDataModel *DataModel) *DaoModel {
	dao := &DaoModel{
		OriginDataModel: originDataModel,
		AddTplVar: &AddTpl{
			ModuleName: originDataModel.ModuleName,
			ModelName:  originDataModel.ModelName,
		},
		UpdateTplVar: &UpdateTpl{
			ModuleName: originDataModel.ModuleName,
			ModelName:  originDataModel.ModelName,
		},
		DeleteTplVar: &DeleteTpl{
			ModuleName: originDataModel.ModuleName,
			ModelName:  originDataModel.ModelName,
		},
		QueryTplVar: &QueryTpl{
			ModuleName: originDataModel.ModuleName,
			ModelName:  originDataModel.ModelName,
		},
	}
	// parse
	updateWheres := make([]string, 0)
	deleteWheres := make([]string, 0)
	queryWheres := make([]string, 0)
	for fieldName, field := range originDataModel.Fields {
		structFieldName := fmt.Sprintf("%s.%s", originDataModel.ModelName, fieldName)
		dbFieldName := field.DbName
		if field.Updatable {
			updateWheres = append(updateWheres, dbFieldName)
			dao.UpdateTplVar.WhereArgs = append(dao.UpdateTplVar.WhereArgs, structFieldName)
		}
		if field.Deletable {
			deleteWheres = append(deleteWheres, dbFieldName)
			dao.DeleteTplVar.WhereArgs = append(dao.DeleteTplVar.WhereArgs, structFieldName)
		}
		if field.Queryable {
			queryWheres = append(queryWheres, dbFieldName)
			dao.QueryTplVar.WhereArgs = append(dao.QueryTplVar.WhereArgs, structFieldName)
		}
	}
	dao.UpdateTplVar.WhereCondition = strings.Join(updateWheres, " AND ")
	dao.DeleteTplVar.WhereCondition = strings.Join(deleteWheres, " AND ")
	dao.QueryTplVar.WhereCondition = strings.Join(queryWheres, " AND ")
	// return
	return dao
}

type AddTpl struct {
	ModuleName string
	ModelName  string
}

type UpdateTpl struct {
	ModuleName     string
	ModelName      string
	WhereCondition string
	WhereArgs      []string
}

type DeleteTpl struct {
	ModuleName     string
	ModelName      string
	WhereCondition string
	WhereArgs      []string
}

type QueryTpl struct {
	ModuleName     string
	ModelName      string
	WhereCondition string
	WhereArgs      []string
}

