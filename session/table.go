package session

import (
	"EORM/log"
	"EORM/schema"
	"fmt"
	"reflect"
	"strings"
)

// Model 设置会话数据表
func (s *Session) Model(value interface{}) *Session {
	// 若未设置表 或者传入了不同的表才更新当前会话的表
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// RefTable 获取数据表
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("未设置 Model")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ", ")
	_, err := s.Raw(fmt.Sprintf("create table %s (%s)", s.RefTable().Name, desc)).Exec()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("drop table if exists %s", s.RefTable().Name)).Exec()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *Session) HasTable() bool {
	tableName := s.RefTable().Name
	sql, values := s.dialect.TableExistSQL(tableName)
	row := s.Raw(sql, values...).Query()
	tmp := ""
	err := row.Scan(&tmp)
	if err != nil {
		log.Error(err)
		return false
	}
	return tmp == tableName
}
