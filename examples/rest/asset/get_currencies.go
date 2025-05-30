package main

import (
	"log"

	"github.com/anhnguyenbk/go-okx/examples/rest"
	"github.com/anhnguyenbk/go-okx/rest/api/asset"
)

func main() {
	param := &asset.GetCurrenciesParam{}
	req, resp := asset.NewGetCurrencies(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*asset.GetCurrenciesResponse))
}
