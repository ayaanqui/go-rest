package routes

import (
	"github.com/gorilla/mux"
)

func CreateRoutes(app *AppBase, router *mux.Router) {
	router.HandleFunc("/", app.Home).Methods("GET")
	router.HandleFunc("/posts", app.CreatePost).Methods("POST")
	router.HandleFunc("/posts", app.GetPosts).Methods("GET")
}