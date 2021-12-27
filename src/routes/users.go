package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
)

func (app *AppBase) CreateUser(w http.ResponseWriter, r *http.Request) {
	post_user := types.CreateUser{}
	if err := json.NewDecoder(r.Body).Decode(&post_user); err != nil {
		utils.JsonResponse(w, types.Response{Message: "Incorrect json formatting"})
		return
	}
	new_user := types.User{
		Username: post_user.Username,
		Email: post_user.Email,
		Password: post_user.Password,
		IsAdmin: false,
		IsActive: true,
	}
	app.DB.Table("users").Create(&new_user)
	utils.JsonResponse(w, &new_user)
}