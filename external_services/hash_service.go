package external_services

import (
	"github.com/ybbus/jsonrpc/v2"
	"os"
)

type HashServiceInterface interface {
	HashCalc(code string) (string, error)
}

type HashServiceLocal struct{}

func (c HashServiceLocal) HashCalc(code string) (string, error) {
	return "9d5d5c7afe1a9efbb5fcf5f903e15808", nil
}

type HashServiceExternal struct{}

type HashServiceExternalResponse struct {
	Hash string `json:"hash"`
}

func (c HashServiceExternal) HashCalc(code string) (string, error) {
	rpcClient := jsonrpc.NewClient(os.Getenv("S2URL"))
	response, err := rpcClient.Call("hash.calc", code)
	if err != nil {
		return "", err
	}

	if response.Error != nil {
		return "", response.Error
	}

	var responseData *HashServiceExternalResponse
	err = response.GetObject(&responseData)
	if err != nil || responseData == nil {
		return "", err
	}

	return responseData.Hash, nil
}
