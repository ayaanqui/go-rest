package utils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/ayaanqui/go-rest-server/src/types"
	_ "github.com/go-sql-driver/mysql"
)

func load_db_config() (types.DbConnection, error) {
	db_config_file_data, err := ioutil.ReadFile("db_config.json")
	if err != nil {
		return types.DbConnection{}, errors.New("db_config.json file not found")
	}

	var db_config types.DbConnection
	err = json.Unmarshal([]byte(db_config_file_data), &db_config)
	if err != nil {
		return types.DbConnection{}, err
	}
	return db_config, nil
}

func DbConnect() (*sql.DB, error) {
	config, err := load_db_config()
	if err != nil {
		return nil, err
	}
	url := config.Username + ":" + config.Password + "@tcp(localhost:3306)/" + config.DbName
	return sql.Open("mysql", url)
}