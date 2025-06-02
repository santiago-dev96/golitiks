package main

import "github.com/gocolly/colly"

func main() {
	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		// Print the response status code
		println("Response status code:", r.StatusCode)
	})
	c.Visit("https://www.elnacional.com/politica")
}
