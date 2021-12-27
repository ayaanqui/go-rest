package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
	"github.com/google/uuid"
)

func (app *AppBase) Register(w http.ResponseWriter, r *http.Request) {
	post_user := types.CreateUser{}
	if err := json.NewDecoder(r.Body).Decode(&post_user); err != nil {
		utils.JsonResponse(w, types.Response{Message: "Incorrect json formatting"})
		return
	}

	// Verify if username or email already exists
	check := types.User{}
	app.DB.Table("users").Find(
		&check, 
		"username = ? OR email = ?", 
		post_user.Username, 
		post_user.Email,
	)
	if check.ID != uuid.Nil {
		utils.JsonResponse(w, types.Response{Message: "An account with the username or email already exists"})
		return
	}

	// User does not exist, so create the account
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