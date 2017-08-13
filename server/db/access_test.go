package db

import (
	"testing"

	"github.com/zssky/log"
)

const (
	accessdbname = "../data/accessdb.accdb"
)

func newAccess(name string) Impl {
	access, err := NewAccess("", 0, "", "", name)
	if err != nil {
		log.Fatalf("Open Access error, err:%v", err)
	}

	return access
}

func TestNewAccess(t *testing.T) {
	access := newAccess(accessdbname)
	defer access.Close()

	t.Logf("access Open Success, access:%v", access)
}

func TestAccessDBType(t *testing.T) {
	access := newAccess(accessdbname)
	defer access.Close()

	dbType := access.DBType()
	t.Logf("dbType:%v", dbType)
}

func TestAccessShowTables(t *testing.T) {
	access := newAccess(accessdbname)
	defer access.Close()

	list, err := access.ShowTables()
	if err != nil {
		t.Fatalf("Show Tables error, err:%v", err)
	}

	t.Logf("tablet list:%v", list)

}
