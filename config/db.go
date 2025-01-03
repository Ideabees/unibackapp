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
	dsn := "gin:ginpassword@tcp(mysql:3306)/mydb"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Database not reachable: %v", err)
	}
	fmt.Println("Database connection established")

}
