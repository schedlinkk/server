package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/schedlinkk/server/pkg/services/users"
)

func Simple(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		uid := users.UID(cookie.Value)
		log.Printf("%s logged in.", cookie.Value)

		ctx := r.Context()
		ctx = context.WithValue(ctx, UID, uid)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
