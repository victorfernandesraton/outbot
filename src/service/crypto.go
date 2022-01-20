package service

import "github.comvictorfernandesraton/outbot/src/entity"

type CyptoServiceInterface interface {
	GetCryptoInfo(symbol string) (*entity.Crypto, error)
	GetCurrencyInfo(symbol string) (*entity.Crypto, error)
}
