package main

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/gocolly/colly"
)

var titleText = flag.String("title", "Title", "Text of the title header cell")
var linkText = flag.String("link", "Link", "Text of the link header cell")
var sourceText = flag.String("source", "Source", "Text of the source header cell")
var help = flag.Bool("help", false, "Show help")
var helpShort = flag.Bool("h", false, "")

var mu sync.Mutex

func main() {
	flag.Usage = printHelp
	flag.Parse()
	log.SetFlags(0)
	if *help || *helpShort {
		flag.Usage()
		os.Exit(0)
	}
	outputFilename := flag.Arg(0)
	if outputFilename == "" {
		log.Println("Output filename must be provided")
		flag.Usage()
		os.Exit(1)
	}
	newsStorage, err := NewNewsStorage(outputFilename, *titleText, *linkText, *sourceText)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := newsStorage.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	elNacionalScrapper := NewNewsScrapper(elNacionalURL)
	elNacionalDatach, elNacionalErrch := elNacionalScrapper.Scrape(elNacionalScrapeFn)
	totalScrapeOps := 1
	scrapeOpsFinished := 0
outer:
	for {
		select {
		case newsData := <-elNacionalDatach:
			err := newsStorage.Store(newsData)
			if err != nil {
				panic(err)
			}
		case err := <-elNacionalErrch:
			if err != nil {
				log.Println(err)
			}
			mu.Lock()
			scrapeOpsFinished++
			mu.Unlock()
			if scrapeOpsFinished == totalScrapeOps {
				break outer
			}
		}
	}
	err = newsStorage.Save()
	if err != nil {
		log.Fatalln(err)
	}
}

func printHelp() {
	log.Println("Usage: golitiks <filename> [-title <title>] [-link <link>] [-source <source>]")
	log.Println()
	log.Println("Flags:")
	log.Println()
	flag.PrintDefaults()
}

func elNacionalScrapeFn(c *colly.Collector, datach chan<- NewsData, errch chan<- error) {
	c.OnHTML(".module-category", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		source := getNewsSource(e.Request.URL.Host)
		title := e.ChildAttr("a", "title")
		newsData := NewsData{
			Title:  title,
			Link:   link,
			Source: source,
		}
		datach <- newsData
	})
	c.OnHTML(".list-articles .article", func(e *colly.HTMLElement) {
		title := e.ChildText(".content .title")
		source := getNewsSource(e.Request.URL.Host)
		link := e.ChildAttr(".image a", "href")
		newsData := NewsData{
			Title:  title,
			Link:   link,
			Source: source,
		}
		datach <- newsData
	})
	c.OnError(func(r *colly.Response, err error) {
		errch <- err
	})
	c.OnScraped(func(r *colly.Response) {
		errch <- nil
	})
}
