package db

import (
	"testing"

	"github.com/zssky/log"
)

const (
	mssqlhost   = "172.17.0.3"
	mssqlport   = 1433
	mssqluser   = "sa"
	mssqlpasswd = "1qaz@WSX"
	mssqldb     = "test"
)

func newMSSql(host string, port int, user, passwd string, dbname string) Impl {
	db, err := NewMSSql(host, port, user, passwd, dbname)
	if err != nil {
		log.Fatalf("Open MSSql error, err:%v", err)
	}

	return db
}

func TestMSSqlDBType(t *testing.T) {
	db := newMSSql(mssqlhost, mssqlport, mssqluser, mssqlpasswd, mssqldb)

	dbType := db.DBType()
	t.Logf("DBType:%v", dbType)
}

func TestMSSqlShowTables(t *testing.T) {
	db := newMSSql(mssqlhost, mssqlport, mssqluser, mssqlpasswd, mssqldb)

	list, err := db.ShowTables()
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	t.Logf("list:%v", list)
}
