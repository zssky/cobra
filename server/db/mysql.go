package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//Register mysql implation to Factory
	RegisterFactory(DB_TYPE_MYSQL, NewMysql)
}

// CobraMysql is a implation for CobraDB to operation mysql
type CobraMysql struct {
	sql.DB
}

// NewMysql
func NewMysql(host string, port int, user, password string, dbname string) (Impl, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, password, host, port, dbname)
	db, err := sql.Open(DB_TYPE_MYSQL, args)
	if err != nil {
		return nil, err
	}

	return &CobraMysql{*db}, nil
}

// DBType return database type
func (db *CobraMysql) DBType() string {
	return DB_TYPE_MYSQL
}

// ShowTables - Show db's table list
func (db *CobraMysql) ShowTables() ([]string, error) {
	rows, err := db.Query("show tables")
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
