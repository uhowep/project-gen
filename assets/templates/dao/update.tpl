{{ define "dao_update.tpl" }}

{{ $model = .Model }}
func ({{.VarName}} *{{.DaoName}}) Update{{$model.ModelName}}(ctx context.Context, record *{{$model.ModuleName}}.{{$model.ModelName}}) error {
    return {{.VarName}}.DB().Where("{{$model.UpdateTplVar.GetWhereConditionStr}}", {{$model.UpdateTplVar.GetWhereArgStr}}).Update(&{{$model.ModuleName}}.{{$model.ModelName}}{}).Error
}
{{ end }}