package schema

import (
	"EORM/dialect"
	"EORM/utils"
	"testing"
)

type User struct {
	Name string `EORM:"PRIMARY KEY"`
	Age  int
}

func TestParse(t *testing.T) {
	testDialect, _ := dialect.GetDialect(utils.SqliteDBDriverName)
	schema := Parse(&User{}, testDialect)
	if schema.Name != "User" || len(schema.FieldNames) != 2 {
		t.Fatal("解析失败")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("获取 tag 失败")
	}
}
