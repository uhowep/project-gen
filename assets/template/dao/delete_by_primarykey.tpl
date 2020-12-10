{{ define "dao_del.tpl" }}

func (db *{{.ModuleName}}) Delete{{.ModelName}}(ctx context.Context, primaryKey {{.ModelPrimaryKeyType}}) error {
	return db.DB().Where("{{.ModelPrimaryKeyName}}=?",primaryKey).Delete(&{{.ModuleName}}.{{.ModelName}}).Error
}

{{ end }}
