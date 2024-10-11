package session

import (
	"EORM/dialect"
	"EORM/log"
	"EORM/utils"
	"database/sql"
	"errors"
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
	testDial, ok := dialect.GetDialect(utils.SqliteDBDriverName)
	if !ok {
		err = errors.New("get dialect failed")
		log.Error(err)
		t.Fatal(err)
	}
	return New(testDB, testDial)
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

type Account struct {
	Id   int64  `EORM:"PRIMARY KEY"`
	Name string `EORM:""`
}

func TestSession_CreateTable(t *testing.T) {
	log.SetLevel(log.InfoLevel)
	session := newSession(t).Model(&Account{})
	if err := session.DropTable(); err != nil {
		t.Fatal(err)
	}
	if err := session.CreateTable(); err != nil {
		t.Fatal(err)
	}
}

func TestSession_HasTable(t *testing.T) {
	log.SetLevel(log.InfoLevel)
	session := newSession(t).Model(&Account{})
	_ = session.CreateTable()
	if !session.HasTable() {
		err := errors.New("expect true, bug false")
		log.Error(err)
		t.Fatal(err)
	}
	err := session.DropTable()
	for session.HasTable() || err != nil {
		if err == nil {
			err = errors.New("expect false, bug true")
		} else if errors.Is(err, sql.ErrNoRows) {
			break
		}
		log.Error(err)
		t.Fatal(err)
	}
}
