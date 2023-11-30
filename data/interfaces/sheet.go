package interfaces

type Sheet interface {
	Set(i, j int, value any) error
	SetNatural(index string, value any) error
}
