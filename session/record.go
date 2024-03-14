package session

import (
	"EugeneORM/clause"
	"EugeneORM/log"
)

func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0) //创建一个空的接口类型切片，该切片不包含任何元素。
	for _, value := range values {
		table := s.Model(value).RefTable()                             // 根据穿入的对象更新数据表
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)      // 构造 INSERT 子句
		recordValues = append(recordValues, table.RecordValues(value)) // 合并整理好的数据
	}

	s.clause.Set(clause.VALUES, recordValues...)              // 生成 VALUE 子句
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES) // 把子句合并为 SQL 语句
	result, err := s.Raw(sql, vars...).Exec()                 // 执行 SQL
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}
