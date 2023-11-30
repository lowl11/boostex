package sheet

import "github.com/xuri/excelize/v2"

type Sheet struct {
	name string
	file *excelize.File
}

func New(name string, file *excelize.File) *Sheet {
	return &Sheet{
		name: name,
		file: file,
	}
}
