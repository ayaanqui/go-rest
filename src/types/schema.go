package types

type Post struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Slug string `json:"slug"`
	Content string `json:"content"`
	Summary string `json:"summary"`
	Date string `json:"date"`
	Updated string `json:"updated"`
}

type Home struct {
	Id string `json:"id"`
	Message string `json:"message"`
	Date string `json:"date"`
}