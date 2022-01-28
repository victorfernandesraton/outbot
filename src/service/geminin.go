package service

import (
	"errors"

	"github.com/jsgoyette/gemini"
	"github.comvictorfernandesraton/outbot/src/config"
	"github.comvictorfernandesraton/outbot/src/entity"
)

type GemininService struct {
	Config *config.GemininConfig
}

func (service *GemininService) GetCryptoInfo(symbol string) (*entity.Crypto, error) {
	return nil, errors.New("Not implemented")
}
func (service *GemininService) GetCurrencyInfo(symbol string) (*entity.Price, error) {
	gemini.New(false, service.Config.Key, service.Config.Secret)
	return nil, errors.New("Not implemented")
}
