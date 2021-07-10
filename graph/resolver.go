package graph

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/kens1n/vitechmssample/external_services"
	"github.com/kens1n/vitechmssample/postgres"
)

type Resolver struct {
	HdataRepo   postgres.HdataRepo
	GuidService external_services.GuidServiceInterface
	HashService external_services.HashServiceInterface
}
