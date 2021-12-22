package routes

import (
	"database/sql"

	"github.com/ayaanqui/go-rest-server/src/types"
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

	conn.AutoMigrate(&types.Post{})
	conn.AutoMigrate(&types.Home{})
}