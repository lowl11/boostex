package excel_file

import (
	"github.com/lowl11/boostex"
	"github.com/lowl11/boostex/internal/celler"
	"github.com/lowl11/boostex/internal/sheet"
)

func (service *Service) Close() error {
	return service.file.Close()
}

func (service *Service) NewTab(tabName string) (boostex.Sheet, error) {
	sheetIndex, err := service.file.NewSheet(tabName)
	if err != nil {
		return nil, err
	}

	service.file.SetActiveSheet(sheetIndex)

	newSheet := sheet.New(tabName, service.file)
	service.tabs[tabName] = newSheet
	return newSheet, nil
}

func (service *Service) GetTab(tabName string) *sheet.Sheet {
	if tabName == service.firstTabName {
		return service.firstTab
	}

	return service.tabs[tabName]
}

func (service *Service) Set(tabName string, i, j int, value any) error {
	return service.file.SetCellValue(tabName, celler.Convert(i, j), value)
}

func (service *Service) SetNatural(tabName, index string, value any) error {
	return service.file.SetCellValue(tabName, index, value)
}

func (service *Service) Create() error {
	return service.file.SaveAs(service.fileName)
}
