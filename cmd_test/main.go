package main

import (
	"EORM"
	"EORM/log"
	"EORM/utils"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.SetLevel(log.InfoLevel)
	engine, err := EORM.NewEngine(utils.SqliteDBDriverName, fmt.Sprintf("../%s", utils.SqliteData))
	if err != nil {
		log.Error(err)
	}
	defer engine.Close()

	session := engine.NewSession()
	_, _ = session.Raw("drop table if exists User").Exec()
	_, _ = session.Raw("create table User(Name text, Age int)").Exec()

	ret, _ := session.Raw("insert into User(`Name`,`Age`) values (?, ?), (?, ?)", "Eugene", 18, "TiTi", 17).Exec()
	var cnt int64
	if cnt, err = ret.RowsAffected(); err != nil || cnt != 2 {
		log.Errorf("expect 2, but got: %d\terr: %v", cnt, err)
		return
	}
	log.Infof("测试完毕, cnt = %d", cnt)
}
