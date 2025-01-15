package http

import (
	"log"
	"net/http"

	"github.com/schedlinkk/server/pkg/config"
)

func Run() error {

	log.Printf("Start loading config...\n")
	if err := config.Load(); err != nil {
		return err
	}

	mux := http.NewServeMux()

	s := http.Server{
		Addr:    ":8888",
		Handler: mux,
	}

	log.Printf("Starting a server...\n")
	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil

}
