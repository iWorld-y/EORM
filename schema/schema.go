package schema

import (
	"EugeneORM/dialect"
	"go/ast"
	"reflect"
)

// Field 代表数据表中一个字段
type Field struct {
	Name string // 字段名
	Type string // 字段类型
	Tag  string // 字段约束条件
}

// Schema 代表一张数据表
type Schema struct {
	Model      interface{}       // 被映射的对象
	Name       string            // 表名
	Fields     []*Field          // 字段
	FieldNames []string          // 包含所有的字段名
	fieldMap   map[string]*Field // 记录字段名 => 字段的映射关系
}

// GetField 获取字段名对应的数据表中的字段
func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// Parse 将任意的 Golang 对象解析为 Schema 实例
// 它的返回值是一个指向 Schema 结构体的指针。
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type() // 获取目标对象 dest 的类型信息

	schema := &Schema{ // 初始化一个 Schema 结构体实例
		Model:    dest,                    // Model 字段是传入的 Golang 对象
		Name:     modelType.Name(),        // Name 字段是传入对象的类型名
		fieldMap: make(map[string]*Field), // fieldMap 是一个从字段名到字段信息的映射
	}

	for i := 0; i < modelType.NumField(); i++ { // 遍历传入对象的所有字段
		p := modelType.Field(i) // 获取第 i 个字段的信息

		if !p.Anonymous && ast.IsExported(p.Name) { // 如果字段不是匿名字段且是导出的字段 (即变量以大写字母开头)

			field := &Field{ // 初始化一个 Field 结构体实例
				Name: p.Name,                                              // Name 字段是字段名
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))), // 利用 dialect.Dialect 接口获取字段的数据类型
			}

			if v, ok := p.Tag.Lookup("EugeneORM"); ok { // 如果字段有 "EugeneORM" 标签，将其值赋给 Field 结构体的 Tag 字段
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)          // 将字段添加到 Schema 的字段列表中
			schema.FieldNames = append(schema.FieldNames, p.Name) // 将字段名添加到 Schema 的字段名列表中
			schema.fieldMap[p.Name] = field                       // 将字段添加到字段映射表中
		}
	}

	return schema // 返回解析得到的 Schema 实例
}
