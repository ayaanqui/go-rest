package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/ayaanqui/go-rest-server/src/types"
)

func DbConfig() (types.DbConnection, error) {
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