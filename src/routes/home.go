package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
)

func select_query(row *sql.Rows, data *types.Home) error {
	var id string
	var message string
	var date string
	
	if err := row.Scan(&id, &message, &date); err != nil {
		return err
	}
	*data = types.Home{
		Id: id, 
		Message: message, 
		Date: date,
	}
	return nil
}

func (app *AppBase) Home(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	if message == "" {
		// Show all messages in home table
		rows, err := app.DB.Query("SELECT BIN_TO_UUID(id) as id, message, date FROM home")
		if err != nil {
			log.Fatal(err)
			utils.JsonResponse(w, types.Response{Message: "Could not process query"})
			return
		}

		parsed_rows := make([]types.Home, 0)
		for rows.Next() {
			var parsed_cols types.Home
			if err := select_query(rows, &parsed_cols); err != nil {
				log.Fatal(err)
				utils.JsonResponse(w, types.Response{Message: "Could not process columns"})
				return
			}
			parsed_rows = append(parsed_rows, parsed_cols)
		}
		rows.Close()
		utils.JsonResponse(w, types.Result{ Data: parsed_rows })
		return
	}

	// Insert message to table
	if _, err := app.DB.Exec("INSERT INTO home (message) VALUES(?)", message); err != nil {
		utils.JsonResponse(w, types.Response{Message: "Could not process insert query"})
		return
	}
	utils.JsonResponse(w, types.Response{Message: message})
}