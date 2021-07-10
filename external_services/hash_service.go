package external_services

type HashServiceInterface interface {
	HashCalc(code string) string
}

type HashServiceLocal struct{}

func (c HashServiceLocal) HashCalc(code string) string {
	return "9d5d5c7afe1a9efbb5fcf5f903e15808"
}
