package main

import (
	"errors"
	"net/http"
)

// HandleShutdown handles request to /shutdown and gracefully shuts it down
func HandleShutdown(w http.ResponseWriter, r *http.Request) {
	// Initial implimentation is not so gracefull I admit but the deferal in tempMonitor should shut everything down properly
	Fatal(errors.New("Shutting Down"))
}

// HandleSubscribe handles requests to /subscribe and upgrades connection to websocket to serve live data readings
func HandleSubscribe(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.Write([]byte("error connecting with websocket"))
		return
	}
	hashKey := RandomKey()
	Subscribers[hashKey] = socket
}

// HandleGetLogs handles requests to /logs and returns a JSON object containing the last 100 logs
func HandleGetLogs(w http.ResponseWriter, r *http.Request) {
	mJSON := GetLatestLogs()
	w.Write(mJSON)
}
