{{ define "dao_query.tpl" }}

{{ $model = .Model }}
func ({{.VarName}} *{{.DaoName}}) Query{{$model.ModelName}}(ctx context.Context, record *{{$model.ModuleName}}.{{$model.ModelName}}) (*{{$model.ModuleName}}.{{$model.ModelName}}, error) {
    var record {{$model.ModuleName}}.{{$model.ModelName}}
    err := {{.VarName}}.DB().Model(&{{$model.ModuleName}}.{{$model.ModelName}}{}).Where("{{$model.QueryTplVar.GetWhereConditionStr}}", {{$model.QueryTplVar.GetWhereArgStr).First(&{{$model.ModuleName}}.{{$model.ModelName}}{}).Error
    if err != nil {
        return nil, err
    }
    // return
    return &record, nil
}
{{ end }}
