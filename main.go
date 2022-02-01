package main

import (
	"github.com/joho/godotenv"
	"github.comvictorfernandesraton/outbot/src/config"
	"github.comvictorfernandesraton/outbot/src/service"
)

func main() {
	godotenv.Load()

	geminConfig := &config.GemininConfig{}
	services := []service.CyptoServiceInterface{
		&service.GemininService{
			Config: geminConfig.New(),
		},
	}

	for _, v := range services {
		v.GetCurrencyInfo("BTC")
	}

}
