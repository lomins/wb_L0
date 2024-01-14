package webserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter(h *handler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/order/{id}", h.getOrder)

	return r
}
