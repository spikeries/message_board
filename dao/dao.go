package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	dB = db
}
