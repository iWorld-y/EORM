package main

import (
	"EORM"
	"EORM/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := EORM.NewEngine("sqlite3", "EORMTest.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text, Age integer);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text, Age integer);").Exec()
	result, _ := s.Raw("INSERT INTO User(Name, Age) VALUES (?, ?), (?, ?);", "Tom", 18, "Sam", 25).Exec()
	s.Raw("")
	count, _ := result.RowsAffected()
	log.Info("执行成功, 改变了 ", count, " 条数据\n")
}
