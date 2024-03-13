package EugeneORM

import (
	"EugeneORM/log"
	"EugeneORM/session"
	"database/sql"
)

// Engine 负责连接/测试数据库等前期准备工作, 以及关闭连接等后期收尾工作
type Engine struct {
	db *sql.DB
}

// NewEngine 返回一个 Engine 对象, 并测试数据库连接
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// ping 一下确保数据库连接存活
	if err = db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}
	e = &Engine{db: db}
	log.Info("数据库连接成功")
	return e, nil
}

// Close 关闭数据库连接
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("关闭数据库连接失败")
	}
	log.Info("数据库连接已关闭")
}

// NewSession 使用 engine.db 创建一个 Session 对象
// 即创建一个会话, 以便于数据库交互
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
