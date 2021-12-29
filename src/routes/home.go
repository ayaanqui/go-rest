package routes

import (
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
)

func (app *AppBase) Home(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	if message == "" {
		// Show all messages in home table
		result := new([]types.Home)
		app.DB.Table("homes").Select("id", "created_at", "updated_at", "message").Scan(&result)
		utils.JsonResponse(w, types.Result{Data: result})
		return
	}
	// Insert message to table
	new_message := types.Home{Message: message}
	app.DB.Create(&new_message)
	utils.JsonResponse(w, new_message)
}