package EugeneORM

import (
	"EugeneORM/dialect"
	"EugeneORM/log"
	"EugeneORM/session"
	"database/sql"
)

// Engine 负责连接/测试数据库等前期准备工作, 以及关闭连接等后期收尾工作
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

// NewEngine 创建 Engine 实例时，获取 driver 对应的 dialect, 并测试数据库连接
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

	// 确保指定的 dialect 存在
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("方言 %s 不存在\n", driver)
	}

	e = &Engine{db: db, dialect: dial}
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
	return session.New(engine.db, engine.dialect)
}
