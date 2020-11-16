package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error
    connStr := "host=localhost port=5432 user=postgres password=11111111 dbname=acckasdq_1 sslmode=disable"
    DB, err = sql.Open("postgres", connStr)
    if err != nil{
    	panic(err)
    }
}