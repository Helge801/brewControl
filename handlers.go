package main

import (
	"fmt"
	"log"
	"net/http"
)

// HandleShutdown gracefully shuts it down
func HandleShutdown(w http.ResponseWriter, r *http.Request) {
	// Start temp recording here
}

// HandleRoot http requests to / and serves react frontend
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	// serve react frontend here
}

// HandleSubscribe handles requests to /subscribe and upgrades connection to websocket to serve live data readings
func HandleSubscribe(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(msgType)
		fmt.Println(string(msg))
		err = socket.WriteMessage(msgType, msg)
	}
}

func HandleGetLogs() {

}
