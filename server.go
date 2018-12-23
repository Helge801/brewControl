package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Subscribers is an active list of subscribers
var Subscribers = map[string]*websocket.Conn{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// StartServer self explanitory
func StartServer() {
	http.HandleFunc("/subscribe", HandleSubscribe)
	http.HandleFunc("/shutdown", HandleShutdown)
	http.HandleFunc("/logs", HandleGetLogs)
	http.ListenAndServe(":4000", nil)
}

// SendEntry will sent an entry with the given temp to all subscribers
// This is an extreamly crude way of handling subscribers but it works for now
func SendEntry(temp float32) {
	entry := Entry{
		Time: fmt.Sprintln(time.Now().Format("2006-01-02 15:04:05")),
		Temp: temp,
	}
	mJSON, e := json.Marshal(entry)
	NonFatal(e)
	for k, v := range Subscribers {
		e := v.WriteMessage(1, mJSON)
		if e != nil {
			delete(Subscribers, k)
		}
	}
}
