package main

import (
	"fmt"
	"log"

	"github.com/schedlinkk/server/pkg/api/http"
	"github.com/schedlinkk/server/pkg/config"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("Mongo string: %s", config.MongoString())

	if err := http.Run(); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
