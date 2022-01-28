package service

import (
	"errors"

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
	return nil, errors.New("Not implemented")
}
