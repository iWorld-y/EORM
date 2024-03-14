package clause

import "strings"

type Type int

// Clause 定义子句, sql 中存放生成的 SQL 语句, sqlVars 中存放数据
type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]interface{}
}

// 操作类型
const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
)

// Set 根据 Type(操作类型) 调用对应的 generator, 生成该子句的 SQL
func (c *Clause) Set(name Type, vars ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}
	// name 是指操作类型, 在 clause/generator.go 中的 init() 函数指明了
	// 每一种操作类型对应的 clause 生成函数
	// 则假设当前操作类型为 select, 那么这句话相当于:
	// sql, vars := _insert(vars...)
	sql, vars := generators[name](vars...)
	c.sql[name] = sql
	c.sqlVars[name] = vars
}

// Build 根据传入的操作顺序生成最终的 SQL
func (c *Clause) Build(orders ...Type) (string, []interface{}) {
	var sqls []string
	var vars []interface{}
	for _, order := range orders {
		if sql, ok := c.sql[order]; ok {
			sqls = append(sqls, sql)
			vars = append(vars, c.sqlVars[order]...)
		}
	}
	return strings.Join(sqls, " "), vars
}
