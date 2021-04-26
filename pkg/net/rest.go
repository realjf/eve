package net

import (
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/realjf/eve/pkg/lib"
	. "github.com/realjf/eve/terminal"
)

type HttpServer struct {
	router   *mux.Router
	analyzer *Analyzer
}

func NewHttpServer(analyzer *Analyzer) *HttpServer {
	instance := new(HttpServer)
	instance.router = mux.NewRouter()
	instance.analyzer = analyzer
	return instance
}

func (h *HttpServer) Listen() {
	h.router.HandleFunc("/analyzer", h.URLHandler)
	h.router.HandleFunc("/analyzer-api", h.APIHandler)
	h.router.HandleFunc("/ping", h.PingHandler)

}

func (h *HttpServer) URLHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *HttpServer) APIHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *HttpServer) PingHandler(w http.ResponseWriter, r *http.Request) {
	Infoln("pong")
	w.Write([]byte("pong"))
}
