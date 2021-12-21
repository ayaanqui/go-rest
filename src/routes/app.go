package routes

import (
	"database/sql"

	"gorm.io/gorm"
)

type AppBase struct {
	DB *gorm.DB
}

type IAppBase interface {
	NewBaseHandler(conn *sql.DB)
}

func (app *AppBase) NewBaseHandler(conn *gorm.DB) {
	app.DB = conn
}