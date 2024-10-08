package session

import (
	"EORM/clause"
	"EORM/dialect"
	"EORM/log"
	"EORM/schema"
	"database/sql"
	"strings"
)

// Session 负责与数据库的交互
type Session struct {
	db       *sql.DB         // 使用 sql.Open() 方法连接数据库成功后返回的指针
	dialect  dialect.Dialect // 所使用的数据库方言
	refTable *schema.Schema  // 需要读写的数据表
	clause   clause.Clause   // 指定操作 + 对应数据 生成的子句
	sql      strings.Builder // 用于拼接 SQL 语句及其中的占位符
	sqlVars  []interface{}   // 占位符所对应的数据
}

// New 创建一个新的 Session 对象，接受一个 *sql.DB 指针作为参数，并返回该对象的指针
func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{db: db,
		dialect: dialect,
	}
}

// Clear 重置 Session 对象的 sql 字段, sqlVars 字段和子句字段
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

// DB 返回 Session 对象的 db 字段，即 *sql.DB 指针
func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw 拼接 SQL 语句和其中的占位符，并将占位符对应的数据存储在 sqlVars 中
// 接受一个 sql 字符串作为参数，并可以接受任意数量的占位符所对应的数据
// 返回 Session 对象的指针，以支持链式调用
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// Exec 使用 sqlVars 执行原始 SQL
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow 从数据库读取一条记录
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows 从数据库读取一组数据
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return rows, err
}
