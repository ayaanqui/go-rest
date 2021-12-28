package routes

import (
	"database/sql"
	"log"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AppBase struct {
	DB *gorm.DB
	Tokens types.Tokens
}

type IAppBase interface {
	NewBaseHandler(conn *sql.DB)
	CreateRoutes(router *mux.Router)
}

func (app *AppBase) NewBaseHandler(conn *gorm.DB) *AppBase {
	app.DB = conn
	tokens, tokens_err := utils.ParseTokens()
	if tokens_err != nil {
		panic(tokens_err)
	}
	app.Tokens = tokens

	err := conn.AutoMigrate(
		&types.Post{}, 
		&types.Home{},
		&types.User{},
	)
	if err != nil {
		log.Fatal("Could not generate schema.\n")
		panic(tokens_err)
	}
	return app
}

// Create routes given a gorilla/mux router instance
func (app *AppBase) CreateRoutes(router *mux.Router) *AppBase {
	router.HandleFunc("/", app.Home).Methods("GET")
	router.HandleFunc("/posts", app.CreatePost).Methods("POST")
	router.HandleFunc("/posts", app.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", app.GetPostFromId).Methods("GET")
	router.HandleFunc("/register", app.Register).Methods("POST")
	router.HandleFunc("/login", app.Login).Methods("POST")
	return app
}

func New(conn *gorm.DB) *AppBase {
	app := AppBase{}
	return app.NewBaseHandler(conn)
}