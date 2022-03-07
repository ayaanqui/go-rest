package app

import (
	"database/sql"
	"log"

	"github.com/ayaanqui/go-rest/src/types"
	"github.com/ayaanqui/go-rest/src/utils"
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

	// database
	err := conn.AutoMigrate(
		&types.Post{},
		&types.User{},
	)
	if err != nil {
		log.Fatal("Could not generate schema.\n")
		panic(tokens_err)
	}

	// oauth providers
	app.SetupOauthProviders()

	return app
}

func New(conn *gorm.DB) *AppBase {
	app := AppBase{}
	return app.NewBaseHandler(conn)
}