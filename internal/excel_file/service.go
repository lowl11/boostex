package excel_file

import (
	"github.com/lowl11/boostex/internal/sheet"
	"github.com/xuri/excelize/v2"
)

type Service struct {
	fileName     string
	file         *excelize.File
	tabs         map[string]*sheet.Sheet
	firstTabName string
	firstTab     *sheet.Sheet
}

func New(fileName, tabName string) *Service {
	file := excelize.NewFile()
	_ = file.SetSheetName("Sheet1", tabName)
	return &Service{
		fileName:     fileName,
		file:         file,
		tabs:         make(map[string]*sheet.Sheet),
		firstTabName: tabName,
		firstTab:     sheet.New(tabName, file),
	}
}
