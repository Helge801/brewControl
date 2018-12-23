package main

import (
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

func main() {
	InitializeDB()
	StartMonitor()
	StartServer()
}

func handler(w http.ResponseWriter, r *http.Request) {
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
