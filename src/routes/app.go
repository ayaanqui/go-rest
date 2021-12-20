package routes

import (
	"database/sql"
	"log"
)

type AppBase struct {
	DB *sql.DB
}

type IAppBase interface {
	NewBaseHandler(conn *sql.DB)
}

func (app *AppBase) NewBaseHandler(conn *sql.DB) {
	app.DB = conn

	_, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS home (
			id BINARY(16) NOT NULL DEFAULT(
				UUID_TO_BIN(UUID())
			) PRIMARY KEY, 
			message TEXT NOT NULL,
			date DATETIME NOT NULL DEFAULT(NOW())
		)
	`)
	if err != nil {
		log.Fatal("Could not execute CREATE query for home table.\n", err)
		return
	}
}