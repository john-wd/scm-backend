package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/john-wd/scm-backend/mock"
	"github.com/john-wd/scm-backend/server"
)

var (
	mockPath = flag.String("mockPath", "", "where mocks are located")
)

func main() {
	flag.Parse()
	if *mockPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	mock.Configure(*mockPath)
	mux := chi.NewMux()
	serv := server.New()

	mux = serv.RegisterRoutes(mux)
	addr := "localhost:9999"
	log.Default().Printf("starting brstm server listening to %s", addr)

	http.ListenAndServe(addr, mux)
}
