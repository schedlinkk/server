package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/schedlinkk/server/pkg/auth"
	"github.com/schedlinkk/server/pkg/services/users"
)

func NewUserscontrollerWithSimpleAuth(us *users.UserService) http.Handler {
	return auth.Simple(NewUsersController(us))
}

func NewUsersController(us *users.UserService) *UsersController {
	uc := &UsersController{s: us}
	return uc
}

type UsersController struct {
	s *users.UserService
}

func (uc *UsersController) FetchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uid, ok := ctx.Value(auth.UID).(users.UID)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := uc.s.FetchUserByUID(uid)
	if err != nil {
		if err == users.ErrUserDontExist {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(userJSON); err != nil {
		log.Printf("Failed to write respone to user | %s\n", user.UID)
	}

}

func (uc *UsersController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/info":
		uc.FetchUser(w, r)
	}
}
