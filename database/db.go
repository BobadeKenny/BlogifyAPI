package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

var Db *sql.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
	}

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
	host, port, user, dbname, password)
db, errSql := sql.Open("postgres", psqlSetup)
if errSql != nil {
   fmt.Println("There is an error while connecting to the database ", err)
   panic(err)
} else {
   Db = db
   fmt.Println("Successfully connected to database!")
}

}