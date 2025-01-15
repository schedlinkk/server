package main

import (
	"github.com/schedlinkk/server/pkg/api/http"
	"log"
)

func main() {
	if err := http.Run(); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
