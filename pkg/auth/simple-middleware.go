package auth

import (
	"context"
	"log"
	"net/http"
)

func Simple(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		uid := cookie.Value
		log.Printf("%s logged in.", cookie.Value)

		ctx := r.Context()
		ctx = context.WithValue(ctx, "uid", uid)
		r = r.WithContext(ctx)

		next(w, r)
	})
}
