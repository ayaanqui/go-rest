package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
)

func (app *AppBase) Posts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		app.create_post(w, r)
	case "GET":
		app.get_posts(w, r)
	default:
		utils.JsonResponse(w, types.Response{
			Message: fmt.Sprintf("Method [%s] not allowed", r.Method),
		})
	}
}

func (app *AppBase) get_posts(w http.ResponseWriter, r *http.Request) {
	posts := make([]types.Post, 0)
	app.DB.Table("posts").Find(&posts)
	utils.JsonResponse(w, types.Result{Data: &posts})
}

func (app *AppBase) create_post(w http.ResponseWriter, r *http.Request) {
	data := types.CreatePost{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.JsonResponse(w, types.Response{Message: "Could not parse body data"})
		return
	}
	
	// Create post
	slug := strings.ReplaceAll(data.Title, " ", "-")
	slug = strings.ToLower(slug)
	new_post := types.Post{
		Title: data.Title,
		Slug: slug,
		Summary: data.Summary,
		Content: data.Content,
	}
	app.DB.Create(&new_post)
	utils.JsonResponse(w, new_post)
}