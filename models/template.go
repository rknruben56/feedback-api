package models

type Template struct {
	ID          int    `json:"id"`
	Class       string `json:"class"`
	Content     string `json:"content"`
	DateCreated string `json:"dateCreated"`
}
