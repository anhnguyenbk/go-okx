package main

import (
	"log"

	"github.com/anhnguyenbk/go-okx/examples/rest"
	"github.com/anhnguyenbk/go-okx/rest/api"
	"github.com/anhnguyenbk/go-okx/rest/api/account"
)

func main() {
	param := &account.PostSetPositionModeParam{
		PosMode: api.PosModeLongShort,
	}
	req, resp := account.NewPostSetPositionMode(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.PostSetPositionModeResponse))
}
