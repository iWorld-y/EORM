package dialect

import (
	"reflect"
)

type Dialect interface {
	DataTypeOf(typ reflect.Value) string                    // 将 Golang 类型转为指定数据库的数据类型
	TableExistSQL(tableName string) (string, []interface{}) // 返回某个表是否存在的 SQL 语句, 参数为表名 table
}

var dialectsMap = map[string]Dialect{}

// RegisterDialect 注册 Dialect 实例
// 若新增对某数据库的支持, 调用 RegisterDialect 即可注册到全局
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

// GetDialect 获取 Dialect 实例
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return dialect, ok
}
