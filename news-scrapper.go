package main

import "github.com/gocolly/colly"

// Entity that scrapes the web to
// find news about politics.
type NewsScrapper struct {
	collector *colly.Collector
	url       string
	Data      []NewsData
}

// NewNewsScrapper creates am already initialized NewsScrapper.
func NewNewsScrapper(url string) *NewsScrapper {
	return &NewsScrapper{
		collector: colly.NewCollector(),
		url:       url,
		Data:      make([]NewsData, 0),
	}
}

// Scrape leaves the scrapping exact process to the client,
// while providing a channel that the user must make use
// of to know if the scrapping is done.
func (s *NewsScrapper) Scrape(handler func(c *colly.Collector, errch chan<- error)) chan<- error {
	errch := make(chan error)
	handler(s.collector, errch)
	return errch
}
