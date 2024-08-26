package models

type ContentTextItem struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type ContentImageItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
