package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Create routes given a gorilla/mux router instance
func (app *AppBase) CreateRoutes(router *mux.Router) *AppBase {
	router.HandleFunc("/", app.Home).Methods("GET")
	router.HandleFunc("/posts", app.CreatePost).Methods("POST")
	router.HandleFunc("/posts", app.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", app.GetPostFromId).Methods("GET")
	router.HandleFunc("/register", app.Register).Methods("POST")
	router.HandleFunc("/login", app.Login).Methods("POST")
	router.Handle("/me", UseJwtAuth(app, http.HandlerFunc(app.Profile))).Methods("GET")
	router.HandleFunc("/auth/{provider}", app.Auth).Methods("GET")
	router.HandleFunc("/auth/{provider}/callback", app.AuthCallback).Methods("GET")
	return app
}