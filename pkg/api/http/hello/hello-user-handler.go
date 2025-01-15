package hello

import (
	"fmt"
	"net/http"

	"github.com/schedlinkk/server/pkg/auth"
)

func Handler() http.Handler {
	return auth.Simple(handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uid, ok := ctx.Value("uid").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	s := fmt.Sprintf("Hello, %s\n", uid)
	w.Write([]byte(s))
	return

}
