package schema

import (
	"EORM/dialect"
	"go/ast"
	"reflect"
)

// Field 数据表的一个字段
type Field struct {
	Name string // 字段名
	Type string // 字段类型
	Tag  string // 约束条件
}

// Schema 一个数据表
type Schema struct {
	Model      interface{}       // 被映射对象
	Name       string            // 表名
	Fields     []*Field          // 表的字段
	FieldNames []string          // 所有字段名
	fieldMap   map[string]*Field // 字段名到字段信息的 map
}

func (s *Schema) GetField(name string) *Field {
	return s.fieldMap[name]
}
func Parse(model interface{}, dialect dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(model)).Type()
	schema := &Schema{
		Model:    model,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		modelField := modelType.Field(i)
		if !modelField.Anonymous && ast.IsExported(modelField.Name) {
			field := &Field{
				Name: modelField.Name,
				Type: dialect.DataTypeOf(reflect.Indirect(reflect.New(modelField.Type))),
			}
			if v, ok := modelField.Tag.Lookup("EORM"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, modelField.Name)
			schema.fieldMap[modelField.Name] = field
		}
	}
	return schema
}
