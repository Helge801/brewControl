package main

import "fmt"

// AdLog ads a log entry to the database that can be later retrieved
func AdLog(message string) {
	statement, e := db.Prepare(fmt.Sprintf("INSERT INTO logs (log, time) VALUES (%v, DATETIME('now'))", message))
	err(e)
	statement.Exec()
}

// Fatal error handling for for errors that must be nil
func Fatal(e error) {
	if e != nil {
		AdLog(fmt.Sprint(e))
		panic(e)
	}
}

// NonFatal error handling for errors that should be nil but do not break the entire program
func NonFatal(e error) bool {
	if e != nil {
		AdLog(fmt.Sprint(e))
		return true
	}
	return false
}
