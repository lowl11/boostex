package excel

func (excel Excel) SetHeaders(tabName string, headers ...string) error {
	tab := excel.service.GetTab(tabName)
	if tab == nil {
		return nil
	}

	return tab.SetHeaders(headers...)
}

func (excel Excel) Set(tab string, i, j int, value any) error {
	return excel.service.GetTab(tab).Set(i, j, value)
}

func (excel Excel) SetNatural(tab, index string, value any) error {
	return excel.service.GetTab(tab).SetNatural(index, value)
}

func (excel Excel) Create() error {
	return excel.service.Create()
}

func (excel Excel) Close() error {
	return excel.service.Close()
}
