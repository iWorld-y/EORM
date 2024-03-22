package session

import (
    "EugeneORM/clause"
    "EugeneORM/log"
    "reflect"
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

// Find 传入一个切片指针, 将数据库查询到的数据存入切片中
func (s *Session) Find(values interface{}) error {
    destSlice := reflect.Indirect(reflect.ValueOf(values))                // 获取目标切片的反射值
    destType := destSlice.Type().Elem()                                   // 获取目标切片的元素类型
    table := s.Model(reflect.New(destType).Elem().Interface()).RefTable() // 获取目标切片元素类型对应的数据库表信息

    s.clause.Set(clause.SELECT, table.Name, table.FieldNames)                              // 设置查询的 SELECT 子句，包括表名和字段名
    sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT) // 构建查询的 SQL 语句和参数列表
    rows, err := s.Raw(sql, vars...).QueryRows()                                           // 执行查询并获取结果集
    if err != nil {
        log.Error(err)
        return err
    }

    for rows.Next() { // 遍历结果集的每一行
        dest := reflect.New(destType).Elem()    // 创建目标元素类型的空实例
        var values []interface{}                // 创建用于存储每个字段值的接口切片
        for _, name := range table.FieldNames { // 遍历表的字段名
            values = append(values, dest.FieldByName(name).Addr().Interface()) // 获取字段值的地址，并将地址添加到接口切片中
        }

        if err := rows.Scan(values...); err != nil { // 扫描结果集的当前行，并将值存储到对应的字段地址中
            log.Error(err)
            return err
        }

        destSlice.Set(reflect.Append(destSlice, dest)) // 将填充好的目标实例添加到目标切片中
    }
    return rows.Close()
}
