package dialect

import (
	"EORM/utils"
	"fmt"
	"reflect"
	"time"
)

type sqlite3 struct{}

var _ Dialect = &sqlite3{}

// 包在第一次加载时，会将 sqlite3 的 dialect 自动注册到全局
func init() {
	RegisterDialect(utils.SqliteDBDriverName, &sqlite3{})
}

// DataTypeOf 将 go 数据类型映射到 SQLite3 的数据类型
func (s *sqlite3) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

// TableExistSQL 获取查询某表是否存在的 SQL
func (s *sqlite3) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "select name from sqlite_master where type='table' and name=?", args
}
