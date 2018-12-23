package main

import (
	"fmt"
	"log"
	"math/rand"
)

// AdLog ads a log entry to the database that can be later retrieved
func AdLog(message string) {
	statement, e := DB.Prepare("INSERT INTO logs (log, time) VALUES (\"" + message + "\", DATETIME('now'))")
	log.Println(e)
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

// RandomKey generates and return a randow hash key to use as connection ids
func RandomKey() string {
	len := 10
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}
