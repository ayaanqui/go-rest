package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/routes"
)

func main() {
	http.HandleFunc("/", routes.Home)
	http.HandleFunc("/post", routes.Post)

	const port = "3001"
	fmt.Printf("Starting server at port %s...\n", port)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatalf("Server already started on port %s\n\n", port)
		log.Fatal(err)
	}
}
