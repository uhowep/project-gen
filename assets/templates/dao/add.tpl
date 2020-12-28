{{ define "dao_add.tpl" }}

{{ $model = .Model }}
func ({{.VarName}} *{{.DaoName}}) Add{{$model.ModelName}}(ctx context.Context, record *{{$model.ModuleName}}.{{$model.ModelName}}) (result *{{$model.ModuleName}}.{{$model.ModelName}}, err error) {
    if err := {{.VarName}}.DB().Model(&{{$model.ModuleName}}.{{$model.ModelName}}{}).Create(record).Error; err != nil {
        return nil, err
    }
    // return
    return record, nil
}
{{ end }}