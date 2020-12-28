{{ define "dao_del_by_primary_key.tpl" }}

{{ $model = .Model }}
func ({{.VarName}} *{{.DaoName}}) Delete{{$model.ModelName}}ByPrimary(ctx context.Context, primaryKey {{.ModelPrimaryKeyType}}) error {
	return db.DB().Where("{{.ModelPrimaryKeyName}}=?",primaryKey).Delete(&{{$model.ModuleName}}.{{$model.ModelName}}).Error
}

{{ end }}
