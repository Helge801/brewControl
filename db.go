package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitializeDB() {
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
	statement, e = database.Prepare("CREATE TABLE IF NOT EXISTS logs (id INTEGER PRIMARY KEY AUTOINCREMENT, log STRING, time SMALLDATETIME)")
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

func GetLatestLogs() []byte {
	logs := []string{}
	rows, e := db.Query("select * from (select * from logs order by time ASC limit 10) order by time DESC")
	if e != nil {
		NonFatal(e)
		return []byte("Error getting logs")
	}
	var time string
	var log string
	rows.Scan(&time)
	logs = append(logs, time+": "+log)
	for rows.Next() {
		rows.Scan(&time)
		logs = append(logs, time+": "+log)
	}
	mJSON, e := json.Marshal(logs)
	if e != nil {
		NonFatal(errors.New("Could notmarshal logs to JSON"))
	}
	return mJSON
}

func err(e error) {
	if e != nil {
		panic(e)
	}
}
