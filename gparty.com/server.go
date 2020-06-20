package main 

import (
	"log"

	party "gparty.com/server/party"
)

func main() {
	log.Println("Starting backend")
	go party.StartBackendServer()
	log.Println("Starting frontend")
	party.StartFrontendServer()
}

