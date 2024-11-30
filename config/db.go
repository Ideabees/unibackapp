package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:Admin@12345@tcp(localhost:3306)/uniapp")
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Database not reachable: %v", err)
	}
	fmt.Println("Database connection established")
}
