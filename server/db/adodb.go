package db

import (
	"database/sql"

	_ "github.com/mattn/go-adodb"
)

// CobraADODB is a implation for CobraDB to operation adodb
type CobraADODB struct {
	sql.DB
}
