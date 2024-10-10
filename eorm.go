package EORM

import (
	"EORM/log"
	"database/sql"
)

type Engine struct {
	db *sql.DB
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
	log.Info("数据库已连接")
	return engine, nil
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("数据库关闭失败")
	}
	log.Info("数据库已关闭")
}
