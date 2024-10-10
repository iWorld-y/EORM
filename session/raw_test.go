package session

import (
	"EORM/log"
	"EORM/utils"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func newSession(t *testing.T) *Session {
	testDB, err := sql.Open(utils.SqliteDBDriverName, fmt.Sprintf("../%s", utils.SqliteData))
	if err != nil {
		log.Error(err)
		t.Fatal(err)
	}
	return New(testDB)
}
func TestSession_Exec(t *testing.T) {
	session := newSession(t)
	_, _ = session.Raw("drop table if exists User").Exec()
	_, _ = session.Raw("create table User(Name text, Age int)").Exec()

	ret, _ := session.Raw("insert into User(`Name`,`Age`) values (?, ?), (?, ?)", "Eugene", 18, "TiTi", 17).Exec()
	if cnt, err := ret.RowsAffected(); err != nil || cnt != 2 {
		t.Fatalf("expect 2, but got: %d\terr: %v", cnt, err)
	}
}

func TestSession_QueryRows(t *testing.T) {
	session := newSession(t)
	//_, _ = session.Raw("drop table if exists User").Exec()
	//_, _ = session.Raw("create table User(Name text, Age int)").Exec()

	row := session.Raw("select count(*) from User").Query()
	cnt := 0
	if err := row.Scan(&cnt); err != nil || cnt == 0 {
		t.Fatal("failed ", err)
	}
}
