package sql2struct

const strutTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	strutTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*TableColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{strutTpl: strutTpl}
}
