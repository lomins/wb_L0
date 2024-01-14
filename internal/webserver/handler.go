package webserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	s storer
}

func newHandler(s storer) *handler {
	return &handler{s}
}

func (h *handler) getOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	oid := vars["id"]

	order, found := h.s.Get(oid)
	if !found {
		http.Error(w, "order_id not found "+oid, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(order)
}
