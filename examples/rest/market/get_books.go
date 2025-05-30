package main

import (
	"log"

	"github.com/anhnguyenbk/go-okx/examples/rest"
	"github.com/anhnguyenbk/go-okx/rest/api/market"
)

func main() {
	param := &market.GetBooksParam{
		InstId: "ETH-USDT",
	}
	req, resp := market.NewGetBooks(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetBooksResponse))
}
