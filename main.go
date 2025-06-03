package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/xuri/excelize/v2"
)

var data = make([]NewsData, 0)
var outputFilename = flag.String("output", "", "XLSX file name to save news data")

func main() {
	flag.Parse()
	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}()
	index, err := file.NewSheet("News Data")
	if err != nil {
		log.Fatalf("Error creating new sheet: %v", err)
	}
	err = file.SetCellValue("News Data", "A1", "Título")
	if err != nil {
		log.Fatalf("Error setting cell value: %v", err)
	}
	err = file.SetCellValue("News Data", "B1", "Enlace")
	if err != nil {
		log.Fatalf("Error setting cell value: %v", err)
	}
	err = file.SetCellValue("News Data", "C1", "Fuente")
	if err != nil {
		log.Fatalf("Error setting cell value: %v", err)
	}
	file.SetActiveSheet(index)
	c := colly.NewCollector()
	c.OnHTML(".module-category", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		source := getNewsSource(e.Request.URL.Host)
		title := e.ChildAttr("a", "title")
		newsData := NewsData{
			Title:  title,
			Link:   link,
			Source: source,
		}
		data = append(data, newsData)
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
		data = append(data, newsData)
	})
	c.OnScraped(func(r *colly.Response) {
		for i, newsData := range data {
			err := file.SetCellValue("News Data", "A"+strconv.Itoa(i+2), newsData.Title)
			if err != nil {
				log.Fatalf("Error setting cell value for title: %v", err)
			}
			err = file.SetCellValue("News Data", "B"+strconv.Itoa(i+2), newsData.Link)
			if err != nil {
				log.Fatalf("Error setting cell value for link: %v", err)
			}
			err = file.SetCellValue("News Data", "C"+strconv.Itoa(i+2), newsData.Source)
			if err != nil {
				log.Fatalf("Error setting cell value for source: %v", err)
			}
		}
		err := file.SaveAs(*outputFilename + ".xlsx")
		if err != nil {
			log.Fatalf("Error saving file: %v", err)
		}
	})
	c.Visit("https://www.elnacional.com/politica")
}
