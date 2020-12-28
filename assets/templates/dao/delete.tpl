{{ define "dao_delete.tpl" }}

{{ $model = .Model }}
func ({{.VarName}} *{{.DaoName}}) Delete{{$model.ModelName}}(ctx context.Context, record *{{$model.ModuleName}}.{{$model.ModelName}}) error {
    return {{.VarName}}.DB().Delete(&{{$model.ModuleName}}.{{$model.ModelName}}{}, "{{$model.DeleteTplVar.GetWhereConditionStr}}", {{$model.DeleteTplVar.GetWhereArgStr}}).Error
}
{{ end }}