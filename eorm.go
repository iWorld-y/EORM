package EORM

import (
	"EORM/dialect"
	"EORM/log"
	"EORM/session"
	"database/sql"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

// NewEngine 连接数据库, 并且通过 ping 测试连接
func NewEngine(driver, source string) (*Engine, error) {
	var engine = &Engine{}
	var err error
	engine.db, err = sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if err = engine.db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}

	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s 获取失败", driver)
		return nil, err
	}
	engine.dialect = dial
	log.Info("数据库已连接")
	return engine, nil
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("数据库关闭失败")
	}
	log.Info("数据库已关闭")
}
func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}
