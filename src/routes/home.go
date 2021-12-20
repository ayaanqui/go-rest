package routes

import (
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
)

func (app *AppBase) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.JsonResponse(w, types.Response{Message: "Only GET request allowed"})
		return
	}

	message := r.URL.Query().Get("message")
	if message == "" {
		message = "Hello world"
	}
	utils.JsonResponse(w, types.Response{Message: message})
}