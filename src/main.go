package main

import (
	"log"
	"net/http"

	"github.com/ayaanqui/go-rest/src/app"
	"github.com/ayaanqui/go-rest/src/utils"
	"github.com/gorilla/mux"
)

func main() {
	// Attempt connection with DB
	conn, err := utils.DbConnect()
	if err != nil {
		log.Fatal("Could not connect to database.\n", err)
		return
	}
	
	// Create router instance
	router := mux.NewRouter()
	// Create server base with DB connection
	server := app.New(conn)
	server.CreateRoutes(router)

	const port = "3001"
	log.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(":" + port, router); err != nil {
		log.Fatalf("Server already started on port %s\n\n", port)
		log.Fatal(err)
	}
}
