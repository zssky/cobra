package db

import (
	"database/sql"
)

func init() {
	//Register access implation to Factory
	RegisterFactory(DB_TYPE_ACCESS, NewAccess)
}

// CobraAccess is a implation for CobraDB to operation Access
type CobraAccess struct {
	CobraADODB
}

// NewAccess
func NewAccess(host string, port int, user, password string, dbname string) (Impl, error) {
	db, err := sql.Open(DB_TYPE_ADO, "Provider=Microsoft.Jet.OLEDB.4.0;Data Source="+dbname+";")
	if err != nil {
		return nil, err
	}

	return &CobraAccess{CobraADODB{*db}}, nil
}

// DBType return database type
func (db *CobraAccess) DBType() string {
	return DB_TYPE_ACCESS
}

// ShowTables - Show db's table list
func (db *CobraAccess) ShowTables() ([]string, error) {
	rows, err := db.Query("select name from msysobjects where type=1 and flags=0")
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
