package session

import (
	"EORM/log"
	"database/sql"
	"strings"
)

type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVals []interface{}
}

func New(db *sql.DB) *Session {
	return &Session{db: db}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVals = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVals = append(s.sqlVals, values...)
	return s
}
func (s *Session) checkSql() *Session {
	sqlStr := s.sql.String()
	if !strings.HasSuffix(sqlStr, ";") {
		s.sql.WriteString(";")
	}
	return s
}
func (s *Session) Exec() (ret sql.Result, err error) {
	defer s.Clear()
	s.checkSql()
	log.Info(s.sql.String(), s.sqlVals)
	if ret, err = s.DB().Exec(s.sql.String(), s.sqlVals...); err != nil {
		log.Error(err)
	}
	return ret, err
}

func (s *Session) Query() *sql.Row {
	defer s.Clear()
	s.checkSql()
	log.Info(s.sql.String(), s.sqlVals)
	return s.DB().QueryRow(s.sql.String(), s.sqlVals...)
}

func (s *Session) QueryRows() (*sql.Rows, error) {
	defer s.Clear()
	s.checkSql()
	var rows = &sql.Rows{}
	var err error
	log.Info(s.sql.String(), s.sqlVals)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVals...); err != nil {
		log.Error(err)
	}
	return rows, err
}
