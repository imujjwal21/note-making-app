package sqldb

import "database/sql"

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}
