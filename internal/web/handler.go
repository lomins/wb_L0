package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lomins/wildberriesL0/internal/config"
)

type WebServer struct {
	http.Server
}

func New(cfg config.Config) *WebServer {
	r := mux.NewRouter()

	return &WebServer{
		http.Server{
			Addr:    cfg.WebSerberAddr,
			Handler: r,
		},
	}
}

func (srv *WebServer) Launch() {
	log.Println("launching web-server on address: ", srv.Server.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start web-server: ", err)
	}
}
