package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string                    // 将 Go 的类型映射为数据库的数据类型
	TableExistSQL(tableName string) (string, []interface{}) //返回某个表是否存在的 SQL 语句, 参数为表名 table
}

// RegisterDialect 注册方言实例
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

// GetDialect 获取方言实例
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return dialect, ok
}
