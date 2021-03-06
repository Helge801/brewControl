package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

// HandleShutdown handles request to /shutdown and gracefully shuts it down
func HandleShutdown(w http.ResponseWriter, r *http.Request) {
	// Initial implimentation is not so gracefull I admit but the deferal in tempMonitor should shut everything down properly
	m := []string{"shutting Down"}
	mJSON, _ := json.Marshal(m)
	go shutdown()
	w.Write(mJSON)
}

func shutdown() {
	RunLoop = false
	time.Sleep(time.Second * 11)
	os.Exit(1)
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
func HandleLogs(w http.ResponseWriter, r *http.Request) {
	mJSON := GetLatestLogs()
	w.Write(mJSON)
}

func HandleGetRecords(w http.ResponseWriter, r *http.Request) {
	mJSON := GetRecords()
	w.Write(mJSON)
}
