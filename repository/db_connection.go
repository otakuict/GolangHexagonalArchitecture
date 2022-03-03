package repository

import "github.com/jmoiron/sqlx"

func Dbconn() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		panic(err)
	}
	return db
}
