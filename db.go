package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDB() {
	DB = setupDB()
}

func setupDB() *sql.DB {
	log.Println("Setting up DataBase")
	database, e := sql.Open("sqlite3", "./brewtemp.db")
	err(e)
	log.Println("Adding table templog")
	statement, e := database.Prepare("CREATE TABLE IF NOT EXISTS templog (id INTEGER PRIMARY KEY AUTOINCREMENT, temp DECIMAL(4,1), time SMALLDATETIME)")
	err(e)
	_, e = statement.Exec()
	err(e)
	log.Println("adding table logs")
	statement, e = database.Prepare("CREATE TABLE IF NOT EXISTS logs (id INTEGER PRIMARY KEY AUTOINCREMENT, log STRING, time SMALLDATETIME)")
	err(e)
	_, e = statement.Exec()
	err(e)
	log.Println("Database setup successfully")
	return database
}

// InsertEntry ads a temp reading to the database
func InsertEntry(temp float32) {
	statement, e := DB.Prepare(fmt.Sprintf("INSERT INTO templog (temp, time) VALUES (%v, DATETIME('now'))", temp))
	err(e)
	statement.Exec()
}

func GetLatestLogs() []byte {
	logs := []string{}
	rows, e := DB.Query("SELECT * FROM (SELECT time, log FROM logs ORDER BY time ASC limit 10) ORDER BY TIME DESC")
	if e != nil {
		NonFatal(e)
		return []byte("Error getting logs")
	}
	var time string
	var log string
	for rows.Next() {
		rows.Scan(&time, &log)
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
