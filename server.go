package main

import "net/http"

// StartServer self explanitory
func StartServer() {
	http.HandleFunc("/", HandleRoot)
	http.ListenAndServe(":4000", nil)
}
