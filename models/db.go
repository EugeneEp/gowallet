package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"os"
)

var DB *sql.DB

func Init() {
	var err error
	connStr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
    DB, err = sql.Open(os.Getenv("DB_CONNECTION"), connStr)
    if err != nil{
    	panic(err)
    }
}