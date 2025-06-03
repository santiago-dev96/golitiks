package main

import "github.com/gocolly/colly"

// Entity that scrapes the web to
// find news about politics.
type NewsScrapper struct {
	collector *colly.Collector
	url       string
}

// NewNewsScrapper creates am already initialized NewsScrapper.
func NewNewsScrapper(url string) *NewsScrapper {
	return &NewsScrapper{
		collector: colly.NewCollector(),
		url:       url,
	}
}

// Scrape leaves the scraping exact process to the client criteria,
// and does the scraping in a separate goroutine while providing a
// channel that the user must make use of to know if the scraping is done.
func (s *NewsScrapper) Scrape(handler func(c *colly.Collector, datach chan<- NewsData, errch chan<- error)) (<-chan NewsData, <-chan error) {
	errch := make(chan error)
	datach := make(chan NewsData)
	go func() {
		handler(s.collector, datach, errch)
		s.collector.Visit(s.url)
	}()
	return datach, errch
}
