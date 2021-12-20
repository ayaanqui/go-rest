package routes

import "database/sql"

type AppBase struct {
	DB *sql.DB
}

type IAppBase interface {
	NewBaseHandler(conn *sql.DB)
}

func (app *AppBase) NewBaseHandler(conn *sql.DB) {
	app.DB = conn
}