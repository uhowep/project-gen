package {{.ModuleName}}

type {{.ModelName}} struct {
    {{- range $fieldName, $field := .Fields}}
    {{$fieldName}} {{$field.Type}} `json:"{{$field.DbName}}" form:"{{$field.DbName}}`
    {{- end}}
}