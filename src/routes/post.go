package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.JsonResponse(w, types.Response{Message: "Method not allowed"})
		return
	}
	
	person := types.Person{}
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		utils.JsonResponse(w, types.Response{Message: "Could not parse body data"})
		return
	}
	utils.JsonResponse(w, person)
}