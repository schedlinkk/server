package hello

import (
	"fmt"
	"log"
	"net/http"

	"github.com/schedlinkk/server/pkg/auth"
)

func Handler() http.Handler {
	return auth.Simple(handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uid, ok := ctx.Value(auth.UUID).(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	s := fmt.Sprintf("Hello, %s\n", uid)
	if _, err := w.Write([]byte(s)); err != nil {
		log.Printf("Failed to write response to %s", uid)
	}
}
