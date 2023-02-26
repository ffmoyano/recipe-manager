package server

import (
	"net/http"
	"recipes/src/handler"
)

func SetHandlers(mux *http.ServeMux) {
	mux.Handle("/user", &handler.UserHandler{})
}
