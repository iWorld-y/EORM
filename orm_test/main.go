package main

import (
	"EugeneORM"
	"EugeneORM/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := EugeneORM.NewEngine("sqlite3", "EugeneORMTest.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?);", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	log.Info("执行成功, 改变了 ", count, " 条数据\n")
}
