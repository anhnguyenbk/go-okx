package main

import (
	"log"

	"github.com/anhnguyenbk/go-okx/ws/public"
)

func main() {
	handler := func(c public.EventIndexTickers) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeIndexTickers("BTC-USDT", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
