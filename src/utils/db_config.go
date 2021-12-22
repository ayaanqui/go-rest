package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/ayaanqui/go-rest-server/src/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func DbConnect() (*gorm.DB, error) {
	config, err := load_db_config()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", config.Username, config.Password, config.DbName)
	return gorm.Open(mysql.Open(url), &gorm.Config{})
}