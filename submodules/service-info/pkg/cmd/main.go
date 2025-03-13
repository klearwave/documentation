package main

import (
	"net/http"

	"github.com/klearwave/service-info/pkg/server"
)

func main() {
	// create the server
	server, err := server.NewServer()
	if err != nil {
		panic(err)
	}

	// ensure we have database connectivity
	db, err := server.Database.Connection.DB()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	// register our routes once we confirm database connectivity is established
	server.RegisterRoutes()

	// start the server
	http.ListenAndServe("0.0.0.0:8888", server.Router)
}
