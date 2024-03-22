package session

import (
    "EugeneORM/dialect"
    "database/sql"
    "testing"
)

var (
    TestDB      *sql.DB
    TestDial, _ = dialect.GetDialect("sqlite3")
)

func TestMain(t *testing.T) {
    TestDB, _ = sql.Open("sqlite3", "EugeneORMTest.db")
    code := t.Run()
    _ = TestDB.Close()
}

func NewSession() *Session {
    return New(TestDB, TestDial)
}

func TestSession_Exec(t *testing.T) {
    s := NewSession()
    _, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
    _, _ = s.Raw("CREATE TABLE User (Name TEXT);").Exec()
    result, _ := s.Raw("INSERT INTO User(`Name`) VALUES (?), (?);", "Tom", "Sam").Exec()

    if count, err := result.RowsAffected(); err != nil || count != 2 {
        t.Fatal("expect 2, but got ", count)
    }
}

func TestSession_QueryRows(t *testing.T) {
    s := NewSession()
    _, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
    _, _ = s.Raw("CREATE TABLE User (Name TEXT);").Exec()
    result, _ := s.Raw("INSERT INTO User(`Name`) VALUES (?), (?);", "Tom", "Sam").Exec()

}
