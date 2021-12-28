package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (app *AppBase) Register(w http.ResponseWriter, r *http.Request) {
	post_user := types.CreateUser{}
	if err := json.NewDecoder(r.Body).Decode(&post_user); err != nil {
		w.WriteHeader(400)
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
		w.WriteHeader(400)
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

func (app *AppBase) Login(w http.ResponseWriter, r *http.Request) {
	login := types.LoginUser{}
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		w.WriteHeader(400)
		utils.JsonResponse(w, types.Response{Message: "Could not parse json"})
		return
	}

	const message string = "Username or password is incorrect"
	user := types.User{}
	app.DB.Table("users").Find(&user, "username = ? OR email = ?", login.Username, login.Username)
	if user.ID == uuid.Nil {
		w.WriteHeader(400)
		utils.JsonResponse(w, types.Response{Message: message})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		w.WriteHeader(400)
		utils.JsonResponse(w, types.Response{Message: message})
		return
	}
	utils.JsonResponse(w, types.Response{Message: "Logged in successfully!"})
}