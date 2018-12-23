package main

// Entry is a database entry as exists in go
type Entry struct {
	Time string  `json:"time"`
	Temp float32 `json:"temp"`
}
