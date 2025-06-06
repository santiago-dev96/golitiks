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
func NewNewsStorage(filename, titleText, dateText, sourceText, linkText string) (*NewsStorage, error) {
	// Setup the excelize file and sheet.
	excelizeHandle := excelize.NewFile()
	index, err := excelizeHandle.NewSheet(sheetName)
	if err != nil {
		return nil, err
	}

	// Set the headers for the columns.
	err = excelizeHandle.SetCellValue(sheetName, "A1", titleText)
	if err != nil {
		return nil, err
	}
	err = excelizeHandle.SetCellValue(sheetName, "B1", dateText)
	if err != nil {
		return nil, err
	}
	err = excelizeHandle.SetCellValue(sheetName, "C1", sourceText)
	if err != nil {
		return nil, err
	}
	err = excelizeHandle.SetCellValue(sheetName, "D1", linkText)
	if err != nil {
		return nil, err
	}
	// Set the column widths for better readability.
	err = excelizeHandle.SetColWidth(sheetName, "A", "A", 60)
	if err != nil {
		return nil, err
	}
	err = excelizeHandle.SetColWidth(sheetName, "B", "D", 20)
	if err != nil {
		return nil, err
	}
	// Set the style for the header row.
	style, err := excelizeHandle.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#D9EAD3"},
			Pattern: 1,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if err != nil {
		return nil, err
	}
	err = excelizeHandle.SetCellStyle(sheetName, "A1", "D1", style)
	if err != nil {
		return nil, err
	}

	// Set the active sheet to the newly created one.
	excelizeHandle.SetActiveSheet(index)

	return &NewsStorage{excelizeHandle: excelizeHandle, filename: filename + ".xlsx"}, nil
}

// Store saves NewsData to the file
// the receiver has internally.
func (ns *NewsStorage) Store(data NewsData) error {
	row := ns.recordsWritten + 2 // +2 to account for header row
	if err := ns.excelizeHandle.SetCellValue(sheetName, "A"+strconv.Itoa(row), data.Title); err != nil {
		return err
	}
	if err := ns.excelizeHandle.SetCellValue(sheetName, "B"+strconv.Itoa(row), data.Date); err != nil {
		return err
	}
	if err := ns.excelizeHandle.SetCellValue(sheetName, "C"+strconv.Itoa(row), string(data.Source)); err != nil {
		return err
	}
	if err := ns.excelizeHandle.SetCellValue(sheetName, "D"+strconv.Itoa(row), "url"); err != nil {
		return err
	}
	display, tooltip := "url", data.Title
	if err := ns.excelizeHandle.SetCellHyperLink(sheetName, "D"+strconv.Itoa(row), data.Link, "External", excelize.HyperlinkOpts{
		Display: &display,
		Tooltip: &tooltip,
	}); err != nil {
		return err
	}
	// Set the font and underline style for the link cell.
	style, err := ns.excelizeHandle.NewStyle(&excelize.Style{
		Font: &excelize.Font{Color: "1265BE", Underline: "single"},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if err != nil {
		return err
	}
	err = ns.excelizeHandle.SetCellStyle(sheetName, "D"+strconv.Itoa(row), "D"+strconv.Itoa(row), style)
	if err != nil {
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
