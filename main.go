package main

import (
	"keyValueStorage/api"
	"keyValueStorage/store"
	"log"
)

func main() {
	service := store.NewKVStore()

	log.Println("Starting server")
	api.StartServer(service)
}
