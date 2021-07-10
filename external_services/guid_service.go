package external_services

type GuidServiceInterface interface {
	GuidGenerate(code string) string
}

type GuidServiceLocal struct{}

func (c GuidServiceLocal) GuidGenerate(code string) string {
	return "f6014c6c-ddc5-11eb-ba80-0242ac130004"
}
