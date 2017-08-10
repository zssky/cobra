package db

import (
	"testing"

	"github.com/zssky/log"
)

const (
	mysqlhost   = "172.17.0.2"
	mysqlport   = 3306
	mysqluser   = "root"
	mysqlpasswd = "123456"
	mysqldb     = "test"
)

func newMysql(host string, port int, user, passwd string, dbname string) Impl {
	db, err := NewMysql(host, port, user, passwd, dbname)
	if err != nil {
		log.Fatalf("Open Mysql error, err:%v", err)
	}

	return db
}

func TestMysqlDBType(t *testing.T) {
	db := newMysql(mysqlhost, mysqlport, mysqluser, mysqlpasswd, mysqldb)

	dbType := db.DBType()
	t.Logf("DBType:%v", dbType)
}

func TestMysqlShowTables(t *testing.T) {
	db := newMysql(mysqlhost, mysqlport, mysqluser, mysqlpasswd, mysqldb)

	list, err := db.ShowTables()
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	t.Logf("list:%v", list)
}
