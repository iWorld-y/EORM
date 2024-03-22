package schema

// Field2 一个字段
type Field2 struct {
	Name string
	Type string
	Tag  string
}
type Schema2 struct {
	Model      interface{}
	TableName  string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}
