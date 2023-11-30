package sheet

import (
	"github.com/lowl11/boostex/internal/celler"
)

func (sheet Sheet) Set(i, j int, value any) error {
	return sheet.file.SetCellValue(sheet.name, celler.Convert(i, j), value)
}

func (sheet Sheet) SetNatural(index string, value any) error {
	return sheet.file.SetCellValue(sheet.name, index, value)
}

func (sheet Sheet) SetHeaders(headers ...string) error {
	i := 1

	for _, header := range headers {
		if err := sheet.Set(i, 1, header); err != nil {
			return err
		}

		i++
	}

	return nil
}
