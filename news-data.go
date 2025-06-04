package main

import "time"

// NewsData stores information
// about a news article.
type NewsData struct {
	Title  string
	Link   string
	Source NewsSource
	Date   time.Time
}
