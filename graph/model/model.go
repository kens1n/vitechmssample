package model

type Hdata struct {
	Guid string `json:"guid" pg:",pk"`
	Hash string `json:"hash" pg:",pk"`
	Data string `json:"data"`
}
