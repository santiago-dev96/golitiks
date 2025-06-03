package main

type NewsData struct {
	Title  string     `json:"title"`
	Link   string     `json:"link"`
	Source NewsSource `json:"source"`
}
