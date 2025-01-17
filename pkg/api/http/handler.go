package http

import (
	"net/http"

	"github.com/schedlinkk/server/pkg/api/http/hello"
	"github.com/schedlinkk/server/pkg/api/http/usercontroller"
	"github.com/schedlinkk/server/pkg/services/users"
)

func Handler() *http.ServeMux {
	userStore := users.NewUserStoreStubWithSampleData()
	userService := users.NewUserServiceWithStore(userStore)
	userController := usercontroller.NewWithSimpleAuth(userService)

	mux := http.NewServeMux()
	mux.Handle("/hello", hello.Handler())
	mux.Handle("/users/", http.StripPrefix("/users", userController))

	return mux

}
