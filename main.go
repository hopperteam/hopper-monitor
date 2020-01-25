package main

import (
	"github.com/hopperteam/hopper-monitor/monitoring"
	"github.com/hopperteam/hopper-monitor/rest"
	"github.com/hopperteam/hopper-monitor/storage"
	"log"
)

func main() {
	prov := storage.LoadConfig()

	log.Print("Starting")

	conn, err := monitoring.Connect(prov, "my-hopper", "localhost:32093")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected")
	go conn.StreamLogs()

	log.Print("Listening")
	log.Fatal(rest.ListenAndServe(prov))
}
