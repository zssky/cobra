package db

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func init() {
	//Register mssql implation to Factory
	RegisterFactory(DB_TYPE_MSSQL, NewMSSql)
}

// CobraMSSql is a implation for CobraDB to operation mssql
type CobraMSSql struct {
	sql.DB
}

// NewMSSql
func NewMSSql(host string, port int, user, password string, dbname string) (Impl, error) {
	args := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", host, user, password, port, dbname)
	db, err := sql.Open(DB_TYPE_MSSQL, args)
	if err != nil {
		return nil, err
	}

	return &CobraMSSql{*db}, nil
}

// DBType return database type
func (db *CobraMSSql) DBType() string {
	return DB_TYPE_MSSQL
}

// ShowTables - Show db's table list
func (db *CobraMSSql) ShowTables() ([]string, error) {
	rows, err := db.Query(" SELECT Name FROM SysObjects Where XType='U' ORDER BY Name")
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
