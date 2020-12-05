{{ define "dao_add.tpl" }}

func (db *{{.ModuleName}}) Add{{.ModelName}}(ctx context.Context, record *{{.ModelPackageName}}.{{.ModelName}}) (result *{{ ModelPackageName }}.{{.ModelName}}, err error) {
    if err := db.DB().Model(&{{.ModelPackageName}}.{{.ModelName}}{}).Create(record).Error; err != nil {
        return nil, err
    }
    // return
    return record, nil
}

{{ end }}