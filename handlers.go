package main

import (
	"fmt"
	"log"
	"net/http"
)

// HandleStartRecording initiates temp recordings
func HandleStartRecording(w http.ResponseWriter, r *http.Request) {
	// Start temp recording here
}

// HandleRoot handles requests to / and upgrades connection to websocket
func HandleRoot(w http.ResponseWriter, r *http.Request) {
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
