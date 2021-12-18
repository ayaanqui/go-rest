package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/types"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(types.Response{Message: "Hello world"})
	})

	const port = "3001"
	fmt.Printf("Starting server at port %s...\n", port)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatalf("Server already started on port %s\n\n", port)
		log.Fatal(err)
	}
}
