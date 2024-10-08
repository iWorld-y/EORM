// Package session 操作数据库表
package session

import (
	"EORM/log"
	"EORM/schema"
	"fmt"
	"reflect"
	"strings"
)

// Model 根据传入的对象决定是否要生成数据表 refTable
func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil ||
		reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		// 若当前 session 的数据表为空
		// 或者传入对象的类型与当前 session 的数据表对应对象的类型不一致
		// 即传入的结构体发生了变化
		// 则更新数据表 refTable
		// 这样做是因为解析操作比较费时, 所以若传入的结构体未发生变化就不更新 refTable
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// RefTable 获取 refTable 的值, 若 refTable 未被赋值则报错
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("模型未设置")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.RefTable()
	var colums []string
	for _, field := range table.Fields {
		colums = append(colums, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(colums, ",")
	_, err := s.Raw(fmt.Sprintf("create table %s (%s);", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("drop table if exists %s;", s.RefTable().Name)).Exec()
	return err
}

func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var temp string
	_ = row.Scan(&temp)
	return temp == s.RefTable().Name
}
