package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/routes"
	"github.com/ayaanqui/go-rest-server/src/utils"
)

func main() {
	// Attempt connection with DB
	conn, err := utils.DbConnect()
	if err != nil {
		log.Fatal("Could not connect to database.\n", err)
		return
	}
	defer conn.Close()
	
	// Create app base with DB connection
	app := routes.AppBase{}
	app.NewBaseHandler(conn)

	// Create server instance
	server := http.NewServeMux()
	server.Handle("/", utils.HandleGet(http.HandlerFunc(app.Home)))
	server.Handle("/post", utils.HandlePost(http.HandlerFunc(routes.Post)))

	const port = "3001"
	fmt.Printf("Starting server at port %s...\n", port)

	if err := http.ListenAndServe(":" + port, server); err != nil {
		log.Fatalf("Server already started on port %s\n\n", port)
		log.Fatal(err)
	}
}
