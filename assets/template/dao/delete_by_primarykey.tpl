{{ define "dao_del.tpl" }}

func (db *{{.ModuleName}}) Delete{{.ModelName}}(ctx context.Context, primaryKey {{.ModelPrimaryKeyType}}) error {
    return db.DB().Delete(&{{.ModelPackageName}}.{{.ModelName}}, "{{.ModelPrimaryKeyName}}=?", primaryKey).Error
}

{{ end }}