package excel

import "github.com/lowl11/boostex/internal/excel_file"

type Excel struct {
	service *excel_file.Service
}

func New(fileName, tabName string) *Excel {
	return &Excel{
		service: excel_file.New(fileName, tabName),
	}
}
