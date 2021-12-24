package routes

import (
	"database/sql"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AppBase struct {
	DB *gorm.DB
}

type IAppBase interface {
	NewBaseHandler(conn *sql.DB)
	CreateRoutes(router *mux.Router)
}

func (app *AppBase) NewBaseHandler(conn *gorm.DB) {
	app.DB = conn

	conn.AutoMigrate(&types.Post{})
	conn.AutoMigrate(&types.Home{})
}

// Create routes given a gorilla/mux router instance
func (app *AppBase) CreateRoutes(router *mux.Router) {
	router.HandleFunc("/", app.Home).Methods("GET")
	router.HandleFunc("/posts", app.CreatePost).Methods("POST")
	router.HandleFunc("/posts", app.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", app.GetPostFromId).Methods("GET")
}

func New(conn *gorm.DB) *AppBase {
	app := AppBase{}
	app.NewBaseHandler(conn)
	return &app
}