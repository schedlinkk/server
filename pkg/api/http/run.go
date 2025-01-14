package http

import "net/http"

func Run() error {

	mux := http.NewServeMux()

	s := http.Server{
		Addr:    ":8888",
		Handler: mux,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil

}
