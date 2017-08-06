package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	//Register sqlite3 implation to Factory
	RegisterFactory(DB_TYPE_SQLITE3, NewSqlite3)
}

// CobraSqlite3 is a implation for CobraDB to operation sqlite3
type CobraSqlite3 struct {
	sql.DB
}

// NewSqlite3
func NewSqlite3(host string, port int, user, password string, dbname string) (Impl, error) {
	db, err := sql.Open(DB_TYPE_SQLITE3, dbname)
	if err != nil {
		return nil, err
	}

	return &CobraSqlite3{*db}, nil
}

// DBType return database type
func (db *CobraSqlite3) DBType() string {
	return DB_TYPE_SQLITE3
}

// ShowTables - Show db's table list
func (db *CobraSqlite3) ShowTables() ([]string, error) {
	rows, err := db.Query("select name from sqlite_master where type='table' order by name")
	if err != nil {
		return nil, err
	}

	list := make([]string, 0)
	for rows.Next() {
		var str string
		if err := rows.Scan(&str); err != nil {
			return nil, err
		}

		list = append(list, str)
	}

	return list, nil
}
