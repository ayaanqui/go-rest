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

	home := `CREATE TABLE IF NOT EXISTS home (
		id BINARY(16) NOT NULL DEFAULT(
			UUID_TO_BIN(UUID())
		) UNIQUE PRIMARY KEY, 
		message TEXT NOT NULL,
		date DATETIME NOT NULL DEFAULT(NOW())
	)`
	if _, err := conn.Exec(home); err != nil {
		log.Fatal("Could not execute CREATE query for home table.\n", err)
		return
	}

	post := `CREATE TABLE IF NOT EXISTS post (
		id BINARY(16) NOT NULL DEFAULT(
			UUID_TO_BIN(UUID())
		) UNIQUE PRIMARY KEY, 
		title VARCHAR(255) NOT NULL,
		slug VARCHAR(255) NOT NULL UNIQUE PRIMARY KEY,
		content TEXT NOT NULL,
		summary TEXT NOT NULL,
		date DATETIME NOT NULL DEFAULT(NOW()),
		updated DATETIME NOT NULL DEFAULT(NOW())
	)`
	if _, err := conn.Exec(post); err != nil {
		log.Fatal("Could not execute CREATE query for post table.\n", err)
		return
	}
}