package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)

	// Open connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB Connection Error:", err)
	}

	err = db.Ping()
	// Ping for check

	if err != nil {
		log.Fatal("DB Ping Error:", err)
	}

	fmt.Println(" Successfully connected to MySQL")
	DB = db
}
