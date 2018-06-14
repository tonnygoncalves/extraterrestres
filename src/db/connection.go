package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/src/github.com/joho/godotenv"
)

var (
	db *sql.DB

	dataUsername string
	dataName     string
	dataPassword string
)

func connnection() (*sql.DB, error) {
	var err error
	godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//get values from my .env file
	dataUsername = os.Getenv("DB_USERNAME")
	dataPassword = os.Getenv("DB_PASSWORD")
	dataName = os.Getenv("DB_NAME")

	fmt.Printf("%s:%s@/%s", dataUsername, dataPassword, dataName)
	mysql.RegisterLocalFile("")
	// Abro la connexion
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", dataUsername, dataPassword, dataName))
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, err
}
