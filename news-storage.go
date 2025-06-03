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
func NewNewsStorage(filename, titleText, linkText, sourceText string) (*NewsStorage, error) {
	file := excelize.NewFile()
	index, err := file.NewSheet(sheetName)
	if err != nil {
		return nil, err
	}
	err = file.SetCellValue(sheetName, "A1", titleText)
	if err != nil {
		return nil, err
	}
	err = file.SetCellValue(sheetName, "B1", linkText)
	if err != nil {
		return nil, err
	}
	err = file.SetCellValue(sheetName, "C1", sourceText)
	if err != nil {
		return nil, err
	}
	file.SetActiveSheet(index)
	return &NewsStorage{excelizeHandle: file, filename: filename + ".xlsx"}, nil
}

// Store saves NewsData to the file
// the receiver has internally.
func (ns *NewsStorage) Store(data NewsData) error {
	row := ns.recordsWritten + 2 // +2 to account for header row
	if err := ns.excelizeHandle.SetCellValue(sheetName, "A"+strconv.Itoa(row), data.Title); err != nil {
		return err
	}
	if err := ns.excelizeHandle.SetCellValue(sheetName, "B"+strconv.Itoa(row), data.Link); err != nil {
		return err
	}
	if err := ns.excelizeHandle.SetCellValue(sheetName, "C"+strconv.Itoa(row), string(data.Source)); err != nil {
		return err
	}
	ns.recordsWritten++
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
