package app

import (
	"log"

	"github.com/ayaanqui/go-rest/src/types"
)

func (app *AppBase) CreateSchemas() {
	err := app.DB.AutoMigrate(
		&types.Post{},
		&types.User{},
	)
	if err != nil {
		log.Fatal("Could not generate schema.\n")
		panic(err)
	}
}