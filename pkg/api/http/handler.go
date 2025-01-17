package http

import (
	"net/http"

	"github.com/schedlinkk/server/pkg/api/http/controllers"
	"github.com/schedlinkk/server/pkg/api/http/hello"
	"github.com/schedlinkk/server/pkg/services/users"
)

func Handler() *http.ServeMux {
	userStore := users.NewUserStoreStubWithSampleData()
	userService := users.NewUserServiceWithStore(userStore)
	userController := controllers.NewUserControllerWithSimpleAuth(userService)

	mux := http.NewServeMux()
	mux.Handle("/hello", hello.Handler())
	mux.Handle("/users/", http.StripPrefix("/users", userController))

	return mux

}
