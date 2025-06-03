package main

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

const sheetName = "News Data"

// NewsStorage represents a storage
// where to put the news data
// collected by the NewsScrappers.
type NewsStorage struct {
	excelizeHandle *excelize.File
	recordsWritten int
	filename       string
}

// NewNewsStorage initializes a NewsStorage
// entity.
func NewNewsStorage(filename string) *NewsStorage {
	file := excelize.NewFile()
	return &NewsStorage{excelizeHandle: file, filename: filename}
}

// Store saves NewsData to the file
// the receiver has internally.
func (ns *NewsStorage) Store(data []NewsData) error {
	for _, newsData := range data {
		row := ns.recordsWritten + 2 // +2 to account for header row
		if err := ns.excelizeHandle.SetCellValue(sheetName, "A"+strconv.Itoa(row), newsData.Title); err != nil {
			return err
		}
		if err := ns.excelizeHandle.SetCellValue(sheetName, "B"+strconv.Itoa(row), newsData.Link); err != nil {
			return err
		}
		if err := ns.excelizeHandle.SetCellValue(sheetName, "C"+strconv.Itoa(row), string(newsData.Source)); err != nil {
			return err
		}
		ns.recordsWritten++
	}
	return nil
}

// Save saves the current data written to the
// storage to the filename associated with the NewsStorage
// with the .xlsx extension.
func (ns *NewsStorage) Save() error {
	return ns.excelizeHandle.SaveAs(ns.filename)
}

// Close closes the excelize handle.
func (ns *NewsStorage) Close() error {
	return ns.excelizeHandle.Close()
}
