package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
)

func (app *AppBase) Post(w http.ResponseWriter, r *http.Request) {
	data := types.CreatePost{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.JsonResponse(w, types.Response{Message: "Could not parse body data"})
		return
	}
	
	// Create post
	slug := strings.ReplaceAll(data.Title, " ", "-")
	slug = strings.ToLower(slug)
	insert_query := `
	INSERT INTO post 
		(title, slug, summary, content) 
		VALUES (?,?,?,?)
	`
	_, err := app.DB.Exec(insert_query, data.Title, slug, data.Summary, data.Content)
	if err != nil {
		utils.JsonResponse(w, types.Response{Message: "Could not create post"})
		return
	}

	// Return new post
	query := `
	SELECT 
		BIN_TO_UUID(id) as id, 
		title, 
		slug, 
		summary, 
		content, 
		date, 
		updated
	FROM post
		WHERE slug=? 
		LIMIT 1
	`
	row, err := app.DB.Query(query, slug)
	if err != nil {
		utils.JsonResponse(w, types.Response{Message: "Could not fetch post"})
		return
	}
	row.Next()
	var id, title, summary, content, date, updated string
	row.Scan(&id, &title, &slug, &summary, &content, &date, &updated)
	utils.JsonResponse(w, types.Post{
		Id: id,
		Title: title,
		Slug: slug,
		Summary: summary,
		Content: content,
		Date: date,
		Updated: updated,
	})
}