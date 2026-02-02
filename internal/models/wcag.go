package models

type WCAGItem struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Keywords string `json:"keywords"`
	Level    string `json:"level"`
}
