package http

import (
	"fmt"
	"github.com/schedlinkk/server/pkg/config"
	"github.com/schedlinkk/server/pkg/db/mngdb"
	"log"
	"net/http"
)

func Run() error {

	log.Printf("Start loading config...\n")
	if err := config.Load(); err != nil {
		return err
	}

	if err := mngdb.Connect(); err != nil {
		return fmt.Errorf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := mngdb.Disconnect(); err != nil {
			log.Printf("Failed to disconnect from MongoDB: %v", err)
		}

	}()

	s := http.Server{
		Addr:    ":8888",
		Handler: Handler(), // http.Handler function return global http handler.
	}

	log.Printf("Starting a server...\n")
	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil

}
