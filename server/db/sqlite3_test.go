package db

import (
	"testing"

	"github.com/zssky/log"
)

const (
	dbname = "../data/sqlite3.db"
)

func newSqlite3(name string) Impl {
	sqlite3, err := NewSqlite3("", 0, "", "", dbname)
	if err != nil {
		log.Fatalf("Open Sqlite3 error, err:%v", err)
	}

	return sqlite3
}

func TestNewSqlite3(t *testing.T) {
	sqlite3 := newSqlite3(dbname)
	defer sqlite3.Close()

	t.Logf("sqlite Open Success, sqlite3:%v", sqlite3)
}

func TestSqlite3DBType(t *testing.T) {
	sqlite3 := newSqlite3(dbname)
	defer sqlite3.Close()

	dbType := sqlite3.DBType()
	t.Logf("dbType:%v", dbType)
}

func TestSqlite3ShowTables(t *testing.T) {
	sqlite3 := newSqlite3(dbname)
	defer sqlite3.Close()

	list, err := sqlite3.ShowTables()
	if err != nil {
		t.Fatalf("Show Tables error, err:%v", err)
	}

	t.Logf("tablet list:%v", list)

}
