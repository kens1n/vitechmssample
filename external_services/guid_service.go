package external_services

import (
	"github.com/ybbus/jsonrpc/v2"
	"os"
)

type GuidServiceInterface interface {
	GuidGenerate(code string) (string, error)
}

type GuidServiceLocal struct{}

func (c GuidServiceLocal) GuidGenerate(code string) (string, error) {
	return "f6014c6c-ddc5-11eb-ba80-0242ac130004", nil
}

type GuidServiceExternal struct{}

type GuidServiceExternalResponse struct {
	Token string `json:"token"`
}

func (c GuidServiceExternal) GuidGenerate(code string) (string, error) {
	rpcClient := jsonrpc.NewClient(os.Getenv("S1URL"))
	response, err := rpcClient.Call("guid.generate", code)
	if err != nil {
		return "", err
	}

	if response.Error != nil {
		return "", response.Error
	}

	var responseData *GuidServiceExternalResponse
	err = response.GetObject(&responseData)
	if err != nil || responseData == nil {
		return "", nil
	}

	return responseData.Token, nil
}
