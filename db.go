package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	db = setupDB()

	rows, e := db.Query("SELECT temp, time FROM templog WHERE time <= '2018-12-23 02:37:00' AND time > '2018-12-23 02:27:32'")
	err(e)
	var temp string
	var time string
	for rows.Next() {
		rows.Scan(&temp, &time)
		fmt.Println(temp + " " + time)
	}

}

func setupDB() *sql.DB {
	database, e := sql.Open("sqlite3", "./brewtemp.db")
	err(e)
	statement, e := database.Prepare("CREATE TABLE IF NOT EXISTS templog (id INTEGER PRIMARY KEY AUTOINCREMENT, temp DECIMAL(4,1), time SMALLDATETIME)")
	err(e)
	statement, e := database.Prepare("CREATE TABLE IF NOT EXISTS logs (id INTEGER PRIMARY KEY AUTOINCREMENT, log STRING, time SMALLDATETIME)")
	err(e)
	statement.Exec()
	return database
}

// InsertEntry ads a temp reading to the database
func InsertEntry(temp float32) {
	statement, e := db.Prepare(fmt.Sprintf("INSERT INTO templog (temp, time) VALUES (%v, DATETIME('now'))", temp))
	err(e)
	statement.Exec()
}

func err(e error) {
	if e != nil {
		panic(e)
	}
}
