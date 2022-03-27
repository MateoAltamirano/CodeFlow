package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type CodeFlowServer struct {
	server *http.Server
}

func CreateServer(mux *chi.Mux) *CodeFlowServer {
	server := &http.Server{
		Addr:           ":9000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &CodeFlowServer{server}
}

func (codeFlowServer *CodeFlowServer) Run() {
	fmt.Print("Server running on port: 9000")
	log.Fatal(codeFlowServer.server.ListenAndServe())
}